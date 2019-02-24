# YAML sample

YAMLから読み込んだキーを動的に指定したい。
`yaml_data[key]` 的な感じで。

```sh
% dep ensure
% go run main.go
```

### benchmark

```sh
% go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/sugamasao/go-misc/yaml-sample
BenchmarkUseStruct-4              100000             17232 ns/op            7768 B/op        102 allocs/op
BenchmarkUseInterface-4           100000             19326 ns/op            8272 B/op        112 allocs/op
BenchmarkUseString-4              100000             18819 ns/op            8272 B/op        112 allocs/op
PASS
ok      github.com/sugamasao/go-misc/yaml-sample        6.146s
```