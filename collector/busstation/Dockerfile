# 빌드 스테이지
FROM --platform=linux/amd64 golang AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# proj 설치
RUN apt-get update && apt-get install -y libproj-dev proj-bin

# 빌드
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -a -ldflags="-w -s" -o main .

# 실행용 디렉토리 세팅
WORKDIR /dist
RUN mkdir resources
RUN cp /build/main .
RUN cp /build/resources/config_local.yaml ./resources/config_local.yaml

EXPOSE 3306

ENTRYPOINT ["/app/main"]
