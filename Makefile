postgresimage: ## add postgres image to docker
	docker pull postgres:13-alpine

postgres: ## setup postgres container
	docker run --name layouts-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=WEBdeveloepr1452 -d postgres:13-alpine

createdb:  ## Seting up db
	docker exec -it layouts-postgres createdb --username=root --owner=root layouts

dropdb:  ## Dropdown db
	docker exec -it layouts-postgres dropdb layouts

migrateup:  ## Dropdown db
	migrate -path db/migration -database "postgresql://root:WEBdeveloepr1452@localhost:5432/layouts?sslmode=disable" -verbose up

migratedown:  ## Dropdown db
	migrate -path db/migration -database "postgresql://root:WEBdeveloepr1452@localhost:5432/layouts?sslmode=disable" -verbose down

sqlc: ## sqlc generates type-safe code from SQL
	sqlc generate

.PHONY: postgresimage postgres createdb dropdb migrateup migratedown sqlc