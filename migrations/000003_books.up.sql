CREATE TABLE IF NOT EXISTS books (
    id serial PRIMARY KEY,
    storage_id INT NOT NULL,
    genre VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    comment TEXT,
    CONSTRAINT fk_storage FOREIGN KEY(storage_id) REFERENCES storage(id)
);