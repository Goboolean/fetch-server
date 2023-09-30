package etcd_test

import (
	"context"
	"os"
	"testing"

	"github.com/Goboolean/fetch-system.master/internal/infrastructure/etcd"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)



var client *etcd.Client

func Setup() *etcd.Client {
	c, err := etcd.New(&resolver.ConfigMap{
		"HOST": os.Getenv("ETCD_HOST"),
	})
	if err != nil {
		panic(err)
	}

	return c
}

func Teardown(c *etcd.Client) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	client = Setup()
	defer Teardown(client)

	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: false,
	})

	code := m.Run()
	os.Exit(code)
}


func Test_Constructor(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		if err := client.Ping(context.Background()); err != nil {
			t.Fatal(err)
		}
	})
}


func Test_Worker(t *testing.T) {

	var workers []*etcd.Worker = []*etcd.Worker{
		{ID: uuid.New().String(), Platform: "polygon", Status:   "active"},
		{ID: uuid.New().String(), Platform: "polygon", Status:   "pending"},
		{ID: uuid.New().String(), Platform: "kis",     Status:   "active"},
		{ID: uuid.New().String(), Platform: "kis",     Status:   "pending"},
	}

	t.Run("InsertWorker", func(t *testing.T) {
		for _, w := range workers {
			err := client.InsertWorker(context.Background(), w)
			assert.NoError(t, err)
		}
	})

	t.Run("GetWorker", func(t *testing.T) {
		for _, w := range workers {
			worker, err := client.GetWorker(context.Background(), w.ID)
			assert.NoError(t, err)
			assert.Equal(t, w, worker)
		}
	})

	t.Run("UpdateWorkerStatus", func(t *testing.T) {
		err := client.UpdateWorkerStatus(context.Background(), workers[0].ID, "dead")
		assert.NoError(t, err)

		w, err := client.GetWorker(context.Background(), workers[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, "dead", w.Status)
	})

	t.Run("DeleteWorker", func(t *testing.T) {
		err := client.DeleteWorker(context.Background(), workers[0].ID)
		assert.NoError(t, err)

		_, err = client.GetWorker(context.Background(), workers[0].ID)
		assert.Error(t, err)
	})

	t.Run("GetAllWorkers", func(t *testing.T) {
		ws, err := client.GetAllWorkers(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, len(workers)-1, len(ws))
	})
}


func Test_Product(t *testing.T) {

	var products []*etcd.Product = []*etcd.Product{
		{ID: "test.goboolean.kor", Platform: "kis",      Symbol: "goboolean", Worker: uuid.New().String(), Status: "onsubscribe"},
		{ID: "test.goboolean.eng", Platform: "polygon",  Symbol: "gofalse",   Worker: uuid.New().String(), Status: "onsubscribe"},
		{ID: "test.goboolean.jpn", Platform: "buycycle", Symbol: "gonil",     Worker: uuid.New().String(), Status: "onsubscribe"},
		{ID: "test.goboolean.chi", Platform: "kis",      Symbol: "gotrue",    Worker: uuid.New().String(), Status: "onsubscribe"},
	}

	t.Run("InsertProducts", func(t *testing.T) {
		err := client.InsertProducts(context.Background(), products)
		assert.NoError(t, err)
	})

	t.Run("GetProduct", func(t *testing.T) {
		for _, p := range products {
			product, err := client.GetProduct(context.Background(), p.ID)
			assert.NoError(t, err)
			assert.Equal(t, p, product)
		}
	})

	t.Run("UpdateProductStatus", func(t *testing.T) {
		err := client.UpdateProductStatus(context.Background(), products[0].ID, "notsubscribed")
		assert.NoError(t, err)

		p, err := client.GetProduct(context.Background(), products[0].ID)
		assert.NoError(t, err)
		assert.Equal(t, "notsubscribed", p.Status)
	})

	t.Run("UpdateProductWorker", func(t *testing.T) {
		err := client.UpdateProductWorker(context.Background(), products[1].ID, uuid.New().String())
		assert.NoError(t, err)

		p, err := client.GetProduct(context.Background(), products[1].ID)
		assert.NoError(t, err)
		assert.NotEqual(t, products[0].Worker, p.Worker)
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		err := client.DeleteProduct(context.Background(), products[2].ID)
		assert.NoError(t, err)

		_, err = client.GetProduct(context.Background(), products[2].ID)
		assert.Error(t, err)
	})

	t.Run("InsertOneProduct", func(t *testing.T) {
		err := client.InsertOneProduct(context.Background(), products[2])
		assert.NoError(t, err)

		p, err := client.GetProduct(context.Background(), products[2].ID)
		assert.NoError(t, err)
		assert.Equal(t, products[2], p)
	})

	t.Run("GetAllProducts", func(t *testing.T) {
		ps, err := client.GetAllProducts(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, len(products)-1, len(ps))
	})
}