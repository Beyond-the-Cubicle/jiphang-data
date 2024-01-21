package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
	CreateBusStations(stationId string, stationName string, arsId string, latitude float64, longitude float64, cityCode string, cityName string) error
	ReadBusStation(stationId string) (BusStation, error)
	ReadBusStations(stationIds []string) ([]BusStation, error)
	ReadAllBusStations() ([]BusStation, error)
	DeleteAllBusStations() error
}

type store struct {
	db *sql.DB
}

func New() *store {
	databaseType := "mysql"
	databaseUrl := "root:@tcp(127.0.0.1:3306)/chulgeun_gil_planner"

	db, err := sql.Open(databaseType, databaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &store{db: db}
}

func (store *store) Close() {
	store.db.Close()
}
