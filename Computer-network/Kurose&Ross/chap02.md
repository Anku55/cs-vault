# 1. Principles of Network Applications

This is **Section 2.1 of Kurose & Ross, Chapter 2 — Application Layer**. The book's central idea is that a network application consists of **programs running on different end systems that communicate with each other over the network**. Importantly, the application software runs at the **end systems**, not inside routers or link-layer switches. Computer Networking- A Top-Down…

The section is divided exactly into the five topics you listed:

1. **Network Application Architectures**
2. **Processes Communicating**
3. **Transport Services Available to Applications**
4. **Transport Services Provided by the Internet**
5. **Application-Layer Protocols** Computer Networking- A Top-Down…

---

# 1. Network Application Architectures

First, distinguish **network architecture** from **application architecture**.

### Network architecture

This is the underlying Internet architecture:

```text
Application
Transport
Network
Link
Physical
```

It is essentially fixed from the application developer's point of view.

### Application architecture

This describes **how the application itself is structured across different end systems**.

The book focuses on two predominant architectures:

1. **Client-server architecture**
2. **Peer-to-peer (P2P) architecture** Computer Networking- A Top-Down…

---

## 1.1 Client-Server Architecture

In a **client-server architecture**, there is an **always-on server** that services requests from many clients.

```text
             ┌── Client 1
             │
             ├── Client 2
Server ──────┼── Client 3
             │
             └── Client 4
```

### Server

The server:

- is always on
- waits for requests
- services requests from clients
- generally has a fixed, well-known IP address

### Client

The client:

- initiates communication
- sends requests to the server
- receives responses from the server

For example, in the Web:

```text
Browser                         Web Server
(Client)                        (Server)
    │                               │
    │────── HTTP request ──────────→│
    │                               │
    │←───── HTTP response ─────────│
```

The browser is the **client process**, and the Web server is the **server process**. Computer Networking- A Top-Down…

### Examples

The book gives these examples of client-server applications:

- Web
- FTP
- Telnet
- e-mail Computer Networking- A Top-Down…

---

## Why use data centers?

A single server may not be able to handle requests from a huge number of clients.

For example:

```text
Millions of clients
       ↓
   One server
       ↓
   Overloaded
```

Therefore, popular Internet services use **data centers containing many servers**.

Conceptually:

```text
             ┌── Server
             ├── Server
Clients ─────┼── Server
             ├── Server
             └── Server
                  │
              Data Center
```

The book gives search engines, Internet commerce, Web mail, and social networking as examples of services using one or more data centers. Computer Networking- A Top-Down…

---

# 1.2 Peer-to-Peer (P2P) Architecture

In a **P2P architecture**, there is little or no dependence on dedicated servers.

Instead, hosts communicate **directly with one another**.

These hosts are called **peers**.

```text
       Peer A
       ↙   ↘
   Peer B ↔ Peer C
       ↘   ↙
       Peer D
```

Peers are typically:

- desktops
- laptops
- other user-controlled hosts

and may be **intermittently connected**. Computer Networking- A Top-Down…

### Examples

The book gives:

- BitTorrent
- Xunlei
- Skype

as examples of applications based on P2P architecture. Computer Networking- A Top-Down…

---

## Important advantage: Self-scalability

This is a very important P2P concept.

Suppose a new peer joins a P2P file-sharing system.

It creates:

### Additional workload

because the new peer may request files.

But it also creates:

### Additional service capacity

because that peer can upload files to other peers.

So:

```text
New peer joins
      ↓
More demand
      +
More uploading capacity
      ↓
Self-scalability
```

The book identifies **self-scalability** as one of the most compelling features of P2P architecture. Computer Networking- A Top-Down…

### Other advantages

- Less dependence on dedicated servers
- Less server infrastructure
- Less server bandwidth requirement
- Cost-effective

### Challenges

Because P2P is highly decentralized, it creates challenges involving:

- security
- performance
- reliability Computer Networking- A Top-Down…

---

# 1.3 Hybrid Architecture

Some applications combine client-server and P2P characteristics.

For example, the book describes some instant-messaging systems where:

```text
Server
  │
  └── Tracks users' IP addresses

User A ←────────→ User B
       Direct message
```

The server helps locate users, but messages may travel directly between user hosts.

Therefore:

> **Hybrid architecture = combination of client-server and P2P elements.** Computer Networking- A Top-Down…

---

# Client-Server vs P2P

| Feature | Client-Server | P2P |
|---|---|---|
| Dedicated server | Yes | Little/no dedicated server |
| Communication | Client ↔ Server | Peer ↔ Peer |
| Server always on | Usually yes | Not necessarily |
| Server infrastructure | Significant | Usually less |
| Scalability | Server can become bottleneck | Self-scalable |
| Architecture | Centralized | Decentralized |
| Example | Web | BitTorrent |

---

# 2. Processes Communicating

Now we move from **application architecture** to the actual entities that communicate.

The book makes an important distinction:

> It is not technically the programs themselves that communicate; it is **processes**.

A **process** is a program that is currently running on an end system. Computer Networking- A Top-Down…

---

## 2.1 Processes on Different Hosts

Suppose two processes are running on different computers:

```text
Host A                         Host B

Process A                      Process B
    │                              ▲
    │                              │
    └────── Computer Network ──────┘
```

The sending process:

1. creates a message
2. sends it into the network

The receiving process:

1. receives the message
2. processes it
3. may send a response

Computer Networking- A Top-Down…

The communicating processes reside at the **application layer**.

---

# 2.2 Client and Server Processes

A network application generally consists of pairs of communicating processes.

For a particular communication session:

### Client

> The process that **initiates communication**.

### Server

> The process that **waits to be contacted**.

Computer Networking- A Top-Down…

### Example: Web

```text
Browser                         Web Server
Client                          Server

   │                               │
   │────── Request ───────────────→│
   │                               │
   │←────── Response ──────────────│
```

### Example: P2P

Suppose Peer A asks Peer B for a file:

```text
Peer A ───── request ─────→ Peer B
client                         server
```

In another session, Peer B could request a file from Peer A.

Therefore, a P2P process can act as:

```text
Client + Server
```

depending on the particular communication session. Computer Networking- A Top-Down…

---

# 2.3 The Socket

This is one of the **most important concepts in Section 2.1**.

A process communicates with the network through a **socket**.

The book describes a socket as the **interface between the application process and the transport-layer protocol**. Computer Networking- A Top-Down…

Think of:

```text
Application Process
       │
       │
    SOCKET
       │
       │
Transport Layer
       │
       ↓
Network
```

### House analogy

The book compares:

- **Process → house**
- **Socket → door**

The process sends a message through its socket, just as someone sends something through a door.

The transport layer is on the other side of the door and is responsible for transporting the message toward the destination socket. Computer Networking- A Top-Down…

---

## Who controls what?

The application developer controls the:

> **application side of the socket**

The transport-layer implementation controls the:

> **transport side of the socket**

Therefore:

```text
Application developer
        ↓
Application
        ↓
     SOCKET
        ↓
TCP / UDP
        ↓
Network
```

The socket is therefore the **API/interface between the application and transport layer**.

---

# 2.4 Addressing a Process

The network needs to know **which process** should receive the data.

It is not enough to identify only the destination host.

The destination host may be running many processes.

Therefore, the Internet uses:

### IP address

Identifies the **host**.

### Port number

Identifies the appropriate **process/application** on that host.

Conceptually:

```text
IP address       → Which host?
Port number      → Which process?
```

For example, the book notes:

- Web server → port **80**
- SMTP mail server → port **25** Computer Networking- A Top-Down…

So:

```text
Destination = IP address + port number
```

This is what allows transport-layer protocols to deliver data to the correct application process.

---

# 3. Transport Services Available to Applications

An application does not directly control the underlying network.

Instead:

```text
Application
     ↓
Transport Layer
     ↓
Network Layer
     ↓
Link Layer
     ↓
Physical Layer
```

The application chooses a **transport-layer protocol** based on the services it needs.

The book classifies possible transport services along **four dimensions**:

1. Reliable data transfer
2. Throughput
3. Timing
4. Security Computer Networking- A Top-Down…

---

# 3.1 Reliable Data Transfer

Packets can be lost or corrupted inside a network.

For example:

```text
Sender
  ↓
Router buffer overflow
  ↓
Packet lost
```

For applications such as:

- e-mail
- file transfer
- remote host access
- Web document transfer
- financial applications

loss of data can be unacceptable.

Therefore, a transport protocol may provide:

> **Reliable data transfer**

This means the sending process can pass data to the transport layer and rely on the data being delivered correctly and completely to the receiving process. Computer Networking- A Top-Down…

---

## Loss-tolerant applications

Some applications can tolerate some data loss.

The book gives **conversational audio/video** as an example.

A lost packet may cause:

```text
Lost packet
    ↓
Small audio/video glitch
```

but may not seriously affect the communication.

Therefore:

> Not every application needs reliable data transfer. Computer Networking- A Top-Down…

---

# 3.2 Throughput

**Throughput** is the rate at which bits can be delivered from the sending process to the receiving process.

Other network sessions compete for the same network resources, so available throughput can fluctuate.

A transport protocol could theoretically provide:

> **Guaranteed throughput of \(r\) bits/sec**

meaning the application could request a minimum rate \(r\), and the transport protocol would guarantee that rate. Computer Networking- A Top-Down…

### But important:

The Internet's current transport protocols **do not provide throughput guarantees**. Computer Networking- A Top-Down…

---

# 3.3 Timing

Some applications have timing requirements.

For example, consider Internet telephony:

```text
User speaks
    ↓
Network delay
    ↓
Other user hears
```

If the delay becomes too large, conversation becomes unnatural.

Similarly, in an interactive game, a large delay between:

```text
Player action
      ↓
Network
      ↓
Game response
```

makes the application feel unrealistic. Computer Networking- A Top-Down…

A transport protocol could theoretically provide:

> **Timing guarantees**

such as ensuring that data arrives within a specified maximum delay.

But:

### The Internet does not provide timing guarantees.

Computer Networking- A Top-Down…

---

# 3.4 Security

A transport protocol could provide security services.

For example, it could:

```text
Sender
  ↓
Encrypt
  ↓
Network
  ↓
Decrypt
  ↓
Receiver
```

This can provide **confidentiality**.

Other possible security services include:

- data integrity
- end-point authentication

Computer Networking- A Top-Down…

The book later discusses these topics in detail in **Chapter 8**.

---

# 3.5 Summary of Possible Transport Services

| Service | Meaning |
|---|---|
| **Reliable data transfer** | Data reaches receiver correctly and completely |
| **Throughput** | Guaranteed/minimum data-delivery rate |
| **Timing** | Guaranteed delay or timing bound |
| **Security** | Confidentiality, integrity, authentication |

---

# 4. Transport Services Provided by the Internet

Now the book moves from:

> **What services could a transport protocol provide?**

to:

> **What does the actual Internet provide?**

The Internet provides two transport protocols to applications:

1. **TCP**
2. **UDP**

The application developer chooses between them. Computer Networking- A Top-Down…

---

# 4.1 TCP

TCP provides:

### 1. Connection-oriented service

Before application data begins flowing, the client and server perform a **handshake**.

```text
Client                     Server
  │                           │
  │──── control information ─→│
  │←─── control information ──│
  │──── control information ─→│
  │                           │
  │════ Application data ════│
```

After this, a TCP connection exists between the two sockets.

The connection is **full-duplex**:

> both processes can send data to each other simultaneously.

When communication finishes, the connection is torn down. Computer Networking- A Top-Down…

---

### 2. Reliable data transfer

TCP guarantees that data delivered to the receiving socket:

- is not missing
- is not duplicated
- is delivered in the correct order

The book describes TCP as providing a reliable byte stream. Computer Networking- A Top-Down…

---

### 3. Congestion control

TCP also has **congestion control**.

If the network becomes congested:

```text
Network congestion
       ↓
TCP detects/handles congestion
       ↓
Sender's rate is reduced
```

This protects the network from being overwhelmed by a TCP connection. Computer Networking- A Top-Down…

---

### TCP does NOT provide

TCP does **not** provide guaranteed:

- throughput
- timing

And ordinary TCP itself does not provide encryption/security; the book discusses security enhancements such as SSL separately. Computer Networking- A Top-Down…

---

# 4.2 UDP

UDP is described by the book as a:

> **no-frills, lightweight transport protocol**

It provides minimal services. Computer Networking- A Top-Down…

### UDP provides:

- connectionless communication
- process-to-process delivery
- error checking

### UDP does NOT provide:

- reliable data transfer
- connection establishment
- congestion control
- ordering guarantee
- throughput guarantee
- timing guarantee

Computer Networking- A Top-Down…

---

## Connectionless

With UDP:

```text
Application
    ↓
UDP
    ↓
Network
```

There is **no handshake** before data transmission.

Therefore:

> UDP is connectionless. Computer Networking- A Top-Down…

---

## Why would anyone choose UDP?

Because some applications benefit from having **more direct control over when data is sent**.

TCP:

- performs congestion control
- retransmits lost data
- can delay transmission while ensuring reliable delivery

For real-time applications that:

- can tolerate some loss
- don't want retransmission delays
- need control over sending rate

UDP can be better suited. Computer Networking- A Top-Down…

UDP also:

- has no connection-establishment delay
- maintains no connection state
- has smaller header overhead

The book notes that TCP has a **20-byte header**, while UDP has an **8-byte header**. Computer Networking- A Top-Down…

---

# TCP vs UDP

| Feature | TCP | UDP |
|---|---|---|
| Connection | **Connection-oriented** | **Connectionless** |
| Handshake | Yes | No |
| Reliable delivery | **Yes** | No |
| Ordered delivery | Yes | No guarantee |
| Congestion control | **Yes** | No |
| Flow control | Yes | No |
| Throughput guarantee | No | No |
| Timing guarantee | No | No |
| Header overhead | 20 bytes | 8 bytes |
| Typical use | Web, e-mail, file transfer | DNS, loss-tolerant/time-sensitive applications |

The core distinction:

> **TCP gives reliability and connection-oriented service; UDP gives a lightweight, connectionless, best-effort service.**

---

# 4.3 IP Underneath TCP and UDP

TCP and UDP are built on top of **IP**.

IP provides:

> **Best-effort host-to-host delivery.**

IP does **not** guarantee:

- delivery
- ordering
- data integrity

Computer Networking- A Top-Down…

So the hierarchy is:

```text
Application Process
        ↓
     TCP / UDP
        ↓
        IP
        ↓
 Physical Network
```

TCP can take IP's unreliable host-to-host service and build a **reliable process-to-process service** on top of it.

This is a very important networking principle:

> **A higher layer can build a stronger service using a weaker service from the layer below.**

Computer Networking- A Top-Down…

---

# 5. Application-Layer Protocols

Now we reach the final topic.

Applications communicate by sending messages through sockets.

But the processes need to agree on:

- What messages look like
- What the fields mean
- When messages are sent
- How the receiver responds

This is the job of an **application-layer protocol**. Computer Networking- A Top-Down…

---

## Definition

> **An application-layer protocol defines how application processes running on different end systems pass messages to each other.**

Computer Networking- A Top-Down…

---

# 5.1 Four Things Defined by an Application-Layer Protocol

## 1. Types of messages

It specifies the types of messages exchanged.

For example:

```text
Client → Request
Server → Response
```

The book gives **request messages** and **response messages** as examples. Computer Networking- A Top-Down…

---

## 2. Syntax

The protocol specifies the **structure/format** of each message.

For example:

```text
Message
┌──────────┬──────────┬──────────┐
│ Field 1  │ Field 2  │ Field 3  │
└──────────┴──────────┴──────────┘
```

It defines:

- what fields exist
- how fields are arranged
- how fields are separated/delimited

Computer Networking- A Top-Down…

---

## 3. Semantics

The protocol specifies the **meaning of each field**.

For example:

```text
Field = 200
```

What does 200 mean?

The protocol must define it.

Thus:

> **Syntax = structure**  
> **Semantics = meaning**

Computer Networking- A Top-Down…

---

## 4. Rules for sending and responding

The protocol defines:

> **When and how a process sends messages and responds to messages.**

