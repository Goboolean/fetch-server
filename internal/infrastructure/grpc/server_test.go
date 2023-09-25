package grpc_test

import (
	"os"
	"testing"

	pb "github.com/Goboolean/fetch-system.master/api/grpc"
	server "github.com/Goboolean/fetch-system.master/internal/infrastructure/grpc"
	"github.com/Goboolean/shared/pkg/resolver"

	_ "github.com/Goboolean/common/pkg/env"
)

var (
	host   *server.Host
	client *server.Client
)

type MockAdapter struct {
	pb.UnimplementedWorkerServer
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
	
}