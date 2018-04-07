package splitbysize

import (
	"fmt"
	"os"
	"strings"
	"errors"
	"reflect"
	"io/ioutil"
	"github.com/shubhamp-zeotap/Upflow/app/utils"
	template "github.com/shubhamp-zeotap/Upflow/app/plugins"
)

type PluginInput struct {
	Source string `{"Default": "0", "CanUseDefault": false}`
	Target string `{"Default": "0", "CanUseDefault": false}`
	Size float64 `{"Default": 0, "CanUseDefault": false}`
	Unit string `{"Default": "m", "CanUseDefault": true}`
	DeleteOriginal bool `{"Default": true, "CanUseDefault": true}`
}

type PluginOutput struct {
	Target string
}

type SplitBySize struct {
	template.BasePlugin
	Input PluginInput
	Output PluginOutput
}

func (s SplitBySize) Execute() (r template.PluginReturnValue, err error) {
	target := splitFilesBySize(s.Input.Source, s.Input.Size, s.Input.Unit, s.Input.Target, s.Input.DeleteOriginal)
	s.Output = PluginOutput{
		Target : target,
	}

	r = template.PluginReturnValue{}
	r.Values = s.ConvertToInterface(target)
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
	utils.GeneratePlugin("SplitBySize", pluginConf)
}


func splitFilesBySize(sourceDir string, size float64, unit string, target string, deleteOriginal bool) (string) {
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		splitCmd := fmt.Sprintf("split -b %v%s %s%s %s", size, unit, sourceDir, strings.Replace(f.Name(), " ", "\\ ", -1), target)
		utils.RunCommand(splitCmd)
		if deleteOriginal == true {
			err := os.Remove(sourceDir + f.Name())
			if err != nil {
				panic(err)
			}
		}
	}
	return target
}


func Create(args... interface{}) (r *SplitBySize, err error) {
	arg0, ok := args[0].(string)
	if !ok {
		return nil, errors.New("string expected as input 1")
	}

	arg1, ok := args[1].(float64)
	if !ok {
		return nil, errors.New("int expected as input 2")
	}

	arg2, ok := args[2].(string)
	if !ok || (arg2 != "k" && arg2 != "m")  {
		return nil, errors.New("string (k | m) expected as input 3")
	}

	arg3, ok := args[3].(string)
	if !ok {
		return nil, errors.New("string expected as input 4")
	}

	arg4, ok := args[4].(bool)
	if !ok {
		return nil, errors.New("bool expected as input 5")
	}

	plugin := SplitBySize{}
	plugin.Input = PluginInput{
		Source : arg0, 
		Size : arg1,
		Unit : arg2,
		Target : arg3,
		DeleteOriginal : arg4,
	}

	plugin.Output = PluginOutput{}
	r = &plugin
	return
}