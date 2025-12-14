User Management API
A production-ready RESTful API for managing users with automatic age calculation. Built with Go, Fiber, MySQL, and fully containerized with Docker.

Features
Full CRUD Operations: Create, Read, Update, and Delete users.

Automatic Age Calculation: Age is dynamically computed from the date of birth.

Production-Ready Setup:

Dockerized services with health checks and retry logic.

Database schema auto-initialization on startup.

Structured logging with Zap.

Connection pooling and graceful shutdown.

Modern Stack: Go 1.25, Fiber v2, MySQL 8.0.

🗂️ Project Structure
text
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

    
Quick Start
Get the API running with a single command. No manual database setup required.

1.Clone the repository
git clone https://github.com/<your-username>/aniyxProject.git
cd aniyxProject

2.Start the application
docker-compose up -d

This command will:Pull the MySQL 8.0 and Go Alpine images.

Start a MySQL container and automatically create the users table.

Build and start the Go API container, which will wait for the database to be ready.

Verify the services are running
docker-compose ps
Both services should show a status of Up or healthy.

3.API Endpoints
Health Check
curl http://localhost:3000/health
Response: {"status":"ok"}

User Management
Method	Endpoint	Description	Example Request Body
GET	/users	List all users	- http://localhost:3000/users
POST	/users	Create a new user	{"name": "Akash Pradhan", "dob": "1990-05-15"} - http://localhost:3000/users
GET	/users/:id	Get a specific user	-http://localhost:3000/users/1
PUT	/users/:id	Update a user	{"name": "Akash Pradhan", "dob": "1992-08-22"} - http://localhost:3000/users
DELETE	/users/:id	Delete a user	- http://localhost:3000/users/1
GET	/health:  application health	- http://localhost:3000/health




4.Docker Details
Services
api: Go Fiber application built from source. Exposes port 3000.

mysql: MySQL 8.0 database with pre-initialized schema from init.sql. Exposes port 3307 on the host.

Database Initialization
The init.sql script is automatically executed when the MySQL container starts for the first time. 
It:Creates the users table if it doesn't exist.
Inserts sample user data for immediate testing.

Health Checks
MySQL: The API container waits for MySQL to be ready before starting.

API: Includes a /health endpoint for container health monitoring.

5.Running Tests & Development
Rebuilding the Application
After making code changes, rebuild the containers:

docker-compose up -d --build


Viewing Logs

# Follow logs for all services
docker-compose logs -f

# View logs for a specific service
docker-compose logs -f api
docker-compose logs -f mysql
