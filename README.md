# ThreadPulse

A public notice board REST API built in Go — designed as a production-style backend project covering caching, background jobs, rate limiting, and CI pipelines.

---

## What is ThreadPulse?

ThreadPulse is a notice board API where users can post threads, reply to discussions, and upvote content. The most upvoted threads surface on a "hot threads" endpoint. Every feature in this project was built to learn and demonstrate a specific production backend concept — not just to ship features.

---

## Tech Stack

| Area | Technology |
|---|---|
| Language | Go |
| Framework | Gin |
| Database | PostgreSQL |
| DB Driver | sqlx |
| Migrations | golang-migrate |
| Caching + Rate Limiting | Redis |
| Auth | JWT (golang-jwt/jwt/v5) |
| Containerisation | Docker + docker-compose |
| CI Pipeline | GitHub Actions |

---

## API Endpoints

### Auth
| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| POST | `/auth/register` | Register a new user | No |
| POST | `/auth/login` | Login and receive JWT token | No |

### Threads
| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| POST | `/threads` | Create a new thread | Yes |
| GET | `/threads` | List all threads (paginated) | No |
| GET | `/threads/:id` | Get a single thread | No |
| PUT | `/threads/:id` | Update a thread | Yes |
| DELETE | `/threads/:id` | Delete a thread | Yes |

### Replies
| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| POST | `/threads/:id/replies` | Post a reply to a thread | Yes |
| GET | `/threads/:id/replies` | List replies for a thread (paginated) | No |
| DELETE | `/threads/:id/replies/:replyId` | Delete a reply | Yes |

### Upvotes
| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| POST | `/threads/:id/upvote` | Upvote a thread (processed via background worker) | Yes |

### Hot Threads
| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| GET | `/threads/hot` | Get most upvoted threads (Redis cached) | No |

---

## Project Structure

```
ThreadPulse/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── middleware.go
│   ├── threads/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   ├── replies/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   ├── upvotes/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── worker.go
│   ├── middleware/
│   │   ├── error.go
│   │   └── ratelimit.go
│   └── db/
│       └── db.go
├── migrations/
├── config/
│   └── config.go
├── docker-compose.yml
├── Dockerfile
├── .env
├── .github/
│   └── workflows/
│       └── ci.yml
└── go.mod
```

---

## Key Concepts Implemented

**Pagination** — All list endpoints use offset-based pagination with `page` and `limit` query params. Responses include metadata (total count, current page, total pages) so clients can navigate results without over-fetching.

**Redis Caching with TTL + Cache Invalidation** — The hot threads endpoint is cached in Redis with a TTL. When a new upvote is processed, the cache is invalidated so the next request fetches fresh data from Postgres. This demonstrates the core read-through cache pattern used in production systems.

**Background Worker Pattern** — Upvotes are not processed synchronously in the HTTP handler. The handler enqueues the job onto a Go channel, and a separate worker goroutine processes it independently. This decouples the HTTP response time from the DB write, which is the foundation of async processing in backend systems.

**Centralised Error Handling** — Errors are not handled ad-hoc inside each handler with scattered `c.JSON(400, ...)` calls. Instead, handlers return structured errors and a single middleware layer formats and sends all error responses. This makes error behaviour consistent and easy to change in one place.

**Redis-Backed Rate Limiting** — The post thread and post reply endpoints are rate limited using Redis as the counter store. Each user gets a request budget per time window. Redis is used because it's fast, supports atomic increment operations, and works across multiple instances of the API — something an in-memory counter cannot do.

**GitHub Actions CI Pipeline** — Every push to the repository triggers an automated pipeline that runs the full test suite. This ensures no broken code reaches the main branch and mirrors how production engineering teams work.

---

## Running Locally

**Prerequisites:** Docker and docker-compose installed.

```bash
# Clone the repo
git clone https://github.com/deeep8250/ThreadPulse
cd ThreadPulse

# Copy the env file and fill in your values
cp .env.example .env

# Start everything
docker-compose up --build
```

The API will be available at `http://localhost:8000`.

---

## What I Learned Building This

This was my Phase 2 backend project, built after completing Phase 1 (REST APIs, CRUD, JWT auth, Docker). The goal was to write code that behaves like real production systems — not just code that works.

The biggest shifts from Phase 1:

- **Thinking beyond the request/response cycle.** The background worker pattern was the most conceptually new thing — decoupling a write operation from the HTTP handler changes how you think about what a handler is actually responsible for.
- **Redis as a tool, not magic.** Understanding TTL, cache misses, and invalidation as explicit decisions rather than automatic behaviour.
- **Centralised error handling.** In Phase 1 every handler managed its own errors. Pulling that into middleware made the codebase significantly cleaner and taught me what middleware is actually for.
- **CI as a discipline.** Setting up GitHub Actions forced me to write tests that actually pass in a clean environment, not just locally.

---

## Status

Built as part of a progressive backend learning curriculum (Phase 2 of 4).
