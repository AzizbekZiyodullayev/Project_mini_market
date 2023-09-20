package models

type CreateStaffTarif struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"` // (fixed, percent)
	AmountForCash float64 `json:"amount_for_cash"`
	AmountForCard float64 `json:"amount_for_card"`
}

type StaffTarif struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"` // (fixed, percent)
	AmountForCash float64 `json:"amount_for_cash"`
	AmountForCard float64 `json:"amount_for_card"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type GetAllStaffTarifRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Name  string `json:"name"`
}

type GetAllStaffTarif struct {
	StaffTarifs []StaffTarif `json:"staff_tarifs"`
	Count       int          `json:"count"`
}
