package vsphere

import (
	"context"
	"errors"
)

// ErrNotImplemented is returned by stub methods
var ErrNotImplemented = errors.New("not implemented")

// Client interface for vSphere operations
type Client interface {
	// Connection
	Login(ctx context.Context) error
	Logout(ctx context.Context) error

	// Info
	Version(ctx context.Context) (string, error)
	CurrentUser(ctx context.Context) (string, error)
}

// StubClient is a stub implementation that returns not implemented errors
type StubClient struct{}

// NewStubClient creates a new stub client
func NewStubClient() *StubClient {
	return &StubClient{}
}

func (c *StubClient) Login(ctx context.Context) error {
	return ErrNotImplemented
}

func (c *StubClient) Logout(ctx context.Context) error {
	return ErrNotImplemented
}

func (c *StubClient) Version(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}

func (c *StubClient) CurrentUser(ctx context.Context) (string, error) {
	return "", ErrNotImplemented
}
