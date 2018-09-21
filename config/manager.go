package config

import (
	"sync"
	"github.com/BurntSushi/toml"
)

type ConfigManager struct {
	Config	TomlConfig
}

func (cm *ConfigManager) Load(path string) {
	_, err := toml.DecodeFile(path, &cm.Config)
	if err != nil {
		panic(err)
	}
}

var CfgInst *ConfigManager
var once sync.Once

func GetInstance() *ConfigManager {
	once.Do(func() {
		CfgInst = &ConfigManager{}
	})
	return CfgInst
}



