migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users





migrate -path=./cmd/migrate/migrations -database="postgres://postgres:1234@localhost:5441/social?sslmode=disable" up 





direnv allow .