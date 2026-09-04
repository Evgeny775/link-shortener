# 🔗 Link Shortener

A simple URL shortening service written in **Go**.

> 🚧 **Work in Progress**
> The project is currently under active development. Some features, including authentication, are not fully implemented yet.

## 📌 About

This project is a backend service for creating and managing shortened URLs.

The main goal of the project is to practice building a backend application in Go with a layered architecture, working with HTTP handlers, middleware, databases and authentication.

### Current functionality

* 🔗 Create shortened links
* ↪️ Redirect from a short link to the original URL
* ✏️ Update links
* 🗑️ Delete links
* 👤 User registration
* 🔐 Authentication middleware
* 🗄️ PostgreSQL database
* 🧩 GORM ORM
* 🌐 CORS middleware
* 📝 Request logging
* 🐳 PostgreSQL via Docker Compose

The project uses Go's standard `net/http` package for the HTTP server and routing.

## 🛠️ Tech Stack

* **Go 1.25+**
* **net/http**
* **PostgreSQL**
* **GORM**
* **Docker / Docker Compose**
* **go-playground/validator**
* **godotenv**

Dependencies and the required Go version are defined in `go.mod`.

## 🏗️ Project Structure

```text
.
├── cmd/
│   └── main.go              # Application entry point
│
├── configs/                 # Application configuration
│
├── internal/
│   ├── auth/                # Authentication
│   ├── link/                # Link management
│   ├── middleware/          # HTTP middleware
│   └── user/                # User management
│
├── migrations/
│   └── auto/                # Database migrations
│
├── pkg/
│   ├── db/                  # Database connection
│   ├── req/                 # HTTP request helpers
│   └── res/                 # HTTP response helpers
│
├── docker-compose.yaml
├── go.mod
└── go.sum
```

The current project separates authentication, links, users and middleware into independent packages under `internal/`.

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Evgeny775/link-shortener.git
cd link-shortener
```

### 2. Configure environment variables

Create a `.env` file in the project root.

Example:

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres_docker
DB_PASSWORD=my_pass
DB_NAME=link_shortener

SERVER_PORT=8080
```

> The exact configuration options may change while the project is under development.

### 3. Start PostgreSQL

The project includes a Docker Compose configuration for PostgreSQL.

```bash
docker compose up -d
```

By default, PostgreSQL is exposed on port `5433` on the host.

### 4. Run the application

```bash
go run ./cmd
```

The HTTP server currently starts on:

```text
http://localhost:8080
```

## 📡 API

### Authentication

#### Register

```http
POST /auth/register
```

Example request:

```json
{
  "email": "user@example.com",
  "username": "user",
  "password": "password"
}
```

#### Login

```http
POST /auth/login
```

> Authentication is currently under development.

### Links

#### Create a link

```http
POST /link
```

Requires authentication.

Example request:

```json
{
  "url": "https://example.com"
}
```

#### Redirect

```http
GET /{hash}
```

Redirects the client to the original URL.

#### Update a link

```http
PATCH /link/{id}
```

Requires authentication.

Example request:

```json
{
  "url": "https://example.org",
  "hash": "example"
}
```

#### Delete a link

```http
DELETE /link/{id}
```

Requires authentication.

The current router defines these endpoints directly using `http.ServeMux`.

## 🔄 How It Works

When a user creates a link, the service generates a short random hash and stores it together with the original URL.

```text
Client
   │
   │ POST /link
   ▼
HTTP Handler
   │
   ▼
Link Service
   │
   ├── Generate short hash
   │
   ├── Check uniqueness
   │
   ▼
Link Repository
   │
   ▼
PostgreSQL
```

When someone opens the shortened URL:

```text
GET /abc123
     │
     ▼
Link Handler
     │
     ▼
Link Service
     │
     ▼
PostgreSQL
     │
     ▼
HTTP Redirect
     │
     ▼
Original URL
```

The current implementation generates six-character hashes using letters and digits and retries when a generated hash conflicts with an existing one.

## 🧱 Architecture

The application follows a simple layered structure:

```text
Handler
   ↓
Service
   ↓
Repository
   ↓
Database
```

This separation keeps HTTP-specific logic, business logic and database operations independent from each other.

Middleware is also composed into a chain before the router:

```text
HTTP Request
     ↓
   CORS
     ↓
  Logging
     ↓
   Router
     ↓
  Handler
```

The application currently uses CORS and logging middleware globally.

## 🗺️ Roadmap

The project is still being developed.

Planned improvements:

* [ ] Complete authentication flow
* [ ] Implement JWT authentication
* [ ] Improve authorization
* [ ] Add user ownership for links
* [ ] Add input validation
* [ ] Improve error handling
* [ ] Add unit tests
* [ ] Add integration tests
* [ ] Improve logging
* [ ] Add API documentation
* [ ] Add Dockerfile for the application
* [ ] Improve database migrations
* [ ] Add graceful server shutdown
* [ ] Add configuration validation
* [ ] Add CI with GitHub Actions

## 🎯 Purpose

This project is primarily a learning project focused on backend development with Go.

While building it, I am practicing:

* Go backend development
* REST API design
* HTTP handlers and middleware
* Layered architecture
* PostgreSQL
* GORM
* Authentication and authorization
* Error handling
* Docker
* Database migrations
* Writing maintainable Go code

## 📄 License

This project is currently intended for educational purposes.
