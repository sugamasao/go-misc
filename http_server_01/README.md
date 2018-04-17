# go-lang HTTPサンプル

## ローカルで試す場合

```
% go run app.go
```

## Dockerで試す場合

```
% GOOS=linux GOARCH=amd64 go build -o bin/app app.go
```

でLiunx向けバイナリを作成する。そのあとdocker-composeで起動させれば良い

```
% docker-compose up
```

## 構成

- 8000ポートでnginx（単一コンテナ）が受け持つ
- 8080ポートでWebアプリケーション（単一コンテナ）が起動する