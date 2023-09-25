package grpc

import (
	"fmt"

	pb "github.com/Goboolean/fetch-system.master/api/grpc"
	"github.com/Goboolean/shared/pkg/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	pb.WorkerClient

	conn *grpc.ClientConn
}

func NewClient(c *resolver.ConfigMap) (*Client, error) {

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf(":%s", port)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Client{
		WorkerClient: pb.NewWorkerClient(conn),
		conn:                    conn,
	}, nil
}

func (c *Client) Ping() error {
	fmt.Println("ping")
	state := c.conn.GetState()
	if state != connectivity.Ready {
		return fmt.Errorf("connection is not ready")
	}
	return nil
}

func (c *Client) Close() {
	c.conn.Close()
}
