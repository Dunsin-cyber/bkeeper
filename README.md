# Bkeeper Finance API

A financial management API server built with Go, Echo framework, and PostgreSQL.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [Database Setup](#database-setup)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Development](#development)

## Overview

Bkeeper is a financial management API that provides user authentication, profile management, and category management capabilities. Built with modern Go practices and the Echo web framework.

## Features

- User authentication (register, login, password reset)
- JWT-based authorization
- Profile management
- Category management
- Email notifications
- PostgreSQL database with GORM ORM
- RESTful API design

## Prerequisites

Before running this server, ensure you have the following installed:

- Go 1.25.3 or higher
- PostgreSQL database
- Git (for cloning the repository)

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Dunsin-cyber/bkeeper.git
   cd bkeeper
   ```

2. **Install dependencies:**
   ```bash
   make install
   # or
   go mod download
   ```

## Configuration

1. **Create environment file:**
   
   Copy the example environment file and configure it:
   ```bash
   cp .env.example .env
   ```

2. **Configure environment variables:**
   
   Edit the `.env` file with your settings:
   ```env
   # Server Configuration
   PORT=8000

   # Database Configuration
   DB_DATABASE=bkeeper
   DB_HOST=localhost
   DB_USERNAME=postgres
   DB_PASSWORD=mysecretpassword
   DB_PORT=5432

   # Application
   APP_NAME="Bkeeper Finance"
   JWT_SECRET_KEY="your-secret-key-here"

   # Mailer Configuration (for password reset emails)
   MAIL_MAILER=smtp
   MAIL_HOST=sandbox.smtp.mailtrap.io
   MAIL_PORT=2525
   MAIL_USERNAME=your_mailtrap_username
   MAIL_PASSWORD=your_mailtrap_password
   MAIL_ENCRYPTION=tls
   MAIL_SENDER="support@bkeeper.finance"
   ```

   **Important:** Replace placeholder values with your actual credentials.

## Running the Server

### Option 1: Using the run script
```bash
./run.sh
```

### Option 2: Using Make (Development mode with auto-reload)
```bash
make run-dev
```

### Option 3: Using Go directly
```bash
go run ./cmd/api
```

The server will start on `http://localhost:8000` (or the port specified in your `.env` file).

## Database Setup

1. **Create PostgreSQL database:**
   ```bash
   createdb bkeeper
   # or using psql
   psql -U postgres -c "CREATE DATABASE bkeeper;"
   ```

2. **Run migrations:**
   ```bash
   make migrate
   # or
   go run ./internal/database/migrate_up.go
   ```

   This will create the following tables:
   - `user_models` - User accounts
   - `app_token_models` - Password reset tokens
   - `category_models` - Financial categories

3. **Seed database (optional):**
   ```bash
   make seeder FILENAME=category
   ```

## API Endpoints

### Health Check
- `GET /` - Health check endpoint

### Authentication (Public)
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/forgot/password` - Request password reset
- `POST /api/v1/auth/reset/password` - Reset password with token

### Profile (Authenticated)
- `GET /api/v1/profile/authenticated/user` - Get authenticated user details
- `PATCH /api/v1/profile/change/password` - Change user password

### Categories (Authenticated)
- `GET /api/v1/categories/all` - List all categories
- `POST /api/v1/categories/store` - Create a new category
- `DELETE /api/v1/categories/delete/:id` - Delete a category

**Note:** All authenticated endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

## Project Structure

```
bkeeper/
├── cmd/
│   └── api/
│       ├── handlers/          # Request handlers
│       ├── middlewares/       # Custom middleware
│       ├── requests/          # Request validation
│       ├── services/          # Business logic
│       ├── main.go           # Application entry point
│       └── routes.go         # Route definitions
├── common/                   # Shared utilities (JWT, database)
├── internal/
│   ├── app_errors/          # Error handling
│   ├── database/            # Migrations and seeders
│   ├── mailer/              # Email service
│   └── models/              # Database models
├── .env.example             # Example environment configuration
├── Makefile                 # Build and run commands
├── go.mod                   # Go dependencies
└── README.md               # This file
```

## Development

### Available Make Commands

```bash
make help          # Show available commands
make install       # Install dependencies
make run-dev       # Run in development mode with auto-reload
make migrate       # Run database migrations
make migrate_fresh # Drop and recreate all tables
make seeder FILENAME=<name>  # Run specific seeder
make test          # Run tests
```

### Testing

Run the test suite:
```bash
make test
# or
go test ./tests/... -v
```

### Code Structure

- **Handlers**: Process HTTP requests and return responses
- **Middlewares**: Authentication, logging, and request processing
- **Models**: Database schema definitions using GORM
- **Services**: Business logic layer
- **Common**: Shared utilities like JWT and database connection

## Troubleshooting

### Common Issues

1. **Database connection failed:**
   - Verify PostgreSQL is running
   - Check database credentials in `.env`
   - Ensure database exists

2. **Port already in use:**
   - Change the `PORT` value in `.env`
   - Or stop the process using the port

3. **JWT errors:**
   - Ensure `JWT_SECRET_KEY` is set in `.env`
   - Check token format in Authorization header

4. **Email sending fails:**
   - Verify MAIL_* credentials in `.env`
   - For development, use a service like Mailtrap

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is part of the Bkeeper Finance application.

## Support

For issues and questions, please open an issue on the GitHub repository.
