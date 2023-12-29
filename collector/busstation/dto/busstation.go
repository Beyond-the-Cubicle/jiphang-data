package dto

import "strconv"

type RawOpenAPIResponse struct {
	BusStation []interface{}
}

type OpenAPIBusStation struct {
	SIGUN_NM            string // 시군명
	SIGUN_CD            string // 시군코드
	STATION_NM_INFO     string // 정류소명
	ENG_STATION_NM_INFO string // 정류소영문명
	STATION_ID          string // 정류소ID
	STATION_MANAGE_NO   string // 정류소번호
	STATION_DIV_NM      string // 중앙차로여부 ex. 노변정류장
	JURISD_INST_NM      string // 관할관청 ex. 경기도 안양시
	LOCPLC_LOC          string // 위치 (주소정보) ex. 경기도 안양시 만안구 석수동
	WGS84_LAT           string // WGS84 위도
	WGS84_LOGT          string // WGS84 경도
}

func (openApiBusStation *OpenAPIBusStation) ToBusStation() BusStation {
	latitude, _ := strconv.ParseFloat(openApiBusStation.WGS84_LAT, 64)
	longitude, _ := strconv.ParseFloat(openApiBusStation.WGS84_LOGT, 64)

	return BusStation{
		CityName:            openApiBusStation.SIGUN_NM,
		CityCode:            openApiBusStation.SIGUN_CD,
		StationName:         openApiBusStation.STATION_NM_INFO,
		EnglishStationName:  openApiBusStation.ENG_STATION_NM_INFO,
		StationId:           openApiBusStation.STATION_ID,
		ArsId:               openApiBusStation.STATION_MANAGE_NO,
		StationDivisionName: openApiBusStation.STATION_DIV_NM,
		GovermentOfficeName: openApiBusStation.JURISD_INST_NM,
		Location:            openApiBusStation.LOCPLC_LOC,
		Latitude:            latitude,
		Longitude:           longitude,
	}
}

type OpenAPIResponseHead struct {
	TotalCount int
	ResultCode OpenAPIResultCode
	ApiVersion string
}

type OpenAPIFailResponse struct {
	Result OpenAPIResultCode
}

type OpenAPIError struct {
	Url    string
	Result OpenAPIResultCode
}

type OpenAPIResultCode struct {
	Code    string
	Message string
}

type BusStation struct {
	StationId           string
	StationName         string
	EnglishStationName  string
	ArsId               string
	StationDivisionName string
	GovermentOfficeName string
	Location            string
	Latitude            float64
	Longitude           float64
	CityCode            string
	CityName            string
}
