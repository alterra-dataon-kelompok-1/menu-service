FROM golang:1.18.2-alpine3.15 AS builder
WORKDIR /menu-service
COPY ./ ./
RUN go mod download
RUN go build -o main

FROM alpine:3.15
WORKDIR /menu-service
COPY --from=builder /menu-service/main .
COPY .env .
EXPOSE 8000
CMD [ "./main" ]