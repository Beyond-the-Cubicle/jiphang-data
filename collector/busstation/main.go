package main

import (
	"flag"
	"fmt"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/app"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

func main() {
	// 배포환경 인자
	var env = flag.String("env", "local", "배포 환경 입력 - ex. local|dev|stg|prd")
	flag.Parse()
	fmt.Printf("env: %v\n", *env)

	// 설정값 로드
	appConfig := config.NewConfig(*env)
	fmt.Printf("appConfig: %+v\n", appConfig)

	standardStore := store.NewStandardStore(appConfig)
	seoulStore := store.NewSeoulStore(appConfig)
	gyunggiStore := store.NewGyunggiStore(appConfig)
	application := app.New(appConfig, standardStore, seoulStore, gyunggiStore)

	// DB 내 데이터 모두 제거
	clearDB(standardStore, seoulStore, gyunggiStore)

	// TODO: go루틴 사용해서 IO 작업 비동기 수행
	collectSeoulBusStations(application)
	collectGyunggiBusStations(application)

	standardStore.Close()
	seoulStore.Close()
	gyunggiStore.Close()
}

func clearDB(standardStore store.StandardStore, seoulStore store.SeoulStore, gyunggiStore store.GyunggiStore) {
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
}

func collectSeoulBusStations(application app.App) {
	fmt.Printf("=============== 서울 버스정류장 데이터 수집 시작 ===============\n")
	seoulOpenApiBusStations, err := application.CollectSeoulBusStations(app.Json)
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

func collectGyunggiBusStations(application app.App) {
	fmt.Printf("=============== 경기 버스정류장 데이터 수집 시작 ===============\n")
	gyunggiOpenApiBusStations, err := application.CollectGyunggiBusStations(app.Json)
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
