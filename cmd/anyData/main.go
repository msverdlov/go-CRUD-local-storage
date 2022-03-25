package main

import (
	"anyData/internal/app/anyData/config"
	"anyData/internal/app/anyData/routes"
)

var url string

func main() {
	config.SetConfig(config.ConfFilePath, config.ConfFileName, config.ConfFileExt)
	url = config.GetViperValueByKey("server.host") + ":" + config.GetViperValueByKey("server.port")
	err := routes.InitRoutes().Run(url)
	if err != nil {
		return
	}
}