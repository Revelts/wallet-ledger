CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    balance NUMERIC(36,18) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS ledgers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    amount NUMERIC(36,18) NOT NULL,
    balance NUMERIC(36,18) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('deposit', 'withdraw', 'transfer_in', 'transfer_out')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );

CREATE INDEX idx_ledgers_user_id ON ledgers(user_id);
CREATE INDEX idx_ledgers_created_at ON ledgers(created_at);

CREATE TABLE IF NOT EXISTS transfers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    amount NUMERIC(36,18) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
    );

CREATE INDEX idx_transfers_sender_id ON transfers(sender_id);
CREATE INDEX idx_transfers_receiver_id ON transfers(receiver_id);
CREATE INDEX idx_transfers_created_at ON transfers(created_at);


INSERT INTO users (id, name, email) VALUES
                                        (gen_random_uuid(), 'Alice Doe', 'alice@example.com'),
                                        (gen_random_uuid(), 'Bob Smith', 'bob@example.com');

INSERT INTO wallets (id, user_id, balance)
SELECT gen_random_uuid(), id, 0 FROM users WHERE email IN ('alice@example.com', 'bob@example.com');
