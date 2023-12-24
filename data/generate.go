package data

import (
	"log"
)

func GenerateAndWriteData() {
	log.Println("Starting data generation and writing process.")

	log.Println("Generating fixed channel types and event types.")
	channelTypes := FixedChannelTypes()
	eventTypes := FixedEventTypes()

	log.Println("Writing channel types to CSV.")
	WriteCSV("csv/channel_types.csv", channelTypes)
	log.Println("Channel types written to CSV successfully.")

	log.Println("Writing event types to CSV.")
	WriteCSV("csv/event_types.csv", eventTypes)
	log.Println("Event types written to CSV successfully.")

	log.Println("Generating customer data")
	customers := GenerateCustomers(300)
	log.Println("Writing customers to CSV.")
	WriteCSV("csv/customers.csv", customers)
	log.Println("Customers written to CSV successfully.")

	log.Println("Generating customer data entries.")
	customerData := GenerateCustomerData(customers, channelTypes, 3)
	log.Println("Writing customer data to CSV.")
	WriteCSV("csv/customer_data.csv", customerData)
	log.Println("Customer data written to CSV successfully.")

	log.Println("Generating customer events.")
	customerEvents := GenerateCustomerEvents(150)
	log.Println("Writing customer events to CSV.")
	WriteCSV("csv/customer_events.csv", customerEvents)
	log.Println("Customer events written to CSV successfully.")

	log.Println("Generating content ")
	contents := GenerateContent(100)
	log.Println("Writing content data to CSV.")
	WriteCSV("csv/contents.csv", contents)
	log.Println("Content data written to CSV successfully.")

	log.Println("Generating content prices.")
	contentPrices := GenerateContentPrices(contents)
	log.Println("Writing content prices to CSV.")
	WriteCSV("csv/content_prices.csv", contentPrices)
	log.Println("Content prices written to CSV successfully.")

	log.Println("Generating customer event data")
	customerEventData := GenerateCustomerEventData(customerEvents, contents, customers, eventTypes)
	log.Println("Writing customer event data to CSV.")
	WriteCSV("csv/customer_event_data.csv", customerEventData)
	log.Println("Customer event data written to CSV successfully.")

	log.Println("Data generation and CSV writing process completed.")
}
