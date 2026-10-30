# Go HTTP Middleware --- Complete Learning Notes

> **Purpose:** A detailed revision guide for understanding,
> implementing, testing, and extending HTTP middleware in Go.
>
> **Stack:** Go standard library --- `net/http`, `net/http/httptest`,
> `sync`, `time`, `context`, and `testing`.
>
> **Learning status:** Core middleware concepts are covered. The final
> section gives the recommended next topic and a practical learning
> sequence.

------------------------------------------------------------------------

## Table of Contents

1.  [What is middleware?](#1-what-is-middleware)
2.  [HTTP fundamentals required for
    middleware](#2-http-fundamentals-required-for-middleware)
3.  [The `http.Handler` interface](#3-the-httphandler-interface)
4.  [`http.HandlerFunc` and
    `HandleFunc`](#4-httphandlerfunc-and-handlefunc)
5.  [Understanding `ServeMux`](#5-understanding-servemux)
6.  [Writing your first middleware](#6-writing-your-first-middleware)
7.  [Understanding `next.ServeHTTP`](#7-understanding-nextservehttp)
8.  [Middleware chaining and execution
    order](#8-middleware-chaining-and-execution-order)
9.  [Logging middleware](#9-logging-middleware)
10. [Authentication middleware](#10-authentication-middleware)
11. [Timing middleware](#11-timing-middleware)
12. [Understanding and wrapping
    `http.ResponseWriter`](#12-understanding-and-wrapping-httpresponsewriter)
13. [Recovery middleware and panic
    handling](#13-recovery-middleware-and-panic-handling)
14. [CORS middleware](#14-cors-middleware)
15. [Rate limiting fundamentals](#15-rate-limiting-fundamentals)
16. [Fixed Window rate limiting](#16-fixed-window-rate-limiting)
17. [Token Bucket rate limiting](#17-token-bucket-rate-limiting)
18. [Testing with `net/http/httptest`](#18-testing-with-nethttphttptest)
19. [Concurrency, mutexes, and race
    detection](#19-concurrency-mutexes-and-race-detection)
20. [Middleware benchmarking](#20-middleware-benchmarking)
21. [Production design
    considerations](#21-production-design-considerations)
22. [Recommended project structure](#22-recommended-project-structure)
23. [Common mistakes and debugging
    checklist](#23-common-mistakes-and-debugging-checklist)
24. [What to learn next in
    middleware](#24-what-to-learn-next-in-middleware)
25. [Practice assignments and completion
    checklist](#25-practice-assignments-and-completion-checklist)

------------------------------------------------------------------------

# 1. What is middleware?

Middleware is a function that receives an HTTP handler, adds some
behaviour around it, and returns another HTTP handler.

A middleware usually performs work **before** the next handler,
**after** the next handler, or both.

Examples:

-   Logging: record the method, path, status code, and request duration.
-   Authentication: verify whether a request has valid credentials.
-   Authorization: check whether an authenticated user may perform an
    operation.
-   Recovery: recover from a panic in the request-handling goroutine.
-   CORS: add cross-origin response headers and handle preflight
    requests.
-   Rate limiting: restrict how frequently a client can call an
    endpoint.
-   Request ID: attach a unique identifier to each request.
-   Metrics: record request counts, response codes, and latency.

## 1.1 Middleware as a wrapper

Think of a handler as the main work and middleware as a layer around
that work.

``` text
Incoming HTTP request
        |
        v
   Logging middleware
        |
        v
 Authentication middleware
        |
        v
    Timing middleware
        |
        v
      Router
        |
        v
    Route handler
        |
        v
  HTTP response
```

The request travels inward through the layers. When the route handler
returns, execution unwinds outward through the middleware layers.

This is why middleware is sometimes described as an **onion model**.

## 1.2 Why not put everything inside every handler?

Without middleware, every route might repeat logging, authentication,
timing, and error handling:

``` go
func getUsers(w http.ResponseWriter, r *http.Request) {
    // Logging code
    // Authentication code
    // Timing code
    // Actual get-users logic
}

func createUser(w http.ResponseWriter, r *http.Request) {
    // Same logging code
    // Same authentication code
    // Same timing code
    // Actual create-user logic
}
```

This creates duplication and makes changes difficult.

With middleware, shared behaviour is implemented once and reused:

``` go
handler := LoggingMiddleware(AuthenticationMiddleware(mux))
```

The route handlers can focus on their actual responsibilities.

------------------------------------------------------------------------

# 2. HTTP fundamentals required for middleware

Before understanding middleware, be comfortable with these concepts.

## 2.1 HTTP request

A client sends a request to a server. An HTTP request contains
information such as:

-   Method: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`, `OPTIONS`.
-   URL and path: for example, `/users/10`.
-   Headers: metadata such as `Authorization` and `Content-Type`.
-   Body: optional data, often JSON in APIs.

Example request:

``` http
POST /users HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Authorization: Bearer example-token

{"name":"Ankush"}
```

## 2.2 HTTP response

The server sends a response containing:

-   Status code: `200`, `201`, `400`, `401`, `404`, `429`, `500`, etc.
-   Headers: metadata about the response.
-   Body: the returned content.

Example:

``` http
HTTP/1.1 201 Created
Content-Type: application/json

{"id":10,"name":"Ankush"}
```

## 2.3 The two important handler parameters

Most HTTP handlers work with this signature:

``` go
func handler(w http.ResponseWriter, r *http.Request) {
    // Handle the request and write a response.
}
```

`w http.ResponseWriter` is used to construct the response.

`r *http.Request` contains the incoming request.

Example:

``` go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
}
```

The `Write` call writes response-body bytes. If no status has been
written yet, the first body write implicitly commits a `200 OK`
response.

------------------------------------------------------------------------

# 3. The `http.Handler` interface

This is the foundation of Go middleware.

The essential interface is:

``` go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

In other words, an `http.Handler` is anything that has a method with
this signature:

``` go
ServeHTTP(http.ResponseWriter, *http.Request)
```

Go uses **implicit interface implementation**. A type does not need to
declare that it implements `http.Handler`. It only needs to implement
the required method.

## 3.1 A custom handler type

``` go
package main

import (
    "fmt"
    "net/http"
)

type WelcomeHandler struct{}

func (h WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to my Go server")
}

func main() {
    handler := WelcomeHandler{}
    http.ListenAndServe(":8080", handler)
}
```

`WelcomeHandler` implements `ServeHTTP`, so it satisfies `http.Handler`.

The HTTP server can call `handler.ServeHTTP(w, r)` for each incoming
request.

## 3.2 Why middleware accepts `http.Handler`

A middleware often has this signature:

``` go
func LoggingMiddleware(next http.Handler) http.Handler
```

It accepts a handler because it needs something to call after performing
its own work.

It also returns a handler because the middleware itself must be usable
by the HTTP server or by another middleware.

This is the key design:

``` text
Handler in  -> Middleware -> Handler out
```

------------------------------------------------------------------------

# 4. `http.HandlerFunc` and `HandleFunc`

Go lets you write ordinary functions as HTTP handlers.

For example:

``` go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello")
}
```

But `http.Handle` expects an `http.Handler`, not an ordinary function
value.

`http.HandlerFunc` is an adapter that makes a compatible function
satisfy `http.Handler`.

Conceptually, its implementation looks like this:

``` go
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

The standard library's real implementation is equivalent in principle.

Use it like this:

``` go
handler := http.HandlerFunc(helloHandler)
```

Now `handler` implements `http.Handler`.

## 4.1 `Handle` versus `HandleFunc`

``` go
mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
```

`Handle` accepts an `http.Handler`.

``` go
mux.HandleFunc("GET /hello", helloHandler)
```

`HandleFunc` accepts a compatible function and adapts it for
registration.

Use `HandleFunc` for ordinary handler functions. Use `Handle` when you
already have a value implementing `http.Handler`, such as a
middleware-wrapped handler.

------------------------------------------------------------------------

# 5. Understanding `ServeMux`

`http.ServeMux` is Go's standard HTTP request multiplexer, commonly used
as a router.

It stores route registrations and dispatches an incoming request to the
matching handler.

``` go
mux := http.NewServeMux()

mux.HandleFunc("GET /hello", helloHandler)
mux.HandleFunc("POST /users", createUserHandler)
mux.HandleFunc("GET /users", listUsersHandler)
```

For Go versions supporting method-aware patterns (Go 1.22 and later),
patterns such as `"GET /hello"` match both a method and a path.

The mux is itself an `http.Handler`, because it implements `ServeHTTP`.

That means it can be passed into middleware:

``` go
handler := LoggingMiddleware(mux)
```

The router does not need to know that middleware is around it.

------------------------------------------------------------------------

# 6. Writing your first middleware

Here is the basic pattern:

``` go
package middleware

import (
    "log"
    "net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Before handler")

        next.ServeHTTP(w, r)

        log.Println("After handler")
    })
}
```

Let's break it down.

### Step 1: Accept the next handler

``` go
func LoggingMiddleware(next http.Handler) http.Handler
```

`next` is the handler that should run after this middleware.

It may be a router or another middleware-wrapped handler.

### Step 2: Return an HTTP handler

``` go
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Middleware logic
})
```

The anonymous function has the standard HTTP handler signature.
`http.HandlerFunc` adapts it to `http.Handler`.

### Step 3: Run code before the next handler

``` go
log.Println("Before handler")
```

This executes before the router or next middleware.

### Step 4: Call the next handler

``` go
next.ServeHTTP(w, r)
```

This transfers control to the next handler in the chain.

### Step 5: Run code after it returns

``` go
log.Println("After handler")
```

This executes after `next.ServeHTTP` returns normally.

If the downstream handler panics, normal statements after the call are
skipped while the panic unwinds. A deferred function can still run.

------------------------------------------------------------------------

# 7. Understanding `next.ServeHTTP`

Consider:

``` go
handler := LoggingMiddleware(mux)
```

The `mux` value is passed as the `next` argument.

When the server receives a request, it invokes the returned handler's
`ServeHTTP` method. The `HandlerFunc` adapter calls the anonymous
function. That function logs the request and then calls:

``` go
next.ServeHTTP(w, r)
```

Because `next` refers to `mux`, the router receives the request and
chooses the matching route.

Important: `next` is not automatically the route function. It is the
handler passed into the middleware. If that handler is a `ServeMux`, the
mux later selects the route handler.

## 7.1 Execution trace

``` text
HTTP server
    |
    v
LoggingMiddleware's returned handler
    |
    v
"Before handler"
    |
    v
next.ServeHTTP(w, r)
    |
    v
ServeMux.ServeHTTP
    |
    v
Matching route handler
    |
    v
Route writes response
    |
    v
ServeMux returns
    |
    v
Logging middleware resumes
    |
    v
"After handler"
```

------------------------------------------------------------------------

# 8. Middleware chaining and execution order

A middleware chain makes it easier to apply several middleware
functions.

First define a middleware type:

``` go
type Middleware func(http.Handler) http.Handler
```

This is a named function type. A function with the same signature can be
used as a `Middleware`.

## 8.1 A chain helper

``` go
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        handler = middlewares[i](handler)
    }

    return handler
}
```

Why iterate backwards?

Suppose the requested order is:

``` go
handler := Chain(
    mux,
    LoggingMiddleware,
    AuthenticationMiddleware,
    TimingMiddleware,
)
```

We want incoming requests to pass through Logging, then Authentication,
then Timing, and finally the mux.

To produce that nesting, wrapping happens in reverse:

``` text
First:  Timing(mux)
Next:   Authentication(Timing(mux))
Last:   Logging(Authentication(Timing(mux)))
```

The final handler is:

``` text
Logging(
    Authentication(
        Timing(
            mux,
        ),
    ),
)
```

The request execution order is:

``` text
Logging before
  Authentication before
    Timing before
      Route handler
    Timing after
  Authentication after
Logging after
```

This is the **onion model**.

## 8.2 Middleware can stop the chain

Authentication or rate limiting may reject a request:

``` go
if invalid {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
}

next.ServeHTTP(w, r)
```

If the middleware returns without calling `next.ServeHTTP`, downstream
middleware and the route handler do not run.

Outer middleware still resumes after the middleware returns.

------------------------------------------------------------------------

# 9. Logging middleware

A basic logging middleware records the method and path.

``` go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Incoming request: method=%s path=%s", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)
    })
}
```

Useful request information includes:

-   `r.Method`: HTTP method.
-   `r.URL.Path`: path without query parameters.
-   `r.URL.RawQuery`: encoded query string.
-   `r.RemoteAddr`: remote network address, often including a port.
-   `r.UserAgent()`: User-Agent header.
-   `r.Header`: request headers; be careful not to log secrets.

A production logger should generally record the response status and
duration too. That requires observing the response writer or using an
appropriate logging/metrics integration.

Do not log passwords, access tokens, session cookies, or sensitive
request bodies.

------------------------------------------------------------------------

# 10. Authentication middleware

Authentication answers:

> Who is making this request, and can their credentials be verified?

A simple learning example:

``` go
func AuthenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        if token != "secret123" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

The example token is deliberately simplistic. Do not use a hard-coded
shared token as production authentication.

Real applications may use:

-   Server-side sessions.
-   Signed access tokens.
-   OAuth 2.0 / OpenID Connect.
-   API keys for appropriate service-to-service use cases.

## 10.1 Authentication versus authorization

These are different concepts.

-   **Authentication:** verifies identity.
-   **Authorization:** checks whether that identity is allowed to
    perform a particular action.

For example, a user may be authenticated but not authorized to delete
another user's account.

A useful next step is to store verified identity in the request context
and let authorization logic inspect it.

------------------------------------------------------------------------

# 11. Timing middleware

Timing middleware measures how long downstream processing takes.

``` go
func TimingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        next.ServeHTTP(w, r)

        duration := time.Since(start)
        log.Printf("Request completed in %s", duration)
    })
}
```

`time.Now()` records the start time.

`time.Since(start)` computes the elapsed duration.

The measured duration includes downstream middleware and the route
handler. It does not necessarily represent the time until a remote
client has received and processed all response bytes.

For production metrics, record method, route template, status code, and
duration with bounded-cardinality labels. Avoid using arbitrary user IDs
or full URLs as metric labels.

------------------------------------------------------------------------

# 12. Understanding and wrapping `http.ResponseWriter`

`http.ResponseWriter` is an interface used to construct an HTTP
response.

Its essential methods are:

``` go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

### `Header()`

Returns the response header map.

``` go
w.Header().Set("Content-Type", "application/json")
```

Set headers before the response is committed.

### `WriteHeader(statusCode)`

Commits the response with a status code.

``` go
w.WriteHeader(http.StatusCreated)
```

Only the first status code is effective. If no explicit status is
written, the first body write implicitly sends `200 OK`.

### `Write(data)`

Writes response-body bytes.

``` go
w.Write([]byte("Hello"))
```

The returned `n` is the number of bytes actually written. It may be
smaller than `len(data)` if a write is incomplete.

## 12.1 Why wrap ResponseWriter?

A logging middleware may need to observe:

-   Which status code was written.
-   How many response bytes were written.
-   Whether the response has already been committed.

A basic wrapper:

``` go
type StatusRecorder struct {
    http.ResponseWriter
    StatusCode   int
    BytesWritten int
    Written      bool
}

func NewStatusRecorder(w http.ResponseWriter) *StatusRecorder {
    return &StatusRecorder{
        ResponseWriter: w,
        StatusCode:     http.StatusOK,
    }
}

func (rw *StatusRecorder) WriteHeader(code int) {
    if rw.Written {
        return
    }

    rw.StatusCode = code
    rw.Written = true
    rw.ResponseWriter.WriteHeader(code)
}

func (rw *StatusRecorder) Write(data []byte) (int, error) {
    if !rw.Written {
        rw.WriteHeader(http.StatusOK)
    }

    n, err := rw.ResponseWriter.Write(data)
    rw.BytesWritten += n

    return n, err
}

func (rw *StatusRecorder) Unwrap() http.ResponseWriter {
    return rw.ResponseWriter
}
```

Explanation:

-   Embedding `http.ResponseWriter` delegates methods not explicitly
    overridden to the underlying writer.
-   `StatusCode` records the first status.
-   `Written` tracks whether a response has been committed.
-   `BytesWritten` records actual bytes returned by the underlying
    `Write`.
-   `Unwrap` exposes the original writer to utilities that know how to
    unwrap it.

A logging middleware can use it:

``` go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        recorder := NewStatusRecorder(w)
        start := time.Now()

        next.ServeHTTP(recorder, r)

        log.Printf(
            "method=%s path=%s status=%d bytes=%d duration=%s",
            r.Method,
            r.URL.Path,
            recorder.StatusCode,
            recorder.BytesWritten,
            time.Since(start),
        )
    })
}
```

## 12.2 Important ResponseWriter caveat

The standard `ResponseWriter` has optional capabilities in addition to
its three core methods. Examples include:

-   `http.Flusher` for flushing buffered data.
-   `http.Hijacker` for taking over a connection in supported server
    contexts.
-   `http.Pusher` for HTTP/2 server push where supported.
-   `io.ReaderFrom` in some implementations.

A wrapper that only implements the three core methods may hide optional
interfaces from downstream handlers. That can break streaming or
connection-upgrade behaviour.

`Unwrap()` is useful, but it does **not** automatically make the wrapper
implement all optional interfaces. Production wrappers need deliberate
support for the capabilities they intend to preserve. Avoid writing a
generic wrapper casually; use a well-tested implementation or a
carefully designed wrapper for the application's requirements.

------------------------------------------------------------------------

# 13. Recovery middleware and panic handling

A panic is an abnormal control-flow event. It is different from an
ordinary returned error.

A deferred function runs when the surrounding function returns,
including while a panic unwinds the stack.

`recover()` can stop a panic only when called directly from a deferred
function in the same goroutine's panic flow.

``` go
func example() {
    defer func() {
        if value := recover(); value != nil {
            log.Println("Recovered:", value)
        }
    }()

    panic("unexpected problem")
}
```

After recovery, the panicking function does not resume at the statement
after `panic`. It returns from the function.

## 13.1 Recovery middleware

``` go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if value := recover(); value != nil {
                log.Printf("panic recovered: %v", value)

                http.Error(
                    w,
                    "Internal Server Error",
                    http.StatusInternalServerError,
                )
            }
        }()

        next.ServeHTTP(w, r)
    })
}
```

Place recovery outside the handlers and middleware whose panics it
should catch:

``` go
handler := Chain(
    mux,
    RecoveryMiddleware,
    LoggingMiddleware,
    AuthenticationMiddleware,
)
```

The recovery middleware wraps the other layers.

## 13.2 Important limitations

-   Recovery only catches panics in the same goroutine. It does not
    catch a panic in a separate goroutine started by a handler.
-   If the response has already been committed, the middleware cannot
    reliably replace it with a clean `500` response.
-   The standard `net/http` server already recovers panics from a
    handler's serving goroutine, logs a stack trace, and handles the
    connection. Custom recovery middleware is useful when the
    application needs consistent logging or controlled behaviour.
-   Avoid exposing panic details to clients. Log appropriate diagnostic
    information securely.

A production recovery middleware should check whether the response has
already been committed before attempting to write an error response.

------------------------------------------------------------------------

# 14. CORS middleware

CORS means **Cross-Origin Resource Sharing**.

A browser enforces the same-origin policy. An origin consists of:

-   Scheme, such as `http` or `https`.
-   Host, such as `localhost`.
-   Port, such as `3000` or `8080`.

These are different origins:

``` text
Frontend: http://localhost:3000
Backend:  http://localhost:8080
```

A browser may prevent frontend JavaScript from reading a cross-origin
response unless the backend supplies appropriate CORS response headers.

CORS is a browser-enforced access policy. It is not authentication and
does not stop non-browser clients from calling the server.

## 14.1 Basic CORS middleware

``` go
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set(
            "Access-Control-Allow-Origin",
            "http://localhost:3000",
        )

        w.Header().Set(
            "Access-Control-Allow-Methods",
            "GET, POST, PUT, PATCH, DELETE, OPTIONS",
        )

        w.Header().Set(
            "Access-Control-Allow-Headers",
            "Content-Type, Authorization",
        )

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

The origin is explicitly restricted to the local frontend.

## 14.2 What is a preflight request?

For certain cross-origin requests, the browser first sends an `OPTIONS`
request to ask whether the intended method and headers are allowed.

Example:

``` http
OPTIONS /users HTTP/1.1
Origin: http://localhost:3000
Access-Control-Request-Method: POST
Access-Control-Request-Headers: authorization,content-type
```

A successful response might contain:

``` http
HTTP/1.1 204 No Content
Access-Control-Allow-Origin: http://localhost:3000
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Headers: Authorization, Content-Type
```

If the browser accepts the preflight response, it can proceed with the
actual request.

Not every cross-origin request triggers preflight. The browser's CORS
rules determine when it is required.

## 14.3 Production CORS considerations

-   Validate the incoming `Origin` against an explicit allowlist.
-   Do not reflect arbitrary origins.
-   If dynamically selecting an allowed origin, use `Vary: Origin`.
-   If cookies or other credentials are allowed, configure
    `Access-Control-Allow-Credentials: true` and do not use `*` as the
    allowed origin.
-   Handle preflight before authentication middleware that expects
    application credentials. A browser preflight generally does not
    contain the application's normal authorization token.
-   CORS is not a substitute for CSRF protection when cookie-based
    authentication is used.

------------------------------------------------------------------------

# 15. Rate limiting fundamentals

Rate limiting controls how frequently a client can make requests.

Example policy:

``` text
Allow at most 5 requests per minute per client IP.
```

If a client exceeds the limit, the server commonly responds with:

``` http
HTTP/1.1 429 Too Many Requests
```

Rate limiting can protect APIs from accidental overuse, abusive clients,
and some forms of traffic spikes. It is not a complete DDoS defence.

A limiter needs to define:

1.  What identifies a client? IP, API key, user ID, or another identity?
2.  What is the allowed request rate?
3.  What happens when the limit is exceeded?
4.  When does capacity become available again?
5.  Is the state local to one process or shared across servers?

For a learning implementation, an in-memory map is useful. In a
multi-instance production deployment, a shared store or gateway-level
limiter is usually needed.

------------------------------------------------------------------------

# 16. Fixed Window rate limiting

A Fixed Window limiter divides time into fixed intervals.

Example:

``` text
Window: 60 seconds
Limit: 5 requests
```

Each client gets a counter for the current window. When the window
expires, the counter resets.

## 16.1 Data model

``` go
type Client struct {
    Count     int
    WindowEnd time.Time
}
```

A map can associate each client key with its state:

``` go
map[string]Client
```

## 16.2 Implementation

``` go
type FixedWindowLimiter struct {
    mu      sync.Mutex
    clients map[string]Client
    limit   int
    window  time.Duration
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
    return &FixedWindowLimiter{
        clients: make(map[string]Client),
        limit:   limit,
        window:  window,
    }
}

func (rl *FixedWindowLimiter) Allow(key string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now := time.Now()

    client, exists := rl.clients[key]

    if !exists || !now.Before(client.WindowEnd) {
        client = Client{
            Count:     0,
            WindowEnd: now.Add(rl.window),
        }
    }

    client.Count++
    rl.clients[key] = client

    return client.Count <= rl.limit
}
```

Explanation:

1.  Lock the shared map.
2.  Read the client's state.
3.  If the client is new or its window expired, start a new window.
4.  Increment the count.
5.  Save the updated client.
6.  Allow only if the count is within the limit.
7.  Unlock when the method returns.

The `defer` ensures the mutex is unlocked on every normal return path.

## 16.3 Fixed Window boundary problem

Suppose the limit is five requests per minute.

A client sends five requests just before a window ends and five more
just after the next window begins.

The client may send ten requests within a very short interval while
technically staying within each separate window's limit.

This is called the **boundary burst problem**.

Fixed Window is simple, but it does not smooth traffic across
boundaries.

------------------------------------------------------------------------

# 17. Token Bucket rate limiting

Token Bucket allows controlled bursts while replenishing capacity over
time.

Imagine a bucket with tokens:

-   Bucket capacity: 5 tokens.
-   Each request consumes 1 token.
-   Refill rate: 0.2 tokens per second.
-   A new client starts with a full bucket.

If no requests arrive, tokens refill until the bucket reaches capacity.
Tokens above capacity are discarded.

## 17.1 Data model

``` go
type Bucket struct {
    Tokens     float64
    LastRefill time.Time
}
```

Limiter configuration:

``` go
type TokenBucketLimiter struct {
    mu         sync.Mutex
    clients    map[string]*Bucket
    capacity   float64
    refillRate float64
}
```

## 17.2 Refill mathematics

Let:

-   `elapsed` = time since the last refill calculation, in seconds.
-   `refillRate` = tokens added per second.
-   `capacity` = maximum token count.

Then:

``` text
new tokens = old tokens + elapsed seconds × refill rate
```

But the bucket cannot exceed capacity:

``` text
new tokens = min(capacity, new tokens)
```

Example:

``` text
Old tokens: 2
Elapsed time: 4 seconds
Refill rate: 0.5 tokens/second

Refilled: 4 × 0.5 = 2
New balance: 2 + 2 = 4 tokens
```

## 17.3 Implementation

``` go
type TokenBucketLimiter struct {
    mu         sync.Mutex
    clients    map[string]*Bucket
    capacity   float64
    refillRate float64
}

func NewTokenBucketLimiter(capacity, refillRate float64) *TokenBucketLimiter {
    return &TokenBucketLimiter{
        clients:    make(map[string]*Bucket),
        capacity:   capacity,
        refillRate: refillRate,
    }
}

func (rl *TokenBucketLimiter) Allow(key string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now := time.Now()

    bucket, exists := rl.clients[key]

    if !exists {
        bucket = &Bucket{
            Tokens:     rl.capacity,
            LastRefill: now,
        }
        rl.clients[key] = bucket
    }

    elapsed := now.Sub(bucket.LastRefill).Seconds()
    if elapsed < 0 {
        elapsed = 0
    }

    bucket.Tokens += elapsed * rl.refillRate

    if bucket.Tokens > rl.capacity {
        bucket.Tokens = rl.capacity
    }

    bucket.LastRefill = now

    if bucket.Tokens >= 1 {
        bucket.Tokens--
        return true
    }

    return false
}
```

This is a basic in-memory implementation for learning. A production
version should validate configuration, manage client-state cleanup, and
consider lock contention and distributed deployment.

## 17.4 Middleware adapter

The limiter's `Allow` method can be wrapped in ordinary HTTP middleware:

``` go
func (rl *TokenBucketLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        key := r.RemoteAddr

        if !rl.Allow(key) {
            w.Header().Set("Retry-After", "1")
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

For real deployments, `RemoteAddr` includes a port and may identify a
proxy rather than the original client. Client identity extraction must
be designed around the network topology and trusted proxies.

------------------------------------------------------------------------

# 18. Testing with `net/http/httptest`

Testing is essential. Do not rely only on starting the server and
manually checking it in a browser.

The standard library package is:

``` go
import "net/http/httptest"
```

Two especially useful functions:

-   `httptest.NewRequest`: creates a request for testing.
-   `httptest.NewRecorder`: creates a response recorder.

The recorder captures the response in memory.

## 18.1 Test a basic handler

Production handler:

``` go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
}
```

Test:

``` go
func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/hello", nil)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(helloHandler)
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", rr.Code)
    }

    if rr.Body.String() != "Hello World" {
        t.Errorf("unexpected body: %q", rr.Body.String())
    }
}
```

Execution:

1.  Create a fake request.
2.  Create a response recorder.
3.  Invoke the handler directly.
4.  Inspect the recorded status and body.

No TCP listener or real browser is needed.

## 18.2 Test authentication middleware

Valid token test:

``` go
func TestAuthenticationMiddleware_ValidToken(t *testing.T) {
    next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Welcome"))
    })

    handler := AuthenticationMiddleware(next)

    req := httptest.NewRequest(http.MethodGet, "/private", nil)
    req.Header.Set("Authorization", "secret123")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", rr.Code)
    }

    if rr.Body.String() != "Welcome" {
        t.Fatalf("expected Welcome, got %q", rr.Body.String())
    }
}
```

Invalid token test:

``` go
func TestAuthenticationMiddleware_InvalidToken(t *testing.T) {
    next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        t.Error("next handler must not run for an invalid token")
    })

    handler := AuthenticationMiddleware(next)

    req := httptest.NewRequest(http.MethodGet, "/private", nil)
    req.Header.Set("Authorization", "wrong-token")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusUnauthorized {
        t.Fatalf("expected 401, got %d", rr.Code)
    }
}
```

The test handler deliberately calls `t.Error` if it runs. This verifies
that the middleware stops the chain when credentials are invalid.

## 18.3 Test a rate limiter

``` go
func TestTokenBucketLimiter(t *testing.T) {
    limiter := NewTokenBucketLimiter(2, 0)

    if !limiter.Allow("client-a") {
        t.Fatal("first request should be allowed")
    }

    if !limiter.Allow("client-a") {
        t.Fatal("second request should be allowed")
    }

    if limiter.Allow("client-a") {
        t.Fatal("third request should be rejected")
    }

    if !limiter.Allow("client-b") {
        t.Fatal("a different client should have its own bucket")
    }
}
```

This test uses a zero refill rate so the result is deterministic and
does not depend on waiting for time to pass.

## 18.4 Run tests

``` bash
go test ./...
```

Verbose output:

``` bash
go test -v ./...
```

Run tests with the race detector:

``` bash
go test -race ./...
```

------------------------------------------------------------------------

# 19. Concurrency, mutexes, and race detection

Go's HTTP server can process multiple requests concurrently. Handlers
must be safe when multiple goroutines access shared state.

A Go map is not safe for concurrent writes without synchronization.

This is unsafe:

``` go
var requests = make(map[string]int)

