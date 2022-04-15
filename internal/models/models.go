package models

import "time"

type JsonResp struct {
	Key string `json:"key"`
	Definition string `json:"definition"`
}

type User struct {
	ID int
	Username string
	Password string
	CreatedAt time.Time
}