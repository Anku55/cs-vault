# URL Shortener

A backend URL Shortener built with **Go**. The project converts long URLs into short links and redirects users from the short code back to the original URL.

I built this project to practise backend development concepts such as HTTP APIs, layered architecture, interfaces, dependency injection, URL validation, random short-code generation, in-memory storage, error handling, and testing.

> **Project status:** Functional learning/prototype project. Storage is currently in memory, so mappings are lost when the server restarts.

---

## Overview

The URL Shortener currently provides two main API operations:

- `POST /shorten` — accepts a URL and returns a short link.
- `GET /{shortCode}` — looks up the original URL and redirects the client.

It also validates URLs before storing them and includes tests for the handler, service, and storage layers. fileciteturn0file0L35-L51

### Current features

- URL shortening
- HTTP redirects
- URL validation
- Random short-code generation
- In-memory URL storage
- Handler, service, and storage tests
- Layered architecture
- Error handling

### Not implemented yet

- Persistent database storage
- Redis caching
- Authentication
- Rate limiting
- Analytics
- Production-scale distributed deployment fileciteturn0file0L55-L69

---

## Tech Stack

- **Language:** Go
- **HTTP:** `net/http`
- **URL parsing:** `net/url`
- **Random code generation:** `crypto/rand`
- **Testing:** Go testing package / `httptest`
- **Storage:** Go map
- **API format:** JSON

The project uses Go's standard library rather than relying on external frameworks. The implementation uses `net/http` for HTTP handling, `net/url` for URL parsing, and `crypto/rand` for random short-code generation. fileciteturn0file0L89-L105

---

## Architecture

The project follows a simple layered architecture:

```text
              Client
        Browser / Postman
               |
               v
        ┌───────────────┐
        │    Handler    │
        │ HTTP / JSON   │
        │ Status Codes  │
        └───────┬───────┘
                |
                v
        ┌───────────────┐
        │    Service    │
        │ Validation    │
        │ Business Logic│
        │ Code Creation │
        └───────┬───────┘
                |
                v
        ┌───────────────┐
        │    Storage    │
        │ shortCode →   │
        │ original URL  │
        └───────────────┘
```

The Handler is responsible for HTTP-specific concerns, the Service contains business logic, and Storage handles saving and retrieving URL mappings. fileciteturn0file0L113-L151

### Why I chose this architecture

I wanted to keep HTTP handling separate from business logic and storage.

This gives me:

- Lower coupling
- Easier testing
- Clear responsibilities
- The ability to replace the storage implementation later

For example, the current in-memory storage can eventually be replaced with PostgreSQL without changing the overall API flow.

---

## Project Structure

```text
URL-Shortener/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   │   ├── url.go
│   │   └── url_test.go
│   │
│   ├── service/
│   │   ├── url.go
│   │   └── url_test.go
│   │
│   └── storage/
│       ├── memory.go
│       └── memory_test.go
│
└── ...
```

### Responsibilities

| File | Responsibility |
|---|---|
| `cmd/server/main.go` | Creates components and starts the server |
| `internal/handler/url.go` | Handles HTTP requests and responses |
| `internal/service/url.go` | URL validation and shortening logic |
| `internal/storage/memory.go` | Stores URL mappings in memory |
| `internal/handler/url_test.go` | Tests HTTP handler behavior |
| `internal/service/url_test.go` | Tests service logic |
| `internal/storage/memory_test.go` | Tests storage behavior |

These responsibilities follow the current project structure described in the project notes. fileciteturn0file0L155-L175

---

# API

## `POST /shorten`

Creates a short URL from an original URL.

### Request

```http
POST /shorten
Content-Type: application/json
```

```json
{
  "url": "https://example.com/articles/long-page"
}
```

### Flow

```text
Client
  |
  v
Handler reads JSON
  |
  v
Service validates URL
  |
  v
Short code generated
  |
  v
Storage saves mapping
  |
  v
Short URL returned
```

The expected creation response is commonly `201 Created`; the exact status should always be verified against the current Handler implementation. fileciteturn0file0L203-L269

---

## `GET /{shortCode}`

Redirects the client to the original URL.

For example:

```text
GET /tVnxW6
```

If the short code exists:

