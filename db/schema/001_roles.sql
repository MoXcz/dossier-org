-- +goose Up
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE, -- 'admin', 'customer'
    description TEXT NOT NULL
);

-- +goose Down
DROP TABLE roles;