For example:

```text
Client
  │
  │ Request
  ↓
Server
  │
  │ Response
  ↓
Client
```

The protocol specifies the rules governing this exchange. Computer Networking- A Top-Down…

---

# 5.2 Application vs Application-Layer Protocol

This distinction is **very important**.

An application-layer protocol is only **one component** of a network application.

### Example: Web

The Web application includes:

- HTML
- Web browser
- Web server
- HTTP

HTTP defines the **format and sequence of messages** exchanged between browser and Web server.

Therefore:

```text
WEB APPLICATION
│
├── HTML
├── Browser
├── Web Server
└── HTTP
```

HTTP is important, but it is **not the entire Web application**. Computer Networking- A Top-Down…

---

# 5.3 Open/Public Protocols

Some application-layer protocols are specified in **RFCs** and are publicly available.

Example:

### HTTP

A browser developer can implement HTTP according to its RFC.

Then:

```text
Browser Developer A
        ↓
     HTTP rules
        ↓
Web Server Developer B
        ↓
     HTTP rules
```

The independently developed programs can communicate because they follow the same protocol.

Computer Networking- A Top-Down…

This is why standardized protocols enable **interoperability**.

---

# 5.4 Proprietary Protocols

Some application-layer protocols are **proprietary**.

Their rules are not publicly available.

The book gives **Skype** as an example of an application using proprietary application-layer protocols. Computer Networking- A Top-Down…

With a proprietary application:

```text
Developer
   ↓
Own client + own server
   ↓
Private protocol
```

An independent developer cannot necessarily create a compatible implementation because the protocol specification is not publicly available. Computer Networking- A Top-Down…

---

# 5.5 Application-Layer Protocol Examples

The book's examples include:

| Application | Application-layer protocol | Typical transport |
|---|---|---|
| Web | **HTTP** | TCP |
| E-mail | **SMTP** | TCP |
| Remote terminal access | **Telnet** | TCP |
| File transfer | **FTP** | TCP |
| DNS | **DNS protocol** | Usually UDP |
| Internet telephony | **SIP/RTP or proprietary protocols** | UDP or TCP |

The book notes that Web, e-mail, remote terminal access, and file transfer use TCP because reliable delivery is important. Internet telephony can use UDP because it can tolerate some loss and is sensitive to transmission behavior. Computer Networking- A Top-Down…

---

# 6. Putting the Entire Section Together

Suppose you open a Web page.

The entire process can be understood using all five concepts.

### Step 1 — Application architecture

The Web uses **client-server architecture**:

```text
Browser                    Web Server
Client                      Server
```

### Step 2 — Processes

The browser and Web server are **processes**.

```text
Browser Process
      ↕
Web Server Process
```

### Step 3 — Socket

The processes communicate through **sockets**:

```text
Browser Process
      ↓
    Socket
      ↓
     TCP
      ↓
   Internet
      ↓
     TCP
      ↓
    Socket
      ↓
Web Server Process
```

### Step 4 — Transport service

The Web chooses **TCP** because reliable data transfer is important.

TCP provides:

- connection-oriented service
- reliable delivery
- congestion control

### Step 5 — Application-layer protocol

The application uses **HTTP**.

HTTP defines:

- request messages
- response messages
- message format
- field meanings
- rules governing requests and responses

So:

```text
Web Application
       │
       ├── Client-server architecture
       │
       ├── Browser process
       │       ↓
       │     Socket
       │       ↓
       ├── TCP
       │       ↓
       ├── IP
       │       ↓
       └── Network
```

And at the application layer:

```text
Browser ─── HTTP request ───→ Web Server
Browser ←── HTTP response ─── Web Server
```

---

# ⭐ Final Exam Summary

## 1. Network Application Architectures

### Client-server

```text
Many clients → Always-on server
```

- centralized
- server has fixed, well-known address
- server can become bottleneck
- data centers solve large-scale server demand

### P2P

```text
Peer ↔ Peer ↔ Peer
```

- decentralized
- little/no dedicated server
- self-scalable
- cost-effective
- security, performance, and reliability challenges

### Hybrid

Combines client-server and P2P characteristics.

---

## 2. Processes Communicating

- **Process** = running program.
- Processes on different hosts communicate by exchanging messages.
- **Client** = initiates communication.
- **Server** = waits for communication.
- **Socket** = interface between application process and transport layer.
- **IP address** identifies the host.
- **Port number** identifies the application/process on that host. Computer Networking- A Top-Down…

---

## 3. Transport Services Available

Four possible dimensions:

```text
Reliable Data Transfer
Throughput
Timing
Security
```

Computer Networking- A Top-Down…

---

## 4. Transport Services Provided by Internet

### TCP

```text
Connection-oriented
Reliable
Ordered
Flow control
Congestion control
```

### UDP

```text
Connectionless
Unreliable
No ordering guarantee
No congestion control
Small overhead
```

Computer Networking- A Top-Down…

Neither TCP nor UDP provides **guaranteed throughput or timing**. Computer Networking- A Top-Down…

---

## 5. Application-Layer Protocol

Defines:

```text
1. Message types
2. Message syntax
3. Field semantics
4. Rules for sending/responding
```

Computer Networking- A Top-Down…

And remember:

> **Application ≠ Application-layer protocol.**

For example:

```text
Web Application
      ↓
Browser + Server + HTML + HTTP

HTTP = protocol
Web = complete application
```

Computer Networking- A Top-Down…

### ⭐ One-line conceptual chain

# 1. The Web and HTTP

This is **Section 2.2 of Kurose & Ross, Computer Networking: A Top-Down Approach, 7th Edition**. The section covers exactly these five topics: **HTTP, persistent/non-persistent connections, HTTP message format, cookies, and Web caching**. Computer Networking- A Top-Down…

The overall picture is:

```text
Browser / HTTP Client
        │
        │ HTTP request
        ↓
      TCP
        ↓
     Internet
        ↓
      TCP
        │
        │ HTTP request
        ↓
Web Server / HTTP Server
        │
        │ HTTP response
        ↓
     Client
```

HTTP is the **application-layer protocol of the Web**. It defines the structure of HTTP messages and how Web clients and servers exchange them. Computer Networking- A Top-Down…

---

# 1.1 HTTP

## What is HTTP?

**HTTP = HyperText Transfer Protocol.**

It is the **application-layer protocol used by the Web**.

HTTP is implemented by two programs:

- **HTTP client program** — usually the browser
- **HTTP server program** — Web server

They run on different end systems and communicate by exchanging HTTP messages. Computer Networking- A Top-Down…

### Basic communication

```text
Client                                  Server
(Browser)                               (Web Server)

   │                                         │
   │────── HTTP Request Message ────────────→│
   │                                         │
   │←───── HTTP Response Message ───────────│
   │                                         │
```

The client normally **initiates** the communication.

The server waits for requests and responds to them.

---

# 1.2 Web Pages and Objects

Before understanding HTTP, understand what the book means by a **Web page**.

A Web page consists of **objects**.

An object is a file that can be addressed by a URL.

Examples:

- HTML file
- JPEG image
- video clip
- Java applet

Computer Networking- A Top-Down…

For example, suppose a page contains:

```text
1 HTML file
5 JPEG images
```

Then the page consists of:

\[
1+5=6\text{ objects}
\]

The base HTML file contains references/URLs to the other objects.

```text
                 Web Page
                    │
          ┌─────────┴─────────┐
          ↓                   ↓
    Base HTML file       Referenced objects
                              │
                    ┌─────────┼─────────┐
                    ↓         ↓         ↓
                 image1     image2    image3 ...
```

This is important because **HTTP transfers these individual objects**.

---

# 1.3 URL

Each Web object has a **URL (Uniform Resource Locator)**.

The book explains a URL as having two important components:

1. **Hostname**
2. **Path name**

Example:

```text
http://www.someSchool.edu/someDepartment/picture.gif
```

Here:

```text
Hostname:
www.someSchool.edu

Path:
 /someDepartment/picture.gif
```

Computer Networking- A Top-Down…

The hostname identifies the server where the object resides, while the path identifies the particular object on that server.

---

# 1.4 HTTP Uses TCP

HTTP uses **TCP as its underlying transport protocol**, rather than UDP. Computer Networking- A Top-Down…

The process is:

```text
HTTP
  ↓
TCP
  ↓
IP
  ↓
Internet
```

TCP provides HTTP with **reliable data transfer**.

Therefore, HTTP does not need to worry about:

- lost TCP data
- retransmission
- reordering

Those are handled by TCP and lower layers. Computer Networking- A Top-Down…

---

# 1.5 HTTP Is Stateless

This is a **very important property**.

HTTP is a **stateless protocol**.

That means:

> The HTTP server does not maintain information about previous requests made by a particular client.

Suppose a browser requests the same object twice:

```text
Request 1 → Server
Response 1 ← Server

Request 2 → Server
Response 2 ← Server
```

The server does not inherently remember that it already sent the object to that client.

It simply sends the object again. Computer Networking- A Top-Down…

### Why is statelessness useful?

It simplifies server design and allows Web servers to handle very large numbers of simultaneous TCP connections. Computer Networking- A Top-Down…

But applications often **need state**.

That is why HTTP uses mechanisms such as **cookies**, which we will discuss below.

---

# 2. Persistent vs Non-Persistent Connections

This is one of the most important topics in HTTP.

The question is:

> If a client requests many objects from the same server, should each object use a separate TCP connection, or should multiple objects use the same TCP connection?

There are two possibilities:

1. **Non-persistent HTTP**
2. **Persistent HTTP**

Computer Networking- A Top-Down…

---

# 2.1 Non-Persistent Connections

With a **non-persistent connection**:

> One TCP connection is used for exactly **one request/response pair**.

Then the connection is closed.

```text
TCP Connection 1
Client ── Request ──→ Server
Client ←─ Response ─ Server
Connection closes

TCP Connection 2
Client ── Request ──→ Server
Client ←─ Response ─ Server
Connection closes
```

---

## Example

Suppose a Web page contains:

```text
1 HTML file
10 JPEG images
```

Total:

\[
11\text{ objects}
\]

With non-persistent HTTP:

```text
HTML       → TCP connection 1
JPEG 1     → TCP connection 2
JPEG 2     → TCP connection 3
...
JPEG 10    → TCP connection 11
```

Therefore:

\[
\boxed{11\text{ TCP connections}}
\]

The book explicitly gives this example. Computer Networking- A Top-Down…

---

## Steps for one object

For each object:

### Step 1 — Establish TCP connection

Client initiates a TCP connection to the server, normally on port **80** for HTTP.

### Step 2 — Send HTTP request

Client sends an HTTP request through its socket.

### Step 3 — Server sends response

Server retrieves the object, places it in an HTTP response, and sends it to the client.

### Step 4 — Close TCP connection

The server tells TCP to close the connection after the response.

### Step 5 — Client receives object

The client extracts the object from the HTTP response.

### Step 6 — Repeat

For every additional object, a new TCP connection is created. Computer Networking- A Top-Down…

---

# 2.2 Delay of Non-Persistent HTTP

The book uses **RTT (Round-Trip Time)**.

### RTT

RTT is the time required for a small packet to travel:

```text
Client → Server → Client
```

It includes:

- propagation delays
- queuing delays
- processing delays

at intermediate routers and switches. Computer Networking- A Top-Down…

For one object using non-persistent HTTP:

```text
1 RTT → TCP connection establishment
1 RTT → HTTP request + response
```

Therefore, approximately:

\[
\boxed{2RTT + \text{object transmission time}}
\]

The book explicitly gives this approximation. Computer Networking- A Top-Down…

### ⭐ Exam point

For a small object where transmission time is negligible:

\[
\boxed{\text{Response time} \approx 2RTT}
\]

---

# 2.3 Problems with Non-Persistent Connections

There are two major problems.

### 1. Connection overhead

A new TCP connection must be established for every object.

This requires TCP state and buffers at both ends. Computer Networking- A Top-Down…

### 2. Additional delay

Each object suffers approximately:

\[
2RTT
\]

of delay before considering object transmission time. Computer Networking- A Top-Down…

---

# 2.4 Persistent Connections

With a **persistent connection**:

> The TCP connection remains open after the server sends a response.

Subsequent requests and responses can use the **same TCP connection**. Computer Networking- A Top-Down…

```text
One TCP connection
─────────────────────────────────

Client ── Request 1 ─────────────→ Server
Client ←─ Response 1 ───────────── Server

Client ── Request 2 ─────────────→ Server
Client ←─ Response 2 ───────────── Server

Client ── Request 3 ─────────────→ Server
Client ←─ Response 3 ───────────── Server

             Connection remains open
```

Thus, the base HTML file and referenced images can be transferred over the same TCP connection. Computer Networking- A Top-Down…

---

# 2.5 Persistent HTTP and Pipelining

HTTP/1.1 persistent connections support **pipelining**.

The client can send multiple requests:

```text
Request 1 ─────→
Request 2 ─────→
Request 3 ─────→
```

without waiting for each response before sending the next request.

The server can then send the corresponding objects back-to-back. Computer Networking- A Top-Down…

The book notes that HTTP/1.1's default mode uses persistent connections with pipelining. It also notes that HTTP/2 later allowed multiple requests and responses to be **interleaved** over a connection. Computer Networking- A Top-Down…

---

# 2.6 Persistent vs Non-Persistent

| Feature | Non-Persistent | Persistent |
|---|---|---|
| TCP connection | New connection for each object | Same connection reused |
| Requests per connection | One | Multiple |
| Connection closed after response? | Yes | No, normally |
| TCP setup overhead | High | Lower |
| Delay | Higher | Lower |
| Server resource overhead | Higher | Lower |
| Multiple objects over one connection | ❌ | ✅ |
| Pipelining | ❌ | ✅ in HTTP/1.1 |

### ⭐ Most important difference

> **Non-persistent: one TCP connection per request/response.**  
> **Persistent: multiple request/response exchanges over the same TCP connection.**

---

# 3. HTTP Message Format

HTTP defines **two types of messages**:

1. **HTTP request message**
2. **HTTP response message** Computer Networking- A Top-Down…

---

# 3.1 HTTP Request Message

A typical request from the book:

```text
GET /somedir/page.html HTTP/1.1
Host: www.someschool.edu
Connection: close
User-agent: Mozilla/5.0
Accept-language: fr

```

Computer Networking- A Top-Down…

The request is composed of:

```text
Request line
Header lines
Blank line
[Optional entity body]
```

---

## Request Line

The first line is the **request line**.

It contains three fields:

```text
Method    URL    HTTP version
```

Example:

```text
GET /somedir/page.html HTTP/1.1
```

Therefore:

```text
Method:
GET

URL/path:
 /somedir/page.html

Version:
HTTP/1.1
```

Computer Networking- A Top-Down…

---

# 3.2 HTTP Methods

The book mentions several methods:

- **GET**
- **POST**
- **HEAD**
- **PUT**
- **DELETE**

The majority of HTTP request messages use **GET**.

GET is used when the browser requests an object. Computer Networking- A Top-Down…

### GET

```text
GET /page.html HTTP/1.1
```

Means approximately:

> Give me this object.

---

# 3.3 Header Lines

After the request line come **header lines**.

Example:

```text
Host: www.someschool.edu
Connection: close
User-agent: Mozilla/5.0
Accept-language: fr
```

The book explains these as follows. Computer Networking- A Top-Down…

### Host

```text
Host: www.someschool.edu
```

Specifies the host where the object resides.

The Host header is particularly important for **Web proxy caches**.

### Connection

```text
Connection: close
```

Tells the server that the client does not want a persistent connection.

### User-agent

```text
User-agent: Mozilla/5.0
```

Identifies the browser/client software.

### Accept-language

```text
Accept-language: fr
```

Indicates the user's preferred language.

---

# 3.4 Blank Line

After the header lines comes a **blank line**.

It indicates the end of the header section.

```text
GET /page.html HTTP/1.1
Host: example.com
Connection: close
User-agent: Mozilla/5.0

← blank line
```

The book notes that each line ends with carriage return + line feed, and the final header line is followed by an additional carriage return + line feed. Computer Networking- A Top-Down…

---

# 3.5 HTTP Request Structure

