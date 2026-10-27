# 1.1 What Is the Internet?

Kurose & Ross explain the Internet from **two complementary viewpoints**:

1. **Nuts-and-Bolts Description** → *What physical and software components make up the Internet?*
2. **Services Description** → *What does the Internet provide to applications?*
3. **What Is a Protocol?** → *What rules allow all these components to communicate correctly?*

The book deliberately uses both descriptions because the Internet is simultaneously **a huge collection of interconnected hardware/software components** and **an infrastructure that provides communication services to distributed applications**. Computer Networking- A Top-Down…

---

# 1.1.1 A Nuts-and-Bolts Description

## What is the Internet?

At the most basic level:

> **The Internet is a computer network that interconnects computing devices throughout the world.**

These devices are called **hosts** or **end systems** in Internet terminology. They include traditional computers and servers as well as smartphones, tablets, TVs, gaming consoles, smart appliances, watches, cars, security systems, and other Internet-connected devices. Computer Networking- A Top-Down…

So, conceptually:

```text
             THE INTERNET
                  |
        +---------+---------+
        |                   |
   End Systems         Network Core
   (Hosts)             (Switches)
        |                   |
        +---- Links --------+
                  |
                 ISPs
```

The important components are:

- **End systems / hosts**
- **Communication links**
- **Packet switches**
- **Routers**
- **Link-layer switches**
- **Internet Service Providers (ISPs)**
- **Protocols**

Let's understand each.

---

## 1. End Systems / Hosts

An **end system** is a device at the edge of the Internet that communicates with other devices.

Examples given by the book include:

- Desktop computers
- Servers
- Laptops
- Smartphones
- Tablets
- TVs
- Gaming consoles
- Thermostats
- Home-security systems
- Watches
- Cars
- Other Internet-connected devices

The book uses **host** and **end system** essentially interchangeably in this context. Computer Networking- A Top-Down…

### Why are they called "end systems"?

Because they are located at the **edge of the network** and are the systems where applications actually run.

For example:

```text
Your Laptop  ←→  Internet  ←→  Web Server
   Host                         Host
```

Your laptop and the Web server are both end systems.

---

# 2. Communication Links

End systems don't communicate by being directly connected to every other end system.

Instead, they are connected through **communication links** and **packet switches**. Computer Networking- A Top-Down…

Communication links can use different physical media, including:

- **Coaxial cable**
- **Copper wire**
- **Optical fiber**
- **Radio spectrum**

Different links support different transmission rates.

The transmission rate of a link is measured in:

\[
\boxed{\text{bits/second (bps)}}
\]

For example:

- 10 Mbps
- 100 Mbps
- 1 Gbps
- 10 Gbps

The book's main point here is that **different links can transmit data at different rates**. Computer Networking- A Top-Down…

---

# 3. Packets

Suppose one end system wants to send a large piece of data to another end system.

The sender does not necessarily send the entire data as one giant unit.

Instead:

### Step 1 — Segment the data

The sending end system **breaks/segments the data into smaller pieces**.

### Step 2 — Add header information

The sender adds **header bytes** to each segment.

### Step 3 — Create packets

The resulting units are called **packets**.

### Step 4 — Send packets through the network

The packets travel through communication links and packet switches.

### Step 5 — Reassemble at destination

At the destination end system, the packets are **reassembled into the original data**. Computer Networking- A Top-Down…

Conceptually:

```text
Original Data
     |
     ↓
+----+----+----+----+
| P1 | P2 | P3 | P4 |
+----+----+----+----+
     |
     ↓
   Network
     |
     ↓
+----+----+----+----+
| P1 | P2 | P3 | P4 |
+----+----+----+----+
     |
     ↓
Reassembled Data
```

This idea of **packetization** is fundamental to understanding the Internet.

---

# 4. Packet Switches

A **packet switch** receives a packet on one incoming communication link and forwards it through one of its outgoing links. Computer Networking- A Top-Down…

The book identifies two prominent types:

### 1. Routers
### 2. Link-layer switches

Both forward packets toward their destinations.

However, they are typically used in different parts of the network:

| Component | Typical location |
|---|---|
| Link-layer switch | Access networks |
| Router | Network core |

Computer Networking- A Top-Down…

---

# 5. Route / Path

A packet normally passes through multiple:

- Communication links
- Packet switches

between its source and destination.

The **sequence of communication links and packet switches** traversed by the packet is called its:

\[
\boxed{\text{route or path}}
\]

For example:

```text
Host A
  |
  | Link
  ↓
Router 1
  |
  | Link
  ↓
Router 2
  |
  | Link
  ↓
Router 3
  |
  ↓
Host B
```

The complete sequence from Host A to Host B is the packet's **path through the network**. Computer Networking- A Top-Down…

---

# 6. Transportation Analogy

The book uses a very useful analogy.

Think of a **transportation network**:

```text
Factory → Roads → Intersections → Roads → Warehouse
```

Large cargo is divided among multiple trucks.

Similarly:

```text
Source Host → Links → Packet Switches → Links → Destination Host
```

The correspondence is:

| Transportation network | Computer network |
|---|---|
| Building | End system |
| Cargo | Data |
| Truck | Packet |
| Highway/Road | Communication link |
| Intersection | Packet switch |
| Route | Network path |
| Warehouse | Destination |

The book's point is that just as trucks travel through roads and intersections, **packets travel through links and packet switches**. Computer Networking- A Top-Down…

This analogy becomes especially useful later when studying **packet switching**.

---

# 7. Internet Service Providers — ISPs

How does your laptop or smartphone actually get connected to the Internet?

Through an **Internet Service Provider (ISP)**.

Examples of access providers described by the book include:

- Residential ISPs
- Corporate ISPs
- University ISPs
- Public Wi-Fi providers
- Cellular data ISPs

Each ISP itself consists of:

- Packet switches
- Communication links

ISPs provide different types of access, including residential broadband, high-speed LAN access, and mobile wireless access. Computer Networking- A Top-Down…

For example:

```text
Your Laptop
     |
     ↓
Local / Residential ISP
     |
     ↓
Regional / National ISP
     |
     ↓
Other ISP
     |
     ↓
Destination Network
     |
     ↓
Web Server
```

### Important point

There isn't one single organization that owns the entire Internet.

The Internet consists of **many independently managed networks that are interconnected**.

Lower-tier ISPs connect through national/international upper-tier ISPs. The book describes upper-tier ISPs as having high-speed routers interconnected by high-speed fiber-optic links. Computer Networking- A Top-Down…

This is why the Internet can be thought of as a **network of networks**.

---

# 8. Protocols

All these components need rules to communicate.

Therefore, end systems, packet switches, and other Internet components run **protocols** that control how information is sent and received. Computer Networking- A Top-Down…

Two particularly important Internet protocols are:

- **TCP — Transmission Control Protocol**
- **IP — Internet Protocol**

Together, the principal Internet protocols are commonly referred to as:

\[
\boxed{\text{TCP/IP}}
\]

The book specifically notes that IP specifies the format of packets sent and received among routers and end systems. Computer Networking- A Top-Down…

---

# 9. Internet Standards

For protocols to work, different manufacturers and systems must **agree on what the protocols mean and how they operate**.

Otherwise, one computer might implement a protocol differently from another computer, preventing interoperability.

This is where **standards** are important.

### IETF

Internet standards are developed by the:

**Internet Engineering Task Force (IETF)**

The IETF publishes standards documents called:

\[
\boxed{\text{RFCs — Requests for Comments}}
\]

RFCs define protocols such as:

- TCP
- IP
- HTTP
- SMTP

The book also mentions the **IEEE 802 LAN/MAN Standards Committee**, which specifies standards such as Ethernet and Wi-Fi. Computer Networking- A Top-Down…

---

# Nuts-and-Bolts View — Complete Picture

So, from the hardware/software perspective:

```text
             INTERNET
                 |
       +---------+---------+
       |                   |
   END SYSTEMS         NETWORK CORE
   / HOSTS             /           \
       |             Routers     Links
       |                 |
       +-------- ISPs ---+
                 |
              Protocols
                 |
        TCP / IP / HTTP / ...
```

A more realistic communication path is:

```text
Sender Host
     |
     ↓
Access Network
     |
     ↓
ISP
     |
     ↓
Routers + Links
     |
     ↓
Destination ISP
     |
     ↓
Destination Host
```

That is the **nuts-and-bolts description**: it tells us **what the Internet is made of and how those components are physically/logically connected**.

---

# 1.1.2 A Services Description

Now the book changes perspective.

Instead of asking:

> **"What components make up the Internet?"**

we ask:

> **"What does the Internet provide to applications?"**

The Internet can be viewed as an **infrastructure that provides services to applications**. Computer Networking- A Top-Down…

This is extremely important.

---

## Internet as an Application Platform

The Internet supports many applications, including:

- E-mail
- Web browsing
- Internet messaging
- Mapping applications
- Real-time traffic information
- Music streaming
- Movie/TV streaming
- Social networking
- Video conferencing
- Multiplayer games
- Location-based recommendation systems

These are called **distributed applications** because they involve multiple end systems that exchange data with one another. Computer Networking- A Top-Down…

For example:

```text
Your Phone
    |
    | Data
    ↓
Internet
    |
    | Data
    ↓
Application Server
```

The application is distributed across multiple end systems.

---

# Very Important: Applications Run on End Systems

The book emphasizes an important distinction:

> **Internet applications run on end systems, not inside the packet switches in the network core.**

For example, suppose you're using a Web browser.

The browser runs on your computer.

The Web server application runs on a server.

Routers in between simply help move the packets.

Conceptually:

```text
Application        Application
   ↓                  ↑
Your PC              Server
   |                  |
   +------ Internet --+
          |
       Routers
```

The routers facilitate the exchange of data but **do not run the application itself**. Computer Networking- A Top-Down…

This distinction is important for understanding the Internet's architecture.

---

# How Does One Application Communicate With Another?

Suppose you create a new distributed application.

You write programs that run on different end systems.

For example:

```text
Computer A                     Computer B
+-----------+                  +-----------+
| Program A |                  | Program B |
+-----------+                  +-----------+
      |                              ↑
      |          Internet            |
      +------------------------------>
```

Now the question becomes:

> **How does Program A tell the Internet to deliver data to Program B?**

The book introduces the **socket interface** to answer this question. Computer Networking- A Top-Down…

---

# Socket Interface

An end system attached to the Internet provides a **socket interface**.

The socket interface specifies how a program running on one end system asks the Internet infrastructure to deliver data to a particular destination program on another end system.

In simple terms:

\[
\boxed{\text{Application} \rightarrow \text{Socket Interface} \rightarrow \text{Internet}}
\]

The application doesn't directly control all the routers and links.

Instead, it uses the defined interface to request network communication.

The book describes this interface as a **set of rules that the sending program must follow** so that the Internet can deliver data to the destination program. Computer Networking- A Top-Down…

---

# Postal Service Analogy

The book gives an excellent analogy.

Suppose **Alice wants to send a letter to Bob**.

Alice cannot simply throw the letter somewhere and expect the postal system to understand it.

She must follow the postal service's rules:

1. Put the letter in an envelope.
2. Write Bob's name.
3. Write Bob's address.
4. Write the ZIP code.
5. Seal the envelope.
6. Add the stamp.
7. Put it into the appropriate mailbox.

These rules form the **postal service interface**.

Similarly, an Internet application must follow the Internet's **socket interface** to have its data delivered to another application. Computer Networking- A Top-Down…

### Analogy

```text
POSTAL SYSTEM                  INTERNET

Letter                         Data
   ↓                              ↓
Envelope                       Packet
   ↓                              ↓
Address                        Destination
   ↓                              ↓
Postal rules                   Socket interface
   ↓                              ↓
Postal network                 Internet
   ↓                              ↓
Bob                            Destination program
```

---

# Internet Provides Multiple Services

The postal system doesn't provide only one service.

It can provide things such as:

- Ordinary delivery
- Express delivery
- Reception confirmation

Similarly, the Internet provides **multiple services to applications**.

When developing an Internet application, the application chooses the Internet service appropriate for its needs. Computer Networking- A Top-Down…

The details of these transport services are developed later in the book, particularly in the Application and Transport Layer chapters.

---

# 1.1.3 What Is a Protocol?

Now we reach one of the **most important definitions in the entire chapter**.

A protocol is essentially a **set of rules governing communication between entities**.

But Kurose & Ross explain this concept using a human analogy first.

---

# Human Protocol

Imagine you want to ask someone the time.

A typical interaction might be:

```text
You:       Hi
Person:    Hi
You:       Got the time?
Person:    2:00
```

There are rules governing this interaction.

You generally:

1. Initiate communication with **Hi**.
2. Wait for the response.
3. If the response is appropriate, continue.
4. Ask your question.
5. Receive the answer.

If the other person responds with something indicating that they don't want to communicate, you don't proceed normally.

So a human interaction involves:

- Messages
- Order of messages
- Responses
- Actions based on responses
- Actions when no response occurs

The book uses this to build the concept of a computer-network protocol. Computer Networking- A Top-Down…

---

# Network Protocol

In a **network protocol**, the entities communicating are not necessarily humans.

They are hardware or software components such as:

- Computers
- Smartphones
- Tablets
- Routers
- Network interfaces
- Other network-capable devices

The book states that essentially all Internet activity involving two or more communicating remote entities is governed by a protocol. Computer Networking- A Top-Down…

Examples:

### Protocol at network interface

Hardware-implemented protocols can control the flow of bits between two physically connected computers.

### Congestion-control protocols

They control how quickly packets are transmitted between sender and receiver.

### Routing protocols

Protocols in routers help determine a packet's path from source to destination.

Thus:

```text
Protocols are everywhere in the Internet.
```

Computer Networking- A Top-Down…

---

# Web Request Example of a Protocol

The book gives a Web request as an example.

Suppose you enter a Web-page URL into your browser.

A simplified sequence is:

### Step 1 — Connection request

Your computer sends a **connection request** to the Web server.

```text
Client → Connection Request → Server
```

### Step 2 — Connection reply

The server receives it and sends a reply.

```text
Client ← Connection Reply ← Server
```

### Step 3 — Request the Web document

Once the client knows that it can proceed, it sends a **GET** message requesting the Web page.

```text
Client → GET /page → Server
```

### Step 4 — Server responds

The server sends the requested Web page/file back.

```text
Client ← Web Page ← Server
```

The important thing is not merely the messages themselves.

The protocol also specifies **what each side should do when a particular message is transmitted or received**. Computer Networking- A Top-Down…

---

# The Formal Definition of a Protocol

This is the **key definition to remember for exams**:

> **A protocol defines the format and the order of messages exchanged between two or more communicating entities, as well as the actions taken on the transmission and/or receipt of a message or other event.** Computer Networking- A Top-Down…

Break this definition into three parts:

### 1. Format

**What does the message look like?**

Example:

```text
Message = Header + Data
```

The protocol determines the structure.

### 2. Order

**Which message comes first, second, third, etc.?**

For example:

```text
Request
   ↓
Reply
   ↓
Data Request
   ↓
Data Response
```

### 3. Actions

**What should an entity do after sending or receiving a message?**

For example:

```text
Receive Request
      ↓
Check Request
      ↓
Send Reply
      ↓
Process Next Message
```

So:

\[
\boxed{
\text{Protocol}
=
\text{Message Format}
+
\text{Message Order}
+
\text{Actions}
}
\]

This is the core idea.

---

# Why Are Protocols Necessary?

Imagine two computers communicating but using different rules.

Computer A thinks:

```text
Message 1 → Message 2 → Message 3
```

while Computer B expects:

```text
Message 3 → Message 1 → Message 2
```

They cannot communicate effectively.

The book makes the same point with the human analogy: if communicating entities follow different protocols, the protocols **do not interoperate**, and useful communication cannot take place. Computer Networking- A Top-Down…

