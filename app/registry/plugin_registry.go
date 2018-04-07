package registry

import (
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/s3download"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/s3transfer"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/splitbysize"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/transform"
	"github.com/shubhamp-zeotap/Upflow/app/plugins/crafts/unzip"
)

func RegisterGlobalPlugins() {
	s3download.Register()
	splitbysize.Register()
	s3transfer.Register()
	transform.Register()
	unzip.Register()
}