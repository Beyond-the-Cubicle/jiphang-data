package main

import (
	"fmt"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/app"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/busstation/store"
)

func main() {
	store := store.New()
	application := app.New(store)

	fmt.Printf("=============== 버스정류장 DB 데이터 초기화 시작 ===============\n")
	err := store.DeleteAllBusStations()
	if err != nil {
		panic("버스정류장 데이터 초기화 실패 - " + err.Error())
	}
	fmt.Printf("=============== 버스정류장 DB 데이터 초기화 완료 ===============\n")

	fmt.Printf("=============== 서울 버스정류장 데이터 수집 시작 ===============\n")
	seoulOpenApiBusStations, err := application.CollectSeoulBusStations(config.SeoulApiKey, app.Json)
	if err != nil {
		panic("서울 버스 정류장 수집 중 오류 발생 - " + err.Error())
	}
	fmt.Printf("=============== 서울 버스정류장 데이터 수집 완료 ===============\n")

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

	store.Close()
}
