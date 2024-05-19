CREATE TABLE IF NOT EXISTS patients (
    id INT PRIMARY KEY,
    name VARCHAR(50),
    phone_number VARCHAR(15),
    birth_date DATE NULL,
    gender_id SMALLINT NULL,
    identity_card_scan_img VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_phone_number ON patients (phone_number);