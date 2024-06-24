-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    name varchar NOT NULL ,
    email varchar UNIQUE,
    password varchar NOT NULL,
    type varchar NOT NULL,
    is_active BOOLEAN DEFAULT false
);
-- +migrate Down
DROP TABLE IF EXISTS "public"."users";
