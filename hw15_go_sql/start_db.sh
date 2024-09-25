#!/bin/bash -x

# Run the database using Docker
docker compose -f ./internal/storage/postgres/docker-compose.yml up