package db

import (
	"github.com/INebotov/JustNets/backend/config"
	"github.com/INebotov/JustNets/backend/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

func (database *DataBase) Init() {
	log := logger.MyLog{}
	log.Init("db_logs")
	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.LogFatal("Cant connect to db ( Error: %s )", err)
	}
	log.LogInfo("Sucsessfly connected to DB!")
	database.DB = db
}

func (database *DataBase) ExecSQL(command string) {

}
