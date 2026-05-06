# ERP Setting Service

🚀 **High-Performance ERP Staff Management System** built with Go, Fiber, and PostgreSQL. Designed with **Clean Architecture** principles and production-ready standards.

## ✨ Key Features

- **Clean Architecture:** Modular design (Handler -> Service -> Repository) for high maintainability.
- **High Performance:** Built with [Fiber](https://gofiber.io/) and [pgx](https://github.com/jackc/pgx) for optimized database interactions.
- **Robust Security:** 
  - Input validation using `go-playground/validator`.
  - SQL Injection protection via parameterized queries.
  - Graceful shutdown for data integrity.
- **Containerized:** Fully Dockerized with optimized multi-stage builds.
- **Developer Experience:** 
  - Structured Logging with `slog`.
  - Automated database initialization.
  - Swagger-ready API documentation.

## 🛠 Tech Stack

- **Language:** Go 1.26+
- **Framework:** Fiber v2
- **Database:** PostgreSQL 18
- **Infrastructure:** Docker & Docker Compose
- **Validation:** Go-Validator v10
- **Logging:** Slog (Structured JSON Logging)

## 📁 Project Structure

```text
├── cmd/api             # Application entry point
├── internal/           # Private library code
│   ├── config          # Environment configuration
│   ├── database        # Database connection logic
│   └── helper          # Shared helpers (Validation, etc.)
├── modules/            # Business modules (Domain-driven)
│   └── staffs          # Staff management module
│       ├── dto         # Data Transfer Objects
│       ├── handler     # HTTP Handlers
│       ├── service     # Business Logic
│       └── repository  # Database access
├── response/           # Unified API response format
└── init.sql            # Database schema & seed
```

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Installation & Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/BreakDown-CS/erp-setting-cs.git
   cd erp-setting-cs
   ```

2. **Run with Docker Compose:**
   ```bash
   docker compose up -d
   ```

3. **Verify the installation:**
   The API will be available at `http://localhost:8080`.

## 📡 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/staffs/saveStaff` | Create a new staff member |
| `GET` | `/staffs/getStaffList` | Get paginated list of staff |
| `POST` | `/staffs/getStaffById` | Get staff details by ID |

## 🛡 Security & Best Practices

- **Validation:** Every request is validated before processing.
- **Concurrency:** Uses `pgxpool` for efficient connection management.
- **Environment:** No secrets in code; all configurations are managed via `.env`.
- **Graceful Shutdown:** The server listens for `SIGINT/SIGTERM` to close database connections and finish active requests safely.

---
*Created by [BreakDown] as a showcase of modern backend engineering in Go.*
