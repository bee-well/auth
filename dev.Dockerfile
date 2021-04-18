FROM golang:alpine

WORKDIR app

COPY . .

RUN go build

CMD ["./auth"]