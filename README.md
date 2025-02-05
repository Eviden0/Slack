<h4 align="center">一款安服集成化工具平台，希望能让你少开几个应用测试</h4>

<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/qiwentaidi/Slack?filename=go.mod">
<img src="https://img.shields.io/badge/wails-v2.9.1-blue">
<a href="https://github.com/qiwentaidi/Slack/releases/"><img src="https://img.shields.io/github/v/release/qiwentaidi/Slack"></a>
<a href="https://github.com/qiwentaidi/Slack/releases/"><img src="https://img.shields.io/github/downloads/qiwentaidi/Slack/total"></a>
</p>






# 运行代码

## 安装环境

Linux/Debian

```sh
# need install Go 1.20+
sudo apt install golang-go

# need install nodejs 15+
sudo apt install nodejs npm

# install gcc && libpcap
sudo apt install build-essential libgtk-3-dev libpcap-dev 

# if apt can't be found 4.0-dev then install 4.1-dev
1、sudo apt install libwebkit2gtk-4.0-dev
2、sudo apt install libwebkit2gtk-4.1-dev

# install wails to run app
go install github.com/wailsapp/wails/v2/cmd/wails@latest

git clone https://github.com/qiwentaidi/Slack.git
```

Windows

```sh
# need install Go 1.20+ && nodejs 15+
# download gcc(https://github.com/niXman/mingw-builds-binaries/releases) and configure environment variables

# install wails to run app
go install github.com/wailsapp/wails/v2/cmd/wails@latest

git clone https://github.com/qiwentaidi/Slack.git
```

## 编译/调试

``````sh
cd Slack
# used for debugging while writing code
wails dev 

# build executable file on /build/bin
wails build

# open web console in executable file
wails build -debug -devtools
``````

`Mac build dmg`

``````sh
# need build Slack.app first

brew install create-dmg

create-dmg --volname "Slack" --window-pos 200 120 --window-size 800 400 --icon-size 100  --icon "Slack.app" 200 190 --app-drop-link 600 185 --hide-extension "Slack.app" --volicon build/bin/Slack.app/Contents/Resources/iconfile.icns  "Slack.dmg" build/bin/Slack.app
``````

# 模块介绍

## 首页

![image-20240510214011206](assets/image-20240530185056848.png)

![image-20240510214011206](assets/image-20240530185121324.png)

## 渗透测试

### 网站扫描

![image-20240510214011206](assets/image-20240510214011206.png)

![image-20240510214144484](assets/image-20240510214144484.png)

### 主机扫描

![image-20240628171638340](assets/image-20240628171638340.png)

### 暴破与未授权检测

![image-20240628173000629](assets/image-20240628173000629.png)

### 目录扫描

![image-20240410135353422](assets/image-20240410135353422.png)

### JSFinder

![image-20240425150345629](assets/image-20240425150345629.png)

## 资产收集

### 公司名称查资产

![image-20240718104209501](assets/image-20240718104209501.png)

![image-20240412150938371](assets/image-20240412150938371.png)

### 子域名暴破

![image-20240222154726647](assets/image-20240222154726647.png)

### 域名信息查询

![image-20240629003953025](assets/image-20240629003953025.png)

## 空间引擎

### FOFA

![image-20240718103758587](assets/image-20240718103758587.png)

### 鹰图

![image-20240620234045723](assets/image-20240620234045723.png)

### Quake

![image-20240628172913333](assets/image-20240628172913333.png)

## 小工具

### 加解密模块

![image-20240628172153383](assets/image-20240628172153383.png)

### 数据处理

![image-20240628172250510](assets/image-20240628172250510.png)

### 杀软/补丁识别

![image-20240628172737827](assets/image-20240628172737827.png)

### 反弹Shell生成

![image-20240628172826365](assets/image-20240628172826365.png)

### 备忘录

![image-20240628172702349](assets/image-20240628172702349.png)

### WeChat/DingDing

![image-20240512003433659](assets/image-20240512003433659.png)

# 常见问题

## Windows

