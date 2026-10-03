# C++ Concurrency — Topic List

## 1. Concurrency Fundamentals
- [ ] Concurrency vs Parallelism
- [ ] Process vs Thread
- [ ] Single-Core vs Multi-Core
- [ ] CPU Scheduling Basics
- [ ] Context Switching
- [ ] Thread Lifecycle
- [ ] Race Conditions
- [ ] Data Races
- [ ] Critical Sections
- [ ] Shared State
- [ ] Thread Safety
- [ ] Thread Affinity

## 2. `std::thread`
- [ ] Creating Threads
- [ ] Starting Threads
- [ ] Passing Arguments
- [ ] Lambda Threads
- [ ] Function Threads
- [ ] Member Function Threads
- [ ] `join()`
- [ ] `detach()`
- [ ] `joinable()`
- [ ] Thread Lifetime
- [ ] Thread IDs
- [ ] `std::this_thread`
- [ ] `sleep_for()`
- [ ] `sleep_until()`
- [ ] `yield()`

## 3. Thread Management
- [ ] Joinable Threads
- [ ] Detached Threads
- [ ] Thread Ownership
- [ ] Thread Lifetime Management
- [ ] RAII Thread Management
- [ ] `std::jthread`
- [ ] Cooperative Cancellation
- [ ] `std::stop_token`

## 4. Race Conditions
- [ ] Race Condition
- [ ] Data Race
- [ ] Read-Read Access
- [ ] Read-Write Access
- [ ] Write-Write Access
- [ ] Critical Section
- [ ] Race Condition Detection
- [ ] Race Prevention

## 5. Mutexes
- [ ] `std::mutex`
- [ ] `lock()`
- [ ] `unlock()`
- [ ] `try_lock()`
- [ ] `std::lock_guard`
- [ ] `std::unique_lock`
- [ ] `std::scoped_lock`
- [ ] Mutex Ownership
- [ ] Mutex Contention
- [ ] Recursive Mutex
- [ ] `std::recursive_mutex`
- [ ] Timed Mutex
- [ ] `std::timed_mutex`
- [ ] `std::recursive_timed_mutex`

## 6. Locking Strategies
- [ ] RAII Locking
- [ ] Lock Granularity
- [ ] Fine-Grained Locking
- [ ] Coarse-Grained Locking
- [ ] Lock Ordering
- [ ] Multiple Mutex Locking
- [ ] `std::lock`
- [ ] `std::try_lock`
- [ ] `std::scoped_lock`
- [ ] Avoiding Lock Contention

## 7. Deadlocks
- [ ] Deadlock
- [ ] Deadlock Conditions
- [ ] Circular Wait
- [ ] Lock Ordering
- [ ] Deadlock Prevention
- [ ] Deadlock Avoidance
- [ ] Deadlock Detection
- [ ] Lock Hierarchies

## 8. Condition Variables
- [ ] `std::condition_variable`
- [ ] `wait()`
- [ ] `notify_one()`
- [ ] `notify_all()`
- [ ] Predicate-Based Waiting
- [ ] Spurious Wakeups
- [ ] Producer-Consumer Pattern
- [ ] Condition Variable + Mutex

## 9. Atomic Operations
- [ ] `std::atomic`
- [ ] Atomic Variables
- [ ] Atomic Load
- [ ] Atomic Store
- [ ] Atomic Exchange
- [ ] Compare-And-Swap
- [ ] `compare_exchange_weak()`
- [ ] `compare_exchange_strong()`
- [ ] Atomic Counters
- [ ] Atomic Flags
- [ ] Lock-Free Programming
- [ ] Wait/Notify on Atomics

## 10. Memory Ordering 🔥
- [ ] Sequential Consistency
- [ ] Relaxed Ordering
- [ ] Acquire Ordering
- [ ] Release Ordering
- [ ] Acquire-Release
- [ ] `memory_order_relaxed`
- [ ] `memory_order_acquire`
- [ ] `memory_order_release`
- [ ] `memory_order_acq_rel`
- [ ] `memory_order_seq_cst`
- [ ] Memory Barriers
- [ ] Happens-Before Relationship

