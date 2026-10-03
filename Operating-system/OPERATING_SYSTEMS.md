The highest-value area:

- **Processes & threads**
- **Concurrency & synchronization**
- **Memory management**
- **I/O & system calls**
- **File systems**
- **Networking from OS perspective**
- **Scheduling**
- **IPC**
- **Linux fundamentals**
- **OS concepts behind Go runtime**

```markdown
# Operating Systems Roadmap

# 1. OS Fundamentals ⭐⭐⭐⭐⭐

- [ ] What is an Operating System?
- [ ] OS responsibilities
- [ ] Kernel
- [ ] User space
- [ ] Kernel space
- [ ] System programs
- [ ] System calls
- [ ] Hardware abstraction
- [ ] Resource management
- [ ] Process management
- [ ] Memory management
- [ ] File system management
- [ ] I/O management
- [ ] Security and protection
- [ ] Virtualization

---

# 2. OS Architecture ⭐⭐⭐⭐

- [ ] Monolithic kernel
- [ ] Microkernel
- [ ] Hybrid kernel
- [ ] Layered architecture
- [ ] Modular kernel
- [ ] Kernel modules
- [ ] User space vs kernel space
- [ ] Privileged mode
- [ ] User mode
- [ ] Kernel mode
- [ ] Mode switching
- [ ] Interrupts
- [ ] Exceptions
- [ ] Traps

---

# 3. System Calls ⭐⭐⭐⭐⭐

- [ ] What is a system call?
- [ ] System call lifecycle
- [ ] User program → kernel
- [ ] System call interface
- [ ] System call overhead
- [ ] Trap instruction
- [ ] Return from system call
- [ ] open()
- [ ] read()
- [ ] write()
- [ ] close()
- [ ] fork()
- [ ] exec()
- [ ] wait()
- [ ] exit()
- [ ] mmap()
- [ ] socket()
- [ ] bind()
- [ ] listen()
- [ ] accept()
- [ ] connect()

---

# 4. Processes ⭐⭐⭐⭐⭐

- [ ] What is a process?
- [ ] Program vs process
- [ ] Process address space
- [ ] Process Control Block
- [ ] PID
- [ ] PPID
- [ ] Process states
- [ ] New
- [ ] Ready
- [ ] Running
- [ ] Waiting
- [ ] Terminated
- [ ] Process creation
- [ ] Process termination
- [ ] Parent process
- [ ] Child process
- [ ] Orphan process
- [ ] Zombie process

---

# 5. Process Creation

- [ ] fork()
- [ ] exec()
- [ ] wait()
- [ ] waitpid()
- [ ] exit()
- [ ] Copy-on-write
- [ ] Parent/child relationship
- [ ] Process inheritance
- [ ] File descriptor inheritance
- [ ] Environment variables
- [ ] Process groups
- [ ] Sessions

---

# 6. Threads ⭐⭐⭐⭐⭐

- [ ] What is a thread?
- [ ] Process vs thread
- [ ] User-level threads
- [ ] Kernel-level threads
- [ ] Thread Control Block
- [ ] Thread stack
- [ ] Thread-local storage
- [ ] Thread creation
- [ ] Thread termination
- [ ] Thread scheduling
- [ ] Multithreading
- [ ] Thread pools
- [ ] Worker threads

---

# 7. Process vs Thread

Understand deeply:

```text
Process
├── Address Space
├── Code
├── Heap
├── Global Data
├── File Descriptors
└── Threads
    ├── Thread 1
    ├── Thread 2
    └── Thread 3
