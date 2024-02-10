-- Create promo_codes table
CREATE TABLE promo_codes (
    promo_id UUID PRIMARY KEY,
    promo_name TEXT UNIQUE NOT NULL,
    promo TEXT NOT NULL,
    active BOOLEAN NOT NULL
);