-- +goose Up
CREATE TABLE users (
    user_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role_id INT NOT NULL
        REFERENCES roles(role_id)
        ON UPDATE CASCADE  -- in case of: rename a role
        ON DELETE RESTRICT -- can’t delete a role while users point at it
);

-- +goose Down
DROP TABLE users;
