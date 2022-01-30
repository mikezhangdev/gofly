package main

import (
	"log"
	"myblog/core"
	"myblog/core/app"
	"myblog/router"
)

var AppName = "myblog"
var Addr = ":8081"
var ConfigPath = "./"
var ConfigName = "config"

func main() {
	r, err := app.NewRunner(&core.Application{
		Name:       AppName,
		Addr:       Addr,
		ConfigPath: ConfigPath,
		ConfigName: ConfigName,
		LoadConfig: nil,
		GinHandler: router.GinHandlerRouter,
	})
	if err != nil {
		log.Fatalf("app.NewRunner err %v", err)
	}
	err = r.ListenGinServer()
	if err != nil {
		log.Fatalf("r.ListenGinServer err %v", err)
	}
}
