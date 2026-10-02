# C++ Error Handling — Topic List

## 1. Error Handling Fundamentals
- [ ] Errors vs Exceptions
- [ ] Compile-Time Errors
- [ ] Linker Errors
- [ ] Runtime Errors
- [ ] Logic Errors
- [ ] Undefined Behavior
- [ ] Error Detection
- [ ] Error Reporting
- [ ] Error Recovery

## 2. Return-Based Error Handling
- [ ] Error Codes
- [ ] Boolean Return Values
- [ ] Integer Error Codes
- [ ] Sentinel Values
- [ ] Output Parameters
- [ ] `errno`
- [ ] `std::error_code`
- [ ] `std::error_condition`

## 3. Exceptions
- [ ] Exception Basics
- [ ] `throw`
- [ ] `try`
- [ ] `catch`
- [ ] Exception Propagation
- [ ] Stack Unwinding
- [ ] Multiple `catch` Blocks
- [ ] Catch by Reference
- [ ] Catch-All `catch(...)`
- [ ] Rethrowing Exceptions
- [ ] Exception Matching

## 4. Standard Exceptions
- [ ] `std::exception`
- [ ] `std::runtime_error`
- [ ] `std::logic_error`
- [ ] `std::invalid_argument`
- [ ] `std::out_of_range`
- [ ] `std::length_error`
- [ ] `std::domain_error`
- [ ] `std::overflow_error`
- [ ] `std::underflow_error`
- [ ] `std::bad_alloc`
- [ ] `std::bad_cast`
- [ ] `std::bad_typeid`

## 5. Custom Exceptions
- [ ] Custom Exception Classes
- [ ] Inheriting from `std::exception`
- [ ] `what()`
- [ ] Exception Messages
- [ ] Exception Categories
- [ ] Nested Exceptions
- [ ] `std::throw_with_nested`
- [ ] `std::rethrow_if_nested`

## 6. Exception Safety
- [ ] Exception-Safe Code
- [ ] Basic Exception Guarantee
- [ ] Strong Exception Guarantee
- [ ] No-Throw Guarantee
- [ ] Exception-Neutral Code
- [ ] Resource Safety
- [ ] State Consistency

## 7. RAII & Error Handling
- [ ] RAII
- [ ] Resource Cleanup During Exceptions
- [ ] Destructor During Stack Unwinding
- [ ] Smart Pointers
- [ ] Lock Guards
- [ ] File Resource Management
- [ ] Socket Resource Management

## 8. `noexcept`
- [ ] `noexcept`
- [ ] Conditional `noexcept`
- [ ] `noexcept` Operator
- [ ] `noexcept` Functions
- [ ] `std::terminate`
- [ ] Move Constructors & `noexcept`
- [ ] Move Assignment & `noexcept`

## 9. Assertions
- [ ] `assert()`
- [ ] Preconditions
- [ ] Postconditions
- [ ] Invariants
- [ ] Debug Assertions
- [ ] Release Builds
- [ ] `static_assert`

## 10. Optional-Based Error Handling
- [ ] `std::optional`
- [ ] Representing Absence
- [ ] `has_value()`
- [ ] `value()`
- [ ] `value_or()`
- [ ] `operator*`
- [ ] `operator->`
- [ ] Optional vs Exceptions

## 11. Expected-Based Error Handling
- [ ] `std::expected`
- [ ] Value/Error Model
- [ ] `std::unexpected`
- [ ] `has_value()`
- [ ] `value()`
- [ ] `error()`
- [ ] `value_or()`
- [ ] Expected vs Exceptions
- [ ] Expected vs Error Codes

## 12. Error Handling with Functions
- [ ] Returning Error Codes
- [ ] Returning `optional`
- [ ] Returning `expected`
- [ ] Throwing Exceptions
- [ ] Error Propagation
- [ ] Error Transformation
- [ ] Error Recovery
- [ ] Error Ownership

## 13. Error Handling with Constructors
- [ ] Constructor Exceptions
- [ ] Failed Construction
- [ ] RAII During Construction
- [ ] Destructor During Failed Construction
- [ ] Factory Functions for Fallible Objects

## 14. Error Handling with Destructors
- [ ] Destructor Safety
- [ ] Why Destructors Should Not Throw
- [ ] `noexcept` Destructors
- [ ] Exceptions During Stack Unwinding
- [ ] `std::terminate`

## 15. Error Handling in Classes
- [ ] Class Invariants
- [ ] Validation
- [ ] Invalid Object States
- [ ] Exception-Safe Classes
- [ ] Error Reporting
- [ ] Error Recovery
- [ ] Strong Exception Guarantee

## 16. Error Handling in File I/O
- [ ] File Open Errors
- [ ] Read Errors
- [ ] Write Errors
- [ ] EOF Handling
- [ ] Stream Error States
- [ ] `fail()`
- [ ] `bad()`
- [ ] `eof()`
- [ ] `good()`
- [ ] Exception-Based Stream Handling

## 17. Error Handling in Networking
- [ ] Socket Errors
- [ ] Connection Errors
- [ ] Read Errors
- [ ] Write Errors
- [ ] Timeout Errors
- [ ] Connection Reset
- [ ] Connection Closed
- [ ] Partial Reads
- [ ] Partial Writes
- [ ] `errno`
- [ ] System Call Error Handling

## 18. Error Handling in Multithreading
- [ ] Exceptions in Threads
- [ ] Exception Propagation Across Threads
- [ ] `std::exception_ptr`
- [ ] `std::current_exception`
- [ ] `std::rethrow_exception`
- [ ] `std::future`
- [ ] `std::promise`
- [ ] Thread-Safe Error Reporting

## 19. System-Level Errors
- [ ] `errno`
- [ ] `strerror`
- [ ] `perror`
- [ ] POSIX Error Codes
- [ ] System Call Failures
- [ ] Resource Exhaustion
- [ ] Permission Errors
- [ ] File Not Found
- [ ] Connection Refused
- [ ] Interrupted System Calls

## 20. Logging & Error Reporting
- [ ] Error Logs
- [ ] Log Levels
- [ ] Error Context
- [ ] Error Messages
- [ ] Error IDs
- [ ] Structured Logging
- [ ] Stack Traces
- [ ] Debug vs Production Errors

## 21. Error Handling Design
- [ ] Fail Fast
- [ ] Error Propagation
- [ ] Error Recovery
- [ ] Error Translation
- [ ] Error Boundaries
- [ ] Local vs Global Error Handling
- [ ] Avoiding Silent Failures
- [ ] Avoiding Over-Catching
- [ ] Avoiding Exception Abuse
- [ ] Consistent Error Strategy

## 22. Debugging & Diagnostics
- [ ] GDB
- [ ] Breakpoints
- [ ] Stack Traces
- [ ] Core Dumps
- [ ] AddressSanitizer
- [ ] UndefinedBehaviorSanitizer
- [ ] Valgrind
- [ ] Assertions
- [ ] Logging
- [ ] Crash Diagnostics

## 23. Redis/System Project Relevant
- [ ] Command Validation Errors
- [ ] Invalid Command Errors
- [ ] Invalid Arguments
- [ ] Type Errors
- [ ] Key Not Found
- [ ] Connection Errors
- [ ] Socket Errors
- [ ] Serialization Errors
- [ ] Memory Allocation Errors
- [ ] Timeout Handling
- [ ] Client Disconnect Handling
- [ ] Partial Request Handling
- [ ] Protocol Errors
- [ ] Graceful Shutdown
- [ ] Error Logging
- [ ] Error Propagation