FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT ["CompileDaemon", "--directory=/app",  "--command=/app/go_migration_rnd"]
