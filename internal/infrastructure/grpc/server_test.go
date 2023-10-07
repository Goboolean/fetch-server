package grpc_test

import (
	"context"
	"os"
	"testing"

	pb "github.com/Goboolean/fetch-system.master/api/grpc"
	server "github.com/Goboolean/fetch-system.master/internal/infrastructure/grpc"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/stretchr/testify/assert"

	_ "github.com/Goboolean/common/pkg/env"
)

var (
	host   *server.Host
	client *server.Client
)


var (
	registered = false
	healthy    = false
)
type MockAdapter struct {
	pb.UnimplementedWorkerServer
}

func (a *MockAdapter) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	registered = true
	return &pb.RegisterResponse{}, nil
}

func (a *MockAdapter) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	healthy = true
	return &pb.HealthCheckResponse{}, nil
}


func SetupHost() *server.Host {
	host, err := server.New(&resolver.ConfigMap{
		"PORT": os.Getenv("GRPC_PORT"),
	}, &MockAdapter{})
	if err != nil {
		panic(err)
	}

	return host
}

func TeardownHost(h *server.Host) {
	h.Close()
}

func SetupClient() *server.Client {
	client, err := server.NewClient(&resolver.ConfigMap{
		"PORT": os.Getenv("GRPC_PORT"),
	})
	if err != nil {
		panic(err)
	}

	return client
}

func TeardownClient(c *server.Client) {
	c.Close()
}


func TestMain(m *testing.M) {
	host = SetupHost()
	client = SetupClient()
	code := m.Run()
	TeardownHost(host)
	TeardownClient(client)
	os.Exit(code)
}


func Test_Constructor(t *testing.T) {

	t.Run("Ping()", func(t *testing.T) {
		err := client.Ping()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}


func Test_Method(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		_, err := client.Register(context.Background(), &pb.RegisterRequest{})
		assert.NoError(t, err)
		assert.True(t, registered)
	})

	t.Run("HealthCheck", func(t *testing.T) {
		_, err := client.HealthCheck(context.Background(), &pb.HealthCheckRequest{})
		assert.NoError(t, err)
		assert.True(t, healthy)
	})
}