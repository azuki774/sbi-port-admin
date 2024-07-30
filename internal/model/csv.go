package model

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	csvElementSize = 12 // CSVファイルの要素数
)

type CSVData struct {
	Fields [][]string
	Date   time.Time
}

func NewCSVRecord(rawStr string, t time.Time) (csvdata CSVData, err error) {
	reader := csv.NewReader(strings.NewReader(rawStr))
	records, err := reader.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		return CSVData{Fields: nil, Date: t}, err
	}
	for _, row := range records {
		// validation
		if len(row) != csvElementSize {
			return CSVData{Fields: nil, Date: t}, fmt.Errorf("invalid format csv")
		}

		// erase ','
		for i, e := range row {
			row[i] = strings.ReplaceAll(e, ",", "")
		}

	}
	return CSVData{Fields: records, Date: t}, nil
}

func (c CSVData) FundsLoad() (fundsInfo []DailyRecord, err error) {
	index := 0
	for _, v := range c.Fields {
		if index != 0 {
			// ラベル部分は取り込まない
			nowfundInfo := DailyRecord{RecordDate: c.Date}
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
