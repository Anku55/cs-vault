
# Go Backend Learning Notes
## HTTP in Go — From Fundamentals

---

# 1. What is HTTP?

**HTTP = HyperText Transfer Protocol.**

It is the protocol used for communication between a **client** and a **server**.

Typical flow:

```text
Client
  │
  │ HTTP Request
  ▼
Server
  │
  │ HTTP Response
  ▼
Client
```

Example:

When you open:

```text
http://localhost:8080/users
```

your browser sends an HTTP request to your Go server.

The Go server processes it and sends an HTTP response.

### Important terms to remember

| Concept | Meaning |
|---|---|
| Client | Sends request |
| Server | Receives request and sends response |
| Request | Message from client → server |
| Response | Message from server → client |
| HTTP | Protocol used for communication |
| URL | Address of the resource |
| Method | What operation the client wants |
| Header | Metadata about request/response |
| Body | Actual data being sent |
| Status Code | Result of the request |

---

# 2. HTTP Request

An HTTP request contains several important parts:

```text
HTTP Request
│
├── Method
├── URL
├── Headers
└── Body
```

Example:

```http
GET /users HTTP/1.1
Host: localhost:8080
Accept: application/json
```

### Method

The method tells the server **what the client wants to do**.

Common methods:

```text
GET
POST
PUT
PATCH
DELETE
```

For example:

```http
GET /users
```

means:

> Give me the users.

While:

```http
POST /users
```

means:

> Create a new user.

---

# 3. HTTP Response

The server sends a response back:

```text
HTTP Response
│
├── Status Code
├── Headers
└── Body
```

Example:

```http
HTTP/1.1 200 OK
Content-Type: text/plain

Hello from Go!
```

Important status codes:

| Code | Meaning |
|---:|---|
| `200` | OK |
| `201` | Created |
| `204` | No Content |
| `400` | Bad Request |
| `401` | Unauthorized |
| `403` | Forbidden |
| `404` | Not Found |
| `500` | Internal Server Error |

You'll use these heavily when building APIs.

---

# 4. Go's HTTP Package

Go's standard library provides HTTP functionality through:

```go
"net/http"
```

Package name:

```text
net/http
```

This is one of the most important Go packages for backend development.

You'll frequently encounter:

```go
import "net/http"
```

### Important `net/http` functions/types to remember

```text
http.ListenAndServe
http.HandleFunc
http.Handle
http.Get
http.Post
http.NewRequest
http.Client
http.Request
http.Response
http.ResponseWriter
```

We'll study each of these properly.

---

# 5. Your First Go HTTP Server

Basic server:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)

    fmt.Println("Server running on :8080")

    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}
```

Let's break this down.

---

# 6. `package main`

```go
package main
```

`main` is a special Go package.

An executable Go program normally uses:

```go
package main
```

and contains:

```go
func main()
```

The program starts executing from:

```go
func main()
```

### Remember

**Package:**

```text
package main
```

**Entry-point function:**

```text
main()
```

---

# 7. `import`

```go
import (
    "fmt"
    "net/http"
)
```

`import` allows your package to use functionality from another package.

Here:

```go
"fmt"
```

provides formatting functionality.

And:

```go
"net/http"
```

provides HTTP functionality.

---

# 8. `http.ListenAndServe`

### Function name

```go
http.ListenAndServe
```

This starts an HTTP server and makes it **listen for incoming connections**.

Example:

```go
http.ListenAndServe(":8080", nil)
```

Think:

```text
ListenAndServe
       │
       ├── Listen
       │
       └── Serve
```

It means:

> Listen on port `8080` and handle incoming HTTP requests.

---

# 9. Port

In:

```go
http.ListenAndServe(":8080", nil)
```

the:

```text
8080
```

is the **port**.

Your server becomes accessible at:

```text
http://localhost:8080
```

Mental model:

```text
localhost
    │
    └── 8080
          │
          └── Your Go HTTP server
```

A machine can run multiple network services using different ports.

Example:

```text
Go API       → 8080
Frontend     → 3000
Database     → 5432
```

---

# 10. `http.HandleFunc`

### Function name

```go
http.HandleFunc
```

This registers a function to handle requests for a particular URL pattern.

Example:

```go
http.HandleFunc("/", homeHandler)
```

Meaning:

> When a request comes to `/`, use `homeHandler`.

Another example:

```go
http.HandleFunc("/users", usersHandler)
```

Now:

```text
GET /users
```

can be handled by:

```go
usersHandler
```

Mental model:

```text
URL path
   ↓
Router
   ↓
Handler function
   ↓
Response
```

---

# 11. Handler Function

A handler is a function that handles an HTTP request.

Example:

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}
```

The important function name here is:

```text
homeHandler
```

But more importantly, notice its parameters:

```go
w http.ResponseWriter
r *http.Request
```

These two are extremely important in Go HTTP programming.

---

# 12. `http.Request`

### Type

```go
http.Request
```

Usually you'll see:

```go
r *http.Request
```

`r` represents the **incoming HTTP request**.

It contains information sent by the client.

For example:

