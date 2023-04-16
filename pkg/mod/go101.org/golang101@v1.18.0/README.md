**[微信群召集帖](https://github.com/golang101/golang101/issues/11)** | **[电子书下载](https://github.com/golang101/golang101/releases)**

----

[<b>Go语言101</b>](https://gfw.go101.org)是关于Go语言编程的系列丛书。
目前本系列丛书包括：

* 


本系列丛书同时适合Go初学者和有一定经验的Go程序员阅读，
希望它们能够帮助Go程序员更深更全面地理解Go语言。

本丛书的英文项目地址为：[github.com/go101/go101](https://github.com/go101/go101)

微信公众号（用于发布一些Go细节、常识和技巧）：

![](not-embedded/website/res/101-group-qrcode-2.jpg?raw=true)

_(若上面二维码未显示出来，请点击此[墙内版链接](https://tool.oschina.net/action/qrcode/generate?data=http%3A%2F%2Fweixin.qq.com%2Fr%2FRy6ju1TE0AmvrRDY93tV&output=image%2Fgif&error=L&type=0&margin=12&size=4)或者在微信中搜索"Go 101"公众号。)_

### 安装、更新以及本地阅读本系列丛书

如果你使用官方Go工具链v1.16+，则不需克隆本项目代码：

```shell
### 安装和更新

$ go install go101.org/golang101@latest

### 本地阅读（GOBIN路径需配置在PATH中。GOBIN路径的默认值为GOPATH/bin。）

$ golang101
Server started:
   http://localhost:12345 (non-cached version)
   http://127.0.0.1:12345 (cached version)
```

如果你使用官方Go工具链v1.15-或者欲做一些本地修改（比如准备提交PR等）：
```shell
### 安装

$ git clone https://github.com/golang101/golang101.git

### 更新. 进入本书项目目录（包含当前`README.md`文件的目录），然后运行：

$ git pull

### 本地阅读系列丛书. 进入本书项目目录，然后运行：

$ go run .
Server started:
   http://localhost:12345 (non-cached version)
   http://127.0.0.1:12345 (cached version)
```

本系列丛书起始页将自动在用户默认浏览器中打开。如果没有，请手动访问http://localhost:12345

命令行选项：
```
-port=1234
-theme=light # 或者 dark （默认）
```

### 一些注意事项

* 如果在线版被墙了，请运行如上所述本地版或者下载[离线版](https://github.com/golang101/golang101/releases)阅读。
* 本人保留本系列丛书的出版权（包括纸质和各种电子版）。请勿随意在线转载（转载之前请在本项目中开一个issue问询）。
* 本系列丛书仍在不断改进中。请关注本系列丛书微信（二维码见上）和[博客](https://gfw.go101.org/blog/101.html)以注意各种更新。

![](pages/fundamentals/res/101-reward-qrcode-5.png?raw=true)
