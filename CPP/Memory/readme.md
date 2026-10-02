# C++ Memory Management — 

## 1. Memory Fundamentals
- [ ] Process Memory Layout
- [ ] Code / Text Segment
- [ ] Read-Only Data
- [ ] Initialized Data Segment
- [ ] BSS Segment
- [ ] Heap
- [ ] Stack
- [ ] Static Storage
- [ ] Memory Addresses
- [ ] Virtual Memory Basics

## 2. Stack Memory
- [ ] Stack Frames
- [ ] Local Variables
- [ ] Function Parameters
- [ ] Return Address
- [ ] Stack Growth
- [ ] Stack Overflow
- [ ] Recursion & Stack
- [ ] Stack vs Heap

## 3. Heap Memory
- [ ] Dynamic Allocation
- [ ] Heap Allocation
- [ ] Heap Deallocation
- [ ] Heap Fragmentation
- [ ] Allocation Overhead
- [ ] Heap vs Stack

## 4. Pointers
- [ ] Pointer Basics
- [ ] Address-of `&`
- [ ] Dereference `*`
- [ ] `nullptr`
- [ ] Pointer Arithmetic
- [ ] Pointer to Pointer
- [ ] Pointer to Object
- [ ] Pointer to Function
- [ ] Pointer to Member
- [ ] `const` with Pointers

## 5. References
- [ ] Lvalue References
- [ ] Rvalue References
- [ ] Reference Lifetime
- [ ] References vs Pointers
- [ ] `const` References
- [ ] Dangling References

## 6. Dynamic Memory
- [ ] `new`
- [ ] `delete`
- [ ] `new[]`
- [ ] `delete[]`
- [ ] Dynamic Arrays
- [ ] Dynamic Objects
- [ ] Allocation Failure
- [ ] `std::nothrow`

## 7. Memory Ownership
- [ ] Ownership
- [ ] Single Ownership
- [ ] Shared Ownership
- [ ] Non-Owning References
- [ ] Ownership Transfer
- [ ] Resource Lifetime
- [ ] Ownership Models

## 8. RAII
- [ ] RAII Concept
- [ ] Resource Acquisition
- [ ] Resource Release
- [ ] Constructor-Based Ownership
- [ ] Destructor-Based Cleanup
- [ ] RAII for Memory
- [ ] RAII for Files
- [ ] RAII for Sockets
- [ ] RAII for Locks

## 9. Smart Pointers
- [ ] `std::unique_ptr`
- [ ] `std::make_unique`
- [ ] `std::shared_ptr`
- [ ] `std::make_shared`
- [ ] `std::weak_ptr`
- [ ] Ownership Semantics
- [ ] Reference Counting
- [ ] Control Block
- [ ] Custom Deleters
- [ ] Smart Pointers with Arrays
- [ ] Smart Pointers with Polymorphism

## 10. Memory Bugs
- [ ] Memory Leak
- [ ] Dangling Pointer
- [ ] Dangling Reference
- [ ] Use-After-Free
- [ ] Double Free
- [ ] Double Delete
- [ ] Buffer Overflow
- [ ] Buffer Underflow
- [ ] Out-of-Bounds Access
- [ ] Uninitialized Memory
- [ ] Invalid Memory Access
- [ ] Null Pointer Dereference
- [ ] Memory Corruption

## 11. Object Lifetime
- [ ] Object Lifetime
- [ ] Storage Duration
- [ ] Automatic Storage Duration
- [ ] Static Storage Duration
- [ ] Thread Storage Duration
- [ ] Dynamic Storage Duration
- [ ] Temporary Objects
- [ ] Lifetime Extension
- [ ] Destruction Order

## 12. Copy & Move Memory Management
- [ ] Shallow Copy
- [ ] Deep Copy
- [ ] Copy Constructor
- [ ] Copy Assignment
- [ ] Move Constructor
- [ ] Move Assignment
- [ ] Rvalue References
- [ ] `std::move`
- [ ] Copy Elision
- [ ] RVO
- [ ] NRVO
- [ ] Rule of 3
- [ ] Rule of 5
- [ ] Rule of 0

