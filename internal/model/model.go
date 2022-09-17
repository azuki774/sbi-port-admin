package model

import "time"

type CSVData [][]string

type DailyRecord struct {
	RecordDate        time.Time // 取り込み日付
	FundName          string    // ファンド名
	Amount            int       // 数量 47947
	AcquisitionPrice  int       // 取得単価 12796
	NowPrice          int       // 現在値 12864
	TheDayBefore      int       // 前日比 -284
	TheDayBeforeRatio float64   // 前日比（％） -2.16
	Profit            float64   // 損益 +326.03
	ProfitRatio       float64   // 損益（％） +0.53
	Valuation         float64   // 評価額 61679.02
}
