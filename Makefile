.PHONY: db
db:
	docker compose -f compose.yaml exec postgres psql -U admin -d todo
