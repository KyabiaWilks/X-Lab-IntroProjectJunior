package main

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Comment struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func initDB() {
	//log.SetOutput(log.txt)
	viper.SetConfigName("config") // 读取json配置文件
	viper.AddConfigPath(".")      // 设置配置文件和可执行二进制文件在用一个目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	 var err error
    var dbUser = viper.GetString("postgres.username")  // 配置文件字段名可自己定义
    var dbPass = viper.GetString("postgres.password")
    var dbHost = viper.GetString("postgres.host")      // 例如 "localhost"
    var dbPort = viper.GetString("postgres.port")      // 例如 "5432"
    var dbName = viper.GetString("postgres.dbname")
    
    // PostgreSQL
    dsn := "host=" + dbHost +
        " user=" + dbUser +
        " password=" + dbPass +
        " dbname=" + dbName +
        " port=" + dbPort +
        " sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&Comment{}) // 自动建表
}
