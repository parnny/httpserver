package pipeline

import (
"fmt"
log "github.com/cihub/seelog"
"io/ioutil"
"net/http"
"os"
"parnny.com/httpserver/config"
"strconv"
"strings"
)

type SeelogPipeline struct {
	BasePipeline
}

func (p *SeelogPipeline) OnInit() (bool) {
	CMInst := config.GetInstance()
	logfile := CMInst.Config.Seelog.Config_path

	cfg_bytes, err := ioutil.ReadFile(logfile)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg_string := string(cfg_bytes)
	cur_pid := int64(os.Getpid())
	cfg_string = strings.Replace(cfg_string, "%processid", strconv.FormatInt( cur_pid, 10), -1)

	logger, err := log.LoggerFromConfigAsString(cfg_string)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	log.ReplaceLogger(logger)
	return true
}

func (p *SeelogPipeline) OnProcess(data string, r *http.Request) (bool) {
	if len(data) == 0 {
		log.Error("Get empty data from request")
		return false
	} else {
		log.Info(data)
		return true
	}
}
