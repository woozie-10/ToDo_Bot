package database

import (
	"ToDo_bot/config"
	"ToDo_bot/entities"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	mysql_database := config.Config.GetString("mysql.database")
	mysql_user := config.Config.GetString("mysql.user")
	mysql_password := config.Config.GetString("mysql.password")
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?parseTime=true", mysql_user, mysql_password, mysql_database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&entities.Task{}); err != nil {
		panic(err)
	}
	return nil
}
