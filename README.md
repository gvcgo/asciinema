# WIP
# asciinema V2 
**This is not a complete project** This repo is a fork of the [asciinema](github.com/moqsien/asciinema) repo under the `golang` branch. It has been refactored a bit so that it can be used as a lib. As the originating branch is quite old, this lib will be behind some of the latest features and improvements that has been made to asciinema overall. 

### Implemented
- Record
- Play

The for was made to use with a local project, and all PR's are welcome. 

## Usage
### Install
```sh
go get -u github.com/moqsien/asciinema # for v1 asciinema format
```

```go
package main

import "github.com/moqsien/asciinema"

func main() {
    cli := asciinema.New()
    cast, err := cli.Rec()
    ...
}
```
