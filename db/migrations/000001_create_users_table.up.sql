CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    nip VARCHAR(13),
    role_id SMALLINT NULL,
    gender_id SMALLINT NULL,
    password VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_nip UNIQUE (nip)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_nip ON users (nip);