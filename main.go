package main

import (
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"os"
	"os/signal"
	"parnny.com/datalog/config"
	"parnny.com/datalog/pipeline"
	"parnny.com/datalog/service"
	"parnny.com/datalog/thirdparty"
)

func main() {
	logger, err := log.LoggerFromConfigAsString(thirdparty.GetSeelogConfigContent())
	if nil != err {
		fmt.Println("Can not load seelog config", err)
		os.Exit(1)
	}
	log.UseLogger(logger)

	cfg_path := flag.String("config", "config/examples/server.toml", "config path")
	flag.Parse()

	CMInst := config.GetInstance()
	if CMInst == nil {
		fmt.Println("Can not get config manager")
		os.Exit(1)
	}
	CMInst.Load(*cfg_path)

	PMInst := pipeline.GetInstance()
	if PMInst == nil {
		fmt.Println("Can not get pipeline manager")
		os.Exit(1)
	}

	signalChan := make(chan os.Signal)
	go func() {
		//阻塞程序运行，直到收到终止的信号
		s := <-signalChan
		fmt.Println("Get signal ", s)
		fmt.Println("Cleaning before stop...")
		PMInst.OnExit()
		os.Exit(0)
	}()
	signal.Notify(signalChan)

	service.Start()
}
