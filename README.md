# Hexagonal Architecture REST API with Go & MongoDB

🏗️ A RESTful API built with Go implementing Hexagonal Architecture (Ports and Adapters) pattern with MongoDB integration. Features clean architecture separation with domain-driven design principles.

## 🚀 Features

- **Clean Architecture**: Hexagonal Architecture (Ports and Adapters)
- **User Management**: Registration and listing endpoints
- **MongoDB Integration**: Proper BSON handling and validation
- **Error Handling**: Comprehensive error responses
- **CORS Support**: Cross-origin resource sharing
- **Validation**: Input validation and password confirmation
- **Health Check**: System health monitoring

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Framework**: Fiber v2
- **Database**: MongoDB
- **Architecture**: Hexagonal/Clean Architecture
- **Patterns**: Repository Pattern, Domain-Driven Design

## 📁 Project Structure

```
├── cmd/                    # Application entrypoint
│   └── main.go
├── internal/
│   ├── adapter/           # External adapters
│   │   ├── handler/       # HTTP handlers
│   │   │   ├── route.go
│   │   │   └── user.go
│   │   └── storage/       # Database implementations
│   │       └── mongo/
│   │           ├── db.go
│   │           └── repository/
│   │               └── user.go
│   └── core/              # Business logic
│       ├── domain/        # Domain entities
│       │   ├── error.go
│       │   └── user.go
│       ├── port/          # Interfaces
│       │   └── user.go
│       └── service/       # Business services
│           └── user.go
├── .env                   # Environment variables
├── go.mod
└── go.sum
```

## 🚀 API Endpoints

### User Management
- `POST /api/v1/user/register` - User registration
- `GET /api/v1/user/list` - List all users

### System
- `GET /health` - Health check

## 🔧 Installation & Setup

### Prerequisites
- Go 1.24 or higher
- MongoDB running locally or remote connection

### Environment Variables
Create a `.env` file in the root directory:

```env
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=hexagonal_test
PORT=8080
```

### Install Dependencies
```bash
go mod tidy
```

### Run the Application
```bash
go run cmd/main.go
```

The server will start on `http://127.0.0.1:8080`

## 📝 API Usage Examples

### Register a User
```bash
curl -X POST http://127.0.0.1:8080/api/v1/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "password_confirmation": "password123"
  }'
```

### List Users
```bash
curl http://127.0.0.1:8080/api/v1/user/list
```

### Health Check
```bash
curl http://127.0.0.1:8080/health
```

## 🏗️ Architecture Overview

This project follows **Hexagonal Architecture** principles:

- **Domain Layer**: Core business logic and entities
- **Port Layer**: Interfaces defining contracts
- **Adapter Layer**: External implementations (HTTP handlers, database repositories)
- **Service Layer**: Business logic orchestration

### Key Components

1. **Domain**: `User` entity with business rules
2. **Ports**: `UserService` and `UserRepository` interfaces
3. **Adapters**: HTTP handlers and MongoDB repository implementations
4. **Services**: Business logic implementation

## 🔍 Error Handling

The API returns structured error responses:

```json
{
  "error": "Error message",
  "details": "Detailed error information (development only)"
}
```

## 🛡️ Validation

- Email format validation
- Password length requirements (minimum 8 characters)
- Password confirmation matching
- Duplicate email prevention

## 🚦 Development

### Running Tests
```bash
go test ./...
```

### Code Structure Guidelines
- Keep business logic in the `core` package
- External dependencies in `adapter` package
- Use interfaces for testability
- Follow Go naming conventions

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

Built with ❤️ using Go and clean architecture principles.
