build:
	go build

db:
	-docker run --name pg_shrt --rm -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:11
	sleep 2
	./shrt -db

setup: build db

destroy:
	-docker container stop pg_shrt

run: build
	./shrt
