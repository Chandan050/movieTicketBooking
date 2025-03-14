package models

import (
	"go.starlark.net/lib/time"
	"gorm.io/gorm"
)

// user login table
type Account struct {
	gorm.Model
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

// person details table
type Person struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type Movies struct {
	gorm.Model
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DurationMins int       `json:"duration_mins"`
	Language     string    `json:"language"`
	ReleaseDate  time.Time `json:"release_date"`
	Country      string    `json:"country"`
	Genre        string    `json:"genre"`
}

type Address struct {
	Name    string `json:"name"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

type Cinema struct {
	gorm.Model
	Name             string  `json:"name"`
	TotalCinemaHalls int     `json:"totalCinemaHalls"`
	Location         Address `json:"location"`
}
