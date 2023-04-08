swag:
	swag init --parseDependency --parseInternal -g ./cmd/main.go

createContainer:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:13-alpine

createdb:
	docker exec -it postgres12 createdb --username=admin --owner=admin postgres

dropdb:
	docker exec -it postgres12 dropdb postgres

migrateup:
	migrate -path storage/postgre/migrations -database "postgresql://admin:admin@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path storage/postgre/migrations -database "postgresql://admin:admin@localhost:5432/postgres?sslmode=disable" -verbose down
