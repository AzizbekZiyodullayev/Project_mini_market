package models

import "time"

type CreateStaff struct {
	BranchID string  `json:"branch_id"`
	TariffID string  `json:"tariff_id"`
	Name     string  `json:"name"`
	Type     string  `json:"staff_type"`
	Balance  float64 `json:"balance"`
}

type Staff struct {
	ID        string    `json:"id"`
	BranchID  string    `json:"branch_id"`
	TariffID  string    `json:"tariff_id"`
	Type      string    `json:"staff_type"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExistsReq struct {
	Phone string `json:"phone"`
}

type StaffTop struct {
	BranchID string `json:"branch_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Money    int    `json:"money"`
}

type ChangeBalance struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type GetAllStaffRequest struct {
	Page        int     `json:"page"`
	Limit       int     `json:"limit"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	BalanceFrom float64 `json:"balance_from"`
	BalanceTo   float64 `json:"balance_to"`
}

type GetAllStaff struct {
	Staffs []Staff `json:"staffs"`
	Count  int     `json:"count"`
}
