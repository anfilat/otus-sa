FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000

CMD ["./main"] 
