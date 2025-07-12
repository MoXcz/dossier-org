-- +goose Up
CREATE TABLE dossiers (
    dossier_id   BIGSERIAL PRIMARY KEY,
    title         TEXT      NOT NULL,
    data          JSONB     NOT NULL,
    assigned_to   BIGINT    NOT NULL
      REFERENCES users(user_id)
      ON DELETE RESTRICT,   -- if you delete an employee first, dossiers must be reassigned or removed
      -- Note that users SHOULD NEVER need to be deleted in the first place
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE dossiers;
