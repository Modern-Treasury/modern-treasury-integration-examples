## API example: Python + Flask

This small application demonstrates how you might set up a web server
using Python and [Flask][flask] with RESTful routes to accept your modern-treasury-js
form submissions without having to handle sensitive data.

This example makes use of the official [Python client library][client] for the Modern Treasury API.

Note that it is not necessary to use the Flask framework. In this example it is
used to organize various API actions into distinct application routes, but one
could just as easily implement these API actions within another application
framework.

### Use

#### Local

1. Start the server

  ```bash
  $ pip install -r requirements.txt
  $ FLASK_APP=app.py flask run -p 9001
  ```
2. Open [http://localhost:9001](http://localhost:9001)

[flask]: https://flask.palletsprojects.com/en/2.2.x/
[client]: https://github.com/Modern-Treasury/modern-treasury-python