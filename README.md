# 🎂 Age Calculator - Production Ready Full Stack Go Application

A production-ready full-stack web application built with **Go**, **Fiber**, **PostgreSQL**, and **SQLC** that calculates a person's exact age in **Years, Months, and Days** while demonstrating modern backend engineering practices such as clean architecture, request tracing, structured logging, validation, and secure database operations.

## 🌐 Live Demo

🚀 **Live Application:** https://age-calculator-703761824443.us-central1.run.app/

📂 **Source Code:** https://github.com/Maheshdindur/Age-Calculator

---

## ⭐ Highlights

- Built using Go, Fiber, PostgreSQL, SQLC, and Uber Zap
- Implemented Clean Architecture with clear separation of concerns
- Developed full CRUD REST APIs
- Added structured logging and request tracing
- Implemented request validation and centralized error handling
- Protected database operations using parameterized SQL queries
- Deployed on Google Cloud Run
- Responsive frontend using HTML, CSS, and JavaScript

---

## 📖 Overview

Many age calculator applications focus only on age computation.

This project was designed to demonstrate how a real-world backend service should be structured and deployed using modern engineering practices.

The application allows users to:

- Create age records
- View stored calculations
- Update user information
- Delete records
- Calculate age accurately in Years, Months, and Days

while maintaining production-quality standards such as:

- Clean Architecture
- Type-safe database access
- Request tracing
- Structured logging
- Validation
- Cloud deployment

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

- Parameterized SQL Queries
- SQL Injection Protection
- Input Validation
- Centralized Error Handling

### 📊 Observability

- Structured JSON Logging
- Unique Request IDs
- Request Tracking
- Error Logging

### 🎨 Responsive Frontend

- Clean User Interface
- Mobile Responsive Design
- Dynamic API Integration

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

## 🔄 Request Flow

```text
Client Request
      ↓
Fiber Router
      ↓
Middleware
(Request ID, Logging, Validation)
      ↓
Handler Layer
      ↓
Service Layer
(Age Calculation Logic)
      ↓
Repository Layer
      ↓
SQLC Generated Queries
      ↓
PostgreSQL Database
      ↓
JSON Response
```

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

## 🚀 Getting Started

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

```env
DATABASE_URL=postgres://username:password@localhost:5432/user_db?sslmode=disable
SERVER_PORT=3000
```

### Run Application

```bash
go run cmd/server/main.go
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

## 🚀 Future Improvements

- JWT Authentication
- Role-Based Access Control (RBAC)
- Swagger Documentation
- Redis Caching
- CI/CD with GitHub Actions
- Kubernetes Deployment

---

## 👨‍💻 Author

**Mahesh Dindur**

GitHub: https://github.com/Maheshdindur

---

⭐ If you found this project useful, consider giving it a star.
