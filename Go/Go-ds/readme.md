# Go Data Structures, Structs, Methods & Interfaces

---

# 2. Go Data Structures ⭐⭐⭐⭐⭐

These are the core data structures you will use throughout Go backend and systems programming.

---

# 2.1 Arrays

## Fundamentals

- [ ] What is an array?
- [ ] Array declaration
- [ ] Array initialization
- [ ] Array literals
- [ ] Array length
- [ ] Array indexing
- [ ] Zero-value initialization
- [ ] Iterating over arrays
- [ ] Nested arrays
- [ ] Multidimensional arrays

## Array Semantics

- [ ] Arrays are value types
- [ ] Array assignment
- [ ] Array copying
- [ ] Passing arrays to functions
- [ ] Returning arrays from functions
- [ ] Comparing arrays
- [ ] Arrays vs pointers

## Memory

- [ ] Contiguous memory
- [ ] Array size in memory
- [ ] Stack vs heap for arrays
- [ ] Fixed-size nature of arrays
- [ ] Array memory layout

## Important

- [ ] Understand why arrays have fixed length
- [ ] Understand array value semantics
- [ ] Understand when arrays are preferable to slices

---

# 2.2 Slices ⭐⭐⭐⭐⭐

## Fundamentals

- [ ] What is a slice?
- [ ] Slice declaration
- [ ] Slice literals
- [ ] Creating slices
- [ ] Indexing
- [ ] Iterating
- [ ] Modifying elements
- [ ] Nil slices
- [ ] Empty slices

## Slice Internals

- [ ] Slice header
- [ ] Pointer
- [ ] Length
- [ ] Capacity
- [ ] Underlying array
- [ ] Relationship between slice and array
- [ ] Slice memory layout
- [ ] Slice descriptor

### Conceptual Structure

