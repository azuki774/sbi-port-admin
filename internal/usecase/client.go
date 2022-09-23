package usecase

import (
	"fmt"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

type UsecaseClient struct {
	Logger *zap.Logger
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

// func (u *UsecaseClient) registJob() (err error) {

// }
