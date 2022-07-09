package database

import (
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	// we use sync.Once for make sure we create connection only once
	once sync.Once
)

// CreateConnection is a function for creating new connection with database
// you can choose you want use mysql or postgresql
func CreateConnection() {

	conf := dbConfig{
		User: "root",
		Pass: "08520852",
		Host: "localhost",
		Port: 3306,
		Name: "cafetaria",
	}

	mysql := mysqlConfig{dbConfig: conf}
	// if you use postgres, you can uncomment code bellow

	// postgres := postgresqlConfig{dbConfig: conf}

	once.Do(func() {
		mysql.Connect()
		// postgres.Connect()
	})
}

// GetConnection is a faction for return connection or return value dbConn
// because we set var dbConn is private
func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
