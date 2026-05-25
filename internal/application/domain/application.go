package domain

import "time"

type Application struct {
	ID           int
	Description  string
	CreatedAt    time.Time
	HouseID      int
	SpecialistID int
	WorkTypeID   int
	StatusID     int
}
