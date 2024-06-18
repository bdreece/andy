# ============================================================================ #
# DEFINITIONS
# ============================================================================ #
project := andy

GO        := go
GOFLAGS   := -v
NPM 	  := npm
NPMFLAGS  := -w @andy/app
AIR       := $(GO) run github.com/air-verse/air@latest
AIRFLAGS  := -c configs/air.toml
SQLC      := $(GO) run github.com/sqlc-dev/sqlc/cmd/sqlc@latest
SQLCFLAGS := -f configs/sqlc.yaml

rootdir := .

bindir  = $(rootdir)/bin
webdir  = $(rootdir)/web
distdir = $(appdir)/dist

appdir = $(webdir)/app
appout = $(distdir)/index.umd.cjs $(distdir)/index.css 
appsrc = $(shell find $(appdir) \
		 	-type d \( \
		 		-path $(appdir)/node_modules -o \
				-path $(distdir) \) \
				-prune -false -o \
			-type f \( \
				-name "*.js" -o \
				-name "*.json" -o \
				-name "*.ts" -o \
				-name "*.css" \))

maindir = $(rootdir)/cmd/$(project)
mainout = $(bindir)/$(project)
mainsrc = $(shell find . -type f -name "*.go")

sqldir = $(rootdir)/pkg/database
sqlout = $(sqldir)/db.sql.go $(sqldir)/models.sql.go $(sqldir)/querier.sql.go
sqlsrc = $(shell find $(sqldir) -type f -name "*.sql")

tmpldir = $(webdir)/templates
tmplsrc = $(shell find $(tmpldir) -type f -name "*.gotmpl")

# ============================================================================ #
# HELPERS
# ============================================================================ #

default: build

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:' 
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## clean: clean all targets
## clean/main: clean the main executable
## clean/app: clean the frontend app
## clean/sql: clean the generated SQL driver code
.PHONY: clean clean/main clean/app
clean: clean/app clean/main clean/sql
clean/main: $(bindir)
	rm -f $(bindir)/*
clean/app: $(distdir)
	rm -f $(distdir)/*
clean/sql: $(sqldir)
	rm -f $(sqlout)

## build: build all targets
## build/main: build the main executable
## build/app: build the frontend app
.PHONY: build build/main build/app
build build/main: generate $(mainout)
build/app: $(appout)

## watch: run and watch the main executable
.PHONY: watch
watch:
	$(AIR) $(AIRFLAGS)

## generate: generate sources
## generate/sql: generate SQL driver sources
.PHONY: generate generate/sql
generate/sql: $(sqlout)
generate:
	@$(GO) generate ./...

# ============================================================================ #
# RECIPES
# ============================================================================ #
$(bindir):
	mkdir -p $@

$(distdir):
	mkdir -p $@

$(mainout): $(mainsrc) $(appout) | $(bindir)
	$(GO) build $(GOFLAGS) -o $(bindir) $(maindir)

$(appout) &: $(appsrc) $(tmplsrc)
	$(NPM) run $(NPMFLAGS) build

$(sqlout) &: $(sqlsrc)
	$(SQLC) generate $(SQLCFLAGS)


