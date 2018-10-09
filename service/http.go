package service

import (
	"fmt"
	"net/http"
	"parnny.com/datalog/config"
	"parnny.com/datalog/pipeline"
)

func OnProcess(done chan bool, w http.ResponseWriter, r *http.Request)  {
	PMInst := pipeline.GetInstance()
	len := r.ContentLength
	bytes := make([]byte, len)
	r.Body.Read(bytes)
	body := string(bytes)
	code, _ := PMInst.OnProcess(&body,r)
	w.WriteHeader(code)
	done <- true
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go OnProcess(done,w,r)
	<-done
	close(done)
}

func Start() {
	http.HandleFunc("/", IndexHandler)
	CMInst := config.GetInstance()
	addr := CMInst.Config.Http.Server_ip_port
	fmt.Printf("Service start @ %s \n", addr)
	http.ListenAndServe(addr, nil)
}
