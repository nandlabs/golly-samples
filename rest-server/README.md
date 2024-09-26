# golly: rest-server - code examples

### Overview
Shows how to use the rest-server library and create REST APIs for your applications.

### Code Examples
---
- [simple-server](/rest-server/simple-server/server.go)
- [server-routing](/rest-server/server-routing/main.go)

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

### Contributing
---
If you want to contribute any more examples related to turbo, please follow the below guidelines on how to create a module.
1. Follow the basec contribution guidelines present [in Readme](/README.md)
2. Create the `example` folder under the `rest-server` folder.
3. Steps to create an individual `go module`
    ```bash
    cd example
    go mod init github.com/nandlabs/golly-samples/rest-server/example

    go get oss.nandlabs.io/golly
    ```
4. Now you should be able to create your example