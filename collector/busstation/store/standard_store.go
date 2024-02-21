package store

import (
	"database/sql"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	_ "github.com/go-sql-driver/mysql"
)

type StandardStore interface {
	CreateBusStations(stationId string, stationName string, arsId string, latitude float64, longitude float64) error
	ReadBusStation(stationId string) (StandardBusStation, error)
	ReadBusStations(stationIds []string) ([]StandardBusStation, error)
	ReadAllBusStations() ([]StandardBusStation, error)
	DeleteAllBusStations() error
}

type standardStore struct {
	db *sql.DB
}

func NewStandardStore(appConfig config.Config) *standardStore {
	db, err := sql.Open(appConfig.DatabaseType, appConfig.DatabaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &standardStore{db: db}
}

func (store *standardStore) Close() {
	store.db.Close()
}
