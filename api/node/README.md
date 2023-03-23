## API example: Node + Express

This small application demonstrates how you might set up a web server
using Node.js and [Express][express] with RESTful routes to accept your modern-treasury-js
form submissions without having to handle sensitive data.

This example makes use of the official [node client library][client] for the Modern Treasury API.

Note that it is not necessary to use the Express framework. In this example it is
used to organize various API actions into distinct application routes, but one
could just as easily implement these API actions within another application
framework.

### Use

#### Local

1. Start the server

  ```bash
  $ npm i
  $ node app
  ```
2. Open [http://localhost:9001](http://localhost:9001)

[express]: https://expressjs.com/
[client]: https://github.com/Modern-Treasury/modern-treasury-node