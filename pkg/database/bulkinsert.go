package database

import (
	"database/sql"
	"fmt"
	"go-project/pkg/model"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func BulkInsert(tx *sql.Tx, tableName string, columnNames []string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open CSV file %s: %v", filename, err)
	}
	defer file.Close()
	mysql.RegisterLocalFile(filename)
	// Register the CSV file as a local file with MySQL
	if _, err := tx.Exec("SET GLOBAL local_infile = true"); err != nil {
		log.Fatalf("Failed to enable local infile: %v", err)
	}

	if _, err := tx.Exec(fmt.Sprintf("LOAD DATA LOCAL INFILE '%s' INTO TABLE %s FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 LINES (%s)", filename, tableName, strings.Join(columnNames, ", "))); err != nil {
		log.Fatalf("Failed to load data from CSV into table %s: %v", tableName, err)
	}
}

// join concatenates a slice of strings with a given separator.
func join(items []string, separator string) string {
	return strings.Join(items, separator)
}

// createPlaceholders creates a string of placeholders for an SQL statement.
func createPlaceholders(numColumns int) string {
	placeholders := make([]string, numColumns)
	for i := range placeholders {
		placeholders[i] = "?"
	}
	return "(" + strings.Join(placeholders, ", ") + ")"
}

// convertInterfaceSlice converts a slice of strings to a slice of empty interfaces.
func convertInterfaceSlice(slice []string) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(slice))
	for i, d := range slice {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func StructFieldNames(v interface{}) ([]string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("provided value is not a struct: %v", val.Kind())
	}

	var fieldNames []string
	for i := 0; i < val.Type().NumField(); i++ {
		fieldNames = append(fieldNames, val.Type().Field(i).Name)
	}
	return fieldNames, nil
}

func InsertData() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error starting transaction for seeding: %v", err)
	}
	log.Println("Starting bulk insert for all tables.")

	tables := []struct {
		Name    string
		Struct  interface{}
		CSVFile string
	}{
		{"ChannelType", model.ChannelType{}, "csv/channel_types.csv"},
		{"EventType", model.EventType{}, "csv/event_types.csv"},
		{"Customer", model.Customer{}, "csv/customers.csv"},
		{"CustomerData", model.CustomerData{}, "csv/customer_data.csv"},
		{"CustomerEvent", model.CustomerEvent{}, "csv/customer_events.csv"},
		{"Content", model.Content{}, "csv/contents.csv"},
		{"ContentPrice", model.ContentPrice{}, "csv/content_prices.csv"},
		{"CustomerEventData", model.CustomerEventData{}, "csv/customer_event_data.csv"},
	}

	for _, table := range tables {
		log.Printf("Processing table: %s\n", table.Name)
		columnNames, err := StructFieldNames(table.Struct)
		if err != nil {
			log.Fatalf("Error retrieving field names for %s: %v", table.Name, err)
			tx.Rollback()
			return
		}

		BulkInsert(tx, table.Name, columnNames, table.CSVFile)
		log.Printf("Data inserted into table: %s\n", table.Name)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	log.Println("All data successfully inserted into the database.")
}