```text
Slice
├── Pointer → underlying array
├── Length
└── Capacity

## Slicing

- [ ] Basic slicing
- [ ] `a[low:high]`
- [ ] Full slice expression
- [ ] `a[low:high:max]`
- [ ] Slice bounds
- [ ] Shared underlying array
- [ ] Sub-slices

## Append

- [ ] `append`
- [ ] Appending one element
- [ ] Appending multiple elements
- [ ] Appending another slice
- [ ] `...` expansion
- [ ] Append with sufficient capacity
- [ ] Append when capacity is exceeded
- [ ] Slice reallocation
- [ ] Underlying array replacement

## Capacity

- [ ] `len`
- [ ] `cap`
- [ ] Initial capacity
- [ ] Capacity growth
- [ ] Reallocation
- [ ] Preallocating capacity
- [ ] `make([]T, len, cap)`

## Copying

- [ ] `copy`
- [ ] Copying between slices
- [ ] Partial copies
- [ ] Overlapping slices
- [ ] Shallow copy behavior
- [ ] Deep copy considerations

## Memory Behavior

- [ ] Shared backing arrays
- [ ] Slice aliasing
- [ ] Memory retention
- [ ] Large backing array problem
- [ ] Reslicing
- [ ] Copying to release memory

## Common Pitfalls

- [ ] Unexpected modification through shared slices
- [ ] Append invalidating assumptions
- [ ] Nil vs empty slice
- [ ] Slice passed to function
- [ ] Slice returned from function
- [ ] Memory retention

---

# 2.3 Maps ⭐⭐⭐⭐⭐

## Fundamentals

- [ ] What is a map?
- [ ] Map declaration
- [ ] Map literals
- [ ] `make(map[K]V)`
- [ ] Insert
- [ ] Lookup
- [ ] Update
- [ ] Delete
- [ ] Iteration
- [ ] `len`

## Lookup

- [ ] Single-value lookup
- [ ] Two-value lookup
- [ ] `value, ok`
- [ ] Missing keys
- [ ] Zero value of missing keys

## Map Behavior

- [ ] Maps are reference-like data structures
- [ ] Map assignment
- [ ] Passing maps to functions
- [ ] Returning maps
- [ ] Nil maps
- [ ] Empty maps

## Map Internals — Conceptual

- [ ] Hash table concept
- [ ] Hash function
- [ ] Buckets
- [ ] Hash collisions
- [ ] Load factor
- [ ] Map growth
- [ ] Rehashing concept
- [ ] Bucket organization
- [ ] Overflow concept

You do NOT need to memorize Go runtime source code.

Understand the conceptual model.

## Concurrency

- [ ] Why regular maps are not safe for concurrent writes
- [ ] Concurrent reads
- [ ] Concurrent read/write problems
- [ ] Mutex-protected maps
- [ ] `sync.RWMutex`
- [ ] `sync.Map`

## Important

- [ ] Understand average O(1) lookup conceptually
- [ ] Understand key comparability
- [ ] Understand map iteration order is not guaranteed
- [ ] Understand nil map behavior

---

# 2.4 Structs ⭐⭐⭐⭐⭐

## Fundamentals

- [ ] Struct definition
- [ ] Struct fields
- [ ] Struct literals
- [ ] Positional struct literals
- [ ] Named struct literals
- [ ] Zero-value struct
- [ ] Accessing fields
- [ ] Modifying fields

## Nested Structs

- [ ] Struct inside struct
- [ ] Nested field access
- [ ] Anonymous nested structs
- [ ] Struct composition

## Struct Embedding

- [ ] Embedded structs
- [ ] Promoted fields
- [ ] Promoted methods
- [ ] Method conflicts
- [ ] Explicit field access
- [ ] Embedding vs inheritance

## Anonymous Fields

- [ ] Anonymous struct fields
- [ ] Embedded types
- [ ] Named vs anonymous fields
- [ ] Field promotion

## Struct Tags

- [ ] Struct tags
- [ ] JSON tags
- [ ] `json:"name"`
- [ ] `omitempty`
- [ ] Reflection and tags
- [ ] Validation tags

Example:

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

## Struct Memory

- [ ] Struct memory layout
- [ ] Field ordering
- [ ] Padding
- [ ] Alignment
- [ ] Struct size
- [ ] `unsafe.Sizeof` concept

---

# 2.5 Pointers ⭐⭐⭐⭐⭐

## Fundamentals

- [ ] What is a pointer?
- [ ] Pointer declaration
- [ ] Address operator `&`
- [ ] Dereference operator `*`
- [ ] Pointer variables
- [ ] Pointer to primitive
- [ ] Pointer to struct
- [ ] Pointer to pointer

## Pointer Semantics

- [ ] Value vs pointer
- [ ] Passing values
- [ ] Passing pointers
- [ ] Modifying through pointers
- [ ] Pointer equality
- [ ] Pointer copying

## Struct Pointers

- [ ] Pointer to struct
- [ ] Accessing fields through pointers
- [ ] `p.Field` syntax
- [ ] Pointer receiver
- [ ] Returning struct pointers

## Nil Pointers

- [ ] `nil` pointer
- [ ] Nil pointer dereference
- [ ] Checking for nil
- [ ] Pointer lifecycle

## Memory

- [ ] Address space
- [ ] Stack
- [ ] Heap
- [ ] Escape analysis
- [ ] Heap allocation
- [ ] Garbage collector
- [ ] Pointer lifetime

---

# 2.6 nil ⭐⭐⭐⭐⭐

Understand `nil` separately because it behaves differently across Go types.

- [ ] Nil pointer
- [ ] Nil slice
- [ ] Nil map
- [ ] Nil channel
- [ ] Nil interface
- [ ] Nil function
- [ ] Nil error
- [ ] Comparing with nil
- [ ] Dereferencing nil
- [ ] Writing to nil map
- [ ] Reading from nil map
- [ ] Appending to nil slice
- [ ] Sending to nil channel
- [ ] Receiving from nil channel
- [ ] Calling nil function
- [ ] Nil interface trap

---

# 2.7 Zero Values ⭐⭐⭐⭐⭐

Understand the zero value of every major Go type.

- [ ] Zero value of int
- [ ] Zero value of float
- [ ] Zero value of bool
- [ ] Zero value of string
- [ ] Zero value of pointer
- [ ] Zero value of struct
- [ ] Zero value of array
- [ ] Zero value of slice
- [ ] Zero value of map
- [ ] Zero value of channel
- [ ] Zero value of interface
- [ ] Zero value of function

## Design Principle

- [ ] Designing useful zero values
- [ ] Zero-value usability
- [ ] Why Go emphasizes zero values

---

# 2.8 make ⭐⭐⭐⭐⭐

Understand exactly where `make` is used.

- [ ] Purpose of `make`
- [ ] `make` for slices
- [ ] `make` for maps
- [ ] `make` for channels
- [ ] Length argument
- [ ] Capacity argument
- [ ] Slice preallocation
- [ ] Map initialization
- [ ] Channel capacity
- [ ] `make` vs literal initialization
- [ ] `make` vs `new`

Examples:

```go
make([]int, 10)

