package database

// import (
// 	"bitereview/config"
// 	"bitereview/model"

// 	"github.com/sirupsen/logrus"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type DBInstance struct {
// 	DB *gorm.DB
// }

// var DB DBInstance

// func InitDatabase(c *config.Config) {
// 	db, err := gorm.Open(postgres.Open(c.Database.ConnectionURL))
// 	if err != nil {
// 		logrus.Fatalf("Failed to initialize database: %s", err.Error())
// 	}

// 	logrus.Info("Connection opened to database")

// 	if err := db.AutoMigrate(&model.User{}); err != nil {
// 		logrus.Fatalf("Failed to auto migrate database: %s", err.Error())
// 	}

// 	DB = DBInstance{
// 		DB: db,
// 	}
// }
