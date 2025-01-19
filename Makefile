LOCAL_BIN:=$(CURDIR)/bin

bin-deps: 
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@latest

go-deps:
	go get github.com/idsulik/go-collections/queue
	go get -u github.com/golang/mock/gomock

add-deps: bin-deps go-deps

generate-mockgen:
	PATH="$(LOCAL_BIN):$$PATH" go generate -x -run=mockgen ./...

generate-mock:
	$(info Generating mocks...)
	find . -name '*_mock.go' -delete
	bin/mockgen -source=internal/input/input.go -destination=internal/input/mocks/input_mock.go -package=mock_input
	bin/mockgen -source=internal/engine/engine.go -destination=internal/engine/mocks/engine_mock.go -package=mock_engine 
test:
	$(info running tests...)
	go test ./tests/
