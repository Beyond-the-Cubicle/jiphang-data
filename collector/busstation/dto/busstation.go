package dto

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
