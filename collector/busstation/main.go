package main

import (
	"fmt"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/app"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

func main() {
	// TODO: 외부 config 파일 로드해서 사용하도록 수정
	var seoulApiKey = "65674a58626f776c36356c51774d6a"
	var gyunggiApiKey = "39ddaa503de8488995343515399f539e"

	standardStore := store.NewStandardStore()
	seoulStore := store.NewSeoulStore()
	gyunggiStore := store.NewGyunggiStore()
	application := app.New(standardStore, seoulStore, gyunggiStore)

	fmt.Printf("=============== 버스정류장 DB 데이터 초기화 시작 ===============\n")
	err := standardStore.DeleteAllBusStations()
	if err != nil {
		panic("표준 버스정류장 데이터 초기화 실패 - " + err.Error())
	}
	err = seoulStore.DeleteAllBusStations()
	if err != nil {
		panic("서울 버스정류장 데이터 초기화 실패 - " + err.Error())
	}
	err = gyunggiStore.DeleteAllBusStations()
	if err != nil {
		panic("경기 버스정류장 데이터 초기화 실패 - " + err.Error())
	}
	fmt.Printf("=============== 버스정류장 DB 데이터 초기화 완료 ===============\n")

	collectSeoulBusStations(application, seoulApiKey)
	collectGyunggiBusStations(application, gyunggiApiKey)

	standardStore.Close()
	seoulStore.Close()
	gyunggiStore.Close()
}

func collectSeoulBusStations(application app.App, seoulApiKey string) {
	fmt.Printf("=============== 서울 버스정류장 데이터 수집 시작 ===============\n")
	seoulOpenApiBusStations, err := application.CollectSeoulBusStations(seoulApiKey, app.Json)
	if err != nil {
		panic("서울 버스 정류장 수집 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 서울 버스정류장 데이터 수집 완료 ===============\n")

	fmt.Printf("=============== 서울 버스정류장 데이터 저장 시작 ===============\n")
	err = application.InsertSeoulBusStations(seoulOpenApiBusStations)
	if err != nil {
		panic("서울 버스정류장 데이터 저장 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 서울 버스정류장 데이터 저장 완료 ===============\n")

	fmt.Printf("=============== 서울 버스정류장 데이터 표준 데이터로 포멧 전환 시작 ===============\n")
	busStations, err := application.ConvertSeoulBusStationsToStandard(seoulOpenApiBusStations)
	if err != nil {
		panic("서울 버스정류장 데이터 표준 데이터로 포멧 전환 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 서울 버스정류장 데이터 표준 데이터로 포멧 전환 완료 ===============\n")

	fmt.Printf("=============== 서울 버스정류장 데이터 DB 저장 시작 ===============\n")
	err = application.InsertBusStations(busStations)
	if err != nil {
		panic("서울 버스정류장 데이터 DB 저장 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 서울 버스정류장 데이터 DB 저장 완료 ===============\n")
}

func collectGyunggiBusStations(application app.App, gyunggiApiKey string) {
	fmt.Printf("=============== 경기 버스정류장 데이터 수집 시작 ===============\n")
	gyunggiOpenApiBusStations, err := application.CollectGyunggiBusStations(gyunggiApiKey, app.Json)
	if err != nil {
		panic("경기 버스 정류장 수집 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 경기 버스정류장 데이터 수집 완료 ===============\n")

	fmt.Printf("=============== 경기 버스정류장 데이터 저장 시작 ===============\n")
	err = application.InsertGyunggiBusStations(gyunggiOpenApiBusStations)
	if err != nil {
		panic("경기 버스정류장 데이터 저장 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 경기 버스정류장 데이터 저장 완료 ===============\n")

	fmt.Printf("=============== 경기 버스정류장 데이터 표준 데이터로 포멧 전환 시작 ===============\n")
	busStations, err := application.ConvertGyunggiBusStationsToStandard(gyunggiOpenApiBusStations)
	if err != nil {
		panic("경기 버스정류장 데이터 표준 데이터로 포멧 전환 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 경기 버스정류장 데이터 표준 데이터로 포멧 전환 완료 ===============\n")

	fmt.Printf("=============== 경기 버스정류장 데이터 DB 저장 시작 ===============\n")
	err = application.InsertBusStations(busStations)
	if err != nil {
		panic("경기 버스정류장 데이터 DB 저장 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 경기 버스정류장 데이터 DB 저장 완료 ===============\n")
}
