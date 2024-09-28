package finout

import (
	"context"
	"net/http"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

const (
	DefaultBaseURL = "https://app.finout.io/v1"
)

func NewSecuredClient(clientID, secretKey string) (*Client, error) {
	headerClientID, err := securityprovider.NewSecurityProviderApiKey("header", "x-finout-client-id", clientID)
	if err != nil {
		return nil, err
	}

	headerSecretKey, err := securityprovider.NewSecurityProviderApiKey("header", "x-finout-secret-key", secretKey)
	if err != nil {
		return nil, err
	}

	return NewClient(
		DefaultBaseURL,
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Accept", "application/json")
			return nil
		}),
		WithRequestEditorFn(headerClientID.Intercept),
		WithRequestEditorFn(headerSecretKey.Intercept),
		// WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		// 	dump, err := httputil.DumpRequestOut(req, true)

		// 	if err != nil {
		// 		return err
		// 	}

		// 	log.Printf("\n\n%q", dump)

		// 	return nil
		// }),
	)
}

func NewSecuredClientWithResponses(clientID, secretKey string) (*ClientWithResponses, error) {
	headerClientID, err := securityprovider.NewSecurityProviderApiKey("header", "x-finout-client-id", clientID)
	if err != nil {
		return nil, err
	}

	headerSecretKey, err := securityprovider.NewSecurityProviderApiKey("header", "x-finout-secret-key", secretKey)
	if err != nil {
		return nil, err
	}

	return NewClientWithResponses(
		DefaultBaseURL,
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Accept", "application/json")
			return nil
		}),
		WithRequestEditorFn(headerClientID.Intercept),
		WithRequestEditorFn(headerSecretKey.Intercept),
		// WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		// 	dump, err := httputil.DumpRequestOut(req, true)

		// 	if err != nil {
		// 		return err
		// 	}

		// 	log.Printf("\n\n%q", dump)

		// 	return nil
		// }),
	)
}
