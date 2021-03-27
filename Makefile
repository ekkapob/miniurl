POSTGRES_URL = postgres://postgres:postgres@localhost:5432/app?sslmode=disable
URL_EXPIRE_SECONDS = 604800
HOSTNAME = http://127.0.0.1:8000

dev:
	HOSTNAME=$(HOSTNAME) \
	URL_EXPIRE_SECONDS=$(URL_EXPIRE_SECONDS) \
 	POSTGRES_URL=$(POSTGRES_URL) \
 	air

run:
	go install
	HOSTNAME=$(HOSTNAME) \
	URL_EXPIRE_SECONDS=$(URL_EXPIRE_SECONDS) \
 	POSTGRES_URL=$(POSTGRES_URL) \
 	miniurl

# example:
# $ make mg-create name=create_users_table
mg-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

mg-up:
	migrate -database $(POSTGRES_URL) -path db/migrations up

mg-down:
	migrate -database $(POSTGRES_URL) -path db/migrations down
