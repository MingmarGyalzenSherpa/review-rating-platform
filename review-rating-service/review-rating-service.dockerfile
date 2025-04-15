#build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY reviewRatingApp /app

CMD ["/app/reviewRatingApp"]