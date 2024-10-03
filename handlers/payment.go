package handlers

import (
	"encoding/json"
	"fmt"
	"merchant_bank_api/models"
	"net/http"
	"os"
)

// PaymentHandler godoc
// @Summary Make a payment
// @Description Allows a logged-in customer to make a payment to a merchant
// @Tags payment
// @Accept  json
// @Produce  json
// @Param payment body map[string]interface{} true "Payment Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /payment [post]
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if loggedInCustomer == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var paymentRequest struct {
		Amount       int    `json:"amount"`
		MerchantName string `json:"merchant_name"`
	}

	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Load merchants
	file, _ := os.ReadFile("data/merchants.json")
	var merchants []models.Merchant
	json.Unmarshal(file, &merchants)

	var validMerchant *models.Merchant
	for _, merchant := range merchants {
		if merchant.Name == paymentRequest.MerchantName {
			validMerchant = &merchant
			break
		}
	}

	if validMerchant == nil {
		http.Error(w, "Merchant not found", http.StatusNotFound)
		return
	}

	if loggedInCustomer.Balance < paymentRequest.Amount {
		http.Error(w, "Insufficient balance", http.StatusBadRequest)
		return
	}

	loggedInCustomer.Balance -= paymentRequest.Amount
	logHistory(fmt.Sprintf("Payment of %d to %s", paymentRequest.Amount, validMerchant.Name))

	json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful"})
}

func logHistory(action string) {
	file, _ := os.ReadFile("data/history.json")
	var history []models.History
	json.Unmarshal(file, &history)

	history = append(history, models.History{
		ID:        len(history) + 1,
		Customer:  loggedInCustomer.Username,
		Action:    action,
		Timestamp: "2024-10-03T00:00:00Z", // mock timestamp for simplicity
	})

	historyData, _ := json.Marshal(history)
	os.WriteFile("data/history.json", historyData, 0644)
}
