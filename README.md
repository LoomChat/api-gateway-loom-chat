# Loom Chat API Gateway 
*NOTE: I use amost exclusively the Go standard library and no thrid-party library or framework
for educational purposes! But in production you better use robust, well-known, production-ready
tools!*

This is the API Gateway for the whole internal backend services of the Loom Chat System!
All the HTTP/WebSocket requests coming from the Front-End/UI goes through this API Gateway
first.

## General project structure
```
gateway-service/
├── cmd/
│   └── gateway/
│       └── main.go
├── internal/
│   ├── server/          # HTTP & WS server
│   ├── middleware/      # auth, logging, rate limit
│   ├── routing/         # route definitions
│   ├── proxy/           # forwarding to services
│   ├── ws/              # websocket connection mgmt
│   └── config/
├── pkg/                 # reusable (optional)
└── README.md
```
