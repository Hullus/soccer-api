package middleware

import (
	"context"
	"net/http"
)

const LanguageKey string = "language"

func Language(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Accept-Language")
		if lang != "ka" {
			lang = "en"
		}
		ctx := context.WithValue(r.Context(), LanguageKey, lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
