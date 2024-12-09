CREATE TABLE IF NOT EXISTS storages (
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS genres (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS storage_genre (
    storage_id INT,
    genre_id INT,
    PRIMARY KEY (storage_id, genre_id),
    CONSTRAINT fk_storage FOREIGN KEY(storage_id) REFERENCES storages(id),
    CONSTRAINT fk_genre FOREIGN KEY(genre_id) REFERENCES genres(id)
);

