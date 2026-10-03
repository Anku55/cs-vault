

```markdown
# Computer Networks Roadmap

# 1. Network Fundamentals ⭐⭐⭐⭐⭐

- [ ] What is a computer network?
- [ ] Network types
- [ ] LAN
- [ ] WAN
- [ ] MAN
- [ ] PAN
- [ ] Internet
- [ ] Intranet
- [ ] Client
- [ ] Server
- [ ] Peer-to-peer
- [ ] Network topology
- [ ] Bus topology
- [ ] Star topology
- [ ] Ring topology
- [ ] Mesh topology
- [ ] Network devices
- [ ] Hub
- [ ] Switch
- [ ] Router
- [ ] Bridge
- [ ] Gateway
- [ ] Modem
- [ ] Access Point
- [ ] Firewall
- [ ] Proxy
- [ ] Reverse Proxy

---

# 2. OSI Model ⭐⭐⭐⭐⭐

- [ ] OSI model
- [ ] Application Layer
- [ ] Presentation Layer
- [ ] Session Layer
- [ ] Transport Layer
- [ ] Network Layer
- [ ] Data Link Layer
- [ ] Physical Layer
- [ ] Responsibilities of each layer
- [ ] Protocols at each layer
- [ ] Encapsulation
- [ ] Decapsulation
- [ ] PDU
- [ ] Segment
- [ ] Packet
- [ ] Frame
- [ ] Bits

---

# 3. TCP/IP Model ⭐⭐⭐⭐⭐

- [ ] TCP/IP model
- [ ] Application Layer
- [ ] Transport Layer
- [ ] Internet Layer
- [ ] Network Access Layer
- [ ] OSI vs TCP/IP
- [ ] Protocol mapping
- [ ] Encapsulation
- [ ] Decapsulation

---

# 4. Data Link Layer

- [ ] Frames
- [ ] MAC addresses
- [ ] Ethernet
- [ ] ARP
- [ ] MAC address table
- [ ] Switch forwarding
- [ ] Broadcast
- [ ] Unicast
- [ ] Multicast
- [ ] Collision domain
- [ ] Broadcast domain
- [ ] VLAN
- [ ] 802.1Q
- [ ] MTU
- [ ] Frame structure
- [ ] Error detection
- [ ] CRC

---

# 5. IP Addressing ⭐⭐⭐⭐⭐

- [ ] IPv4
- [ ] IPv6
- [ ] IP address
- [ ] Public IP
- [ ] Private IP
- [ ] Loopback
- [ ] localhost
- [ ] 127.0.0.1
- [ ] 0.0.0.0
- [ ] Subnet mask
- [ ] CIDR
- [ ] Network address
- [ ] Broadcast address
- [ ] Host address
- [ ] Subnetting
- [ ] Supernetting
- [ ] Default gateway
- [ ] Address resolution

---

# 6. Subnetting ⭐⭐⭐⭐

- [ ] Binary representation of IP
- [ ] Network bits
- [ ] Host bits
- [ ] CIDR notation
- [ ] /8
- [ ] /16
- [ ] /24
- [ ] /32
- [ ] Subnet calculation
- [ ] Number of hosts
- [ ] Number of subnets
- [ ] Network range
- [ ] Broadcast address
- [ ] VLSM
- [ ] Subnetting problems

---

# 7. Routing ⭐⭐⭐⭐⭐

- [ ] Routing
- [ ] Routing table
- [ ] Default route
- [ ] Static routing
- [ ] Dynamic routing
- [ ] Next hop
- [ ] Route selection
- [ ] Longest prefix match
- [ ] TTL
- [ ] ICMP
- [ ] Traceroute
- [ ] Routing protocols
- [ ] RIP
- [ ] OSPF
- [ ] BGP
- [ ] Autonomous System
- [ ] Internet routing

---

# 8. ARP

- [ ] ARP
- [ ] ARP request
- [ ] ARP response
- [ ] ARP cache
- [ ] MAC resolution
- [ ] Gratuitous ARP
- [ ] ARP spoofing
- [ ] IPv4 address → MAC address

---

# 9. ICMP

- [ ] ICMP
- [ ] Echo request
- [ ] Echo reply
- [ ] ping
- [ ] Destination unreachable
- [ ] Time exceeded
- [ ] traceroute
- [ ] ICMP in network debugging

---

# 10. UDP ⭐⭐⭐⭐⭐

- [ ] UDP
- [ ] UDP header
- [ ] Source port
- [ ] Destination port
- [ ] Length
- [ ] Checksum
- [ ] Connectionless communication
- [ ] UDP characteristics
- [ ] UDP limitations
- [ ] UDP use cases
- [ ] UDP vs TCP
- [ ] UDP in Go

---

# 11. TCP ⭐⭐⭐⭐⭐

- [ ] TCP
- [ ] TCP header
- [ ] Source port
- [ ] Destination port
- [ ] Sequence number
- [ ] Acknowledgment number
- [ ] Window size
- [ ] Flags
- [ ] Checksum
- [ ] SYN
- [ ] ACK
- [ ] FIN
- [ ] RST
- [ ] PSH
- [ ] URG

---

# 12. TCP Three-Way Handshake ⭐⭐⭐⭐⭐

- [ ] SYN
- [ ] SYN-ACK
- [ ] ACK
- [ ] Sequence numbers
- [ ] Initial sequence number
- [ ] Connection establishment
- [ ] Why three steps?
- [ ] TCP connection state

Understand:

```text
Client                  Server

   SYN  ------------->

        <------------- SYN + ACK

   ACK  ------------->
