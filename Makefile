DB_URL := "postgres://postgres:mubina2007@localhost:5432/migration?sslmode=disable"

gorn:
	go run api/api.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up 

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down 

migrate-file:
	migrate create -ext sql -dir migrations/ -seq table