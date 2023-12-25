package exporter

import (
	"database/sql"
	"fmt"
	"go-project/pkg/model"
	"log"
	"strings"
	"time"
)

func ExportCustomersData(db *sql.DB, customers []model.CustomerExport) error {
	currentDate := time.Now().Format("20060102")
	tableName := fmt.Sprintf("test_export_%s", currentDate)
	// Create the export table for the current day
	err := CreateExportTable(db)
	if err != nil {
		return fmt.Errorf("error creating export table: %v", err)
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Prepare bulk insert and update statements
	var insertValues []string
	var updateValues []string
	var params []interface{}
	paramId := 1

	for _, customer := range customers {
		exists, err := CheckCustomerExists(tx, tableName, customer.CustomerID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error checking customer existence: %v", err)
		}

		if exists {
			// Prepare update
			updateValues = append(updateValues, fmt.Sprintf("(%d, ?, ?)", customer.CustomerID))
			params = append(params, customer.Email, customer.CA)
		} else {
			// Prepare insert
			insertValues = append(insertValues, fmt.Sprintf("(?, ?, ?)"))
			params = append(params, customer.CustomerID, customer.Email, customer.CA)
		}
		paramId += 3
	}

	// Execute bulk insert
	if len(insertValues) > 0 {
		// Use the tableName variable in the SQL query
		insertSQL := fmt.Sprintf("INSERT INTO %s (CustomerID, Email, CA) VALUES %s", tableName, strings.Join(insertValues, ","))
		_, err = tx.Exec(insertSQL, params...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error bulk inserting customer data: %v", err)
		}
	}

	// Execute bulk update
	if len(updateValues) > 0 {
		// Use the tableName variable in the SQL query for update
		updateSQL := fmt.Sprintf("UPDATE %s SET Email = ?, CA = ? WHERE CustomerID IN (%s)", tableName, strings.Join(updateValues, ","))
		_, err = tx.Exec(updateSQL, params...)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error bulk updating customer data: %v", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit error: %v", err)
	}

	log.Println("Successfully exported customer data")
	return nil
}

func CheckCustomerExists(tx *sql.Tx, tableName string, customerID int) (bool, error) {
	// SQL statement to check if a record with the given CustomerID exists
	checkExistsSQL := `SELECT COUNT(*) FROM ` + tableName + ` WHERE CustomerID = ?`

	// Variable to store the count result
	var count int

	// Execute the query using the transaction object
	err := tx.QueryRow(checkExistsSQL, customerID).Scan(&count)
	if err != nil {
		// Handle the error appropriately; it might be more than just sql.ErrNoRows
		return false, fmt.Errorf("error checking if customer exists: %v", err)
	}

	// Check if the count is greater than 0, which means the customer exists
	return count > 0, nil
}
