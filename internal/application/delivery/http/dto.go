package http

type CreateApplicationDTO struct {
	Description  string `json:"description"`
	HouseID      int    `json:"house_id"`
	SpecialistID int    `json:"specialist_id"`
	WorkTypeID   int    `json:"work_type_id"`
	StatusID     int    `json:"status_id"`
}

type DeleteApplicationDTO struct {
	ID int `json:"id"`
}

type CloseApplicationDTO struct {
	ID int `json:"id"`
}
