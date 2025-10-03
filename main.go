package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := User{
		ID:    1,
		Name:  "Minh Duc",
		Email: "minhduc@example.com",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", getUser)

	log.Println("Server listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}