# ShopLite

Lightweight Go REST API for a small shop domain using Gin, GORM (PostgreSQL), and validator.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![Build](https://img.shields.io/badge/build-local-green)
![Tests](https://img.shields.io/badge/tests-go%20test%20./...-brightgreen)

## Tech Stack
- Go, Gin (router)
- GORM + PostgreSQL driver (ORM)
- go-playground/validator (validation)
- Testify (assertions, mocks)

## Project Structure
- cmd/shoplite/main.go – app entrypoint
- config/config.go – env-based configuration
- internal/database – DB connect and auto-migrate
- internal/models – GORM models
- internal/repositories – data access
- internal/services – business logic
- internal/handlers – HTTP handlers with validation
- internal/routes – routes registration
- internal/utils – middleware and JSON responses
- internal/testutil – testing helpers (DB setup, migrations)
- migrations/ – placeholder (AutoMigrate used)

## Setup
1) Clone the repo

```powershell
git clone <your-repo-url>
cd shoplite
```

2) Install dependencies

```powershell
go mod tidy
```

3) Configure environment (option A: env vars)

- SERVER_PORT (default: 8080)
- DB_HOST (default: localhost)
- DB_PORT (default: 5432)
- DB_USER (default: postgres)
- DB_PASSWORD (default: postgres)
- DB_NAME (default: shoplite)
- DB_SSLMODE (default: disable)

Option B: YAML for tests only – edit `config.test.yaml` (uses a separate test database by default: `shoplite_test`). Ensure the DB exists.

4) Run migrations

AutoMigrate runs on startup and in tests. If you prefer manual migration tools (e.g., golang-migrate), place files in `migrations/` and run them before starting the server.

5) Start the server

```powershell
# From repository root
$env:DB_HOST = "localhost"
$env:DB_USER = "postgres"
$env:DB_PASSWORD = "postgres"
$env:DB_NAME = "shoplite"
$env:SERVER_PORT = "8080"

go run ./cmd/shoplite
```

On first start, AutoMigrate creates/updates tables.

## Testing

- Configure a PostgreSQL test database (default from `config.test.yaml` is `shoplite_test`). Create it if missing.
- Run tests:

```powershell
go test ./...
```

With VS Code tasks:
- Open the Command Palette > "Run Task" > "ShopLite: Test All" or "ShopLite: Run Server".
- Debug server: use the launch config "ShopLite: Launch Server".

## API Endpoints

- Customers
  - POST /customers – create
  - GET /customers – list
  - GET /customers/:id – detail
- Products
  - POST /products – create
  - GET /products – list
  - GET /products/:id – detail
- Orders
  - POST /orders – create with items
  - GET /orders – list
  - GET /orders/:id – detail

Response format:

```json
{ "status": "success|error", "message": "...", "data": {}}
```

### Example Requests

Create customer

```http
POST http://localhost:8080/customers
Content-Type: application/json

{ "name": "Alice", "email": "alice@example.com" }
```

Create product

```http
POST http://localhost:8080/products
Content-Type: application/json

{ "name": "Widget", "price": 9.99, "stock": 10 }
```

Create order

```http
POST http://localhost:8080/orders
Content-Type: application/json

{
  "customer_id": 1,
  "order_date": "2025-11-04T12:00:00Z",
  "status": "pending",
  "items": [
    { "product_id": 1, "quantity": 2, "price": 9.99 }
  ]
}
```

## Contribution Guidelines

1. Create a feature branch from main.
2. Add/adjust unit tests for your change.
3. Run `go test ./...` and ensure all pass.
4. Submit a PR with a clear description and screenshots/logs if applicable.

## Notes
- Validation is enforced on inputs; order statuses: pending | paid | shipped | cancelled.
- Business logic is minimal by design; extend with inventory checks and totals as needed.
