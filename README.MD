
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
- Standard JSON API response formatting (success, error)
- Ready for scalable production API projects
- Background jobs using **Goroutines** (e.g., sending emails asynchronously)

---

## ⚙️ Database Configuration

You can enable one or multiple databases via `.env`:

```env
# Database Drivers (choose: mysql, pgsql, mongo, or combinations like mysql+mongo)
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
```

Only one active database is required, but you can enable multiple at once if needed.

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
│   │   └── email_jobs.go
├── config/
├── migrations/
├── routes/
├── .env
├── main.go
└── go.mod
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

3. Setup your `.env` file.

4. Run the server:

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

---

## 🧑‍💻 Future Improvements

- Add Role & Permission System
- CRUD Generator
- Queue & Jobs system (like Laravel Queues)
- Caching (Redis integration)
- WebSocket support for real-time apps

---

## 📄 License

This project is open-source and free to use.  
Feel free to modify it for your needs.

---

> Created with ❤️ by **Anwar** | Contact: **aan.elzam93@gmail.com**
