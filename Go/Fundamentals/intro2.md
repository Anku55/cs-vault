# Go Fundamentals — Variables, Types, Control Flow & Functions

> Practical syntax notes for building Go fluency.

---

# 1. Variables

In Go, variables are named containers for values.

Go is **statically typed**, meaning every variable has a type known at compile time. A variable's type does not automatically change later.

For example:

```go
var age int = 30
```

Here:

```text
age → variable
int → type
30  → value
```

---

## 1.1 `var` — Variable Declaration

Go provides the `var` keyword for declaring variables.

`var` can be used:

- Inside functions
- At package level

You can explicitly specify the type:

```go
var age int = 30
```

Or let Go infer the type:

```go
var city = "Etawah"
```

Go determines:

```text
city → string
```

### Multiple variables with `var`

You can group declarations:

```go
var (
    name    string = "Alice"
    isAdmin bool   = true
)
```

---

# 2. `:=` — Short Variable Declaration

Inside functions, Go commonly uses the short variable declaration:

```go
count := 10
age := 18
class := "frontier"
```

Go determines the type from the value.

For example:

```go
count := 10
```

means:

```text
count → int
```

And:

```go
class := "frontier"
```

means:

```text
class → string
```

### Important

`:=` can be used for **declaring new variables inside functions**.

It cannot be used for package-level variable declarations.

---

# 3. Zero Values

If you declare a variable without explicitly giving it an initial value, Go automatically gives it a **zero value**.

Go does not leave variables with arbitrary/garbage values.

Common zero values:

| Type | Zero value |
|---|---|
| `int`, `float64`, etc. | `0` |
| `string` | `""` |
| `bool` | `false` |
| pointer | `nil` |
| slice | `nil` |
| map | `nil` |
| channel | `nil` |
| interface | `nil` |
| function | `nil` |

Example:

```go
var age int
var name string
var active bool
```

Conceptually:

```text
age    → 0
name   → ""
active → false
```

Zero values are an important part of Go's design.

---

# 4. Built-in Data Types

Go's built-in types can be understood as several broad categories.

> The categories below are useful for learning, although "basic", "aggregate", "reference", and "interface" are informal organizational categories rather than Go's official type-system taxonomy.

---

## 4.1 Basic Types

### Boolean

```go
bool
```

Can contain:

```text
true
false
```

Example:

```go
isAdmin := true
```

---

### String

```go
string
```

Represents a sequence of bytes, conventionally containing UTF-8 encoded text.

Example:

```go
name := "Alice"
```

---

### Signed Integers

```go
int
int8
int16
int32
int64
```

Signed integers can represent positive and negative whole numbers.

Example:

```go
age := 20
```

`int` is implementation-dependent in size and is typically 32 or 64 bits. Its exact size depends on the Go implementation/architecture.

---

### Unsigned Integers

```go
uint
uint8
uint16
uint32
uint64
```

Unsigned integers represent non-negative integer values.

Example:

```go
var count uint = 10
```

---

### Floating-Point Numbers

```go
float32
float64
```

Used for floating-point/decimal values.

Example:

```go
price := 10.5
```

An untyped floating-point constant such as `10.5` is normally given type `float64` when a concrete type is required.

---

### `byte`

```go
byte
```

`byte` is an alias for:

```go
uint8
```

It is commonly used when working with raw binary data.

For example:

```go
data := []byte("hello")
```

This will be particularly important when working with:

- files
- network protocols
- serialization
- event payloads

---

### `rune`

```go
rune
```

`rune` is an alias for:

```go
int32
```

It represents a Unicode code point.

Example:

```go
var r rune = 'A'
```

---

### Complex Numbers

```go
complex64
complex128
```

Represent numbers with real and imaginary components.

Example:

```go
z := complex(2, 3)
```

---

# 5. Composite / Other Important Types

### Arrays

Fixed-size sequences of elements.

```go
var numbers [5]int
```

---

### Structs

Custom types containing multiple fields.

