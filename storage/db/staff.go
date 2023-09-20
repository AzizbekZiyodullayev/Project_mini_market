package db

import (
	"context"
	"file/models"
	"file/pkg/helper"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{db: db}
}

func (s *staffRepo) CreateStaff(req *models.CreateStaff) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "staffs" ("id", "branch_id", "tariff_id", "staff_type", "name", "balance", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING "id"
	`

	result := s.db.QueryRow(context.Background(), query, id, req.BranchID, req.TariffID, req.Type, req.Name, req.Balance)

	var createdID string
	if err := result.Scan(&createdID); err != nil {
		return "", fmt.Errorf("failed to create staff: %w", err)
	}

	return createdID, nil
}

func (s *staffRepo) UpdateStaff(req *models.Staff) (string, error) {
	query := `
		UPDATE "staffs"
		SET "branch_id" = $1, "tariff_id" = $2, 
		"staff_type" = $3, "balance" = $4, 
		"name" = $5, "updated_at" = NOW()
		WHERE "id" = $6
		RETURNING "id"
	`

	result, err := s.db.Exec(context.Background(),
		query, req.BranchID, req.TariffID, req.Type, req.Balance, req.Name, req.ID)

	if err != nil {
		return "", fmt.Errorf("failed to update staff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID %s not found", req.ID)
	}

	return req.ID, nil
}
func (s *staffRepo) GetStaff(req *models.IdRequest) (*models.Staff, error) {
	query := `
		SELECT "id", "branch_id", "tariff_id", "staff_type", "name", "balance", "created_at", "updated_at"
		FROM "staffs"
		WHERE "id" = $1
	`

	staff := models.Staff{}

	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staff.ID,
		&staff.BranchID,
		&staff.TariffID,
		&staff.Type,
		&staff.Name,
		&staff.Balance,
		&staff.CreatedAt,
		&staff.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &models.Staff{}, fmt.Errorf("staff not found")
		}
		return &models.Staff{}, fmt.Errorf("failed to get staff: %w", err)
	}

	return &staff, nil
}

func (s *staffRepo) GetAllStaff(req *models.GetAllStaffRequest) (*models.GetAllStaff, error) {
	params := make(map[string]interface{})
	filter := ""

	query := `
		SELECT "id", "branch_id", "tariff_id", "staff_type", "name", "balance", "created_at", "updated_at"
		FROM "staffs"
	`
	if req.Name != "" {
		filter += ` WHERE "name" ILIKE '%' || :search || '%' `
		params["search"] = req.Name
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}

	offset := (req.Page - 1) * limit
	params["limit"] = limit
	params["offset"] = offset

	query = query + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &models.GetAllStaff{}
	resp.Staffs = make([]models.Staff, 0)
	count := 0
	for rows.Next() {
		var staff models.Staff
		count++
		err := rows.Scan(
			&staff.ID,
			&staff.BranchID,
			&staff.TariffID,
			&staff.Type,
			&staff.Name,
			&staff.Balance,
			&staff.CreatedAt,
			&staff.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		resp.Staffs = append(resp.Staffs, staff)
	}

	resp.Count = count
	return resp, nil
}

func (s *staffRepo) DeleteStaff(req *models.IdRequest) (string, error) {
	query := `
		DELETE FROM "staffs"
		WHERE "id" = $1
		RETURNING "id"
	`

	var deletedID string

	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(&deletedID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("staff not found")
		}
		return "", fmt.Errorf("failed to delete staff: %w", err)
	}

	return deletedID, nil
}
