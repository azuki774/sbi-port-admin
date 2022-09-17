package repository

import (
	"azuki774/sbiport-server/internal/model"

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

func (dbR *DBRepository) SaveRecords(records []model.DailyRecord) (err error) {
	dbR.Conn.Create(&records)
	return nil
}
