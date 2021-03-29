REDIS_URL = localhost:6379
REDIS_PASSWORD =
REDIS_URL_EXPIRE_SECONDS = 3600
POSTGRES_URL = postgres://postgres:postgres@localhost:5432/app?sslmode=disable
POSTGRES_URL_EXPIRE_SECONDS = 604800
POPULAR_URL_TIMELAPSE_MINS = 15
HOSTNAME = http://127.0.0.1:8000

dev:
	REDIS_URL=$(REDIS_URL) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	REDIS_URL_EXPIRE_SECONDS=$(REDIS_URL_EXPIRE_SECONDS) \
 	POSTGRES_URL=$(POSTGRES_URL) \
	POSTGRES_URL_EXPIRE_SECONDS=$(POSTGRES_URL_EXPIRE_SECONDS) \
	POPULAR_URL_TIMELAPSE_MINS=$(POPULAR_URL_TIMELAPSE_MINS) \
	HOSTNAME=$(HOSTNAME) \
 	air -c .air.toml

run:
	go install
	REDIS_URL=$(REDIS_URL) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	REDIS_URL_EXPIRE_SECONDS=$(REDIS_URL_EXPIRE_SECONDS) \
 	POSTGRES_URL=$(POSTGRES_URL) \
	POSTGRES_URL_EXPIRE_SECONDS=$(POSTGRES_URL_EXPIRE_SECONDS) \
	POPULAR_URL_TIMELAPSE_MINS=$(POPULAR_URL_TIMELAPSE_MINS) \
	HOSTNAME=$(HOSTNAME) \
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
