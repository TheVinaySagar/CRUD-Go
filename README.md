# Go Gin Web API

This is a modular API boilerplate using **Gin Framework** and **GORM**.

## 🚀 How to Run

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Start the server:
   ```bash
   go run cmd/main.go
   ```

3. Test API in **Postman**:
   - GET `/users`
   - POST `/users` with `{ "name": "Alice", "email": "alice@example.com" }`
