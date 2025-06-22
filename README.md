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

## User Object

To create a user:

Body of the request should include:

```json
{
  "name": "string",     // Name of user
  "email": "string",    // Email address of user
  "password": "string"  // Password of user
}
```

With `curl`:

```sh
curl --request POST \
  --url http://localhost:8080/user \
  --header 'content-type: application/json' \
  --data '{
  "name": "Pedro",
  "email": "pedro@mail.com",
  "password": "Pedro123!"
}'
```

And then it returns:

```json
{
  "id": 1,
  "name": "Pedro",
  "email": "pedro@mail.com",
  "encryptedpassword": "$2a$12$EmCGIgRvO9M0ZGeF.5D8GOjtqMw6yDgJukH9u1Kq5eTzcA1AZGTeq"
}
```

The error response (TODO):

```json
{
  "error": ["string",]     // Description of the errors
}
```

## References

1. https://github.com/vishnubob/wait-for-it
