# Movies CRUD Application

A RESTful API for managing movies built with **Go**, **Gin**, **GORM**, and **UberFx**. This application provides secure endpoints for performing CRUD operations on movies, complete with JWT-based authentication and role-based authorization.

---

## Features

- **CRUD Operations for Movies:** Create, read, update, and delete movie records.
- **JWT-based Authentication:** Secure endpoints with JSON Web Tokens.
- **Role-based Authorization:** Distinguish between regular and admin users.
- **PostgreSQL Integration:** Reliable data storage using PostgreSQL.
- **Swagger API Documentation:** Interactive API docs available for testing endpoints.
- **Docker Support:** Containerized deployment with Docker Compose for ease of setup.

---

## Project Structure

```plaintext
Copymovies-crud-app/
├── cmd/
│   └── api/
│       └── main.go             # Application entry point with UberFx
├── internal/
│   ├── config/                 # Configuration management
│   ├── controller/             # HTTP request handlers
│   ├── dto/                    # Data Transfer Objects
│   ├── middleware/             # HTTP middleware (auth)
│   ├── model/                  # Database models
│   ├── repository/             # Database operations
│   └── service/                # Business logic
├── pkg/
│   ├── auth/                   # Authentication utilities
│   ├── database/               # Database connection
│   └── validator/              # Input validation
├── api/
│   └── swagger.yaml            # API documentation
├── .env                        # Environment variables (not committed)
├── .env.example                # Example environment variables
├── Dockerfile                  # Docker image definition
├── docker-compose.yml          # Docker Compose configuration
├── go.mod                      # Go module definition
└── go.sum                      # Go module checksums

```
## Prerequisites

- **Go 1.21+**
- **PostgreSQL**
- **Docker & Docker Compose** (for containerized deployment)

---

## Running the Application

### Local Development (Using Go)

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/movies-crud-app.git
    cd movies-crud-app
    ```
   
2. **Modify the `.env` file** with your settings if needed.

3. **Run the application:**

    ```bash
    go run cmd/api/main.go
    ```

---

### Using Docker Compose

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/movies-crud-app.git
    cd movies-crud-app
    ```

2. **Start the services:**

    ```bash
    docker-compose up -d
    ```

3. **Access the application:**
    - **API:** [http://localhost:8080](http://localhost:8080)

---

## Default Admin Credentials

On the first run, the application automatically creates an admin user if one doesn't exist:

- **Username:** `admin`
- **Password:** `adminpassword`

> **Note:** In production, change these credentials by setting the `ADMIN_USERNAME` and `ADMIN_PASSWORD` environment variables.

---

## API Documentation

Swagger documentation is available at:

```bash
http://localhost:8080/swagger/index.html
