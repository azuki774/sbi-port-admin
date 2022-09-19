package repository

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type DBRepository struct {
	Conn *gorm.DB
}

func (dbR *DBRepository) CloseDB() (err error) {
	sqlDB, err := dbR.Conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (dbR *DBRepository) SaveRecords(ctx context.Context, records []model.DailyRecord, update bool) (result model.CreateRecordResult, err error) {
	for _, record := range records {
		exists := false
		err = dbR.Conn.WithContext(ctx).First(&record).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Not Found
		} else if err == nil {
			// Already registed
			exists = true
		} else {
			// SELECT error
			result.FailedNumber = result.FailedNumber + 1
			return result, fmt.Errorf("failed to SELECT error: %w", err)
		}

		if exists {
			if update {
				// update
				err = dbR.Conn.WithContext(ctx).Updates(&record).Error
				if err != nil {
					result.FailedNumber = result.FailedNumber + 1
					return result, err
				}
				result.UpdatedNumber = result.UpdatedNumber + 1
			} else {
				// skip
				result.SkippedNumber = result.SkippedNumber + 1
			}
		} else {
			// create new data
			err = dbR.Conn.WithContext(ctx).Create(&record).Error
			if err != nil {
				result.FailedNumber = result.FailedNumber + 1
				return result, err
			}
			result.CreatedNumber = result.CreatedNumber + 1
		}
	}

	return result, nil
}
