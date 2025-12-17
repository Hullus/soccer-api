package requests

type UpdatePlayerRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
}
