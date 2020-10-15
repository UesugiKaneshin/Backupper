package backupper

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigAliyun struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

type ConfigMysql struct {
	Host         string
	Port         string
	User         string
	Password     string
	Databasename string
}

type ConfigRoutiner struct {
	Hour   int
	Minute int
	Number int
}

type Config struct {
	ConfigAliyun   ConfigAliyun
	ConfigMysqls   []ConfigMysql
	ConfigRoutiner ConfigRoutiner
}

type Configer struct {
	configAliyun   *ConfigAliyun
	configMysqls   *[]ConfigMysql
	configRoutiner *ConfigRoutiner
}

func NewConfiger() *Configer {
	this := new(Configer)

	return this
}

func (this *Configer) Read() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	} else {
		var config Config
		viper.Unmarshal(&config)

		this.configAliyun = &config.ConfigAliyun
		this.configMysqls = &config.ConfigMysqls
		this.configRoutiner = &config.ConfigRoutiner
	}
}

func (this *Configer) GetConfigAliyun() *ConfigAliyun {
	return this.configAliyun
}

func (this *Configer) GetConfigMysqls() *[]ConfigMysql {
	return this.configMysqls
}

func (this *Configer) GetConfigRoutiner() *ConfigRoutiner {
	return this.configRoutiner
}
