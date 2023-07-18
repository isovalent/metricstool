GO := go
TARGET := metricstool
TEST_FLAGS ?= -v -cover -timeout=5m -race

.PHONY: all
all: $(TARGET)

.PHONY: $(TARGET)
$(TARGET):
	$(GO) build -o $(TARGET) ./cmd

.PHONY: test
test:
	$(GO) test $(TEST_FLAGS) ./...

.PHONY: vendor-test
vendor-test:
	$(MAKE) -C pkg/promlinter/testdata vendor-src

.PHONY: vendor
vendor:
	$(GO) mod tidy && $(GO) mod vendor && $(GO) mod verify
