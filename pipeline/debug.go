package pipeline

import (
"net/http"
)

type DebugPipeline struct {
	BasePipeline
}

func (p *DebugPipeline) OnProcess(data string, r *http.Request) (bool) {
	return true
}
