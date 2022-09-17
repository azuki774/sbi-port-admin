package repository

import (
	"azuki774/sbi-port-admin/internal/model"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"go.uber.org/zap"
)

const (
	csvElementSize = 12 // CSVファイルの要素数
)

type CSVRecord struct {
	Logger   *zap.Logger
	FilePath string
}

func (c *CSVRecord) Load() (fundsInfo []model.DailyRecord, err error) {
	f, err := os.Open(c.FilePath)
	if err != nil {
		return []model.DailyRecord{}, err
	}
	defer f.Close()

	csvData, err := c.portCSVToString(f)
	if err != nil {
		return []model.DailyRecord{}, err
	}

	fundsInfo, err = c.fundsLoad(csvData)
	if err != nil {
		return []model.DailyRecord{}, err
	}

	return fundsInfo, nil
}

func (c *CSVRecord) portCSVToString(osf *os.File) (records [][]string, err error) {
	r := csv.NewReader(osf)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}
	return records, nil
}

func (c *CSVRecord) fundsLoad(csvData [][]string) (fundsInfo []model.DailyRecord, err error) {
	index := 0
	for _, v := range csvData {
		if index != 0 {
			// ラベル部分は取り込まない
			var nowfundInfo model.DailyRecord
			err := fundLoad(&nowfundInfo, v)
			if err != nil {
				c.Logger.Warn("parse error", zap.Error(err))
				return nil, nil
			}
			fundsInfo = append(fundsInfo, nowfundInfo)
		}
		index++
	}
	return fundsInfo, nil
}

func fundLoad(fundInfo *model.DailyRecord, rowData []string) (err error) {
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

	fundInfo.TheDayBefore, err = strconv.Atoi(rowData[6]) // 前日比
	if err != nil {
		return fmt.Errorf("fundInfo.TheDayBefore Atoi error: %w", err)
	}

	fundInfo.TheDayBeforeRatio, err = strconv.ParseFloat(rowData[7], 64) // 前日比（％）
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
