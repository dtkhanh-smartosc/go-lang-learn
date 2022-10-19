package database

import (
	"log"
	"restAPI/model"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB
var once sync.Once

func Migrate(database *gorm.DB) error {
	err := database.AutoMigrate(&model.Person{})
	if err != nil {
		return err
	}

	return nil
}

func ConnectSqlLite() error {
	once.Do(func() {
		Connection, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}

		if err = Migrate(Connection); err != nil {
			log.Fatal(err)
		}
	})

	return nil
}
