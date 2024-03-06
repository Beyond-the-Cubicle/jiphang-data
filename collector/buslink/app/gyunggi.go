package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type GyunggiOpenApiResponse struct {
	TBBMSROUTESTATIONM []interface{}
}

type GyunggiOpenApiResponseHead struct {
	TotalCount int
	ResultCode OpenAPIResultCode
	ApiVersion string
}
type GyunggiOpenApiBusLink struct {
	RouteId                  string `mapstructure:"ROUTE_ID"`                    // 노선 아이디
	StationOrder             int    `mapstructure:"STTN_ORDER"`                  // 정류장 순서번호
	StationId                string `mapstructure:"STTN_ID"`                     // 정류장 아이디
	GisDistance              int    `mapstructure:"GIS_DSTN"`                    // GIS 거리
	AccumulatedDistance      int    `mapstructure:"ACCMLT_DSTN"`                 // 누적거리
	RealDistance             int    `mapstructure:"REAL_DSTN"`                   // 실제거리
	DecidedDistance          int    `mapstructure:"DCSN_DSTN"`                   // 확정거리
	ProgressDivisionCode     string `mapstructure:"PROGRS_DIV_CD"`               // 진행구분코드
	RegisteredBy             string `mapstructure:"REGIST_ID"`                   // 등록아이디
	RegisteredAt             string `mapstructure:"REGIST_DE"`                   // 등록일자
	UseDivision              string `mapstructure:"USE_DIV"`                     // 사용구분
	IsRegionalLine           string `mapstructure:"UNWEL_HNO_STATN_ROUTE_EXTNO"` // 벽지노선 유무
	ProgressDivisionCodeName string `mapstructure:"PROGRS_DIV_CD_NM"`            // 진행구분코드명
	UseDivisionName          string `mapstructure:"USE_DIV_NM"`                  // 벽지노선 유무명
}

func (app *app) CollectGyunggiBusLinks(docType DocType) ([]GyunggiOpenApiBusLink, error) {
	var apiError OpenAPIError
	pageIndex := 1
	pageSize := 1000
	var gyunggiOpenAPIBusStations []GyunggiOpenApiBusLink

	responseForCount, apiError, _ := requestGyunggiBusLinks(app.gyunggiApiKey, docType, 1, 1)
	if (apiError != OpenAPIError{}) {
		errorMessage := fmt.Sprintf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
		return nil, errors.New(errorMessage)
	}
	headForTotalCount := responseForCount.extractGyunggiOpenApiResponseHead()
	fmt.Printf("경기도 수집 대상 정류장 구간 개수: %d\n", headForTotalCount.TotalCount)

	for {
		response, apiError, url := requestGyunggiBusLinks(app.gyunggiApiKey, docType, pageIndex, pageSize)
		if (apiError != OpenAPIError{}) {
			errorMessage := fmt.Sprintf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
			return nil, errors.New(errorMessage)
		}

		head := responseForCount.extractGyunggiOpenApiResponseHead()
		if head.ResultCode.Code != "INFO-000" {
			fmt.Printf("[정상 처리되지 않은 응답코드 수신] URL: %s, ResultCode: %#v\n", url, head.ResultCode)
			continue
		}

		// 정류장 정보 수집리스트에 추가
		gyunggiOpenAPIBusStations = append(gyunggiOpenAPIBusStations, response.extractGyunggiOpenAPIBusLinks()...)

		pageIndex += 1
		// 페이지 크기와 인덱스를 곱한 값이 전체 데이터 수보다 커지면 중단
		if pageIndex*pageSize > headForTotalCount.TotalCount {
			break
		}
	}

	fmt.Printf("경기도 수집된 정류장 구간 개수: %d\n", len(gyunggiOpenAPIBusStations))

	return gyunggiOpenAPIBusStations, nil
}

func requestGyunggiBusLinks(apiKey string, docType DocType, pageIndex int, pageSize int) (GyunggiOpenApiResponse, OpenAPIError, string) {
	var apiError OpenAPIError
	var openAPIFailResponse OpenAPIFailResponse
	var rawOpenApiResponse GyunggiOpenApiResponse

	url := fmt.Sprintf("https://openapi.gg.go.kr/TBBMSROUTESTATIONM?key=%s&type=%s&pIndex=%d&pSize=%d", apiKey, docType, pageIndex, pageSize)
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

func (response *GyunggiOpenApiResponse) extractGyunggiOpenApiResponseHead() GyunggiOpenApiResponseHead {
	rawOpenApiHeadMapList := response.TBBMSROUTESTATIONM[0].(map[string]interface{})["head"].([]interface{})
	totalCount := rawOpenApiHeadMapList[0].(map[string]interface{})["list_total_count"]

	openApiResultcode := rawOpenApiHeadMapList[1].(map[string]interface{})["RESULT"].(map[string]interface{})
	resultCode := OpenAPIResultCode{
		Code:    openApiResultcode["CODE"].(string),
		Message: openApiResultcode["MESSAGE"].(string),
	}

	apiVersion := rawOpenApiHeadMapList[2].(map[string]interface{})["api_version"].(string)

	head := GyunggiOpenApiResponseHead{
		TotalCount: int(totalCount.(float64)),
		ResultCode: resultCode,
		ApiVersion: apiVersion,
	}
	return head
}

func (response *GyunggiOpenApiResponse) extractGyunggiOpenAPIBusLinks() []GyunggiOpenApiBusLink {
	var gyunggiOpenAPIBusStations []GyunggiOpenApiBusLink

	rawOpenApiRowMapList := response.TBBMSROUTESTATIONM[1].(map[string]interface{})["row"].([]interface{})
	err := mapstructure.Decode(rawOpenApiRowMapList, &gyunggiOpenAPIBusStations)

	if err != nil {
		panic(err)
	}
	return gyunggiOpenAPIBusStations
}
