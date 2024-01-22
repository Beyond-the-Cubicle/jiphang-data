package app

import "github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"

type GyunggiOpenAPIBusStation struct {
	// TBD
}

func (app *app) CollectGyunggiBusStations(apiKey string, docType DocType) ([]GyunggiOpenAPIBusStation, error) {
	return nil, nil
}
func (app *app) ConvertGyunggiBusStationsToStandard(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) ([]store.BusStation, error) {
	return nil, nil
}
