# Go+Gin+Gorm+Docker+DDDのREST APIサンプル
こちらはタイトル通りGo+Gin+Gorm+Docker+DDDのREST APIサンプルです。Golangの学習にあたり、アウトプットを目的に作成しました。少しづつ追加実装していく予定です。

## スタートガイド
※ 事前にDocker Desktop for MacやDocker Desktop for Windowsなどでご自身のPC環境にDockerのインストールは済ませておいてください

はじめにリポジトリをクローンして、対象のディレクトリに移動します
```
$ git clone https://github.com/fumiakikobayashi/go-ddd-rest-api-sample.git
$ cd go-learning-environment
```
下記コマンドを実行してコンテナを立ち上げてください。（`make`コマンドを使用することで、`Makefile`に定義しているコマンドを実行することができます。）

※ このコマンドは初回のみ実行します
```
$ make build
$ make up
```
続いて`make ps` コマンドを実行し、コンテナが`runnning`になっていることを確認してください
```
$ make ps
NAME                            COMMAND             SERVICE             STATUS              PORTS
go-learning-environment-app-1   "/bin/sh"           app                 running             

```
これで環境の構築は完了です。最後に実際にgolangを実行できるかどうか確認します。


下記コマンドを実行して、コンテナに入ります。
```
$ make app
```
`go run main.go`で`main.go`を実行します。下記のように現在時刻が表示されれば成功です。
```
$ go run main.go
Welcome to the playground!
The time is 2023-01-23 12:20:52.048513751 +0000 UTC m=+0.000050001
```

## Tips
- 各種コマンドはMakefileをご覧ください [Makefile](https://github.com/PicoCELA/onpremis-api/blob/main/Makefile)# go-ddd-rest-api-sample
