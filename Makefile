up:
	docker compose up -d
build:
	docker compose build
stop:
	docker compose stop
down:
	docker compose down --remove-orphans
down-v:
	docker compose down --remove-orphans --volumes
restart:
	@make down
	@make up
destroy:
	docker compose down --rmi all --volumes --remove-orphans
ps:
	docker compose ps
app:
	docker compose exec app /bin/sh
db:
	docker compose exec db /bin/sh