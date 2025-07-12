# dossier-org

API for dossier management

App:

```sh
docker build . -t dossier-org
```

Postgres:

```sh
docker run --name dossier-db -e POSTGRES_PASSWORD=<password> -p 5432:5432 -d postgres
```

Migrations:

```sh
goose up
```

Docker compose:

```sh
docker compose up
```

## Endpoints

| Method | Path         | Description                  |
| ------ | ------------ | ---------------------------- |
| GET    | `/user/{id}` | Retrieve a single user by ID |
| GET    | `/user`      | Retrieve a list of all users |
| POST   | `/user`      | Create a new user            |
| GET   | `/user/{id}/dossier`      | Get dossiers from user by ID|
| POST   | `/dossier`      | Create dossier and assign it to user |


# TO-DO

- [ ] Add auth
- [ ] First create dosser, then assign it?
- [ ] Add permissions
