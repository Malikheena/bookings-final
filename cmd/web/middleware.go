package main

import (
	"net/http"

	"github.com/Malikheena/bookingsfinal/internal/helpers"
	"github.com/justinas/nosurf"
)

//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			fmt.Println("Hit the page")
//			next.ServeHTTP(w, r)
//		})
//	}
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// // SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			session.Put(r.Context(), "error", "Log in first!")
			http.Redirect(w, r, "/user/login", http.StatusResetContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
