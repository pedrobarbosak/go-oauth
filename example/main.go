package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedrobarbosak/go-oauth"
)

func main() {
	port := 8080

	clientID := "<client_id>"
	secret := "<client_secret>"
	redirectURL := fmt.Sprintf("http://localhost:%d/oauth/google", port)
	scopes := []string{"profile", "email"}

	auth := oauth.NewGoogle(clientID, secret, redirectURL, scopes...)

	// Login Handler
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		url, err := auth.Login()
		if err != nil {
			log.Println("login failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	// Callback from provider handler
	http.HandleFunc("/oauth/google", func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.Callback(r)
		if err != nil {
			log.Println("callback failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := auth.GetUser(r.Context(), token.AccessToken)
		if err != nil {
			log.Println("callback failed get user:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s - %s - %s - %s", user.ID, user.Email, user.Name, user.VerifiedEmail)
	})

	// Start server
	fmt.Println("Starting server on :", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
