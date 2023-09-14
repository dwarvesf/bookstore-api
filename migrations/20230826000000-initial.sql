
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS topics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    topic_id INTEGER,
    author VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS favorites (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE books ADD CONSTRAINT fk_books_topics FOREIGN KEY (topic_id) REFERENCES topics (id);
ALTER TABLE favorites ADD CONSTRAINT fk_favorites_books FOREIGN KEY (book_id) REFERENCES books (id);
ALTER TABLE favorites ADD CONSTRAINT fk_favorites_users FOREIGN KEY (user_id) REFERENCES users (id);

CREATE UNIQUE INDEX IF NOT EXISTS idx_favorites_book_user ON favorites (book_id, user_id);

INSERT INTO topics (name, code, created_at, updated_at)
VALUES
    ('Programming', 'programming', '2023-09-14 10:00:00', '2023-09-14 10:30:00'),
    ('Database', 'database', '2023-09-14 11:00:00', '2023-09-14 11:30:00'),
    ('DevOps', 'devops', '2023-09-14 12:00:00', '2023-09-14 12:30:00');

-- +migrate Down
DROP TABLE IF EXISTS favorites;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS topics;
DROP TABLE IF EXISTS users;