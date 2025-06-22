#!/bin/bash

echo "Waiting for PostgreSQL..."
./wait-for-it.sh db:5432 --timeout=30 --strict -- echo "PostgreSQL is up."

echo "Running migrations..."
goose -dir /db/schema up

echo "Starting application..."
exec /usr/local/bin/app
