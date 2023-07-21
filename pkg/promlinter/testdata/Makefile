GO := go

# This is needed because analysistest framework doesn't support Go modules,
# only GOPATH-like trees: https://github.com/golang/go/issues/37054
# So to parse the test package, we vendor its dependencies under src/
# directory. For the same reason, we use relative import paths to import local
# packages, so we have to ignore `go mod` errors complaining about them.
vendor-src:
	$(GO) mod tidy -e && $(GO) mod vendor -e -o src/ && $(GO) mod verify