Therefore:

\[
\boxed{\text{Both communicating entities must understand/use the same protocol.}}
\]

This is fundamental to interoperability.

---

# How the Three Concepts Fit Together

Now combine the entire section.

## 1. Nuts-and-Bolts View

Answers:

> **What is the Internet made of?**

It consists of:

```text
End Systems
     ↓
Communication Links
     ↓
Packet Switches
     ↓
Routers / Link-Layer Switches
     ↓
ISPs
     ↓
Protocols
```

Computer Networking- A Top-Down…

---

## 2. Services View

Answers:

> **What does the Internet do for applications?**

It provides infrastructure that allows distributed applications running on different end systems to exchange data.

```text
Application A
     ↓
Socket Interface
     ↓
Internet
     ↓
Socket Interface
     ↓
Application B
```

Computer Networking- A Top-Down…

---

## 3. Protocol View

Answers:

> **How do all these entities communicate correctly?**

Through protocols that specify:

```text
        PROTOCOL
           |
   +-------+-------+
   |       |       |
 Format   Order   Actions
```

Computer Networking- A Top-Down…

---

# 🔥 One Complete Example

Suppose you open a Web page.

### At the application level

Your browser wants a Web page from a server.

```text
Browser → Internet → Web Server
```

### At the services level

The browser uses the Internet's communication services through its **socket interface** to request delivery of data to the destination application. Computer Networking- A Top-Down…

### At the nuts-and-bolts level

The data travels through:

```text
Your Computer
     ↓
Access Network
     ↓
ISP
     ↓
Routers
     ↓
Communication Links
     ↓
Destination Network
     ↓
Web Server
```

The data is divided into **packets**, and packet switches forward those packets along a path. Computer Networking- A Top-Down…

### At the protocol level

Protocols govern the communication.

For example:

- **HTTP** → Web/application communication
- **TCP** → transport communication
- **IP** → packet delivery across the network

The book identifies TCP and IP as two of the most important Internet protocols and collectively refers to the principal Internet protocols as **TCP/IP**. Computer Networking- A Top-Down…

---

# 📌 Exam-Oriented Summary

### What is the Internet?

The Internet can be described in **two ways**:

### **1. Nuts-and-Bolts Description**
The Internet is a computer network consisting of interconnected:

- Hosts/end systems
- Communication links
- Packet switches
- Routers
- Link-layer switches
- ISPs
- Protocols

Data is segmented into packets, which travel through communication links and packet switches along a route/path to the destination. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

### **2. Services Description**
The Internet is an infrastructure that provides communication services to **distributed applications** running on end systems. Applications communicate by using interfaces such as the Internet's **socket interface**. Computer Networking- A Top-Down…

### **3. Protocol**
A network protocol defines:

1. **Format** of messages
2. **Order** of messages
3. **Actions** taken when messages are sent/received or events occur

Computer Networking- A Top-Down…

### The simplest way to remember it:

\[
\boxed{
\text{Internet}
=
\text{Infrastructure}
+
\text{Services}
+
\text{Protocols}
}
\]

**Infrastructure** tells us **what the Internet is made of**.  
**Services** tell us **what the Internet provides to applications**.  
**Protocols** tell us **how the communicating entities follow rules to make communication possible**.

# 1.2 The Network Edge

In Kurose & Ross, **The Network Edge** is the first part of the Internet that we study after the general overview.

The authors begin at the **edge of the network**, where the end systems and applications are located. They then explain how these end systems connect to the first router through **access networks**, and finally discuss the **physical media** that carry the bits. Computer Networking- A Top-Down…

The section has two main parts:

1. **1.2.1 Access Networks**
2. **1.2.2 Physical Media** Computer Networking- A Top-Down…

---

# 1.2.1 Access Networks

## What is an Access Network?

An **access network** is the network that **physically connects an end system to the first router**, also called the **edge router**, on the path to another distant end system. Computer Networking- A Top-Down…

So, if you are using a laptop at home:

```text
Laptop
   |
   | Access Network
   ↓
Edge Router
   |
   ↓
Internet Core
   |
   ↓
Destination
```

The access network is the portion between your end system and the first router that connects you toward the Internet.

### Why is it called the "edge"?

Because end systems are located at the **edge of the Internet**, while routers and links that interconnect these end systems form the **network core**.

```text
             INTERNET CORE
       Routers + Communication Links
                  |
       -------------------------
       |           |           |
   Edge Router Edge Router Edge Router
       |           |           |
      / \         / \         / \
    Hosts       Hosts       Hosts
```

The book divides access networks primarily according to where they are used:

- **Home access**
- **Enterprise access**
- **Wide-area wireless access** Computer Networking- A Top-Down…

---

# A. Home Access

The book discusses these major home-access technologies:

1. **DSL**
2. **Cable**
3. **FTTH**
4. **Dial-up**
5. **Satellite** Computer Networking- A Top-Down…

Let's understand each.

---

# 1. DSL — Digital Subscriber Line

DSL uses the **existing telephone infrastructure** to provide Internet access.

A residence typically gets DSL from the same telephone company that provides its wired telephone service. Computer Networking- A Top-Down…

### Basic structure

```text
Home
+----------------+
| PC             |
| DSL Modem      |
+-------+--------+
        |
        | Twisted-pair copper
        |
        ↓
      DSLAM
        |
        ↓
   Telephone Company
        |
        ↓
     Internet
```

### Important components

#### DSL Modem

The modem is located at the customer's home.

It takes digital data and converts it into high-frequency signals suitable for transmission over the telephone line.

#### DSLAM

At the telephone company's central office (**CO**) is a:

**Digital Subscriber Line Access Multiplexer (DSLAM)**

The DSLAM receives signals from many homes and separates the data from the telephone signals. It then sends the data toward the Internet. Hundreds or even thousands of households can connect to a single DSLAM. Computer Networking- A Top-Down… Computer Networking- A Top-Down…

---

## DSL Uses Different Frequency Bands

One important feature of DSL is that the same telephone line can carry **telephone communication and Internet data simultaneously**.

The book describes three frequency ranges:

| Channel | Frequency range | Purpose |
|---|---:|---|
| Downstream | 50 kHz – 1 MHz | Internet data toward home |
| Upstream | 4 kHz – 50 kHz | Internet data toward ISP |
| Telephone | 0 – 4 kHz | Ordinary telephone |

Computer Networking- A Top-Down…

Therefore, one physical telephone line behaves as though it contains three separate logical channels.

```text
                 DSL Telephone Line
                       |
       +---------------+---------------+
       |               |               |
   0–4 kHz          4–50 kHz       50 kHz–1 MHz
       |               |               |
   Telephone       Upstream        Downstream
```

This allows:

> **Telephone call + Internet connection at the same time.**

The book identifies this as an example of **frequency-division multiplexing**. Computer Networking- A Top-Down…

---

## DSL Is Asymmetric

The downstream rate is generally higher than the upstream rate.

Therefore:

\[
\boxed{\text{DSL is asymmetric}}
\]

The book gives examples of DSL standards including:

- 12 Mbps downstream / 1.8 Mbps upstream
- 55 Mbps downstream / 15 Mbps upstream

Actual rates can be lower because the provider may limit the rate and because performance depends on:

- Distance between home and central office
- Gauge/thickness of the twisted-pair line
- Electrical interference

DSL is designed for relatively short distances; the book notes that beyond roughly 5–10 miles from the central office, another access technology may be necessary. Computer Networking- A Top-Down…

---

# 2. Cable Internet Access

Cable Internet uses the existing **cable television infrastructure**.

The physical network is a combination of:

- **Fiber optics**
- **Coaxial cable**

Therefore, it is called:

\[
\boxed{\text{HFC — Hybrid Fiber Coax}}
\]

Computer Networking- A Top-Down…

### Structure

```text
                   Internet
                      |
                    CMTS
                      |
                 Cable Head End
                      |
                  Fiber Cable
                      |
                Fiber Node
                 /   |   \
                /    |    \
          Coaxial  Coaxial  Coaxial
             |        |        |
           Home     Home     Home
```

Fiber connects the **cable head end** to neighborhood-level junctions.

From those junctions, **coaxial cable** connects to individual homes. Computer Networking- A Top-Down…

A neighborhood junction typically supports **500–5,000 homes** according to the book. Computer Networking- A Top-Down…

---

## Cable Modem and CMTS

The home uses a **cable modem**.

At the cable head end is a:

**CMTS — Cable Modem Termination System**

The CMTS performs a role similar to the DSLAM in DSL networks: it converts the signals coming from many cable modems back into digital format. Computer Networking- A Top-Down…

---

## Downstream and Upstream

Cable Internet has:

- **Downstream channel**
- **Upstream channel**

Typically:

\[
\text{Downstream rate} > \text{Upstream rate}
\]

So cable Internet access is also generally **asymmetric**. Computer Networking- A Top-Down…

---

## Most Important Property: Cable is Shared

This is an important exam point.

Cable Internet is a **shared broadcast medium**.

Every packet transmitted by the cable head end travels downstream over the shared channel toward the homes.

Similarly, packets sent by homes travel upstream toward the cable head end. Computer Networking- A Top-Down…

Suppose the cable network has 100 Mbps downstream capacity.

If only one user is actively downloading:

```text
100 Mbps
   |
   ↓
 User A
```

The user can potentially receive close to the available rate.

But if many users simultaneously download:

```text
             100 Mbps
                |
       +--------+--------+
       |        |        |
      A        B        C
```

the aggregate capacity is shared among them, so each user's effective rate can decrease significantly. Computer Networking- A Top-Down…

### Upstream collision issue

Because the upstream channel is also shared, a **distributed multiple-access protocol** is required to coordinate transmissions and avoid collisions. Computer Networking- A Top-Down…

This becomes important later when the book discusses **multiple-access protocols**.

---

# 3. FTTH — Fiber to the Home

**FTTH** stands for:

\[
\boxed{\text{Fiber To The Home}}
\]

The idea is straightforward:

> Provide an optical-fiber path directly from the central office to the home.

Computer Networking- A Top-Down…

FTTH can provide Internet access in the **gigabits-per-second range**. Computer Networking- A Top-Down…

---

## FTTH Architectures

The book describes two optical-distribution architectures:

### 1. Active Optical Network — AON

AON is essentially **switched Ethernet**.

### 2. Passive Optical Network — PON

PON is the architecture discussed in more detail in this section. Computer Networking- A Top-Down…

---

## PON Structure

In a PON:

```text
                    Central Office
                         |
                        OLT
                         |
                 Shared Optical Fiber
                         |
                    Splitter
                  /    |    \
                 /     |     \
               ONT    ONT    ONT
                |      |      |
              Home   Home   Home
```

### ONT

Each home has an:

**Optical Network Terminator (ONT)**

The ONT connects to the neighborhood optical splitter using dedicated optical fiber.

### Optical Splitter

The splitter combines the traffic from multiple homes onto a shared optical fiber toward the central office.

### OLT

At the central office is the:

**Optical Line Terminator (OLT)**

The OLT converts between optical and electrical signals and connects toward the Internet through a telco router. Computer Networking- A Top-Down…

---

## FTTH Capacity

FTTH can potentially provide **gigabit-per-second** Internet access.

The book emphasizes that actual ISP offerings depend on the service plan, with higher rates generally costing more. Computer Networking- A Top-Down…

---

# 4. Dial-Up

Dial-up Internet uses traditional telephone lines.

The basic model is:

```text
Home Modem
     |
Telephone Line
     |
ISP Modem
     |
Internet
```

It is conceptually similar to DSL in that a home modem connects over the telephone line to equipment at the ISP.

However, dial-up is **extremely slow compared with broadband technologies**.

The book gives:

\[
\boxed{56\text{ kbps}}
\]

as the dial-up rate. Computer Networking- A Top-Down…

---

# 5. Satellite Access

Satellite Internet can be used where technologies such as:

- DSL
- Cable
- FTTH

are unavailable, particularly in some rural locations. Computer Networking- A Top-Down…

The book states that satellite access can provide speeds of **more than 1 Mbps** in the context discussed there. Computer Networking- A Top-Down…

Satellite communication is discussed more fully under physical media.

---

# B. Enterprise Access — Ethernet and WiFi

In:

- Corporate networks
- University networks
- Increasingly, home networks

a **LAN (Local Area Network)** is used to connect an end system to the edge router. Computer Networking- A Top-Down…

---

# Ethernet

The book identifies **Ethernet** as by far the most prevalent access technology in corporate, university, and home networks. Computer Networking- A Top-Down…

Typical structure:

```text
PC ───────┐
Laptop ───┤
          ↓
    Ethernet Switch
          |
          ↓
   Institutional Router
          |
          ↓
       ISP
          |
       Internet
```

Ethernet users typically use **twisted-pair copper wire** to connect to an Ethernet switch. Computer Networking- A Top-Down…

The book gives typical access rates of:

- **100 Mbps**
- **1 Gbps**

for users, while servers may have:

- **1 Gbps**
- **10 Gbps**

access. Computer Networking- A Top-Down…

---

# WiFi

Wireless LAN users connect to an **access point**.

```text
Laptop )))))
          \
Phone ))))) → WiFi Access Point
          /
Tablet)))))
              |
              ↓
       Wired Ethernet
              |
              ↓
        Institutional ISP
```

The access point connects into the organization's wired network, typically using Ethernet. Computer Networking- A Top-Down…

The technology is:

\[
\boxed{\text{IEEE 802.11 = WiFi}}
\]

WiFi users typically need to be within **a few tens of meters** of the access point. Computer Networking- A Top-Down…

The book notes that 802.11 provides a shared transmission rate of **more than 100 Mbps** in the context of the edition. Computer Networking- A Top-Down…

---

# C. Wide-Area Wireless Access — 3G and LTE

Mobile devices such as smartphones use cellular infrastructure to send and receive Internet packets through a **base station** operated by the cellular provider.

Unlike WiFi, the coverage extends much farther:

- WiFi → typically tens of meters
- Cellular → potentially tens of kilometers

Computer Networking- A Top-Down…

### 3G

The book describes 3G as providing **packet-switched wide-area wireless Internet access** at speeds exceeding 1 Mbps. Computer Networking- A Top-Down…

### 4G / LTE

LTE stands for:

\[
\boxed{\text{Long-Term Evolution}}
\]

The book describes LTE as a fourth-generation wide-area wireless technology with rates exceeding 10 Mbps, with commercial deployments reporting downstream rates of many tens of Mbps. Computer Networking- A Top-Down…

---

# Access Networks — Important Comparison

| Technology | Main medium | Typical setting | Important property |
|---|---|---|---|
| DSL | Twisted-pair copper | Home | Asymmetric; telephone + Internet simultaneously |
| Cable/HFC | Fiber + coaxial | Home | Shared broadcast medium |
| FTTH | Optical fiber | Home | Very high potential rate |
| Dial-up | Telephone twisted pair | Home | Very slow, ~56 kbps |
| Satellite | Radio | Home/rural | Useful where wired broadband unavailable |
| Ethernet | Copper/fiber | Enterprise/home | Dominant wired LAN access |
| WiFi | Radio | Enterprise/home | Wireless LAN; short range |
| 3G | Radio | Wide-area/mobile | Cellular Internet |
| LTE | Radio | Wide-area/mobile | Higher-speed cellular access |

The book's review questions explicitly ask students to identify and classify six access technologies into **home, enterprise, and wide-area wireless access**, so this classification is important for exams. Computer Networking- A Top-Down…

---

# 1.2.2 Physical Media

Now we move one level deeper.

An access network tells us **how an end system gets connected**.

Physical media tells us:

> **What actually carries the bits between the transmitter and receiver?**

The book explains this using the journey of a single bit.

Suppose a bit travels:

