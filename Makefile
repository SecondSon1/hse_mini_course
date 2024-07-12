http_server:
	go build -C ./cmd/http/server -o ../../../bin/http_server

http_client:
	go build -C ./cmd/http/client -o ../../../bin/http_client

grpc_server:
	go build -C ./cmd/grpc/server -o ../../../bin/grpc_server

grpc_client:
	go build -C ./cmd/grpc/client -o ../../../bin/grpc_client

http: http_server http_client

grpc: grpc_server grpc_client
