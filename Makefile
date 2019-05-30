.ONESHELL:

up:
	docker-compose -f ./build/docker-compose.yaml up --build -d

down:
	docker-compose -f ./build/docker-compose.yaml down

test: up
	docker exec app go run cmd/migrations/migrate_test_db.go
	docker exec app go test ./pkg/...

migrate:
	docker exec app go run cmd/migrations/migrate_db.go

start:
	docker exec app go run ./pkg/main.go