## 11. C++ Memory Model
- [ ] C++ Memory Model
- [ ] Thread Visibility
- [ ] Synchronization
- [ ] Happens-Before
- [ ] Synchronizes-With
- [ ] Modification Order
- [ ] Atomicity
- [ ] Ordering
- [ ] Visibility

## 12. Thread-Safe Data Structures
- [ ] Thread-Safe Counter
- [ ] Thread-Safe Queue
- [ ] Thread-Safe Stack
- [ ] Thread-Safe Map
- [ ] Concurrent Queue
- [ ] Bounded Queue
- [ ] Blocking Queue
- [ ] Lock-Based Data Structures
- [ ] Lock-Free Data Structures

## 13. Futures & Promises
- [ ] `std::future`
- [ ] `std::promise`
- [ ] `std::async`
- [ ] `std::packaged_task`
- [ ] `get()`
- [ ] `wait()`
- [ ] `wait_for()`
- [ ] `wait_until()`
- [ ] Exception Propagation
- [ ] Asynchronous Results

## 14. `std::async`
- [ ] Async Tasks
- [ ] `std::launch::async`
- [ ] `std::launch::deferred`
- [ ] Async Execution
- [ ] Deferred Execution
- [ ] Future Management

## 15. Thread Local Storage
- [ ] `thread_local`
- [ ] Thread-Local Variables
- [ ] Thread-Local Objects
- [ ] Thread-Local Lifetime
- [ ] Thread-Local Storage Use Cases

## 16. Thread Pools
- [ ] Thread Pool Architecture
- [ ] Worker Threads
- [ ] Task Queue
- [ ] Task Submission
- [ ] Worker Loop
- [ ] Condition Variables
- [ ] Graceful Shutdown
- [ ] Dynamic Worker Management
- [ ] Fixed Thread Pool
- [ ] Work Distribution

## 17. Producer-Consumer
- [ ] Producer-Consumer Pattern
- [ ] Blocking Queue
- [ ] Condition Variables
- [ ] Bounded Buffer
- [ ] Backpressure
- [ ] Multiple Producers
- [ ] Multiple Consumers

## 18. Parallelism
- [ ] Data Parallelism
- [ ] Task Parallelism
- [ ] Work Distribution
- [ ] Load Balancing
- [ ] Parallel Algorithms
- [ ] `std::execution`
- [ ] Parallel `for_each`
- [ ] Parallel `sort`

## 19. C++ Parallel Algorithms
- [ ] Execution Policies
- [ ] `std::execution::seq`
- [ ] `std::execution::par`
- [ ] `std::execution::par_unseq`
- [ ] Parallel Algorithms
- [ ] Thread Safety of Algorithms

## 20. Synchronization Primitives
- [ ] Mutex
- [ ] Condition Variable
- [ ] Semaphore
- [ ] Latch
- [ ] Barrier
- [ ] Atomic
- [ ] Spinlock

## 21. Semaphores
- [ ] `std::counting_semaphore`
- [ ] `std::binary_semaphore`
- [ ] `acquire()`
- [ ] `release()`
- [ ] `try_acquire()`
- [ ] Semaphore-Based Synchronization

## 22. Latches & Barriers
- [ ] `std::latch`
- [ ] `count_down()`
- [ ] `wait()`
- [ ] `arrive_and_wait()`
- [ ] `std::barrier`
- [ ] Reusable Synchronization Points
- [ ] Phase-Based Synchronization

## 23. Spinlocks
- [ ] Spinlock Concept
- [ ] Atomic-Based Spinlock
- [ ] Busy Waiting
- [ ] Spinlock vs Mutex
- [ ] CPU Usage
- [ ] Lock Contention

## 24. Lock-Free Programming
- [ ] Lock-Free vs Wait-Free
- [ ] CAS
- [ ] Compare-and-Swap
- [ ] ABA Problem
- [ ] Lock-Free Stack
- [ ] Lock-Free Queue
- [ ] Memory Reclamation
- [ ] Hazard Pointers
- [ ] Epoch-Based Reclamation

