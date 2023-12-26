
#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#
run-test:
	go test -cover  ./...   

migration-up:
	$ docker run -v /Users/mpratama/Documents/Coding/Personal/todolistapi/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable  up

migration-down:
	$ docker run -v /Users/mpratama/Documents/Coding/Personal/todolistapi/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable  down

.PHONY: migration-up migration-down


