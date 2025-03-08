FROM golang:1.22

WORKDIR /app

COPY . .
RUN go build

CMD ["/app/speakit", "serve", "--redis", "redis:6379", "--voicevox", "voicevox:50021"]
