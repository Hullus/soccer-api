package domain

import (
	"context"
	"soccer-api/internal/http/middleware"
)

type User struct {
	ID           int64
	Email        string
	PasswordHash string
}

// აქ წესით არ უნდა დაგტოვო შენ
func GetUserID(ctx context.Context) int64 {
	userID, _ := ctx.Value(middleware.UserIDKey).(int64)
	return userID
}
