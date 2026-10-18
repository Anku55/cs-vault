
# Go —  Quick Notes

## 1. Basic Go Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Remember

- `package main` → executable program package.
- `func main()` → program's entry point.
- `import` → brings another package into your code.
- `fmt.Println()` → prints to the terminal.
- Go source files use `.go`.

---

## 2. Package

A **package** groups related Go code.

```go
package queue
```

All `.go` files in the same directory normally belong to the same package.

Example:

```text
queue/
├── queue.go
├── worker.go
└── retry.go
```

---

## 3. Module

A **module** is a collection of Go packages managed together.

Create one:

```bash
go mod init example/hello
```

This creates:

```text
go.mod
```

Typical project:

```text
event-engine/
├── go.mod
├── broker/
├── queue/
├── storage/
└── consumer/
```

### Package vs Module

```text
Module
 ├── Package
 ├── Package
 └── Package
```

---

## 4. `go.mod`

`go.mod` describes your module and its dependencies.

Example:

```go
module github.com/username/event-engine

go 1.24
```

### Remember

- Identifies your module.
- Records required dependencies.
- Specifies the Go version.
- Usually committed to Git.

---

## 5. External Packages

Example:

```go
import "rsc.io/quote"
```

Then:

```go
fmt.Println(quote.Go())
```

Go can download external modules automatically.

Find Go packages:

**pkg.go.dev**

---

## 6. `go.sum`

`go.sum` contains cryptographic checksums used to verify downloaded modules.

### Remember

```text
go.mod → What dependencies/versions are required
go.sum → Checksums for dependency verification
```

Usually commit both files:

```text
go.mod
go.sum
```

---

## 7. Important Go Commands

### Run

```bash
go run .
```

Compile and run the current package.

### Build

```bash
go build
```

Compile the program.

### Initialize module

```bash
go mod init <module-path>
```

### Clean/update dependencies

```bash
go mod tidy
```

Adds missing dependencies and removes unused ones.

### Format

```bash
go fmt ./...
```

### Test

```bash
go test ./...
```

### Static analysis

```bash
go vet ./...
```

### Get help

```bash
go help
```

---

## 8. Standard Library vs External Module

### Standard library

Comes with Go:

```go
import "fmt"
import "os"
import "net"
import "sync"
```

No separate installation required.

### External module

Provided by another project:

```go
import "rsc.io/quote"
```

Managed through Go Modules.

---

## 9. Important Mental Model

```text
.go file
   ↓
Package
   ↓
Module
   ↓
Dependencies
   ↓
Go toolchain
   ↓
Executable
```

---

# For Your Event Streaming Engine

These concepts will become:

```text
event-engine/
│
├── go.mod
│
├── cmd/
│   └── broker/
│       └── main.go       # Program entry point
│
├── broker/               # Broker package
├── queue/                # Job queue package
├── storage/              # Persistent log
├── consumer/             # Consumer logic
└── protocol/             # Network protocol
```

Example:

```go
package main

import "github.com/username/event-engine/broker"

func main() {
    b := broker.New()
    b.Start()
}
```

---

#  Must Remember

- **Package** → groups Go source files.
- **`main` package** → creates an executable.
- **`main()`** → program entry point.
- **`import`** → use another package.
- **Module** → collection of packages.
- **`go.mod`** → module + dependency information.
- **`go.sum`** → dependency checksums.
- **`go mod tidy`** → synchronizes dependencies.
- **`go run .`** → run current package.
- **`go build`** → compile.
- **`go test`** → run tests.
- **`go fmt`** → format code.
- **`go vet`** → detect suspicious code.
- **pkg.go.dev** → discover/read Go package documentation.

---

---
How do you create your own Go module containing reusable code that another Go application can import?
---
```text
Module 1
greetings
   ↓
contains reusable code

Module 2
hello
   ↓
imports greetings
```
Conceptually
```
┌─────────────────────┐
│ greetings module    │
│                     │
│ package greetings   │
│                     │
│ Hello()             │
└──────────┬──────────┘
           │
           │ import
           ▼
┌─────────────────────┐
│ application module  │
│                     │
│ package main        │
│                     │
│ calls Hello()       │
└─────────────────────┘
```
---
One Go module can provide reusable packages that another module can import.
---
# Go: Package, Module, and Exported Names

