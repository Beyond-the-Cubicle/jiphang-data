package store

import (
	"database/sql"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
)

type SeoulStore interface {
	CreateBusStations(stationId string, stationName string, stationType string, arsId string, coordinateX float64, coordinateY float64, busArrivalInfoGuideInstallYn string) error
	ReadBusStation(stationId string) (SeoulBusStation, error)
	ReadBusStations(stationIds []string) ([]SeoulBusStation, error)
	ReadAllBusStations() ([]SeoulBusStation, error)
	DeleteAllBusStations() error
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
