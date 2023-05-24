# Demo

## Plan

- [x] Scaffold React + Go gRPC
- [x] Create a basic CRUD frontend
- [x] Connect to postgres
- [ ] Create repository pattern
- [ ] Add tests

## Get up and running locally

```shell
make install
make serve/backend
make serve/frontend
make serve/db
```

## Get up and running in Docker

** Uninstall docker-desktop ()

```shell
brew install colima
colima start
docker-compose up --build
````