[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema)
[![GitHub License](https://github.com/gvcgo/asciinema?style=for-the-badge)](LICENSE)
[![GitHub Release](https://github.com/gvcgo/asciinema?display_name=tag&style=for-the-badge)](https://github.com/gvcgo/asciinema/releases)
[![PRs Card](https://img.shields.io/badge/PRs-vm-cyan.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema/pulls)
[![Issues Card](https://img.shields.io/badge/Issues-vm-pink.svg?style=for-the-badge)](https://github.com/gvcgo/asciinema/issues)

[中文](https://github.com/gvcgo/asciinema/blob/main/docs/README_CN.md) | [En](https://github.com/gvcgo/asciinema)

------------
## What is asciinema?

**asciinema** [as-kee-nuh-muh] is a free and open source solution for recording terminal sessions and sharing them on the web.
To learn about **asciinema**, you can visit [asciinema.org](https://asciinema.org).

And this project is a **cross-platform** version of **asciinema** writtern in go with full features. You can use it to **create, edit, upload, convert(to gif animation)** an asciinema cast on **MacOS/Linux/Windows**. 

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

## Thanks To

- [go-asciinema](https://github.com/securisec/asciinema) provided most of the code for unix-like platforms.
- [PowerSession-rs](https://github.com/Watfaq/PowerSession-rs) inspired me the conpty fixes.
- [conpty-go](https://github.com/qsocket/conpty-go)
- [conpty](https://github.com/UserExistsError/conpty)
- [asciinema-edit](https://github.com/cirocosta/asciinema-edit)
