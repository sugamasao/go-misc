# WebDriverを使ったブラウザの操作

chromedriverを使うので、brewでインストールする

```sh
% brew cask install chromedriver
```

# go-langのバージョンについて

- 1.11以上
- 外部ライブラリ管理には公式のdebを使っている

Docker用にビルドする場合は

```sh
% GOOS=linux GOARCH=amd64 go build main.go
```

# Dockerについて

これを参考に使って見る

https://hub.docker.com/r/robcherry/docker-chromedriver/

```
% docker build -t sample:latest ./
```

ビルドして、IDを使ってexecでコマンドを叩く

```sh
% docker exec -it 0b33 /bin/bash
```
