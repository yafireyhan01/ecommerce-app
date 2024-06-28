serve:
	go run main.go

migrate_new:
	migrate create -ext sql -dir db/migrations -seq $(name)

export_url:
	export POSTGRESQL_URL='postgres://postgres:root@localhost:5432/synapsis_db?sslmode=disable'

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

.PHONY: migrate_new migrate_up export_url serve