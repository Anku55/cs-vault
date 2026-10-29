# Go Backend Development — Notes

## Table of Contents

1. [Path Parameters in Go](#1-path-parameters-in-go)
2. [Extracting and Validating IDs](#2-extracting-and-validating-ids)
3. [Path Parameters vs Query Parameters](#3-path-parameters-vs-query-parameters)
4. [Go Maps for In-Memory Storage](#4-go-maps-for-in-memory-storage)
5. [Concurrency and Mutexes](#5-concurrency-and-mutexes)
6. [Building a CRUD REST API](#6-building-a-crud-rest-api)
7. [POST — Create a User](#7-post--create-a-user)
8. [GET — Retrieve All Users](#8-get--retrieve-all-users)
9. [GET — Retrieve a User by ID](#9-get--retrieve-a-user-by-id)
10. [DELETE — Delete a User](#10-delete--delete-a-user)
11. [PUT — Update a Complete Resource](#11-put--update-a-complete-resource)
12. [PATCH — Update Selected Fields](#12-patch--update-selected-fields)
13. [HTTP Status Codes](#13-http-status-codes)
14. [JSON Encoding and Decoding](#14-json-encoding-and-decoding)
15. [HTTP Middleware](#15-http-middleware)
16. [Logging Middleware Implementation](#16-logging-middleware-implementation)
17. [Request Lifecycle and Middleware Flow](#17-request-lifecycle-and-middleware-flow)
18. [Testing the API](#18-testing-the-api)
19. [Important Go Concepts Used](#19-important-go-concepts-used)
20. [Current Learning Roadmap](#20-current-learning-roadmap)

---

# 1. Path Parameters in Go

## 1.1 What are path parameters?

Path parameters are values embedded inside a URL path that identify a particular resource.

For example:

```text
GET /users/1
GET /users/25
GET /users/100
```

Here, `1`, `25`, and `100` are user IDs.

Instead of creating a separate route for every user, we define one route that accepts a variable value.

In Go 1.22 and later, the standard `net/http` package supports method-aware routes and named wildcards.

```go
mux.HandleFunc("GET /users/{id}", getUserByIDHandler)
```

Explanation:

- `GET` specifies the HTTP method.
- `/users/` is the fixed path.
- `{id}` is a wildcard path segment.
- `getUserByIDHandler` handles matching requests.

Examples that match:

```text
GET /users/1
GET /users/25
GET /users/abc
```

The final example matches the route, but `abc` is not a valid integer ID. We must validate it inside the handler.

## 1.2 Extracting a path parameter with `PathValue`

Go provides `r.PathValue()` to retrieve the value of a named wildcard.

```go
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("id")

    fmt.Println(idString)

    fmt.Fprintln(w, "User ID:", idString)
}
```

For this request:

```http
GET /users/25
```

The result is:

```text
idString = "25"
```

**Important:** `PathValue()` returns a string, not an integer.

Even though the URL contains a number, HTTP path values are text.

## 1.3 When should you use path parameters?

Use path parameters when identifying a specific resource.

Examples:

```text
GET    /users/10
PUT    /users/10
PATCH  /users/10
DELETE /users/10
GET    /products/42
GET    /orders/100
```

The ID identifies the resource on which the operation acts.

---

# 2. Extracting and Validating IDs

A path parameter is initially a string. If the application expects an integer ID, convert and validate it.

## 2.1 Convert a string into an integer

Go's `strconv` package provides `strconv.Atoi()`.

```go
import "strconv"

idString := r.PathValue("id")

id, err := strconv.Atoi(idString)
if err != nil {
    http.Error(w, "Invalid user ID", http.StatusBadRequest)
    return
}
```

Explanation:

- `strconv.Atoi()` converts a string to an integer.
- `id` contains the converted integer if successful.
- `err` indicates whether conversion failed.
- `http.StatusBadRequest` represents HTTP 400.
- `return` stops the handler from continuing after invalid input.

For example:

| Path value | Result |
|---|---|
| `"12"` | Integer `12` |
| `"0"` | Integer `0` |
| `"-5"` | Integer `-5` |
| `"abc"` | Conversion error |
| `"12abc"` | Conversion error |

Conversion alone does not establish that an ID is valid. An application may also require the ID to be positive.

## 2.2 Complete ID validation

```go
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("id")

    id, err := strconv.Atoi(idString)
    if err != nil || id <= 0 {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Continue with the database or storage lookup.
    fmt.Fprintln(w, "Valid user ID:", id)
}
```

This handler checks two conditions:

1. The ID must be convertible to an integer.
2. The ID must be greater than zero.

The `||` operator means OR. If either condition is true, the request is rejected.

## 2.3 Why is `return` important?

Consider this incorrect pattern:

```go
if err != nil {
    http.Error(w, "Invalid user ID", http.StatusBadRequest)
}

// The handler continues even after the error.
fmt.Fprintln(w, "Processing user")
```

The handler may continue processing invalid input.

Correct approach:

```go
if err != nil {
    http.Error(w, "Invalid user ID", http.StatusBadRequest)
    return
}
```

**Rule:** After writing an error response, return immediately unless there is a specific reason to continue.

## 2.4 `http.Error()`

```go
http.Error(w, "User not found", http.StatusNotFound)
```

This function writes an error message and the specified HTTP status code.

Common examples:

```go
http.Error(w, "Invalid input", http.StatusBadRequest)
http.Error(w, "Unauthorized", http.StatusUnauthorized)
http.Error(w, "User not found", http.StatusNotFound)
http.Error(w, "Internal server error", http.StatusInternalServerError)
```

By default, `http.Error()` produces a plain-text error response, not a JSON error object.

---

# 3. Path Parameters vs Query Parameters

Both carry information through a URL, but they serve different purposes.

## 3.1 Path parameters

Example:

```text
GET /users/25
```

Route:

```go
mux.HandleFunc("GET /users/{id}", getUserByIDHandler)
```

Read the value:

```go
id := r.PathValue("id")
```

Typical purpose: identify one specific resource.

## 3.2 Query parameters

Example:

```text
GET /users?age=20
GET /users?page=2&limit=10
GET /users?name=Ankit
```

Read query parameters using:

```go
age := r.URL.Query().Get("age")
```

For pagination:

```go
page := r.URL.Query().Get("page")
limit := r.URL.Query().Get("limit")
```

A query parameter is not automatically converted to an integer. Use `strconv.Atoi()` if an integer is required.

## 3.3 Useful query methods

### `r.URL.Query()`

Returns query values as a `url.Values` value, which is map-like.

```go
query := r.URL.Query()
```

### `query.Get("name")`

Returns the first value associated with the key, or an empty string if no value exists.

```go
name := r.URL.Query().Get("name")
```

### Direct map-style access

```go
values := r.URL.Query()["name"]
```

Because a query key can occur more than once, the value is a slice of strings.

For example:

```text
GET /users?tag=go&tag=backend
```

Conceptually:

```go
map[string][]string{
    "tag": {"go", "backend"},
}
```

To read all values:

```go
tags := r.URL.Query()["tag"]
```

To read just the first value, use `Get()`.

## 3.4 Which one should you choose?

| Requirement | Example |
|---|---|
| Retrieve a specific user | `/users/25` |
| Filter users by age | `/users?age=20` |
| Paginate a list | `/users?page=2&limit=10` |
| Search users by name | `/users?name=Ankit` |
| Delete a specific user | `DELETE /users/25` |

**Rule of thumb:** Path parameters identify resources; query parameters commonly filter, search, sort, or paginate collections.

---

# 4. Go Maps for In-Memory Storage

Before connecting a database, we used a Go map to store users temporarily.

## 4.1 What is a map?

A map stores key-value pairs.

```go
users := make(map[int]string)

users[1] = "Ankit"
users[2] = "Rahul"

fmt.Println(users[1])
```

Output:

```text
Ankit
```

In this example:

- Key type: `int`
- Value type: `string`

For the user management API, we used a map whose values are `User` structs.

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var users = make(map[int]User)
```

The map associates each user ID with a complete user object.

Conceptually:

```text
users
  |
  +-- 1 -> User{ID: 1, Name: "Ankit", Age: 21}
  |
  +-- 2 -> User{ID: 2, Name: "Rahul", Age: 22}
```

## 4.2 Insert a value

```go
user := User{
    ID:   1,
    Name: "Ankit",
    Age:  21,
}

users[user.ID] = user
```

The ID becomes the map key.

## 4.3 Retrieve a value

```go
user, exists := users[1]

if exists {
    fmt.Println(user.Name)
}
```

A map lookup can return two values:

```go
value, exists := mapName[key]
```

- `value` contains the stored value if the key exists.
- `exists` is `true` if the key is present.

This distinction matters because a missing key returns the value type's zero value. Checking only the value can be ambiguous.

## 4.4 Delete a value

```go
delete(users, 1)
```

`delete()` removes the entry for the specified key. Deleting a missing key is safe and does nothing.

## 4.5 Why use a map instead of a slice?

A slice-based search may require checking each user until the ID matches.

A map provides direct key-based lookup, which is typically average-case \(O(1)\).

| Operation | Map | Slice |
|---|---|---|
| Find by ID | Average \(O(1)\) | Usually \(O(n)\) |
| Insert by ID | Average \(O(1)\) | Depends on implementation |
| Preserve order | No guaranteed order | Yes, by index |
| Simple ID lookup | Convenient | Requires searching |

Map iteration order is unspecified. If an API needs sorted output, sort the slice before encoding it.

## 4.6 Generate IDs

We used a counter:

```go
var nextID = 1
```

When creating a user:

```go
user.ID = nextID
nextID++
```

This is suitable for a small in-memory learning project. It is not persistent: restarting the process loses the map and resets the counter. A real database can provide durable storage and an appropriate ID-generation strategy.

---

# 5. Concurrency and Mutexes

Go HTTP servers can handle multiple requests concurrently. This means multiple handlers may access shared data at the same time.

If two requests modify a regular map concurrently, the program can encounter a data race or a concurrent map access failure.

We used `sync.RWMutex` to protect shared state.

## 5.1 Define a mutex

```go
import "sync"

var (
    users   = make(map[int]User)
    nextID  = 1
    usersMu sync.RWMutex
)
```

`sync.RWMutex` supports two types of locking:

- `Lock()` / `Unlock()` — exclusive access, suitable for writes.
- `RLock()` / `RUnlock()` — shared read access, suitable for reads.

## 5.2 Reading data

```go
usersMu.RLock()
user, exists := users[id]
usersMu.RUnlock()
```

A read lock allows multiple readers to access the protected data simultaneously, provided no writer holds the exclusive lock.

## 5.3 Writing data

```go
usersMu.Lock()
users[id] = user
usersMu.Unlock()
```

An exclusive lock prevents other readers and writers from accessing the protected data while the lock is held.

## 5.4 Using `defer`

You may see this pattern:

```go
usersMu.RLock()
defer usersMu.RUnlock()

user, exists := users[id]
```

`defer` schedules a function call to run when the surrounding function returns.

This helps ensure the lock is released even when the function has multiple return paths.

However, for short critical sections, release the lock before doing slow work such as writing an HTTP response.

## 5.5 Keep related operations atomic

This pattern is unsafe if the lookup and deletion are separated by an unlock:

```go
usersMu.RLock()
_, exists := users[id]
usersMu.RUnlock()

if exists {
    usersMu.Lock()
    delete(users, id)
    usersMu.Unlock()
}
```

Another request could modify the map between the lookup and deletion.

Instead, perform the existence check and deletion under the same exclusive lock:

```go
usersMu.Lock()

_, exists := users[id]
if exists {
    delete(users, id)
}

usersMu.Unlock()
```

**Important:** Protect the shared map and the `nextID` counter consistently. Every access to those shared variables must follow the same synchronization policy.

---

# 6. Building a CRUD REST API

CRUD represents the four fundamental data operations:

- **C — Create**
- **R — Read**
- **U — Update**
- **D — Delete**

We implemented these operations for a `User` resource.

## 6.1 CRUD endpoint design

| HTTP method | Endpoint | Purpose | Success status |
|---|---|---|---|
| `POST` | `/users` | Create a user | `201 Created` |
| `GET` | `/users` | Retrieve all users | `200 OK` |
| `GET` | `/users/{id}` | Retrieve one user | `200 OK` |
| `PUT` | `/users/{id}` | Replace editable user fields | `200 OK` |
| `PATCH` | `/users/{id}` | Update selected fields | `200 OK` |
| `DELETE` | `/users/{id}` | Delete a user | `204 No Content` |

This is the API design used in our learning project.

## 6.2 Define the data model

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

The JSON struct tags define how fields appear in JSON.

For example:

```json
{
  "id": 1,
  "name": "Ankit",
  "age": 21
}
```

The API uses JSON as its request and response format.

## 6.3 Register routes

```go
func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", healthHandler)
    mux.HandleFunc("POST /users", createUserHandler)
    mux.HandleFunc("GET /users", getUsersHandler)
    mux.HandleFunc("GET /users/{id}", getUserByIDHandler)
    mux.HandleFunc("PUT /users/{id}", updateUserHandler)
    mux.HandleFunc("PATCH /users/{id}", patchUserHandler)
    mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

`http.NewServeMux()` creates a request multiplexer, which matches incoming requests to registered handlers.

`http.ListenAndServe()` starts the HTTP server on port `8080` and serves requests using the supplied mux.

`log.Fatal()` logs an error and terminates the program if `ListenAndServe()` returns an error.

---

# 7. POST — Create a User

## 7.1 What does POST do?

A `POST /users` request creates a new user.

Example request body:

```json
{
  "name": "Ankit",
  "age": 21
}
```

The client does not provide the ID in this design. The server generates it.

## 7.2 Decode the request body

```go
var input struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

err := json.NewDecoder(r.Body).Decode(&input)
if err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
}
```

`json.NewDecoder(r.Body)` creates a JSON decoder that reads the request body.

`Decode(&input)` parses the JSON and populates the struct.

The `&input` expression supplies a pointer to the destination, allowing the decoder to populate it.

## 7.3 Validate the input

```go
if input.Name == "" {
    http.Error(w, "Name is required", http.StatusBadRequest)
    return
}

if input.Age <= 0 {
    http.Error(w, "Age must be positive", http.StatusBadRequest)
    return
}
```

Validation ensures the input satisfies the rules defined by our application.

For a production API, consider whether whitespace-only names should be rejected, whether age has an upper bound, and how malformed or unknown JSON fields should be handled.

## 7.4 Store the user and return JSON

```go
usersMu.Lock()

user := User{
    ID:   nextID,
    Name: input.Name,
    Age:  input.Age,
}

users[user.ID] = user
nextID++

usersMu.Unlock()

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)

if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Printf("Failed to encode created user: %v", err)
}
```

The server creates the user, stores it in the map, and returns the created object.

The successful response has status `201 Created`.

### Important details

- `w.Header().Set()` sets an HTTP response header.
- `w.WriteHeader()` sets the response status.
- `json.NewEncoder(w).Encode(user)` writes JSON to the HTTP response.
- A response body should be written only after the relevant status and headers have been selected.

For robust production code, finish all request validation before acquiring the lock and keep the critical section short.

---

# 8. GET — Retrieve All Users

## 8.1 Handler implementation

```go
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    usersMu.RLock()

    result := make([]User, 0, len(users))

    for _, user := range users {
        result = append(result, user)
    }

    usersMu.RUnlock()

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(result); err != nil {
        log.Printf("Failed to encode users: %v", err)
    }
}
```

## 8.2 Why convert the map into a slice?

The map is useful for storing and retrieving users by ID. The slice is convenient for returning a collection in JSON.

This code:

```go
for _, user := range users {
    result = append(result, user)
}
```

iterates over map values and appends each user to the result slice.

## 8.3 Why use `make([]User, 0, len(users))`?

```go
result := make([]User, 0, len(users))
```

This creates:

- A slice of type `[]User`.
- An initial length of `0`.
- A capacity equal to the current number of users.

If the map is empty, JSON encoding produces:

```json
[]
```

By contrast, a nil slice generally encodes as:

```json
null
```

For a collection endpoint, returning an empty array is often a more convenient API contract.

## 8.4 Why release the lock before encoding?

JSON encoding is response work, not map access.

Copy the data while holding the read lock, then release the lock before writing the response. This reduces how long other requests must wait for access to the map.

Remember that map iteration order is unspecified. The returned array may not have a stable order unless you sort it.

---

# 9. GET — Retrieve a User by ID

## 9.1 Handler implementation

```go
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("id")

    id, err := strconv.Atoi(idString)
    if err != nil || id <= 0 {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    usersMu.RLock()
    user, exists := users[id]
    usersMu.RUnlock()

    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(user); err != nil {
        log.Printf("Failed to encode user: %v", err)
    }
}
```

## 9.2 Execution flow

1. Read the ID from the path.
2. Convert it into an integer.
3. Reject invalid or nonpositive IDs.
4. Acquire the read lock.
5. Look up the user in the map.
6. Release the lock.
7. Return `404 Not Found` if the user does not exist.
8. Otherwise, encode the user as JSON.

## 9.3 Why use a map lookup?

```go
user, exists := users[id]
```

This is preferable to iterating through every user because the ID is already the map key.

Do not confuse these two cases:

- Invalid ID syntax: `400 Bad Request`.
- Valid ID that does not exist: `404 Not Found`.

---

# 10. DELETE — Delete a User

## 10.1 Handler implementation

```go
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("id")

    id, err := strconv.Atoi(idString)
    if err != nil || id <= 0 {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    usersMu.Lock()

    _, exists := users[id]
    if exists {
        delete(users, id)
    }

    usersMu.Unlock()

    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
```

## 10.2 Why use an exclusive lock?

Deleting changes shared state, so we use:

```go
usersMu.Lock()
```

The existence check and deletion occur under the same lock. This prevents another handler from changing the map between those operations.

## 10.3 Why return `204 No Content`?

`204 No Content` indicates that the operation succeeded and the response has no body.

Correct:

```go
w.WriteHeader(http.StatusNoContent)
```

Avoid this:

```go
w.WriteHeader(http.StatusNoContent)
fmt.Fprintln(w, "User deleted")
```

A `204` response must not contain a response body.

---

# 11. PUT — Update a Complete Resource

## 11.1 PUT semantics

In our project, `PUT /users/{id}` updates the user's editable fields: `name` and `age`.

Example:

```http
PUT /users/1
Content-Type: application/json
```

```json
{
  "name": "Ankit Sharma",
  "age": 22
}
```

Our implementation requires both editable fields. It does not let the client replace the ID.

## 11.2 Handler implementation

```go
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
    idString := r.PathValue("id")

    id, err := strconv.Atoi(idString)
    if err != nil || id <= 0 {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var input struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if input.Name == "" || input.Age <= 0 {
        http.Error(w, "Valid name and positive age are required", http.StatusBadRequest)
        return
    }

    usersMu.Lock()

    user, exists := users[id]
    if exists {
        user.Name = input.Name
        user.Age = input.Age
        users[id] = user
    }

    usersMu.Unlock()

    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(user); err != nil {
        log.Printf("Failed to encode updated user: %v", err)
    }
}
```

## 11.3 Why read, modify, and save?

```go
user, exists := users[id]

if exists {
    user.Name = input.Name
    user.Age = input.Age
    users[id] = user
}
```

A map lookup returns a copy of the stored struct value. Updating the local `user` variable does not automatically update the map entry, so we assign the modified struct back to the map.

The lookup, modification, and save occur under one exclusive lock.

---

# 12. PATCH — Update Selected Fields

## 12.1 PATCH semantics

`PATCH /users/{id}` changes only the fields included in the request.

For example, changing only the age:

```json
{
  "age": 23
}
```

The name should remain unchanged.

## 12.2 Why use pointer fields?

We used pointers to distinguish an omitted field from an ordinary zero value.