func handler(w http.ResponseWriter, r *http.Request) {
    requests[r.RemoteAddr]++
}
```

Two goroutines may update the map simultaneously, causing a data race or
concurrent map access failure.

## 19.1 Protect shared state with a mutex

``` go
var (
    mu       sync.Mutex
    requests = make(map[string]int)
)

func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    requests[r.RemoteAddr]++
    count := requests[r.RemoteAddr]
    mu.Unlock()

    fmt.Fprintln(w, count)
}
```

`Lock` enters the critical section. `Unlock` leaves it.

Prefer `defer` for methods with multiple return paths:

``` go
func increment(key string) int {
    mu.Lock()
    defer mu.Unlock()

    requests[key]++
    return requests[key]
}
```

Keep critical sections short. Do not hold a global mutex while
performing slow network calls or writing a large response.

## 19.2 Race detector

Run:

``` bash
go test -race ./...
```

The race detector reports data races that occur during the test
execution.

It cannot prove that the program is race-free for every possible
execution. Tests must exercise relevant concurrent code paths.

## 19.3 Concurrent limiter test

``` go
func TestTokenBucketLimiterConcurrent(t *testing.T) {
    limiter := NewTokenBucketLimiter(100, 0)

    var wg sync.WaitGroup

    for i := 0; i < 200; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            limiter.Allow("same-client")
        }()
    }

    wg.Wait()
}
```

This test exercises concurrent calls against the same client bucket. Run
it with `go test -race ./...`.

For stronger tests, also count successful requests using an atomic
counter and verify that exactly the expected number were accepted.

------------------------------------------------------------------------

# 20. Middleware benchmarking

Testing checks correctness. Benchmarking measures performance.

Go benchmark functions begin with `Benchmark` and accept `*testing.B`.

``` go
func BenchmarkTokenBucketAllow(b *testing.B) {
    limiter := NewTokenBucketLimiter(100000, 0)

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        limiter.Allow("benchmark-client")
    }
}
```

Run:

``` bash
go test -bench=. ./...
```

For memory allocation information:

``` bash
go test -bench=. -benchmem ./...
```

Benchmarking reminders:

-   Use repeatable inputs.
-   Avoid logging inside the measured loop.
-   Understand whether the benchmark measures a hot client or many
    clients.
-   Run benchmarks more than once before drawing conclusions.
-   Optimize only after measuring a real bottleneck.

------------------------------------------------------------------------

# 21. Production design considerations

A learning middleware is not automatically production-ready. Consider
these points before deploying.

## 21.1 Validate configuration

For a token bucket:

-   Capacity should be greater than zero.
-   Refill rate should not be negative.
-   Extremely large values should be considered carefully.

## 21.2 Clean up client state

In-memory per-IP maps can grow indefinitely as new client keys arrive.

Possible approaches:

-   Periodic cleanup of expired entries.
-   Bounded caches.
-   Shared storage with TTL.
-   A gateway or dedicated rate-limiting service.

Cleanup itself must be concurrency-safe.

## 21.3 Multiple application instances

If two application instances each keep their own in-memory rate-limit
map, a client may effectively receive a separate quota on each instance.

For shared limits across instances, use shared state, often Redis or an
API gateway.

## 21.4 Client IP and reverse proxies

`r.RemoteAddr` may represent a reverse proxy or load balancer.

Headers such as `X-Forwarded-For` and `X-Real-IP` can carry
client-address information, but clients can forge them unless a trusted
proxy overwrites and controls them.

Only trust forwarding headers from explicitly configured trusted
proxies.

## 21.5 Security and observability

-   Do not log credentials or secrets.
-   Avoid returning internal panic details to clients.
-   Use structured logs where appropriate.
-   Record request IDs to correlate logs.
-   Use bounded metric labels.
-   Apply authentication and authorization at the correct layer.
-   Make CORS origin policy explicit.
-   Decide which routes should be excluded from or included in rate
    limiting.

------------------------------------------------------------------------

# 22. Recommended project structure

A small learning project could use:

``` text
middleware-project/
│
├── go.mod
├── main.go
│
├── middleware/
│   ├── chain.go
│   ├── logging.go
│   ├── authentication.go
│   ├── timing.go
│   ├── response_writer.go
│   ├── recovery.go
│   ├── cors.go
│   ├── ratelimit_fixed_window.go
│   ├── ratelimit_token_bucket.go
│   ├── request_id.go
│   │
│   ├── chain_test.go
│   ├── authentication_test.go
│   ├── cors_test.go
│   ├── recovery_test.go
│   ├── ratelimit_test.go
│   └── benchmark_test.go
│
└── README.md
```

This is a suggested learning structure, not a requirement. Keep package
names and import paths consistent with the module path declared in
`go.mod`.

## 22.1 Example `main.go`

``` go
package main

