package s3download

import (
	"fmt"
	"os"
	"errors"
	"reflect"
	"github.com/shubhamp-zeotap/Upflow/app/utils"
	template "github.com/shubhamp-zeotap/Upflow/app/plugins"
)

type PluginInput struct {
	Source string `{"Default": "0", "CanUseDefault": false}`
	ExportGuid string `{"Default": "0", "CanUseDefault": false}`
	IsSourceDir bool `{"Default": true, "CanUseDefault": true}`
	Profile string `{"Default": "", "CanUseDefault": true}`
}

type PluginOutput struct {
	Target string
}

type S3Download struct {
	template.BasePlugin
	Input PluginInput
	Output PluginOutput
}

func (s3d S3Download) Execute() (r template.PluginReturnValue, err error) {
	target := downloadFromS3(s3d.Input.Source, s3d.Input.ExportGuid, s3d.Input.IsSourceDir, s3d.Input.Profile)
	s3d.Output = PluginOutput{
		Target : target,
	}

	r = template.PluginReturnValue{}
	r.Values = s3d.ConvertToInterface(target)
	return		
}


func Register() {
	pluginConf := make(map[string]interface{})
	pluginConf["CanBeFirst"] = true
	pluginConf["ShouldBeFirstOnly"] = false
	pluginConf["Description"] = "Some description"

	i := reflect.ValueOf(&PluginInput{}).Elem()
	pluginConf["InputConfig"] = utils.ReflectValueToJsonArray(i)

	o := reflect.ValueOf(&PluginOutput{}).Elem()
	pluginConf["OutputConfig"] = utils.ReflectValueToJsonArray(o)

	//Change const name to struct derived name
	utils.GeneratePlugin("S3Download", pluginConf)
}


func downloadFromS3(s3Path string, exportGuid string, isSourceDir bool, profile string) (target string) {
	target = fmt.Sprintf(os.ExpandEnv("$HOME/input/%s/"), exportGuid)
	mkdirCmd := fmt.Sprintf("mkdir -p %s", target)
	utils.RunCommand(mkdirCmd)

	s3CpCmd := fmt.Sprintf("aws s3 cp %s %s", s3Path, target)

	if isSourceDir {
		s3CpCmd = fmt.Sprintf("%s --recursive", s3CpCmd)
	}

	if profile != "" {
		s3CpCmd = fmt.Sprintf("%s --profile %s", s3CpCmd, profile)
	}

	utils.RunCommand(s3CpCmd)
	return
}


func Create(args... interface{}) (r *S3Download, err error) {
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

	plugin := S3Download{}
	plugin.Input = PluginInput{
		Source : arg0, 
		ExportGuid : arg1, 
		IsSourceDir : arg2, 
		Profile : arg3,				
	}

	plugin.Output = PluginOutput{}
	r = &plugin
	return
}
