package model

import (
	"reflect"
	"testing"
	"time"
)

func TestNewCSVRecord(t *testing.T) {
	type args struct {
		rawStr string
		date   time.Time
	}
	tests := []struct {
		name        string
		args        args
		wantRecords CSVData
		wantErr     bool
	}{
		{
			name: "ok",
			args: args{
				rawStr: `取引,ファンド名,買付日,数量,取得単価,現在値,前日比,前日比（％）,損益,損益（％）,評価額,編集
積立  売却,AAA,--/--/--,26231,13000,11403,-258,-2.21,-4189.09,-12.28,29911.2,詳細 

積立  売却,BBB,--/--/--,10946,31610,29726,+235,+0.80,-2062.22,-5.96,32538.07,詳細 
`,
			},
			wantRecords: CSVData{
				Fields: [][]string{
					{"取引", "ファンド名", "買付日", "数量", "取得単価", "現在値", "前日比", "前日比（％）", "損益", "損益（％）", "評価額", "編集"},
					{"積立  売却", "AAA", "--/--/--", "26231", "13000", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 "},
					{"積立  売却", "BBB", "--/--/--", "10946", "31610", "29726", "+235", "+0.80", "-2062.22", "-5.96", "32538.07", "詳細 "},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := NewCSVRecord(tt.args.rawStr, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCSVRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("NewCSVRecord() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func Test_fundLoad(t *testing.T) {
	type args struct {
		fundInfo *DailyRecord
		rowData  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *DailyRecord
	}{
		{
			name: "ok",
			args: args{
				fundInfo: &DailyRecord{},
				rowData:  []string{"積立  売却", "AAA", "--/--/--", "26231", "13000", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 "},
			},
			wantErr: false,
			want: &DailyRecord{
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
			err := fundLoad(tt.args.fundInfo, tt.args.rowData)
			if (err != nil) != tt.wantErr {
				t.Errorf("fundLoad() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.fundInfo, tt.want) {
				t.Errorf("fundLoad() = %v, want %v", tt.args.fundInfo, tt.want)
			}
		})
	}
}
