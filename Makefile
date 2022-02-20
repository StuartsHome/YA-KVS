GOGEN=go generate
GOGEN_MOCKERY=$(GOGEN) -run="mockery"
include mock.mk

test:
	go test ./...
