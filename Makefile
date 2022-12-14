postgresql:
	docker run --name postgre14.5 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createDB:
	docker exec -it postgre14.5 createdb --username=root --owner=root simple_bank

dropDB:
	docker exec -it postgre14.5 dropdb simple_bank

migrateUp:
	./migrate --path /home/nima/GolandProjects/excercise/db/migration --database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrateDown:
	./migrate --path /home/nima/GolandProjects/excercise/db/migration --database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

migrateUpW:
	migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -path ./db/migration up

migrateDownW:
	migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -path ./db/migration down

sqlc:
	sqlc generate

sqlcW:
	docker run --rm -v "C:\Users\Nima\GolandProjects\simpleBank:/src" -w /src kjconroy/sqlc generate

test:
	go test -v ./...

.PHONY:postgresql createDB dropDB migrateUp migrateDownW migrateDownW migrateUp sqlc test