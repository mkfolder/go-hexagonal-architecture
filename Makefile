gen:
	@wire gen ./internal/bootstrap

compose:
	@mkdir -p volumes/postgres_data
	@mkdir -p volumes/pgadmin_data
	@docker compose up -d

down:
	@docker compose down
	@docker image prune -af
	@docker volume prune -af

recompose:
	@make down
	@make compose

.PHONY: gen compose down recompose
