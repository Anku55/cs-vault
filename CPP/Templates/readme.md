# C++ Templates — Topic List

## 1. Template Fundamentals
- [ ] What Are Templates
- [ ] Generic Programming
- [ ] Function Templates
- [ ] Class Templates
- [ ] Template Parameters
- [ ] Template Arguments
- [ ] Template Instantiation
- [ ] Template Compilation Model
- [ ] Template Type Deduction

## 2. Function Templates
- [ ] Basic Function Templates
- [ ] Multiple Template Parameters
- [ ] Template Parameter Types
- [ ] Explicit Template Arguments
- [ ] Template Argument Deduction
- [ ] Return Type Deduction
- [ ] `auto` with Templates
- [ ] `decltype` with Templates
- [ ] Function Template Overloading
- [ ] Default Template Arguments

## 3. Class Templates
- [ ] Basic Class Templates
- [ ] Multiple Template Parameters
- [ ] Member Functions
- [ ] Static Members
- [ ] Nested Types
- [ ] Constructors
- [ ] Destructors
- [ ] Class Template Instantiation
- [ ] Default Template Parameters

## 4. Non-Type Template Parameters
- [ ] Non-Type Template Parameters
- [ ] Integer Parameters
- [ ] Enum Parameters
- [ ] Pointer Parameters
- [ ] Reference Parameters
- [ ] `auto` Non-Type Parameters
- [ ] Compile-Time Values
- [ ] Template Parameter Objects

## 5. Template Specialization
- [ ] Full Specialization
- [ ] Partial Specialization
- [ ] Function Template Specialization
- [ ] Class Template Specialization
- [ ] Variable Template Specialization
- [ ] Specialization Selection

## 6. Variadic Templates 
- [ ] Variadic Templates
- [ ] Parameter Packs
- [ ] Template Parameter Packs
- [ ] Function Parameter Packs
- [ ] Pack Expansion
- [ ] Fold Expressions
- [ ] Unary Fold
- [ ] Binary Fold
- [ ] Recursive Variadic Templates
- [ ] `sizeof...`

## 7. Template Type Deduction
- [ ] Type Deduction Rules
- [ ] Deduction by Value
- [ ] Deduction by Reference
- [ ] Deduction by Pointer
- [ ] `const` Deduction
- [ ] Array Deduction
- [ ] Function Deduction
- [ ] Forwarding References
- [ ] Reference Collapsing

## 8. `auto` & `decltype`
- [ ] `auto`
- [ ] `decltype`
- [ ] `decltype(auto)`
- [ ] Type Deduction
- [ ] Value Categories
- [ ] `decltype` with Expressions
- [ ] `auto` Return Types

## 9. Perfect Forwarding 
- [ ] Lvalues
- [ ] Rvalues
- [ ] Rvalue References
- [ ] Forwarding References
- [ ] `std::move`
- [ ] `std::forward`
- [ ] Perfect Forwarding
- [ ] Reference Collapsing
- [ ] Universal Reference Terminology
- [ ] Forwarding Constructor
- [ ] Forwarding Function

## 10. Template Metaprogramming
- [ ] Compile-Time Programming
- [ ] Template Recursion
- [ ] Compile-Time Computation
- [ ] Type Computation
- [ ] Value Computation
- [ ] Type Traits
- [ ] Compile-Time Conditions
- [ ] `constexpr`
- [ ] `consteval`
- [ ] `constinit`

## 11. Type Traits 
- [ ] `std::is_same`
- [ ] `std::is_integral`
- [ ] `std::is_floating_point`
- [ ] `std::is_pointer`
- [ ] `std::is_reference`
- [ ] `std::is_const`
- [ ] `std::is_array`
- [ ] `std::is_class`
- [ ] `std::is_function`
- [ ] `std::is_base_of`
- [ ] `std::is_constructible`
- [ ] `std::is_copy_constructible`
- [ ] `std::is_move_constructible`
- [ ] `std::is_default_constructible`

## 12. Type Transformations
- [ ] `std::remove_reference`
- [ ] `std::remove_const`
- [ ] `std::remove_cv`
- [ ] `std::remove_pointer`
- [ ] `std::add_reference`
- [ ] `std::add_const`
- [ ] `std::add_pointer`
- [ ] `std::decay`
- [ ] `std::enable_if`
- [ ] `std::conditional`
- [ ] `std::common_type`
- [ ] `std::underlying_type`

## 13. SFINAE
- [ ] SFINAE Concept
- [ ] Substitution Failure
- [ ] `std::enable_if`
- [ ] `std::void_t`
- [ ] Detection Idiom
- [ ] Function SFINAE
- [ ] Class SFINAE
- [ ] Return-Type SFINAE
- [ ] Parameter-Type SFINAE
- [ ] SFINAE vs Concepts

## 14. Concepts & Constraints 
- [ ] Concepts
- [ ] `concept`
- [ ] `requires`
- [ ] Requires Expressions
- [ ] Simple Requirements
- [ ] Type Requirements
- [ ] Compound Requirements
- [ ] Nested Requirements
- [ ] Constrained Templates
- [ ] Constrained Functions
- [ ] Constrained Classes
- [ ] Standard Concepts
- [ ] Custom Concepts
- [ ] Concepts vs SFINAE

## 15. `if constexpr`
- [ ] Compile-Time Branching
- [ ] `if constexpr`
- [ ] Type-Dependent Code
- [ ] Template Instantiation
- [ ] `if constexpr` vs Runtime `if`

