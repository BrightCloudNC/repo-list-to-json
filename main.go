package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BrightCloudNC/repo-list-to-json/csvtojson"
)

func main(){
	rawDataPath := "./data/rawdata.csv"
    f, err := os.Open(rawDataPath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

	dataXList := csvtojson.DataXList(data)
	jsonData, err := json.MarshalIndent(dataXList, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

	outputPath := "./output/dataX.json"
	jsonFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)

	fmt.Println(string(jsonData))
}