package domain

type Alert struct {
	ID          string  `json:"alertId"`
	UserID      string  `json:"userId"`
	ProductID   string  `json:"productId"`
	TargetPrice float64 `json:"targetPrice"`
	IsActive    bool    `json:"isActive"`
}
