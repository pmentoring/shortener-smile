up:
	docker compose up -d
down:
	docker compose down --remove-orphans
exec:
	docker compose exec -it go-app bash