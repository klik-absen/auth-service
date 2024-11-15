# Klik Absen - Auth Service

## Features

- Authenticate: Authenticate users and generate token for session management.
- Password Security: Hash and store password securely with SHA256.

## Prerequisites

- Go 1.23.2
- PostgreSQL (or compatible SQL database)

## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/klik-absen/auth-service.git
   ```

2. Install Dependencies
   ```bash
   go mod tidy
   ```

3. Environment Variables: Configure a .env file in the project root:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=12345678
   DB_NAME=postgres
   ```

## Running the Service

1. Run Locally:
   ```bash
   go run .\cmd\auth-service\main.go
   ```

2. Build Executable:
   ```bash
   go build .\cmd\auth-service\
   .\auth-service
   ```

## API Endpoints

Endpoint | Method | Description
---|---|---
/api/v1/auth | POST | Login and receive a token

## License

This project is licensed under the MIT License.