DB_URL=postgresql://myuser:mypassword@localhost:5432/authenticate?sslmode=disable

infras:
	docker-compose up -d

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: infras migrateup migratedown new_migration sqlc test