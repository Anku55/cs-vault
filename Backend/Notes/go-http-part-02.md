# Go HTTP — 
## Topic: `net/http` Request Handling

---

# 1. HTTP Mental Model

The basic HTTP communication model:

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

An HTTP **Request** contains:

```text
Method
URL
Headers
Body
```

An HTTP **Response** contains:

```text
Status Code
Headers
Body
```

---

# 2. Go HTTP Package

Go's standard HTTP package is:

```go
"net/http"
```

Important types/functions learned:

```go
http.Request
http.ResponseWriter
http.HandleFunc
http.ListenAndServe
http.Error
```

---

# 3. `http.Request`

Inside a handler:

```go
func handler(w http.ResponseWriter, r *http.Request) {
}
```

`r` represents the **incoming HTTP request**.

Think:

```text
r = everything the client sent to the server
```

Important fields:

```go
r.Method
r.URL
r.Header
r.Body
```

We have covered the first three in detail and are about to start `r.Body`.

---

# 4. `r.Method`

`r.Method` tells you which HTTP method the client used.

Example:

```go
fmt.Fprintln(w, r.Method)
```

Possible values:

```text
GET
POST
PUT
PATCH
DELETE
```

Prefer Go's predefined constants:

```go
http.MethodGet
http.MethodPost
http.MethodPut
http.MethodPatch
http.MethodDelete
```

Example:

```go
if r.Method == http.MethodGet {
    fmt.Fprintln(w, "GET request")
}
```

---

# 5. Handling Multiple HTTP Methods

Instead of a long `if/else if` chain, Go can use a `switch`:

```go
switch r.Method {
case http.MethodGet:
    fmt.Fprintln(w, "GET")

case http.MethodPost:
    fmt.Fprintln(w, "POST")

case http.MethodDelete:
    fmt.Fprintln(w, "DELETE")

case http.MethodPatch:
    fmt.Fprintln(w, "PATCH")

default:
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
```

## Important Go behavior

Go `switch` cases **do not automatically fall through**.

You normally don't write:

```go
break
```

after every case.

---

# 6. HTTP 405 — Method Not Allowed

When a server understands the request but doesn't allow that HTTP method for a route:

```go
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
```

Important symbol:

```go
http.StatusMethodNotAllowed
```

This represents:

```text
405
```

`http.Error()` writes an error response.

General form:

```go
http.Error(w, message, statusCode)
```

---

# 7. `r.URL`

`r.URL` contains information about the request URL.

Example:

```text
/request?id=123
```

Important URL fields we learned:

```go
r.URL.Path
r.URL.RawQuery
r.URL.Query()
```

---

# 8. `r.URL.Path`

Gets only the URL path.

Example request:

```text
GET /users/123
```

Then:

```go
r.URL.Path
```

returns:

```text
/users/123
```

Example:

```go
fmt.Fprintln(w, r.URL.Path)
```

---

# 9. `r.URL.RawQuery`

Gets the raw query string.

Example:

```text
/request?id=123&name=ankush
```

Then:

```go
r.URL.Path
```

gives:

```text
/request
```

and:

```go
r.URL.RawQuery
```

gives:

```text
id=123&name=ankush
```

Important:

`RawQuery` does **not** include the `?`.

---

# 10. `r.URL.Query()`

This is the more useful way to work with query parameters.

```go
query := r.URL.Query()
```

The result is:

```go
url.Values
```

Conceptually:

```go
type Values map[string][]string
```

So query parameters are represented approximately as:

```text
key → multiple values
```

Example:

```text
/request?id=123&name=ankush
```

Conceptually:

```text
id   → ["123"]
name → ["ankush"]
```

---

# 11. `query.Get()`

To retrieve one query parameter:

```go
id := query.Get("id")
```

Example:

```text
/request?id=123
```

Result:

```go
id == "123"
```

If the parameter doesn't exist:

```text
/request
```

then:

```go
query.Get("id")
```

returns:

```text
""
```

Important function to remember:

```go
url.Values.Get()
```

---

# 12. Multiple Values for One Query Parameter

A query parameter can appear multiple times:

```text
/request?id=123&id=456&id=789
```

