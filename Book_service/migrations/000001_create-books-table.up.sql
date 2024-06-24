-- +migrate Up
CREATE TABLE IF NOT EXISTS books (
    isbn SERIAL primary key,
    Title varchar UNIQUE,
    Author varchar NOT NULL,
    genres varchar NOT NULL,
    quantity varchar NOT NULL,
    publication_date varchar NOT NULL,
    next_available varchar NOT NULL,
    is_active BOOLEAN DEFAULT false
);
