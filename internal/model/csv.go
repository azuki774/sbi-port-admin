package model

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	csvElementSize = 12 // CSVファイルの要素数
)

func NewCSVRecord(rawStr string) (csvdata CSVData, err error) {
	var records [][]string
	rowRecord := strings.Split(rawStr, "\n")
	for _, v := range rowRecord {
		comRec := strings.Split(v, ",")
		if len(comRec) != csvElementSize {
			// 空行をスキップする
			continue
		}
		records = append(records, comRec)
	}
	return CSVData(records), nil
}

func (c CSVData) FundsLoad() (fundsInfo []DailyRecord, err error) {
	index := 0
	for _, v := range c {
		if index != 0 {
			// ラベル部分は取り込まない
			var nowfundInfo DailyRecord
			err := fundLoad(&nowfundInfo, v)
			if err != nil {
				return nil, fmt.Errorf("parse error: %w", err)
			}
			fundsInfo = append(fundsInfo, nowfundInfo)
		}
		index++
	}
	return fundsInfo, nil
}

func fundLoad(fundInfo *DailyRecord, rowData []string) (err error) {
	if len(rowData) != csvElementSize {
		return fmt.Errorf("fundLoad parse error")
	}

	fundInfo.FundName = rowData[1] // ファンド名

	fundInfo.Amount, err = strconv.Atoi(rowData[3]) // 数量
	if err != nil {
		return fmt.Errorf("fundInfo.Count Atoi error: %w", err)
	}

	fundInfo.AcquisitionPrice, err = strconv.Atoi(rowData[4]) // 取得単価
	if err != nil {
		return fmt.Errorf("fundInfo.PurchaseUnitPrice Atoi error: %w", err)
	}

	fundInfo.NowPrice, err = strconv.Atoi(rowData[5]) // 現在値
	if err != nil {
		return fmt.Errorf("fundInfo.NowPrice Atoi error: %w", err)
	}

	fundInfo.ThedayBefore, err = strconv.Atoi(rowData[6]) // 前日比
	if err != nil {
		return fmt.Errorf("fundInfo.TheDayBefore Atoi error: %w", err)
	}

	fundInfo.ThedayBeforeRatio, err = strconv.ParseFloat(rowData[7], 64) // 前日比（％）
	if err != nil {
		return fmt.Errorf("fundInfo.TheDayBeforeRatio ParseFloat error: %w", err)
	}

	fundInfo.Profit, err = strconv.ParseFloat(rowData[8], 64) // 損益
	if err != nil {
		return fmt.Errorf("fundInfo.Profit ParseFloat error: %w", err)
	}

	fundInfo.ProfitRatio, err = strconv.ParseFloat(rowData[9], 64) // 損益（％）
	if err != nil {
		return fmt.Errorf("fundInfo.ProfitRatio ParseFloat error: %w", err)
	}

	fundInfo.Valuation, err = strconv.ParseFloat(rowData[10], 64) // 評価額
	if err != nil {
		return fmt.Errorf("fundInfo.Valuation ParseFloat error: %w", err)
	}

	return nil
}
