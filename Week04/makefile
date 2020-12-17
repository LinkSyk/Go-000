proto:
	cd api/account/ && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    v1/account.proto 

run:
	cd cmd/ && go run account/*