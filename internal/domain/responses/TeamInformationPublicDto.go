package responses

type TeamInformationPublic struct {
	Name             string `json:"name"`
	BudgetCents      int64  `json:"budget_cents"`
	Country          string `json:"country"`
	TotalMarketValue int64  `json:"total_market_value"`
}
