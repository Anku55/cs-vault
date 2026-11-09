# Top 150 Go Interview Problems

## Go Fundamentals & Language Tricky Questions

- [ ] **1.** What is the difference between var, :=, and const?
- [ ] **2.** Explain zero values for all major Go types.
- [ ] **3.** What is the difference between new() and make()?
- [ ] **4.** What is the difference between byte and rune?
- [ ] **5.** What happens when you divide an integer by zero?
- [ ] **6.** Predict the execution order of multiple defer statements.
- [ ] **7.** What happens when defer captures a variable versus its value?
- [ ] **8.** How do panic, recover, and defer interact?
- [ ] **9.** What is variable shadowing? Find a bug in nested scopes.
- [ ] **10.** What is the difference between an array and a slice?
- [ ] **11.** How does Go pass arguments: pass-by-value or pass-by-reference?
- [ ] **12.** What is the difference between a type alias and a defined type?
- [ ] **13.** What is iota? Implement a set of bit flags using iota.
- [ ] **14.** How do named return values interact with defer?
- [ ] **15.** What happens when you compare two structs? Which structs are comparable?

## Slices, Maps, Strings & Memory Traps

- [ ] **16.** Explain the internal representation of a slice: pointer, length, and capacity.
- [ ] **17.** Predict the output when two slices share the same backing array.
- [ ] **18.** What happens when append() exceeds a slice's capacity?
- [ ] **19.** Explain the difference between len() and cap().
- [ ] **20.** What is the difference between a nil slice and an empty slice?
- [ ] **21.** Predict the output of modifying a slice after reslicing.
- [ ] **22.** Why can a small subslice retain a huge backing array in memory?
- [ ] **23.** How do you safely clone a slice?
- [ ] **24.** Implement deletion from a slice without preserving order.
- [ ] **25.** Implement deletion from a slice while preserving order.
- [ ] **26.** How do maps work? What happens when you access a missing key?
- [ ] **27.** Why are maps not safe for concurrent writes?
- [ ] **28.** What happens when you iterate over a map while inserting elements?
- [ ] **29.** Why can't you take the address of a map element?
- [ ] **30.** Explain string immutability, UTF-8, byte indexing, and rune iteration.

## Structs, Interfaces, Pointers & Type System

- [ ] **31.** Explain value receivers vs pointer receivers.
- [ ] **32.** What are method sets in Go?
- [ ] **33.** Can a pointer type implement an interface when the value type cannot? Explain.
- [ ] **34.** Explain the typed nil interface trap with code.
- [ ] **35.** What is the difference between a nil interface and an interface holding a nil pointer?
- [ ] **36.** How does struct embedding work? Is it inheritance?
- [ ] **37.** What happens when embedded structs have conflicting method names?
- [ ] **38.** Explain interface composition.
- [ ] **39.** How does type assertion work? What happens if it fails?
- [ ] **40.** Explain type switch and its use cases.
- [ ] **41.** Why should Go interfaces usually be small?
- [ ] **42.** What is the difference between any and interface{}?
- [ ] **43.** Implement a custom error type that supports errors.As().
- [ ] **44.** Explain errors.Is(), errors.As(), and error wrapping using %w.
- [ ] **45.** Design a generic Stack[T] with Push, Pop, and IsEmpty methods.

## Goroutines, Channels & Concurrency

