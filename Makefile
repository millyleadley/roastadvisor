build:
	go build . && ./roastadvisor

swagger:
	swag init --overridesFile=.swaggo

gen:
	openapi-generator generate -i docs/swagger.json -g typescript-fetch -o client/gen
