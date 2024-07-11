http_server:
	go build -C ./cmd/http/server -o ../../../http_server.out

http_client:
	go build -C ./cmd/http/client -o ../../../http_client.out

grpc_server:
	go build -C ./cmd/grpc/server -o ../../../grpc_server.out

grpc_client:
	go build -C ./cmd/grpc/client -o ../../../grpc_client.out

http: http_server http_client

grpc: grpc_server grpc_client
