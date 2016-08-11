package models
type Schedule struct {
	UserID            int64       `json:"userid"`
	ID                int         `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	StartDate         string      `json:"startdate"`
	EndDate           string      `json:"enddate"`
	Destination       string      `json:"destination"`
}
