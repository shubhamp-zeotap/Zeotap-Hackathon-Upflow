package registry

import (
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/s3download"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/s3transfer"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/splitbysize"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/transform"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/unzip"
	"github.com/shubhamp-zeotap/Upflow/app/plugins"
)


func check(err error) {
	if err != nil {
		panic(err)
	}	
}


func GetPlugin(pluginName string, args... interface{}) (p plugins.Plugin) {
	switch pluginName {
		case "S3Download" : 
			p, err := s3download.Create(args...)
			check(err)
			return p
		case "SplitBySize" : 
			p, err := splitbysize.Create(args...)
			check(err)
			return p
		case "S3Transfer" : 
			p, err := s3transfer.Create(args...)
			check(err)
			return p
		case "Transform" : 
			p, err := transform.Create(args...)
			check(err)
			return p
		case "Unzip" :
			p, err := unzip.Create(args...)
			check(err)
			return p
		default:
			panic("No such plugin found !")
	}
}
