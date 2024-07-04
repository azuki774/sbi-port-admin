package usecase

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
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

	// get YYYYYMMDD from YYYYMMDD_<category_tag>
	basefileNameArr := strings.Split(basefileName, "_")
	basefileName = basefileNameArr[0]

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

func getCategoryTag(filePath string) (categoryTag string) {
	basefileName := filepath.Base(filePath[:len(filePath)-len(filepath.Ext(filePath))])
	arr := strings.Split(basefileName, "_")
	categoryTag = arr[len(arr)-1]
	return categoryTag
}

// YYYYMMDD_<category_tag>.csv をサーバに登録する
// 例外として、<category_tag> = 1 --> "nisa24", <category_tag> = 2 --> "nisa" に変換する。
func (u *UsecaseClient) RegistJob(ctx context.Context, filePath string) (err error) {
	var reg RegistFileInformation
	reg.FilePath = filePath
	u.Logger.Info("process file", zap.String("filename", filePath))

	err = reg.FillDateByFileName()
	if err != nil {
		u.Logger.Error("failed to get date", zap.Error(err))
		return err
	}

	// get category tag
	categoryTag := getCategoryTag(filePath)

	func() { // 一時例外処理
		if categoryTag == "1" {
			categoryTag = "nisa24"
		} else if categoryTag == "2" {
			categoryTag = "nisa"
		}
	}()

	endPoint := "/regist/" + categoryTag + "/" + reg.Date // /regist/<category_tag>/YYYYMMDD
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
	u.Logger.Info("regist CSV file", zap.String("filename", filePath))
	return nil
}
