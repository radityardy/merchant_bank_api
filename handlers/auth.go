package handlers

import (
	"encoding/json"
	"merchant_bank_api/models"
	"net/http"
	"os"
)

var loggedInCustomer *models.Customer

// LoginHandler godoc
// @Summary Login customer
// @Description Logs in a customer using username and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param customer body models.Customer true "Customer Login"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds models.Customer
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	file, _ := os.ReadFile("data/customers.json")
	var customers []models.Customer
	json.Unmarshal(file, &customers)

	for _, customer := range customers {
		if customer.Username == creds.Username && customer.Password == creds.Password {
			loggedInCustomer = &customer
			json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

// LogoutHandler godoc
// @Summary Logout customer
// @Description Logs out the currently logged-in customer
// @Tags auth
// @Success 200 {object} map[string]string
// @Router /logout [post]
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if loggedInCustomer == nil {
		http.Error(w, "No customer logged in", http.StatusBadRequest)
		return
	}

	loggedInCustomer = nil
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
