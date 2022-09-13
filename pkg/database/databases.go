package database

import (
	"dealljobs/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func InitDatabase() {
	cfg := config.Get()
	db, err := gorm.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		panic(err)
	}

	DB = db

	DB.DB().SetMaxOpenConns(10)
	DB.DB().SetConnMaxLifetime(time.Duration(10) * time.Minute)
	DB.DB().SetMaxIdleConns(5)

	// set gorm configuration
	DB.LogMode(true)
	DB.SingularTable(false)

	//DB, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dealljobs-test")
	err = DB.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"`id` bigint(20) NOT NULL AUTO_INCREMENT, " +
		"username varchar(50) not null, " +
		"name varchar(50) not null, " +
		"password varchar(50) not null, " +
		"address varchar(50) not null, " +
		"role varchar(50) not null, " +
		"created_at datetime not null default current_timestamp, " +
		"updated_at datetime not null default current_timestamp, " +
		"deleted_at datetime default null," +
		"PRIMARY KEY (`id`)" +
		")").Error
	if err != nil {
		panic(err)
	}

	// Seeder Admin
	err = DB.Exec("INSERT INTO users (id, username, name, password, address, role) " +
		"VALUES(1, 'admin', '', '+OyZfszwFbIyrCuXmS7ObK8oBg2V0Mv6bagDBk6UFYM=', '', 'ADMIN') " +
		"ON DUPLICATE KEY UPDATE username='admin'").Error
	if err != nil {
		panic(err)
	}

}
