# 구간정보 조회

## pre-requisite

1. [bun](https://bun.sh/docs/installation) 설치
2. depenency 설치
    ```bash
    bun install
    ```
3. .env 파일 수정 (mysql 접속 정보)
    ```dotenv
    DATABASE_URL="mysql://${user}:${password}@localhost:3306/{database}"
    ```
4. database 세팅
    ```bash
    bunx prisma db pull # 기존 데이터베이스가 존재하는 경우 prisma 스키마 세팅용
    bunx prisma db push # 노선정보 스키마 적용
    ```

## 실행

### 서울 구간정보 조회
```bash
bun run seoul
```