Using:

```go
query.Get("id")
```

is not the right approach when you want all values.

Instead:

```go
ids := query["id"]
```

Result:

```go
[]string{"123", "456", "789"}
```

Then use `range`:

```go
for _, value := range ids {
    fmt.Fprintln(w, "ID:", value)
}
```

Output:

```text
ID: 123
ID: 456
ID: 789
```

---

# 13. `for range`

We used:

```go
for _, value := range ids {
    fmt.Fprintln(w, "ID:", value)
}
```

Meaning:

```text
ids
 ↓
take each element
 ↓
put it into value
 ↓
execute loop body
```

For:

```go
[]string{"123", "456", "789"}
```

the loop runs three times.

The `_` means:

> I don't need the index.

Without `_`:

```go
for index, value := range ids {
}
```

With `_`:

```go
for _, value := range ids {
}
```

---

# 14. Two-Value Map Lookup

We learned an important Go pattern:

```go
values, ok := query["id"]
```

This is called a **two-value map lookup**.

`values` contains the value:

```go
[]string
```

`ok` tells whether the key exists:

```go
true
false
```

Example:

```go
values, ok := query["id"]

if ok {
    for _, value := range values {
        fmt.Fprintln(w, "ID:", value)
    }
} else {
    fmt.Fprintln(w, "ID is required")
}
```

---

# 15. `query.Get()` vs `query["id"]`

This distinction is important.

### `Get()`

```go
id := query.Get("id")
```

Returns:

```go
string
```

Useful when you want **one value**.

Missing parameter:

```text
""
```

---

### Map lookup

```go
ids, ok := query["id"]
```

Returns:

```text
[]string
bool
```

Useful when:

- you want multiple values
- you need to know whether the key actually exists

Example:

```text
/request?id=123&id=456
```

gives:

```go
ids = []string{"123", "456"}
ok = true
```

---

# 16. Applying the Same Pattern to `name`

Example URL:

```text
/request?name=ankush&name=rahul&name=amit
```

Code:

```go
names, ok := query["name"]

if ok {
    for _, value := range names {
        fmt.Fprintln(w, "Name:", value)
    }
} else {
    fmt.Fprintln(w, "Name is required")
}
```

Output:

```text
Name: ankush
Name: rahul
Name: amit
```

---

# 17. `r.Header`

HTTP requests also contain **headers**.

Access all headers:

```go
headers := r.Header
```

`r.Header` is of type:

```go
http.Header
```

Conceptually:

```go
type Header map[string][]string
```

So headers are also represented as:

```text
Header name → multiple values
```

Example output:

```text
map[
    Accept:[...]
    Accept-Encoding:[...]
    Accept-Language:[...]
    User-Agent:[...]
]
```

---

# 18. Common Browser Request Headers

When you opened:

```text
http://localhost:8080/request
```

your browser sent headers such as:

```text
Accept
Accept-Encoding
Accept-Language
User-Agent
Sec-Ch-Ua
Sec-Ch-Ua-Mobile
Sec-Ch-Ua-Platform
```

These are automatically sent by the browser.

---

# 19. `r.Header.Get()`

To retrieve one specific header:

```go
userAgent := r.Header.Get("User-Agent")
```

Then:

```go
fmt.Fprintln(w, "User-Agent:", userAgent)
```

Example output:

```text
User-Agent: Mozilla/5.0 ...
```

Another example:

```go
accept := r.Header.Get("Accept")
```

Then:

```go
fmt.Fprintln(w, "Accept:", accept)
```

---

# 20. `Header.Get()` Mental Model

Think:

```text
r
│
└── Header
      │
      ├── Accept
      ├── User-Agent
      ├── Accept-Language
      └── ...
```

Calling:

```go
r.Header.Get("User-Agent")
```

means:

```text
Go to request
   ↓
go to headers
   ↓
find "User-Agent"
   ↓
return its value
```

---

# 21. Header Lookup With `ok`

We also started learning the same two-value lookup pattern with headers:

```go
value, ok := r.Header["User-Agent"]
```

Here:

```go
value
```

contains the header values.

And:

```go
ok
```

