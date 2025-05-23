run:
	go run cmd/main.go

docker:

	sudo docker run --name=url -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

migrate:
	
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up