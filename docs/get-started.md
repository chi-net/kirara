# 开始使用

你可以通过以下方式安装kirara至你的服务器

## 安装
### 1. 通过容器服务(推荐)

在确保安装了Podman(推荐)直接执行以下命令即可

```sh
podman run -dit -p 3888:8080 -m 128M ghcr.io/chi-net/kirara
```

拉取运行后在3888端口暴露，如果你的3888端口被占用也可以更改成其他的端口。

当然，你用Docker也不是不行，但是[我不推荐](https://blog.chihuo2104.dev/posts/comment-fixed-and-hi-podman)

```sh
docker run -dit -p 3888:8080 -m 128M ghcr.io/chi-net/kirara
```

### 2. 通过下载执行可执行文件

在[Releases](https://github.com/chi-net/kirara/releases/)下载适合您服务器的可执行文件(请注意您的服务器系统和CPU架构)

::: info 
我们在`Releases`中提供`x86`和`i386`版本的二进制分发包，如果您的服务器为特殊架构，您可以前往[这里](https://repo.chinet.work/chinet-portal/kirara-portbuilds/branch/master/)下载适合你服务器的架构安装包
请注意，我们目前只提供**正式版本**的`arm64(built by Rockchip RK3399)`安装包可供下载，如果您亲爱的小创车不支持我们编译的架构的话，您也可以从源码编译。
:::

随后直接执行可执行文件，输入`http://[Your Server IP]:8080/`即可

### 3. 通过从源码编译至可执行文件

::: info 
请确保你已经安装了`golang 1.18+`和`git`，否则有无法编译成功的风险
:::

1. 如果你位于中国大陆地区，建议先[改源](https://goproxy.cn)
2. 拉取本库并使用go编译
```sh
git pull https://github.com/chi-net/kirara.git
cd kirara/backend
go build
```
3. 此时工作目录会出现`kirara`或`kirara.exe`，这时候就大功告成了！