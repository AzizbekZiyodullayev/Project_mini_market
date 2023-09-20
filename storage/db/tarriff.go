package db

import (
	"context"
	"database/sql"
	"file/models"
	"file/pkg/helper"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffTarifRepo struct {
	db *pgxpool.Pool
}

func NewStaffTarifRepo(db *pgxpool.Pool) *staffTarifRepo {
	return &staffTarifRepo{db: db}
}

func (s *staffTarifRepo) CreateStaffTarif(req *models.CreateStaffTarif) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "tariffs" 
		("id", "name", "type", "amount_for_cash", "amount_for_card", "created_at")
		VALUES 
		($1, $2, $3, $4, $5, NOW())
	`

	_, err := s.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff tariff: %w", err)
	}

	return id, nil
}

func (s *staffTarifRepo) GetStaffTarif(req *models.IdRequest) (resp *models.StaffTarif, err error) {
	query := `
		SELECT  "id", "name", "type", "amount_for_cash", "amount_for_card", "created_at", "updated_at"
		FROM "tariffs" WHERE "id" = $1
	`
	var (
		created_at time.Time
		updated_at sql.NullTime
	)

	tariff := models.StaffTarif{}
	err = s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&tariff.Id,
		&tariff.Name,
		&tariff.Type,
		&tariff.AmountForCard,
		&tariff.AmountForCash,
		&created_at,
		&updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staff tariff not found")
		}
		return nil, fmt.Errorf("failed to get staff tariff: %w", err)
	}

	tariff.CreatedAt = created_at.Format(time.RFC3339)
	if updated_at.Valid {
		tariff.UpdatedAt = updated_at.Time.Format(time.RFC3339)
	}

	return &tariff, nil
}
func (s *staffTarifRepo) GetAllStaffTarif(req *models.GetAllStaffTarifRequest) (resp *models.GetAllStaffTarif, err error) {
	params := make(map[string]interface{})
	filter := ""
	created_at := time.Time{}
	updated_at := sql.NullTime{}

	sekect := `
		SELECT
		"id", "name", "type", "amount_for_cash", "amount_for_card", "created_at", "updated_at"
		FROM "tariffs"
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

	query := sekect + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := s.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp = &models.GetAllStaffTarif{}
	resp.StaffTarifs = make([]models.StaffTarif, 0)
	count := 0
	for rows.Next() {
		var tariff models.StaffTarif
		count++
		err := rows.Scan(
			&tariff.Id,
			&tariff.Name,
			&tariff.Type,
			&tariff.AmountForCard,
			&tariff.AmountForCash,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tariff.CreatedAt = created_at.Format(time.RFC3339)
		if updated_at.Valid {
			tariff.UpdatedAt = updated_at.Time.Format(time.RFC3339)
		}
		resp.StaffTarifs = append(resp.StaffTarifs, tariff)
	}

	resp.Count = count
	return resp, nil
}

func (s *staffTarifRepo) UpdateStaffTarif(req *models.StaffTarif) (string, error) {
	query := `
		UPDATE "tariffs"
		SET "name" = $1, "type" = $2, "amount_for_cash" = $3, "amount_for_card" = $4, "updated_at" = NOW()
		WHERE "id" = $5
	`

	result, err := s.db.Exec(context.Background(), query,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
		req.Id,
	)
	if err != nil {
		return "", fmt.Errorf("failed to update staff tariff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff tariff with ID %s not found", req.Id)
	}

	return req.Id, nil
}

func (s *staffTarifRepo) DeleteStaffTarif(req *models.IdRequest) (string, error) {
	query := `
		DELETE FROM "tariffs"
		WHERE "id" = $1
	`

	result, err := s.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete staff tariff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff tariff with ID %s not found", req.Id)
	}

	return req.Id, nil
}
