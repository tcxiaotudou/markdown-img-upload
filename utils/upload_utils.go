package utils

import (
	"bufio"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	upload "markdown-img-upload/interface"
	"os"
	"regexp"
	"strings"
)

var UploadConfig *Config

type Config struct {
	FilePath         string `ini:"file_path"`
	TencentBucketURL string `ini:"tencent_bucket_url"`
	TencentSecretID  string `ini:"tencent_secret_id"`
	TencentSecretKey string `ini:"tencent_secret_key"`
	SmsToken         string `ini:"sms_token"`
}

// 通过文件名，加载配置后拼接完整文件路径
func GetFilePath(fileName string) string {
	err := ReadConfig("config.ini") //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	return UploadConfig.FilePath + fileName
}

// 读取配置文件.设置配置
func ReadConfig(path string) error {
	//var config Config
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return err
	}
	UploadConfig = new(Config)
	conf.MapTo(UploadConfig)
	return nil
}

// 查找和替换markdown文件中的图片路径
func ModifyImg(filepath string, source upload.Upload) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	content := string(data)
	imgRegexp := regexp.MustCompile(`!\[.*?\)`)
	params := imgRegexp.FindAllString(content, -1)
	for _, param := range params {
		imgRegexp = regexp.MustCompile(`\((.*?)\)`)
		imgOriginUrl := imgRegexp.FindAllString(param, -1)
		r := strings.NewReplacer("(", "", ")", "")
		targetUrl := r.Replace(imgOriginUrl[0])
		// 上传图片, 返回cloudPath
		cloudPath, err := source.Upload(targetUrl)
		if err != nil {
			return err
		}
		content = strings.ReplaceAll(content, targetUrl, cloudPath)
	}
	fw, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(content)
	if err != nil {
		panic(err)
	}
	w.Flush()
	return nil
}
