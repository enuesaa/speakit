FROM golang:1.21

WORKDIR /workspace
COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go install golang.org/x/tools/gopls@latest

CMD ["air", "-c", ".air.toml"]
