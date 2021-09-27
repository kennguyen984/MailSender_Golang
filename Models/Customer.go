package Models

import (
	// "errors"
	"encoding/csv"
	"os"
)

type Customer struct {
	Title      string
	First_name string
	Last_name  string
	Email      string
}

func (c *Customer) Get(path string) ([]Customer, error) {
	var results []Customer

	csvFile, err := os.Open(path)
	if err != nil {
		return results, err
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return results, err
	}
	// skip first row
	for i := 1; i < len(csvLines); i++ {
		line := csvLines[i]
		row := Customer{
			Title:      line[0],
			First_name: line[1],
			Last_name:  line[2],
			Email:      line[3],
		}
		results = append(results, row)
	}
	return results, nil
}
