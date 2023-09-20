run:
	go run cmd/main.go

swag:
	swag init -g api/main.go -o api/docs
	