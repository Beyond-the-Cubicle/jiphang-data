package store

import (
	"database/sql"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/config"
	_ "github.com/go-sql-driver/mysql"
)

type SeoulStore interface {
	CreateBusLinks()
	DeleteAllBusLinks()
}

type seoulStore struct {
	db *sql.DB
}

func NewSeoulStore(appConfig config.Config) *seoulStore {
	db, err := sql.Open(appConfig.DatabaseType, appConfig.DatabaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &seoulStore{db: db}
}

func (store *seoulStore) Close() {
	store.db.Close()
}
