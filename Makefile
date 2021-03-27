POSTGRES_URL = postgres://postgres:postgres@localhost:5432/app?sslmode=disable

dev:
	POSTGRES_URL=$(POSTGRES_URL) air

# example:
# $ make mg-create name=create_users_table
mg-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

mg-up:
	migrate -database $(POSTGRES_URL) -path db/migrations up

mg-down:
	migrate -database $(POSTGRES_URL) -path db/migrations down
