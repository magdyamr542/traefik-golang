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


# TLS and certificates 
ca:
	# Generate a private key for the CA:
	cd tls && openssl genpkey -algorithm RSA -out ca.key
	# Create a self-signed certificate for the CA:
	cd tls && openssl req -new -x509 -key ca.key -out ca.crt

serverkey:
	# Generate a private key for the server:
	cd tls && openssl genpkey -algorithm RSA -out server.key
	# Create a certificate signing request (CSR) for the server:
	cd tls && openssl req -new -key server.key -out server.csr -config san.cnf
	# Create a server certificate by signing the CSR with the CA:
	cd tls && openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extensions reqexts -extfile san.cnf

# clientkey:
# 	# Generate a private key for the client:
# 	openssl genpkey -algorithm RSA -out client.key
# 	# Create a certificate signing request (CSR) for the client:
# 	openssl req -new -key client.key -out client.csr
# 	# Create a client certificate by signing the CSR with the CA:
# 	openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt

