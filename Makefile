dev/run/import:
	. ./.env
	docker compose up -d
	# 上のコマンドの後に25秒待つ
	sleep 25
	docker compose exec server go run src/cmd/insertcsv/main.go

dev/run/server:
	. ./.env
	docker compose up -d
