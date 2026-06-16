package domain

import "time"

type FullApplication struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	DateTime          time.Time `json:"datetime"`
	Street            string    `json:"street"`
	HouseNumber       int       `json:"house_number"`
	SpecialistName    string    `json:"specialist_name"`
	SpecialistPost    string    `json:"specialist_post"`
	WorkType          string    `json:"work_type"`
	ApplicationStatus string    `json:"application_status"`
}

type Statistics struct {
	TotalRequests  int64 `json:"total_requests"`
	OpenRequests   int64 `json:"open_requests"`
	ClosedRequests int64 `json:"closed_requests"`
}
