# Asciinema V2 
This repo is a fork of the [asciinema](https://github.com/asciinema/asciinema) repo under the `golang` branch.

## Supported Platforms
- Unix-like: Shell
- Windows: PowerShell

## Implemented
- Record
- Play
- Auth
- Upload

The fork was made to use with a local project, and all PR's are welcome. 

## Maybe, the only thing you need is **gvc**.
[gvc](https://github.com/moqsien/gvc) is a powerful command-line tool with asciinema features implemented.

## Usage
### Install
```sh
go get -u github.com/moqsien/asciinema
```

### Import
```go
package main

import (
	"os"

	"github.com/moqsien/asciinema/cmd"
)

func main() {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	cli := cmd.New(args...)
	// fmt.Println(cli.Auth())
	// cli.Rec()
	cli.Play()
	// info, _ := cli.Upload()
	// fmt.Println(info)
}
```
## Thanks To
- [go-asciinema](https://github.com/securisec/asciinema) provided most of the code for unix-like platforms.
- [PowerSession-rs](https://github.com/Watfaq/PowerSession-rs) inspired me the conpty fixes.
- [conpty-go](https://github.com/qsocket/conpty-go)
- [conpty](https://github.com/UserExistsError/conpty)
