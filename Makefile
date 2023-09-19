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
	docker compose exec app bash
db:
	docker compose exec db bash
logs:
	docker compose logs -f --tail=100
app-log:
	docker compose logs -f --tail=100 app
db-log:
	docker compose logs -f --tail=100 db
migrate-up:
	docker compose exec app migrate -database "mysql://user:password@tcp(db:3306)/go-test" -path /app/src/Infrastructures/Migrations/ up
migrate-down:
	docker compose exec app migrate -database "mysql://user:password@tcp(db:3306)/go-test" -path /app/src/Infrastructures/Migrations/ down
test:
	docker compose exec app go test ./src/Tests/...