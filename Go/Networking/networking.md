# Go Networking & net/http

Complete roadmap for learning networking in Go, from TCP/UDP socket programming to production-grade HTTP servers, clients, middleware, connection management, streaming, and graceful shutdown.

---

# PART 1 — NETWORKING FUNDAMENTALS

## 1. Networking Basics

- [ ] Client-Server Architecture
- [ ] Peer-to-Peer Architecture
- [ ] IP Address
- [ ] IPv4
- [ ] IPv6
- [ ] Private IP
- [ ] Public IP
- [ ] Loopback Address
- [ ] `localhost`
- [ ] Port
- [ ] Socket
- [ ] MAC Address
- [ ] DNS
- [ ] Domain Name
- [ ] Hostname
- [ ] Network Interface
- [ ] Network Routing
- [ ] NAT
- [ ] Firewall Basics

---

# 2. OSI & TCP/IP

- [ ] OSI Model
- [ ] TCP/IP Model
- [ ] Application Layer
- [ ] Transport Layer
- [ ] Network Layer
- [ ] Data Link Layer
- [ ] TCP/IP Layer Mapping
- [ ] Encapsulation
- [ ] Decapsulation
- [ ] Packets
- [ ] Frames
- [ ] Segments
- [ ] MTU
- [ ] Network addressing

---

# 3. TCP

- [ ] What is TCP?
- [ ] Connection-oriented communication
- [ ] TCP 3-Way Handshake
- [ ] SYN
- [ ] SYN-ACK
- [ ] ACK
- [ ] TCP Connection Lifecycle
- [ ] TCP Sequence Numbers
- [ ] TCP Acknowledgements
- [ ] TCP Retransmission
- [ ] TCP Reliability
- [ ] TCP Flow Control
- [ ] TCP Congestion Control
- [ ] TCP Receive Buffer
- [ ] TCP Send Buffer
- [ ] TCP Connection Termination
- [ ] FIN
- [ ] RST
- [ ] TIME_WAIT
- [ ] Keep-Alive
- [ ] Half-Closed Connections
- [ ] TCP Backlog
- [ ] Connection Timeout
- [ ] TCP Streams

---

# 4. UDP

- [ ] What is UDP?
- [ ] Connectionless communication
- [ ] UDP Datagram
- [ ] UDP vs TCP
- [ ] UDP reliability
- [ ] UDP packet loss
- [ ] UDP ordering
- [ ] UDP broadcasting
- [ ] UDP multicasting
- [ ] UDP use cases
- [ ] Datagram size
- [ ] UDP sockets

---

# 5. TCP vs UDP

- [ ] Connection-oriented vs connectionless
- [ ] Reliability
- [ ] Ordering
- [ ] Flow control
- [ ] Congestion control
- [ ] Latency
- [ ] Overhead
- [ ] Streaming vs datagrams
- [ ] When to use TCP
- [ ] When to use UDP

---

# PART 2 — GO NETWORKING

# 6. Go `net` Package

- [ ] `net` package
- [ ] `net.Conn`
- [ ] `net.Listener`
- [ ] `net.TCPConn`
- [ ] `net.TCPListener`
- [ ] `net.UDPConn`
- [ ] `net.UDPAddr`
- [ ] `net.IP`
- [ ] `net.IPAddr`
- [ ] `net.TCPAddr`
- [ ] `net.ResolveTCPAddr`
- [ ] `net.ResolveUDPAddr`

---

# 7. TCP Server

- [ ] `net.Listen`
- [ ] Listening on an address
- [ ] Port binding
- [ ] `Listener.Accept`
- [ ] Accept loop
- [ ] Connection handling
- [ ] `Conn.Read`
- [ ] `Conn.Write`
- [ ] `Conn.Close`
- [ ] Handling multiple connections
- [ ] Connection-per-goroutine
- [ ] Server shutdown

Basic architecture:

    Client
       |
       | TCP
       v
    Listener
       |
       v
    Accept()
       |
       v
    Connection
       |
       v
    Goroutine

---

# 8. TCP Client