```

Study:

- [ ] Memory sharing
- [ ] Isolation
- [ ] Creation cost
- [ ] Context switching
- [ ] Communication
- [ ] Synchronization
- [ ] Failure isolation

---

# 8. CPU Scheduling ⭐⭐⭐⭐⭐

- [ ] CPU burst
- [ ] I/O burst
- [ ] Scheduler
- [ ] Dispatcher
- [ ] Context switch
- [ ] Preemptive scheduling
- [ ] Non-preemptive scheduling
- [ ] FCFS
- [ ] SJF
- [ ] SRTF
- [ ] Priority Scheduling
- [ ] Round Robin
- [ ] Multilevel Queue
- [ ] Multilevel Feedback Queue

---

# 9. Scheduling Metrics

- [ ] CPU utilization
- [ ] Throughput
- [ ] Turnaround time
- [ ] Waiting time
- [ ] Response time
- [ ] Completion time
- [ ] Scheduling overhead
- [ ] Context-switch overhead

---

# 10. Context Switching ⭐⭐⭐⭐⭐

- [ ] What is context switching?
- [ ] Process context
- [ ] Thread context
- [ ] CPU registers
- [ ] Program counter
- [ ] Stack pointer
- [ ] Context save
- [ ] Context restore
- [ ] Context-switch overhead
- [ ] Process context switch
- [ ] Thread context switch

Understand why:

```text
Thread A
   ↓
Context Switch
   ↓
Thread B
   ↓
Context Switch
   ↓
