package responses

type MarketListing struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Country          string `json:"country"`
	Position         string `json:"position"`
	AskingPriceCents int64  `json:"asking_price_cents"`
}
