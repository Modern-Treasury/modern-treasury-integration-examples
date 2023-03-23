# Contributing

### Contributing to this project

Contributors are highly valued and welcomed! This project is designed in a manner that facilitates the expansion of both frontend and language-specific backends. Every backend example is compatible with the same set of frontend examples, enabling seamless integration of new frontend implementations with any backend server.

If you have any questions or need assistance regarding the guidelines or making a contribution, we are here to help. Simply [create a Pull Request][new-pr] with your proposal, and we will gladly provide support and guidance to ensure a smooth collaboration.

#### Creating new frontend examples

1. Begin by creating a fresh directory in the [public folder](public) to house your example. Make sure to keep all your HTML, CSS, and JS files within this directory.
2. Determine the desired action for your form and submit it to the appropriate endpoint mentioned in the [API server specifications](#api-server-specifications).
3. Modify [index.html](public/index.html) to include a link to your newly created example.
4. Update the [README](README.md) to include a link to the code directory of your new example, ensuring easy access for other users.

#### Creating new backend examples

1. Start by creating a fresh directory within the [api directory](api), and name it after the programming language you intend to add.
2. Develop endpoints that align with the specifications mentioned in the [API server specifications](#api-server-specifications).
3. Craft a clear and descriptive README that explains how to launch your server and provides guidance on accessing the examples through a web browser.
4. Update both the [main README](README.md) and the [API README](api/README.md)  to include links to the code directory of your newly added example, making it easily accessible to others.

### API Server specifications

| Endpoint | Action |
| -------- | ------ |
| POST `/api/create-cp-acf` | New counterparty and associated account collection flow |
| POST `/api/create-cp-pf` | New counterparty and associated payment flow |

All GET requests should directly serve files from the [public directory](public).

### External examples


If you have developed an application or website that utilizes modern-treasury-js in a new or innovative way, we encourage you to [create an issue][new-issue] containing a link to your site. We will then include a link to it in the README, showcasing your unique implementation.


[new-issue]: https://github.com/modern-treasury/modern-treasury-js-examples/issues/new
[new-pr]: https://github.com/modern-treasury/modern-treasury-js-examples/pulls/new