- [ ] `net.Dial`
- [ ] `net.DialTimeout`
- [ ] `net.Dialer`
- [ ] Establishing TCP connection
- [ ] Reading responses
- [ ] Writing requests
- [ ] Closing connections
- [ ] Connection timeout
- [ ] Retry connection
- [ ] Reconnecting

---

# 9. TCP Message Framing

TCP is a byte stream, so learn:

- [ ] Why TCP has no message boundaries
- [ ] Message framing
- [ ] Fixed-length messages
- [ ] Delimiter-based protocol
- [ ] Length-prefixed messages
- [ ] Header + payload
- [ ] Binary protocol
- [ ] Text protocol
- [ ] Partial reads
- [ ] Partial writes
- [ ] Buffering
- [ ] Message size limits

This is extremely important for your Event Streaming Engine.

---

# 10. TCP Connection Management

- [ ] Connection timeout
- [ ] Read deadline
- [ ] Write deadline
- [ ] `SetDeadline`
- [ ] `SetReadDeadline`
- [ ] `SetWriteDeadline`
- [ ] Keep-alive
- [ ] Idle connections
- [ ] Connection limits
- [ ] Connection pooling
- [ ] Connection reuse
- [ ] Reconnection
- [ ] Connection failure
- [ ] Graceful connection close

---

# 11. Go `io` Package

Learn these interfaces deeply:

- [ ] `io.Reader`
- [ ] `io.Writer`
- [ ] `io.ReadWriter`
- [ ] `io.ReadCloser`
- [ ] `io.WriteCloser`
- [ ] `io.ReadWriteCloser`
- [ ] `io.Copy`
- [ ] `io.CopyN`
- [ ] `io.LimitReader`
- [ ] `io.MultiReader`
- [ ] `io.MultiWriter`
- [ ] `io.TeeReader`
- [ ] `io.Pipe`

---

# 12. Buffered Networking

- [ ] `bufio.Reader`
- [ ] `bufio.Writer`
- [ ] `Read`
- [ ] `ReadBytes`
- [ ] `ReadString`
- [ ] `ReadLine`
- [ ] `Write`
- [ ] `Flush`
- [ ] Buffer size
- [ ] Buffered network communication
- [ ] Performance implications

---

# 13. UDP in Go

- [ ] `net.ListenUDP`
- [ ] `net.DialUDP`
- [ ] `ReadFromUDP`
- [ ] `WriteToUDP`
- [ ] `UDPAddr`
- [ ] UDP server
- [ ] UDP client
- [ ] Datagram handling
- [ ] UDP timeouts
- [ ] Packet size

---

# 14. DNS in Go

- [ ] DNS basics
- [ ] DNS resolution
- [ ] `net.LookupHost`
- [ ] `net.LookupIP`
- [ ] `net.LookupAddr`
- [ ] `net.Resolver`
- [ ] Custom DNS resolver
- [ ] DNS caching concepts
- [ ] DNS timeout
- [ ] DNS failure handling

---

# 15. TLS / HTTPS

- [ ] TLS basics
- [ ] HTTPS
- [ ] Certificates
- [ ] Certificate Authority
- [ ] Public/private keys
- [ ] TLS handshake
- [ ] Certificate verification
- [ ] `crypto/tls`
- [ ] `tls.Config`
- [ ] TLS server
- [ ] TLS client
- [ ] `ListenAndServeTLS`
- [ ] `tls.Dial`
- [ ] Server certificates
- [ ] Client certificates
- [ ] Mutual TLS
- [ ] TLS versions
- [ ] Certificate rotation

---

# 16. Network Errors

- [ ] Connection refused
- [ ] Connection reset
- [ ] Connection timeout
- [ ] DNS failure
- [ ] Broken pipe
- [ ] EOF
- [ ] Temporary network errors
- [ ] Retryable errors
- [ ] Non-retryable errors
- [ ] Error wrapping
- [ ] Network error classification

---

# 17. Network Timeouts

- [ ] Dial timeout
- [ ] Connection timeout
- [ ] Read timeout
- [ ] Write timeout
- [ ] Request timeout
- [ ] Idle timeout
- [ ] Server timeout
- [ ] Client timeout
- [ ] Context timeout
- [ ] Deadline vs timeout

---

