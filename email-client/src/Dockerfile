FROM golang:alpine

WORKDIR /server

RUN  go mod init email-client

COPY ./ ./

RUN go build

CMD ["go", "run", "."]

EXPOSE 8080