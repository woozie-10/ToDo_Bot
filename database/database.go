package database

import (
	"ToDo_bot/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	println(config.Config.GetString("bot.token"))
	mysql_database := config.Config.GetString("mysql.database")
	mysql_user := config.Config.GetString("mysql.user")
	mysql_password := config.Config.GetString("mysql.password")
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", mysql_user, mysql_password, mysql_database))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS Tasks (
            Id INT AUTO_INCREMENT PRIMARY KEY,
            TgId INT,
            Text MEDIUMTEXT
        )
    `)
	if err != nil {
		return err
	}
	return nil
}
