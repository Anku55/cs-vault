# Go package vs module

## 1. Go Project Structure

A typical Go project is organized into:

```text
Project
│
├── go.mod
├── main.go
└── packages/
    ├── file1.go
    └── file2.go
```

The three important concepts are:

```text
File → Package → Module
```

---

# 2. Package in Go

A **package** is a group of related Go source files that belong together.

Every `.go` file starts with a package declaration:

```go
package main
```

or:

```go
package users
```

### Example

```text
calculator/
├── add.go
└── subtract.go
```

`add.go`:

```go
package calculator

func Add(a int, b int) int {
    return a + b
}
```

`subtract.go`:

```go
package calculator

func Subtract(a int, b int) int {
    return a - b
}
```

Both files belong to the same:

```text
calculator package
```

### Important rule

> **Files in the same directory normally belong to the same package.**

You cannot normally have:

```text
Arrays/
├── main.go       → package main
└── main2.go      → package arrays
```

This produces:

```text
found packages main and arrays
```

Instead, either make both:

```go
package main
```

or put them in separate directories.

---

# 3. `package main`

`main` is a special package.

An executable Go program uses:

```go
package main
```

and contains:

```go
func main() {
    // program starts here
}
```

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go")
}
```

When you run:

```bash
go run .
```

Go starts execution from:

```go
main()
```

Mental model:

```text
package main
      ↓
   main()
      ↓
Program starts
```

---

# 4. Import

`import` allows one package to use code from another package.

Example:

```go
import "fmt"
```

Then:

```go
fmt.Println("Hello")
```

Here:

```text
fmt     → package
Println → function
```

Another example:

```go
import "net/http"

http.ListenAndServe(":8080", nil)
```

Here:

```text
http             → package
ListenAndServe   → function
```

General pattern:

```text
packageName.FunctionName()
```

---

# 5. Exported vs Unexported

Go uses **capitalization** to control whether something can be accessed from another package.

### Exported

Starts with a capital letter:

```go
func Add(a int, b int) int {
    return a + b
}
```

Another package can use:

```go
calculator.Add(10, 20)
```

### Unexported

Starts with a lowercase letter:

```go
func add(a int, b int) int {
    return a + b
}
```

Another package cannot use:

```go
calculator.add(10, 20) // ❌
```

### Remember

```text
Uppercase → Exported → accessible outside package

lowercase → Unexported → package-internal
```

This is Go's basic mechanism for encapsulation.

---

# 6. Module in Go

A **module** represents a Go project and its dependency boundary.

A module is defined by a:

```text
go.mod
```

file.

Think:

```text
MODULE
  │
  ├── PACKAGE
  │     ├── file.go
  │     └── file.go
  │
  ├── PACKAGE
  │     └── file.go
  │
  └── PACKAGE
        └── file.go
```

So:

```text
Module
  ↓
contains packages
  ↓
packages contain Go files
```

---

# 7. `go.mod`

`go.mod` is the configuration/manifest file for a Go module.

Example:

```go
module arrays

go 1.27.0
```

There are two important parts here.

### Module name

```go
module arrays
```

This identifies the module.

For a real project, it is common to use a repository-style path:

```go
module github.com/username/mybackend
```

### Go version

```go
go 1.27.0
```

This specifies the Go version associated with the module.

---

# 8. Module vs Package

This distinction is very important.

### Module

The overall project:

```text
mybackend/
├── go.mod
├── main.go
├── users/
└── database/
```

### Package

A group of Go files in a directory:

```text
users/
├── user.go
└── auth.go
```

Both might contain:

```go
package users
```

### Visual model

```text
                 MODULE
                   │
       ┌───────────┼───────────┐
       ↓           ↓           ↓
    package     package     package
      main        users      database
       │           │           │
    main.go     user.go      db.go
                auth.go
```

---

# 9. Module and Import Paths

Suppose your `go.mod` contains:

```go
module github.com/ankush/mybackend
```

And you have:

```text
mybackend/
├── go.mod
├── main.go
└── users/
    └── user.go
```

Then the `users` package can be imported as:

```go
import "github.com/ankush/mybackend/users"
```

Conceptually:

```text
Module path
     +
Package directory
     ↓
Import path

github.com/ankush/mybackend/users
```

---

# 10. Dependencies

A Go module can depend on external packages.

For example:

```bash
go get github.com/gin-gonic/gin
```

Go records dependencies in:

```text
go.mod
```

For example:

```go
require github.com/gin-gonic/gin v1.10.0
```

Go also maintains:

```text
go.sum
```

`go.sum` contains checksums used to verify module dependencies.

### Remember

```text
go.mod
    ↓
Module information + dependencies

go.sum
    ↓
