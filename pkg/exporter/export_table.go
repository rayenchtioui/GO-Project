package exporter

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateExportTable(db *sql.DB) error {
	// Constructing the table name using the current date
	currentDate := time.Now().Format("20060102")
	tableName := fmt.Sprintf("test_export_%s", currentDate)

	// SQL statement to create the table if it doesn't already exist
	createTableSQL := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            CustomerID INT PRIMARY KEY,
            Email VARCHAR(255),
            CA DECIMAL(10, 2)
        );
    `, tableName)

	// Executing the SQL statement
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating table %s: %v", tableName, err)
	}

	return nil
}
