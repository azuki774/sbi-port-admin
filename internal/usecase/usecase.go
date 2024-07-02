package usecase

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"time"

	"go.uber.org/zap"
)

type DBRepository interface {
	SaveRecords(ctx context.Context, records []model.DailyRecord, categoryTag string, update bool) (result model.CreateRecordResult, err error)
	GetDailyRecords(ctx context.Context, date string, categoryTag string) (recordsRepl []model.DailyRecordRepl, err error)
}

type Usecase struct {
	Logger *zap.Logger
	DBRepo DBRepository
}

func (u *Usecase) RegistDailyRecords(ctx context.Context, rawStr string, t time.Time, categoryTag string) (result model.CreateRecordResult, err error) {
	csvData, err := model.NewCSVRecord(rawStr, t)
	if err != nil {
		u.Logger.Error("failed to parse CSV file", zap.Error(err))
		return model.CreateRecordResult{}, err
	}

	fundInfos, err := csvData.FundsLoad()
	if err != nil {
		u.Logger.Error("failed to load fund Information", zap.Error(err))
		return model.CreateRecordResult{}, err
	}

	result, err = u.DBRepo.SaveRecords(ctx, fundInfos, categoryTag, false)
	if err != nil {
		u.Logger.Error("failed to save records", zap.Error(err))
	}
	u.Logger.Info("register daily record")
	return result, nil
}

func (u *Usecase) GetDailyRecords(ctx context.Context, date string, categoryTag string) ([]model.DailyRecordRepl, error) {
	// validation
	err := model.ValidateDate(date)
	if err != nil {
		return []model.DailyRecordRepl{}, ErrInvalidDate
	}

	recordsRepl, err := u.DBRepo.GetDailyRecords(ctx, date, categoryTag)
	if err != nil {
		u.Logger.Error("failed to get records", zap.Error(err))
	}

	return recordsRepl, nil
}
