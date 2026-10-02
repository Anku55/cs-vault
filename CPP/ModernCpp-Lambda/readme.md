# C++ Lambda Expressions — Topic List

## 1. Lambda Fundamentals
- [ ] Lambda Expression Syntax
- [ ] Lambda Function Structure
- [ ] Lambda Parameters
- [ ] Lambda Return Type
- [ ] Lambda Body
- [ ] Calling a Lambda
- [ ] Storing a Lambda in `auto`
- [ ] Lambda as a Function Object

## 2. Lambda Captures
- [ ] Capture List `[]`
- [ ] Capture by Value `[=]`
- [ ] Capture by Reference `[&]`
- [ ] Capture Specific Variables `[x, y]`
- [ ] Mixed Capture `[=, &x]`
- [ ] Mixed Capture `[&, x]`
- [ ] Capture `this`
- [ ] Capture `*this`
- [ ] Init Capture / Generalized Capture
- [ ] Capturing `const` Variables
- [ ] Capturing References

## 3. Lambda Parameters
- [ ] No Parameters
- [ ] Single Parameter
- [ ] Multiple Parameters
- [ ] Default Parameters
- [ ] Parameter Types
- [ ] Reference Parameters
- [ ] `const` Parameters
- [ ] Variadic Lambda Parameters
- [ ] Parameter Packs

## 4. Lambda Return Types
- [ ] Automatic Return Type
- [ ] Explicit Return Type
- [ ] Trailing Return Type
- [ ] `decltype` with Lambdas
- [ ] `auto` Return Type
- [ ] Returning References
- [ ] Returning Objects

## 5. `mutable` Lambdas
- [ ] `mutable` Keyword
- [ ] Modifying Captured-by-Value Variables
- [ ] Mutable Lambda State
- [ ] Mutable vs Non-Mutable Lambdas

## 6. Generic Lambdas 
- [ ] Generic Lambda
- [ ] `auto` Parameters
- [ ] Multiple `auto` Parameters
- [ ] Generic Lambda Type Deduction
- [ ] Generic Lambda with References
- [ ] Generic Lambda with `const`
- [ ] Generic Lambda with Templates

## 7. Template Lambdas
- [ ] Template Parameter Lists
- [ ] Explicit Template Parameters
- [ ] Generic Template Lambdas
- [ ] Variadic Template Lambdas
- [ ] Template Lambda + Concepts
- [ ] Template Lambda + `requires`

## 8. Lambda Type System
- [ ] Lambda Closure Type
- [ ] Closure Object
- [ ] Unique Lambda Type
- [ ] Lambda Type Is Unnamed
- [ ] `decltype(lambda)`
- [ ] Lambda Object Storage
- [ ] Lambda Object Lifetime

## 9. `std::function`
- [ ] `std::function`
- [ ] Storing Lambdas
- [ ] Passing Lambdas to Functions
- [ ] Lambda Conversion to `std::function`
- [ ] `std::function` Overhead
- [ ] Lambda vs `std::function`

## 10. Lambda with STL 
- [ ] Lambda with `std::sort`
- [ ] Lambda with `std::stable_sort`
- [ ] Lambda with `std::find_if`
- [ ] Lambda with `std::count_if`
- [ ] Lambda with `std::for_each`
- [ ] Lambda with `std::transform`
- [ ] Lambda with `std::remove_if`
- [ ] Lambda with `std::partition`
- [ ] Lambda with `std::accumulate`
- [ ] Lambda with `std::min_element`
- [ ] Lambda with `std::max_element`
- [ ] Lambda as Custom Comparator

## 11. Lambda & Algorithms
- [ ] Custom Sorting
- [ ] Custom Searching
- [ ] Filtering
- [ ] Transformation
- [ ] Aggregation
- [ ] Predicate Functions
- [ ] Comparison Functions
- [ ] Stateful Algorithms

## 12. Lambda & Objects
- [ ] Lambda Capturing Class Members
- [ ] Capturing `this`
- [ ] Capturing `*this`
- [ ] Lambda Inside Member Functions
- [ ] Lambda Access to Private Members
- [ ] Lambda with Static Members
- [ ] Lambda Object Lifetime

## 13. Lambda & Lifetime 
- [ ] Capture Lifetime
- [ ] Dangling Reference Capture
- [ ] Dangling `this` Capture
- [ ] Lambda Returned from Function
- [ ] Lambda Stored for Later Execution
- [ ] Safe Capture by Value
- [ ] Safe Capture by Reference

