package tencent

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

type Tencent struct {
	BucketURL string
	SecretID  string
	SecretKey string
}

func (t *Tencent) Upload(path string) (string, error) {
	t = ParseConfig()
	u, _ := url.Parse(t.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretID,
			SecretKey: t.SecretKey,
		},
	})
	fileName := filepath.Base(path)
	response, err := c.Object.PutFromFile(context.Background(), fileName, path, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	resultURL := response.Request.URL.String()
	log.Println("resultURL: ", resultURL)
	response.Body.Close()
	return resultURL, nil
}

func ParseConfig() *Tencent {
	conf, err := ini.Load("config.ini") //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		panic(err)
	}
	bucketURL := conf.Section("tencent").Key("tencent.bucket_url").String()
	secretID := conf.Section("tencent").Key("tencent.secret_id").String()
	secretKey := conf.Section("tencent").Key("tencent.secret_key").String()
	return &Tencent{
		BucketURL: bucketURL,
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}
