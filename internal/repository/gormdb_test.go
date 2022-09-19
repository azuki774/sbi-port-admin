package repository

import (
	"azuki774/sbiport-server/internal/model"
	"context"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func NewDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		return nil, mock, err
	}

	db, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn:                      mockDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	return db, mock, err
}

func TestDBRepository_SaveRecords(t *testing.T) {
	type fields struct {
		Conn *gorm.DB
	}
	type args struct {
		ctx     context.Context
		records []model.DailyRecord
		update  bool
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult model.CreateRecordResult
		wantErr    bool
		before     func(mock sqlmock.Sqlmock)
	}{
		{
			name: "new data",
			args: args{
				ctx: context.Background(),
				records: []model.DailyRecord{{
					RecordDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					FundName:   "AAA",
					Amount:     100,
					Valuation:  123.45,
				}},
				update: false,
			},
			wantResult: model.CreateRecordResult{
				CreatedNumber: 1,
			},
			wantErr: false,
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `daily_records` WHERE `daily_records`.`record_date` = ? AND `daily_records`.`fund_name` = ? ORDER BY `daily_records`.`record_date` LIMIT 1")).
					WithArgs(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), "AAA").WillReturnError(gorm.ErrRecordNotFound)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `daily_records` (`record_date`,`fund_name`,`amount`,`acquisition_price`,`now_price`,`theday_before`,`theday_before_ratio`,`profit`,`profit_ratio`,`valuation`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
					WillReturnResult(sqlmock.NewResult(100, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "skipped",
			args: args{
				ctx: context.Background(),
				records: []model.DailyRecord{{
					RecordDate: time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
					FundName:   "BBB",
					Amount:     100,
					Valuation:  123.45,
				}},
				update: false,
			},
			wantResult: model.CreateRecordResult{
				SkippedNumber: 1,
			},
			wantErr: false,
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `daily_records` WHERE `daily_records`.`record_date` = ? AND `daily_records`.`fund_name` = ? ORDER BY `daily_records`.`record_date` LIMIT 1")).
					WithArgs(time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local), "BBB").
					WillReturnRows(sqlmock.NewRows([]string{"record_date", "fund_name"}).AddRow(time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local), "BBB"))
			},
		},
		{
			name: "new data (overwrite mode)",
			args: args{
				ctx: context.Background(),
				records: []model.DailyRecord{{
					RecordDate: time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
					FundName:   "AAA",
					Amount:     100,
					Valuation:  123.45,
				}},
				update: true,
			},
			wantResult: model.CreateRecordResult{
				CreatedNumber: 1,
			},
			wantErr: false,
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `daily_records` WHERE `daily_records`.`record_date` = ? AND `daily_records`.`fund_name` = ? ORDER BY `daily_records`.`record_date` LIMIT 1")).
					WithArgs(time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local), "AAA").WillReturnError(gorm.ErrRecordNotFound)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `daily_records` (`record_date`,`fund_name`,`amount`,`acquisition_price`,`now_price`,`theday_before`,`theday_before_ratio`,`profit`,`profit_ratio`,`valuation`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
					WillReturnResult(sqlmock.NewResult(102, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "update (overwrite mode)",
			args: args{
				ctx: context.Background(),
				records: []model.DailyRecord{{
					RecordDate: time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local),
					FundName:   "BBB",
					Amount:     100,
					Valuation:  123.45,
				}},
				update: true,
			},
			wantResult: model.CreateRecordResult{
				UpdatedNumber: 1,
			},
			wantErr: false,
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `daily_records` WHERE `daily_records`.`record_date` = ? AND `daily_records`.`fund_name` = ? ORDER BY `daily_records`.`record_date` LIMIT 1")).
					WithArgs(time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local), "BBB").
					WillReturnRows(sqlmock.NewRows([]string{"record_date", "fund_name"}).AddRow(time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local), "BBB"))
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `daily_records` SET `amount`=?,`valuation`=? WHERE `record_date` = ? AND `fund_name` = ?")).WithArgs(100, 123.45, time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local), "BBB").
					WillReturnResult(sqlmock.NewResult(103, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "failed(INSERT)",
			args: args{
				ctx: context.Background(),
				records: []model.DailyRecord{{
					RecordDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
					FundName:   "AAA",
					Amount:     100,
					Valuation:  123.45,
				}},
				update: false,
			},
			wantResult: model.CreateRecordResult{
				FailedNumber: 1,
			},
			wantErr: true,
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `daily_records` WHERE `daily_records`.`record_date` = ? AND `daily_records`.`fund_name` = ? ORDER BY `daily_records`.`record_date` LIMIT 1")).
					WithArgs(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local), "AAA").WillReturnError(gorm.ErrRecordNotFound)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `daily_records` (`record_date`,`fund_name`,`amount`,`acquisition_price`,`now_price`,`theday_before`,`theday_before_ratio`,`profit`,`profit_ratio`,`valuation`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
					WillReturnError(gorm.ErrInvalidData)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := NewDBMock()
			tt.before(mock)
			dbR := &DBRepository{
				Conn: db,
			}
			gotResult, err := dbR.SaveRecords(tt.args.ctx, tt.args.records, tt.args.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.SaveRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("DBRepository.SaveRecords() = %v, want %v", gotResult, tt.wantResult)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDBRepository_GetDailyRecords(t *testing.T) {
	type fields struct {
		Conn *gorm.DB
	}
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantRecordsRepl []model.DailyRecordRepl
		wantErr         bool
		before          func(mock sqlmock.Sqlmock)
	}{
		{
			name: "2 records",
			args: args{
				ctx:  context.Background(),
				date: "20060102",
			},
			wantRecordsRepl: []model.DailyRecordRepl{
				{
					RecordDate:        "20060102",
					FundName:          "AAA",
					Amount:            26231,
					AcquisitionPrice:  13000,
					NowPrice:          11403,
					ThedayBefore:      -258,
					ThedayBeforeRatio: -2.21,
					Profit:            -4189.09,
					ProfitRatio:       -12.28,
					Valuation:         29911.2,
				},
				{
					RecordDate:        "20060102",
					FundName:          "BBB",
					Amount:            26231,
					AcquisitionPrice:  13000,
					NowPrice:          11403,
					ThedayBefore:      -258,
					ThedayBeforeRatio: -2.21,
					Profit:            -4189.09,
					ProfitRatio:       -12.28,
					Valuation:         29911.2,
				},
			},
			wantErr: false,
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `daily_records` WHERE record_date = ?")).
					WithArgs("20060102").
					WillReturnRows(sqlmock.NewRows([]string{"record_date", "fund_name", "amount", "acquisition_price", "now_price", "theday_before", "theday_before_ratio", "profit", "profit_ratio", "valuation"}).
						AddRow(time.Date(2006, 1, 2, 0, 0, 0, 0, time.Local), "AAA", 26231, 13000, 11403, -258, -2.21, -4189.09, -12.28, 29911.2).
						AddRow(time.Date(2006, 1, 2, 0, 0, 0, 0, time.Local), "BBB", 26231, 13000, 11403, -258, -2.21, -4189.09, -12.28, 29911.2))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := NewDBMock()
			tt.before(mock)
			dbR := &DBRepository{
				Conn: db,
			}
			gotRecordsRepl, err := dbR.GetDailyRecords(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.GetDailyRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecordsRepl, tt.wantRecordsRepl) {
				t.Errorf("DBRepository.GetDailyRecords() = %#+v, want %#+v", gotRecordsRepl, tt.wantRecordsRepl)
			}
		})
	}
}