Thread A
```

---

# 11. Concurrency ⭐⭐⭐⭐⭐

- [ ] Concurrency
- [ ] Parallelism
- [ ] Sequential execution
- [ ] Concurrent execution
- [ ] Parallel execution
- [ ] Race conditions
- [ ] Critical section
- [ ] Shared state
- [ ] Atomic operations
- [ ] Synchronization
- [ ] Thread safety

---

# 12. Critical Section ⭐⭐⭐⭐⭐

- [ ] Critical section
- [ ] Entry section
- [ ] Critical section
- [ ] Exit section
- [ ] Remainder section
- [ ] Mutual exclusion
- [ ] Progress
- [ ] Bounded waiting

---

# 13. Synchronization ⭐⭐⭐⭐⭐

- [ ] Mutex
- [ ] Spinlock
- [ ] Semaphore
- [ ] Binary semaphore
- [ ] Counting semaphore
- [ ] Condition variable
- [ ] Monitor
- [ ] Read-write lock
- [ ] Barrier
- [ ] Atomic operations

---

# 14. Race Conditions ⭐⭐⭐⭐⭐

- [ ] Race condition
- [ ] Data race
- [ ] Read-modify-write
- [ ] Atomicity
- [ ] Lost update
- [ ] Check-then-act
- [ ] Thread-safe code
- [ ] Lock-free programming

---

# 15. Deadlocks ⭐⭐⭐⭐⭐

- [ ] What is deadlock?
- [ ] Mutual exclusion
- [ ] Hold and wait
- [ ] No preemption
- [ ] Circular wait
- [ ] Deadlock prevention
- [ ] Deadlock avoidance
- [ ] Deadlock detection
- [ ] Deadlock recovery
- [ ] Resource Allocation Graph
- [ ] Banker's Algorithm

---

# 16. Livelock and Starvation

- [ ] Starvation
- [ ] Livelock
- [ ] Deadlock vs starvation
- [ ] Deadlock vs livelock
- [ ] Priority inversion
- [ ] Priority inheritance
- [ ] Fairness

---

# 17. IPC — Inter Process Communication ⭐⭐⭐⭐⭐

- [ ] IPC
- [ ] Pipes
- [ ] Named pipes
- [ ] Message queues
- [ ] Shared memory
- [ ] Semaphores
- [ ] Signals
- [ ] Unix domain sockets
- [ ] Sockets

---

# 18. Pipes

- [ ] Anonymous pipe
- [ ] Named pipe
- [ ] Pipe buffer
- [ ] Producer-consumer
- [ ] Blocking read
- [ ] Blocking write
- [ ] Pipe lifecycle

---

# 19. Shared Memory

- [ ] Shared memory
- [ ] Memory mapping
- [ ] mmap()
- [ ] Shared memory synchronization
- [ ] Shared memory vs message passing
- [ ] Shared memory performance

---

# 20. Signals

- [ ] Signal
- [ ] SIGINT
- [ ] SIGTERM
- [ ] SIGKILL
- [ ] SIGHUP
- [ ] SIGCHLD
- [ ] Signal handler
- [ ] Signal masking
- [ ] Signal delivery
- [ ] Graceful shutdown

---

# 21. Memory Management ⭐⭐⭐⭐⭐

- [ ] Physical memory
- [ ] Virtual memory
- [ ] Address space
- [ ] Logical address
- [ ] Physical address
- [ ] Memory allocation
- [ ] Memory deallocation
- [ ] Memory fragmentation
- [ ] Internal fragmentation
- [ ] External fragmentation

---

# 22. Process Address Space ⭐⭐⭐⭐⭐

Understand:

```text
High Address
+----------------+
| Stack          |
+----------------+
|                |
| Memory Mapping |
|                |
+----------------+
| Heap           |
+----------------+
| BSS            |
+----------------+
| Data           |
+----------------+
| Text / Code    |
+----------------+
Low Address
```

Study:

- [ ] Text segment
- [ ] Data segment
- [ ] BSS
- [ ] Heap
- [ ] Stack
- [ ] Memory mapped region
- [ ] Shared libraries

---

# 23. Virtual Memory ⭐⭐⭐⭐⭐

- [ ] Virtual memory
- [ ] Virtual address
- [ ] Physical address
- [ ] Address translation
- [ ] Page
- [ ] Frame
- [ ] Page table
- [ ] Multi-level page tables
- [ ] Page Table Entry
- [ ] TLB
- [ ] TLB hit
- [ ] TLB miss
- [ ] Page fault

---

# 24. Paging

- [ ] Paging
- [ ] Page size
- [ ] Page number
- [ ] Page offset
- [ ] Frame number
- [ ] Page table
- [ ] Multi-level paging
- [ ] Inverted page table
- [ ] Page replacement

---

# 25. Page Replacement

- [ ] FIFO
- [ ] Optimal
- [ ] LRU
- [ ] Clock
- [ ] Second Chance
- [ ] Page fault
- [ ] Page fault rate
- [ ] Thrashing

---

# 26. Memory Allocation

- [ ] malloc()
- [ ] calloc()
- [ ] realloc()
- [ ] free()
- [ ] Heap allocation
- [ ] Stack allocation
- [ ] Memory allocator
- [ ] Fragmentation
- [ ] Memory leak
- [ ] Double free
- [ ] Use-after-free

---

# 27. mmap ⭐⭐⭐⭐⭐

- [ ] mmap()
- [ ] munmap()
- [ ] File-backed memory mapping
- [ ] Anonymous mapping
- [ ] Shared mapping
- [ ] Private mapping
- [ ] Memory mapped files
- [ ] mmap vs read/write
- [ ] mmap for IPC

---

# 28. Copy-on-Write

- [ ] Copy-on-write
- [ ] fork() + COW
- [ ] Shared pages
- [ ] Page fault during COW
- [ ] COW performance

---

# 29. Memory-Mapped Files

Important for systems projects.

- [ ] File mapping
- [ ] mmap
- [ ] Random access
- [ ] Page cache
- [ ] Dirty pages
- [ ] msync
- [ ] Persistent data
- [ ] mmap vs normal file I/O

---

# 30. Garbage Collection ⭐⭐⭐⭐⭐

- [ ] Garbage collection
- [ ] Mark and sweep
- [ ] Generational GC
- [ ] Concurrent GC
- [ ] Stop-the-world
- [ ] GC roots
- [ ] Heap allocation
- [ ] GC pressure
- [ ] Memory retention
- [ ] GC latency

---

# 31. Go Runtime Memory

Understand:

- [ ] Go heap
- [ ] Go stack
- [ ] Goroutine stack
- [ ] Stack growth
- [ ] Escape analysis
- [ ] Heap allocation
- [ ] Garbage collector
- [ ] Garbage collection pauses
- [ ] Memory allocation
- [ ] Memory reuse
- [ ] Runtime scheduler

---

# 32. Go Scheduler ⭐⭐⭐⭐⭐

Understand:

- [ ] Goroutines
- [ ] OS threads
- [ ] G
- [ ] M
- [ ] P
- [ ] G-M-P model
- [ ] Goroutine scheduling
- [ ] Work stealing
- [ ] Preemption
- [ ] Scheduler queues
- [ ] Blocking system calls
- [ ] Network poller

Concept:

```text
Goroutines
    |
    v
    P
    |
    v
   M
    |
    v
 CPU
