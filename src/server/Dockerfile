FROM golang:alpine

WORKDIR /server

RUN  go mod init server

RUN go get github.com/gin-gonic/gin

RUN go get go.mongodb.org/mongo-driver/mongo 

RUN go get github.com/joho/godotenv 

RUN go get github.com/golang-jwt/jwt

RUN go get github.com/go-playground/validator/v10

RUN go get github.com/segmentio/kafka-go

COPY ./ ./

RUN go build

CMD ["go", "run", "."]

EXPOSE 8080