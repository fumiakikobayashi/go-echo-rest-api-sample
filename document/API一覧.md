# API一覧

- タスク作成API
  - `POST /tasks`
- タスク一覧取得API
  - `GET /tasks?sort={sort_field}&order={order}`
- タスク削除API
  - `DELETE /tasks/{task_id}`
- タスク編集API
  - `PUT /tasks/{task_id}`
- お気に入り設定API
  - `PATCH /tasks/{task_id}/favorite`
- タスク完了状態設定API
  - `PATCH /tasks/{task_id}/complete`