# speakit
## コンセプト
- RSSの情報を読み上げたい
- RSSの先の情報（Webページ）も読み取って欲しい
- で、要約とかもして欲しい

## Memo
音声を流すには
### Voicevox
- https://github.com/VOICEVOX/voicevox_engine
- https://qiita.com/mikito/items/21aa74c3850a70c647f7

see https://github.com/VOICEVOX/voicevox_engine
```bash
echo -n "こんにちは" > text.txt

curl -s -X POST "127.0.0.1:50021/audio_query?speaker=1" --get --data-urlencode text@text.txt > query.json
curl -s -H "Content-Type: application/json" -X POST -d @query.json "127.0.0.1:50021/synthesis?speaker=1" > audio.wav
```
### GCP Speech To Text
- https://zenn.dev/tatsuyasusukida/articles/gcp-text-to-speech

## Development Plan
- sonosのapiを使う
- rssや音声への変換は事前にしておけば良い
- コンテンツ配信にウェブサーバーを立てる
