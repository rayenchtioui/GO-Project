package main

import (
	"go-project/data"
	"go-project/pkg/database"
	"go-project/pkg/exporter"
	"go-project/pkg/helper"
	"go-project/pkg/processing"
	"log"
)

func main() {
	db := database.GetDB()
	database.SetupDatabase(db)
	data.GenerateAndWriteData()
	database.InsertData()
	log.Println("Calculating Customer Revenues")
	customerRevenues, err := processing.CalculateCustomerRevenues(db)
	if err != nil {
		log.Fatalf("Error calculating customer revenues: %v", err)
		return
	}

	// Step 2: Creating and Sorting Revenue Data
	log.Println("Sorting Customers")
	sortedCustomers := processing.SortCustomersByRevenue(customerRevenues)

	// Step 3: Identifying Top Customers
	log.Println("Top Customers")
	topCustomers := processing.IdentifyTopCustomers(sortedCustomers)

	// Step 4: Revenue Distribution Analysis Across Quantiles
	log.Println("Quantile Analysis")
	quantileAnalysis := processing.AnalyzeRevenueDistribution(sortedCustomers, 0.025)

	log.Println("Top Customers:", topCustomers)
	log.Println("Quantile Analysis:", quantileAnalysis)
	// step 5: Exporting Customer Data
	log.Println("Retrieving emails for customers")
	exportCustomers, err := helper.GetExportCustomers(db, topCustomers)
	if err != nil {
		log.Fatalf("Error in GetExportCustomers: %v", err)
		return
	}
	log.Println("Exporting Top Customers")
	err = exporter.ExportCustomersData(db, exportCustomers)
	if err != nil {
		log.Fatalf("Error exporting customer data: %v", err)
		return
	}
	database.CloseDB()
}
