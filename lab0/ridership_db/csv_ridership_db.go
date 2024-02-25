package ridershipDB

import (
	"encoding/csv"
	"fmt"
	"io"
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

func (c *CsvRidershipDB) GetRidership(lineId string) ([]int64, error) {
	rc := make([]int64, 9)

	// Skip the header row
	_, err := c.csvReader.Read()
	if err != nil {
		return nil, err
	}
	for {
		data, err := c.csvReader.Read()
		if err != nil {
			if err == io.EOF {
				return rc, nil
			}
			return nil, err
		}
		if len(data) != 5 {
			continue
		}
		if data[0] == lineId {
			one := data[len(data)-1]
			covOne, err := strconv.ParseInt(one, 10, 64)
			if err != nil {
				return nil, err
			}
			rc[c.idIdxMap[data[2]]] += covOne
		}
	}
	return rc, nil

}

func (c *CsvRidershipDB) Close() error {
	return c.csvFile.Close()
}