```go
type User struct {
    Name string
    Age  int
}
```

---

### Pointers

Store an address/reference to a value.

```go
var p *int
```

---

### Slices

Flexible, dynamically sized views over arrays.

```go
numbers := []int{1, 2, 3}
```

---

### Maps

Key/value collections.

```go
ages := map[string]int{
    "Alice": 20,
}
```

---

### Channels

Used for communication between goroutines.

```go
ch := make(chan int)
```

---

### Functions

Functions are also values in Go.

They can be:

- stored in variables
- passed to other functions
- returned from functions

Example:

```go
add := func(a, b int) int {
    return a + b
}
```

---

### Interfaces

An interface defines a set of methods that a type can satisfy.

Example:

```go
type Writer interface {
    Write([]byte) error
}
```

A type satisfies an interface by implementing the required methods.

---

# 6. Strict Go Rules

Go intentionally has several strict compiler rules.

---

## 6.1 Unused Variables

A local variable declared inside a function generally must be used.

This causes a compile error:

```go
func main() {
    age := 20
}
```

because `age` is declared but unused.

Go's compiler treats unused local variables as errors.

---

## 6.2 No Implicit Numeric Conversion

Go does not automatically convert numeric types.

For example, you cannot freely combine:

```go
int
```

and:

```go
float64
```

without conversion.

Example:

```go
var x int = 10
var y float64 = 20.5

// x + y  // invalid
```

You need an explicit conversion:

```go
result := float64(x) + y
```

This strictness helps prevent accidental type conversions.

---

# 7. `if` Statements

Go's `if` syntax is straightforward.

Example:

```go
score := 85

if score >= 60 {
    fmt.Println("You passed!")
}
```

### Important syntax rules

Go does **not** use parentheses around the condition:

```go
if score >= 60 {
}
```

Not:

```go
if (score >= 60) {
}
```

Curly braces are required:

```go
if score >= 60 {
    // code
}
```

The condition must evaluate to a `bool`.

Go does **not** have implicit truthy/falsy behavior like some other languages.

For example, you cannot use an integer directly as a condition:

```go
// if 1 { }       // invalid
// if 0 { }       // invalid
```

---

# 8. `if` with `else if` and `else`

Example:

```go
score := 85

if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("Fail")
}
```

Go requires the opening `{` to be on the same line as the `if`, `else if`, or `else` statement.

---

# 9. `if` Initialization Statement

Go allows a short initialization statement before the condition.

Syntax:

```go
if initialization; condition {
}
```

Example:

```go
if val := getCount(); val > 10 {
    fmt.Println("High count:", val)
} else {
    fmt.Println("Low count:", val)
}
```

Here:

```go
val := getCount()
```

initializes `val`.

Then:

```go
val > 10
```

checks the condition.

### Scope

`val` is available inside the `if` statement and its corresponding `else if` / `else` blocks.

It is not available after the entire `if` statement:

```go
if val := getCount(); val > 10 {
    fmt.Println(val)
}

 // val is not accessible here
```

This is useful because temporary variables don't leak into the surrounding scope.

---

# 10. No Ternary Operator

Go does not have the traditional ternary operator found in languages such as C/C++/Java/JavaScript:

```text
condition ? value1 : value2
```

Instead, use normal `if` statements.

---

# 11. `for` Loops

Go has one looping keyword:

```go
for
```

It can represent several different loop styles.

---

# 12. Classic `for` Loop

The traditional form contains:

```text
initialization
condition
post-statement
```

Example:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Execution:

```text
i = 0
 ↓
check i < 5
 ↓
execute body
 ↓
i++
 ↓
check again
```

Output:

```text
0
1
2
3
4
```

Unlike languages such as C/C++/Java, Go does not put parentheses around the loop header.

---

# 13. `for` as a While Loop

Go does not have a separate `while` keyword.

Instead:

```go
n := 1

for n <= 5 {
    fmt.Println(n)
    n++
}
```

