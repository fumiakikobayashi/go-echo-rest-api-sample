# Go+echo+Gorm+Docker+DDDのREST APIサンプル
こちらはタイトル通りGo+echo+Gorm+Docker+DDD(オニオンアーキテクチャ)のREST APIサンプルです。Golangの学習にあたり、アウトプットを目的に作成しました。少しづつ追加実装していく予定です。

## スタートガイド
※ 事前にDocker Desktop for MacやDocker Desktop for Windowsなどでご自身のPC環境にDockerのインストールは済ませておいてください

はじめにリポジトリをクローンして、対象のディレクトリに移動します
```
$ git clone https://github.com/fumiakikobayashi/go-echo-rest-api-sample.git
$ cd go-echo-rest-api-sample
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
NAME                           COMMAND                  SERVICE             STATUS              PORTS
go-echo-rest-api-sample-app-1   "go run ."               app                 running             0.0.0.0:8080->8080/tcp
go-echo-rest-api-sample-db-1    "docker-entrypoint.s…"   db                  running             0.0.0.0:3306->3306/tcp
go-echo-rest-api-sample-pma-1   "/docker-entrypoint.…"   pma                 running             0.0.0.0:8088->80/tcp
```
これで環境の構築は完了です。

## API一覧
- タスク作成API
  - `POST /tasks`
- タスク一覧取得API
  - `GET /tasks?sort={sort_field}&order={order}`
    - sort：name, deadline, favorite
    - order：asc、desc
- タスク削除API
  - `DELETE /tasks/{task_id}`
- タスク編集API
  - `PUT /tasks/{task_id}`
- お気に入り設定API
  - `PATCH /tasks/{task_id}/favorite`
- タスク完了状態設定API
  - `PATCH /tasks/{task_id}/complete`
- 提案タスク取得API
  - `GET /tasks/suggestion`

## Tips
- 各種コマンドはMakefileをご覧ください [Makefile](https://github.com/fumiakikobayashi/customer-management-sample/blob/main/Makefile)
