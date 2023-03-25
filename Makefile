# Generate protobuf -> go file
buf:
	buf generate

run: run-user run-proxy

run-user:
	cd cmd/user && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags migrate github.com/zsmartex/zsmartex/cmd/user

run-proxy:
	cd cmd/proxy && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run github.com/zsmartex/zsmartex/cmd/proxy

.PHONY: wire
wire:
	cd internal/user/app && wire && cd -