```

---

# 33. Goroutines vs OS Threads ⭐⭐⭐⭐⭐

- [ ] Goroutine
- [ ] OS thread
- [ ] Goroutine stack
- [ ] Thread stack
- [ ] Scheduling cost
- [ ] Context switching
- [ ] Blocking behavior
- [ ] Thread creation
- [ ] Goroutine creation
- [ ] M:N scheduling

---

# 34. Go Network Poller ⭐⭐⭐⭐⭐

Understand conceptually:

- [ ] Blocking I/O
- [ ] Non-blocking I/O
- [ ] Network polling
- [ ] epoll
- [ ] kqueue
- [ ] IOCP
- [ ] Go netpoller
- [ ] Goroutine parking
- [ ] Goroutine wakeup
- [ ] File descriptor readiness

---

# 35. I/O ⭐⭐⭐⭐⭐

- [ ] Input/output
- [ ] Blocking I/O
- [ ] Non-blocking I/O
- [ ] Synchronous I/O
- [ ] Asynchronous I/O
- [ ] Buffered I/O
- [ ] Unbuffered I/O
- [ ] Direct I/O
- [ ] Sequential I/O
- [ ] Random I/O

---

# 36. File Descriptors ⭐⭐⭐⭐⭐

- [ ] File descriptor
- [ ] stdin
- [ ] stdout
- [ ] stderr
- [ ] File descriptor table
- [ ] open()
- [ ] read()
- [ ] write()
- [ ] close()
- [ ] dup()
- [ ] dup2()
- [ ] File descriptor inheritance
- [ ] Socket as file descriptor
- [ ] Pipe as file descriptor

This is extremely important for backend systems.

---

# 37. File Systems ⭐⭐⭐⭐⭐

- [ ] File
- [ ] Directory
- [ ] Path
- [ ] File descriptor
- [ ] Inode
- [ ] File metadata
- [ ] Permissions
- [ ] Links
- [ ] Hard link
- [ ] Symbolic link
- [ ] Mount
- [ ] File system hierarchy

---

# 38. Linux File System

- [ ] /
- [ ] /bin
- [ ] /etc
- [ ] /home
- [ ] /tmp
- [ ] /var
- [ ] /dev
- [ ] /proc
- [ ] /sys
- [ ] /usr
- [ ] /opt

---

# 39. File System Internals

- [ ] Inode
- [ ] Directory entry
- [ ] File metadata
- [ ] Data blocks
- [ ] Block allocation
- [ ] Free space management
- [ ] Journaling
- [ ] File system cache
- [ ] Page cache
- [ ] Write buffering

---

# 40. File I/O

- [ ] open
- [ ] read
- [ ] write
- [ ] close
- [ ] lseek
- [ ] fsync
- [ ] fdatasync
- [ ] stat
- [ ] fstat
- [ ] rename
- [ ] unlink
- [ ] mkdir
- [ ] rmdir

---

# 41. Page Cache ⭐⭐⭐⭐⭐

- [ ] Page cache
- [ ] File cache
- [ ] Read cache
- [ ] Write cache
- [ ] Dirty pages
- [ ] Cache eviction
- [ ] fsync
- [ ] Buffered write
- [ ] Direct I/O

Important for Event Streaming Engine.

Understand:

```text
Application
     |
     v
write()
     |
     v
Page Cache
     |
     v