- [ ] **46.** What is a goroutine? How is it different from an OS thread?
- [ ] **47.** Explain the G-M-P scheduler model.
- [ ] **48.** What happens when a goroutine blocks on an unbuffered channel?
- [ ] **49.** What is the difference between buffered and unbuffered channels?
- [ ] **50.** What happens when you send to a closed channel?
- [ ] **51.** What happens when you receive from a closed channel?
- [ ] **52.** What happens when you send to or receive from a nil channel?
- [ ] **53.** Can multiple goroutines safely send to the same channel?
- [ ] **54.** Who should close a channel? Why is closing from a receiver dangerous?
- [ ] **55.** Explain select. What happens when multiple cases are ready?
- [ ] **56.** How can you implement a timeout using select?
- [ ] **57.** Implement a worker pool using goroutines, channels, and WaitGroup.
- [ ] **58.** Implement fan-in: merge multiple channels into one.
- [ ] **59.** Implement fan-out: distribute tasks to multiple workers.
- [ ] **60.** Implement a producer-consumer system with bounded buffering.
- [ ] **61.** Find and fix a goroutine leak caused by a blocked channel send.
- [ ] **62.** Find and fix a deadlock in a channel-based program.
- [ ] **63.** How do you safely share a map between goroutines?
- [ ] **64.** Explain sync.Mutex vs sync.RWMutex vs channels.
- [ ] **65.** Implement a rate limiter using a token bucket or time.Ticker.

## Synchronization, Context & Concurrency Bugs

- [ ] **66.** Explain sync.WaitGroup. Why should Add() happen before starting a goroutine?
- [ ] **67.** Why should a WaitGroup generally be passed by pointer?
- [ ] **68.** Implement a thread-safe counter using Mutex.
- [ ] **69.** Implement a thread-safe counter using atomic operations.
- [ ] **70.** What is the difference between sync.Once and init()?
- [ ] **71.** When would you use sync.Pool? What are its limitations?
- [ ] **72.** What is sync.Cond? Implement a producer-consumer condition.
- [ ] **73.** Explain race condition vs deadlock vs livelock.
- [ ] **74.** Use go test -race to find and fix a data race.
- [ ] **75.** Implement a concurrent-safe cache with expiration.
- [ ] **76.** Implement cancellation of workers using context.WithCancel().
- [ ] **77.** Implement a timeout-based operation using context.WithTimeout().
- [ ] **78.** Why should context.Context usually be the first function parameter?
- [ ] **79.** Why should context.Context not be stored in a struct?
- [ ] **80.** Implement a concurrent task runner that cancels remaining tasks when one fails using errgroup.

## Go Runtime, Memory, Generics & Testing

- [ ] **81.** Explain stack vs heap allocation in Go.
- [ ] **82.** What is escape analysis? How do you inspect it using go build -gcflags=-m?
- [ ] **83.** How does garbage collection work in Go at a high level?
- [ ] **84.** What is the purpose of GOGC and GOMEMLIMIT?
- [ ] **85.** How do goroutine stacks grow dynamically?
- [ ] **86.** How can you identify memory leaks in a Go application?
- [ ] **87.** How do you profile CPU and memory using pprof?
- [ ] **88.** What is the difference between benchmarking and profiling?
- [ ] **89.** Implement a generic Set[T comparable].
- [ ] **90.** Explain generic type constraints and the ~ operator.
- [ ] **91.** When should you use generics instead of interfaces?
- [ ] **92.** Implement table-driven tests for a function with multiple edge cases.
- [ ] **93.** Write a benchmark for a slice operation and compare alternative implementations.
- [ ] **94.** Implement a fuzz test for a string parser.
- [ ] **95.** How do you test HTTP handlers using httptest and detect data races?

## Go HTTP, REST APIs & Middleware

- [ ] **96.** Explain the difference between http.Handler and http.HandlerFunc.
- [ ] **97.** How does http.ServeMux route requests? Explain method-aware routing.
- [ ] **98.** Implement a REST API using net/http without a framework.
- [ ] **99.** Implement a middleware chain for logging, recovery, and authentication.
- [ ] **100.** Implement middleware that tracks response status code and response size.
- [ ] **101.** Implement request ID middleware and propagate the ID through context.
- [ ] **102.** How do you decode JSON safely and reject unknown fields?
- [ ] **103.** Implement pagination, filtering, and sorting for an API endpoint.
- [ ] **104.** How do you configure HTTP server timeouts to defend against slow clients?
- [ ] **105.** Implement graceful shutdown using signal.NotifyContext and http.Server.Shutdown().
- [ ] **106.** Implement a rate-limiting middleware using golang.org/x/time/rate.
- [ ] **107.** How do you handle CORS correctly? What are preflight requests?
- [ ] **108.** Implement an HTTP client with timeout, cancellation, and connection reuse.
- [ ] **109.** How do you handle retries without retrying non-idempotent requests unsafely?
- [ ] **110.** Implement a custom error response format and centralized API error handling.

