-- +goose Up
CREATE TABLE users(
  id BIGINT GENERATED ALWAYS AS IDENTITY (CACHE 200) PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  encryptedPassword TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
