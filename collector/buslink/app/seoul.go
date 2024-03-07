package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

type SeoulOpenApiResponse struct {
	MasterRouteNode interface{}
}

type SeoulOpenApiResponseHead struct {
	TotalCount int
	ResultCode OpenAPIResultCode
}

type SeoulOpenApiBusLink struct {
	RouteId       string  `mapstructure:"ROUTE_ID"`
	StationId     string  `mapstructure:"STTN_ID"`
	DistanceMeter float64 `mapstructure:"STTN_DSTNC_MTR"`
	StationOrder  float64 `mapstructure:"STTN_ORD"`
}

func (app *app) InsertSeoulBusLinks(seoulOpenApiBusLinks []SeoulOpenApiBusLink) error {
	for _, link := range seoulOpenApiBusLinks {
		routeId, err := strconv.ParseInt(link.RouteId, 10, 64)
		if err != nil {
			return err
		}

		stationId, err := strconv.ParseInt(link.StationId, 10, 64)
		if err != nil {
			return err
		}

		err = app.seoulStore.CreateBusLinks(
			routeId,
			stationId,
			int(link.DistanceMeter),
			int(link.StationOrder),
		)

		if err != nil {
			return err
		}
	}
	return nil
}

func (app *app) CollectSeoulBusLinks(docType DocType) ([]SeoulOpenApiBusLink, error) {
	var apiError OpenAPIError
	pageIndex := 1
	pageSize := 1000
	var seoulOpenAPIBusLinks []SeoulOpenApiBusLink

	responseForCount, apiError, _ := requestSeoulBusLinks(app.seoulApiKey, docType, 1, 1)
	if (apiError != OpenAPIError{}) {
		errorMessage := fmt.Sprintf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
		return nil, errors.New(errorMessage)
	}
	head := responseForCount.extractHead()
	fmt.Printf("서울 수집 대상 정류장 구간 개수: %d\n", head.TotalCount)

	// TODO: 비동기
	for {
		response, apiError, url := requestSeoulBusLinks(app.seoulApiKey, docType, pageIndex, pageSize)
		if (apiError != OpenAPIError{}) {
			errorMessage := fmt.Sprintf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
			return nil, errors.New(errorMessage)
		}

		head := responseForCount.extractHead()
		if head.ResultCode.Code != "INFO-000" {
			fmt.Printf("[정상 처리되지 않은 응답코드 수신] URL: %s, ResultCode: %#v\n", url, head.ResultCode)
			continue
		}

		// 구간 정보 수집리스트에 추가
		seoulOpenAPIBusLinks = append(seoulOpenAPIBusLinks, response.extractOpenApiBusLink()...)

		pageIndex += 1
		// 페이지 크기와 인덱스를 곱한 값이 전체 데이터의 수보다 커지면 중단
		if pageIndex*pageSize > head.TotalCount {
			break
		}
	}
	fmt.Printf("서울 수집된 정류장 구간 개수: %d\n", len(seoulOpenAPIBusLinks))

	return seoulOpenAPIBusLinks, nil
}

func requestSeoulBusLinks(apiKey string, docType DocType, pageIndex int, pageSize int) (SeoulOpenApiResponse, OpenAPIError, string) {
	var apiError OpenAPIError
	var openAPIFailResponse OpenAPIFailResponse
	var rawOpenApiResponse SeoulOpenApiResponse

	url := fmt.Sprintf("http://openapi.seoul.go.kr:8088/%s/%s/masterRouteNode/%d/%d", apiKey, docType, pageIndex, pageSize)

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
		return SeoulOpenApiResponse{}, apiError, url
	}

	json.Unmarshal(jsonByte, &rawOpenApiResponse)
	return rawOpenApiResponse, apiError, url
}

func (response *SeoulOpenApiResponse) extractHead() SeoulOpenApiResponseHead {
	totalCount := response.MasterRouteNode.(map[string]interface{})["list_total_count"]

	openApiResultCode := response.MasterRouteNode.(map[string]interface{})["RESULT"].(map[string]interface{})
	resultCode := OpenAPIResultCode{
		Code:    openApiResultCode["CODE"].(string),
		Message: openApiResultCode["MESSAGE"].(string),
	}

	head := SeoulOpenApiResponseHead{
		TotalCount: int(totalCount.(float64)),
		ResultCode: resultCode,
	}
	return head
}

func (response *SeoulOpenApiResponse) extractOpenApiBusLink() []SeoulOpenApiBusLink {
	var seoulOpenApiBusLinks []SeoulOpenApiBusLink

	rawOpenApiRowMapList := response.MasterRouteNode.(map[string]interface{})["row"].([]interface{})
	err := mapstructure.Decode(rawOpenApiRowMapList, &seoulOpenApiBusLinks)

	if err != nil {
		panic(err)
	}
	return seoulOpenApiBusLinks
}