```text
302 Found
Location: https://example.com/articles/long-page
```

If the short code doesn't exist:

```text
404 Not Found
```

An unsupported HTTP method should result in:

```text
405 Method Not Allowed
```

The current successful redirect test checks for `302 Found` and the `Location` header. fileciteturn0file0L273-L289

---

# URL Validation

Before storing a URL, the Service performs basic structural validation.

The current rules are:

1. The URL must parse successfully.
2. The scheme must be `http` or `https`.
3. The host must not be empty.

### Examples

| Input | Result |
|---|---|
| `https://google.com` | Valid |
| `http://example.com/page` | Valid |
| `example.com` | Invalid — missing scheme |
| `ftp://example.com` | Invalid — unsupported scheme |
| `https:///page` | Invalid — missing host |
| `random text` | Invalid |

When validation fails, the Service returns an error and the Handler translates it into a `400 Bad Request`. fileciteturn0file0L293-L331

### Important limitation

This is only **structural URL validation**.

It does not prove that:

- The domain exists
- The server is reachable
- The page exists
- The destination is safe

I intentionally keep this separate from network-level validation because checking the destination would introduce additional latency, failure cases, and security considerations.

---

# Short-Code Generation

The project uses:

```go
crypto/rand
```

to generate random short codes.

I chose `crypto/rand` because it provides cryptographically secure randomness and makes generated codes harder to predict compared with ordinary pseudorandom generation.

However, randomness alone does **not** guarantee uniqueness.

A robust production implementation should:

1. Generate a code.
2. Check whether it already exists.
3. Retry if a collision occurs.

The current implementation should be described accurately rather than claiming collision retry unless it has actually been implemented. fileciteturn0file0L439-L445

### Possible code space

If the alphabet contains:

```text
a-z
A-Z
0-9
```

there are 62 possible characters.

For six-character codes:

```text
62^6 = 56,800,235,584
```

This is a large space, but collisions are still theoretically possible.

---

# Storage

The current project uses an in-memory Go map:

```text
short code → original URL
```

Conceptually:

```text
tVnxW6 → https://example.com/articles/long-page
a8KpQ2 → https://example.com/products/item
```

### Why in-memory storage?

I wanted to first focus on:

- API design
- Layered architecture
- Interfaces
- Validation
- HTTP handling
- Testing

without introducing database configuration and persistence complexity.

The downside is important:

> **All mappings are lost when the server restarts.**

A persistent database would be required for a real production system. fileciteturn0file0L669-L685

---

# Go Concepts Used

## Interfaces

An interface defines a set of method signatures.

The Service depends on a Storage abstraction rather than a specific storage implementation.

This means I can have:

```text
Service
   |
   v
Storage interface
   |
   ├── In-memory implementation
   |
   └── Future PostgreSQL implementation
```

A type satisfies a Go interface by implementing its required methods; it doesn't need an explicit declaration. fileciteturn0file0L395-L407

---

## Dependency Injection

Instead of making each component create its dependencies internally, dependencies are provided to it.

Conceptually:

```text
Storage → Service → Handler
```

This makes dependencies explicit and makes testing easier because I can provide alternative implementations.

For example, a test can inject a controlled storage implementation instead of relying on the real application storage.

---

## `cmd` and `internal`

The project follows a common Go layout:

```text
cmd/server/
```

contains the application entry point.

```text
internal/
```

contains implementation packages that should not be imported by unrelated external projects.

This helps keep application internals separated from any public API surface. fileciteturn0file0L415-L421

---

# HTTP Concepts

## Why `POST /shorten`?

The endpoint accepts input and creates a new URL mapping.

`POST` is appropriate because the operation creates server-side state.

---

## Why `GET /{shortCode}`?

Opening a short URL retrieves the destination and navigates to it.

The client is requesting a resource/action rather than creating a new mapping, so `GET` is appropriate.

---

## HTTP Status Codes

| Status | Meaning | Example |
|---|---|---|
| `201` | Created | Short URL created |
| `302` | Found | Redirect to original URL |
| `400` | Bad Request | Invalid URL |
| `404` | Not Found | Unknown short code |
| `405` | Method Not Allowed | Unsupported HTTP method |
| `500` | Internal Server Error | Unexpected server-side failure |

