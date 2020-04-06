## 💡 简介

- 自动上传markdown文件中的图片至图床（目前仅支持sm.ms），并替换链接
- 方便markdown文档中图片的处理，例如用于hexo博客编写、笔记等

## 🔮 下载使用

你可以选择直接  [下载](https://github.com/tcxiaotudou/markdown-img-upload/releases/tag/v1.0.0)，然后进行解压

#### 修改配置

1. 修改`config.ini`文件中的`file_path`字段: markdown文档所在目录：

```properties
file_path = "D:\blog\source\_posts\"
```

#### 执行

Mac/Unix：

```bash
./markdown-img-upload -file <filename.md>
```

Windows:

```bash
markdown-img-upload.exe -file <filename.md>
```

## 📽️ 本地运行

#### 拉取代码 ####

```bash
git clone git@github.com:tcxiaotudou/markdown-img-upload.git
```
#### 修改配置
1. 修改`config.ini`文件中的`file_path`字段: markdown文档所在目录：
```properties
file_path = "D:\blog\source\_posts\"
```
#### 下载依赖并编译
```bash
go mod tidy
go build
```
#### 执行
Mac/Unix：
```bash
./markdown-img-upload -file <filename.md>
```
Windows:
```bash
markdown-img-upload.exe -file <filename.md>
```

### 后续

- 接入阿里云，腾讯云，七牛等云存储

### 演示

![example.gif](https://i.loli.net/2020/04/06/Id3thvYFbaWpclB.gif)