## 1. Package vs Module

### Package

A **package** groups related Go source files.

Example:

```text
greetings/
└── greetings.go
```

Inside `greetings.go`:

```go
package greetings
```

So, `greetings` is a package.

### Module

A **module** is a collection of one or more related packages.

Example:

```text
greetings/
├── go.mod
├── greetings.go
└── ...
```

The `go.mod` file makes the directory a Go module.

### Remember

```text
Packages → organize code
Modules  → organize packages + dependencies
```

---

## 2. Creating a Module

Create a directory:

```bash
mkdir greetings
cd greetings
```

Initialize it as a Go module:

```bash
go mod init example.com/greetings
```

This creates:

```text
greetings/
└── go.mod
```

---

## 3. `go mod init`

```bash
go mod init example.com/greetings
```

This initializes a new Go module.

The argument:

```text
example.com/greetings
```

is the **module path**.

The resulting `go.mod` looks approximately like:

```go
module example.com/greetings

go 1.xx
```

The exact Go version depends on the installed Go version.

---

## 4. Module Path

The **module path** is the identity/name used to refer to a module.

For example:

```go
module github.com/ankush/event-engine
```

A package inside the module:

```text
github.com/ankush/event-engine
└── queue
```

can be imported as:

```go
import "github.com/ankush/event-engine/queue"
```

### Important relationship

```text
Module path
     +
Package directory
     ↓
Import path
```

For example:

```text
github.com/ankush/event-engine
                │
                └── queue
```

becomes:

```go
import "github.com/ankush/event-engine/queue"
```

---

## 5. Why Module Path Matters

The module path tells Go the identity/location associated with the module.

For a GitHub project:

```go
module github.com/ankush/event-engine
```

This corresponds to the repository location.

For tutorials, you may see:

```go
module example.com/greetings
```

You do **not** need to own `example.com`.

---

## 6. `go.mod`

`go.mod` describes the Go module.

Basic example:

```go
module example.com/greetings

go 1.xx
```

It can also contain dependencies:

```go
module example.com/greetings

go 1.xx

require some-module v1.2.3
```

### `go.mod` tells Go:

- What is this module?
- Which Go version does it use?
- Which external modules does it depend on?

---

## 7. Creating a Package

Example:

```text
greetings/
├── go.mod
└── greetings.go
```

`greetings.go`:

```go
package greetings

import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

---

## 8. `package greetings`

```go
package greetings
```

This tells Go that the source file belongs to the `greetings` package.

It is different from:

```go
package main
```

`package main` is normally used for an executable program.

`package greetings` is being used here as a **reusable package/library**.

Conceptually:

```text
greetings
    │
    └── Hello()
```

Another application can import and use this package.

---

## 9. Importing `fmt`

```go
import "fmt"
```

`fmt` is part of Go's **standard library**.

The code needs it because it uses:

```go
fmt.Sprintf()
```

`fmt.Sprintf()` creates a formatted string.

---

## 10. Understanding the `Hello` Function

```go
func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

Function signature:

```text
func Hello(name string) string
     │     │            │
     │     │            └── return type
     │     └─────────────── parameter
     └───────────────────── function name
```

Meaning:

> Define a function named `Hello` that accepts a string called `name` and returns a string.

---

## 11. Function Parameters

```go
name string
```

means:

```text
name   → variable/parameter name
string → parameter type
```

Calling:

```go
Hello("Ankush")
```

means that inside the function:

```text
name = "Ankush"
```

---

## 12. Return Type

The final:

```go
string
```

means the function returns a string.

Example:

```go
func Hello(name string) string {
    return "Hello"
}
```

Conceptually:

```text
Hello("Ankush")
       ↓
"Hello"
```

The caller receives the returned string.

---

## 13. Exported vs Unexported Names ⭐

Go uses **capitalization** to control whether identifiers are accessible from another package.

### Exported

Starts with a capital letter:

```go
func Hello() {
}
```

Other packages can use it.

### Unexported

Starts with a lowercase letter:

```go
func hello() {
}
```

It cannot be accessed directly from another package.

