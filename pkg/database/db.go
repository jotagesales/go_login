package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// Connect to database and configure connection pool
func Connect(user, password, dbname string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	log.Debug(connectionString)
	db, err := gorm.Open("postgres", connectionString)
	// db.LogMode(true)

	//connection pool
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 60)
	return db, err
}
