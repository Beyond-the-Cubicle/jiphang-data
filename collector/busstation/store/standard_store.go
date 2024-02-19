package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type StandardStore interface {
	CreateBusStations(stationId string, stationName string, arsId string, latitude float64, longitude float64, cityCode string, cityName string) error
	ReadBusStation(stationId string) (StandardBusStation, error)
	ReadBusStations(stationIds []string) ([]StandardBusStation, error)
	ReadAllBusStations() ([]StandardBusStation, error)
	DeleteAllBusStations() error
}

type standardStore struct {
	db *sql.DB
}

func NewStandardStore() *standardStore {
	databaseType := "mysql"
	databaseUrl := "root:@tcp(127.0.0.1:3306)/chulgeun_gil_planner"

	db, err := sql.Open(databaseType, databaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &standardStore{db: db}
}

func (store *standardStore) Close() {
	store.db.Close()
}
