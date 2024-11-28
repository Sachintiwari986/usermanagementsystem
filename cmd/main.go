package main

import (
	"log"
	"net/http"

	"github.com/Sachintiwari986/user-management-system/internal/db"
	"github.com/Sachintiwari986/user-management-system/internal/handlers"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up HTTP routes
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("/reset-password", handlers.ResetPassword)
	http.HandleFunc("/forgot-password", handlers.ForgotPassword)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
