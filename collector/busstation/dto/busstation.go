package dto

import "strconv"

type GyunggiRawOpenAPIResponse struct {
	BusStation []interface{}
}

type GyunggiOpenAPIBusStation struct {
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

func (openApiBusStation *GyunggiOpenAPIBusStation) ToBusStation() BusStation {
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

type SeoulRawOpenAPIResponse struct {
	TbisMasterStation TbisMasterStation
}

type TbisMasterStation struct {
	List_total_count int
	Result           OpenAPIResultCode
	Row              []SeoulOpenAPIBusStation
}

type SeoulOpenAPIBusStation struct {
	STTN_ID               string  // 정류장 ID
	STTN_NM               string  // 정류장 명칭
	STTN_TYPE             string  // 정류장 유형 (중앙차로 여부 (일반차로, 중앙차로 ...))
	STTN_NO               string  // 정류장 번호(ARS ID)
	CRDNT_X               float64 // 경도
	CRDNT_Y               float64 // 위도
	BUSINFO_FCLT_INSTL_YN string  // 버스도착정보아낸기 설치 여부
}

func (openApiBusStation *SeoulOpenAPIBusStation) ToBusStation() BusStation {
	return BusStation{
		StationName: openApiBusStation.STTN_NM,
		StationId:   openApiBusStation.STTN_ID,
		ArsId:       openApiBusStation.STTN_NO,
		Latitude:    openApiBusStation.CRDNT_Y,
		Longitude:   openApiBusStation.CRDNT_X,
	}
}

type TagoOpenAPIResponse struct {
	CurrentCount int
	Page         int
	PerPage      int
	TotalCount   int
	MatchCount   int
	Data         []TagoOpenAPIBusStation
}

type TagoOpenAPIBusStation struct {
	StationId         string `json:"정류장번호"`
	StationName       string `json:"정류장명"`
	Latitude          string `json:"위도"`
	Longitude         string `json:"경도"`
	CollectedDate     string `json:"정보수집일"`
	MobileShortNumber int    `json:"모바일단축번호"`
	CityCode          int    `json:"도시코드"`
	CityName          string `json:"도시명"`
	ManagerCityName   string `json:"관리도시명"`
}

func (openApiBusStation *TagoOpenAPIBusStation) ToBusStation() BusStation {
	latitude, _ := strconv.ParseFloat(openApiBusStation.Latitude, 64)
	longitude, _ := strconv.ParseFloat(openApiBusStation.Longitude, 64)

	return BusStation{
		CityName:            openApiBusStation.CityName,
		CityCode:            strconv.Itoa(openApiBusStation.CityCode),
		StationName:         openApiBusStation.StationName,
		StationId:           openApiBusStation.StationId,
		ArsId:               strconv.Itoa(openApiBusStation.MobileShortNumber),
		GovermentOfficeName: openApiBusStation.ManagerCityName,
		Latitude:            latitude,
		Longitude:           longitude,
	}
}

type TagoFailResponse struct {
	Code int
	Msg  string
}
