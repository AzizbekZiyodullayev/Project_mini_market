package db

import (
	"context"
	"database/sql"
	"errors"
	"file/models"
	"file/pkg/helper"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *saleRepo {
	return &saleRepo{db: db}
}

func (c *saleRepo) CreateSale(req *models.CreateSales) (string, error) {
	id := uuid.NewString()

	query := `
        INSERT INTO "sales" 
        ("id", "client_name", "branch_id", "shop_assistant_id",
        "cashier_id", "price", "payment_type", "created_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
    `

	var err error
	var insertedID string

	if req.Cashier_id != "" && req.Shop_assistant_id != "" {
		_, err = c.db.Exec(context.Background(),
			query,
			id,
			req.Client_name,
			req.Branch_id,
			req.Shop_assistant_id,
			req.Cashier_id,
			req.Price,
			req.Payment_Type,
		)
		insertedID = id
	} else if req.Cashier_id != "" {
		_, err = c.db.Exec(context.Background(),
			query,
			id,
			req.Client_name,
			req.Branch_id,
			nil,
			req.Cashier_id,
			req.Price,
			req.Payment_Type,
		)
		insertedID = req.Cashier_id
	} else if req.Shop_assistant_id != "" {
		_, err = c.db.Exec(context.Background(),
			query,
			id,
			req.Client_name,
			req.Branch_id,
			req.Shop_assistant_id,
			nil,
			req.Price,
			req.Payment_Type,
		)
		insertedID = req.Shop_assistant_id
	} else {
		return "", errors.New("either cashier_id or shop_assistant_id should be provided")
	}

	if err != nil {
		return "", fmt.Errorf("failed to create sale: %w", err)
	}

	return insertedID, nil
}

func (c *saleRepo) GetSale(req *models.IdRequest) (resp *models.Sales, err error) {
	query := `
    SELECT "id", "client_name", "branch_id", "shop_assistant_id",
    "cashier_id", "price", "payment_type", "status", "created_at", "updated_at"
    FROM "sales" WHERE "id" = $1
    `
	var cashier_id sql.NullString
	var shop_assistant_id sql.NullString
	sale := models.Sales{}
	err = c.db.QueryRow(context.Background(), query, req.Id).Scan(
		&sale.Id,
		&sale.Client_name,
		&sale.Branch_id,
		&shop_assistant_id,
		&cashier_id,
		&sale.Price,
		&sale.Payment_Type,
		&sale.Status,
		&sale.Created_at,
		&sale.Updated_at,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("sale not found")
		}
		return nil, fmt.Errorf("failed to get sale: %w", err)
	}

	if cashier_id.Valid {
		sale.Cashier_id = cashier_id.String
	}
	if shop_assistant_id.Valid {
		sale.Shop_assistant_id = shop_assistant_id.String
	}

	return &sale, nil
}

func (c *saleRepo) GetAllSale(req *models.GetAllSalesRequest) (*models.GetAllSalesResponse, error) {
	params := make(map[string]interface{})
	filter := ""

	var cashier_id sql.NullString
	var shop_assistant_id sql.NullString

	selectQuery := `
    SELECT "id", "client_name", "branch_id", "shop_assistant_id",
    "cashier_id", "price", "payment_type", "status", "created_at", "updated_at"
    FROM "sales"
    `

	if req.Client_name != "" {
		filter += ` WHERE "client_name" ILIKE '%' || :search || '%' `
		params["search"] = req.Client_name
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

	query := selectQuery + filter + " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := c.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	resp := &models.GetAllSalesResponse{}
	resp.Sales = make([]models.Sales, 0)
	count := 0
	for rows.Next() {
		var sale models.Sales
		count++
		err := rows.Scan(
			&sale.Id,
			&sale.Client_name,
			&sale.Branch_id,
			&shop_assistant_id,
			&cashier_id,
			&sale.Price,
			&sale.Payment_Type,
			&sale.Status,
			&sale.Created_at,
			&sale.Updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if cashier_id.Valid {
			sale.Cashier_id = cashier_id.String
		}
		if shop_assistant_id.Valid {
			sale.Shop_assistant_id = shop_assistant_id.String
		}

		resp.Sales = append(resp.Sales, sale)
	}

	resp.Count = count
	return resp, nil
}

func (c *saleRepo) UpdateSale(req *models.Sales) (string, error) {
	query := `
	UPDATE "sales" SET
	"client_name" = $1,
	"branch_id" = $2,
	"shop_assistant_id" = $3,
	"cashier_id" = $4,
	"price" = $5,
	"payment_type" = $6,
	"updated_at" = NOW()
	WHERE "id" = $7
	RETURNING "id"
	`

	var updatedID string
	err := error(nil)

	if req.Shop_assistant_id != "" && req.Cashier_id != "" {
		err = c.db.QueryRow(
			context.Background(),
			query,
			req.Client_name,
			req.Branch_id,
			req.Shop_assistant_id,
			req.Cashier_id,
			req.Price,
			req.Payment_Type,
			req.Id,
		).Scan(&updatedID)
	} else if req.Shop_assistant_id == "" {
		err = c.db.QueryRow(
			context.Background(),
			query,
			req.Client_name,
			req.Branch_id,
			nil,
			req.Cashier_id,
			req.Price,
			req.Payment_Type,
			req.Id,
		).Scan(&updatedID)
	} else if req.Cashier_id == "" {
		err = c.db.QueryRow(
			context.Background(),
			query,
			req.Client_name,
			req.Branch_id,
			req.Shop_assistant_id,
			nil,
			req.Price,
			req.Payment_Type,
			req.Id,
		).Scan(&updatedID)
	}

	if err != nil {
		return "", fmt.Errorf("failed to update sale: %w", err)
	}

	return updatedID, nil
}

func (c *saleRepo) DeleteSale(req *models.IdRequest) (resp string, err error) {
	query := `DELETE FROM "sales" WHERE "id" = $1`

	resul, err := c.db.Exec(context.Background(), query, req.Id)
	if err != nil {
		return "", fmt.Errorf("failed to delete sale: %w", err)
	}

	if resul.RowsAffected() == 0 {
		return "", fmt.Errorf("sale with ID %s not found", req.Id)
	}

	return req.Id, nil
}
