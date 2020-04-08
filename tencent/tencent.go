package tencent

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"markdown-img-upload/utils"
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
	bucketURL := utils.UploadConfig.TencentBucketURL
	secretID := utils.UploadConfig.TencentSecretID
	secretKey := utils.UploadConfig.TencentSecretKey
	t = &Tencent{
		BucketURL: bucketURL,
		SecretID:  secretID,
		SecretKey: secretKey,
	}
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
	defer response.Body.Close()
	resultURL := response.Request.URL.String()
	log.Println("resultURL: ", resultURL)
	return resultURL, nil
}
