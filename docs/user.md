# `/user` Endpoints

## `POST /user` to create a new user:

Body of the request should include:

```json
{
  "name": "string",
  "email": "string",
  "password": "string"
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

To simulate an error:
```sh
curl --request POST \
  --url http://localhost:8080/user \
  --header 'content-type: application/json' \
  --data '{
  "name": "a",
  "email": "@mail.com",
  "password": "a1234567"
}'
```

The error response when the name, email, password are *incorrect*:

```json
{
  "error": [
    "name lenght should be at least 2 characters",
    "email is invalid"
  ]
}
```

The parameters for what is incorrect are the following:
- Name should be at last 2 characters
- Password should be at least 7 characters
- Email has to match: `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$` (start with lowercase characters and/or symbols followed by `@` and then more lowercase characters and/or symbols followed by a `.` and ending with 2-4 lowercase characters).

## `GET /user/{id}` to retrieve a single user by ID

The body does not require anything at the moment:

```sh
curl --request GET \
  --url http://localhost:8080/user/1 \
  --header 'content-type: application/json'
```

Which returns:

```json
{
  "id": 1,
  "name": "Pedro",
  "email": "pedro@mail.com",
  "encryptedpassword": "$2a$12$EmCGIgRvO9M0ZGeF.5D8GOjtqMw6yDgJukH9u1Kq5eTzcA1AZGTeq"
}
```

## `/GET /user

The body does not require anything at the moment:
```sh
curl --request GET \
  --url http://localhost:8080/user \
  --header 'content-type: application/json'
```

Which returns a list of objects of the users:

```json
[
  {
    "id": 1,
    "name": "Juans",
    "email": "j@a.com",
    "encryptedpassword": "$2a$12$9fxMcXqYtfgWgqKjdQiGrOQgn2jX83dOYwXuQPnF9nAtLIb4EDWnO"
  }
]
```