# 18. Connection Pooling

- [ ] Why connection pooling?
- [ ] TCP connection reuse
- [ ] HTTP connection reuse
- [ ] Database connection pooling
- [ ] Maximum connections
- [ ] Idle connections
- [ ] Connection lifetime
- [ ] Pool exhaustion
- [ ] Connection cleanup

---

# PART 3 — HTTP FUNDAMENTALS

# 19. HTTP Basics

- [ ] What is HTTP?
- [ ] HTTP request
- [ ] HTTP response
- [ ] Request line
- [ ] Status line
- [ ] Headers
- [ ] Body
- [ ] HTTP methods
- [ ] HTTP status codes
- [ ] HTTP versions

---

# 20. HTTP Methods

- [ ] GET
- [ ] POST
- [ ] PUT
- [ ] PATCH
- [ ] DELETE
- [ ] HEAD
- [ ] OPTIONS
- [ ] CONNECT
- [ ] TRACE
- [ ] Safe methods
- [ ] Idempotent methods

---

# 21. HTTP Status Codes

## 1xx

- [ ] 100 Continue
- [ ] 101 Switching Protocols

## 2xx

- [ ] 200 OK
- [ ] 201 Created
- [ ] 202 Accepted
- [ ] 204 No Content

## 3xx

- [ ] 301 Moved Permanently
- [ ] 302 Found
- [ ] 304 Not Modified
- [ ] 307 Temporary Redirect
- [ ] 308 Permanent Redirect

## 4xx

- [ ] 400 Bad Request
- [ ] 401 Unauthorized
- [ ] 403 Forbidden
- [ ] 404 Not Found
- [ ] 405 Method Not Allowed
- [ ] 408 Request Timeout
- [ ] 409 Conflict
- [ ] 413 Content Too Large
- [ ] 415 Unsupported Media Type
- [ ] 422 Unprocessable Content
- [ ] 429 Too Many Requests

## 5xx

- [ ] 500 Internal Server Error
- [ ] 501 Not Implemented
- [ ] 502 Bad Gateway
- [ ] 503 Service Unavailable
- [ ] 504 Gateway Timeout

---

# 22. HTTP Headers

- [ ] Request headers
- [ ] Response headers
- [ ] `Content-Type`
- [ ] `Content-Length`
- [ ] `Accept`
- [ ] `Authorization`
- [ ] `User-Agent`
- [ ] `Host`
- [ ] `Connection`
- [ ] `Cache-Control`
- [ ] `ETag`
- [ ] `Last-Modified`
- [ ] `Location`
- [ ] `Cookie`
- [ ] `Set-Cookie`
- [ ] `Origin`
- [ ] `Referer`
- [ ] CORS headers
- [ ] Security headers

---

# PART 4 — GO `net/http`

# 23. `net/http` Package

- [ ] `http.Server`
- [ ] `http.Client`
- [ ] `http.Request`
- [ ] `http.Response`
- [ ] `http.Handler`
- [ ] `http.HandlerFunc`
- [ ] `http.ServeMux`
- [ ] `http.DefaultServeMux`
- [ ] `http.DefaultClient`
- [ ] `http.DefaultTransport`

---

# 24. HTTP Server

- [ ] `http.ListenAndServe`
- [ ] `http.Server`
- [ ] `Server.ListenAndServe`
- [ ] Server configuration
- [ ] Server address
- [ ] Server timeouts
- [ ] ReadTimeout
- [ ] ReadHeaderTimeout
- [ ] WriteTimeout
- [ ] IdleTimeout
- [ ] MaxHeaderBytes

---

# 25. HTTP Handlers

- [ ] `http.Handler`
- [ ] `http.HandlerFunc`
- [ ] `ServeHTTP`
- [ ] Handler composition
- [ ] Handler registration
- [ ] Handler lifecycle
- [ ] Request context
- [ ] Response writer

Example:

    func handler(w http.ResponseWriter, r *http.Request) {
        // request handling
    }

---

# 26. ServeMux

- [ ] `http.NewServeMux`
- [ ] Route registration
- [ ] Path matching
- [ ] Method-based routing
- [ ] Path parameters
- [ ] Wildcards
- [ ] Route precedence
- [ ] Subrouters concept
- [ ] Custom mux

