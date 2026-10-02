# C++ Networking & System Programming — Topic List

## 1. Linux & OS Fundamentals
- [ ] Linux Basics
- [ ] Processes
- [ ] Threads
- [ ] Process vs Thread
- [ ] Process Address Space
- [ ] User Space vs Kernel Space
- [ ] System Calls
- [ ] Context Switching
- [ ] CPU Scheduling Basics
- [ ] Virtual Memory
- [ ] File Descriptors
- [ ] Signals
- [ ] Environment Variables
- [ ] Exit Codes

## 2. Linux System Calls
- [ ] `open()`
- [ ] `close()`
- [ ] `read()`
- [ ] `write()`
- [ ] `pread()`
- [ ] `pwrite()`
- [ ] `lseek()`
- [ ] `stat()`
- [ ] `fstat()`
- [ ] `fcntl()`
- [ ] `ioctl()`
- [ ] `dup()`
- [ ] `dup2()`
- [ ] `pipe()`
- [ ] `poll()`
- [ ] `select()`
- [ ] `epoll()`

## 3. File & I/O System Programming
- [ ] File Descriptors
- [ ] Standard Input/Output/Error
- [ ] Blocking I/O
- [ ] Non-Blocking I/O
- [ ] Buffered I/O
- [ ] Direct I/O
- [ ] File Permissions
- [ ] File Metadata
- [ ] Directory Operations
- [ ] `mmap()`
- [ ] Memory-Mapped Files
- [ ] `fsync()`
- [ ] `fdatasync()`

## 4. Processes
- [ ] `fork()`
- [ ] `exec()`
- [ ] `wait()`
- [ ] `waitpid()`
- [ ] Parent Process
- [ ] Child Process
- [ ] Process IDs
- [ ] Process Groups
- [ ] Zombie Processes
- [ ] Orphan Processes
- [ ] Process Creation
- [ ] Process Termination

## 5. Inter-Process Communication
- [ ] Pipes
- [ ] Named Pipes / FIFO
- [ ] Shared Memory
- [ ] Memory-Mapped IPC
- [ ] Unix Domain Sockets
- [ ] Signals
- [ ] Message Queues
- [ ] Semaphores
- [ ] Synchronization Between Processes

## 6. Signals
- [ ] Signal Basics
- [ ] Signal Handlers
- [ ] `signal()`
- [ ] `sigaction()`
- [ ] `SIGINT`
- [ ] `SIGTERM`
- [ ] `SIGKILL`
- [ ] `SIGCHLD`
- [ ] `SIGPIPE`
- [ ] `SIGALRM`
- [ ] Signal Masks
- [ ] `sigprocmask()`
- [ ] Signal Safety

## 7. Linux Threads
- [ ] POSIX Threads
- [ ] `pthread_create()`
- [ ] `pthread_join()`
- [ ] `pthread_detach()`
- [ ] Thread Attributes
- [ ] Thread IDs
- [ ] Thread Local Storage
- [ ] CPU Affinity
- [ ] Thread Scheduling
- [ ] Thread Synchronization

## 8. Concurrency System Programming
- [ ] Mutexes
- [ ] Condition Variables
- [ ] Semaphores
- [ ] Read-Write Locks
- [ ] Atomics
- [ ] Spinlocks
- [ ] Race Conditions
- [ ] Deadlocks
- [ ] Lock Ordering
- [ ] Lock-Free Programming

## 9. Socket Programming Fundamentals 🔥
- [ ] Socket Concept
- [ ] Client-Server Model
- [ ] TCP
- [ ] UDP
- [ ] IPv4
- [ ] IPv6
- [ ] IP Address
- [ ] Port
- [ ] Socket Address
- [ ] `socket()`
- [ ] `bind()`
- [ ] `listen()`
- [ ] `accept()`
- [ ] `connect()`
- [ ] `send()`
- [ ] `recv()`
- [ ] `sendto()`
- [ ] `recvfrom()`
- [ ] `shutdown()`
- [ ] `close()`

## 10. TCP Networking
- [ ] TCP Connection
- [ ] Three-Way Handshake
- [ ] TCP Streams
- [ ] Reliable Delivery
- [ ] Sequence Numbers
- [ ] Acknowledgements
- [ ] Retransmission
- [ ] Flow Control
- [ ] Congestion Control
- [ ] Connection Termination
- [ ] Half-Closed Connections
- [ ] Keepalive
- [ ] TCP Buffering