```

---

# 13. TCP Connection Termination

- [ ] FIN
- [ ] ACK
- [ ] Four-way termination
- [ ] Half-close
- [ ] TIME_WAIT
- [ ] CLOSE_WAIT
- [ ] FIN_WAIT
- [ ] LAST_ACK
- [ ] RST
- [ ] Graceful shutdown
- [ ] Connection reset

---

# 14. TCP Reliability ⭐⭐⭐⭐⭐

- [ ] Reliable delivery
- [ ] Sequence numbers
- [ ] Acknowledgments
- [ ] Retransmission
- [ ] Timeout
- [ ] Duplicate packets
- [ ] Duplicate ACK
- [ ] Out-of-order packets
- [ ] Packet loss
- [ ] Checksum
- [ ] Sliding window

---

# 15. Flow Control ⭐⭐⭐⭐⭐

- [ ] Flow control
- [ ] Receive window
- [ ] Send window
- [ ] Sliding window
- [ ] Receiver buffer
- [ ] Window advertisement
- [ ] Zero window
- [ ] Window scaling

---

# 16. TCP Congestion Control ⭐⭐⭐⭐⭐

- [ ] Congestion
- [ ] Congestion window
- [ ] Slow start
- [ ] Congestion avoidance
- [ ] Fast retransmit
- [ ] Fast recovery
- [ ] AIMD
- [ ] cwnd
- [ ] ssthresh
- [ ] TCP Reno
- [ ] TCP Cubic
- [ ] TCP BBR
- [ ] Flow control vs congestion control

---

# 17. Ports ⭐⭐⭐⭐⭐

- [ ] Port number
- [ ] Source port
- [ ] Destination port
- [ ] Well-known ports
- [ ] Registered ports
- [ ] Dynamic ports
- [ ] Ephemeral ports
- [ ] Port binding
- [ ] Port listening
- [ ] Port conflicts
- [ ] Socket address

Important ports:

- [ ] 22 SSH
- [ ] 53 DNS
- [ ] 80 HTTP
- [ ] 443 HTTPS
- [ ] 5432 PostgreSQL
- [ ] 6379 Redis

---

# 18. Sockets ⭐⭐⭐⭐⭐

- [ ] Socket
- [ ] IP + Port
- [ ] Socket address
- [ ] TCP socket
- [ ] UDP socket
- [ ] Client socket
- [ ] Server socket
- [ ] bind
- [ ] listen
- [ ] accept
- [ ] connect
- [ ] send
- [ ] receive
- [ ] close
- [ ] shutdown

---

# 19. TCP Server Architecture

Understand:

```text
Client
   |
   v
Socket
   |
   v
Server
   |
   +---- accept()
   |
   +---- Connection
   |
   +---- Read
   |
   +---- Process
   |
   +---- Write
   |
   +---- Close
