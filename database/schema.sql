-- Tabela de payout
CREATE TABLE IF NOT EXISTS payouts (
    id UUID PRIMARY KEY,
    payout JSON NOT NULL,
    route TEXT NOT NULL
);