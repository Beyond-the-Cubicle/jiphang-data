package store

import "strings"

type GyunggiBusStation struct {
	StationId               string  // 정류장 ID
	StationName             string  // 정류장 이름
	CoordinateY             float64 // Y좌표
	CoordinateX             float64 // X좌표
	GpsCoordinateY          float64 // GPS Y좌표
	GpsCoordinateX          float64 // GPS X좌표
	RinkId                  string  // 링크아이디
	StationType             string  // 정류장 유형
	TransferStationExtNo    string  // 환승 정류장 유무 ex. C, N
	MedianBusLaneYn         string  // 중앙차로 여부
	StationEnglishName      string  // 정류장 영어 이름
	ArsId                   string  // ARS ID
	InstitutionCode         string  // 기관 코드
	DataDisplayYn           string  // 데이터 표출 유무(Y/N)
	RegisteredBy            string  // 등록 주체 아이디
	RegisteredAt            string  // 등록 일시 YYYYMMDDHHmmss
	Memo                    string  // 비고
	SignPostType            string  // 표지판 유형
	DongCode                string  // 행정동 코드
	RegionCode              string  // 권역 코드
	UseYn                   string  // 사용구분(Y/N)
	StationChineseName      string  // 정류장 중국어 이름
	StationJapaneseName     string  // 정류장 일본어 이름
	StationVietnamName      string  // 정류장 베트남어 이름
	DrtYn                   string  // DRT 유무
	StationTypeName         string  // 정류장 유형 이름(ex. 미지정, 시내)
	TransferStationTypeName string  // 환승유무를 나타내는 문자열(ex. 일반, 환승)
	SignPostTypeName        string  // 표지판 유형 이름(ex. 표지판 없음)
}

func (store *gyunggiStore) CreateBusStations(
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
) error {
	_, err := store.db.Exec(`INSERT INTO gyunggi_bus_station VALUES (
		?, ?, ?, ?, ?,
		?, ?, ?, ?, ?,
		?, ?, ?, ?, ?,
		?, ?, ?, ?, ?,
		?, ?, ?, ?, ?,
		?, ?, ?
		)`,
		stationId,
		stationName,
		coordinateX,
		coordinateY,
		gpsCoordinateX,
		gpsCoordinateY,
		rinkId,
		stationType,
		transferStationExtNo,
		medianBusLaneYn,
		stationEnglishName,
		arsId,
		institutionCode,
		dataDisplayYn,
		registeredBy,
		registeredAt,
		memo,
		signPostType,
		dongCode,
		regionCode,
		useYn,
		stationChineseName,
		stationJapaneseName,
		stationVietnamName,
		drtYn,
		stationTypeName,
		transferStationTypeName,
		signPostTypeName,
	)
	return err
}

func (store *gyunggiStore) ReadBusStation(stationId string) (GyunggiBusStation, error) {
	result := GyunggiBusStation{}
	err := store.db.QueryRow("SELECT * FROM gyunggi_bus_station WHERE stationId = ?", stationId).Scan(
		&result.StationId,
		&result.StationName,
		&result.CoordinateX,
		&result.CoordinateY,
		&result.GpsCoordinateX,
		&result.GpsCoordinateY,
		&result.RinkId,
		&result.StationType,
		&result.TransferStationExtNo,
		&result.MedianBusLaneYn,
		&result.StationEnglishName,
		&result.ArsId,
		&result.InstitutionCode,
		&result.DataDisplayYn,
		&result.RegisteredBy,
		&result.RegisteredAt,
		&result.Memo,
		&result.SignPostType,
		&result.DongCode,
		&result.RegionCode,
		&result.UseYn,
		&result.StationChineseName,
		&result.StationJapaneseName,
		&result.StationVietnamName,
		&result.DrtYn,
		&result.StationTypeName,
		&result.TransferStationTypeName,
		&result.SignPostTypeName,
	)
	return result, err
}

func (store *gyunggiStore) ReadBusStations(stationIds []string) ([]GyunggiBusStation, error) {
	results := []GyunggiBusStation{}
	rows, err := store.db.Query("SELECT * FROM gyunggi_bus_station WHERE stationId in (?" + strings.Repeat(",?", len(stationIds)-1) + ")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := GyunggiBusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.CoordinateX,
			&result.CoordinateY,
			&result.GpsCoordinateX,
			&result.GpsCoordinateY,
			&result.RinkId,
			&result.StationType,
			&result.TransferStationExtNo,
			&result.MedianBusLaneYn,
			&result.StationEnglishName,
			&result.ArsId,
			&result.InstitutionCode,
			&result.DataDisplayYn,
			&result.RegisteredBy,
			&result.RegisteredAt,
			&result.Memo,
			&result.SignPostType,
			&result.DongCode,
			&result.RegionCode,
			&result.UseYn,
			&result.StationChineseName,
			&result.StationJapaneseName,
			&result.StationVietnamName,
			&result.DrtYn,
			&result.StationTypeName,
			&result.TransferStationTypeName,
			&result.SignPostTypeName,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *gyunggiStore) ReadAllBusStations() ([]GyunggiBusStation, error) {
	results := []GyunggiBusStation{}
	rows, err := store.db.Query("SELECT * FROM gyunggi_bus_station")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		result := GyunggiBusStation{}
		err := rows.Scan(
			&result.StationId,
			&result.StationName,
			&result.CoordinateX,
			&result.CoordinateY,
			&result.GpsCoordinateX,
			&result.GpsCoordinateY,
			&result.RinkId,
			&result.StationType,
			&result.TransferStationExtNo,
			&result.MedianBusLaneYn,
			&result.StationEnglishName,
			&result.ArsId,
			&result.InstitutionCode,
			&result.DataDisplayYn,
			&result.RegisteredBy,
			&result.RegisteredAt,
			&result.Memo,
			&result.SignPostType,
			&result.DongCode,
			&result.RegionCode,
			&result.UseYn,
			&result.StationChineseName,
			&result.StationJapaneseName,
			&result.StationVietnamName,
			&result.DrtYn,
			&result.StationTypeName,
			&result.TransferStationTypeName,
			&result.SignPostTypeName,
		)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, err
}

func (store *gyunggiStore) DeleteAllBusStations() error {
	_, err := store.db.Exec("DELETE FROM gyunggi_bus_station")
	return err
}
