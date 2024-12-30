package utils

import (
	"encoding/csv"
	"os"
)

// Load solution in csv format
func LoadSolution(path string) [][]string {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if fileInfo.IsDir() {
		panic("Provided path is a directory")
	}
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}
