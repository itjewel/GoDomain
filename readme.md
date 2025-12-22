# GoDomain – Golang Backend Project

## Project Run

```bash
go run ./cmd/server/main.go
```

## Project URLs

API Base URL: http://localhost:7000/
Swagger UI: http://localhost:7000/swagger/index.html

## Swagger Generate

swag init -g cmd/server/main.go -o internal/interfaces/http/swagger

---

## Key Features

- RESTful API implementation using Clean Architecture and Use Case pattern
- Full CRUD operations (Create, Read, Update, Delete) with proper validation
- PostgreSQL integration with connection pooling and structured schema design
- Repository Pattern for database abstraction and testability
- Domain-Driven Design (DDD) with clear separation of concerns
- Strongly typed data models using Go structs
- Interface-based design for loose coupling and scalability
- Swagger integration for API documentation and interactive testing
- Redis integration for caching, performance optimization, and rate limiting
- RabbitMQ support for asynchronous and event-driven processing
- gRPC support for high-performance service-to-service communication
- Centralized singleton-based configuration management
- Struct-level and business-rule validation strategy
- Docker-ready application for containerized deployment
- Grafana integration for monitoring and metrics visualization
- PGAdmin 4 support for PostgreSQL database management
- PostgreSQL export and import support for backup and restore
- Scalable and maintainable architecture following SOLID principles
- Production-ready project structure suitable for real-world systems
- Designed for learning, extensibility, and enterprise-grade backend development

---

## Architecture Principles

- Clean Architecture
- SOLID Principles
- Separation of Concerns
- Scalable and Maintainable Design

---

## Project Goal

Build a production-ready Golang backend application using modern architecture patterns, best practices, and real-world tooling.