---

# 27. Request Object

Master:

- [ ] `http.Request`
- [ ] `Method`
- [ ] `URL`
- [ ] `Header`
- [ ] `Body`
- [ ] `Host`
- [ ] `RemoteAddr`
- [ ] `RequestURI`
- [ ] `ContentLength`
- [ ] `TransferEncoding`
- [ ] `TLS`
- [ ] `Context`

---

# 28. URL Handling

- [ ] `net/url`
- [ ] URL parsing
- [ ] Scheme
- [ ] Host
- [ ] Path
- [ ] RawQuery
- [ ] Fragment
- [ ] Query parameters
- [ ] `url.Parse`
- [ ] `url.Values`
- [ ] URL encoding
- [ ] URL decoding

Important for your URL Shortener.

---

# 29. Query Parameters

- [ ] Reading query parameters
- [ ] `r.URL.Query()`
- [ ] Multiple values
- [ ] URL encoding
- [ ] Validation
- [ ] Default values
- [ ] Pagination parameters
- [ ] Filtering parameters
- [ ] Sorting parameters

---

# 30. Path Parameters

- [ ] Path variables
- [ ] Route matching
- [ ] Path extraction
- [ ] Parameter validation
- [ ] URL decoding
- [ ] Path traversal considerations

---

# 31. Request Body

- [ ] Reading request body
- [ ] `r.Body`
- [ ] `io.ReadAll`
- [ ] Streaming request body
- [ ] Request size limits
- [ ] `http.MaxBytesReader`
- [ ] JSON request body
- [ ] Form data
- [ ] Multipart form
- [ ] File uploads

---

# 32. ResponseWriter

- [ ] `http.ResponseWriter`
- [ ] `Write`
- [ ] `WriteHeader`
- [ ] Headers
- [ ] Status code
- [ ] Response body
- [ ] Header ordering
- [ ] Content-Type
- [ ] Response streaming

---

# 33. JSON APIs

- [ ] `encoding/json`
- [ ] JSON encoding
- [ ] JSON decoding
- [ ] Struct tags
- [ ] `omitempty`
- [ ] Request validation
- [ ] Response models
- [ ] Error responses
- [ ] Consistent API response format
- [ ] JSON content type

---

# 34. HTTP Middleware

- [ ] What is middleware?
- [ ] Middleware chaining
- [ ] Logging middleware
- [ ] Authentication middleware
- [ ] Authorization middleware
- [ ] Recovery middleware
- [ ] CORS middleware
- [ ] Rate limiting middleware
- [ ] Request ID middleware
- [ ] Metrics middleware
- [ ] Timeout middleware

Pattern:

    Request
       |
       v
    Logger
       |
       v
    Auth
       |
       v
    Rate Limiter
       |
       v
    Handler

---

# 35. HTTP Client

- [ ] `http.Client`
- [ ] `http.Get`
- [ ] `http.Post`
- [ ] `http.NewRequest`
- [ ] `client.Do`
- [ ] Request headers
- [ ] Request body
- [ ] Response body
- [ ] Response status
- [ ] Client timeout
- [ ] Context cancellation
- [ ] Connection reuse

---

# 36. HTTP Transport

- [ ] `http.Transport`
- [ ] Connection pooling
- [ ] Keep-alive
- [ ] MaxIdleConns
- [ ] MaxIdleConnsPerHost
- [ ] MaxConnsPerHost
- [ ] IdleConnTimeout
- [ ] TLS configuration
- [ ] Proxy configuration
- [ ] DialContext
- [ ] ResponseHeaderTimeout
- [ ] ExpectContinueTimeout

---

# 37. HTTP Timeouts

Server:

- [ ] ReadTimeout
- [ ] ReadHeaderTimeout
- [ ] WriteTimeout
- [ ] IdleTimeout

Client:

- [ ] Client.Timeout
- [ ] Context timeout
- [ ] Dial timeout
- [ ] TLS handshake timeout
- [ ] Response header timeout
- [ ] Request deadline

---

# 38. HTTP Cookies

