// Package kap provides a Go client for the KAP (Public Disclosure Platform)
// REST API operated by MKK.
//
// KAP publishes disclosures, financial reports, and corporate actions for
// Borsa Istanbul listed companies. This client covers all 12 API endpoints
// using only the Go standard library.
//
// # Production usage
//
// In production, authenticate with your API key and generate a bearer token:
//
//	client := kap.NewClient("YOUR-API-KEY")
//	token, err := client.GenerateToken(ctx)
//
// # Test environment
//
// For the test environment, use basic authentication:
//
//	client := kap.NewClient("", kap.WithBasicAuth("user", "pass"))
//
// # Fetching disclosures
//
//	disclosures, err := client.Disclosures(ctx, 1092228, nil)
//	detail, err := client.DisclosureDetail(ctx, 1211180, "data", "")
//
// All methods accept a context.Context for cancellation and timeout control.
// Errors returned by the API are represented as *APIError values which
// can be unwrapped to sentinel errors (ErrUnauthorized, ErrTokenExpired, etc.)
// using errors.Is.
package kap
