package store

import "strings"

type BusStation struct {
	StationId   string
	StationName string
	ArsId       string
	Latitude    float64
	Longitude   float64
	CityCode    string
	CityName    string
}

func (store *store) CreateBusStations(stationId string, stationName string, arsId string, latitude float64, longitude float64, cityCode string, cityName string) error {
	_, err := store.db.Exec("INSERT INTO bus_station VALUES (?, ?, ?, ?, ?, ?, ?)",
		stationId,
		stationName,
		arsId,
		latitude,
		longitude,
		cityCode,
		cityName,
	)
	return err
}

func (store *store) ReadBusStation(stationId string) (BusStation, error) {
	result := BusStation{}
	err := store.db.QueryRow("SELECT * FROM bus_station WHERE stationId = ?", stationId).Scan(
		&result.StationId,
		&result.StationName,
		&result.Latitude,
		&result.Longitude,
		&result.CityCode,
		&result.CityName,
	)
	return result, err
}

func (store *store) ReadBusStations(stationIds []string) ([]BusStation, error) {
	results := []BusStation{}
	rows, err := store.db.Query("SELECT * FROM bus_station WHERE stationId in (?" + strings.Repeat(",?", len(stationIds)-1) + ")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := BusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.Latitude,
			&result.Longitude,
			&result.CityCode,
			&result.CityName,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *store) ReadAllBusStations() ([]BusStation, error) {
	results := []BusStation{}
	rows, err := store.db.Query("SELECT * FROM bus_station")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := BusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.Latitude,
			&result.Longitude,
			&result.CityCode,
			&result.CityName,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *store) DeleteAllBusStations() error {
	_, err := store.db.Exec("DELETE FROM bus_station")
	return err
}
