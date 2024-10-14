run:
	go build . && ./roastadvisor

swagger:
	swag init --requiredByDefault

api:
	goa gen github.com/millyleadley/roastadvisor/api/design -o api

gen:
	openapi-generator generate -i docs/swagger.json -g typescript-fetch -o client/gen