```

Study:

- [ ] TCP server lifecycle
- [ ] TCP client lifecycle
- [ ] accept loop
- [ ] connection handling
- [ ] blocking sockets
- [ ] non-blocking sockets
- [ ] connection timeout
- [ ] read timeout
- [ ] write timeout
- [ ] graceful shutdown

---

# 20. DNS ⭐⭐⭐⭐⭐

- [ ] DNS
- [ ] Domain name
- [ ] Resolver
- [ ] DNS server
- [ ] Root server
- [ ] TLD server
- [ ] Authoritative server
- [ ] Recursive resolver
- [ ] DNS lookup
- [ ] DNS caching
- [ ] DNS TTL
- [ ] DNS records
- [ ] A
- [ ] AAAA
- [ ] CNAME
- [ ] MX
- [ ] NS
- [ ] TXT
- [ ] SRV
- [ ] PTR
- [ ] Reverse DNS

Understand:

```text
google.com
     |
     v
DNS Resolver
     |
     v
Root
     |
     v
.com
     |
     v
Authoritative DNS
     |
     v
IP Address
```

---

# 21. HTTP ⭐⭐⭐⭐⭐

- [ ] HTTP
- [ ] Request
- [ ] Response
- [ ] HTTP methods
- [ ] GET
- [ ] POST
- [ ] PUT
- [ ] PATCH
- [ ] DELETE
- [ ] HEAD
- [ ] OPTIONS
- [ ] HTTP headers
- [ ] Request body
- [ ] Response body
- [ ] Status codes
- [ ] Content-Type
- [ ] Content-Length
- [ ] Accept
- [ ] Authorization
- [ ] User-Agent
- [ ] Host
- [ ] Cache-Control
- [ ] Cookie
- [ ] Set-Cookie

---

# 22. HTTP Status Codes ⭐⭐⭐⭐⭐

- [ ] 1xx
- [ ] 2xx
- [ ] 3xx
- [ ] 4xx
- [ ] 5xx
- [ ] 200 OK
- [ ] 201 Created
- [ ] 202 Accepted
- [ ] 204 No Content
- [ ] 301
- [ ] 302
- [ ] 304
- [ ] 307
- [ ] 308
- [ ] 400
- [ ] 401
- [ ] 403
- [ ] 404
- [ ] 405
- [ ] 409
- [ ] 429
- [ ] 500
- [ ] 502
- [ ] 503
- [ ] 504

---

# 23. HTTP Versions

- [ ] HTTP/1.0
- [ ] HTTP/1.1
- [ ] HTTP/2
- [ ] HTTP/3
- [ ] HTTP/1.1 persistent connections
- [ ] Keep-Alive
- [ ] HTTP pipelining
- [ ] HTTP/2 multiplexing
- [ ] HTTP/2 streams
- [ ] HTTP/2 frames
- [ ] HPACK
- [ ] HTTP/3
- [ ] QUIC
- [ ] HTTP/3 over QUIC

---

# 24. HTTP Connection Management

- [ ] Short-lived connections
- [ ] Persistent connections
- [ ] Keep-alive
- [ ] Connection pooling
- [ ] Idle connections
- [ ] Connection timeout
- [ ] Read timeout
- [ ] Write timeout
- [ ] Idle timeout
- [ ] Connection reuse
- [ ] Connection exhaustion

---

# 25. HTTPS ⭐⭐⭐⭐⭐

- [ ] HTTPS
- [ ] TLS
- [ ] TLS handshake
- [ ] Certificates
- [ ] Certificate Authority
- [ ] Public key
- [ ] Private key
- [ ] Symmetric encryption
- [ ] Asymmetric encryption
- [ ] Session keys
- [ ] Certificate validation
- [ ] Certificate chain
- [ ] TLS versions
- [ ] TLS 1.2
- [ ] TLS 1.3
- [ ] SNI
- [ ] ALPN

---

# 26. HTTP Authentication

- [ ] Basic Authentication
- [ ] Bearer Authentication
- [ ] API Keys
- [ ] Session Authentication
- [ ] Cookies
- [ ] JWT
- [ ] OAuth 2.0
- [ ] OpenID Connect
- [ ] Token expiration
- [ ] Refresh tokens

---

# 27. HTTP Cookies

- [ ] Cookie
- [ ] Set-Cookie
- [ ] Cookie attributes
- [ ] Domain
- [ ] Path
- [ ] Expires
- [ ] Max-Age
- [ ] Secure
- [ ] HttpOnly
- [ ] SameSite
- [ ] Session cookies
- [ ] Persistent cookies

---

# 28. HTTP Caching

- [ ] Browser cache
- [ ] Server cache
- [ ] Proxy cache
- [ ] Cache-Control
- [ ] max-age
- [ ] no-cache
- [ ] no-store
- [ ] ETag
- [ ] If-None-Match
- [ ] Last-Modified
- [ ] If-Modified-Since
- [ ] 304 Not Modified

---

# 29. REST API Networking

- [ ] REST
- [ ] Resources
- [ ] Resource URLs
- [ ] HTTP methods
- [ ] Statelessness
- [ ] Idempotency
- [ ] Safe methods
- [ ] API versioning
- [ ] Pagination
- [ ] Filtering
- [ ] Sorting
- [ ] Rate limiting
- [ ] Error responses

---

# 30. WebSockets

- [ ] WebSocket
- [ ] WebSocket handshake
- [ ] HTTP Upgrade
- [ ] Persistent connection
- [ ] Full-duplex communication
- [ ] Frames
- [ ] Ping
- [ ] Pong
- [ ] Close
- [ ] WebSocket vs HTTP
- [ ] WebSocket vs SSE
- [ ] WebSocket in Go

---

# 31. Server-Sent Events

- [ ] SSE
- [ ] Event stream
- [ ] text/event-stream
- [ ] Long-lived HTTP connection
- [ ] Event ID
- [ ] Reconnection
- [ ] SSE vs WebSocket
- [ ] SSE in Go

---

# 32. gRPC

- [ ] gRPC
- [ ] RPC
- [ ] Protocol Buffers
- [ ] .proto files
- [ ] Service definition
- [ ] Unary RPC
- [ ] Server streaming
- [ ] Client streaming
- [ ] Bidirectional streaming
- [ ] HTTP/2
- [ ] gRPC deadlines
- [ ] gRPC metadata
- [ ] gRPC status codes
- [ ] gRPC interceptors
- [ ] gRPC in Go

---

# 33. Proxies ⭐⭐⭐⭐⭐

- [ ] Proxy
- [ ] Forward proxy
- [ ] Reverse proxy
- [ ] Reverse proxy architecture
- [ ] Nginx
- [ ] Load balancing
- [ ] TLS termination
- [ ] Request forwarding
- [ ] Header forwarding
- [ ] X-Forwarded-For
- [ ] X-Forwarded-Proto
- [ ] Proxy timeout

---

# 34. Load Balancing ⭐⭐⭐⭐⭐

- [ ] Load balancer
- [ ] Layer 4 load balancing
- [ ] Layer 7 load balancing
- [ ] Round Robin
- [ ] Weighted Round Robin
- [ ] Least Connections
- [ ] Consistent Hashing
- [ ] Health Checks
- [ ] Active Health Checks
- [ ] Passive Health Checks
- [ ] Failover
- [ ] Session Affinity
- [ ] Sticky Sessions

---

# 35. NAT

- [ ] NAT
- [ ] Private IP
- [ ] Public IP
- [ ] Source NAT
- [ ] Destination NAT
- [ ] Port Address Translation
- [ ] NAT traversal
- [ ] Why NAT exists

---

# 36. Firewall

- [ ] Firewall
- [ ] Packet filtering
- [ ] Stateful firewall
- [ ] Stateless firewall
- [ ] Inbound rules
- [ ] Outbound rules
- [ ] Ports
- [ ] Network security groups

---

# 37. Network Security

- [ ] TLS
- [ ] HTTPS
- [ ] Certificate validation
- [ ] Man-in-the-middle attack
- [ ] DNS spoofing
- [ ] ARP spoofing
- [ ] IP spoofing
- [ ] DDoS
- [ ] SYN flood
- [ ] Rate limiting
- [ ] Firewall
- [ ] Network segmentation

---

# 38. Reliability in Distributed Systems ⭐⭐⭐⭐⭐

- [ ] Network failures
- [ ] Packet loss
- [ ] Packet duplication
- [ ] Packet reordering
- [ ] Network timeout
- [ ] Connection failure
- [ ] Partial failure
- [ ] Service unavailable
- [ ] Retry
- [ ] Exponential backoff
- [ ] Jitter
- [ ] Timeout
- [ ] Circuit breaker
- [ ] Bulkhead
- [ ] Idempotency
- [ ] Request deduplication

---

# 39. Timeouts ⭐⭐⭐⭐⭐

Understand separately:

- [ ] Connection timeout
- [ ] DNS timeout
- [ ] TLS handshake timeout
- [ ] Request timeout
- [ ] Read timeout
- [ ] Write timeout
- [ ] Idle timeout
- [ ] Context deadline
- [ ] Server timeout
- [ ] Client timeout

---

# 40. Retry Mechanisms ⭐⭐⭐⭐⭐

- [ ] Why retries?
- [ ] Retryable errors
- [ ] Non-retryable errors
- [ ] Fixed retry
- [ ] Exponential backoff
- [ ] Exponential backoff + jitter
- [ ] Retry limits
- [ ] Retry storms
- [ ] Idempotent requests
- [ ] Retry + timeout interaction

---

# 41. Backpressure ⭐⭐⭐⭐⭐

- [ ] What is backpressure?
- [ ] Producer faster than consumer
- [ ] Consumer slower than producer
- [ ] Buffering
- [ ] Bounded queues
- [ ] Blocking producers
- [ ] Dropping messages
- [ ] Rate limiting
- [ ] Flow control
- [ ] Backpressure in streaming systems

---

# 42. Network Buffers

- [ ] Send buffer
- [ ] Receive buffer
- [ ] Kernel socket buffers
- [ ] Application buffers
- [ ] Buffer overflow
- [ ] Buffer sizing
- [ ] TCP receive window
- [ ] TCP send window

---

# 43. Network Debugging Tools ⭐⭐⭐⭐⭐

- [ ] ping
- [ ] traceroute
- [ ] tracert
- [ ] nslookup
- [ ] dig
- [ ] curl
- [ ] wget
- [ ] netstat
- [ ] ss
- [ ] ip
- [ ] ipconfig
- [ ] ifconfig
- [ ] route
- [ ] arp
- [ ] tcpdump
- [ ] Wireshark
- [ ] lsof

---

# 44. curl

Learn:

- [ ] GET request
- [ ] POST request
- [ ] PUT request
- [ ] DELETE request
- [ ] Headers
- [ ] Request body
- [ ] JSON
- [ ] Cookies
- [ ] Authentication
- [ ] Redirects
- [ ] TLS debugging
- [ ] Connection debugging
- [ ] Timing information

---

# 45. Wireshark

- [ ] Packet capture
- [ ] Capture filters
- [ ] Display filters
- [ ] TCP handshake
- [ ] TCP termination
- [ ] DNS packets
- [ ] HTTP packets
- [ ] TLS packets
- [ ] Retransmissions
- [ ] Duplicate ACKs
- [ ] TCP streams
- [ ] Packet timing

---

# 46. Go Networking Fundamentals ⭐⭐⭐⭐⭐

- [ ] net package
- [ ] net.Conn
- [ ] net.Listener
- [ ] net.Dial
- [ ] net.DialTCP
- [ ] net.Listen
- [ ] net.ListenTCP
- [ ] Accept
- [ ] Read
- [ ] Write
- [ ] Close
- [ ] TCP connections
- [ ] UDP connections
- [ ] IP addresses
- [ ] TCPAddr
- [ ] UDPAddr
- [ ] DNS resolution
- [ ] Network timeouts
- [ ] Context cancellation

---

# 47. Go TCP Server

Build:

```text
TCP Client
     |
     v
