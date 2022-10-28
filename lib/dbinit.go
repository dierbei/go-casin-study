package lib

import (
	"fmt"
	"log"

	"go-casbin-study/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func initDB() {
	Gorm = gormDB()
}
func gormDB() *gorm.DB {
	cf := config.GetConfig().Mysql
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cf.Username,
		cf.Password,
		cf.Address,
		cf.Port,
		cf.Database,
	)), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db
}