Dependency checksums
```

---

# 11. Useful Go Module Commands

### Create a module

```bash
go mod init <module-name>
```

Example:

```bash
go mod init github.com/username/mybackend
```

### Add a dependency

```bash
go get <package>
```

Example:

```bash
go get github.com/gin-gonic/gin
```

### Clean/update dependencies

```bash
go mod tidy
```

`go mod tidy` adds missing dependencies and removes unnecessary ones.

### Download dependencies

```bash
go mod download
```

### View modules/dependencies

```bash
go list -m all
```

---

# 12. HTTP Basics

HTTP is the protocol used for communication between a client and a server.

For backend development:

```text
Client
   │
   │ HTTP Request
   ↓
Go Backend
   │
   │ HTTP Response
   ↓
Client
```

For example:

```text
Browser
   ↓
GET /users
   ↓
Go Server
   ↓
JSON response
```

---

# 13. HTTP Request

An HTTP request contains things such as:

```text
HTTP Method
URL/Path
Headers
Body
```

Example:

```http
POST /users HTTP/1.1
Content-Type: application/json

{
    "name": "Ankush"
}
```

Here:

```text
POST
    ↓
HTTP method

/users
    ↓
Path

Content-Type
    ↓
Header

JSON
    ↓
Request body
```

---

# 14. Common HTTP Methods

These are important for backend development:

| Method | Purpose |
|---|---|
| GET | Retrieve data |
| POST | Create data |
| PUT | Replace/update data |
| PATCH | Partially update data |
| DELETE | Delete data |

Example API:

```text
GET     /users
GET     /users/10
POST    /users
PATCH   /users/10
DELETE  /users/10
```

---

# 15. Go's `net/http` Package

Go has a built-in HTTP package:

```go
import "net/http"
```

You can create a server without using a framework.

Basic example:

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}

func main() {
    http.HandleFunc("/", handler)

    http.ListenAndServe(":8080", nil)
}
```

Run:

```bash
go run .
```

Then visit:

```text
http://localhost:8080
```

---

# 16. `http.Request`

A handler receives:

```go
func handler(w http.ResponseWriter, r *http.Request)
```

`r` represents the **incoming HTTP request**.

You can access:

```go
r.Method
r.URL
r.Header
r.Body
```

For example:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
    fmt.Println(r.URL.Path)
}
```

If the client requests:

```text
GET /users
```

you can get:

```text
GET
/users
```

---

# 17. `http.ResponseWriter`

`w` represents the HTTP response that your server sends back.

You can write a response:

```go
fmt.Fprintln(w, "Hello")
```

Set a response header:

```go
w.Header().Set("Content-Type", "application/json")
```

Set a status code:

```go
w.WriteHeader(http.StatusCreated)
```

Mental model:

```text
Incoming Request
       ↓
*r http.Request
       ↓
   Your Handler
       ↓
http.ResponseWriter
       ↓
Outgoing Response
```

---

# 18. HTTP Status Codes

Important status codes:

### Success

```text
200 OK
201 Created
204 No Content
```

### Client errors

```text
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
```

### Server error

```text
500 Internal Server Error
```

In Go:

```go
w.WriteHeader(http.StatusCreated)
```

is equivalent to sending:

```text
201 Created
```

---

# 19. Routing

Routing determines which code handles a particular URL.

For example:

```text
GET /users
      ↓
usersHandler()

GET /products
      ↓
productsHandler()
```

In Go:

```go
http.HandleFunc("/users", usersHandler)
http.HandleFunc("/products", productsHandler)
```

---

# 20. The Big Picture

Everything you've learned so far fits together:

```text
                 GO MODULE
                 go.mod
                    │
       ┌────────────┼────────────┐
       ↓            ↓            ↓
    PACKAGE      PACKAGE      PACKAGE
      main         users       database
       │
       ↓
   HTTP SERVER
       │
       ↓
  net/http
       │
       ↓
    Request
       │
       ↓
    Handler
       │
       ↓
   Response
```

---

# Quick Revision

### Package

> A group of related Go source files.

```go
package users
```

### Module

> The overall Go project/dependency boundary.

```go
module github.com/username/mybackend
```

### `go.mod`

> Defines the module and manages its dependencies.

```go
module mybackend

go 1.27.0
```

### Import

> Allows your code to use another package.

```go
import "net/http"
```

### Exported

```go
Add()
```

Accessible outside the package.

### Unexported

```go
add()
```

Only accessible within the package.

### HTTP Request

```text
Method + URL + Headers + Body
```

### HTTP Response

```text
Status Code + Headers + Body
```

### Go HTTP Handler

```go
func handler(w http.ResponseWriter, r *http.Request)
```

```text
r → incoming request
w → outgoing response
```

---

# 🎯 What to Learn Next

Your next logical topic is **Go HTTP in depth**:

```text
HTTP
  ↓
Request & Response
  ↓
Headers
  ↓
URL
  ↓
Query Parameters
  ↓
Path Parameters
  ↓
Request Body
  ↓
JSON
  ↓
Routing
  ↓
Middleware
  ↓
REST API
```

Once you understand these, you'll have the foundation needed to start building a real Go backend.