Disk
```

---

# 42. Storage

- [ ] HDD
- [ ] SSD
- [ ] NVMe
- [ ] Disk latency
- [ ] IOPS
- [ ] Sequential access
- [ ] Random access
- [ ] Storage throughput
- [ ] Write amplification
- [ ] Durability

---

# 43. Disk Scheduling

- [ ] FCFS
- [ ] SSTF
- [ ] SCAN
- [ ] C-SCAN
- [ ] LOOK
- [ ] C-LOOK

Know the concepts for interviews.

---

# 44. Linux Processes ⭐⭐⭐⭐⭐

Learn practically:

- [ ] ps
- [ ] top
- [ ] htop
- [ ] pid
- [ ] ppid
- [ ] kill
- [ ] kill -9
- [ ] jobs
- [ ] bg
- [ ] fg
- [ ] nohup
- [ ] nice
- [ ] renice

---

# 45. Linux Threads

- [ ] Thread identification
- [ ] ps -T
- [ ] top thread view
- [ ] Thread CPU usage
- [ ] Thread states
- [ ] Thread debugging

---

# 46. Linux Memory Debugging

- [ ] free
- [ ] vmstat
- [ ] top
- [ ] htop
- [ ] /proc/meminfo
- [ ] /proc/[pid]/status
- [ ] RSS
- [ ] VSZ
- [ ] Virtual memory
- [ ] Resident memory
- [ ] Memory pressure

---

# 47. Linux Networking from OS Perspective

- [ ] Socket
- [ ] File descriptor
- [ ] TCP socket
- [ ] UDP socket
- [ ] listen
- [ ] accept
- [ ] connect
- [ ] send
- [ ] recv
- [ ] socket buffer
- [ ] TCP state
- [ ] epoll
- [ ] network namespace

---

# 48. epoll ⭐⭐⭐⭐⭐

Very important for backend/system programming.

- [ ] epoll
- [ ] epoll_create
- [ ] epoll_ctl
- [ ] epoll_wait
- [ ] File descriptor readiness
- [ ] Level-triggered
- [ ] Edge-triggered
- [ ] Non-blocking sockets
- [ ] Event loop
- [ ] epoll vs select
- [ ] epoll vs poll

Understand:

```text
Many Connections
      |
      v
    epoll
      |
      v
Ready FDs
      |
      v
 Application
```

---

# 49. select / poll / epoll

- [ ] select
- [ ] poll
- [ ] epoll
- [ ] File descriptor monitoring
- [ ] Blocking vs non-blocking
- [ ] Scalability
- [ ] Event-driven I/O

---

# 50. Event-Driven Architecture ⭐⭐⭐⭐⭐

- [ ] Event loop
- [ ] Event queue
- [ ] I/O readiness
- [ ] Non-blocking I/O
- [ ] Event handler
- [ ] Reactor pattern
- [ ] Proactor pattern
- [ ] Event-driven server

---

# 51. Reactor Pattern

Understand:

```text
Connections
     |
     v
Event Loop
     |
     v
Ready Event
     |
     v
Handler
     |
     v
Process
```

Learn:

- [ ] Reactor
- [ ] Event demultiplexer
- [ ] Event loop
- [ ] Handlers
- [ ] Non-blocking I/O

---

# 52. Backend Process Architecture

Understand:

```text
Load Balancer
      |
      v
+-------------+
| Go Server 1 |
+-------------+
| Go Server 2 |
+-------------+
| Go Server 3 |
+-------------+
      |
      v
 Database
```

Study:

- [ ] Process model
- [ ] Thread model
- [ ] Worker model
- [ ] Connection model
- [ ] Process isolation
- [ ] Graceful shutdown
- [ ] Restart
- [ ] Health checks

---

# 53. Job Queue — OS Concepts

Understand:

```text
Producer
    |
    v
Job Queue
    |
    +---- Worker 1
    |
    +---- Worker 2
    |
    +---- Worker 3
