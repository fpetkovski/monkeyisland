### Prerequisites
* docker
* docker-compose
* make

### Setup instructions
Execute the following commands: 
* `make up` to boot the environment
* `make migrate` to migrate the database. Make sure the database inside the mysql container is ready to receive connections
* `make up` to start the api 

The API should be available on http://localhost:8080

### Tests
* `make test` runs tests against a testing database

### API Documentation
Swagger-based documentation is available on http://localhost:8082

### Tear down
* `make down` to tear down the environment

### Notable libraries used
* gorilla/mux
* jinzhu/gorm
* testify/assert

