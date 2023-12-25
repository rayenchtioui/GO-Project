package helper

import (
	"database/sql"
	"fmt"
	"go-project/pkg/model"
	"strconv"
)

func GetCustomerEmail(db *sql.DB, customerID string) (string, error) {
	var email string
	query := `SELECT ChannelValue FROM CustomerData WHERE CustomerID = ? AND ChannelTypeID = (SELECT ChannelTypeID FROM ChannelType WHERE Name = 'Email')`
	err := db.QueryRow(query, customerID).Scan(&email)
	if err != nil {
		return "", nil
	}
	return email, nil
}

func GetExportCustomers(db *sql.DB, topCustomers []model.CustomerRevenue) ([]model.CustomerExport, error) {
	var exportCustomers []model.CustomerExport
	for _, customer := range topCustomers {
		email, err := GetCustomerEmail(db, customer.CustomerID)
		if err != nil {
			return nil, fmt.Errorf("error retrieving email for customer %s: %v", customer.CustomerID, err)
		}

		customerID, err := strconv.Atoi(customer.CustomerID)
		if err != nil {
			return nil, fmt.Errorf("error converting customer ID to int: %v", err)
		}

		exportCustomer := model.CustomerExport{
			CustomerID: customerID,
			Email:      email,
			CA:         customer.CA,
		}
		exportCustomers = append(exportCustomers, exportCustomer)
	}
	return exportCustomers, nil
}
