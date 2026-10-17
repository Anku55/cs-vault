# URL Shortener 

## Table of Contents

1. [Project Overview](#1-project-overview)
2. [60–90 Second Project Introduction](#2-60–90-second-project-introduction)
3. [Architecture](#3-architecture)
4. [Implemented Milestones](#4-implemented-milestones)
5. [Important Interview Questions and Answers](#5-important-interview-questions-and-answers)
6. [Practical Code Walkthrough](#6-practical-code-walkthrough)
7. [Handling Questions You Cannot Answer](#7-handling-questions-you-cannot-answer)
8. [Pre-Interview Checklist](#8-pre-interview-checklist)

---

## 1. Project Overview

A URL Shortener converts a long URL into a short link. When someone opens that short link, the server finds the original URL and redirects the user to it.

### Implemented features

- `POST /shorten` — accepts a URL and returns a short link.
- `GET /{shortCode}` — looks up the original URL and redirects to it.
- URL validation — accepts only `http` and `https` URLs with a non-empty host.
- In-memory storage — stores short-code-to-URL mappings in a Go map.
- Tests — covers handler, service, and storage behavior.

### Not implemented yet

- Persistent database storage
- Redis caching
- Authentication
- Rate limiting
- Analytics
- Production-scale distributed deployment

Be clear about the difference between implemented features and future improvements.

---

## 2. 60–90 Second Project Introduction

Use this as a guide; understand the meaning rather than memorizing it word-for-word.

> I have been building a URL Shortener using Go. The main purpose is to convert long URLs into shorter links that redirect users to the original URLs.
>
> So far, I have implemented three main features. First, a `POST /shorten` endpoint that accepts a URL, validates it, generates a random short code, and stores the mapping in memory. Second, a `GET /{shortCode}` endpoint that retrieves the original URL and redirects the user to it. Third, URL validation to reject invalid URLs before storing them.
>
> I use Go's `net/http` package to handle HTTP requests, `net/url` for URL parsing and validation, and `crypto/rand` for generating random short codes.
>
> The project follows a layered architecture with Handler, Service, and Storage layers. This separation makes the code easier to test and maintain.
>
> Currently, the project uses in-memory storage, so mappings are lost when the server restarts. My next goal is to add persistent storage and improve the system further.

---

## 3. Architecture

```text
Client (Browser / Postman)
          |
          v
    Handler Layer
  HTTP, JSON, status codes
          |
          v
     Service Layer
 Validation, business logic,
     short-code generation
          |
          v
     Storage Layer
 Map: short code -> original URL
```

The response travels back to the client. For a redirect request, the Handler, Service, and Storage work together to retrieve the original URL, then the server sends an HTTP redirect response.

### Project files

| File | Responsibility |
|---|---|
| `cmd/server/main.go` | Creates components and starts the server |
| `internal/handler/url.go` | Handles HTTP requests and responses |
| `internal/service/url.go` | URL validation and shortening logic |
| `internal/storage/memory.go` | Stores URL mappings in memory |
| `internal/handler/url_test.go` | Tests HTTP handler behavior |
| `internal/service/url_test.go` | Tests service logic |
| `internal/storage/memory_test.go` | Tests storage behavior |

### Responsibilities of each layer

- **Handler:** Parses the request, reads JSON, calls the Service, and writes an HTTP response.
- **Service:** Implements business rules such as validation and short-code generation.
- **Storage:** Saves and retrieves the mapping between a short code and its original URL.

### Why this architecture?

It separates HTTP-specific logic from business logic and storage. That lowers coupling, improves testability, and makes it easier to replace the storage implementation later.

---

## 4. Implemented Milestones

### Milestone 1 — URL Shortening

Example request:

```http
POST /shorten
Content-Type: application/json
```

```json
{
  "url": "https://example.com/articles/long-page"
}
```

Flow:

1. The Handler reads the JSON request.
2. The Service validates the URL.
3. A random short code is generated.
4. Storage saves the short-code-to-original-URL mapping.
5. The client receives a short URL.

Example response (illustrative; check the actual response format in your code):

```json
{
  "short_url": "http://localhost:8080/tVnxW6"
}
```

The expected success status for a creation endpoint is `201 Created`; verify what your Handler actually returns.

### Milestone 2 — URL Redirect

When a user opens a URL such as `http://localhost:8080/tVnxW6`:

- If the short code exists, the server retrieves the original URL and redirects the client.
- If the short code does not exist, the server returns `404 Not Found`.
- If the request uses an unsupported HTTP method, the server returns `405 Method Not Allowed`.
- The redirect response uses a `Location` header to identify the destination.
- The current successful redirect test checks `302 Found` and the `Location` header.

### Milestone 3 — URL Validation

The current validation rules are:

- The URL must parse successfully.
- The scheme must be `http` or `https`.
- The host must not be empty.

| Input | Expected result |
|---|---|
| `https://google.com` | Valid |
| `http://example.com/page` | Valid |
| `example.com` | Invalid: missing scheme |
| `ftp://example.com` | Invalid: unsupported scheme |
| `https:///page` | Invalid: missing host |
| `random text` | Invalid |

When validation fails, the Service returns an error and the Handler returns `400 Bad Request`.

This is basic structural validation. It does **not** prove that a destination exists, is reachable, or is safe.

---

## 5. Important Interview Questions and Answers

### A. Project overview and design

**1. Why did you build a URL Shortener?**

A URL Shortener is a practical backend project that helped me learn HTTP APIs, request handling, validation, random code generation, storage abstraction, error handling, and testing. It also gives me a foundation for understanding backend system design.

**2. What features have you implemented?**

I have implemented URL shortening through `POST /shorten`, redirection through `GET /{shortCode}`, and URL validation. I have also written tests for service, storage, and handler behavior.

**3. Explain the architecture.**

The Handler handles HTTP requests and responses. The Service contains business logic, including URL validation and short-code generation. The Storage layer saves and retrieves URL mappings. The Handler calls the Service, and the Service uses the Storage abstraction.

**4. Why did you use Go?**

Go provides a straightforward language design, built-in concurrency support, a standard HTTP library, interfaces, and a built-in testing framework. These features make it suitable for backend and networking applications.

**5. What was an important technical decision?**

Separating Handler, Service, and Storage. It isolates HTTP-specific code from business logic and data storage, making components easier to test and replace.

**6. What are the current limitations?**

Mappings are stored in memory and disappear when the process restarts. The current project also does not yet include persistent storage, rate limiting, analytics, authentication, or distributed deployment.

### B. Go concepts

**7. What is an interface in Go?**

An interface defines a set of method signatures. A type satisfies an interface by implementing its methods; it does not need an explicit declaration. A Storage interface allows the Service to depend on the storage contract rather than a particular implementation.

**8. What is dependency injection?**

Dependency injection means providing a component with its dependencies instead of making it create them internally. Passing Storage to Service and Service to Handler makes dependencies explicit and allows alternative implementations in tests.

**9. Why use `cmd` and `internal`?**

`cmd/server` contains the application entry point. Go's `internal` directory restricts which packages may import those implementation packages, helping separate internal implementation from public APIs.

**10. What is `net/http`?**

It is Go's standard package for HTTP clients and servers. This project uses it to handle requests, write responses, and send redirects.

**11. What is `net/url`?**

It provides URL parsing and URL-related utilities. The project uses it to inspect the URL scheme and host during validation.

**12. Why use `crypto/rand` instead of `math/rand`?**

`crypto/rand` produces cryptographically secure random values and is useful when codes should be difficult to predict. `math/rand` is for ordinary pseudorandom use cases and is not designed for security-sensitive randomness. Neither one alone guarantees uniqueness.

**13. Map versus slice?**

A map stores key-value pairs and supports lookup by key. A slice is a dynamically sized view over an underlying array. A map is a natural fit for looking up an original URL by short code.

**14. Why use errors instead of exceptions?**

Go commonly represents recoverable failures with explicit `error` return values. The caller checks and handles the error. In this project, validation errors propagate to the Handler, which translates them into HTTP responses.

### C. HTTP and API questions

**15. Why is `/shorten` a POST endpoint?**

It accepts input and creates a URL mapping. POST is appropriate for this operation, with the original URL commonly sent in a JSON request body.

**16. Why is the redirect endpoint GET?**

Opening a short URL retrieves the destination and performs navigation. GET is appropriate because the client is requesting the destination rather than creating a mapping.

**17. Explain these status codes.**

| Code | Meaning | Typical use |
|---|---|---|
| `201` | Created | URL mapping created |
| `302` | Found | Temporary redirect |
| `400` | Bad Request | Invalid URL or request |
| `404` | Not Found | Short code does not exist |
| `405` | Method Not Allowed | Unsupported method |
| `500` | Internal Server Error | Unexpected server-side failure |

Check the actual Handler before claiming every status is implemented.

**18. What is the `Location` header?**

It identifies the destination in a redirect response. A browser receiving a redirect status and the `Location` header can request the destination.

**19. Difference between `301`, `302`, `307`, and `308`?**

`301` and `308` are permanent redirects; `302` and `307` are temporary redirects. `307` and `308` preserve the HTTP method. `301` and `302` may have method-changing behavior in common clients.

**20. Why return JSON?**

JSON is a widely supported way to exchange structured data between clients and backend services. It lets the API return the short URL in a predictable response format.

**21. What if the client sends GET to `/shorten`?**

The endpoint should return `405 Method Not Allowed` for an unsupported method. An `Allow` header listing supported methods is also useful if implemented.

### D. Validation and short-code generation

**22. How do you validate a URL?**

I parse it with `net/url`, check that parsing succeeds, ensure the scheme is `http` or `https`, and verify that the host is non-empty. Invalid input produces a Service error, which the Handler maps to `400 Bad Request`.

**23. Is `https://example.com` valid? What about `example.com`?**

The first passes the current validation rules. The second is rejected because it lacks an explicit HTTP or HTTPS scheme.

**24. How do you generate a short code?**

The implementation uses `crypto/rand` and an alphanumeric character set to generate a random code. Verify the actual code length and character set in the implementation before stating exact values.

**25. Can two URLs get the same short code?**

Yes. Random generation does not guarantee uniqueness. A robust implementation should check whether the code already exists and generate another one if a collision occurs. Do not claim collision retry is implemented unless you have verified it in the code.

**26. How many six-character codes are possible?**

If the alphabet contains 62 characters (`a-z`, `A-Z`, `0-9`), the number of six-character codes is:

`62^6 = 56,800,235,584`

This assumes six positions and a 62-character alphabet.

**27. Does validation guarantee the website exists?**

No. Structural validation does not prove the domain resolves, the server is reachable, or the page exists. Network checks would introduce additional latency, timeouts, and security considerations.

### E. Testing

**28. What is unit testing?**

Unit testing checks a component independently. The project has tests for service logic, storage operations, and handler behavior.

**29. What is `net/http/httptest`?**

It provides tools to test HTTP handlers without starting a real server. Tests can construct requests and inspect recorded response status codes, headers, and bodies.

**30. Which cases have you tested?**

The described test suite covers successful redirects, missing short codes (`404`), wrong methods (`405`), and URL validation cases such as missing schemes, unsupported schemes, and missing hosts. Confirm the exact cases in the test files.

**31. How do you run tests?**

```bash
go test ./...
```

For package-level coverage:

```bash
go test -cover ./...
```

**32. Unit testing versus integration testing?**

Unit tests check components in isolation. Integration tests check whether multiple components work together. Testing the complete HTTP flow through Handler, Service, and Storage is closer to an integration test.

**33. How would you test invalid JSON?**

Send malformed JSON to `POST /shorten` and verify a `400 Bad Request` response. Also test missing and empty URL fields, unsupported content types if enforced, and oversized request bodies.

### F. Storage and system design

**34. Why use in-memory storage?**

It lets me focus first on API design, architecture, and core functionality without introducing database complexity. It is useful for learning and prototyping but not sufficient for durable production storage.

**35. What happens after a server restart?**

Mappings held in memory are lost. Persistent storage such as PostgreSQL would preserve mappings across restarts.

**36. Why not use a database from the start?**

I wanted to implement the core flow and layer separation first. Because the Service depends on a Storage abstraction, I can add a database-backed implementation later without fundamentally changing the API.

**37. What is the lookup time complexity?**

Go map lookup is expected `O(1)` on average. This describes the map lookup, not the total request latency.

**38. Is a regular Go map safe for concurrent access?**

A regular map is not safe for concurrent reads and writes without synchronization. Verify whether the storage implementation uses a mutex or another safe approach. If not, add synchronization and concurrent tests; the race detector can help find data races.

**39. How would you scale the project?**

I would add persistent storage such as PostgreSQL, enforce unique short codes, add indexes and rate limiting, and introduce monitoring and deployment automation. Redis could be evaluated as a cache for frequently accessed mappings. Multiple application instances should share a consistent persistent store.

**40. Why PostgreSQL instead of only Redis?**

PostgreSQL offers durable relational storage and transaction support, making it a good source of truth. Redis can provide fast lookups and caching; its durability depends on configuration. The choice depends on the requirements.

### G. Security and production readiness

**41. What security issues can a URL Shortener face?**

It can hide phishing, malicious destinations, and spam. Other concerns include code guessing and abuse of the shortening endpoint. Production systems may need rate limiting, abuse reporting, monitoring, and destination screening.

**42. What is rate limiting?**

It limits how many requests a client can make during a defined period. It can reduce automated abuse and resource exhaustion.

**43. Does URL validation prevent SSRF?**

No. Checking the scheme and host is not enough to prevent Server-Side Request Forgery. SSRF is especially relevant if the server fetches user-supplied URLs. Such a design needs careful destination/IP validation, redirect controls, and network restrictions. The current basic shortening flow does not establish that a destination is safe.

**44. Are random short codes authentication tokens?**

No. They identify URL mappings and are not a replacement for authentication or authorization. Randomness can make codes harder to guess, but it does not provide full access control.

**45. What would you improve before production?**

I would add persistent storage, concurrency-safe operations, collision handling, request-size limits, rate limiting, structured logging, graceful shutdown, stronger error handling, integration tests, and monitoring.

---

## 6. Practical Code Walkthrough

Be ready to open these files and explain the relevant code:

| Interview question | File to show |
|---|---|
| Where does the application start? | `cmd/server/main.go` |
| Where is URL validation implemented? | `internal/service/url.go` |
| Where is the short code generated? | Service code using `crypto/rand` |
| Where are URLs stored? | `internal/storage/memory.go` |
| How are HTTP requests handled? | `internal/handler/url.go` |
| How does redirect work? | Handler redirect logic |
| How are errors translated to status codes? | Service error returns and Handler |
| How is redirect tested? | `internal/handler/url_test.go` |
| How are components wired together? | `cmd/server/main.go` |

### Code concepts to revise

- `http.Handler` and `http.HandlerFunc`
- `http.Request` and `http.ResponseWriter`
- `json.NewDecoder` and `json.NewEncoder`
- `url.Parse`
- `crypto/rand`
- Interfaces and dependency injection
- Error propagation
- `httptest.NewRecorder()` and `httptest.NewRequest()`
- Go maps and concurrent access

**Important:** Explain what the code currently does, not only what a robust design should do. For example, if collision retry is not implemented, describe it as a planned improvement.

---

## 7. Handling Questions You Cannot Answer

Use this structure:

1. State what you know.
2. Identify what needs verification.
3. Explain how you would approach it.

Example:

> Currently, I am using an in-memory map for storage. I know that a regular Go map is not safe for concurrent writes without synchronization. I would verify the current implementation and, if needed, protect map access with a mutex. I would also add concurrent tests and run the Go race detector.

Useful phrases:

- “I haven't implemented that feature yet, but I understand how I would approach it.”
- “I would need to check the current implementation before making that claim.”
- “This is a limitation of the current version that I plan to address.”

---

## 8. Pre-Interview Checklist

- [ ] Explain the project in 60–90 seconds.
- [ ] Walk through `main.go`, Handler, Service, and Storage.
- [ ] Explain `POST /shorten` from request to response.
- [ ] Explain redirect lookup, `302`, and the `Location` header.
- [ ] Explain URL validation and `400`, `404`, and `405`.
- [ ] Revise interfaces, dependency injection, maps, errors, and `crypto/rand`.
- [ ] Explain tests and `go test ./...`.
- [ ] Explain in-memory persistence, concurrency, collisions, and future improvements.

### Final practice

1. Run `go test ./...`.
2. Start the server and shorten a valid URL.
3. Open the generated short URL and verify the redirect.
4. Submit an invalid URL and verify the `400` response.
5. Try a nonexistent short code and verify the `404` response.
6. Open the code and practise explaining the architecture aloud.

**Main goal:** Be able to explain the decisions behind your implemented code and openly identify its limitations. You do not need to pretend the project is finished.
