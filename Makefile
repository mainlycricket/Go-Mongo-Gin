include secrets.sh

.PHONY:build-server
build-server:
	go build -o bin/ ./cmd/server/

.PHONY:run-server
run-server: build-server
	./bin/server

.PHONY:watch-server
watch-server:
	reflex -r '\.go$$' -s -- sh -c 'make run-server'

DB_NAME ?= $(DB_NAME)

.PHONY:reset-db
reset-db:
	mongosh --eval "use $(DB_NAME)" --eval "db.dropDatabase()"
	mongosh --eval "use $(DB_NAME)" --eval "load('./internal/data/models/mongo-indexes.js')"

.PHONY:test-code
test-code:
	mongosh --eval "use $(DB_NAME)_test" --eval "db.dropDatabase()"
	mongosh --eval "use $(DB_NAME)_test" --eval "load('./internal/data/models/mongo-indexes.js')"
	go test -count=1 ./...
	mongosh --eval "use $(DB_NAME)_test" --eval "db.dropDatabase()"
