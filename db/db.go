package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMSQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

// // A hypothetical local config struct that includes a method to format the DSN
// type config struct {
// 	// ... fields like Host, Port, User, Password ...
// }

// func (c *config) FormatDSN() string {
// 	// Example DSN formatting logic
// 	return "user=user password=password dbname=mydb sslmode=disable"
// }

// func NewStorage(cfg config) (*sql.DB, error) {
// 	// Note: Changed "postgres" to "pgx" for the recommended driver
// 	db, err := sql.Open("pgx", cfg.FormatDSN())
// 	if err != nil {
// 		// FIX: Return the error instead of crashing the app
// 		return nil, fmt.Errorf("error opening database", err)
// 	}

// 	// You should also call Ping() to verify the connection is live:
// 	if err = db.Ping(); err != nil {
// 		db.Close() // Close the bad connection
// 		return nil, fmt.Errorf("error connecting to database", err)
// 	}

// 	return db, nil
// }