## 13. Alignment & Object Representation
- [ ] Object Representation
- [ ] `sizeof`
- [ ] Alignment
- [ ] `alignof`
- [ ] Padding
- [ ] Struct Padding
- [ ] Alignment Requirements
- [ ] `alignas`
- [ ] `std::byte`
- [ ] Object Lifetime & Storage

## 14. Memory Layout
- [ ] Struct Memory Layout
- [ ] Class Memory Layout
- [ ] Base Class Layout
- [ ] Virtual Table
- [ ] Virtual Pointer
- [ ] Multiple Inheritance Layout
- [ ] Empty Base Optimization
- [ ] Data Locality

## 15. Allocators
- [ ] Memory Allocators
- [ ] `std::allocator`
- [ ] Custom Allocators
- [ ] Allocation Strategies
- [ ] Pool Allocation
- [ ] Arena Allocation
- [ ] Object Pools
- [ ] Placement `new`

## 16. Low-Level Memory
- [ ] Raw Memory
- [ ] Raw Storage
- [ ] Placement `new`
- [ ] Explicit Destructor Calls
- [ ] `memcpy`
- [ ] `memmove`
- [ ] `memset`
- [ ] `memcmp`
- [ ] `std::copy`
- [ ] `std::move`
- [ ] Strict Aliasing
- [ ] Type Punning
- [ ] Object Lifetime Rules

## 17. Virtual Memory
- [ ] Virtual Address
- [ ] Physical Address
- [ ] Page
- [ ] Page Table
- [ ] Page Fault
- [ ] TLB
- [ ] Memory Mapping
- [ ] Address Space
- [ ] Copy-on-Write
- [ ] Memory Protection

## 18. Memory-Mapped I/O
- [ ] Memory-Mapped Files
- [ ] `mmap`
- [ ] `munmap`
- [ ] Shared Memory
- [ ] File-Backed Memory
- [ ] Anonymous Mapping

## 19. Memory Debugging
- [ ] GDB Memory Inspection
- [ ] AddressSanitizer
- [ ] LeakSanitizer
- [ ] UndefinedBehaviorSanitizer
- [ ] Valgrind
- [ ] Heap Profiling
- [ ] Memory Leak Detection
- [ ] Core Dumps

## 20. Memory Performance
- [ ] Cache Locality
- [ ] Spatial Locality
- [ ] Temporal Locality
- [ ] Cache Lines
- [ ] False Sharing
- [ ] Allocation Overhead
- [ ] Fragmentation
- [ ] Memory Pooling
- [ ] Object Reuse
- [ ] Zero-Copy
- [ ] Memory Bandwidth

## 21. Systems Programming Memory
- [ ] Process Address Space
- [ ] Heap Management
- [ ] Stack Management
- [ ] Memory-Mapped Files
- [ ] Shared Memory
- [ ] IPC Memory
- [ ] Kernel vs User Memory
- [ ] File Descriptor Related Memory
- [ ] Buffer Management
- [ ] Ring Buffers

## 22. Redis-Relevant Memory Topics 🔥
- [ ] In-Memory Data Structures
- [ ] Hash Tables
- [ ] Dynamic Strings
- [ ] Memory Ownership
- [ ] Custom Allocators
- [ ] Memory Pooling
- [ ] Object Reuse
- [ ] Reference Counting
- [ ] Copy-on-Write
- [ ] Memory Fragmentation
- [ ] Memory Limits
- [ ] Eviction Strategies
- [ ] LRU Cache
- [ ] LFU Cache
- [ ] TTL-Based Memory Management
- [ ] Memory Profiling
- [ ] Zero-Copy Techniques
- [ ] Efficient Serialization
- [ ] Buffer Management
- [ ] Memory-Aware Data Structures