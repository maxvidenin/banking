package dto

type NewAccountResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}
