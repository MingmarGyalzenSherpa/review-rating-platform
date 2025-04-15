#build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY productApp /app

CMD ["/app/productApp"]