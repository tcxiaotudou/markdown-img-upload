# markdown-img-upload
- 自动上传markdown文件中的图片至图床，并替换链接

- 方便markdown文档中图片的处理，例如用于hexo博客编写、笔记等

## 使用
### 下载
```bash
git clone git@github.com:tcxiaotudou/markdown-img-upload.git
```
### 修改配置
1. 修改`config.ini`文件中的`file_path`字段: markdown文档所在目录：
```shell script
file_path = "D:\blog\source\_posts\"
```
### 安装依赖并编译
```bash
go mod tidy
go build
```
### 执行
Mac/Unix：
```bash
./markdown-img-upload
```
Windows:
```bash
markdown-img-upload.exe
```
### 演示
![example.gif](https://i.loli.net/2020/04/06/Id3thvYFbaWpclB.gif)