make([]int, 0, 100)

make(map[string]int)

make(chan int, 10)
```

---

# 2.9 new ⭐⭐⭐⭐

- [ ] Purpose of `new`
- [ ] `new(T)`
- [ ] Zero-value allocation
- [ ] Pointer returned by `new`
- [ ] `new` vs `make`
- [ ] When to use `new`
- [ ] Why `new` is less common in idiomatic Go
- [ ] `new` with structs
- [ ] `new` with primitive types

Example:

```go
p := new(int)
```

---

# 3. Structs + Methods ⭐⭐⭐⭐⭐

---

# 3.1 Methods

## Fundamentals

- [ ] What is a method?
- [ ] Method declaration
- [ ] Receiver
- [ ] Receiver type
- [ ] Method invocation
- [ ] Methods on structs
- [ ] Methods on named types

Example:

```go
type User struct {
    Name string
}

func (u User) GetName() string {
    return u.Name
}
```

---

# 3.2 Value Receivers

- [ ] What is a value receiver?
- [ ] Receiver copying
- [ ] Modifying receiver
- [ ] When to use value receivers
- [ ] Value receiver with small structs
- [ ] Value receiver and interfaces

---

# 3.3 Pointer Receivers ⭐⭐⭐⭐⭐

- [ ] What is a pointer receiver?
- [ ] Modifying original struct
- [ ] Avoiding large struct copies
- [ ] Pointer receiver syntax
- [ ] Automatic address-taking
- [ ] Pointer receiver and interfaces
- [ ] Consistent receiver choice

Example:

```go
func (u *User) SetName(name string) {
    u.Name = name
}
```

---

# 3.4 Value Receiver vs Pointer Receiver

Understand:

```text
Value Receiver
     ↓
Copies receiver

Pointer Receiver
     ↓
Works with original object
```

Study:

- [ ] Mutation
- [ ] Copying
- [ ] Performance
- [ ] Method sets
- [ ] Interface implementation
- [ ] Receiver consistency

---

# 3.5 Method Sets ⭐⭐⭐⭐⭐

- [ ] Method set of value type
- [ ] Method set of pointer type
- [ ] Value receiver methods
- [ ] Pointer receiver methods
- [ ] Interface implementation
- [ ] Why `T` and `*T` can implement different interfaces
- [ ] Method set rules

---

# 3.6 Constructor Patterns

Go does not have constructors like C++/Java.

Learn:

- [ ] Constructor convention
- [ ] `NewType()` pattern
- [ ] Returning pointer
- [ ] Returning value
- [ ] Constructor validation
- [ ] Functional options
- [ ] Default configuration
- [ ] Constructor dependency injection

Example:

```go
func NewServer(port int) *Server {
    return &Server{
        Port: port,
    }
}
```

---

# 3.7 Composition ⭐⭐⭐⭐⭐

- [ ] Composition over inheritance
- [ ] Struct embedding
- [ ] Interface composition
- [ ] Dependency composition
- [ ] Building larger components
- [ ] Avoiding deep inheritance-style designs

---

# 3.8 Dependency Injection

- [ ] Dependency injection concept
- [ ] Constructor injection
- [ ] Interface-based dependencies
- [ ] Dependency inversion
- [ ] Manual dependency injection
- [ ] Testing with injected dependencies
- [ ] Mock implementations

Example:

```text
Handler
   ↓
Service
   ↓
Repository
```

---

# 4. Interfaces ⭐⭐⭐⭐⭐

Interfaces are one of the most important Go concepts for your backend architecture.

---

# 4.1 Interface Fundamentals

- [ ] What is an interface?
- [ ] Interface declaration
- [ ] Interface methods
- [ ] Implementing an interface
- [ ] Implicit implementation
- [ ] No explicit `implements`
- [ ] Interface values
- [ ] Interface as abstraction

Example:

```go
type Storage interface {
    Save(key string, value string) error
    Get(key string) (string, error)
}
```

---

# 4.2 Implicit Interface Implementation

- [ ] Structural typing
- [ ] Method matching
- [ ] Implementing interfaces automatically
- [ ] Compile-time interface satisfaction
- [ ] Small interfaces
- [ ] Interface composition

Example:

```text
Storage interface
       ↑
       |
