-- Create users table
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    telegram_id VARCHAR(255) NOT NULL UNIQUE,
    invite_code_id UUID REFERENCES invite_codes(code_id) NOT NULL
);

