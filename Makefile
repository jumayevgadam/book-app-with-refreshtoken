run:
	swag init --output ./docs --parseDependency --generalInfo ./cmd/app/main.go
	go run cmd/app/main.go

migrate_create:
	migrate create -ext sql -dir ./internal/infrastructure/migrations/ -seq tables

migration_up:
	migrate -path ./internal/infrastructure/migrations/ -database "postgresql://postgres:12345@localhost:5432/jwt_app?sslmode=disable" -verbose up
