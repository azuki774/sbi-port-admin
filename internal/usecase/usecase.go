package usecase

import (
	"azuki774/sbiport-server/internal/model"
	"context"

	"go.uber.org/zap"
)

type DBRepository interface {
	SaveRecords(ctx context.Context, records []model.DailyRecord, update bool) (result model.CreateRecordResult, err error)
}

type Usecase struct {
	Logger *zap.Logger
	DBRepo DBRepository
}

func (u *Usecase) RegistDailyRecords(ctx context.Context, rawStr string) (result model.CreateRecordResult, err error) {
	csvData, err := model.NewCSVRecord(rawStr)
	if err != nil {
		u.Logger.Error("failed to parse CSV file", zap.Error(err))
		return model.CreateRecordResult{}, err
	}

	fundInfos, err := csvData.FundsLoad()
	if err != nil {
		u.Logger.Error("failed to load fund Information", zap.Error(err))
		return model.CreateRecordResult{}, err
	}

	result, err = u.DBRepo.SaveRecords(ctx, fundInfos, false)
	if err != nil {
		u.Logger.Error("failed to save records", zap.Error(err))
	}
	u.Logger.Info("register daily record")
	return result, nil
}
