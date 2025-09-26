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
    data              JSONB,                          -- 非 indexed 参数原始 JSON
    created_at        TIMESTAMP DEFAULT NOW()         -- 插入时间
);
ALTER TABLE event_raw ADD CONSTRAINT uq_tx_log UNIQUE (tx_hash, log_index);
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
			event_signature, event_name, topic0, topic1, topic2, topic3, data, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		ev.TxHash, ev.LogIndex, ev.BlockNumber, ev.BlockTime, ev.ContractAddress,
		ev.EventSignature, ev.EventName, ev.Topic0, ev.Topic1, ev.Topic2, ev.Topic3, ev.Data, ev.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) InsertEventParsed(ctx context.Context, ep *EventParsed) error {
	tokenIDStr := ""
	if ep.TokenID != nil {
		tokenIDStr = ep.TokenID.String()
	}
	valueStr := ""
	if ep.Value != nil {
		valueStr = ep.Value.String()
	}
	_, err := s.db.ExecContext(ctx, `
        INSERT INTO contract_event_parsed (
            tx_hash, log_index, block_number, block_time, contract_address,
            event_name, from_address, to_address, token_id, value, metadata, created_at
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
    `,
		ep.TxHash, ep.LogIndex, ep.BlockNumber, ep.BlockTime, ep.ContractAddress, ep.EventName,
		ep.FromAddress, ep.ToAddress, tokenIDStr, valueStr, ep.Metadata, ep.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
