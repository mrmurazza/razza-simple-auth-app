package pkg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDatabase() {
	Database, _ = sql.Open("sqlite3", "dealljobs.db")
	statement, _ := Database.Prepare("CREATE TABLE IF NOT EXISTS users (" +
		"id integer primary key, " +
		"username varchar(50) not null, " +
		"name varchar(50) not null, " +
		"password varchar(50) not null, " +
		"address varchar(50) not null, " +
		"role varchar(50) not null, " +
		"created_at datetime not null default current_timestamp, " +
		"updated_at datetime not null default current_timestamp, " +
		"deleted_at datetime default null);")
	statement.Exec()

	statement.Close()
}
