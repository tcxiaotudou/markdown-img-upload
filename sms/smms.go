package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"markdown-img-upload/utils"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	SMS_URL = "https://sm.ms/api/v2/upload"
)

type SMMS struct {
	token string
}

func (smms *SMMS) Upload(path string) (string, error) {
	cloudPath := ""
	// 打开要上传的文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer file.Close()
	body := &bytes.Buffer{}
	// 创建一个multipart类型的写文件
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("smfile", path)
	if err != nil {
		return "", err
	}
	// 将源复制到目标，将file写入到part   是按默认的缓冲区32k循环操作的，不会将内容一次性全写入内存中,这样就能解决大文件的问题
	_, err = io.Copy(formFile, file)
	err = writer.Close()
	if err != nil {
		fmt.Println("write file err:", err)
	}
	request, err := http.NewRequest("POST", SMS_URL, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	token := utils.UploadConfig.SmsToken
	request.Header.Set("Authorization", token)
	clt := http.Client{}
	resp, err := clt.Do(request)
	// mutex.Unlock()
	if err != nil {
		return cloudPath, err
	}
	defer resp.Body.Close()
	ret := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return cloudPath, err
	}
	success := ret["success"].(bool)
	if success {
		cloudPath = ret["data"].(map[string]interface{})["url"].(string)
	} else {
		log.Fatal("upload fail: ", ret)
	}
	log.Println("resultURL: ", cloudPath)
	return cloudPath, nil
}