This behaves like a traditional `while` loop.

The loop continues while:

```go
n <= 5
```

is true.

---

# 14. Infinite `for` Loop

Omit the condition:

```go
for {
    fmt.Println("Running indefinitely...")
}
```

This creates an infinite loop.

You can stop it using:

```go
break
```

Example:

```go
for {
    fmt.Println("Running indefinitely...")
    break
}
```

`return` can also terminate the surrounding function.

Infinite loops are common in long-running systems such as servers and workers.

---

# 15. `for range`

`range` is used to iterate over collections and other iterable values such as:

- arrays
- slices
- strings
- maps
- channels

Example:

```go
names := []string{"Alice", "Bob", "Charlie"}

for index, value := range names {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}
```

The loop provides:

```text
index → position
value → element
```

---

## 15.1 Ignore a Value with `_`

If you don't need the index:

```go
for _, value := range names {
    fmt.Println(value)
}
```

`_` is called the **blank identifier**.

It means:

> I intentionally don't need this value.

---

## 15.2 Go 1.22+ Integer `range`

Modern Go also allows:

```go
for i := range 5 {
    fmt.Println(i)
}
```

This iterates over:

```text
0
1
2
3
4
```

---

# 16. Functions

A function is a reusable block of code declared using:

```go
func
```

A basic function can have:

- a name
- parameters
- return values
- a function body

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

Calling it:

```go
result := add(5, 3)
fmt.Println("Sum:", result)
```

Output:

```text
Sum: 8
```

---

# 17. Function Parameters

This:

```go
func add(a int, b int) int
```

means:

```text
Function name → add
Parameter     → a
Type          → int
Parameter     → b
Type          → int
Return type   → int
```

When consecutive parameters have the same type, Go allows a shorter syntax:

```go
func add(a, b int) int {
    return a + b
}
```

This is equivalent to:

```go
func add(a int, b int) int {
    return a + b
}
```

---

# 18. Multiple Return Values

One of Go's important features is that functions can return multiple values.

Example:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }

    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)

if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

This pattern is extremely common in Go.

A function can return:

```text
result + error
```

instead of throwing an exception.

---

# 19. Named Return Values

Go allows return values to be named in the function signature.

Example:

```go
func rectangleArea(width, height int) (area int) {
    area = width * height
    return
}
```

Here:

```go
(area int)
```

declares a named return variable.

The bare:

```go
return
```

returns the current value of `area`.

### Important

Named returns are a Go feature, but they should be used when they improve clarity. They are not required for normal functions.

---

# 20. Variadic Functions

A variadic function accepts a variable number of arguments.

Use:

```go
...
```

before the parameter type:

```go
func sumAll(nums ...int) int {
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
```

Call it:

```go
result := sumAll(1, 2, 3, 4)
```

Inside the function:

```go
nums
```

behaves like a slice:

```go
[]int
```

---

# 21. Passing Arguments to Functions

Go uses **pass-by-value** semantics.

That means a function receives a copy of the value passed to it.

Example:

```go
func update(x int) {
    x = 100
}

func main() {
    value := 10

    update(value)

    fmt.Println(value)
}
```

The original `value` remains:

```text
10
```

because `update` received a copy.

---

# 22. Passing a Pointer

If you want a function to modify the caller's variable, you can pass a pointer.

Example:

```go
func update(x *int) {
    *x = 100
}

func main() {
    value := 10

    update(&value)

    fmt.Println(value)
}
```

Now:

```text
value → 100
```

because the function received a pointer to the original variable.

### Important distinction

The function still receives a **copy of the pointer value**. That copied pointer points to the same underlying variable.

So it is more precise to say:

> Go always passes arguments by value; when the value is a pointer, the copied pointer still refers to the same underlying object.

---

# 23. Functions Are First-Class Values

In Go, functions are first-class values.

They can be:

- assigned to variables
- passed as arguments
- returned from functions

Example:

