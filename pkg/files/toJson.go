package files

import (
	"encoding/json"
	"log"
	"os"
)

func ToJson(data interface{}) []byte{
	jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
	return jsonData
}

func WriteJsonFile(jsonData []byte,outputPath string){
	jsonFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}