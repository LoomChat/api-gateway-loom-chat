# Loom Chat API Gateway 
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
