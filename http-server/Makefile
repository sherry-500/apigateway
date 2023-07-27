migrateup:
	migrate -path db/migration -database "postgresql://postgres:eat12345@localhost:5432/idlmap?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:eat12345@localhost:5432/idlmap?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "D:\OS\apigateway\http-server:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: dropdb migrateup migratedown sqlc test