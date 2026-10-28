# Go HTTP Backend Notes --- Request Body to Routing

Detailed revision notes covering the concepts studied from `r.Body`
through JSON handling, validation, response headers, status codes, and
route matching.

## Table of Contents

1.  [HTTP request overview](#1-http-request-overview)
2.  [`r.Body` and `io.ReadAll`](#2-rbody-and-ioreadall)
3.  [Bytes and strings](#3-bytes-and-strings)
4.  [JSON request bodies](#4-json-request-bodies)
5.  [Structs and JSON tags](#5-structs-and-json-tags)
6.  [Decoding JSON](#6-decoding-json)
7.  [Why `&user` is used](#7-why-user-is-used)
8.  [Method checks and validation](#8-method-checks-and-validation)
9.  [JSON syntax, type, and validation
    errors](#9-json-syntax-type-and-validation-errors)
10. [JSON success responses](#10-json-success-responses)
11. [Consistent JSON error
    responses](#11-consistent-json-error-responses)
12. [HTTP status codes](#12-http-status-codes)
13. [Response headers and ordering](#13-response-headers-and-ordering)
14. [`200`, `201`, and `204`](#14-200-201-and-204)
15. [`404` vs `405`](#15-404-vs-405)
16. [Go `ServeMux` route matching](#16-go-servemux-route-matching)
17. [Path routes such as `/users/123`](#17-path-routes-such-as-users123)
18. [Complete handler example](#18-complete-handler-example)
19. [Symbols to revise](#19-symbols-to-revise)
20. [Practice checklist](#20-practice-checklist)

------------------------------------------------------------------------

## 1. HTTP request overview

An HTTP request contains a method, URL, headers, and sometimes a body.

``` text
HTTP Request
├── Method: GET, POST, PUT, PATCH, DELETE
├── URL
│   ├── Path: /users/123
│   └── Query: ?active=true
├── Headers: Content-Type, Accept, User-Agent, ...
└── Body: JSON or other data sent by the client
```

A Go HTTP handler receives the request and writes the response:

``` go
func handler(w http.ResponseWriter, r *http.Request) {
    // r: incoming request
    // w: outgoing response
}
```

-   `r *http.Request` contains request information from the client.
-   `w http.ResponseWriter` is used to write the response.

## 2. `r.Body` and `io.ReadAll`

`r.Body` is the incoming request body. It implements `io.ReadCloser` and
can be read as a stream. Request bodies are commonly used with `POST`,
`PUT`, and `PATCH`.

``` go
func handleBody(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }

    fmt.Fprintln(w, string(body))
}
```

Imports for this snippet:

``` go
import (
    "fmt"
    "io"
    "net/http"
)
```

Important names: - `r.Body`: incoming body stream. -
`io.ReadAll(r.Body)`: reads the body and returns `[]byte`. -
`string(body)`: converts the bytes to a string. -
`fmt.Fprintln(w, ...)`: writes text to the HTTP response. -
`http.Error(w, message, status)`: writes a plain-text error response and
status.

`io.ReadAll` reads the body into memory. For production APIs, consider
request-size limits to avoid reading arbitrarily large bodies.

## 3. Bytes and strings

`io.ReadAll` returns a byte slice: `[]byte`. Go displays this type as
`[]uint8` because `byte` is an alias for `uint8`.

``` go
body, err := io.ReadAll(r.Body)
if err != nil {
    http.Error(w, "Failed to read request body", http.StatusBadRequest)
    return
}

fmt.Fprintf(w, "Body type: %T\n", body)
fmt.Fprintln(w, "Body content:", body)
fmt.Fprintln(w, "Body as string:", string(body))
```

For plain text `Hello`, byte values begin:

``` text
[72 101 108 108 111]
```

`72` is the byte value for `H`.

### Why did the body sometimes start with `34`?

In Postman, the JSON string value:

``` json
"Hello from POST body"
```

includes quote characters. A double quote has byte value `34`, so the
body starts with `34`.

In Text mode, sending:

``` text
Hello from POST body
```

without quotation marks means the first byte is `72` (`H`).

**Key idea:** `io.ReadAll` reads the bytes sent by the client. It does
not remove quotes or parse JSON.

## 4. JSON request bodies

JSON (JavaScript Object Notation) is a text format for exchanging
structured data.

``` json
{
  "name": "Ankush",
  "age": 21
}
```

In Postman, select `POST`, then **Body → raw → JSON**. The request
should have:

``` http
Content-Type: application/json
```

The request `Content-Type` tells the server the format of the body. It
does not parse the body automatically.

## 5. Structs and JSON tags

Represent the JSON with a Go struct:

``` go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

-   `type User struct`: defines a struct type named `User`.
-   `Name string`: a text field.
-   `Age int`: an integer field.
-   `` `json:"name"` `` maps JSON key `name` to Go field `Name`.
-   `` `json:"age"` `` maps JSON key `age` to Go field `Age`.

The fields are exported because their names begin with capital letters.
JSON tags allow JSON keys to remain lowercase.

## 6. Decoding JSON

Import `encoding/json` and decode the request body directly:

``` go
var user User

err := json.NewDecoder(r.Body).Decode(&user)
if err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
}
```

-   `json.NewDecoder(r.Body)` creates a decoder that reads JSON from the
    request stream.
-   `Decode(&user)` decodes the JSON into `user`.
-   A non-nil `err` means decoding failed.
-   `return` stops the handler after sending the error response.

This is often simpler than reading all bytes first and then decoding
them separately. Related function: `json.Unmarshal(data, &user)`, which
decodes from an existing byte slice.

## 7. Why `&user` is used

``` go
var user User
err := json.NewDecoder(r.Body).Decode(&user)
```

`user` is the variable to populate. `&user` means "take the address of
`user`," giving a pointer of type `*User`. The decoder needs a
destination it can update.

Before decoding, Go's zero values are:

``` text
user.Name = ""
user.Age  = 0
```

After decoding the example JSON:

``` text
user.Name = "Ankush"
user.Age  = 21
```

Remember:

``` go
user   // struct value
&user  // pointer to the struct
```

Pointers are a larger Go topic; for this handler, remember to pass
`&user` to `Decode`.

## 8. Method checks and validation

A JSON endpoint should check its HTTP method before trying to decode the
body:

``` go
if r.Method != http.MethodPost {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
```

Useful constants include `http.MethodGet`, `http.MethodPost`,
`http.MethodPut`, `http.MethodPatch`, and `http.MethodDelete`.

A browser address-bar visit normally sends `GET`. It does not
automatically send the JSON body entered in Postman. A `POST`-only
endpoint should return `405 Method Not Allowed` for a GET request.

After decoding, validate application data:

``` go
if user.Name == "" {
    http.Error(w, "Name is required", http.StatusBadRequest)
    return
}

if user.Age <= 0 {
    http.Error(w, "Age must be greater than 0", http.StatusBadRequest)
    return
}
```

A missing `age` field becomes `0` because `0` is the zero value for
`int`, so the age check rejects missing age as well.

## 9. JSON syntax, type, and validation errors

These are three different failure categories.

### A. Malformed JSON

``` json
{
  "name": "Ankush",
  "age":
}
```

The JSON syntax is invalid, so `Decode` returns an error.

### B. Type mismatch

The struct expects `Age int`, but the request contains:

``` json
{
  "name": "Ankush",
  "age": "twenty-one"
}
```

This is syntactically valid JSON, but `age` is a string instead of a
number. An error may say:

``` text
json: cannot unmarshal string into Go struct field User.age of type int
```

### C. Validation error

``` json
{
  "name": "",
  "age": 21
}
```

This JSON can decode successfully, but the application rejects the empty
name.

A controlled decoding error response:

``` go
err := json.NewDecoder(r.Body).Decode(&user)
if err != nil {
    http.Error(w, "Request body contains invalid or incompatible JSON", http.StatusBadRequest)
    return
}
```

During learning, `err.Error()` helps inspect the failure. Public APIs
should generally return controlled messages instead of exposing
implementation details.

## 10. JSON success responses

To return a JSON response rather than plain text, use
`json.NewEncoder(w)`:

``` go
w.Header().Set("Content-Type", "application/json")

err = json.NewEncoder(w).Encode(user)
if err != nil {
    return
}
```

With the `User` struct and JSON tags above, the response is:

``` json
{
  "name": "Ankush",
  "age": 21
}
```

  Symbol                      Direction   Purpose
  --------------------------- ----------- ------------------------
  `json.NewDecoder(r.Body)`   Incoming    Read JSON from request
  `json.NewEncoder(w)`        Outgoing    Write JSON to response

`Encode` appends a newline after the encoded JSON value.

## 11. Consistent JSON error responses

A success response may be JSON while `http.Error` returns plain text. To
keep errors in JSON too, define:

``` go
type APIError struct {
    Error string `json:"error"`
}
```

Create a helper:

``` go
func writeJSONError(w http.ResponseWriter, message string, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    _ = json.NewEncoder(w).Encode(APIError{
        Error: message,
    })
}
```

Example use:

``` go
if user.Name == "" {
    writeJSONError(w, "Name is required", http.StatusBadRequest)
    return
}
```

Response body:

``` json
{
  "error": "Name is required"
}
```

The status is `400 Bad Request`, and `Content-Type` is
`application/json`.

Use the helper for other errors as well:

``` go
writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
writeJSONError(w, "Age must be greater than 0", http.StatusBadRequest)
```

The example ignores the encoder error for brevity. In production,
consider logging or handling response-writing failures.

## 12. HTTP status codes

  -------------------------------------------------------------------------------------
                        Status Go constant                        Typical meaning
  ---------------------------- ---------------------------------- ---------------------
                           200 `http.StatusOK`                    Request succeeded

                           201 `http.StatusCreated`               Resource created

                           204 `http.StatusNoContent`             Success with no
                                                                  response body

                           400 `http.StatusBadRequest`            Invalid request/input

                           401 `http.StatusUnauthorized`          Authentication
                                                                  required or invalid

                           403 `http.StatusForbidden`             Client lacks
                                                                  permission

                           404 `http.StatusNotFound`              Route/resource not
                                                                  found

                           405 `http.StatusMethodNotAllowed`      Method unsupported
                                                                  for this route

                           500 `http.StatusInternalServerError`   Unexpected
                                                                  server-side failure
  -------------------------------------------------------------------------------------

Status codes communicate the outcome; the response body can provide more
detail.

## 13. Response headers and ordering

Response headers provide metadata. For JSON:

``` go
w.Header().Set("Content-Type", "application/json")
```

A custom header can also be set:

``` go
w.Header().Set("X-Learning-Topic", "HTTP-Headers")
```

In Postman, inspect the response **Headers** tab to see it.

A normal response order is:

``` go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)

err := json.NewEncoder(w).Encode(user)
if err != nil {
    return
}
```

Remember: 1. `w.Header().Set(...)` sets headers. 2.
`w.WriteHeader(status)` commits the status and headers. 3. Writing the
body also commits the response if not already committed.

If the body is written first, Go normally commits `200 OK`. Once the
response is committed, changing its status or headers is too late.

`http.Error` is a convenient plain-text error helper. Use a JSON helper
when the error body must be JSON.

## 14. `200`, `201`, and `204`

### `200 OK`

A successful request with a response body is appropriate for many
operations:

``` go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(user)
```

### `201 Created`

Use when a resource was successfully created. A real create endpoint
typically saves the resource before returning this status. Merely
echoing the submitted user does not actually create a persistent
resource.

``` go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(user)
```

### `204 No Content`

Use when the request succeeded and no body should be returned, often
after deletion:

``` go
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
```

Register it with:

``` go
http.HandleFunc("/deleteUser", handleDeleteUser)
```

Test `DELETE /deleteUser`. Expected: `204 No Content` and an empty body.
This practice handler does not actually delete a database record. A
`204` response must not contain response content.

## 15. `404` vs `405`

-   `404 Not Found`: the requested path/resource was not found.
-   `405 Method Not Allowed`: the path is recognized, but the method is
    unsupported.

Method-check example:

``` go
if r.Method != http.MethodDelete {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
```

404 example:

``` go
http.NotFound(w, r)
return
```

`http.NotFound(w, r)` sends a 404 response.

## 16. Go `ServeMux` route matching

Register routes with `http.HandleFunc`:

``` go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/users", usersHandler)
http.HandleFunc("/products", productsHandler)
```

In Go's default `http.ServeMux`, `/` is a catch-all pattern and can
match paths not matched by more specific registered patterns. Therefore
`/unknown` can reach `homeHandler` instead of automatically returning
404.

To ensure the home handler serves only `/`:

``` go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    fmt.Fprintln(w, "Welcome to my API")
}
```

`http.NotFound(w, r)` sends `404 Not Found`. The `return` prevents the
handler from continuing to write a success response.

## 17. Path routes such as `/users/123`

Register a subtree pattern:

``` go
http.HandleFunc("/users/", userByIDHandler)
```

Handler:

``` go
func userByIDHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Path:", r.URL.Path)
}
```

Request:

``` http
GET /users/123
```

Response:

``` text
Path: /users/123
```

In the default `ServeMux`, the trailing-slash pattern `/users/` matches
paths beneath `/users/`. Registering `/users` alone does not implement
`/users/123`.

This handler only prints the path; it does not extract or validate the
ID yet. The next step is to extract `123` from the path.

## 18. Complete handler example

This example combines method checks, JSON decoding, validation, JSON
errors, headers, and a JSON success response.

``` go
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type APIError struct {
    Error string `json:"error"`
}

func writeJSONError(w http.ResponseWriter, message string, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    _ = json.NewEncoder(w).Encode(APIError{
        Error: message,
    })
}

func handleBody(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        writeJSONError(
            w,
            "Request body contains invalid or incompatible JSON",
            http.StatusBadRequest,
        )
        return
    }

    if user.Name == "" {
        writeJSONError(w, "Name is required", http.StatusBadRequest)
        return
    }

    if user.Age <= 0 {
        writeJSONError(w, "Age must be greater than 0", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Learning-Topic", "HTTP-Headers")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(user); err != nil {
        // The response may already be committed. A real application
        // should log or otherwise handle response-writing failures.
        return
    }
}

func main() {
    http.HandleFunc("/inspectBody", handleBody)
    http.ListenAndServe(":8080", nil)
}
```

This is a learning example. Production APIs should also consider
request-body size limits, rejecting unexpected extra JSON values,
authentication/authorization when needed, logging, and more complete
validation.

### Postman test

Method: `POST`

URL:

``` text
http://localhost:8080/inspectBody
```

Header:

``` http
Content-Type: application/json
```

Valid body:

``` json
{
  "name": "Ankush",
  "age": 21
}
```

Expected response:

``` json
{
  "name": "Ankush",
  "age": 21
}
```

Invalid body example:

``` json
{
  "name": "",
  "age": 21
}
```

Expected error:

``` json
{
  "error": "Name is required"
}
```

Status: `400 Bad Request`.

## 19. Symbols to revise

### Request body and I/O

-   `http.Request`
-   `r.Body`
-   `io.ReadAll`
-   `io.ReadCloser`
-   `io.Writer`
-   `[]byte`
-   `byte`
-   `uint8`

### JSON

-   `encoding/json`
-   `json.NewDecoder`
-   `(*json.Decoder).Decode`
-   `json.NewEncoder`
-   `(*json.Encoder).Encode`
-   `json.Unmarshal`
-   JSON struct tags

### Structs and pointers

-   `struct`
-   `&` address-of operator
-   `*User` pointer type
-   zero values

### HTTP methods and errors

-   `r.Method`
-   `http.MethodGet`
-   `http.MethodPost`
-   `http.MethodPut`
-   `http.MethodPatch`
-   `http.MethodDelete`
-   `http.Error`
-   `http.NotFound`
-   `http.StatusBadRequest`
-   `http.StatusMethodNotAllowed`
-   `http.StatusInternalServerError`

### Responses

-   `http.ResponseWriter`
-   `w.Header()`
-   `Header.Set`
-   `w.WriteHeader`
-   `w.Write`
-   `Content-Type`
-   `application/json`

### Routing

-   `http.HandleFunc`
-   `http.ServeMux`
-   catch-all `/`
-   subtree pattern `/users/`
-   `r.URL.Path`

## 20. Practice checklist

-   [ ] Read a request body with `io.ReadAll(r.Body)`.
-   [ ] Explain why `io.ReadAll` returns `[]byte`.
-   [ ] Explain why a JSON string can begin with quote byte `34`.
-   [ ] Handle read errors rather than discarding them.
-   [ ] Set `Content-Type: application/json` in Postman.
-   [ ] Define a struct with JSON tags.
-   [ ] Decode JSON with `json.NewDecoder(r.Body).Decode(&user)`.
-   [ ] Explain at a basic level why `&user` is passed to `Decode`.
-   [ ] Reject unsupported methods with `405`.
-   [ ] Validate required fields and values.
-   [ ] Distinguish malformed JSON, type mismatch, and validation
    errors.
-   [ ] Return JSON with `json.NewEncoder(w).Encode(...)`.
-   [ ] Return JSON errors using a helper.
-   [ ] Explain why headers/status must be set before writing the
    response body.
-   [ ] Explain `200`, `201`, and `204`.
-   [ ] Explain the difference between `404` and `405`.
-   [ ] Explain why `/` can match unknown paths in the default mux.
-   [ ] Register and test a `/users/` path handler.

------------------------------------------------------------------------

## Final mental model

``` text
Incoming HTTP Request
        |
        v
Check r.Method
        |
        v
Read r.Body
        |
        v
Decode JSON into a Go struct
        |
        v
Validate the values
        |
        +---- invalid ----> JSON error + suitable status
        |
        v
Set response headers
        |
        v
Set response status
        |
        v
Encode JSON response
```

Core pattern:

``` go
var user User

if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
    return
}
```

This reads JSON from the request, populates a Go struct, and lets the
handler validate the data before responding.