```text
Source
  ↓
Router 1
  ↓
Router 2
  ↓
Router 3
  ↓
Destination
```

At every hop, the bit is transmitted from one device to another through a **physical medium**.

The transmission occurs through:

- **Electromagnetic waves**, or
- **Optical pulses**

depending on the medium. Computer Networking- A Top-Down…

---

# Guided vs Unguided Media

The book divides physical media into **two categories**.

## 1. Guided Media

The signal is guided through a **solid physical medium**.

Examples:

- Twisted-pair copper wire
- Coaxial cable
- Fiber-optic cable

## 2. Unguided Media

The signal propagates through:

- Atmosphere
- Outer space

Examples:

- Wireless LAN radio
- Terrestrial radio
- Satellite radio

Computer Networking- A Top-Down…

### Remember:

```text
Physical Media
     |
     +-------------------+
     |                   |
  Guided              Unguided
     |                   |
  Wire/Cable           Wireless
     |                   |
 Copper/Coax/Fiber    Radio/Satellite
```

---

# 1. Twisted-Pair Copper Wire

This is the **least expensive and most commonly used guided transmission medium** discussed by the book. Computer Networking- A Top-Down…

It consists of:

> **Two insulated copper wires twisted together.**

Each wire is approximately 1 mm thick.

The wires are twisted in a spiral pattern primarily to **reduce electrical interference** from nearby wire pairs. Computer Networking- A Top-Down…

---

## UTP — Unshielded Twisted Pair

**UTP** is commonly used for computer networks within buildings.

The book gives LAN data rates ranging from:

\[
10\text{ Mbps} \rightarrow 10\text{ Gbps}
\]

depending on factors including:

- Wire thickness
- Distance between transmitter and receiver Computer Networking- A Top-Down…

Modern twisted-pair technology such as **Category 6a** can achieve **10 Gbps for distances up to 100 meters**, according to the book. Computer Networking- A Top-Down…

Twisted pair is used for:

- Ethernet LANs
- DSL
- Dial-up

Computer Networking- A Top-Down…

---

# 2. Coaxial Cable

Coaxial cable also contains **two copper conductors**, but unlike twisted pair, they are **concentric** rather than parallel. Computer Networking- A Top-Down…

Its construction includes special insulation and shielding, allowing high data transmission rates.

It is commonly used in:

- Cable television
- Cable Internet

The book notes that coaxial cable can operate as a **guided shared medium**.

That means multiple end systems can be connected to the same cable and receive transmissions sent by other end systems. Computer Networking- A Top-Down…

This is directly related to why cable Internet has a **shared-medium** characteristic.

---

# 3. Fiber Optics

An optical fiber is a **thin, flexible medium that carries pulses of light**.

Each light pulse represents a bit. Computer Networking- A Top-Down…

Conceptually:

```text
Bit 1 → Light pulse
Bit 0 → No/light-level difference
```

A single optical fiber can support extremely high bit rates, potentially **tens or hundreds of Gbps** according to the book. Computer Networking- A Top-Down…

---

## Advantages of Fiber

The book emphasizes three major characteristics:

### 1. High capacity

Very high bit rates.

### 2. Immunity to electromagnetic interference

Electrical interference does not affect optical transmission in the same way it affects copper.

### 3. Low attenuation

The book notes very low signal attenuation over distances up to approximately **100 km** in the context discussed.

### 4. Difficult to tap

Fiber is relatively difficult to tap, which provides an advantage in certain security contexts. Computer Networking- A Top-Down…

Because of these characteristics, fiber is the preferred **long-haul guided transmission medium**, especially for overseas links, and is extensively used in the Internet backbone. Computer Networking- A Top-Down…

---

# 4. Terrestrial Radio Channels

Radio channels transmit signals using the **electromagnetic spectrum**.

They are attractive because they:

- Do not require physical wires
- Can penetrate walls
- Can support mobile users
- Can potentially carry signals over long distances Computer Networking- A Top-Down…

But radio transmission has important challenges.

Its characteristics depend strongly on:

- Propagation environment
- Distance
- Obstacles
- Other transmissions

The book identifies:

### Path loss

Signal strength decreases as the signal travels.

### Shadow fading

Signal strength can decrease because of objects obstructing the path.

### Multipath fading

Signals can reflect from objects and arrive through multiple paths.

### Interference

Other transmissions and electromagnetic signals can interfere with the communication. Computer Networking- A Top-Down…

---

## Three broad terrestrial radio ranges

The book classifies terrestrial radio channels into:

| Range | Example |
|---|---|
| Very short distance | Wireless headsets, keyboards, medical devices |
| Local area | WiFi |
| Wide area | Cellular networks |

The ranges are roughly:

- **1–2 meters** for very short-range systems
- **10–few hundred meters** for local-area systems
- **Tens of kilometers** for wide-area systems Computer Networking- A Top-Down…

---

# 5. Satellite Radio Channels

A communication satellite connects two or more **Earth-based microwave transmitter/receivers**, called **ground stations**.

The satellite:

1. Receives a signal on one frequency band.
2. Regenerates the signal using a repeater.
3. Transmits it on another frequency. Computer Networking- A Top-Down…

There are two types discussed:

- **Geostationary satellites**
- **Low-Earth Orbit (LEO) satellites**

---

## Geostationary Satellites

A geostationary satellite remains above approximately the same point on Earth.

The book describes the orbit as approximately:

\[
36,000\text{ km}
\]

above Earth's surface. Computer Networking- A Top-Down…

The large distance causes a significant propagation delay.

The book gives approximately:

\[
\boxed{280\text{ ms}}
\]

for the ground → satellite → ground propagation path discussed. Computer Networking- A Top-Down…

Despite this delay, satellite links can operate at **hundreds of Mbps** and are useful in areas without DSL or cable Internet access. Computer Networking- A Top-Down…

---

## LEO — Low-Earth Orbit Satellites

LEO satellites:

- Are much closer to Earth
- Do not remain above one fixed location
- Orbit around Earth
- May communicate with ground stations and with other satellites

Because a single LEO satellite does not provide permanent coverage of one area, **many satellites are needed for continuous coverage**. Computer Networking- A Top-Down…

---

# 🔥 Access Network vs Physical Media

This distinction is very important.

### Access Network

Answers:

> **How is an end system connected to the Internet?**

Examples:

- DSL
- Cable
- FTTH
- Ethernet
- WiFi
- 3G/LTE

### Physical Media

Answers:

> **What physical/electromagnetic medium actually carries the signal?**

Examples:

- Twisted-pair copper
- Coaxial cable
- Fiber optics
- Terrestrial radio
- Satellite radio

For example:

```text
DSL
 ↓
Access Network Technology
 ↓
Twisted-Pair Copper
 ↓
Physical Medium
```

Another example:

```text
Cable Internet
 ↓
Access Network Technology
 ↓
Fiber + Coaxial Cable
 ↓
Physical Media
```

And:

```text
WiFi
 ↓
Access Network Technology
 ↓
Radio Spectrum
 ↓
Physical Medium
```

The book explicitly connects these technologies with their underlying media: HFC uses fiber + coaxial cable, DSL/Ethernet use copper wire, and mobile access uses radio spectrum. Computer Networking- A Top-Down…

---

# 🧠 Complete Mental Model

Think of the Internet as:

```text
                 NETWORK CORE
          Routers + Long-haul Links
                     |
        -----------------------------
        |             |             |
     Edge Router   Edge Router   Edge Router
        |             |             |
   ACCESS NETWORKS
        |
  +-----+---------+-----------+
  |     |         |           |
 DSL  Cable      FTTH      WiFi/LTE
  |     |         |           |
Copper HFC       Fiber       Radio
  |   /   \        |           |
  | Fiber Coax     |           |
  +-----+----------+-----------+
        |
   END SYSTEMS
 PC / Laptop / Phone / Server
```

---

# 🎯 What You Must Remember for Exams

## Access Networks

### Home
- **DSL** → twisted-pair telephone line; asymmetric; DSLAM
- **Cable/HFC** → fiber + coax; shared broadcast medium; CMTS
- **FTTH** → optical fiber to home; very high capacity; AON/PON
- **Dial-up** → telephone line; 56 kbps
- **Satellite** → useful where wired broadband isn't available

### Enterprise/Home LAN
- **Ethernet** → dominant wired LAN technology
- **WiFi/802.11** → wireless LAN; access point; short range

### Wide-area wireless
- **3G**
- **LTE/4G**
- Cellular base stations provide connectivity over much larger areas than WiFi. Computer Networking- A Top-Down…

---

## Physical Media

### Guided
1. **Twisted-pair copper**
2. **Coaxial cable**
3. **Fiber optics**

### Unguided
4. **Terrestrial radio**
5. **Satellite radio**

Computer Networking- A Top-Down…

---

## ⭐ High-value exam distinctions

| Concept | Key point |
|---|---|
| **DSL** | Dedicated connection to DSLAM; asymmetric |
| **Cable/HFC** | **Shared** broadcast medium |
| **FTTH** | Fiber directly toward home; very high capacity |
| **Ethernet** | Dominant wired LAN access |
| **WiFi** | Wireless LAN using 802.11 |
| **Twisted pair** | Cheap, common copper medium |
| **Coaxial** | Copper, shielded, can be shared |
| **Fiber** | Highest capacity; long-haul; low attenuation |
| **Terrestrial radio** | Wireless; affected by fading/interference |
| **Satellite** | Large coverage; geostationary systems have significant propagation delay |

# 1.3 The Network Core

The **network core** is the collection of **packet switches and communication links** that interconnect the end systems at the network edge.

The basic job of the network core is:

> **Move data from a source end system to a destination end system through a sequence of communication links and packet switches.**

The book divides this section into three parts:

1. **Packet Switching**
2. **Circuit Switching**
3. **A Network of Networks**

Computer Networking- A Top-Down…

---

# 1.3.1 Packet Switching

## 1. What is packet switching?

In packet switching, a source divides a long message into smaller units called **packets**.

These packets then travel through:

**Source → Packet switch → Packet switch → ... → Packet switch → Destination**

The two predominant types of packet switches mentioned by the book are:

- **Routers**
- **Link-layer switches**

Each packet is transmitted over each link at the **full transmission rate of that link**. If a packet contains \(L\) bits and the link has transmission rate \(R\) bits/sec, the **transmission delay** for that packet on that link is:

\[
d_{trans}=\frac{L}{R}
\]

Computer Networking- A Top-Down…

### Example

Suppose:

- Packet size = \(L=1000\) bits
- Link rate = \(R=1\) Mbps

Then:

\[
d_{trans}=\frac{1000}{1,000,000}=0.001\text{ sec}=1\text{ ms}
\]

So it takes **1 ms to push all 1000 bits onto the link**.

---

# 2. Store-and-Forward Transmission

This is one of the most important concepts in packet switching.

Most packet switches use **store-and-forward transmission**.

It means:

> A packet switch must receive the **entire packet** before it can begin transmitting that packet onto the outgoing link.

Computer Networking- A Top-Down…

### How it works

Suppose:

```text
Source ───── Link 1 ───── Router ───── Link 2 ───── Destination
```

The source starts transmitting a packet.

### Step 1
The source transmits the packet onto Link 1.

The router **cannot immediately forward** the first bits.

It must first receive the **complete packet**.

### Step 2
After receiving the entire packet, the router stores it.

### Step 3
The router begins transmitting the complete packet onto Link 2.

Therefore, with two links of rate \(R\), and ignoring propagation delay:

\[
d_{end-to-end}=\frac{L}{R}+\frac{L}{R}
\]

\[
\boxed{d_{end-to-end}=\frac{2L}{R}}
\]

Computer Networking- A Top-Down…

---

## Store-and-forward over \(N\) links

If a packet travels through \(N\) links, where every link has transmission rate \(R\):

\[
\boxed{d_{end-to-end}=N\frac{L}{R}}
\]

This is the book's Equation 1.1 when propagation, queuing, and processing delays are ignored. Computer Networking- A Top-Down…

### Example

Suppose:

- \(L=1000\) bits
- \(R=1\) Mbps
- \(N=4\) links

Then:

\[
d=\frac{4(1000)}{1,000,000}
\]

\[
\boxed{d=4\text{ ms}}
\]

---

# 3. Queuing Delay

A router usually has **multiple links** attached to it.

For each outgoing link, the router has an **output buffer**, also called an **output queue**.

Its purpose is to temporarily store packets waiting to be transmitted.

Computer Networking- A Top-Down…

Consider:

```text
             ┌───────────────┐
Packets ───► │    Router     │ ───► Output link
             │               │
             │ Output buffer │
             └───────────────┘
```

Suppose the outgoing link is currently transmitting another packet.

A newly arriving packet cannot immediately use the link.

Therefore:

```text
Incoming packet
      ↓
   Buffer/Queue
      ↓
Outgoing link
```

The packet experiences **queuing delay**.

### Important property

Unlike transmission delay, queuing delay is **variable**.

It depends on the level of congestion in the network.

- Low traffic → little or no queuing
- Heavy traffic → long queue → large delay

Computer Networking- A Top-Down…

---

# 4. Packet Loss

The output buffer of a router has **finite capacity**.

Therefore, if packets arrive while the buffer is already full, there is nowhere to store the new packet.

The result is **packet loss**.

The arriving packet—or one of the packets already waiting in the queue—is dropped. Computer Networking- A Top-Down…

So:

```text
Packets arrive
     ↓
Output buffer
     ↓
 ┌───────────────┐
 │ Packet Packet │
 │ Packet Packet │
 │ Packet Packet │
 └───────────────┘
     ↓
Buffer full
     ↓
New packet → DROPPED
```

This is why packet-switched networks can experience both:

- **Queuing delay**
- **Packet loss**

---

# 5. Forwarding in a Packet-Switched Network

A router receives a packet on one of its incoming links and must determine:

> **Which outgoing link should this packet be sent through?**

In the Internet, the router uses a **forwarding table**.

The forwarding table tells the router which outgoing link to use for a packet based on information in the packet.

The book then distinguishes the actual forwarding operation from the routing mechanisms that help create the forwarding information. Computer Networking- A Top-Down…

Conceptually:

```text
                 Router
                   │
        ┌──────────┼──────────┐
        ↓          ↓          ↓
      Link 1     Link 2     Link 3
```

The router examines the packet and selects the appropriate outgoing link.

---

# 6. Statistical Multiplexing

This is a **very important reason packet switching is efficient**.

Packet switching uses **statistical multiplexing**.

The basic idea is:

> Link capacity is allocated **on demand**, packet by packet, only to users who currently have packets to transmit.

Computer Networking- A Top-Down…

Consider a 1 Mbps link shared by several users.

If only User A has packets:

```text
User A ─────┐
User B ─────┤
User C ─────┼──► 1 Mbps link
User D ─────┘
```

User A can use essentially the **full 1 Mbps** while the other users are inactive.

When another user becomes active, the link capacity is shared according to the packets that need transmission.

---

## Why is this efficient?

Internet traffic is often **bursty**.

A user might:

- transmit for a short period,
- then remain idle,
- then transmit again.

If we permanently reserve capacity for every user, much of that capacity would remain unused.

Packet switching avoids this by allocating the link **when packets actually exist**.

The book explicitly contrasts this with circuit switching:

- **Circuit switching:** pre-allocates link use regardless of demand.
- **Packet switching:** allocates link use on demand.

Computer Networking- A Top-Down…

---

# 7. Why Packet Switching Can Be Better Than Circuit Switching

The book gives two important examples.

### Example 1 — Bursty users

Suppose:

- Link capacity = 1 Mbps
- Each user generates data at 100 kbps
- Each user is active only 10% of the time

With circuit switching:

\[
\frac{1\text{ Mbps}}{100\text{ kbps}}=10
\]

