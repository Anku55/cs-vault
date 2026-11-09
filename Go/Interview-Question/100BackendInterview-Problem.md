# Top 100 Go Backend Interview Problems

A focused checklist for Go backend interviews. Covers API development, HTTP, security, PostgreSQL, concurrency, architecture, testing, production operations, and system design.

## 1. REST API Design and Implementation (1–15)

- [ ] 1. Design and implement CRUD REST APIs for a user resource using Go.
- [ ] 2. Implement pagination, sorting, and filtering for a list endpoint.
- [ ] 3. Design consistent request and response JSON structures for an API.
- [ ] 4. Validate request payloads and return field-level validation errors.
- [ ] 5. Implement proper HTTP status codes for common API outcomes.
- [ ] 6. Build an endpoint that supports partial updates using PATCH.
- [ ] 7. Implement API versioning without breaking existing clients.
- [ ] 8. Design an API for uploading and downloading files.
- [ ] 9. Implement bulk-create and bulk-update endpoints safely.
- [ ] 10. Prevent duplicate resource creation when clients retry requests.
- [ ] 11. Implement idempotency keys for a payment-like API.
- [ ] 12. Design a search endpoint with multiple optional filters.
- [ ] 13. Return structured errors with error codes and request identifiers.
- [ ] 14. Implement graceful handling of malformed JSON and oversized bodies.
- [ ] 15. Design REST endpoints for a simple order management system.

## 2. HTTP, Routing, and Middleware (16–25)

- [ ] 16. Explain the Go `net/http` request lifecycle from connection to handler.
- [ ] 17. Implement a logging middleware that records method, path, status, and latency.
- [ ] 18. Implement middleware chaining in a reusable way.
- [ ] 19. Write a middleware that captures response status and response size.
- [ ] 20. Implement panic-recovery middleware and explain its limitations.
- [ ] 21. Add request IDs and propagate them through logs and responses.
- [ ] 22. Implement CORS correctly for a browser-facing API.
- [ ] 23. Configure `http.Server` timeouts to protect against slow clients.
- [ ] 24. Implement graceful shutdown on SIGINT/SIGTERM.
- [ ] 25. Explain the differences between `Handle`, `HandleFunc`, `ServeMux`, and `Handler`.

## 3. Authentication, Authorization, and Security (26–35)

- [ ] 26. Implement authentication using secure, hashed passwords.
- [ ] 27. Design access-token and refresh-token flows.
- [ ] 28. Validate JWT claims, expiry, issuer, and signing algorithm safely.
- [ ] 29. Implement role-based access control for API routes.
- [ ] 30. Explain authentication vs authorization and enforce both.
- [ ] 31. Prevent SQL injection using parameterized queries.
- [ ] 32. Protect an API from brute-force login attempts.
- [ ] 33. Store secrets safely using environment/configuration management.
- [ ] 34. Explain CSRF, XSS, and when each threat applies to a Go backend.
- [ ] 35. Design secure password-reset and email-verification flows.

## 4. PostgreSQL and Database Integration (36–50)

- [ ] 36. Connect a Go service to PostgreSQL using `database/sql` and a driver.
- [ ] 37. Implement repository methods for create, read, update, and delete.
- [ ] 38. Configure and explain database connection pooling.
- [ ] 39. Implement pagination using LIMIT/OFFSET and compare it with keyset pagination.
- [ ] 40. Design tables and relationships for users, orders, and order items.
- [ ] 41. Write a query to find the top N customers by total purchase amount.
- [ ] 42. Write a query to find duplicate records and safely remove duplicates.
- [ ] 43. Explain and demonstrate INNER JOIN, LEFT JOIN, and self JOIN.
- [ ] 44. Diagnose and fix the N+1 query problem.
- [ ] 45. Use PostgreSQL indexes to optimize a slow query.
- [ ] 46. Explain composite indexes and how column order affects query plans.
- [ ] 47. Analyze a query using `EXPLAIN ANALYZE`.
- [ ] 48. Implement transactions in Go with commit and rollback paths.
- [ ] 49. Handle database errors such as unique violations and foreign-key violations.
- [ ] 50. Prevent lost updates when multiple requests modify the same row.

## 5. Transactions, Consistency, and Data Modeling (51–60)

