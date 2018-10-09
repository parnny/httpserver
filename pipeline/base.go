package pipeline

import "net/http"

type BaseInterface interface {
	OnProcess(data *string, r *http.Request) bool
	OnInit() bool
	OnExit()
}

type BasePipeline struct {
	BaseInterface
}

func (p *BasePipeline) OnInit() (bool) {
	return true
}

func (p *BasePipeline) OnProcess(data *string, r *http.Request) (bool) {
	return true
}

func (p *BasePipeline)OnExit()  {

}
