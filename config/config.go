package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/scliang-strive/webServerTools/common/global"
)

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" ini:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"log-file" ini:"log-file" yaml:"log-file" toml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" ini:"stdout"`
	File    string `mapstructure:"file" json:"file" ini:"file"`
}

func LoadConfigFromJson() {
	// 打开文件
	file, _ := os.Open("config.json")
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&global.globalConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(global.CONFIG)
}
