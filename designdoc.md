# designdoc

## Apps
### Containers
- manager (golang)
- voicevox
- speaker (python using soco)

### API Gateway
managerコンテナが、API Gateway (Proxy)的な役割も受け持つ
- POST /api/load ... fetch rss feed and request to convert
- GET /api/jobs
- POST /api/speak ... python sonos api
- POST /api/next ... python sonos api
- POST /api/pause ... python sonos api
- GET /contents/{id}.wav ... wav file

## Memo
- 単に rss を一つずつ流すのではなく、文脈でグルーピングしたい
