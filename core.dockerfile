FROM alpine:3.21

EXPOSE 80
WORKDIR /service

COPY ./bin/core-service /service/core-service

CMD ["/service/core-service"]