package app

import (
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

type App interface {
	CollectSeoulBusStations(docType DocType) ([]SeoulOpenAPIBusStation, error)
	CollectGyunggiBusStations(docType DocType) ([]GyunggiOpenAPIBusStation, error)
	InsertSeoulBusStations(seoulOpenApiBusStations []SeoulOpenAPIBusStation) error
	InsertGyunggiBusStations(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) error
	ConvertSeoulBusStationsToStandard(seoulOpenApiBusStations []SeoulOpenAPIBusStation) ([]store.StandardBusStation, error)
	ConvertGyunggiBusStationsToStandard(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) ([]store.StandardBusStation, error)
	InsertBusStations(busStations []store.StandardBusStation) error
}

type app struct {
	seoulApiKey   string
	gyunggiApiKey string
	store         store.StandardStore
	seoulStore    store.SeoulStore
	gyunggiStore  store.GyunggiStore
}

func New(appConfig config.Config, store store.StandardStore, seoulStore store.SeoulStore, gyunggiStore store.GyunggiStore) *app {
	return &app{
		seoulApiKey:   appConfig.SeoulApiKey,
		gyunggiApiKey: appConfig.GyunggiApiKey,
		store:         store,
		seoulStore:    seoulStore,
		gyunggiStore:  gyunggiStore,
	}
}
