# PublicIP-API [![Go Report Card](https://goreportcard.com/badge/github.com/TimoSLE/go-dyndns)](https://goreportcard.com/report/github.com/TimoSLE/publicip-api)
A simple HTTP API returning the API of the requesting client

##Building
If you are using make you can build the executable with
>make build

Otherwise, you can use
>go build ./...

##Usage
There are to command line flags which modify the functionality of the API

Flag | Explanation | Default
--- | --- | ---
b | Binding IP Address for HTTP Server | 127.0.0.1:8080
h | Header to retrieve IP address from | X-Real-IP

The request to the API should be made in the following format
> https://exampleapi.com/?format={format}

Formats json, text and xml are available, if given format is not found/there is no format specified, the api will be defaulting to plain text

## Contributing
Every Contribution is welcome, feel free to open pull requests and/or suggest changes in issues