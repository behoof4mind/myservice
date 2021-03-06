package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"myservice/environment"
)

var Db *sqlx.DB

func init() {
	environment.InitEnvVariables()
	db, err := sql.Open("mysql", environment.E.DB_USERNAME+":"+environment.E.DB_PASSWORD+"@("+environment.E.DB_URL+")/myservice?parseTime=true")
	if err != nil {
		logrus.Fatalf("Fail to connect to database: %v", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "models/migrations",
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		logrus.Fatalf("Fail apply migrations: %v", err)
	}
	logrus.Printf("Applied %d migrations!\n", n)
	db.Close()

	Db, err = sqlx.Connect("mysql", environment.E.DB_USERNAME+":"+environment.E.DB_PASSWORD+"@("+environment.E.DB_URL+")/myservice?parseTime=true")
	if err != nil {
		logrus.Fatalf("Fail to connect to database: %v", err)
	}
}