MemoryStorage
RedisStorage
PostgresStorage
```

---

# 4.3 Interface Values ⭐⭐⭐⭐⭐

Understand the conceptual structure:

```text
Interface Value
 ├── Dynamic Type
 └── Dynamic Value
```

Study:

- [ ] Static type
- [ ] Dynamic type
- [ ] Dynamic value
- [ ] Interface assignment
- [ ] Interface comparison
- [ ] Interface equality
- [ ] Nil interface

---

# 4.4 nil Interface Problem ⭐⭐⭐⭐⭐

Understand this deeply:

```go
var p *MyType = nil

var x MyInterface = p

fmt.Println(x == nil)
```

Study:

- [ ] Nil interface
- [ ] Typed nil
- [ ] Interface containing nil pointer
- [ ] Why interface != nil
- [ ] Detecting typed nil
- [ ] Avoiding typed nil bugs

---

# 4.5 Empty Interface / any

- [ ] `interface{}`
- [ ] `any`
- [ ] Empty interface semantics
- [ ] Storing arbitrary types
- [ ] Type assertions
- [ ] Type switches
- [ ] JSON usage
- [ ] When `any` is appropriate
- [ ] When `any` should be avoided

---

# 4.6 Type Assertions

- [ ] Type assertion syntax
- [ ] Safe type assertion
- [ ] Two-value assertion
- [ ] Failed assertion
- [ ] Pointer type assertion
- [ ] Interface type assertion
- [ ] Type assertion vs type conversion

Example:

```go
value, ok := x.(string)
```

---

# 4.7 Type Switches

- [ ] Type switch syntax
- [ ] Multiple types
- [ ] Default case
- [ ] Type switch with interfaces
- [ ] Type switch with `any`
- [ ] Type switch vs type assertion

Example:

```go
switch v := x.(type) {
case string:
    // ...
case int:
    // ...
default:
    // ...
}
```

---

# 4.8 Interface Composition

- [ ] Combining interfaces
- [ ] Embedded interfaces
- [ ] Small interfaces
- [ ] Interface segregation
- [ ] Building capability-based interfaces

Example:

```go
type Reader interface {
    Read()
}

type Writer interface {
    Write()
}

type ReadWriter interface {
    Reader
    Writer
}
```

---

# 4.9 Interface Design

- [ ] Keep interfaces small
- [ ] Define interfaces near consumers
- [ ] Avoid unnecessary interfaces
- [ ] One-method interfaces
- [ ] Dependency inversion
- [ ] Interface ownership
- [ ] Concrete types vs interfaces
- [ ] Avoid interface pollution
- [ ] Interface composition

---

# 4.10 Interfaces + Dependency Injection

Study:

```text
Handler
   |
   v
Service Interface
   |
   v
Service Implementation
   |
   v
Repository Interface
   |
   +---- Memory Repository
   +---- PostgreSQL Repository
   +---- Redis Repository
```

Topics:

- [ ] Dependency injection
- [ ] Constructor injection
- [ ] Interface-based dependencies
- [ ] Mock dependencies
- [ ] Fake implementations
- [ ] Test isolation

---

# 4.11 Interfaces + Testing

- [ ] Mocking
- [ ] Fake implementations
- [ ] Stub implementations
- [ ] Test doubles
- [ ] Dependency substitution
- [ ] Interface-based testing
- [ ] Testing services independently
- [ ] Testing handlers independently

---

# 4.12 Interfaces + Standard Library

Understand common Go interfaces:

- [ ] `io.Reader`
- [ ] `io.Writer`
- [ ] `io.Closer`
- [ ] `io.ReadWriter`
- [ ] `fmt.Stringer`
- [ ] `error`
- [ ] `http.Handler`
- [ ] `sort.Interface`
- [ ] `context.Context`

These are extremely important because Go's ecosystem relies heavily on interfaces.

---

# 5. Data Structures + Memory Model

After learning the individual structures, understand how they interact.

---

# 5.1 Value Semantics

- [ ] Values
- [ ] Copies
- [ ] Assignment
- [ ] Function arguments
- [ ] Return values
- [ ] Struct copying
- [ ] Array copying
- [ ] Slice copying
- [ ] Map assignment
- [ ] Interface copying

---

# 5.2 Reference-like Behavior

Understand the difference between:

- [ ] Pointer
- [ ] Slice
- [ ] Map
- [ ] Channel
- [ ] Interface
- [ ] Function

Understand what is actually copied when these are passed around.

---

# 5.3 Memory Relationships

Understand:

```text
Array
  |
  +---- contiguous memory

