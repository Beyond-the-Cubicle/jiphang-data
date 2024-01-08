package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/database"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/dto"
	"github.com/mitchellh/mapstructure"
)

func main() {
	database.ConnectDb()
	deleteAllBusStations()
	// TODO: TAGO가 아닌 서울, 경기만 별도로 하려면, 겹치는 stationId 해결해야함 -> 문제는 stationId 겹칠 때 ARS ID가 다른 경우가 있음
	// collectSeoulBusStations()
	// collectGynggiBusStations()
	collectTagoBusStations()
	database.Db.Close()

}

func collectSeoulBusStations() {
	fmt.Printf("=============== 서울 버스 정류장 수집 시작 ===============\n")
	var apiError dto.OpenAPIError
	startIndex := 1
	endIndex := 1000
	var busStations []dto.BusStation

	// 정류장 전체 카운트 구하기
	responseForCount, apiError, _ := requestSeoulBusStations(1, 1)
	if (responseForCount.TbisMasterStation.List_total_count == 0 && apiError != dto.OpenAPIError{}) {
		fmt.Printf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
		return
	}
	totalCount := responseForCount.TbisMasterStation.List_total_count
	fmt.Printf("서울 수집 대상 버스 정류장 갯수: %d\n", totalCount)

	for {
		response, apiError, url := requestSeoulBusStations(startIndex, endIndex)
		if (response.TbisMasterStation.List_total_count == 0 && apiError != dto.OpenAPIError{}) {
			fmt.Printf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
			return
		}

		if response.TbisMasterStation.Result.Code != "INFO-000" {
			fmt.Printf("[정상 처리되지 않은 응답코드 수신] URL: %s, ResultCode: %#v\n", url, response.TbisMasterStation.Result)
			continue
		}

		for _, seoulOpenAPIBusStation := range response.TbisMasterStation.Row {
			busStations = append(busStations, seoulOpenAPIBusStation.ToBusStation())
		}

		// 전체 카운트보다 현재까지 수행한 endIndex가 크면 완료 처리
		if endIndex > totalCount {
			break
		}

		startIndex += 1000
		endIndex += 1000
	}

	fmt.Printf("[서울 버스 정상 수집 완료] Collected BusStation Count: %d\n", len(busStations))

	// DB에 저장
	insertBusStations(busStations)
	fmt.Printf("서울 버스 정류장 정보 DB 저장 완료\n")
	fmt.Printf("=============== 서울 버스 정류장 수집 종료 ===============\n")
}

func requestSeoulBusStations(startIndex int, endIndex int) (dto.SeoulRawOpenAPIResponse, dto.OpenAPIError, string) {
	var apiError dto.OpenAPIError
	var openAPIFailResponse dto.OpenAPIFailResponse
	var rawOpenApiResponse dto.SeoulRawOpenAPIResponse

	url := fmt.Sprintf("http://openapi.seoul.go.kr:8088/%s/%s/tbisMasterStation/%d/%d", config.SeoulApiKey, config.DocType, startIndex, endIndex)
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
		return dto.SeoulRawOpenAPIResponse{}, apiError, url
	}

	json.Unmarshal(jsonByte, &rawOpenApiResponse)
	return rawOpenApiResponse, apiError, url
}

