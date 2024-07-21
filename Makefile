run:
	go run cmd/main.go

db/migrate:
	sql-migrate up

g/migrate:
	sql-migrate new ${name}