net.Dial()
     |
     v
TCP Server
     |
     v
net.Listen()
     |
     v
Accept()
     |
     v
net.Conn
     |
     v
Read / Write
```

Practice:

- [ ] TCP echo server
- [ ] TCP client
- [ ] Multiple clients
- [ ] Goroutine per connection
- [ ] Connection timeout
- [ ] Graceful shutdown
- [ ] Framing messages
- [ ] Protocol design

---

# 48. Go UDP Server

- [ ] UDP listener
- [ ] UDP client
- [ ] ReadFromUDP
- [ ] WriteToUDP
- [ ] UDP packets
- [ ] Packet size
- [ ] Timeouts
- [ ] UDP concurrency

Build:

- [ ] UDP echo server
- [ ] UDP metrics collector
- [ ] Simple UDP discovery service

---

# 49. Go net/http ⭐⭐⭐⭐⭐

- [ ] http.Server
- [ ] http.Client
- [ ] http.Request
- [ ] http.Response
- [ ] http.Handler
- [ ] http.HandlerFunc
- [ ] ServeMux
- [ ] Handle
- [ ] HandleFunc
- [ ] ListenAndServe
- [ ] ListenAndServeTLS
- [ ] Serve
- [ ] Request routing
- [ ] HTTP middleware
- [ ] Request context
- [ ] ResponseWriter

---

# 50. net/http Server Configuration

- [ ] ReadTimeout
- [ ] ReadHeaderTimeout
- [ ] WriteTimeout
- [ ] IdleTimeout
- [ ] MaxHeaderBytes
- [ ] TLSConfig
- [ ] Server shutdown
- [ ] Graceful shutdown
- [ ] Connection state
- [ ] HTTP keep-alive

---

# 51. Go HTTP Client

- [ ] http.Get
- [ ] http.Post
- [ ] http.NewRequest
- [ ] http.NewRequestWithContext
- [ ] http.Client
- [ ] http.Transport
- [ ] Request headers
- [ ] Request body
- [ ] Response body
- [ ] Status codes
- [ ] Client timeout
- [ ] Connection pooling
- [ ] Keep-alive
- [ ] TLS configuration
- [ ] Proxy configuration

---

# 52. Go HTTP Transport

- [ ] http.Transport
- [ ] DialContext
- [ ] MaxIdleConns
- [ ] MaxIdleConnsPerHost
- [ ] MaxConnsPerHost
- [ ] IdleConnTimeout
- [ ] TLSHandshakeTimeout
- [ ] ResponseHeaderTimeout
- [ ] ExpectContinueTimeout
- [ ] Connection reuse

---

# 53. Go HTTP Middleware

Implement:

- [ ] Logging middleware
- [ ] Authentication middleware
- [ ] Authorization middleware
- [ ] Request ID middleware
- [ ] Recovery middleware
- [ ] Rate limiting middleware
- [ ] CORS middleware
- [ ] Timeout middleware
- [ ] Metrics middleware
- [ ] Tracing middleware

---

# 54. Go HTTP Request Lifecycle

Understand:

```text
Client
  ↓
