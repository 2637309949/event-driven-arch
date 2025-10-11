package main

import (
	"context"
	"database/sql"
)

const migration = `
CREATE TABLE IF NOT EXISTS cqrs_file (
	file_id     BIGINT PRIMARY KEY,
	save_path   VARCHAR(255) NOT NULL,  
	new_name    VARCHAR(128) NOT NULL, 
	orig_name   VARCHAR(255) NOT NULL,  
	ext         VARCHAR(16)  NOT NULL,   
	mime_type   VARCHAR(64)  NOT NULL,  
	created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO cqrs_file (file_id, save_path, new_name, orig_name, ext, mime_type)
VALUES (764428779249475584, '/uploads/2025/10', 'example.txt', '原始文件.txt', '.txt', 'text/plain')
ON CONFLICT (file_id) DO NOTHING;
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

func (s *Repository) FileByID(ctx context.Context, id int64) (File, error) {
	row := s.db.QueryRowContext(ctx, `
		SELECT file_id, save_path, new_name, orig_name, ext, mime_type, created_at
		FROM cqrs_file WHERE file_id = $1
	`, id)
	file, err := scanFile(row)
	if err != nil {
		return File{}, err
	}
	return file, nil
}

func (s *Repository) CreateFile(ctx context.Context, file *File) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO cqrs_file (file_id, save_path, new_name, orig_name, ext, mime_type)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, file.FileId, file.SavePath, file.NewName, file.OrigName, file.Ext, file.MimeType)
	return err
}

func (s *Repository) UpdateFile(ctx context.Context, id int64, updateFn func(file *File)) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	row := tx.QueryRowContext(ctx, `
		SELECT file_id, save_path, new_name, orig_name, ext, mime_type, created_at
		FROM cqrs_file WHERE file_id = $1 FOR UPDATE
	`, id)

	file, err := scanFile(row)
	if err != nil {
		return err
	}

	updateFn(&file)

	_, err = tx.ExecContext(ctx, `
		UPDATE cqrs_file 
		SET save_path = $1, new_name = $2, orig_name = $3, ext = $4, mime_type = $5
		WHERE file_id = $6
	`, file.SavePath, file.NewName, file.OrigName, file.Ext, file.MimeType, file.FileId)

	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanFile(s scanner) (File, error) {
	var f File
	err := s.Scan(&f.FileId, &f.SavePath, &f.NewName, &f.OrigName, &f.Ext, &f.MimeType, &f.CreatedAt)
	if err != nil {
		return File{}, err
	}
	return f, nil
}
