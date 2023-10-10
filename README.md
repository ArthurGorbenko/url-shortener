# Golang CRUD REST API

## Overview

This project is a simple url shortener developed using Golang, aiming to provide a hands-on experience with the language and its ecosystem. The API interacts with a PostgreSQL database, and the application is containerized using Docker for deployment on Fly.io.
Url to play around https://cold-night-186.fly.dev

## Prerequisites

- Golang v1.20 (or the latest version available)
- PostgreSQL
- Docker
- Fly CLI (for deployment to Fly.io)

## Local Development

### 1. Clone the Repository

```sh
git clone git@github.com:ArthurGorbenko/url-shortener.git
cd url-shortener
```

### 2. Setup PostgreSQL
Ensure PostgreSQL is installed and running. Create a database and user for the application:

```sql
CREATE DATABASE your_database_name;
CREATE USER your_user_name WITH ENCRYPTED PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_user_name;
```

### 3. Configure Environment Variables
Create a .env file in the root of the project directory and configure the database connection string and other necessary environment variables:

```
CONNECTION_STRING=host=localhost user=your_user_name password=your_password dbname=your_database_name port=5432 sslmode=disable TimeZone=America/Toronto
PORT=8080
```

### 4. Install Dependencies and Run the Application
```sh
go mod download
go run .
```

## Deployment to Fly.io
Ensure you are logged in to Fly.io using the CLI. Deploy the Docker image:
```sh
flyctl deploy
```