```go
add := func(a, b int) int {
    return a + b
}

result := add(2, 3)
```

Here:

```text
add
```

holds a function value.

---

# 24. Closures

A **closure** is a function that captures variables from its surrounding scope.

Example:

```go
func sequenceGenerator() func() int {
    i := 0

    return func() int {
        i++
        return i
    }
}
```

Use it:

```go
nextInt := sequenceGenerator()

fmt.Println(nextInt()) // 1
fmt.Println(nextInt()) // 2
```

The returned function remembers the variable:

```text
i
```

Even after `sequenceGenerator()` has returned.

A new generator has separate state:

```go
newInt := sequenceGenerator()

fmt.Println(newInt()) // 1
```

So:

```text
nextInt → own captured i
newInt  → separate captured i
```

Closures are useful for maintaining state and building callbacks/functions with captured context.

---

# 25. `defer`

`defer` schedules a function call to execute when the surrounding function is about to return.

Example:

```go
func example() {
    defer fmt.Println("cleanup")

    fmt.Println("working")
}
```

Conceptually:

```text
working
cleanup
```

### Important details

Arguments to a deferred function are evaluated when the `defer` statement executes.

The deferred function itself runs later, when the surrounding function returns.

---

## 25.1 Multiple `defer` Statements

Multiple deferred calls execute in **LIFO order**:

```text
Last In → First Out
```

Example:

```go
func example() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
}
```

Output:

```text
third
second
first
```

---

# 26. `panic`

Go does not use traditional `try/catch` exceptions.

`panic` stops the normal control flow of the current goroutine and begins panicking.

Example:

```go
func dangerousOperation() {
    panic("critical system failure")
}
```

After a panic, deferred functions still run while the panic is unwinding.

---

# 27. `recover`

`recover` can regain control from a panic.

It is useful **inside a deferred function**.

Example:

```go
func safeExecute() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()

    fmt.Println("Executing task...")

    panic("critical system failure")

    fmt.Println("This line will never execute.")
}
```

Then:

```go
func main() {
    safeExecute()

    fmt.Println("Program continues normally.")
}
```

Conceptually:

```text
safeExecute()
     ↓
panic
     ↓
deferred function runs
     ↓
recover()
     ↓
safeExecute returns
     ↓
main continues
```

### Important

`recover` only recovers a panic in the **same goroutine** in which `recover` is called.

It cannot recover a panic occurring in a different goroutine.

---

# 28. Methods

A method is a function associated with a type through a **receiver**.

Example:

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return v.X + v.Y
}
```

Here:

```go
(v Vertex)
```

is the receiver.

The method can be called on a `Vertex`:

```go
v := Vertex{X: 3, Y: 4}

result := v.Abs()
```

---

# 29. Value Receivers

Example:

```go
func (v Vertex) Abs() float64 {
    return v.X + v.Y
}
```

The receiver is passed as a value.

The method receives its own copy of the `Vertex` value.

Changes to the receiver itself do not modify the original value.

> A value receiver copies the struct value; saying it always performs a "deep copy" is incorrect because Go's value-copy semantics are shallow with respect to fields such as slices, maps, pointers, etc.

---

# 30. Pointer Receivers

Example:

```go
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```

The receiver is a pointer:

```go
*v
```

This allows the method to modify the original `Vertex`.

Example:

```go
v := Vertex{X: 2, Y: 3}

v.Scale(2)
```

Now:

```text
X = 4
Y = 6
```

Pointer receivers are commonly used when:

- the method needs to modify the receiver
- copying the value would be undesirable
- the type has methods that logically operate on one mutable instance

The compiler can often automatically take the address of an addressable value when calling a pointer-receiver method.

---

# 31. Higher-Order Functions

A function that accepts another function as an argument or returns a function is commonly called a **higher-order function**.

Example:

```go
func apply(x int, f func(int) int) int {
    return f(x)
}
```

Here:

```go
f func(int) int
```

is a function parameter.

---

# 32. Generics

Go introduced generics in Go 1.18.

Generics allow functions and types to work with multiple types while retaining compile-time type checking.

Example:

```go
func Map[T any, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))

    for i, v := range slice {
        result[i] = f(v)
    }

    return result
}
```

Here:

```text
T → input element type
U → output element type
```

Usage:

```go
nums := []int{1, 2, 3, 4}