### Example

```go
package greetings

func Hello() {
}

func goodbye() {
}
```

Another package can use:

```go
greetings.Hello()
```

But cannot use:

```go
greetings.goodbye()
```

because `goodbye` is unexported.

---

## 14. Simple Mental Model

```text
greetings package
│
├── Hello()       ← exported / accessible
│
└── goodbye()     ← unexported / package-private
```

### Rule to remember

```text
Capital first letter → Exported
Lowercase first letter → Unexported
```

This applies to many Go identifiers, including:

- Functions
- Types
- Variables
- Constants
- Struct fields

---

## 15. Why This Matters for an Event Engine

When designing packages, expose only what other packages need.

Example:

```go
type Queue struct {
}

func (q *Queue) Push() {
}

func (q *Queue) pushInternal() {
}
```

Here:

```text
Push()        → exported API
pushInternal() → internal implementation
```

This allows you to keep the package's internal implementation hidden while exposing a clean public API.

---

## Key Takeaways

```text
Package
→ Groups related Go source files.

Module
→ Groups one or more packages and manages dependencies.

go mod init
→ Creates/initializes a Go module.

go.mod
→ Describes the module, Go version, and dependencies.

Module path
→ The module's identity and basis for import paths.

Function parameter
→ name string

Function return type
→ The type after the parameter list.

Capitalized identifier
→ Exported.

Lowercase identifier
→ Unexported.
```
# Go Module Tutorial — Notes

## 1. `:=` — Short Variable Declaration

`:=` declares **and initializes** a variable.

```go
message := fmt.Sprintf(...)
```

Go automatically determines the variable's type from the value.

Examples:

```go
message := "Hello"  // string
count := 10         // int
price := 10.5       // floating-point type
```

Approximately equivalent to:

```go
var message string
message = "Hello"
```

---

## 2. `:=` vs `var`

### Short form

```go
message := "Hello"
```

### Explicit form

```go
var message string
message = "Hello"
```

Or:

```go
var message string = "Hello"
```

In normal Go code, `:=` is commonly used when the type is obvious from the assigned value.

---

## 3. `fmt.Sprintf`

`fmt.Sprintf` **creates and returns a formatted string**.

It does **not** print the result.

```go
message := fmt.Sprintf("Hi, %v. Welcome!", name)
```

If:

```go
name := "Ankush"
```

then:

```text
message = "Hi, Ankush. Welcome!"
```

Compare:

```go
fmt.Println("Hello")
```

→ prints something.

```go
fmt.Sprintf("Hello, %v", name)
```

→ creates a string and returns it.

---

## 4. `%v`

`%v` is a formatting verb that means approximately:

> Format the value using its default representation.

Example:

```go
name := "Ankush"

message := fmt.Sprintf("Hi, %v. Welcome!", name)
```

Result:

```text
Hi, Ankush. Welcome!
```

---

## 5. Function Execution Flow

When another package calls:

```go
greetings.Hello("Ankush")
```

the flow is:

```text
Hello("Ankush")
      ↓
name = "Ankush"
      ↓
fmt.Sprintf(...)
      ↓
"Hi, Ankush. Welcome!"
      ↓
return message
      ↓
caller receives the string
```

---

## 6. Why Packages Are Useful

The important idea isn't the greeting itself.

The important idea is **reusable code**.

Instead of putting everything inside:

```go
func main() {
    // 500 lines of code
}
```

you can divide functionality into packages:

```text
broker/
queue/
storage/
consumer/
protocol/
```

Each package provides functionality that other packages can use.

---

# 7. Event Engine Package Structure

A future event engine might look like:

```text
event-engine/
│
├── go.mod
│
├── broker/
│   └── broker.go
│
├── queue/
│   └── queue.go
│
├── storage/
│   └── log.go
│
├── consumer/
│   └── consumer.go
│
└── cmd/
    └── broker/
        └── main.go
```

### Queue package

```go
package queue

type Job struct {
    ID      string
    Payload []byte
}

func New() *Queue {
    // ...
}
```

### Broker package

```go
package broker

import "github.com/ankush/event-engine/queue"

func New() *Broker {
    // ...
}
```

### Main application

