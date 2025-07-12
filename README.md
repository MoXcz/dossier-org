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


