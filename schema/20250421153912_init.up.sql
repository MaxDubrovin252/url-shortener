CREATE TABLE url(
    id SERIAL PRIMARY KEY,
    url varchar(255) NOT NULL,
    alias varchar(255) NOT NULL UNIQUE
);
CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);