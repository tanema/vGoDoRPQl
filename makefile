up:
	docker-compose up

up-clean:
	 docker-compose up --force-recreate

down:
	docker-compose down

gen-schema: #required gqlgen installed locally
	gqlgen -out api/graph_api/generated.go \
		-package graph_api \
		-typemap api/graph_api/types.json \
		-schema api/graph_api/schema.graphql

rebuild:
	docker-compose build api
	docker-compose build frontend
	docker-compose build db

build-prod:
	docker build ./api --build-arg app_env=production
	docker build ./frontend --build-arg app_env=production
	docker build ./db
