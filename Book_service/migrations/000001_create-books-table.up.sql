-- +migrate Up
CREATE TABLE IF NOT EXISTS books (
    isbn SERIAL primary key,
    Title varchar UNIQUE,
    total_page INT,
    Author varchar NOT NULL,
    genres varchar NOT NULL,
    quantity int NOT NULL,
    publication_date varchar NOT NULL,
    next_available varchar NOT NULL,
    is_active BOOLEAN DEFAULT false
);
