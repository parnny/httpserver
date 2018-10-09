package pipeline

import (
	"net/http"
	"reflect"
	"sync"

	utils "github.com/parnny/utils4go"
)

type PipelineManager struct {
	TempDataMap 	*utils.SafeMap
	Processors      [16]BaseInterface
}

func (pm *PipelineManager) Init() {
	pm.TempDataMap = utils.NewSafeMap()
	pm.Processors[0] = &DebugPipeline{}
	pm.Processors[1] = &FlashlogPipeline{}

	for _, pipeline := range pm.Processors {
		if pipeline != nil {
			pipeline.OnInit()
		}
	}
}

func (pm *PipelineManager) SetTempData(key string, value interface{}) {
	pm.TempDataMap.Set(key, value)
}

func (pm *PipelineManager) GetTempData(key string) (interface{}, bool) {
	value, ok := pm.TempDataMap.Get(key)
	return value, ok
}

func (pm *PipelineManager) OnProcess(data *string, r *http.Request) (int, string) {
	if nil == data {
		return 202, "[OnProcess] Failed@data check"
	}

	pm.TempDataMap.Clear()

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

	return code, err
}

func (pm *PipelineManager)OnExit()  {
	for _, pipeline := range pm.Processors {
		if pipeline != nil {
			pipeline.OnExit()
		}
	}
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

