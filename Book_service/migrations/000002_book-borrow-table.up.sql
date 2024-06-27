-- +migrate Up
CREATE TABLE IF NOT EXISTS borrow (
    user_id INT,
    book_title VARCHAR(255),
    page_readed INT,
    page_in_book INT,
    return_date DATE NOT NULL,
    return_status BOOLEAN DEFAULT false,
    issue_date DATE NOT NULL,
    is_active BOOLEAN
);