- [ ] Cookies
- [ ] `http.Cookie`
- [ ] `Set-Cookie`
- [ ] Reading cookies
- [ ] Secure cookies
- [ ] HttpOnly
- [ ] SameSite
- [ ] Cookie expiration
- [ ] Session cookies
- [ ] Cookie security

---

# 39. Sessions

- [ ] Session concept
- [ ] Session IDs
- [ ] Server-side sessions
- [ ] Cookie-based sessions
- [ ] Session expiration
- [ ] Session invalidation
- [ ] Session storage
- [ ] Distributed sessions

---

# 40. Authentication

- [ ] Basic Authentication
- [ ] Bearer tokens
- [ ] JWT
- [ ] API keys
- [ ] Session authentication
- [ ] Authentication middleware
- [ ] Token validation
- [ ] Token expiration
- [ ] Refresh tokens

---

# 41. CORS

- [ ] Same-Origin Policy
- [ ] CORS
- [ ] Simple requests
- [ ] Preflight requests
- [ ] OPTIONS
- [ ] `Access-Control-Allow-Origin`
- [ ] `Access-Control-Allow-Methods`
- [ ] `Access-Control-Allow-Headers`
- [ ] Credentials
- [ ] CORS middleware

---

# 42. HTTP Redirects

- [ ] HTTP redirects
- [ ] 301
- [ ] 302
- [ ] 303
- [ ] 307
- [ ] 308
- [ ] `http.Redirect`
- [ ] Redirect behavior
- [ ] Redirect loops

Important for URL Shortener.

---

# 43. HTTP Streaming

- [ ] Streaming response
- [ ] `http.Flusher`
- [ ] Flush
- [ ] Chunked transfer
- [ ] Server-Sent Events
- [ ] Long polling
- [ ] Streaming request body
- [ ] Streaming response body
- [ ] Backpressure
- [ ] Connection lifetime

---

# 44. WebSockets

- [ ] WebSocket concept
- [ ] HTTP Upgrade
- [ ] Persistent connection
- [ ] Full-duplex communication
- [ ] WebSocket client
- [ ] WebSocket server
- [ ] Ping/Pong
- [ ] Connection lifecycle
- [ ] WebSocket timeout
- [ ] WebSocket concurrency

---

# 45. HTTP/2

- [ ] HTTP/2 basics
- [ ] Multiplexing
- [ ] Streams
- [ ] Frames
- [ ] Header compression
- [ ] Server push concept
- [ ] HTTP/2 connection reuse
- [ ] HTTP/2 in Go

---

# 46. HTTP/3

- [ ] HTTP/3 basics
- [ ] QUIC
- [ ] UDP-based transport
- [ ] HTTP/3 vs HTTP/2
- [ ] QUIC streams
- [ ] Connection migration
- [ ] HTTP/3 use cases

Advanced — learn after HTTP/1.1 and HTTP/2.

---

# 47. Graceful HTTP Shutdown

- [ ] `Server.Shutdown`
- [ ] Shutdown context
- [ ] OS signals
- [ ] Stop accepting new connections
- [ ] Finish active requests
- [ ] Close idle connections
- [ ] Shutdown timeout
- [ ] Cleanup background workers
- [ ] Flush pending data

---

# 48. HTTP Security

- [ ] HTTPS
- [ ] TLS
- [ ] Input validation
- [ ] Request size limits
- [ ] Header validation
- [ ] Authentication
- [ ] Authorization
- [ ] CORS
- [ ] CSRF
- [ ] XSS
- [ ] SQL injection
- [ ] Path traversal
- [ ] SSRF
- [ ] Request smuggling concepts
- [ ] Security headers
- [ ] Rate limiting

---

# 49. HTTP Reliability

- [ ] Timeouts
- [ ] Retries
- [ ] Retryable status codes
- [ ] Exponential backoff
- [ ] Jitter
- [ ] Circuit breaker
- [ ] Rate limiting
- [ ] Connection pooling
- [ ] Idempotency
- [ ] Request cancellation
- [ ] Graceful degradation

---

# 50. HTTP Observability

