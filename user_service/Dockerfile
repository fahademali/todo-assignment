FROM golang:1.21

WORKDIR /user_service/cmd/server

COPY . .

RUN go get -u github.com/lib/pq

RUN go build -o main ./user_service/cmd/server

EXPOSE 8080

CMD ["./main"]