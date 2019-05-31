### Prerequisites
* docker
* docker-compose
* make

### Setup instructions
Execute the following commands: 
* `make up` to boot the environment
* `make migrate` to migrate the database. Make sure the database inside the container is ready to receive connections
* `make up` to start the api 

### Tests
* `make test` runs tests against a different database

### Tear down
* `make down` to tear down the environment

### Notable libraries used
* gorilla/mux
* gorm
* testify/assert
