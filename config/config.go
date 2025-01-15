// https://cobra.dev/
package config

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"os"
)

type DBConfig struct {
	schmea			string	`json:"schema" description:"database type" default:"postgres"`
	Host			string	`json:"host" description:"database host" default:"172.25.1.22"`
	Port			int		`json:"port" description:"database port" default:"5432"`
	UserName		string	`json:"username" description:"database username" default:"habitformation"`
	Password		string	`json:"password" description:"database password" default:"habitformation"`
	DBName			string	`json:"dbname" description:"database name" default:"1"`
	MaxIdleConns	int		`json:"maxIdleConns" description:"database connect max idle" default:"10"`
	MaxOpenConns	int		`json:"maxOpenConns" description:"database connect max open" default:"10"`
	SSLMode			string	`json:"sslmode" description:"database SSL mode" default:"disable"`
}

type WEBConfig struct {
	Port int `json:"port" description:"web server port" default:"8080"`
}

type Configuration struct {
	ConfigFile string
	DBConfig DBConfig `json:"database" description:"database config" default:"{}"`
	WEBConfig WEBConfig `json:"web" description:"web server config" default:"{}"`
}

var Config *Configuration

func (config *Configuration) loadData() error {
	log.Printf("解析配置文件开始")
	// 打开json文件
	jsonFile, err := os.Open(config.ConfigFile)

	// 最好要处理以下错误
	if err != nil {
		log.Printf("加载配置文件%v失败, err: %v\n", config.ConfigFile, err)
		return err
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Printf("配置文件解析为Json报错, err: %v\n", err)
		return err
	}

	json.MarshalIndent(config, "", " ")
	log.Printf("解析配置文件结束，配置数据为: %v", config)
	return err
}

