package config

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/go-ini/ini"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Log 日志基本配置
type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" ini:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"log-file" ini:"log-file" yaml:"log-file" toml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" ini:"stdout"`
	File    string `mapstructure:"file" json:"file" ini:"file"`
}

// KafkaConfig 配置
type KafkaConfig struct {
	Topic            string `json:"topic"`
	GroupId          string `json:"group.id"`
	BootstrapServers string `json:"bootstrap.servers"`
	SecurityProtocol string `json:"security.protocol"`
	SaslMechanism    string `json:"sasl.mechanism"`
	SaslUsername     string `json:"sasl.username"`
	SaslPassword     string `json:"sasl.password"`
}

// Database 连接配置
type Database struct {
	*gorm.DB
	Driver       string `json:"driver" yaml:"driver"`
	Host         string `json:"host" yaml:"host"`
	Port         int    `json:"port" yaml:"port"`
	Database     string `json:"database" yaml:"database"`
	UserName     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	Charset      string `json:"charset" yaml:"charset"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"`
}

// RedisConfig 连接配置
type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// WebServer 服务地址端口配置
type WebServer struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Config struct {
	Logger    Log         `json:"logger"`
	WebConfig WebServer   `json:"web_config"`
	Kafka     KafkaConfig `json:"kafka"`
	DB        Database    `json:"db"`
	Redis     RedisConfig `json:"redis"`
}

var C *Config

func GetConfig() *Config {
	return C
}

func LoadConfigFromJson(configPath string) {
	// 打开文件
	file, _ := os.Open(configPath)
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(C)
	if err != nil {
		panic(err)
	}

	fmt.Println(C)
}

func LoadConfigFromIni(configPath string) {
	err := ini.MapTo(C, configPath)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(C)
}

func LoadConfigFromYaml(configPath string) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, C)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(C)
}

func LoadConfigFromToml(configPath string) {
	_, err := toml.DecodeFile(configPath, C)
	if err != nil {
		panic(err)
	}
	fmt.Println(C)
}

func ParseConfigByViper(configPath, configName, configType string) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	})
	//直接反序列化为Struct
	if err := v.Unmarshal(C); err != nil {
		fmt.Println(err)
	}
	return
}
