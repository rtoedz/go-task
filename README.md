API Endpoints

- Base URL lokal: http://localhost:{PORT} (default 8080)
- Base URL udag deploy : go-task-production.up.railway.app

Products
- GET /api/product
- POST /api/product
  - Body: { "name": string, "price": int, "stock": int }
- GET /api/product/{id}
- PUT /api/product/{id}
  - Body: { "name": string, "price": int, "stock": int }
- DELETE /api/product/{id}

Categories
- GET /api/category
- POST /api/category
  - Body: { "name": string }
- GET /api/category/{id}
- PUT /api/category/{id}
  - Body: { "name": string }
- DELETE /api/category/{id}

