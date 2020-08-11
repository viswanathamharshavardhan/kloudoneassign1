FROM golang:1.12.7-alpine3.9
MAINTAINER harshateju
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