```

Study:

- [ ] Processes
- [ ] Threads
- [ ] Goroutines
- [ ] Worker pools
- [ ] Blocking
- [ ] Synchronization
- [ ] Mutex
- [ ] Channels
- [ ] Semaphores
- [ ] Backpressure
- [ ] CPU scheduling
- [ ] I/O waiting
- [ ] Graceful shutdown
- [ ] Signals
- [ ] Retry
- [ ] Deadlock prevention

---

# 54. Event Streaming Engine — OS Concepts ⭐⭐⭐⭐⭐

Study deeply:

- [ ] File descriptors
- [ ] TCP sockets
- [ ] Non-blocking I/O
- [ ] epoll
- [ ] Event loops
- [ ] Goroutines
- [ ] Goroutine scheduling
- [ ] Memory buffers
- [ ] Ring buffers
- [ ] File I/O
- [ ] Page cache
- [ ] mmap
- [ ] fsync
- [ ] Sequential writes
- [ ] Batch writes
- [ ] Disk durability
- [ ] Concurrent consumers
- [ ] Backpressure
- [ ] CPU vs I/O bottlenecks

---

# 55. Event Streaming Storage

Understand:

```text
Producer
   |
   v
Memory Buffer
   |
   v
Batch
   |
   v
Page Cache
   |
   v
Disk
```

Study:

- [ ] Append-only log
- [ ] Sequential writes
- [ ] File segments
- [ ] File rotation
- [ ] Offset
- [ ] Index
- [ ] mmap
- [ ] Page cache
- [ ] fsync
- [ ] Crash recovery
- [ ] Log recovery
- [ ] Data durability

---

# 56. Memory and Performance

- [ ] CPU cache
- [ ] Cache line
- [ ] Locality
- [ ] Spatial locality
- [ ] Temporal locality
- [ ] False sharing
- [ ] Memory alignment
- [ ] CPU vs memory bottleneck
- [ ] Memory bandwidth
- [ ] Cache miss
- [ ] Context switching overhead

---

# 57. CPU Cache

- [ ] L1 cache
- [ ] L2 cache
- [ ] L3 cache
- [ ] Cache line
- [ ] Cache hit
- [ ] Cache miss
- [ ] Locality
- [ ] False sharing
- [ ] Cache coherence

---

# 58. CPU vs I/O Bound

- [ ] CPU-bound workload
- [ ] I/O-bound workload
- [ ] CPU utilization
- [ ] I/O wait
- [ ] Throughput
- [ ] Latency
- [ ] Worker count
- [ ] Concurrency tuning

---

# 59. Performance Profiling

- [ ] CPU profiling
- [ ] Memory profiling
- [ ] I/O profiling
- [ ] Goroutine profiling
- [ ] Lock contention
- [ ] System call tracing

Tools:

- [ ] top
- [ ] htop
- [ ] vmstat
- [ ] iostat
- [ ] pidstat
- [ ] strace
- [ ] perf
- [ ] lsof
- [ ] pprof

---

# 60. strace

Learn:

- [ ] Trace system calls
- [ ] open
- [ ] read
- [ ] write
- [ ] close
- [ ] socket
- [ ] connect
- [ ] accept
- [ ] mmap
- [ ] futex
- [ ] clone
- [ ] execve

Use strace to understand:

```text
Go Application
      |
      v
System Calls
      |
      v
Linux Kernel
```

---

# 61. Linux Permissions

- [ ] Users
- [ ] Groups
- [ ] File permissions
- [ ] Read
- [ ] Write
- [ ] Execute
- [ ] chmod
- [ ] chown
- [ ] umask
- [ ] Process permissions
- [ ] Root user

---

# 62. Process Isolation

- [ ] Process isolation
- [ ] Namespaces
- [ ] PID namespace
- [ ] Network namespace
- [ ] Mount namespace
- [ ] User namespace
- [ ] Control groups
- [ ] cgroups
- [ ] Resource limits

---

# 63. Containers ⭐⭐⭐⭐⭐

- [ ] What is a container?
- [ ] Containers vs VMs
- [ ] Linux namespaces
- [ ] cgroups
- [ ] Container networking
- [ ] Container filesystem
- [ ] Container processes
- [ ] Resource limits
- [ ] Docker fundamentals

---

# 64. Graceful Shutdown ⭐⭐⭐⭐⭐

Understand:

```text
SIGTERM
   |
   v
