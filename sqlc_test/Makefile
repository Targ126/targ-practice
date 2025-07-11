sqlc repo:
	sqlc generate



local-run:
	export POSTGRES_APPNAME=targ_test && \
	export POSTGRES_DBNAME=targ_test && \
	export POSTGRES_USERNAME=targ_test && \
	export POSTGRES_HOST=localhost && \
	export POSTGRES_PORT=5432 && \
	export POSTGRES_PASSWORD=targ_test && \
	go run ./cmd/main.go