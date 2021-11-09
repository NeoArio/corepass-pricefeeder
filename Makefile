PROJECT_NAME :=  $(shell basename pwd)
BINANARY_NAME := $($(PROJECT_NAME)d)

.PHONY: build clean start test test-local compile-test

build: 
	## Build 
	@echo "  >  Building binary..."
	@echo $(BINANARY_NAME)
	- go build -o pricefeederd main.go
	@echo "  >  Build done ..."

	# Do all build jobs


clean: 
	## to remove generated files
	- rm -f pricefeederd
	# note: call scripts from /scripts
	#
start: 
	- ./pricefeederd serve


compile-test:
	- go test -c  ./tests/... -o test-pricefeederd
	- @echo "Tests Compiled"

test: compile-test
	- ./test-pricefeederd
	- @echo "Tests Run Finished"

test-local: compile-test
	- ./test-pricefeederd -config config.yaml
	- @echo "Tests Run Finished"
