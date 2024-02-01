package files

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadFile(path string)[][]string {
	f, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

	return data
}