doubled := Map(nums, func(n int) int {
    return n * 2
})

fmt.Println(doubled)
```

Output:

```text
[2 4 6 8]
```

---

# 33. Important Syntax Patterns to Memorize

### Variable

```go
name := "Ankush"
```

### Explicit variable

```go
var age int = 20
```

### Zero-value variable

```go
var count int
```

### Condition

```go
if age >= 18 {
}
```

### `if` initialization

```go
if value := getValue(); value > 10 {
}
```

### Classic loop

```go
for i := 0; i < 10; i++ {
}
```

### While-style loop

```go
for condition {
}
```

### Infinite loop

```go
for {
}
```

### Range loop

```go
for index, value := range items {
}
```

### Ignore a value

```go
for _, value := range items {
}
```

### Function

```go
func add(a, b int) int {
    return a + b
}
```

### Multiple return values

```go
func divide(a, b int) (int, error) {
}
```

### Variadic function

```go
func sum(nums ...int) int {
}
```

### Pointer parameter

```go
func update(value *int) {
}
```

### Method

```go
func (q *Queue) Push() {
}
```

### Anonymous function

```go
func(x int) int {
    return x * 2
}
```

### Closure

```go
func() int {
    // captures variables from outer scope
}
```

---

# 34. Core Mental Model

```text
                         Go
                          │
        ┌─────────────────┼─────────────────┐
        │                 │                 │
    Variables          Control Flow      Functions
        │                 │                 │
   ┌────┼────┐       ┌────┼────┐       ┌────┼────────┐
   │    │    │       │    │    │       │    │        │
  var  :=  zero     if   for  range  params return  methods
 values types       │                  │
                    └── scope         ├── multiple returns
                                      ├── variadic
                                      ├── closures
                                      └── function values
```

---

# 35. Most Important Rules to Remember

1. **Go is statically typed.**
2. **Local unused variables cause compilation errors.**
3. **Go does not perform implicit numeric conversions.**
4. **`:=` declares and initializes a new variable inside a function.**
5. **Every variable gets a zero value when no initial value is provided.**
6. **`if` conditions do not use parentheses.**
7. **`if` and `for` blocks require braces.**
8. **Go has no separate `while` keyword. Use `for`.**
9. **Go has no ternary operator.**
10. **`range` is used to iterate over collections and channels.**
11. **`_` is the blank identifier used when a value is intentionally ignored.**
12. **Functions can return multiple values.**
13. **Returning `(value, error)` is a fundamental Go pattern.**
14. **Variadic parameters use `...`.**
15. **Go passes arguments by value.**
16. **Pointers allow functions/methods to operate on the original value.**
17. **Functions are first-class values.**
18. **Closures can capture variables from their surrounding scope.**
19. **`defer` executes deferred calls when the surrounding function returns.**
20. **Multiple `defer` calls execute in LIFO order.**
21. **`panic` interrupts normal control flow of a goroutine.**
22. **`recover` can recover a panic only from the same goroutine and is normally used inside a deferred function.**
23. **Value receivers receive a copy of the receiver value.**
24. **Pointer receivers can modify the receiver and can avoid copying the receiver value.**
25. **Generics provide reusable, compile-time type-safe abstractions.**

---

# 36. Minimal Practice Program

A small program combining the fundamentals:

```go
package main

import "fmt"

func processJob(job string) string {
    return fmt.Sprintf("Processing: %s", job)
}

func main() {
    jobs := []string{
        "send-email",
        "process-payment",
        "generate-report",
    }

    for i, job := range jobs {
        fmt.Println(i, processJob(job))
    }
}
```

