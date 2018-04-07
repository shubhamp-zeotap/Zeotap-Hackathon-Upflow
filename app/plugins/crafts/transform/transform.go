package transform

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"io/ioutil"
	"errors"
	"reflect"
	"github.com/shubhamp-zeotap/Upflow/app/utils"
	template "github.com/shubhamp-zeotap/Upflow/app/plugins"
)

type PluginInput struct {
	Source string `{"Default": "0", "CanUseDefault": false}`
	ExportGuid string `{"Default": "0", "CanUseDefault": false}`
	SegmentMapping map[string]interface{} `{"Default": {}, "CanUseDefault": false}`
	DoParallel bool `{"Default": true, "CanUseDefault": true}`
}

type PluginOutput struct {
	Target string
}

type Transform struct {
	template.BasePlugin
	Input PluginInput
	Output PluginOutput
}

func (t Transform) Execute() (r template.PluginReturnValue, err error) {
	target := transform(t.Input.Source, t.Input.ExportGuid, t.Input.SegmentMapping, t.Input.DoParallel)
	t.Output = PluginOutput{
		Target : target,
	}

	r = template.PluginReturnValue{}
	r.Values = t.ConvertToInterface(target)
	return		
}

func Register() {
	pluginConf := make(map[string]interface{})
	pluginConf["CanBeFirst"] = false
	pluginConf["ShouldBeFirstOnly"] = false
	pluginConf["Description"] = "Daap Segment ID to Channel Segment ID transformation"

	i := reflect.ValueOf(&PluginInput{}).Elem()
	pluginConf["InputConfig"] = utils.ReflectValueToJsonArray(i)

	o := reflect.ValueOf(&PluginOutput{}).Elem()
	pluginConf["OutputConfig"] = utils.ReflectValueToJsonArray(o)

	//Change const name to struct derived name
	utils.GeneratePlugin("Transform", pluginConf)
}


func transform(source string, exportGuid string, segmentMapping map[string]interface{}, doParallel bool) (target string) {
	target = fmt.Sprintf(os.ExpandEnv("$HOME/output/%s/"), exportGuid)
	cleanCmd := fmt.Sprintf("rm -rf %s", target)
	utils.RunCommand(cleanCmd)	

	mkdirCmd := fmt.Sprintf("mkdir -p %s", target)
	utils.RunCommand(mkdirCmd)

	files, err := ioutil.ReadDir(source)
	if err != nil {
		panic(err)
	}

	jobs := make(chan int, len(files))
	results := make(chan bool, len(files))

	var workers int
	if doParallel == true {
		workers = len(files)
	} else {
		workers = 1
	}

	//create workers
	for w := 0; w < workers; w++ {
		go transformSingleFile(w, source, files, target, segmentMapping, jobs, results)
	}

	//distribute jobs	
	for w := 0; w < workers; w++ {
		jobs <- w
	}
	close(jobs)

	//wait for completion
	for w := 0; w < workers; w++ {
		<-results
	}
	return
}


func transformSingleFile(id int, source string, files []os.FileInfo, target string, segmentMapping map[string]interface{}, jobs <- chan int, results chan <- bool) {
	for j := range jobs {
		fmt.Println("Executing worker Id : ", j)
		sf, err := os.Open(source + files[id].Name())
		if err != nil {
			panic(err)
		}
		defer sf.Close()
		
		df, err := os.OpenFile(target + files[id].Name(), os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer df.Close()

		reader := bufio.NewReader(sf)
		var line string
		for {
			line, err = reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			transformedLine := transformSingleLine(line, segmentMapping)
			df.Write([]byte(transformedLine))
		}

		results <- true
	}
}


func transformSingleLine(line string, segmentMapping map[string]interface{}) (string) {
	data := strings.Split(strings.TrimSpace(line), "\t")
	segments := data[len(data)-1]
	daapIds := strings.Split(segments, ",")
	channelIds := make([]string, 0, len(daapIds))
	for _, daapId := range daapIds {
		if daapId != "" {	
			channelIds = append(channelIds, segmentMapping[daapId].(string))			
		}
	}
	data[len(data)-1] = strings.Join(channelIds, ",")
	return strings.Join(data, "\t") + "\n"
}


func Create(args... interface{}) (r *Transform, err error) {
	arg0, ok := args[0].(string)
	if !ok {
		return nil, errors.New("string expected as input 1")
	}

	arg1, ok := args[1].(string)
	if !ok {
		return nil, errors.New("string expected as input 2")
	}

	arg2, ok := args[2].(map[string]interface{})
	if !ok {
		return nil, errors.New("map[string][string] expected as input 3")
	}

	arg3, ok := args[3].(bool)
	if !ok {
		return nil, errors.New("bool expected as input 4")
	}

	plugin := Transform{}
	plugin.Input = PluginInput{
		Source : arg0, 
		ExportGuid : arg1, 
		SegmentMapping : arg2,
		DoParallel : arg3,
	}

	plugin.Output = PluginOutput{}
	r = &plugin
	return
}

