# designdoc

## Steps
### Step 1
- extract rss ... converter
- voicevox のウェブサーバーを立て音声メッセージに変換する

### Step 2
- コンテンツ配信用のウェブサーバーを立て音声メッセージを配信する

### Step 3
- スマホでボタンを押す
- sonosのapiを操作する (Python)
- sonosから音声が流れる

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
