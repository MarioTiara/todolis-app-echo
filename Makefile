createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todolistwebapi

dropdb:
	docker exec -it postgres12 dropdb todolistwebapi

migration_new:
	migrate create -ext sql -dir internal/platform/database/migrations -seq todolist_schema

migration_up:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable" -verbose up

migration_down:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable" -verbose down

migration_status:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable" status

migration_fix:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable" force 9

atlas-migration:
	atlas migrate diff --env gorm 

atlas-push-migration:
	atlas migrate push app --dev-url "docker://postgres/15/dev?search_path=public"

atlas-apply-migration:
	atlas migrate apply --dir "atlas://app"  --url "postgres://root:secret@:5432/todolistwebapi?search_path=public&sslmode=disable"

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

docker-db-migration-up:
	$ docker run -v /Users/mpratama/Documents/Coding/Personal/GOLANG/todolistapp/internal/platform/database/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://root:secret@localhost:5432/todolistwebapi?sslmode=disable  up

.PHONY: createdb dropdb atlas-migration atlas-apply-migration docker-db-migration-up


