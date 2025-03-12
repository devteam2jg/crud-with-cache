init-dotenv:
	mkdir -p .env.local
	find .env.example -type f ! -name '*.md' -exec cp {} .env.local/ \;

build:
	docker compose -f compose/local.yml build

up-infra:
	docker compose -f compose/infra.yml up -d

up:
	docker compose -f compose/local.yml up

up-app:
	docker compose -f compose/local.yml up app