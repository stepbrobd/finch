package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCSV(filename string) ([][]float32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var data [][]float32
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		var row []float32
		for _, valueStr := range record {
			value, err := strconv.ParseFloat(valueStr, 32)
			if err != nil {
				return nil, err
			}
			row = append(row, float32(value))
		}

		data = append(data, row)
	}

	return data, nil
}
