GOOS ?= darwin
GOARCH ?= arm64
TAG := $(shell git describe --tags --abbrev=0)
GO_TAG := $(shell echo "$(TAG)" | sed 's/^go\///')
DIRTY := $(shell [ -n "$(git status -s)" ] && echo '-snapshot')
VERSIONED_PREFIX := $(if $(DIRTY),$(GO_TAG)_$(GOOS)_$(GOARCH),$(GO_TAG)_SNAPSHOT_$(GOOS)_$(GOARCH))

version:
	echo $(VERSION)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/sample-app_$(VERSIONED_PREFIX) ./cmd

artifact: build
	mkdir -p ./dist/artifacts
	tar zcvf ./dist/artifacts/sample-app_$(VERSIONED_PREFIX).tar.gz -C ./bin sample-app_$(VERSIONED_PREFIX)
	shasum -a 256 ./dist/artifacts/sample-app_$(VERSIONED_PREFIX).tar.gz > ./dist/artifacts/sample-app_$(VERSIONED_PREFIX).sha256

# echo $PAT | gh auth login --with-token -h github.com
release: artifact
	@(gh release --repo yana1205-test/sample-app view $(TAG) ;\
	if [[ "$$?" != "0" ]];then \
		echo create release $(TAG) ;\
		gh release --repo yana1205-test/sample-app create $(TAG) --generate-notes ;\
	fi)
	gh release --repo yana1205-test/sample-app upload $(TAG) ./dist/artifacts/sample-app_$(VERSIONED_PREFIX).*