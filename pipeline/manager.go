package pipeline

import (
	"net/http"
	"parnny.com/httpserver/utils"
	"reflect"
	"sync"
)

type PipelineManager struct {
	TempDataMap 	*utils.SafeMap
	Processors      [16]BaseInterface
	Deocder         BaseInterface
	Writer          BaseInterface
}

func (pm *PipelineManager) Init() {
	pm.TempDataMap = utils.NewSafeMap()
	pm.Processors[0] = &DebugPipeline{}
	pm.Writer = &SeelogPipeline{}

	for _, pipeline := range pm.Processors {
		if pipeline != nil {
			pipeline.OnInit()
		}
	}

	if pm.Writer != nil {
		pm.Writer.OnInit()
	}
}

func (pm *PipelineManager) SetTempData(key string, value interface{}) {
	pm.TempDataMap.Set(key, value)
}

func (pm *PipelineManager) GetTempData(key string) (interface{}, bool) {
	value, ok := pm.TempDataMap.Get(key)
	return value, ok
}

func (pm *PipelineManager) OnProcess(data string, r *http.Request) (int, string) {
	code := 200
	err := ""
	for _, pipeline := range pm.Processors {
		if pipeline != nil {
			if !pipeline.OnProcess(data, r) {
				ptype := reflect.TypeOf(pipeline)
				code = 201
				err = "[OnProcess] Failed@" + ptype.String()
				break
			}
		}
	}

	if pm.Writer != nil {
		if !pm.Writer.OnProcess(data,r){
			ptype := reflect.TypeOf(pm.Writer)
			code = 202
			err = "[OnProcess] Failed@" + ptype.String()
		}
	}
	return code, err
}

var PMInst *PipelineManager
var once sync.Once

func GetInstance() *PipelineManager {
	once.Do(func() {
		PMInst = &PipelineManager{}
		PMInst.Init()
	})
	return PMInst
}