tells whether the header key exists.

This follows the same Go map pattern we learned with:

```go
query["id"]
```

---

# 22. Important Connection: Query Parameters vs Headers

Both use a map-like structure.

### Query parameters

```go
query := r.URL.Query()

values, ok := query["id"]
```

### Headers

```go
values, ok := r.Header["User-Agent"]
```

Both support multiple values.

Conceptually:

```text
Query parameters                 Headers

r.URL.Query()                    r.Header
      ↓                              ↓
 url.Values                       http.Header
      ↓                              ↓
map[string][]string              map[string][]string
```

This is an important Go pattern.

---

# 23. `http.ResponseWriter`

The first handler argument:

```go
w http.ResponseWriter
```

represents the outgoing HTTP response.

For example:

```go
fmt.Fprintln(w, "Hello")
```

writes:

```text
Hello
```

to the HTTP response.

Important functions used with `w`:

```go
fmt.Fprintln(w, ...)
fmt.Fprint(w, ...)
fmt.Fprintf(w, ...)
http.Error(w, ...)
```

---

# 24. `fmt.Fprintln`

We used this constantly:

```go
fmt.Fprintln(w, "Hello")
```

Because `w` implements:

```go
io.Writer
```

`fmt.Fprintln` can write to the HTTP response.

Compare:

```go
fmt.Println("Hello")
```

This writes to the **terminal**.

Whereas:

```go
fmt.Fprintln(w, "Hello")
```

writes to the **HTTP response**.

---

# 25. Current Practice Structure

Your practice project currently has:

```text
Practice-02/
└── http-net/
    ├── go.mod
    ├── main.go
    └── main2.go
```

Both files use:

```go
package main
```

`main.go` contains:

```go
func main()
```

and starts the server.

`main2.go` contains additional handlers and registration logic.

---

# 26. Current Routes

You currently have routes such as:

```go
http.HandleFunc("/request", handlefunction)
http.HandleFunc("/header", handleheader)
```

So:

```text
/request → handlefunction
/header  → handleheader
```

---

# 27. `http.HandleFunc`

General form:

```go
http.HandleFunc("/path", handlerFunction)
```

Example:

```go
http.HandleFunc("/header", handleheader)
```

This registers a handler for `/header`.

---

# 28. `http.ListenAndServe`

Starts the HTTP server:

```go
http.ListenAndServe(":8080", nil)
```

Meaning:

```text
:8080 → listen on port 8080
nil   → use the default ServeMux
```

---

# 29. Important Symbols to Search During Revision

Keep this list. These are the exact Go symbols you have encountered:

### Package

```go
"net/http"
```

### Request

```go
http.Request
```

### Response

```go
http.ResponseWriter
```

### Server

```go
http.ListenAndServe
```

### Routing

```go
http.HandleFunc
```

### Methods

```go
r.Method

http.MethodGet
http.MethodPost
http.MethodPut
http.MethodPatch
http.MethodDelete
```

### URL

```go
r.URL
r.URL.Path
r.URL.RawQuery
r.URL.Query()
```

### Query parameters

```go
url.Values
query.Get()
query["key"]
```

### Map lookup

```go
value, ok := map[key]
```

### Headers

```go
r.Header
http.Header
r.Header.Get()
r.Header["Header-Name"]
```

### Response writing

```go
fmt.Fprintln
fmt.Fprint
fmt.Fprintf
```

### HTTP errors

```go
http.Error
http.StatusMethodNotAllowed
```

### Loop

```go
for range
```

---

# 30. Current Mental Model

At this point, you should be able to look at:

```text
POST /request?id=123&name=ankush
```

and mentally break it down as:

```text
POST
 ↓
r.Method

/request
 ↓
r.URL.Path

id=123&name=ankush
 ↓
r.URL.RawQuery
 ↓
r.URL.Query()

id
 ↓
query.Get("id")
or
query["id"]

name
 ↓
query.Get("name")
or
query["name"]
```

And for headers:

```text
User-Agent: Mozilla/5.0 ...
Accept: text/html...
```

you can access them with:

```go
r.Header.Get("User-Agent")
r.Header.Get("Accept")
```

---

