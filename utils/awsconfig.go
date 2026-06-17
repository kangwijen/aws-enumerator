package utils

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// CustomEndpointConfigured reports whether a global custom AWS endpoint is active.
func CustomEndpointConfigured() bool {
	return os.Getenv("AWS_ENDPOINT_URL") != ""
}

// LoadAWSConfig loads credentials and region via the standard AWS SDK chain.
// When AWS_ENDPOINT_URL is set, it is applied with config.WithBaseEndpoint so
// per-service overrides (AWS_ENDPOINT_URL_S3, etc.) still work through the SDK.
func LoadAWSConfig(ctx context.Context) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}

	if endpointURL := os.Getenv("AWS_ENDPOINT_URL"); endpointURL != "" {
		opts = append(opts, config.WithBaseEndpoint(endpointURL))
	}

	return config.LoadDefaultConfig(ctx, opts...)
}
