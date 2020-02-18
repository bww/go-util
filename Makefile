
TEST_PKGS = ./...

.PHONY: all test

all: test

test:
	go test -v $(TEST_PKGS)
