
FROM golang:1.23

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /web3app-serv
CMD ["/web3app-serv"]