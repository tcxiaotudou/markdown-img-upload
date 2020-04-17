## 💡 简介

- 自动上传markdown文件中的图片至图床，并替换链接
- 方便markdown文档中图片的处理，例如用于hexo博客编写、笔记等
- 目前支持smms（默认），腾讯云OSS

## 🔮 下载使用

你可以选择直接  [下载](https://github.com/tcxiaotudou/markdown-img-upload/releases/tag/v1.0.0)，然后进行解压

#### 修改配置

将压缩包内的文件解压到同一目录，修改`config.ini`文件中的`file_path`字段: markdown文档所在目录：

```properties
file_path = "D:\blog\source\_posts\"
```

如需接入`腾讯云OSS`等，也请配置对应参数，否则无法使用

#### 执行

Mac/Unix：

```bash
./markdown-img-upload -source smms -file <filename.md>
```

Windows:

```bash
markdown-img-upload.exe -source smms -file <filename.md>
```

**注：**

`-source`：可选参数，指定上传源，目前支持 `smms`，`腾讯云oss`，不指定时，默认为`smms`

- `smms` 
- `tencent` 

`-file`：必选参数，指定markdown文件名，例：`测试.md`

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
./markdown-img-upload -source smms -file <filename.md>
```
Windows:
```bash
markdown-img-upload.exe -source smms -file <filename.md>
```

**注：**

`-source`：可选参数，指定上传源，目前支持 `smms`，`腾讯云oss`，不指定时，默认为`smms`

- `smms` 
- `tencent` 

`-file`：必选参数，指定markdown文件名，例：`测试.md`

### 后续

- [x] 接入腾讯云OSS 
- [ ] 接入阿里云OSS
- [ ] 接入七牛云OSS 

### 演示

![example.gif](https://i.loli.net/2020/04/06/Id3thvYFbaWpclB.gif)
