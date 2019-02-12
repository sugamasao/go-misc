# MySQL接続サンプル

Dockerでテスト用DBを立ち上げ、main.goで検証用のコードを動かす

```
% docker-compose up # 別ターミナルで
% dep ensure
% go run main.go
```

sql以下のディレクトリはDocker起動時にseedなどを設定するSQLを実行するためのフックディレクトリ