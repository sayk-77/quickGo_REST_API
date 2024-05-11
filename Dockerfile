FROM golang:latest
WORKDIR /QuickGo_backend

COPY . .

RUN go build -o main .

CMD ["./main"]