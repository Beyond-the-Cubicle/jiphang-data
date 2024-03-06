package main

import (
	"flag"
	"fmt"

	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/app"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/config"
	"github.com/Beyond-the-Cubicle/cgp-data/collector/buslink/store"
)

func main() {
	// 배포환경 인자
	var env = flag.String("env", "local", "배포 환경 입력 - ex. local|dev|stg|prd")
	flag.Parse()
	fmt.Printf("env: %v\n", *env)

	// 설정값 로드
	appConfig := config.NewConfig(*env)
	fmt.Printf("appConfig: %+v\n", appConfig)

	seoulStore := store.NewSeoulStore(appConfig)
	gyunggiStore := store.NewGyunggiStore(appConfig)
	application := app.New(appConfig, seoulStore, gyunggiStore)

	seoulBusLinks, _ := application.CollectSeoulBusLinks(app.Json)
	fmt.Printf("seoul link count: %d", len(seoulBusLinks))
	gyunggiBusLinks, _ := application.CollectGyunggiBusLinks(app.Json)
	fmt.Printf("gyunggi link count: %d", len(gyunggiBusLinks))
}