DNS
  ↓
TCP
  ↓
TLS
  ↓
HTTP
  ↓
Go net/http
  ↓
Router
  ↓
Middleware
  ↓
Handler
  ↓
Service
  ↓
Database / Redis
  ↓
Response
```

Know exactly what happens at every stage.

---

# 55. URL Shortener Networking

Understand:

```text
Browser
   |
   v
DNS
   |
   v
TCP
   |
   v
TLS
   |
   v
HTTP
   |
   v
Go Server
   |
   v
Router
   |
   v
Database / Redis
```

Study:

- [ ] DNS resolution
- [ ] HTTP redirect
- [ ] 301
- [ ] 302
- [ ] 307
- [ ] 308
- [ ] Short URL lookup
- [ ] Cache lookup
- [ ] Database lookup
- [ ] Connection pooling
- [ ] Redirect latency
- [ ] HTTP keep-alive

---

# 56. Job Queue Networking

Understand:

```text
Producer
    |
    | HTTP / gRPC
    v
Queue Server
    |
    v
Workers
```

Study:

- [ ] Producer communication
- [ ] Worker communication
- [ ] TCP connections
- [ ] HTTP APIs
- [ ] gRPC
- [ ] Message acknowledgement
- [ ] Retry
- [ ] Timeout
- [ ] Backpressure
- [ ] Connection failure
- [ ] Worker disconnection
- [ ] Heartbeats

---

# 57. Event Streaming Engine Networking

Understand:

```text
Producer
    |
    | TCP / HTTP / gRPC
    v
