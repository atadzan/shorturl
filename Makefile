run-api:
	go run api/shorturl.go
run-rpc:
	go run rpc/transform/transform.go
create-migrate:
	migrate create -ext sql -dir rpc/transform/models/migrations -tz Local $(name)
migrate-up:
	migrate -database 'mysql://user:password@tcp(localhost:6606)/db' -path rpc/transform/model/migrations up
migrate-down:
	migrate -database 'mysql://user:password@tcp(localhost:6606)/db' -path rpc/transform/model/migrations down