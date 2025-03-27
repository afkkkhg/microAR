# Компиляция приложения в контейнере, если нет Golang'а
FROM golang:1.23.7-alpine3.21 AS builder
WORKDIR /appsource
COPY ./main.go ./
COPY ./group-handler.go ./
COPY ./contact-handlers.go ./
COPY ./go.mod ./

RUN go mod tidy
RUN go build -o app .

# Само приложение
FROM alpine:3.21
WORKDIR /myapp
COPY --from=builder /appsource/app ./
EXPOSE 6080
ENTRYPOINT ["/myapp/app"]