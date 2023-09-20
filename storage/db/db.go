package db

import (
	"context"
	"file/config"
	"file/storage"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db          *pgxpool.Pool
	branches    *branchRepo
	tariffs     *staffTarifRepo
	staffes     *staffRepo
	sales       *saleRepo
	transaction *transactionRepo
	// staffTarifs *staffTarifRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnections

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (b *store) Branch() storage.BranchesI {
	if b.branches == nil {
		b.branches = NewBranchRepo(b.db)
	}
	return b.branches
}

func (b *store) Tariff() storage.TariffsI {
	if b.tariffs == nil {
		b.tariffs = NewStaffTarifRepo(b.db)
	}
	return b.tariffs
}

func (b *store) Staff() storage.StaffesI {
	if b.staffes == nil {
		b.staffes = NewStaffRepo(b.db)
	}
	return b.staffes
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Sales() storage.SalesI {
	if s.sales == nil {
		s.sales = NewSaleRepo(s.db)
	}
	return s.sales
}

func (s *store) Transaction() storage.TransactionI {
	if s.transaction == nil {
		s.transaction = NewTransactionRepo(s.db)
	}
	return s.transaction
}
