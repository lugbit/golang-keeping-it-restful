package drivers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// ConnectDB opens a DB handler
func ConnectDB() *sql.DB {
	// Build datasource from environmental variables
	dsn := buildDataSrc()

	// Open databse using the mysql driver and the configuration above
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// Ping the database to test connection
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// buildDataSrc builds the data source name(DSN) to the following format:
// username:password@protocol(address)/dbname?param=value
func buildDataSrc() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}
