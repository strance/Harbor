package master

import (
	"encoding/json"
	"io/ioutil"
)

// 单例
var G_config *Config

type Config struct {
	ApiPort         int `json:"apiPort"`
	ApiReadTimeOut  int `json:"apiReadTimeOut"`
	ApiWriteTimeOut int `json:"apiWriteTimeOut"`
}

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		conf Config
		content []byte
	)

	// 1.读取配置文件
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2.反序列化json
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3.赋值单例
	G_config = &conf

	return
}