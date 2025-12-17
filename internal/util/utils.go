package util

import (
	"context"
	"soccer-api/internal/http/middleware"
)

func GetUserID(ctx context.Context) int64 {
	userID, _ := ctx.Value(middleware.UserIDKey).(int64)
	return userID
}