## 25. False Sharing & Cache
- [ ] CPU Cache
- [ ] Cache Lines
- [ ] Cache Coherency
- [ ] False Sharing
- [ ] Cache Locality
- [ ] Data Alignment
- [ ] `std::hardware_destructive_interference_size`
- [ ] Padding for Concurrency

## 26. Advanced Synchronization
- [ ] Read-Write Locks
- [ ] `std::shared_mutex`
- [ ] `std::shared_lock`
- [ ] Reader-Writer Problem
- [ ] Upgrade/Downgrade Locking
- [ ] Lock-Free Synchronization
- [ ] Wait-Free Algorithms

## 27. Concurrency Patterns
- [ ] Producer-Consumer
- [ ] Reader-Writer
- [ ] Thread Pool
- [ ] Work Stealing
- [ ] Fork-Join
- [ ] Pipeline
- [ ] Actor Model
- [ ] Active Object
- [ ] Future/Promise
- [ ] Publish-Subscribe

## 28. Cancellation & Shutdown
- [ ] Cooperative Cancellation
- [ ] `std::stop_token`
- [ ] `std::stop_source`
- [ ] `std::stop_callback`
- [ ] Graceful Thread Shutdown
- [ ] Thread Pool Shutdown
- [ ] Resource Cleanup
- [ ] Cancellation Safety

## 29. Concurrency & Exceptions
- [ ] Exceptions in Threads
- [ ] `std::exception_ptr`
- [ ] `std::current_exception`
- [ ] `std::rethrow_exception`
- [ ] Exception Propagation
- [ ] Future-Based Exception Handling
- [ ] Thread Failure Handling

## 30. Concurrency Debugging
- [ ] Thread Debugging with GDB
- [ ] Thread Sanitizer
- [ ] Address Sanitizer
- [ ] Race Detection
- [ ] Deadlock Detection
- [ ] Thread Dumps
- [ ] Core Dumps
- [ ] Logging Concurrent Systems

## 31. Concurrency Performance
- [ ] Thread Creation Cost
- [ ] Context Switching
- [ ] Lock Contention
- [ ] Synchronization Overhead
- [ ] Atomic Performance
- [ ] Cache Locality
- [ ] False Sharing
- [ ] Scalability
- [ ] Throughput
- [ ] Latency
- [ ] Amdahl's Law
- [ ] Benchmarking

## 32. Linux Concurrency
- [ ] POSIX Threads
- [ ] `pthread_create`
- [ ] `pthread_join`
- [ ] POSIX Mutex
- [ ] POSIX Condition Variables
- [ ] Linux Scheduler Basics
- [ ] CPU Affinity
- [ ] Thread Priorities
- [ ] Signals & Threads

## 33. Networking Concurrency 🔥
- [ ] Concurrent TCP Server
- [ ] Multi-Threaded Server
- [ ] Thread-Per-Connection
- [ ] Worker Thread Model
- [ ] Thread Pool Server
- [ ] Non-Blocking Sockets
- [ ] Event Loop
- [ ] Reactor Pattern
- [ ] Connection Management
- [ ] Graceful Client Disconnect

## 34. Redis-Relevant Concurrency 🔥
- [ ] Concurrent Client Handling
- [ ] Shared Data Protection
- [ ] Command Execution
- [ ] Thread Pools
- [ ] Background Tasks
- [ ] Connection Management
- [ ] Event Loop
- [ ] Producer-Consumer Queue
- [ ] Atomic Counters
- [ ] Lock Contention
- [ ] Cache-Line Effects
- [ ] Graceful Shutdown
- [ ] Background Persistence
- [ ] Concurrent Logging

## 35. Concurrency Projects
- [ ] Thread-Safe Counter
- [ ] Producer-Consumer Queue
- [ ] Thread-Safe Queue
- [ ] Thread Pool
- [ ] Parallel File Processor
- [ ] Concurrent Logger
- [ ] Multi-Threaded Web Crawler
- [ ] Concurrent TCP Server
- [ ] Thread-Pool TCP Server
- [ ] Lock-Free Queue
- [ ] Concurrent Cache
- [ ] Redis-Style Command Server