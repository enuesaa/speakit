# designdoc
## URL体系
APIが主軸で、
管理画面もある、と言う形にしたい

- GET /feeds ... feed list
- GET /feeds/{id}
- POST /feeds
- DELETE /feeds/{id}

- POST /jobs ... fetch rss feed and request to convert. 202 を返したい
- GET /jobs
- GET /jobs/{id}

- GET /contents ... 一覧
- GET /contents/{id} ... asset id を返す

- GET /assets/{id}  ... wav file

## 画面
- GET /admin
- GET /admin/feeds
- GET /admin/player
  - start
  - next
  - prev
  - stop

## Development Plan
- 単に rss を一つずつ流すのではなく、文脈でグルーピングしたい