- [ ] 51. Explain ACID properties using a practical transfer operation.
- [ ] 52. Implement a money transfer between two accounts transactionally.
- [ ] 53. Explain PostgreSQL isolation levels and demonstrate a race anomaly.
- [ ] 54. Use row-level locking (`SELECT ... FOR UPDATE`) appropriately.
- [ ] 55. Design a schema for a URL shortener and its analytics.
- [ ] 56. Design a schema for a job queue with status, retries, and scheduling.
- [ ] 57. Implement soft deletion and explain its tradeoffs.
- [ ] 58. Perform a database schema migration with backward compatibility.
- [ ] 59. Design a strategy for handling concurrent unique-key creation.
- [ ] 60. Explain when to use JSONB versus normalized relational tables.

## 6. Go Backend Concurrency and Resilience (61–70)

- [ ] 61. Process multiple backend tasks concurrently with bounded worker pools.
- [ ] 62. Implement a worker pool with jobs, results, and graceful shutdown.
- [ ] 63. Prevent goroutine leaks in request handlers and background workers.
- [ ] 64. Propagate request cancellation using `context.Context`.
- [ ] 65. Add timeouts to downstream HTTP and database calls.
- [ ] 66. Implement rate limiting using a token bucket or leaky bucket.
- [ ] 67. Implement retries with exponential backoff and jitter.
- [ ] 68. Explain why retries can amplify outages and how to limit retry storms.
- [ ] 69. Implement a circuit breaker for a failing downstream service.
- [ ] 70. Coordinate concurrent cache misses to avoid a thundering herd.

## 7. Backend Architecture and Design Patterns (71–80)

- [ ] 71. Structure a Go backend into handler, service, and repository layers.
- [ ] 72. Define interfaces at the consumer boundary and inject dependencies.
- [ ] 73. Explain when dependency injection is useful and when it becomes overengineering.
- [ ] 74. Design a modular monolith for an e-commerce backend.
- [ ] 75. Explain monolith vs microservices and identify migration signals.
- [ ] 76. Implement configuration loading and validation for multiple environments.
- [ ] 77. Design a background job system separated from HTTP request processing.
- [ ] 78. Implement the transactional outbox pattern for reliable event publishing.
- [ ] 79. Explain CQRS and where it may be appropriate.
- [ ] 80. Design service-to-service communication using HTTP or gRPC.

## 8. Testing and Code Quality (81–88)

- [ ] 81. Write table-driven unit tests for a service layer.
- [ ] 82. Test HTTP handlers using `httptest`.
- [ ] 83. Mock repository dependencies without making tests brittle.
- [ ] 84. Write integration tests against a real PostgreSQL database.
- [ ] 85. Test concurrent code and use the Go race detector.
- [ ] 86. Test timeouts, cancellation, and retry behavior deterministically.
- [ ] 87. Test API validation, error responses, and edge cases.
- [ ] 88. Explain unit, integration, contract, and end-to-end tests.

## 9. Deployment, Observability, and Production Debugging (89–95)

- [ ] 89. Build a multi-stage Docker image for a Go API.
- [ ] 90. Design a CI pipeline that runs formatting, linting, tests, and builds.
- [ ] 91. Implement structured logging with useful fields and correlation IDs.
- [ ] 92. Instrument an API with metrics such as request rate, latency, and error rate.
- [ ] 93. Explain logs, metrics, and traces; identify when each is useful.
- [ ] 94. Debug high CPU, memory growth, slow endpoints, and goroutine buildup.
- [ ] 95. Design health, readiness, and liveness endpoints.

## 10. Backend System Design Scenarios (96–100)

- [ ] 96. Design a URL shortener with redirects, analytics, and abuse prevention.
- [ ] 97. Design a notification service supporting email, SMS, and push delivery.
- [ ] 98. Design a rate-limiting service for multiple API instances.
- [ ] 99. Design a job queue with delayed jobs, retries, dead-letter handling, and workers.
- [ ] 100. Design a file-upload service with object storage, metadata, and access control.

---

## Suggested Practice Method

For each problem:

1. Write a concise explanation in your notes.
2. Implement the solution in Go where applicable.
3. Add tests, including failure and edge cases.
4. Explain tradeoffs aloud as if in an interview.
5. Record common pitfalls and revisit them during revision.
