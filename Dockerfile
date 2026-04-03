FROM golang:latest AS builder

SHELL ["/bin/bash", "-c"]

# Устанавливаем переменные
ENV GOARCH=amd64
ENV CC=musl-gcc
ENV TARGETOS=linux
ENV ENTRY_POINT=./cmd/entrypoint
ENV PROGRAM=./shortner
ENV WKDIR=/build

# Устанавливаем зависимости
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y musl musl-dev musl-tools

# Рабочая директория
WORKDIR ${WKDIR}
COPY . ${WKDIR}

# Компилируем
RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${GOARCH} go build -o ${PROGRAM} ${ENTRY_POINT}

# Создаем финальный образ
FROM alpine:latest

# Устанавливаем переменные
ENV PROGRAM=shortner
ENV WKDIR=/app
ENV BUILDIR=/build

# Рабочая директория
WORKDIR ${WKDIR}

# Копируем исполняемый файл из предыдущего образа
COPY --from=builder ${BUILDIR}/${PROGRAM} ./${PROGRAM}

# Добавляем сертификаты
RUN apk add --upgrade --no-cache ca-certificates && update-ca-certificates

# Устанавливаем время
RUN apk add tzdata && echo "Europe/Moscow" > /etc/timezone && ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

COPY config/prod.yaml .
COPY internal/database/migrations ./migrations

RUN update-ca-certificates
