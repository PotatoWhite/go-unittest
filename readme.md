# 의존성 설치

## viper

```shell
go get github.com/spf13/viper
```

## testify

```shell
go get github.com/stretchr/testify
```

## UUID 
```shell
go get github.com/google/uuid
```

## gorm & gorm mysql

```shell
go get gorm.io/gorm
go get gorm.io/driver/mysql

```

## gin

``` shell
go get github.com/gin-gonic/gin
```

# 환경 설정

## docker mysql

```shell
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1234 --name mysql mysql
```

### mysql 설정

```sql
create database user_service;

create user 'user_account'@'%' identified by '1234';

grant all privileges on user_service.* to'user_account'@'%';

flush privileges;
```