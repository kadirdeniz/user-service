# User Service

User service is a basic service that has CRUD operation endpoints. 

## Installation

```bash
# create mongodb docker container
$ make docker-run-mongo

# create redis docker container
$ make docker-run-redis

# run application
$ make run
```

## Architectural Decisions

This repository is created with using `Test Driven Development`.`Dockertest` package for integration testing with `mongodb` and `redis`, `Ginkgo` was used for increase the understanding of test codes. All test can be run with using `make tests` command. Before run the tests mongodb and redis needs to be created.


## Tech Stack

Golang, MongoDB, Redis, Docker, Pactflow, Dockertest

## Contributing

Pull requests are welcome.
