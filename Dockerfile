FROM golang:1.21

WORKDIR /workspace
COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go install -v golang.org/x/tools/gopls@latest

RUN go get github.com/getkin/kin-openapi/openapi3
RUN go get github.com/getkin/kin-openapi/openapi3gen
RUN go get github.com/go-yaml/yaml
RUN go get github.com/iancoleman/strcase

CMD ["air", "-c", ".air.toml"]
