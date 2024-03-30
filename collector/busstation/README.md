# 개요

공공데이터 API를 이용해서 서울, 경기도 버스정류장 정보를 DB에 적재하고, 
미리 정의한 어플리케이션 표준 버스정류장 데이터 구조로 변환하여 DB에 적재하는 기능 수행

사용 API
- 서울: https://data.seoul.go.kr/dataList/OA-21231/S/1/datasetView.do
- 경기: https://data.gg.go.kr/portal/data/service/selectServicePage.do?page=1&rows=10&sortColumn=&sortDirection=&infId=LSQV0RTU9NXPA8RCZLV933248158&infSeq=2&order=&loc=&searchWord=%EB%B2%84%EC%8A%A4

표준 데이터 변환 시 데이터 보정
- 경기도 데이터의 경우 bessel(EPSG:5174) 좌표계 좌표를 WGS84 좌표계 좌표로 변경하여 저장
- 경기도 데이터에 포함된 서울 소재 정류장들은 제거 (서울 open api 데이터와 중복 방지)

<br>

# 어플리케이션 구조
TBD

<br>

# 데이터 구조
일단 resource/database_ddl.sql 파일 참고해주세요.

<br>

# 실행 방법

## docker build & run
[build]<br>
docker buildx build --platform linux/amd64 --tag busstation-collector:0.0.1 .

[run]<br>
docker run -p 3306:3306 busstation-collector:0.0.1

## mac local 환경
mac 로컬환경에서는 config_local.yaml 파일의 database.url 프로퍼티 값을 host.docker.internal 로 세팅해야합니다.
- https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host

[build]<br>
docker build -f DockerfileForMac --tag busstation-collector-mac:0.0.1 .

[run]<br>
docker run -p 3306:3306 busstation-collector-mac:0.0.1

<br>

## 설정 파일 준비
빌드 후 실행 시 빌드결과물 파일(busstation)과 같은 경로에 위치해야 합니다. <br>
빌드 없이 실행 시 프로젝트 루트 디렉토리 하위 resource 디렉토리 내에 위치해야 합니다.

파일명: config_{local|dev|stg|prd}.yaml <br>
ex. config_local.yaml

```yaml
api:
  key:
    seoul: seoulApiKey # 서울 API 키
    gyunggi: gyunggiApiKey # 경기 API 키
database:
  type: mysql # 데이터베이스 타입
  url: 127.0.0.1 # 데이터베이스 url
  port: 3306 # 데이터베이스 port
  id: id # 데이터베이스 접근 ID
  password: password # 데이터베이스 접근 패스워드
  database-name: dbName # 데이터베이스 DB명
```

<br>

## 빌드 및 실행
env 옵션 입력하지 않으면 배포환경이 디폴트로 "local"로 지정되고, 빌드 결과물 파일(busstation)과 같은 경로에 config_local.yaml 파일이 있어야합니다.

```console
$ go build
$ ./busstation -env {local|dev|stg|prd}
```

<br>

## 빌드 없이 실행

```console
$ go run main.go
```

<br>

# TODO
- 비동기 로직 추가로 성능 개선
- 환경에 따른 docker build (local, dev 등)
- 도커 이미지 실행용 스테이지 추가
