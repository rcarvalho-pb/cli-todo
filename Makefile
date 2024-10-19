MIGRATION_PATH=./db/migrations
DB_PATH=./db/db.db
migration:
	@migrate create -ext sql -dir ${MIGRATION_PATH} -seq ${name}
migrate:
	@migrate -path ${MIGRATION_PATH} -database sqlite3://${DB_PATH} ${type}
run:
	@env DB_PATH=${DB_PATH} go run *.go
