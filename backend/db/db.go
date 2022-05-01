package db

import (
	"fmt"

	"github.com/INebotov/JustNets/backend/config"
	"github.com/INebotov/JustNets/backend/logger"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	DB    *gorm.DB
	DEBUG bool
	log   logger.MyLog
}

func (database *DataBase) Init(dbname string, m ...interface{}) {
	database.log = logger.MyLog{}
	database.log.Init("db_logs", "DataBase")
	if database.DEBUG {
		db, err := gorm.Open(sqlite.Open(dbname+".db"), &gorm.Config{})
		if err != nil {
			database.log.LogFatal("Cant connect to db ( Error: %s )", err)
		}
		database.DB = db
	} else {
		dsn := fmt.Sprintf(config.GetDSN(), dbname)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			database.log.LogFatal("Cant connect to db ( Error: %s )", err)
		}
		database.DB = db
	}
	for _, el := range m {
		database.DB.AutoMigrate(el)
	}

	database.log.LogInfo("Sucsessfly connected to DB %s!", dbname)
}
