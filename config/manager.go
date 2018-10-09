package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"
)

type ConfigManager struct {
	Config	TomlConfig
}

func (cm *ConfigManager) Load(path string) {
	_, err := toml.DecodeFile(path, &cm.Config)
	if err != nil {
		panic(err)
	}
	fmt.Println(cm)
}

func (cm *ConfigManager)String() string {
	b, err := json.Marshal(cm.Config)
	if err != nil {
		return fmt.Sprintf("%+v", cm.Config)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", cm.Config)
	}
	return out.String()
}

var CfgInst *ConfigManager
var once sync.Once

func GetInstance() *ConfigManager {
	once.Do(func() {
		CfgInst = &ConfigManager{}
	})
	return CfgInst
}



