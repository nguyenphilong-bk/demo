include .env
export $(shell sed 's/=.*//' .env)

down:
	docker compose down
up:
	docker compose up -d --remove-orphans --force-recreate && echo "Your server is running at port $(PORT)"

info:
	echo "DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) DB_NAME=$(DB_NAME) PORT=$(PORT)"
