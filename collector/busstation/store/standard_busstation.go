package store

import "strings"

type StandardBusStation struct {
	StationId   string
	StationName string
	ArsId       string
	Latitude    float64
	Longitude   float64
}

func (store *standardStore) CreateBusStations(stationId string, stationName string, arsId string, latitude float64, longitude float64) error {
	_, err := store.db.Exec("INSERT INTO standard_bus_station VALUES (?, ?, ?, ?, ?)",
		stationId,
		stationName,
		arsId,
		latitude,
		longitude,
	)
	return err
}

func (store *standardStore) ReadBusStation(stationId string) (StandardBusStation, error) {
	result := StandardBusStation{}
	err := store.db.QueryRow("SELECT * FROM standard_bus_station WHERE stationId = ?", stationId).Scan(
		&result.StationId,
		&result.StationName,
		&result.ArsId,
		&result.Latitude,
		&result.Longitude,
	)
	return result, err
}

func (store *standardStore) ReadBusStations(stationIds []string) ([]StandardBusStation, error) {
	results := []StandardBusStation{}
	rows, err := store.db.Query("SELECT * FROM standard_bus_station WHERE stationId in (?" + strings.Repeat(",?", len(stationIds)-1) + ")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := StandardBusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.ArsId,
			&result.Latitude,
			&result.Longitude,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *standardStore) ReadAllBusStations() ([]StandardBusStation, error) {
	results := []StandardBusStation{}
	rows, err := store.db.Query("SELECT * FROM standard_bus_station")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := StandardBusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.ArsId,
			&result.Latitude,
			&result.Longitude,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *standardStore) DeleteAllBusStations() error {
	_, err := store.db.Exec("DELETE FROM standard_bus_station")
	return err
}
