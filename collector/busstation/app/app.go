package app

import "github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"

type App interface {
	CollectSeoulBusStations(apiKey string, docType DocType) ([]SeoulOpenAPIBusStation, error)
	CollectGyunggiBusStations(apiKey string, docType DocType) ([]GyunggiOpenAPIBusStation, error)
	ConvertSeoulBusStationsToStandard(seoulOpenApiBusStations []SeoulOpenAPIBusStation) ([]store.BusStation, error)
	ConvertGyunggiBusStationsToStandard(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) ([]store.BusStation, error)
	InsertBusStations(busStations []store.BusStation) error
}

type app struct {
	store store.Store
}

func New(store store.Store) *app {
	return &app{
		store: store,
	}
}
