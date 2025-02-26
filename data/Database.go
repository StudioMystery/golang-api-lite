package data

import (
	"golang-api-lite/utility"
	"log"

	"golang-api-lite/library"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	var config = utility.GetConfig()

	db, err := gorm.Open(sqlite.Open(config.Connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// CRUD Functions

func ReadFooBar(ID int) (library.FooBar, error) {
	var db = GetConnection()
	fooBar := library.FooBar{}

	err := db.Model(&library.FooBar{}).Where("ID = ?", ID).First(&fooBar).Error

	return fooBar, err
}

func UpdateFooBar(fooBar library.FooBar) (int64, error) {
	var db = GetConnection()

	result := db.Model(&library.FooBar{}).Where("ID = ?", fooBar.ID).Updates(&fooBar)

	return result.RowsAffected, result.Error
}

func CreateFooBar(fooBar library.FooBar) (library.FooBar, error) {
	var db = GetConnection()

	err := db.Model(&library.FooBar{}).Create(&fooBar).Error
	if err != nil {
		log.Fatalf("[%s]", err)
	}

	return fooBar, err
}

func DeleteFooBar(ID int) error {
	var db = GetConnection()

	err := db.Model(&library.FooBar{}).Delete("ID = ?", ID).Error

	return err
}
