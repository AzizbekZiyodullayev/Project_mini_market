package db

import (
	"context"
	"file/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) *transactionRepo {
	return &transactionRepo{db: db}
}

func (t *transactionRepo) CreateTransaction(req *models.CreateTransaction) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO "transactions" ("id", "type", "amount", "source_type", "text", "sale_id", "staff_id", "created_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := t.db.Exec(context.Background(), query, id, req.Type, req.Amount, req.Source_type, req.Text, req.Sale_id, req.Staff_id)
	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %w", err)
	}

	return id, nil
}

func (t *transactionRepo) GetTransaction(req *models.IdRequest) (*models.Transaction, error) {
	var transaction models.Transaction

	query := `
		SELECT "id", "type", "amount", "source_type", "text", "sale_id", "staff_id", "created_at", "updated_at"
		FROM "transactions"
		WHERE "id" = $1
	`

	err := t.db.QueryRow(context.Background(), query, req.Id).Scan(
		&transaction.Id,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Source_type,
		&transaction.Text,
		&transaction.Sale_id,
		&transaction.Staff_id,
		&transaction.Created_at,
		&transaction.Updated_at,
	)
	if err != nil {

		return nil, fmt.Errorf("transaction not found")
	}

	return &transaction, nil
}

func (t *transactionRepo) GetAllTransaction(req *models.GetAllTransactionRequest) (*models.GetAllTransactionResponse, error) {
	var response models.GetAllTransactionResponse
	response.Transactions = make([]models.Transaction, 0)

	query := `
		SELECT id, type, amount, source_type, text, sale_id, staff_id, created_at, updated_at
		FROM transactions
		WHERE text ILIKE '%' || $1 || '%'
		LIMIT $2 OFFSET $3
	`

	rows, err := t.db.Query(context.Background(), query, req.Text, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Source_type,
			&transaction.Text,
			&transaction.Sale_id,
			&transaction.Staff_id,
			&transaction.Created_at,
			&transaction.Updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}
		response.Transactions = append(response.Transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over transactions: %w", err)
	}

	countQuery := `
		SELECT COUNT(*)
		FROM transactions
		WHERE text ILIKE '%' || $1 || '%'
	`
	err = t.db.QueryRow(context.Background(), countQuery, req.Text).Scan(&response.Count)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction count: %w", err)
	}

	return &response, nil
}

func (t *transactionRepo) UpdateTransaction(req *models.Transaction) (string, error) {
	query := `
		UPDATE "transactions"
		SET "type" = $1, "amount" = $2, "source_type" = $3, "text" = $4, "sale_id" = $5, "staff_id" = $6,  "updated_at" = NOW()
		WHERE "id" = $7
	`

	_, err := t.db.Exec(context.Background(), query, req.Type, req.Amount, req.Source_type, req.Text, req.Sale_id, req.Staff_id, req.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("transaction not found")
		}
		return "", fmt.Errorf("failed to update transaction: %w", err)
	}

	return req.Id, nil
}

func (t *transactionRepo) DeleteTransaction(req *models.IdRequest) (string, error) {
	query := `
		DELETE FROM transactions
		WHERE id = $1
	`

	result, err := t.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete transaction: %w", err)
	}

	affectedRows := result.RowsAffected()
	if affectedRows != 1 {
		return "", fmt.Errorf("transaction not found")
	}

	return req.Id, nil
}

// func (t *transactionRepo) GetTopStaffs(req *models.TopWorkerRequest) (resp *map[string]models.StaffTop, err error) {

// }
