package model

type ContentPrice struct {
	ContentPriceID uint32
	ContentID      uint32
	Price          float64
	Currency       string
	InsertDate     string
}