## PostgreSQL, SQL & Database Integration

- [ ] **111.** Explain how database/sql manages connections through sql.DB.
- [ ] **112.** Implement CRUD operations using PostgreSQL and pgx or database/sql.
- [ ] **113.** How do you prevent SQL injection in Go?
- [ ] **114.** Implement a transaction with BeginTx, Commit, and Rollback.
- [ ] **115.** Explain ACID properties and PostgreSQL isolation levels.
- [ ] **116.** How do you handle sql.ErrNoRows and nullable database values?
- [ ] **117.** How do you configure and tune a database connection pool?
- [ ] **118.** What causes connection pool exhaustion? How would you diagnose it?
- [ ] **119.** Explain indexes, composite indexes, and index selectivity.
- [ ] **120.** How do you diagnose a slow query using EXPLAIN ANALYZE?
- [ ] **121.** What is the N+1 query problem? Fix it using joins or batch queries.
- [ ] **122.** Implement cursor-based pagination for a PostgreSQL API.
- [ ] **123.** How do you implement optimistic locking using a version column?
- [ ] **124.** How do you handle deadlocks and serialization failures in database transactions?
- [ ] **125.** Implement a repository layer with interfaces and test it using a fake repository.

## Backend Architecture, Security & Production

- [ ] **126.** Design a clean layered Go backend with handler, service, and repository layers.
- [ ] **127.** How do you implement dependency injection in Go without a framework?
- [ ] **128.** Implement JWT authentication and token validation middleware.
- [ ] **129.** How do you securely hash and verify passwords?
- [ ] **130.** Implement role-based authorization for API endpoints.
- [ ] **131.** How do you manage secrets and configuration in a Go service?
- [ ] **132.** Design a caching strategy using Redis and PostgreSQL. Explain cache invalidation.
- [ ] **133.** How do you prevent cache stampedes when many requests miss the same key?
- [ ] **134.** Implement a background job queue with retry and dead-letter handling.
- [ ] **135.** How do you design idempotent API operations?
- [ ] **136.** How do you implement structured logging with log/slog and request correlation IDs?
- [ ] **137.** How would you instrument a Go API using Prometheus metrics and OpenTelemetry tracing?
- [ ] **138.** How do you diagnose high CPU, high memory, or high latency in a Go service?
- [ ] **139.** How do you containerize a Go application using a multi-stage Docker build?
- [ ] **140.** Design a resilient API with timeouts, retries, circuit breakers, rate limiting, and graceful shutdown.

## Famous Go Coding Challenges

- [ ] **141.** Implement a concurrent-safe LRU cache using a map and doubly linked list.
- [ ] **142.** Implement a bounded worker pool that supports graceful shutdown and context cancellation.
- [ ] **143.** Implement a concurrent URL checker with bounded concurrency and per-request timeouts.
- [ ] **144.** Implement a singleflight-style duplicate request suppression mechanism.
- [ ] **145.** Implement a thread-safe in-memory key-value store with TTL expiration.
- [ ] **146.** Implement a producer-consumer pipeline with fan-out, fan-in, and error propagation.
- [ ] **147.** Implement a concurrent HTTP crawler with a maximum depth and deduplication.
- [ ] **148.** Implement a token bucket rate limiter that is safe for concurrent use.
- [ ] **149.** Implement a job queue API with PostgreSQL persistence, workers, retries, and job status tracking.
- [ ] **150.** Design and implement a production-oriented Go HTTP API with authentication, PostgreSQL, Redis caching, middleware, tests, metrics, and graceful shutdown.

---