Only **10 users** can be supported because capacity must be reserved for each user.

With packet switching, many more users can share the link because they are unlikely to all be active simultaneously.

The book gives an example with **35 users**, where the probability of 11 or more users being simultaneously active is approximately 0.0004. Computer Networking- A Top-Down…

### Example 2 — One user suddenly becomes active

Suppose 10 users share a 1 Mbps link, but only one user suddenly generates a large amount of data.

With TDM circuit switching, that user receives only its assigned time slot, while the other slots remain unused.

With packet switching, the active user can use the link at essentially the **full 1 Mbps** because the other users aren't generating packets.

Computer Networking- A Top-Down…

### Core idea

\[
\boxed{\text{Packet switching = capacity shared according to current demand}}
\]

---

# 1.3.2 Circuit Switching

Now consider the other approach.

## 1. What is circuit switching?

In a circuit-switched network, the network establishes a **dedicated end-to-end connection**, called a **circuit**, between the source and destination.

Before data transmission begins:

1. A connection is established.
2. Resources along the path are reserved.
3. Data is transmitted using those reserved resources.
4. When communication ends, the resources are released.

Computer Networking- A Top-Down…

Traditional **telephone networks** are the classic example given by the book.

---

## 2. Dedicated resources

Consider:

```text
Host A ── Switch ── Switch ── Host B
```

Suppose every link has four circuits:

```text
Link capacity
┌───┬───┬───┬───┐
│ C1│ C2│ C3│ C4│
└───┴───┴───┴───┘
```

If Host A communicates with Host B, the network reserves one circuit on **every link along the path**.

For example:

```text
Host A
  │
 [C2] ─── [C2] ─── [C2]
  │
Host B
```

That circuit remains dedicated to the connection for its entire duration.

Computer Networking- A Top-Down…

### Example

Suppose:

- Link capacity = 1 Mbps
- Each link has 4 circuits

Then each circuit receives:

\[
\frac{1\text{ Mbps}}{4}=250\text{ kbps}
\]

So that connection receives a **dedicated 250 kbps** transmission rate.

---

# 3. Multiplexing in Circuit Switching

The book explains two ways of dividing a link among circuits:

1. **Frequency-Division Multiplexing (FDM)**
2. **Time-Division Multiplexing (TDM)**

---

## A. Frequency-Division Multiplexing — FDM

In FDM, the frequency spectrum of the link is divided among connections.

Each connection receives a **frequency band** for the duration of the connection.

For example:

```text
Frequency
──────────────────────────────────►

| Conn A | Conn B | Conn C | Conn D |
```

Each connection continuously receives its assigned frequency range.

The book gives the telephone example of a **4-kHz frequency band** per connection. Computer Networking- A Top-Down…

### Key point

**FDM divides the link by frequency.**

---

# 4. Time-Division Multiplexing — TDM

In TDM, time is divided into fixed-duration **frames**.

Each frame is divided into **time slots**.

```text
Frame 1
┌────┬────┬────┬────┐
│ A  │ B  │ C  │ D  │
└────┴────┴────┴────┘

Frame 2
┌────┬────┬────┬────┐
│ A  │ B  │ C  │ D  │
└────┴────┴────┴────┘
```

If Connection A is assigned Slot A, then it gets that slot in **every frame**.

Thus:

> **TDM divides the link by time.**

Computer Networking- A Top-Down…

---

# 5. Packet Switching vs Circuit Switching

This is one of the **most important comparisons in Section 1.3**.

| Feature | Packet Switching | Circuit Switching |
|---|---|---|
| Connection | No dedicated connection required | Dedicated end-to-end connection |
| Resource reservation | No | Yes |
| Transmission | Packet-by-packet | Continuous reserved circuit |
| Link sharing | On demand | Pre-allocated |
| Queuing | Possible | Generally avoided after circuit establishment |
| Congestion | Can cause queuing and loss | Connection may be blocked if resources unavailable |
| Efficiency for bursty data | High | Lower |
| Resource utilization | Resources used when packets exist | Reserved resources may remain unused |
| Typical example | Internet | Traditional telephone network |

The book's central distinction is:

> **Circuit switching pre-allocates link use regardless of demand, whereas packet switching allocates link use on demand.** Computer Networking- A Top-Down…

---

# 6. An Important Conceptual Difference

Suppose User A has no data to send.

### Circuit switching

Its circuit is still reserved:

```text
User A ───── RESERVED ─────►
                ↑
             unused
```

The capacity cannot normally be used by another user.

### Packet switching

No packets means no capacity is being consumed by that user.

```text
User A ── no packets ──►

User B ── packet ───────►  Link
```

The link is used by whoever currently has packets.

That is why packet switching is particularly suitable for **bursty computer-data traffic**.

---

# 1.3.3 A Network of Networks

Now we move from:

> **How does data move through the core?**

to:

> **How are all the networks themselves interconnected to form the Internet?**

This is what the book means by **"a network of networks."**

---

# 1. What is a Network of Networks?

End systems connect to an **access ISP**.

For example:

```text
Laptop
   │
Wi-Fi
   │
Access ISP
   │
   ?
   │
Other networks
   │
Destination
```

But there are enormous numbers of access ISPs.

Therefore, simply connecting users to access ISPs is not enough.

The **access ISPs themselves must be interconnected**.

This creates:

\[
\boxed{\text{Internet = Network of Networks}}
\]

Computer Networking- A Top-Down…

The book emphasizes that today's Internet structure evolved significantly because of **economic and national-policy considerations**, not simply technical performance. Computer Networking- A Top-Down…

---

# 2. Why not connect every ISP directly to every other ISP?

Imagine there are hundreds of thousands of access ISPs.

A naive design would be:

```text
ISP A ───────── ISP B
  │  \          /  │
  │   \        /   │
  │    \      /    │
  │     \    /     │
ISP C ───────── ISP D
```

Every ISP would need a separate connection to every other ISP.

This is a **full mesh**.

The book says this would be far too expensive because each access ISP would need separate links to hundreds of thousands of other access ISPs. Computer Networking- A Top-Down…

Therefore, the Internet developed a hierarchy and interconnection structure.

---

# 3. Network Structure 1 — One Global Transit ISP

The first simplified structure is:

```text
          Global Transit ISP
       /    /    |    \    \
      /    /     |     \    \
   ISP   ISP    ISP    ISP   ISP
```

All access ISPs connect to one **global transit ISP**.

The access ISP pays the global transit ISP for connectivity.

Therefore:

- Access ISP = **customer**
- Global transit ISP = **provider**

Computer Networking- A Top-Down…

### Problem

A single global transit ISP would be extremely expensive and would have too much control.

---

# 4. Network Structure 2 — Multiple Global Transit ISPs

Other companies could compete by creating their own global transit ISPs.

So:

```text
             Transit ISP 1
             /     |     \
            /      |      \
         Access  Access   Access

             Transit ISP 2
             /     |     \
         Access  Access   Access
```

Now access ISPs can choose between competing providers based on:

- pricing
- services

But another problem appears:

> The global transit ISPs themselves must be interconnected.

Otherwise, an access ISP connected to Transit ISP 1 could not communicate with an access ISP connected to Transit ISP 2.

Computer Networking- A Top-Down…

---

# 5. Network Structure 3 — Regional ISPs + Tier-1 ISPs

The book then introduces a more realistic hierarchy.

```text
                 Tier-1 ISP
                /         \
               /           \
       Regional ISP     Regional ISP
        /    |    \       /    |    \
      ISP   ISP   ISP   ISP   ISP   ISP
```

### Access ISP

Provides Internet connectivity directly to end users.

Examples include:

- residential ISP
- university network
- company network

### Regional ISP

Aggregates/connects access ISPs within a region.

### Tier-1 ISP

A high-level ISP that connects to regional ISPs and other tier-1 ISPs.

The book notes that there is **no official organization that sanctions tier-1 status**. Computer Networking- A Top-Down…

---

# 6. Customer-Provider Relationship

At each level of the hierarchy, a lower-level ISP can be a **customer** of a higher-level ISP.

For example:

```text
Access ISP
    │
    │ pays
    ▼
Regional ISP
    │
    │ pays
    ▼
Tier-1 ISP
```

So:

- Access ISP → customer of regional ISP
- Regional ISP → customer of tier-1 ISP
- Tier-1 ISP → at the top and does not pay another ISP as a provider

Computer Networking- A Top-Down…

---

# 7. Points of Presence — PoPs

The book next introduces **Points of Presence (PoPs)**.

A PoP is:

> A group of one or more routers, at the same location, in a provider's network where customer ISPs can connect to the provider ISP.

Computer Networking- A Top-Down…

Conceptually:

```text
Customer ISP
     │
     │ high-speed link
     ▼
 ┌───────────────┐
 │      PoP       │
 │ Router Router  │
 └───────────────┘
     │
Provider ISP
```

---

# 8. Multi-Homing

An ISP does not necessarily need to depend on only one provider.

It can connect to **two or more provider ISPs**.

This is called **multi-homing**.

Example:

```text
             Regional ISP 1
                  │
                  │
Access ISP ───────┤
                  │
                  │
             Regional ISP 2
```

If one provider fails, the ISP can continue sending and receiving packets through the other provider.

Therefore:

\[
\boxed{\text{Multi-homing → greater resilience}}
\]

Computer Networking- A Top-Down…

---

# 9. Peering

Suppose two nearby ISPs are at the **same level** of the hierarchy.

Instead of sending their traffic through an upstream provider, they can directly connect to each other.

This is called **peering**.

```text
ISP A ───────────────── ISP B
          Direct link
```

Traffic between the two ISPs can then use this direct connection.

The book says peering is typically **settlement-free**, meaning neither ISP pays the other. Computer Networking- A Top-Down…

### Why peer?

Main reason:

**Reduce the cost of sending traffic through upstream providers.**

---

# 10. Internet Exchange Point — IXP

An **Internet Exchange Point (IXP)** is a meeting point where multiple ISPs can peer with one another.

Conceptually:

```text
 ISP A ──┐
 ISP B ──┼──► IXP
 ISP C ──┤
 ISP D ──┘
```

The IXP typically operates in a standalone facility containing switches.

So instead of every ISP individually establishing many direct connections, multiple ISPs can connect through the exchange point.

Computer Networking- A Top-Down…

---

# 11. Network Structure 4

The book calls the following combination **Network Structure 4**:

- Access ISPs
- Regional ISPs
- Tier-1 ISPs
- PoPs
- Multi-homing
- Peering
- IXPs

Computer Networking- A Top-Down…

This is a much better approximation of the actual Internet than the simple hierarchy.

---

# 12. Network Structure 5 — Content-Provider Networks

The book finally adds **content-provider networks**.

A major example discussed is **Google**.

The basic idea is:

```text
                Tier-1 ISP
                /        \
               /          \
        Regional ISP    Regional ISP
          /    \           /    \
       Access  Access    Access  Access
          \      \       /      /
           \      Google       /
              Private Network
```

Large content providers can build their own private networks connecting their data centers.

The book describes Google's private TCP/IP network as spanning the globe and being separate from the public Internet. It can peer directly with lower-tier ISPs or connect through IXPs, while also connecting to tier-1 ISPs where necessary. Computer Networking- A Top-Down…

### Why build such a network?

The book identifies two major motivations:

1. **Reduce payments to upper-tier ISPs**
2. **Have greater control over how services are delivered to end users**

Computer Networking- A Top-Down…

---

# 13. Final Structure of Today's Internet

The book summarizes today's Internet as a complex **network of networks**.

It consists of:

```text
                ┌─────────────────────┐
                │     Tier-1 ISPs     │
                └─────────┬───────────┘
                          │
                ┌─────────┴───────────┐
                │    Regional ISPs    │
                └─────────┬───────────┘
                          │
                ┌─────────┴───────────┐
                │     Access ISPs     │
                └─────────┬───────────┘
                          │
                     End Systems
```

But the real structure also includes:

- **Peering**
- **IXPs**
- **PoPs**
- **Multi-homing**
- **Content-provider networks**
- Interconnections between different tiers

The book's final summary describes the Internet as consisting of tier-1 ISPs and many lower-tier ISPs, with higher-tier ISPs interconnected and major content providers increasingly operating their own networks. Computer Networking- A Top-Down…

---

# ⭐ The Three Concepts Together

You should keep these three ideas clearly separated:

### 1. Packet Switching
**How data is transported through the network core.**

```text
Message
   ↓
Packets
   ↓
Router → Router → Router
   ↓
Destination
```

- No dedicated resources
- Store-and-forward
- Queuing possible
- Statistical multiplexing
- Efficient for bursty data

---

### 2. Circuit Switching
**Another method for transporting data through the core.**

```text
Source
  │
  ═══ Dedicated Circuit ═══
  │
Destination
```

- Connection established first
- Resources reserved
- Dedicated transmission rate
- FDM or TDM
- Traditional telephone networks

---

### 3. Network of Networks
**How the different networks/ISPs are interconnected to create the Internet.**

```text
End Systems
     ↓
Access ISP
     ↓
Regional ISP
     ↓
Tier-1 ISP
     ↕
Peering / IXP
     ↕
Other ISP networks
```

---

## 🔥 Exam-Critical Points

Remember these especially:

1. **Packet switching**
   - Message → packets
   - Store-and-forward
   - Output buffers
   - Queuing delay
   - Packet loss
   - Statistical multiplexing

2. **Circuit switching**
   - Dedicated end-to-end connection
   - Resources reserved for duration of connection
   - FDM = divide by **frequency**
   - TDM = divide by **time**

3. **Main difference**
   \[
   \boxed{\text{Circuit switching = reservation}}
   \]
   \[
   \boxed{\text{Packet switching = on-demand sharing}}
   \]

4. **Network of networks**
   - Access ISP → Regional ISP → Tier-1 ISP
   - Customer-provider relationships
   - PoP
   - Multi-homing
   - Peering
   - IXP
   - Content-provider networks

# 1.4 Delay, Loss, and Throughput in Packet-Switched Networks

According to **Kurose & Ross, Computer Networking: A Top-Down Approach, 7th Edition**, a packet-switched network has three fundamental performance characteristics:

- **Delay** — how long it takes data to travel from source to destination.
- **Packet loss** — packets can be dropped inside the network.
- **Throughput** — how much data can be transferred per second.

The book introduces these because an ideal network would transfer unlimited data **instantaneously and without loss**, which is impossible in reality. Computer Networking- A Top-Down…

---

# 1.4.1 Overview of Delay

Consider a packet traveling from a source to a destination:

```text
Source
  │
  ▼
Router 1
  │
  ▼
Router 2
  │
  ▼
Router 3
  │
  ▼
Destination
```

At every node, the packet can experience four types of delay:

1. **Processing delay**
2. **Queuing delay**
3. **Transmission delay**
4. **Propagation delay**

These together constitute the **nodal delay**. Computer Networking- A Top-Down…

A useful picture is:

```text
Packet arrives
      │
      ▼
┌───────────────┐
│ Processing    │
└───────┬───────┘
        │
        ▼
┌───────────────┐
│ Queue         │ ← Queuing delay
└───────┬───────┘
        │
        ▼
┌───────────────┐
│ Transmission  │ ← Transmission delay
└───────┬───────┘
        │
        ▼
═══════════════════  Link
   Propagation
═══════════════════
        │
        ▼
Next router
```

---

# 2. Processing Delay

## Definition

**Processing delay** is the time required by a router to:

- examine the packet's header,
- determine where to direct the packet,
- and potentially perform error checking.

Computer Networking- A Top-Down…

For example:

```text
Packet arrives
      ↓
Router examines header
      ↓
Determines outgoing link
      ↓
Packet placed in output queue
```

### What happens during processing?

The router needs to inspect information in the packet header and determine the appropriate outgoing link.

