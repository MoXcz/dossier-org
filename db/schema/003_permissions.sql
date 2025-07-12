-- +goose Up
CREATE TABLE permissions (
    permission_id  SERIAL PRIMARY KEY,
    name           TEXT NOT NULL UNIQUE,    -- e.g. 'user.create', 'report.view'
    description    TEXT
);

-- +goose Down
DROP TABLE permissions;
