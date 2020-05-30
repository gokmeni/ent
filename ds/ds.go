package ds

import (
	"database/sql"
	"github.com/gokmeni/ent/ent"
	"os"
	"sync"
	"time"

	entsql "github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/jackc/pgx/stdlib"
)

const (
	postgresEnvVariable string = "POSTGRES_CONNECTION_STRING"
)

var client *ent.Client
var once sync.Once

func GetConnection() *ent.Client {
	once.Do(func() {
		db, err := sql.Open("pgx", os.Getenv(postgresEnvVariable))

		if err != nil {
			panic(err)
		}

		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
		db.SetConnMaxLifetime(1 * time.Minute)

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		drv := entsql.OpenDB("postgres", db)
		client = ent.NewClient(ent.Driver(drv))
	})

	return client
}

