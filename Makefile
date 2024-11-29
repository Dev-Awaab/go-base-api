.PHONY: postgresinit startdb stopdb clean postgres createdb migrateup migratedown server sqlc newmigration

# =========================================================
# Project Commands
# =========================================================

# Initialize a new PostgreSQL container
postgresinit:
	@echo "Initializing a new PostgreSQL container..."
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -d postgres:15-alpine
	@echo "PostgreSQL container initialized. Accessible on port 5433."

# Start the PostgreSQL container if it is stopped
startdb:
	@echo "Starting the PostgreSQL container..."
	docker start postgres15
	@echo "PostgreSQL container started."

# Access the PostgreSQL CLI
postgres:
	@echo "Accessing PostgreSQL CLI..."
	docker exec -it postgres15 psql -U root
	@echo "Exited PostgreSQL CLI."

# Create a new database
createdb:
	@echo "Creating the database 'go-server-db'..."
	docker exec -it postgres15 createdb --username=root --owner=root go-server-db
	@echo "Database 'go-server-db' created."

# =========================================================
# Migration Commands
# =========================================================

# Run all migrations to apply the latest database schema
migrateup:
	@echo "Applying migrations..."
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-server-db?sslmode=disable" -verbose up
	@echo "Migrations applied."

# Rollback the most recent migration
migratedown:
	@echo "Rolling back migrations..."
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-server-db?sslmode=disable" -verbose down
	@echo "Migrations rolled back."

# Generate a new migration file
newmigration:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name is required. Usage: make newmigration name=your_migration_name"; \
		exit 1; \
	fi
	@echo "Creating new migration files with name '$(name)'..."
	migrate create -seq -ext sql -dir db/migrations "$(name)"
	@echo "Migration files created in db/migrations/"

# =========================================================
# Development Server
# =========================================================

# Run the Go application
server:
	@echo "Starting the Go application..."
	go run cmd/server/main.go

# =========================================================
# SQLC Commands
# =========================================================

# Generate Go code from SQL files
sqlc:
	@echo "Running sqlc to generate database code..."
	sqlc generate
	@echo "Database code generated successfully."

# =========================================================
# Utility Commands
# =========================================================

# Clean up the PostgreSQL container (stops and removes it)
clean:
	@echo "Stopping and removing the PostgreSQL container..."
	docker stop postgres15 && docker rm postgres15
	@echo "PostgreSQL container removed."

# Stop the running PostgreSQL container
stopdb:
	@echo "Stopping the PostgreSQL container..."
	docker stop postgres15
	@echo "PostgreSQL container stopped."

# Display all available commands
help:
	@echo "Available Makefile Commands:"
	@echo "  postgresinit     - Initialize a new PostgreSQL container."
	@echo "  startdb          - Start the PostgreSQL container if it is stopped."
	@echo "  stopdb           - Stop the running PostgreSQL container."
	@echo "  clean            - Stop and remove the PostgreSQL container."
	@echo "  postgres         - Access the PostgreSQL CLI."
	@echo "  createdb         - Create a new database named 'go-server-db'."
	@echo "  migrateup        - Apply all migrations to update the database schema."
	@echo "  migratedown      - Rollback the most recent migration."
	@echo "  newmigration     - Create a new migration file. Usage: make newmigration name=your_migration_name"
	@echo "  sqlc             - Generate Go code from SQL files."
	@echo "  server           - Run the Go application."
	@echo "  help             - Display this help message."
