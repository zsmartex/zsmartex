.PHONY: buf generate wire backend user

# Generate protobuf -> go file
buf:
	buf generate

generate: buf

wire:
	@cd cmd/user/client && wire
	@cd cmd/user/server && wire

backend:
	docker-compose up -d postgres mongodb mongo-express nats

user: backend
	docker-compose up --build user-server user-client
