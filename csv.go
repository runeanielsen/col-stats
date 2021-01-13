package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	cr := csv.NewReader(r)
	column--

	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Cannot read data from file: %s", err)
	}

	var data []float64

	for i, row := range allData {
		if i == 0 {
			continue
		}

		if len(row) <= column {
			// File does not have that many columns
			return nil, fmt.Errorf("Invalid column #. File has only %d columns", len(row))
		}

		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("Data is not numeric: %s", err)
		}

		data = append(data, v)
	}

	return data, nil
}