import (
    "log"
    "net/http"
    "time"

    "middleware-project/middleware"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })

    limiter := middleware.NewTokenBucketLimiter(5, 0.2)

    handler := middleware.Chain(
        mux,
        middleware.RecoveryMiddleware,
        middleware.CORSMiddleware,
        middleware.LoggingMiddleware,
        limiter.Middleware,
    )

    log.Println("Server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
```

The import path `"middleware-project/middleware"` is only correct if
`go.mod` declares `module middleware-project`. Replace it with your
actual module path if different.

The example assumes the functions and types shown in the earlier
sections exist in the `middleware` package.

------------------------------------------------------------------------

# 23. Common mistakes and debugging checklist

## Mistake 1: Forgetting `next.ServeHTTP`

If a middleware never calls `next.ServeHTTP` and does not intentionally
write a response and return, the request may appear to hang or produce
an empty response.

## Mistake 2: Calling `next.ServeHTTP` twice

Normally, call the next handler once. Calling it twice can execute route
logic twice and create duplicate side effects.

## Mistake 3: Wrong middleware order

If authentication blocks an OPTIONS preflight request, CORS may fail.
Place preflight handling where it can run before application
authentication.

## Mistake 4: Writing headers too late

Set response headers before the response is committed. Once
`WriteHeader` or the first body write commits the response, changing
headers generally has no effect.

## Mistake 5: Assuming `recover` catches every panic

It does not catch panics in other goroutines. A goroutine that starts
its own work needs its own deliberate panic-handling strategy where
appropriate.

## Mistake 6: Assuming a mutex makes all code safe

A mutex only protects data consistently when every access to that shared
data follows the same synchronization strategy.

## Mistake 7: Treating a demo rate limiter as production-ready

Consider client identity, cleanup, multiple instances, trusted proxies,
lock contention, and configuration validation.

## Debugging checklist

-   Does the middleware return an `http.Handler`?
-   Is `next` the handler you expect?
-   Is `next.ServeHTTP(w, r)` called at the right point?
-   Does an early `return` intentionally stop the chain?
-   Are response headers set before the response is committed?
-   Is shared mutable state protected?
-   Does `go test -race ./...` pass?
-   Are tests covering both success and rejection paths?
-   Does the middleware preserve the response-writer capabilities
    required by the application?

------------------------------------------------------------------------

# 24. What to learn next in middleware

You have covered the main middleware foundations: handler interfaces,
chaining, logging, authentication, timing, response-writer wrapping,
recovery, CORS, rate limiting, and introductory testing.

The **recommended next topic is Request ID Middleware and
`context.Context`**.

Why this next?

Every incoming request can be assigned a unique ID. That ID can travel
through the handler chain and be used by downstream services and logs.
This introduces an important Go backend concept: request-scoped context.

## 24.1 What problem does a Request ID solve?

Imagine your server receives three requests at nearly the same time:

``` text
Request A -> /users
Request B -> /orders
Request C -> /payments
```

Their logs may interleave:

``` text
GET /users started
GET /orders started
GET /payments started
GET /users completed
GET /payments completed
GET /orders completed
```

It becomes difficult to tell which log lines belong to which request.

Assign each request a unique ID:

``` text
Request A -> ID: req-a12
Request B -> ID: req-b34
Request C -> ID: req-c56
```

Then logs become easier to correlate:

``` text
[req-a12] GET /users started
[req-b34] GET /orders started
[req-c56] GET /payments started
[req-a12] GET /users completed
[req-c56] GET /payments completed
[req-b34] GET /orders completed
```

## 24.2 What is `context.Context`?

`context.Context` carries request-scoped values, cancellation signals,
and deadlines across API boundaries.

An HTTP request already has a context:

``` go
ctx := r.Context()
```

To create a derived context with a value:

``` go
ctx := context.WithValue(
    r.Context(),
    key,
    requestID,
)
```

Then attach it to a derived request:

``` go
r = r.WithContext(ctx)
```

Downstream handlers can access that context through:

``` go
requestID := r.Context().Value(key)
```

Use a private custom key type rather than a plain string key to avoid
collisions.

## 24.3 Request ID middleware --- learning implementation

This example uses a simple generated ID for demonstration. In a real
application, use a robust UUID or random-ID generator, or a carefully
validated trusted upstream request ID.

``` go
package middleware

import (
    "context"
    "fmt"
    "net/http"
    "sync/atomic"
)

type requestIDKey struct{}

var requestCounter atomic.Uint64

func RequestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := fmt.Sprintf("req-%d", requestCounter.Add(1))

        ctx := context.WithValue(
            r.Context(),
            requestIDKey{},
            id,
        )

        r = r.WithContext(ctx)

        w.Header().Set("X-Request-ID", id)

        next.ServeHTTP(w, r)
    })
}

