package backupper

import (
	"fmt"
	"log"
	"time"
)

type Routiner struct {
	date time.Time

	hour   int
	minute int
	number int

	configer *Configer
	dumper   *Dumper
	uploader *Uploader
}

func NewRoutiner() *Routiner {
	this := new(Routiner)

	this.date, _ = time.Parse("2016-01-02 15:04:05", "2010-10-10 10:10:10")

	this.configer = NewConfiger()
	this.dumper = NewDumper()
	this.uploader = NewUploader()

	return this
}

func (this *Routiner) SetDefaultValue(hour int, minute int, number int) {
	this.hour = hour
	this.minute = minute
	this.number = number
}

func (this *Routiner) Run() {
	for {
		this.getConfig()

		now := time.Now()
		if now.Format("2006-01-02") != this.date.Format("2006-01-02") && now.Hour() > this.hour && now.Minute() > this.minute {
			this.do()
			this.date = now

			log.Println(now)
		} else {
		}

		time.Sleep(time.Minute)
	}
}

func (this *Routiner) getConfig() {
	this.configer.Read()

	configRoutiner := *this.configer.GetConfigRoutiner()

	this.SetDefaultValue(configRoutiner.Hour, configRoutiner.Minute, configRoutiner.Number)
}

func (this *Routiner) do() {
	configAliyun := *this.configer.GetConfigAliyun()
	configMysqls := *this.configer.GetConfigMysqls()

	this.uploader.SetDefaultValue(configAliyun.EndPoint, configAliyun.AccessKeyId, configAliyun.AccessKeySecret, configAliyun.BucketName)
	this.uploader.Connect()

	for _, val := range configMysqls {
		timeStr := time.Now().Format("20060102150405")

		filename := fmt.Sprintf("%s%s.sql", val.Databasename, timeStr)
		filepath := fmt.Sprintf("./tmp/%s", filename)

		objectpath := fmt.Sprintf("database/%s", val.Databasename)
		objectName := fmt.Sprintf("%s/%s", objectpath, filename)

		this.dumper.SetDefaultValue(val.Host, val.Port, val.User, val.Password, val.Databasename, filepath)
		this.dumper.Download()

		this.uploader.Upload(objectName, filepath)
		list := this.uploader.Retrieve(objectpath)

		this.uploader.Delete(list, this.number)

		this.dumper.DeleteFile()
	}
}