Slice
  |
  +---- pointer
  +---- length
  +---- capacity
          |
          v
      Array

Pointer
  |
  +---- address
          |
          v
       Memory

Map
  |
  +---- runtime-managed hash table

Channel
  |
  +---- runtime-managed synchronization structure
```

---

# 6. Practical Data Structure Patterns

Practice implementing:

- [ ] Dynamic array using slices
- [ ] Stack using slices
- [ ] Queue using slices
- [ ] Circular queue
- [ ] Set using map
- [ ] Frequency map
- [ ] LRU cache
- [ ] Concurrent map
- [ ] Thread-safe counter
- [ ] In-memory key-value store

---

# 7. Apply to Your Projects

## URL Shortener

Implement:

```text
Handler
   ↓
Service
   ↓
Storage Interface
   ↓
Memory Storage
```

Learn:

- [ ] Structs
- [ ] Methods
- [ ] Interfaces
- [ ] Pointers
- [ ] Maps
- [ ] Slices
- [ ] Error handling
- [ ] Dependency injection

---

# Job Queue

Implement:

```text
Producer
    ↓
Queue
    ↓
Workers
    ↓
Handler
```

Learn:

- [ ] Structs
- [ ] Methods
- [ ] Interfaces
- [ ] Channels
- [ ] Slices
- [ ] Maps
- [ ] Pointers
- [ ] Context
- [ ] Concurrency-safe structures

---

# Event Streaming Engine

Implement:

```text
Producer
    ↓
Broker
    ↓
Topic
    ↓
Partition
    ↓
Consumer
```

Learn:

- [ ] Struct composition
- [ ] Interfaces
- [ ] Maps
- [ ] Slices
- [ ] Pointers
- [ ] Memory management
- [ ] Concurrent data structures
- [ ] Channel-based communication
- [ ] Thread-safe state
- [ ] Buffer management

---

# 8. Mastery Checklist

Before moving to advanced Go, I should be able to explain:

- [ ] Array vs Slice
- [ ] Slice length vs capacity
- [ ] Slice header
- [ ] Slice backing array
- [ ] What happens during `append`
- [ ] Slice aliasing
- [ ] Slice memory retention
- [ ] Map lookup
- [ ] Map zero value
- [ ] Map internals conceptually
- [ ] Map concurrency limitations
- [ ] Struct value semantics
- [ ] Struct pointers
- [ ] Pointer vs value
- [ ] Nil values
- [ ] Zero values
- [ ] `make`
- [ ] `new`
- [ ] Value receiver
- [ ] Pointer receiver
- [ ] Method sets
- [ ] Struct embedding
- [ ] Composition
- [ ] Interface implementation
- [ ] Interface values
- [ ] Dynamic type
- [ ] Dynamic value
- [ ] Nil interface
- [ ] Typed nil
- [ ] Type assertion
- [ ] Type switch
- [ ] Interface composition
- [ ] Dependency injection
- [ ] Mocking with interfaces
- [ ] `io.Reader`
- [ ] `io.Writer`
- [ ] `http.Handler`

---

# 9. Recommended Learning Sequence

Follow this order:

1. [ ] Arrays
2. [ ] Slices
3. [ ] Slice internals
4. [ ] Append & Capacity
5. [ ] Slice memory behavior
6. [ ] Maps
7. [ ] Map internals
8. [ ] Structs
9. [ ] Struct embedding
10. [ ] Pointers
11. [ ] Struct pointers
12. [ ] nil
13. [ ] Zero values
14. [ ] make
15. [ ] new
16. [ ] Methods
17. [ ] Value receivers
18. [ ] Pointer receivers
19. [ ] Method sets
20. [ ] Composition
21. [ ] Constructor patterns
22. [ ] Dependency injection
23. [ ] Interfaces
24. [ ] Implicit implementation
25. [ ] Interface values
26. [ ] Nil interfaces
27. [ ] Type assertions
28. [ ] Type switches
29. [ ] Interface composition
30. [ ] Interface design
31. [ ] Interfaces + testing
32. [ ] Standard library interfaces
33. [ ] Memory/value semantics
34. [ ] Concurrent data structures
35. [ ] Apply everything to projects
```

.