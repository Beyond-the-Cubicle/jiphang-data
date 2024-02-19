package store

import "strings"

type SeoulBusStation struct {
	StationId                    string  // 정류장 ID
	StationName                  string  // 정류장 명칭
	StationType                  string  // 정류장 유형 (중앙차로 여부 (일반차로, 중앙차로 ...))
	ArsId                        string  // 정류장 번호(ARS ID)
	CoordinateX                  float64 // 경도
	CoordinateY                  float64 // 위도
	BusArrivalInfoGuideInstallYn string  // 버스도착정보안내기 설치 여부
}

func (store *seoulStore) CreateBusStations(stationId string, stationName string, stationType string, arsId string, coordinateX float64, coordinateY float64, busArrivalInfoGuideInstallYn string) error {
	_, err := store.db.Exec("INSERT INTO seoul_bus_station VALUES (?, ?, ?, ?, ?, ?, ?)",
		stationId,
		stationName,
		stationType,
		arsId,
		coordinateX,
		coordinateY,
		busArrivalInfoGuideInstallYn,
	)
	return err
}

func (store *seoulStore) ReadBusStation(stationId string) (SeoulBusStation, error) {
	result := SeoulBusStation{}
	err := store.db.QueryRow("SELECT * FROM seoul_bus_station WHERE stationId = ?", stationId).Scan(
		&result.StationId,
		&result.StationName,
		&result.StationType,
		&result.ArsId,
		&result.CoordinateX,
		&result.CoordinateY,
		&result.BusArrivalInfoGuideInstallYn,
	)
	return result, err
}

func (store *seoulStore) ReadBusStations(stationIds []string) ([]SeoulBusStation, error) {
	results := []SeoulBusStation{}
	rows, err := store.db.Query("SELECT * FROM seoul_bus_station WHERE stationId in (?" + strings.Repeat(",?", len(stationIds)-1) + ")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := SeoulBusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.StationType,
			&result.ArsId,
			&result.CoordinateX,
			&result.CoordinateY,
			&result.BusArrivalInfoGuideInstallYn,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *seoulStore) ReadAllBusStations() ([]SeoulBusStation, error) {
	results := []SeoulBusStation{}
	rows, err := store.db.Query("SELECT * FROM seoul_bus_station")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := SeoulBusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.StationType,
			&result.ArsId,
			&result.CoordinateX,
			&result.CoordinateY,
			&result.BusArrivalInfoGuideInstallYn,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *seoulStore) DeleteAllBusStations() error {
	_, err := store.db.Exec("DELETE FROM seoul_bus_station")
	return err
}
