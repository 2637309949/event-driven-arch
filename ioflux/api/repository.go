package main

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const migration = `
CREATE TABLE IF NOT EXISTS cqrs_trx (
	trxid BIGINT PRIMARY KEY,
	type INT NOT NULL DEFAULT 0,
	state INT NOT NULL DEFAULT 0,
	name VARCHAR NOT NULL DEFAULT '',
	progress INT NOT NULL DEFAULT 0,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO cqrs_trx (trxid, type, state, name, progress) VALUES
	(178953182742379521, 101, 1001, '接受请求', 1),
	(178953182742379521, 101, 1002, '正在锁定库存', 10),
	(178953182742379521, 101, 1003, '锁定库存成功', 20),
	(178953182742379521, 101, 1004, '生成待支付单', 30)
ON CONFLICT (trxid) DO NOTHING;
`

func MigrateDB(db *sql.DB) error {
	_, err := db.Exec(migration)
	return err
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) TrxByID(ctx context.Context, id int64) (Trx, error) {
	row := s.db.QueryRowContext(ctx, `SELECT trxid, type, state, name, progress, created_at FROM cqrs_trx WHERE trxid = $1`, id)
	trx, err := scanTrx(row)
	if err != nil {
		return Trx{}, err
	}

	return trx, nil
}

func (s *Repository) SaveTrx(ctx context.Context, id int64, updateFn func(trx *Trx)) (err error) {
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

	row := s.db.QueryRowContext(ctx, `SELECT trxid, type, state, name, progress, created_at FROM cqrs_trx WHERE trxid = $1 FOR UPDATE`, id)
	trx, err := scanTrx(row)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	updateFn(&trx)

	if trx.TrxId > 0 {
		_, err = tx.ExecContext(ctx, `UPDATE cqrs_trx SET state = $1, name = $2, progress = $3 WHERE trxid = $4`, trx.State, trx.Name, trx.Progress, trx.TrxId)
		if err != nil {
			return err
		}
	} else {
		_, err = tx.ExecContext(ctx, `INSERT INTO cqrs_trx (trxid, type, state, name, progress) VALUES ($1, $2, $3, $4, $5)`, id, trx.Type, trx.State, trx.Name, trx.Progress)
		if err != nil {
			return err
		}
	}

	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanTrx(s scanner) (Trx, error) {
	var type0, state, progress int
	var trxid int64
	var name string
	var createdAt time.Time

	err := s.Scan(&trxid, &type0, &state, &name, &progress, &createdAt)
	if err != nil {
		return Trx{}, err
	}

	return Trx{
		TrxId:     trxid,
		Type:      type0,
		State:     state,
		Name:      name,
		Progress:  progress,
		CreatedAt: createdAt,
	}, nil
}
