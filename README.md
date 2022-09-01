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

# Keycloak Server

- **[27/08/2022]**
  - I'm stuck. Keycloak Server is running well on my local machine (I run it on Docker for Mac), but it's going irrelevant since I hit an api for get access token for my specifics client with `client_credentials` grant type from localhost `127.0.0.1` with Postman, while the Keycloack Server is running on internal docker host [`host.docker.internal`](https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host), so token that I received is not valid cause it's a different host.
  - Depends on the issue that I explain before, I decide to setting up my own remote VPS using my raspi (Raspberry Pi 4 - 4GB - model B - Raspbian Buster OS - armv7 32-bit arch). When my raspi was ready (plug-in to power source), I connect from my Mac to my litle raspi using ssh. First think I setup a docker (not that easy) and run Keycloak Server container on docker, but keycloak does'nt have docker image that support linux with arm base arch, so I searching and I have some option:
    1. First, is using an good-people docker image that support linux with arm base arch (still have some issue).
    2. Second, is build my own keycloak container from this [reference](https://github.com/keycloak/keycloak-containers) (but, it took time).
- **[31/08/2022]**
  - I think I wanna create my own Keycloak Server from this [reference](https://www.keycloak.org/server/configuration).
