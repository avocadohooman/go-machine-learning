package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func hanleError(err error) {
	if err != nil {
		log.Fatal(err.Error);
	}
}

func ingest(file io.Reader) (header []string, data [][]string, indices []map[string][]int, err error) {
	reader := csv.NewReader(file);

	// handle reader
	if header, err = reader.Read(); err != nil {
		return
	}

	indices = make([]map[string][]int, len(header))
	var rowCount , colCount int = 0, len(header)
	fmt.Println("rowCount", rowCount)
	fmt.Println("colCount", colCount)

	for rec, err := reader.Read(); err == nil; rec, err = reader.Read() {
		if len(rec) != colCount {
			return nil, nil, nil, errors.Errorf("Expected Columns: %d. Got %d columns in row %d", colCount, len(rec), rowCount)
		}
		data = append(data, rec)
		for j, val := range rec {
			if indices[j] == nil {
				indices[j] = make(map[string][]int)
			}
			indices[j][val] = append(indices[j][val], rowCount)
		}
		rowCount++
	}
	return
}
// cardinality counts the number of unique values in a column.
// This assumes that the index i of indices represents a column.
func cardinality(indices []map[string][]int) []int {
	returnValue := make([]int, len(indices))
	for i, m := range indices {
		returnValue[i] = len(m)
	}
	return returnValue
}

func main() {
	file, err := os.Open("./data/train.csv")
	hanleError(err)
	hdr, data, indices, err := ingest(file)
	hanleError(err)
	cardinality := cardinality(indices)
	fmt.Printf("Original Data: \nRows: %d, Cols: %d\n======\n", len(data), len(hdr))
	for i, h := range hdr {
		fmt.Printf("%v: %v\n", h, cardinality[i])
	}
	fmt.Println("")
}