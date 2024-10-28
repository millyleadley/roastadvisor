run:
	go build . && ./roastadvisor

gen-api:
	goa gen github.com/millyleadley/roastadvisor/api/design -o api	

gen-ts:
	openapi-generator generate -i api/gen/http/openapi3.json -g typescript-fetch -o client/gen

psql:
	psql -U postgres -h localhost -d postgres

brew-start:
	brew services start postgresql@14

brew-stop:
	brew services stop postgresql@14