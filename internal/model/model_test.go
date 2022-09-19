package model

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDailyRecordRepl(t *testing.T) {
	type args struct {
		d DailyRecord
	}
	tests := []struct {
		name string
		args args
		want DailyRecordRepl
	}{
		{
			name: "ok",
			args: args{d: DailyRecord{
				RecordDate:        time.Date(2006, 1, 23, 0, 0, 0, 0, time.Local),
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
			},
			want: DailyRecordRepl{
				RecordDate:        "20060123",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDailyRecordRepl(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDailyRecordRepl() = %v, want %v", got, tt.want)
			}
		})
	}
}
