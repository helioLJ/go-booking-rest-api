# Event Booking REST API

A comprehensive Event Booking REST API built with Go, featuring JWT authentication, SQLite database, and full CRUD operations for events and user registrations.

## Features

- User authentication with JWT tokens
- Password hashing using bcrypt
- Full CRUD operations for events
- Event registration system
- SQLite database integration
- Route protection and authorization
- Ownership-based access control

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: Gin Gonic
- **Database**: SQLite3
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt

## Project Structure

```
project-root/
├── api-test/          # HTTP test files for API endpoints
├── db/                # Database initialization and connection
├── middlewares/       # Custom middlewares (Authentication)
├── models/            # Data models and database methods
├── routes/            # HTTP handlers and route registration
├── utils/             # Utility functions (JWT, password hashing)
└── main.go            # Application entry point
```

## API Endpoints

| Method | Endpoint | Auth Required | Description |
| :--- | :--- | :--- | :--- |
| **GET** | `/events` | No | List all events |
| **GET** | `/events/:id` | No | Get a single event by ID |
| **POST** | `/signup` | No | Create a new user account |
| **POST** | `/login` | No | Login and retrieve JWT |
| **POST** | `/events` | **Yes** | Create a new event |
| **PUT** | `/events/:id` | **Yes** | Update an event (Creator only) |
| **DELETE**| `/events/:id` | **Yes** | Delete an event (Creator only) |
| **POST** | `/events/:id/register`| **Yes** | Register for an event |
| **DELETE**| `/events/:id/register`| **Yes** | Cancel registration |

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/helioLJ/go-booking-rest-api.git
cd go-booking-rest-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run .
```

The server will start on `http://localhost:8080`

### Building

To build the application:
```bash
go build -o booking-api
./booking-api
```

## Usage Examples

### 1. Register a new user

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }'
```

### 2. Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }'
```

Response will include a JWT token:
```json
{
  "message": "Login successful!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Create an event (authenticated)

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: YOUR_JWT_TOKEN" \
  -d '{
    "name": "Go Conference 2024",
    "description": "Annual Go programming conference",
    "location": "San Francisco, CA",
    "dateTime": "2024-09-15T09:00:00Z"
  }'
```

### 4. Get all events

```bash
curl http://localhost:8080/events
```

### 5. Register for an event (authenticated)

```bash
curl -X POST http://localhost:8080/events/1/register \
  -H "Authorization: YOUR_JWT_TOKEN"
```

## Testing

HTTP test files are provided in the `api-test/` directory. You can use VSCode with the REST Client extension or any HTTP client to test the endpoints.

## Database Schema

### Users Table
- `id` (INTEGER, PRIMARY KEY)
- `email` (TEXT, UNIQUE)
- `password` (TEXT, hashed)

### Events Table
- `id` (INTEGER, PRIMARY KEY)
- `name` (TEXT)
- `description` (TEXT)
- `location` (TEXT)
- `dateTime` (DATETIME)
- `user_id` (INTEGER, FOREIGN KEY)

### Registrations Table
- `id` (INTEGER, PRIMARY KEY)
- `event_id` (INTEGER, FOREIGN KEY)
- `user_id` (INTEGER, FOREIGN KEY)

## Development Phases

This project was built in 6 phases:

1. **Phase 1**: Basic API setup with Gin framework
2. **Phase 2**: SQLite database integration
3. **Phase 3**: Full CRUD operations for events
4. **Phase 4**: User authentication with JWT
5. **Phase 5**: Route protection and authorization
6. **Phase 6**: Event registrations

## Security

- Passwords are hashed using bcrypt with a cost of 14
- JWT tokens expire after 2 hours
- Protected routes require valid JWT tokens
- Event modifications restricted to creators only

## License

MIT License

## Author

Built as part of "The Complete Guide to Go" course by Maximilian Schwarzmüller
