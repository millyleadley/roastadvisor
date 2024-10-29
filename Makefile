# Override these for different configurations
export PGUSER ?= postgres
export PGPASSWORD ?= postgres
export PGDATABASE ?= development
export PGHOST ?= localhost
export PGSSLMODE ?= disable

run:
	go build . && ./roastadvisor

# gen generates the OpenAPI spec and then the TypeScript client code from that spec
gen:
	goa gen github.com/millyleadley/roastadvisor/api/design -o api	
	openapi-generator generate -i api/gen/http/openapi3.json -g typescript-fetch -o client/gen

brew-start:
	brew services start postgresql@14

brew-stop:
	brew services stop postgresql@14

################################################################################
# Database
################################################################################

psql:
	psql -U ${PGUSER} -h ${PGHOST} -d ${PGDATABASE}

GOOSE=goose -dir lib/database/migrations

# goose picks these up when running migrations
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=${GOOSE_DRIVER}://${PGUSER}@${PGHOST}:5432/${PGDATABASE}?sslmode=${PGSSLMODE}

create-migration:
	@if [ -z "${migration_name}" ] ; then echo "No migration name supplied. Invocation: make migration_name=migration_name_here create-migration" ; false ; fi
	$(GOOSE) create $(migration_name) sql

migrate:
	$(GOOSE) up
	$(MAKE) dump-structure

rollback:
	$(GOOSE) down
	$(MAKE) dump-structure

# Dump the current version of the database, so we can easily diff it
dump-structure:
	pg_dump --schema-only --no-owner -U ${PGUSER} -h ${PGHOST} -d ${PGDATABASE} | \
		grep -Ev '^-- Dumped from|^-- Dumped by' > lib/database/structure.sql
