package store

// 버스 구간정보 DB 테이블 매핑용 구조
type GyunggiBusLink struct {
	RouteId             int64  // 노선 ID
	StationOrder        int    // 정류장 순서 번호
	StationId           int64  // 정류장 ID
	GisDistance         int    // GIS 거리
	AccumulatedDistance int    // 누적거리
	RealDistance        int    // 실제거리
	DecidedDistance     int    // 확정거리
	ProgressCode        string // 진행구분코드
	RegisteredBy        string // 등록아이디
	RegisteredAt        string // 등록일자
	UseDivision         string // 사용구분
	RegionalLineYn      string // 벽지노선 유무
	ProgressCodeName    string // 진행구분코드명
	UseDivisionName     string // 벽지노선 유무명
}

// 버스 구간정보 생성
func (gyunggiStore *gyunggiStore) CreateBusLinks(
	routeId int64,
	stationOrder int,
	stationId int64,
	gisDistance, accumulatedDistance, realDistance, decidedDistance int,
	progressCode, registeredBy, registeredAt, useDivision, regionalLineYn, progressCodeName, useDivisionName string,
) error {
	_, err := gyunggiStore.db.Exec(`INSERT INTO gyunggi_bus_link VALUES (
		?, ?, ?, ?, ?,
		?, ?, ?, ?, ?,
		?, ?, ?, ?)`,
		routeId, stationOrder, stationId,
		gisDistance, accumulatedDistance, realDistance, decidedDistance,
		progressCode, registeredBy, registeredAt, useDivision, regionalLineYn, progressCodeName, useDivisionName,
	)
	return err
}

// 버스 구간정보 전체 삭제
func (gyunggiStore *gyunggiStore) DeleteAllBusLinks() error {
	_, err := gyunggiStore.db.Exec("DELETE FROM gyunggi_bus_link")
	return err
}
