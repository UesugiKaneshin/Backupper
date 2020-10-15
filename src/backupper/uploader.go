package backupper

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Uploader struct {
	endPoint        string
	accessKeyId     string
	accessKeySecret string
	bucketName      string

	client *oss.Client
	bucket *oss.Bucket
}

func NewUploader() *Uploader {
	this := new(Uploader)

	return this
}

func (this *Uploader) SetDefaultValue(endPoint string, accessKeyId string, accessKeySecret string, bucketName string) {
	this.endPoint = endPoint
	this.accessKeyId = accessKeyId
	this.accessKeySecret = accessKeySecret
	this.bucketName = bucketName
}

func (this *Uploader) Connect() bool {
	result := true

	client, err := oss.New(this.endPoint, this.accessKeyId, this.accessKeySecret)
	if err != nil {
		result = false
		log.Fatal(err)
	} else {
	}

	bucket, err := client.Bucket(this.bucketName)
	if err != nil {
		result = false
		log.Fatal(err)
	} else {
	}

	this.client = client
	this.bucket = bucket

	return result
}

func (this *Uploader) Upload(objectName string, objectPath string) bool {
	result := true

	err := this.bucket.PutObjectFromFile(objectName, objectPath)
	if err != nil {
		result = false
		log.Fatal(err)
	} else {
	}

	return result
}

func (this *Uploader) Retrieve(prefix string) []string {
	list, err := this.bucket.ListObjects(oss.Prefix(prefix))
	if err != nil {
		log.Fatal(err)
	} else {
	}

	result := []string{}
	for _, val := range list.Objects {
		result = append(result, val.Key)
	}

	return result
}

func (this *Uploader) Delete(list []string, number int) bool {
	result := true

	if len(list) > number {
		list = list[0 : len(list)-number]

		_, err := this.bucket.DeleteObjects(list)
		if err != nil {
			result = false
			log.Fatal(err)
		}
	} else {
	}

	return result
}
