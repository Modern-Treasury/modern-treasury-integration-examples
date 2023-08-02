Modern Treasury integration examples
===================
<p align="center">
  <img src="https://files.readme.io/a49b14e-account-collection.gif" align="center">
</p>

This repository contains a set of example implementations of
[modern-treasury-js][modern-treasury-js] using HTML, CSS, and JavaScript, and a set of API usage
examples to demonstrate how to leverage our [pre-built UIs][pre-built-uis] to quickly integrate common payments workflows.

This code is intended as an example implementation and should not be used directly in a production environment without thorough review and modification. While this code provides a functional implementation that demonstrates certain concepts or techniques, it may lack important considerations for security, performance, scalability, and reliability.

### Account Collection form example

- [Account Collection form example](public/acf.html)

### Payment form example

- [Payment form example](public/pf.html)

### User Onboarding form example

- [User Onboarding form example](public/onboarding.html)

### API usage examples

- [Python](api/python)
- [Node](api/node)
- [Go](api/go)
- [Java](api/java)

#### Configuring the examples

Note that each API example in this code will retrieve configuration values from environment variables. You can set these variables accordingly to swiftly configure the example and establish a connection to your Modern Treasury organization.

| Environment variable | description |
| -------------------- | ----------- |
| MT_ORG_ID | The [ID][api-keys]  of your Modern Treasury organization |
| MT_API_KEY | Your [private API key][api-keys] |
| MT_PUB_KEY | Your [publishable API key][publishable-keys] |
| FS_KEY | Used to encrypt session storage


### How to run

Each example in this code is designed to be executed locally.

While running these examples, it is important to modify the code according to your specific requirements for redirection and error handling. However, please note that the example applications were primarily created to demonstrate fundamental API functionalities upon initial execution.

**Note**: Use this code as a reference or starting point for your own implementation, but exercise caution and diligence when adapting it for production use. 


#### Local


To run this code locally, follow these steps:

1. Clone the repository to your local machine.
2. Take some time to go through the simple application code to become familiar with it.
3. Refer to the startup instructions provided in one of the [API usage examples](api) mentioned above to get started.


#### Deploy Immediately to AWS App Runner

You can swiftly deploy this repository on [AWS App Runner][apprunner] by utilizing the provided [Python example](api/python) and [configuration file](apprunner.yaml). If you opt for this deployment method, you are welcome to remove the other language backends from your cloned repository. This will streamline your setup and focus solely on the Python implementation.

### Contributing

[See CONTRIBUTING file](CONTRIBUTING.md).

### License

[MIT](LICENSE.md)

### Special Thanks

This repository draws inspiration from [Recurly's](https://www.recurly.com) excellent [recurly-integration-examples project](https://github.com/recurly/recurly-integration-examples).

[modern-treasury-js]: https://docs.moderntreasury.com/reference/modern-treasury-js
[api-keys]: https://app.moderntreasury.com/developers/api_keys
[publishable-keys]: https://app.moderntreasury.com/settings/developers/publishable_keys
[pre-built-uis]: https://docs.moderntreasury.com/docs/prebuilt-uis-overview
[apprunner]: https://aws.amazon.com/apprunner/