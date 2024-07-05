dev/run/import:
	. ./.env
	docker compose up -d
	docker compose exec server go run src/cmd/insertcsv/main.go

dev/run/server:
	. ./.env
	docker compose up -d

restart: dev/run/down dev/run/server

dev/run/down:
	docker compose down
