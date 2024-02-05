-- Create invite_codes table
CREATE TABLE invite_codes (
    code_id UUID PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    active BOOLEAN NOT NULL
);