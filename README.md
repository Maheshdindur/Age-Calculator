# 🎂 Age Calculator API & Frontend 🚀

  A full-stack User Management system built with **Go** that calculates precise human age in **Years, Months, and Days**. It combines a clean backend architecture, PostgreSQL data storage, and a responsive frontend for managing age
  records.

  ---

  ## 🛠️ Features
  - **Accurate Age Logic**: Dynamically calculates age down to the day, handling varying month lengths and leap years.
  - **Full CRUD**: Create, Read, Update, and Delete users through a modern web UI.
  - **Data Integrity**: Prevents future dates from being stored and validates all inputs.
  - **Observability**: Structured JSON logging (Uber Zap) and unique Request IDs for every transaction.
  - **Database Safety**: Uses parameterized SQL queries to reduce SQL Injection risk.

  ---

  ## 🏗️ Project Architecture
  The project follows a **Layered Architecture** so each part of the application has a clear responsibility:
  - **`cmd/server/`**: Application entry point and dependency injection.
  - **`internal/handler/`**: HTTP layer (GoFiber) – handles requests and responses.
  - **`internal/service/`**: Business logic layer – handles the complex age calculations.
  - **`internal/repository/`**: Data access layer – communicates with PostgreSQL.
  - **`db/sqlc/`**: Type-safe database code generated from raw SQL.
  - **`public/`**: Frontend assets (HTML/CSS/JS).

  ---

  ## 🚀 Getting Started

  ### 1. Prerequisites
  - **Go** (v1.22 or higher)
  - **PostgreSQL** (Running locally or in the cloud)
  - **Git**

  ### 2. Database Setup
  1. Create a database named `user_db` in your PostgreSQL instance.
  2. Run the migration script to create the `users` table:
     ```sql
     -- Path: db/migrations/000001_create_users_table.up.sql
     CREATE TABLE users (
         id SERIAL PRIMARY KEY,
         name TEXT NOT NULL,
         dob DATE NOT NULL
     );

  ### 3. Installation & Configuration

  1. Clone the repository:

     git clone https://github.com/Maheshdindur/Age-Calculator.git
     cd Age-Calculator

  2. Install dependencies:

     go mod tidy

  3. Set up Environment Variables:
     Create a .env file in the root directory:

     DATABASE_URL=postgres://your_username:your_password@localhost:5432/user_db?sslmode=disable
     SERVER_PORT=3000

  ### 4. Running the Application

  Start the server:

  go run cmd/server/main.go

  Open your browser to: http://localhost:3000 (http://localhost:3000)

  ———

  ## 🧪 API Endpoints

   Method    Endpoint      Description
  ━━━━━━━━  ━━━━━━━━━━━━  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
   POST      /users        Create a user & calculate age
  ────────  ────────────  ───────────────────────────────
   GET       /users/:id    Get specific user details
  ────────  ────────────  ───────────────────────────────
   PUT       /users/:id    Update user details
  ────────  ────────────  ───────────────────────────────
   DELETE    /users/:id    Remove a user
  ────────  ────────────  ───────────────────────────────
   GET       /users        List all saved calculations

  ———

  ## 📦 Deployment

  This app is deployed on Google Cloud Platform. Try the live version here: https://age-calculator-703761824443.us-central1.run.app/ (https://age-calculator-703761824443.us-central1.run.app/)
