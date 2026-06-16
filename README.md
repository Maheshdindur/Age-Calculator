# 🎂 User Lifecycle & Dynamic Age API 🚀

Welcome to the **User Management API**! This project is a robust, production-ready backend system built with Go. It doesn't just store names; it accurately tracks the human lifecycle by calculating a user's exact age in **Years, Months, and Days** in real-time.

---

## 🌟 What does this project do?
Imagine you are running a service where you need to know exactly how old your members are—not just "30 years old," but "30 years, 2 months, and 5 days old." 

This API allows you to:
1.  **Register Users**: Save their name and Date of Birth (DOB).
2.  **Instant Calculation**: Automatically calculate their detailed age every time you view their profile.
3.  **Manage Data**: Update or remove users through a clean, modern web interface.

---

## 🏗️ How it Works (The "Restaurant Analogy")
To keep the code clean and professional, I used a **Layered Architecture**. Think of it like a high-end restaurant:

1.  **The Waiter (Handler Layer)**: 
    *   Takes your order (HTTP Request).
    *   Checks if you provided a valid name and date (Validation).
    *   If everything looks good, the waiter takes the order to the kitchen.
2.  **The Chef (Service Layer)**: 
    *   This is the "Brain" of the operation. 
    *   The Chef performs the complex math to calculate the age in Years, Months, and Days. 
    *   The Chef doesn't care how the waiter got the order; they just focus on the logic.
3.  **The Pantry (Repository Layer)**: 
    *   The Chef asks for ingredients (Data).
    *   The Pantry talks to the "Fridge" (Database) to fetch or store information safely.
4.  **The Fridge (PostgreSQL Database)**: 
    *   Where the data lives permanently. Even if the restaurant closes (server restarts), the data stays fresh.

---

## 🛠️ The Tech Stack
*   **Go (Golang)**: Chosen for its incredible speed and reliability.
*   **GoFiber**: A modern web framework that makes the API fast and responsive.
*   **PostgreSQL**: An industry-standard database for secure data storage.
*   **SQLC**: A tool that ensures our database commands (SQL) are 100% accurate and safe.
*   **Uber Zap**: A professional logging system that tracks every action the API takes.

---

## 🚀 Getting Started

### 1. Prerequisites
*   **Go** installed on your machine.
*   **PostgreSQL** database running.

### 2. Setup
1.  **Clone the project** to your local machine.
2.  **Configure the Environment**:
    Create a `.env` file and add your database details:
    ```env
    DATABASE_URL=postgres://username:password@localhost:5432/user_db?sslmode=disable
    ```
3.  **Prepare the Database**:
    Run the SQL script found in `db/migrations/000001_create_users_table.up.sql` to create the `users` table.

### 3. Run the App
```bash
go run cmd/server/main.go
```
Now, open your browser and visit: **[http://localhost:3000](http://localhost:3000)**

---

## 🧪 Detailed Age Logic: Why it's Special
Calculating age is harder than it looks! February has 28 days, while March has 31. A simple subtraction of years would be wrong.

**My Logic Handles:**
*   **Month Rollovers**: If today is June 10th and you were born June 15th, you aren't a year older yet! My code checks the day and month carefully.
*   **Varying Month Lengths**: The code automatically calculates how many days were in the *previous* month to give an exact "Days" count.
*   **Accuracy**: It’s accurate down to the single day.

---

## 🛡️ Security & Reliability
*   **SQL Injection Protection**: I use "Placeholders" ($1, $2) so hackers can't send dangerous commands to the database.
*   **Request Tracking**: Every request gets a unique "ID tag" (Request ID) so we can track exactly what happened if an error occurs.
*   **Validation**: The system rejects empty names or impossible dates automatically.

---

## 👨‍💻 Developer Note
This project demonstrates the ability to build a **Full-Stack experience**—from raw SQL database queries to a polished Go backend, and finally a clean user interface. It follows professional software engineering standards used in top tech companies.
