DB_PATH=./db/db.db
run:
	@env DB_PATH=${DB_PATH} go run *.go
