## API example: Go + net/http

This small application demonstrates how you can set up a web server using Go and [net/http](https://pkg.go.dev/net/http), along with RESTful routes, to accept form submissions from a modern-treasury-js enabled application without the hassle of handling sensitive data.

To achieve this, the example utilizes the official [Go client library][client] for the Modern Treasury API.

It's important to note that using the net/http package is not mandatory. In this example, it's used to organize different API actions into separate application routes. However, you can just as easily implement these API actions within another application framework if you prefer.

### Use

#### Local

1. Start the server

  ```bash
  $ go get .
  $ go run app.go
  ```
2. Open [http://localhost:9001](http://localhost:9001)

[client]: https://github.com/Modern-Treasury/modern-treasury-go