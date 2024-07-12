[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema)
[![GitHub License](https://img.shields.io/github/license/gvcgo/asciinema?style=for-the-badge)](LICENSE)
[![GitHub Release](https://img.shields.io/github/v/release/gvcgo/asciinema?display_name=tag&style=for-the-badge)](https://github.com/gvcgo/asciinema/releases)
[![PRs Card](https://img.shields.io/badge/PRs-acast-cyan.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema/pulls)
[![Issues Card](https://img.shields.io/badge/Issues-acast-pink.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema/issues)

[中文](https://github.com/gvcgo/asciinema/blob/main/docs/README_CN.md) | [En](https://github.com/gvcgo/asciinema)

------------

- [什么是asciinema?](#什么是asciinema)
- [如何安装](#如何安装)
- [子命令介绍](#子命令介绍)
- [效果演示](#效果演示)
- [感谢以下项目](#感谢以下项目)

------------
## 什么是asciinema?

**asciinema** [as-kee-nuh-muh] 是一个免费开源的终端会话录制和分享工具。
你可以访问 [asciinema.org](https://asciinema.org) 了解更多关于 **asciinema** 的信息。

本项目是 **asciinema** 的跨平台版本，使用 **go** 语言编写，拥有完整的功能，比[官方asciinema工具](https://github.com/asciinema/asciinema)更强大，更方便好用。你可以在MacOS/Linux/Windows上使用它来 **创建、编辑、上传、转换(转换为gif)** 一个 **asciinema**格式的录像。

当你需要写文档，做教程，分享终端相关的操作时，**asciinema**会非常实用。

------------
## 如何安装
- **推荐安装方法**：通过版本管理器[vmr](https://vdocs.vmr.us.kg/zh-cn/)进行安装.

- 通过**go**自带命令进行安装**acast**.
```bash
go install github.com/gvcgo/asciinema/cmd/acast@latest
```

- 从**releases**页面下载后手动解压.
[releases](https://github.com/gvcgo/asciinema/releases)

------------
## 子命令介绍
| subcommand | args example | desc |
|-------|-------|-------|
| **auth** | - | 将本地ID授权到你注册的asciinema.org账户，这样你就可以使用本地ID来上传cast文件到官网了. |
| **convert-to-gif** | input.cast output.gif | 将cast文件转换为gif动图，需要用到[agg](https://github.com/asciinema/agg)，建议使用[vm](https://github.com/gvcgo/version-manager)一键安装agg |
| **cut** | --start=0.0 --end=2.9 input.cast output.cast | 剪切掉cast文件中不需要的时间段，单位是秒. |
| **play** | input.cast | 播放cast文件. |
| **quantize** | --ranges=1.0,5.0 input.cast output.cast | 更新特定区间内的延迟. |
| **record** | xxx.cast | 录制cast文件. |
| **speed** | --start=0.0 --end=2.9 --factor=0.7 input.cast output.cast | 通过一个参数因子，调节某个指定时间区间内的播放速度. |
| **upload** | xxx.cast | 上传cast文件到asciinema.org，需要**auth**授权. |
| **version** | - | 显示acast的版本信息. |

------------
## 效果演示

- **Normal Speed**
[![asciicast](https://asciinema.org/a/651138.svg)](https://asciinema.org/a/651138)
- **Normal Speed Converted to GIF**
![normal](https://github.com/moqsien/img_repo/raw/main/test.gif)

- **Speed x2**
[![asciicast](https://asciinema.org/a/651140.svg)](https://asciinema.org/a/651140)
- **Speed x2 Converted to GIF**
![speedup](https://github.com/moqsien/img_repo/raw/main/test-speedup.gif)

------------

## 感谢以下项目

- [go-asciinema](https://github.com/securisec/asciinema) provided most of the code for unix-like platforms.
- [PowerSession-rs](https://github.com/Watfaq/PowerSession-rs) inspired me the conpty fixes.
- [conpty-go](https://github.com/qsocket/conpty-go)
- [conpty](https://github.com/UserExistsError/conpty)
- [asciinema-edit](https://github.com/cirocosta/asciinema-edit)
- [agg](https://github.com/asciinema/agg)