The exact status behavior should always be verified against the current Handler implementation before claiming that every status is implemented. fileciteturn0file0L463-L505

---

## `Location` Header

For redirects, the response contains a `Location` header.

For example:

```http
HTTP/1.1 302 Found
Location: https://example.com/articles/long-page
```

The client can use this header to determine where it should navigate. fileciteturn0file0L509-L513

---

# Testing

Testing is an important part of the project.

The repository contains tests for:

- Handler behavior
- Service logic
- Storage behavior

The project uses Go's HTTP testing tools so handlers can be tested without starting a real server.

For example:

```go
httptest.NewRequest()
httptest.NewRecorder()
```

allow tests to create HTTP requests and inspect the resulting status codes, headers, and response bodies. fileciteturn0file0L601-L625

## Run tests

```bash
go test ./...
```

## Run tests with coverage

```bash
go test -cover ./...
```

---

## Test Cases

The current test coverage described for the project includes:

- Successful redirects
- Missing short codes → `404`
- Unsupported methods → `405`
- Valid HTTP URLs
- Valid HTTPS URLs
- Missing schemes
- Unsupported schemes
- Missing hosts

Additional useful tests would include:

- Invalid JSON
- Missing `url` field
- Empty URL
- Duplicate short-code collisions
- Concurrent access
- Large request bodies

The exact test suite should be verified against the current test files before claiming a particular case is implemented. fileciteturn0file0L605-L665

---

# Concurrency

A regular Go map is **not safe for concurrent reads and writes without synchronization**.

This matters because an HTTP server can process multiple requests concurrently.

If the current storage implementation does not protect the map, I would improve it using synchronization such as a mutex and add concurrent tests.

The Go race detector would also be useful for finding data races.

This is an important production consideration even though the current project is primarily a learning/prototype implementation. fileciteturn0file0L697-L709

---

# Scaling the Project

If I wanted to take this from a learning project toward a production system, I would improve it in stages.

### 1. Persistent storage

Move mappings from an in-memory map to something such as PostgreSQL.

```text
Handler
   ↓
Service
   ↓
Storage interface
   ↓
PostgreSQL
```

This would preserve mappings across restarts.

### 2. Collision handling

Make short-code creation explicitly check for existing codes and retry on collisions.

### 3. Concurrency safety

Ensure storage operations are safe when multiple HTTP requests access the system simultaneously.

### 4. Rate limiting

Limit how many requests a client can make within a period of time.

### 5. Monitoring

Add:

- Structured logging
- Metrics
- Health checks
- Error monitoring

### 6. Caching

Redis could be introduced for frequently accessed mappings.

A sensible architecture would use PostgreSQL as the durable source of truth and Redis as an optional cache. fileciteturn0file0L711-L725

---

# Security Considerations

A URL Shortener has several security and abuse concerns.

## Malicious destinations

Short links can hide phishing or malicious URLs.

A production system could require:

- Abuse reporting
- Destination screening
- Monitoring
- Rate limiting

## Predictable codes

Short codes should not be treated as authentication credentials.

Random codes can make guessing harder, but they do not provide authentication or authorization.

## SSRF

Basic URL validation does **not** prevent SSRF.

This project currently doesn't fetch user-supplied URLs from the server, so SSRF is not established as a vulnerability in the current shortening flow.

If the system were extended to fetch or inspect destinations server-side, it would need stronger controls around:

- Private IP ranges
- Redirects
- DNS resolution
- Network access
- Destination validation

fileciteturn0file0L729-L765

---

# Limitations

The current implementation is intentionally small.

The biggest limitations are:

### 1. In-memory storage

Mappings disappear after a restart.

### 2. No explicit persistence

There is currently no database.

### 3. Collision handling needs attention

Random generation does not itself guarantee uniqueness.

### 4. Concurrency safety needs to be verified

A normal Go map requires synchronization when concurrent writes are possible.

### 5. No rate limiting

The API can be abused without request throttling.

### 6. No analytics

The current application doesn't track click counts, referrers, geographic data, or other usage metrics.

### 7. No authentication

Short URLs are not protected by user accounts or authorization.

These limitations are intentional and help define the current scope of the project. fileciteturn0file0L55-L69

