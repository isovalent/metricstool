GO := go

# This is needed because analysistest framework doesn't support Go modules,
# only GOPATH-like trees: https://github.com/golang/go/issues/37054
# So to parse the test package, we vendor its dependencies under src/
# directory.
vendor-src:
	$(GO) mod tidy && $(GO) mod vendor -o src/ && $(GO) mod verify
