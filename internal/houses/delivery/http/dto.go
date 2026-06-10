package http

type CreateHouseDTO struct {
	Street              string `json:"street"`
	HouseNumber         int    `json:"house_number"`
	EntrancesNumber     int    `json:"entrances_number"`
	ManagingCampaiginId int    `json:"managing_campaigin_id"`
}

type DeleteHouseDTO struct {
	ID int `json:"id"`
}
