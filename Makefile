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

sqlc:
	sqlc generate

.PHONY:postgresql createDB dropDB migrateUp migrateUp sqlc