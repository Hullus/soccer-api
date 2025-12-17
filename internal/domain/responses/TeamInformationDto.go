package responses

type TeamInformation struct {
	ID               int64  `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	BudgetCents      int64  `db:"budget_cents" json:"budget_cents"`
	Country          string `db:"country" json:"country"`
	OwnerID          int64  `db:"owner_id" json:"owner_id"`
	TotalMarketValue int64  `db:"total_market_value" json:"total_market_value"`
}
