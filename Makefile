
TEST_PKGS = github.com/bww/go-util/env \
						github.com/bww/go-util/text \
						github.com/bww/go-util/rand \
						github.com/bww/go-util/qname \
						github.com/bww/go-util/uuid \
						github.com/bww/go-util/scan \
						github.com/bww/go-util/slug

.PHONY: all test

all: test

test:
	go test -test.v $(TEST_PKGS)
