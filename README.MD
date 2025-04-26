
# go-laravel-like

A Laravel-inspired backend framework built with **Golang** using **Gin** framework.  
Designed to provide a clean, modular, and scalable structure similar to Laravel.

---

## ✨ Project Information

- **Project Name:** go-laravel-like
- **Author:** Anwar  
- **Email:** aan.elzam93@gmail.com
- **Language:** Golang
- **Framework:** Gin Gonic

---

## 🚀 Features

- Clean and modular structure inspired by Laravel
- Environment variables support via `.env` file
- Built-in **JWT Authentication** support
- Middleware support (public & protected routes)
- Support **multiple database connections**:
  - MySQL
  - PostgreSQL
  - MongoDB
- Flexible database activation (`DB_DRIVERS` can choose one or more)
- Auto migration system (currently for `users` table)
- Background jobs using **Goroutines** (e.g., sending emails asynchronously)
- Standard JSON API response formatting (success, error)
- Ready for scalable production API projects
- **Swagger API Documentation** auto-generated
- **Redis Caching System** (optional)
- **Queue Worker** for background job processing (optional)
- **SMTP Email Service Integration** (optional)
- **Custom Validation** with error binding
- **Multi-Environment Support** (development, staging, production)
- **Rate Limiter Middleware** to prevent abuse
- **Seeder & Factory** for testing data

---

## ⚙️ Configuration in .env

```env
# Database Drivers
DB_DRIVERS=mysql

# MySQL
DB_MYSQL_HOST=127.0.0.1
DB_MYSQL_PORT=3306
DB_MYSQL_DATABASE=go-laravel-like
DB_MYSQL_USERNAME=root
DB_MYSQL_PASSWORD=

# PostgreSQL
DB_PGSQL_HOST=127.0.0.1
DB_PGSQL_PORT=5432
DB_PGSQL_DATABASE=laravel_like
DB_PGSQL_USERNAME=postgres
DB_PGSQL_PASSWORD=

# MongoDB
DB_MONGO_URI=mongodb://localhost:27017
DB_MONGO_DATABASE=laravel_like

# Redis
REDIS_ENABLED=false
REDIS_HOST=localhost:6379

# Mail SMTP
MAIL_ENABLED=false
MAIL_FROM=no-reply@example.com
MAIL_HOST=smtp.example.com
MAIL_PORT=587
MAIL_USERNAME=username@example.com
MAIL_PASSWORD=password

# JWT
JWT_SECRET=your_jwt_secret
JWT_EXPIRE=3600

# App Environment
APP_ENV=development
```

---

## 📦 Folder Structure

```
go-laravel-like/
├── app/
│   ├── controllers/
│   ├── middlewares/
│   ├── models/
│   ├── helpers/
│   ├── jobs/
│   ├── logs/
│   ├── queue/
├── config/
├── database/
│   ├── seeder/
├── docs/
├── migrations/
├── routes/
├── .env
├── main.go
├── go.mod
├── README.md
```

---

## 🛠 Installation

1. Clone the project:

```bash
git clone https://github.com/your-repo/go-laravel-like.git
cd go-laravel-like
```

2. Install dependencies:

```bash
go mod tidy
```

3. Setup your `.env` or `.env.{environment}` file.

4. Generate Swagger docs:

```bash
swag init
```

5. Run the server:

```bash
go run main.go
```

✅ The server will start on `http://localhost:8000` (or as set in `.env`).

---

## 📋 API Endpoints

| Method | Endpoint         | Description                |
|:-------|:------------------|:----------------------------|
| POST   | `/api/register`    | Register new user           |
| POST   | `/api/login`       | Login and retrieve JWT token |
| GET    | `/api/profile`     | Get logged-in user profile (Protected route, requires JWT token) |
| GET    | `/swagger/index.html` | Swagger API Documentation |

---

## 🛡 Middleware

- **JWT Authentication Middleware**
- **Global Error Recovery Middleware**
- **Rate Limiter Middleware** (per IP control)

---

## 🚀 Background Jobs & Queue

The project supports background jobs using Go's Goroutines and Redis queue.

- Push background jobs
- Process queue workers
- Optional: Redis disable fallback

---

## 🚀 Caching System (Redis)

Redis can be optionally enabled for caching data or handling queues.

Ensure to enable Redis by setting:
```env
REDIS_ENABLED=true
```

---

## 📧 Email Service (SMTP)

SMTP email sending is supported for sending emails like verification or notifications.

Ensure to enable Mail by setting:
```env
MAIL_ENABLED=true
```

---

## 🛠 Seeder & Factory

Seeder available for creating dummy data like users.

Example:

```bash
Seeder.SeedUsers()
```

Generate multiple dummy users quickly for development/testing.

---

## 🧑‍💻 Future Improvements

- Add Role & Permission System
- CRUD Generator
- Queue & Jobs system with retries and delay
- WebSocket support for real-time features
- Full unit and integration tests
- Dockerization support

---

## 📄 License

This project is open-source and free to use.  
Feel free to modify it for your needs.

---

> Created with ❤️ by **Anwar** | Contact: **aan.elzam93@gmail.com**
