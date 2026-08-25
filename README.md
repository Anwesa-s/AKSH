# AKSH — Production-Grade API Gateway

> A production-oriented API Gateway built from scratch in Go, designed to explore routing, reverse proxying, middleware, microservice communication, JWT authentication, and role-based authorization.

![Go](https://img.shields.io/badge/Go-1.26.5-00ADD8?logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-in%20development-orange)
![Architecture](https://img.shields.io/badge/architecture-API%20Gateway-blue)

---

## Overview

**AKSH** is a custom API Gateway built in Go with a focus on production-oriented backend and infrastructure engineering.

The gateway acts as a single entry point between clients and multiple backend services. It handles request processing, routing, authentication, authorization, logging, and reverse proxying before forwarding requests to the appropriate microservice.

The long-term goal is to evolve AKSH into a production-grade gateway with advanced capabilities such as rate limiting, caching, load balancing, circuit breaking, observability, service discovery, and cloud deployment.

---

## Architecture

```text
                         ┌─────────────────────┐
                         │       Client        │
                         │ Browser / curl / API│
                         └──────────┬──────────┘
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │        AKSH         │
                         │   API Gateway :8080 │
                         └──────────┬──────────┘
                                    │
                              Middleware
                                    │
                     ┌──────────────┴──────────────┐
                     │                             │
                     ▼                             ▼
               Logger Middleware             JWT Authentication
                                                   │
                                            Claims / Role
                                                   │
                                                   ▼
                                            Authorization
                                                   │
                                                   ▼
                                               Router
                                                   │
                         ┌─────────────────────────┼─────────────────────┐
                         │                         │                     │
                         ▼                         ▼                     ▼
                 User Service :9001       Order Service :9002    Payment Service
```

### Request Flow

```text
Client
  ↓
AKSH :8080
  ↓
Logger
  ↓
JWT Authentication
  ↓
JWT Claims → Request Context
  ↓
Role-Based Authorization
  ↓
Route Matching
  ↓
Reverse Proxy
  ↓
Backend Microservice
```

---

## Features

### Configuration-Driven Routing

AKSH uses configuration-based route registration to map incoming API paths to backend services.

```text
/users   → User Service
/orders  → Order Service
```

This keeps service locations separate from the gateway's core routing logic.

### Reverse Proxy

Requests received by AKSH are forwarded to the appropriate backend service through a reverse proxy.

```text
Client
   ↓
AKSH :8080
   ↓
Reverse Proxy
   ↓
Backend Service
```

This allows clients to interact with a single gateway instead of communicating directly with individual microservices.

### Configuration Management

Backend services and routes are defined using YAML configuration.

Example:

```yaml
services:
  user-service:
    url: http://localhost:9001

routes:
  - path: /users
    service: user-service
```

### Middleware Pipeline

Current middleware includes:

- Request logging
- JWT authentication
- Role-based authorization

The middleware architecture is designed to support additional gateway capabilities in the future.

### JWT Authentication

AKSH provides a login endpoint that generates signed JSON Web Tokens after successful authentication.

JWT claims currently include:

```json
{
  "user_id": 101,
  "role": "user"
}
```

Token validation includes:

- Signature verification
- Signing-method validation
- Token validity checks
- Expiration validation
- Bearer token extraction

Protected endpoints require:

```text
Authorization: Bearer <JWT>
```

### JWT Claims & Request Context

After successful authentication, AKSH extracts information from the JWT and stores it in the Go request context.

Currently available claims include:

```text
user_id
role
```

This allows downstream middleware and handlers to access authenticated user information without decoding the JWT again.

### Role-Based Authorization

AKSH separates authentication from authorization.

**Authentication:** Who are you?

**Authorization:** What are you allowed to access?

Example:

```text
No JWT
   ↓
401 Unauthorized

Valid user JWT
   ↓
Admin route
   ↓
403 Forbidden

Valid admin JWT
   ↓
Admin route
   ↓
200 OK
```

---

## Security Testing

The authentication and authorization layer has been tested against:

| Scenario | Expected Result |
|---|---:|
| Missing JWT | `401 Unauthorized` |
| Invalid JWT | `401 Unauthorized` |
| Expired JWT | `401 Unauthorized` |
| Valid user JWT → protected route | `200 OK` |
| Valid user JWT → admin route | `403 Forbidden` |
| Valid admin JWT → admin route | `200 OK` |

Run the Go test suite with:

```bash
go test ./...
```

---

## Project Structure

```text
AKSH/
│
├── backend/
│   ├── user-service/
│   │   └── main.go
│   ├── order-service/
│   │   └── main.go
│   └── payment-service/
│       └── main.go
│
├── cmd/
│   └── server/
│       └── main.go
│
├── config/
│   └── config.yaml
│
├── internal/
│   ├── auth/
│   │   └── jwt.go
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   └── auth.go
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── logger.go
│   │   └── role.go
│   ├── models/
│   ├── proxy/
│   │   └── proxy.go
│   └── router/
│
├── go.mod
├── go.sum
└── README.md
```

---

## Getting Started

### Prerequisites

- Go 1.26+
- Git

### Clone the repository

```bash
git clone https://github.com/Anwesa-s/AKSH.git
cd AKSH
```

### Start the gateway

```bash
go run ./cmd/server
```

The gateway runs on:

```text
http://localhost:8080
```

### Start the backend services

Run the required backend services according to the configuration.

Current development services include:

```text
AKSH Gateway      → :8080
User Service      → :9001
Order Service     → :9002
Payment Service   → configured backend service
```

### Run tests

```bash
go test ./...
```

---

## Authentication Examples

### Login as a user

```powershell
curl.exe -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username":"anwesa","password":"password123"}'
```

### Login as an administrator

```powershell
curl.exe -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username":"admin","password":"admin123"}'
```

The response contains:

```json
{
  "token": "YOUR_JWT"
}
```

### Access a protected route

```powershell
curl.exe -i -H "Authorization: Bearer YOUR_JWT" http://localhost:8080/users
```

### Access an administrator route

```powershell
curl.exe -i -H "Authorization: Bearer YOUR_ADMIN_JWT" http://localhost:8080/admin
```

> **Note:** The credentials shown above are development/test credentials and are not intended for production use.

---

## Configuration

AKSH uses:

```text
config/config.yaml
```

to define backend services and their routes.

Conceptually:

```text
Incoming Request
       ↓
Route Lookup
       ↓
Configured Service
       ↓
Reverse Proxy
       ↓
Backend Service
```

This configuration-driven approach makes it easier to add or modify services without changing the core gateway architecture.

---

## Roadmap

### Core Gateway

- [x] HTTP server
- [x] Request routing
- [x] Configuration-driven routes
- [x] Reverse proxy
- [x] Multiple backend services
- [x] Middleware pipeline

### Security

- [x] JWT authentication
- [x] JWT validation
- [x] Token expiration handling
- [x] JWT claims
- [x] Request context
- [x] Role-based authorization
- [x] Authentication and authorization testing

### Reliability & Performance

- [ ] Rate limiting
- [ ] Redis integration
- [ ] Health checks
- [ ] Load balancing
- [ ] Retry engine
- [ ] Circuit breaker
- [ ] Response caching
- [ ] Compression
- [ ] CORS

### Advanced Gateway Capabilities

- [ ] Service discovery
- [ ] Dynamic configuration / hot reload
- [ ] Plugin system
- [ ] Canary routing
- [ ] Blue-green deployment
- [ ] A/B routing

### Observability & Deployment

- [ ] Metrics
- [ ] Prometheus integration
- [ ] Grafana dashboards
- [ ] Dockerization
- [ ] CI/CD pipeline
- [ ] Cloud deployment
- [ ] Production hardening

---

## Engineering Goals

AKSH is being developed to explore the engineering concepts behind real-world API gateways and distributed backend systems.

Key areas include:

- HTTP request lifecycle
- API gateway architecture
- Reverse proxies
- Routing systems
- Middleware pipelines
- Microservice communication
- Authentication
- Authorization
- Rate limiting
- Load balancing
- Reliability engineering
- Caching
- Observability
- Service discovery
- Production deployment

---

## Development Philosophy

The project is developed incrementally with an emphasis on understanding the underlying systems rather than simply integrating existing gateway software.

Each capability follows a practical engineering cycle:

```text
Understand
    ↓
Design
    ↓
Implement
    ↓
Test
    ↓
Refactor
    ↓
Document
```

The goal is to continuously evolve AKSH from a functional gateway into a production-oriented infrastructure component.

---

## Future Vision

The long-term architecture is intended to evolve toward:

```text
                         ┌─────────────────────┐
                         │       Clients       │
                         └──────────┬──────────┘
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │        AKSH         │
                         │    API Gateway      │
                         └──────────┬──────────┘
                                    │
              ┌─────────────────────┼─────────────────────┐
              │                     │                     │
              ▼                     ▼                     ▼
        Authentication        Rate Limiting          Routing
              │                     │                     │
              └─────────────────────┼─────────────────────┘
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │ Reliability Layer  │
                         │ Retry / Circuit     │
                         │ Breaker / Timeout   │
                         └──────────┬──────────┘
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │  Load Balancing &   │
                         │ Service Discovery   │
                         └──────────┬──────────┘
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │ Microservices       │
                         │ User / Order / etc. │
                         └─────────────────────┘
```

The eventual goal is a gateway capable of handling security, traffic management, reliability, observability, and deployment concerns in a unified architecture.

---

## Author

**Anwesa Sahu**

Built as a hands-on backend and infrastructure engineering project using Go.

---

## Why AKSH?

AKSH is built to understand what happens inside an API Gateway rather than treating the gateway as a black box.

The project focuses on progressively implementing the fundamental systems that make modern gateways reliable, secure, scalable, and observable.
