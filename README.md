# SecureGo Login System

### Matthew Reiley

**Date:** December 17, 2023

---

## 📋 Overview

**SecureGo Login System** is a robust and efficient authentication solution developed using Go (Golang). This project demonstrates how to build a secure, scalable, and user-friendly login system that protects user credentials from unauthorized access. It showcases advanced techniques such as JWT-based authentication, password hashing with bcrypt, role-based access control (RBAC), and audit logging.

---

## 🚩 Problem Statement

In an age of pervasive digital interactions, user credentials are under constant threat. Traditional login mechanisms often fail to adapt to modern security requirements, leaving accounts vulnerable. SecureGo tackles these challenges by providing:

* Strong encryption and hashing to safeguard passwords.
* JWT token-based session handling.
* RBAC to control access to sensitive routes.
* Audit trails to monitor and respond to suspicious activity.

---

## 🎯 Goals

✅ **Enhanced Security:** Leverages Go's capabilities to implement secure authentication and encryption techniques.
✅ **User-Friendly Experience:** Balances security with a seamless login and registration flow.
✅ **Scalable & Performant:** Built to handle growth without sacrificing performance.
✅ **Adaptable & Integratable:** Easily integrates into diverse tech stacks.
✅ **Auditable:** Includes logging and monitoring for accountability.

---

## 🛠️ Key Technologies

* **Go (Golang):** High-performance backend language.
* **MySQL:** Database for persistent user and role storage.
* **bcrypt:** For secure password hashing.
* **JWT:** Stateless authentication tokens.
* **Gorilla Mux:** HTTP router.
* **CORS & Middleware:** For secure HTTP access control.

---

## 📝 Project Structure & Components

### Core Files

* **`main.go`** — Entry point. Initializes MySQL, configures routes, middleware, and starts the server.

---

### Database Layer

* **`db/database.go`** — Initializes and connects to MySQL database.
* **`models/user.go`** — Defines the `User` struct with username, hashed password, and role.
* **`repositories/user_repository.go`** — Provides CRUD operations for user data.

---

### Handlers

* **`handlers/homeHandler.go`** — Serves the home route (`/`), confirming server status.
* **`handlers/loginHandler.go` & `handlers/adminLoginHandler.go`** — Authenticate users and admins. Generate JWT tokens with claims.
* **`handlers/registerHandler.go`** — Registers new users, hashes their passwords, and issues a JWT token.

---

### Helpers

* **`helpers/passwordHelper.go`** — Checks hashed vs. plain-text passwords.
* **`helpers/tokenHelper.go`** — Creates, validates, and extracts JWT tokens from requests.

---

### Middleware

* **`middleware/authmiddleware.go`** — Ensures routes are only accessible to authenticated users.
* **`middleware/rbacmiddleware.go`** — Restricts routes to specific user roles.

---

## 🔄 Features

✨ **JWT Authentication:** Stateless and secure user sessions.
✨ **Password Hashing:** bcrypt used to protect stored passwords.
✨ **Role-Based Access Control (RBAC):** Grants different access levels to admin and regular users.
✨ **CORS Handling:** Safe cross-origin resource sharing.
✨ **Audit Trails:** Logs all login attempts and access for monitoring.

---

## 🚀 How to Run

### Prerequisites

* Go (>=1.18)
* MySQL

### Steps

1️⃣ Clone the repository:

```bash
git clone <repo-url>
cd securego-login-system
```

2️⃣ Configure your MySQL connection in `db/database.go`.

3️⃣ Run the server:

```bash
go run main.go
```

4️⃣ The server runs on `http://localhost:8080`.

---

## 🔗 API Endpoints

| Endpoint       | Method | Description                           |
| -------------- | ------ | ------------------------------------- |
| `/`            | GET    | Home page                             |
| `/register`    | POST   | Register a new user                   |
| `/login`       | POST   | User login                            |
| `/admin-login` | POST   | Admin login                           |
| `/profile`     | GET    | Protected route (authenticated users) |
| `/admin`       | GET    | Admin-only route                      |

---

## 📄 Summary

SecureGo Login System is a thoughtfully designed authentication system that balances modern security practices with usability and scalability. With JWT, bcrypt, RBAC, and solid middleware practices, it exemplifies industry standards for building secure login systems. Whether for small startups or large-scale enterprise apps, SecureGo ensures user data remains safe while providing a seamless login experience.

---

**Author:** Matthew Reiley
**Date:** December 17, 2023