For exam purposes:

```text
┌─────────────────────────────────────┐
│ Request Line                         │
│ Method + URL + HTTP Version          │
├─────────────────────────────────────┤
│ Header Line                          │
│ Header Line                          │
│ Header Line                          │
├─────────────────────────────────────┤
│ Blank Line                           │
├─────────────────────────────────────┤
│ Entity Body (if present)             │
└─────────────────────────────────────┘
```

---

# 3.6 HTTP Response Message

A typical response:

```text
HTTP/1.1 200 OK
Connection: close
Date: Tue, 18 Aug 2015 15:44:04 GMT
Server: Apache/2.2.3
Last-Modified: Tue, 18 Aug 2015 15:11:03 GMT
Content-Length: 6821
Content-Type: text/html

(data data data...)
```

Computer Networking- A Top-Down…

It contains three major parts:

1. **Status line**
2. **Header lines**
3. **Entity body**

Computer Networking- A Top-Down…

---

# 3.7 Status Line

The status line has:

```text
HTTP version
Status code
Status phrase
```

Example:

```text
HTTP/1.1 200 OK
```

Therefore:

```text
Version:
HTTP/1.1

Status code:
200

Phrase:
OK
```

Computer Networking- A Top-Down…

---

# 3.8 Important HTTP Status Codes

### 200 OK

Request succeeded and requested information is returned.

### 301 Moved Permanently

Requested object has permanently moved.

The new URL is specified using the **Location** header.

### 400 Bad Request

The request could not be understood by the server.

### 404 Not Found

Requested document does not exist on the server.

### 505 HTTP Version Not Supported

The requested HTTP version is not supported by the server.

Computer Networking- A Top-Down…

### ⭐ Exam table

| Code | Meaning |
|---|---|
| **200** | OK / successful |
| **301** | Moved Permanently |
| **400** | Bad Request |
| **404** | Not Found |
| **505** | HTTP Version Not Supported |

---

# 3.9 Response Header Lines

Some important headers:

### Date

Indicates when the HTTP response was created and sent by the server.

### Server

Identifies the Web server software.

### Last-Modified

Indicates when the requested object was created or last modified.

This is important for **Web caching**.

### Content-Length

Indicates the number of bytes in the object.

### Content-Type

Indicates the type of object in the entity body.

For example:

```text
Content-Type: text/html
```

The book emphasizes that the object type is officially indicated by **Content-Type**, not merely by the file extension. Computer Networking- A Top-Down…

---

# 3.10 Entity Body

The **entity body** contains the actual object being transferred.

For example:

```text
HTTP/1.1 200 OK
Content-Type: text/html

<html>
   ...
</html>
```

The HTML data is the entity body.

For an image response, the entity body contains the image data.

---

# 4. Cookies

Now we encounter an apparent contradiction:

> HTTP is stateless, but websites often need to remember users.

For example, a website may need to remember:

- a user's identity
- shopping cart
- login session
- previous activity

HTTP uses **cookies** for this purpose. Computer Networking- A Top-Down…

---

# 4.1 What is a Cookie?

A cookie allows a website to **keep track of users**.

The book identifies **four components** of cookie technology:

1. Cookie header line in the HTTP response
2. Cookie header line in the HTTP request
3. Cookie file on the user's end system, managed by the browser
4. Back-end database at the website Computer Networking- A Top-Down…

---

# 4.2 How Cookies Work

Consider a user visiting a website for the first time.

### Step 1 — Client requests page

```text
Browser ─── HTTP request ───→ Server
```

### Step 2 — Server creates an ID

Suppose the server creates:

```text
ID = 1678
```

It stores this ID in its back-end database.

### Step 3 — Server sends Set-cookie

The server includes:

```text
Set-cookie: 1678
```

in the HTTP response. Computer Networking- A Top-Down…

### Step 4 — Browser stores cookie

The browser stores the server's hostname and cookie ID in its cookie file.

```text
Browser cookie file

Amazon → 1678
```

### Step 5 — Future requests include Cookie

When the browser requests another page from that server:

```text
Cookie: 1678
```

is included in the HTTP request.

Computer Networking- A Top-Down…

---

# 4.3 Complete Cookie Flow

```text
          First visit
Browser ───────────────→ Server
        HTTP request

Browser ←─────────────── Server
        Set-cookie: 1678

Browser
┌─────────────────┐
│ Cookie file     │
│ Server → 1678   │
└─────────────────┘


          Later visit
Browser ───────────────→ Server
        Cookie: 1678

Server
   ↓
Looks up 1678
   ↓
Identifies user's session/activity
```

The server can associate the cookie ID with information stored in its back-end database. Computer Networking- A Top-Down…

---

# 4.4 What Can Cookies Be Used For?

The book gives examples such as:

- identifying users
- maintaining shopping carts
- maintaining user sessions
- providing one-click shopping
- remembering user activity/preferences

Cookies effectively create a **user-session layer on top of stateless HTTP**. Computer Networking- A Top-Down…

---

# 4.5 Cookies and Privacy

Cookies can also raise **privacy concerns**.

By combining:

```text
Cookie ID
+
User account information
+
Browsing activity
```

a website may learn substantial information about a user.

The book therefore notes the controversy surrounding cookies and their potential privacy implications. Computer Networking- A Top-Down…

---

# 5. Web Caching

## What is a Web Cache?

A **Web cache**, also called a **proxy server**, is a network entity that satisfies HTTP requests on behalf of an origin Web server.

It stores copies of recently requested objects. Computer Networking- A Top-Down…

```text
                ┌───────────────┐
                │ Origin Server │
                └───────▲───────┘
                        │
                   Internet
                        │
                 ┌──────▼──────┐
                 │  Web Cache  │
                 └──────▲──────┘
                        │
                      LAN
                        │
                 ┌──────▼──────┐
                 │   Browser   │
                 └─────────────┘
```

---

# 5.1 How Web Caching Works

Suppose the browser requests:

```text
http://www.someschool.edu/campus.gif
```

### Step 1

Browser establishes a TCP connection to the cache.

```text
Browser → Cache
```

### Step 2

Browser sends an HTTP request to the cache.

### Step 3 — Cache checks storage

The cache asks:

> Do I already have this object?

---

## Cache Hit

If the object is already stored:

```text
Browser → Cache
             │
             ├── Object exists
             ↓
Browser ← Cache
```

The cache immediately sends the object back.

No communication with the origin server is necessary. Computer Networking- A Top-Down…

---

## Cache Miss

If the cache does not have the object:

```text
Browser → Cache
             │
             │ Cache miss
             ↓
          Origin Server
```

The cache:

1. Opens a TCP connection to the origin server.
2. Sends an HTTP request.
3. Receives the object.
4. Stores a copy locally.
5. Sends the object to the browser. Computer Networking- A Top-Down…

---

# 5.2 A Cache Is Both Client and Server

This is an important conceptual point.

From the browser's perspective:

```text
Browser → Cache
```

The cache acts as a **server**.

From the origin server's perspective:

```text
Cache → Origin Server
```

The cache acts as a **client**.

Therefore:

> **A Web cache is simultaneously a server and a client.** Computer Networking- A Top-Down…

---

# 5.3 Why Use Web Caching?

The book gives two major reasons.

## 1. Reduce response time

If the cache is geographically/network-topologically close to the client:

```text
Browser → nearby cache
```

is much faster than:

```text
Browser → Internet → distant origin server
```

especially when the link to the origin server is a bottleneck. Computer Networking- A Top-Down…

---

## 2. Reduce Internet traffic

If many users request the same object:

```text
User 1 ──┐
User 2 ──┼──→ Cache ──→ Internet
User 3 ──┤
User 4 ──┘
```

the cache can serve repeated requests locally.

This reduces traffic on the institution's Internet access link.

That can reduce the need for expensive bandwidth upgrades. Computer Networking- A Top-Down…

---

# 5.4 Cache Hit Rate

The **hit rate** is the fraction of requests that can be satisfied by the cache.

For example:

\[
\text{Hit rate}=0.4
\]

means:

\[
40\%
\]

of requests are served directly by the cache.

Therefore:

\[
\text{Miss rate}=1-0.4=0.6
\]

or:

\[
60\%
\]

must go to origin servers.

The book uses exactly such an example and shows that reducing the fraction of traffic crossing the access link can dramatically reduce delay. Computer Networking- A Top-Down…

---

# 5.5 The Cache Consistency Problem

Caching creates a new problem:

> **What if the cached copy is old?**

Suppose:

```text
Origin server:
Version 2

Cache:
Version 1
```

The cache contains a **stale object**.

The book solves this with the **conditional GET** mechanism. Computer Networking- A Top-Down…

---

# 5.6 Conditional GET

An HTTP request is a **conditional GET** when:

1. It uses the **GET** method, and
2. It contains an **If-Modified-Since** header.

Computer Networking- A Top-Down…

Example:

```text
GET /fruit/kiwi.gif HTTP/1.1
Host: www.exotiquecuisine.com
If-modified-since: Wed, 9 Sep 2015 09:23:24
```

The date corresponds to the object's previous **Last-Modified** value. Computer Networking- A Top-Down…

The meaning is:

> "Send me the object only if it has been modified since this date."

---

## If object has NOT changed

Server responds:

```text
HTTP/1.1 304 Not Modified
```

with an empty entity body.

The cache then uses its existing copy rather than downloading the object again. Computer Networking- A Top-Down…

So:

```text
Cache ── Conditional GET ──→ Server

Server:
"No change"

       ↓

304 Not Modified

       ↓

Cache uses existing copy
```

This saves:

- bandwidth
- transmission time
- response time

---

## If object HAS changed

The server sends the updated object.

So the cache replaces its old copy with the new version.

---

# 5.7 Web Caching Summary

```text
                  HTTP Request
Browser ─────────────────────────→ Cache
                                     │
                              ┌──────┴──────┐
                              │             │
                           HIT           MISS
                              │             │
                              ↓             ↓
                         Return copy    Request origin
                                            │
                                            ↓
                                      Origin Server
                                            │
                                            ↓
                                      New object
                                            │
                              ┌─────────────┘
                              ↓
                         Store in cache
                              │
                              ↓
                         Send to browser
```

---

# 6. Cookies vs Web Caching

These two are often confused, so keep them separate.

| Feature | Cookies | Web Cache |
|---|---|---|
| Main purpose | Maintain user state/identify users | Store copies of Web objects |
| Solves | HTTP statelessness | Delay and bandwidth problems |
| Stored at | Browser + server database | Cache storage |
| Uses HTTP headers | Yes | Yes |
| Example | Shopping cart/session | Cached image/HTML |
| Main benefit | User/session tracking | Faster delivery + reduced traffic |

### Key distinction

> **Cookies remember the user/session.**  
> **Caches remember copies of objects.**

---

# 7. Complete HTTP Operation

Let's combine everything.

Suppose you enter:

```text
http://www.example.com/index.html
```

### Step 1 — Browser identifies the server

The URL contains:

```text
Hostname: www.example.com
Path: /index.html
```

### Step 2 — Browser establishes TCP connection

```text
Browser ───── TCP ─────→ Server/Cache
```

### Step 3 — Browser sends HTTP request

```text
GET /index.html HTTP/1.1
Host: www.example.com
```

### Step 4 — Cache checks its storage

If there is a valid cached copy:

```text
Cache → Browser
```

Otherwise:

```text
Cache → Origin Server
```

### Step 5 — Server returns HTTP response

```text
HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: ...
...
HTML DATA
```

### Step 6 — Browser receives HTML

The HTML may contain references to other objects:

```text
HTML
 ├── image.jpg
 ├── logo.png
 ├── video...
 └── stylesheet...
```

The browser sends additional HTTP requests for those objects.

### Step 7 — Cookies may be included

If the browser has a cookie for the server:

```text
Cookie: 1678
```

may be included in the request.

### Step 8 — Persistent TCP connection may be reused

Instead of creating a new TCP connection for every object, persistent HTTP can reuse the same connection.

---

# ⭐ Final Exam-Oriented Summary

## HTTP

- HTTP = **HyperText Transfer Protocol**
- Application-layer protocol of the Web.
- Client and server exchange **HTTP request and response messages**.
- HTTP uses **TCP**.
- HTTP itself is **stateless**. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

---

## Non-Persistent HTTP

```text
1 request
     ↓
1 TCP connection
     ↓
1 response
     ↓
connection closes
```

For each object:

\[
\boxed{\approx 2RTT+\text{transmission time}}
\]

A page containing 11 objects can require **11 TCP connections** if all use non-persistent HTTP. Computer Networking- A Top-Down…

---

## Persistent HTTP

```text
One TCP connection
       ↓
Request 1 → Response 1
Request 2 → Response 2
Request 3 → Response 3
       ↓
Connection remains open
```

Advantages:

- fewer TCP connection establishments
- lower delay
- lower connection-management overhead
- multiple objects can use the same connection

HTTP/1.1 supports persistent connections and pipelining. Computer Networking- A Top-Down…

---

## HTTP Request

```text
Request line
Header lines
Blank line
Entity body (if any)
```

Request line:

\[
\boxed{\text{Method + URL + HTTP version}}
\]

Example:

```text
GET /page.html HTTP/1.1
```

Computer Networking- A Top-Down…

---

## HTTP Response

```text
Status line
Header lines
Blank line
Entity body
```

Status line:

\[
\boxed{\text{HTTP version + status code + phrase}}
\]

Example:

```text
HTTP/1.1 200 OK
```

Computer Networking- A Top-Down…

---

## Important Status Codes

```text
200 → OK
301 → Moved Permanently
400 → Bad Request
404 → Not Found
505 → HTTP Version Not Supported
```

Computer Networking- A Top-Down…

---

## Cookies

Four components:

```text
1. Set-cookie response header
2. Cookie request header
3. Browser cookie file
4. Server back-end database
```

Cookies create a **user-session layer on top of stateless HTTP**. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

---

## Web Cache

A Web cache/proxy:

```text
Browser → Cache → Origin Server
```

### Cache hit

```text
Browser → Cache
         ← Object
```

### Cache miss

```text
Browser → Cache → Origin Server
                    ↓
                  Object
                    ↓
Browser ← Cache ←───┘
```

Main benefits:

1. **Reduce response time**
2. **Reduce traffic on Internet access links**
3. **Reduce bandwidth-upgrade costs** Computer Networking- A Top-Down…

---

## Conditional GET

Used to determine whether a cached object is still current.

```text
GET + If-Modified-Since
            ↓
       Origin Server
        ↙         ↘
Not modified     Modified
    ↓                ↓
304 response      New object
```

**304 Not Modified** tells the cache that its existing copy can be used. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

### ⭐ One-line memory aid

> **HTTP transfers Web objects; persistent connections reduce connection overhead; HTTP messages define requests/responses; cookies add user state to stateless HTTP; and Web caches store copies of objects to reduce delay and traffic.**
# 2.3 Electronic Mail in the Internet

According to **Kurose & Ross, *Computer Networking: A Top-Down Approach*, 7th Edition**, Internet e-mail has three major components:

1. **User agents** — used to compose, read, reply to, forward, and save messages.
2. **Mail servers** — maintain users' mailboxes and outgoing message queues.
3. **SMTP** — transfers mail between mail servers and also from a user's mail client to the user's mail server. Computer Networking- A Top-Down…

A typical e-mail path is:

```text
Alice's User Agent
       |
      SMTP
       ↓
Alice's Mail Server
       |
      SMTP
       ↓
Bob's Mail Server
       |
   POP3 / IMAP / HTTP
       ↓
Bob's User Agent
```

The important point is that **SMTP is used for sending/pushing mail**, while **POP3, IMAP, or HTTP are used by the recipient to access/pull mail**. Computer Networking- A Top-Down…

---

# 1. SMTP — Simple Mail Transfer Protocol

## What is SMTP?

**SMTP (Simple Mail Transfer Protocol)** is the principal application-layer protocol for Internet electronic mail. It is defined in **RFC 5321** and uses **TCP** for reliable data transfer. Computer Networking- A Top-Down…

SMTP operates between mail servers:

- Sending mail server → **SMTP client**
- Receiving mail server → **SMTP server**