## 14. Lambda & Move Semantics
- [ ] Move Capture
- [ ] Init Capture
- [ ] Capturing `std::move`
- [ ] Move-Only Lambdas
- [ ] Capturing `unique_ptr`
- [ ] Capturing `shared_ptr`
- [ ] Capturing `weak_ptr`

## 15. Lambda & Concurrency 
- [ ] Lambda with `std::thread`
- [ ] Lambda with `std::jthread`
- [ ] Lambda with `std::async`
- [ ] Lambda with `std::future`
- [ ] Lambda with Thread Pools
- [ ] Lambda with Task Queues
- [ ] Lambda with Mutexes
- [ ] Lambda with Condition Variables
- [ ] Lambda Capture in Multithreading
- [ ] Thread-Safe Lambda State

## 16. Lambda & Callbacks
- [ ] Callback Functions
- [ ] Lambda Callbacks
- [ ] Asynchronous Callbacks
- [ ] Event Callbacks
- [ ] Timer Callbacks
- [ ] Network Callbacks
- [ ] Error Callbacks

## 17. Lambda & Functional Programming
- [ ] Higher-Order Functions
- [ ] Functions as Arguments
- [ ] Functions as Return Values
- [ ] Closures
- [ ] Function Composition
- [ ] Predicates
- [ ] Stateful Functions
- [ ] Pure Functions

## 18. Lambda & `constexpr`
- [ ] `constexpr` Lambdas
- [ ] Compile-Time Lambdas
- [ ] `consteval` Lambdas
- [ ] Lambda in Constant Expressions
- [ ] Compile-Time Algorithms

## 19. Lambda & Concepts
- [ ] Constrained Generic Lambdas
- [ ] `requires` with Lambdas
- [ ] Standard Concepts
- [ ] Custom Concepts
- [ ] Type Constraints

## 20. Lambda & Ranges
- [ ] Lambdas with `std::ranges`
- [ ] `filter`
- [ ] `transform`
- [ ] `take`
- [ ] `drop`
- [ ] Lambda-Based Views
- [ ] Lazy Evaluation
- [ ] Range Pipelines

## 21. Lambda Performance
- [ ] Lambda Object Size
- [ ] Capture Overhead
- [ ] Value vs Reference Capture
- [ ] Lambda Inlining
- [ ] `std::function` Overhead
- [ ] Function Pointer vs Lambda
- [ ] Zero-Cost Abstractions
- [ ] Allocation with `std::function`

## 22. Lambda vs Function Pointer
- [ ] Function Pointer
- [ ] Lambda
- [ ] Non-Capturing Lambda
- [ ] Capturing Lambda
- [ ] Conversion to Function Pointer
- [ ] Runtime vs Compile-Time Behavior

## 23. Lambda vs Functor
- [ ] Function Objects
- [ ] Lambda Objects
- [ ] Operator `()`
- [ ] Stateful Functors
- [ ] Lambda vs Functor
- [ ] When to Use Each

## 24. Advanced Lambda Topics
- [ ] Nested Lambdas
- [ ] Recursive Lambdas
- [ ] Self-Referential Lambdas
- [ ] Lambda Recursion with `std::function`
- [ ] Lambda Recursion with `auto`
- [ ] Lambda Returning Lambda
- [ ] Lambda Capturing Lambda
- [ ] Higher-Order Lambdas

## 25. System Programming / Redis-Relevant 
- [ ] Lambda-Based Event Handlers
- [ ] Lambda-Based Command Dispatch
- [ ] Lambda-Based Callbacks
- [ ] Lambda with Socket Events
- [ ] Lambda with Event Loops
- [ ] Lambda with Thread Pools
- [ ] Lambda with Task Queues
- [ ] Lambda with Timers
- [ ] Lambda with Async I/O
- [ ] Lambda with Custom Comparators
- [ ] Lambda-Based Resource Cleanup
- [ ] Lambda + RAII
- [ ] Lambda + Smart Pointers
- [ ] Lambda + `std::variant`
- [ ] Lambda + `std::visit`

## 26. Practice
- [ ] Sort Objects Using Lambda
- [ ] Filter a Vector
- [ ] Transform a Vector
- [ ] Count Elements Using Predicate
- [ ] Custom Priority Queue Comparator
- [ ] Custom `map` Comparator
- [ ] Recursive Lambda
- [ ] Generic Lambda
- [ ] Lambda-Based Thread Pool
- [ ] Lambda-Based Event Dispatcher
- [ ] Lambda-Based Task Queue
- [ ] Lambda-Based Command Dispatcher