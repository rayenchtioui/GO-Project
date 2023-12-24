package data

import (
	"go-project/pkg/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var (
	startDate = time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC)
	endDate   = time.Now()
)

func GenerateCustomers(count int) []model.Customer {
	var customers []model.Customer
	for i := 0; i < count; i++ {
		randomDate := gofakeit.DateRange(startDate, endDate)
		formattedDate := randomDate.Format("2006-01-02 15:04:05")
		customer := model.Customer{
			CustomerID:       uint64(i + 1),
			ClientCustomerID: uint64(gofakeit.Number(1000, 9999)),
			InsertDate:       formattedDate,
		}
		customers = append(customers, customer)
	}
	return customers
}

func GenerateCustomerData(customers []model.Customer, channelTypes []model.ChannelType, entriesPerCustomer int) []model.CustomerData {
	var customerData []model.CustomerData
	var id = 0
	for _, c := range customers {
		for i := 0; i < entriesPerCustomer; i++ {
			ct := channelTypes[gofakeit.Number(0, len(channelTypes)-1)]
			randomDate := gofakeit.DateRange(startDate, endDate)
			formattedDate := randomDate.Format("2006-01-02 15:04:05")
			var channelValue string
			switch ct.Name {
			case "Email":
				channelValue = gofakeit.Email()
			case "PhoneNumber":
				channelValue = gofakeit.Phone()
			case "Postal":
				channelValue = gofakeit.Zip()
			case "MobileID":
				channelValue = gofakeit.UUID()
			case "Cookie":
				channelValue = gofakeit.UUID()
			}
			customerDatum := model.CustomerData{
				CustomerChannelID: uint64(id + 1),
				CustomerID:        c.CustomerID,
				ChannelTypeID:     ct.ChannelTypeID,
				ChannelValue:      channelValue,
				InsertDate:        formattedDate,
			}
			customerData = append(customerData, customerDatum)
			id++
		}
	}
	return customerData
}

func GenerateCustomerEvents(count int) []model.CustomerEvent {
	var customerEvents []model.CustomerEvent
	for i := 0; i < count; i++ {
		randomDate := gofakeit.DateRange(startDate, endDate)
		formattedDate := randomDate.Format("2006-01-02 15:04:05")
		customerEvent := model.CustomerEvent{
			EventID:       uint64(i + 1),
			ClientEventID: uint64(gofakeit.Number(1000, 9999)),
			InsertDate:    formattedDate,
		}
		customerEvents = append(customerEvents, customerEvent)
	}
	return customerEvents
}

func GenerateContent(count int) []model.Content {
	var contents []model.Content
	for i := 0; i < count; i++ {
		randomDate := gofakeit.DateRange(startDate, endDate)
		formattedDate := randomDate.Format("2006-01-02 15:04:05")
		content := model.Content{
			ContentID:       uint32(i + 1),
			ClientContentID: uint64(gofakeit.Number(1000, 9999)),
			InsertDate:      formattedDate,
		}
		contents = append(contents, content)
	}
	return contents
}

func GenerateContentPrices(contents []model.Content) []model.ContentPrice {
	var contentPrices []model.ContentPrice
	for i, c := range contents {
		randomDate := gofakeit.DateRange(startDate, endDate)
		formattedDate := randomDate.Format("2006-01-02 15:04:05")
		contentPrice := model.ContentPrice{
			ContentPriceID: uint32(i + 1),
			ContentID:      c.ContentID,
			Price:          gofakeit.Price(1, 1000),
			Currency:       gofakeit.CurrencyShort(),
			InsertDate:     formattedDate,
		}
		contentPrices = append(contentPrices, contentPrice)
	}
	return contentPrices
}

func GenerateCustomerEventData(customerEvents []model.CustomerEvent, contents []model.Content, customers []model.Customer, eventTypes []model.EventType) []model.CustomerEventData {
	var customerEventData []model.CustomerEventData
	var id = 1
	for _, e := range customerEvents {
		for _, c := range contents {
			for _, cu := range customers {
				eventType := eventTypes[gofakeit.Number(0, len(eventTypes)-1)]
				randomEventDate := gofakeit.DateRange(startDate, endDate).Format("2006-01-02 15:04:05")
				randomInsertDate := gofakeit.DateRange(startDate, endDate).Format("2006-01-02 15:04:05")
				eventDatum := model.CustomerEventData{
					EventDataID: uint64(id),
					EventID:     e.EventID,
					ContentID:   c.ContentID,
					CustomerID:  cu.CustomerID,
					EventTypeID: eventType.EventTypeID,
					EventDate:   randomEventDate,
					Quantity:    uint16(gofakeit.Number(1, 10)),
					InsertDate:  randomInsertDate,
				}
				customerEventData = append(customerEventData, eventDatum)
				id++
			}
		}
	}
	return customerEventData
}
