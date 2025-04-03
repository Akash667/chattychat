postgres:
	docker exec --user="postgres" -it  ed0 psql
migrate:
	 migrate -path server/db/migrations -database "postgresql://root:mysecretpassword@localhost:5433/?sslmode=disable" -verbose up
runbackend:
	go run server/cmd/main.go &
runfrontend:
	cd client/my-app/ && npx next dev

runapp: runbackend runfrontend
.PHONY: postgres  migrate runbackend runfrontend runapp
	@echo "Running both backend and frontend"
	@echo "Backend running on http://localhost:8080"
	@echo "Frontend running on http://localhost:3000"
	@echo "Press Ctrl+C to stop both servers"
	@wait