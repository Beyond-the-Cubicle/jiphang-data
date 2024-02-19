package app

import "github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"

type App interface {
	CollectSeoulBusStations(apiKey string, docType DocType) ([]SeoulOpenAPIBusStation, error)
	CollectGyunggiBusStations(apiKey string, docType DocType) ([]GyunggiOpenAPIBusStation, error)
	InsertSeoulBusStations(seoulOpenApiBusStations []SeoulOpenAPIBusStation) error
	InsertGyunggiBusStations(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) error
	ConvertSeoulBusStationsToStandard(seoulOpenApiBusStations []SeoulOpenAPIBusStation) ([]store.StandardBusStation, error)
	ConvertGyunggiBusStationsToStandard(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) ([]store.StandardBusStation, error)
	InsertBusStations(busStations []store.StandardBusStation) error
}

type app struct {
	store        store.StandardStore
	seoulStore   store.SeoulStore
	gyunggiStore store.GyunggiStore
}

func New(store store.StandardStore, seoulStore store.SeoulStore, gyunggiStore store.GyunggiStore) *app {
	return &app{
		store:        store,
		seoulStore:   seoulStore,
		gyunggiStore: gyunggiStore,
	}
}
