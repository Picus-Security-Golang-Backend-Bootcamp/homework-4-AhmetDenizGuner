package csv_helper

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filename string, startIndex int) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result [][]string

	for _, line := range lines[startIndex:] {
		result = append(result, line)
	}

	return result, nil
}
