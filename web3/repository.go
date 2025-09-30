package main

import (
	"context"
	"database/sql"
)

const migration = `
CREATE TABLE IF NOT EXISTS event_raw (
    id                BIGSERIAL PRIMARY KEY,          -- 表内唯一主键
    tx_hash           CHAR(66) NOT NULL,              -- 交易哈希
    log_index         INT NOT NULL,                   -- 交易内日志序号
    block_number      BIGINT NOT NULL,                -- 区块高度
    block_time        TIMESTAMP NOT NULL,             -- 区块时间
    contract_address  VARCHAR(42) NOT NULL,           -- 合约地址
    event_signature   VARCHAR(128) NOT NULL,          -- 事件签名 keccak256
    event_name        VARCHAR(64) NOT NULL,           -- 事件名称
    topic0            CHAR(66),                       -- indexed 参数1
    topic1            CHAR(66),                       -- indexed 参数2
    topic2            CHAR(66),                       -- indexed 参数3
    topic3            CHAR(66),                       -- indexed 参数4
    data              BYTEA,                          -- 非 indexed 参数原始 
    created_at        TIMESTAMP DEFAULT NOW()         -- 插入时间
);
ALTER TABLE event_raw ADD CONSTRAINT uq_raw_tx_log UNIQUE (tx_hash, log_index);
CREATE INDEX idx_event_block_number ON event_raw(block_number);
CREATE INDEX idx_event_contract_name ON event_raw(contract_address, event_name);
CREATE INDEX idx_event_topic0 ON event_raw(topic0);
CREATE INDEX idx_event_data_jsonb ON event_raw USING GIN (data);
CREATE TABLE IF NOT EXISTS event_parsed (
    id                BIGSERIAL PRIMARY KEY,
    tx_hash           CHAR(66) NOT NULL,
    log_index         INT NOT NULL,
    block_number      BIGINT NOT NULL,
    block_time        TIMESTAMP NOT NULL,
    contract_address  VARCHAR(42) NOT NULL,
    event_name        VARCHAR(64) NOT NULL,
    from_address      VARCHAR(42),
    to_address        VARCHAR(42),
    token_id          NUMERIC,                         -- 适配 ERC721 / ERC1155
    value             NUMERIC,                         -- 适配 ERC20 Transfer
    metadata          JSONB,                           -- 其他扩展字段
    created_at        TIMESTAMP DEFAULT NOW()
);
ALTER TABLE event_parsed ADD CONSTRAINT uq_parsed_tx_log UNIQUE (tx_hash, log_index);
CREATE INDEX idx_parsed_from ON event_parsed(from_address);
CREATE INDEX idx_parsed_to ON event_parsed(to_address);
CREATE INDEX idx_parsed_tokenid ON event_parsed(token_id);
CREATE INDEX idx_parsed_event_name ON event_parsed(contract_address, event_name);
CREATE TABLE IF NOT EXISTS event_stats (
    id                BIGSERIAL PRIMARY KEY,          -- 表内唯一主键
    event_name        VARCHAR(64) NOT NULL,           -- 事件名称
    event_label       VARCHAR(64) NOT NULL,           -- 事件标签
    event_count       INT NOT NULL,					  -- 事件统计
    created_at        TIMESTAMP DEFAULT NOW()         -- 插入时间
);
ALTER TABLE event_stats ADD CONSTRAINT uq_stats_ev_name UNIQUE (event_name);
CREATE INDEX idx_event_name ON event_stats(event_name);
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

func (s *Repository) InsertEventRaw(ctx context.Context, ev *EventRaw) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO event_raw (
			tx_hash, log_index, block_number, block_time, contract_address, 
			event_signature, event_name, topic0, topic1, topic2, topic3, data
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        ON CONFLICT ON CONSTRAINT uq_raw_tx_log DO NOTHING`,
		ev.TxHash, ev.LogIndex, ev.BlockNumber, ev.BlockTime, ev.ContractAddress,
		ev.EventSignature, ev.EventName, ev.Topic0, ev.Topic1, ev.Topic2, ev.Topic3, ev.Data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) InsertEventParsed(ctx context.Context, ep *EventParsed) error {
	_, err := s.db.ExecContext(ctx, `
        INSERT INTO event_parsed (
            tx_hash, log_index, block_number, block_time, contract_address,
            event_name, from_address, to_address, token_id, value, metadata
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
        ON CONFLICT ON CONSTRAINT uq_parsed_tx_log DO NOTHING`,
		ep.TxHash, ep.LogIndex, ep.BlockNumber, ep.BlockTime, ep.ContractAddress, ep.EventName,
		ep.FromAddress, ep.ToAddress, ep.TokenID, ep.Value, ep.Metadata,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) QueryEventStats(ctx context.Context) ([]EventStats, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT event_name, event_label, event_count
		FROM event_stats
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stats []EventStats
	for rows.Next() {
		var es EventStats
		err := rows.Scan(
			&es.EventName,
			&es.EventLabel,
			&es.EventCount,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, es)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}


func (s *Repository) UpsertEventStats(ctx context.Context, ep *EventStats) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO event_stats (event_name, event_label, event_count)
		VALUES ($1, $2, $3)
		ON CONFLICT (event_name) 
		DO UPDATE SET 
		event_count = event_stats.event_count + EXCLUDED.event_count`, ep.EventName, ep.EventLabel, ep.EventCount)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) UpdateEventStats(ctx context.Context, ep *EventStats) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO event_stats (event_name, event_label, event_count)
		VALUES ($1, $2, $3)
		ON CONFLICT (event_name) 
		DO UPDATE SET 
		event_count = EXCLUDED.event_count`, ep.EventName, ep.EventLabel, ep.EventCount)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) QueryActiveUsers(ctx context.Context) (uint64, error) {
	query := `
		SELECT COUNT(DISTINCT addr) AS active_users
		FROM (
			SELECT from_address AS addr
			FROM event_parsed
			WHERE from_address IS NOT NULL
			  AND block_time >= NOW() - INTERVAL '1 month'
			UNION
			SELECT to_address AS addr
			FROM event_parsed
			WHERE to_address IS NOT NULL
			  AND block_time >= NOW() - INTERVAL '1 month'
		) AS u
	`
	var count uint64
	err := s.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