## 11. UDP Networking
- [ ] UDP Datagram
- [ ] Connectionless Communication
- [ ] `sendto()`
- [ ] `recvfrom()`
- [ ] Packet Loss
- [ ] Packet Ordering
- [ ] UDP Broadcasting
- [ ] UDP Multicast
- [ ] When to Use TCP vs UDP

## 12. Socket Configuration
- [ ] `setsockopt()`
- [ ] `getsockopt()`
- [ ] `SO_REUSEADDR`
- [ ] `SO_REUSEPORT`
- [ ] `SO_KEEPALIVE`
- [ ] `SO_RCVBUF`
- [ ] `SO_SNDBUF`
- [ ] Socket Timeouts
- [ ] `fcntl()`
- [ ] Non-Blocking Sockets

## 13. Address Resolution
- [ ] `sockaddr`
- [ ] `sockaddr_in`
- [ ] `sockaddr_in6`
- [ ] `sockaddr_storage`
- [ ] `inet_pton()`
- [ ] `inet_ntop()`
- [ ] `getaddrinfo()`
- [ ] `getnameinfo()`
- [ ] DNS Basics
- [ ] Hostname Resolution

## 14. Blocking vs Non-Blocking I/O
- [ ] Blocking Socket
- [ ] Non-Blocking Socket
- [ ] `O_NONBLOCK`
- [ ] `EAGAIN`
- [ ] `EWOULDBLOCK`
- [ ] Partial Read
- [ ] Partial Write
- [ ] Retry Logic
- [ ] Backpressure

## 15. I/O Multiplexing 🔥
- [ ] `select()`
- [ ] `poll()`
- [ ] `epoll()`
- [ ] File Descriptor Sets
- [ ] Readiness
- [ ] Level Triggered
- [ ] Edge Triggered
- [ ] `epoll_create1()`
- [ ] `epoll_ctl()`
- [ ] `epoll_wait()`
- [ ] Event Loop

## 16. Event-Driven Programming
- [ ] Event Loop
- [ ] Reactor Pattern
- [ ] Event Dispatcher
- [ ] Event Handlers
- [ ] Non-Blocking Architecture
- [ ] Connection State
- [ ] Timer Events
- [ ] Signal Events
- [ ] Backpressure
- [ ] Graceful Shutdown

## 17. Network Protocols
- [ ] Protocol Design
- [ ] Message Framing
- [ ] Request/Response
- [ ] Binary Protocols
- [ ] Text Protocols
- [ ] Headers
- [ ] Payload
- [ ] Length-Prefix Framing
- [ ] Delimiter-Based Framing
- [ ] Serialization
- [ ] Deserialization
- [ ] Endianness
- [ ] Network Byte Order

## 18. HTTP Networking
- [ ] HTTP Request
- [ ] HTTP Response
- [ ] HTTP Methods
- [ ] HTTP Headers
- [ ] HTTP Status Codes
- [ ] HTTP/1.0
- [ ] HTTP/1.1
- [ ] Persistent Connections
- [ ] Keep-Alive
- [ ] Chunked Transfer
- [ ] HTTP Parsing
- [ ] HTTP Server
- [ ] HTTP Client

## 19. TLS & Secure Networking
- [ ] TLS Basics
- [ ] Certificates
- [ ] Public Key Cryptography
- [ ] TLS Handshake
- [ ] HTTPS
- [ ] Certificate Validation
- [ ] OpenSSL Basics
- [ ] Secure Socket Communication

## 20. Network Concurrency
- [ ] Thread-Per-Connection
- [ ] Thread Pool
- [ ] Worker Threads
- [ ] Event Loop
- [ ] Reactor Model
- [ ] Connection Pool
- [ ] Concurrent Connections
- [ ] Shared State
- [ ] Synchronization
- [ ] Graceful Connection Shutdown

## 21. C++ Networking Libraries
- [ ] POSIX Sockets
- [ ] Boost.Asio
- [ ] Standalone Asio
- [ ] Networking TS Concepts
- [ ] Synchronous I/O
- [ ] Asynchronous I/O
- [ ] Async Callbacks
- [ ] Completion Handlers
- [ ] Coroutines with Networking