Event Broker
    |
    +---- Topic A
    |
    +---- Topic B
    |
    +---- Topic C
    |
    v
Consumers
```

Study:

- [ ] Persistent connections
- [ ] TCP
- [ ] Message framing
- [ ] Request/response protocol
- [ ] Streaming protocol
- [ ] Batching
- [ ] Compression
- [ ] Backpressure
- [ ] Flow control
- [ ] Consumer acknowledgement
- [ ] Heartbeats
- [ ] Connection failure
- [ ] Reconnection
- [ ] Retry
- [ ] Consumer timeout
- [ ] Producer timeout
- [ ] Message ordering
- [ ] Network partition
- [ ] Partial failure

---

# 58. Message Framing ⭐⭐⭐⭐⭐

Important for your Event Streaming Engine.

Learn:

- [ ] Why TCP has no message boundaries
- [ ] Byte stream
- [ ] Message framing
- [ ] Length-prefix framing
- [ ] Delimiter-based framing
- [ ] Fixed-size messages
- [ ] Header + payload
- [ ] Serialization
- [ ] Deserialization
- [ ] Partial reads
- [ ] Partial writes
- [ ] Buffer management

Example:

```text
+----------+----------+----------------+
| Length   | Type     | Payload        |
+----------+----------+----------------+
```

---

# 59. Serialization

- [ ] JSON
- [ ] XML
- [ ] Protocol Buffers
- [ ] MessagePack
- [ ] Binary protocols
- [ ] Encoding
- [ ] Decoding
- [ ] Schema evolution
- [ ] Backward compatibility
- [ ] Forward compatibility
- [ ] Serialization performance

---

# 60. Network Protocol Design ⭐⭐⭐⭐⭐

Learn how to design your own protocol:

- [ ] Request format
- [ ] Response format
- [ ] Header
- [ ] Payload
- [ ] Message ID
- [ ] Request ID
- [ ] Status code
- [ ] Error format
- [ ] Versioning
- [ ] Length prefix
- [ ] Heartbeat
- [ ] Timeout
- [ ] Retry
- [ ] Acknowledgement
- [ ] Connection lifecycle

---

# 61. Distributed Systems Networking

- [ ] Client-server communication
- [ ] Service-to-service communication
- [ ] RPC
- [ ] REST
- [ ] gRPC
- [ ] Service discovery
- [ ] Load balancing
- [ ] Health checks
- [ ] Heartbeats
- [ ] Timeouts
- [ ] Retries
- [ ] Circuit breakers
- [ ] Backpressure
- [ ] Rate limiting
- [ ] Idempotency
- [ ] Partial failures
- [ ] Network partitions
- [ ] Eventual consistency

---

# 62. Service Discovery

- [ ] Service discovery
- [ ] Client-side discovery
- [ ] Server-side discovery
- [ ] DNS-based discovery
- [ ] Service registry
- [ ] Health checks
- [ ] Registration
- [ ] Deregistration
- [ ] Load balancing

---

# 63. Rate Limiting

- [ ] Why rate limiting?
- [ ] Fixed window
- [ ] Sliding window
- [ ] Token bucket
- [ ] Leaky bucket
- [ ] Per-IP rate limiting
- [ ] Per-user rate limiting
- [ ] Global rate limiting
- [ ] Distributed rate limiting
- [ ] Redis-based rate limiting

---

# 64. Observability for Networking

- [ ] Request logs
- [ ] Request ID
- [ ] Correlation ID
- [ ] Latency
- [ ] Throughput
- [ ] Error rate
- [ ] Connection count
- [ ] Active connections
- [ ] Network errors
- [ ] Timeout metrics
- [ ] Retry metrics
- [ ] Distributed tracing

---

# 65. Networking Interview Questions

- [ ] What happens when you type google.com?
- [ ] What happens when you enter a URL in a browser?
- [ ] TCP vs UDP
- [ ] TCP three-way handshake
- [ ] TCP four-way termination
- [ ] Why TIME_WAIT?
- [ ] What is a socket?
- [ ] What is a port?
- [ ] What is DNS?
- [ ] How DNS resolution works
- [ ] What is ARP?
- [ ] What is NAT?
- [ ] What is CIDR?
- [ ] What is subnetting?
- [ ] What is HTTP?
- [ ] HTTP/1.1 vs HTTP/2
- [ ] HTTP/2 vs HTTP/3
- [ ] HTTP vs HTTPS
- [ ] TLS handshake
- [ ] What is a reverse proxy?
- [ ] What is a load balancer?
- [ ] L4 vs L7 load balancing
- [ ] What is WebSocket?
- [ ] REST vs gRPC
- [ ] What is connection pooling?
- [ ] What happens when TCP packet is lost?
- [ ] Flow control vs congestion control
- [ ] What is backpressure?
- [ ] What is a timeout?
- [ ] Why do distributed systems need retries?
- [ ] Why can retries be dangerous?
- [ ] What is idempotency?
- [ ] What is a network partition?
- [ ] What is a partial failure?

---

# 66. Practical Networking Projects

## Project 1 — TCP Echo Server

- [ ] TCP server
- [ ] TCP client
- [ ] Multiple clients
- [ ] Goroutines
- [ ] Connection handling
- [ ] Timeouts
- [ ] Graceful shutdown

## Project 2 — HTTP Server from Scratch

- [ ] TCP
- [ ] HTTP request parsing
- [ ] HTTP response
- [ ] Headers
- [ ] Status codes
- [ ] Routing
- [ ] Middleware

## Project 3 — Reverse Proxy

- [ ] HTTP forwarding
- [ ] Request forwarding
- [ ] Response forwarding
- [ ] Connection pooling
- [ ] Timeouts
- [ ] Load balancing
- [ ] Health checks

## Project 4 — Rate Limiter

- [ ] Token bucket
- [ ] Per-IP limits
- [ ] Redis
- [ ] HTTP middleware
- [ ] Distributed rate limiting

## Project 5 — URL Shortener

- [ ] HTTP
- [ ] DNS
- [ ] TCP
- [ ] HTTPS
- [ ] Redirects
- [ ] Connection pooling
- [ ] Redis
- [ ] PostgreSQL

## Project 6 — Job Queue

- [ ] Producer
- [ ] Consumer
- [ ] TCP/HTTP
- [ ] Worker connections
- [ ] Acknowledgement
- [ ] Retry
- [ ] Timeout
- [ ] Heartbeat
- [ ] Backpressure

## Project 7 — Event Streaming Engine

- [ ] TCP server
- [ ] Custom protocol
- [ ] Message framing
- [ ] Serialization
- [ ] Persistent connections
- [ ] Producer
- [ ] Consumer
- [ ] Topics
- [ ] Partitions
- [ ] Consumer groups
- [ ] Heartbeats
- [ ] Acknowledgements
- [ ] Backpressure
- [ ] Flow control
- [ ] Retry
- [ ] Reconnection
- [ ] Event ordering

---

# 67. Recommended Learning Order

## Phase 1 — Networking Fundamentals

1. [ ] OSI Model
2. [ ] TCP/IP Model
3. [ ] Ethernet
4. [ ] MAC
5. [ ] IP
6. [ ] IPv4
7. [ ] IPv6
8. [ ] Subnetting
9. [ ] ARP
10. [ ] ICMP
11. [ ] Routing
12. [ ] Ports

## Phase 2 — Transport Layer

13. [ ] TCP
14. [ ] UDP
15. [ ] TCP Handshake
16. [ ] TCP Termination
17. [ ] Reliability
18. [ ] Flow Control
19. [ ] Congestion Control
20. [ ] Sockets

## Phase 3 — Application Layer

21. [ ] DNS
22. [ ] HTTP
23. [ ] HTTPS
24. [ ] TLS
25. [ ] HTTP/1.1
26. [ ] HTTP/2
27. [ ] HTTP/3
28. [ ] WebSockets
29. [ ] SSE
30. [ ] gRPC

## Phase 4 — Backend Networking

31. [ ] REST
32. [ ] Reverse Proxy
33. [ ] Load Balancing
34. [ ] Connection Pooling
35. [ ] Timeouts
36. [ ] Retries
37. [ ] Backpressure
38. [ ] Rate Limiting
39. [ ] Circuit Breakers
40. [ ] Service Discovery

## Phase 5 — Go Networking

41. [ ] net package
42. [ ] TCP server
43. [ ] TCP client
44. [ ] UDP server
45. [ ] net/http
46. [ ] HTTP middleware
47. [ ] HTTP client
48. [ ] HTTP Transport
49. [ ] Connection pooling
50. [ ] TLS
51. [ ] Graceful shutdown

## Phase 6 — Distributed Systems Networking

52. [ ] RPC
53. [ ] gRPC
54. [ ] Service-to-service communication
55. [ ] Network failures
56. [ ] Timeouts
57. [ ] Retries
58. [ ] Idempotency
59. [ ] Backpressure
60. [ ] Heartbeats
61. [ ] Network partitions
62. [ ] Partial failures
63. [ ] Load balancing
64. [ ] Service discovery

## Phase 7 — Systems Projects

65. [ ] TCP Echo Server
66. [ ] HTTP Server
67. [ ] Reverse Proxy
68. [ ] Rate Limiter
69. [ ] URL Shortener
70. [ ] Job Queue
71. [ ] Event Streaming Engine
```

