package utils

import (
	"os"
	"fmt"
	"reflect"
	"io/ioutil"
	"encoding/json"
)

//Check again
const confPath string = "./github.com/shubhamp-zeotap/Upflow/app/plugins.json"


func ReflectValueToJsonArray(input reflect.Value) (output []interface{}) {
	typeOfT := input.Type()
	for i:=0; i < input.NumField(); i++ {
		f := typeOfT.Field(i)
		var metadata map[string]interface{}
		if err := json.Unmarshal([]byte(string(f.Tag)), &metadata); err != nil {
			metadata = make(map[string]interface{})
		}
		metadata["Name"] = string(f.Name)		
		metadata["DataType"] = fmt.Sprintf("%s", f.Type)	//Check this later
		output = append(output, metadata)
	}	
	return
}


func GeneratePlugin(pluginName string, pluginConf map[string]interface{}) {
	file, err := os.OpenFile(confPath, os.O_RDONLY | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	configTextData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	file.Close()

	var configData map[string]interface{}
	if err := json.Unmarshal(configTextData, &configData); err != nil {
		configData = make(map[string]interface{})
	}		

	configData[pluginName] = pluginConf
	configTextData, _ = json.Marshal(configData)
	err = ioutil.WriteFile(confPath, []byte(string(configTextData)), 0666)
	if err != nil {
		panic("Plugin generation broke")
	}
}