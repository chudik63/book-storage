build:
	docker-compose --env-file ./configs/.env up --build

run:
	go run cmd/main.go

dev up:
	docker-compose --env-file ./configs/.env up postgres

dev down:

	docker-compose --env-file ./configs/.env down postgres
	
dev down v:
	docker-compose --env-file ./configs/.env down -v
