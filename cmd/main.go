package main

import (
	"database/sql"
	"log"

	"vehicle-docs-in-go/cmd/api"
	"vehicle-docs-in-go/db"
	"vehicle-docs-in-go/db/config"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initializeDatabase(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal((err))
	}
}

func initializeDatabase(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Println("Database connected successfully")
	}
}
