package store

import (
	"database/sql"
)

type GyunggiStore interface {
	CreateBusStations(
		stationId string,
		stationName string,
		coordinateX float64,
		coordinateY float64,
		gpsCoordinateX float64,
		gpsCoordinateY float64,
		rinkId string,
		stationType string,
		transferStationExtNo string,
		medianBusLaneYn string,
		stationEnglishName string,
		arsId string,
		institutionCode string,
		dataDisplayYn string,
		registeredBy string,
		registeredAt string,
		memo string,
		signPostType string,
		dongCode string,
		regionCode string,
		useYn string,
		stationChineseName string,
		stationJapaneseName string,
		stationVietnamName string,
		drtYn string,
		stationTypeName string,
		transferStationTypeName string,
		signPostTypeName string,
	) error
	ReadBusStation(stationId string) (GyunggiBusStation, error)
	ReadBusStations(stationIds []string) ([]GyunggiBusStation, error)
	ReadAllBusStations() ([]GyunggiBusStation, error)
	DeleteAllBusStations() error
}

type gyunggiStore struct {
	db *sql.DB
}

func NewGyunggiStore() *gyunggiStore {
	databaseType := "mysql"
	databaseUrl := "root:@tcp(127.0.0.1:3306)/chulgeun_gil_planner"

	db, err := sql.Open(databaseType, databaseUrl)
	if err != nil {
		panic("failed to connect database - " + err.Error())
	}

	return &gyunggiStore{db: db}
}

func (store *gyunggiStore) CloseGyunggiStore() {
	store.db.Close()
}
