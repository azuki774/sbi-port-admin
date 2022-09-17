package repository

// type CSVRecord struct {
// 	Logger   *zap.Logger
// 	FilePath string
// }

// func (c *CSVRecord) Load() (fundsInfo []model.DailyRecord, err error) {
// 	f, err := os.Open(c.FilePath)
// 	if err != nil {
// 		return []model.DailyRecord{}, err
// 	}
// 	defer f.Close()

// 	csvData, err := c.portCSVToString(f)
// 	if err != nil {
// 		return []model.DailyRecord{}, err
// 	}

// 	fundsInfo, err = c.fundsLoad(csvData)
// 	if err != nil {
// 		return []model.DailyRecord{}, err
// 	}

// 	return fundsInfo, nil
// }

// func (c *CSVRecord) portCSVToString(osf *os.File) (records [][]string, err error) {
// 	r := csv.NewReader(osf)
// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}

// 		records = append(records, record)
// 	}
// 	return records, nil
// }
