# 🚀 Interview Cheat Sheet: Go User Management API

## 1. The Big Picture (Architecture)
**Pattern:** Layered Architecture (N-Tier)
*   **CMD (`main.go`)**: The Glue. Wires DB, Services, and Handlers together.
*   **Handler**: The "Waiter." Handles HTTP requests/responses and input validation.
*   **Service**: The "Chef." Contains the **Business Logic** (Detailed Age Calculation).
*   **Repository**: The "Pantry." The ONLY layer that talks to the Database.
*   **SQLC**: A tool that converts SQL into Go code. **Why?** It prevents typos and makes DB access type-safe.

## 2. The Tech Stack (The "Why")
*   **Go (Golang)**: Fast, typed, and great for concurrent backend systems.
*   **GoFiber**: The web framework. It’s inspired by Express (Node.js) and is extremely fast.
*   **PostgreSQL**: A reliable, industry-standard relational database.
*   **Uber Zap**: A structured logger. It logs in **JSON format**, making it easy for machines to read.
*   **Validator (v10)**: Uses struct tags (like `validate:"required"`) to ensure user data is clean before processing.

## 3. The "Killer" Logic: Detailed Age Calculation
*   **The Problem:** Months have different lengths (28, 30, 31 days). A simple subtraction of years isn't accurate.
*   **The Solution:** 
    1.  Subtract the birth year from the current year.
    2.  Check the month difference. If the current month is before the birth month, subtract 1 year and add 12 to months.
    3.  Check the day difference. If the current day is before the birth day, subtract 1 month and add the number of days from the *previous* month.
*   **Result:** Provides a precise string: `"X years, Y months, Z days."`

## 4. Key File Map (Where to show code)
*   **Logic:** `internal/service/user_service.go` -> Show the `CalculateAge` function.
*   **Routes:** `internal/routes/routes.go` -> Show how endpoints are mapped.
*   **Database:** `db/queries.sql` -> Show your raw SQL knowledge.
*   **Frontend:** `public/index.html` -> Show the `fetch()` calls to the API.

## 5. Potential Interview Questions & Answers
**Q: "What happens if a user enters a future date for DOB?"**
*   *A:* "The `validator` in the Handler layer will catch it. I used `datetime=2006-01-02` validation, and the Service layer would return an error if the date logic results in a negative age."

**Q: "Why use a connection pool (pgxpool)?"**
*   *A:* "Opening a new DB connection for every request is slow. `pgxpool` keeps a set of open connections ready to use, which drastically improves performance under load."

**Q: "How do you handle errors?"**
*   *A:* "I implemented a **Global Error Handler** in Fiber. Any error in the Service or Repo layer bubbles up to the Handler, which then returns a consistent JSON response like `{"error": "message"}`."

**Q: "What is the purpose of the Request ID middleware?"**
*   *A:* "Observability. It tags every request with a UUID. If there’s a bug, I can search the logs for that specific ID to see the entire lifecycle of that request."
