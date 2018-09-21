package service

import (
	"fmt"
	"net/http"
	"parnny.com/httpserver/config"
	"parnny.com/httpserver/pipeline"
	"parnny.com/httpserver/utils"
)

func OnProcess(done chan bool, w http.ResponseWriter, r *http.Request)  {
	PMInst := pipeline.GetInstance()
	len := r.ContentLength
	bytes := make([]byte, len)
	r.Body.Read(bytes)
	body := string(bytes)
	code, err := PMInst.OnProcess(body,r)
	w.WriteHeader(code)
	if code == 200 {
		fmt.Fprintf(w,"hello world from goroutine(%s):success", utils.GetGID())
	} else {
		fmt.Fprintf(w,"hello world from goroutine(%s):error %s", utils.GetGID(), err)
	}
	done <- true
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go OnProcess(done,w,r)
	<-done
}

func Start() {
	http.HandleFunc("/", IndexHandler)
	CMInst := config.GetInstance()
	addr := CMInst.Config.Http.Server_ip_port
	fmt.Printf("Service start @ %s \n", addr)
	http.ListenAndServe(addr, nil)
}
