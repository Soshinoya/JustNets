package db

import (
	"github.com/INebotov/JustNets/backend/config"
	"github.com/INebotov/JustNets/backend/logger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBase struct {
	DB    *gorm.DB
	DEBUG bool
	log   logger.MyLog
}

func (database *DataBase) Init() {
	database.log = logger.MyLog{}
	database.log.Init("db_logs", "DataBase")
	if database.DEBUG == true {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			database.log.LogFatal("Cant connect to db ( Error: %s )", err)
		}
		database.DB = db
	} else {
		dsn := config.GetDSN()
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			database.log.LogFatal("Cant connect to db ( Error: %s )", err)
		}
		database.DB = db
	}
	database.log.LogInfo("Sucsessfly connected to DB!")
}