## 22. C++ Coroutines for Networking
- [ ] Coroutine Basics
- [ ] `co_await`
- [ ] `co_return`
- [ ] `co_yield`
- [ ] Awaitable
- [ ] Promise Type
- [ ] Coroutine Frame
- [ ] Async Networking
- [ ] Coroutine-Based Server

## 23. Performance Networking
- [ ] Throughput
- [ ] Latency
- [ ] Jitter
- [ ] Connection Scalability
- [ ] Socket Buffering
- [ ] Zero-Copy
- [ ] Scatter/Gather I/O
- [ ] `readv()`
- [ ] `writev()`
- [ ] `sendmsg()`
- [ ] `recvmsg()`
- [ ] `sendfile()`
- [ ] `splice()`
- [ ] CPU Affinity
- [ ] Cache Locality

## 24. Advanced Linux I/O
- [ ] `io_uring`
- [ ] Asynchronous I/O
- [ ] Submission Queue
- [ ] Completion Queue
- [ ] SQE
- [ ] CQE
- [ ] Async File I/O
- [ ] Async Network I/O
- [ ] `mmap()`
- [ ] `sendfile()`
- [ ] `splice()`

## 25. Debugging & Observability
- [ ] GDB
- [ ] `strace`
- [ ] `ltrace`
- [ ] `tcpdump`
- [ ] Wireshark
- [ ] `ss`
- [ ] `netstat`
- [ ] `lsof`
- [ ] `ip`
- [ ] `curl`
- [ ] `nc`
- [ ] Core Dumps
- [ ] Stack Traces
- [ ] Sanitizers
- [ ] Performance Profiling

## 26. System Performance
- [ ] CPU Profiling
- [ ] Memory Profiling
- [ ] I/O Profiling
- [ ] `perf`
- [ ] CPU Cache
- [ ] Context Switches
- [ ] System Call Overhead
- [ ] Lock Contention
- [ ] Network Bottlenecks
- [ ] Latency Profiling
- [ ] Throughput Benchmarking

## 27. Production Server Concepts
- [ ] Graceful Shutdown
- [ ] Signal Handling
- [ ] Connection Limits
- [ ] Timeout Management
- [ ] Rate Limiting
- [ ] Backpressure
- [ ] Resource Limits
- [ ] Error Recovery
- [ ] Logging
- [ ] Metrics
- [ ] Health Checks
- [ ] Configuration Management

## 28. Redis-Relevant System Programming 🔥🔥
- [ ] TCP Server
- [ ] Client Connections
- [ ] Event Loop
- [ ] Non-Blocking Sockets
- [ ] `epoll`
- [ ] Protocol Parsing
- [ ] Command Dispatch
- [ ] In-Memory Data Structures
- [ ] Hash Tables
- [ ] Dynamic Strings
- [ ] Memory Management
- [ ] Persistence
- [ ] File I/O
- [ ] `mmap`
- [ ] Background Threads
- [ ] Timers
- [ ] TTL
- [ ] Eviction
- [ ] Connection Management
- [ ] Graceful Shutdown
- [ ] Performance Profiling

## 29. Projects

### Beginner
- [ ] TCP Echo Server
- [ ] TCP Echo Client
- [ ] UDP Chat
- [ ] File Transfer Server
- [ ] HTTP Client

### Intermediate
- [ ] HTTP/1.1 Server
- [ ] Multi-Threaded TCP Server
- [ ] Thread Pool Server
- [ ] Concurrent Chat Server
- [ ] Proxy Server
- [ ] Load Balancer

### Advanced
- [ ] Event-Driven HTTP Server
- [ ] `epoll` Web Server
- [ ] Custom Binary Protocol Server
- [ ] High-Performance TCP Server
- [ ] Async Networking Library
- [ ] Mini Redis Server

### Target Project

```text
C++ Fundamentals
       ↓
Memory Management
       ↓
Concurrency
       ↓
Linux System Programming
       ↓
Socket Programming
       ↓
Non-Blocking I/O
       ↓
epoll
       ↓
Event Loop
       ↓
Protocol Design
       ↓
Persistence
       ↓
Redis Clone
```