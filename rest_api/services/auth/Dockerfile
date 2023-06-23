FROM golang:1.16-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./



RUN export GO111MODULE=on
RUN go get

RUN go build -o /server

EXPOSE 8080

CMD ["server"]
ENTRYPOINT ["/server"]