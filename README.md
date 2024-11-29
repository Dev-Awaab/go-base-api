# **Go-Base-API**

A starter template for building RESTful APIs using Go with the Gin framework, following clean architecture principles. This project is structured for scalability and maintainability.

---

## **Author**

Created by **Abdul-Fattah Abdul-Kareem**

---

## **Features**

- Clean and modular folder structure.
- SQLC for type-safe database queries.
- Configurable via `.env` file.
- Centralized error handling and logging.
- Middleware for request validation and security.
- Docker support for easy deployment.
- PostgreSQL integration with migrations.

---

## **Folder Structure**

```plaintext
GO-BASE-API/
├── cmd/
│   └── server/
│       └── main.go           # Entry point for the application
├── config/
│   └── config.go             # Configuration loader
├── db/
│   ├── migrations/           # Database migration files
│   ├── query/                # SQL queries for complex operations
│   ├── sqlc/                 # Auto-generated SQLC database queries
│   └── db.go                 # Database connection initializer
├── internal/
│   ├── user/                 # User-related functionality
│   │   ├── user_handler.go   # HTTP handlers for user routes
│   │   ├── user_repository.go# Database interactions for users
│   │   ├── user_service.go   # Business logic for user operations
│   │   └── user.go           # User model and DTOs
├── pkg/
│   ├── logger/               # Centralized logging
│   ├── middleware/           # Middleware for validation and security
│   ├── router/               # HTTP router initialization
│   └── utils/                # Utility functions (e.g., response helpers)
├── .env                      # Environment variables
├── .env.example              # Example .env file for setup
├── .gitignore                # Ignored files for git
├── Dockerfile                # Dockerfile for containerizing the app
├── go.mod                    # Go module file
├── go.sum                    # Dependencies lock file
├── Makefile                  # Task automation for migrations and server
└── README.md                 # Project documentation
```

## **Getting Started**

### **Prerequisites**

- **Go** (v1.20+)
- **Docker** (for running PostgreSQL)
- **`migrate` CLI** (for database migrations)
- **PostgreSQL**

## **Setup**

### **1. Clone the repository**

```bash
git clone https://github.com/your-username/go-base-api.git
cd go-base-api
```

### **2. Install dependencies**

```bash
go mod tidy
```

### **3. Create a .env file: Copy the .env.example file and configure the variables**

```bash
cp .env.example .env
```

### **4. Check the Makefile for command**

```bash
make help
```

## **Makefile Commands**

The project includes a `Makefile` to simplify common tasks:

| Command        | Description                                                                       |
| -------------- | --------------------------------------------------------------------------------- |
| `postgresinit` | Initialize a new PostgreSQL container.                                            |
| `startdb`      | Start the PostgreSQL container if it is stopped.                                  |
| `stopdb`       | Stop the running PostgreSQL container.                                            |
| `clean`        | Stop and remove the PostgreSQL container.                                         |
| `postgres`     | Access the PostgreSQL CLI.                                                        |
| `createdb`     | Create a new database named `go-server-db`.                                       |
| `migrateup`    | Apply all migrations to update the database schema.                               |
| `migratedown`  | Rollback the most recent migration.                                               |
| `newmigration` | Create a new migration file. Usage: `make newmigration name=your_migration_name`. |
| `sqlc`         | Generate Go code from SQL files.                                                  |
| `server`       | Run the Go application.                                                           |
| `help`         | Display the list of available Makefile commands.                                  |

```bash
make help
```

## **Build and Run the Docker Image**

### **1. Build the Docker image:**

```bash
docker build -t go-base-api .
```

### **2. Run the container:**

```bash
docker run -p 8080:8080 --env-file=.env go-base-api
```

bash
Copy code
docker build -t go-base-api .
Run the container:

bash
Copy code
docker run -p 8080:8080 --env-file=.env go-base-api

## **Technologies Used**

- **Go**: Programming language.
- **Gin**: Web framework for Go.
- **PostgreSQL**: Relational database.
- **SQLC**: Type-safe database queries.
- **Docker**: Containerization.
- **Makefile**: Task automation.

## **Contributing**

Feel free to submit issues or pull requests for improvements.
