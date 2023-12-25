package processing

import (
	"database/sql"
	"go-project/pkg/model"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/schollz/progressbar/v3"
)

func CalculateCustomerRevenues(db *sql.DB) ([]model.CustomerRevenue, error) {
	eventDataQuery := "SELECT CustomerID, ContentID, Quantity FROM CustomerEventData WHERE EventTypeID = 6 AND EventDate >= '2020-04-01'"
	eventDataRows, err := db.Query(eventDataQuery)
	if err != nil {
		log.Printf("Error executing eventData query: %v\n", err)
		return nil, err
	}
	defer eventDataRows.Close()

	// Retrieve ContentPrice
	priceQuery := "SELECT ContentID, Price FROM ContentPrice"
	priceRows, err := db.Query(priceQuery)
	if err != nil {
		log.Printf("Error executing price query: %v\n", err)
		return nil, err
	}
	defer priceRows.Close()

	// Initialize the progress bar
	bar := progressbar.Default(-1)

	// Map to store prices for each content
	contentPrices := make(map[string]float64)
	for priceRows.Next() {
		var contentID string
		var price float64
		if err := priceRows.Scan(&contentID, &price); err != nil {
			log.Printf("Error scanning price row: %v\n", err)
			return nil, err
		}
		contentPrices[contentID] = price
		bar.Add(1) // Update the progress bar
	}

	if err = priceRows.Err(); err != nil {
		log.Printf("Error iterating price rows: %v\n", err)
		return nil, err
	}

	bar.Finish() // Finish the progress bar for price rows
	bar.Reset()  // Reset the progress bar for eventData rows

	// Map to temporarily store the total revenue for each customer
	tempRevenue := make(map[string]float64)
	for eventDataRows.Next() {
		var customerID, contentID string
		var quantity float64
		if err := eventDataRows.Scan(&customerID, &contentID, &quantity); err != nil {
			log.Printf("Error scanning eventData row: %v\n", err)
			return nil, err
		}

		if price, exists := contentPrices[contentID]; exists {
			tempRevenue[customerID] += quantity * price
		}

		bar.Add(1) // Update the progress bar
	}

	if err = eventDataRows.Err(); err != nil {
		log.Printf("Error iterating eventData rows: %v\n", err)
		return nil, err
	}

	bar.Finish() // Finish the progress bar for eventData rows

	// Convert the map to a slice of CustomerRevenue
	var customerRevenues []model.CustomerRevenue
	for customerID, revenue := range tempRevenue {
		customerRevenues = append(customerRevenues, model.CustomerRevenue{
			CustomerID: customerID,
			CA:         revenue,
		})
	}
	printRandomEntries(customerRevenues, 10)

	return customerRevenues, nil
}

// printRandomEntries prints n random entries from the customer revenue slice.
func printRandomEntries(revenues []model.CustomerRevenue, n int) {
	log.Printf("Printing %d random entries from the slice:\n", n)
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := 0; i < n && i < len(revenues); i++ {
		idx := r.Intn(len(revenues))
		log.Printf("CustomerID: %s, TotalRevenue: %f\n", revenues[idx].CustomerID, revenues[idx].CA)
	}
}

func SortCustomersByRevenue(customers []model.CustomerRevenue) []model.CustomerRevenue {
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].CA > customers[j].CA
	})
	return customers
}
