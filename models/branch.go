package models

type CreateBranch struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	FoundedAt int    `json:"founded_at"`
}

type UpdateBranch struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	FoundedAt int    `json:"founded_at"`
}

type Branch struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Year      int    `json:"year"`
	FoundedAt int    `json:"founded_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IdRequest struct {
	Id string `json:"id"`
}

type GetAllBranchRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Name  string `json:"name"`
}

type GetAllBranch struct {
	Branches []Branch `json:"branches"`
	Count    int      `json:"count"`
}
