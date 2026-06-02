# go-echo-server

Simple Go HTTP server that logs request URL and body to stdout and responds "ok".

Usage:

- Build locally: `go build -o server .`
- Run: `PORT=8080 ./server`

Docker:

Build image:

```
docker build -t go-echo-server .
```

Run container:

```
docker run -p 8080:8080 -e PORT=8080 go-echo-server
```

Environment variables:

- `PORT` (default `8080`)
- `READ_TIMEOUT` (default `5s`)
- `WRITE_TIMEOUT` (default `10s`)