It may also check for **bit-level errors** that occurred while the packet was transmitted to the router. Computer Networking- A Top-Down…

### Magnitude

The book says that processing delays in high-speed routers are typically on the order of **microseconds or less**. Computer Networking- A Top-Down…

So, in many situations:

\[
d_{proc}\approx\text{very small}
\]

However, the book notes that processing delay can influence a router's **maximum throughput**, i.e., the maximum rate at which the router can forward packets. Computer Networking- A Top-Down…

---

# 3. Queuing Delay

After processing, the packet may have to wait in a queue before it can be transmitted.

```text
              Output link
                  │
                  ▼
Packet → [ Packet Packet Packet ] → Link
          ↑
        Queue
```

## Definition

**Queuing delay** is the time that a packet waits in the output queue before being transmitted onto the outgoing link.

Computer Networking- A Top-Down…

### If the queue is empty

The packet does not have to wait:

\[
\boxed{d_{queue}=0}
\]

### If the queue is heavily loaded

The packet may have to wait behind many packets:

\[
\boxed{d_{queue}\text{ becomes large}}
\]

The delay depends on:

- number of packets already waiting,
- traffic arrival rate,
- transmission rate of the outgoing link,
- whether traffic arrives periodically or in bursts.

Computer Networking- A Top-Down…

---

## Traffic intensity

The book defines **traffic intensity** as:

\[
\boxed{\frac{La}{R}}
\]

where:

- \(a\) = average packet arrival rate, in packets/sec
- \(L\) = packet size, in bits
- \(R\) = transmission rate, in bits/sec

Therefore:

\[
La
\]

is the average incoming bit rate.

So:

\[
\text{Traffic intensity}
=
\frac{\text{incoming bit rate}}{\text{outgoing transmission rate}}
\]

Computer Networking- A Top-Down…

### Important condition

If:

\[
\frac{La}{R}>1
\]

then packets arrive, on average, faster than they can be transmitted.

Therefore the queue tends to grow without bound in the idealized infinite-buffer model, and queuing delay approaches infinity.

The book's rule is essentially:

\[
\boxed{\text{Design the system so traffic intensity is no greater than 1}}
\]

Computer Networking- A Top-Down…

---

# 4. Transmission Delay

This is one of the most commonly confused concepts.

## Definition

**Transmission delay** is the amount of time required to **push all the bits of a packet onto the link**.

Suppose:

- Packet length = \(L\) bits
- Link transmission rate = \(R\) bits/sec

Then:

\[
\boxed{d_{trans}=\frac{L}{R}}
\]

Computer Networking- A Top-Down…

---

## Example

Suppose:

\[
L=10,000\text{ bits}
\]

and:

\[
R=10\text{ Mbps}
\]

Then:

\[
d_{trans}
=
\frac{10,000}{10,000,000}
\]

\[
\boxed{d_{trans}=1\text{ ms}}
\]

So it takes 1 ms to push the entire packet onto the link.

---

# 5. Propagation Delay

Once a bit has been pushed onto the link, it physically travels through the medium toward the next router.

The time required for this travel is the **propagation delay**.

Computer Networking- A Top-Down…

The formula is:

\[
\boxed{d_{prop}=\frac{d}{s}}
\]

where:

- \(d\) = physical distance between the two nodes
- \(s\) = propagation speed of the medium

The book gives propagation speeds roughly in the range:

\[
2\times10^8 \text{ m/s}
\quad\text{to}\quad
3\times10^8 \text{ m/s}
\]

depending on the physical medium. Computer Networking- A Top-Down…

---

## Example

Suppose:

\[
d=1000\text{ km}=1,000,000\text{ m}
\]

and:

\[
s=2\times10^8\text{ m/s}
\]

Then:

\[
d_{prop}
=
\frac{1,000,000}{2\times10^8}
\]

\[
=0.005\text{ sec}
\]

\[
\boxed{d_{prop}=5\text{ ms}}
\]

---

# 6. Transmission Delay vs Propagation Delay

This distinction is **extremely important for exams**.

| Transmission Delay | Propagation Delay |
|---|---|
| Time to push packet's bits onto link | Time for bits to travel across link |
| \(\frac{L}{R}\) | \(\frac{d}{s}\) |
| Depends on packet length \(L\) | Depends on physical distance \(d\) |
| Depends on transmission rate \(R\) | Depends on propagation speed \(s\) |
| Does **not** depend on distance | Does **not** depend on packet length |
| Concerned with putting bits onto link | Concerned with movement of bits through link |

The book explicitly emphasizes this difference. Computer Networking- A Top-Down…

### Easy way to remember

> **Transmission = pushing the packet onto the wire.**

> **Propagation = traveling through the wire.**

---

# 7. A Very Important Example

Suppose a packet has:

\[
L=1000\text{ bits}
\]

and the link rate is:

\[
R=1\text{ Mbps}
\]

Then:

\[
d_{trans}=\frac{1000}{10^6}=1\text{ ms}
\]

Now suppose the link is:

\[
d=200\text{ km}
\]

and:

\[
s=2\times10^8\text{ m/s}
\]

Then:

\[
d_{prop}
=
\frac{200,000}{2\times10^8}
=
1\text{ ms}
\]

Thus:

```text
Transmission delay = 1 ms
Propagation delay  = 1 ms
```

They happen for **different reasons**, even though their numerical values happen to be equal here.

---

# 8. The Book's Highway Analogy

Kurose & Ross use a **caravan and tollbooth** analogy.

Imagine:

- Each car = one bit
- Entire caravan = one packet
- Tollbooth = router
- Highway = communication link

The tollbooth must first push all cars of the caravan onto the highway.

That corresponds to:

### Transmission delay

Then the cars travel along the highway to the next tollbooth.

That corresponds to:

### Propagation delay

Computer Networking- A Top-Down…

This is a good conceptual distinction:

```text
Tollbooth                  Next tollbooth
   │                             │
   │ ← transmission →            │
   ├────────────────────────────►│
   │       propagation           │
```

---

# 9. Nodal Delay

The four delays combine to form the **nodal delay**.

The book gives:

\[
\boxed{
d_{nodal}
=
d_{proc}
+
d_{queue}
+
d_{trans}
+
d_{prop}
}
\]

Computer Networking- A Top-Down…

So for a single router/link:

```text
Nodal delay
     =
Processing
   +
Queuing
   +
Transmission
   +
Propagation
```

### Important observation

These components do not necessarily have equal importance.

For example:

- Processing delay may be very small.
- Transmission delay can be important on low-speed links.
- Propagation delay can be very large over long-distance/satellite links.
- Queuing delay can become dominant during congestion.

The book specifically notes that propagation delay can be hundreds of milliseconds for a geostationary satellite link, while transmission delay can also become hundreds of milliseconds over low-speed dial-up links. Computer Networking- A Top-Down…

---

# 10. Packet Loss

So far we assumed that packets can wait indefinitely in a queue.

Real routers don't have infinite buffers.

They have **finite queue capacity**.

Suppose:

```text
              Router
                 │
                 ▼
        ┌─────────────────┐
        │ Packet Packet   │
        │ Packet Packet   │
        │ Packet Packet   │
        └─────────────────┘
                 ↑
             FULL QUEUE
```

If another packet arrives:

```text
New packet
    ↓
FULL QUEUE
    ↓
   DROP
```

The router has nowhere to store it, so the packet is **dropped**.

That is **packet loss**. Computer Networking- A Top-Down…

---

## Why does packet loss occur?

Primarily because:

\[
\boxed{\text{Packets arrive faster than they can be transmitted}}
\]

which causes the queue to fill.

As traffic intensity increases, the fraction of lost packets generally increases. Computer Networking- A Top-Down…

From the end system's point of view, the packet appears to have entered the network but never arrived at the destination.

---

# 11. What happens to a lost packet?

The book notes that a lost packet **may be retransmitted on an end-to-end basis** so that the data can eventually be transferred from source to destination. Computer Networking- A Top-Down…

This distinction is important:

> **Packet loss occurs inside the network, but recovery from loss can be handled by end systems/protocols.**

The details of mechanisms such as TCP retransmission are covered later in the book.

---

# 12. End-to-End Delay

So far we've discussed delay at **one node**.

But normally we want to know:

> How long does it take for a packet to travel from the source all the way to the destination?

That is **end-to-end delay**.

Suppose:

```text
Source → Router 1 → Router 2 → ... → Router N−1 → Destination
```

There are \(N\) links.

For each node/link, the packet experiences the appropriate delays.

---

## Simplified equation from the book

The book considers an **uncongested network**, so queuing delay is negligible.

It assumes:

- \(N-1\) routers between source and destination
- processing delay = \(d_{proc}\)
- transmission rate = \(R\)
- propagation delay = \(d_{prop}\)
- packet length = \(L\)

Since:

\[
d_{trans}=\frac{L}{R}
\]

the book gives:

\[
\boxed{
d_{end-end}
=
N(d_{proc}+d_{trans}+d_{prop})
}
\]

or:

\[
\boxed{
d_{end-end}
=
N\left(
d_{proc}+\frac{L}{R}+d_{prop}
\right)
}
\]

Computer Networking- A Top-Down…

---

# 13. More General End-to-End Delay

In a real network, links and routers can have different characteristics.

So the general idea is:

\[
\boxed{
d_{end-end}
=
\sum_{\text{nodes/links}}
(d_{proc}+d_{queue}+d_{trans}+d_{prop})
}
\]

The important point is that the individual delays **accumulate along the entire path**.

The book's simplified Equation 1.2 omits queuing because it assumes an uncongested network. Computer Networking- A Top-Down…

---

# 14. Example of End-to-End Delay

Suppose a packet travels over **3 links**.

For each link/node, assume:

\[
d_{proc}=1\text{ ms}
\]

\[
d_{trans}=2\text{ ms}
\]

\[
d_{prop}=5\text{ ms}
\]

and queuing delay is negligible.

Then:

\[
d_{end-end}
=
3(1+2+5)
\]

\[
=3(8)
\]

\[
\boxed{24\text{ ms}}
\]

---

# 15. Throughput

Now we move from **delay** to **how much data can actually be delivered**.

## Definition

The book defines **end-to-end throughput** in the context of transferring a large file from one host to another.

### Instantaneous throughput

The instantaneous throughput is:

> The rate, in bits/sec, at which the destination is receiving the file at a particular instant.

### Average throughput

If:

- File size = \(F\) bits
- Transfer time = \(T\) seconds

then:

\[
\boxed{
\text{Average throughput}=\frac{F}{T}
}
\]

Computer Networking- A Top-Down…

---

# 16. Example of Average Throughput

Suppose:

\[
F=100\text{ Mbits}
\]

and the complete transfer takes:

\[
T=20\text{ sec}
\]

Then:

\[
\text{Throughput}
=
\frac{100\text{ Mbits}}{20\text{ sec}}
\]

\[
\boxed{5\text{ Mbps}}
\]

So the average throughput is **5 Mbps**.

---

# 17. Bottleneck Link

This is the central concept for understanding throughput.

Consider:

```text
Server
  │
  │ Rs = 10 Mbps
  ▼
Router
  │
  │ Rc = 2 Mbps
  ▼
Client
```

The server can send at 10 Mbps.

But the router can forward toward the client at only 2 Mbps.

Therefore, the end-to-end throughput cannot exceed 2 Mbps.

\[
\boxed{
Throughput=\min(R_s,R_c)
}
\]

The slowest link is called the **bottleneck link**. Computer Networking- A Top-Down…

---

# 18. Throughput Across Multiple Links

Suppose a path contains \(N\) links:

```text
Server
  │ R₁
  ▼
Router
  │ R₂
  ▼
Router
  │ R₃
  ▼
...
  │ Rₙ
  ▼
Client
```

If:

\[
R_1,R_2,\ldots,R_N
\]

are the transmission rates, then the book gives:

\[
\boxed{
Throughput=\min\{R_1,R_2,\ldots,R_N\}
}
\]

That is, the throughput is approximately the transmission rate of the **bottleneck link** when there is no competing traffic. Computer Networking- A Top-Down…

---

# 19. Example: Finding the Bottleneck

Suppose:

| Link | Rate |
|---|---:|
| \(R_1\) | 10 Mbps |
| \(R_2\) | 5 Mbps |
| \(R_3\) | 20 Mbps |
| \(R_4\) | 8 Mbps |

Then:

\[
Throughput
=
\min(10,5,20,8)
\]

\[
\boxed{Throughput=5\text{ Mbps}}
\]

Therefore:

\[
\boxed{R_2\text{ is the bottleneck link}}
\]

---

# 20. Throughput and Shared Traffic

The bottleneck isn't necessarily just the slowest physical link.

**Other traffic sharing the network can also reduce throughput.**

The book gives an example involving 10 simultaneous downloads sharing a common 5 Mbps link.

If the common link divides its capacity equally:

\[
\frac{5\text{ Mbps}}{10}
=
0.5\text{ Mbps}
\]

Therefore each download gets:

\[
\boxed{500\text{ kbps}}
\]

even though the individual access links may support higher rates. Computer Networking- A Top-Down…

Thus:

> Throughput depends not only on the transmission rates along the path, but also on **intervening traffic**.

Computer Networking- A Top-Down…

---

# 21. Throughput vs Delay

These two are often confused.

### Delay

Asks:

> **How long does it take for a packet/data to reach the destination?**

Measured in:

\[
\boxed{\text{seconds, milliseconds, microseconds}}
\]

### Throughput

Asks:

> **How much data can be delivered per second?**

Measured in:

\[
\boxed{\text{bits/sec}}
\]

Example:

```text
Network A:
Delay = 20 ms
Throughput = 5 Mbps

Network B:
Delay = 100 ms
Throughput = 100 Mbps
```

Network B has **higher throughput but greater delay**.

They are different performance characteristics.

The book notes that for applications such as Internet telephony, both low delay and sufficiently high instantaneous throughput can matter, while for file transfers, delay is generally less critical and high throughput is desirable. Computer Networking- A Top-Down…

---

# 22. Complete Picture of a Packet's Delay

Put everything together:

```text
                    ONE LINK

Packet arrives
     │
     ▼
┌──────────────┐
│ Processing   │
│    delay     │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│   Queue      │
│    delay     │
└──────┬───────┘
       │
       ▼
┌──────────────┐
│ Transmission │
│    delay     │
└──────┬───────┘
       │
       ▼
════════════════════
  Propagation delay
════════════════════
       │
       ▼
   Next router
```

Therefore:

\[
\boxed{
d_{nodal}
=
d_{proc}
+
d_{queue}
+
d_{trans}
+
d_{prop}
}
\]

And:

\[
\boxed{
d_{trans}=\frac{L}{R}
}
\]

\[
\boxed{
d_{prop}=\frac{d}{s}
}
\]

For the entire path:

\[
\boxed{
d_{end-end}
=
\sum \text{all nodal delays along the path}
}
\]

And for a large file transfer with no competing traffic:

\[
\boxed{
Throughput
\approx
\min(R_1,R_2,\ldots,R_N)
}
\]

---

# ⭐ Exam-Focused Comparison

| Concept | Meaning | Formula | Mainly depends on |
|---|---|---|---|
| **Processing delay** | Router examines/processes packet | — | Router processing |
| **Queuing delay** | Packet waits in output queue | — | Traffic/congestion |
| **Transmission delay** | Time to push all packet bits onto link | \(\frac{L}{R}\) | Packet size + link rate |
| **Propagation delay** | Time for bits to travel across link | \(\frac{d}{s}\) | Distance + propagation speed |
| **Nodal delay** | Total delay at one node/link | \(d_{proc}+d_{queue}+d_{trans}+d_{prop}\) | All four |
| **End-to-end delay** | Total source-to-destination delay | Sum of nodal delays | Entire path |
| **Packet loss** | Packet dropped because queue is full | — | Congestion/buffer capacity |
| **Throughput** | Data delivered per second | \(\frac{F}{T}\) | Bottleneck + competing traffic |

