FROM alpine:3.21 

EXPOSE 80
WORKDIR /service

COPY ./bin/storage-service /service/storage-service

CMD ["/service/storage-service"]
