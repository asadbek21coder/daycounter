package handler

import (
	"html/template"
	"net/http"
)

const authToken = "30aadcf6-94cd-4032-8ca6-423e0f396826" // random uuid

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != authToken {
			tmpl, err := template.ParseFiles("templates/unauthorized.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}
