package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

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

func (openApiBusStation *SeoulOpenAPIBusStation) ToBusStation() store.BusStation {
	return store.BusStation{
		StationName: openApiBusStation.STTN_NM,
		StationId:   openApiBusStation.STTN_ID,
		ArsId:       openApiBusStation.STTN_NO,
		Latitude:    openApiBusStation.CRDNT_Y,
		Longitude:   openApiBusStation.CRDNT_X,
	}
}

func (app *app) CollectSeoulBusStations(apiKey string, docType DocType) ([]SeoulOpenAPIBusStation, error) {
	var apiError OpenAPIError
	startIndex := 1
	endIndex := 1000
	var seoulOpenApiBusStations []SeoulOpenAPIBusStation

	// 정류장 전체 카운트 구하기
	responseForCount, apiError, _ := requestSeoulBusStations(apiKey, docType, 1, 1)
	if (responseForCount.TbisMasterStation.List_total_count == 0 && apiError != OpenAPIError{}) {
		errorMessage := fmt.Sprintf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
		return nil, errors.New(errorMessage)
	}
	totalCount := responseForCount.TbisMasterStation.List_total_count
	fmt.Printf("서울 수집 대상 버스정류장 개수: %d\n", totalCount)

	for {
		response, apiError, url := requestSeoulBusStations(apiKey, docType, startIndex, endIndex)
		if (response.TbisMasterStation.List_total_count == 0 && apiError != OpenAPIError{}) {
			errorMessage := fmt.Sprintf("[에러 응답 수신] URL: %s, code: %s, message: %s\n", apiError.Url, apiError.Result.Code, apiError.Result.Message)
			return seoulOpenApiBusStations, errors.New(errorMessage)
		}

		if response.TbisMasterStation.Result.Code != "INFO-000" {
			fmt.Printf("[정상 처리되지 않은 응답코드 수신] URL: %s, ResultCode: %#v\n", url, response.TbisMasterStation.Result)
			continue
		}

		// 응답으로 받은 정류장 정보 수집리스트에 추가
		seoulOpenApiBusStations = append(seoulOpenApiBusStations, response.TbisMasterStation.Row...)

		// 전체 카운트보다 현재까지 수행한 endIndex가 크면 완료 처리
		if endIndex > totalCount {
			break
		}

		startIndex += 1000
		endIndex += 1000
	}

	fmt.Printf("서울 수집된 버스정류장 개수: %d\n", len(seoulOpenApiBusStations))

	return seoulOpenApiBusStations, nil
}

func (app *app) ConvertSeoulBusStationsToStandard(seoulOpenApiBusStations []SeoulOpenAPIBusStation) ([]store.BusStation, error) {
	var busStations []store.BusStation
	for _, seoulOpenApiBusStation := range seoulOpenApiBusStations {
		busStations = append(busStations, seoulOpenApiBusStation.ToBusStation())
	}
	return busStations, nil
}

func (app *app) InsertBusStations(busStations []store.BusStation) error {
	for _, busstation := range busStations {
		err := app.store.CreateBusStations(
			busstation.StationId,
			busstation.StationName,
			busstation.ArsId,
			busstation.Latitude,
			busstation.Longitude,
			busstation.CityCode,
			busstation.CityName,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func requestSeoulBusStations(apiKey string, docType DocType, startIndex int, endIndex int) (SeoulRawOpenAPIResponse, OpenAPIError, string) {
	var apiError OpenAPIError
	var openAPIFailResponse OpenAPIFailResponse
	var rawOpenApiResponse SeoulRawOpenAPIResponse

	url := fmt.Sprintf("http://openapi.seoul.go.kr:8088/%s/%s/tbisMasterStation/%d/%d", apiKey, docType, startIndex, endIndex)
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
		return SeoulRawOpenAPIResponse{}, apiError, url
	}

	json.Unmarshal(jsonByte, &rawOpenApiResponse)
	return rawOpenApiResponse, apiError, url
}
