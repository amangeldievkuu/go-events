# Go REST Events API

A robust REST API server built with Go for managing events and attendees. The application provides user authentication, event management, and attendee tracking functionality.

## Features

- 🔐 **User Authentication**
  - JWT-based authentication
  - User registration and login
  - Secure password hashing

- 📅 **Event Management**
  - Create, read, update, and delete events
  - Event ownership and authorization
  - Event details including title, description, date, and location

- 👥 **Attendee Management**
  - Add/remove attendees to events
  - View event attendees
  - Track user event attendance

- 📚 **API Documentation**
  - Swagger/OpenAPI documentation
  - Interactive API testing interface

## Technology Stack

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [SQLite](https://www.sqlite.org/) - Database
- [JWT](https://github.com/golang-jwt/jwt) - Authentication
- [Swagger](https://github.com/swaggo/swag) - API documentation
- [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations

## Getting Started

### Prerequisites

- Go
- SQLite3

### Installation

1. Clone the repository:
```sh
git clone https://github.com/amangeldievkuu/go-events.git
cd go-events
```
2. Run the migrations
```sh
go run ./cmd/migrate up
```
3. Run the project
```sh
go run ./cmd/api
```