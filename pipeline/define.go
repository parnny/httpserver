package pipeline

import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary


type BaseJsonMsg struct {
	Appname		string
	Msgtype		string
	Timestamp	int64
}

