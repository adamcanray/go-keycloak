# Project Structure

this prooject structure is using domain driven design.

- /src
  - /config
  - /controllers -- contains the actual api’s
  - /services -- Handles business logic
  - /repositories -- Handles data access layer
  - /domains -- contains database related structs
- main.go

# Run with Docker

You can Create and Start container with this command `MODE=<mode> MYSQL_PROVIDER_HOST=<host> MYSQL_PROVIDER_PORT=<port> KEYCLOAK_CLIENT_ID=<id> KEYCLOAK_CLIENT_SECRET=<secret> KEYCLOAK_REALM=<realm> KEYCLOAK_HOST=<host> SERVICE_TAG_VERSION=<tag> docker compose up -d`
