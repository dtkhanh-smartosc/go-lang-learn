package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gormGinAPI/model"
	"log"
	"sync"
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

func ConnectSqlite() (connection *gorm.DB, err error) {
	/*
		this make a connection happened one time only
	*/
	once.Do(func() {

		connection, err = gorm.Open(sqlite.Open("sqlite-gorm.db"), &gorm.Config{})

		//defer func() {
		//	if r := recover(); r != nil {
		//		fmt.Println("Recover in ConnectSqlite", r)
		//	}
		//}()
		if err != nil {
			log.Fatalln(err)
		}
		if err := Migrate(connection); err != nil {
			log.Fatalln(err)
		}
	})
	log.Println("connect to database success", connection)
	return connection, nil
}
