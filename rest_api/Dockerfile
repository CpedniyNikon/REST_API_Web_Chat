FROM golang:1.19.9

WORKDIR /app

COPY . .


RUN go mod download


RUN go env -w GO111MODULE=on
RUN go build -o /rest_api

EXPOSE 8080

CMD [ "/rest_api" ]