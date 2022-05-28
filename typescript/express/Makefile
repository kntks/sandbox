install:
	docker compose run --rm app npm ci

start:
	docker compose up

fmt:
	docker compose run --rm app npm run lint-fix

pull:
	docker compose exec app npx prisma db pull

generate:
	 docker compose exec app npx prisma generate

restart:
	docker compose restart app

down:
	docker compose down