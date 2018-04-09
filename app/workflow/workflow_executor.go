package workflow

import (
	"fmt"
	"strings"
	"sort"
	"github.com/shubhamp-zeotap/Upflow/app/registry"
	"github.com/shubhamp-zeotap/Upflow/app/utils"
)

func ExecuteWorkflow(workflow map[string]interface{}, globalInput map[string]interface{}) {
	globalArgs := jsonToArgsList(globalInput)
	steps := workflow["StepsSequence"].([]interface{})
	outputCache := make([][]interface{}, 0, len(steps))

	fmt.Println("Executing workflow ...")
	for _, s := range steps {	
		execArgs := make([]interface{}, 0)
		step := s.(map[string]interface{})
		stepName := step["StepName"].(string)
		inputConfigs := step["InputConfig"].([]interface{})

		fmt.Println("Executing step : " + stepName)
		for inputIndex, i := range inputConfigs {
			fmt.Println("Input Config :", inputIndex)
			input := i.(map[string]interface{})
			givenValue := fmt.Sprintf("%s", input["Input"])
			inputRef := strings.Split(givenValue, ".")			
			var usedValue interface{}
			if len(inputRef) == 2  {
				i0, i1 := utils.StrToInt(inputRef[0]), utils.StrToInt(inputRef[1])
				if i0 == -1 {
					usedValue = globalArgs[i1]
				} else {
					usedValue = outputCache[i0][i1]
				}
			} else if givenValue != "" {
				usedValue = input["Input"]
			} else if input["CanUseDefault"].(bool) == true {
				usedValue = input["Default"]
			} else {
				panic("No suitable input found for this plugin !")
			}
			execArgs = append(execArgs, usedValue)
		}

		plugin := registry.GetPlugin(stepName, execArgs...)
		retValues, err := plugin.Execute()
		if err != nil {
			panic("Worflow execution broke at step :" + stepName)
		}
		outputCache = append(outputCache, retValues.Values)
	}
	fmt.Println("Workflow execution complete ...")	
}


func jsonToArgsList(jsonData map[string]interface{}) (argsList []interface{}) {
	var keys []string
	for k := range jsonData {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		argsList = append(argsList, jsonData[k])
	}
	return
}