```go
package main

import "github.com/ankush/event-engine/broker"

func main() {
    b := broker.New()
    b.Start()
}
```

The idea is:

```text
main application
       ↓
     broker
       ↓
     queue
```

Each package has a focused responsibility.

---

# 8. Exported vs Unexported Names

Go uses **capitalization** instead of keywords like `public` and `private`.

### Exported

Starts with a capital letter:

```go
func Hello() {
}
```

Accessible from another package.

### Unexported

Starts with a lowercase letter:

```go
func hello() {
}
```

Not directly accessible from another package.

---

## 9. Exported and Unexported Types

The same rule applies to types.

```go
type Queue struct {}
```

`Queue` is exported.

```go
type queue struct {}
```

`queue` is unexported.

---

## 10. Exported and Unexported Struct Fields

Example:

```go
type Job struct {
    ID      string
    payload []byte
}
```

Here:

```text
ID       → exported
payload  → unexported
```

Another package can directly access:

```go
job.ID
```

but cannot directly access:

```go
job.payload
```

---

# 11. Package Architecture

The basic architecture being introduced is:

```text
                    Go Module
                       │
            ┌──────────┴──────────┐
            │                     │
        Package A             Package B
            │                     │
       exported API          exported API
            │                     │
            └──────────┬──────────┘
                       │
                  Application
```

For an event engine:

```text
                 event-engine module
                         │
       ┌─────────────────┼─────────────────┐
       │                 │                 │
     queue             broker            storage
       │                 │                 │
       └─────────────────┼─────────────────┘
                         │
                    broker app
```

The key idea is:

> A module contains packages, and packages expose APIs that other packages can use.

---

# 12. Function Signature

Example:

```go
func Hello(name string) string
```

This means:

```text
Function: Hello
Input:    string
Output:   string
```

Breakdown:

```text
func Hello(name string) string
     │     │            │
     │     │            └── return type
     │     └─────────────── parameter
     └───────────────────── function name
```

---

# 13. `return`

`return` sends a value back to the caller.

```go
func Hello(name string) string {
    message := "Hello " + name
    return message
}
```

Conceptually:

```text
caller
  ↓
Hello("Ankush")
  ↓
"Hello Ankush"
```

---

# 14. Tutorial Concepts

The tutorial introduces these concepts:

### Create a module

```text
Modules
Packages
Functions
Exported functions
Parameters
Return values
:=
fmt.Sprintf
```

### Call code from another module

Example:

```go
import "example.com/greetings"
```

Then:

```go
greetings.Hello(...)
```

This teaches package/module usage across module boundaries.

### Return and handle errors

Go uses the `error` type.

Example:

```go
func Publish(msg Message) error
```

### Return a random greeting

Introduces **slices**:

```go
[]string
```

Slices are fundamental in Go.

Examples:

```go
[]byte
[]Message
[]Job
```

### Greetings for multiple people

Introduces **maps**:

```go
map[string]string
```

### Add a test

Introduces Go's built-in testing system:

```go
func TestHello(t *testing.T)
```

### Compile and install

Introduces Go's build tooling:

```bash
go build
go install
```

---

# 15. Key Things to Remember

## Module

A collection of related Go packages.

## Package

A group of related Go source files.

## `go.mod`

Defines the module and tracks its dependencies.

## Module path

The identity used to reference the module.

Example:

```text
github.com/ankush/event-engine
```

## Exported function

Starts with uppercase:

```go
func Hello()
```

Accessible from another package.

## Unexported function

Starts with lowercase:

```go
func hello()
```

Not accessible directly from another package.

## `:=`

Short variable declaration:

```go
message := "Hello"
```

## `fmt.Sprintf`

Creates a formatted string:

```go
message := fmt.Sprintf("Hi, %v", name)
```

It does **not** print the result.

## `return`

Sends a value back to the caller:

```go
return message
```

---

# Mental Model

```text
Module
  │
  ├── Package
  │     ├── exported API
  │     └── internal implementation
  │
  ├── Package
  │     ├── exported API
  │     └── internal implementation
  │
  └── Package
```

The core idea:

> **Modules organize packages. Packages organize code. Exported identifiers define what other packages can use.**