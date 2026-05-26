# ConcurrentURLStatusCheck

Check HTTP status codes for multiple URLs concurrently (max 5, 10s timeout per request).

`go run . https://example.com https://google.com`

`go build -o urlcheck . && ./urlcheck https://example.com`
