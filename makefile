up:
	docker-compose up
up-clean:
	 docker-compose up --force-recreate
down:
	docker-compose down
test-go:
	vgo test ./api/...
test-js:
	cd frontend && npm test
gen-schema: #required gqlgen installed locally
	gqlgen -out api/gql/gql.go \
		-package gql \
		-typemap api/gql/types.json \
		-schema api/gql/schema.graphql
rebuild:
	docker-compose build api
	docker-compose build frontend
build-prod:
	docker build ./api --build-arg app_env=production
	docker build ./frontend --build-arg app_env=production
