package s3transfer

import (
	"fmt"
	"errors"
	"reflect"
	"github.com/shubhamp-zeotap/Upflow/app/utils"
	template "github.com/shubhamp-zeotap/Upflow/app/plugins"
)

type PluginInput struct {
	Source string `{"Default": "0", "CanUseDefault": false}`
	Target string `{"Default": "0", "CanUseDefault": false}`
	IsSourceDir bool `{"Default": true, "CanUseDefault": true}`
	Profile string `{"Default": "", "CanUseDefault": true}`
}

type PluginOutput struct {
	Target string
}

type S3Transfer struct {
	template.BasePlugin
	Input PluginInput
	Output PluginOutput
}

func (s3t S3Transfer) Execute() (r template.PluginReturnValue, err error) {
	transferFromS3(s3t.Input.Source, s3t.Input.Target, s3t.Input.IsSourceDir, s3t.Input.Profile)
	s3t.Output = PluginOutput{
		Target : s3t.Input.Target,
	}

	r = template.PluginReturnValue{}
	r.Values = s3t.ConvertToInterface(s3t.Input.Target)
	return		
}


func Register() {
	pluginConf := make(map[string]interface{})
	pluginConf["CanBeFirst"] = false
	pluginConf["ShouldBeFirstOnly"] = false
	pluginConf["Description"] = "Some description"

	i := reflect.ValueOf(&PluginInput{}).Elem()
	pluginConf["InputConfig"] = utils.ReflectValueToJsonArray(i)

	o := reflect.ValueOf(&PluginOutput{}).Elem()
	pluginConf["OutputConfig"] = utils.ReflectValueToJsonArray(o)

	//Change const name to struct derived name
	utils.GeneratePlugin("S3Transfer", pluginConf)
}


func transferFromS3(source string, target string, isSourceDir bool, profile string) {
	s3CpCmd := fmt.Sprintf("aws s3 cp %s %s", source, target)

	if isSourceDir {
		s3CpCmd = fmt.Sprintf("%s --recursive", s3CpCmd)
	}

	if profile != "" {
		s3CpCmd = fmt.Sprintf("%s --profile %s", s3CpCmd, profile)
	}

	utils.RunCommand(s3CpCmd)
}


func Create(args... interface{}) (r *S3Transfer, err error) {
	arg0, ok := args[0].(string)
	if !ok {
		return nil, errors.New("string expected as input 1")
	}

	arg1, ok := args[1].(string)
	if !ok {
		return nil, errors.New("string expected as input 2")
	}

	arg2, ok := args[2].(bool)
	if !ok {
		return nil, errors.New("bool expected as input 3")
	}

	arg3, ok := args[3].(string)
	if !ok {
		arg3 = ""
	}

	plugin := S3Transfer{}
	plugin.Input = PluginInput{
		Source : arg0, 
		Target : arg1, 
		IsSourceDir : arg2, 
		Profile : arg3,				
	}

	plugin.Output = PluginOutput{}
	r = &plugin
	return
}