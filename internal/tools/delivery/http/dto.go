package http

type CreateToolDTO struct {
	Name                string `json:"name"`
	ManagingCampaiginId int    `json:"managing_campaigin_id"`
	Quantity            int    `json:"quantity"`
}

type DeleteToolDTO struct {
	ID int `json:"id"`
}

type UpdateToolDTO struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
