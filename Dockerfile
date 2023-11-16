# Compile stage
FROM golang:1.21 AS build-env

ADD . /algorithm
WORKDIR /algorithm

RUN go build -o /server

# Final stage
FROM ubuntu:latest
LABEL authors="Crane"

EXPOSE 8000

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]