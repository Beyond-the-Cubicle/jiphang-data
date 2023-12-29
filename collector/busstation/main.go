package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/database"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/dto"
	"github.com/mitchellh/mapstructure"
)

func main() {
	database.ConnectDb()
	collectGynggiBusStations()

	database.Db.Close()
}

func collectGynggiBusStations() {

	var apiError dto.OpenAPIError
	pageSize := 1000
	pageIndex := 1
	var openAPIFailResponse dto.OpenAPIFailResponse
	var rawOpenApiResponse dto.RawOpenAPIResponse
	var busStations []dto.BusStation
	for {
		url := fmt.Sprintf("https://openapi.gg.go.kr/BusStation?Key=%s&Type=%s&pIndex=%d&pSize=%d", config.ApiKey, config.DocType, pageIndex, pageSize)
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
			apiError = dto.OpenAPIError{
				Url:    url,
				Result: openAPIFailResponse.Result,
			}
			break
		}

		json.Unmarshal(jsonByte, &rawOpenApiResponse)
		// 정상 응답 케이스 - head 파싱
		head := makeOpenApiResponseHead(rawOpenApiResponse)
		if head.ResultCode.Code != "INFO-000" {
			fmt.Printf("[정상 처리되지 않은 응답코드 수신] URL: %s, ResultCode: %#v\n", url, head.ResultCode)
			continue
		}

		// 정상 응답 케이스 - row 파싱
		busStations = append(busStations, makeBusStations(rawOpenApiResponse)...)

		pageIndex++
	}

	if (len(busStations) == 0 && apiError != dto.OpenAPIError{}) {
		fmt.Printf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
		return
	}

	fmt.Printf("[정상 수집 완료] Collected BusStation Count: %d\n", len(busStations))

	// DB에 저장
	insertBusStations(busStations)
	fmt.Printf("경기도 버스 정류장 정보 DB 저장 완료")
}

func makeOpenApiResponseHead(rawOpenApiResponse dto.RawOpenAPIResponse) dto.OpenAPIResponseHead {
	rawOpenApiHeadMapList := rawOpenApiResponse.BusStation[0].(map[string]interface{})["head"].([]interface{})

	// list_total_count
	totalCount := rawOpenApiHeadMapList[0].(map[string]interface{})["list_total_count"]

	// RESULT - CODE, MESSAGE
	openApiResultCode := rawOpenApiHeadMapList[1].(map[string]interface{})["RESULT"].(map[string]interface{})
	resultCode := dto.OpenAPIResultCode{
		Code:    openApiResultCode["CODE"].(string),
		Message: openApiResultCode["MESSAGE"].(string),
	}

	// api_version
	apiVersion := rawOpenApiHeadMapList[2].(map[string]interface{})["api_version"]

	head := dto.OpenAPIResponseHead{
		TotalCount: int(totalCount.(float64)),
		ResultCode: resultCode,
		ApiVersion: apiVersion.(string),
	}

	return head
}

func makeBusStations(rawOpenApiResponse dto.RawOpenAPIResponse) []dto.BusStation {
	// openApiResponse의 두번째 요소가 row를 포함하는 map 구조로 실제 정류소 정보 리스트를 포함
	rawOpenApiRowMapList := rawOpenApiResponse.BusStation[1].(map[string]interface{})["row"]

	var openApiBusStations []dto.OpenAPIBusStation
	err := mapstructure.Decode(rawOpenApiRowMapList, &openApiBusStations)
	if err != nil {
		panic(err)
	}

	var busStations []dto.BusStation

	for _, openApiBusStation := range openApiBusStations {
		busStations = append(busStations, openApiBusStation.ToBusStation())
	}

	return busStations
}

// TODO: 벌크 insert
func insertBusStations(busStations []dto.BusStation) {
	for _, busstation := range busStations {
		_, err := database.Db.Exec("INSERT INTO bus_station VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			busstation.StationId, busstation.StationName, busstation.EnglishStationName, busstation.ArsId,
			busstation.StationDivisionName, busstation.GovermentOfficeName, busstation.Location,
			busstation.Latitude, busstation.Longitude, busstation.CityCode, busstation.CityName,
		)
		if err != nil {
			panic(err)
		}
	}
}
