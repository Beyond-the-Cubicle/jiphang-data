package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

type GyunggiOpenApiResponse struct {
	TBBMSSTATIONM interface{}
}

type GyunggiOpenApiResponseBody struct {
	BusStations []GyunggiOpenAPIBusStation `json:"row"`
}

type GyunggiOpenAPIBusStation struct {
	StationId               string  `json:"STTN_ID"`            // 정류장 ID
	StationName             string  `json:"STTN_NM"`            // 정류장 이름
	CoordinateY             float64 `json:"Y_CRDNT"`            // Y좌표
	CoordinateX             float64 `json:"X_CRDNT"`            // X좌표
	GpsCoordinateY          float64 `json:"GPS_X_CRDNT"`        // GPS Y좌표
	GpsCoordinateX          float64 `json:"GPS_Y_CRDNT"`        // GPS X좌표
	RinkId                  string  `json:"RINK_ID"`            // 링크아이디
	StationType             string  `json:"STTN_TYPE"`          // 정류장 유형
	TransferStationExtNo    string  `json:"TRANSIT_STTN_EXTNO"` // 환승 정류장 유무 ex. C, N
	MedianBusLaneYn         string  `json:"CNTR_CARTRK_YN"`     // 중앙차로 여부
	StationEnglishName      string  `json:"STTN_ENG_NM"`        // 정류장 영어 이름
	ArsId                   string  `json:"ARS_ID"`             // ARS ID
	InstitutionCode         string  `json:"INST_CD"`            // 기관 코드
	DataDisplayYn           string  `json:"DATA_EXPRS_EXTNO"`   // 데이터 표출 유무(Y/N)
	RegisteredBy            string  `json:"REGIST_ID"`          // 등록 주체 아이디
	RegisteredAt            string  `json:"REGIST_DE"`          // 등록 일시 YYYYMMDDHHmmss
	Memo                    string  `json:"RM"`                 // 비고
	SignPostType            string  `json:"SIGNPOST_TYPE"`      // 표지판 유형
	DongCode                string  `json:"ADMINIST_DONG_CD"`   // 행정동 코드
	RegionCode              string  `json:"VOLM_STATN_CD"`      // 권역 코드
	UseYn                   string  `json:"USE_DIV"`            // 사용구분(Y/N)
	StationChineseName      string  `json:"STTN_CHN_NM"`        // 정류장 중국어 이름
	StationJapaneseName     string  `json:"STTN_JPNLANG_NM"`    // 정류장 일본어 이름
	StationVietnamName      string  `json:"STTN_VIETNAM_NM"`    // 정류장 베트남어 이름
	DrtYn                   string  `json:"DRT_EXTNO"`          // DRT 유무
	StationTypeName         string  `json:"STATION_TP_NM"`      // 정류장 유형 이름(ex. 미지정, 시내)
	TransferStationTypeName string  `json:"CHNG_STATION_YN_NM"` // 환승유무를 나타내는 문자열(ex. 일반, 환승)
	SignPostTypeName        string  `json:"MARK_TYPE_NM"`       // 표지판 유형 이름(ex. 표지판 없음)

}

func (app *app) CollectGyunggiBusStations(apiKey string, docType DocType) ([]GyunggiOpenAPIBusStation, error) {
	return nil, nil
}

func (app *app) ConvertGyunggiBusStationsToStandard(gyunggiOpenApiBusStations []GyunggiOpenAPIBusStation) ([]store.BusStation, error) {
	return nil, nil
}

func requestGyunggiBusStations(apiKey string, docType DocType, pageIndex int, pageSize int) (GyunggiOpenApiResponse, OpenAPIError, string) {
	var apiError OpenAPIError
	var openAPIFailResponse OpenAPIFailResponse
	var rawOpenApiResponse GyunggiOpenApiResponse

	url := fmt.Sprintf("https://openapi.gg.go.kr/TBBMSSTATIONM?KEY=%s&Type=%s&pIndex=%d&pSize=%d", apiKey, docType, pageIndex, pageSize)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	jsonByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 정상 응답이 아닌 케이스
	json.Unmarshal(jsonByte, &openAPIFailResponse)
	if openAPIFailResponse.Result.Code != "" {
		apiError = OpenAPIError{
			Url:    url,
			Result: openAPIFailResponse.Result,
		}
		return GyunggiOpenApiResponse{}, apiError, url
	}

	json.Unmarshal(jsonByte, &rawOpenApiResponse)
	return rawOpenApiResponse, apiError, url
}
