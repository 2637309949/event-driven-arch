package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const migration = `
CREATE TABLE IF NOT EXISTS cqrs_order (
	order_id BIGINT PRIMARY KEY,
	user_id BIGINT NOT NULL DEFAULT 0,
	state INT NOT NULL DEFAULT 0,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO cqrs_order (order_id, user_id, state) VALUES
	(178953182742379520, 20251101, 1001)
ON CONFLICT (order_id) DO NOTHING;
`

func MigrateDB(db *sql.DB) {
	_, err := db.Exec(migration)
	if err != nil {
		fmt.Println(err)
	}
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) OrderByID(ctx context.Context, id int) (Order, error) {
	row := s.db.QueryRowContext(ctx, `SELECT order_id, user_id, state, created_at FROM cqrs_order WHERE order_id = $1`, id)
	order, err := scanOrder(row)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (s *Repository) CreateOrder(ctx context.Context, order *Order) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO cqrs_order (order_id, user_id, state) VALUES ($1, $2, $3)`, order.OrderId, order.UserId, order.State)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) UpdatePost(ctx context.Context, id int, updateFn func(order *Order)) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			txErr := tx.Rollback()
			if txErr != nil {
				err = txErr
			}
		}
	}()

	row := s.db.QueryRowContext(ctx, `SELECT order_id, user_id, state, created_at FROM cqrs_order WHERE id = $1 FOR UPDATE`, id)
	order, err := scanOrder(row)
	if err != nil {
		return err
	}

	updateFn(&order)

	_, err = tx.ExecContext(ctx, `UPDATE cqrs_order SET state = $1 WHERE order_id = $2`, order.State, order.OrderId)
	if err != nil {
		return err
	}

	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanOrder(s scanner) (Order, error) {
	var orderId, userId int64
	var state int
	var createdAt time.Time

	err := s.Scan(&orderId, &userId, &state, &createdAt)
	if err != nil {
		return Order{}, err
	}

	return Order{
		OrderId:   orderId,
		UserId:    userId,
		State:     state,
		CreatedAt: createdAt,
	}, nil
}