> Q：grdp库无法打包，出现如下报错
>
> ```
> Wails CLI v2.7.1
> 
> Executing: go mod tidy
> • Generating bindings: 
> ERROR  
> package slack-wails
>  imports slack-wails/core/portscan
>  imports github.com/tomatome/grdp/protocol/pdu
>  imports github.com/tomatome/grdp/protocol/t125/gcc
>  imports github.com/tomatome/grdp/plugin: build constraints exclude all Go files in C:\xx\go\pkg\mod\github.com\tomatome\grdp@v0.1.0\plugin
> ```
>
> A：检查GCC环境环境是否安装，可以通过命令行输入gcc查看命令提示
>
> ```
> 1、查看文档上方安装教程
> 2、go env查看CGO_ENABLED是否为1，若不是则go env -w CGO_ENABLED=1 
> ```
## Linux
>
>Q1: gopacket打包失败
>
>``````sh
># Building target: linux/amd64
>
>• Generating bindings:   ERROR   
># github.com/google/gopacket/pcap
>/root/go/pkg/mod/github.com/google/gopacket@v1.1.19/pcap/pcap_unix.go:34:10: fatal error: pcap.h: No such file or directory                     
>   34 | #include <pcap.h>
>      |          ^~~~~~~~
>compilation terminated.
>
>exit status 1
>
>ERROR   
># github.com/google/gopacket/pcap
>/root/go/pkg/mod/github.com/google/gopacket@v1.1.19/pcap/pcap_unix.go:34:10: fatal error: pcap.h: No such file or directory                     
>   34 | #include <pcap.h>
>      |          ^~~~~~~~
>compilation terminated.
>
>exit status 1
>♥   If Wails is useful to you or your company, please consider sponsoring the project:                                                                   
>https://github.com/sponsors/leaanthony
>``````
>
>A: 安装libpcap-dev包
>
>``````sh
>sudo apt install libpcap-dev 
>``````
>
>
>
>Q2: nodejs编译内存溢出
>
>``````sh
><--- Last few GCs --->
>    
>    [7287:0x55767da1ede0]     9749 ms: Mark-sweep (reduce) 979.8 (1002.0) -> 979.4 (999.5) MB, 256.4 / 0.0 ms  (+ 11.1 ms in 5 steps since start of marking, biggest step 5.5 ms, walltime since start of marking 272 ms) (average mu = 0.190, current mu = 0.155) 
>    
>    <--- JS stacktrace --->
>    
>    FATAL ERROR: Ineffective mark-compacts near heap limit Allocation failed - JavaScript heap out of memory
>     1: 0x7ff1401fbdd8 node::Abort() [/lib/x86_64-linux-gnu/libnode.so.109]
>     2: 0x7ff1400fe362  [/lib/x86_64-linux-gnu/libnode.so.109]
>     3: 0x7ff1405d9f80 v8::Utils::ReportOOMFailure(v8::internal::Isolate*, char const*, bool) [/lib/x86_64-linux-gnu/libnode.so.109]
>     4: 0x7ff1405da33c v8::internal::V8::FatalProcessOutOfMemory(v8::internal::Isolate*, char const*, bool) [/lib/x86_64-linux-gnu/libnode.so.109]
>     5: 0x7ff1407adbb5  [/lib/x86_64-linux-gnu/libnode.so.109]
>     6: 0x7ff1407adc8b  [/lib/x86_64-linux-gnu/libnode.so.109]
>     7: 0x7ff1407c1a5a v8::internal::Heap::PerformGarbageCollection(v8::internal::GarbageCollector, v8::internal::GarbageCollectionReason, char const*, v8::GCCallbackFlags) [/lib/x86_64-linux-gnu/libnode.so.109]
>     8: 0x7ff1407c25c9 v8::internal::Heap::CollectGarbage(v8::internal::AllocationSpace, v8::internal::GarbageCollectionReason, v8::GCCallbackFlags) [/lib/x86_64-linux-gnu/libnode.so.109]
>     9: 0x7ff14079ee4a v8::internal::HeapAllocator::AllocateRawWithLightRetrySlowPath(int, v8::internal::AllocationType, v8::internal::AllocationOrigin, v8::internal::AllocationAlignment) [/lib/x86_64-linux-gnu/libnode.so.109]
>    10: 0x7ff14079ff55 v8::internal::HeapAllocator::AllocateRawWithRetryOrFailSlowPath(int, v8::internal::AllocationType, v8::internal::AllocationOrigin, v8::internal::AllocationAlignment) [/lib/x86_64-linux-gnu/libnode.so.109]
>    11: 0x7ff1407810ce v8::internal::Factory::NewFillerObject(int, v8::internal::AllocationAlignment, v8::internal::AllocationType, v8::internal::AllocationOrigin) [/lib/x86_64-linux-gnu/libnode.so.109]
>    12: 0x7ff140b5ad2c v8::internal::Runtime_AllocateInYoungGeneration(int, unsigned long*, v8::internal::Isolate*) [/lib/x86_64-linux-gnu/libnode.so.109]
>    13: 0x7ff1404e37b9  [/lib/x86_64-linux-gnu/libnode.so.109]
>    Aborted
>``````
>
>A: 设置nodejs最大内存
>
>``````sh
>export NODE_OPTIONS="--max_old_space_size=4096"
>``````
>
>
>
>Q3: 新版Linux没有安装未找到 libwebkit2gtk-4.0-dev
>
>A: 安装libwebkit2gtk-4.1-dev 并安装Wails CLI 2.9.0+ ，通过设置编译tags支持
>
>``````sh
>apt install libwebkit2gtk-4.1-dev
>wails build -tags webkit2_41
>``````
## MacOS

>
> Q: Mac 启动出现安装包损坏
>
> A: 在隐私和安全性处运行安装应用为任意来源
>
> ``````sh
> xattr -c slack-macos.app
> ``````

# 联系方式

如果有问题或者好的提议可以Issue提问或者加我联系方式（请备注来意 进群或者问题交流）

![image-20231006124944803](assets/image-20231006124944803.png)

# 免责声明

本工具仅面向**合法授权**的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

为避免被恶意使用，本项目所有收录的poc均为漏洞的理论判断，不存在漏洞利用过程，不会对目标发起真实攻击和漏洞利用。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。**请勿对非授权目标进行扫描。**

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。

在安装并使用本工具前，请您**务必审慎阅读、充分理解各条款内容**，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要安装并使用本工具。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。
