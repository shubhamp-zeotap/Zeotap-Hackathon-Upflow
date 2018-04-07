package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFileToJson(filePath string) map[string]interface{} {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		panic(err)
	}

	return jsonData
}

func WriteJsonToFile(filePath string, jsonData map[string]interface{}) {
	data, _ := json.MarshalIndent(jsonData, "", "  ")
	err := ioutil.WriteFile(filePath, []byte(string(data)), 0666)
	if err != nil {
		panic(err)
	}
}
