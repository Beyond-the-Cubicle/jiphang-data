package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/database"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/dto"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

func main() {
	store := store.New()

	database.ConnectDb()
	err := store.DeleteAllBusStations()
	if err != nil {
		panic("delete all bus stations failed - " + err.Error())
	}
	collectSeoulBusStations(store)
	store.Close()
}

func collectSeoulBusStations(store store.Store) {
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
	insertBusStations(store, busStations)
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

// TODO: 벌크 insert
func insertBusStations(store store.Store, busStations []dto.BusStation) {
	for _, busstation := range busStations {
		err := store.CreateBusStations(
			busstation.StationId,
			busstation.StationName,
			busstation.ArsId,
			busstation.Latitude,
			busstation.Longitude,
			busstation.CityCode,
			busstation.CityName,
		)
		if err != nil {
			panic(err)
		}
	}
}
