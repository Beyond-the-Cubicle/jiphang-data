package store

import (
	"database/sql"
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

func NewSeoulStore() *seoulStore {
	databaseType := "mysql"
	databaseUrl := "root:@tcp(127.0.0.1:3306)/chulgeun_gil_planner"

	db, err := sql.Open(databaseType, databaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &seoulStore{db: db}
}

func (store *seoulStore) CloseSeoulStore() {
	store.db.Close()
}
