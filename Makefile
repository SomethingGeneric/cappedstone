# Makefile for building capyagent and capydaemon and creating self-signed certs for dev testing
# Usage:
#   make build         # builds local binaries into bin/
#   make release       # builds release tarballs for linux/darwin (amd64/arm64) into release/
#   make test_agent    # create certs/agent.crt and certs/agent.key for local testing
#   make test_daemon   # create certs/daemon.crt and certs/daemon.key for local testing
#   make test_all      # create both agent and daemon certs
#   make clean         # remove bin/, release/, certs/

BINS := capyagent capydaemon
BIN_DIR.capyagent := capyagent
BIN_DIR.capydaemon := capydaemon
BIN_PACKAGE.capyagent := ./cmd/capyagent
BIN_PACKAGE.capydaemon := ./cmd/capydaemon
OUTDIR := bin
RELEASEDIR := release
CERTDIR := certs
AGENT_CERT := $(CERTDIR)/agent.crt
AGENT_KEY := $(CERTDIR)/agent.key
DAEMON_CERT := $(CERTDIR)/daemon.crt
DAEMON_KEY := $(CERTDIR)/daemon.key
RELEASE_PLATFORMS := linux_amd64 linux_arm64 darwin_amd64 darwin_arm64
RELEASE_ARTIFACTS := $(foreach bin,$(BINS),$(foreach platform,$(RELEASE_PLATFORMS),$(RELEASEDIR)/$(bin)_$(platform).tar.gz))
GOCACHE ?= $(CURDIR)/.cache/go-build
GOMODCACHE ?= $(CURDIR)/.cache/go-mod
GOPATH ?= $(CURDIR)/.cache/gopath
CACHE_ENV := GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) GOPATH=$(GOPATH)

GO ?= go
GOFLAGS ?=
LDFLAGS ?= -ldflags="-s -w"

.PHONY: all build release clean test_agent test_daemon test_all

all: build

build: $(OUTDIR) $(BINS:%=$(OUTDIR)/%)

$(OUTDIR):
	mkdir -p $(OUTDIR)

# build each binary from its module-local cmd directory
$(OUTDIR)/%: | $(OUTDIR)
	mkdir -p $(GOCACHE) $(GOMODCACHE) $(GOPATH)
	$(CACHE_ENV) $(GO) build -C $(BIN_DIR.$*) $(GOFLAGS) $(LDFLAGS) -o $(abspath $@) $(BIN_PACKAGE.$*)

# build release tarballs for common platforms
release: $(RELEASE_ARTIFACTS)

$(RELEASEDIR)/:
	mkdir -p $(RELEASEDIR)

$(CERTDIR):
	mkdir -p $(CERTDIR)

# Self-signed certs for local/dev testing.
# Uses openssl with subjectAltName for localhost and 127.0.0.1 (requires openssl supporting -addext).
test_agent: $(CERTDIR)
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	  -keyout $(AGENT_KEY) -out $(AGENT_CERT) \
	  -subj "/CN=capyagent" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"
	mkdir -p capyagent/$(CERTDIR)
	install -m 0644 $(AGENT_CERT) capyagent/$(CERTDIR)/
	install -m 0600 $(AGENT_KEY) capyagent/$(CERTDIR)/

test_daemon: $(CERTDIR)
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	  -keyout $(DAEMON_KEY) -out $(DAEMON_CERT) \
	  -subj "/CN=capydaemon" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"
	mkdir -p capydaemon/$(CERTDIR)
	install -m 0644 $(DAEMON_CERT) capydaemon/$(CERTDIR)/
	install -m 0600 $(DAEMON_KEY) capydaemon/$(CERTDIR)/

test_all: test_agent test_daemon

clean:
	rm -rf $(OUTDIR) $(RELEASEDIR) $(CERTDIR) $(GOCACHE) $(GOMODCACHE) $(GOPATH)

define release_template
$(RELEASEDIR)/$(1)_$(2).tar.gz: | $(RELEASEDIR)/
	mkdir -p $(RELEASEDIR)/$(1)_$(2) $(GOCACHE) $(GOMODCACHE) $(GOPATH)
	GOOS=$(word 1,$(subst _, ,$(2))) GOARCH=$(word 2,$(subst _, ,$(2))) $(CACHE_ENV) $(GO) build -C $(BIN_DIR.$(1)) $(GOFLAGS) $(LDFLAGS) -o $(abspath $(RELEASEDIR)/$(1)_$(2)/$(1)) $(BIN_PACKAGE.$(1))
	tar -C $(RELEASEDIR)/$(1)_$(2) -czf $$@ $(1)
	rm -rf $(RELEASEDIR)/$(1)_$(2)
endef

$(foreach bin,$(BINS),$(foreach platform,$(RELEASE_PLATFORMS),$(eval $(call release_template,$(bin),$(platform)))))
