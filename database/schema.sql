-- Tabela de payout
CREATE TABLE IF NOT EXISTS payout_table (
    id UUID PRIMARY KEY,
    payout JSON NOT NULL,
    route TEXT NOT NULL
);