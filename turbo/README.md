# golly: turbo - code examples

### Overview
Shows how to use the turbo library and create servers for your applications.

### Code Examples
---
- [simple-server](/turbo/simple-server/server.go)

### How to run the examples
---
1. simple-server
    ```bash
    cd simple-server
    go run server.go

    Server Output ----------
    [timestamp] INFO Initiating Turbo
    [timestamp] INFO Registering New Route: /api/v1/healthz

    Invoke API ----------
    curl http://localhost:8080/api/v1/healthz
    > server is up and running
    ```