A mail server can therefore act as an SMTP client when sending mail and as an SMTP server when receiving mail. Computer Networking- A Top-Down…

### Why mail servers are needed

Suppose Alice wants to send an e-mail to Bob.

Alice's computer does **not normally communicate directly with Bob's personal computer**. Instead:

1. Alice composes the message using her user agent.
2. Her user agent sends it to **Alice's mail server**.
3. Alice's mail server puts it in its outgoing queue.
4. The SMTP client on Alice's mail server establishes a TCP connection with Bob's mail server.
5. SMTP transfers the message.
6. Bob's mail server places the message in Bob's mailbox.
7. Bob later accesses the mailbox using a mail-access protocol. Computer Networking- A Top-Down…

### Why this two-step procedure?

Suppose Bob's mail server is temporarily down.

If Alice's user agent had to communicate directly with Bob's server, it would have no convenient place to keep trying.

Instead, Alice's mail server stores the message in its **message queue** and can retry delivery later. Computer Networking- A Top-Down…

---

## SMTP uses TCP

The sending SMTP client establishes a **TCP connection to port 25** of the receiving SMTP server.

After TCP connection establishment, the SMTP client and server perform application-layer handshaking and then transfer the message. Computer Networking- A Top-Down…

### Basic SMTP dialogue

The book gives an example involving these commands:

```text
HELO
MAIL FROM
RCPT TO
DATA
QUIT
```

For example:

```text
C: MAIL FROM: <alice@crepes.fr>
S: 250 alice@crepes.fr ... Sender ok

C: RCPT TO: <bob@hamburger.edu>
S: 250 bob@hamburger.edu ... Recipient ok

C: DATA
S: 354 Enter mail, end with "." on a line by itself

C: Do you like ketchup?
C: How about pickles?
C: .

S: 250 Message accepted for delivery

C: QUIT
S: 221 hamburger.edu closing connection
```

Here:

- `MAIL FROM` identifies the sender.
- `RCPT TO` identifies the recipient.
- `DATA` indicates that the actual mail message follows.
- A line containing only `.` marks the end of the message.
- `QUIT` terminates the SMTP session. Computer Networking- A Top-Down…

### Important: SMTP persistent connections

SMTP uses **persistent TCP connections**.

If one sending mail server has several messages for the same receiving mail server, it can send them all through the same TCP connection rather than establishing a new connection for every message. Computer Networking- A Top-Down…

---

# 2. HTTP vs SMTP

Both HTTP and SMTP transfer information from one host to another, and both can use persistent TCP connections. But their operation is significantly different. Computer Networking- A Top-Down…

| Feature | HTTP | SMTP |
|---|---|---|
| Full form | HyperText Transfer Protocol | Simple Mail Transfer Protocol |
| Main use | Web objects | E-mail messages |
| Communication style | Mainly **pull** | Primarily **push** |
| Who initiates TCP connection? | Machine wanting to **receive** data | Machine wanting to **send** mail |
| Data restriction | No 7-bit ASCII restriction | Message, including body, restricted to **7-bit ASCII** |
| Multiple objects | Each object has its own HTTP response | All objects are placed into one mail message |
| Persistent connections | Supported | Supported |

### A. Pull vs Push

This is the **most important difference**.

### HTTP → Pull

A Web client wants an object from a Web server.

```text
Client ───── request ─────> Server
Client <──── object ─────── Server
```

The client initiates the TCP connection because it **wants to receive** the object.

### SMTP → Push

A sending mail server wants to deliver an e-mail.

```text
Sending Mail Server ───── message ─────> Receiving Mail Server
```

The sending side initiates the TCP connection because it **wants to send/push** the message. Computer Networking- A Top-Down…

---

## B. 7-bit ASCII restriction

SMTP requires the **entire mail message, including its body, to be in 7-bit ASCII**.

Therefore, if an e-mail contains:

- an image,
- audio,
- video,
- or characters outside 7-bit ASCII,

the data must first be encoded into ASCII before SMTP transmission and decoded afterward. Computer Networking- A Top-Down…

HTTP does **not** impose this restriction.

---

## C. Handling multiple objects

Suppose an e-mail contains:

```text
Text + Image + Audio
```

SMTP places these objects into **one mail message**.

HTTP handles a Web document differently: individual objects are transferred in separate HTTP response messages. Computer Networking- A Top-Down…

---

# 3. Mail Message Formats

This is different from the **SMTP commands** we just studied.

An SMTP session has commands such as:

```text
MAIL FROM
RCPT TO
DATA
```

But the **mail message itself** contains its own headers and body.

The book describes the mail message format using **RFC 5322**. Computer Networking- A Top-Down…

## Structure

```text
Mail Message
│
├── Header lines
│
├── Blank line
│
└── Message body
```

The header and body are separated by a **blank line (CRLF)**.

### Example

```text
From: alice@crepes.fr
To: bob@hamburger.edu
Subject: Searching for the meaning of life.

Dear Bob,

Here is the message body.
```

The book's example similarly contains `From`, `To`, and `Subject` header lines followed by a blank line and then the body. Computer Networking- A Top-Down…

---

## Mail headers

Each header has the general form:

```text
keyword: value
```

For example:

```text
From: alice@crepes.fr
To: bob@hamburger.edu
Subject: Searching for the meaning of life.
```

### Important headers

- **From:** sender
- **To:** recipient
- **Subject:** subject of the message

The book states that every message must have `From:` and `To:` headers; `Subject:` is optional. Computer Networking- A Top-Down…

---

## Very important distinction: SMTP commands vs mail headers

Do **not** confuse:

```text
MAIL FROM:
RCPT TO:
```

with:

```text
From:
To:
```

They serve different purposes.

| SMTP command | Mail header |
|---|---|
| `MAIL FROM:` | `From:` |
| `RCPT TO:` | `To:` |
| Part of SMTP **conversation** | Part of the actual **mail message** |
| Used during SMTP handshaking | Stored as message information |

The book explicitly emphasizes this distinction. Computer Networking- A Top-Down…

---

# 4. Mail Access Protocols

Now we have a problem:

SMTP delivers the message to **Bob's mail server**, but Bob's user agent is running on his local computer.

How does Bob retrieve the message?

SMTP cannot normally be used for this because:

- SMTP is a **push** protocol.
- Bob needs to **pull** messages from his mail server.

Therefore, a **mail access protocol** is used.

The book discusses:

1. **POP3**
2. **IMAP**
3. **HTTP** for Web-based e-mail. Computer Networking- A Top-Down…

---

# 4.1 POP3

**POP3 = Post Office Protocol Version 3**

It is a relatively simple mail-access protocol defined in **RFC 1939**.

The user agent opens a **TCP connection to port 110** on the mail server. Computer Networking- A Top-Down…

POP3 has **three phases**:

```text
1. Authorization
        ↓
2. Transaction
        ↓
3. Update
```

## Phase 1: Authorization

The user agent authenticates itself using:

```text
USER <username>
PASS <password>
```

Example:

```text
C: user bob
S: +OK

C: pass hungry
S: +OK user successfully logged on
```

The book notes that the username and password are sent **in the clear** in this basic POP3 description. Computer Networking- A Top-Down…

---

## Phase 2: Transaction

The user agent can:

- retrieve messages,
- mark messages for deletion,
- remove deletion marks,
- obtain mail statistics.

Important commands include:

```text
LIST
RETR
DELE
QUIT
```

For example:

```text
C: list
S: 1 498
S: 2 912

C: retr 1
S: <message 1>

C: dele 1

C: retr 2
S: <message 2>

C: dele 2
```

Computer Networking- A Top-Down… Computer Networking- A Top-Down…

---

## Download-and-delete vs Download-and-keep

POP3 can commonly be configured in two ways.

### Download-and-delete

```text
Mail Server
   |
   | download
   ↓
Local Computer
   |
   X message deleted from server
```

Problem: if Bob reads the message on his office PC and then wants to read it from his home computer, the message may no longer be on the server.

### Download-and-keep

```text
Mail Server
   |
   | download
   ↓
Local Computer

Message remains on server
```

Now Bob can access the message from multiple computers. Computer Networking- A Top-Down…

---

# 4.2 IMAP

**IMAP = Internet Mail Access Protocol**

IMAP is also a mail-access protocol, but it provides **much more functionality than POP3** and is consequently more complex. It is defined in **RFC 3501**. Computer Networking- A Top-Down…

The key idea is:

> **With IMAP, the mailbox and its organization remain on the mail server.**

An IMAP server associates messages with folders.

Initially:

```text
INBOX
 ├── Message 1
 ├── Message 2
 └── Message 3
```

The user can create folders:

```text
INBOX
 ├── Message 1
 └── Message 2

College
 └── Message 3
```

IMAP provides commands to:

- create folders,
- move messages between folders,
- delete messages,
- search remote folders,
- maintain mailbox state across sessions. Computer Networking- A Top-Down…

### Why IMAP is useful

Suppose you use:

- laptop,
- desktop,
- phone.

Because the mailbox organization is maintained on the server, all devices can access the same organized mailbox.

This is a major difference from the simple POP3 download-and-delete model.

---

## IMAP can retrieve only parts of a message

Another important IMAP feature is that the user agent can request only specific components, such as:

- message headers,
- one part of a multipart MIME message.

This is useful when the connection has low bandwidth because the user doesn't have to download a huge attachment or entire message unnecessarily. Computer Networking- A Top-Down…

---

# 4.3 Web-Based E-Mail

With Web-based e-mail, the user's **Web browser acts as the user agent**.

Instead of using POP3 or IMAP directly:

```text
Browser ←──── HTTP ────→ Mail Server
```

When receiving:

```text
Bob's Mail Server
        |
       HTTP
        ↓
   Bob's Browser
```

When sending:

```text
Alice's Browser
        |
       HTTP
        ↓
Alice's Mail Server
```

However, when Alice's mail server communicates with another mail server, it still uses **SMTP**. Computer Networking- A Top-Down…

So:

```text
                 SMTP
Alice Server ─────────────→ Bob Server
     ↑                         ↓
    HTTP                      HTTP
     ↑                         ↓
Alice Browser              Bob Browser
```

---

# POP3 vs IMAP vs HTTP

| Feature | POP3 | IMAP | HTTP/Web mail |
|---|---|---|---|
| Purpose | Mail access | Mail access | Web-based mail access |
| Basic model | Download messages | Manage remote mailbox | Browser accesses mailbox |
| Mail organization | Mainly local after download | **Remote/server-side** | Server-side through Web application |
| Remote folders | No | **Yes** | Yes through Web interface |
| Search remote mailbox | Limited | **Yes** | Yes through Web interface |
| Access message parts | Limited | **Yes** | Through Web interface |
| Multiple-device use | Less convenient in download-delete mode | **Well suited** | **Well suited** |
| Protocol used to browser/client | POP3 | IMAP | HTTP |

The book's central distinction is that **POP3 is simple and limited**, whereas **IMAP maintains a richer remote mailbox state**, and Web-based e-mail uses **HTTP between the browser and mail server**. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

---

# Complete E-Mail Flow — Exam View

Suppose **Alice sends an e-mail to Bob**:

### Step 1 — Composition

```text
Alice → User Agent
```

Alice writes the message.

### Step 2 — Alice's user agent → Alice's mail server

SMTP is used to push the message to Alice's mail server.

### Step 3 — Alice's mail server → Bob's mail server

SMTP establishes a TCP connection and transfers the e-mail.

```text
Alice Mail Server
       |
      SMTP
       ↓
Bob Mail Server
```

### Step 4 — Message enters Bob's mailbox

Bob's mail server stores the message in his mailbox.

### Step 5 — Bob retrieves the message

Using one of:

```text
POP3
IMAP
HTTP
```

depending on the mail-access method. Computer Networking- A Top-Down…

---

## ⭐ Most Important Exam Points

1. **SMTP is the principal application-layer protocol for Internet e-mail.**
2. **SMTP uses TCP.**
3. SMTP is primarily a **push** protocol.
4. HTTP is mainly a **pull** protocol.
5. SMTP transfers mail **from sending mail server to receiving mail server**.
6. **POP3/IMAP/HTTP** provide ways for the recipient to access mail.
7. SMTP uses **7-bit ASCII** for the mail message; binary/multimedia data must be encoded.
8. Mail message = **headers + blank line + body**.
9. `MAIL FROM` and `RCPT TO` are **SMTP commands**, whereas `From:` and `To:` are **mail headers**.
10. POP3 has **authorization → transaction → update** phases.
11. POP3 supports **download-and-delete** and **download-and-keep**.
12. IMAP maintains the **remote mailbox state and folders on the server**.
13. Web-based e-mail uses **HTTP between browser and mail server**, while mail servers still communicate with each other using SMTP. Computer Networking- A Top-Down… Computer Networking- A Top-Down… Computer Networking- A Top-Down…
# 2.4 DNS — Domain Name System

According to **Kurose & Ross, *Computer Networking: A Top-Down Approach*, 7th Edition**, DNS is the Internet's **directory service**.

The central problem is:

- Humans prefer names such as `www.example.com`.
- Routers and network communication ultimately need **IP addresses**.

DNS provides the translation between these two. It is both:

1. a **distributed database** implemented by a hierarchy of DNS servers, and
2. an **application-layer protocol** that allows hosts to query that database.

The book describes DNS as an application-layer protocol using **UDP port 53**. Computer Networking- A Top-Down…

---

# 1. DNS Services

DNS does more than simply convert a hostname into an IP address. The book identifies several important services.

## 1.1 Hostname-to-IP Address Translation

This is the **main service of DNS**.

For example:

```text
Hostname:
www.someschool.edu

        ↓ DNS

IP address:
128.119.245.12
```

Applications such as browsers and mail readers use DNS to obtain the IP address before communicating with the destination host. Computer Networking- A Top-Down…

### Example: accessing a Web page

Suppose the browser wants:

```text
http://www.someschool.edu/index.html
```

The browser cannot directly establish the HTTP connection using only the hostname.

The sequence is:

```text
1. Browser extracts hostname
          ↓
2. DNS client is invoked
          ↓
3. DNS query sent for www.someschool.edu
          ↓
4. DNS reply returns IP address
          ↓
5. Browser establishes TCP connection
          ↓
6. HTTP request is sent
```

The DNS lookup therefore occurs **before** the browser can establish the TCP connection to the Web server. Computer Networking- A Top-Down…

---

# 1.2 Host Aliasing

A host can have:

- a complicated **canonical hostname**, and
- one or more easier **alias hostnames**.

For example:

```text
Canonical hostname:
relay1.west-coast.enterprise.com

Aliases:
enterprise.com
www.enterprise.com
```

The alias is easier for users to remember.

DNS can provide the canonical hostname corresponding to an alias and its associated IP address. Computer Networking- A Top-Down…

---

# 1.3 Mail Server Aliasing

DNS can also provide aliases for **mail servers**.

Users want simple e-mail addresses, while the actual mail server may have a complicated hostname.

For example:

```text
Simple:
bob@yahoo.mail

Actual mail server:
relay1.west-coast.yahoo.com
```

DNS allows the mail application to obtain the actual/canonical mail-server hostname.

This functionality is provided through an **MX (Mail Exchange) record**. Computer Networking- A Top-Down…

---

# 1.4 Load Distribution

DNS can distribute traffic among multiple replicated servers.

Suppose a popular Web site has three replicated servers:

```text
                 www.example.com
                       |
              DNS returns addresses
                 /      |      \
                /       |       \
          Server A   Server B   Server C
        1.2.3.10    1.2.3.11   1.2.3.12
```

One canonical hostname can therefore correspond to **multiple IP addresses**.

DNS can return the complete set of addresses but **rotate their ordering** between replies.

Since clients commonly use the first address, this rotation distributes requests among the replicated servers. Computer Networking- A Top-Down…

The same technique can also be used for multiple mail servers.

---

# 2. How DNS Works

This is the most important part of the topic.

## Why can't there be one DNS server?

The Internet has an enormous number of hosts.

A single centralized DNS server would have problems with:

- **scale** — enormous database;
- **traffic volume** — huge number of queries;
- **single point of failure**;
- **geographical distance**;
- **maintenance and updating**.

