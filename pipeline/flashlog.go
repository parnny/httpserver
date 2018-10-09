package pipeline

import (
	log "github.com/cihub/seelog"
	"io/ioutil"
	"net/http"
	"os"
	"parnny.com/datalog/config"
	"parnny.com/datalog/thirdparty"
	"path/filepath"
	"time"

	utils "github.com/parnny/utils4go"
)

const default_timetick_timeout_logfile = 10 * time.Second
const default_timetick_empty_directory  = 60 * time.Second

type FlashlogPipeline struct {
	BasePipeline
	LoggerMap			*utils.SafeMap
	FlashConfig			config.FlashlogCoreConfig
	RecycleFlag			bool
}

func (p *FlashlogPipeline) OnInit() (bool) {
	CMInst := config.GetInstance()
	p.FlashConfig = CMInst.Config.Flashlog

	p.LoggerMap = utils.NewSafeMap()

	go func() {
		interval := default_timetick_timeout_logfile
		if p.FlashConfig.Timertick.Timeout_logfile > 0 {
			interval = p.FlashConfig.Timertick.Timeout_logfile * time.Second
		}
		for range time.Tick(interval) {
			p.TimeoutCheck()
		}
	}()

	go func() {
		interval := default_timetick_empty_directory
		if p.FlashConfig.Timertick.Empty_directory > 0 {
			interval = p.FlashConfig.Timertick.Empty_directory * time.Second
		}
		for range time.Tick(interval) {
			p.RecycleFlag = true
		}
	}()


	if p.FlashConfig.Monitor.Active {
		M, err := thirdparty.NewFileMonitor()
		if err != nil {
			log.Error(err)
		} else {
			M.Start()
			M.Watch(p.FlashConfig.Logpath)
		}
	}

	return true
}

func (p *FlashlogPipeline) GetLogger(flsinfo *utils.FlashlogInfo ) (*utils.FlashlogObj) {
	logger, ok := p.LoggerMap.Get(flsinfo.Filepath)
	if ok {
		return logger.(*utils.FlashlogObj)
	} else {

		logger := new(utils.FlashlogObj)
		if logger.Init(flsinfo) {
			p.LoggerMap.Set(flsinfo.Filepath, logger)
		}
		return logger
	}
	return nil
}

func (p *FlashlogPipeline) OnProcess(data *string, r *http.Request) (bool) {
	if p.RecycleFlag {
		p.Recycle(p.FlashConfig.Logpath)
		p.RecycleFlag = false
	}

	var msgs []BaseJsonMsg
	err := json.Unmarshal([]byte(*data), &msgs)

	if err != nil {
		log.Error(err)
		return false
	}

	for _, obj := range msgs {
		flsinfo := p.GenFlashlogInfo(obj)
		logger := p.GetLogger(flsinfo)
		if logger != nil {
			logger.Write([]byte(*data))
		}
	}
	return true
}

func (p *FlashlogPipeline) OnExit() {
	p.LoggerMap.Foreach(func(key string, val interface{}) {
		logger := val.(*utils.FlashlogObj)
		logger.Close(utils.LogCloseReasonExit)
	})
	p.LoggerMap.Clear()
}

func (p *FlashlogPipeline)TimeoutCheck() {
	p.LoggerMap.Foreach(func(key string, val interface{}) {
		 logger := val.(*utils.FlashlogObj)
		 if logger.IsTimeout() {
		 	logger.Close(utils.LogCloseReasonTimeout)
		 	delete(p.LoggerMap.Data, key)
		 }
	})
}

func(p *FlashlogPipeline)Recycle(path string) {
	childs, _ := ioutil.ReadDir(path)

	if len(childs) == 0 {
		os.RemoveAll(path)
		return
	}

	for _, child := range childs {
		if child.IsDir() {
			p.Recycle(filepath.Join(path,child.Name()))
		}
	}
}

func (p *FlashlogPipeline)GenFlashlogInfo(obj BaseJsonMsg) *utils.FlashlogInfo {
	ts_now := time.Now().Unix()
	ts_msg := int64(obj.Timestamp)
	var ts_diff int64
	if ts_now > ts_msg {
		ts_diff = ts_now - ts_msg
	} else {
		ts_diff = ts_msg - ts_now
	}

	if ts_diff <= p.FlashConfig.Threshold {
		return utils.GenFlashlogInfo(obj.Timestamp,
			p.FlashConfig.Standard.Timestep,
			p.FlashConfig.Standard.Rollsize,
			p.FlashConfig.Logpath,
			obj.Msgtype)
	} else {
		return utils.GenFlashlogInfo(obj.Timestamp,
			p.FlashConfig.Nonstandard.Timestep,
			p.FlashConfig.Nonstandard.Rollsize,
			p.FlashConfig.Logpath,
			obj.Msgtype)
	}
	return nil
}