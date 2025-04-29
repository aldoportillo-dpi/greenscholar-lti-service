new-migration:
	migrate create -ext sql -dir db/migrations -seq

sqlc-generate:
	sqlc generate --file ./db/sqlc.yaml

openapi-generate:
	npx @redocly/cli build-docs openapi/openapi.yaml
	mv ./redoc-static.html docs/index.html

	go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=openapi/config.yaml openapi/openapi.yaml
	npx openapi-typescript ./openapi/openapi.yaml -o openapi/schema.ts