## 16. Template Overloading
- [ ] Function Template Overloading
- [ ] Class Template Overloading Concepts
- [ ] Overload Resolution
- [ ] Template vs Non-Template Overloads
- [ ] Partial Ordering
- [ ] Specialization vs Overloading
- [ ] Constrained Overloads

## 17. Dependent Names 
- [ ] Dependent Types
- [ ] Dependent Names
- [ ] `typename`
- [ ] `template` Disambiguator
- [ ] Type-Dependent Expressions
- [ ] Value-Dependent Expressions
- [ ] Two-Phase Name Lookup

## 18. Template Inheritance
- [ ] Template Base Classes
- [ ] Derived Class Templates
- [ ] Dependent Base Classes
- [ ] CRTP
- [ ] Static Polymorphism
- [ ] Template Mixins
- [ ] Template-Based Interfaces

## 19. CRTP 
- [ ] Curiously Recurring Template Pattern
- [ ] Static Polymorphism
- [ ] Compile-Time Dispatch
- [ ] CRTP Interfaces
- [ ] CRTP Mixins
- [ ] CRTP vs Virtual Functions

## 20. Template Aliases
- [ ] `using`
- [ ] Alias Templates
- [ ] Template Type Aliases
- [ ] Alias Specialization
- [ ] Nested Template Aliases

## 21. Variable Templates
- [ ] Variable Templates
- [ ] `template <typename T>`
- [ ] Static Compile-Time Variables
- [ ] Variable Template Specialization

## 22. Template Template Parameters
- [ ] Template Template Parameters
- [ ] Passing Templates as Arguments
- [ ] Template Template Parameter Packs
- [ ] STL-Style Generic Containers

## 23. Templates & `constexpr`
- [ ] `constexpr` Functions
- [ ] `constexpr` Variables
- [ ] `constexpr` Constructors
- [ ] Compile-Time Evaluation
- [ ] `consteval`
- [ ] `constinit`
- [ ] Templates + `constexpr`
- [ ] Compile-Time Data Structures

## 24. Templates & Lambdas
- [ ] Generic Lambdas
- [ ] Lambda Template Parameters
- [ ] Explicit Template Parameters
- [ ] Variadic Generic Lambdas
- [ ] Template Lambdas
- [ ] Lambdas + Concepts
- [ ] Lambdas + Perfect Forwarding

## 25. Templates & STL 
- [ ] Generic Containers
- [ ] Generic Algorithms
- [ ] Iterators
- [ ] Allocators
- [ ] Custom Comparators
- [ ] Function Objects
- [ ] `std::function`
- [ ] Ranges
- [ ] Concepts in STL
- [ ] Template-Based Data Structures

## 26. Templates & Memory Management
- [ ] Generic Smart Pointers
- [ ] Custom Allocators
- [ ] Memory Pools
- [ ] Generic Resource Management
- [ ] RAII Templates
- [ ] Policy-Based Memory Management

## 27. Templates & Concurrency
- [ ] Generic Thread Functions
- [ ] Variadic Thread Arguments
- [ ] Generic Locks
- [ ] Generic Thread Pools
- [ ] Generic Task Queues
- [ ] Compile-Time Concurrency Utilities
- [ ] Template-Based Synchronization

## 28. Templates & System Programming 
- [ ] Generic Buffers
- [ ] Generic Socket Wrappers
- [ ] Generic File Wrappers
- [ ] Generic Resource Handles
- [ ] Policy-Based Design
- [ ] Zero-Cost Abstractions
- [ ] Compile-Time Configuration
- [ ] Type-Safe System APIs
- [ ] Generic Serialization
- [ ] Generic Protocol Parsers

## 29. Advanced Template Design
- [ ] Policy-Based Design
- [ ] Traits Classes
- [ ] Type Erasure
- [ ] Static Polymorphism
- [ ] Compile-Time Dispatch
- [ ] Tag Dispatch
- [ ] Detection Idiom
- [ ] Expression Templates
- [ ] CRTP
- [ ] Metafunctions
- [ ] Compile-Time Reflection Concepts

## 30. Template Compilation & Linking
- [ ] Template Instantiation
- [ ] Implicit Instantiation
- [ ] Explicit Instantiation
- [ ] Explicit Specialization
- [ ] One Definition Rule
- [ ] Header-Only Templates
- [ ] Template Linker Errors
- [ ] Separate Compilation
- [ ] Code Bloat
- [ ] Template Compilation Time

## 31. Template Performance
- [ ] Zero-Cost Abstractions
- [ ] Compile-Time Optimization
- [ ] Inlining
- [ ] Code Generation
- [ ] Template Code Bloat
- [ ] Binary Size
- [ ] Compilation-Time Optimization
- [ ] Runtime vs Compile-Time Tradeoffs

## 32. Redis / Systems Project Relevant 
- [ ] Generic Data Structures
- [ ] Generic Hash Table
- [ ] Generic Linked List
- [ ] Generic Dynamic Array
- [ ] Generic Memory Pool
- [ ] Custom Allocator
- [ ] Generic Command Handler
- [ ] Variadic Command Dispatch
- [ ] Type-Safe Command Registration
- [ ] Compile-Time Configuration
- [ ] Policy-Based Design
- [ ] Zero-Cost Abstractions
- [ ] Concepts for API Constraints
- [ ] Perfect Forwarding
- [ ] `std::variant` + Templates
- [ ] Generic Serialization
- [ ] Generic Protocol Parsing