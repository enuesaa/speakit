# speakit

## Future plan
- create examples
  - load dotenv for development
  - rename env vars
- lib
  - sonosctl
  - serve
  - ai
- framework チック
- スピード調整

```go
func main() {
  // signature は変えるかもだが、だいたいこんな感じで feed -> transform -> speaker みたいにデータを流していく
  // statemachine チックかな？
  app := speakit.New(speakit.NewRSSFeedGenerator())
  app.Transform(speakit.NewTransformer()) // text to mp3, voice settings.

  app.Speak(speakit.NewSonosSpeaker())
}
```
