package etcd

import (
	"context"
	"time"

	"github.com/Goboolean/shared/pkg/resolver"
	"go.etcd.io/etcd/client/v3"
)




type Client struct {
	client *clientv3.Client
}

func New(c *resolver.ConfigMap) (*Client, error) {

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	config := clientv3.Config{
		Endpoints:   []string{host},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}


func (c *Client) Ping(ctx context.Context) error {
	mapi := clientv3.NewMaintenance(c.client)
	_, err := mapi.Status(ctx, c.client.Endpoints()[0])
	return err
}


func (c *Client) Close() error {
	return c.client.Close()
}