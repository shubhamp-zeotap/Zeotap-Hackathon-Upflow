package app

import (
	"github.com/shubhamp-zeotap/Upflow/app/utils"
	"github.com/shubhamp-zeotap/Upflow/app/workflow"
)

//Check later
const pluginsConfPath string = "./github.com/shubhamp-zeotap/Upflow/app/plugins.json"
const workflowsConfPath string = "./github.com/shubhamp-zeotap/Upflow/app/workflows.json"
const systemInputsConfPath string = "./github.com/shubhamp-zeotap/Upflow/app/systeminputs.json"


func SendPlugins() map[string]interface{} {
	return utils.ReadFileToJson(pluginsConfPath)
}

func SendWorkflows() map[string]interface{} {
	return utils.ReadFileToJson(workflowsConfPath)
}

func SendSystemInputs() map[string]interface{} {
	return utils.ReadFileToJson(systemInputsConfPath)
}


func UpdateWorkflow(channel string, newConfig map[string]interface{}) map[string]interface{} {
	workflowsConfig := utils.ReadFileToJson(workflowsConfPath)
	workflowsConfig[channel] = newConfig
	utils.WriteJsonToFile(workflowsConfPath, workflowsConfig)
	response := make(map[string]interface{})
	response["status"] = true
	return response
}


func RunWorkflow(channel string, inputData map[string]interface{}) map[string]interface{} {
	workflowsConfig := utils.ReadFileToJson(workflowsConfPath)
	channelConfig := workflowsConfig[channel].(map[string]interface{})
	go workflow.ExecuteWorkflow(channelConfig, inputData)
	response := make(map[string]interface{})
	response["status"] = true
	return response
}


