package DBModels

type Team struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	BudgetCents int64  `db:"budget_cents"`
	Country     string `db:"country"`
	OwnerID     int64  `db:"owner_id"`
}
