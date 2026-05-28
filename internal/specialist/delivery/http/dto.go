package http

type CreateSpecialistDTO struct {
	FullName           string `json:"full_name"`
	Position           string `json:"position"`
	PhoneNumber        string `json:"phone_number"`
	ManagingCamaiginID int    `json:"managing_campaigin_id"`
}

type DeleteSpecialistDTO struct {
	ID int `json:"id"`
}