---

# 🔥 The Most Important Distinctions

### 1. Transmission vs Propagation

\[
\boxed{d_{trans}=\frac{L}{R}}
\]

Depends on **packet size and transmission rate**, **not distance**.

\[
\boxed{d_{prop}=\frac{d}{s}}
\]

Depends on **distance and propagation speed**, **not packet size or transmission rate**. Computer Networking- A Top-Down…

---

### 2. Queuing delay vs Transmission delay

**Transmission delay** is determined by:

\[
\frac{L}{R}
\]

It takes that long to push the packet onto the link.

**Queuing delay** depends on how many packets are waiting before it.

So:

> Transmission delay is fundamentally a property of the **packet and link**.

> Queuing delay is fundamentally a property of **traffic/congestion**.

---

### 3. Delay vs Throughput

\[
\boxed{\text{Delay}=\text{time}}
\]

\[
\boxed{\text{Throughput}=\text{data/time}}
\]

And throughput is generally limited by the **bottleneck link**.

---

### 4. Packet Loss

\[
\boxed{\text{Finite buffer + excessive incoming traffic}\rightarrow\text{packet loss}}
\]

As congestion increases, queuing delay and packet-loss probability can increase. Computer Networking- A Top-Down…

---

# 1. Protocol Layers and Service Models

In **Kurose & Ross, Chapter 1.5**, the authors introduce protocol layering because the Internet is an extremely complicated system containing many applications, protocols, end systems, packet switches, and different link technologies. Layering provides a way to organize this complexity. Computer Networking- A Top-Down…

There are two main ideas you need to understand:

1. **Layered Architecture**
2. **Encapsulation**

---

# 1.1 Layered Architecture

## Why do we need layers?

The Internet contains a huge number of components and protocols. If we tried to understand the entire Internet as one single system, it would become extremely difficult.

So, network designers **organize protocols and the hardware/software implementing them into layers**.

The basic idea is:

> **Each layer performs a specific set of functions and provides services to the layer above it.**

Each layer also uses the services provided by the layer immediately below it. Computer Networking- A Top-Down…

### Simple structure

```text
        Application Layer
               ↓
        Transport Layer
               ↓
         Network Layer
               ↓
          Link Layer
               ↓
         Physical Layer
```

This is the **five-layer Internet protocol stack** used throughout the book:

1. Application
2. Transport
3. Network
4. Link
5. Physical

The collection of protocols in these layers is called the **protocol stack**. Computer Networking- A Top-Down…

---

# 1.2 Service Model of a Layer

This is an important concept.

A layer provides a **service model** to the layer immediately above it.

The **service model** describes **what service a layer provides**, without necessarily specifying exactly how that service is implemented.

For example:

```text
Layer n
   │
   │ provides service
   ↓
Layer n+1
```

At the same time, Layer \(n\) itself uses the service of Layer \(n-1\):

```text
        Layer n+1
            ↑
      service from n
            ↑
        Layer n
            ↑
      service from n-1
            ↑
        Layer n-1
```

The book gives an important example:

Suppose Layer \(n\) wants to provide **reliable delivery**. Layer \(n-1\) might provide only unreliable delivery. Layer \(n\) can implement reliability by detecting lost messages and retransmitting them. Computer Networking- A Top-Down…

So:

> **A layer can provide a more sophisticated service by using the service of the layer below it and adding its own functionality.**

---

# 1.3 Why Layering Is Useful

The book identifies several important advantages.

## 1. Modularity

Each layer has a relatively well-defined responsibility.

This makes the network architecture **modular**.

For example, if the implementation of one layer changes, other layers do not necessarily need to change as long as the service provided to the neighboring layer remains the same.

The book's airline analogy illustrates exactly this idea: the implementation of a function can change while the function/service offered by that layer remains the same. Computer Networking- A Top-Down…

### Example

Suppose a layer provides:

> "Reliable delivery."

The internal mechanism used to achieve reliability can change, but the layer above still receives the same reliable-delivery service.

Therefore:

```text
Same service
     ↑
Different implementation
```

This is one of the major benefits of layering.

---

## 2. Easier system design and understanding

Layering gives us a **structured way to discuss network components**.

Instead of asking:

> "How does the entire Internet work?"

we can ask:

- What does the application layer do?
- What does the transport layer do?
- What does the network layer do?
- What does the link layer do?
- What does the physical layer do?

This breaks a complex problem into manageable pieces. Computer Networking- A Top-Down…

---

## 3. Easier updating

Because layers are modular, a component can often be modified without changing the entire network architecture.

The book specifically points out that **modularity makes it easier to update system components**. Computer Networking- A Top-Down…

---

# 1.4 Internet Protocol Stack

The book's Internet architecture has **five layers**:

```text
┌─────────────────────────────┐
│      Application Layer      │
├─────────────────────────────┤
│       Transport Layer       │
├─────────────────────────────┤
│        Network Layer        │
├─────────────────────────────┤
│          Link Layer         │
├─────────────────────────────┤
│        Physical Layer       │
└─────────────────────────────┘
```

Let's understand the function of each layer.

---

## 1. Application Layer

This is where **network applications and their application-layer protocols** reside.

Examples from the book:

- **HTTP** — Web document request and transfer
- **SMTP** — e-mail transfer
- **FTP** — file transfer
- **DNS** — name translation

An application-layer protocol is distributed across multiple end systems, where applications use the protocol to exchange information.

The book calls the packet of information at this layer a **message**. Computer Networking- A Top-Down…

### PDU/name

**Message**

---

## 2. Transport Layer

The transport layer transports application-layer messages between **application endpoints/processes**.

The Internet has two transport protocols:

- **TCP**
- **UDP**

### TCP

TCP provides:

- connection-oriented service
- reliable delivery
- flow control
- congestion control
- segmentation of long messages

### UDP

UDP provides:

- connectionless service
- no reliability
- no flow control
- no congestion control

The book refers to a transport-layer packet as a **segment**. Computer Networking- A Top-Down…

### PDU/name

**Segment**

---

## 3. Network Layer

The network layer is responsible for moving packets from **one host to another**.

Its principal protocol is **IP (Internet Protocol)**.

The transport layer gives the network layer:

- a transport-layer segment
- the destination address

The network layer then attempts to deliver that segment to the destination host.

The network layer also contains **routing protocols**, which determine routes through the network. Computer Networking- A Top-Down…

The book calls a network-layer packet a:

### PDU/name

**Datagram**

---

## 4. Link Layer

The network layer gets a datagram from source to destination through a sequence of links.

The link layer is responsible for moving a packet from:

> **one node → adjacent node**

The node may be a:

- host
- router
- switch
- WiFi access point

The network layer passes the datagram down to the link layer, which delivers it to the next node. At the next node, the datagram is passed back up to the network layer. Computer Networking- A Top-Down…

Examples:

- Ethernet
- WiFi
- DOCSIS

The book calls a link-layer packet a:

### PDU/name

**Frame**

---

## 5. Physical Layer

The physical layer moves the **individual bits** within a frame from one node to the next.

Its operation depends on:

- the link
- the transmission medium

Examples of media include:

- twisted-pair copper
- fiber optics
- coaxial cable

The physical layer determines how the individual bits are physically transmitted across the link. Computer Networking- A Top-Down…

### PDU/name

**Bits**

---

# 1.5 Important Layer Summary

| Layer | Main responsibility | Example protocols | Unit |
|---|---|---|---|
| **Application** | Network applications and application protocols | HTTP, SMTP, FTP, DNS | **Message** |
| **Transport** | Process-to-process delivery | TCP, UDP | **Segment** |
| **Network** | Host-to-host delivery/routing | IP, routing protocols | **Datagram** |
| **Link** | Node-to-adjacent-node delivery | Ethernet, WiFi, DOCSIS | **Frame** |
| **Physical** | Move individual bits across link | Link-specific physical protocols | **Bits** |

This table is **very important for exams**.

---

# 1.6 Protocol Layer Implementation

A layer is not necessarily implemented entirely in software.

The book explains:

- Application layer → usually software in end systems
- Transport layer → usually software in end systems
- Network layer → often a mixture of hardware and software
- Link and physical layers → typically associated with network interface hardware such as Ethernet/WiFi interface cards

Also, different network components implement different sets of layers. Computer Networking- A Top-Down…

For example:

### End system/host

```text
Application
Transport
Network
Link
Physical
```

A host implements **all five layers**. Computer Networking- A Top-Down…

### Router

```text
Network
Link
Physical
```

A router implements layers **1–3**.

### Link-layer switch

```text
Link
Physical
```

A link-layer switch implements layers **1–2**. Computer Networking- A Top-Down…

This explains why:

- A router understands IP/network-layer information.
- A link-layer switch does not implement IP; it works at the link layer.

---

# 1.7 Drawback of Layering

Layering is useful, but the book also discusses criticisms.

### Drawback 1: Duplicate functionality

A function may appear in multiple layers.

For example, error recovery can exist:

- on a per-link basis
- on an end-to-end basis

So functionality can be duplicated. Computer Networking- A Top-Down…

### Drawback 2: Information may cross layer boundaries

A layer may need information that exists only in another layer.

For example, one layer may need a timestamp maintained elsewhere.

This conflicts with the goal of **separation of layers**. Computer Networking- A Top-Down…

---

# 1.8 OSI Reference Model

The book also briefly discusses the **OSI model**.

OSI has **seven layers**:

```text
Application
Presentation
Session
Transport
Network
Data Link
Physical
```

The Internet protocol stack has five layers:

```text
Application
Transport
Network
Link
Physical
```

The important difference is that OSI has two additional layers:

- **Presentation**
- **Session**

The presentation layer provides services such as:

- data compression
- data encryption
- data representation/description

The session layer provides:

- delimiting
- synchronization
- checkpointing and recovery

The Internet does not have separate presentation and session layers; when an application needs such functionality, the book says it is the **application developer's responsibility** to implement it. Computer Networking- A Top-Down…

---

# 2. Encapsulation

Now comes the second major concept.

## Definition

**Encapsulation** is the process in which each layer takes the packet/data from the layer above and adds its own header information.

In the sending host, data moves:

```text
Application
     ↓
Transport
     ↓
Network
     ↓
Link
     ↓
Physical
```

At each step, additional control information is added.

The book explicitly describes the added information as **header information**. Computer Networking- A Top-Down…

---

# 2.1 Step-by-Step Encapsulation

Suppose an application wants to send a message:

```text
M
```

where \(M\) is the application-layer message.

---

## Step 1 — Application Layer

The application creates:

```text
[M]
```

This is the **message**.

No lower-layer header has been added yet.

---

## Step 2 — Transport Layer

The message is passed to the transport layer.

The transport layer adds a **transport-layer header**:

```text
[ Ht | M ]
```

where:

- \(H_t\) = transport-layer header
- \(M\) = application message

Together:

```text
[ Ht | M ] = Segment
```

The transport-layer header can contain information used by the receiving transport layer, such as information for delivering the message to the appropriate application and error-detection information. Computer Networking- A Top-Down…

Therefore:

> **Message + Transport Header = Segment**

---

## Step 3 — Network Layer

The transport-layer segment is passed to the network layer.

The network layer adds its own header:

```text
[ Hn | Ht | M ]
```

where:

- \(H_n\) = network-layer header
- \(H_t\) = transport-layer header
- \(M\) = application message

Together this becomes a:

```text
Datagram
```

The network-layer header contains information such as **source and destination end-system addresses**. Computer Networking- A Top-Down…

Therefore:

> **Network Header + Segment = Datagram**

---

## Step 4 — Link Layer

The datagram is passed to the link layer.

The link layer adds its own header:

```text
[ Hl | Hn | Ht | M ]
```

This becomes a:

```text
Frame
```

Therefore:

> **Link Header + Datagram = Frame**

The book notes that a packet at each layer contains **header fields and a payload field**, where the payload is typically the packet received from the layer above. Computer Networking- A Top-Down…

---

## Step 5 — Physical Layer

The physical layer takes the frame and transmits its individual bits across the physical medium.

Conceptually:

```text
Application:
        M

Transport:
        Ht | M
        └──────┘
        Segment

Network:
        Hn | Ht | M
        └──────────┘
        Datagram

Link:
        Hl | Hn | Ht | M
        └───────────────┘
        Frame

Physical:
        010101010101...
```

---

# 2.2 The Most Important Idea in Encapsulation

Notice what happens at every layer:

```text
Higher-layer packet
        ↓
becomes payload
        ↓
lower layer adds its own header
```

So:

```text
Application message
        ↓
Transport adds Ht
        ↓
Segment
        ↓
Network adds Hn
        ↓
Datagram
        ↓
Link adds Hl
        ↓
Frame
        ↓
Bits
```

This is **encapsulation**.

---

# 2.3 De-encapsulation at the Receiver

At the destination, the reverse process occurs.

The receiving host moves upward:

```text
Physical
   ↓
Link
   ↓
Network
   ↓
Transport
   ↓
Application
```

Suppose the received frame is:

```text
[ Hl | Hn | Ht | M ]
```

### Link layer

It processes/removes the link-layer information and passes the datagram upward:

```text
[ Hn | Ht | M ]
```

### Network layer

It processes its header and passes the segment upward:

```text
[ Ht | M ]
```

### Transport layer

It processes the transport header and delivers the application message:

```text
[ M ]
```

### Application layer

The application receives:

```text
M
```

This reverse operation is called **de-encapsulation**. The book illustrates it using the postal-envelope analogy. Computer Networking- A Top-Down…

---

# 2.4 Postal Analogy from the Book

The book uses an **interoffice memo and postal service** analogy.

Imagine Alice sends a memo to Bob.

### Original memo

```text
Memo
```

Equivalent to:

```text
Application message
```

Alice puts the memo into an interoffice envelope containing Bob's name and department.

```text
[Transport information | Memo]
```

Equivalent to:

```text
[Ht | M]
```

Then the branch mailroom puts that envelope inside another envelope suitable for the public postal service, with the source and destination branch addresses.

```text
[Network information | Transport envelope]
```

Equivalent to:

```text
[Hn | Ht | M]
```

At the destination, the outer envelope is removed, then the inner envelope, and finally Bob gets the original memo.

That is exactly analogous to:

```text
Encapsulation → transmission → de-encapsulation
```

The book explicitly uses this analogy to explain the nested nature of encapsulation. Computer Networking- A Top-Down…

---

# 2.5 Encapsulation Through Routers and Switches

An important point: **the packet does not simply travel through all five layers at every device.**

Consider:

```text
Host A
   ↓
Link-layer Switch
   ↓
Router
   ↓
Host B
```

### At Host A

All five layers are involved:

```text
Application
Transport
Network
Link
Physical
```

The data is encapsulated:

```text
Message
 ↓
Segment
 ↓
Datagram
 ↓
Frame
 ↓
Bits
```

### At Link-Layer Switch

The switch implements only:

```text
Link
Physical
```

So it handles the **frame**, not the transport or application message.

### At Router

The router implements:

```text
Network
Link
Physical
```

So it can process the **IP/network-layer datagram**, then send it onward through another link.

### At Host B

The data moves upward through all five layers until the application receives the original message.

This is why the book says different network components contain different sets of layers. Computer Networking- A Top-Down…

---

# 2.6 One Important Complication

Encapsulation is not always simply:

```text
1 message → 1 segment → 1 datagram → 1 frame
```

A large application message may be divided into **multiple transport-layer segments**.

Those segments may themselves be divided into multiple network-layer datagrams.

So conceptually:

```text
Large Message
      ↓
 ┌────┴────┐
 ↓         ↓
Segment  Segment
 ↓         ↓
Datagrams Datagrams
```

