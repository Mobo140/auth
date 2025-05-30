# Auth Service

**Auth Service** is an authentication and authorization microservice written in Go using gRPC. It provides user management, authentication, and access control functionality.

---

## 🚀 Quick Start

### Requirements

- Go 1.21+
- Docker & Docker Compose (optional)
- CLI tools: protoc, grpc-gateway, etc. (Makefile automates installation)

### 1. Clone the project

```bash
git clone https://github.com/your-org/auth-service.git
cd auth-service
````

### 2. Run setup

```bash
make setup
```

This will:

- Install all necessary tools (if missing)
- Generate gRPC, gateway, validators, swagger, etc.
- Start services (via Docker Compose if configured)

---

## 📦 Features

- Username/password authentication
- JWT tokens with access and refresh tokens
- User management (CRUD)
- Role-based access control (Admin/User/Guest)
- gRPC Gateway for REST API
- Swagger documentation

---

## 🧰 Makefile Commands

| Command                  | Description                                                |
| ------------------------ | ---------------------------------------------------------- |
| `make setup`             | Install dependencies, generate code, start services        |
| `make install-deps`      | Install CLI tools (protoc, grpc-gateway, validate, linter) |
| `make generate`          | Generate gRPC, gateway, swagger, validators                |
| `make generate-auth-api` | Generate code from `auth.proto`                            |
| `make lint`              | Run static code analysis (golangci-lint)                   |
| `make test`              | Run unit tests                                             |

---

## 🔧 Tech Stack

- Go 1.21+
- gRPC + Protobuf
- PostgreSQL (user storage)
- Redis (caching)
- Jaeger (tracing)
- gRPC Gateway + Swagger (REST API + docs)
- Zap (logging)
- Prometheus + Grafana (monitoring)

---

## 🔐 Security

- Password hashing
- JWT tokens (access token: 15 minutes, refresh token: 24 hours)
- Role-based access control (Admin, User, Guest)
- Refresh token mechanism

---

## 🛠 API

### User Service (`user_v1`)

- Create — create a new user
- Get — get user info by ID
- GetUsers — list users with pagination
- Update — update user info
- Delete — delete user

### Auth Service (`auth_v1`)

- Login — authenticate and get refresh token
- GetRefreshToken — refresh the refresh token
- GetAccessToken — get access token using refresh token

### Access Service (`access_v1`)

- Check — check access rights to protected endpoints

---

## 🔄 Authentication Flow

1. Client calls `Login` with username and password
2. Receives a refresh token
3. Uses refresh token to get access token
4. Access token is used to access protected resources
5. Refresh token expiration requires re-authentication

---

## 🔄 Interaction with Other Services

- Client authenticates with Auth Service
- Other services receive access token and validate it through Auth Service (`Check` method)
- Auth Service validates token and user permissions

---

## 👥 Roles and Permissions

| Role  | Access Level                             |
| ----- | ---------------------------------------- |
| Admin | Full access to all endpoints             |
| User  | Limited access to user-related functions |
| Guest | Access to public endpoints only          |

---

## 📁 Project Layout

```
├── api/                      # Protobuf definitions
├── cmd/                      # Entrypoints (main.go)
├── internal/                 # Private application code
│   ├── app/                  # Application wiring (DI, lifecycle)
│   ├── client/               # External clients (e.g., Redis)
│   ├── config/               # Config loading
│   ├── converter/            # Data transformers between transport and service layers
│   ├── interceptor/          # gRPC interceptors (auth, logging, etc.)
│   ├── model/                # Domain models and constants
│   ├── metric/               # Initialization of metrics
│   ├── ratelimiter/          # Rate limiting logic
│   ├── repository/           # Storage access (Postgres, etc.)
│   ├── service/              # Business logic
│   └── transport/
│       └── handlers/         # gRPC handlers
├── migrations/               # Database schema (Goose)
├── pkg/                      # Generated code and shared helpers
├── vendor.protogen/          # External proto dependencies
├── Makefile                  # Dev utility commands
├── env/                      # Environment variables


---

## 📄 License

MIT
