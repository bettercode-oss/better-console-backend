# better-console Back-end Service

Golang 으로 구현한 better-console Back-end Service


## 데이터베이스

### Sqlite
별도 설정을 하지 않는다면 기본적으로는 Sqlite file 데이터베이스를 사용한다.

### MySQL
* 데이터 베이스 생성
```sql
-- 아래 데이터베이스명은 예시
CREATE SCHEMA IF NOT EXISTS `better_console` DEFAULT CHARACTER SET utf8mb4;
```

* 애플리케이션 실행 환경 변수 설정
```
BETTER_CONSOLE_DB_HOST=localhost:3306
BETTER_CONSOLE_DB_DRIVER=mysql
BETTER_CONSOLE_DB_NAME=better_admin
BETTER_CONSOLE_DB_USER=root
BETTER_CONSOLE_DB_PASSWORD=1111
```
