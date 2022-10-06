package main

import (
	"encoding/csv"
	"io"
	"log"
)


func ingest(file io.Reader) (header []string, data [][]string, indeices []map[string][]int, err error) {
	reader := csv.NewReader(file);

}

func hanleError(err error) {
	if err != nil {
		log.Fatal(err.Error);
	}
}

func main() {

}