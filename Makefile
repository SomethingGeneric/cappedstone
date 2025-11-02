# Makefile for building capyagent and capydaemon and creating self-signed certs for dev testing
# Usage:
#   make build         # builds local binaries into bin/
#   make release       # builds release tarballs for linux/darwin (amd64/arm64) into release/
#   make test_agent    # create certs/agent.crt and certs/agent.key for local testing
#   make test_daemon   # create certs/daemon.crt and certs/daemon.key for local testing
#   make test_all      # create both agent and daemon certs
#   make clean         # remove bin/, release/, certs/

BINS := capyagent capydaemon
OUTDIR := bin
RELEASEDIR := release
CERTDIR := certs

GO ?= go
GOFLAGS ?=
LDFLAGS ?= -ldflags="-s -w"

.PHONY: all build release clean test_agent test_daemon test_all

all: build

build: $(OUTDIR) $(BINS:%=$(OUTDIR)/%)

$(OUTDIR):
	mkdir -p $(OUTDIR)

# expects each binary to have a corresponding ./cmd/<name> package
$(OUTDIR)/%:
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $@ ./cmd/$*

# build release tarballs for common platforms
release: $(RELEASEDIR)
	@for GOOS in linux darwin; do \
	  for GOARCH in amd64 arm64; do \
		for bin in $(BINS); do \
		  out=$(RELEASEDIR)/$${bin}_$${GOOS}_$${GOARCH}; \
		  mkdir -p $$out; \
		  env GOOS=$$GOOS GOARCH=$$GOARCH $(GO) build $(GOFLAGS) $(LDFLAGS) -o $$out/$$bin ./cmd/$$bin || exit 1; \
		  tar -C $$out -czf $(RELEASEDIR)/$${bin}_$${GOOS}_$${GOARCH}.tar.gz -C $$out $$bin; \
		  rm -rf $$out; \
		done; \
	  done; \
	done

$(RELEASEDIR):
	mkdir -p $(RELEASEDIR)

$(CERTDIR):
	mkdir -p $(CERTDIR)

# Self-signed certs for local/dev testing.
# Uses openssl with subjectAltName for localhost and 127.0.0.1 (requires openssl supporting -addext).
test_agent: $(CERTDIR)
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	  -keyout $(CERTDIR)/agent.key -out $(CERTDIR)/agent.crt \
	  -subj "/CN=capyagent" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

test_daemon: $(CERTDIR)
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	  -keyout $(CERTDIR)/daemon.key -out $(CERTDIR)/daemon.crt \
	  -subj "/CN=capydaemon" -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"

test_all: test_agent test_daemon

clean:
	rm -rf $(OUTDIR) $(RELEASEDIR) $(CERTDIR)