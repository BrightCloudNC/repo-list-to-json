package main

import (
	"fmt"

	"github.com/BrightCloudNC/repo-list-to-json/pkg/dataX"
	"github.com/BrightCloudNC/repo-list-to-json/pkg/files"
)

func main(){
	rawDataPath := "./data/rawdata.csv"
	outputPath := "./output/dataX.json"

    // Read data in CSV file.
	data := files.ReadFile(rawDataPath)

	// Clean data.
	dataXList := dataX.DataXList(data)

	// Cleaned data to JSON format.
	jsonData := files.ToJson(dataXList)
	
	//Write JSON data to JSON file.
	files.WriteJsonFile(jsonData,outputPath)

	fmt.Println(string(jsonData))
}