package store

import (
	"database/sql"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/config"
)

type GyunggiStore interface {
	CreateBusLinks()
	DeleteAllBusLinks()
}

type gyunggiStore struct {
	db *sql.DB
}

func NewGyunggiStore(appConfig config.Config) *gyunggiStore {
	db, err := sql.Open(appConfig.DatabaseType, appConfig.DatabaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &gyunggiStore{db: db}
}

func (store *gyunggiStore) Close() {
	store.db.Close()
}
