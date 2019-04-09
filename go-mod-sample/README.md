# go mod の動作確認するくん

`go mod`で特殊なパッケージを更新できるか確認するくん

```
% go version
go version go1.12.3 darwin/amd64
% GO111MODULE=on go run main.go
go: finding golang.org/x/text/encoding/japanese latest
go: finding golang.org/x/text/encoding latest
build command-line-arguments: cannot load golang.org/x/text/encoding/japanese: cannot find module providing package golang.org/x/text/encoding/japanese
```
