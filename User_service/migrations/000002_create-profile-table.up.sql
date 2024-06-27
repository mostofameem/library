-- +migrate Up
CREATE TABLE IF NOT EXISTS user_profile (
    user_id INT,
    book_title varchar ,
    page_readed INT,
    page_in_book INT,
    return_date varchar NOT NULL,
    return_status varchar NOT NULL
);