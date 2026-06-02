FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM alpine:3.22.4
RUN apk add --no-cache ca-certificates
COPY --from=builder /server /server
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/server"]