- [ ] Request logging
- [ ] Structured logging
- [ ] Request ID
- [ ] Trace ID
- [ ] Request duration
- [ ] Status code metrics
- [ ] Request count
- [ ] Error rate
- [ ] Latency
- [ ] P50
- [ ] P95
- [ ] P99
- [ ] OpenTelemetry
- [ ] Prometheus metrics

---

# 51. HTTP Testing

- [ ] `httptest`
- [ ] `httptest.NewServer`
- [ ] `httptest.NewRecorder`
- [ ] Handler testing
- [ ] Integration testing
- [ ] HTTP client testing
- [ ] Mock HTTP server
- [ ] Request validation testing
- [ ] Error response testing
- [ ] Middleware testing
- [ ] Timeout testing
- [ ] Concurrent request testing

---

# 52. Advanced Networking

- [ ] Connection pooling
- [ ] Load balancing
- [ ] Reverse proxy
- [ ] Proxy servers
- [ ] Forward proxy
- [ ] Reverse proxy
- [ ] Health checks
- [ ] Service discovery
- [ ] DNS-based discovery
- [ ] Connection draining
- [ ] Network backpressure
- [ ] Network rate limiting
- [ ] Network retries
- [ ] Circuit breakers

---

# 53. Go Reverse Proxy

- [ ] `httputil`
- [ ] `httputil.NewSingleHostReverseProxy`
- [ ] Reverse proxy architecture
- [ ] Request forwarding
- [ ] Response forwarding
- [ ] Header forwarding
- [ ] Proxy timeouts
- [ ] Load balancing
- [ ] Health checks

---

# 54. Production Networking

- [ ] TCP connection limits
- [ ] HTTP connection limits
- [ ] Timeouts
- [ ] Keep-alive
- [ ] Connection pooling
- [ ] Backpressure
- [ ] Rate limiting
- [ ] Retries
- [ ] Circuit breakers
- [ ] Graceful shutdown
- [ ] TLS
- [ ] Observability
- [ ] Network metrics
- [ ] Load testing
- [ ] Network debugging

---

# 55. Networking Tools

Learn to use:

- [ ] `curl`
- [ ] `wget`
- [ ] `ping`
- [ ] `traceroute`
- [ ] `nslookup`
- [ ] `dig`
- [ ] `netstat`
- [ ] `ss`
- [ ] `lsof`
- [ ] `telnet`
- [ ] `nc` / netcat
- [ ] Wireshark
- [ ] tcpdump

---

# 56. Practical Projects

## Beginner

- [ ] TCP Echo Server
- [ ] TCP Echo Client
- [ ] UDP Echo Server
- [ ] UDP Echo Client
- [ ] HTTP Server using `net/http`
- [ ] HTTP Client
- [ ] REST API

## Intermediate

- [ ] Concurrent TCP Server
- [ ] HTTP Reverse Proxy
- [ ] Rate Limiter
- [ ] Connection Pool
- [ ] WebSocket Server
- [ ] File Upload Server
- [ ] Streaming HTTP Server

## Advanced

- [ ] URL Shortener
- [ ] Load Balancer
- [ ] Job Queue
- [ ] TCP Message Broker
- [ ] Event Streaming Engine
- [ ] HTTP Proxy
- [ ] Distributed Service

---

# 57. Learning Order

Follow this order:

1. [ ] Networking fundamentals
2. [ ] TCP/IP
3. [ ] TCP
4. [ ] UDP
5. [ ] Socket concepts
6. [ ] Go `net`
7. [ ] TCP server
8. [ ] TCP client
9. [ ] Message framing
10. [ ] `io.Reader` / `io.Writer`
11. [ ] Buffered I/O
12. [ ] Connection management
13. [ ] Timeouts
14. [ ] DNS
15. [ ] TLS
16. [ ] HTTP fundamentals
17. [ ] HTTP methods
18. [ ] HTTP status codes
19. [ ] HTTP headers
20. [ ] `net/http`
21. [ ] HTTP handlers
22. [ ] ServeMux
23. [ ] Request/Response
24. [ ] JSON APIs
25. [ ] Middleware
26. [ ] HTTP Client
27. [ ] HTTP Transport
28. [ ] Connection pooling
29. [ ] Authentication
30. [ ] CORS
31. [ ] HTTP streaming
32. [ ] WebSockets
33. [ ] Graceful shutdown
34. [ ] HTTP security
35. [ ] HTTP reliability
36. [ ] HTTP observability
37. [ ] HTTP testing
38. [ ] Reverse proxy
39. [ ] HTTP/2
40. [ ] HTTP/3
41. [ ] Production networking

