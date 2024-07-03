package repository

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const categoryTagMasterName = "category_tag_master"

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

func (dbR *DBRepository) SaveRecords(ctx context.Context, records []model.DailyRecord, categoryTag string, update bool) (result model.CreateRecordResult, err error) {
	// categoryTag から 挿入すべきテーブル名を取得する
	var categoryMaster model.CategoryTagMaster
	err = dbR.Conn.Table(categoryTagMasterName).WithContext(ctx).Where("category_tag_name = ?", categoryTag).First(&categoryMaster).Error
	if err != nil {
		return model.CreateRecordResult{}, err
	}
	tableName := categoryMaster.TableName

	for _, record := range records {
		exists := false
		err = dbR.Conn.Table(tableName).WithContext(ctx).First(&record).Error
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
				err = dbR.Conn.Table(tableName).WithContext(ctx).Updates(&record).Error
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
			err = dbR.Conn.Table(tableName).WithContext(ctx).Create(&record).Error
			if err != nil {
				result.FailedNumber = result.FailedNumber + 1
				return result, err
			}
			result.CreatedNumber = result.CreatedNumber + 1
		}
	}

	return result, nil
}

func (dbR *DBRepository) GetDailyRecords(ctx context.Context, date string, categoryTag string) (recordsRepl []model.DailyRecordRepl, err error) {
	var records []model.DailyRecord
	err = dbR.Conn.WithContext(ctx).Where("record_date = ?", date).Find(&records).Error
	if err != nil {
		return []model.DailyRecordRepl{}, fmt.Errorf("failed to SELECT error: %w", err)
	}

	for _, v := range records {
		recordRepl := model.NewDailyRecordRepl(v)
		recordsRepl = append(recordsRepl, recordRepl)
	}

	return recordsRepl, nil
}
