package clientoptions

import (
	"context"

	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The ExamplePackage interface is an example interface.
type ClientOptions interface {
	Parse(ctx context.Context, grpcUrl string) (string, []grpc.DialOption, error)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------
