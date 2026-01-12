# Event Booking REST API - Project Plan

This document outlines the plan for building a complete Event Booking REST API using Go. This project integrates key concepts including HTTP handling with Gin, database interactions with SQLite, and JWT-based authentication.

## 1. Project Overview
We will build a backend API that allows users to:
- Browse available events.
- Sign up and log in.
- Create, update, and delete their own events.
- Register for an event and cancel registrations.

## 2. Technology Stack
- **Language**: Go (Golang)
- **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin) (High-performance HTTP web framework)
- **Database**: SQLite (via `modernc.org/sqlite` or `github.com/mattn/go-sqlite3`)
- **Authentication**: JWT (JSON Web Tokens)
- **Security**: `bcrypt` for password hashing

## 3. Architecture & Folder Structure
We will adopt a modular structure to keep code organized:
```
project-root/
├── db/             # Database initialization and connection logic
├── middlewares/    # Custom middlewares (e.g., Auth verification)
├── models/         # Data structures and database interaction methods
├── routes/         # HTTP request handlers and route registration
├── utils/          # Utility functions (Crypto, JWT generation/parsing)
├── api-test/       # .http files for testing API endpoints
└── main.go         # Application entry point
```

## 4. Implementation Steps

### Phase 1: Foundation & Basic API
1. **Project Setup**: Initialize Go module and install Gin.
2. **Basic Routes**: Create a simple GET `/events` and POST `/events` endpoint using in-memory storage first.
3. **Event Model**: Define the `Event` struct with fields: `ID`, `Name`, `Description`, `Location`, `DateTime`, `UserID`.

### Phase 2: Database Integration
1. **SQLite Setup**: Initialize the database connection in `db/db.go`.
2. **Tables**: Create `events` and `users` tables.
3. **Refactor Storage**: Update `Event` methods (`Save`, `GetAll`) to read/write from the SQLite database instead of memory.

### Phase 3: Full CRUD for Events
1. **Get Single Event**: Implement GET `/events/:id`.
2. **Update Event**: Implement PUT `/events/:id`.
3. **Delete Event**: Implement DELETE `/events/:id`.
4. **Refactoring**: Move routing logic into a dedicated `routes` package.

### Phase 4: User Authentication
1. **User Model**: Define `User` struct (`ID`, `Email`, `Password`).
2. **Password Hashing**: Use `bcrypt` to hash passwords on signup and compare on login.
3. **Auth Routes**:
   - POST `/signup`: Create a new user.
   - POST `/login`: Authenticate credentials.
4. **JWT Implementation**:
   - Create utilities to **generate** tokens upon login.
   - Create utilities to **verify** tokens for protected requests.

### Phase 5: Route Protection & Authorization
1. **Middleware**: Create an `Authenticate` middleware to validate JWTs in headers.
2. **Protect Routes**: Apply middleware to `create`, `update`, and `delete` event routes.
3. **Ownership Check**: Ensure a user can only update or delete events *they* created.

### Phase 6: Event Registrations
1. **Registrations Table**: Create a `registrations` table (linking `EventID` and `UserID`).
2. **Registration Routes**:
   - POST `/events/:id/register`: Authenticated users can register for an event.
   - DELETE `/events/:id/register`: Authenticated users can cancel registration.
3. **Logic**: Implement `Register` and `CancelRegistration` methods in models.

## 5. API Endpoints Reference

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

## 6. Next Steps
- Implement the basic server in `main.go`.
- Set up the database connection.
- Begin iteratively building routes as per Phase 1.