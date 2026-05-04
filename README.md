# 🚀 ERP Microservices (Go + Fiber + PostgreSQL)

## 📌 Overview

ระบบ ERP (Enterprise Resource Planning) ที่ออกแบบด้วยแนวคิด **Microservices Architecture**
พัฒนาโดยใช้ **Golang (Fiber)** และ **PostgreSQL** รองรับการทำงานด้าน:

* 👤 User / Staff Management
* 📦 Inventory / Stock Management
* 🧾 Sales / Order / Invoice

---

## 🧱 Architecture

```text
Client
  ↓
Nginx (Reverse Proxy)
  ↓
-----------------------------
| setting | stock | sell |
-----------------------------
        ↓
     PostgreSQL
```

### 🔹 Services

| Service          | Description                 | Port |
| ---------------- | --------------------------- | ---- |
| `erp-setting-cs` | User / Role / Master Data   | 8001 |
| `erp-stock-cs`   | Product / Warehouse / Stock | 8002 |
| `erp-sell-cs`    | Order / Invoice / Payment   | 8003 |

---

## ⚙️ Tech Stack

* **Backend:** Golang (Fiber)
* **Database:** PostgreSQL
* **Container:** Docker / Docker Compose
* **Architecture:** Clean Architecture
* **API:** RESTful

---

## 📦 Project Structure (Example)

```
erp-setting-cs/
├── internal/
│   ├── domain/
│   ├── usecase/
│   ├── repository/
│   └── handler/
├── main.go
└── go.mod
```

---

## 🔗 Service Communication

บริการแต่ละตัวสื่อสารกันผ่าน HTTP (REST API)

### Example:

```bash
sell → stock → check stock
sell → stock → deduct stock
```

```go
http.Post("http://stock-service:8000/api/check-stock", ...)
```

---

## 🐳 Run with Docker

### 1. Clone Project

```bash
git clone https://github.com/your-username/erp-project.git
cd erp-project
```

### 2. Run All Services

```bash
docker-compose up --build
```

---

## 🌐 API Endpoints (Example)

### Setting Service

```
GET    /api/users
POST   /api/login
```

### Stock Service

```
GET    /api/products
POST   /api/stock/in
POST   /api/stock/out
```

### Sell Service

```
POST   /api/orders
GET    /api/orders
POST   /api/invoice
```

---

## 🧠 Business Flow

### 🧾 Create Order

1. User creates order (sell-service)
2. sell-service → check stock (stock-service)
3. stock-service → validate stock
4. sell-service → confirm order
5. stock-service → deduct stock

---

## 🔐 Authentication

* JWT-based authentication
* Access Token + Refresh Token (planned)

---

## 📊 Database Design

* Single PostgreSQL instance
* Separated by schema:

  * `setting.*`
  * `stock.*`
  * `sell.*`

---

## 🚀 Deployment

* Dockerized services
* Reverse Proxy (Nginx)
* Deploy on VPS
* Domain + HTTPS (production)

---

## 👨‍💻 Demo Account

```
username: admin
password: 1234
```

---

## 📌 Future Improvements

* [ ] Add CI/CD (GitHub Actions)
* [ ] Add Logging (Zap)
* [ ] Add Unit Test
* [ ] Add API Gateway
* [ ] Add Distributed Transaction (Saga Pattern)

---

## 🧑‍💼 Author

**ARM (BreakDown-CS)**
Backend Developer (Golang)

---

## ⭐ Highlight

* Microservices Design
* Clean Architecture
* Real-world ERP Flow
* Docker-based Deployment

---
