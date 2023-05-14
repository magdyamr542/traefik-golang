server1:
	go run cmd/server/server.go --server-id "Worker 1" --port 1234

server2:
	go run cmd/server/server.go --server-id "Worker 2" --port 1235

server3:
	go run cmd/server/server.go --server-id "Worker 3" --port 1236
