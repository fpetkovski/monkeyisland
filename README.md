### Prerequisites
* docker
* docker-compose
* make

### Setup instructions
Execute the following commands: 
* `make up` to boot the environment
* `make migrate` to migrate the database
* `make up` to start the api 

### Tests
* `make test` runs tests against a different database

### Notable libraries used
* gorilla/mux
* gorm
* testify/assert
