package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

func Init() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		// TODO : Loger process
	}

	//db.AutoMigrate()
}

func (database *DataBase) GetDSN() string {

}
