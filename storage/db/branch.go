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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (b *branchRepo) CreateBranch(req *models.CreateBranch) (string, error) {
	id := uuid.NewString()
	yearNow := time.Now().Year()
	year := yearNow - req.FoundedAt

	query := `
		INSERT INTO branches(id, name, adress, year, founded_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		year,
		req.FoundedAt,
		time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create branch: %w", err)
	}

	return id, nil
}

func (b *branchRepo) GetBranch(req *models.IdRequest) (resp *models.Branch, err error) {
	query := `SELECT id, name, adress, year, founded_at, created_at, updated_at FROM branches WHERE id=$1`
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	branch := models.Branch{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.Year,
		&branch.FoundedAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("sale not found")
		}
		return nil, fmt.Errorf("failed to get sale: %w", err)
	}
	branch.CreatedAt = createdAt.Format(time.RFC3339)
	if updatedAt.Valid {
		branch.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &branch, nil
}

func (b *branchRepo) UpdateBranch(req *models.Branch) (string, error) {
	yearNow := time.Now().Year()
	year := yearNow - req.FoundedAt
	query := `UPDATE branches SET name = $1, adress = $2, year = $3, founded_at = $4, updated_at = NOW() WHERE id = $5 RETURNING id`
	result, err := b.db.Exec(context.Background(), query, req.Name, req.Address, year, req.FoundedAt, req.ID)
	if err != nil {
		return "", fmt.Errorf("failed to update branch: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch with ID %s not found", req.ID)
	}

	return req.ID, nil
}

func (b *branchRepo) GetAllBranch(req *models.GetAllBranchRequest) (*models.GetAllBranch, error) {
	params := make(map[string]interface{})
	filter := ""
	offset := (req.Page - 1) * req.Limit
	createdAt := time.Time{}
	updatedAt := sql.NullTime{}

	s := `SELECT id, name, adress, year, founded_at, created_at, updated_at FROM branches`

	if req.Name != "" {
		filter += ` WHERE name ILIKE '%' || :search || '%' `
		params["search"] = req.Name
	}

	limit := fmt.Sprintf(" LIMIT %d", req.Limit)
	offsetQ := fmt.Sprintf(" OFFSET %d", offset)
	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp := &models.GetAllBranch{}
	resp.Branches = make([]models.Branch, 0)
	count := 0
	for rows.Next() {
		var branch models.Branch
		count++
		err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.Year,
			&branch.FoundedAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		branch.CreatedAt = createdAt.Format(time.RFC3339)
		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}
		resp.Branches = append(resp.Branches, branch)
	}

	resp.Count = count
	return resp, nil
}

func (b *branchRepo) DeleteBranch(req *models.IdRequest) (resp string, err error) {
	query := `DELETE FROM branches WHERE id = $1 RETURNING id`

	result, err := b.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch with ID %s not found", req.Id)

	}

	return req.Id, nil
}
