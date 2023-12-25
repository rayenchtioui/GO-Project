package processing

import (
	"go-project/pkg/model"
)

func IdentifyTopCustomers(sortedCustomers []model.CustomerRevenue) []model.CustomerRevenue {
	quantileSize := 0.025 // Top 2.5% of customers
	numberOfTopCustomers := int(float64(len(sortedCustomers)) * quantileSize)

	if numberOfTopCustomers == 0 && len(sortedCustomers) > 0 {
		numberOfTopCustomers = 1
	}

	return sortedCustomers[:numberOfTopCustomers]
}

func AnalyzeRevenueDistribution(sortedCustomers []model.CustomerRevenue, quantileSize float64) map[int]model.QuantileInfo {
	quantileMap := make(map[int]model.QuantileInfo)
	totalCustomers := len(sortedCustomers)
	quantileIndex := int(quantileSize * float64(totalCustomers))

	for i, customer := range sortedCustomers {
		currentQuantile := i / quantileIndex
		if currentQuantile >= 100 {
			currentQuantile = 99
		}

		quantileInfo := quantileMap[currentQuantile]
		quantileInfo.NumberOfCustomers++

		// Update Max and Min Revenue for the quantile
		if customer.CA > quantileInfo.MaxRevenue {
			quantileInfo.MaxRevenue = customer.CA
		}
		if quantileInfo.MinRevenue == 0 || customer.CA < quantileInfo.MinRevenue {
			quantileInfo.MinRevenue = customer.CA
		}

		quantileMap[currentQuantile] = quantileInfo
	}

	return quantileMap
}
