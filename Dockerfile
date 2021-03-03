FROM golang:1.16-alpine as builder

WORKDIR /app

COPY . .
RUN go mod download && go build -v -o server

FROM alpine

WORKDIR /app

COPY quizzs.json .
COPY --from=builder /app/server .

EXPOSE 8000

CMD ["./server"]