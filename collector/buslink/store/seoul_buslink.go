package store

// 버스 구간정보 DB 테이블 매핑용 구조
type SeoulBusLink struct {
	RouteId              int64 // 노선 ID
	StationId            int64 // 정류장 ID
	StationDistanceMeter int   // 링크 구간거리
	StationOrder         int   // 정류장 순서 번호
}

// 버스 구간정보 생성
func (seoulStore *seoulStore) CreateBusLinks(routeId, stationId int64, stationDistanceMeter, stationOrder int) error {
	_, err := seoulStore.db.Exec(`INSERT INTO seoul_bus_link VALUES (?, ?, ?, ?)`,
		routeId, stationId, stationDistanceMeter, stationOrder)
	return err
}

// 버스 구간정보 전체 삭제
func (seoulStore *seoulStore) DeleteAllBusLinks() error {
	_, err := seoulStore.db.Exec("DELETE FROM seoul_bus_link")
	return err
}
