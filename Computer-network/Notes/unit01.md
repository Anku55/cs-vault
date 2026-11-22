# Computer Networks -- Unit 1 Notes

## Table of Contents

-   [1.1 Network Fundamentals](#11-network-fundamentals)
    -   [1.1.1 Definition of a Computer
        Network](#111-what-is-a-computer-network)
    -   [1.1.2 Characteristics, Advantages and
        Disadvantages](#112-characteristics-advantages-and-disadvantages)
    -   [1.1.3 Components of a Computer
        Network](#113-components-of-a-computer-network)
    -   [1.1.4 Network Devices](#114-network-devices)
    -   [1.1.5 Types of Networks](#115-types-of-networks)
    -   [1.1.6 Client-Server
        Architecture](#116-client-server-architecture)
    -   [1.1.7 Peer-to-Peer
        Architecture](#117-peer-to-peer-architecture)
    -   [1.1.8 Network Performance
        Metrics](#118-network-performance-metrics)
-   [1.2 Network Topologies](#12-network-topologies)
    -   [1.2.1 Bus Topology](#121-bus-topology)
    -   [1.2.2 Star Topology](#122-star-topology)
    -   [1.2.3 Ring Topology](#123-ring-topology)
    -   [1.2.4 Mesh Topology](#124-mesh-topology)
    -   [1.2.5 Tree Topology](#125-tree-topology)
    -   [1.2.6 Hybrid Topology](#126-hybrid-topology)
    -   [1.2.7 Physical vs Logical
        Topology](#127-physical-vs-logical-topology)
-   [1.3 OSI Reference Model](#13-osi-reference-model)
    -   [1.3.1 What is the OSI Model?](#131-what-is-the-osi-model)
    -   [1.3.2 The Seven Layers](#132-the-seven-layers-of-the-osi-model)
    -   [1.3.3 Protocols Associated with OSI
        Layers](#133-protocols-associated-with-osi-layers)
    -   [1.3.4 PDU](#134-pdu--protocol-data-unit)
    -   [1.3.5 Encapsulation and
        Decapsulation](#135-encapsulation-and-decapsulation)
    -   [1.3.6 Service, Interface and
        Protocol](#136-service-interface-and-protocol)
    -   [1.3.7 Connection-Oriented vs Connectionless
        Communication](#137-connection-oriented-vs-connectionless-communication)
    -   [1.3.8 Peer-to-Peer Communication Between
        Layers](#138-peer-to-peer-communication-between-layers)
    -   [1.3.9 OSI vs TCP/IP](#139-osi-model-vs-tcpip-model)
-   [1.4 TCP/IP Protocol Suite](#14-tcpip-protocol-suite)
    -   [1.4.1 TCP/IP Architecture](#141-what-is-tcpip)
    -   [1.4.2 How Data Travels Across a
        Network](#142-how-data-travels-from-one-computer-to-another)
-   [1.5 Physical Layer Fundamentals](#15-physical-layer-fundamentals)
    -   [1.5.1 Analog Signals](#151-analog-signals)
    -   [1.5.2 Digital Signals](#152-digital-signals)
    -   [1.5.3 Periodic and Non-Periodic
        Signals](#153-periodic-and-non-periodic-signals)
    -   [1.5.4 Signal Characteristics](#154-characteristics-of-a-signal)
    -   [1.5.5 Bit Rate vs Baud Rate](#155-bit-rate-vs-baud-rate)
    -   [1.5.6 Bandwidth and Frequency](#156-bandwidth-and-frequency)
    -   [1.5.7 Noise and Attenuation](#157-noise-and-attenuation)
    -   [1.5.8 Signal-to-Noise Ratio
        (SNR)](#158-signal-to-noise-ratio-snr)
    -   [1.5.9 Nyquist Bit Rate Theorem](#159-nyquist-bit-rate-theorem)
    -   [1.5.10 Shannon Capacity
        Theorem](#1510-shannon-capacity-theorem)
    -   [1.5.11 Encoding and
        Modulation](#1511-encoding-and-modulation-basics)
    -   [1.5.12 Multiplexing](#1512-multiplexing)
-   [1.6 Transmission Media](#16-transmission-media)
    -   [1.6.1 Guided Media](#161-guided-transmission-media)
    -   [1.6.2 Unguided Media](#162-unguided-transmission-media)
    -   [1.6.3 Guided vs Unguided Media](#163-guided-vs-unguided-media)
    -   [1.6.4 Simplex, Half-Duplex and
        Full-Duplex](#164-simplex-half-duplex-and-full-duplex)
-   [1.7 Formula Sheet](#17-formula-sheet)
-   [1.8 Textbook References](#18-textbook-references)

------------------------------------------------------------------------

# 1.1 Network Fundamentals

## 1.1.1 What is a Computer Network?

A **computer network** is a collection of two or more computing devices
connected through communication links so that they can exchange data and
share resources.

These devices communicate using predefined rules called **network
protocols**.

Examples of connected devices:

-   Desktop computers
-   Laptops
-   Smartphones
-   Servers
-   Printers
-   Routers
-   Smart TVs
-   IoT devices

According to Kurose and Ross, the Internet connects billions of
computing devices. These devices are called **hosts or end systems**.

### Basic diagram

``` text
 Computer A                 Computer B
     |                           |
     +-------- Switch -----------+
                  |
                Router
                  |
               Internet
                  |
                Server
```

### How does it work?

Suppose you want to send a file from your laptop to another computer.

1.  Your laptop generates the data.
2.  Network software divides the data into smaller units called packets.
3.  Packets receive relevant control information in headers.
4.  Packets travel through communication links and networking devices.
5.  The destination computer receives the packets.
6.  The data is reconstructed at the destination.

This is the fundamental idea behind packet-switched communication.

**Real-life example:** When you send a photograph through WhatsApp, the
application sends data over the network. The receiving device eventually
reconstructs the photograph from the transmitted data.

## 1.1.2 Characteristics, Advantages and Disadvantages

### Characteristics

  -----------------------------------------------------------------------
  Characteristic                      Explanation
  ----------------------------------- -----------------------------------
  Connectivity                        Connects multiple devices

  Resource sharing                    Allows sharing of printers, files
                                      and storage

  Communication                       Supports email, messaging and video
                                      calls

  Scalability                         New devices can be added

  Reliability                         Redundancy can provide alternative
                                      communication paths

  Distributed processing              Work can be distributed across
                                      multiple machines

  Security                            Authentication, encryption and
                                      access control protect resources
  -----------------------------------------------------------------------

### Advantages

1.  **Resource sharing:** Multiple users can use the same printer or
    storage system.
2.  **Data sharing:** Users can exchange files without physically
    transferring storage devices.
3.  **Communication:** Supports video conferencing, messaging and email.
4.  **Centralized management:** Organizations can manage users and
    resources centrally.
5.  **Cost reduction:** Shared resources reduce infrastructure costs.
6.  **Remote access:** Employees can access organizational resources
    remotely.
7.  **Cloud computing:** Applications and storage can be accessed over
    networks.

### Disadvantages

-   Security vulnerabilities.
-   Network failures can interrupt services.
-   Initial infrastructure costs.
-   Network administration complexity.
-   Malware can spread between connected devices.
-   Congestion can reduce performance.

## 1.1.3 Components of a Computer Network

A computer network consists of several important components.

### 1. End devices (Hosts)

These are devices that communicate over a network.

Examples: laptop, smartphone, web server and desktop computer.

A host can act as a sender, receiver, or both.

### 2. Network Interface Card (NIC)

A NIC provides the interface through which a device connects to a
network.

It can be:

-   Wired Ethernet NIC.
-   Wireless Wi-Fi NIC.

A NIC generally provides a link-layer address called a **MAC address**.

``` text
Laptop
  |
 Ethernet / Wi-Fi
  |
 NIC
  |
Network
```

### 3. Transmission media

Transmission media carry data between devices.

Two major categories:

-   Guided media: copper cables and optical fiber.
-   Unguided media: radio waves, microwaves and infrared.

### 4. Networking devices

These devices connect computers, forward data, amplify signals or
translate between networks.

Examples include routers, switches, hubs and access points.

### 5. Network protocols

Protocols define how devices communicate.

  Protocol   Purpose
  ---------- ----------------------------------------
  HTTP       Web communication
  HTTPS      Secure web communication
  TCP        Reliable transport
  UDP        Connectionless transport
  IP         Logical addressing and packet delivery
  DNS        Domain name resolution
  Ethernet   Wired LAN communication
  Wi-Fi      Wireless LAN communication
  DHCP       Automatic IP configuration

## 1.1.4 Network Devices

### A. Hub

A **hub** is a basic networking device that connects multiple devices in
a LAN. It operates at the Physical Layer (Layer 1) of the OSI model.

**Working:** Suppose Computer A sends data to Computer B. The hub
receives the electrical signal from Computer A and repeats it to all
other connected ports. It does not examine destination MAC addresses.

``` text
             Computer A
                  |
             +---- HUB ----+
             |      |      |
           PC B    PC C   PC D
```

All connected devices receive the transmitted signal.

**Characteristics:**

-   Operates at Layer 1.
-   Does not maintain a MAC address table.
-   Does not filter frames.
-   Creates one shared collision domain.
-   Usually operates in half-duplex mode.

**Real-life analogy:** Announcing something through a loudspeaker in a
classroom. Everyone hears the announcement, even if it is intended for
only one student.

### B. Switch

A **network switch** connects multiple devices within a LAN and forwards
Ethernet frames based on MAC addresses. A traditional Ethernet switch
operates at Layer 2.

**How does it work?**

Suppose PC1 wants to communicate with PC3.

1.  PC1 sends an Ethernet frame.
2.  The switch receives the frame.
3.  It examines the source MAC address and learns which port the source
    device is connected to.
4.  It checks the destination MAC address.
5.  It looks up the destination in its MAC address table.
6.  It forwards the frame through the appropriate port.

Example MAC address table:

  MAC address   Switch port
  ------------- -------------
  AA:AA:AA      Port 1
  BB:BB:BB      Port 2
  CC:CC:CC      Port 3

If the destination MAC is CC:CC:CC, the switch forwards the frame to
Port 3.

If the destination MAC is unknown, the switch generally floods the frame
through other ports in the same VLAN.

**Advantages:**

-   Reduces unnecessary traffic compared with hubs.
-   Each switch port is a separate collision domain.
-   Supports full-duplex communication.
-   Improves LAN efficiency.

**Real-life analogy:** A receptionist who knows every employee's office
number and forwards a message only to the intended employee.

### C. Router

A **router** is a networking device that connects different IP networks
and forwards packets based on destination IP addresses. It primarily
operates at Layer 3 (Network Layer).

**How does it work?**

Suppose:

-   Network A: `192.168.1.0/24`
-   Network B: `192.168.2.0/24`

A computer in Network A wants to communicate with a computer in Network
B.

1.  The source computer identifies that the destination is outside its
    local subnet.
2.  It sends the packet to its default gateway.
3.  The router examines the destination IP address.
4.  It checks its routing table.
5.  It selects the appropriate outgoing interface or next hop.
6.  It forwards the packet.

``` text
Network A                       Network B
192.168.1.0/24                  192.168.2.0/24

 PC1 -------- Router ----------- PC2
               |
          Routing Table
```

**Important distinction:**

-   Switch: Forwards frames using MAC addresses.
-   Router: Forwards packets using IP addresses.

**Real-life example:** Your home Wi-Fi router connects your private home
network to your ISP's network and ultimately to the Internet.

### D. Bridge

A bridge connects two or more LAN segments and filters traffic using MAC
addresses. It operates at Layer 2.

Its purpose is to reduce unnecessary traffic between network segments.

``` text
 LAN Segment 1
 PC1 ---- PC2
      |
    Bridge
      |
 PC3 ---- PC4
 LAN Segment 2
```

A bridge learns MAC addresses and forwards frames according to its
forwarding table.

A modern Ethernet switch is essentially a multiport bridge.

### E. Repeater

A repeater regenerates or repeats signals to extend the transmission
distance. It operates at Layer 1.

``` text
Sender ---- Weak signal ---- Repeater ---- Regenerated signal ---- Receiver
```

**Example:** A Wi-Fi range extender can receive a wireless signal and
retransmit it to improve coverage. It is not necessarily identical
internally to a basic physical-layer repeater.

### F. Gateway

A gateway connects networks or systems that may use different protocols
or communication architectures.

It can perform protocol conversion, address translation, or
application-level processing.

Example: An organization may use a gateway to connect an internal
network to an external network.

**Important:** In everyday networking, the term default gateway commonly
refers to the router that forwards traffic outside the local subnet. A
general-purpose gateway has a broader meaning.

### G. Modem

Modem stands for **MO**dulator + **DEM**odulator.

A modem converts signals between forms suitable for different
communication systems.

For example, a cable modem communicates over the cable provider's
network and provides an interface to the customer's local network.

``` text
Home Network
     |
   Router
     |
   Modem
     |
ISP Network
     |
  Internet
```

Modern equipment may combine modem and router functionality in one
physical device.

### H. Wireless Access Point (AP)

An access point allows wireless devices to connect to a wired network
using Wi-Fi.

Example: Your college Wi-Fi access point connects your laptop wirelessly
to the college's LAN.

An AP generally bridges wireless and wired LAN traffic at Layer 2.

### Comparison of networking devices

  Device                  Main OSI layer Primary function
  -------------- ----------------------- --------------------------------------
  Hub                                  1 Repeats signals to all ports
  Repeater                             1 Regenerates signals
  Switch                               2 Forwards frames using MAC
  Bridge                               2 Connects and filters LAN segments
  Router                               3 Routes IP packets
  Gateway                        Depends Connects different systems/protocols
  Modem            Depends on technology Modulates/demodulates signals
  Access Point               Primarily 2 Connects wireless clients to LAN

## 1.1.5 Types of Networks

Networks can be classified according to geographical coverage.

### A. PAN -- Personal Area Network

A PAN is a network that connects devices within a small area around an
individual.

Typical range: a few meters, depending on the technology.

Examples:

-   Smartphone connected to Bluetooth earbuds.
-   Smartwatch connected to a smartphone.
-   Laptop connected to a phone hotspot.

Technologies: Bluetooth, USB, NFC and personal Wi-Fi connections.

### B. LAN -- Local Area Network

A LAN connects devices within a limited geographical area.

Examples:

-   Home network.
-   College computer laboratory.
-   Office network.
-   Computer lab.

Characteristics:

-   Usually privately managed.
-   High-speed communication.
-   Limited geographical coverage.
-   Ethernet and Wi-Fi are common technologies.

``` text
        College LAN
     +---------------+
     |               |
   PC 1            PC 2
     |               |
     +---- Switch ---+
              |
           Server
```

### C. MAN -- Metropolitan Area Network

A MAN connects multiple LANs across a city or metropolitan region.

Example: A university with multiple campuses across the same city can
connect its campuses through a metropolitan network.

Characteristics:

-   Larger coverage than LAN.
-   Smaller coverage than WAN.
-   Often uses high-capacity fiber-optic links.

### D. WAN -- Wide Area Network

A WAN connects networks over large geographical distances.

Examples:

-   A bank connecting branches across India.
-   A multinational organization connecting offices worldwide.
-   The Internet.

Characteristics:

-   Large geographical coverage.
-   Uses routers and long-distance communication links.
-   Can use leased lines, fiber, satellite and other technologies.

### Comparison

  -----------------------------------------------------------------------------------
  Feature        PAN            LAN               MAN                 WAN
  -------------- -------------- ----------------- ------------------- ---------------
  Coverage       Personal       Building/campus   City                Country/world

  Example        Bluetooth      College network   City network        Internet

  Ownership      Individual     Organization      Multiple            Multiple
                                                  entities/provider   providers

  Typical        Bluetooth      Ethernet/Wi-Fi    Metro               MPLS, leased
  technologies                                    Ethernet/fiber      lines, IP
  -----------------------------------------------------------------------------------

## 1.1.6 Client-Server Architecture

Client-server architecture is a networking model in which one system,
called the client, requests services or resources from another system,
called the server.

``` text
       Mobile Client
             |
             | HTTP Request
             v
       Application Server
             |
       Process Request
             |
             v
        HTTP Response
             |
             v
       Mobile Client
```

**How it works:**

1.  Client sends a request.
2.  Server receives the request.
3.  Server processes it.
4.  Server sends a response.
5.  Client displays or processes the result.

**Example:** Opening YouTube.

Other examples:

-   Banking applications.
-   E-commerce applications.
-   College management systems.
-   Web applications.

**Go backend connection:** When you build an HTTP server using Go's
`net/http` package, your Go program acts as a server. A browser or
Postman acts as the client.

### Advantages

-   Centralized management.
-   Easier security enforcement.
-   Centralized data storage.
-   Easier backup and maintenance.

### Disadvantages

-   Server can become a bottleneck.
-   Server failure can affect many clients.
-   Infrastructure can become expensive as traffic grows.

## 1.1.7 Peer-to-Peer Architecture

In a peer-to-peer (P2P) architecture, computers communicate directly and
can act as both clients and servers.

``` text
       Peer A
       /    \
      /      \
  Peer B ---- Peer C
      \       /
       \     /
        Peer D
```

Each peer can:

-   Request data.
-   Provide data.
-   Share computing resources.
-   Communicate with other peers.

Examples:

-   BitTorrent.
-   Some blockchain networks.
-   Distributed file-sharing systems.

### Client-server vs P2P

  -----------------------------------------------------------------------
  Feature                 Client-server           P2P
  ----------------------- ----------------------- -----------------------
  Architecture            Centralized             Distributed

  Main roles              Separate client/server  Each peer can perform
                          roles                   both roles

  Management              Centralized             Distributed

  Scalability             Server capacity can     Resources can grow as
                          limit growth            peers join

  Failure                 Central server can be a Other peers may
                          single point of failure continue operating

  Example                 Traditional web         BitTorrent
                          application             
  -----------------------------------------------------------------------

Kurose and Ross discuss both client-server and P2P application
architectures in Chapter 2, Section 2.1.

## 1.1.8 Network Performance Metrics

Network performance determines how efficiently data moves from source to
destination.

Five important terms:

1.  Bandwidth
2.  Throughput
3.  Latency
4.  Jitter
5.  Packet loss

### A. Bandwidth

Bandwidth can refer to different things depending on context.

In digital networking, it commonly means the maximum data-carrying
capacity of a link, measured in bits per second.

Examples:

-   100 Mbps Ethernet.
-   1 Gbps Ethernet.
-   100 Mbps Internet plan.

In signal processing, bandwidth specifically means the range of
frequencies a signal occupies, measured in Hz.

Do not confuse these two meanings.

### B. Throughput

Throughput is the actual rate at which data is successfully transferred.

Example: Your Internet plan provides 100 Mbps, but your speed test shows
75 Mbps.

-   Nominal link capacity: 100 Mbps.
-   Measured throughput: 75 Mbps.

Throughput is generally affected by congestion, protocol overhead,
receiver limitations and other bottlenecks.

### C. Latency

Latency is the time taken for data to travel from one point to another.
It is commonly measured in milliseconds.

The delay depends on:

-   Propagation delay.
-   Transmission delay.
-   Processing delay.
-   Queuing delay.

### D. Jitter

Jitter is the variation in packet delay over time.

Suppose packets arrive with these delays:

  Packet     Delay
  -------- -------
  1          20 ms
  2          22 ms
  3          45 ms
  4          21 ms
  5          60 ms

The variation in arrival delays is jitter.

Jitter is especially important in video conferencing, online gaming,
voice-over-IP and live streaming.

### E. Packet loss

Packet loss occurs when packets fail to reach their destination.

Possible reasons:

-   Congestion.
-   Faulty cables.
-   Wireless interference.
-   Hardware failures.
-   Buffer overflow.

Example: A video call may freeze when packets are dropped because a
router's queue becomes full.

### Performance metrics comparison

  Metric        Meaning                       Unit
  ------------- ----------------------------- ------
  Bandwidth     Available link capacity       bps
  Throughput    Actual successful data rate   bps
  Latency       Communication delay           ms
  Jitter        Variation in delay            ms
  Packet loss   Fraction of lost packets      \%

### Numerical example

Suppose a 100 MB file is transferred over a link with a throughput of 20
Mbps.

Convert MB to megabits:

\[ 100 `\times 8`{=tex} = 800`\text{ Mb}`{=tex} \]

Transfer time:

\[ T = `\frac{\text{Data Size}}{\text{Throughput}}`{=tex} \]

\[ T = `\frac{800}{20}`{=tex} = 40`\text{ seconds}`{=tex} \]

**Answer: 40 seconds.** This assumes a constant throughput of 20 Mbps
and ignores additional transfer overhead.

------------------------------------------------------------------------

# 1.2 Network Topologies

A **network topology** describes how computers, networking devices and
communication links are arranged in a network.

There are two important ways to understand topology:

1.  Physical topology -- How devices and cables are physically
    connected.
2.  Logical topology -- How data actually flows between devices.

Six major topologies are included in this syllabus.

## 1.2.1 Bus Topology

In bus topology, all devices are connected to a single main
communication cable called the **backbone cable**.

``` text
 [PC1]      [PC2]      [PC3]      [PC4]
   |          |          |          |
===+==========+==========+==========+===
              Backbone
```

### Working

1.  A computer sends data onto the backbone cable.
2.  The signal travels along the shared communication medium.
3.  All connected devices can receive the signal.
4.  Only the device with the matching destination address processes the
    frame.
5.  Terminators at the ends prevent signal reflections.

### Advantages

-   Simple installation.
-   Requires less cable than some other topologies.
-   Low initial cost for small networks.
-   Does not require a central switch.

### Disadvantages

-   Backbone failure can bring down the entire network.
-   Performance decreases as network traffic increases.
-   Fault identification can be difficult.
-   Limited cable length and number of devices.
-   Shared-medium collisions can occur in traditional bus Ethernet.

**Real-life example:** Traditional coaxial Ethernet networks, such as
10BASE2 and 10BASE5.

## 1.2.2 Star Topology

In star topology, all devices connect to a central device, usually a
switch or hub.

``` text
             PC1
              |
 PC2 -------- Switch -------- PC3
             /     \
           PC4     PC5
              |
             PC6
```

### Working

Suppose PC1 sends data to PC3.

1.  PC1 sends a frame to the central switch.
2.  The switch examines the destination MAC address.
3.  It forwards the frame to PC3's port.
4.  PC3 receives the frame.

### Advantages

-   Easy to install and maintain.
-   Failure of one endpoint cable usually affects only that device.
-   Easy troubleshooting.
-   Easy to expand.
-   Good performance with modern switches.

### Disadvantages

-   Central switch failure can affect the entire network.
-   Requires more cable than bus topology.
-   Additional cost for the central device.

**Real-life example:** A college computer laboratory, where every
computer connects to a central Ethernet switch.

## 1.2.3 Ring Topology

In ring topology, each device connects to two neighboring devices,
forming a closed loop.

``` text
        PC1 -------- PC2
         |            |
         |            |
        PC4 -------- PC3
```

### Working

In a traditional unidirectional ring:

1.  Data travels from one device to the next.
2.  Each device examines the destination address.
3.  If it is not the destination, the data is forwarded.
4.  The destination receives the data.

Some ring networks use a token-passing mechanism.

A token is a special control frame that grants permission to transmit.

Only the device holding the token can transmit, reducing collisions.

### Advantages

-   Predictable access in token-based networks.
-   Avoids collisions in token-passing implementations.
-   Can provide orderly communication.

### Disadvantages

-   A single break can disrupt a basic ring.
-   Adding or removing devices can be difficult.
-   Troubleshooting can be complex.
-   Delay can increase as the number of nodes increases.

**Real-life example:** Token Ring networks and certain fiber-optic ring
infrastructures.

## 1.2.4 Mesh Topology

In mesh topology, devices are interconnected through multiple links.

There are two types:

-   **Full mesh:** Every device has a direct connection to every other
    device.
-   **Partial mesh:** Only selected devices have multiple direct
    connections.

### Full mesh connection formula

For (n) devices, the number of required links is:

\[ L = `\frac{n(n-1)}{2}`{=tex} \]

Example: Five devices.

\[ L = `\frac{5(5-1)}{2}`{=tex} \]

\[ L = `\frac{20}{2}`{=tex}=10 \]

Therefore, 10 links are required.

### Advantages

-   High reliability.
-   Multiple alternative paths.
-   Good fault tolerance.
-   Failure of one link does not necessarily disconnect devices.

### Disadvantages

-   Expensive installation.
-   Requires many cables and ports.
-   Complex configuration and maintenance.

**Real-life example:** Mesh wireless networks and redundant
interconnections between core routers in enterprise networks.

## 1.2.5 Tree Topology

Tree topology is a hierarchical network arrangement that combines
star-like structures.

``` text
               Core Switch
                /       \
          Switch A     Switch B
          /    \        /    \
        PC1   PC2     PC3   PC4
```

### Working

-   A central/root device connects to lower-level switches.
-   Each lower-level switch connects to additional devices.
-   Communication passes through the hierarchy.

### Advantages

-   Easy to expand.
-   Suitable for large organizations.
-   Supports hierarchical administration.
-   Faults can sometimes be isolated to individual branches.

### Disadvantages

-   Failure of an important upper-level device affects multiple
    branches.
-   Requires more cabling.
-   Configuration is more complex.

**Real-life example:** A university network with a core switch,
departmental switches and laboratory computers.

## 1.2.6 Hybrid Topology

Hybrid topology combines two or more different network topologies.

Example: An organization can have:

-   Star topology within each department.
-   A redundant backbone connecting departments.
-   Mesh connections between important routers.

### Advantages

-   Flexible.
-   Scalable.
-   Can be designed for reliability.
-   Suitable for large organizations.

### Disadvantages

-   Expensive.
-   Difficult to design.
-   Requires skilled network administrators.

## 1.2.7 Physical vs Logical Topology

**Physical topology** describes the actual physical arrangement of
devices and cables.

**Logical topology** describes how data moves between devices,
regardless of their physical arrangement.

Example: A switched Ethernet network physically follows a star
arrangement, but the logical communication between devices is determined
by Ethernet switching and the network protocols.

Another example is Wi-Fi: devices may physically connect to the same
access point, while the wireless medium is shared for communication.

  Physical topology              Logical topology
  ------------------------------ ----------------------------------
  Physical arrangement           Data-flow arrangement
  Describes cables and devices   Describes communication behavior
  Concerned with installation    Concerned with transmission
  Example: Star                  Example: Shared wireless medium

### Comparison of topologies

  --------------------------------------------------------------------------------
  Topology       Cost             Reliability        Scalability    Main weakness
  -------------- ---------------- ------------------ -------------- --------------
  Bus            Low              Low                Limited        Backbone
                                                                    failure

  Star           Moderate         Good               Good           Central device
                                                                    failure

  Ring           Moderate         Depends on design  Moderate       Ring
                                                                    disruption

  Mesh           High             Very high          Expensive to   Many
                                                     expand         connections

  Tree           Moderate--high   Depends on         High           Root-level
                                  hierarchy                         failure

  Hybrid         High             Design-dependent   High           Complexity
  --------------------------------------------------------------------------------

------------------------------------------------------------------------

# 1.3 OSI Reference Model

## 1.3.1 What is the OSI Model?

OSI stands for **Open Systems Interconnection**.

It is a conceptual reference model developed by the International
Organization for Standardization (ISO).

It divides network communication into seven layers.

The main purpose is to standardize network communication and divide
complex communication tasks into smaller, manageable responsibilities.

### Why do we need layers?

Imagine developing a video-calling application.

The application requires:

-   User interface and application protocols.
-   Data representation and encryption.
-   Session management.
-   Reliable or real-time transport.
-   Logical addressing and routing.
-   Local network delivery.
-   Physical transmission of bits.

If one component had to handle everything, the system would become
extremely complex.

Layering separates responsibilities.

## 1.3.2 The Seven Layers of the OSI Model

Remember the layers from top to bottom:

**A P S T N D P**

    Layer Name           Main responsibility
  ------- -------------- ----------------------------------------
        7 Application    Network services for applications
        6 Presentation   Data representation and transformation
        5 Session        Session management
        4 Transport      End-to-end process communication
        3 Network        Logical addressing and routing
        2 Data Link      Framing and local link delivery
        1 Physical       Transmission of bits

### Layer 7 -- Application Layer

The Application Layer provides network services to application programs.
It is the layer closest to the end user.

It does not mean the application itself is the entire layer. Rather, it
provides protocols through which applications communicate over a
network.

**Responsibilities:**

-   Web communication.
-   Email services.
-   File transfer.
-   Name resolution.
-   Remote access.
-   Network resource sharing.

**Important protocols:**

  Protocol   Purpose
  ---------- ---------------------------
  HTTP       Transfers web resources
  HTTPS      Secure HTTP communication
  FTP        File transfer
  SMTP       Sending email
  DNS        Domain name resolution
  DHCP       Dynamic IP configuration
  SSH        Secure remote access
  SNMP       Network management

**Real-life example:** You open `https://www.google.com`. Your browser
uses application-layer protocols to communicate with Google's
infrastructure.

Example HTTP request:

``` http
GET / HTTP/1.1
Host: www.google.com
```

This is an application-layer HTTP message.

**Important:** HTTP is an application-layer protocol, not a
transport-layer protocol.

### Layer 6 -- Presentation Layer

The Presentation Layer handles the representation of data exchanged
between communicating systems.

Think of it as the data translation layer.

**Responsibilities:**

1.  Data translation.
2.  Data formatting.
3.  Character encoding.
4.  Compression.
5.  Encryption and decryption.

**1. Data translation**

Different systems may represent information differently.

Example:

``` text
Character: A

ASCII: 65
Unicode: U+0041
UTF-8: 41 (hexadecimal)
```

**2. Encryption and decryption**

Encryption converts readable data into an unreadable form. Decryption
converts it back into readable information.

``` text
Original data:
Hello

       Encryption

Encrypted data:
Ciphertext

       Decryption

Original data:
Hello
```

**Important real-world clarification:** TLS performs encryption and
related security functions in modern Internet communication. Although
encryption is traditionally associated with the OSI Presentation Layer
in conceptual teaching, TLS does not map neatly to a single OSI layer.

**3. Compression**

Compression reduces data size.

Examples: JPEG image compression, PNG encoding, GZIP compression and
video compression.

**Important:** Presentation-layer examples are conceptual mappings. Real
protocols can operate across or outside strict OSI boundaries.

### Layer 5 -- Session Layer

The Session Layer manages communication sessions between applications.

A session represents an organized interaction between communicating
systems.

**Responsibilities:**

-   Session establishment.
-   Session maintenance.
-   Session termination.
-   Dialog control.
-   Synchronization and checkpoints.

**Example:** Imagine transferring a very large file. If the transfer is
interrupted, a session-oriented system may use checkpoints to resume
from a previously completed point rather than restarting everything.

``` text
Client                     Server

   |------ Session Setup ----->|
   |                           |
   |------ Data Transfer ----->|
   |------ Checkpoint -------->|
   |                           |
   |------ Session Close ----->|
```

**Real-life example:** A remote desktop session allows you to interact
with another computer continuously. The Session Layer is a conceptual
way of understanding the management of that ongoing interaction.

### Layer 4 -- Transport Layer

The Transport Layer provides communication between application processes
running on different hosts.

It supports end-to-end transport services.

**Responsibilities:**

-   Segmentation and reassembly.
-   Process-to-process delivery.
-   Port addressing.
-   Multiplexing and demultiplexing.
-   Reliable data transfer when provided by the protocol.
-   Flow control.
-   Congestion control.

**Important protocols: TCP and UDP**

**TCP -- Transmission Control Protocol**

TCP provides:

-   Reliable data delivery.
-   Ordered byte-stream delivery.
-   Error detection.
-   Retransmission of lost data.
-   Flow control.
-   Congestion control.
-   Connection establishment and termination.

Example: Loading a website over HTTP/1.1 or HTTP/2 generally uses TCP.

**UDP -- User Datagram Protocol**

UDP provides:

-   Connectionless communication.
-   Datagram-based delivery.
-   Minimal transport overhead.
-   No built-in guarantee of delivery.
-   No built-in guarantee of ordering.

Example: DNS commonly uses UDP, although it can also use TCP.

  -----------------------------------------------------------------------
  Feature                 TCP                     UDP
  ----------------------- ----------------------- -----------------------
  Connection              Connection-oriented     Connectionless

  Reliability             Built-in reliable       No built-in delivery
                          delivery                guarantee

  Ordering                Ordered byte stream     No ordering guarantee

  Speed/overhead          More overhead           Less overhead

  Flow control            Yes                     No

  Congestion control      Yes                     No built-in transport
                                                  congestion control

  Examples                HTTP/1.1, HTTP/2        DNS, real-time media
  -----------------------------------------------------------------------

**Port numbers**

A port number identifies a transport-layer endpoint associated with an
application process.

     Port Protocol
  ------- ------------------------------
    20/21 FTP
       22 SSH
       25 SMTP
       53 DNS
       80 HTTP
      443 HTTPS
     3306 MySQL
     5432 PostgreSQL
     8080 Common alternative HTTP port

Example:

``` text
Client:
IP = 192.168.1.10
Port = 50000

Server:
IP = 142.250.183.14
Port = 443
```

A network connection is identified using information such as source and
destination IP addresses and port numbers, together with the transport
protocol.

**Go backend connection:** When you run:

``` go
http.ListenAndServe(":8080", nil)
```

Your application listens for HTTP connections on port 8080. The
operating system and transport stack handle the underlying TCP
communication.

### Layer 3 -- Network Layer

The Network Layer is responsible for logical addressing and routing
packets across interconnected networks.

Its major protocol is IP.

**Responsibilities:**

-   Logical addressing.
-   Routing.
-   Packet forwarding.
-   Fragmentation in IPv4 (where applicable).
-   Internetwork communication.

**Important protocols:**

-   IPv4
-   IPv6
-   ICMP
-   Routing protocols such as OSPF and BGP

An IP address identifies a network interface in an IP network.

Example:

``` text
192.168.1.10
```

Suppose:

``` text
Computer A
IP: 192.168.1.10

       |
       | Router
       |

Computer B
IP: 192.168.2.20
```

The router forwards packets between these different networks.

**Routing vs Forwarding**

-   **Routing:** Determining the paths packets should take through a
    network.
-   **Forwarding:** Moving an arriving packet from a router's input
    interface to the appropriate output interface.

Kurose and Ross explain this distinction in Chapter 4.

### Layer 2 -- Data Link Layer

The Data Link Layer provides communication between directly connected
network nodes over a particular link.

It takes a network-layer datagram and encapsulates it into a frame.

**Responsibilities:**

-   Framing.
-   MAC addressing.
-   Error detection.
-   Medium access control.
-   Link-level delivery.

**Important protocols and technologies:**

-   Ethernet.
-   Wi-Fi (IEEE 802.11).
-   PPP.
-   HDLC.
-   ARP is closely associated with IP-to-link-address resolution in IPv4
    networks, although its exact OSI mapping is not universal.

A MAC address is a link-layer address used to identify an interface on a
local network.

Example:

``` text
00:1A:2B:3C:4D:5E
```

Suppose your laptop communicates with your home router. The IP packet is
encapsulated inside an Ethernet or Wi-Fi frame, depending on the
connection. The frame is delivered over the local link to the next node.

**Important:** MAC addresses are used for local-link delivery. Routers
generally replace the link-layer frame when forwarding a packet onto
another link.

### Layer 1 -- Physical Layer

The Physical Layer transmits individual bits over a physical
communication medium.

It defines how bits are represented and transmitted.

**Responsibilities:**

-   Signal representation.
-   Bit transmission.
-   Electrical characteristics.
-   Optical characteristics.
-   Radio characteristics.
-   Connectors and physical interfaces.
-   Transmission timing.

Examples: copper Ethernet cable, optical fiber and wireless radio
transmission.

If a computer sends:

``` text
10110101
```

The physical layer represents these bits using electrical, optical or
radio signals. The receiving device detects the signal and reconstructs
the bit sequence.

## 1.3.3 Protocols Associated with OSI Layers

  -----------------------------------------------------------------------
  OSI Layer                           Protocols / technologies
  ----------------------------------- -----------------------------------
  Application                         HTTP, HTTPS, FTP, SMTP, DNS, DHCP,
                                      SSH

  Presentation                        TLS-related data protection, JPEG,
                                      ASCII, UTF-8, compression formats

  Session                             RPC, NetBIOS session services
                                      (conceptual examples)

  Transport                           TCP, UDP

  Network                             IPv4, IPv6, ICMP, OSPF, BGP

  Data Link                           Ethernet, PPP, HDLC, Wi-Fi MAC

  Physical                            100BASE-TX, optical fiber PHYs,
                                      radio PHYs
  -----------------------------------------------------------------------

Remember: Some protocols do not fit perfectly into one OSI layer. This
table represents common educational associations.

## 1.3.4 PDU -- Protocol Data Unit

PDU refers to the unit of data handled by a particular layer.

  Layer          PDU
  -------------- --------------------------------
  Application    Data / Message
  Presentation   Data
  Session        Data
  Transport      Segment (TCP) / Datagram (UDP)
  Network        Packet / IP Datagram
  Data Link      Frame
  Physical       Bits

Example: When you send a message:

``` text
Hello Ankush
```

The data is transformed as it moves down the protocol stack.

``` text
Application Layer
        DATA
         |
Transport Layer
     SEGMENT
         |
Network Layer
      PACKET
         |
Data Link Layer
       FRAME
         |
Physical Layer
        BITS
```

This process is called encapsulation.

## 1.3.5 Encapsulation and Decapsulation

Encapsulation is the process of adding protocol-specific headers (and
sometimes trailers) to data as it moves down the protocol stack.

### Step-by-step example

Suppose your browser sends:

``` text
Hello Server
```

**Step 1: Application Layer**

``` text
[Application Data]
```

**Step 2: Transport Layer**

TCP adds a TCP header.

``` text
[TCP Header][Application Data]
```

Now it is a TCP segment.

**Step 3: Network Layer**

IP adds an IP header.

``` text
[IP Header][TCP Header][Application Data]
```

Now it is an IP packet.

**Step 4: Data Link Layer**

Ethernet adds its header and trailer.

``` text
[Ethernet Header]
[IP Header]
[TCP Header]
[Application Data]
[Ethernet Trailer]
```

Now it is an Ethernet frame.

**Step 5: Physical Layer**

The frame is transmitted as bits.

``` text
101101010101010101...
```

### Decapsulation

At the destination, the reverse process happens.

``` text
Bits
  |
Frame
  |
IP Packet
  |
TCP Segment
  |
Application Data
```

Each layer examines and removes its relevant protocol information and
passes the remaining data upward.

### Why is encapsulation necessary?

-   Each layer can perform its own responsibilities.
-   Protocols can operate independently.
-   Different technologies can work together.
-   Communication becomes modular and manageable.

## 1.3.6 Service, Interface and Protocol

These three terms are closely related but have different meanings.

### Service

A service defines what a layer provides to the layer above it.

Example: The transport layer provides communication services to the
application layer.

### Interface

An interface defines how the upper layer accesses the services of the
lower layer.

Example: A socket API allows an application to use transport-layer
services.

### Protocol

A protocol defines the rules followed by peer entities at the same layer
on different systems.

Example: TCP defines how communicating transport-layer entities exchange
information.

  Term        Meaning
  ----------- --------------------------------------------
  Service     What a layer provides
  Interface   How a layer accesses that service
  Protocol    Rules used by peer entities to communicate

An important distinction: protocols operate logically between
corresponding peer layers, while services are provided upward through
interfaces.

## 1.3.7 Connection-Oriented vs Connectionless Communication

### Connection-oriented communication

A logical connection is established before data transfer.

Example: TCP.

``` text
Client                 Server
  |------ SYN ---------->|
  |<----- SYN-ACK -------|
  |------ ACK ---------->|
  |                      |
  |===== Data Transfer ==|
  |                      |
```

TCP uses a three-way handshake to establish a connection.

**Advantages:**

-   Reliable delivery.
-   Ordered byte-stream delivery.
-   Retransmission when necessary.
-   Flow control.

**Disadvantages:**

-   Additional overhead.
-   Connection establishment takes time.

### Connectionless communication

Data is transmitted without establishing a transport connection first.

Example: UDP.

``` text
Client                 Server

  |------ Datagram ------>|
  |------ Datagram ------>|
  |------ Datagram ------>|
```

**Advantages:**

-   Low overhead.
-   No transport connection establishment.
-   Useful for time-sensitive communication.

**Disadvantages:**

-   No built-in delivery guarantee.
-   No built-in ordering.
-   Applications must implement additional reliability if required.

## 1.3.8 Peer-to-Peer Communication Between Layers

The OSI model uses the concept of peer entities.

``` text
Computer A                 Computer B

Application  <----------> Application
Transport    <----------> Transport
Network      <----------> Network
Data Link    <----------> Data Link
Physical     <----------> Physical
```

This diagram represents logical communication between corresponding
layers.

In reality, data moves down the source stack, across the network, and up
the destination stack.

A layer uses the services of the layer directly below it and provides
services to the layer directly above it.

## 1.3.9 OSI Model vs TCP/IP Model

  Feature                      OSI               TCP/IP
  ---------------------------- ----------------- ------------------------------
  Developed by                 ISO               DARPA/DoD research community
  Layers                       7                 Commonly 4
  Nature                       Reference model   Protocol architecture
  Application-related layers   3                 1
  Transport                    1                 1
  Network                      1                 1
  Link and Physical            2                 1

The four-layer TCP/IP model:

1.  Application.
2.  Transport.
3.  Internet.
4.  Network Access.

The five-layer model commonly used for teaching:

1.  Application.
2.  Transport.
3.  Network.
4.  Data Link.
5.  Physical.

### Mapping

  OSI            TCP/IP four-layer
  -------------- -------------------
  Application    Application
  Presentation   Application
  Session        Application
  Transport      Transport
  Network        Internet
  Data Link      Network Access
  Physical       Network Access

**Exam point:** Do not confuse the four-layer and five-layer TCP/IP
representations. Both are used in networking education, but they group
the lower layers differently.

------------------------------------------------------------------------

# 1.4 TCP/IP Protocol Suite

## 1.4.1 What is TCP/IP?

TCP/IP stands for Transmission Control Protocol / Internet Protocol.

It is the fundamental protocol architecture used for communication over
the Internet.

It defines how data is:

-   Generated by applications.
-   Transported between processes.
-   Addressed across networks.
-   Delivered over individual links.

### TCP/IP architecture

The four-layer TCP/IP model:

  Layer            Main responsibility                Examples
  ---------------- ---------------------------------- ----------------------
  Application      Network application services       HTTP, DNS, SMTP, FTP
  Transport        Process-to-process communication   TCP, UDP
  Internet         Logical addressing and routing     IPv4, IPv6, ICMP
  Network Access   Local network communication        Ethernet, Wi-Fi, PPP

The five-layer representation separates the Network Access Layer into
Data Link and Physical Layers.

## 1.4.2 How Data Travels from One Computer to Another

**Scenario:** You open `www.example.com` on your laptop.

### Step 1: Application layer

The browser needs to retrieve a web page. It generates an HTTP request.

``` http
GET / HTTP/1.1
Host: www.example.com
```

### Step 2: DNS resolution

The browser needs the server's IP address. It uses DNS to resolve the
domain name.

``` text
www.example.com
       |
       v
DNS Server
       |
       v
93.184.216.34
```

The IP shown is an illustrative example, not a claim about the current
address of the domain.

### Step 3: Transport layer

The browser establishes a TCP connection if using HTTP over TCP.

TCP:

-   Establishes a connection.
-   Segments the application data.
-   Adds source and destination port numbers.
-   Provides reliable, ordered delivery.

For HTTPS over TCP, TLS negotiation occurs before protected HTTP
communication.

### Step 4: Network layer

IP adds source and destination IP addresses.

``` text
Source IP:      192.168.1.10
Destination IP: 93.184.216.34
```

The packet is routed toward the destination.

### Step 5: Data Link Layer

The IP packet is encapsulated into a frame. The frame contains
link-layer addressing for the local hop.

For example, your laptop sends the frame to the MAC address of its
default gateway.

### Step 6: Physical Layer

The frame is transmitted through the physical medium.

Depending on the connection, this could be:

-   Wi-Fi radio signals.
-   Ethernet electrical signals.
-   Optical signals over fiber.

### Step 7: Router processing

Your home router receives the frame.

It:

1.  Processes the incoming frame.
2.  Extracts the IP packet.
3.  Examines the destination IP.
4.  Checks its routing table.
5.  Forwards the packet toward the ISP.

The frame is rebuilt for the next link.

### Step 8: Destination server

The destination server receives the packet.

The receiving protocol stack performs decapsulation and delivers the
application data to the web server.

### Step 9: Response

The web server generates an HTTP response. The response travels back
through the network to your laptop. Your browser processes the response
and displays the webpage.

**Important:** This example describes the conceptual protocol flow.
Actual Internet paths may involve many routers, caches, proxies, CDNs
and different link technologies.

------------------------------------------------------------------------

# 1.5 Physical Layer Fundamentals

The Physical Layer is responsible for transmitting raw bits from one
device to another through a communication medium.

It deals with the physical representation of information using
electrical, optical or electromagnetic signals.

A fundamental idea:

**Computers process digital data, but communication media carry physical
signals.**

``` text
Computer A
    |
Binary Data: 101101
    |
Signal Conversion
    |
Electrical / Optical / Radio Signal
    |
Communication Medium
    |
Signal Detection
    |
Binary Data: 101101
    |
Computer B
```

## 1.5.1 Analog Signals

An analog signal is a continuous-time signal whose amplitude can vary
continuously over a range of values.

Examples:

-   Human voice.
-   Traditional radio signals.
-   Temperature variations.
-   Sound waves.

An analog signal can have infinitely many possible amplitude values
within its operating range.

For example, a microphone converts sound pressure variations into a
corresponding electrical signal.

### Characteristics

-   Continuous amplitude.
-   Can vary smoothly with time.
-   Can be affected by noise.
-   Commonly represented using sine waves.

## 1.5.2 Digital Signals

A digital signal represents information using discrete signal levels.

In binary communication, the two basic logical values are:

-   0
-   1

For example:

``` text
Binary data:

1 0 1 1 0 1
```

A digital transmission system represents these bits using defined
electrical, optical or radio signal states.

### Characteristics

-   Discrete signal levels.
-   Suitable for digital processing.
-   Can support reliable regeneration.
-   Noise can still affect transmission.
-   Widely used in computer networks.

### Analog vs Digital

  -----------------------------------------------------------------------
  Feature                 Analog                  Digital
  ----------------------- ----------------------- -----------------------
  Nature                  Continuous              Discrete

  Values                  Continuous range        Defined levels

  Noise handling          Often accumulates       Can regenerate signals
                          degradation             

  Representation          Continuous waveform     Digital symbols/bits

  Example                 Microphone output       Ethernet data
  -----------------------------------------------------------------------

**Important:** Digital communication is not necessarily transmitted
using square waves. Different physical-layer technologies use different
signal encodings and modulation schemes.

## 1.5.3 Periodic and Non-Periodic Signals

### Periodic signal

A periodic signal repeats itself after a fixed time interval. This
interval is called the time period.

Examples: sine wave and regular clock signal.

Mathematically:

\[ x(t)=x(t+T) \]

Where:

-   (x(t)) = signal at time (t).
-   \(T\) = time period.

### Non-periodic signal

A non-periodic signal does not repeat itself at regular intervals.

Examples: a typical speech waveform, a random data signal and a
non-repeating information signal.

  Periodic                        Non-periodic
  ------------------------------- ------------------------------------
  Repeats after fixed intervals   Does not repeat at fixed intervals
  Has a defined period            No fixed repetition period
  Example: Sine wave              Example: Random data waveform

## 1.5.4 Characteristics of a Signal

Four important signal characteristics:

1.  Amplitude.
2.  Frequency.
3.  Phase.
4.  Wavelength.

### A. Amplitude

Amplitude is the maximum magnitude or displacement of a signal from its
reference level.

For a voltage signal, it can be measured in volts.

Example: If a signal varies between +5 V and -5 V, its peak amplitude is
5 V.

### B. Frequency

Frequency is the number of complete cycles of a periodic signal per
second. It is measured in Hertz (Hz).

\[ f=`\frac{1}{T}`{=tex} \]

Where:

-   \(f\) = frequency.
-   \(T\) = time period.

Example: If a signal completes 100 cycles per second:

\[ f=100`\text{ Hz}`{=tex} \]

### C. Time Period

Time period is the time required to complete one full cycle.

\[ T=`\frac{1}{f}`{=tex} \]

Example: If frequency is 50 Hz:

\[
T=`\frac{1}{50}`{=tex}=0.02`\text{ seconds}`{=tex}=20`\text{ ms}`{=tex}
\]

### D. Phase

Phase describes the position of a periodic signal relative to a
reference point. It is generally measured in degrees or radians.

One complete cycle:

\[ 360\^`\circ=2`{=tex}`\pi`{=tex}`\text{ radians}`{=tex} \]

Two signals can have the same frequency and amplitude but different
phases.

### E. Wavelength

Wavelength is the physical distance covered by one complete cycle of a
wave. It is represented by (`\lambda`{=tex}).

\[ `\lambda`{=tex}=`\frac{v}{f}`{=tex} \]

Where:

-   (`\lambda`{=tex}) = wavelength in meters.
-   \(v\) = propagation velocity in m/s.
-   \(f\) = frequency in Hz.

Example:

\[ v=3`\times10`{=tex}\^8`\text{ m/s}`{=tex} \]

\[ f=100`\text{ MHz}`{=tex}=100`\times10`{=tex}\^6`\text{ Hz}`{=tex} \]

\[ `\lambda`{=tex}=`\frac{3\times10^8}{100\times10^6}`{=tex} \]

\[ `\boxed{\lambda=3\text{ meters}}`{=tex} \]

## 1.5.5 Bit Rate vs Baud Rate

These terms are frequently confused.

### Bit rate

Bit rate is the number of bits transmitted per second.

Unit: bits per second (bps).

### Baud rate

Baud rate is the number of signal symbols transmitted per second.

Unit: baud.

A symbol can represent one or multiple bits.

If every symbol represents the same number of bits:

\[
`\text{Bit Rate}`{=tex}=`\text{Baud Rate}`{=tex}`\times`{=tex}`\text{Bits per Symbol}`{=tex}
\]

Example 1: A system transmits 1 bit per symbol at 1000 baud.

\[ 1000`\times1`{=tex}=1000`\text{ bps}`{=tex} \]

Example 2: A system transmits 2 bits per symbol at 1000 baud.

\[ 1000`\times2`{=tex}=2000`\text{ bps}`{=tex} \]

Example 3: A system transmits 4 bits per symbol at 2400 baud.

\[ 2400`\times4`{=tex}=9600`\text{ bps}`{=tex} \]

  Bit rate                      Baud rate
  ----------------------------- --------------------------------
  Bits transmitted per second   Symbols transmitted per second
  Measured in bps               Measured in baud
  Depends on bits per symbol    Depends on symbol rate

**Remember:** Bit rate and baud rate are equal only when one symbol
represents exactly one bit.

## 1.5.6 Bandwidth and Frequency

These terms have different meanings depending on the context.

### Frequency

Frequency tells us how many cycles occur per second. Unit: Hertz.

### Bandwidth

In signal communication, bandwidth is the difference between the highest
and lowest frequencies in a frequency band.

\[ B=f_H-f_L \]

Example: A communication channel supports frequencies from 100 MHz to
110 MHz.

\[ B=110-100 \]

\[ `\boxed{B=10\text{ MHz}}`{=tex} \]

A channel's available frequency bandwidth affects how much information
it can carry, together with signal quality and modulation.

In networking advertisements, bandwidth is also commonly used to mean
the data rate of a link, such as 100 Mbps.

### Important distinction

-   Frequency: Hz.
-   Frequency bandwidth: Hz.
-   Digital data rate: bps.

## 1.5.7 Noise and Attenuation

### Noise

Noise refers to unwanted disturbances that affect a transmitted signal.

It can interfere with the receiver's ability to identify the original
information.

**Types of noise:**

1.  **Thermal noise:** Generated by the random thermal motion of
    electrons in electronic components.
2.  **Induced noise:** Caused by electromagnetic interference from
    external sources. Example: Electrical machinery affecting nearby
    communication cables.
3.  **Crosstalk:** Unwanted signal coupling from one communication
    channel into another. Example: Signals in one copper cable pair
    interfering with a neighboring pair.
4.  **Impulse noise:** Sudden, high-amplitude disturbances. Example:
    Electrical switching or lightning-related interference.

### Attenuation

Attenuation is the reduction in signal strength as the signal travels
through a transmission medium.

Example: A signal transmitted through a long copper cable may become
weaker as distance increases.

Reasons:

-   Cable resistance.
-   Energy loss.
-   Frequency-dependent absorption.
-   Connector losses.

### Signal amplification and regeneration

Repeaters or amplifiers can help extend communication distance.

However, amplification also amplifies noise, whereas a digital
regenerator attempts to reconstruct the original digital symbols.

## 1.5.8 Signal-to-Noise Ratio (SNR)

SNR measures the ratio of signal power to noise power.

\[ SNR=`\frac{P_s}{P_n}`{=tex} \]

Where:

-   (P_s) = Signal power.
-   (P_n) = Noise power.

A higher SNR generally means a cleaner signal.

SNR is often expressed in decibels (dB).

\[
SNR\_{dB}=10`\log`{=tex}\_{10}`\left`{=tex}(`\frac{P_s}{P_n}`{=tex}`\right`{=tex})
\]

### Numerical example

Suppose:

-   Signal power = 100 mW.
-   Noise power = 1 mW.

\[ SNR=`\frac{100}{1}`{=tex}=100 \]

Convert into decibels:

\[ SNR\_{dB}=10`\log`{=tex}\_{10}(100) \]

\[ =10`\times2`{=tex} \]

\[ `\boxed{20\text{ dB}}`{=tex} \]

### Interpretation

  SNR        General interpretation
  ---------- ------------------------------------
  High       Signal is strong relative to noise
  Low        Noise has a greater effect
  Very low   Communication becomes difficult

## 1.5.9 Nyquist Bit Rate Theorem

The Nyquist theorem gives the maximum theoretical bit rate for a
noiseless band-limited channel when using a specified number of discrete
signal levels.

Formula:

\[ C=2B`\log`{=tex}\_2 L \]

Where:

-   \(C\) = Maximum bit rate in bps.
-   \(B\) = Channel bandwidth in Hz.
-   \(L\) = Number of discrete signal levels.

### Numerical example

Suppose:

-   Bandwidth = 3000 Hz.
-   Number of signal levels = 2.

\[ C=2`\times3000`{=tex}`\times`{=tex}`\log`{=tex}\_2 2 \]

Since:

\[ `\log`{=tex}\_2 2=1 \]

Therefore:

\[ C=6000`\text{ bps}`{=tex} \]

Now suppose there are 4 signal levels.

\[ C=2`\times3000`{=tex}`\times`{=tex}`\log`{=tex}\_2 4 \]

\[ =6000`\times2`{=tex} \]

\[ `\boxed{C=12000\text{ bps}}`{=tex} \]

### Key observation

Increasing the number of signal levels can increase the theoretical bit
rate for a fixed bandwidth.

However, distinguishing more signal levels becomes harder in the
presence of noise.

## 1.5.10 Shannon Capacity Theorem

The Shannon theorem gives the theoretical maximum data rate of a
communication channel with a specified bandwidth and signal-to-noise
ratio, under the additive white Gaussian noise model.

Formula:

\[ C=B`\log`{=tex}\_2(1+SNR) \]

Where:

-   \(C\) = Channel capacity in bps.
-   \(B\) = Bandwidth in Hz.
-   SNR = Linear signal-to-noise ratio.

**Important:** Use SNR as a linear ratio, not directly as a decibel
value.

### Numerical example

Suppose:

-   Bandwidth = 3000 Hz.
-   SNR = 30 dB.

First, convert SNR into a linear ratio.

\[ SNR=10\^{30/10} \]

\[ SNR=1000 \]

Now apply Shannon's formula:

\[ C=3000`\log`{=tex}\_2(1+1000) \]

\[ C=3000`\log`{=tex}\_2(1001) \]

\[ C`\approx3000`{=tex}`\times9.967`{=tex} \]

\[ `\boxed{C\approx29901\text{ bps}}`{=tex} \]

### Nyquist vs Shannon

  -----------------------------------------------------------------------
  Nyquist                             Shannon
  ----------------------------------- -----------------------------------
  Assumes noiseless channel           Considers channel noise

  Depends on bandwidth and signal     Depends on bandwidth and SNR
  levels                              

  (2B`\log`{=tex}\_2L)                (B`\log`{=tex}\_2(1+SNR))

  Determines a limit under its        Gives a theoretical noisy-channel
  noiseless model                     capacity
  -----------------------------------------------------------------------

**Exam tip:** The Nyquist formula and Shannon formula describe different
constraints. Do not treat them as interchangeable.

## 1.5.11 Encoding and Modulation Basics

Computers generate digital information, but physical communication
systems must represent that information using signals.

Two broad concepts are encoding and modulation.

### Encoding

Encoding maps data into signal patterns suitable for transmission.

Examples:

-   NRZ.
-   Manchester encoding.
-   Differential Manchester.
-   4B/5B.
-   8b/10b.

**NRZ (Non-Return-to-Zero):** Signal levels represent bits without
necessarily returning to a zero reference between bits.

**Manchester encoding:** Each bit period contains a transition in the
middle. The direction of transition represents the bit value under the
selected convention. It also helps with clock synchronization.

### Modulation

Modulation changes a carrier signal according to the information being
transmitted.

The three classical digital modulation methods are:

**1. ASK -- Amplitude Shift Keying**

The amplitude of the carrier changes according to the digital data.

**2. FSK -- Frequency Shift Keying**

The carrier frequency changes according to the digital data.

**3. PSK -- Phase Shift Keying**

The phase of the carrier changes according to the digital data.

  Method   Parameter changed
  -------- -------------------
  ASK      Amplitude
  FSK      Frequency
  PSK      Phase

Modern communication systems also use QAM (Quadrature Amplitude
Modulation), which varies combinations of amplitude and phase.

For example, Wi-Fi uses sophisticated modulation and coding schemes to
adapt transmission rates to channel conditions.

## 1.5.12 Multiplexing

Multiplexing is a technique that allows multiple signals or data streams
to share a common communication medium.

A device called a multiplexer combines multiple input signals into a
shared channel. At the receiving end, a demultiplexer separates the
signals.

Imagine a highway. Instead of constructing a separate highway for every
vehicle, multiple vehicles share the same road using an organized
system. Multiplexing follows a similar principle.

### A. FDM -- Frequency Division Multiplexing

In FDM, the available frequency spectrum is divided into separate
frequency bands. Each signal uses a different frequency band.

Example: Traditional radio broadcasting. Different stations transmit on
different carrier frequencies.

### B. TDM -- Time Division Multiplexing

In TDM, multiple users share a communication channel by transmitting
during different time slots.

Example: A communication system assigns separate time slots to different
data sources.

Types:

-   Synchronous TDM.
-   Statistical TDM.

### C. WDM -- Wavelength Division Multiplexing

WDM is used in optical fiber communication.

Different data streams are transmitted using different wavelengths of
light through the same fiber.

Example: Telecommunication companies use WDM to carry multiple
high-capacity communication channels over optical fiber.

### D. CDM -- Code Division Multiplexing

CDM allows multiple users to transmit over the same frequency band at
the same time using different codes.

Each user is assigned a code that helps the receiver distinguish its
signal.

Example: CDMA cellular systems use code-based separation.

### Multiplexing comparison

  Technique   Separation method             Example
  ----------- ----------------------------- -------------------------------
  FDM         Different frequencies         Radio broadcasting
  TDM         Different time slots          Digital communication systems
  WDM         Different light wavelengths   Optical fiber
  CDM         Different codes               CDMA systems

------------------------------------------------------------------------

# 1.6 Transmission Media

Transmission media are the physical paths or wireless channels through
which data travels from one device to another.

They are divided into two major categories:

1.  Guided transmission media.
2.  Unguided transmission media.

## 1.6.1 Guided Transmission Media

In guided media, signals travel through a physical medium such as a
cable.

Three important types:

-   Twisted pair cable.
-   Coaxial cable.
-   Optical fiber.

### A. Twisted Pair Cable

Twisted pair cable contains two insulated copper wires twisted around
each other.

### Why are the wires twisted?

Twisting helps reduce electromagnetic interference and crosstalk.

The two wires are twisted in a way that helps induced interference
affect both conductors similarly, allowing differential signaling to
reject much of the common interference.

### Types

**1. UTP -- Unshielded Twisted Pair**

-   No additional metallic shielding.
-   Relatively inexpensive.
-   Easy to install.
-   Common in Ethernet LANs.

**2. STP -- Shielded Twisted Pair**

-   Includes shielding.
-   Offers improved protection against external interference when
    properly installed and grounded.
-   Usually more expensive.

### Categories

  Category   Typical applications
  ---------- -----------------------------------------------------------
  Cat5e      Gigabit Ethernet
  Cat6       Gigabit Ethernet and some 10 Gigabit applications
  Cat6a      10 Gigabit Ethernet up to 100 m
  Cat8       High-speed data-center connections over shorter distances

Actual performance depends on the cable standard, installation quality,
channel length and networking equipment.

### Advantages

-   Affordable.
-   Easy installation.
-   Flexible.
-   Widely available.

### Disadvantages

-   More susceptible to interference than fiber.
-   Distance limitations.
-   Lower capacity than many fiber systems.

**Real-life example:** The Ethernet cable connecting your desktop
computer to your home or college switch.

### B. Coaxial Cable

A coaxial cable consists of:

1.  Central copper conductor.
2.  Insulating dielectric layer.
3.  Metallic shield.
4.  Outer protective jacket.

The central conductor carries the signal, while the surrounding shield
helps reduce interference.

### Advantages

-   Good shielding.
-   Better interference resistance than basic unshielded cable.
-   Can support relatively high-frequency signals.
-   Suitable for broadband distribution.

### Disadvantages

-   Bulkier than twisted pair.
-   More difficult to install.
-   Less flexible in some applications.

**Real-life example:** Cable television and cable broadband Internet
connections.

### C. Optical Fiber Cable

Optical fiber transmits information using light.

It consists mainly of:

1.  Core.
2.  Cladding.
3.  Protective coating.

### How does it work?

The transmitter converts digital information into optical signals.

The light travels through the fiber core.

The receiver detects the optical signal and converts it back into
electrical/digital information.

### Total Internal Reflection

Optical fiber uses total internal reflection to guide light through the
core.

The core has a higher refractive index than the cladding.

When the light meets the core-cladding boundary at an appropriate angle,
it remains guided within the core.

### Types of optical fiber

**1. Single-mode fiber (SMF)**

-   Very small core.
-   Supports propagation primarily in one spatial mode.
-   Suitable for long-distance communication.
-   High bandwidth-distance capability.
-   Common in telecom backbone networks.

**2. Multimode fiber (MMF)**

-   Larger core.
-   Supports multiple propagation modes.
-   Commonly used for shorter-distance communication.
-   Often used inside buildings and data centers.

### Single-mode vs Multimode

  -----------------------------------------------------------------------
  Feature                 Single-mode             Multimode
  ----------------------- ----------------------- -----------------------
  Core size               Smaller                 Larger

  Light propagation       One spatial mode        Multiple modes

  Dispersion              Lower modal dispersion  Modal dispersion
                                                  present

  Distance                Longer                  Generally shorter

  Applications            Telecom backbone        Data centers, buildings

  Equipment               Typically laser-based   LED or VCSEL-based,
                                                  depending on system
  -----------------------------------------------------------------------

### Advantages of optical fiber

-   Very high capacity.
-   Low attenuation over long distances.
-   Immune to electromagnetic interference.
-   Difficult to tap without detection in many implementations.
-   Suitable for long-distance communication.

### Disadvantages

-   Installation can be expensive.
-   Fragile compared with some copper cables.
-   Requires specialized splicing and testing equipment.
-   Optical transceivers can be expensive.

**Real-life examples:**

-   Undersea Internet cables.
-   ISP backbone networks.
-   Fiber-to-the-home (FTTH).
-   Data-center interconnections.

## 1.6.2 Unguided Transmission Media

Unguided media transmit signals through free space rather than through a
physical cable.

They commonly use electromagnetic waves.

Major types:

1.  Radio waves.
2.  Microwaves.
3.  Infrared.
4.  Satellite communication.

### A. Radio Waves

Radio waves are electromagnetic waves used for wireless communication.

**Characteristics:**

-   Can travel through free space.
-   Some frequency bands can penetrate walls.
-   Can support omnidirectional communication.
-   Performance depends on frequency, power, antennas and environment.

Applications:

-   Wi-Fi.
-   Bluetooth.
-   FM radio.
-   Cellular communication.

**Real-life example:** Your smartphone communicates with a nearby
cellular base station using radio waves.

### B. Microwaves

Microwaves are electromagnetic waves commonly used for high-frequency
wireless communication.

**Characteristics:**

-   Often use directional antennas.
-   Terrestrial links commonly require line of sight.
-   Can provide high data rates.
-   Can be affected by obstacles and atmospheric conditions.

Applications:

-   Point-to-point wireless links.
-   Cellular backhaul.
-   Satellite communication.

**Real-life example:** Two telecom towers communicating through a
directional microwave link.

### C. Infrared

Infrared communication uses electromagnetic radiation at infrared
frequencies.

**Characteristics:**

-   Often short-range.
-   Commonly requires a direct or reflected optical path.
-   Does not generally pass through opaque walls.
-   Can be affected by obstacles.

Applications:

-   Television remote controls.
-   Infrared sensors.
-   Short-range device communication.

### D. Satellite Communication

Satellite communication uses satellites to relay signals between
geographically separated locations.

### Working

1.  An earth station transmits a signal to a satellite.
2.  This transmission is called the uplink.
3.  The satellite receives and retransmits the signal.
4.  The retransmission toward Earth is called the downlink.
5.  The receiving earth station gets the signal.

### Types of communication satellites

**1. GEO -- Geostationary Earth Orbit**

-   Approximately 35,786 km above the equator.
-   Appears stationary relative to Earth.
-   Large coverage area.
-   High propagation delay.

**2. MEO -- Medium Earth Orbit**

-   Located between LEO and GEO.
-   Used in navigation and some communication systems.

**3. LEO -- Low Earth Orbit**

-   Much closer to Earth.
-   Lower propagation delay than GEO.
-   Requires many satellites for continuous wide-area coverage.

### Advantages

-   Large geographical coverage.
-   Useful in remote areas.
-   Supports broadcasting.
-   Can connect geographically isolated locations.

### Disadvantages

-   Propagation delay, particularly with GEO.
-   Weather-related signal degradation in some frequency bands.
-   Infrastructure and launch costs.
-   Limited spectrum resources.

## 1.6.3 Guided vs Unguided Media

  -----------------------------------------------------------------------
  Feature                 Guided                  Unguided
  ----------------------- ----------------------- -----------------------
  Medium                  Physical cable          Free space

  Signal path             Cable-guided            Wireless

  Examples                Twisted pair, fiber     Radio, microwave

  Interference            Depends on medium and   Can be affected by
                          shielding               external signals

  Mobility                Limited by cable        Supports mobility

  Security                Physical access         Signals may be
                          required for many       intercepted within
                          attacks                 coverage
  -----------------------------------------------------------------------

## 1.6.4 Simplex, Half-Duplex and Full-Duplex

These terms describe the direction in which communication can take
place.

### A. Simplex

Simplex communication allows data transmission in only one direction.

Example: Traditional television broadcasting. The broadcasting station
sends information to the television receiver. The receiver does not send
data back through that same one-way broadcast channel.

### B. Half-Duplex

Half-duplex communication allows transmission in both directions, but
not simultaneously.

Example: Walkie-talkies. One person talks while the other listens. They
take turns.

### C. Full-Duplex

Full-duplex communication allows both sides to transmit simultaneously.

Example: A telephone conversation. Both people can speak and listen
simultaneously.

### Comparison

  Feature                     Simplex                 Half-duplex     Full-duplex
  --------------------------- ----------------------- --------------- -------------
  Direction                   One-way                 Both ways       Both ways
  Simultaneous transmission   No                      No              Yes
  Example                     Broadcast TV            Walkie-talkie   Telephone
  Efficiency                  Application-dependent   Moderate        High

**Important networking distinction:** Traditional shared Ethernet was
half-duplex, while modern switched Ethernet commonly supports
full-duplex links.

------------------------------------------------------------------------

# 1.7 Formula Sheet

Keep this section separately for revision. These formulas are
particularly useful for university examinations.

  Concept               Formula
  --------------------- -----------------------------------------------------
  Frequency             (f=`\frac{1}{T}`{=tex})
  Time period           (T=`\frac{1}{f}`{=tex})
  Wavelength            (`\lambda`{=tex}=`\frac{v}{f}`{=tex})
  Bit rate              (`\frac{\text{Number of bits}}{\text{Time}}`{=tex})
  Bit rate and baud     (R_b=R_s`\log`{=tex}\_2 L)
  Frequency bandwidth   (B=f_H-f_L)
  SNR                   (SNR=`\frac{P_s}{P_n}`{=tex})
  SNR in dB             (10`\log`{=tex}\_{10}(SNR))
  Nyquist               (C=2B`\log`{=tex}\_2 L)
  Shannon               (C=B`\log`{=tex}\_2(1+SNR))
  Transmission delay    (d\_{trans}=`\frac{L}{R}`{=tex})
  Propagation delay     (d\_{prop}=`\frac{d}{s}`{=tex})
  Mesh links            (`\frac{n(n-1)}{2}`{=tex})

Where:

-   (R_b) = Bit rate.
-   (R_s) = Symbol rate.
-   \(L\) = Number of signal levels in the Nyquist and baud formulas.
-   \(B\) = Bandwidth in Hz.
-   \(d\) = Distance.
-   \(s\) = Propagation speed.
-   \(L\) in the transmission-delay formula = Packet size in bits.
-   \(R\) = Link transmission rate in bps.

Notice that the letter (L) is used in different formulas for different
quantities. Always identify its meaning before calculating.

------------------------------------------------------------------------

# 1.8 Textbook References

Primary reference:

**Computer Networking: A Top-Down Approach, 7th Edition**\
James F. Kurose and Keith W. Ross.

The textbook's Chapter 1 is particularly relevant:

  Topic                               Textbook section     Printed pages
  ----------------------------------- ------------------ ---------------
  Internet and network fundamentals   1.1                           2--8
  Network edge and access networks    1.2                          9--20
  Network core                        1.3                         21--34
  Delay, loss, throughput             1.4                         35--46
  Protocol layers and encapsulation   1.5                         47--54
  Network security                    1.6                         55--58

The OSI model is discussed in the textbook's Chapter 1, Section 1.5. The
book focuses primarily on the Internet's five-layer architecture. The
seven-layer OSI syllabus requires additional detail.

Similarly, analog/digital signal theory, Nyquist and Shannon theorems,
and multiplexing require supplementary physical-layer material.

### Further textbook reading

-   **Chapter 1, Section 1.1:** Internet definition, end systems,
    communication links, packet switches, protocols and network
    services.
-   **Section 1.2 -- The Network Edge:** End systems, access networks
    and physical media.
-   **Section 1.3 -- The Network Core:** Packet switching, circuit
    switching and networks of networks.
-   **Section 1.4 -- Delay, Loss and Throughput:** Processing delay,
    queuing delay, transmission delay, propagation delay and throughput.
-   **Section 1.5 -- Protocol Layers and Service Models:** Layered
    architecture, service models and encapsulation.
-   **Chapter 2, Section 2.1:** Client-server and P2P application
    architectures.
-   **Chapter 4:** Routing and forwarding.
-   **Chapter 6:** Link-layer services, Ethernet, switches and LANs.
-   **Chapter 7:** Wireless and mobile networks.

### Practical learning

Connect networking concepts to a Go HTTP server. Use Wireshark to
inspect:

-   TCP handshakes.
-   HTTP requests.
-   IP addresses.
-   Port numbers.
-   Ethernet frames.

Practical packet inspection helps connect the theoretical protocol
layers to actual network traffic.
