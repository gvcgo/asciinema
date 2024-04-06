[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema)
[![GitHub License](https://img.shields.io/github/license/gvcgo/asciinema?style=for-the-badge)](LICENSE)
[![GitHub Release](https://img.shields.io/github/v/release/gvcgo/asciinema?display_name=tag&style=for-the-badge)](https://github.com/gvcgo/asciinema/releases)
[![PRs Card](https://img.shields.io/badge/PRs-vm-cyan.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema/pulls)
[![Issues Card](https://img.shields.io/badge/Issues-vm-pink.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema/issues)

[中文](https://github.com/gvcgo/asciinema/blob/main/docs/README_CN.md) | [En](https://github.com/gvcgo/asciinema)

------------
- [What is asciinema?](#what-is-asciinema)
- [Installation](#installation)
- [Subcommands](#subcommands)
- [Demo](#demo)
- [Thanks To](#thanks-to)

------------
## What is asciinema?

**asciinema** [as-kee-nuh-muh] is a free and open source solution for recording terminal sessions and sharing them on the web.
To learn about **asciinema**, you can visit [asciinema.org](https://asciinema.org).

And this project is a **cross-platform** version of **asciinema** writtern in go with full features. You can use it to **create, edit, upload, convert(to gif animation)** an asciinema cast on **MacOS/Linux/Windows**. 

------------
## Installation
- **Recommanded**: Install **acast** using **version manager** [vm](https://github.com/gvcgo/version-manager).
```bash
vm use asciinema@v03.9
```

- Install **acast** using **go**.
```bash
go install github.com/gvcgo/asciinema/cmd/acast@latest
```

- Download **acast** from **releases**.
[releases](https://github.com/gvcgo/asciinema/releases)

------------
## Subcommands
| subcommand | args example | desc |
|-------|-------|-------|
| **auth** | - | Authorizes to your asciinema.org account. |
| **convert-to-gif** | input.cast output.gif | Converts a cast to gif animation. |
| **cut** | --start=0.0 --end=2.9 input.cast output.cast | Removes a certain range of a cast. |
| **play** | input.cast | Plays a cast. |
| **quantize** | --ranges=1.0,5.0 input.cast output.cast | Updates the cast delays following quantization ranges. |
| **record** | xxx.cast | Starts recording a cast. |
| **speed** | --start=0.0 --end=2.9 --factor=0.7 input.cast output.cast | Updates the speed of a cast by certain factor. |
| **upload** | xxx.cast | Uploads a cast to asciinema.org. |
| **version** | - | Shows version info of acast. |

------------
## Demo

- **Normal Speed**
[![asciicast](https://asciinema.org/a/651138.svg)](https://asciinema.org/a/651138)
- **Normal Speed Converted to GIF**
![normal](https://github.com/moqsien/img_repo/raw/main/test.gif)

- **Speed x2**
[![asciicast](https://asciinema.org/a/651140.svg)](https://asciinema.org/a/651140)
- **Speed x2 Converted to GIF**
![speedup](https://github.com/moqsien/img_repo/raw/main/test-speedup.gif)

------------

## Thanks To

- [go-asciinema](https://github.com/securisec/asciinema) provided most of the code for unix-like platforms.
- [PowerSession-rs](https://github.com/Watfaq/PowerSession-rs) inspired me the conpty fixes.
- [conpty-go](https://github.com/qsocket/conpty-go)
- [conpty](https://github.com/UserExistsError/conpty)
- [asciinema-edit](https://github.com/cirocosta/asciinema-edit)
- [agg](https://github.com/asciinema/agg)
