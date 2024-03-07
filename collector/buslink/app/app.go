package app

import (
	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/store"
)

type App interface {
	CollectSeoulBusLinks(docType DocType) ([]SeoulOpenApiBusLink, error)
	InsertSeoulBusLinks(seoulOpenApiBusLinks []SeoulOpenApiBusLink) error
	CollectGyunggiBusLinks(docType DocType) ([]GyunggiOpenApiBusLink, error)
	InsertGyunggiBusLinks(gyunggiOpenApiBusLinks []GyunggiOpenApiBusLink) error
}

type app struct {
	seoulApiKey   string
	gyunggiApiKey string
	seoulStore    store.SeoulStore
	gyunggiStore  store.GyunggiStore
}

func New(appConfig config.Config, seoulStore store.SeoulStore, gyunggiStore store.GyunggiStore) *app {
	return &app{
		seoulApiKey:   appConfig.SeoulApiKey,
		gyunggiApiKey: appConfig.GyunggiApiKey,
		seoulStore:    seoulStore,
		gyunggiStore:  gyunggiStore,
	}
}
