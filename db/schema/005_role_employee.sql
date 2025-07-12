-- +goose Up
CREATE TABLE employee_profiles (
    user_id BIGINT PRIMARY KEY
      REFERENCES users(user_id)
      ON DELETE CASCADE
    -- TODO: add any employee‑specific fields here, e.g. hire_date, department, etc.
    -- maybe delete, still thinking on how to manage permissions, but this seems
    -- promising
);

-- +goose Down
DROP TABLE employee_profiles;