---

# What I Learned

This project helped me understand backend development beyond simply writing an HTTP endpoint.

The main concepts I practised were:

- HTTP APIs
- Request and response handling
- JSON
- HTTP status codes
- Redirects
- URL parsing
- Input validation
- Go interfaces
- Dependency injection
- Layered architecture
- In-memory storage
- Error propagation
- Unit testing
- HTTP testing
- Concurrency considerations
- Basic backend security
- System design trade-offs

The biggest architectural lesson for me was the value of separating the **Handler**, **Service**, and **Storage** layers.

It makes the current implementation easier to understand and gives me a clear path for replacing the in-memory map with persistent storage later.

---

# Running the Project

## Prerequisites

Install Go and verify:

```bash
go version
```

## Run

From the project root:

```bash
go run ./cmd/server
```

The server should then be available on the configured local address.

## Test

```bash
go test ./...
```

## Test with coverage

```bash
go test -cover ./...
```

---

# Example Workflow

### 1. Start the server

```bash
go run ./cmd/server
```

### 2. Create a short URL

```http
POST /shorten
Content-Type: application/json
```

```json
{
  "url": "https://example.com/articles/long-page"
}
```

### 3. Receive the short URL

Example:

```json
{
  "short_url": "http://localhost:8080/tVnxW6"
}
```

### 4. Open the short URL

```text
http://localhost:8080/tVnxW6
```

The server looks up the code and redirects to the original URL.

---

# Project Status

**Current status: Functional learning/prototype project**

Implemented:

- URL shortening
- Redirects
- URL validation
- Random short-code generation
- In-memory storage
- Layered architecture
- Service, storage, and handler tests

Planned improvements:

- Persistent database
- Collision handling
- Concurrency-safe storage
- Rate limiting
- Analytics
- Better production observability
- Distributed deployment

---

# Interview Preparation

If I discuss this project in an interview, I should be able to explain the implementation rather than only describe the features.

The main things I should know are:

1. **Why I chose Go**
2. **Why I separated Handler, Service, and Storage**
3. **How `POST /shorten` works from request to response**
4. **How the redirect flow works**
5. **How URL validation works**
6. **Why I used an interface for Storage**
7. **How dependency injection is used**
8. **Why `crypto/rand` is used**
9. **Why in-memory storage was chosen initially**
10. **What happens after a server restart**
11. **Why a regular map needs synchronization for concurrent access**
12. **How I would handle short-code collisions**
13. **How I tested the HTTP handlers**
14. **What I would change before production**

The important part is to be honest about what is implemented and what is still a planned improvement.

---

## If I Don't Know an Interview Question

I would avoid guessing.

A good way to answer is:

> I haven't implemented that yet, but I understand the general approach. I would first verify the current implementation, then make the change in the appropriate layer and add tests for it.

Or:

> I would need to check the current implementation before making that claim. My understanding is that the current limitation is ____, and I would address it by ____.

This is better than claiming that a feature exists when it doesn't.

---

## Final Checklist

Before discussing the project, I should be able to:

- [ ] Explain the project in 60–90 seconds
- [ ] Walk through `main.go`
- [ ] Explain the Handler layer
- [ ] Explain the Service layer
- [ ] Explain the Storage layer
- [ ] Explain `POST /shorten`
- [ ] Explain `GET /{shortCode}`
- [ ] Explain `302` and the `Location` header
- [ ] Explain `400`, `404`, and `405`
- [ ] Explain URL validation
- [ ] Explain interfaces
- [ ] Explain dependency injection
- [ ] Explain `crypto/rand`
- [ ] Explain map lookup
- [ ] Explain the current in-memory limitation
- [ ] Explain concurrency concerns
- [ ] Explain collision handling
- [ ] Run `go test ./...`
- [ ] Explain the existing tests
- [ ] Explain how I would scale the system

---

## Final Note

I built this project as a practical way to learn backend development with Go.

The current implementation is intentionally small, but I understand the architecture, the trade-offs, and the areas that would need improvement for production.

My goal is not to present it as a finished large-scale system. The value of the project is that I can explain **why I made the decisions I made, how the current implementation works, and how I would evolve it as the requirements grow.**
