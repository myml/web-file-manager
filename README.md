# web-file-manager

## 介绍

简单的在线文件管理器，在浏览器管理文件，并计划支持远程下载，支持桌面和移动端。

## 截图

![截图](./screenshots.png)

## 使用

Use golang

`go get github.com/myml/web-file-manager`

Use docker/podman

`podman run -d -v $(pwd):/data -p 8080:8080 myml/web-file-manager`

## 路线图

- [x] 文件查看
- [x] 文件上传
- [x] 文件下载
- [x] 文件夹创建
- [x] 文件删除
- [x] 文件复制
- [ ] 文件选择
- [x] 文件剪贴
- [ ] 文件属性
- [ ] 远程下载
- [ ] bt 下载
- [x] 图片预览
- [x] pdf 预览
- [ ] word 预览
- [ ] tar.gz 预览
- [ ] ssh 文件管理
- [ ] 国际化支持
- [x] systemd socket activation
