package models

import "time"

type CreateSales struct {
	Client_name       string  `json:"client_name"`
	Branch_id         string  `json:"branch_id"`
	Shop_assistant_id string  `json:"shop_assistant_id"`
	Cashier_id        string  `json:"cashier_id"`
	Price             float64 `json:"price"`
	Payment_Type      string  `json:"payment_type"` // card, cash
	Status            string  `json:"status"`       // success, cancel
}

type Sales struct {
	Id                string    `json:"id"`
	Client_name       string    `json:"client_name"`
	Branch_id         string    `json:"branch_id"`
	Shop_assistant_id string    `json:"shop_assistant_id"`
	Cashier_id        string    `json:"cashier_id"`
	Price             float64   `json:"price"`
	Payment_Type      string    `json:"payment_type"` // card, cash
	Status            string    `json:"status"`       // success, cancel
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}

type GetAllSalesRequest struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	Client_name string `json:"client_name"`
}

type GetAllSalesResponse struct {
	Sales []Sales `json:"sales"`
	Count int     `json:"count"`
}

type SaleTopBranch struct {
	Day         string  `json:"day"`
	BranchId    string  `json:"branch_id"`
	SalesAmount float64 `json:"sales_amount"`
}

type SaleCountSumBranch struct {
	BranchId    string  `json:"branch_id"`
	Count       int     `json:"count"`
	SalesAmount float64 `json:"sales_amount"`
}
