package DBModels

type User struct {
	ID           int64  `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}
