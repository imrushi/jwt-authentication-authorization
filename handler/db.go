package handler

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "userdb"
)

func GetDatabase() *gorm.DB {

	datanaseurl := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := gorm.Open(user, datanaseurl)
	if err != nil {
		log.Fatalf("Wrong Database Url: %v", err)
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Errorf("Error Database Connection : %v", err)
	}

	log.Info("Connected to Database")
	return connection
}

func InitalMigration() {
	connection := GetDatabase()
	defer Closedatabse(connection)
	connection.AutoMigrate(User{})
}

func Closedatabse(connction *gorm.DB) {
	sqldb := connction.DB()
	sqldb.Close()
	log.Info("Database Connection is closed")
}
