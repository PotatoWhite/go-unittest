package main

import (
	"flag"
	"github.com/pkg/errors"
	"go-unittest/infrastructure/config"
	"go-unittest/infrastructure/gin"
	"go-unittest/infrastructure/mysql"
	"go-unittest/internal/user"
	"gorm.io/gorm"
)

var (
	port       int
	configPath string
	db         *gorm.DB
)

func main() {
	defer close()

	// user 모듈
	userHandler := user.NewHandler(db)

	// gin 서버 실행
	engine := gin.New()

	// 라우터 등록
	userHandler.RouteGroup("users", engine)

	gin.Service(engine)

}

// init 함수는 main 함수보다 먼저 실행된다.
// init 함수는 패키지가 로드될 때 한번만 실행된다.
func init() {
	// 커맨드라인 인자를 입력받는다.
	if err := cmdlineInput(); err != nil {
		panic(err)
	}

	// 설정파일을 로드한다.
	if err := config.Load(configPath); err != nil {
		panic(err)
	}

	// mysql 연결
	_db, err := mysql.Open()
	if err != nil {
		panic(err)
	}
	db = _db
}

func close() {

	mysql.Close(db)
}

func cmdlineInput() error {
	// flag 라이브러리를 사용하여 커맨드라인 인자를 입력받는다.
	flag.StringVar(&configPath, "config", "./config/config.yaml", "config file path")
	flag.IntVar(&port, "port", 8080, "gin port")
	flag.Parse()

	if configPath == "" {
		return errors.New("invalid config path")
	}

	if port <= 0 {
		return errors.New("invalid port")
	}

	return nil
}
