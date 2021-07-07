help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

start: ## Up the docker-compose without cache or orphans
	docker-compose up \
	--detach \
	--build \
	--remove-orphans

stop: ## Down the docker-compose 
	docker-compose stop

logs: ## Display logs of your containers 
	docker-compose logs --follow

lint:
	gofmt -e -l -s -w .

.PHONY: help

seed: ## Populates Elasticsearch with sample data
	curl -H 'Content-Type: application/x-ndjson' -XPOST 'localhost:9200/shakespeare/doc/_bulk?pretty' --data-binary @dataset.json