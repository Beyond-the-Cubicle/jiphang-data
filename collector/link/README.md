# 구간정보 스크랩핑

## 전제조건
1. pnpm 설치
2. 도커에서 접근할 수 있도록 mysql url이 갖춰져야 함

## pnpm 설치방법
https://pnpm.io/ko/installation 참고

## 도커 실행방법
1. 프로젝트 디렉토리 아래에 `.env.sample`을 보고 `.env` 파일 생성
   1. 도커로 mysql이 올라와 있다면 `localhost` 대신 `host.docker.internal`로 변경
2. `pnpm run docker` 명령어 실행
