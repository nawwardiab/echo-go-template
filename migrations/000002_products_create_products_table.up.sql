CREATE TABLE IF NOT EXISTS products (
    id            SERIAL PRIMARY KEY,
    title         TEXT    NOT NULL,
    year          INT     NOT NULL CHECK (year >= 0),
    artist        TEXT    NOT NULL,
    img           TEXT,
    price         NUMERIC(12,2) NOT NULL CHECK (price >= 0),
    genre         TEXT
);