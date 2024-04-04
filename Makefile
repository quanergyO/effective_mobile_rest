.PHONY: run build

run:
	go run cmd/main.go

build:
	go build cmd/main.go

dbrun:
	docker run --name=effective-mobile-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

migrate_up:
	migrate -path ./schema/ -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema/ -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
