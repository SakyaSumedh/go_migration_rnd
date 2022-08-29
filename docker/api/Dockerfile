FROM golang:1.19-alpine

RUN apk update && apk add --no-cache --virtual .build-deps curl git python3

RUN addgroup rnd && adduser -D rnd_user -G rnd -h /home/rnd_user

ENV HOME /home/rnd_user

ENV APP_DIR ${HOME}/project

WORKDIR ${APP_DIR}

COPY --chown=rnd_user:rnd ./ ${APP_DIR}

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENV PATH="${APP_DIR}:${PATH}"

ENTRYPOINT ["CompileDaemon", "-color=true", "-build=go build -buildvcs=false .",  "-command=./go_migration_rnd"]
