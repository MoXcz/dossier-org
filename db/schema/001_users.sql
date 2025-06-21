-- +goose Up
CREATE TABLE users(
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  encryptedPassword TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
