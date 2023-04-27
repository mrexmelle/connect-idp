package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/authx/account"
	"github.com/mrexmelle/connect-iam/authx/session"
)

func main() {
	r := chi.NewRouter()

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", account.Post)
		r.Patch("/{ehid}/tenures/{tenureId}/endDate", account.PatchEndDate)
	})

	r.Route("/sessions", func(r chi.Router) {
		r.Post("/", session.Post)
	})

	r.Group(func(r chi.Router) {
		secretToken := jwtauth.New("HS256", []byte("1nt3rst3ll4r-*-a5tR0"), nil)
		r.Use(jwtauth.Verifier(secretToken))

		r.Route("/accounts/me/profile", func(r chi.Router) {
			r.Get("/", account.GetMyProfile)
		})

		r.Route("/accounts/me/tenures", func(r chi.Router) {
			r.Get("/", account.GetMyTenures)
		})
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
