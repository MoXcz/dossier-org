-- +goose Up
CREATE TABLE role_permissions (
    role_id       INT NOT NULL
      REFERENCES roles(role_id)
      ON DELETE CASCADE,
    permission_id INT NOT NULL
      REFERENCES permissions(permission_id)
      ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- +goose Down
DROP TABLE role_permissions;
