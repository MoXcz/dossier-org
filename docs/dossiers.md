# Dossiers

For Many-to-many relationships on dossiers (currently a dossier can only assigned to one employee):
```sql
CREATE TABLE dossiers (
    dossier_id   BIGSERIAL PRIMARY KEY,
    title         TEXT     NOT NULL,
    data          JSONB    NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
    -- remove relation to users
);

-- and create a dossiers-to-users table (or just leave it at the dossiers table)
CREATE TABLE dossier_assignments (
    dossier_id BIGINT NOT NULL
      REFERENCES dossiers(dossier_id)
      ON DELETE CASCADE,
    user_id    BIGINT NOT NULL
      REFERENCES users(user_id)
      ON DELETE CASCADE,
    PRIMARY KEY (dossier_id, user_id),
);
```

