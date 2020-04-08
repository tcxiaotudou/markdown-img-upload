package utils

import (
	"bufio"
	"io/ioutil"
	upload "markdown-img-upload/interface"
	"os"
	"regexp"
	"strings"
)

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