func collectGynggiBusStations() {
	fmt.Printf("=============== 경기도 버스 정류장 수집 시작 ===============\n")
	var apiError dto.OpenAPIError
	pageSize := 1000
	pageIndex := 1
	var openAPIFailResponse dto.OpenAPIFailResponse
	var rawOpenApiResponse dto.GyunggiRawOpenAPIResponse
	var busStations []dto.BusStation
	for {
		url := fmt.Sprintf("https://openapi.gg.go.kr/BusStation?Key=%s&Type=%s&pIndex=%d&pSize=%d", config.GyunggiApiKey, config.DocType, pageIndex, pageSize)
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

	fmt.Printf("[경기도 버스 정상 수집 완료] Collected BusStation Count: %d\n", len(busStations))

	// DB에 저장
	insertBusStations(busStations)
	fmt.Printf("경기도 버스 정류장 정보 DB 저장 완료\n")
	fmt.Printf("=============== 경기도 버스 정류장 수집 종료 ===============\n")
}

func makeOpenApiResponseHead(rawOpenApiResponse dto.GyunggiRawOpenAPIResponse) dto.OpenAPIResponseHead {
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

func makeBusStations(rawOpenApiResponse dto.GyunggiRawOpenAPIResponse) []dto.BusStation {
	// openApiResponse의 두번째 요소가 row를 포함하는 map 구조로 실제 정류소 정보 리스트를 포함
	rawOpenApiRowMapList := rawOpenApiResponse.BusStation[1].(map[string]interface{})["row"]

	var openApiBusStations []dto.GyunggiOpenAPIBusStation
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

func deleteAllBusStations() {
	fmt.Printf("=============== 버스 정류장 전체 데이터 제거 시작 ===============\n")
	_, err := database.Db.Exec("DELETE FROM bus_station")
	if err != nil {
		panic(err)
	}
	fmt.Printf("=============== 버스 정류장 전체 데이터 제거 종료 ===============\n")
}

func collectTagoBusStations() {
	fmt.Printf("=============== TAGO 버스 정류장 수집 시작 ===============\n")
	var apiError dto.OpenAPIError
	pageIndex := 1
	pageSize := 5000
	var busStations []dto.BusStation

	// 정류장 전체 카운트 구하기
	responseForCount, apiError := requestTagoBusStations(1, 1)
	if (responseForCount.TotalCount == 0 && apiError != dto.OpenAPIError{}) {
		fmt.Printf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
		return
	}
	totalCount := responseForCount.TotalCount
	fmt.Printf("TAGO 수집 대상 버스 정류장 갯수: %d\n", totalCount)

	for {
		response, apiError := requestTagoBusStations(pageIndex, pageSize)
		if (response.TotalCount == 0 && apiError != dto.OpenAPIError{}) {
			fmt.Printf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
			return
		}

		for _, tagoOpenAPIBusStation := range response.Data {
			busStations = append(busStations, tagoOpenAPIBusStation.ToBusStation())
		}

		fmt.Printf("pageIndex: %d\n", pageIndex)

		// 전체 카운트만큼 수집되면 중지
		if len(busStations) >= totalCount {
			break
		}

		pageIndex += 1
	}

	fmt.Printf("[TAGO 버스 정상 수집 완료] Collected BusStation Count: %d\n", len(busStations))

	// DB에 저장
	insertBusStations(busStations)
	fmt.Printf("TAGO 버스 정류장 정보 DB 저장 완료\n")
	fmt.Printf("=============== TAGO 버스 정류장 수집 종료 ===============\n")
}

func requestTagoBusStations(pageIndex int, pageSize int) (dto.TagoOpenAPIResponse, dto.OpenAPIError) {
	var apiError dto.OpenAPIError
	var tagoFailResponse dto.TagoFailResponse
	var tagoOpenAPIResponse dto.TagoOpenAPIResponse

	url := fmt.Sprintf("https://api.odcloud.kr/api/15067528/v1/uddi:eb02ec03-6edd-4cb0-88b8-eda22ca55e80?page=%d&perPage=%d&serviceKey=%s", pageIndex, pageSize, config.TagoApiKey)
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
	json.Unmarshal(jsonByte, &tagoFailResponse)
	if tagoFailResponse.Code != 0 {
		apiError = dto.OpenAPIError{
			Url: url,
			Result: dto.OpenAPIResultCode{
				Code:    strconv.Itoa(tagoFailResponse.Code),
				Message: tagoFailResponse.Msg,
			},
		}
		return dto.TagoOpenAPIResponse{}, apiError
	}

	json.Unmarshal(jsonByte, &tagoOpenAPIResponse)
	return tagoOpenAPIResponse, apiError
}
