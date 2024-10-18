FROM golang:1.22 as go-base

ENV PROJECT_DIR=/app \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . ./

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go mod download
RUN go build -o /app/http-init /app/http-init.go

RUN chmod +x /app/http-init

EXPOSE 8000

USER ${OS_USER}:${OS_USER}

FROM go-base as go-dev

ENV GO111MODULE=on \
    CGO_ENABLED=0

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o /app/http-init" -command="/app/http-init"