package main

import (
	"flag"
	"os"
	"parnny.com/httpserver/config"
	"parnny.com/httpserver/pipeline"
	"parnny.com/httpserver/service"
)

func main() {
	cfg_path := flag.String("config", "config/examples/server.toml", "config path")
	flag.Parse()

	CMInst := config.GetInstance()
	if CMInst == nil {
		panic("Can not get config manager")
		os.Exit(1)
	}
	CMInst.Load(*cfg_path)

	PMInst := pipeline.GetInstance()
	if PMInst == nil {
		panic("Can not get pipeline manager")
		os.Exit(1)
	}

	print()

	service.Start()
}