---

# 58. Project Mapping

## URL Shortener

Required:

- [ ] HTTP
- [ ] `net/http`
- [ ] Request/Response
- [ ] URL parsing
- [ ] Path parameters
- [ ] HTTP redirects
- [ ] JSON
- [ ] Middleware
- [ ] HTTP client
- [ ] Database
- [ ] Redis
- [ ] Rate limiting
- [ ] Graceful shutdown

---

## Job Queue

Required:

- [ ] TCP/HTTP
- [ ] Goroutines
- [ ] Channels
- [ ] Worker pools
- [ ] Context
- [ ] Timeouts
- [ ] Retry
- [ ] Backoff
- [ ] Connection management
- [ ] Graceful shutdown

---

## Event Streaming Engine

Required:

- [ ] TCP
- [ ] `net`
- [ ] TCP server
- [ ] TCP client
- [ ] Message framing
- [ ] Binary/text protocol
- [ ] `io.Reader`
- [ ] `io.Writer`
- [ ] Buffering
- [ ] Connection management
- [ ] Goroutines
- [ ] Channels
- [ ] Backpressure
- [ ] Timeouts
- [ ] Append-only log
- [ ] Topics
- [ ] Partitions
- [ ] Offsets
- [ ] Consumer groups
- [ ] Message ordering
- [ ] Graceful shutdown

---

# 59. Mastery Checklist

Before considering Go networking strong, I should be able to:

- [ ] Build a TCP server from scratch
- [ ] Build a TCP client from scratch
- [ ] Handle multiple TCP clients concurrently
- [ ] Design a simple TCP protocol
- [ ] Implement message framing
- [ ] Handle partial reads
- [ ] Handle connection failures
- [ ] Implement timeouts
- [ ] Implement graceful connection shutdown
- [ ] Build a UDP server
- [ ] Resolve DNS using Go
- [ ] Build a TLS server
- [ ] Build an HTTP server using only `net/http`
- [ ] Build an HTTP client
- [ ] Implement middleware
- [ ] Implement authentication middleware
- [ ] Implement rate limiting
- [ ] Implement request timeouts
- [ ] Implement graceful HTTP shutdown
- [ ] Implement HTTP streaming
- [ ] Build a reverse proxy
- [ ] Test HTTP services using `httptest`
- [ ] Understand connection pooling
- [ ] Understand HTTP/1.1
- [ ] Understand HTTP/2
- [ ] Understand HTTP/3 conceptually
- [ ] Debug network problems using CLI tools
- [ ] Load-test a network service
- [ ] Monitor network service latency and throughput

---

# Priority for My Projects

## Must Master

- [ ] TCP
- [ ] UDP basics
- [ ] Go `net`
- [ ] `io.Reader`
- [ ] `io.Writer`
- [ ] TCP Server
- [ ] TCP Client
- [ ] Message Framing
- [ ] Connection Management
- [ ] Timeouts
- [ ] HTTP
- [ ] `net/http`
- [ ] HTTP Handlers
- [ ] HTTP Client
- [ ] Middleware
- [ ] JSON
- [ ] URL Handling
- [ ] HTTP Redirects
- [ ] Context
- [ ] Graceful Shutdown
- [ ] Testing

## Important

- [ ] TLS
- [ ] DNS
- [ ] Connection Pooling
- [ ] Reverse Proxy
- [ ] Rate Limiting
- [ ] HTTP Streaming
- [ ] WebSockets
- [ ] HTTP/2
- [ ] Observability

## Advanced / Later

- [ ] HTTP/3
- [ ] QUIC
- [ ] Custom DNS Resolver
- [ ] Advanced TCP tuning
- [ ] Zero-copy networking
- [ ] Kernel networking
- [ ] eBPF networking
- [ ] Advanced network optimization