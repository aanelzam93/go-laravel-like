
# go-laravel-like

A Laravel-inspired backend framework built with **Golang** using **Gin** framework.  
Designed to provide a clean, modular, and scalable structure for large-scale backend development.

---

## ✨ Project Information

- **Project Name:** go-laravel-like
- **Author:** Anwar  
- **Email:** aan.elzam93@gmail.com
- **Language:** Golang
- **Framework:** Gin Gonic

---

## 🚀 Major Features

- Clean and modular structure inspired by Laravel & DDD
- Environment variables support via `.env` file
- Built-in **JWT Authentication** support
- Middleware support (Authentication, Rate Limiter, Recovery)
- Support **multiple database connections** (MySQL, PostgreSQL, MongoDB)
- Auto migration system
- Advanced background jobs with **Redis Queue Worker**
- Response standardization (JSON Format)
- Custom Validation Helper
- Multi-Environment Config Support (`.env.development`, `.env.production`)
- Graceful Shutdown Support
- Cluster Ready (stateless, Redis centralized)
- Swagger API Documentation
- Seeder & Factory for Dummy Data
- Health Check Endpoint
- Modular System per Feature/Domain
- Optional Redis, SMTP, Queue System Integration
- **Dockerfile & Docker Compose Support**

---

## 📦 Modular Folder Structure

```
go-laravel-like/
├── app/
│   ├── modules/
│   │   ├── auth/
│   │   ├── user/
│   ├── middlewares/
│   ├── helpers/
│   ├── queue/
│   ├── jobs/
├── config/
├── database/
│   ├── migrations/
│   ├── seeders/
├── routes/
├── docs/
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── .env
├── main.go
├── go.mod
├── README.md
```

---

## 📋 API Endpoints (example)

| Method | Endpoint          | Description |
|:-------|:------------------|:------------|
| POST   | `/api/register`    | Register new user |
| POST   | `/api/login`       | Login user |
| GET    | `/swagger/index.html` | Swagger Documentation |
| GET    | `/health` | Health Check (DB, Redis) |

---

## 🛠 Setup & Installation

### Local Development

1. Clone the project:

```bash
git clone https://github.com/your-repo/go-laravel-like.git
cd go-laravel-like
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create `.env` or `.env.development` file.

4. Generate Swagger Documentation:

```bash
swag init
```

5. Run the server:

```bash
go run main.go
```

---

### Docker Development

1. Ensure Docker and Docker Compose are installed.

2. Run the project using:

```bash
docker-compose up --build
```

✅ The services will start:
- App on `http://localhost:8000`
- MySQL on `localhost:3307`
- Redis on `localhost:6379`

3. Docker Compose will handle:
- Go app container
- MySQL database container
- Redis server container

---

## 🧩 How Modular Works

- Controller handles HTTP request.
- Controller calls Service.
- Service manipulates Model (DB).
- Controller returns JSON Response.

Each module (auth, user, etc.) is isolated in its own domain.

---

## 🛠 Future Improvements

- Add Role & Permission system (RBAC)
- CRUD Generator
- Real-time WebSocket Support
- API Versioning (v1, v2)
- Multi-language support (i18n)

---

## 📄 License

This project is open-source and free to use.  
Feel free to modify it for your needs.

---

> Created with ❤️ by **Anwar** | Contact: **aan.elzam93@gmail.com**
