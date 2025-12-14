# User Management API

A **production-ready RESTful API** for managing users with **automatic age calculation**. Built using **Go**, **Fiber**, **MySQL**, and fully **containerized with Docker**.

---

## 🚀 Features

* **Full CRUD Operations**
  Create, Read, Update, and Delete users.

* **Automatic Age Calculation**
  Age is dynamically calculated from the user's date of birth.

* **Production-Ready Setup**

  * Dockerized services with health checks and retry logic
  * Database schema auto-initialization on startup
  * Structured logging using **Zap**
  * Connection pooling and graceful shutdown

* **Modern Tech Stack**

  * Go **1.25**
  * Fiber **v2**
  * MySQL **8.0**

---

## 🗂️ Project Structure

```text
aniyxProject/
├── docker-compose.yml          # Multi-service Docker definition
├── Dockerfile                  # Go API build instructions
├── init.sql                    # Database schema & sample data
├── .env.example                # Environment variables template
├── cmd/
│   └── server/
│       └── main.go             # Application entry point
└── internal/                   # Core application logic
    ├── handler/                # HTTP request handlers
    ├── middleware/             # Fiber middleware (logging, etc.)
    ├── repository/             # Database operations layer
    ├── service/                # Business logic layer
    ├── models/                 # Data structures (request/response)
    └── logger/                 # Logging configuration
```

---

## ⚡ Quick Start

Get the API running with a **single command**. No manual database setup required.

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/<your-username>/aniyxProject.git
cd aniyxProject
```

### 2️⃣ Start the Application

```bash
docker-compose up -d
```

This command will:

* Pull MySQL **8.0** and Go Alpine images
* Start the MySQL container and automatically create the `users` table
* Build and start the Go API container (waits for DB readiness)

### 3️⃣ Verify Services

```bash
docker-compose ps
```

Both services should show a status of **Up** or **healthy**.

---

## 📡 API Endpoints

### 🔍 Health Check

```bash
GET /health
```

Example:

```bash
curl http://localhost:3000/health
```

Response:

```json
{"status":"ok"}
```

---

### 👤 User Management

| Method | Endpoint   | Description       | Example                                                        |
| ------ | ---------- | ----------------- | -------------------------------------------------------------- |
| GET    | /users     | List all users    | [http://localhost:3000/users](http://localhost:3000/users)     |
| POST   | /users     | Create a new user | `{ "name": "Akash Pradhan", "dob": "1990-05-15" }`             |
| GET    | /users/:id | Get a user by ID  | [http://localhost:3000/users/1](http://localhost:3000/users/1) |
| PUT    | /users/:id | Update a user     | `{ "name": "Akash Pradhan", "dob": "1992-08-22" }`             |
| DELETE | /users/:id | Delete a user     | [http://localhost:3000/users/1](http://localhost:3000/users/1) |

---

## 🐳 Docker Details

### Services

* **api**
  Go Fiber application built from source. Exposes port **3000**.

* **mysql**
  MySQL 8.0 database initialized using `init.sql`. Exposes port **3307** on the host.

---

### Database Initialization

The `init.sql` script runs automatically on first startup:

* Creates the `users` table if it does not exist
* Inserts sample user data for immediate testing

---

### Health Checks

* **MySQL**: API container waits until the database is ready
* **API**: `/health` endpoint available for monitoring

---

## 🧪 Development & Testing

### Rebuild After Code Changes

```bash
docker-compose up -d --build
```

### View Logs

```bash
# All services
docker-compose logs -f

# API logs only
docker-compose logs -f api

# MySQL logs only
docker-compose logs -f mysql
```

---

## 📌 Notes

* Ensure Docker and Docker Compose are installed
* Copy `.env.example` to `.env` and adjust values if needed
* Suitable for production-ready backend.

---


