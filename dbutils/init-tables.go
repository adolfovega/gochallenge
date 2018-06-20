package dbutils

import "log"
import "database/sql"

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(task)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}
	log.Println("All tables created/initialized successfully!")
}
