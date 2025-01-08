postgres:
	docker exec --user="postgres" -it  ed0 psql
migrate:
	 migrate -path server/db/migrations -database "postgresql://root:mysecretpassword@localhost:5433/?sslmode=disable" -verbose up
.PHONY: postgres  migrate
