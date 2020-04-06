package main

import (
	"flag"
	"gopkg.in/ini.v1"
	sms "hexo-img-upload/upload"
	"hexo-img-upload/utils"
	"log"
)

const (
	FILE_PATH = "config.ini"
)

type Config struct {
	filePath string `ini:"file_path"`
}

func main() {
	var source string
	var file string
	flag.StringVar(&source, "source", "sms", "上传源")
	flag.StringVar(&file, "file", "", "文件名")
	//解析命令行参数
	flag.Parse()
	path := GetFilePath(file)
	UploadImg(path, source)
}

func GetFilePath(fileName string) string {
	filePath, err := ReadConfig(FILE_PATH) //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	return filePath + fileName
}

//读取配置文件
func ReadConfig(path string) (string, error) {
	//var config Config
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return "", err
	}
	filePath := conf.Section("").Key("file_path").String()
	return filePath, nil
}

func UploadImg(filePath string, source string) error {
	log.Println("path:", filePath)
	switch source {
	case "sms":
		utils.ModifyImg(filePath, new(sms.SMMS))
	}
	return nil
}