At the receiving side, the necessary reconstruction takes place.

The book explicitly notes this more complex case. Computer Networking- A Top-Down…

---

# 2.7 Service Model vs Encapsulation

These two concepts are easy to confuse.

### Service Model

Answers:

> **What does a layer provide to the layer above it?**

Example:

```text
TCP → reliable, connection-oriented service
UDP → connectionless, no-reliability service
```

The service is what the upper layer can use.

### Encapsulation

Answers:

> **How does data move down through the protocol stack?**

Example:

```text
Message
   ↓ + Ht
Segment
   ↓ + Hn
Datagram
   ↓ + Hl
Frame
   ↓
Bits
```

So remember:

> **Service model = what a layer offers.**  
> **Encapsulation = how each layer adds its information around the data from the layer above.**

---

# 2.8 Exam-Focused Summary

### Layered Architecture

- The Internet is organized into **five protocol layers**.
- Each layer has specific responsibilities.
- Each layer provides a **service model** to the layer above.
- A layer uses the services of the layer immediately below.
- Layering provides **modularity**, making systems easier to understand and update.
- The collection of protocols in the layers is called the **protocol stack**. Computer Networking- A Top-Down…

### Five Internet Layers

```text
Application → Message
Transport   → Segment
Network     → Datagram
Link        → Frame
Physical    → Bits
```

### Encapsulation

```text
Message
   ↓
Ht + Message
   ↓
Hn + Ht + Message
   ↓
Hl + Hn + Ht + Message
   ↓
Bits
```

At the destination:

```text
Bits
 ↓
Frame
 ↓
Datagram
 ↓
Segment
 ↓
Message
```

### One-line definition

> **Encapsulation is the process by which each protocol layer adds its own header information to the packet received from the layer above, forming the packet used by that layer.** Computer Networking- A Top-Down…

### ⭐ Most important distinction

| Concept | Question it answers |
|---|---|
| **Layer** | What function is being performed? |
| **Service model** | What does this layer provide to the layer above? |
| **Protocol** | What rules govern communication within the layer? |
| **Encapsulation** | How is data wrapped with layer-specific headers? |
| **De-encapsulation** | How is that wrapping removed at the destination? |
| **Protocol stack** | What are all the layers/protocols taken together? |
# 1. Networks Under Attack

In **Section 1.6 of Kurose & Ross**, the authors introduce network security by asking a basic question:

> **What can go wrong when computers and networks are connected to the Internet?**

The section is a survey of major attacks. Detailed defenses are developed later, especially in **Chapter 8, Security in Computer Networks**. Computer Networking- A Top-Down…

The main attacks discussed are:

1. **Malware**
2. **Denial-of-Service (DoS) attacks**
3. **Packet sniffing**
4. **IP spoofing / masquerading**
5. **Modification, insertion, and deletion of messages**
6. **Man-in-the-middle attacks** are mentioned as a broader security challenge

---

# 1.1 Malware Attacks

## What is malware?

**Malware** is malicious software that can enter and infect an Internet-connected host.

Once installed, malware can perform various malicious actions, such as:

- deleting files
- installing spyware
- collecting private information
- collecting passwords and keystrokes
- sending collected information back to the attacker
- turning the infected machine into part of a **botnet** Computer Networking- A Top-Down…

### Botnet

A **botnet** is a network of many compromised devices controlled by an attacker.

The attacker can use these compromised hosts for activities such as:

- spam distribution
- Distributed Denial-of-Service (**DDoS**) attacks

So the attack can develop like this:

```text
Attacker
   ↓
Compromises Host A
   ↓
Host A + many other infected hosts
   ↓
Botnet
   ↓
Used for attacks
```

---

# 1.2 Self-Replicating Malware

A particularly dangerous property of some malware is **self-replication**.

Once one host is infected:

```text
Host A infected
      ↓
infects Host B, C, D
      ↓
B, C, D infect more hosts
      ↓
rapid expansion
```

The book notes that this can allow malware to spread **exponentially fast**. Computer Networking- A Top-Down…

The book discusses two important forms:

### Virus

A **virus** is malware that requires some form of **user interaction** to infect the device.

Example from the book:

```text
Malicious e-mail attachment
          ↓
User opens attachment
          ↓
Malicious code executes
          ↓
Device infected
```

An e-mail virus may then send copies of itself to people in the user's address book. Computer Networking- A Top-Down…

### Worm

A **worm** can enter a device **without explicit user interaction**.

For example:

```text
Vulnerable network application
          ↓
Attacker sends malware
          ↓
Application accepts/runs it
          ↓
Host becomes infected
          ↓
Worm searches Internet
          ↓
Finds other vulnerable hosts
          ↓
Infects them
```

The key distinction in the book is:

| Virus | Worm |
|---|---|
| Requires some user interaction | Does not require explicit user interaction |
| Example: malicious e-mail attachment opened by user | Exploits a vulnerable network application |
| Can self-replicate after execution | Can automatically search for and infect other vulnerable hosts |

Computer Networking- A Top-Down…

### ⭐ Exam point

> **Virus → user interaction required.**  
> **Worm → no explicit user interaction required.**

---

# 1.3 Denial-of-Service (DoS) Attacks

A **Denial-of-Service attack** attempts to make a network, host, or infrastructure component **unusable by legitimate users**.

Targets can include:

- Web servers
- E-mail servers
- DNS servers
- institutional networks Computer Networking- A Top-Down…

The book divides Internet DoS attacks into **three categories**.

---

## A. Vulnerability Attack

The attacker sends a **small number of carefully constructed messages** to a vulnerable application or operating system.

If the packets trigger the vulnerability:

```text
Attacker
   ↓
Specially crafted packets
   ↓
Vulnerable application/OS
   ↓
Service stops
      OR
Host crashes
```

The important point is that this attack does **not necessarily require a huge amount of traffic**.

A few carefully crafted packets may be sufficient. Computer Networking- A Top-Down…

---

## B. Bandwidth Flooding

Here, the attacker sends a **large amount of traffic** toward the target.

```text
Attacker
   ↓↓↓↓↓↓↓↓↓
   Huge traffic
       ↓
Target access link
       ↓
Becomes congested
       ↓
Legitimate packets cannot reach server
```

The target's access link becomes clogged, preventing legitimate traffic from getting through. Computer Networking- A Top-Down…

### Connection to Chapter 1.4

This attack directly relates to the earlier concept of **queuing and congestion**.

If packets arrive faster than the target/access link can transmit them, queues fill and legitimate packets can be lost.

The book makes this connection explicitly when discussing bandwidth-flooding attacks. Computer Networking- A Top-Down…

---

## C. Connection Flooding

The attacker establishes a large number of:

- half-open TCP connections, or
- fully open TCP connections

at the target.

```text
Many bogus TCP connections
          ↓
Target must maintain them
          ↓
Resources become exhausted
          ↓
Legitimate connections rejected
```

The target becomes so overloaded with bogus connections that it can stop accepting legitimate ones. Computer Networking- A Top-Down…

---

# 1.4 DoS vs DDoS

A normal DoS attack may originate from a **single source**.

But a single machine may not have enough capacity to overwhelm a target with a very high-rate access link.

Therefore, an attacker can use many compromised machines.

This is a **Distributed Denial-of-Service (DDoS)** attack.

```text
             ┌── Slave 1 ──┐
             │              │
Attacker ────┼── Slave 2 ───┼──→ Victim
             │              │
             ├── Slave 3 ───┤
             │              │
             └── Slave 4 ───┘
```

The attacker controls multiple sources, and each source sends traffic toward the target.

The aggregate traffic from all sources can be approximately equal to the target's access rate \(R\), potentially crippling the service. Computer Networking- A Top-Down…

### Why DDoS is harder to defend against

With a single-source attack:

```text
One attacker
     ↓
Router can potentially identify/block source
```

With DDoS:

```text
Many compromised sources
     ↓
Traffic comes from many locations
     ↓
Much harder to distinguish from legitimate traffic
```

The book therefore notes that DDoS attacks are much harder to detect and defend against than single-source DoS attacks. Computer Networking- A Top-Down…

---

# 1.5 Packet Sniffing

A **packet sniffer** is a passive receiver that records copies of packets traveling through a communication medium.

This is particularly important in wireless networks.

```text
Sender ─────────────→ Receiver
            ↘
             ↘
          Sniffer
```

The sniffer does not need to interfere with the communication. It simply obtains copies of transmitted packets.

Packets may contain sensitive information such as:

- passwords
- personal information
- trade secrets
- private messages Computer Networking- A Top-Down…

---

## Why is sniffing dangerous?

Suppose a user transmits:

```text
Username: Alice
Password: XXXXX
```

If the communication is not adequately protected, a nearby attacker could potentially record the transmitted packets and analyze them.

The important point is:

> **Sniffing is a passive attack.**

The attacker is observing/copying traffic rather than injecting traffic.

Because the attacker is passive, sniffing can be difficult to detect. Computer Networking- A Top-Down…

---

## Sniffing is not limited to wireless

The book also explains that sniffers can operate in wired environments.

Examples include:

- broadcast Ethernet environments
- cable access networks
- an institution's access router/access link if an attacker gains access there

Computer Networking- A Top-Down…

### Defense mentioned by the book

The book points toward **cryptography** as an important defense against packet sniffing. Detailed cryptographic mechanisms are covered later in Chapter 8. Computer Networking- A Top-Down…

---

# 1.6 IP Spoofing / Masquerading

Another attack is to **pretend to be another host**.

The book calls this **IP spoofing**.

An attacker can construct a packet with an arbitrary:

- source address
- destination address
- packet contents

and send it into the Internet.

The network may forward it based on the supplied information.

```text
Attacker
   |
   | Packet:
   | Source = Fake/Someone else's IP
   ↓
Internet
   ↓
Victim
```

The victim sees the false source address and may incorrectly believe the packet came from the claimed source. Computer Networking- A Top-Down…

### Definition

> **IP spoofing is the injection of packets into the Internet with a false source address.**

Computer Networking- A Top-Down…

---

## Why is spoofing dangerous?

If a receiver trusts the source address without authentication, the attacker can **masquerade as another user or host**.

Therefore, the book says we need **end-point authentication** to determine whether a message actually originated from the claimed source. Computer Networking- A Top-Down…

### ⭐ Exam point

```text
Packet Sniffing
      ↓
Attacker observes/copies traffic

IP Spoofing
      ↓
Attacker creates/injects traffic
using a false source identity
```

These are different attacks.

---

# 1.7 Message Modification, Insertion, and Deletion

An attacker may not merely observe a communication.

An attacker positioned between communicating parties can potentially:

- **modify** messages
- **insert** new messages
- **delete** messages

The book later describes these capabilities explicitly in its security discussion. Computer Networking- A Top-Down…

For example:

```text
Alice ───────→ Bob
       ↑
    Attacker
```

The attacker could potentially alter:

```text
Original:
"Transfer ₹100"

Modified:
"Transfer ₹10,000"
```

Or:

```text
Alice → Message 1 → Attacker → [deleted] → Bob
```

Or insert a message that Alice never sent.

This is why secure communication needs **message integrity**: Alice and Bob should be able to determine whether the content was altered during transmission. Computer Networking- A Top-Down…

---

# 1.8 Man-in-the-Middle Attack

The book mentions **man-in-the-middle attacks** as one of the security challenges that must be defended against. Computer Networking- A Top-Down…

Conceptually:

```text
Alice
  ↓
Attacker
  ↓
Bob
```

The attacker is positioned between the communicating parties and can potentially observe or interfere with their communication.

This can combine several capabilities:

- eavesdropping/sniffing
- modification
- insertion
- deletion
- impersonation

The detailed mechanisms for defending against such attacks are covered later in the book rather than in Section 1.6.

---

# 1.9 Why Was the Internet Originally Vulnerable?

This is an important conceptual point from the book.

The original Internet was designed around a model of:

> **a group of mutually trusting users connected to a transparent network**

Under that assumption, security mechanisms were not built into the architecture by default. Computer Networking- A Top-Down…

For example:

- One user could send packets to another by default.
- User identity was generally accepted based on the declared identity.
- Authentication was not automatically required.

But today's Internet does **not** consist of mutually trusting users.

Therefore, modern networking needs mechanisms to deal with:

- sniffing
- endpoint masquerading
- malware
- DDoS
- man-in-the-middle attacks
- other attacks Computer Networking- A Top-Down…

---

# 1.10 How These Attacks Relate to Security Goals

The attacks introduced in Section 1.6 motivate the security requirements developed later in the book.

| Attack | What attacker does | Security property threatened |
|---|---|---|
| **Malware** | Infects a host and performs malicious actions | Host/data security |
| **DoS** | Makes service unavailable | **Availability** |
| **Packet sniffing** | Records packets | **Confidentiality** |
| **IP spoofing** | Pretends to be another source | **Authentication** |
| **Message modification** | Changes transmitted data | **Integrity** |
| **Message insertion** | Adds unauthorized messages | **Integrity/authentication** |
| **Message deletion** | Removes messages | **Integrity/availability** |
| **Man-in-the-middle** | Intercepts and potentially alters communication | Confidentiality, integrity, authentication |

The book later formalizes three major security properties:

### 1. Confidentiality

Only the intended parties should understand the communication.

### 2. Message integrity

The communicating parties should be able to determine whether the message was altered.

### 3. End-point authentication

The communicating parties should be able to verify who the other party actually is. Computer Networking- A Top-Down…

---

# 1.11 Complete Attack Flow to Remember

For exams, think of Section 1.6 as covering **different ways an attacker can interfere with a network**:

```text
                    NETWORK ATTACKS
                          │
       ┌──────────────────┼──────────────────┐
       │                  │                  │
    Attack host       Attack service      Attack communication
       │                  │                  │
    Malware              DoS             Sniffing
       │                  │                  │
   Virus/Worm        ┌────┼────┐          Passive
                     │    │    │         observation
                  Vuln. Bandwidth
                  attack flooding
                        │
                   Connection
                    flooding
```

And:

```text
Communication attacks
        │
        ├── Sniffing
        │      → observe/copy packets
        │
        ├── IP spoofing
        │      → fake source identity
        │
        ├── Modification
        │      → alter messages
        │
        ├── Insertion
        │      → add messages
        │
        └── Deletion
               → remove messages
```

---

# ⭐ Exam-Focused Summary

### **1. Malware**
Malicious software infects hosts. It can steal information, delete files, install spyware, or turn hosts into a botnet.

- **Virus:** requires user interaction.
- **Worm:** does not require explicit user interaction.
- Self-replicating malware can spread rapidly. Computer Networking- A Top-Down…

### **2. DoS**
Makes a network/host/service unavailable to legitimate users.

Three types:

1. **Vulnerability attack** → exploit vulnerable software with carefully crafted packets.
2. **Bandwidth flooding** → overwhelm access link with traffic.
3. **Connection flooding** → create huge numbers of TCP connections. Computer Networking- A Top-Down…

### **3. DDoS**
A distributed DoS attack uses **multiple compromised sources** to attack one target. Computer Networking- A Top-Down…

### **4. Packet Sniffing**
A passive attacker records copies of packets traveling through a network.

> **Sniffing = listening/observing.**

### **5. IP Spoofing**
Attacker sends packets using a **false source IP address**.

> **Spoofing = pretending to be someone else.**

### **6. Message Modification/Insertion/Deletion**
An attacker interferes with messages while they are in transit.

> **Integrity mechanisms** are needed to detect such tampering. Computer Networking- A Top-Down…

### **7. Core security goals**

```text
Confidentiality      → prevent unauthorized reading
Integrity            → detect unauthorized modification
Authentication       → verify communicating identity
Availability         → keep services usable
```

# 1. History of Computer Networking and the Internet

