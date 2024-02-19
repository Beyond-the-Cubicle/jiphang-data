package app

import "github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"

func (app *app) InsertBusStations(busStations []store.StandardBusStation) error {
	for _, busstation := range busStations {
		err := app.store.CreateBusStations(
			busstation.StationId,
			busstation.StationName,
			busstation.ArsId,
			busstation.Latitude,
			busstation.Longitude,
			busstation.CityCode,
			busstation.CityName,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
