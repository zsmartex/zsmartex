.PHONY: buf generate backend user

# Generate protobuf -> go file
buf:
	buf generate

generate: ent buf

backend:
	docker-compose up -d postgres mongodb mongo-express nats

user: backend
	docker-compose up --build user-server user-client
