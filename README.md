# kap-go

Go client for the [KAP (Public Disclosure Platform)](https://www.kap.org.tr/) REST API operated by MKK. Fetch public disclosures, financial reports, and corporate actions for Borsa Istanbul listed companies.

**Note:** API last changed on March 27, 2025. For details, please refer to the [official documentation](https://apiportal.mkk.com.tr/).

## Installation

```bash
go get github.com/knckknckknck/kap-go
```

## Quick Start

### Production

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/knckknckknck/kap-go"
)

func main() {
	client := kap.NewClient(os.Getenv("MKK_API_KEY"))

	ctx := context.Background()

	// Generate a bearer token (valid for 24 hours).
	_, err := client.GenerateToken(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch disclosures starting from index 1092228.
	disclosures, err := client.Disclosures(ctx, 1092228, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range disclosures {
		fmt.Println(d.DisclosureIndex, d.Title)
	}
}
```

### Test Environment

```go
client := kap.NewClient("",
	kap.WithBasicAuth("user", "pass"),
	kap.WithBaseURL("https://apigwdev.mkk.com.tr"),
)
```

## Error Handling

All API errors are returned as `*kap.APIError` and can be matched against sentinel errors:

```go
import "errors"

_, err := client.Disclosures(ctx, 1092228, nil)
if errors.Is(err, kap.ErrTokenExpired) {
	// Re-generate token and retry.
	client.GenerateToken(ctx)
}
if errors.Is(err, kap.ErrUnauthorized) {
	// Check API key or permissions.
}
```

Available sentinel errors: `ErrNoPermission`, `ErrUnauthorized`, `ErrIPRestricted`, `ErrInvalidToken`, `ErrIPVerification`, `ErrTokenExpired`, `ErrTokenValidation`, `ErrNotFound`, `ErrUnexpectedStatus`.

## Configuration Options

```go
kap.WithBaseURL(url)           // Set API base URL
kap.WithTimeout(duration)      // Set HTTP timeout
kap.WithHTTPClient(client)     // Use custom http.Client
kap.WithToken(token)           // Set pre-existing bearer token
kap.WithBasicAuth(user, pass)  // Use basic auth (test environment)
```

## Documentation

- [API Reference (docs/kap-rest-api.md)](docs/kap-rest-api.md)
- [Go Package Documentation](https://pkg.go.dev/github.com/knckknckknck/kap-go)

## Contributing

We welcome contributions to this project! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Submit a pull request.

## License

This project is licensed under the MIT License. See the [MIT](LICENSE) file for details.
