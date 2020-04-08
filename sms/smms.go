package sms

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	SMS_URL = "https://sm.ms/api/v2/sms"
)

type SMMS struct{}

func (smms *SMMS) Upload(path string) (string, error) {
	cloudPath := ""
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fn, _ := filepath.Abs(path)
	formFile, err := writer.CreateFormFile("smfile", fn)

	if err != nil {
		return "", err
	}
	// 读取文件并定稿表单
	srcFile, err := os.Open(fn)
	if err != nil {
		return cloudPath, err
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		return cloudPath, err
	}
	//_ = writer.WriteField("file_id", "0") // 额外参数，比如七牛上传时需要提供token
	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() //发送之前必须调用Close()以写入结尾行
	resp, err := http.Post(SMS_URL, contentType, buf)
	// mutex.Unlock()
	if err != nil {
		return cloudPath, err
	}
	defer resp.Body.Close()
	ret := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return cloudPath, err
	}
	log.Println("response: ", ret)
	success := ret["success"].(bool)
	if success {
		cloudPath = ret["data"].(map[string]interface{})["url"].(string)
	} else {
		cloudPath = ret["images"].(string)
	}
	return cloudPath, nil
}
