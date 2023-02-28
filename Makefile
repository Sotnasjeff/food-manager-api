postgres:
	docker run --name postgres14.5 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root movies_manager

createtable:
	docker exec -it postgres14.5 psql -c 'CREATE TABLE movies(id TEXT, name VARCHAR(30), description VARCHAR(60))' -d movies_manager

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

server: 
	go run main.go

.PHONY: postgres createdb createtable server proto evans
