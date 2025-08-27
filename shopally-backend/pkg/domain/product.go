package domain

import "time"

// Price represents the price of a product in different currencies.
type Price struct {
	ETB         float64   `json:"etb"`
	USD         float64   `json:"usd"`
	FXTimestamp time.Time `json:"fxTimestamp"`
}

// Product represents a product found on an e-commerce platform.
type Product struct {
	ID                string   `json:"id"`
	Title             string   `json:"title"`
	ImageURL          string   `json:"imageUrl"`
	AIMatchPercentage int      `json:"aiMatchPercentage"`
	Price             Price    `json:"price"`
	ProductRating     float64  `json:"productRating"`
	SellerScore       int      `json:"sellerScore"`
	DeliveryEstimate  string   `json:"deliveryEstimate"`
	SummaryBullets    []string `json:"summaryBullets"`
	DeeplinkURL       string   `json:"deeplinkUrl"`
}
