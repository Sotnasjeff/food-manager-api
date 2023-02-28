postgres:
	docker run --name postgres14.5 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root movies_manager

createtable:
	docker exec -it postgres14.5 bash
	psql -c 'CREATE TABLE movies(id VARCHAR(5), name VARCHAR(30), description VARCHAR(60))'

server: 
	go run main.go

.PHONY: postgres createdb createtable server
