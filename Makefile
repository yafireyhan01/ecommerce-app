serve:
	go run cmd/api/main.go

migrate_new:
	migrate create -ext sql -dir migrations -seq $(name)

export_url:
	export POSTGRESQL_URL='postgres://postgres:root@localhost:5432/synapsis_db?sslmode=disable'

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

.PHONY: migrate_new migrate_up export_url serve