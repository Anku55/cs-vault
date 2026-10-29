# Go User Management API — Detailed Code Explanation

## Table of Contents

1. [Project Overview](#1-project-overview)
2. [Package and Imports](#2-package-and-imports)
3. [The `main()` Function and Routing](#3-the-main-function-and-routing)
4. [The User Struct](#4-the-user-struct)
5. [Global Variables and In-Memory Storage](#5-global-variables-and-in-memory-storage)
6. [HTTP Handler Fundamentals](#6-http-handler-fundamentals)
7. [Health Check Handler](#7-health-check-handler)
8. [POST — Create a User](#8-post--create-a-user)
9. [GET — Retrieve All Users](#9-get--retrieve-all-users)
10. [GET — Retrieve a User by ID](#10-get--retrieve-a-user-by-id)
11. [DELETE — Delete a User](#11-delete--delete-a-user)
12. [PUT — Update a User](#12-put--update-a-user)
13. [PATCH — Update Selected Fields](#13-patch--update-selected-fields)
14. [Why Set `Content-Type`](#14-why-set-content-type)
15. [Understanding `Lock`, `RLock`, and `Unlock`](#15-understanding-lock-rlock-and-unlock)
16. [Why Do We Use `return`](#16-why-do-we-use-return)
17. [JSON Encoding and Decoding](#17-json-encoding-and-decoding)
18. [HTTP Status Codes](#18-http-status-codes)
19. [Complete Request Lifecycle](#19-complete-request-lifecycle)
20. [Important Improvements](#20-important-improvements)
21. [Revision Checklist](#21-revision-checklist)

---

# 1. Project Overview

We are building a REST API for managing users using Go's standard `net/http` package.

The API supports creating, reading, updating, and deleting users.

## 1.1 Available endpoints

| Method | Endpoint | Purpose |
|---|---|---|
| GET | `/health` | Check whether the server is running |
| POST | `/users` | Create a new user |
| GET | `/users` | Retrieve all users |
| GET | `/users/{id}` | Retrieve one user |
| PUT | `/users/{id}` | Update the user's name and age |
| PATCH | `/users/{id}` | Update selected fields |
| DELETE | `/users/{id}` | Delete a user |

## 1.2 How the application works

```text
Client (Browser / Postman / curl)
              |
              | HTTP Request
              v
       Go HTTP Server
              |
              v
         ServeMux Router
              |
              v
        Matched Handler
              |
              v
       Validate Input
              |
              v
    Read or Modify User Map
              |
              v
       Create HTTP Response
              |
              v
             Client
```

For example, when the client sends `GET /users/1`, Go matches the route, calls `getUserByIDHandler`, validates the ID, looks up the user in the map, and returns the user as JSON.

## 1.3 What storage are we using?

Currently, the API stores data in an in-memory Go map.

```go
var users = make(map[int]User)
```

This means:

- Data is stored in the server's process memory.
- The map is lost when the program stops.
- The API does not yet use PostgreSQL or another persistent database.
- The map is protected by a mutex because multiple HTTP requests may access it concurrently.

---

# 2. Package and Imports

Your program starts with:

```go
package main
```

## What does `package main` mean?

Every Go source file belongs to a package.

The special package name `main` is used to build an executable program. The program starts by calling the `main()` function in this package.

Your imports are:

```go
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)
```

## What is each package used for?

| Package | Purpose | Used in your API |
|---|---|---|
| `encoding/json` | Converts between JSON and Go values | Decoding request bodies and encoding responses |
| `fmt` | Formatting and printing | Writing the health-check response |
| `log` | Logging messages and errors | Server startup and JSON encoding errors |
| `net/http` | HTTP server, requests, responses, routing | Routes, handlers, status codes |
| `strconv` | String and number conversions | Converting path IDs into integers |
| `sync` | Synchronization primitives | Protecting the shared user map |

For example:

```go
id, err := strconv.Atoi(stringID)
```

uses `strconv` to convert a string to an integer.

And:

```go
json.NewEncoder(w).Encode(user)
```

uses `encoding/json` to write a Go `User` value as JSON.

---

# 3. The `main()` Function and Routing

Your code:

```go
func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", healthHandler)
    mux.HandleFunc("POST /users", createUserHandler)
    mux.HandleFunc("GET /users", getUsersHandler)
    mux.HandleFunc("GET /users/{id}", getUserByIDHandler)
    mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)
    mux.HandleFunc("PUT /users/{id}", updateUserHandler)
    mux.HandleFunc("PATCH /users/{id}", patchUserHandler)

    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

## 3.1 What is `http.NewServeMux()`?

```go
mux := http.NewServeMux()
```

`http.NewServeMux()` creates a request multiplexer, commonly called a router.

Its job is to match incoming HTTP requests to the correct handler based on their method and path.

For example:

```text
GET /users
```

matches:

```go
mux.HandleFunc("GET /users", getUsersHandler)
```

Whereas:

```text
POST /users
```

matches:

```go
mux.HandleFunc("POST /users", createUserHandler)
```

The path is the same, but the HTTP method is different.

## 3.2 What does `HandleFunc()` do?

```go
mux.HandleFunc("GET /users", getUsersHandler)
```

It registers a handler for a particular request pattern.

The first argument defines the pattern. The second argument is the function that should handle matching requests.

Your route:

```go
mux.HandleFunc("GET /users/{id}", getUserByIDHandler)
```

uses `{id}` as a named wildcard path segment.

For example:

```text
GET /users/5
GET /users/20
```

Both requests match the route, and the handler can read the ID with:

```go
idString := r.PathValue("id")
```

**Version note:** Method-aware patterns and `Request.PathValue()` require Go 1.22 or later.

## 3.3 What does `http.ListenAndServe()` do?

```go
http.ListenAndServe(":8080", mux)
```

This starts an HTTP server that listens on port `8080` and uses `mux` to handle incoming requests.

The string:

```go
":8080"
```

means the server listens on port 8080 on the available local interfaces.

You can access it locally using:

```text
http://localhost:8080
```

## 3.4 Why use `log.Fatal()`?

```go
log.Fatal(http.ListenAndServe(":8080", mux))
```

`ListenAndServe()` normally blocks while the server is running. If it returns an error, `log.Fatal()` logs the error and terminates the program.

For example, the server may fail to start because port 8080 is already occupied.

`log.Fatal()` also exits immediately, so deferred functions are not run. This matters in more advanced server lifecycle management.

---

# 4. The User Struct

Your code:

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

## 4.1 What is a struct?

A struct groups related fields into a single type.

Instead of maintaining separate variables for each user's ID, name, and age, we represent the user as one value.

Example:

```go
user := User{
    ID:   1,
    Name: "Ankit",
    Age:  21,
}
```

The value contains all three fields.

## 4.2 What are JSON struct tags?

Consider:

```go
Name string `json:"name"`
```

The `json:"name"` tag tells Go's JSON package to use the JSON property name `name` when encoding or decoding the field.

The corresponding JSON object is:

```json
{
  "id": 1,
  "name": "Ankit",
  "age": 21
}
```

Without these tags, Go normally uses exported field names such as `ID`, `Name`, and `Age` when encoding JSON.

**Important:** The fields are exported because their names begin with uppercase letters. The standard JSON encoder does not encode unexported struct fields.

---

# 5. Global Variables and In-Memory Storage

Your code:

```go
var (
    users   = make(map[int]User)
    nextId  = 1
    usersMu sync.RWMutex
)
```

These variables are declared outside `main()` and the handlers, so they are shared by the functions in this package.

## 5.1 The user map

```go
users = make(map[int]User)
```

The map stores users using their IDs as keys.

The key type is `int`, and the value type is `User`.

Conceptually:

```text
users
  |
  +-- 1 --> User{ID: 1, Name: "Ankit", Age: 21}
  |
  +-- 2 --> User{ID: 2, Name: "Rahul", Age: 22}
```

This makes lookups by ID convenient:

```go
user, exists := users[id]
```

## 5.2 The ID counter

```go
nextId = 1
```

This variable stores the ID to assign to the next user.

After creating a user:

```go
users[user.ID] = user
nextId++
```

The current ID is used first, and then the counter is incremented.

This counter is appropriate for the current learning project. It is not a durable ID-generation strategy: restarting the server resets it, and the data itself is lost.

## 5.3 The mutex

```go
usersMu sync.RWMutex
```

The mutex coordinates access to shared state.

Go's HTTP server may process multiple requests concurrently. If two handlers modify the same map without synchronization, they can cause a data race or concurrent map access failure.

We use:

- `usersMu.RLock()` for protected reads.
- `usersMu.Lock()` for exclusive access and writes.
- `usersMu.RUnlock()` to release a read lock.
- `usersMu.Unlock()` to release an exclusive lock.

We will examine these in detail later.

---

# 6. HTTP Handler Fundamentals

All your handlers have this structure:

```go
func someHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the request.
}
```

## 6.1 What is `http.ResponseWriter`?

```go
w http.ResponseWriter
```

`w` is the object used to construct the HTTP response.

You can use it to:

- Set response headers.
- Set the HTTP status code.
- Write the response body.

Examples:

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
fmt.Fprintln(w, "Server is running")
```

## 6.2 What is `*http.Request`?

```go
r *http.Request
```

`r` is a pointer to the incoming HTTP request.

It contains information such as:

```go
r.Method
r.URL.Path
r.URL.Query()
r.PathValue("id")
r.Header
r.Body
```

For example:

```go
stringID := r.PathValue("id")
```

retrieves a path parameter from a matching route.

And:

```go
json.NewDecoder(r.Body).Decode(&input)
```

reads JSON data from the incoming request body.

The request and response have different roles:

| Value | Role |
|---|---|
| `r` | Read information sent by the client |
| `w` | Construct the response sent to the client |

---

# 7. Health Check Handler

Your code:

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Server is running")
}
```

## How does it work?

When a client sends:

```http
GET /health
```

Go routes the request to `healthHandler`.

Then:

```go
fmt.Fprintln(w, "Server is running")
```

writes the text to the HTTP response body.

The response body is:

```text
Server is running
```

### Why `Fprintln()` instead of `Println()`?

```go
fmt.Println("Server is running")
```

prints to standard output, usually the terminal.

But:

```go
fmt.Fprintln(w, "Server is running")
```

writes to the supplied writer. Here, `w` is the HTTP response writer, so the text goes to the client.

This distinction is important when writing HTTP handlers.

A health endpoint is useful for checking whether the server is reachable. It does not, by itself, prove that a database or other dependency is healthy.

---

# 8. POST — Create a User

Your handler begins with:

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
```

This defines an anonymous struct to hold the fields expected in the request body.

## 8.1 Why do we define `input`?

The client may send:

```json
{
  "name": "Ankit",
  "age": 21
}
```

We need a Go value into which the JSON decoder can place the supplied fields.

That is the purpose of:

```go
var input struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

The fields have types:

- `Name`: `string`
- `Age`: `int`

The `ID` is not included because the server generates it.

## 8.2 Decode the request body

```go
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w, "Invalid json body ", http.StatusBadRequest)
    return
}
```

Let's break this down.

### `r.Body`

Contains the incoming HTTP request body as a stream of bytes.

### `json.NewDecoder(r.Body)`

Creates a JSON decoder that reads from the body.

### `.Decode(&input)`

Parses the JSON and fills the `input` struct.

The `&input` expression passes the address of `input`, allowing the decoder to populate it.

### `if err := ...; err != nil`

This is Go's short `if` initialization syntax.

The decoder returns an error if it cannot decode the JSON into the target value.

If an error occurs, the handler responds with HTTP 400 and stops.

For example, this body is invalid JSON:

```text
{name: Ankit}
```

JSON requires quoted property names:

```json
{
  "name": "Ankit",
  "age": 21
}
```

## 8.3 Validate the name

```go
if input.Name == "" {
    http.Error(w, "Name is required", http.StatusBadRequest)
    return
}
```

A missing `name` field usually leaves the string at its zero value, `""`.

This condition rejects an empty name.

However, a whitespace-only name such as `"   "` passes this check. You can address that later with `strings.TrimSpace()`.

## 8.4 Validate the age

```go
if input.Age <= 0 {
    http.Error(w, "Age can not be negaive", http.StatusBadRequest)
    return
}
```

This rejects zero and negative ages.

The validation reflects the business rule that age must be positive. The error message has a spelling mistake, but that does not affect program execution.

There is also a limitation: because `Age` is a regular integer field, a missing age field becomes `0` and is rejected, but the code does not distinguish a missing age from an explicitly supplied zero.

## 8.5 Acquire the write lock

```go
usersMu.Lock()
```

The handler is about to modify shared state.

An exclusive lock prevents other handlers from accessing the protected map while the lock is held.

## 8.6 Create the user

```go
user := User{
    ID:   nextId,
    Name: input.Name,
    Age:  input.Age,
}
```

This creates a `User` value.

- `ID` comes from the server's counter.
- `Name` comes from the decoded request.
- `Age` comes from the decoded request.

## 8.7 Store the user

```go
users[user.ID] = user
nextId++
```

The first line saves the user under its ID.

The second increments the counter so the next user receives a different ID.

Both operations happen while the exclusive lock is held. This protects the map and counter against concurrent creation requests.

## 8.8 Release the write lock

```go
usersMu.Unlock()
```

The handler has finished modifying the shared data, so it releases the lock.

The response can now be constructed without holding the map lock.

## 8.9 Set the response content type

```go
w.Header().Set("Content-type", "application/json")
```

This tells the client that the response body is JSON.

The standard spelling is `Content-Type`. HTTP header names are case-insensitive, so `Content-type` works too. Use the conventional spelling for consistency.

## 8.10 Set the status code

```go
w.WriteHeader(http.StatusCreated)
```

This sets the response status to `201 Created`, indicating that the resource was created successfully.

## 8.11 Encode the user

```go
if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Printf("Failed to encode user response%v", err)
}
```

This converts the `User` struct into JSON and writes it to the response.

Example:

```json
{
  "id": 1,
  "name": "Ankit",
  "age": 21
}
```

The encoder writes to `w`, so the JSON goes to the HTTP client.

The error is logged if encoding fails. Since the response status may already have been written, changing the status afterward is generally not possible.

---

# 9. GET — Retrieve All Users

Your handler:

```go
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    usersMu.RLock()

    var result []User

    for _, user := range users {
        result = append(result, user)
    }

    usersMu.RUnlock()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(result); err != nil {
        log.Printf("Failed to encode users response: %v", err)
    }
}
```

## 9.1 Why use `RLock()`?

```go
usersMu.RLock()
```

This handler reads the map but does not modify it.

A read lock permits multiple readers to access the protected data simultaneously while preventing a writer from modifying the map during the read.

## 9.2 Why create a slice?

```go
var result []User
```

This declares a slice of users.

The map is useful for looking up individual users by ID. A slice is convenient for returning a collection in JSON.

## 9.3 Understand the loop

```go
for _, user := range users {
    result = append(result, user)
}
```

The `range` loop iterates through the map.

For a map, `range` provides the key and value. We use `_` to ignore the key and `user` to receive the value.

The underscore is Go's blank identifier. It tells Go that we intentionally do not need that value.

`append()` adds each user to the result slice.

Map iteration order is unspecified, so the output array may not be sorted.

## 9.4 Release the read lock

```go
usersMu.RUnlock()
```

We have finished reading the map, so we release the lock.

JSON encoding happens after the lock is released. This avoids holding the map lock while writing the HTTP response.

## 9.5 Set the response headers and status

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
```

- `Content-Type` tells the client to interpret the body as JSON.
- `200 OK` indicates successful retrieval.

The encoder then serializes the slice into a JSON array.

### A small improvement

Your code declares:

```go
var result []User
```

When there are no users, this is a nil slice, which encodes as:

```json
null
```

If you want the endpoint to return an empty JSON array instead, use:

```go
result := make([]User, 0, len(users))
```

Then an empty collection encodes as:

```json
[]
```

This is often a more convenient API response for a collection endpoint.

---

# 10. GET — Retrieve a User by ID

Your handler:

```go
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    stringID := r.PathValue("id")
```

## 10.1 Read the path parameter

For a request such as:

```text
GET /users/5
```

the route `/users/{id}` matches, and:

```go
stringID := r.PathValue("id")
```

returns the string `"5"`.

Path values are strings, so conversion is necessary when the application expects an integer.

## 10.2 Convert the ID

```go
id, err := strconv.Atoi(stringID)
if err != nil || id <= 0 {
    http.Error(w, "User ID must be a positive integer", http.StatusBadRequest)
    return
}
```

`strconv.Atoi()` converts the string into an integer.

The condition rejects the ID when:

- Conversion fails, or
- The resulting integer is zero or negative.

The `||` operator means OR.

If either condition is true, the handler sends a `400 Bad Request` response and returns.

## 10.3 Look up the user

```go
usersMu.RLock()
user, exists := users[id]
usersMu.RUnlock()
```

This is a map lookup.

- `user` receives the stored user value if the ID exists.
- `exists` is `true` if the key is present.
- `RLock()` protects the read.
- `RUnlock()` releases the read lock.

The comma-ok map lookup is useful because a missing map entry returns the zero value of the value type. The `exists` result tells us whether the key was actually present.

## 10.4 Handle a missing user

```go
if !exists {
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

The `!` operator means NOT.

Therefore, `!exists` is true when the user does not exist.

The server responds with `404 Not Found` and stops processing the handler.

## 10.5 Return the user

```go
w.Header().Set("Content-Type", "application/json")
if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Printf("Failed to encode user response: %v", err)
}
```

The handler returns the user as JSON.

Notice that your code does not explicitly call:

```go
w.WriteHeader(http.StatusOK)
```

That is valid. If no status is written explicitly, the first body write normally commits the default status, `200 OK`.

Explicitly setting `200 OK` is optional here.

---

# 11. DELETE — Delete a User

Your handler uses:

```go
usersMu.Lock()
defer usersMu.Unlock()
```

This is slightly different from your GET and PUT handlers.

## 11.1 Why use an exclusive lock?

Deleting modifies the map, so it requires exclusive access.

```go
usersMu.Lock()
```

This prevents concurrent readers and writers from accessing the protected map while the deletion operation is in progress.

## 11.2 What does `defer usersMu.Unlock()` do?

```go
defer usersMu.Unlock()
```

`defer` schedules `Unlock()` to execute when the surrounding function returns.

This means the lock is released when the handler exits normally, including when it returns early after finding that the user does not exist.

It reduces the chance of accidentally leaving the lock held on an error path.

## 11.3 Check whether the user exists

```go
_, exists := users[id]
```

We only need the existence flag, so we discard the user value with `_`.

If the user does not exist:

```go
if !exists {
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

The handler returns a 404 response. The deferred `Unlock()` still runs when the handler returns.

## 11.4 Delete the user

```go
delete(users, id)
```

Go's built-in `delete()` removes the entry for the specified key.

The existence check and deletion are both performed while the exclusive lock is held. This avoids a race between checking the key and deleting it.

## 11.5 Return `204 No Content`

```go
w.WriteHeader(http.StatusNoContent)
```

A `204 No Content` response means the deletion succeeded and the server is not returning a body.

Do not append a message body after writing a 204 response.

---

# 12. PUT — Update a User

The PUT handler changes both editable fields: `Name` and `Age`.

## 12.1 Read and validate the ID

```go
stringID := r.PathValue("id")

id, err := strconv.Atoi(stringID)
if err != nil || id <= 0 {
    http.Error(w, "User ID must be a positive integer", http.StatusBadRequest)
    return
}
```

This follows the same path-parameter validation used in GET and DELETE.

## 12.2 Decode the body

```go
var input struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w, "Invalid JSON body", http.StatusBadRequest)
    return
}
```

The request body must contain valid JSON that can be decoded into these fields.

## 12.3 Validate both fields

```go
if input.Name == "" {
    http.Error(w, "Name is required", http.StatusBadRequest)
    return
}

if input.Age <= 0 {
    http.Error(w, "Age must be greater than zero", http.StatusBadRequest)
    return
}
```

This API requires both fields for PUT. That is consistent with the intended behavior of replacing the editable fields as a unit.

## 12.4 Acquire the lock and retrieve the user

```go
usersMu.Lock()

user, exists := users[id]
if !exists {
    usersMu.Unlock()
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

The exclusive lock protects the entire read-modify-write operation.

The user lookup happens under the lock.

If the user does not exist, the code explicitly releases the lock before returning the error response.

This explicit unlock is necessary in the current code because the function does not use `defer` here.

## 12.5 Update the fields

```go
user.Name = input.Name
user.Age = input.Age
users[id] = user
```

The map lookup returns a copy of the stored `User` struct.

Changing `user.Name` and `user.Age` modifies that local copy. The final assignment:

```go
users[id] = user
```

stores the updated struct back in the map.

The update is completed while the exclusive lock is held.

## 12.6 Release the lock and return the result

```go
usersMu.Unlock()

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)

if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Printf("Failed to encode user response: %v", err)
}
```

The lock is released before response encoding.

The server returns `200 OK` with the updated user as JSON.

### Important design note

The lock protects the shared map while it is being read and modified. Once the updated value has been copied into the local variable `user`, encoding that local value does not require holding the map lock.

---

# 13. PATCH — Update Selected Fields

PATCH is different from PUT in this implementation because the client may supply only the fields it wants to change.

## 13.1 Why use pointer fields?

Your input struct is:

```go
var input struct {
    Name *string `json:"name"`
    Age  *int    `json:"age"`
}
```

The `*string` and `*int` types are pointers.

We use pointers because we need to distinguish an omitted field from a field containing a normal zero value.

For example:

```json
{
  "age": 23
}
```

The decoder sets `input.Age` to a non-nil pointer to `23`, while `input.Name` remains nil.

That lets the handler update only the age.

## 13.2 Reject invalid IDs and JSON

The handler first extracts and validates the path ID, then decodes the JSON request body.

These steps are the same as in the other handlers.

## 13.3 Reject an empty patch

```go
if input.Name == nil && input.Age == nil {
    http.Error(w, "At least one field is required", http.StatusBadRequest)
    return
}
```

The `&&` operator means AND.

If both fields are nil, neither field was supplied in a way that produced a non-nil pointer, so there is nothing to update.

The handler rejects the request rather than returning success without changing anything.

Note that ordinary pointer fields do not distinguish a missing field from an explicit JSON `null`; both decode to nil.

## 13.4 Validate supplied values

```go
if input.Name != nil && *input.Name == "" {
    http.Error(w, "Name cannot be empty", http.StatusBadRequest)
    return
}

if input.Age != nil && *input.Age <= 0 {
    http.Error(w, "Age must be greater than zero", http.StatusBadRequest)
    return
}
```

Each condition has two parts.

For the name:

1. `input.Name != nil` checks that a name was supplied.
2. `*input.Name == ""` checks the value stored behind the pointer.

The `*` operator dereferences a pointer to access its value.

The same principle applies to age.

## 13.5 Lock and retrieve the user

```go
usersMu.Lock()

user, exists := users[id]
if !exists {
    usersMu.Unlock()
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

PATCH changes shared state, so it uses the exclusive lock.

The lock is explicitly released before returning if the user does not exist.

## 13.6 Update only the supplied fields

```go
if input.Name != nil {
    user.Name = *input.Name
}

if input.Age != nil {
    user.Age = *input.Age
}

users[id] = user
usersMu.Unlock()
```

If the client sends only an age, the name is not modified.

If the client sends only a name, the age is not modified.

After modifying the local copy, the code saves it back to the map and releases the lock.

## 13.7 Return the updated user

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)

if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Printf("Failed to encode updated user: %v", err)
}
```

The response contains the complete updated user, not merely the field that changed.

---

# 14. Why Set `Content-Type`?

This is one of the most important HTTP concepts in your API.

Consider:

```go
w.Header().Set("Content-Type", "application/json")
```

The `Content-Type` response header tells the client what kind of data the response body contains.

For example, if the server returns:

```json
{
  "id": 1,
  "name": "Ankit",
  "age": 21
}
```

the response should indicate that the body is JSON.

A response might look like:

```http
HTTP/1.1 200 OK
Content-Type: application/json

{"id":1,"name":"Ankit","age":21}
```

## 14.1 Why does the client need this header?

The client needs to know how to interpret the response.

Depending on the client, the content type can influence:

- How the response body is parsed.
- Whether the client treats the body as JSON, plain text, HTML, or another format.
- How browser tools and API clients display the response.

For example:

```go
w.Header().Set("Content-Type", "application/json")
```

means the response is JSON.

For a plain-text response, you could use:

```go
w.Header().Set("Content-Type", "text/plain; charset=utf-8")
```

## 14.2 Does `Content-Type` convert data into JSON?

No.

This header only describes the response body. It does not transform the body into JSON.

Incorrect:

```go
w.Header().Set("Content-Type", "application/json")
fmt.Fprintln(w, "Hello")
```

The body is still the plain text `Hello`; setting the header does not encode it as a JSON string.

Correct:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(user)
```

Here, the JSON encoder actually serializes the Go value into JSON.

## 14.3 Why must headers be set before writing the body?

In Go, response headers are normally committed when the handler writes the response status or begins writing the body.

Therefore, set headers first:

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(user)
```

After the response has been committed, changing headers generally will not change the headers already sent to the client.

## 14.4 What about error responses?

Your code uses:

```go
http.Error(w, "User not found", http.StatusNotFound)
```

`http.Error()` sets an error status and writes a plain-text error response. It sets an appropriate plain-text content type itself.

Your successful CRUD responses use JSON, while your current error responses use plain text. That is valid, although many APIs choose a consistent JSON error format.

---

# 15. Understanding `Lock`, `RLock`, and `Unlock`

This is critical because your server can handle concurrent HTTP requests.

## 15.1 What problem does a mutex solve?

Suppose two clients send requests at nearly the same time:

```text
Request A: POST /users
Request B: POST /users
```

Both handlers may attempt to update:

```go
users[user.ID] = user
nextId++
```

Without synchronization, concurrent modifications to a Go map can cause a data race or a runtime failure.

A mutex coordinates access to shared data.

## 15.2 `Lock()` — exclusive access

```go
usersMu.Lock()
```

Use `Lock()` when a handler must modify shared state.

Examples:

- Creating a user.
- Updating a user.
- Deleting a user.

While an exclusive lock is held, other readers and writers using the same mutex must wait.

Conceptually:

```text
Handler A                 Handler B
---------                 ---------
Lock()
Modify map
                          Wait for lock...
Unlock()
                          Lock()
                          Modify map
                          Unlock()
```

The actual scheduling can vary, but the mutex prevents simultaneous protected access.

## 15.3 `RLock()` — shared read access

```go
usersMu.RLock()
```

Use `RLock()` when the handler needs to read protected shared state without modifying it.

Examples:

- Retrieving all users.
- Looking up a single user.

Multiple handlers can hold a read lock at the same time, provided no writer has acquired the exclusive lock.

Conceptually:

```text
GET Handler A: RLock() ---- Read
GET Handler B: RLock() ---- Read
                           |
                  Both may read concurrently
```

A writer cannot acquire the exclusive lock while read locks remain held.

## 15.4 `Unlock()` — release an exclusive lock

```go
usersMu.Lock()

users[id] = user

usersMu.Unlock()
```

`Unlock()` releases a lock previously acquired with `Lock()`.

If you forget to unlock, other handlers may wait indefinitely for access to the protected data.

## 15.5 `RUnlock()` — release a read lock

```go
usersMu.RLock()

user, exists := users[id]

usersMu.RUnlock()
```

`RUnlock()` releases a read lock acquired with `RLock()`.

Do not mix the lock types: release `Lock()` with `Unlock()`, and release `RLock()` with `RUnlock()`.

## 15.6 Why not use `Lock()` for every operation?

You technically could protect all access with an exclusive lock, but readers would block other readers unnecessarily.

For example, multiple requests retrieving users do not modify the map, so they can safely share a read lock.

`RWMutex` allows those reads to proceed concurrently while still coordinating them with writes.

For small applications, a plain `sync.Mutex` may also be sufficient. The important requirement is consistent synchronization, not using the most complicated mutex type.

## 15.7 Why must a read-modify-write operation stay under one lock?

Consider the PUT handler:

```go
usersMu.Lock()

user, exists := users[id]
if exists {
    user.Name = input.Name
    user.Age = input.Age
    users[id] = user
}

usersMu.Unlock()
```

The lookup, modification, and assignment happen while the exclusive lock is held.

If the handler unlocked after reading and later reacquired the lock to save, another request could change the user in between. The first request might then overwrite the other request's changes.

**Rule:** Protect the entire logical operation that must be atomic, not just an individual map statement.

## 15.8 Why does DELETE use `defer`?

Your DELETE handler uses:

```go
usersMu.Lock()
defer usersMu.Unlock()
```

This ensures that the lock is released when the function returns, including early returns.

For example:

```go
if !exists {
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

Even though `return` exits the function, the deferred `Unlock()` runs first.

By contrast, your PUT and PATCH handlers explicitly unlock before their early return:

```go
if !exists {
    usersMu.Unlock()
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

Both patterns are valid.

For short critical sections, explicit unlocks can make the protected region easy to see. `defer` can reduce mistakes on multiple return paths. Avoid holding a lock across unrelated or slow work.

## 15.9 Which variables must be protected?

In this API:

```go
var (
    users  = make(map[int]User)
    nextId = 1
    usersMu sync.RWMutex
)
```

Both `users` and `nextId` are shared mutable state.

All reads and writes of those variables should follow the synchronization policy. Your current handlers protect their relevant map operations, and user creation increments `nextId` while holding the exclusive lock.

---

# 16. Why Do We Use `return`?

In your handlers, `return` usually stops processing after a request fails validation or a resource is not found.

Consider:

```go
if err != nil || id <= 0 {
    http.Error(w, "User ID must be a positive integer", http.StatusBadRequest)
    return
}
```

There are two separate operations here:

1. `http.Error()` writes an error response.
2. `return` exits the handler.

**Writing an error response does not automatically stop the Go function.**

Without `return`, the code below the condition would still execute.

For example, this is dangerous:

```go
if err != nil {
    http.Error(w, "Invalid ID", http.StatusBadRequest)
}

// Incorrect: execution continues with an invalid ID.
user, exists := users[id]
```

The corrected version is:

```go
if err != nil {
    http.Error(w, "Invalid ID", http.StatusBadRequest)
    return
}

user, exists := users[id]
```

## 16.1 Common reasons to return early

| Condition | Why return? |
|---|---|
| Invalid JSON | The handler cannot safely use the decoded input |
| Missing name | The request violates validation rules |
| Invalid ID | The ID cannot be used for the intended lookup |
| User not found | There is no resource to update or delete |
| Empty PATCH body | No field has been supplied for an update |

## 16.2 Early returns make handlers easier to read

Instead of nesting the entire handler inside many `else` blocks, reject invalid requests immediately.

This pattern is often called guard clauses.

```go
if invalidInput {
    http.Error(w, "Invalid input", http.StatusBadRequest)
    return
}

// Continue with valid input.
```

The main business logic can then remain less deeply nested.

## 16.3 A `return` does not undo a response

Once a response has been written, `return` does not erase it. It only stops further execution of the current function.

Similarly, returning from a handler does not mean the entire HTTP server stops. The server continues serving other requests.

---

# 17. JSON Encoding and Decoding

## 17.1 Decode a request

```go
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w, "Invalid JSON body", http.StatusBadRequest)
    return
}
```

The decoder reads the request body and converts the JSON into a Go value.

## 17.2 Encode a response

```go
if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Printf("Failed to encode user response: %v", err)
}
```

The encoder converts the Go value into JSON and writes it to the HTTP response.

## 17.3 Why do we check errors?

Both decoding and encoding can fail.

For request decoding, malformed JSON or incompatible field types can cause an error. That is why the handler rejects the request.

For response encoding, an error may occur while writing the response. Logging is useful for diagnosing such failures.

There is an important difference: request decoding errors can normally be reported before the response starts, but an encoding error may occur after part of the response has already been sent. At that point, the handler may not be able to replace the response with a clean error status.

## 17.4 What does `&input` mean?

```go
json.NewDecoder(r.Body).Decode(&input)
```

The `&` operator obtains the address of `input`.

The decoder needs a destination it can populate. Passing `&input` provides a pointer to that destination.

## 17.5 What do `*string` and `*int` mean?

In PATCH:

```go
Name *string `json:"name"`
Age  *int    `json:"age"`
```

These fields store pointers.

For example:

```go
age := 23
agePointer := &age
```

Now:

```go
*agePointer
```

accesses the integer value `23`.

The two operators serve different roles:

- `&age`: obtain the address of a variable.
- `*agePointer`: access the value through a pointer.

---

# 18. HTTP Status Codes

Your API uses status codes to tell clients the outcome of a request.

| Status | Constant | Meaning | Where used |
|---|---|---|---|
| 200 | `http.StatusOK` | Request succeeded | Successful PUT and PATCH |
| 201 | `http.StatusCreated` | Resource created | Successful POST |
| 204 | `http.StatusNoContent` | Success with no response body | Successful DELETE |
| 400 | `http.StatusBadRequest` | Invalid client input | Invalid ID or JSON |
| 404 | `http.StatusNotFound` | Resource does not exist | Missing user |

Example:

```go
http.Error(w, "User not found", http.StatusNotFound)
```

This sends a 404 response.

Compare:

- `400`: The request is invalid under the API's rules.
- `404`: The request identifies a resource that cannot be found.
- `500`: The server encountered an unexpected internal failure.

Using accurate status codes helps clients decide whether to retry, correct their input, or handle a missing resource.

---

# 19. Complete Request Lifecycle

Let's trace an example request:

```http
PATCH /users/1
Content-Type: application/json

{
  "age": 23
}
```

## Step-by-step execution

1. The HTTP server receives the request.
2. `ServeMux` matches `PATCH /users/{id}`.
3. Go invokes `patchUserHandler`.
4. `r.PathValue("id")` extracts `"1"`.
5. `strconv.Atoi()` converts it to integer `1`.
6. `json.NewDecoder(r.Body).Decode(&input)` decodes the body.
7. The handler verifies that at least one field was supplied.
8. The handler validates the supplied age.
9. `usersMu.Lock()` obtains exclusive access.
10. The handler retrieves the user from the map.
11. The handler changes the age without changing the name.
12. The updated user is saved back to the map.
13. `usersMu.Unlock()` releases the lock.
14. The response `Content-Type` is set to `application/json`.
15. The response status is set to `200 OK`.
16. `json.NewEncoder(w).Encode(user)` sends the updated user to the client.

The important sequence is:

```text
Receive Request
      |
      v
Match Route
      |
      v
Extract and Validate Input
      |
      v
Decode JSON
      |
      v
Lock Shared State
      |
      v
Read / Modify / Save User
      |
      v
Unlock Shared State
      |
      v
Set Response Headers and Status
      |
      v
Encode JSON Response
```

---

# 20. Important Improvements

Your API is a good learning implementation. Before treating it as production-ready, consider the following.

## 20.1 Return an empty array for an empty collection

Current:

```go
var result []User
```

Alternative:

```go
result := make([]User, 0, len(users))
```

This returns `[]` instead of `null` when there are no users.

## 20.2 Normalize names before validating

Current:

```go
if input.Name == "" {
    // Reject.
}
```

A name containing only spaces passes.

A possible improvement is:

```go
name := strings.TrimSpace(input.Name)
if name == "" {
    http.Error(w, "Name is required", http.StatusBadRequest)
    return
}
```

Add the `strings` import if you implement this change.

## 20.3 Keep JSON error responses consistent

Your current `http.Error()` responses are plain text, while successful responses are JSON.

You could later define a JSON error shape:

```json
{
  "error": "User not found"
}
```

A helper function can standardize error responses across handlers.

## 20.4 Consider stricter JSON validation

The current decoder handles ordinary JSON requests, but it does not enforce every possible API validation rule.

Later, consider:

- Limiting request-body size.
- Rejecting unknown fields.
- Detecting trailing JSON data.
- Distinguishing missing fields from explicit `null` when required.
- Validating business rules centrally.

## 20.5 Reduce duplicated ID parsing

Your GET, PUT, PATCH, and DELETE handlers repeat the same ID extraction and conversion code.

Once you understand the current implementation, you can refactor that into a helper function or a shared validation layer.

## 20.6 Add automated tests

The next testing milestone should cover:

- Creating a valid user.
- Rejecting malformed JSON.
- Rejecting invalid IDs.
- Retrieving a missing user.
- Updating only one field with PATCH.
- Deleting a user.
- Concurrent requests.

Go's `net/http/httptest` package is useful for testing handlers without starting a real network server.

## 20.7 Add middleware

The next topic in your learning path is HTTP middleware.

Middleware can add logging, request timing, panic recovery, authentication, and other cross-cutting behavior without repeating the same code in every handler.

## 20.8 Add persistent storage later

Your map is temporary in-memory storage. PostgreSQL will introduce different concerns:

- SQL queries.
- Database connection management.
- Constraints and indexes.
- Transactions.
- Error handling for database operations.
- Persistence across restarts.

Keep the current implementation simple until you are comfortable with the HTTP and concurrency concepts.

---

# 21. Revision Checklist

Use this list to test your understanding.

- [ ] Explain the role of `http.NewServeMux()`.
- [ ] Explain how `HandleFunc()` registers a route.
- [ ] Explain how `r.PathValue("id")` extracts a path parameter.
- [ ] Explain why `strconv.Atoi()` is needed.
- [ ] Explain what the `exists` result means in a map lookup.
- [ ] Explain why `users` is a map rather than a slice.
- [ ] Explain why concurrent map access requires synchronization.
- [ ] Explain the difference between `Lock()` and `RLock()`.
- [ ] Explain when to use `Unlock()` and `RUnlock()`.
- [ ] Explain how `defer` releases the DELETE handler's lock.
- [ ] Explain why `return` follows an error response.
- [ ] Explain why JSON decoding uses `&input`.
- [ ] Explain why PATCH uses pointer fields.
- [ ] Explain the difference between `fmt.Println()` and `fmt.Fprintln(w, ...)`.
- [ ] Explain why JSON responses use `Content-Type: application/json`.
- [ ] Explain the difference between `200`, `201`, `204`, `400`, and `404`.
- [ ] Explain why a 204 response must not have a body.
- [ ] Explain why the in-memory map loses data after a restart.

## Final Summary

Your user management API demonstrates several foundational backend engineering concepts:

1. **Routing:** Directing requests to the appropriate handler.
2. **Request processing:** Reading path parameters and JSON bodies.
3. **Validation:** Rejecting invalid input before modifying data.
4. **Data structures:** Using a map for ID-based lookups.
5. **Concurrency:** Protecting shared mutable state with `sync.RWMutex`.
6. **HTTP semantics:** Returning appropriate status codes, headers, and response bodies.
7. **Serialization:** Converting between Go values and JSON.
8. **Control flow:** Using early returns to stop invalid requests from continuing.

The next logical step is to complete logging middleware, understand how middleware wraps handlers, and then move toward server timeouts, graceful shutdown, automated tests, and PostgreSQL.
