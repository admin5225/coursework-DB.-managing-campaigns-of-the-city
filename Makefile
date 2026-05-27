include .env
export


env-up:
	docker compose up -d postgres

env-down:
	docker compose down postgres

env-cleanup:
	docker compose down postgres
	rmdir /s /q out\pgdata

migrate-create:
	@if "$(seq)"=="" ( \
		echo Seq no enter! & exit /b 1 \
	)

	docker compose run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq $(seq)

migrate-up:
	$(MAKE) migrate-action action=up

migrate-down:
	$(MAKE) migrate-action action=down

migrate-action:
	@if "$(action)"=="" ( echo Action no enter! & exit /b 1 )

	docker compose run --rm postgres-migrate -path /migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable "${action}"



servise-run:
	go run ./cmd/app/main.go