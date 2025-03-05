build:
		docker-compose up --build
run:
		go run cmd/main.go
dev-up:
		docker-compose up postgres
dev-down:
		docker-compose down postgres
dev-down-v:
		docker-compose down -v
