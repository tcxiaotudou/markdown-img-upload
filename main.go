package main

import (
	"flag"
	"log"
	"markdown-img-upload/sms"
	"markdown-img-upload/tencent"
	"markdown-img-upload/utils"
)

func main() {
	var source string
	var file string
	flag.StringVar(&source, "source", "sms", "上传源")
	flag.StringVar(&file, "file", "", "文件名")
	//解析命令行参数
	flag.Parse()
	path := utils.GetFilePath(file)
	err := UploadImg(path, source)
	if err != nil {
		log.Fatal("上传失败")
		return
	}
}

func UploadImg(filePath string, source string) error {
	log.Println("path:", filePath)
	switch source {
	case "sms":
		return utils.ModifyImg(filePath, new(sms.SMMS))
	case "tencent":
		return utils.ModifyImg(filePath, new(tencent.Tencent))
	}
	return nil
}
