package main

type Data struct {
	ID          uint   `gorm:"primaryKey"`
	Ticker      string `gorm:"not null" json:"ticker"`
	Company     string `gorm:"not null" json:"company"`
	Brokerage   string `gorm:"not null" json:"brokerage"`
	Action      string `gorm:"not null" json:"action"`
	Rating_from string `gorm:"not null" json:"rating_from"`
	Rating_to   string `gorm:"not null" json:"rating_to"`
	Target_from string `gorm:"not null" json:"target_from"`
	Target_to   string `gorm:"not null" json:"target_to"`
}

type APIResponse struct {
	Items    []Data `json:"items"`
	NextPage string `json:"next_page"`
}
