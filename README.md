# Getting Started

## Run on your local

You simply just create `.env` file in root of the project, then run this command `go run main.go`

## Run with docker

You can create `.env.dev` file first in root of the project, then run this command `docker compose --env-file ./.env.dev up -d`

# Project Structure

this prooject structure is using domain driven design.

- /src
  - /config
  - /controllers -- contains the actual api’s
  - /services -- Handles business logic
  - /repositories -- Handles data access layer
  - /domains -- contains database related structs
- main.go
