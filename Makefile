# Generate protobuf -> go file
buf:
	buf generate

run: run-users run-proxy

run-users:
	cd cmd/users && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags migrate github.com/zsmartex/zsmartex/cmd/users

run-proxy:
	cd cmd/proxy && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run github.com/zsmartex/zsmartex/cmd/proxy

.PHONY: wire
wire:
	cd internal/users/app && wire && cd -
