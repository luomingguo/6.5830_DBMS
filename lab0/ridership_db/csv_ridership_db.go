package ridershipDB

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CsvRidershipDB struct {
	idIdxMap      map[string]int
	csvFile       *os.File
	csvReader     *csv.Reader
	num_intervals int
}

func (c *CsvRidershipDB) Open(filePath string) error {
	c.num_intervals = 9

	// Create a map that maps MBTA's time period ids to indexes in the slice
	c.idIdxMap = make(map[string]int)
	for i := 1; i <= c.num_intervals; i++ {
		timePeriodID := fmt.Sprintf("time_period_%02d", i)
		c.idIdxMap[timePeriodID] = i - 1
	}

	// create csv reader
	csvFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	c.csvFile = csvFile
	c.csvReader = csv.NewReader(c.csvFile)

	return nil
}

// TODO: some code goes here
// Implement the remaining RidershipDB methods
func (s *CsvRidershipDB) GetRidership(lineId string) ([]int64, error) {

	records, err := s.csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	ret := make([]int64, len(s.idIdxMap))
	var timeStamp string
	for _, record := range records {
		if record[0] == lineId {
			val, err := strconv.Atoi(record[4])
			if err != nil {
				return nil, err
			}
			timeStamp = record[2]
			ret[s.idIdxMap[timeStamp]] += int64(val)
		}
	}
	return ret, nil
}

func (s *CsvRidershipDB) Close() error {
	return s.csvFile.Close()
}
