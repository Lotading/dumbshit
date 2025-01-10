FROM golang:1.23.4-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app . 
EXPOSE 8000
CMD [ "./app" ]
