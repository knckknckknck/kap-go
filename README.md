# KAP Go Client - RESTful APIs
[![rest-docs][rest-doc-img]][rest-doc]

Go client for KAP (Public Disclosure Platform) — fetch public disclosures from Borsa Istanbul listed companies

**Note:** API last changed on March 27, 2025. For details, please refer to the [official documentation](https://apiportal.mkk.com.tr/).

## Getting Started

First, make a new directory for your project and navigate into it:

```bash
mkdir myproject && cd myproject
```

Next, initialize a new module for dependency management. This creates a go.mod file to track your dependencies:

```bash
go mod init example
```

Then, create a main.go file. For quick start, you can find example code snippets that demonstrate connecting the REST APIs. Here's an example that fetches the ...

Please remember to set your MKK API key, which you can find on the [MKK API Portal](https://apiportal.mkk.com.tr/), in the environment variable MKK_API_KEY. Or, as a less secure option, by hardcoding it in your code. But please note that hardcoding the API key can be risky if your code is shared or exposed. You can configure the environment variable by running:

```bash
export MKK_API_KEY=your_api_key_here
```

Then, run go mod tidy to automatically download and install the necessary dependencies. This command ensures your go.mod file reflects all dependencies used in your project:

```bash
go mod tidy
```

Finally, execute your application:

```bash
go run main.go
```

## Contributing

We welcome contributions to this project! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Submit a pull request.

## License

This project is licensed under the MIT License. See the [MIT](LICENSE) file for details.
