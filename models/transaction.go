package models

import (
	"time"
)

type CreateTransaction struct {
	Type        string `json:"type"`
	Amount      int    `json:"amount"`
	Source_type string `json:"source_type"`
	Text        string `json:"text"`
	Sale_id     string `json:"sale_id"`
	Staff_id    string `json:"staff_id"`
}

type Transaction struct {
	Id          string    `json:"id"`
	Type        string    `json:"type"`
	Amount      int       `json:"amount"`
	Source_type string    `json:"source_type"`
	Text        string    `json:"text"`
	Sale_id     string    `json:"sale_id"`
	Staff_id    string    `json:"staff_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type GetAllTransactionRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Text  string `json:"text"`
}

type GetAllTransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"count"`
}

type TopWorkerRequest struct {
	Type     string `json:"type"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type TopWorkerRespond struct {
	Staff []TopWorker `json:"staff"`
}

type TopWorker struct {
	BranchName string `json:"branch_name"`
	StaffName  string `json:"staff_name"`
	StaffType  string `json:"staff_type"`
	EarnedSum  int    `json:"earned_sum"`
}
