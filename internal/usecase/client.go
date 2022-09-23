package usecase

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

type HTTPClient interface {
	PostFile(ctx context.Context, endPoint string, filePath string) (resBody []byte, statusCode int, err error)
}

type UsecaseClient struct {
	Logger     *zap.Logger
	HTTPClient HTTPClient
}

type RegistFileInformation struct {
	FilePath string
	Date     string // YYYYMMDD
}

func (r *RegistFileInformation) FillDateByFileName() (err error) {
	if r.FilePath == "" {
		return fmt.Errorf("filename empty")
	}
	basefileName := filepath.Base(r.FilePath[:len(r.FilePath)-len(filepath.Ext(r.FilePath))])

	// try to YYYYMMDD parse
	date := func() (date string) {
		t, err := time.ParseInLocation("20060102", basefileName, time.Local)
		if err != nil {
			return ""
		}
		return t.Local().Format("20060102")
	}()
	if date != "" {
		r.Date = date
		return nil
	}

	// try to YYYY-MM-DD parse
	date = func() (date string) {
		t, err := time.ParseInLocation("2006-01-02", basefileName, time.Local)
		if err != nil {
			return ""
		}
		return t.Local().Format("20060102")
	}()
	if date != "" {
		r.Date = date
		return nil
	}

	// otherwise
	return fmt.Errorf("unrecognize file name")
}

func (u *UsecaseClient) RegistJob(ctx context.Context, filePath string) (err error) {
	var reg RegistFileInformation
	reg.FilePath = filePath

	err = reg.FillDateByFileName()
	if err != nil {
		u.Logger.Error("failed to get date", zap.Error(err))
		return err
	}

	endPoint := "/regist/" + reg.Date
	resBody, statusCode, err := u.HTTPClient.PostFile(ctx, endPoint, reg.FilePath)
	if err != nil {
		u.Logger.Error("failed to post CSV file", zap.Error(err))
		return err
	}
	if statusCode != 200 {
		u.Logger.Error("status code is not 200", zap.Int("status_code", statusCode))
		return errors.New("error response found")
	}

	var crr model.CreateRecordResult
	err = json.Unmarshal(resBody, &crr)
	if err != nil {
		u.Logger.Warn("failed to parse response", zap.Error(err))
		return err
	}

	u.Logger.Info("response from server", zap.String("body", string(resBody)))
	return nil
}
