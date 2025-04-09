# GoTask API 🧠

A fully-featured RESTful API for collaborative task and project management, built with Go, PostgreSQL, Docker, and JWT Authentication.

## 🔧 Tech Stack

- **Golang** (Gin)
- **PostgreSQL**
- **JWT** for auth
- **Bcrypt** for password hashing
- **Docker** + **Docker Compose**
- **GORM**
- **Makefile**
- **GitHub Actions** for CI
- **Swagger / Postman** for API docs

## 🚀 Features

- ✅ User registration and login with secure password hashing
- ✅ JWT authentication with refresh token support
- ✅ Create and manage projects with team collaboration
- ✅ Task CRUD with priority, status, and due dates
- ✅ Comment system on tasks
- ✅ Role-based permissions (owner vs member)
- ✅ Activity logging for auditing actions
- ✅ Pagination for tasks and comments
- ✅ Dockerized development environment
- ✅ CI pipeline with GitHub Actions
- ✅ Swagger/OpenAPI documentation
- ✅ Rate limiting to prevent abuse

<!-- ## 📁 Project Structure 

go-task-api/ 
├── cmd/ # Entry point (main.go) 
├── internal/ # Domain logic 
├── pkg/ # Shared utilities 
├── config/ # Configuration files 
├── scripts/ # DB migrations / seeds 
├── test/ # Test files 
├── Dockerfile 
├── docker-compose.yml 
├── Makefile 
├── .env.example 
└── README.md
-->


## 📦 Getting Started

1. **Clone the repo:**
   ```bash
   git clone https://github.com/yourusername/go-task-api.git
   cd go-task-api
   ```
2. **Set up environment:**
   ```bash
cp .env.example .env
   ```

3. **Run with Docker:**
   ```bash
make docker-up
   ```

4. **Access API:**
   ```bash
http://localhost:8080/api/v1/
   ```
5. **Run tests:**
   ```bash
make test
   ```

📄 API Documentation
Swagger UI: http://localhost:8080/docs

<!-- Postman Collection: Download -->

🤝 Contributing
PRs are welcome! If you find bugs or want to request features, feel free to open an issue.
