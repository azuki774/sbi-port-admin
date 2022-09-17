package usecase

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type DBRepository interface {
}

type Usecase struct {
	Logger *zap.Logger
	DBRepo DBRepository
}

func (u *Usecase) RegistDailyRecords(ctx context.Context, rawStr string) (err error) {
	csvData, err := model.NewCSVRecord(rawStr)
	if err != nil {
		return err
	}

	fundInfos, err := csvData.FundsLoad()
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", fundInfos)

	u.Logger.Info("register daily recorded")
	return nil
}
