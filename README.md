# Go - Backend Development Project (User Management API)

This is a RESTful API built with Go using the Fiber framework. It manages users with a name and date of birth, and dynamically calculates their age.

## 🚀 Getting Started

### Prerequisites
- **Go** (v1.22+)
- **PostgreSQL** or **MySQL** (PostgreSQL is default in config)
- **SQLC** (if you want to regenerate DB code)

### 🛠️ Setup

1. **Clone the repository** (or navigate to the project folder).
2. **Install dependencies**:
   ```bash
   go mod tidy
   ```
3. **Environment Variables**:
   Copy `.env.example` to `.env` and update your database credentials.
   ```bash
   cp .env.example .env
   ```
4. **Database Migration**:
   Run the SQL found in `db/migrations/000001_create_users_table.up.sql` in your database.

### 🏃 Running the Server

```bash
go run cmd/server/main.go
```
The server will start on port `3000` by default.

## 🔄 API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/users` | Create a new user |
| GET | `/users/:id` | Get user by ID (includes age) |
| PUT | `/users/:id` | Update user details |
| DELETE | `/users/:id` | Delete a user |
| GET | `/users` | List all users (includes age) |

### Sample Request (POST `/users`)
```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

## 🏗️ Project Structure
- `cmd/server/`: Entry point.
- `internal/handler/`: HTTP request handling.
- `internal/service/`: Business logic (Age calculation).
- `internal/repository/`: Database access.
- `db/sqlc/`: Generated database code.
- `internal/middleware/`: Request ID and Logging.
