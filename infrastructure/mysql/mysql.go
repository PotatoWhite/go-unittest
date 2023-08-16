package mysql

import (
	"database/sql"
	"fmt"
	"go-unittest/infrastructure/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"runtime"
	"time"
)

func Open() (*gorm.DB, error) {
	// gorm
	_gorm, err := initGORM()
	if err != nil {
		panic(err)
	}

	return _gorm, nil
}

func Close(db *gorm.DB) error {
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	return nil
}

func initGORM() (*gorm.DB, error) {
	// sqldb
	conn, err := initDBConn()
	if err != nil {
		panic(err)
	}

	// setting up gorm
	_gorm, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		Conn:       conn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return _gorm, nil
}

func initDBConn() (*sql.DB, error) {
	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.USER, cfg.PASSWD, cfg.HOST, cfg.PORT, cfg.DBNAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// SetMaxOpenConns 설정
	size := runtime.NumCPU()
	db.SetMaxOpenConns(size*2 + 2) // 취향 따라 조절
	db.SetMaxIdleConns(size)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func initConfig() (*Config, error) {
	return &Config{
		HOST:   config.GetString("db.host"),
		PORT:   config.GetInt("db.port"),
		USER:   config.GetString("db.username"),
		PASSWD: config.GetString("db.password"),
		DBNAME: config.GetString("db.dbname"),
	}, nil
}
