package models

import (
	uuid "github.com/satori/go.uuid"
)

// REQUEST -------

type GetLocationReq struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    int     `json:"radius"`
	Type      int     `json:"type"` // 0 -> is circle, 1 -> is square
}

// REQUEST -------

type Locations struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name        string    `json:"name"`
	Website     string    `json:"website"`
	Coordinates string    `json:"coordinates"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
}

type LocationsCount struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Count   int    `json:"count"`
}
