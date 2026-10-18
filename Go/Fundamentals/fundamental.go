package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
	/*
		In Go (Golang), variables are explicitly declared containers
		used by the compiler to enforce type-correctness. Go is a statically
		 typed language, meaning every variable has a fixed type determined
		 at compile time, and it cannot change later
	*/
	/*
		Go offers two main ways to declare and initialize variable

	*/
	// the Var key word
	/*
		Can be used inside or outside functions (package level).
		You can explicitly state the type or let Go infer it
	*/
	//var age int = 30
	//var city = "Etawah"
	/*var (
		name    string = "Alice"
		isAdmin bool   = true
	)*/
	//=============Short Variable Declaration (:=)============
	//count := 10
	//agee := 18
	//class := "frontier"
	/*

		Zero Values

		If you declare a variable without giving it an initial value,
		Go automatically assigns it a "zero value" instead of leaving
		it uninitialized or garbage-filled.
		• 0 for integers and floats
		• "" (empty string) for strings
		• false for booleans
		• nil for pointers, slices, maps, channels, and interfaces
	*/
	/*

		2. Built-in Data Types

		Go's built-in data types are categorized into four categories: Basic, Aggregate, Reference, and Interface.

		Basic Types


		Category	Type	Description
		Boolean	bool	Holds either true or false.
		String	string	A sequence of characters enclosed in double quotes "".
		Signed Integers	int, int8, int16, int32, int64	Holds negative or positive whole
		numbers. int is 32-bit on 32-bit systems and 64-bit on 64-bit systems.
		Unsigned Integers	uint, uint8, uint16, uint32, uint64	Holds positive-only integers.
		Floating-Point	float32, float64	Decimal numbers. float64 is the default for
		inferred decimals.
		Aliases	byte, rune	byte is an alias for uint8. rune is an alias for int32
		 (represents a Unicode code point).
		Complex	complex64, complex128	Numbers with real and imaginary parts.

		Advanced/Composite Types

		• Aggregate Types: Arrays (fixed-size lists) and Structs (custom collections of fields).
		• Reference Types: Pointers (memory addresses), Slices (dynamic lists), Maps (key/value pairs), Channels (concurrency pipes), and Functions.
		• Interface Types: Defines a set of method signatures that other types must implement.

	*/
	/*

	    Strict Rules to Keep in Mind

	   1. No Unused Variables: If you declare a variable inside a function and do not use it,
	    the Go compiler will flag an error and refuse to compile.
	   2. No Implicit Type Conversion: Go does not convert numbers automatically.
	    You cannot add an int to a float64 without an explicit conversion:
	*/

	/*

			You do not use parentheses () around the condition, but curly braces {}
			 are strictly required, even for single-line blocks. The condition must
			 evaluate directly to a bool (Go has no implicit truthy/falsy values like 0, "", or nil).
			package main

		import "fmt"

		func main() {
		    score := 85

		    if score >= 60 {
		        fmt.Println("You passed!") // Braces are mandatory
		    }
		}

	*/
	/*

		3. The Init Statement (Statement Before Condition)

		Go allows you to execute a short initialization statement right before the condition,
		 separated by a semicolon ;. Variables declared in this init block are only scoped within
		  the if, else if, and else blocks. This keeps your code clean and prevents variables from leaking
			// val is declared, initialized, and evaluated all on one line
		if val := getCount(); val > 10 {
		    fmt.Println("High count:", val)
		} else {
		    fmt.Println("Low count:", val) // val is still accessible here
		}
		// val is NOT accessible out here
	*/
	//No Ternary Operator:

	/*

		1. The Classic for Loop

		Uses three components separated by semicolons: initialization, condition, and post-statement. Parentheses are not used, and curly braces are required.
		go
		package main
		import "fmt"

		func main() {
		    // Standard 3-component loop
		    for i := 0; i < 5; i++ {
		        fmt.Println(i)
		    }
		}
		Use code with caution.

		2. The "While" Loop Pattern

		To write a while loop, omit the initialization and post-statement, keeping only the condition.
		go
		n := 1
		for n <= 5 { // Runs as long as this condition is true
		    fmt.Println(n)
		    n++
		}
		Use code with caution.

		3. Infinite Loop

		Omitting all conditions creates an infinite loop, which runs until an explicit break or return is triggered.
		go
		for {
		    fmt.Println("Running indefinitely...")
		    break // Stops the loop immediately
		}
		Use code with caution.

		4. The for-range Loop

		Used to iterate over collections like slices, arrays, strings, maps, and channels. It returns the index (or key) and the value.
		go
		// Looping over a slice
		names := []string{"Alice", "Bob", "Charlie"}
		for index, value := range names {
		    fmt.Printf("Index: %d, Value: %s\n", index, value)
		}

		// Discarding the index using the blank identifier (_)
		for _, value := range names {
		    fmt.Println(value)
		}

		// Iterating a specific number of times (Go 1.22+)
		for i := range 5 {
		    fmt.Println(i) // Prints 0 through 4
		}
	*/
	/*
		In the Go programming language (Golang), a function is a reusable block of code initialized using the func keyword. Go functions are highly flexible, featuring unique behaviors like multiple return values, named returns, and variadic parameters. [1] (https://www.geeksforgeeks.org/go-language/functions-in-go-language/), [2] (https://www.w3schools.com/go/go_functions.php), [3] (https://www.youtube.com/watch?v=CrgD_otSzDI&t=6)

	1. Basic Function Syntax

	A standard function includes a name, input parameters with their types, a return type, and the body of code. [1] (https://www.youtube.com/watch?v=HDgV7CA8Ae8), [2] (https://www.geeksforgeeks.org/go-language/functions-in-go-language/)
	go
	package main

	import "fmt"

	// A basic function with parameters and a single return value
	func add(a int, b int) int {
	    return a + b
	}

	func main() {
	    result := add(5, 3)
	    fmt.Println("Sum:", result) // Output: Sum: 8
	}
	Use code with caution.
	Note: If consecutive parameters share the same type, you can omit the type for the earlier parameters and just declare it at the end (e.g., func add(a, b int) int). [1] (https://gobyexample.com/functions), [2] (https://www.youtube.com/watch?v=a6V1Aic0HYQ&t=20)

	2. Core Features & Core Concepts


	↩️ Multiple Return Values

	Unlike many languages, Go functions can natively return more than one result. This pattern is heavily utilized across the language for returning a result alongside an error. [1] (https://www.youtube.com/watch?v=CrgD_otSzDI&t=6), [2] (https://www.geeksforgeeks.org/go-language/functions-in-go-language/)
	go
	func divide(a, b int) (int, error) {
	    if b == 0 {
	        return 0, fmt.Errorf("cannot divide by zero")
	    }
	    return a / b, nil
	}
	Use code with caution.

	🏷️ Named Return Values

	You can explicitly name your return variables in the function signature. They act as variables defined at the top of the function and allow for a "naked" or "bare" return statement. [1] (https://www.youtube.com/watch?v=mKPiCLjgrG0&vl=en), [2] (https://www.youtube.com/watch?v=CrgD_otSzDI&t=6)
	go
	func rectangleArea(width, height int) (area int) {
	    area = width * height
	    return // automatically returns 'area'
	}
	Use code with caution.



	A variadic function can accept an arbitrary number of trailing arguments by trailing the parameter type with .... Inside the function, the parameter acts as a slice. [1] (https://www.geeksforgeeks.org/go-language/functions-in-go-language/), [2] (https://www.youtube.com/watch?v=mKPiCLjgrG0&vl=en)
	go
	func sumAll(nums ...int) int {
	    total := 0
	    for _, num := range nums {
	        total += num
	    }
	    return total
	}
	// Usage: sumAll(1, 2, 3, 4)
	Use code with caution.



	Go always passes arguments by value, meaning the function gets a local copy of the data. If you need a function to modify the original variable, you must pass a pointer (*type) instead. [1] (https://www.geeksforgeeks.org/go-language/functions-in-go-language/)

	Feature	Passing by Value	Passing by Pointer
	What is passed	A complete copy of the data	A copy of the memory address
	Impact	Original variable remains safe and unchanged	Original variable can be directly modified
	Syntax example	func update(val int)	func update(val *int)





	*/
	/*
		To write high-performance and idiomatic Go, you must leverage its unique type system and runtime lifecycle controls.
	Here are the advanced concepts that define professional Go function architecture.

	1. First-Class Functions and Closures

	In Go, functions are first-class citizens. This means they can be assigned to variables, passed as arguments to other functions, and returned from functions.
	A closure is a special type of anonymous function that "closes over" and retains access to local variables from its outer scope, even after the outer function has finished executing.
	go
	package main

	import "fmt"

	// sequenceGenerator returns a function that returns an int
	func sequenceGenerator() func() int {
	    i := 0 // This variable is captured by the closure below
	    return func() int {
	        i++
	        return i
	    }
	}

	func main() {
	    nextInt := sequenceGenerator()

	    fmt.Println(nextInt()) // Output: 1
	    fmt.Println(nextInt()) // Output: 2

	    // A new generator gets its own isolated state
	    newInt := sequenceGenerator()
	    fmt.Println(newInt())  // Output: 1
	}
	Use code with caution.

	2. Lifecycles: defer, panic, and recover

	Go does not use traditional try/catch exceptions. Instead, it manages runtime lifecycles and unexpected exceptions through a specific control flow mechanism:
	• defer: Schedules a function call to run immediately after the surrounding function returns. Arguments to deferred functions are evaluated immediately, but the execution happens at the end. Multiple defers execute in LIFO (Last-In, First-Out) order.
	• panic: Stops the ordinary control flow of the goroutine and begins panicking.
	• recover: A built-in function that regains control of a panicking goroutine. It is only useful inside deferred functions.
	go
	package main

	import "fmt"

	func safeExecute() {
	    // recover must be inside a deferred function
	    defer func() {
	        if r := recover(); r != nil {
	            fmt.Printf("Recovered from panic: %v\n", r)
	        }
	    }()

	    fmt.Println("Executing task...")
	    panic("critical system failure") // Triggers control flow interruption
	    fmt.Println("This line will never execute.")
	}

	func main() {
	    safeExecute()
	    fmt.Println("Program continues normally.")
	}
	Use code with caution.

	3. Pointer Receivers vs. Value Receivers (Methods)

	Methods are functions containing an explicit receiver argument before the function name. Deciding whether to use a value or pointer receiver impacts mutations and memory performance.
	go
	type Vertex struct {
	    X, Y float64
	}

	// Value receiver: Operates on a deep copy of Vertex. Cannot modify originals.
	func (v Vertex) Abs() float64 {
	    return v.X + v.Y
	}

	// Pointer receiver: Operates on the direct reference. Modifies original data.
	// It also avoids copying the struct on every function call.
	func (v *Vertex) Scale(f float64) {
	    v.X = v.X * f
	    v.Y = v.Y * f
	}
	Use code with caution.

	4. Advanced Patterns: Higher-Order Functions & Generics

	With the introduction of Generics (Go 1.18+), you can write type-agnostic higher-order utility functions that work on arbitrary structures while preserving strict compile-time type safety.

	Map/Filter/Reduce Paradigm with Generics

	go
	package main

	import "fmt"

	// Map transforms a slice of type T into a slice of type U using function f
	func Map[T any, U any](slice []T, f func(T) U) []U {
	    result := make([]U, len(slice))
	    for i, v := range slice {
	        result[i] = f(v)
	    }
	    return result
	}

	func main() {
	    nums := []int{1, 2, 3, 4}

	    // Double each number dynamically
	    doubled := Map(nums, func(n int) int {
	        return n * 2
	    })

	    fmt.Println(doubled) // Output: [2, 4, 6, 8]
	}
	Use code with caution.

	Summary: Execution Performance Behavior



	*/
}
