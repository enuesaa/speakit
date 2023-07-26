# designdoc

## Apps
### Containers
- api (golang)
- voicevox api
- web (nextjs) ... 作り込まない

### API Gateway
managerコンテナが、API Gateway (Proxy)的な役割も受け持つ
- POST /api/load ... fetch rss feed and request to convert
- GET /api/jobs
- POST /api/speak
- POST /api/next
- POST /api/pause
- GET /contents/{id}.wav ... wav file

## Memo
- 単に rss を一つずつ流すのではなく、文脈でグルーピングしたい
- Macでサーバーを立て、Logitech Spotlight と組み合わせる
