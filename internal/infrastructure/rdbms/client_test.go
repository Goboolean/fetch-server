package rdbms_test

import (
	"context"
	"os"
	"testing"

	_ "github.com/Goboolean/common/pkg/env"
	"github.com/Goboolean/common/pkg/resolver"
	"github.com/Goboolean/fetch-system.master/internal/infrastructure/rdbms"
	"github.com/stretchr/testify/assert"
)

func Setup() *rdbms.DB {
	db, err := rdbms.NewDB(&resolver.ConfigMap{
		"HOST":     os.Getenv("PSQL_HOST"),
		"PORT":     os.Getenv("PSQL_PORT"),
		"USER":     os.Getenv("PSQL_USER"),
		"PASSWORD": os.Getenv("PSQL_PASS"),
		"DATABASE": os.Getenv("PSQL_DATABASE"),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func Teardown(db *rdbms.DB) {
	db.Close()
}

var db *rdbms.DB

func TestMain(m *testing.M) {
	db = Setup()
	defer Teardown(db)
	code := m.Run()
	os.Exit(code)
}

func TestPing(t *testing.T) {
	t.Run("Ping()", func(t *testing.T) {
		err := db.Ping()
		assert.NoError(t, err)
	})
}

func TestQueries(t *testing.T) {
	t.Run("GetAllMetadata()", func(t *testing.T) {
		_, err := db.GetAllMetadata(context.Background())
		assert.NoError(t, err)
	})
}
