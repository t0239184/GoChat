package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/t0239184/GoChat/app/tool"
)

func init() {
	fmt.Println("[init] main")
}

func main() {
	r := gin.New()
	r.Use(tool.AccessLogMiddleware())
	// db := InitDatabase()

	serverPort := ":" + tool.GetString("server.port")
	r.Run(serverPort)
}

func InitDatabase() *gorm.DB {
	user := tool.GetString("database.connection-info.user")
	password := tool.GetString("database.connection-info.password")
	host := tool.GetString("database.connection-info.host")
	port := tool.GetInt("database.connection-info.port")
	database := tool.GetString("database.connection-info.database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", user, password, host, port, database)
	fmt.Println(dsn)
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		logrus.Fatalf("[main] database.New failed: %v", err)
	}
	db.AutoMigrate()
	return db
}