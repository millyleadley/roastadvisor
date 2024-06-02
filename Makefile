run:
	go build . && ./roastadvisor

swagger:
	swag init --requiredByDefault

gen:
	openapi-generator generate -i docs/swagger.json -g typescript-fetch -o client/gen
