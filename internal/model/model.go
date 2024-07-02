package model

import "time"

// DailyRecord is DB model
type DailyRecord struct {
	RecordDate        time.Time `gorm:"primaryKey"` // 取り込み日付
	FundName          string    `gorm:"primaryKey"` // ファンド名
	Amount            int       // 数量 47947
	AcquisitionPrice  int       // 取得単価 12796
	NowPrice          int       // 現在値 12864
	ThedayBefore      int       // 前日比 -284
	ThedayBeforeRatio float64   // 前日比（％） -2.16
	Profit            float64   // 損益 +326.03
	ProfitRatio       float64   // 損益（％） +0.53
	Valuation         float64   // 評価額 61679.02
}

type CategoryTagMaster struct {
	CategoryTagName string
	TableName       string
}

type CreateRecordResult struct {
	CreatedNumber int `json:"created_number,omitempty"`
	UpdatedNumber int `json:"updated_number,omitempty"`
	SkippedNumber int `json:"skipped_number,omitempty"`
	FailedNumber  int `json:"failed_number,omitempty"`
}

// DailyRecordRepl uses when json output
type DailyRecordRepl struct {
	RecordDate        string  // YYYYMMDD
	FundName          string  // ファンド名
	Amount            int     // 数量 47947
	AcquisitionPrice  int     // 取得単価 12796
	NowPrice          int     // 現在値 12864
	ThedayBefore      int     // 前日比 -284
	ThedayBeforeRatio float64 // 前日比（％） -2.16
	Profit            float64 // 損益 +326.03
	ProfitRatio       float64 // 損益（％） +0.53
	Valuation         float64 // 評価額 61679.02
}

func NewDailyRecordRepl(d DailyRecord) DailyRecordRepl {
	dRel := DailyRecordRepl{
		FundName:          d.FundName,
		Amount:            d.Amount,
		AcquisitionPrice:  d.AcquisitionPrice,
		NowPrice:          d.NowPrice,
		ThedayBefore:      d.ThedayBefore,
		ThedayBeforeRatio: d.ThedayBeforeRatio,
		Profit:            d.Profit,
		ProfitRatio:       d.ProfitRatio,
		Valuation:         d.Valuation,
	}
	dRel.RecordDate = d.RecordDate.Local().Format("20060102")
	return dRel
}
