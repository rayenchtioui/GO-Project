package database

import (
	"database/sql"
	"log"
)

func createCustomerTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS Customer (
		CustomerID bigint UNSIGNED AUTO_INCREMENT NOT NULL,
		ClientCustomerID bigint UNSIGNED NOT NULL,
		InsertDate timestamp NOT NULL,
		PRIMARY KEY (CustomerID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createCustomerDataTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS CustomerData (
		CustomerChannelID bigint UNSIGNED AUTO_INCREMENT NOT NULL,
		CustomerID bigint UNSIGNED NOT NULL,
		ChannelTypeID smallint UNSIGNED NOT NULL,
		ChannelValue varchar(600) NOT NULL,
		InsertDate timestamp NOT NULL,
		PRIMARY KEY (CustomerChannelID),
		FOREIGN KEY (CustomerID) REFERENCES Customer (CustomerID),
		FOREIGN KEY (ChannelTypeID) REFERENCES ChannelType (ChannelTypeID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createChannelTypeTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS ChannelType (
		ChannelTypeID smallint UNSIGNED AUTO_INCREMENT NOT NULL,
		Name varchar(30) NOT NULL,
		PRIMARY KEY (ChannelTypeID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createEventTypeTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS EventType (
		EventTypeID smallint UNSIGNED AUTO_INCREMENT NOT NULL,
		Name varchar(30) NOT NULL,
		PRIMARY KEY (EventTypeID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createCustomerEventTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS CustomerEvent (
		EventID bigint UNSIGNED AUTO_INCREMENT NOT NULL,
		ClientEventID bigint NOT NULL,
		InsertDate timestamp NOT NULL,
		PRIMARY KEY (EventID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createContentTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS Content (
		ContentID int UNSIGNED AUTO_INCREMENT NOT NULL,
		ClientContentID bigint UNSIGNED NOT NULL,
		InsertDate timestamp NOT NULL,
		PRIMARY KEY (ContentID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createContentPriceTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS ContentPrice (
		ContentPriceID mediumint UNSIGNED AUTO_INCREMENT NOT NULL,
		ContentID int UNSIGNED NOT NULL,
		Price decimal(8,2) UNSIGNED NOT NULL,
		Currency char(3) NOT NULL,
		InsertDate timestamp NOT NULL,
		PRIMARY KEY (ContentPriceID),
		FOREIGN KEY (ContentID) REFERENCES Content (ContentID)
	);`
	_, err := tx.Exec(query)
	return err
}

func createCustomerEventDataTable(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS CustomerEventData (
		EventDataId bigint UNSIGNED AUTO_INCREMENT NOT NULL,
		EventID bigint UNSIGNED NOT NULL,
		ContentID int UNSIGNED NOT NULL,
		CustomerID bigint UNSIGNED NOT NULL,
		EventTypeID smallint UNSIGNED NOT NULL,
		EventDate timestamp NOT NULL,
		Quantity smallint UNSIGNED NOT NULL,
		InsertDate timestamp NOT NULL,
		PRIMARY KEY (EventDataId),
		FOREIGN KEY (EventID) REFERENCES CustomerEvent (EventID),
		FOREIGN KEY (ContentID) REFERENCES Content (ContentID),
		FOREIGN KEY (CustomerID) REFERENCES Customer (CustomerID),
		FOREIGN KEY (EventTypeID) REFERENCES EventType (EventTypeID)
	);`
	_, err := tx.Exec(query)
	return err
}

func SetupDatabase(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error starting transaction: %v", err)
	}
	tableCreationFunctions := []func(*sql.Tx) error{
		createChannelTypeTable,
		createEventTypeTable,
		createCustomerTable,
		createContentTable,
		createCustomerEventTable,
		createContentPriceTable,
		createCustomerDataTable,
		createCustomerEventDataTable,
	}
	for _, f := range tableCreationFunctions {
		err := f(tx)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Error creating table: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Error committing transaction: %v", err)
	}
	log.Println("Database setup completed successfully.")
}
