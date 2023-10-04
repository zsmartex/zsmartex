# Generate ent orm file
ent:
	go generate ./...

# Generate protobuf -> go file
buf:
	buf generate

generate: ent buf
