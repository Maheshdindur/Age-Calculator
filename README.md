# 🎂 Age Calculator - Production Ready Full Stack Go Application

A production-ready full-stack web application built with **Go**, **Fiber**, **PostgreSQL**, and **SQLC** that calculates a person's exact age in **Years, Months, and Days** while demonstrating modern backend engineering practices such as clean architecture, request tracing, structured logging, validation, and secure database operations.

## 🌐 Live Demo

**Application:** https://age-calculator-703761824443.us-central1.run.app/

**Source Code:** https://github.com/Maheshdindur/Age-Calculator

---

## ✨ Features

### 🎯 Accurate Age Calculation

- Calculates age precisely in Years, Months, and Days
- Handles leap years
- Handles varying month lengths
- Prevents future dates
- Validates user input

### 👤 User Management (CRUD)

- Create users
- View saved calculations
- Update user information
- Delete users

### 🔒 Secure Backend

- Parameterized SQL queries
- SQL injection protection
- Input validation
- Centralized error handling

### 📊 Observability

- Structured JSON logging with Uber Zap
- Unique Request IDs
- Request tracing
- Error monitoring

---

## 🏗️ Architecture

```text
├── cmd
│   └── server
│       └── main.go
│
├── internal
│   ├── handler
│   ├── service
│   ├── repository
│   ├── middleware
│   └── logger
│
├── db
│   ├── migrations
│   ├── query
│   └── sqlc
│
├── public
│
├── Dockerfile
├── docker-compose.yml
└── README.md
```

### Layer Responsibilities

#### Handler Layer
- HTTP request handling
- Request validation
- Response formatting

#### Service Layer
- Business logic
- Age calculation
- Validation rules

#### Repository Layer
- Database operations
- Query execution
- Data persistence

#### SQLC Layer
- Type-safe generated queries
- Compile-time query validation

---

## 🛠️ Tech Stack

### Backend
- Go 1.22+
- Fiber
- PostgreSQL
- SQLC
- Uber Zap
- Validator

### Frontend
- HTML
- CSS
- JavaScript

### DevOps & Cloud
- Docker
- Google Cloud Run
- GitHub

---

## 🗄️ Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22+
- PostgreSQL
- Git

### Clone Repository

```bash
git clone https://github.com/Maheshdindur/Age-Calculator.git

cd Age-Calculator
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Environment Variables

Create a `.env` file:

```env
DATABASE_URL=postgres://username:password@localhost:5432/user_db?sslmode=disable
SERVER_PORT=3000
```

### Run Application

```bash
go run cmd/server/main.go
```

Open:

```text
http://localhost:3000
```

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|----------|----------|-------------|
| POST | /users | Create User |
| GET | /users | Get All Users |
| GET | /users/:id | Get User By ID |
| PUT | /users/:id | Update User |
| DELETE | /users/:id | Delete User |

---

## ☁️ Deployment

Deployed on Google Cloud Run:

https://age-calculator-703761824443.us-central1.run.app/

---

## 📚 Key Learnings

- REST API Development
- Clean Architecture
- PostgreSQL Integration
- SQLC Query Generation
- Middleware Design
- Structured Logging
- Docker Containerization
- Cloud Deployment
- Input Validation
- Production Backend Practices

---

## 👨‍💻 Author

**Mahesh Dindur**

GitHub: https://github.com/Maheshdindur

LinkedIn: https://linkedin.com/in/mahesh-dindur

---

⭐ If you found this project useful, consider giving it a star.
