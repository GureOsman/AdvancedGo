package models
type Passengers struct {
	ID           int64  `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	UserID       int64     `json:"userid"`
	Tickets      []Tickets
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	PasPhone     int      `json:"pasphone"`
}