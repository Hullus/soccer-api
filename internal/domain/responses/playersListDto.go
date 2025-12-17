package responses

type Player struct {
	FirstName        string `db:"first_name" json:"first_name"`
	LastName         string `db:"last_name" json:"last_name"`
	Country          string `db:"country" json:"country"`
	Age              int    `db:"age" json:"age"`
	MarketValueCents int64  `db:"market_value_cents" json:"market_value_cents"`
	Position         string `db:"position" json:"position"`
}