```text
Method
URL
Headers
Body
Cookies
Query parameters
```

You can think:

```text
r = What did the client send me?
```

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
}
```

If the browser sends:

```http
GET /users
```

then:

```go
r.Method
```

will contain:

```text
GET
```

---

# 13. `http.ResponseWriter`

### Interface

```go
http.ResponseWriter
```

Usually you'll see:

```go
w http.ResponseWriter
```

It represents the **response being sent back to the client**.

Think:

```text
r → incoming request
w → outgoing response
```

This is one of the most important mental models in Go HTTP:

```text
Client
  │
  │ request
  ▼
r *http.Request
  │
  │
  ▼
Handler
  │
  │
  ▼
w http.ResponseWriter
  │
  │ response
  ▼
Client
```

---

# 14. `fmt.Fprintln`

### Function name

```go
fmt.Fprintln
```

`Fprintln` writes formatted text to a destination and adds a newline.

Example:

```go
fmt.Fprintln(w, "Hello from Go!")
```

Here:

```go
w
```

is the destination.

Because `w` represents the HTTP response, this sends:

```text
Hello from Go!
```

to the client.

### Compare

```go
fmt.Println("Hello")
```

prints to the terminal.

While:

```go
fmt.Fprintln(w, "Hello")
```

writes to the HTTP response.

Mental model:

```text
fmt.Println
     ↓
Terminal

fmt.Fprintln(w, ...)
     ↓
HTTP Response
     ↓
Browser / API client
```

---

# 15. `io.Writer`

This is an important Go concept behind `fmt.Fprintln`.

`fmt.Fprintln` accepts a destination implementing:

```go
io.Writer
```

`http.ResponseWriter` provides the necessary writing behavior.

So:

```go
fmt.Fprintln(w, "Hello")
```

works because `w` can be written to.

Don't worry about mastering interfaces yet, but remember:

```text
io.Writer
    ↓
Something you can write data into
```

You'll encounter this concept constantly in Go backend development.

---

# 16. HTTP Handler Mental Model

Whenever you see:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    
}
```

think:

```text
r
│
│ "What did the client request?"
│
▼
handler
│
│ "How should I respond?"
│
▼
w
│
│ response
▼
Client
```

This mental model is extremely important.

---

# 17. Checking the HTTP Method

You can inspect the request method using:

```go
r.Method
```

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        fmt.Fprintln(w, "GET request")
    }
}
```

A better practice is to use the standard constants:

```go
http.MethodGet
```

Example:

```go
if r.Method == http.MethodGet {
    fmt.Fprintln(w, "GET request")
}
```

Other useful constants:

```go
http.MethodGet
http.MethodPost
http.MethodPut
http.MethodPatch
http.MethodDelete
```

These are worth remembering.

---

# 18. Simple API Example

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/products", productsHandler)

    fmt.Println("Server running on :8080")

    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Home")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Users")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Products")
}
```

Routes:

```text
GET /
    ↓
homeHandler

GET /users
    ↓
usersHandler

GET /products
    ↓
productsHandler
```

---

# 19. Important Functions / Types Learned So Far

This is the **reference list** you can return to later.

### Package

```text
net/http
```

### Server

```text
http.ListenAndServe
```

### Routing

```text
http.HandleFunc
http.Handle
```

### Request

```text
http.Request
```

### Response

```text
http.ResponseWriter
```

### HTTP methods

```text
http.MethodGet
http.MethodPost
http.MethodPut
http.MethodPatch
http.MethodDelete
```

### Formatting/output

```text
fmt.Println
fmt.Fprintln
fmt.Fprint
fmt.Fprintf
```

### Go interface related to writing

```text
io.Writer
```

---

# 20. Current Learning Map

We'll build HTTP in Go progressively like this:

```text
HTTP Fundamentals
       ↓
Request / Response
       ↓
net/http
       ↓
Server
       ↓
http.ListenAndServe
       ↓
http.HandleFunc
       ↓
Handler Functions
       ↓
http.Request
       ↓
http.ResponseWriter
       ↓
HTTP Methods
       ↓
URL
       ↓
Path Parameters
       ↓
Query Parameters
       ↓
Request Headers
       ↓
Request Body
       ↓
JSON
       ↓
POST APIs
       ↓
PUT / PATCH
       ↓
DELETE
       ↓
Status Codes
       ↓
Error Handling
       ↓
Routing Design
       ↓
Middleware
       ↓
Context
       ↓
Authentication
       ↓
Database
       ↓
Production API
```

---

# 🧠 Most Important Mental Models So Far

### HTTP

```text
Client → Request → Server
Client ← Response ← Server
```

### Go HTTP handler

```text
r = incoming request

w = outgoing response
```

### Routing

```text
URL
 ↓
Route
 ↓
Handler
 ↓
Response
```

### Server

```text
http.ListenAndServe(":8080", ...)
```

means:

> Start an HTTP server and listen for requests on port 8080.

### Sending a response

```go
fmt.Fprintln(w, "Hello")
```

means:

> Write `"Hello"` into the HTTP response.

---