Therefore, DNS uses a **distributed and hierarchical database**.

The book identifies three major classes of DNS servers:

1. **Root DNS servers**
2. **TLD DNS servers**
3. **Authoritative DNS servers**

There is also a very important fourth type:

4. **Local DNS server** Computer Networking- A Top-Down…

---

# 2.1 DNS Server Hierarchy

## A. Root DNS Servers

Root servers are at the top of the DNS hierarchy.

Their important job is to provide information about **TLD servers**.

For example, for:

```text
www.amazon.com
```

the root server can direct the query toward the DNS servers responsible for `.com`.

The book describes hundreds of root-server instances distributed around the world. Computer Networking- A Top-Down…

---

## B. TLD DNS Servers

**TLD = Top-Level Domain**

Examples:

```text
.com
.org
.net
.edu
.gov
```

There are also country-code TLDs such as:

```text
.uk
.fr
.ca
.jp
```

A TLD server provides information about the **authoritative DNS server** for the requested domain.

For example:

```text
www.amazon.com
      ↓
Root
      ↓
.com TLD
      ↓
Authoritative server for amazon.com
```

Computer Networking- A Top-Down…

---

## C. Authoritative DNS Servers

An **authoritative DNS server** contains the DNS records for hosts belonging to a particular organization/domain.

For example:

```text
University of Massachusetts
          ↓
Authoritative DNS server
          ↓
cs.umass.edu hosts
```

If the requested hostname is:

```text
gaia.cs.umass.edu
```

the authoritative DNS server for that domain can provide the corresponding IP address. Computer Networking- A Top-Down…

---

## D. Local DNS Server

A **local DNS server**, also called a **default name server**, is typically operated by an ISP or institution.

When a host connects to its ISP, the ISP provides the host with the IP address of one or more local DNS servers, typically through DHCP.

The local DNS server acts as a **proxy**, forwarding queries into the DNS hierarchy when necessary. Computer Networking- A Top-Down…

---

# 2.2 Complete DNS Lookup

Suppose:

```text
Host:
cse.nyu.edu

wants IP address of:

gaia.cs.umass.edu
```

Assume the local DNS server is:

```text
dns.nyu.edu
```

and the authoritative DNS server is:

```text
dns.umass.edu
```

The lookup proceeds approximately as follows.

### Step 1 — Host → Local DNS server

```text
cse.nyu.edu
     |
     | DNS query:
     | gaia.cs.umass.edu
     ↓
dns.nyu.edu
```

The requesting host asks its local DNS server. Computer Networking- A Top-Down…

### Step 2 — Local DNS → Root DNS server

The local DNS server asks a root server.

The root server sees the `.edu` portion and returns information about the `.edu` TLD servers.

```text
dns.nyu.edu
     |
     | query
     ↓
Root DNS
     |
     | .edu TLD server information
     ↓
dns.nyu.edu
```

### Step 3 — Local DNS → TLD server

The local DNS server then queries a `.edu` TLD server.

The TLD server identifies the authoritative DNS server for `umass.edu`.

```text
dns.nyu.edu
     |
     ↓
.edu TLD server
     |
     | authoritative server for umass.edu
     ↓
dns.nyu.edu
```

### Step 4 — Local DNS → Authoritative DNS server

Finally:

```text
dns.nyu.edu
     |
     | gaia.cs.umass.edu?
     ↓
dns.umass.edu
```

The authoritative server returns the IP address of `gaia.cs.umass.edu`.

The local DNS server then returns that address to the requesting host. Computer Networking- A Top-Down…

### Overall

```text
                    Root DNS
                       ↑
                       |
                 .edu TLD DNS
                       ↑
                       |
              Authoritative DNS
                dns.umass.edu
                       ↑
                       |
                Local DNS
                 dns.nyu.edu
                       ↑
                       |
                 Requesting Host
                  cse.nyu.edu
```

---

# 2.3 Recursive vs Iterative Queries

This distinction is important for exams.

## Recursive Query

The client asks a DNS server to **obtain the requested mapping on its behalf**.

Example:

```text
Host → Local DNS:
"Find the IP address of gaia.cs.umass.edu for me."
```

The local DNS server is responsible for obtaining the final answer.

The book's normal example uses a **recursive query from the requesting host to its local DNS server**. Computer Networking- A Top-Down…

---

## Iterative Query

A DNS server responds with the information it currently knows, often directing the querying DNS server toward another DNS server.

For example:

```text
Local DNS → Root
Root → "Ask the .edu TLD server."

Local DNS → .edu TLD
TLD → "Ask dns.umass.edu."

Local DNS → Authoritative
Authoritative → "Here is the IP address."
```

The subsequent queries in the book's typical example are **iterative**. Computer Networking- A Top-Down…

### Difference

| Recursive | Iterative |
|---|---|
| Server is asked to obtain final answer | Server provides information it has |
| "Find this for me" | "Here is where you should look next" |
| Typical host → local DNS interaction | Typical local DNS → hierarchy interaction |

---

# 2.4 DNS Caching

DNS would generate a huge amount of traffic if every lookup always went through:

```text
Local → Root → TLD → Authoritative
```

Therefore, DNS makes extensive use of **caching**.

When a DNS server receives a DNS reply, it can store the mapping in its cache.

For example:

```text
dns.nyu.edu cache:

cnn.com → 157.166.226.25
```

If another host later asks the same local DNS server for `cnn.com`, the server can answer directly from its cache without querying the hierarchy again. Computer Networking- A Top-Down…

### TTL

Cached DNS information cannot remain forever because mappings can change.

Therefore, each resource record has a **TTL (Time To Live)**.

When the TTL expires, the cached record is discarded. Computer Networking- A Top-Down…

### Result of caching

Caching:

- reduces DNS traffic;
- reduces DNS lookup delay;
- means root servers are bypassed for most queries.

Computer Networking- A Top-Down…

---

# 3. DNS Records and Messages

DNS servers store information using **Resource Records (RRs)**.

A resource record has four fields:

\[
\boxed{(Name,\ Value,\ Type,\ TTL)}
\]

Computer Networking- A Top-Down…

The four major record types discussed in the book are:

- **A**
- **NS**
- **CNAME**
- **MX**

---

# 3.1 Type A Record

**A = Address**

An A record maps a hostname to an IP address.

```text
Name                         Value
------------------------------------------------
relay1.bar.foo.com           145.37.93.126
```

Conceptually:

```text
relay1.bar.foo.com
        ↓
145.37.93.126
```

Thus:

\[
\boxed{\text{A record: hostname} \rightarrow \text{IP address}}
\]

Computer Networking- A Top-Down…

---

# 3.2 Type NS Record

**NS = Name Server**

An NS record identifies the **authoritative DNS server** for a particular domain.

Example:

```text
Name:
foo.com

Value:
dns.foo.com
```

Meaning:

```text
foo.com
   ↓
Authoritative DNS server:
dns.foo.com
```

NS records help direct DNS queries further along the DNS hierarchy. Computer Networking- A Top-Down…

---

# 3.3 Type CNAME Record

**CNAME = Canonical Name**

A CNAME record maps an **alias hostname** to its canonical hostname.

Example:

```text
Alias:
foo.com

Canonical hostname:
relay1.bar.foo.com
```

Conceptually:

```text
foo.com
   ↓ CNAME
relay1.bar.foo.com
```

The canonical hostname can then be resolved to its IP address. Computer Networking- A Top-Down…

---

# 3.4 Type MX Record

**MX = Mail Exchange**

An MX record identifies the canonical hostname of the **mail server** associated with a domain.

Example:

```text
Name:
foo.com

Value:
mail.bar.foo.com
```

So:

```text
foo.com
   ↓ MX
mail.bar.foo.com
```

MX records are particularly important for e-mail because they allow a domain to use a convenient alias for its mail server. Computer Networking- A Top-Down…

---

# DNS Record Types — Quick Comparison

| Type | Meaning | Purpose |
|---|---|---|
| **A** | Address | Hostname → IP address |
| **NS** | Name Server | Domain → authoritative DNS server |
| **CNAME** | Canonical Name | Alias hostname → canonical hostname |
| **MX** | Mail Exchange | Domain/mail alias → mail server hostname |

### Easy memory trick

```text
A     → Address
NS    → Name Server
CNAME → Canonical Name
MX    → Mail eXchange
```

---

# 3.5 DNS Messages

DNS has only **two types of messages**:

1. **Query**
2. **Reply**

Importantly, **both use the same message format**. Computer Networking- A Top-Down…

A DNS message contains four major sections:

```text
DNS Message
│
├── Header
├── Question
├── Answer
├── Authority
└── Additional
```

---

## A. Header Section

The header is **12 bytes**.

It contains:

- identification number;
- flags;
- number of questions;
- number of answer RRs;
- number of authority RRs;
- number of additional RRs.

Computer Networking- A Top-Down…

### Identification

The DNS query contains a **16-bit identification number**.

The server copies this identifier into the reply.

This allows the client to match a received reply with the corresponding query. Computer Networking- A Top-Down…

---

## B. Question Section

The question section contains information about what the client wants.

It includes:

1. **Name** — hostname/domain being queried.
2. **Type** — type of record being requested.

For example:

```text
Name:
www.example.com

Type:
A
```

means:

> Find the A record for `www.example.com`.

A query can also request an MX record, for example. Computer Networking- A Top-Down…

---

## C. Answer Section

The answer section contains the **resource records corresponding to the original query**.

For example:

```text
Query:
www.example.com, Type A

Answer:
www.example.com → 93.184.216.34
```

A reply can contain **multiple resource records**.

For example, a hostname may correspond to multiple IP addresses when several replicated servers provide the same service. Computer Networking- A Top-Down…

---

## D. Authority Section

The authority section contains records identifying **authoritative DNS servers**.

This can tell the querying DNS server where authoritative information can be obtained. Computer Networking- A Top-Down…

---

## E. Additional Section

The additional section contains **helpful additional records**.

A good example is an MX lookup.

Suppose the query asks:

```text
foo.com → MX
```

The answer may say:

```text
foo.com → mail.foo.com
```

But the client also needs the IP address of `mail.foo.com`.

The **Additional** section can contain:

```text
mail.foo.com → IP address
```

through an A record. Computer Networking- A Top-Down…

---

# Complete DNS Picture

```text
                 DNS
                  |
      ┌───────────┴───────────┐
      |                       |
 Distributed Database      Protocol
      |                       |
 Root                      Query/Reply
 TLD                       UDP
 Authoritative             Port 53
 Local
      |
      └── Resource Records
             |
      ┌──────┼──────┬──────┐
      A      NS    CNAME   MX
```

And the lookup process:

```text
Application
    |
    | "What is the IP of www.example.com?"
    ↓
Local DNS Server
    |
    ↓
Root DNS
    |
    ↓
.com TLD DNS
    |
    ↓
Authoritative DNS
    |
    | A record
    ↓
IP address
    |
    ↓
Application
```

---

# ⭐ Exam-Critical Summary

### DNS Services

1. **Hostname-to-IP translation** — primary purpose.
2. **Host aliasing** — alias → canonical hostname.
3. **Mail server aliasing** — uses MX records.
4. **Load distribution** — multiple IP addresses for replicated servers; DNS can rotate their ordering. Computer Networking- A Top-Down…

### How DNS Works

- DNS is a **distributed, hierarchical database**.
- Main hierarchy:
  **Root → TLD → Authoritative**
- **Local DNS server** acts as a proxy for the requesting host.
- Typical lookup:
  **Host → Local DNS → Root → TLD → Authoritative → back**.
- **Recursive query**: server obtains the answer on behalf of requester.
- **Iterative query**: server provides information directing the requester to the next server.
- **Caching** reduces DNS traffic and delay.
- **TTL** determines how long cached records remain valid. Computer Networking- A Top-Down…

### DNS Records

\[
\boxed{RR=(Name,Value,Type,TTL)}
\]

- **A** → hostname → IP
- **NS** → domain → authoritative name server
- **CNAME** → alias → canonical hostname
- **MX** → domain → mail server

### DNS Message

```text
Header
Question
Answer
Authority
Additional
```

There are only **two DNS message types: query and reply**, and both have the same format. Computer Networking- A Top-Down…

**One distinction to remember especially well:**  
**DNS is not merely a hostname-to-IP lookup table.** In Kurose & Ross, it is a **distributed database + application-layer protocol**, and its records also support aliasing, mail-server discovery, and load distribution. Computer Networking- A Top-Down…
# 2.5.1 P2P File Distribution

In **Kurose & Ross, *Computer Networking: A Top-Down Approach*, 7th Edition**, P2P file distribution is presented as an application in which **peers themselves participate in distributing the file**, rather than relying entirely on a centralized server. The book uses **BitTorrent** as the main example. Computer Networking- A Top-Down…

---

# 1. What is P2P File Distribution?

In a traditional **client-server architecture**:

```text
                 Server
              /    |    \
             /     |     \
          Peer   Peer   Peer
```

The server must send a complete copy of the file to every peer.

If there are \(N\) peers, the server has to upload:

\[
NF
\]

bits, where \(F\) is the file size.

This places a large burden on the server's upload bandwidth. Computer Networking- A Top-Down…

---

In **P2P file distribution**:

```text
                  Server
                /       \
              Peer ←→ Peer
               ↑  \   /  ↑
               |   \ /   |
              Peer ←→ Peer
```

The server initially provides the file to peers, but once a peer has received some portion of the file, it can **upload that portion to other peers**.

Therefore:

> **Peers are both consumers and redistributors of the file.**

This is the fundamental reason P2P systems can be **self-scalable**. Computer Networking- A Top-Down…

---

# 2. Why P2P is Self-Scalable

This is one of the most important concepts in this section.

Suppose a new peer joins the system.

It creates two things:

- **additional demand** — it wants to download the file;
- **additional capacity** — after receiving chunks, it can upload those chunks to other peers.

So:

```text
New peer joins
      ↓
Needs file
      ↓
Creates download demand
      ↓
Receives chunks
      ↓
Can upload those chunks
      ↓
Adds upload capacity
```

Therefore, as the number of users increases, the system's available upload capacity also increases.

The book calls this **self-scalability**. Computer Networking- A Top-Down…

---

# 3. Client-Server vs P2P File Distribution

The book quantitatively compares the two architectures.

Let:

- \(F\) = file size in bits
- \(N\) = number of peers
- \(u_s\) = server upload rate
- \(u_i\) = upload rate of peer \(i\)
- \(d_i\) = download rate of peer \(i\)
- \(d_{\min}\) = minimum download rate among all peers

The distribution time is the time required for **all \(N\) peers to obtain a complete copy**. Computer Networking- A Top-Down…

---

## 3.1 Client-Server Distribution

The server must send one complete copy to every peer.

Therefore, it must upload:

\[
NF
\]

bits.

Since the server uploads at rate \(u_s\):

\[
\frac{NF}{u_s}
\]

is a lower bound on distribution time.

Also, the slowest peer has download rate \(d_{\min}\), so it cannot download the entire file faster than:

\[
\frac{F}{d_{\min}}
\]

Therefore:

\[
\boxed{
D_{CS}=
\max
\left\{
\frac{NF}{u_s},
\frac{F}{d_{\min}}
\right\}
}
\]

This is Equation 2.1 in the book. Computer Networking- A Top-Down…

### Important consequence

For large \(N\):

\[
D_{CS}\approx\frac{NF}{u_s}
\]

Therefore:

\[
\boxed{D_{CS}\text{ increases linearly with }N}
\]

If the number of peers increases by 1000×, the server's distribution time can also increase by approximately 1000×. Computer Networking- A Top-Down…

---

# 3.2 P2P Distribution

In P2P, the server does **not** have to send the entire file separately to every peer.

Three limitations determine the minimum distribution time.

### Limitation 1 — Server must initially inject the file

At the beginning, only the server has the file.

Therefore, the server must upload all \(F\) bits at least once:

\[
\frac{F}{u_s}
\]

So:

\[
D_{P2P}\geq\frac{F}{u_s}
\]

Computer Networking- A Top-Down…

---

### Limitation 2 — Slowest peer

The slowest peer has download rate \(d_{\min}\).

Therefore:

\[
D_{P2P}\geq\frac{F}{d_{\min}}
\]

