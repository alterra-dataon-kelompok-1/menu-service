package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port int
		Name string
	}

	mysqlConfig struct {
		dbConfig
	}

	// postgresqlConfig struct {
	// 	dbConfig
	// }
)

func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// func (conf postgresqlConfig) Connect() {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
// 		conf.Host,
// 		conf.User,
// 		conf.Pass,
// 		conf.Name,
// 		conf.Port,
// 	)

// 	var err error

// 	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// }
