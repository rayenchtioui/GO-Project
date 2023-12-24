package model

type CustomerEventData struct {
	EventDataID uint64
	EventID     uint64
	ContentID   uint32
	CustomerID  uint64
	EventTypeID uint16
	EventDate   string
	Quantity    uint16
	InsertDate  string
}
