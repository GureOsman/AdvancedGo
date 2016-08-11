package models

type Planes struct {
	UserID      int64       `json:"userid"`
	ID          int       `gorm:"AUTO_INCREMENT,primary_key" json:"id"`
	Name         string     `json:"name"`
	Capacity    int       `json:"capacity"`
	Type        string      `json:"type"`
}
