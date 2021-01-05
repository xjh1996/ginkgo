package client

import (
	"github.com/caicloud/resource/pkg/server/client/v20201010"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes a versioned client.
type Interface interface {
	// V20201010 returns v20201010 client.
	V20201010() v20201010.Interface
}

// Client contains versioned clients.
type Client struct {
	v20201010 *v20201010.Client
}

// NewClient creates a new client.
func NewClient(cfg *rest.Config) (Interface, error) {
	c := &Client{}
	var err error

	c.v20201010, err = v20201010.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// MustNewClient creates a new client or panic if an error occurs.
func MustNewClient(cfg *rest.Config) Interface {
	return &Client{
		v20201010: v20201010.MustNewClient(cfg),
	}
}

// V20201010 returns a versioned client.
func (c *Client) V20201010() v20201010.Interface {
	return c.v20201010
}
