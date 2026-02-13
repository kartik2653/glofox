# Glofox API

## Project Setup Requirements

- **Go version:** 1.18 or higher
- **Database:** PostgreSQL (configured in `internal/config/config.go`)
- **Go Modules:** Enabled (see `go.mod`)
- **Environment Variables:**
  - Use a `.env` file or set environment variables as needed for DB connection, etc.
- **Install dependencies:**
  - Run `go mod tidy` to install all Go dependencies

## Main Dependencies Used

- [github.com/gofiber/fiber/v2](https://github.com/gofiber/fiber) — Web framework
- [gorm.io/gorm](https://gorm.io/gorm) — ORM for Go
- [github.com/joho/godotenv](https://github.com/joho/godotenv) — Loads environment variables from `.env`

## How to Run

1. Clone the repository
2. Set up your PostgreSQL database and update config if needed
3. Run `go mod tidy` to install dependencies
4. Start the server:
   ```sh
   go run cmd/server/main.go
   ```
   The server will run on `localhost:3001` by default

---

# API ROUTES

## Health Check

### `GET /api/v1/ping`
- **Description:** Health check endpoint
- **Response:** `{ "status": "OK" }`

---

## Classes

### `POST /api/v1/classes/`
- **Description:** Create a new class
- **Body (JSON):**
  ```json
  {
    "class_name": "Yoga",
    "start_date": "2024-02-01", // YYYY-MM-DD
    "end_date": "2024-02-28",   // YYYY-MM-DD
    "capacity": 20,
    "instructor_id": 1
  }
  ```
- **Response:**
  - `201 Created` on success
  - `400 Bad Request` for validation errors

### `GET /api/v1/classes`
- **Description:** List classes with optional filters and pagination
- **Query Params:**
  - `class_name` (optional): filter by class name
  - `id` (optional): filter by class ID
  - `skip` (optional, default 0): offset for pagination
  - `limit` (optional, default 10): limit for pagination
- **Response:**
  ```json
  {
    "data": {
      "classes": [ ... ],
      "total": 100,
      "limit": 10,
      "skip": 0
    },
    "message": "Classes retrieved successfully"
  }
  ```

---

## Bookings

### `POST /api/v1/bookings/`
- **Description:** Create a new booking
- **Body (JSON):**
  ```json
  {
    "class_id": 1,
    "user_id": 2,
    "booking_date": "2024-02-10" // YYYY-MM-DD
  }
  ```
- **Response:**
  - `201 Created` on success
  - `400 Bad Request` for validation errors

### `GET /api/v1/bookings`
- **Description:** List bookings with optional filters and pagination
- **Query Params:**
  - `class_id` (optional): filter by class ID
  - `user_id` (optional): filter by user ID
  - `booking_date` (optional) : filter by booking date
  - `skip` (optional, default 0): offset for pagination
  - `limit` (optional, default 10): limit for pagination
  - `id` (optional) : filter by booking id
- **Response:**
  ```json
  {
    "data": {
      "bookings": [ ... ],
      "total": 100,
      "limit": 10,
      "skip": 0
    },
    "message": "Bookings retrieved successfully"
  }
  ```

---

## Error Responses
- All error responses are in the form:
  ```json
  { "error": "Error message here" }
  ```

---

## Notes
- All dates must be in `YYYY-MM-DD` format.
- All endpoints return JSON responses.