Computer Networking- A Top-Down…

---

### Limitation 3 — Total upload capacity

Now comes the key P2P advantage.

The total upload capacity is:

\[
u_s+u_1+u_2+\cdots+u_N
\]

The system must ultimately deliver \(NF\) bits to all peers.

Therefore:

\[
D_{P2P}
\geq
\frac{NF}{u_s+\sum_{i=1}^{N}u_i}
\]

Computer Networking- A Top-Down…

Combining the three constraints:

\[
\boxed{
D_{P2P}
=
\max
\left\{
\frac{F}{u_s},
\frac{F}{d_{\min}},
\frac{NF}{u_s+\sum_{i=1}^{N}u_i}
\right\}
}
\]

This is Equation 2.3 in the book. Computer Networking- A Top-Down…

---

# 4. Why P2P Scales Better

Compare:

### Client-server

\[
D_{CS}
=
\max
\left\{
\frac{NF}{u_s},
\frac{F}{d_{\min}}
\right\}
\]

The server is the main source of upload capacity.

### P2P

\[
D_{P2P}
=
\max
\left\{
\frac{F}{u_s},
\frac{F}{d_{\min}},
\frac{NF}{u_s+\sum u_i}
\right\}
\]

The upload capacity of **all peers contributes**.

Thus, as more peers join:

```text
More peers
    ↓
More download demand
    +
More upload capacity
    ↓
System can scale
```

This is why the book says P2P applications can be **self-scaling**. Computer Networking- A Top-Down…

---

# 5. BitTorrent

The book's main example of P2P file distribution is **BitTorrent**.

A BitTorrent distribution of a particular file is called a **torrent**.

The peers participating in that torrent exchange portions of the file with one another. Computer Networking- A Top-Down…

---

# 6. File Chunks

Instead of treating the file as one huge object, BitTorrent divides it into **equal-size chunks**.

The book gives a typical chunk size of:

\[
\boxed{256\text{ KBytes}}
\]

When a peer first joins:

```text
Peer
 ↓
Has no chunks
```

As it participates:

```text
Peer
 ↓
Chunk 1
Chunk 7
Chunk 12
Chunk 20
...
```

It simultaneously:

- downloads chunks from other peers;
- uploads chunks it already possesses to other peers. Computer Networking- A Top-Down…

Once a peer obtains the entire file, it can:

- leave the torrent, or
- remain and continue uploading.

A peer can also leave before completing the file and later rejoin. Computer Networking- A Top-Down…

---

# 7. Tracker

Each torrent has an infrastructure node called a **tracker**.

The tracker keeps track of peers participating in the torrent. Computer Networking- A Top-Down…

### When a new peer joins

Suppose Alice joins.

```text
                 Tracker
                    |
             list of peers
                    ↓
                  Alice
```

The tracker:

1. receives Alice's registration;
2. selects a subset of participating peers;
3. sends their IP addresses to Alice.

The book's example uses **50 peers** as the selected subset. Computer Networking- A Top-Down…

Alice then attempts to establish TCP connections with those peers.

The peers with which she successfully establishes connections are her **neighboring peers**.

---

# 8. How a Peer Obtains Chunks

Suppose Alice has:

```text
Alice:
Chunk 1
Chunk 4
Chunk 8
```

Her neighboring peers may have different chunks:

```text
Peer B: 1, 2, 3, 7
Peer C: 2, 4, 5, 8
Peer D: 3, 5, 6, 9
```

Alice periodically asks her neighbors which chunks they possess.

She then knows:

- what she has;
- what each neighbor has;
- which chunks she is missing.

She requests missing chunks over the TCP connections. Computer Networking- A Top-Down…

---

# 9. Rarest-First Technique

The first important decision is:

> **Which missing chunk should Alice request first?**

BitTorrent uses the **rarest-first** technique.

Alice examines the chunks available among her neighbors and identifies the chunks that have the **fewest copies**.

She requests those chunks first. Computer Networking- A Top-Down…

### Example

Suppose Alice is missing:

```text
Chunk 2 → available at 5 neighbors
Chunk 3 → available at 3 neighbors
Chunk 7 → available at 1 neighbor
Chunk 9 → available at 4 neighbors
```

Then:

```text
Chunk 7 ← rarest
```

Alice prioritizes **Chunk 7**.

### Why?

If the rare chunk is not redistributed quickly, it could become unavailable if the few peers possessing it leave.

So rarest-first helps:

- redistribute rare chunks;
- maintain a relatively balanced number of copies of chunks throughout the torrent.

Computer Networking- A Top-Down…

---

# 10. Which Peers Receive Alice's Chunks?

The second important decision is:

> **To which neighbors should Alice upload?**

BitTorrent uses a **trading algorithm**.

Alice measures the rate at which each neighboring peer is sending data to her.

She gives priority to the peers that are currently supplying data to her at the highest rate. Computer Networking- A Top-Down…

---

# 11. Unchoking

Alice selects the **four peers** that are currently sending her data at the highest rates.

She sends chunks to those four peers in return.

These four peers are called:

\[
\boxed{\text{unchoked peers}}
\]

Alice recalculates the rates approximately every **10 seconds**. Computer Networking- A Top-Down…

So:

```text
Alice's neighbors

Peer A ── high rate ──→ Alice  ✓ Unchoked
Peer B ── high rate ──→ Alice  ✓ Unchoked
Peer C ── high rate ──→ Alice  ✓ Unchoked
Peer D ── high rate ──→ Alice  ✓ Unchoked

Peer E ────────────────→ Alice  ✗ Choked
Peer F ────────────────→ Alice  ✗ Choked
```

A **choked** peer does not receive chunks from Alice.

---

# 12. Optimistically Unchoked Peer

There is an important exception.

Every **30 seconds**, Alice randomly selects one additional neighboring peer and sends chunks to it.

This peer is called the:

\[
\boxed{\text{optimistically unchoked peer}}
\]

Computer Networking- A Top-Down…

### Why randomly choose someone?

Suppose Bob is a new peer.

Bob currently has very little data, so he cannot upload much to Alice.

Therefore, Bob will probably not be among Alice's top four peers.

If Alice never uploaded to Bob, Bob could never acquire enough chunks to become a useful uploader.

The random optimistic unchoking gives Bob an opportunity to receive chunks.

If Bob then uploads at a good rate, he may become one of Alice's top four peers. Computer Networking- A Top-Down…

---

# 13. Tit-for-Tat

The trading mechanism is commonly called **tit-for-tat**.

The basic principle is:

> **Peers that contribute upload capacity are given better access to other peers' data.**

So a peer that uploads at a high rate is more likely to receive data from other peers.

This discourages **free-riding**—participating in downloads while refusing to upload. Computer Networking- A Top-Down…

---

# 14. Complete BitTorrent Process

Put everything together:

```text
             ┌─────────────┐
             │   Tracker   │
             └──────┬──────┘
                    │
             list of peers
                    ↓
             ┌─────────────┐
             │    Alice    │
             └──────┬──────┘
                    │
        establishes TCP connections
          with neighboring peers
          /        |        \
         ↓         ↓         ↓
      Peer B     Peer C    Peer D
         ↕         ↕         ↕
       chunks    chunks    chunks
```

Alice then:

### 1. Gets peer list
Tracker provides IP addresses of participating peers.

### 2. Establishes connections
Alice creates TCP connections with neighboring peers.

### 3. Learns available chunks
Alice asks neighbors which chunks they possess.

### 4. Uses rarest-first
Alice requests the rarest missing chunks first.

### 5. Downloads chunks
She receives chunks from neighboring peers.

### 6. Uploads chunks
She simultaneously uploads chunks to other peers.

### 7. Uses tit-for-tat
She preferentially uploads to peers that provide data to her at high rates.

### 8. Optimistically unchokes
Every 30 seconds, she randomly tries another neighbor.

Computer Networking- A Top-Down…

---

# 15. Client-Server vs P2P — Final Comparison

| Feature | Client-Server | P2P |
|---|---|---|
| Main uploader | Central server | Server + peers |
| Peer uploads to other peers? | No | **Yes** |
| Server burden | Very high as \(N\) grows | Reduced |
| Scalability | Limited by server upload capacity | **Self-scalable** |
| Distribution time | \(\max\{\frac{NF}{u_s},\frac{F}{d_{\min}}\}\) | \(\max\{\frac{F}{u_s},\frac{F}{d_{\min}},\frac{NF}{u_s+\sum u_i}\}\) |
| File organization | Complete file from server | File divided into chunks |
| Example | Traditional centralized distribution | **BitTorrent** |
| Major challenge | Server bandwidth | Decentralization, security, performance, reliability |

The book also notes that P2P's decentralized nature creates challenges in **security, performance, and reliability**. Computer Networking- A Top-Down…

---

# ⭐ Most Important Exam Points

1. **P2P architecture:** peers communicate directly with one another with minimal reliance on always-on infrastructure servers.
2. **P2P file distribution:** peers both **download and upload** file data.
3. **Main advantage:** **self-scalability**.
4. In client-server distribution:

\[
\boxed{
D_{CS}=
\max
\left\{
\frac{NF}{u_s},
\frac{F}{d_{\min}}
\right\}
}
\]

5. In P2P distribution:

\[
\boxed{
D_{P2P}=
\max
\left\{
\frac{F}{u_s},
\frac{F}{d_{\min}},
\frac{NF}{u_s+\sum u_i}
\right\}
}
\]

6. **BitTorrent** divides files into chunks; peers exchange chunks with one another.
7. A **torrent** is the collection of peers participating in the distribution of a particular file.
8. A **tracker** keeps track of participating peers and provides a joining peer with a list of other peers.
9. **Rarest-first** determines which missing chunks a peer requests first.
10. **Unchoking:** a peer normally uploads to its four highest-rate suppliers.
11. **Optimistic unchoking:** every 30 seconds, one additional neighbor is randomly selected.
12. The upload incentive mechanism is called **tit-for-tat**.
13. **Choked peers** do not receive chunks from that peer.
14. # 2.6 Video Streaming & Content Distribution Networks

According to **Kurose & Ross, *Computer Networking: A Top-Down Approach*, 7th Edition**, this section explains how large amounts of stored video can be delivered over the Internet while maintaining **continuous playback**, adapting to changing bandwidth, and serving millions of users efficiently. The section has four main parts: **Internet Video, HTTP Streaming and DASH, CDNs, and case studies of Netflix, YouTube, and Kankan**. Computer Networking- A Top-Down…

---

# 1. Internet Video

## 1.1 What is Internet Video?

The book focuses primarily on **streaming stored video**.

In this model:

- Video is prerecorded.
- The video is stored on a server.
- A user requests the video.
- The server streams the video to the user.
- The user can watch it on demand.

Examples include:

- movies,
- TV shows,
- prerecorded sporting events,
- user-generated videos such as YouTube videos. Computer Networking- A Top-Down…

Basic architecture:

```text
              Video Server
                  |
                  | video stream
                  ↓
               Internet
                  |
                  ↓
                Client
                  |
                  ↓
              Video Player
```

---

# 1.2 Characteristics of Video

A video is essentially a **sequence of images**, usually displayed at a constant rate.

For example:

- 24 images/second
- 30 images/second

Each image contains pixels representing luminance and color. Computer Networking- A Top-Down…

The major networking issue is the **very high bit rate** of video.

The book gives examples ranging from roughly:

- **100 kbps** for low-quality video
- **over 3 Mbps** for high-definition streaming
- **more than 10 Mbps** for 4K streaming

A 2-Mbps video lasting 67 minutes consumes approximately **1 GB** of storage and traffic. Computer Networking- A Top-Down…

---

# 1.3 Compression and Video Quality

Video can be compressed to reduce its bit rate.

There is a trade-off:

\[
\boxed{\text{Higher bit rate} \rightarrow \text{higher video quality}}
\]

For example, the same video might be encoded as:

```text
300 kbps  → Low quality
1 Mbps    → Medium quality
3 Mbps    → High quality
```

A user with a fast Internet connection can use the higher-rate version, while a user with limited bandwidth can use the lower-rate version. Computer Networking- A Top-Down…

This idea becomes extremely important in **DASH**.

---

# 1.4 Most Important Performance Measure: Throughput

For streaming video, the book identifies **average end-to-end throughput** as a particularly important performance measure.

Suppose the compressed video has a bit rate:

\[
r_v
\]

Then, for continuous playback, the network needs to provide average throughput at least as large as the video's consumption rate:

\[
\boxed{\text{Average throughput} \geq r_v}
\]

Otherwise, the client will eventually run out of video data and playback will freeze. Computer Networking- A Top-Down…

---

# 1.5 Client-Side Buffering

Network throughput and delay can vary over time.

Therefore, the client does **not normally play every received byte immediately**.

Instead:

```text
Server
  |
  | video data
  ↓
Client Application Buffer
  |
  | after sufficient data is accumulated
  ↓
Video Player
  |
  ↓
Screen
```

The client first builds up a reserve of video in an **application buffer**.

Once the buffer reaches a predetermined threshold, playback begins. Computer Networking- A Top-Down…

### Why buffering?

Suppose the network temporarily slows down.

If the client already has several seconds of video buffered:

```text
Buffered video
████████████████
      ↓
Playback
```

the player can continue playing while waiting for more data.

Thus buffering helps absorb:

1. variations in network delay;
2. temporary reductions in available bandwidth.

The book explains that as long as the buffer does not become completely empty, playback can continue without interruption. Computer Networking- A Top-Down…

---

# 2. HTTP Streaming

In **HTTP streaming**, the video is stored at an HTTP server as an ordinary file with a URL.

The basic process is:

### Step 1

Client establishes a TCP connection to the server.

### Step 2

Client sends:

```text
HTTP GET
```

for the video URL.

### Step 3

Server sends the video file in an HTTP response.

### Step 4

Client stores incoming bytes in its application buffer.

### Step 5

Once enough data has accumulated, playback begins.

Computer Networking- A Top-Down…

---

## HTTP Streaming Architecture

```text
                 HTTP GET
Client ─────────────────────────→ HTTP Server
       ←─────────────────────────
          Video data

             ↓
       Client Buffer
             ↓
       Video Decoder
             ↓
          Display
```

The server sends the video as quickly as the underlying network conditions permit. Computer Networking- A Top-Down…

---

# 2.1 Problem with Basic HTTP Streaming

There is a major problem.

With ordinary HTTP streaming, **all clients receive the same encoding of the video**.

But users can have very different bandwidths:

```text
User A → 10 Mbps
User B → 3 Mbps
User C → 500 kbps
```

Even the same user can experience changing bandwidth over time.

For example:

```text
Time →     t1     t2     t3     t4
Bandwidth  5M     2M     8M     1M
```

If the server continuously sends a high-bitrate version, a low-bandwidth client may experience buffering interruptions.

This limitation motivated **DASH — Dynamic Adaptive Streaming over HTTP**. Computer Networking- A Top-Down…

---

# 3. DASH — Dynamic Adaptive Streaming over HTTP

DASH solves the fixed-quality problem of basic HTTP streaming.

The central idea is:

> **Store multiple versions of the same video at different bit rates, divide each version into chunks, and allow the client to choose the appropriate chunk dynamically.**

Computer Networking- A Top-Down…

---

# 3.1 Multiple Versions of the Video

Suppose a video is encoded into:

```text
Version 1 → 300 kbps
Version 2 → 1 Mbps
Version 3 → 3 Mbps
Version 4 → 5 Mbps
```

Each version has a different quality level.

The server stores all versions. Computer Networking- A Top-Down…

---

# 3.2 Video is Divided into Chunks

Each version is divided into **small video segments/chunks**, typically a few seconds long.

Conceptually:

```text
3 Mbps version:
[Chunk 1][Chunk 2][Chunk 3][Chunk 4]...

1 Mbps version:
[Chunk 1][Chunk 2][Chunk 3][Chunk 4]...

300 kbps version:
[Chunk 1][Chunk 2][Chunk 3][Chunk 4]...
```

The client does **not** have to download the entire video at one fixed quality.

Instead, it requests one chunk at a time. Computer Networking- A Top-Down…

---

# 3.3 Manifest File

