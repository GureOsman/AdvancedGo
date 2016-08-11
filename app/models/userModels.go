package models

import "time"

type User  struct{
	ID            int64 `"json:"AUTO_INCREMENT,primary_key"`
	Name          string `json:"name"`
	Email         string  `json:"email"`
	Password      string   `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time

}
