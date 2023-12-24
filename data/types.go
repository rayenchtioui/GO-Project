package data

import (
	"go-project/pkg/model"
)

func FixedChannelTypes() []model.ChannelType {
	return []model.ChannelType{
		{ChannelTypeID: 1, Name: "Email"},
		{ChannelTypeID: 2, Name: "PhoneNumber"},
		{ChannelTypeID: 3, Name: "Postal"},
		{ChannelTypeID: 4, Name: "MobileID"},
		{ChannelTypeID: 5, Name: "Cookie"},
	}
}

func FixedEventTypes() []model.EventType {
	return []model.EventType{
		{EventTypeID: 1, Name: "sent"},
		{EventTypeID: 2, Name: "view"},
		{EventTypeID: 3, Name: "click"},
		{EventTypeID: 4, Name: "visit"},
		{EventTypeID: 5, Name: "cart"},
		{EventTypeID: 6, Name: "purchase"},
	}
}