The HTTP server also provides a **manifest file**.

The manifest tells the client:

- what versions are available;
- the bit rate of each version;
- the URL associated with each version.

Computer Networking- A Top-Down…

Example:

```text
Manifest
│
├── 300 kbps → URL 1
├── 1 Mbps   → URL 2
├── 3 Mbps   → URL 3
└── 5 Mbps   → URL 4
```

---

# 3.4 Client-Driven Adaptation

This is the key feature of DASH.

The **client**, not the server, decides which version to request next.

The client considers factors such as:

- measured receiving bandwidth;
- amount of video currently buffered.

Then it selects the next chunk.

Computer Networking- A Top-Down…

### High bandwidth + large buffer

```text
High bandwidth
      +
Large buffer
      ↓
High-bitrate chunk
```

### Low bandwidth + small buffer

```text
Low bandwidth
      +
Small buffer
      ↓
Low-bitrate chunk
```

Thus the client can dynamically switch quality. Computer Networking- A Top-Down…

---

# 3.5 Example of DASH Adaptation

Suppose:

```text
Available versions:
500 kbps
1 Mbps
2 Mbps
4 Mbps
```

Initially:

```text
Bandwidth = 1 Mbps
Buffer = small
```

Client chooses:

```text
500 kbps or 1 Mbps
```

Later:

```text
Bandwidth = 5 Mbps
Buffer = large
```

Client may switch to:

```text
4 Mbps
```

Later the network becomes congested:

```text
Bandwidth = 700 kbps
Buffer = decreasing
```

Client switches down:

```text
500 kbps
```

Therefore:

\[
\boxed{\text{DASH adapts video quality to current network conditions}}
\]

This is particularly useful for mobile users whose available bandwidth can fluctuate as they move relative to cellular base stations. Computer Networking- A Top-Down…

---

# 3.6 DASH Uses HTTP GET and Byte Ranges

The client requests each chunk using an HTTP GET request.

The request specifies:

- the URL;
- the appropriate byte range for the desired chunk.

While downloading, the client measures received bandwidth and determines the quality of the next chunk. Computer Networking- A Top-Down…

So:

```text
Manifest
   ↓
Client learns available versions
   ↓
Measure bandwidth + buffer
   ↓
Choose next chunk quality
   ↓
HTTP GET
   ↓
Receive chunk
   ↓
Repeat
```

---

# HTTP Streaming vs DASH

| Feature | HTTP Streaming | DASH |
|---|---|---|
| Video versions | Typically one encoding | **Multiple bit-rate versions** |
| Chunk selection | Fixed stream | **Client chooses each chunk** |
| Adapts to bandwidth? | No | **Yes** |
| Quality can change during session? | No | **Yes** |
| Manifest | Not central | **Yes** |
| Uses HTTP | Yes | Yes |
| Client buffering | Yes | Yes |

The fundamental difference:

\[
\boxed{\text{HTTP Streaming = fixed encoding}}
\]

\[
\boxed{\text{DASH = adaptive, client-selected encoding}}
\]

Computer Networking- A Top-Down…

---

# 4. Content Distribution Networks — CDN

Now consider the scale of services such as YouTube.

A single massive data center could theoretically store all videos and send them to users around the world.

But the book identifies **three major problems** with this approach.

---

## Problem 1 — Long Network Paths

A client far away from the data center may have to traverse:

- many communication links;
- many ISPs;
- potentially multiple continents.

If one link becomes a bottleneck:

\[
\text{End-to-end throughput}
<
\text{Video consumption rate}
\]

and the video can freeze. Computer Networking- A Top-Down…

---

## Problem 2 — Repeated Traffic

Suppose one popular video is requested by 10,000 users in the same region.

The same video data may repeatedly travel over the same network links.

This:

- wastes network bandwidth;
- increases the content provider's bandwidth cost. Computer Networking- A Top-Down…

---

## Problem 3 — Single Point of Failure

One massive data center is a single point of failure.

If:

```text
Data center fails
      ↓
No video delivery
```

or its Internet connections fail, the service can become unavailable. Computer Networking- A Top-Down…

---

# 4.1 What is a CDN?

A **Content Distribution Network (CDN)** solves these problems by placing servers in **multiple geographically distributed locations**.

A CDN:

1. stores copies of videos/content at multiple locations;
2. receives requests from users;
3. directs each request to an appropriate CDN location.

Computer Networking- A Top-Down…

Architecture:

```text
                       Origin
                         |
              ┌──────────┼──────────┐
              ↓          ↓          ↓
           CDN A       CDN B       CDN C
              ↑          ↑          ↑
              |          |          |
           Users       Users       Users
```

Instead of:

```text
Every user
     ↓
One giant data center
```

we have:

```text
Users
 ↓
Nearby / appropriate CDN server
 ↓
Video
```

---

# 4.2 Benefits of CDNs

CDNs improve video delivery by:

### 1. Reducing path length

Users can obtain video from a nearby CDN server.

### 2. Reducing bottlenecks

Traffic does not have to cross the same long paths from one central data center.

### 3. Reducing repeated long-distance traffic

Popular content can be served from locations closer to users.

### 4. Improving scalability

Traffic is distributed across many servers.

### 5. Improving reliability

Failure of one CDN location does not necessarily take down the entire service.

The book therefore describes CDNs as central to large-scale Internet video delivery. Computer Networking- A Top-Down…

---

# 4.3 Private vs Third-Party CDN

The book describes two possibilities.

### Private CDN

Owned by the content provider.

Examples:

- Google's CDN for YouTube
- Netflix's private CDN

### Third-party CDN

A separate company distributes content for multiple providers.

Examples discussed by the book include:

- Akamai
- Limelight
- Level-3

Computer Networking- A Top-Down…

---

# 4.4 CDN Server Placement

The book describes two placement philosophies.

## Enter Deep

Place many CDN clusters **deep inside access networks**, close to end users.

Advantages:

- shorter paths;
- lower delay;
- potentially better user experience.

Disadvantage:

- many locations are required;
- management becomes more complicated.

## Bring Home

Place relatively fewer but larger clusters closer to access networks.

This reduces management complexity but can result in longer paths to users. Computer Networking- A Top-Down…

---

# 4.5 CDN Cluster Selection

A major CDN problem is:

> **Which CDN server should serve a particular client?**

This is called the **cluster selection strategy**.

The CDN can use information about the client's **LDNS IP address** obtained through DNS lookup. Computer Networking- A Top-Down…

One simple approach:

```text
Client's LDNS
     ↓
Determine geographic location
     ↓
Choose geographically closest CDN cluster
```

But geographically closest does **not necessarily mean network closest**.

For example:

```text
Cluster A → geographically closer
Cluster B → fewer network hops / lower delay

                 ↓
             Choose B
```

The book also describes strategies based on **real-time measurements of delay and loss**. Computer Networking- A Top-Down…

---

# 5. Netflix Case Study

Netflix's architecture in the **7th-edition book** has two major components:

1. **Amazon cloud**
2. **Netflix's private CDN**

Computer Networking- A Top-Down…

---

## 5.1 Amazon Cloud

Netflix uses the Amazon cloud for functions including:

- Web site;
- user registration/login;
- billing;
- movie catalogue;
- searching;
- recommendations;
- content ingestion;
- content processing.

Computer Networking- A Top-Down…

### Content processing

Netflix creates multiple versions of each movie:

```text
Movie
 ↓
Different formats
 +
Different bit rates
 ↓
DASH-compatible versions
```

This supports adaptive streaming over HTTP using DASH. Computer Networking- A Top-Down…

---

# 5.2 Netflix Private CDN

Netflix originally used third-party CDNs.

It later developed its **own private CDN** for video delivery.

Netflix placed server racks:

- at IXPs;
- within residential ISPs.

Computer Networking- A Top-Down…

This brings video content very close to users.

---

# 5.3 Netflix Content Distribution

Netflix does **push caching**, rather than pull caching.

That means:

```text
Amazon Cloud
     |
     | scheduled content push
     ↓
Netflix CDN servers
```

Content is pushed to CDN servers during **off-peak hours**.

Locations that cannot store the entire library receive the most popular videos. Computer Networking- A Top-Down…

---

# 5.4 Netflix Playback

When a user selects a movie:

### Step 1

Netflix software in the Amazon cloud determines which CDN servers contain the movie.

### Step 2

It selects the appropriate/best CDN server.

Typically:

```text
Residential ISP has Netflix server
          ↓
Use that server
```

Otherwise:

```text
Nearby IXP server
```

may be selected. Computer Networking- A Top-Down…

### Step 3

Netflix gives the client:

- IP address of the selected CDN server;
- manifest file containing URLs for different video versions.

### Step 4

Client communicates directly with the CDN server using Netflix's proprietary version of DASH.

### Step 5

The client requests chunks using HTTP GET and byte ranges.

Netflix's chunks in the book's example are approximately **4 seconds long**.

The client measures received throughput and selects the quality of the next chunk. Computer Networking- A Top-Down…

---

## Netflix — Key Characteristics

```text
Amazon Cloud
   |
   | creates multiple versions
   ↓
Netflix Private CDN
   |
   | push caching
   ↓
Client
   |
   | proprietary DASH
   ↓
Adaptive video
```

Important points:

- **Private CDN**
- **DASH**
- **Multiple bit-rate versions**
- **Push caching**
- **No DNS redirect required for CDN selection** in the architecture described by the book
- Client receives the selected CDN server's IP from Netflix software. Computer Networking- A Top-Down…

---

# 6. YouTube Case Study

YouTube also uses a **private CDN**.

Google has server clusters in many IXP and ISP locations and also distributes videos directly from large data centers. Computer Networking- A Top-Down…

But there are important differences from Netflix.

---

## 6.1 YouTube Uses Pull Caching

Unlike Netflix:

```text
Netflix → Push caching
YouTube → Pull caching
```

With pull caching, content can be retrieved into a CDN server when needed, rather than being proactively pushed to every location. Computer Networking- A Top-Down…

---

## 6.2 YouTube Uses DNS Redirect

YouTube uses **DNS redirect** to direct clients toward an appropriate CDN cluster.

The cluster-selection strategy generally attempts to select a cluster with low RTT, although it may sometimes choose a more distant cluster to balance load. Computer Networking- A Top-Down…

---

## 6.3 YouTube and HTTP Streaming

The 7th-edition case study describes YouTube as using HTTP streaming with several versions of a video at different bit rates.

However, the book states that in this architecture YouTube **does not use adaptive streaming such as DASH**; instead, the user manually selects the version. Computer Networking- A Top-Down…

YouTube also uses **HTTP byte-range requests** to avoid unnecessarily transmitting data after the desired amount of video has been prefetched or when the user repositions/stops playback. Computer Networking- A Top-Down…

---

# Netflix vs YouTube

| Feature | Netflix | YouTube |
|---|---|---|
| CDN | Private | Private |
| CDN location | IXPs + residential ISPs | Many IXPs + ISPs + large data centers |
| Cache strategy | **Push caching** | **Pull caching** |
| CDN selection | Netflix software directly selects server | **DNS redirect** |
| Streaming | Proprietary DASH | HTTP streaming |
| Adaptive DASH in book's case study | **Yes** | **No** |
| Multiple bit-rate versions | Yes | Yes |
| Client selects quality dynamically | **Yes** | **No, manual selection in book's description** |
| Byte-range requests | Yes | Yes |

Computer Networking- A Top-Down… Computer Networking- A Top-Down…

---

# 7. Kankan — Briefly

The book also discusses **Kankan**, operated by Xunlei, as a P2P video-delivery example.

Kankan originally used P2P delivery:

```text
Tracker
   ↓
Peers
 ↙ ↓ ↘
Video chunks
```

A peer requesting a video:

1. contacts a tracker;
2. discovers peers having the video;
3. requests video chunks in parallel.

Unlike ordinary BitTorrent file downloading, Kankan preferentially requests chunks that will be played **in the near future**, because continuous playback is required. Computer Networking- A Top-Down…

Kankan later moved to a **hybrid CDN-P2P system**.

```text
              CDN
             /   \
            ↓     ↓
         Client ←→ Peers
```

The CDN helps with startup. Once P2P traffic becomes sufficient, the client can rely primarily on peers. If P2P traffic becomes insufficient, CDN connections are restarted. Computer Networking- A Top-Down…

---

# ⭐ Complete Conceptual Flow

This entire section can be understood as a progression:

```text
             INTERNET VIDEO
                   │
                   ↓
          High video bit rate
                   │
                   ↓
          Need sufficient throughput
                   │
                   ↓
          Client-side buffering
                   │
                   ↓
          HTTP Streaming
                   │
                   ↓
       Problem: fixed video quality
                   │
                   ↓
                DASH
                   │
        ┌──────────┴──────────┐
        ↓                     ↓
 Multiple bit rates      Client adaptation
        │                     │
        └──────────┬──────────┘
                   ↓
          Need massive scale
                   │
                   ↓
                 CDN
                   │
       ┌───────────┴───────────┐
       ↓                       ↓
 Distributed servers      Server selection
       │
       ↓
 Netflix / YouTube / Kankan
```

---

# ⭐ Exam-Critical Points

### Internet Video

- Video is a sequence of images.
- Compression trades **video quality against bit rate**.
- Video has very high bit-rate requirements.
- Continuous playback requires:

\[
\boxed{\text{Average throughput} \geq \text{video consumption rate}}
\]

- Client-side buffering absorbs variations in delay and temporary bandwidth reductions. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

### HTTP Streaming

- Video is stored as an ordinary file on an HTTP server.
- Client uses TCP + HTTP GET.
- Video bytes are stored in a client-side buffer.
- Playback begins after sufficient buffering.
- Basic HTTP streaming uses a fixed encoding for all clients. Computer Networking- A Top-Down…

### DASH

\[
\boxed{\text{Multiple versions}+\text{chunks}+\text{client adaptation}}
\]

- Multiple bit rates.
- Chunks of a few seconds.
- Manifest file.
- Client chooses each chunk.
- Uses HTTP GET.
- Can switch quality as bandwidth changes. Computer Networking- A Top-Down…

### CDN

- Distributed servers geographically close to users.
- Reduces long network paths and bottlenecks.
- Avoids repeatedly sending popular content across the same long-distance links.
- Removes dependence on one massive data center.
- CDN must select an appropriate server/cluster. Computer Networking- A Top-Down…

### Netflix

\[
\boxed{\text{Amazon Cloud}+\text{Private CDN}+\text{DASH}+\text{Push caching}}
\]

### YouTube

\[
\boxed{\text{Private CDN}+\text{DNS redirect}+\text{Pull caching}+\text{HTTP streaming}}
\]

### Kankan

\[
\boxed{\text{P2P video} \rightarrow \text{Hybrid CDN-P2P}}
\]

These distinctions—especially **Netflix vs YouTube: push vs pull caching, direct server selection vs DNS redirect, and DASH vs the book's described non-adaptive HTTP streaming**—are the ones most worth memorizing for an exam. Computer Networking- 
# 2.7 Socket Programming — Kurose & Ross

In Section **2.7, “Socket Programming: Creating Network Applications,”** Kurose and Ross show how an application-layer program can communicate with another application running on a different host using the **socket API**. The book demonstrates this using simple **Python 3 client-server programs**, first with UDP and then with TCP. Computer Networking- A Top-Down…

The same simple application is used for both:

1. Client reads a line from the keyboard.
2. Client sends it to the server.
3. Server converts it to uppercase.
4. Server sends it back.
5. Client displays it. Computer Networking- A Top-Down…

---

# 1. What is a Socket?

Recall the book's analogy:

> **Process = house**  
> **Socket = door**

The application process is on one side of the socket, while the transport-layer protocol is on the other side. The application developer controls the application side, but has much less control over the transport-layer side. Computer Networking- A Top-Down…

So:

```text
Application Process
       |
       |  Application data
       ↓
   +---------+
   | Socket  |   ← interface
   +---------+
       |
       ↓
 Transport Layer
       |
       ↓
    Internet
```

A socket therefore provides the **interface between the application and transport layer**.

The application chooses whether to use **TCP or UDP**, and then uses the appropriate socket operations.

---

# 2. UDP Socket Programming

