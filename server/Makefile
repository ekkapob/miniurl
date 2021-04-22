REDIS_URL = localhost:6379
REDIS_PASSWORD =
REDIS_URL_EXPIRE_SECONDS = 3600
POSTGRES_URL = postgres://postgres:postgres@localhost:5432/app?sslmode=disable
POSTGRES_URL_EXPIRE_SECONDS = 604800
POPULAR_URL_TIMELAPSE_MINS = 15
HOSTNAME = http://localhost:4000
PORT = :4000

dev:
	REDIS_URL=$(REDIS_URL) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	REDIS_URL_EXPIRE_SECONDS=$(REDIS_URL_EXPIRE_SECONDS) \
 	POSTGRES_URL=$(POSTGRES_URL) \
	POSTGRES_URL_EXPIRE_SECONDS=$(POSTGRES_URL_EXPIRE_SECONDS) \
	POPULAR_URL_TIMELAPSE_MINS=$(POPULAR_URL_TIMELAPSE_MINS) \
	HOSTNAME=$(HOSTNAME) \
	PORT=$(PORT) \
 	air -c .air.toml

run:
	REDIS_URL=$(REDIS_URL) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	REDIS_URL_EXPIRE_SECONDS=$(REDIS_URL_EXPIRE_SECONDS) \
 	POSTGRES_URL=$(POSTGRES_URL) \
	POSTGRES_URL_EXPIRE_SECONDS=$(POSTGRES_URL_EXPIRE_SECONDS) \
	POPULAR_URL_TIMELAPSE_MINS=$(POPULAR_URL_TIMELAPSE_MINS) \
	HOSTNAME=$(HOSTNAME) \
	PORT=$(PORT) \
 	miniurl

test:
	go test ./... -v

# example:
# $ make mg-create name=create_users_table
mg-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

mg-up:
	migrate -database $(POSTGRES_URL) -path db/migrations up

mg-down:
	migrate -database $(POSTGRES_URL) -path db/migrations down

docker-up:
	docker-compose up postgres redis app

docker-migrate:
	docker-compose up migrate