Stop accepting requests
   |
   v
Finish active requests
   |
   v
Flush buffers
   |
   v
Commit data
   |
   v
Close connections
   |
   v
Exit
```

Study:

- [ ] SIGTERM
- [ ] SIGINT
- [ ] Signal handling
- [ ] Shutdown context
- [ ] Connection draining
- [ ] Worker shutdown
- [ ] Database shutdown
- [ ] Queue shutdown
- [ ] Buffer flushing

---

# 65. Fault Tolerance

- [ ] Process crash
- [ ] Thread failure
- [ ] Network failure
- [ ] Disk failure
- [ ] Database failure
- [ ] Timeout
- [ ] Retry
- [ ] Recovery
- [ ] Checkpointing
- [ ] Durable storage
- [ ] Replication
- [ ] Health checks

---

# 66. OS Interview Questions

- [ ] Process vs thread
- [ ] Process vs program
- [ ] User mode vs kernel mode
- [ ] What is a system call?
- [ ] What happens during a system call?
- [ ] What is context switching?
- [ ] Why is context switching expensive?
- [ ] What is a race condition?
- [ ] What is a critical section?
- [ ] Mutex vs semaphore
- [ ] What is deadlock?
- [ ] Four conditions of deadlock
- [ ] Deadlock vs starvation
- [ ] Process vs thread memory
- [ ] What is virtual memory?
- [ ] What is paging?
- [ ] What is a page fault?
- [ ] What is TLB?
- [ ] What is a page table?
- [ ] What is copy-on-write?
- [ ] What is mmap?
- [ ] What is a file descriptor?
- [ ] What is an inode?
- [ ] What is page cache?
- [ ] What is fork?
- [ ] What is exec?
- [ ] What is IPC?
- [ ] Pipe vs shared memory
- [ ] Blocking vs non-blocking I/O
- [ ] select vs poll vs epoll
- [ ] What is an event loop?
- [ ] What is a socket?
- [ ] How does a server handle multiple connections?
- [ ] What is a zombie process?
- [ ] What is an orphan process?
- [ ] What is a memory leak?
- [ ] Stack vs heap
- [ ] What is garbage collection?
- [ ] CPU-bound vs I/O-bound
- [ ] What happens when a process crashes?
- [ ] What happens when a server receives SIGTERM?

---

# 67. Practical OS Projects

## Project 1 — Process Monitor

- [ ] Process listing
- [ ] PID
- [ ] CPU usage
- [ ] Memory usage
- [ ] Process state
- [ ] Linux /proc

## Project 2 — Mini Shell

- [ ] fork()
- [ ] exec()
- [ ] wait()
- [ ] Pipes
- [ ] Redirection
- [ ] File descriptors
- [ ] Signals

## Project 3 — Thread Pool

- [ ] Worker threads
- [ ] Task queue
- [ ] Mutex
- [ ] Condition variables
- [ ] Shutdown
- [ ] Work stealing

## Project 4 — TCP Server

- [ ] Socket
- [ ] bind
- [ ] listen
- [ ] accept
- [ ] read
- [ ] write
- [ ] epoll
- [ ] Non-blocking I/O

## Project 5 — Job Queue

- [ ] Worker pool
- [ ] Goroutines
- [ ] Channels
- [ ] Synchronization
- [ ] Backpressure
- [ ] Graceful shutdown
- [ ] Retry
- [ ] Timeouts

## Project 6 — Event Streaming Engine

- [ ] TCP server
- [ ] Event loop
- [ ] Goroutines
- [ ] Message buffers
- [ ] Ring buffer
- [ ] Append-only log
- [ ] File segments
- [ ] mmap
- [ ] Page cache
- [ ] fsync
- [ ] Consumer offsets
- [ ] Concurrent consumers
- [ ] Backpressure
- [ ] Crash recovery

---

# 68. Recommended Learning Order

## Phase 1 — Core OS

1. [ ] OS fundamentals
2. [ ] Kernel
3. [ ] User space / kernel space
4. [ ] System calls
5. [ ] Processes
6. [ ] Threads
7. [ ] Scheduling
8. [ ] Context switching

## Phase 2 — Concurrency

9. [ ] Concurrency
10. [ ] Critical sections
11. [ ] Mutex
12. [ ] Semaphore
13. [ ] Condition variables
14. [ ] Atomic operations
15. [ ] Race conditions
16. [ ] Deadlocks
17. [ ] Starvation
18. [ ] Livelock

## Phase 3 — Memory

19. [ ] Process address space
20. [ ] Stack
21. [ ] Heap
22. [ ] Virtual memory
23. [ ] Paging
24. [ ] Page tables
25. [ ] TLB
26. [ ] Page faults
27. [ ] mmap
28. [ ] Copy-on-write
29. [ ] Page cache

## Phase 4 — I/O

30. [ ] File descriptors
31. [ ] File I/O
32. [ ] Blocking I/O
33. [ ] Non-blocking I/O
34. [ ] select
35. [ ] poll
36. [ ] epoll
37. [ ] Event loops
38. [ ] Reactor pattern

## Phase 5 — File Systems

39. [ ] Files
40. [ ] Directories
41. [ ] Inodes
42. [ ] File permissions
43. [ ] Page cache
44. [ ] Journaling
45. [ ] fsync
46. [ ] SSD
47. [ ] Sequential vs random I/O

## Phase 6 — Linux

48. [ ] Linux processes
49. [ ] Linux threads
50. [ ] /proc
51. [ ] Signals
52. [ ] File descriptors
53. [ ] Networking tools
54. [ ] strace
55. [ ] perf
56. [ ] pprof

## Phase 7 — Go Runtime

57. [ ] Goroutines
58. [ ] G-M-P scheduler
59. [ ] Goroutine stacks
60. [ ] Garbage collector
61. [ ] Escape analysis
62. [ ] Network poller
63. [ ] Blocking system calls
64. [ ] Runtime scheduling

## Phase 8 — Systems Engineering

65. [ ] Worker pools
66. [ ] Backpressure
67. [ ] Connection pooling
68. [ ] Graceful shutdown
69. [ ] Fault tolerance
70. [ ] Crash recovery
71. [ ] Resource limits
72. [ ] Observability

## Phase 9 — Your Projects

73. [ ] URL Shortener
74. [ ] Job Queue
75. [ ] Event Streaming Engine
76. [ ] Reverse Proxy
77. [ ] Real-world Go backend
```

## 🎯 Depth priority for your projects

You **do not need equal depth everywhere**.

| Topic | Interview | URL Shortener | Job Queue | Event Streaming |
|---|---:|---:|---:|---:|
| Processes | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| Threads | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Scheduling | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐ |
| Synchronization | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Deadlocks | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| Virtual Memory | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| File Systems | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| File Descriptors | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| System Calls | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| `epoll` / I/O multiplexing | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Page Cache | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| `mmap` | ⭐⭐⭐ | ⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| Go Scheduler | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Linux | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Graceful Shutdown | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

### The most important connection for your Event Streaming Engine

You should eventually be able to understand this entire path:

```text
Producer
   ↓
TCP Socket
   ↓
File Descriptor
   ↓
Go netpoller / epoll
   ↓
Goroutine
   ↓
Application Buffer
   ↓
Event Batch
   ↓
write()
   ↓
Linux Page Cache
   ↓
fsync()
   ↓
Disk
```

And on the consumer side:

```text
Disk
  ↓
Page Cache
  ↓
read() / mmap()
  ↓
Event Buffer
  ↓
Consumer
  ↓
Goroutine
  ↓
TCP Socket
  ↓
Network
```
