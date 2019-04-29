GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build
GOTEST   := CGO_ENABLED=1 $(GO) test
PACKAGES  := $$(go list ./... | grep -vE 'tests|cmd|vendor')
FAILPOINT_DIR := $$(for p in $(PACKAGES); do echo $${p\#"github.com/amyangfei/fpcov/"}; done)
FAILPOINT := bin/failpoint-ctl
FAILPOINT_ENABLE  := $$(echo $(FAILPOINT_DIR) | xargs $(FAILPOINT) enable >/dev/null)
FAILPOINT_DISABLE := $$(find $(FAILPOINT_DIR) | xargs $(FAILPOINT) disable >/dev/null)
TEST_DIR := /tmp/failpoint_test

.PHONY: unit_test coverage build

build:
	$(GOBUILD) -o bin/main ./cmd

unit_test:
	which $(FAILPOINT) >/dev/null 2>&1 || $(GOBUILD) -o $(FAILPOINT) github.com/pingcap/failpoint/failpoint-ctl
	mkdir -p $(TEST_DIR)
	$(FAILPOINT_ENABLE)
	$(GOTEST) -covermode=atomic -coverprofile="$(TEST_DIR)/cov.unit_test.out" $(PACKAGES) \
	|| { $(FAILPOINT_DISABLE); exit 1; }
	$(FAILPOINT_DISABLE)

coverage:
	GO111MODULE=off go get github.com/zhouqiang-cl/gocovmerge
	gocovmerge "$(TEST_DIR)"/cov.* | grep -vE ".*.pb.go|.*.__failpoint_binding__.go" > "$(TEST_DIR)/all_cov.out"
ifeq ("$(GL_TRAVIS_CI)", "on")
    GO111MODULE=off go get github.com/mattn/goveralls
    goveralls -coverprofile=$(TEST_DIR)/all_cov.out -service=travis-ci
else
	go tool cover -html "$(TEST_DIR)/all_cov.out" -o "$(TEST_DIR)/all_cov.html"
endif

failpoint-enable:
	$(FAILPOINT_ENABLE)

failpoint-disable:
	$(FAILPOINT_DISABLE)