func RequestIDFromContext(ctx context.Context) (string, bool) {
    id, ok := ctx.Value(requestIDKey{}).(string)
    return id, ok
}
```

Explanation:

1.  `atomic.Uint64` safely increments a counter across concurrent
    requests.
2.  A private `requestIDKey` type is used as the context key.
3.  `context.WithValue` creates a derived context containing the ID.
4.  `r.WithContext(ctx)` returns a shallow copy of the request with the
    new context.
5.  The response includes `X-Request-ID`.
6.  Downstream handlers can retrieve the ID from the context.

This counter is suitable only as a simple teaching example. It is not a
globally unique ID across server restarts or multiple instances. Use
UUIDs or cryptographically random identifiers when uniqueness across
processes matters.

## 24.4 How to use it

Add it to your middleware chain:

``` go
handler := middleware.Chain(
    mux,
    middleware.RecoveryMiddleware,
    middleware.RequestIDMiddleware,
    middleware.LoggingMiddleware,
)
```

A downstream handler can retrieve the ID:

``` go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    id, ok := middleware.RequestIDFromContext(r.Context())

    if !ok {
        http.Error(w, "Request ID missing", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Hello! Request ID: %s", id)
}
```

## 24.5 What to study after Request ID Middleware

Follow this sequence:

1.  **Request ID middleware** --- generate and attach an ID.
2.  **`context.Context` fundamentals** --- values, cancellation,
    deadlines, and propagation.
3.  **Context-aware database calls** --- use methods such as
    `QueryContext` and `ExecContext`.
4.  **Timeout middleware** --- understand request deadlines and
    cancellation.
5.  **Authorization middleware** --- store authenticated identity in
    context and enforce permissions.
6.  **Trusted proxy handling** --- safely determine client IP behind
    reverse proxies.
7.  **Redis-based rate limiting** --- share rate-limit state across
    multiple server instances.
8.  **Advanced ResponseWriter wrappers** --- preserve required optional
    interfaces.
9.  **Middleware testing and benchmarks** --- test order, cancellation,
    concurrency, and performance.

Do not treat all of these as prerequisites for writing ordinary CRUD
APIs. Learn the request ID and context fundamentals next, then continue
with backend development and revisit advanced middleware as your
projects need it.

------------------------------------------------------------------------

# 25. Practice assignments and completion checklist

## Assignment A --- Handler fundamentals

-   [ ] Create a custom type implementing `http.Handler`.
-   [ ] Convert an ordinary handler function using `http.HandlerFunc`.
-   [ ] Register routes using `ServeMux`.
-   [ ] Explain what `next` refers to in a middleware function.

## Assignment B --- Middleware implementation

-   [ ] Implement logging middleware.
-   [ ] Implement authentication middleware.
-   [ ] Implement timing middleware.
-   [ ] Implement recovery middleware.
-   [ ] Implement CORS middleware.
-   [ ] Implement a middleware chain helper.
-   [ ] Draw the execution order for three chained middleware functions.

## Assignment C --- Rate limiting

-   [ ] Implement a Fixed Window limiter.
-   [ ] Explain the boundary burst problem.
-   [ ] Implement a Token Bucket limiter.
-   [ ] Explain capacity and refill rate.
-   [ ] Protect shared limiter state with synchronization.
-   [ ] Think through client-key selection and expired-state cleanup.

## Assignment D --- Testing and concurrency

-   [ ] Test a handler with `httptest.NewRequest`.
-   [ ] Capture its response using `httptest.NewRecorder`.
-   [ ] Test valid and invalid authentication.
-   [ ] Test CORS preflight behaviour.
-   [ ] Test rate-limit allowed and rejected requests.
-   [ ] Add concurrent tests.
-   [ ] Run `go test ./...`.
-   [ ] Run `go test -race ./...`.
-   [ ] Write at least one benchmark.

## Assignment E --- Next topic

-   [ ] Learn `context.Context`.
-   [ ] Implement Request ID middleware.
-   [ ] Retrieve the request ID in a downstream handler.
-   [ ] Include the request ID in logs.
-   [ ] Test that every request receives an ID.

## When can you consider middleware fundamentals complete?

You should be able to do these tasks without copying a tutorial:

1.  Explain why `http.Handler` is an interface.
2.  Explain how `http.HandlerFunc` adapts a function.
3.  Write middleware with the signature
    `func(http.Handler) http.Handler`.
4.  Explain request entry order and response unwind order.
5.  Stop a middleware chain intentionally.
6.  Implement basic logging, authentication, recovery, and CORS
    middleware.
7.  Explain the trade-offs of Fixed Window and Token Bucket rate
    limiting.
8.  Protect shared state from concurrent access.
9.  Test middleware without starting a real HTTP server.
10. Debug a middleware ordering or response-writer issue.

------------------------------------------------------------------------

## Final learning advice

Do not measure progress by the number of middleware functions you have
copied. Measure it by whether you can explain the execution flow,
implement a new middleware from the handler interface, test its
behaviour, and reason about concurrency and failure cases.

**Your next lesson: `context.Context` and Request ID Middleware in Go.**

Once you understand request-scoped context, you will be better prepared
to build authentication, authorization, request cancellation, timeouts,
database operations, and traceable backend services.
