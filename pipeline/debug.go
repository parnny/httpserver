package pipeline

import (
	"net/http"
)

type DebugPipeline struct {
	BasePipeline
}

func (p *DebugPipeline) OnProcess(data *string, r *http.Request) (bool) {
	if nil == data {
		return false
	}
	if len(*data) == 0{
//		*data = "debug fixed data"
	}
	return true
}
