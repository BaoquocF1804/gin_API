server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=openapiv2\
        proto/*.proto
proto2:
	rm -f pb_account/*.go
	protoc --proto_path=proto2 --go_out=pb_account --go_opt=paths=source_relative \
        --go-grpc_out=pb_account --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb_account --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=openapiv2\
        proto2/*.proto
proto3:
	protoc --proto_path=proto3 --go_out=pb_connect --go_opt=paths=source_relative \
        --go-grpc_out=pb_connect --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb_connect --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=openapiv2\
        proto3/*.proto
evans:
	evans --host localhost --port 9090 -r repl
.PHONY: server proto evans proto2