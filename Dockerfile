FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.* .

RUN go mod tidy

COPY ./ ./

RUN go build -o api-devbook

FROM alpine as binary

WORKDIR /app

COPY --from=builder /app/api-devbook .

RUN apk add --no-cache tzdata && \
    chmod +x ./api-devbook

CMD ["./api-devbook"]