## 2.1 Basic idea

UDP is **connectionless**.

There is **no connection establishment** before sending data. For each message, the sender must specify where the message should go.

The destination address consists of:

\[
\boxed{\text{Destination IP address + Destination port number}}
\]

The IP address identifies the destination host, while the port number identifies the particular socket/application process on that host. Computer Networking- A Top-Down…

The basic communication is:

```text
             UDP
Client --------------------> Server
       IP + Port

Client <-------------------- Server
             UDP
```

Unlike TCP, there is no handshake and no dedicated connection between the two processes.

---

# 3. UDP Client Program

The book's UDP client is essentially:

```python
from socket import *

serverName = 'hostname'
serverPort = 12000

clientSocket = socket(AF_INET, SOCK_DGRAM)

message = input('Input lowercase sentence:')

clientSocket.sendto(
    message.encode(),
    (serverName, serverPort)
)

modifiedMessage, serverAddress = \
    clientSocket.recvfrom(2048)

print(modifiedMessage.decode())

clientSocket.close()
```

The book uses `raw_input()` in its original Python presentation; conceptually, the important socket operations are the same. The book's UDP example uses port **12000** for the server. Computer Networking- A Top-Down…

Let's understand the important lines.

---

## 3.1 Import socket module

```python
from socket import *
```

The Python `socket` module provides the functions needed to create and use sockets. Computer Networking- A Top-Down…

---

## 3.2 Specify server

```python
serverName = 'hostname'
serverPort = 12000
```

`serverName` can contain either:

- the server's IP address, or
- its hostname.

If a hostname is supplied, a DNS lookup can be performed to obtain the IP address.

`serverPort` identifies the server's socket. Computer Networking- A Top-Down…

---

## 3.3 Create UDP socket

```python
clientSocket = socket(AF_INET, SOCK_DGRAM)
```

This is extremely important.

### `AF_INET`

Specifies the **IPv4 address family**.

### `SOCK_DGRAM`

Specifies a **UDP socket**.

So:

```text
socket(AF_INET, SOCK_DGRAM)
          │       │
          │       └── UDP
          └────────── IPv4
```

The client does **not** explicitly specify its port number. The operating system assigns one automatically. Computer Networking- A Top-Down…

---

# 4. Sending Data with UDP

```python
clientSocket.sendto(
    message.encode(),
    (serverName, serverPort)
)
```

There are two important operations here.

### `encode()`

The message is converted from a string into **bytes**, because data sent through the socket is represented as bytes.

### `sendto()`

The client supplies the destination address:

```text
(serverName, serverPort)
```

That means:

```text
Destination:
    IP address = serverName
    Port       = serverPort
```

The UDP application explicitly supplies the destination address when sending. The source address is automatically attached by the underlying operating system. Computer Networking- A Top-Down…

---

# 5. Receiving Data with UDP

```python
modifiedMessage, serverAddress = \
    clientSocket.recvfrom(2048)
```

`recvfrom()` waits for a UDP packet to arrive.

It returns:

1. the received data → `modifiedMessage`
2. the source address → `serverAddress`

The source address contains:

```text
Server IP address
+
Server port number
```

`2048` specifies the buffer size used for receiving the data. Computer Networking- A Top-Down…

Then:

```python
print(modifiedMessage.decode())
```

converts the received bytes back into a string and displays it. Computer Networking- A Top-Down…

Finally:

```python
clientSocket.close()
```

closes the socket. Computer Networking- A Top-Down…

---

# 6. UDP Server Program

The server is:

```python
from socket import *

serverPort = 12000

serverSocket = socket(AF_INET, SOCK_DGRAM)

serverSocket.bind(('', serverPort))

print("The server is ready to receive")

while True:
    message, clientAddress = serverSocket.recvfrom(2048)

    modifiedMessage = message.decode().upper()

    serverSocket.sendto(
        modifiedMessage.encode(),
        clientAddress
    )
```

The book's implementation is intentionally minimal so that the key socket concepts are clear. Computer Networking- A Top-Down…

---

## 6.1 Create UDP server socket

```python
serverSocket = socket(AF_INET, SOCK_DGRAM)
```

Again:

- `AF_INET` → IPv4
- `SOCK_DGRAM` → UDP

---

## 6.2 Bind the server socket

```python
serverSocket.bind(('', serverPort))
```

This is a crucial difference between the client and server.

The server explicitly associates **port 12000** with its socket.

Therefore:

```text
Server Host
     |
     +---- Port 12000
               |
         serverSocket
```

When a UDP packet arrives at the server's IP address with destination port `12000`, the transport layer delivers it to this socket. Computer Networking- A Top-Down…

---

# 7. Receiving a UDP Packet at the Server

```python
message, clientAddress = serverSocket.recvfrom(2048)
```

When a packet arrives:

- `message` contains the data.
- `clientAddress` contains the client's IP address and port number.

The server needs `clientAddress` because it must know **where to send the reply**. Computer Networking- A Top-Down…

This is analogous to a return address in postal mail.

---

# 8. Processing the Message

```python
modifiedMessage = message.decode().upper()
```

The server:

1. converts bytes into a string using `decode()`;
2. converts the string to uppercase using `upper()`.

For example:

```text
Client:
hello networking

        ↓

Server:
HELLO NETWORKING
```

The `upper()` operation is simply the application logic used by this example. Computer Networking- A Top-Down…

---

# 9. Sending the UDP Reply

```python
serverSocket.sendto(
    modifiedMessage.encode(),
    clientAddress
)
```

The server:

1. converts the modified message into bytes;
2. uses `clientAddress` as the destination;
3. sends the resulting UDP packet.

After sending, the server returns to the `while` loop and waits for another UDP packet. It can therefore receive packets from **any client**. Computer Networking- A Top-Down…

---

# 10. Complete UDP Communication

The complete sequence is:

```text
                  CLIENT                         SERVER

             Create socket                 Create socket
             SOCK_DGRAM                    SOCK_DGRAM
                  |                              |
                  |                              |
                  |---- UDP datagram ---------->|
                  |     destination IP + port   |
                  |                              |
                  |                       recvfrom()
                  |                              |
                  |                       Convert to
                  |                       uppercase
                  |                              |
                  |<---- UDP datagram -----------|
                  |                              |
             recvfrom()                         |
                  |                              |
             Display result                     |
                  |                              |
             close socket                       |
```

The important point is that **every UDP send identifies the destination address**. Computer Networking- A Top-Down…

---

# 11. TCP Socket Programming

Now the book moves to TCP.

The fundamental difference is:

> **TCP is connection-oriented.**

Before client and server exchange application data, they must first establish a TCP connection using a **three-way handshake**. Computer Networking- A Top-Down…

Conceptually:

```text
Client                              Server
  |                                   |
  |       Connection request          |
  |---------------------------------->|
  |                                   |
  |       Three-way handshake         |
  |<--------------------------------->|
  |                                   |
  |======== TCP connection ==========|
  |                                   |
  |             Data                  |
  |---------------------------------->|
  |             Data                  |
  |<----------------------------------|
```

Once the connection exists, the application simply writes data into the socket. It does **not** attach a destination address to every piece of data as it does with UDP. Computer Networking- A Top-Down…

---

# 12. TCP Has Two Server-Side Sockets

This is one of the **most important concepts in the section**.

The TCP server has:

### 1. Welcoming socket

Called:

```python
serverSocket
```

Its purpose is to accept incoming connection requests.

### 2. Connection socket

Called:

```python
connectionSocket
```

It is created specifically for communicating with a particular client.

So:

```text
                 TCP SERVER

          +---------------------+
          |                     |
Client -->| Welcoming socket    |
          |   serverSocket      |
          |         |           |
          |         ↓           |
          | Connection socket   |
          | connectionSocket    |
          |         ↕           |
          +---------|-----------+
                    |
                  Client
```

The welcoming socket remains available for other clients, while the connection socket handles the current client. Computer Networking- A Top-Down…

This distinction is **very exam-important**.

---

# 13. TCP Client Program

The book's TCP client is:

```python
from socket import *

serverName = 'servername'
serverPort = 12000

clientSocket = socket(AF_INET, SOCK_STREAM)

clientSocket.connect((serverName, serverPort))

sentence = input('Input lowercase sentence:')

clientSocket.send(sentence.encode())

modifiedSentence = clientSocket.recv(1024)

print('From Server: ', modifiedSentence.decode())

clientSocket.close()
```

The book's original code uses `raw_input()` in its Python presentation. The socket operations are what matter conceptually. Computer Networking- A Top-Down…

---

# 14. Create a TCP Socket

```python
clientSocket = socket(AF_INET, SOCK_STREAM)
```

Here:

```text
AF_INET
   ↓
IPv4

SOCK_STREAM
   ↓
TCP
```

Compare:

```python
socket(AF_INET, SOCK_DGRAM)   # UDP
socket(AF_INET, SOCK_STREAM)  # TCP
```

The client port is again automatically assigned by the operating system. Computer Networking- A Top-Down…

---

# 15. Establish TCP Connection

```python
clientSocket.connect((serverName, serverPort))
```

This is the major difference from UDP.

The client specifies the server's:

- IP address/hostname
- port number

and `connect()` initiates the TCP connection.

The **three-way handshake** occurs at the transport layer, establishing the TCP connection. The application program does not itself implement the handshake. Computer Networking- A Top-Down…

So:

```text
clientSocket.connect()
        ↓
TCP three-way handshake
        ↓
TCP connection established
        ↓
Application data can be exchanged
```

---

# 16. Send Data over TCP

```python
clientSocket.send(sentence.encode())
```

Notice what is **not** present:

```text
No destination IP
No destination port
No sendto()
```

Why?

Because the TCP connection has already associated the client and server socket addresses.

The application simply puts bytes into the TCP connection. Computer Networking- A Top-Down…

This is a major distinction:

### UDP

```python
sendto(data, destination)
```

### TCP

```python
send(data)
```

because the connection already identifies the destination.

---

# 17. Receive Data over TCP

```python
modifiedSentence = clientSocket.recv(2048)
```

The client receives bytes from the established TCP connection.

The returned bytes are converted to a string:

```python
modifiedSentence.decode()
```

and displayed. Computer Networking- A Top-Down…

---

# 18. Close TCP Connection

```python
clientSocket.close()
```

This closes the socket and therefore closes the TCP connection.

The closing operation causes TCP at the client to send an appropriate TCP message to TCP at the server. Computer Networking- A Top-Down…

---

# 19. TCP Server Program

The TCP server from the book is:

```python
from socket import *

serverPort = 12000

serverSocket = socket(AF_INET, SOCK_STREAM)

serverSocket.bind(('', serverPort))

serverSocket.listen(1)

print('The server is ready to receive')

while True:
    connectionSocket, addr = serverSocket.accept()

    sentence = connectionSocket.recv(1024).decode()

    capitalizedSentence = sentence.upper()

    connectionSocket.send(capitalizedSentence.encode())

    connectionSocket.close()
```

Computer Networking- A Top-Down…

Now understand the important operations.

---

# 20. `bind()` in TCP

```python
serverSocket.bind(('', serverPort))
```

The server associates port `12000` with its welcoming socket.

Clients therefore know where to contact the server. Computer Networking- A Top-Down…

---

# 21. `listen()`

```python
serverSocket.listen(1)
```

This makes `serverSocket` a **welcoming socket** that listens for incoming TCP connection requests.

The parameter specifies the maximum number of queued connections; in this example it is `1`. Computer Networking- A Top-Down…

---

# 22. `accept()`

```python
connectionSocket, addr = serverSocket.accept()
```

When a client connects, `accept()` creates a **new socket**:

```text
connectionSocket
```

This socket is dedicated to that particular client.

The original:

```text
serverSocket
```

remains open as the welcoming socket. Computer Networking- A Top-Down…

So:

```text
                 SERVER

        serverSocket
       (welcoming socket)
              |
          accept()
              |
              ↓
       connectionSocket
       (client-specific)
              |
              ↕
            Client
```

---

# 23. Communication through `connectionSocket`

Once the TCP connection is established:

```python
sentence = connectionSocket.recv(1024).decode()
```

The server receives data from that particular client.

Then:

```python
capitalizedSentence = sentence.upper()
```

performs the application-level processing.

Then:

```python
connectionSocket.send(
    capitalizedSentence.encode()
)
```

sends the result back to that client.

TCP guarantees that bytes sent from one side arrive at the other side **and arrive in the same order**. Thus TCP provides a reliable service between the client and server processes. Computer Networking- A Top-Down…

---

# 24. Why Does the Server Close `connectionSocket` but Not `serverSocket`?

The book's server does:

```python
connectionSocket.close()
```

but does **not** close:

```python
serverSocket
```

because:

- `connectionSocket` belongs to the current client connection.
- `serverSocket` is the welcoming socket.

After closing `connectionSocket`, the server can return to:

```python
accept()
```

and allow another client to connect. Computer Networking- A Top-Down…

---

# 25. TCP Communication Flow

The complete sequence is:

```text
CLIENT                                  SERVER

Create clientSocket              Create serverSocket
       |                                 |
       |                           bind(port)
       |                                 |
       |                           listen()
       |                                 |
       |--------- connect() ------------>|
       |                                 |
       |       Three-way handshake       |
       |<------------------------------->|
       |                                 |
       |                           accept()
       |                                 |
       |                    connectionSocket created
       |                                 |
       |------ send(data) -------------->|
       |                                 |
       |                         recv(data)
       |                                 |
       |                         process data
       |                                 |
       |<----- send(modified) -----------|
       |                                 |
       |------ recv(modified)            |
       |                                 |
       |                                 |
       |------ close() ------------------|
       |                                 |
```

The crucial point is that the **connection is established before application data is exchanged**. Computer Networking- A Top-Down…

---

# 26. UDP vs TCP Socket Programming

| Feature | UDP Socket | TCP Socket |
|---|---|---|
| Transport type | Connectionless | Connection-oriented |
| Socket type | `SOCK_DGRAM` | `SOCK_STREAM` |
| Connection establishment | None | Required |
| Handshake | No | Three-way handshake |
| Sending function | `sendto()` | `send()` |
| Destination address during send | Explicitly specified | Already associated with connection |
| Receiving | `recvfrom()` | `recv()` |
| Server socket | One socket can receive from many clients | Welcoming socket + connection socket |
| Reliability | UDP does not guarantee delivery | TCP provides reliable byte-stream service |
| Ordering | No guarantee from UDP | Bytes arrive in order |
| Server operation | `bind()` then `recvfrom()` | `bind()` → `listen()` → `accept()` |
| Closing | `close()` | `close()` terminates socket/connection |

The book explicitly summarizes the distinction as: TCP is connection-oriented and provides a reliable byte-stream channel, while UDP is connectionless and sends independent packets without delivery guarantees. Computer Networking- A Top-Down…

---

# 27. The Most Important Exam Distinctions ⭐

### UDP

```text
Create socket
     ↓
sendto(data, destination)
     ↓
recvfrom()
     ↓
close()
```

The **destination address is supplied with the packet**.

### TCP

```text
Create socket
     ↓
connect()
     ↓
send()/recv()
     ↓
close()
```

The **connection itself identifies the communicating endpoints**.

---

### TCP Server

Remember this sequence:

\[
\boxed{\text{socket} \rightarrow \text{bind} \rightarrow \text{listen} \rightarrow \text{accept}}
\]

Then:

\[
\boxed{\text{connectionSocket} \rightarrow \text{recv/send} \rightarrow \text{close}}
\]

### UDP Server

Remember:

\[
\boxed{\text{socket} \rightarrow \text{bind} \rightarrow \text{recvfrom/sendto}}
\]

There is **no `listen()` and no `accept()`**.

---

## Final conceptual picture

```text
                  SOCKET PROGRAMMING

                         Socket API
                             |
              +--------------+--------------+
              |                             |
             UDP                           TCP
              |                             |
       Connectionless                Connection-oriented
              |                             |
          SOCK_DGRAM                   SOCK_STREAM
              |                             |
          sendto()                       connect()
          recvfrom()                     send()
                                        recv()
              |                             |
       No handshake                 Three-way handshake
              |                             |
       Independent packets          Reliable byte stream
```

