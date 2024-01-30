start:
	swag init -g ./cmd/main.go -o cmd/docs/
	go run cmd/main.go

gen-swag:
	swag init -g ./cmd/main.go -o cmd/docs