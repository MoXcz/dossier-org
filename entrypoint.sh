#!/bin/bash

echo "Running migrations..." && \
  goose up || \
  echo "Error when running migrations"

echo "Starting application..." && \
  exec /usr/local/bin/app || \
  echo "Error when starting application"
