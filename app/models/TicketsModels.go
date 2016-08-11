package models

type Tickets struct {
	ID         int     `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	PasId      int      `json:"pasid"`
	PlaneID    int      `json:"planeid"`
	UserID     int64
	Price      int        `json:"price"`
	Type       string     `json:"type"`
}
