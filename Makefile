backend1:
	go run cmd/server/server.go --role backend --server-id 1 --port 1234

backend2:
	go run cmd/server/server.go --role backend --server-id 2 --port 1235

backend3:
	go run cmd/server/server.go --role backend --server-id 3 --port 1236

frontend1:
	go run cmd/server/server.go --role frontend --server-id 1 --port 1237

frontend2:
	go run cmd/server/server.go --role frontend --server-id 2 --port 1238