This is **Section 1.7 of Kurose & Ross, 7th Edition**. The book divides the history into **five periods**: Computer Networking- A Top-Down…

1. **Development of Packet Switching — 1961–1972**
2. **Proprietary Networks and Internetworking — 1972–1980**
3. **Proliferation of Networks — 1980–1990**
4. **Internet Explosion — 1990s**
5. **New Millennium**

The key story is:

> **Circuit switching → packet switching → ARPANET → multiple independent networks → internetworking/TCP-IP → widespread Internet → Web → broadband, wireless, social networks, private networks, and cloud.**

---

# 1.7.1 Development of Packet Switching: 1961–1972

## Why was packet switching needed?

In the early 1960s, the **telephone network** was the dominant communication network.

Telephone networks used **circuit switching**, which was appropriate for voice because voice traffic is relatively continuous and has a fairly constant transmission rate.

But computers generated a very different kind of traffic.

Computer users often generated **bursty traffic**:

```text
Send command
     ↓
Activity
     ↓
Wait for response
     ↓
Activity again
```

There could be long periods of inactivity between bursts.

Therefore, reserving a dedicated circuit continuously for such traffic would be inefficient. Computer Networking- A Top-Down…

This motivated researchers to investigate **packet switching**.

---

## Development of packet switching

The book identifies **three research groups**, working independently:

### 1. Leonard Kleinrock — MIT

Kleinrock's early work used **queuing theory** to demonstrate the effectiveness of packet switching for bursty traffic.

His work was among the first published work on packet-switching techniques. Computer Networking- A Top-Down…

### 2. Paul Baran — RAND

In 1964, Paul Baran at RAND investigated packet switching for **secure voice communication over military networks**. Computer Networking- A Top-Down…

### 3. Donald Davies and Roger Scantlebury — NPL, England

They were independently developing packet-switching ideas at the **National Physical Laboratory (NPL)**. Computer Networking- A Top-Down…

### Important point

```text
MIT             → Kleinrock
RAND            → Baran
NPL, England    → Davies + Scantlebury
```

Their work laid important foundations for the Internet. Computer Networking- A Top-Down…

---

# ARPA and ARPANET

In the 1960s, **J. C. R. Licklider** and **Lawrence Roberts**, both associated with MIT, went on to lead the computer science program at **ARPA** in the United States.

Lawrence Roberts published an overall plan for **ARPANET** in 1967.

The book describes ARPANET as:

> the first packet-switched computer network and a direct ancestor of today's public Internet. Computer Networking- A Top-Down…

---

## ARPANET in 1969

On **Labor Day in 1969**, the first packet switch was installed at **UCLA**.

Three additional packet switches were then installed at:

- Stanford Research Institute (**SRI**)
- UC Santa Barbara
- University of Utah

Thus, by the end of 1969:

```text
        UCLA
         |
         |
        SRI
       /   \
      /     \
   UCSB    Utah
```

The fledgling network had **four nodes**. Computer Networking- A Top-Down…

One of the first uses was a **remote login from UCLA to SRI**.

Interestingly, the first attempt crashed the remote system. 😄 Computer Networking- A Top-Down…

---

## ARPANET by 1972

By 1972:

- ARPANET had approximately **15 nodes**
- Robert Kahn gave its first public demonstration
- The first host-to-host protocol, **NCP (Network Control Protocol)**, had been completed
- Applications could now be developed on top of the network
- **Ray Tomlinson wrote the first e-mail program in 1972** Computer Networking- A Top-Down…

So:

```text
1969 → 4 nodes
1972 → ~15 nodes + NCP + first e-mail
```

---

# 1.7.2 Proprietary Networks and Internetworking: 1972–1980

The original ARPANET was a **single, closed network**.

To communicate with an ARPANET host, a machine had to be attached to an ARPANET IMP.

But during the 1970s, many other packet-switched networks appeared.

Examples mentioned by the book include:

- **ALOHANet**
- DARPA packet-satellite network
- packet-radio networks
- Telenet
- Cyclades
- Tymnet
- GE Information Services
- IBM's SNA Computer Networking- A Top-Down…

The problem was now:

> **How can different networks communicate with one another?**

---

# The Birth of Internetworking

Each network could have its **own protocol**.

For example:

```text
Network A
   |
 Protocol A

Network B
   |
 Protocol B

Network C
   |
 Protocol C
```

A machine on Network A could not automatically communicate with a machine on Network B.

The solution was to develop an architecture for **connecting different networks together**.

Vinton **Cerf** and Robert **Kahn**, under DARPA sponsorship, pioneered this work.

The term **internetting** was coined for this work. Computer Networking- A Top-Down…

The fundamental idea was to create a:

> **network of networks**

This is the conceptual foundation of today's Internet.

---

# TCP/IP

Cerf and Kahn developed the principles that eventually became **TCP/IP**.

The early TCP was different from today's TCP.

Initially, TCP combined:

- reliable, in-order delivery
- end-system retransmission
- forwarding functions

Later, forwarding functionality was separated from TCP and became the responsibility of **IP**.

The need for an unreliable, non-flow-controlled, end-to-end transport service also led to the development of **UDP**. Computer Networking- A Top-Down…

Thus, by the end of the 1970s, the three major Internet protocols were conceptually in place:

```text
TCP
UDP
IP
```

### ⭐ Important exam point

**TCP + IP were not always one permanently inseparable protocol.**

The early TCP combined functions that later became separated into:

```text
TCP → transport functions
IP  → network/forwarding functions
```

UDP emerged as an alternative transport service.

---

# ALOHANet and Ethernet

Another important development occurred in Hawaii.

**Norman Abramson** developed **ALOHANet**, a packet-based radio network connecting remote sites in the Hawaiian Islands.

The **ALOHA protocol** was the first multiple-access protocol, allowing geographically distributed users to share a single broadcast communication medium. Computer Networking- A Top-Down…

Later, **Metcalfe and Boggs** built on this work to develop **Ethernet** for wire-based shared broadcast networks.

Ethernet was motivated by the need to connect:

- PCs
- printers
- shared disks

This work laid foundations for today's LANs. Computer Networking- A Top-Down…

---

# 1.7.3 Proliferation of Networks: 1980–1990

The 1980s were a period of **very rapid growth**.

At the end of the 1970s:

```text
~200 hosts on ARPANET
```

By the end of the 1980s:

```text
~100,000 hosts on the public Internet
```

The Internet was becoming a **confederation of networks**, increasingly resembling the architecture we recognize today. Computer Networking- A Top-Down…

---

## University Networks

Several important networks contributed to this growth.

### BITNET

Provided:

- e-mail
- file transfer

among universities in the northeastern United States.

### CSNET

Created to connect computer-science researchers who did not have access to ARPANET.

### NSFNET

Created in **1986** to provide access to NSF-sponsored supercomputing centers.

Its backbone started at:

\[
56\text{ kbps}
\]

and reached:

\[
1.5\text{ Mbps}
\]

by the end of the decade.

NSFNET became an important backbone connecting regional networks. Computer Networking- A Top-Down…

---

# TCP/IP Becomes the ARPANET Standard

One of the most important dates in networking history is:

## **January 1, 1983**

On this date, **TCP/IP officially became the standard host protocol for ARPANET**, replacing NCP.

The transition was a **flag day**:

> all hosts were required to switch to TCP/IP on that date. Computer Networking- A Top-Down…

So remember:

```text
Before 1983 → NCP
January 1, 1983 → TCP/IP
```

---

## Other important developments during the 1980s

### TCP congestion control

Important extensions were made to TCP in the late 1980s to implement **host-based congestion control**. Computer Networking- A Top-Down…

### DNS

The **Domain Name System (DNS)** was developed.

It maps human-readable names such as:

```text
gaia.cs.umass.edu
```

to IP addresses. Computer Networking- A Top-Down…

---

# Minitel

The book also discusses a significant development outside the United States.

In the early 1980s, France launched **Minitel**.

It consisted of:

- a public packet-switched network
- X.25 protocol suite
- Minitel servers
- inexpensive terminals with low-speed modems

In 1984, the French government gave a free Minitel terminal to households that wanted one.

At its peak in the mid-1990s, Minitel offered **more than 20,000 services**, including home banking and research databases. Computer Networking- A Top-Down…

---

# 1.7.4 Internet Explosion: The 1990s

The 1990s marked the transition toward the **commercial and mass-market Internet**.

Two major events:

### 1. ARPANET disappeared

ARPANET, the progenitor of the Internet, ceased to exist.

### 2. NSFNET became commercial

In **1991**, NSFNET restrictions on commercial use were lifted.

NSFNET was eventually decommissioned in **1995**, and commercial ISPs carried Internet backbone traffic. Computer Networking- A Top-Down…

But the biggest development of the 1990s was:

# The World Wide Web

---

# World Wide Web

The Web was invented at **CERN by Tim Berners-Lee between 1989 and 1991**.

The initial Web included four key components:

1. **HTML**
2. **HTTP**
3. **Web server**
4. **Browser**

Computer Networking- A Top-Down…

The Web dramatically increased Internet adoption because it provided an easy-to-use platform for accessing information and applications.

---

## Web growth

Around the end of 1993:

```text
~200 Web servers
```

By the mid-1990s:

- GUI-based browsers became popular
- Netscape became widely used
- companies began operating Web servers
- companies began conducting commerce over the Web

Microsoft entered the browser market in 1996, leading to the **browser war between Netscape and Microsoft**. Computer Networking- A Top-Down…

---

# Internet Applications Explode

The second half of the 1990s saw enormous growth.

By the end of the millennium, the book identifies four **killer applications**:

1. **E-mail**
   - attachments
   - Web-accessible e-mail

2. **The Web**
   - Web browsing
   - Internet commerce

3. **Instant messaging**
   - contact lists

4. **Peer-to-peer MP3 file sharing**
   - pioneered by Napster

Computer Networking- A Top-Down…

---

# Dot-com Bubble

Between **1995 and 2001**, Internet companies experienced enormous financial speculation.

Many startups:

- went public
- received extremely high valuations
- had little or no significant revenue

Internet stocks then collapsed in **2000–2001**, causing many startups to shut down.

Nevertheless, several companies emerged as major winners, including:

- Microsoft
- Cisco
- Yahoo
- eBay
- Google
- Amazon Computer Networking- A Top-Down…

---

# 1.7.5 The New Millennium

The book describes continued rapid innovation after 2000.

Major developments included:

1. **Broadband Internet**
2. **High-speed WiFi and 4G**
3. **Mobile/handheld Internet**
4. **Social networks**
5. **Private networks of large online service providers**
6. **Cloud computing**

---

## 1. Broadband Internet

Broadband access expanded rapidly through:

- cable modems
- DSL
- fiber-to-the-home (**FTTH**)

Higher access speeds enabled applications such as:

- YouTube
- Netflix
- video conferencing

Computer Networking- A Top-Down…

---

## 2. Wireless Internet and Mobile Devices

High-speed public WiFi and 4G cellular networks enabled users to remain connected while moving.

The book notes that by **2011**, the number of wireless devices connecting to the Internet had surpassed the number of wired devices. Computer Networking- A Top-Down…

This enabled rapid growth of:

- smartphones
- tablets
- location-based applications

Examples in the book include mapping, check-in, dating, and traffic applications.

---

## 3. Social Networks

Social networks created another layer of connectivity **on top of the Internet**.

Examples given by the book include:

- Facebook
- Instagram
- Twitter
- WeChat

They supported:

- messaging
- photo sharing
- distributed games
- new applications through APIs Computer Networking- A Top-Down…

So the Internet was no longer just:

```text
Computer → Network → Computer
```

It increasingly supported:

```text
People
  ↓
Social networks
  ↓
Applications
  ↓
Internet infrastructure
```

---

# 4. Private Networks of Large Online Services

Large online service providers such as Google and Microsoft built extensive **private networks**.

These networks:

- connect their globally distributed data centers
- allow them to bypass portions of the public Internet
- peer directly with lower-tier ISPs

This improves the delivery of services such as search and e-mail. Computer Networking- A Top-Down…

This connects directly with the book's earlier discussion of the Internet as a **network of networks**.

---

# 5. Cloud Computing

Many Internet companies began running applications in the **cloud**.

Examples in the book include:

- Amazon EC2
- Google Application Engine
- Microsoft Azure

Companies and universities also migrated applications such as:

- e-mail
- Web hosting

to cloud environments.

Cloud providers offer:

- scalable computing
- scalable storage
- access to high-performance private networks Computer Networking- A Top-Down…

---

# Complete Timeline

This is the **most important timeline to memorize** for exams.

```text
1961–1964
    ↓
Packet-switching research
Kleinrock, Baran, Davies/Scantlebury
    ↓
1967
    ↓
Roberts publishes ARPANET plan
    ↓
1969
    ↓
First ARPANET packet switches
UCLA, SRI, UCSB, Utah
    ↓
1972
    ↓
~15 ARPANET nodes
NCP
First e-mail
    ↓
1970s
    ↓
Multiple independent networks appear
ALOHANet, packet radio, packet satellite,
Telenet, Cyclades, etc.
    ↓
1970s
    ↓
Cerf + Kahn
Internetworking
TCP/IP concepts
    ↓
1980s
    ↓
BITNET, CSNET, NSFNET
Rapid network growth
    ↓
Jan 1, 1983
    ↓
TCP/IP replaces NCP on ARPANET
    ↓
Late 1980s
    ↓
TCP congestion control
DNS
    ↓
1989–1991
    ↓
World Wide Web invented
Tim Berners-Lee at CERN
    ↓
1990s
    ↓
Commercial Internet
NSFNET restrictions removed
Commercial ISPs
    ↓
Mid/Late 1990s
    ↓
Web explosion
Netscape, Web commerce
Instant messaging
P2P file sharing
    ↓
2000–2001
    ↓
Dot-com crash
    ↓
New Millennium
    ↓
Broadband + WiFi + 4G
    ↓
Smartphones + mobile apps
    ↓
Social networks
    ↓
Private provider networks
    ↓
Cloud computing
```

---

# ⭐ What You Should Remember for the Exam

| Period | Most important development |
|---|---|
| **1961–1972** | Packet switching + ARPANET |
| **1969** | First four ARPANET nodes |
| **1972** | NCP + first e-mail |
| **1972–1980** | Multiple networks + internetworking |
| **1970s** | Cerf & Kahn develop TCP/IP concepts |
| **1980–1990** | BITNET, CSNET, NSFNET; massive growth |
| **Jan. 1, 1983** | TCP/IP replaces NCP on ARPANET |
| **Late 1980s** | TCP congestion control + DNS |
| **1989–1991** | World Wide Web developed at CERN |
| **1990s** | Internet commercialization + Web explosion |
| **1995–2001** | Internet startup boom and dot-com crash |
| **New millennium** | Broadband, wireless, smartphones, social networks, private networks, cloud |

---

## ⭐ The central evolution

If you need to explain the entire history in **one logical chain**, write it this way:

> **The increasing use of computers created a need for efficient communication of bursty computer traffic, leading to packet-switching research in the 1960s. ARPANET became the first major packet-switched computer network. During the 1970s, many independent packet-switched networks appeared, creating the need to interconnect heterogeneous networks. Cerf and Kahn developed the internetworking architecture that led to TCP/IP. During the 1980s, networks such as NSFNET expanded Internet connectivity, and TCP/IP replaced NCP on ARPANET in 1983. In the 1990s, the World Wide Web and commercial ISPs caused explosive Internet growth. In the new millennium, broadband, wireless access, smartphones, social networks, private provider networks, and cloud computing transformed the Internet into the globally distributed infrastructure we use today.** Computer Networking- A Top-Down… Computer Networking- A Top-Down… Computer Networking- A Top-Down… Computer Networking- A Top-Down… Computer Networking- A Top-Down…