CREATE TABLE IF NOT EXISTS books(
    id serial PRIMARY KEY,
    genre VARCHAR(255),
    title VARCHAR(255),
    author VARCHAR(255),
    comment TEXT
);