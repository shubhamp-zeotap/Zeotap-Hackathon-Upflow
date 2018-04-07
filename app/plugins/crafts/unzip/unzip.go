package unzip

import (
	"fmt"
	"errors"
	"reflect"
	"github.com/shubhamp-zeotap/Upflow/app/utils"
	template "github.com/shubhamp-zeotap/Upflow/app/plugins"
)


type PluginInput struct {
	Source string `{"Default": "0", "CanUseDefault": false}`
}

type PluginOutput struct {
	Target string
}


type Unzip struct {
	template.BasePlugin
	Input PluginInput
	Output PluginOutput
}


func (u Unzip) Execute() (r template.PluginReturnValue, err error) {
	gunzipFiles(u.Input.Source)
	u.Output = PluginOutput{
		Target : u.Input.Source,
	}

	r = template.PluginReturnValue{}
	r.Values = u.ConvertToInterface(u.Input.Source)
	return		
}


func gunzipFiles(source string) {
	gunzipCmd := fmt.Sprintf("gunzip %s*", source)
	utils.RunCommand(gunzipCmd)
	return
}


func Register() {
	pluginConf := make(map[string]interface{})
	pluginConf["CanBeFirst"] = false
	pluginConf["ShouldBeFirstOnly"] = false
	pluginConf["Description"] = "Unzip gz files in directory"

	i := reflect.ValueOf(&PluginInput{}).Elem()
	pluginConf["InputConfig"] = utils.ReflectValueToJsonArray(i)

	o := reflect.ValueOf(&PluginOutput{}).Elem()
	pluginConf["OutputConfig"] = utils.ReflectValueToJsonArray(o)

	//Change const name to struct derived name
	utils.GeneratePlugin("Unzip", pluginConf)
}


func Create(args... interface{}) (r *Unzip, err error) {
	arg0, ok := args[0].(string)
	if !ok {
		return nil, errors.New("string expected as input 1")
	}

	plugin := Unzip{}
	plugin.Input = PluginInput{
		Source : arg0, 
	}

	plugin.Output = PluginOutput{}
	r = &plugin
	return
}


