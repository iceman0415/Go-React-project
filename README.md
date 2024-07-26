# Go + React example

A simple Golang REST API + React UI (node server).

![alt text](https://raw.githubusercontent.com/bogdanguranda/go-react-example/master/dashboard.png)

## Installation

After cloning the repo, assuming you have `docker` installed on your machine, do the following steps in the project folder:

1. Set the env var `MYSQL_PASS` to a password of choice for MySQL. 
2. Run `docker-compose up` to spin all services.
3. Start a browser and check the UI at `localhost:3000`.

Note 1: `MySQL server` will run on port `3306`, the `Go API` on port `8080` and `React UI` (node server) on port `3000`.

Note 2: to run services individually, check the `docker-compose.yml` to see the commands needed or run `docker-compose up <name-of-service>`.

## Tests

Steps:
1. API: generate Go mocks with `go generate ./...` and then run all Go tests with: `go test -v ./...`, assuming you have Go installed.
2. UI: go to the root folder with `cd ui` and then run `npm test`, assuming you have node installed.
