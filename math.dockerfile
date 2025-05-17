FROM alpine:3.21

EXPOSE 80
WORKDIR /service

COPY ./bin/math-service /service/math-service

CMD ["/service/math-service"]