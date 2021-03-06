include ../../Makefile.common_go

PACKAGE_PATH = github.com/chef/automate/components/teams-service

BINS = ${PACKAGE_PATH}/cmd/teams-service
MIGRATION_READMES = storage/postgres/migration/sql/README.md storage/postgres/datamigration/sql/README.md
GIT_SHA = $(shell git rev-parse HEAD)
BUILD_TIME ?= $(shell date -u '+%Y%m%d%H%M%S')
GO_LDFLAGS = --ldflags "-X ${LIBRARY_PATH}/version.Version=${BUILD_TIME} -X ${LIBRARY_PATH}/version.GitSHA=${GIT_SHA} -X ${LIBRARY_PATH}/version.BuildTime=${BUILD_TIME}"

packages:=${PACKAGE_PATH}/...
ifdef CI
  verbose?="-v"
endif

all: lint build test

static: lint ${MIGRATION_READMES}
unit: build test

.PHONY: ${MIGRATION_READMES}
${MIGRATION_READMES}:
	../../scripts/generate_and_check_migration_files.sh $@

${BINS}: bin
	@echo "GO $@"
	@cd bin; go build ${GO_LDFLAGS} $@

bin:
	mkdir -p bin

build: ${BINS}

# Regenerate all *.pb.go files
proto:
	cd ../../ && hab studio run 'source .studiorc; compile_go_protobuf_component teams-service'

# This will use memstore for server tests and skip other PG tests
# by default unless in ciMode or PG_URL is set.
test:
	@go test $(verbose) $(packages) -p 1 --parallel 1 -cover

test_with_db:
	@docker ps | grep teams-postgres || (echo "Docker postgres not up. Run make setup_docker_pg and try this command again."; exit 1)
	@PG_URL="postgresql://postgres@127.0.0.1:5432/teams_test?sslmode=disable&timezone=UTC" go test $(packages) -p 1 --parallel 1 -cover
	@echo "Docker containers still up, run 'make kill_docker_pg' to bring them down or test again with make test_with_db."

setup_docker_pg:
	docker run --name teams-postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=teams_test -p 5432:5432 -d postgres:9
	sleep 5 # let docker come up
	# This creates the extension we need to use UUIDs in the migrations.
	# Done in habitat in prod. Not done in code because you must be a superuser.
	docker exec teams-postgres psql -d "postgresql://postgres@127.0.0.1:5432/teams_test?sslmode=disable" -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""

kill_docker_pg:
	docker rm -f teams-postgres || true

# this command lists all the changes since master, and looks for modifications
# to the migration files -- if there's any of (M)odify, (R)ename, or (D)elete,
# it will trigger an error
migrations-append-only-check:
	@git diff --name-status master... | awk '/[RMD][0-9]*\tcomponents\/teams-service\/postgres\/migration\/sql\//{ print "modified/deleted/renamed: " $$2; ec=1 } END{ exit ec }'

.PHONY: all static unit build compile proto test dep-ensure have-dep lint
