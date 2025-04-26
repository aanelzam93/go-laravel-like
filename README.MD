# go-laravel-like

A Laravel-inspired backend framework built with **Golang** using the **Gin** framework.  
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
- Support for **multiple database connections**:
  - MySQL
  - PostgreSQL
  - MongoDB
- Flexible database activation (choose one or more with `DB_DRIVERS`)
- Auto migration system (currently for `users` table)
- Standard JSON API response formatting (success, error)
- Ready for scalable production API projects

---

## ⚙️ Database Configuration

Set database connections via the `.env` file:

```env
# Database Drivers (options: mysql, pgsql, mongo, or combinations like mysql+mongo)
DB_DRIVERS=mysql

# MySQL Configuration
DB_MYSQL_HOST=127.0.0.1
DB_MYSQL_PORT=3306
DB_MYSQL_DATABASE=go-laravel-like
DB_MYSQL_USERNAME=root
DB_MYSQL_PASSWORD=

# PostgreSQL Configuration
DB_PGSQL_HOST=127.0.0.1
DB_PGSQL_PORT=5432
DB_PGSQL_DATABASE=laravel_like
DB_PGSQL_USERNAME=postgres
DB_PGSQL_PASSWORD=

# MongoDB Configuration
DB_MONGO_URI=mongodb://localhost:27017
DB_MONGO_DATABASE=laravel_like
```

Only one database is required to be active, but multiple can be used simultaneously if needed.

---

## 📦 Folder Structure

```bash
go-laravel-like/
├── app/
│   ├── controllers/
│   ├── middlewares/
│   ├── models/
│   ├── helpers/
├── config/
├── migrations/
├── routes/
├── .env
├── main.go
└── go.mod
```

---

## 🛠️ Installation

1. **Clone the project:**

```bash
git clone https://github.com/your-repo/go-laravel-like.git
cd go-laravel-like
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Configure `.env` file** according to your database setup.

4. **Run the server:**

```bash
go run main.go
```

✅ The server will start at `http://localhost:8000` (or according to your `.env`).

---

## 📋 API Endpoints

| Method | Endpoint       | Description                          |
| :----- | :-------------- | :----------------------------------- |
| POST   | `/api/register` | Register a new user                  |
| POST   | `/api/login`    | Authenticate and retrieve JWT token  |
| GET    | `/api/profile`  | Retrieve authenticated user profile (Protected, requires JWT token) |

---

## 🧑‍💻 Future Improvements

- Role & Permission System
- CRUD Generator
- Queue & Jobs system (similar to Laravel Queues)
- Caching (Redis integration)
- WebSocket support for real-time applications

---

## 📄 License

This project is open-source and free to use.  
Feel free to modify it to suit your needs.

---

> Created with ❤️ by **Anwar** | Contact: **aan.elzam93@gmail.com**