protoc:
	cd proto && protoc --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ *.proto