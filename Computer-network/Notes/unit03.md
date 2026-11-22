# **Unit 3 — Network Layer, IP Addressing, Routing and Network-on-Chip**

This unit is about how data travels between different networks. Unit 2 focused on moving frames across a local link. Unit 3 moves one level higher: it explains how packets travel from a source network to a destination network, how routers choose paths, how IP addresses are divided into subnets, and how routing protocols exchange network information.

The last section, **Network-on-Chip (NoC)**, is a separate hardware topic. It applies routing concepts inside a processor or System-on-Chip rather than across the Internet.

## **Unit 3 — Complete topic checklist**

Use this list to track the syllabus while studying.

**Study progress**

# **0/66**

**3.1 Network Layer Fundamentals**

0 / 10 completed

Responsibilities of Network Layer

Logical addressing

Routing and forwarding

Packet delivery

Connectionless vs connection-oriented services

Datagram networks

Virtual circuit networks

Routing table

Next-hop forwarding

Longest Prefix Match (LPM)

**3.2 IP Addressing**

0 / 12 completed

IPv4 structure, Network ID and Host ID

Classful addressing: A, B, C, D and E

Private and public IP addresses

Loopback, APIPA and special IP addresses

IPv4 packet header and TTL

Fragmentation, reassembly and MTU

IPv6 structure and notation

IPv6 address types

IPv6 packet header

IPv4 vs IPv6

IPv6 transition mechanisms

Link-local and global unicast addresses

**3.3 Subnetting and CIDR**

0 / 9 completed

Subnet masks and CIDR notation

Network and broadcast addresses

First and last usable IP

Number of hosts and subnets

Classful vs classless addressing

VLSM

Supernetting and route aggregation

Subnetting numerical problems

Longest-prefix-match problems

**3.4 Supporting Network Protocols**

0 / 6 completed

ARP and RARP

ICMP and ICMP message types

Ping and Traceroute / Tracert

DHCP

NAT and PAT

Private vs public IP communication

**3.5 Routing Fundamentals**

0 / 7 completed

Static and dynamic routing

Routing metrics and hop count

Administrative distance

Convergence and routing loops

Count-to-infinity problem

Split horizon and route poisoning

Default route

**3.6 Routing Algorithms**

0 / 7 completed

Distance Vector and Bellman-Ford

Distance vector routing table and updates

Count-to-infinity, split horizon and poison reverse

Link State routing

Dijkstra's shortest path algorithm

LSA, LSDB and SPF

Link State vs Distance Vector

**3.7 Routing Protocols**

0 / 6 completed

RIP

OSPF and OSPF areas

BGP and path selection basics

IGP vs EGP

Autonomous Systems

RIP vs OSPF vs BGP

**3.8 Network-on-Chip and SoC Communication**

0 / 9 completed

NoC architecture, routers, processing elements and network interfaces

NoC topologies

XY deterministic and adaptive routing

Static vs dynamic routing

Packet vs circuit switching

Wormhole and store-and-forward switching

Virtual channels

Deadlock and livelock

Latency, throughput and power

# **3.1 Network Layer Fundamentals**

## **1. What is the Network Layer?**

The Network Layer is Layer 3 of the OSI model. It is responsible for delivering packets from a source host to a destination host across one or more interconnected networks.

Think of a situation where your laptop is connected to your home Wi-Fi and you access a server in another country.

Your laptop and the server are not connected to the same physical network. The packet must travel through multiple routers.

The Network Layer provides the logical addressing and routing mechanisms needed to make this communication possible.

**Real-world example: Opening a website**

**A packet travelling across networks**

**Your laptop**

`192.168.1.10`

Home LAN

**Home router**

Default gateway

Internet Service Provider (ISP)

**ISP routers**

Select routes toward the destination

**Destination network router**

**Web server**

Destination IP address

The packet can pass through many routers. Each router examines the destination IP address and decides where the packet should go next.

## **2. Responsibilities of the Network Layer**

### **A. Logical addressing**

The Network Layer uses logical addresses, mainly IP addresses, to identify source and destination hosts or interfaces.

Example:

```
Source IP:      192.168.1.10
Destination IP: 142.250.183.14
```

An IP address is hierarchical: its prefix identifies a network or subnet, while the remaining bits identify an interface within that subnet.

Unlike a MAC address, an IP address is used to support communication across multiple networks.

### **B. Routing**

Routing is the process of determining which path packets should take through an internetwork.

A router uses information from its routing table to select a route.

Example:

```
             Router A
             /      \
            /        \
        Router B    Router C
            \        /
             \      /
           Destination
```

There may be multiple possible paths to the destination. Routing protocols or static configuration help determine which routes are available and preferred.

### **C. Forwarding**

Forwarding is the actual operation of moving a received packet from an incoming interface to an appropriate outgoing interface.

For example:

1. Router receives an IP packet on Interface 1.
2. It reads the destination IP address.
3. It looks up the destination in its forwarding table.
4. It selects an outgoing interface and next hop.
5. It sends the packet onward.

**Routing vs forwarding:**

- Routing determines and maintains available paths.
- Forwarding applies the selected path to each arriving packet.

Routing is commonly considered part of the control plane; packet forwarding is part of the data plane.

### **D. Packet delivery**

Packet delivery means transferring an IP packet from source to destination.

Two basic cases are important:

**Direct delivery:** The destination is on the same local network. The sender can deliver the packet directly over the link.

**Indirect delivery:** The destination is on another network. The sender forwards the packet to its default gateway, which routes it onward.

Example:

```
Laptop A                    Laptop B
192.168.1.10                192.168.1.20
       \                      /
        \                    /
          Same local subnet
```

Laptop A can communicate directly with Laptop B at the link layer after resolving its MAC address.

Now consider:

```
Laptop A                    Remote Server
192.168.1.10                142.250.183.14
       |                          |
    Router  ---------------------+
       |
   Internet
```

Laptop A sends the frame to the router's MAC address. The IP packet is addressed to the remote server.

### **E. Internetworking**

The Network Layer allows different physical networks and link technologies to communicate through a common network-layer protocol such as IP.

For example, an IP packet may travel through:

- Ethernet.
- Wi-Fi.
- Fibre-optic ISP infrastructure.
- Another Ethernet network.

The same IP packet can traverse these different link technologies, although link-layer frames are rebuilt at each routed hop.

## **3. Connectionless vs Connection-Oriented Network Services**

There are two broad ways of providing network-layer service.

### **A. Connectionless service**

In a connectionless network, the sender does not establish a dedicated connection through the network before sending each packet.

Each packet is treated as an independent datagram.

IP uses a connectionless datagram model.

For example:

```
Packet 1: A → B → D
Packet 2: A → C → D
Packet 3: A → B → D
```

Packets belonging to the same communication can potentially follow different paths.

**Characteristics:**

- No network-layer connection setup is required.
- Each packet carries destination addressing information.
- Routers make forwarding decisions for each packet.
- Different packets may experience different delays or paths.
- Delivery is not guaranteed by IP itself.

### **B. Connection-oriented service**

In a connection-oriented network service, a logical connection is established before data transfer.

Network devices maintain state associated with the connection.

The network may assign a connection identifier or virtual-circuit identifier, rather than making every forwarding decision solely from the full destination address.

Example:

```
Sender
  |
  | Connection setup
  v
Router A → Router B → Router C
  |
  | Connection established
  v
Data transfer follows the established path
```

### **Comparison**

| Feature         | Connectionless                                           | Connection-oriented                      |
| --------------- | -------------------------------------------------------- | ---------------------------------------- |
| Setup           | Not required at the network layer                        | Required                                 |
| Packet identity | Destination address                                      | Can use virtual-circuit identifier       |
| Path            | Can vary packet by packet                                | Usually follows established logical path |
| Router state    | No per-flow connection state required by the basic model | Maintains connection/VC state            |
| Example         | IP datagram network                                      | ATM virtual-circuit network              |

**Important distinction:** TCP is connection-oriented at the Transport Layer, but IP underneath it is connectionless. Do not confuse TCP connections with network-layer virtual circuits.

## **4. Datagram Networks**

A datagram network is a connectionless packet-switched network.

Each packet is called a datagram.

It carries the destination IP address, allowing each router to make an independent forwarding decision.

### **Example**

Suppose Host A wants to send three packets to Host D.

```
                   +------ Router B ------+
                  /                        \
Host A ----------                            ---------- Host D
                  \                        /
                   +------ Router C ------+
```

Possible delivery:

- Packet 1: A → B → D
- Packet 2: A → C → D
- Packet 3: A → B → D

The packets may arrive out of order. The receiving Transport Layer can handle ordering when a reliable protocol such as TCP is used.

### **Real-world application**

The Internet Protocol (IP) is the standard example of a datagram-based network-layer service.

## **5. Virtual Circuit Networks**

A virtual circuit (VC) is a logical connection established through a packet-switched network.

Before data transmission, the network establishes a route and assigns a virtual-circuit identifier.

Routers maintain forwarding state for that virtual circuit.

### **Working**

A virtual circuit generally has three phases:

1. Connection setup.
2. Data transfer.
3. Connection release.

```
Source
  |
  v
Router A
  |
  v
Router B
  |
  v
Router C
  |
  v
Destination

A logical path is established before data transfer.
```

Every packet belonging to the virtual circuit follows the established logical path, subject to the network's connection-management behaviour.

### **Example: ATM networks**

Asynchronous Transfer Mode (ATM) used virtual paths and virtual channels identified by fields such as VPI and VCI.

Virtual-circuit networking is useful to understand because it illustrates how packet switching can still maintain connection-oriented logical paths.

## **6. Routing Table**

A routing table is a data structure used by a router to decide where to forward packets.

It contains information such as:

- Destination network or prefix.
- Subnet mask or prefix length.
- Next-hop address.
- Outgoing interface.
- Route metric or preference.
- Route source.

### **Example routing table**

Imagine Router A has this simplified table:

| Destination prefix | Next hop           | Interface |
| ------------------ | ------------------ | --------- |
| `192.168.1.0/24`   | Directly connected | G0/0      |
| `10.0.0.0/8`       | `192.168.1.2`      | G0/0      |
| `172.16.0.0/16`    | `192.168.1.3`      | G0/1      |
| `0.0.0.0/0`        | `192.168.1.1`      | G0/2      |

The last entry is the default route.

### **What happens when a packet arrives?**

Suppose the destination IP is:

```
10.20.30.40
```

Router A looks for a matching route.

`10.20.30.40` belongs to `10.0.0.0/8`.

Therefore, the packet is forwarded to next hop `192.168.1.2` through G0/0.

## **7. Next-Hop Forwarding**

A next hop is the next router or network-layer destination to which a packet should be sent.

Routers usually do not need to know the complete path to the destination. They need to know where to send the packet next.

### **Example**

```
Laptop
  |
  v
Router A
  |
  | Next hop = Router B
  v
Router B
  |
  | Next hop = Router C
  v
Router C
  |
  v
Destination
```

Router A sends the packet to Router B.

Router B independently examines the packet and chooses Router C as its next hop.

This process continues until the packet reaches the destination network.

## **8. Longest Prefix Match (LPM)**

Very important for interviews

Longest Prefix Match is the rule used by IP routers to select the most specific matching route.

A routing table may contain multiple prefixes that match a destination address.

The router selects the matching prefix with the greatest number of leading bits.

### **Example routing table**

| Destination prefix | Next hop |
| ------------------ | -------- |
| `10.0.0.0/8`       | Router A |
| `10.10.0.0/16`     | Router B |
| `10.10.10.0/24`    | Router C |
| `0.0.0.0/0`        | Router D |

Destination:

```
10.10.10.25
```

This destination matches all four entries.

But the most specific matching route is:

```
10.10.10.0/24
```

Therefore, Router C is selected.

### **Visual understanding**

```
10.0.0.0/8
  └── 10.10.0.0/16
        └── 10.10.10.0/24
```

The `/24` prefix contains more matching network bits than `/16` or `/8`.

### **Interview point**

LPM does not simply select the route with the lowest hop count or the numerically smallest next-hop address. It first selects the most specific matching prefix. Route preference and metrics are used when choosing between candidate routes for the same prefix according to the router's routing rules.

CIDR's forwarding model explicitly uses longest-match forwarding.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



# **3.2 IP Addressing**

IP addressing gives devices logical addresses that allow communication across interconnected networks.

There are two major versions:

- IPv4 — Internet Protocol version 4.
- IPv6 — Internet Protocol version 6.

## **Part A — IPv4**

## **1. IPv4 Address Structure**

IPv4 uses a 32-bit address.

It is normally written in dotted-decimal notation, with four octets separated by dots.

Example:

```
192.168.1.10
```

Each octet contains 8 bits.

\\[ 4\times8=32\text{ bits} \\]

Each octet can represent values from 0 to 255.

Example conversion:

```
192       168       1         10
11000000.10101000.00000001.00001010
```

**IPv4 is a 32-bit address**

# **192**

8 bits

# **168**

8 bits

# **1**

8 bits

# **10**

8 bits

Four octets, each containing eight bits.

## **2. Network ID and Host ID**

An IPv4 address can be understood as having two logical parts:

- Network ID: identifies the network.
- Host ID: identifies an interface within that network.

The subnet mask or prefix length tells us where the network portion ends and the host portion begins.

Example:

```
IP Address:  192.168.1.10/24
Mask:        255.255.255.0
```

Binary representation:

```
IP:
11000000.10101000.00000001.00001010

Mask:
11111111.11111111.11111111.00000000
```

The first 24 bits represent the network prefix. The final 8 bits are available for host addressing.

```
192.168.1 | 10
 Network  | Host
```

This is the basis of subnetting.

## **3. Classful Addressing**

Originally, IPv4 used a classful addressing system.

Addresses were divided into fixed classes based on their leading bits.

| Class | First-octet range | Default mask          | Purpose               |
| ----- | ----------------- | --------------------- | --------------------- |
| A     | 1–126\*           | `255.0.0.0` (/8)      | Large networks        |
| B     | 128–191           | `255.255.0.0` (/16)   | Medium networks       |
| C     | 192–223           | `255.255.255.0` (/24) | Smaller networks      |
| D     | 224–239           | Not applicable        | Multicast             |
| E     | 240–255           | Not applicable        | Reserved/experimental |

*The historical Class A range has special exclusions; `0.0.0.0/8` is reserved for special purposes and `127.0.0.0/8` is loopback.*

### **Class A**

First bit is `0`.

```
0 | Network ID (7 bits) | Host ID (24 bits)
```

Default mask:

```
255.0.0.0
```

Example:

```
10.1.2.3
```

### **Class B**

The first two bits are `10`.

```
10 | Network ID (14 bits) | Host ID (16 bits)
```

Default mask:

```
255.255.0.0
```

Example:

```
172.20.5.10
```

### **Class C**

The first three bits are `110`.

```
110 | Network ID (21 bits) | Host ID (8 bits)
```

Default mask:

```
255.255.255.0
```

Example:

```
192.168.1.20
```

### **Class D**

Range:

```
224.0.0.0 – 239.255.255.255
```

Class D is used for multicast addressing, not ordinary classful unicast network/host assignment.

### **Class E**

Range:

```
240.0.0.0 – 255.255.255.255
```

Historically reserved for experimental or future purposes, with some special-use exceptions such as the limited broadcast address.

### **Why classful addressing was replaced**

Classful addressing allocated networks in fixed sizes. This wasted address space for organizations whose requirements did not match the predefined classes.

For example, an organization needing 500 host addresses could not receive an exactly sized Class C network of 254 usable hosts. It needed a much larger allocation.

CIDR introduced flexible prefix lengths, allowing addresses to be allocated and routed more efficiently.

## **4. Private and Public IP Addresses**

### **Private IP addresses**

Private IPv4 addresses are intended for use inside private networks. They are not globally routable across the public Internet.

The three RFC 1918 private ranges are:

| Range                           | CIDR             | Total addresses |
| ------------------------------- | ---------------- | --------------- |
| `10.0.0.0 – 10.255.255.255`     | `10.0.0.0/8`     | 16,777,216      |
| `172.16.0.0 – 172.31.255.255`   | `172.16.0.0/12`  | 1,048,576       |
| `192.168.0.0 – 192.168.255.255` | `192.168.0.0/16` | 65,536          |

Examples:

```
192.168.1.10
10.0.0.5
172.16.2.20
```

A home Wi-Fi router usually gives devices private addresses.

### **Public IP addresses**

A public IP address is globally unique and can be routed on the public Internet, subject to routing and filtering policies.

Example:

```
Client private IP:  192.168.1.10
Router public IP:   203.0.113.10  (documentation example)
```

When the laptop accesses an external server, the home router may translate the private source address to a public address using NAT.

The RFC 1918 ranges and IPv4 special-use blocks are maintained in IANA/RFC registries.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.iana.org\&sz=32)

IANA



## **5. Loopback Address**

The IPv4 loopback range is:

```
127.0.0.0/8
```

The most familiar loopback address is:

```
127.0.0.1
```

It refers to the local machine itself.

### **Example**

When a developer runs a local web server:

```
http://127.0.0.1:8080
```

The browser communicates with a server process on the same computer. The traffic does not need to travel over the physical network.

**Real-world use:** Local backend testing, database development, web development and debugging.

## **6. APIPA Address**

APIPA stands for Automatic Private IP Addressing. The standard IPv4 link-local range is:

```
169.254.0.0/16
```

A device may automatically select an address in this range when it cannot obtain normal configuration through DHCP.

Example:

```
169.254.25.120
```

It can communicate with suitable devices on the same physical or logical link, but the address is not intended for routing across networks.

### **Real-world scenario**

You connect your laptop to a router, but the router's DHCP service is unavailable.

Your laptop may assign itself an address beginning with:

```
169.254.x.x
```

This is a useful troubleshooting clue: the machine did not obtain its expected DHCP configuration.

RFC 3927 defines IPv4 link-local addressing and states that these packets must not be forwarded by routers.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **7. Special IPv4 Addresses**

| Address / prefix  | Purpose                                         |
| ----------------- | ----------------------------------------------- |
| `0.0.0.0`         | Unspecified source address in relevant contexts |
| `0.0.0.0/0`       | Default route                                   |
| `127.0.0.1`       | Loopback address                                |
| `127.0.0.0/8`     | Loopback block                                  |
| `169.254.0.0/16`  | Link-local addressing                           |
| `255.255.255.255` | Limited broadcast                               |
| `10.0.0.0/8`      | Private address range                           |
| `172.16.0.0/12`   | Private address range                           |
| `192.168.0.0/16`  | Private address range                           |
| `224.0.0.0/4`     | Multicast address space                         |

A few important distinctions:

- `0.0.0.0` as an address is not the same thing as `0.0.0.0/0` as a default route prefix.
- `255.255.255.255` is a limited broadcast address, restricted to the local network.
- `127.0.0.1` is not a remote server. It is the local host.
- Not every address outside the RFC 1918 ranges is necessarily available for public host assignment; IANA reserves many additional special-purpose blocks.&#x20;

  [image](https://www.google.com/s2/favicons?domain=https://www.iana.org\&sz=32)

  IANA



## **8. IPv4 Packet Header**

The IPv4 header contains the information needed to deliver an IP datagram.

Its minimum size is 20 bytes. With options, it can extend up to 60 bytes.

**IPv4 header — conceptual structure**

**Version**

4 bits

**IHL**

4 bits

**DSCP / ECN**

8 bits

**Total length**

16 bits

**Identification**

16 bits

**Flags**

3 bits

**Fragment offset**

13 bits

**TTL**

8 bits

**Protocol**

8 bits

**Header checksum**

16 bits

**Source IP address — 32 bits**

**Destination IP address — 32 bits**

**Options + Padding (optional)**

**Payload / upper-layer data**

Conceptual field grouping. The IPv4 header is transmitted in 32-bit rows as specified by the protocol format.

### **Important fields**

| Field               | Size     | Purpose                                            |
| ------------------- | -------- | -------------------------------------------------- |
| Version             | 4 bits   | Identifies IPv4 (value 4)                          |
| IHL                 | 4 bits   | Header length in 32-bit words                      |
| DSCP / ECN          | 8 bits   | Traffic classification and congestion notification |
| Total Length        | 16 bits  | Entire IP datagram length, including header        |
| Identification      | 16 bits  | Identifies fragments of a datagram                 |
| Flags               | 3 bits   | Fragmentation controls                             |
| Fragment Offset     | 13 bits  | Position of fragment data                          |
| TTL                 | 8 bits   | Limits packet lifetime in hops                     |
| Protocol            | 8 bits   | Identifies upper-layer protocol, e.g. TCP or UDP   |
| Header Checksum     | 16 bits  | Checks IPv4 header corruption                      |
| Source Address      | 32 bits  | Sender IPv4 address                                |
| Destination Address | 32 bits  | Receiver IPv4 address                              |
| Options             | Variable | Optional control information                       |

The IPv4 format and fragmentation fields are specified in RFC 791.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **9. TTL — Time to Live**

TTL is an 8-bit field in the IPv4 header.

It prevents packets from circulating indefinitely through routing loops.

Each router that forwards the packet decreases TTL by at least one.

When TTL reaches zero, the router discards the packet and normally sends an ICMP Time Exceeded message to the source.

### **Example**

Suppose a packet is sent with:

```
TTL = 4
```

It travels through routers:

```
Source
  |
  | TTL = 4
  v
Router A → TTL = 3
  |
  v
Router B → TTL = 2
  |
  v
Router C → TTL = 1
  |
  v
Next router discards it when TTL expires
```

In practice, if a router receives a packet with TTL 1 and would reduce it to zero, that router discards it.

### **Real-world use: Traceroute**

Traceroute deliberately sends probes with increasing TTL values.

- Probe 1: TTL = 1, reveals the first router.
- Probe 2: TTL = 2, reveals the second router.
- Probe 3: TTL = 3, reveals the third router.

This reveals the path toward the destination.

TTL is therefore both a loop-prevention mechanism and a useful diagnostic mechanism.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.cisco.com\&sz=32)

Cisco+1



## **10. Fragmentation and Reassembly**

Fragmentation occurs when an IP datagram is larger than the maximum packet size supported by a link along its route.

That limit is called the MTU.

In IPv4, routers may fragment a packet when necessary, unless the Don't Fragment (DF) flag prohibits fragmentation.

### **Example**

Suppose:

- IP datagram size = 4000 bytes.
- IPv4 header size = 20 bytes.
- Outgoing link MTU = 1500 bytes.

The maximum fragment data size must not exceed:

\\[ 1500-20=1480\text{ bytes} \\]

The payload is:

\\[ 4000-20=3980\text{ bytes} \\]

Since fragment offsets are measured in units of 8 bytes, intermediate fragment payloads must be multiples of 8 bytes.

For this example:

| Fragment | Payload bytes | Total length | Fragment offset |
| -------- | ------------- | ------------ | --------------- |
| 1        | 1480          | 1500         | 0               |
| 2        | 1480          | 1500         | 185             |
| 3        | 1020          | 1040         | 370             |

Why are the offsets 185 and 370?

\\[ 1480/8=185 \\]

\\[ 2960/8=370 \\]

The destination uses the identification, fragment offset and More Fragments (MF) flag to reassemble the original datagram.

### **Fragmentation diagram**

```
Original IP Datagram
+--------------------------------------------------+
| Header (20 B) | Payload (3980 B)                 |
+--------------------------------------------------+
                         |
                         v
                  Fragmentation
                         |
           +-------------+-------------+
           |             |             |
           v             v             v
       Fragment 1    Fragment 2    Fragment 3
       1500 bytes    1500 bytes    1040 bytes
           |             |             |
           +-------------+-------------+
                         |
                         v
                    Destination
                         |
                    Reassembly
                         |
                         v
                Original IP Datagram
```

**Important:** Fragmentation may reduce efficiency because every fragment carries its own IP header. Losing a fragment can prevent successful reassembly of the entire datagram.

## **11. MTU — Maximum Transmission Unit**

MTU is the maximum size of the network-layer packet that a link can carry in one frame without fragmentation.

A common Ethernet IP MTU is:

```
1500 bytes
```

This is the IP packet size, not the total Ethernet frame size.

### **Example**

Suppose:

```
IP packet = 1600 bytes
Link MTU  = 1500 bytes
```

The packet is too large to pass over that link as one IPv4 packet.

Depending on the DF bit and network conditions, the packet may be fragmented or dropped.

### **Path MTU**

A path can include links with different MTUs.

The Path MTU is the smallest MTU along the path between two endpoints.

```
Host A
  |
 MTU 1500
  |
Router A
  |
 MTU 1400
  |
Router B
  |
 MTU 1500
  |
Host B

Path MTU = 1400 bytes
```

Path MTU Discovery techniques attempt to identify the maximum packet size that can cross the route without fragmentation.

## **Part B — IPv6**

## **12. Why IPv6 Was Introduced**

IPv4 uses 32-bit addresses.

The theoretical address space is:

\\[ 2^{32}=4,294,967,296 \\]

That is approximately 4.29 billion possible addresses.

As the Internet expanded, IPv4 address availability became a major constraint.

IPv6 was designed with a 128-bit address space:

\\[ 2^{128} \\]

This provides a vastly larger number of possible addresses.

IPv6 also simplifies aspects of the IP header and supports features such as stateless address autoconfiguration and multicast addressing.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://datatracker.ietf.org\&sz=32)

datatracker.ietf.org+1



## **13. IPv6 Address Structure**

IPv6 uses a 128-bit address.

It is represented as eight groups of 16 bits, each written in hexadecimal.

Example:

```
2001:0db8:0000:0000:0000:ff00:0042:8329
```

Each group is called a hextet.

**IPv6 address = 128 bits**

## **2001**

16 bits

## **0db8**

16 bits

## **0000**

16 bits

## **0000**

16 bits

## **0000**

16 bits

## **ff00**

16 bits

## **0042**

16 bits

## **8329**

16 bits

Eight hextets × 16 bits = 128 bits.

## **14. IPv6 Notation Rules**

IPv6 supports compression to make addresses shorter and easier to read.

### **Rule 1: Remove leading zeros within a hextet**

```
2001:0db8:0000:0000:0000:ff00:0042:8329
```

Becomes:

```
2001:db8:0:0:0:ff00:42:8329
```

### **Rule 2: Replace one continuous sequence of zero hextets with&#x20;**`::`

```
2001:db8:0:0:0:ff00:42:8329
```

Becomes:

```
2001:db8::ff00:42:8329
```

### **Important rule**

The `::` abbreviation can be used only once in an address. Otherwise, the number of omitted zero groups would be ambiguous.

### **Another example**

Full address:

```
2001:0db8:0000:0000:0000:0000:0000:0001
```

Compressed:

```
2001:db8::1
```

## **15. IPv6 Address Types**

IPv6 defines three major address types:

1. Unicast.
2. Multicast.
3. Anycast.

Unlike IPv4, IPv6 does not use broadcast addresses. Multicast is used for many functions that would otherwise require broadcast.

### **A. Unicast**

Identifies one interface.

A packet sent to a unicast address is delivered to the interface associated with that address.

### **B. Multicast**

Identifies a group of interfaces.

A packet sent to a multicast address is delivered to members of that multicast group.

The IPv6 multicast range begins with:

```
FF00::/8
```

### **C. Anycast**

An anycast address identifies a set of interfaces.

A packet sent to an anycast address is delivered to one member of the set, typically the nearest according to the routing system's route-selection behaviour.

### **Comparison**

| Type      | Delivery                                          |
| --------- | ------------------------------------------------- |
| Unicast   | One interface                                     |
| Multicast | A group of interfaces                             |
| Anycast   | One interface from a group, typically the nearest |
| Broadcast | Not defined in IPv6                               |

## **16. Important IPv6 Address Ranges**

| Address / prefix | Meaning                                          |
| ---------------- | ------------------------------------------------ |
| `::`             | Unspecified address                              |
| `::1`            | Loopback address                                 |
| `FE80::/10`      | Link-local unicast                               |
| `FC00::/7`       | Unique Local Address range                       |
| `2000::/3`       | Currently allocated global unicast address block |
| `FF00::/8`       | Multicast                                        |

The IANA IPv6 address registry identifies the global unicast, unique-local, link-local and multicast ranges.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.iana.org\&sz=32)

IANA+1



## **17. Link-Local vs Global Unicast IPv6 Addresses**

### **Link-local address**

A link-local address is used for communication within the same local link.

Example:

```
FE80::1
```

Link-local addresses are important for IPv6 neighbour discovery and other local-link operations.

Routers do not forward packets with link-local source or destination addresses to other links.

### **Global unicast address**

A global unicast address is designed for globally routable communication.

Example:

```
2001:db8:1234:1::10
```

`2001:db8::/32` is reserved for documentation examples, so this specific example is not a real globally routable address. A real assigned global address generally comes from the allocated global-unicast space.

### **Difference**

| Feature           | Link-local                                  | Global unicast                           |
| ----------------- | ------------------------------------------- | ---------------------------------------- |
| Scope             | One local link                              | Global routing scope                     |
| Prefix            | `FE80::/10`                                 | Currently allocated from `2000::/3`      |
| Router forwarding | Not forwarded beyond the link               | Can be routed                            |
| Typical use       | Neighbour discovery and local communication | Internet or routed-network communication |

## **18. IPv6 Packet Header**

IPv6 uses a fixed 40-byte base header.

**IPv6 base header**

**Version**

4 bits

**Traffic Class**

8 bits

**Flow Label**

20 bits

**Payload Length**

16 bits

**Next Header**

8 bits

**Hop Limit**

8 bits

**Source Address — 128 bits**

**Destination Address — 128 bits**

**Extension headers (optional)**

**Upper-layer payload**

### **Important IPv6 fields**

| Field               | Size     | Purpose                                                      |
| ------------------- | -------- | ------------------------------------------------------------ |
| Version             | 4 bits   | Identifies IPv6 (value 6)                                    |
| Traffic Class       | 8 bits   | Traffic classification and congestion indication             |
| Flow Label          | 20 bits  | Identifies packets belonging to a flow                       |
| Payload Length      | 16 bits  | Length after the base IPv6 header                            |
| Next Header         | 8 bits   | Identifies the next extension header or upper-layer protocol |
| Hop Limit           | 8 bits   | Limits packet lifetime across routers                        |
| Source Address      | 128 bits | Source IPv6 address                                          |
| Destination Address | 128 bits | Destination IPv6 address                                     |

### **Why is there no header checksum?**

IPv6 does not include a checksum in the base header.

This avoids requiring routers to recalculate a header checksum at every hop.

### **What happened to TTL?**

IPv6 uses a field called **Hop Limit** instead of TTL.

Each forwarding router decreases the Hop Limit by one. If it reaches zero, the packet is discarded.

### **Extension headers**

IPv6 uses extension headers to carry optional information, such as:

- Fragmentation information.
- Routing information.
- Authentication information.
- Destination options.

RFC 8200 defines the IPv6 packet format and extension-header structure.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



## **19. IPv4 vs IPv6**

| Feature            | IPv4                            | IPv6                                                |
| ------------------ | ------------------------------- | --------------------------------------------------- |
| Address size       | 32 bits                         | 128 bits                                            |
| Address notation   | Dotted decimal                  | Colon-separated hexadecimal                         |
| Header size        | 20–60 bytes                     | Fixed 40-byte base header                           |
| Header checksum    | Present                         | Absent                                              |
| Fragmentation      | Routers and source can fragment | Only source nodes fragment                          |
| Broadcast          | Supported                       | Not supported; multicast used instead               |
| Address resolution | ARP for Ethernet                | Neighbour Discovery using ICMPv6                    |
| TTL equivalent     | TTL                             | Hop Limit                                           |
| Configuration      | Manual or DHCP                  | Manual, SLAAC, DHCPv6                               |
| NAT requirement    | Common in IPv4 deployments      | Not inherently required by the address architecture |

### **Fragmentation difference**

This is particularly important.

**IPv4:** A router may fragment a packet if necessary, unless fragmentation is prohibited.

**IPv6:** Routers do not fragment packets. If fragmentation is needed, the source uses an IPv6 Fragment extension header.

If a packet is too large for a link, the router sends an ICMPv6 Packet Too Big message so the source can adjust its packet size or fragment it when appropriate.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



## **20. IPv6 Transition Mechanisms**

IPv4 and IPv6 do not directly interoperate because they use different address formats and packet structures.

Transition mechanisms help networks operate during migration.

### **A. Dual Stack**

A device runs IPv4 and IPv6 simultaneously.

```
             Laptop
          /           \
     IPv4 stack     IPv6 stack
          \           /
           Network interface
```

If the destination supports IPv6 and the system selects IPv6, it uses IPv6. Otherwise, IPv4 may be used.

**Example:** A server has both an IPv4 address and an IPv6 address.

### **B. Tunnelling**

Tunnelling carries packets of one network protocol inside packets of another protocol.

Example:

```
IPv6 network                 IPv6 network
     |                            |
     |        IPv4 network        |
     +---- Router A ===== Router B+
          IPv6 packet encapsulated
          inside IPv4 packets
```

The IPv6 packet is encapsulated inside IPv4 so it can cross an IPv4-only section of the network.

### **C. Translation**

Translation allows communication between IPv4-only and IPv6-only systems using a gateway that converts between the protocol versions.

One approach is NAT64, often used together with DNS64.

### **Comparison**

| Method      | Main idea                             |
| ----------- | ------------------------------------- |
| Dual Stack  | Run both IPv4 and IPv6                |
| Tunnelling  | Carry one IP version inside the other |
| Translation | Convert packets between IPv4 and IPv6 |

**Study takeaway:** Dual stack is often the easiest conceptual transition mechanism to understand; tunnelling handles incompatible network segments; translation supports communication between endpoints using different IP versions.

# **3.3 Subnetting and CIDR**

Very important for interviews and numerical exams

Subnetting is one of the most practical topics in computer networking. It is used by network engineers, cloud engineers, system administrators and backend infrastructure teams.

It helps you understand how IP addresses are divided, how routers identify networks, and how organisations allocate IP addresses efficiently.

## **1. What is subnetting?**

Subnetting is the process of dividing a larger IP network into smaller logical networks called **subnets**.

Suppose an organisation has this network:

```
192.168.1.0/24
```

It has 256 total IPv4 addresses.

Instead of keeping all devices in one large subnet, the organisation can divide it into smaller subnets:

```
Original network: 192.168.1.0/24
               |
       +-------+-------+
       |               |
    Subnet 1         Subnet 2
    /26              /26
       |               |
    Subnet 3         Subnet 4
    /26              /26
```

This helps with:

- Separating departments.
- Controlling broadcast domains.
- Organising IP address allocation.
- Applying security policies.
- Improving network administration.
- Reducing unnecessary broadcast traffic within each subnet.

A subnet is a logical IP network. It is not necessarily a separate physical switch.

## **2. Subnet Mask**

A subnet mask identifies which bits of an IPv4 address represent the network prefix and which bits represent the host portion.

Example:

```
IP address:  192.168.1.10
Subnet mask: 255.255.255.0
CIDR:        /24
```

Binary:

```
IP:
11000000.10101000.00000001.00001010

Mask:
11111111.11111111.11111111.00000000
```

- `1` bits in the subnet mask represent the network portion.
- `0` bits represent the host portion.

### **Common subnet masks**

| CIDR  | Subnet mask       | Host bits | Total addresses | Traditional usable hosts |
| ----- | ----------------- | --------- | --------------- | ------------------------ |
| `/8`  | `255.0.0.0`       | 24        | 16,777,216      | 16,777,214               |
| `/16` | `255.255.0.0`     | 16        | 65,536          | 65,534                   |
| `/24` | `255.255.255.0`   | 8         | 256             | 254                      |
| `/25` | `255.255.255.128` | 7         | 128             | 126                      |
| `/26` | `255.255.255.192` | 6         | 64              | 62                       |
| `/27` | `255.255.255.224` | 5         | 32              | 30                       |
| `/28` | `255.255.255.240` | 4         | 16              | 14                       |
| `/29` | `255.255.255.248` | 3         | 8               | 6                        |
| `/30` | `255.255.255.252` | 2         | 4               | 2                        |

The traditional usable-host formula assumes a normal IPv4 subnet where the network and broadcast addresses are reserved:

\\[ \boxed{\text{Usable hosts}=2^h-2} \\]

Here, \\(h\\) is the number of host bits.

There are special cases such as `/31` point-to-point links and `/32` host routes, where this traditional formula does not apply in the ordinary way.

## **3. CIDR Notation**

CIDR stands for **Classless Inter-Domain Routing**.

Instead of using fixed classful network sizes, CIDR represents the network prefix with a slash followed by the number of prefix bits.

Example:

```
192.168.1.0/24
```

This means:

- IPv4 address length = 32 bits.
- Network prefix = 24 bits.
- Host portion = 8 bits.

Therefore:

\\[ h=32-24=8 \\]

Total addresses:

\\[ 2^8=256 \\]

Traditional usable host addresses:

\\[ 2^8-2=254 \\]

### **More examples**

| CIDR  | Network bits | Host bits |
| ----- | ------------ | --------- |
| `/8`  | 8            | 24        |
| `/16` | 16           | 16        |
| `/24` | 24           | 8         |
| `/26` | 26           | 6         |
| `/30` | 30           | 2         |
| `/32` | 32           | 0         |

A larger prefix length means a smaller address block.

For example:

- `/24` contains 256 addresses.
- `/25` contains 128 addresses.
- `/26` contains 64 addresses.
- `/27` contains 32 addresses.

CIDR and prefix-based route aggregation replaced the limitations of classful Internet routing.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **4. Network Address**

The network address identifies the subnet itself.

It is calculated by performing a bitwise AND between the IP address and subnet mask.

\\[ \boxed{\text{Network address}=\text{IP address AND subnet mask}} \\]

Example:

```
IP:   192.168.1.130
Mask: 255.255.255.192
```

The last octet in binary:

```
IP:       10000010  (130)
Mask:     11000000  (192)
          --------
AND:      10000000  (128)
```

Therefore:

```
Network address = 192.168.1.128
```

## **5. Broadcast Address**

The broadcast address is the address used to reach all hosts on an IPv4 subnet.

It is calculated by setting all host bits to `1`.

For `192.168.1.130/26`:

```
Network address = 192.168.1.128
Host bits       = 6
```

All six host bits set to 1:

```
11 1111 = 63
```

Therefore:

```
Broadcast address = 192.168.1.191
```

A subnet's broadcast address is not normally assignable to an individual host.

## **6. First and Last Usable IP**

In a conventional IPv4 subnet:

\\[ \text{First usable IP}=\text{Network address}+1 \\]

\\[ \text{Last usable IP}=\text{Broadcast address}-1 \\]

For:

```
192.168.1.130/26
```

We calculated:

```
Network:   192.168.1.128
Broadcast: 192.168.1.191
```

Therefore:

```
First usable: 192.168.1.129
Last usable:  192.168.1.190
```

## **7. How to Solve Subnetting Questions**

There are two useful methods.

### **Method A — Binary method**

Convert the address and subnet mask into binary, then perform a bitwise AND.

This method gives the most direct mathematical understanding of subnetting.

### **Method B — Block-size method**

This method is faster for many examination and interview questions.

The formula is:

\\[ \boxed{\text{Block size}=256-\text{interesting mask octet}} \\]

The interesting octet is the first non-255, non-zero octet of the subnet mask.

Example:

```
/26 = 255.255.255.192
```

Block size:

\\[ 256-192=64 \\]

Therefore, subnet boundaries in the fourth octet are:

```
0, 64, 128, 192
```

A fourth-octet value of 130 belongs to the block starting at 128.

This gives:

```
Network:   192.168.1.128
Broadcast: 192.168.1.191
```

## **Numerical Problem 1 — Divide a /24 Network into Four Equal Subnets**

Solved example

Given:

```
192.168.1.0/24
```

Divide it into four equal-sized subnets.

### **Step 1: Find how many bits to borrow**

We need four subnets.

\\[ 2^s=4 \\]

\\[ s=2 \\]

Borrow two bits from the host portion.

Original prefix:

```
/24
```

New prefix:

\\[ 24+2=\boxed{/26} \\]

### **Step 2: Calculate the new subnet mask**

```
/26 = 255.255.255.192
```

### **Step 3: Calculate the block size**

\\[ 256-192=64 \\]

So the subnets start at:

```
0, 64, 128, 192
```

### **Step 4: Write all four subnets**

| Subnet | Network address    | First usable | Last usable | Broadcast |
| ------ | ------------------ | ------------ | ----------- | --------- |
| 1      | `192.168.1.0/26`   | `.1`         | `.62`       | `.63`     |
| 2      | `192.168.1.64/26`  | `.65`        | `.126`      | `.127`    |
| 3      | `192.168.1.128/26` | `.129`       | `.190`      | `.191`    |
| 4      | `192.168.1.192/26` | `.193`       | `.254`      | `.255`    |

**Visual subnet division**

Original network: 192.168.1.0/24

**1**

.0–.63

**2**

.64–.127

**3**

.128–.191

**4**

.192–.255

Each subnet has 64 total addresses and 62 traditional usable host addresses.

Final answer:

- Number of subnets = 4.
- Prefix = `/26`.
- Subnet mask = `255.255.255.192`.
- Total addresses per subnet = 64.
- Usable hosts per subnet = 62.

## **Numerical Problem 2 — Find Network and Broadcast Address**

Given:

```
IP = 192.168.1.130/26
```

Find:

1. Network address.
2. Broadcast address.
3. First usable IP.
4. Last usable IP.
5. Number of usable hosts.

### **Step 1: Convert the prefix into a mask**

```
/26 = 255.255.255.192
```

### **Step 2: Find block size**

\\[ 256-192=64 \\]

Subnet ranges:

| Block | Address range |
| ----- | ------------- |
| 1     | `.0–.63`      |
| 2     | `.64–.127`    |
| 3     | `.128–.191`   |
| 4     | `.192–.255`   |

The address `.130` lies between `.128` and `.191`.

Therefore:

```
Network address   = 192.168.1.128
Broadcast address = 192.168.1.191
```

### **Step 3: Find usable range**

```
First usable = 192.168.1.129
Last usable  = 192.168.1.190
```

### **Step 4: Find usable hosts**

Host bits:

\\[ 32-26=6 \\]

Total addresses:

\\[ 2^6=64 \\]

Usable hosts:

\\[ 64-2=\boxed{62} \\]

Final answer

| Property     | Result             |
| ------------ | ------------------ |
| Network      | `192.168.1.128/26` |
| Broadcast    | `192.168.1.191`    |
| First usable | `192.168.1.129`    |
| Last usable  | `192.168.1.190`    |
| Usable hosts | `62`               |

## **Numerical Problem 3 — Calculate Usable Hosts for /27**

Given:

```
192.168.10.0/27
```

### **Step 1: Find host bits**

\\[ h=32-27=5 \\]

### **Step 2: Calculate total addresses**

\\[ 2^5=32 \\]

### **Step 3: Calculate usable hosts**

\\[ 2^5-2=32-2 \\]

\\[ \boxed{30\text{ usable hosts}} \\]

Other details:

```
Subnet mask:       255.255.255.224
Block size:        256 - 224 = 32
Network:           192.168.10.0
Broadcast:         192.168.10.31
First usable:      192.168.10.1
Last usable:       192.168.10.30
```

## **8. Number of Subnets**

If we borrow \\(s\\) bits from the host portion:

\\[ \boxed{\text{Number of subnets}=2^s} \\]

Example:

A `/24` network is divided into `/27` subnets.

Borrowed bits:

\\[ 27-24=3 \\]

Number of subnets:

\\[ 2^3=\boxed{8} \\]

Each `/27` subnet has:

\\[ 2^{32-27}=32 \\]

Total addresses, or 30 traditional usable host addresses.

### **Number of subnet bits vs host bits**

| Requirement                         | Formula     |
| ----------------------------------- | ----------- |
| Number of subnets                   | \\(2^s\\)   |
| Total addresses per subnet          | \\(2^h\\)   |
| Traditional usable hosts per subnet | \\(2^h-2\\) |

Here:

- \\(s\\) = borrowed subnet bits.
- \\(h\\) = remaining host bits.

## **9. Classful vs Classless Addressing**

| Feature              | Classful                         | Classless               |
| -------------------- | -------------------------------- | ----------------------- |
| Network size         | Fixed classes                    | Variable prefix lengths |
| Prefix information   | Historically inferred from class | Explicitly specified    |
| Subnet masks         | Default class masks              | Flexible subnet masks   |
| Efficient allocation | Less efficient                   | More efficient          |
| Supports VLSM        | No, not in the classful model    | Yes                     |
| Routing              | Class-based assumptions          | Prefix-based routing    |

Example:

Classful thinking treats a Class C address as a `/24` network.

CIDR can represent that network as:

```
192.168.1.0/24
```

or divide it into:

```
192.168.1.0/26
192.168.1.64/26
192.168.1.128/26
192.168.1.192/26
```

The `/26` prefixes provide more granular network allocation.

## **10. Variable Length Subnet Masking (VLSM)**

VLSM allows different subnets within a larger address block to use different prefix lengths.

This is useful when different departments need different numbers of IP addresses.

### **Real-world example: Company network**

Suppose an organisation owns:

```
192.168.10.0/24
```

It needs:

- Engineering: 50 host addresses.
- HR: 25 host addresses.
- Guest network: 10 host addresses.

We allocate the largest requirement first.

### **Step 1: Engineering needs 50 hosts**

Try `/26`:

- Host bits = 6.
- Total addresses = 64.
- Usable hosts = 62.

So Engineering receives:

```
192.168.10.0/26
```

Range:

```
Network:   192.168.10.0
Usable:    192.168.10.1 – 192.168.10.62
Broadcast: 192.168.10.63
```

### **Step 2: HR needs 25 hosts**

Try `/27`:

- Host bits = 5.
- Total addresses = 32.
- Usable hosts = 30.

Next available aligned block:

```
192.168.10.64/27
```

Range:

```
Network:   192.168.10.64
Usable:    192.168.10.65 – 192.168.10.94
Broadcast: 192.168.10.95
```

### **Step 3: Guest network needs 10 hosts**

Try `/28`:

- Host bits = 4.
- Total addresses = 16.
- Usable hosts = 14.

Next available aligned block:

```
192.168.10.96/28
```

Range:

```
Network:   192.168.10.96
Usable:    192.168.10.97 – 192.168.10.110
Broadcast: 192.168.10.111
```

### **Final VLSM allocation**

| Department  | Requirement | Assigned subnet    | Usable hosts |
| ----------- | ----------- | ------------------ | ------------ |
| Engineering | 50          | `192.168.10.0/26`  | 62           |
| HR          | 25          | `192.168.10.64/27` | 30           |
| Guest       | 10          | `192.168.10.96/28` | 14           |

```
192.168.10.0/24
|
+-- 192.168.10.0/26     Engineering
|
+-- 192.168.10.64/27    HR
|
+-- 192.168.10.96/28    Guest
|
+-- Remaining space
```

This is VLSM: the same larger network is divided into differently sized subnets.

## **11. Supernetting**

Supernetting combines multiple smaller contiguous networks into a larger address block.

It is the opposite of subnetting.

### **Example**

Suppose a router has these four routes:

```
192.168.0.0/24
192.168.1.0/24
192.168.2.0/24
192.168.3.0/24
```

They can be represented by one aggregate route:

```
192.168.0.0/22
```

Why `/22`?

Four `/24` networks contain:

\\[ 4\times256=1024 \\]

A `/22` also contains:

\\[ 2^{32-22}=1024 \\]

The first address block is correctly aligned at `192.168.0.0`, and the four `/24` networks are contiguous.

### **Visual understanding**

```
192.168.0.0/24  \
192.168.1.0/24   \
192.168.2.0/24    +---- 192.168.0.0/22
192.168.3.0/24   /
```

## **12. Route Aggregation**

Route aggregation is the process of representing multiple specific routes as a shorter summary prefix.

It reduces routing-table size and can reduce routing-update information.

Example:

Instead of advertising:

```
10.1.0.0/24
10.1.1.0/24
10.1.2.0/24
10.1.3.0/24
```

A router may advertise a suitable summary:

```
10.1.0.0/22
```

The summary is correct only if the component prefixes are contiguous, fit the aggregate boundary, and the advertising router can handle traffic for the whole summary range appropriately.

## **13. Longest Prefix Match — Subnetting and Routing Numerical**

Consider this routing table:

| Destination prefix | Next hop |
| ------------------ | -------- |
| `192.168.1.0/24`   | R1       |
| `192.168.1.128/25` | R2       |
| `192.168.1.160/27` | R3       |
| `0.0.0.0/0`        | R4       |

Find the next hop for:

```
192.168.1.170
```

Check the matching ranges:

| Prefix             | Address range      | Match? |
| ------------------ | ------------------ | ------ |
| `192.168.1.0/24`   | `.0–.255`          | Yes    |
| `192.168.1.128/25` | `.128–.255`        | Yes    |
| `192.168.1.160/27` | `.160–.191`        | Yes    |
| `0.0.0.0/0`        | All IPv4 addresses | Yes    |

All entries match, but `/27` is the most specific.

Therefore:

\\[ \boxed{\text{Next hop}=R3} \\]

### **Another example**

Destination:

```
192.168.1.200
```

The `/27` range is `.160–.191`, so it does not match.

The `/25` range `.128–.255` matches.

Therefore:

\\[ \boxed{\text{Next hop}=R2} \\]

## **14. Additional Subnetting Practice Problems**

Solve these without looking at the solutions above.

**Subnetting quiz**

0/5 answered

**Q1. Divide 172.16.0.0/24 into 8 equal subnets. What is the new prefix?**

/25

/26

/27

/28

**Q2. How many traditional usable hosts are available in a /28 subnet?**

14

16

30

32

**Q3. What is the broadcast address of 192.168.5.75/26?**

192.168.5.63

192.168.5.127

192.168.5.95

192.168.5.255

**Q4. Are 192.168.10.35/27 and 192.168.10.62/27 in the same subnet?**

Yes

No

**Q5. Which prefix is selected by LPM when /16 and /24 both match?**

/16

/24

The default route

The one with the lower next-hop IP

Check answersReset

# **3.4 Supporting Network Protocols**

IP provides addressing and packet delivery, but several supporting protocols help devices discover each other, configure their addresses, report network problems and communicate across private and public networks.

We will study:

1. ARP
2. RARP
3. ICMP
4. Ping
5. Traceroute
6. DHCP
7. NAT
8. PAT

## **1. ARP — Address Resolution Protocol**

ARP stands for Address Resolution Protocol.

It is used in IPv4 local networks to find the MAC address associated with an IPv4 address.

Remember:

- IP address = logical address.
- MAC address = link-layer address.

When a computer wants to send an Ethernet frame to another local device, it needs the destination MAC address.

But suppose it knows only the destination IP address.

ARP solves this problem.

### **Real-world example**

Your laptop has:

```
IP: 192.168.1.10
```

Another local device has:

```
IP: 192.168.1.20
```

Your laptop wants to send data to `192.168.1.20`, but does not know its MAC address.

### **ARP working**

**Step 1: Check ARP cache**

The laptop first checks its local ARP cache to see whether it already knows the IP-to-MAC mapping.

**Step 2: Send ARP Request**

If no entry exists, the laptop sends a broadcast ARP request:

```
Who has 192.168.1.20?
Tell 192.168.1.10.
```

Destination Ethernet MAC:

```
FF:FF:FF:FF:FF:FF
```

**Step 3: ARP Reply**

The device owning `192.168.1.20` sends a reply containing its MAC address.

```
192.168.1.20 is at AA:BB:CC:DD:EE:FF
```

**Step 4: Update ARP cache**

The laptop stores this mapping and can now send the Ethernet frame.

**ARP address resolution**

**Laptop A**

`192.168.1.10`

Broadcast ARP Request

Who has 192.168.1.20?

**Laptop B**

`192.168.1.20`

Unicast ARP Reply

My MAC is AA:BB:CC:DD:EE:FF

**Laptop A updates its ARP cache**

### **ARP cache**

An ARP cache stores recently learned IP-to-MAC mappings.

On Windows, run:

```
arp -a
```

This displays known ARP entries.

### **Important real-world distinction**

Suppose the destination IP is on another network.

Your laptop does **not** ARP for the remote server's MAC address.

It ARPs for the MAC address of the default gateway.

For example:

```
Laptop:       192.168.1.10
Gateway:      192.168.1.1
Remote host:  142.250.183.14
```

The laptop creates an IP packet for `142.250.183.14`, but the Ethernet frame is addressed to the router's MAC address.

The router will then forward the IP packet to the next network.

ARP is specified in RFC 826.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **2. RARP — Reverse Address Resolution Protocol**

RARP stands for Reverse Address Resolution Protocol.

It was designed to allow a device that knew its hardware address (MAC) to request its IPv4 address from a RARP server.

### **Why was it required?**

Early diskless workstations sometimes needed to discover their IP address when starting up.

The device knew its MAC address but did not know its IP configuration.

RARP allowed it to request an IP address.

### **Working**

```
Diskless Workstation              RARP Server
        |                              |
        |----- My MAC is XX:XX ------->|
        |                              |
        |<---- Your IP is 10.0.0.5 ---|
```

### **Why is RARP not commonly used today?**

RARP had limited capabilities. It supplied an IP address but did not provide the richer network configuration information required by modern hosts.

BOOTP and then DHCP became more capable alternatives.

**Exam point:** RARP translates MAC address to IP address, whereas ARP resolves an IPv4 address to MAC address.

## **3. ICMP — Internet Control Message Protocol**

ICMP is a network-layer support protocol used to report errors and provide diagnostic information about IP packet delivery.

ICMP does not make IP itself reliable.

Instead, it provides feedback about network conditions.

Example:

A router cannot deliver a packet because the destination network is unreachable.

It may send an ICMP Destination Unreachable message to the sender.

### **Common ICMP message types**

| ICMP message            | IPv4 type | Purpose                                          |
| ----------------------- | --------- | ------------------------------------------------ |
| Echo Reply              | 0         | Responds to an echo request                      |
| Destination Unreachable | 3         | Reports delivery failure                         |
| Redirect                | 5         | Informs a host about a potentially better router |
| Echo Request            | 8         | Used by ping                                     |
| Time Exceeded           | 11        | TTL expired or fragment reassembly timed out     |
| Parameter Problem       | 12        | Reports a problem with the IP header             |

These are IPv4 ICMP type numbers. ICMPv6 has its own message definitions and type numbers.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



### **Example: Destination Unreachable**

```
Laptop A
   |
   | IP packet for 10.10.10.10
   v
Router
   |
   | No route to destination
   |
   v
ICMP Destination Unreachable
   |
   v
Laptop A
```

The sender receives information that delivery failed.

### **ICMP and IP**

ICMP is carried inside IP packets.

In IPv4, the IPv4 Protocol field has value 1 for ICMP.

In IPv6, ICMPv6 is identified by Next Header value 58.

## **4. Ping Command**

Ping is a network diagnostic utility that commonly uses ICMP Echo Request and Echo Reply messages.

It helps test whether a destination responds over the network.

### **Example**

Windows:

```
ping google.com
```

Linux/macOS:

```
ping google.com
```

Simplified process:

```
Your Laptop                     Destination
     |                               |
     |----- ICMP Echo Request ------>|
     |                               |
     |<----- ICMP Echo Reply --------|
     |                               |
```

Ping output often shows:

- Whether replies were received.
- Round-trip time.
- Packet loss.
- Time-to-live in replies on some systems.

### **Sample output**

```
Reply from 142.250.183.14:
bytes=32 time=18ms TTL=116
```

Interpretation:

- Bytes: size of the ICMP payload reported.
- Time: observed round-trip time.
- TTL: remaining TTL value in the reply packet.

**Important:** A successful ping confirms that the destination responded to the echo request. It does not guarantee that every application or service on that host is available. Firewalls may also block or rate-limit ICMP.

## **5. Traceroute / Tracert**

Traceroute is a diagnostic tool that identifies the sequence of routers (hops) a probe traverses toward a destination.

Windows uses:

```
tracert google.com
```

Linux/macOS commonly use:

```
traceroute google.com
```

### **How does traceroute work?**

It uses probes with gradually increasing TTL values.

First probe:

```
TTL = 1
```

The first router reduces the TTL to zero and discards the probe. It normally returns an ICMP Time Exceeded message.

Second probe:

```
TTL = 2
```

The first router forwards it, and the second router expires it.

Third probe:

```
TTL = 3
```

The third router expires it.

This continues until the destination responds or the maximum hop count is reached.

**Traceroute discovers the path**

**Source laptop**

TTL = 1 → ICMP Time Exceeded

**Router 1**

TTL = 2 → ICMP Time Exceeded

**Router 2**

TTL = 3 → ICMP Time Exceeded

**Router 3**

**Destination**

### **Why is traceroute useful?**

It helps diagnose:

- Routing problems.
- Unexpected network paths.
- High latency at a particular hop.
- Where packet probes stop receiving responses.

A `*` in traceroute output does not necessarily mean that the router is broken. It may simply not return the expected ICMP response.

Cisco explains the TTL-increment and ICMP Time Exceeded process in its traceroute documentation.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.cisco.com\&sz=32)

Cisco+1



## **6. DHCP — Dynamic Host Configuration Protocol**

DHCP automatically provides network configuration to hosts.

Without DHCP, an administrator would have to manually configure each device's IP address, subnet mask, default gateway and other settings.

DHCP can provide:

- IP address.
- Subnet mask / prefix.
- Default gateway.
- DNS server addresses.
- Lease duration.
- Other configuration options.

### **Real-world example**

You connect your laptop to a Wi-Fi network.

The router or a separate DHCP server automatically assigns an IP address.

Example:

```
IP address:      192.168.1.105
Subnet mask:     255.255.255.0
Default gateway: 192.168.1.1
DNS server:      192.168.1.1
```

### **DHCP DORA Process**

The common DHCPv4 address-allocation sequence is known as DORA:

- Discover.
- Offer.
- Request.
- Acknowledge.

**DHCP DORA**

**DHCP Client**

No address configured yet

**1. DHCP Discover**

Client searches for DHCP servers

**DHCP Server**

**2. DHCP Offer**

Server offers an address and configuration

**3. DHCP Request**

Client requests the offered configuration

**4. DHCP ACK**

Server confirms the lease

**Client configures its network interface**

### **Explanation of DORA**

**1. Discover:** The client broadcasts a DHCP Discover message to locate available DHCP servers.

**2. Offer:** A DHCP server offers an IP address and other configuration.

**3. Request:** The client requests the offered address.

**4. Acknowledge:** The server confirms the lease, and the client can configure the address.

### **DHCP ports**

| Device      | UDP port |
| ----------- | -------- |
| DHCP Server | 67       |
| DHCP Client | 68       |

### **What is a DHCP lease?**

A lease is the period for which a DHCP server allows a client to use the assigned address.

The client can renew its lease before it expires.

### **What if the DHCP server is on another subnet?**

DHCP Discover messages are generally broadcast on the local network and are not forwarded by ordinary routers.

A router can use a DHCP relay function to forward client requests to a DHCP server on a different subnet.

DHCP is specified in RFC 2131.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **7. NAT — Network Address Translation**

NAT stands for Network Address Translation.

It modifies IP addressing information as packets pass through a translating device, commonly a border router or firewall.

A common purpose is to allow devices using private IPv4 addresses to communicate with external networks through a public IPv4 address.

### **Why is NAT used?**

Suppose a home has five devices:

| Device   | Private IP     |
| -------- | -------------- |
| Laptop   | `192.168.1.10` |
| Phone    | `192.168.1.11` |
| Tablet   | `192.168.1.12` |
| Desktop  | `192.168.1.13` |
| Smart TV | `192.168.1.14` |

The home router has a public IP:

```
203.0.113.10
```

For illustration, `203.0.113.0/24` is a documentation-only range.

The router can translate private addresses into a public address when traffic leaves the local network.

```
Laptop       192.168.1.10
Phone        192.168.1.11
Desktop      192.168.1.13
      \           |          /
       \          |         /
         +----------------+
         | Home Router    |
         | NAT            |
         +----------------+
                 |
            Public IP
          203.0.113.10
                 |
              Internet
```

### **Types of NAT**

**A. Static NAT**

A fixed one-to-one mapping between a private IP address and a public IP address.

Example:

```
192.168.1.10 ↔ 203.0.113.10
```

This mapping remains fixed.

**B. Dynamic NAT**

Private addresses are mapped to public addresses selected from a configured pool.

The mapping is created dynamically when internal devices initiate communication.

**C. PAT — Port Address Translation**

PAT translates private IP addresses and transport-layer port numbers to a public IP address and translated port numbers.

It is also commonly called NAT overload.

This allows many internal devices to share a single public IPv4 address.

## **8. PAT — Port Address Translation**

PAT is a particularly important practical form of NAT.

Imagine two devices access the same external website.

Laptop:

```
Private IP: 192.168.1.10
Source port: 51000
```

Phone:

```
Private IP: 192.168.1.11
Source port: 51000
```

Both may use the same original source port, but the router can assign different translated source ports.

Example translation table:

| Internal source      | Translated public source |
| -------------------- | ------------------------ |
| `192.168.1.10:51000` | `203.0.113.10:40001`     |
| `192.168.1.11:51000` | `203.0.113.10:40002`     |

The router uses the translated port information to direct response packets back to the correct internal device.

### **PAT diagram**

**Multiple devices sharing one public address**

**Laptop**

`192.168.1.10:51000`

**Phone**

`192.168.1.11:51000`

Router performs PAT

**Public address: 203.0.113.10**

`:40001` → Laptop

`:40002` → Phone

**Internet destination**

### **NAT vs PAT**

| Feature         | NAT                             | PAT                                                                |
| --------------- | ------------------------------- | ------------------------------------------------------------------ |
| Translation     | IP address                      | IP address plus port                                               |
| Mapping         | May be one-to-one or pool-based | Many internal sessions share public IP using different identifiers |
| Public IPv4 use | Can require more public IPs     | Can share one public IP across many sessions                       |
| Common usage    | Static and dynamic NAT          | Home/office router Internet access                                 |

Traditional NAT and NAPT behaviour is described in RFC 3022.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **9. Private vs Public IP Communication — Complete Example**

Suppose you open a website from a home laptop.

Your laptop:

```
Private IP: 192.168.1.10
```

Your home router:

```
LAN IP:     192.168.1.1
Public IP:  203.0.113.10  (example only)
```

Website server:

```
198.51.100.25 (example documentation address)
```

The process:

1. The laptop creates an IP packet whose destination is the website server.
2. It identifies that the destination is outside the local subnet.
3. It sends the Ethernet frame to its default gateway's MAC address.
4. The router receives the frame and performs routing.
5. The router applies NAT/PAT to the outgoing flow.
6. The packet travels toward the website.
7. The website responds to the translated public address and port.
8. The router checks its translation table.
9. It translates the response destination back to the internal address and port.
10. It forwards the packet to the laptop.

**Important:** NAT is not the same as a firewall. NAT changes addresses and/or ports; a firewall enforces security policies. Many home routers perform both functions, which can make them seem like one feature.

# **3.5 Routing Fundamentals**

Routing is the process of selecting paths through a network so that packets can travel from source networks to destination networks.

A router receives a packet, examines its destination IP address, consults its routing information and forwards the packet toward the next hop.

## **1. What is Routing?**

Consider this network:

```
                 Router B
                /        \
               /          \
Host A ---- Router A      Router D ---- Host D
               \          /
                \        /
                 Router C
```

Host A wants to communicate with Host D.

There are two possible paths:

```
Path 1: A → Router A → Router B → Router D → Host D

Path 2: A → Router A → Router C → Router D → Host D
```

The routing process determines which path should be used.

The router's forwarding process then sends each packet toward the selected next hop.

## **2. Static Routing**

Static routing means routes are manually configured by a network administrator.

Example:

A router is configured with:

```
Destination: 192.168.20.0/24
Next hop:    10.0.0.2
```

This tells the router:

> To reach network `192.168.20.0/24`, forward packets to `10.0.0.2`.

### **Example topology**

```
LAN A                  LAN B
192.168.1.0/24         192.168.20.0/24
    |                        |
 Router A ---------------- Router B
          10.0.0.1  10.0.0.2
```

Router A can have a static route:

```
192.168.20.0/24 via 10.0.0.2
```

### **Advantages**

- Simple to understand and configure in small networks.
- Predictable forwarding.
- No routing protocol update traffic.
- Gives the administrator precise control over route configuration.

### **Disadvantages**

- Manual changes are needed when topology changes.
- Does not automatically discover alternate paths.
- Configuration becomes difficult in large networks.
- An incorrect static route can create black holes or routing loops.

### **Real-world use**

Static routes are common for:

- Small branch networks.
- Stub networks with one Internet connection.
- Default routes toward an upstream provider.
- Special routing requirements.

## **3. Dynamic Routing**

Dynamic routing allows routers to learn routes automatically through routing protocols.

Instead of manually configuring every remote network, routers exchange routing information.

Common routing protocols include:

- RIP.
- OSPF.
- BGP.

### **Working**

```
Router A <------ Routing updates ------> Router B
    ^                                     |
    |                                     |
    +------ Routing updates ------ Router C
```

The routers exchange information about reachable networks and update their routing tables.

When the network topology changes, the routers can calculate new paths.

### **Advantages**

- Suitable for larger networks.
- Can respond to link failures.
- Automatically learns remote networks.
- Can choose alternate paths.

### **Disadvantages**

- Requires more processing and memory.
- Routing updates consume resources.
- Configuration and troubleshooting can be more complex.
- Incorrect routing information can spread across a network.

## **4. Static vs Dynamic Routing**

| Feature                | Static Routing                                     | Dynamic Routing                |
| ---------------------- | -------------------------------------------------- | ------------------------------ |
| Configuration          | Manual                                             | Learned through protocols      |
| Adaptation to failures | Requires manual changes or alternate static routes | Can adapt automatically        |
| Protocol overhead      | No routing-protocol updates                        | Uses routing-protocol messages |
| Scalability            | Limited for large changing networks                | Better for large networks      |
| Administrative control | Very explicit                                      | Depends on protocol and policy |
| Example                | Manually configured route                          | OSPF, RIP, BGP                 |

## **5. Routing Metrics**

A routing metric is a value used by a routing protocol to evaluate the desirability of a route.

A lower metric often means a more preferred path within that protocol, but the exact rule depends on the protocol.

### **Common metrics**

**A. Hop count**

The number of routers a packet must pass through.

Example:

```
Path A: Router 1 → Router 2 → Router 3
Hop count = 3

Path B: Router 1 → Router 4
Hop count = 1
```

If hop count is the metric, Path B is preferred.

RIP uses hop count.

**B. Bandwidth**

The capacity of a link to carry data.

A route through higher-bandwidth links may be preferred by a protocol that considers bandwidth.

**C. Delay**

The time required for data to cross a link or network.

Lower delay may be preferred.

**D. Reliability**

An indication of how dependable a link or path is.

**E. Load**

An indication of how heavily a link or path is being used.

**F. Cost**

A protocol-specific numerical value assigned to links or routes.

OSPF uses a cost metric, commonly derived from interface bandwidth according to implementation and configuration.

### **Example**

Consider two paths:

| Path | Hop count | Bandwidth |
| ---- | --------- | --------- |
| A    | 2         | 10 Mbps   |
| B    | 3         | 1 Gbps    |

The preferred path depends on the routing protocol's metric.

If using hop count only, Path A may win.

If using a bandwidth-based metric, Path B may be preferred.

**Important:** The path with the fewest hops is not necessarily the fastest path.

## **6. Administrative Distance**

Administrative Distance (AD) is a router-specific preference value used by some routing implementations to compare route sources when the same destination prefix is learned from multiple sources.

It is not the same as a routing protocol's metric.

### **Example**

A router learns the same destination network through:

- A static route.
- OSPF.
- RIP.

The router can compare the administrative preference of these sources to determine which route source to trust.

For Cisco IOS, common default administrative-distance examples are:

| Route source       | Common Cisco AD |
| ------------------ | --------------- |
| Directly connected | 0               |
| Static route       | 1               |
| eBGP               | 20              |
| EIGRP internal     | 90              |
| OSPF               | 110             |
| RIP                | 120             |
| iBGP               | 200             |

Lower administrative distance is preferred.

These values are Cisco conventions, not universal Internet routing rules. Other vendors may use different route-preference systems.

### **Administrative distance vs metric**

| Administrative Distance          | Metric                                   |
| -------------------------------- | ---------------------------------------- |
| Compares route sources           | Compares paths within a routing protocol |
| Local to a router implementation | Defined by the routing protocol          |
| Lower usually preferred          | Depends on the protocol                  |
| Example: OSPF vs RIP             | Example: one OSPF path vs another        |

## **7. Convergence**

Convergence is the process by which routers reach a consistent view of the routing topology and usable routes after a network change.

Suppose Router B becomes unavailable.

```
Before failure:

Router A ---- Router B ---- Router C

After failure:

Router A       Router B       Router C
   \                            /
    +----------- New path ------+
```

The routers must detect the failure, exchange updated information and calculate usable alternative routes.

The time required for this process is called convergence time.

### **Why does convergence matter?**

During convergence, packets may:

- Be dropped.
- Follow an outdated route.
- Experience temporary routing loops.
- Take an alternate path with more delay.

Faster convergence helps reduce network disruption.

## **8. Routing Loops**

A routing loop occurs when packets are forwarded repeatedly between routers instead of reaching their destination.

### **Example**

```
        Router A
        /      \
       /        \
  Router B ---- Router C
```

Suppose:

- Router A forwards a packet to Router B.
- Router B believes Router C is the correct next hop.
- Router C believes Router A is the correct next hop.

The packet can circulate:

```
A → B → C → A → B → C → ...
```

In IPv4, TTL eventually reaches zero and the packet is discarded.

In IPv6, the Hop Limit performs the equivalent role.

### **Causes of routing loops**

- Outdated routing information.
- Slow convergence.
- Incorrect static routes.
- Inconsistent routing updates.
- Distance-vector routing problems.

## **9. Count-to-Infinity Problem**

Count-to-infinity is a classic problem associated with distance-vector routing.

It occurs when routers incorrectly believe that a destination is reachable through one another after a route has failed.

### **Example**

Suppose there are three routers:

```
A -------- B -------- C
           |
       Destination
```

Assume Router B has a direct route to the destination with metric 1.

Router A reaches it through B with metric 2.

Now the link between B and the destination fails.

But A and B may temporarily have outdated information.

B hears from A that the destination is reachable with metric 2, so B thinks it can reach the destination through A with metric 3.

Then A receives this updated information and increases the metric again.

The metric keeps increasing:

```
A says: Destination = 2
B says: Destination = 3
A says: Destination = 4
B says: Destination = 5
A says: Destination = 6
...
```

This continues until the protocol reaches its defined infinity value.

In RIP, a metric of 16 represents an unreachable destination.

### **Why is this dangerous?**

- It wastes bandwidth on routing updates.
- Convergence becomes slow.
- Routers temporarily forward packets into loops.
- Network recovery is delayed.

## **10. Split Horizon**

Split horizon is a technique used in distance-vector routing to reduce routing loops.

The rule is:

**A router does not advertise a route back through the interface from which it learned that route.**

### **Example**

```
Router A -------- Router B -------- Router C
                     |
                Destination
```

Suppose B learns about the destination from C.

B should not advertise that same route back to C, because C is the router from which B learned it.

This reduces the possibility of A and B or B and C incorrectly believing that they can reach a destination through each other.

## **11. Route Poisoning**

Route poisoning is a technique in which a router advertises a failed route with an unreachable metric.

Example using RIP:

```
Before failure:

Network X → Metric 2

After failure:

Network X → Metric 16
```

The metric 16 tells other routers that the destination is unreachable.

### **Poison Reverse**

Poison reverse is a variation of split horizon.

Instead of simply omitting a route in the direction from which it was learned, a router advertises that route back with an infinite metric.

In RIP, that means advertising metric 16.

### **Difference**

| Technique       | Behaviour                                                                           |
| --------------- | ----------------------------------------------------------------------------------- |
| Split horizon   | Do not advertise the route back through the interface from which it was learned     |
| Route poisoning | Advertise the failed destination with an infinite metric                            |
| Poison reverse  | Advertise an infinite metric back to the neighbour from which the route was learned |

These techniques reduce routing-loop problems, but they do not solve every possible multi-router failure scenario.

## **12. Default Route**

A default route is used when the routing table has no more-specific matching route for a destination.

In IPv4:

```
0.0.0.0/0
```

In IPv6:

```
::/0
```

The `/0` means the prefix length is zero, so the route can match any destination.

### **Example**

A home laptop sends traffic to a remote Internet server.

Its routing table may contain:

```
192.168.1.0/24 → Local network
0.0.0.0/0      → 192.168.1.1
```

When the destination is not part of `192.168.1.0/24`, the laptop uses its default gateway.

### **Default route diagram**

```
Laptop
  |
  v
Home Router
  |
  | Default route
  v
ISP Router
  |
  v
Internet
```

### **When is it used?**

- Home routers.
- Small office networks.
- Branch routers.
- Stub networks with one upstream provider.
- Devices that send external traffic to a known gateway.

# **3.6 Routing Algorithms**

Routing algorithms are the mathematical and logical procedures used to select paths through a network.

Two major categories are:

1. Distance Vector Routing.
2. Link State Routing.

## **1. Distance Vector Routing**

Distance-vector routing is based on the idea that every router knows:

- The destinations it can reach.
- The estimated distance to each destination.
- The next hop used to reach each destination.

Routers exchange routing information with their directly connected neighbours.

They do not need to know the complete network topology.

### **Example topology**

```
             Cost 2
       A --------------- B
       |                 |
       | Cost 7          | Cost 3
       |                 |
       +------- C -------+
```

Assume:

- A to B cost = 2.
- B to C cost = 3.
- A to C cost = 7.

Router A initially knows:

| Destination | Cost | Next hop |
| ----------- | ---- | -------- |
| A           | 0    | —        |
| B           | 2    | B        |
| C           | 7    | C        |

Router A also receives a routing update from B.

B tells A:

```
Destination C = 3
```

A calculates:

\\[ \text{Cost}(A\rightarrow B\rightarrow C) =2+3=5 \\]

Compare the alternatives:

| Path      | Cost |
| --------- | ---- |
| A → C     | 7    |
| A → B → C | 5    |

Router A selects the lower-cost path through B.

Updated routing table:

| Destination | Cost | Next hop |
| ----------- | ---- | -------- |
| A           | 0    | —        |
| B           | 2    | B        |
| C           | 5    | B        |

## **2. Bellman-Ford Algorithm**

Distance-vector routing is based on the Bellman-Ford shortest-path principle.

The formula is:

\\[ \boxed{D_x(y)=\min_v\\{c(x,v)+D_v(y)\\}} \\]

Where:

- \\(D_x(y)\\) = current best-known distance from router \\(x\\) to destination \\(y\\).
- \\(c(x,v)\\) = cost of the direct link between router \\(x\\) and neighbour \\(v\\).
- \\(D_v(y)\\) = neighbour \\(v\\)'s advertised distance to destination \\(y\\).
- The minimum is taken over all neighbours \\(v\\).

### **Numerical example**

Router A has two neighbours: B and C.

Link costs:

```
A → B = 2
A → C = 5
```

B advertises its cost to destination X as 3.

C advertises its cost to destination X as 2.

Calculate the route to X.

Via B:

\\[ 2+3=5 \\]

Via C:

\\[ 5+2=7 \\]

Therefore:

\\[ D_A(X)=\min(5,7)=5 \\]

Router A chooses B as its next hop.

## **3. Distance Vector Routing Table**

A simplified distance-vector routing table looks like:

| Destination | Distance | Next hop |
| ----------- | -------- | -------- |
| Network A   | 0        | Direct   |
| Network B   | 2        | Router B |
| Network C   | 5        | Router B |
| Network D   | 3        | Router C |

The router updates these entries as it receives information from neighbours.

## **4. Routing Updates**

In distance-vector protocols, routers periodically exchange route information and may send updates after topology changes.

Example:

```
Router A
    |
    | Sends its distance vector
    v
Router B

Router B compares the received metrics
with its own known routes and updates its table.
```

The router does not receive the complete topology. It receives a summary of destinations and distances from its neighbours.

## **5. Count-to-Infinity in Distance Vector Routing**

Consider the following:

```
A -------- B -------- C
           |
           D
```

Suppose B directly reaches Network D.

Router A learns about D from B.

Now the connection from B to D fails.

If routers have outdated information, A might continue advertising a route to D through B, while B starts believing it can reach D through A.

The metrics can keep increasing.

```
A: D is at cost 2
B: D is at cost 3
A: D is at cost 4
B: D is at cost 5
...
```

This is the count-to-infinity problem.

In RIP, the unreachable metric is 16, limiting the count.

## **6. Split Horizon and Poison Reverse in Distance Vector**

### **Split horizon**

Do not advertise a route through the interface from which it was learned.

### **Poison reverse**

Advertise the route back to that neighbour with an infinite metric.

For RIP:

```
Metric 16 = unreachable
```

Both methods help prevent certain routing loops.

RIP is a distance-vector routing protocol based on Bellman-Ford principles.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **7. Link State Routing**

Link-state routing works differently from distance-vector routing.

Instead of simply learning distances from neighbours, each router builds a map of the network topology within its routing area.

It uses this information to independently calculate the shortest paths.

### **General working**

1. Discover neighbouring routers.
2. Learn link costs.
3. Create Link State Advertisements (LSAs).
4. Flood LSAs through the routing area.
5. Build a Link State Database (LSDB).
6. Run Dijkstra's Shortest Path First algorithm.
7. Build routing entries from the resulting shortest-path tree.

### **Diagram**

```
                  Router A
                 /        \
                /          \
           Router B ------ Router C
                \          /
                 \        /
                  Router D
```

Routers exchange information about their links, allowing them to build a common topology database within the area.

## **8. Link State Advertisements (LSA)**

An LSA is information describing a router's links or certain other topology and routing information.

An LSA may contain information such as:

- Router identity.
- Neighbouring routers.
- Link state.
- Link cost.
- Network information.

When network topology changes, new LSAs are generated and distributed according to the protocol.

### **Example**

Router A reports:

```
A is connected to B with cost 2.
A is connected to C with cost 5.
```

Router B reports:

```
B is connected to A with cost 2.
B is connected to C with cost 1.
```

The routers use the distributed information to build their topology database.

## **9. LSDB — Link State Database**

LSDB stands for Link State Database.

It stores the link-state information received by a router.

In OSPF, routers within the same area normally maintain a synchronized link-state database for that area, subject to area-specific database rules.

The database describes the area's topology.

The router then uses it to calculate shortest paths.

## **10. Dijkstra's Shortest Path Algorithm**

Dijkstra's algorithm calculates shortest paths from one starting vertex to all reachable vertices in a graph with non-negative edge weights.

OSPF uses a Shortest Path First calculation based on Dijkstra's algorithm.

### **Example topology**

Consider this weighted graph:

**Routing topology with link costs**

Link costs are written next to the connections. Start the shortest-path calculation at S.

We want the shortest paths from S.

Link costs:

- S → A = 2.
- S → B = 5.
- A → B = 2.
- A → C = 1.
- B → C = 3.
- C → D = 2.
- B → D = 6.

### **Step 1: Initialize distances**

```
S = 0
A = ∞
B = ∞
C = ∞
D = ∞
```

Start at S.

Direct neighbours:

```
A = 2
B = 5
```

### **Step 2: Select A**

A has the smallest tentative distance of 2.

Check paths through A:

To B:

\\[ 2+2=4 \\]

Previously B was 5, so update B to 4.

To C:

\\[ 2+1=3 \\]

Previously C was infinity, so update C to 3.

### **Step 3: Select C**

C has the smallest tentative distance of 3.

Check the path to D:

\\[ 3+2=5 \\]

Therefore:

```
D = 5
```

### **Step 4: Select B**

B has distance 4.

Check D through B:

\\[ 4+6=10 \\]

The existing cost to D is 5, so do not update.

Check C through B:

\\[ 4+3=7 \\]

The existing cost to C is 3, so do not update.

### **Step 5: Final result**

| Destination | Shortest cost | Shortest path |
| ----------- | ------------- | ------------- |
| S           | 0             | S             |
| A           | 2             | S → A         |
| B           | 4             | S → A → B     |
| C           | 3             | S → A → C     |
| D           | 5             | S → A → C → D |

Final shortest-path tree:

```
S
|
A
| \
B  C
    |
    D
```

The path to D has the cost:

\\[ 2+1+2=\boxed{5} \\]

## **11. Shortest Path First (SPF)**

SPF stands for Shortest Path First.

It is the algorithmic calculation used by a link-state router to determine the best paths from itself to known destinations.

In OSPF:

- The LSDB provides the topology information.
- Dijkstra's algorithm calculates the shortest-path tree.
- The router uses that tree to derive routing entries.

## **12. Distance Vector vs Link State**

| Feature              | Distance Vector                     | Link State                                      |
| -------------------- | ----------------------------------- | ----------------------------------------------- |
| Knowledge            | Learns distances from neighbours    | Builds a topology map                           |
| Information exchange | Routing tables or route information | Link-state information                          |
| Algorithm            | Bellman-Ford principle              | Dijkstra's SPF                                  |
| Updates              | Typically periodic and triggered    | Topology-driven flooding and refresh mechanisms |
| Convergence          | Can be slower in some protocols     | Often faster in well-designed networks          |
| Complexity           | Generally simpler                   | More computationally complex                    |
| Example              | RIP                                 | OSPF                                            |

**Interview point:** Distance Vector routers ask their neighbours what they know about destinations. Link State routers learn the topology and calculate shortest paths themselves.

# **3.7 Routing Protocols**

Routing algorithms explain how routes can be calculated. Routing protocols define how routers exchange the information needed to calculate and maintain those routes.

The three important protocols in this syllabus are:

- RIP — Routing Information Protocol.
- OSPF — Open Shortest Path First.
- BGP — Border Gateway Protocol.

Before studying them, understand the difference between IGP and EGP.

## **1. Autonomous Systems (AS)**

An Autonomous System is a collection of IP networks and routers under a common administrative routing policy.

An AS is identified by an Autonomous System Number (ASN).

For example:

```
             The Internet
                  |
        +---------+---------+
        |                   |
    AS 100                AS 200
   Company A               ISP A
        |                   |
     Router              Router
        |                   |
    Internal             Internal
    networks             networks
```

An enterprise or service provider may operate its own Autonomous System.

Routing between networks inside one AS and routing between different ASes are handled by different categories of protocols.

## **2. IGP vs EGP**

IGP stands for **Interior Gateway Protocol**.

It is used to exchange routing information within one Autonomous System.

Examples:

- RIP.
- OSPF.
- IS-IS.
- EIGRP.

EGP stands for **Exterior Gateway Protocol**.

It is used to exchange routing information between Autonomous Systems.

The most important exterior routing protocol on the public Internet is BGP.

### **Comparison**

| Feature       | IGP                                                  | EGP                                              |
| ------------- | ---------------------------------------------------- | ------------------------------------------------ |
| Routing scope | Within an AS                                         | Between ASes                                     |
| Main purpose  | Internal network routing                             | Inter-domain routing                             |
| Main concern  | Reachability and path metrics inside an organisation | Reachability, policy and AS-level path selection |
| Examples      | RIP, OSPF                                            | BGP                                              |

**Important:** IGP and EGP describe protocol roles, not OSI layers. Both operate to distribute network-layer reachability information.

## **3. RIP — Routing Information Protocol**

RIP stands for Routing Information Protocol.

RIP is a distance-vector routing protocol based on the Bellman-Ford algorithm.

It is designed for relatively small and simple IP networks.

### **Key characteristics**

- Distance-vector routing.
- Uses hop count as its routing metric.
- Maximum usable hop count is 15.
- Hop count 16 means unreachable.
- RIPv2 uses UDP port 520.
- Routers periodically exchange routing information.

RIPv2 is specified in RFC 2453.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



### **What is hop count?**

Hop count is the number of routers a packet must pass through to reach a destination.

Example:

```
Host A
  |
Router 1
  |
Router 2
  |
Router 3
  |
Destination
```

The route has three router hops.

A route with a smaller hop count is preferred by RIP, assuming other route-selection considerations are equal.

### **RIP metric example**

Suppose Router A knows two paths to Network X.

```
Path 1: A → B → X
Path 2: A → C → D → E → X
```

| Path   | Hop count |
| ------ | --------- |
| Path 1 | 2         |
| Path 2 | 4         |

RIP prefers Path 1.

### **RIP routing updates**

RIP routers exchange route information with their neighbours.

In traditional RIP, periodic updates occur approximately every 30 seconds. Triggered updates can also communicate changes more quickly.

A router receives a neighbour's distance vector and calculates the cost of reaching each destination through that neighbour.

### **RIP versions**

**RIPv1:**

- Classful.
- Does not carry subnet masks in route updates.
- Does not support classless routing information in the way RIPv2 does.

**RIPv2:**

- Classless.
- Carries subnet mask information.
- Supports CIDR and VLSM.
- Supports authentication mechanisms.

### **Limitations of RIP**

- Maximum path length is 15 hops.
- Slower convergence compared with many modern IGP deployments.
- Hop count does not account for bandwidth or delay.
- Less suitable for large, complex networks.

### **Real-world understanding**

RIP is useful for understanding how distance-vector routing works. It is less common in modern large enterprise networks, where protocols such as OSPF are often used instead.

## **4. OSPF — Open Shortest Path First**

OSPF is a link-state Interior Gateway Protocol.

It is used within an Autonomous System to distribute internal routing information and calculate shortest paths.

OSPF uses Dijkstra's Shortest Path First algorithm.

### **Main characteristics**

- Link-state routing.
- Uses a cost metric.
- Builds a Link State Database.
- Floods Link State Advertisements.
- Calculates a shortest-path tree using SPF.
- Supports hierarchical routing using areas.
- Supports CIDR and variable-length prefixes.

OSPFv2 for IPv4 is specified in RFC 2328.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



### **How OSPF works**

Consider this network:

```
             Router A
             /      \
            /        \
        Router B ---- Router C
            \          /
             \        /
              Router D
```

Each router discovers its neighbours and learns about links.

The routers exchange LSAs, and the network builds a link-state database.

Each router runs Dijkstra's algorithm to calculate its own shortest-path tree.

### **OSPF working process**

1. Routers discover neighbouring OSPF routers.
2. They establish neighbour relationships and, where required, adjacency.
3. They exchange link-state information.
4. LSAs are flooded within the relevant scope.
5. Routers build and synchronize their LSDBs.
6. Each router runs SPF.
7. Routing entries are installed based on the calculated paths.

### **What is OSPF cost?**

OSPF uses cost as its route metric.

A common Cisco-style cost calculation is:

\\[ \boxed{\text{Cost}=\frac{\text{Reference bandwidth}}{\text{Interface bandwidth}}} \\]

Actual cost behaviour depends on the implementation and configuration. Modern network links may require an appropriately configured reference bandwidth to distinguish very fast interfaces.

### **Example**

Suppose:

```
Reference bandwidth = 100 Mbps
```

Link A:

```
Bandwidth = 10 Mbps
Cost = 100/10 = 10
```

Link B:

```
Bandwidth = 100 Mbps
Cost = 100/100 = 1
```

Link B has a lower cost.

Therefore, if the competing paths have no other preference difference, OSPF favours the lower-cost path.

## **5. OSPF Areas**

One of the most important OSPF features is hierarchical routing using areas.

In a large network, maintaining a single enormous topology database everywhere can become inefficient.

OSPF areas divide the routing domain into smaller logical regions.

### **Area 0 — Backbone Area**

Area 0 is the OSPF backbone area.

Other OSPF areas connect to the backbone, either directly or through supported virtual-link arrangements.

### **Example OSPF architecture**

**OSPF multi-area architecture**

## **Area 0 — Backbone**

**R1**

**R2**

**ABR**

**Area 1**

Internal routers

R3 — R4

**Area 2**

Internal routers

R5 — R6

Conceptual layout: the Area Border Router (ABR) connects an area to the backbone and exchanges appropriate routing information between areas.

### **Important OSPF router roles**

**Internal Router:** All OSPF interfaces of the router belong to the same area.

**Backbone Router:** Has an interface in Area 0.

**Area Border Router (ABR):** Connects one or more non-backbone areas to the backbone.

**Autonomous System Boundary Router (ASBR):** Redistributes routes from another routing source into OSPF.

### **Why use areas?**

- Reduce the amount of topology information needed by routers outside a particular area.
- Improve scalability.
- Provide route summarization opportunities at area boundaries.
- Help control routing information distribution.

**Exam point:** Area 0 is the OSPF backbone. ABRs connect areas to the backbone.

## **6. BGP — Border Gateway Protocol**

BGP stands for Border Gateway Protocol.

It is the main inter-domain routing protocol used to exchange network reachability information between Autonomous Systems on the Internet.

BGP is often called a **path-vector routing protocol**.

Unlike RIP and OSPF, BGP is strongly policy-driven. It does not simply choose the path with the lowest physical distance or the fewest routers.

BGP-4 is specified in RFC 4271 and supports classless inter-domain routing.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



### **Why is BGP required?**

The Internet consists of many independently administered Autonomous Systems.

For example:

```
          AS 100
          ISP A
            |
            |
          AS 200
          ISP B
            |
            |
          AS 300
        Company C
```

AS 100 needs to exchange reachability information with AS 200.

AS 200 needs to exchange reachability information with AS 300.

BGP makes inter-AS route exchange possible.

### **How does BGP work?**

BGP speakers exchange information about reachable IP prefixes.

A route advertisement can include path attributes such as:

- AS_PATH.
- NEXT_HOP.
- LOCAL_PREF.
- MED.
- ORIGIN.

These attributes help the receiving router make policy-based path-selection decisions.

## **7. Path Vector Routing**

Path-vector routing advertises not only the destination prefix but also information about the sequence of Autonomous Systems the route traverses.

Example:

```
Destination: 203.0.113.0/24

AS_PATH:
100 → 200 → 300
```

This tells the receiving AS that the advertised route passes through AS 100, AS 200 and AS 300.

### **Why is AS_PATH important?**

It helps with:

- Loop prevention at the AS level.
- Path selection.
- Policy decisions.

For example, if AS 200 receives a route advertisement whose AS_PATH already includes AS 200, it can recognize that accepting the route would create an AS-level routing loop.

## **8. BGP Path Selection Basics**

BGP may learn multiple routes to the same destination prefix.

It must choose which path to install and use.

The decision is based on path attributes and routing policy.

Important attributes include:

### **A. LOCAL_PREF**

LOCAL_PREF expresses the preference for an exit path within an AS.

Higher LOCAL_PREF is preferred.

Example:

| Route     | LOCAL_PREF |
| --------- | ---------- |
| Via ISP A | 200        |
| Via ISP B | 100        |

All else being equal, the route through ISP A is preferred inside that AS.

### **B. AS_PATH**

AS_PATH records the sequence of AS numbers the route has traversed.

A shorter AS_PATH is often preferred, all else being equal.

However, network policy can override a shorter-path preference.

### **C. NEXT_HOP**

Identifies the next-hop IP address used to reach the advertised destination.

### **D. MED — Multi-Exit Discriminator**

MED is an attribute used to suggest a preferred entry point into an AS when multiple links connect the same neighbouring AS.

Lower MED is normally preferred when the comparison is applicable and other conditions are satisfied.

### **E. ORIGIN**

Indicates how the route was originally introduced into BGP.

It is one of the attributes considered during route selection.

### **Important qualification**

BGP path-selection order is implementation- and policy-dependent. Common vendor implementations have additional attributes and decision steps, so do not assume that every network follows one universal ranking sequence.

## **9. eBGP vs iBGP**

**eBGP — External BGP**

Used between BGP routers in different Autonomous Systems.

```
AS 100  -------- eBGP --------  AS 200
```

**iBGP — Internal BGP**

Used between BGP routers within the same Autonomous System.

```
              AS 100
       +------------------+
       |                  |
    Router A -------- Router B
             iBGP
       +------------------+
```

### **BGP transport**

BGP uses TCP port 179.

TCP provides reliable transport for exchanging BGP messages.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor



## **10. RIP vs OSPF vs BGP**

| Feature                     | RIP                         | OSPF                          | BGP                                |
| --------------------------- | --------------------------- | ----------------------------- | ---------------------------------- |
| Category                    | IGP                         | IGP                           | EGP                                |
| Algorithm type              | Distance Vector             | Link State                    | Path Vector                        |
| Main metric / decision      | Hop count                   | Cost                          | Policy and path attributes         |
| Maximum RIP metric          | 15 reachable hops           | No equivalent fixed hop limit | No equivalent fixed hop limit      |
| Topology knowledge          | Neighbour-advertised routes | Link-state topology database  | Advertised AS paths and attributes |
| Scope                       | Small internal networks     | Internal networks             | Between Autonomous Systems         |
| Common transport / protocol | UDP 520                     | IP protocol 89 for OSPFv2     | TCP 179                            |
| Example use                 | Simple LAN routing          | Enterprise/campus routing     | ISP and Internet interconnection   |

## **11. When Would You Use Each?**

### **RIP**

Use it as a simple example of distance-vector routing. It is easy to understand but has significant scalability limitations.

### **OSPF**

Use it for internal routing in larger enterprise networks where dynamic path calculation and hierarchical areas are useful.

### **BGP**

Use it to exchange routes between Autonomous Systems and apply routing policies across the Internet or large multi-provider networks.

### **The Internet example**

```
        ISP A — AS 100
              |
            eBGP
              |
        ISP B — AS 200
              |
            eBGP
              |
        ISP C — AS 300

Within each ISP, an IGP such as OSPF
may be used to route internal traffic.
```

An ISP may use OSPF or IS-IS internally and BGP to exchange routes with other providers.

**Key distinction:** OSPF figures out internal routes based on topology and cost. BGP decides which external routes to use based on reachability and policy.

# **3.8 Network-on-Chip (NoC) and SoC Communication**

Specialized hardware topic — study separately

So far, the unit has focused on communication between computers, routers and networks.

Network-on-Chip (NoC) applies similar communication ideas inside a single integrated circuit.

It is used in systems where many processing components need to communicate efficiently.

Examples include:

- CPUs.
- GPUs.
- AI accelerators.
- Memory controllers.
- Digital Signal Processors (DSPs).
- Image-processing units.
- Input/output controllers.

## **1. What is a System-on-Chip (SoC)?**

An SoC is an integrated circuit that combines multiple functional components into a single chip.

A smartphone processor, for example, may contain:

- CPU cores.
- GPU.
- Neural Processing Unit (NPU).
- Memory controller.
- Image Signal Processor (ISP).
- DMA controllers.
- Cache and memory interfaces.
- Communication interfaces.

### **Example: Smartphone SoC**

**Simplified System-on-Chip architecture**

**CPU cores**

**GPU**

**NPU**

**Image processor**

# **Network-on-Chip**

Routers + links + network interfaces

**Memory controller**

**I/O controller**

All these components need to exchange data.

A NoC provides an organised communication network that connects them.

## **2. What is Network-on-Chip?**

A Network-on-Chip is an on-chip communication architecture that connects processing and storage components through routers and communication links.

Instead of connecting every component to every other component with separate wires, components communicate through a shared network fabric.

### **Why is NoC required?**

Consider a chip containing 16 processing components.

If every component needs a direct connection to every other component, the number of connections grows rapidly.

The number of undirected pairwise connections in a fully connected arrangement of \\(N\\) components is:

\\[ \boxed{\frac{N(N-1)}{2}} \\]

For 16 components:

\\[ \frac{16(16-1)}{2}=120 \\]

A direct connection between every pair would require 120 separate connections.

This creates problems:

- Complex wiring.
- High routing congestion.
- Large physical design area.
- Difficult clock and timing management.
- More complex interconnection.

A NoC provides a structured way to connect many components using routers and links.

### **Real-world example: CPU and GPU communication**

Suppose a CPU needs to send data to a GPU.

```
CPU
 |
Network Interface
 |
NoC Router
 |
NoC Links
 |
NoC Router
 |
Network Interface
 |
GPU
```

The CPU sends a packet through its network interface. NoC routers forward it toward the GPU.

## **3. NoC Architecture**

A basic NoC contains four important components:

1. Processing Elements.
2. Network Interfaces.
3. Routers.
4. Links.

### **A. Processing Elements (PEs)**

A Processing Element is a computational or functional component connected to the NoC.

Examples:

- CPU core.
- GPU compute unit.
- AI accelerator.
- Memory controller.
- DSP.

A PE produces data that must be transferred to another PE or block.

### **B. Network Interface (NI)**

The Network Interface connects a processing component to the NoC.

Its responsibilities can include:

- Accepting data from the processing element.
- Creating network packets.
- Adding routing or transaction metadata.
- Converting packets into flits.
- Managing communication between the PE and router.
- Reassembling received data.

A network interface acts as a bridge between the component's local communication protocol and the NoC protocol.

### **C. Router**

An on-chip router forwards packets from its input ports to suitable output ports.

A router may have:

- Input buffers.
- Routing computation logic.
- Switch allocation logic.
- A crossbar.
- Output ports.
- Virtual-channel management.

### **D. Links**

Links are the communication paths connecting routers and processing nodes.

They may carry data between neighbouring on-chip routers.

### **Complete architecture**

```
+------+     +------+     +------+
| PE 0 |     | PE 1 |     | PE 2 |
+--+---+     +--+---+     +--+---+
   |            |            |
  NI           NI           NI
   |            |            |
 Router 0 ---- Router 1 ---- Router 2
   |            |            |
  Links        Links        Links
```

A typical NoC design contains many interconnected routers rather than one central controller.

## **4. Routers and Switches in NoC**

The terms router and switch describe different hardware roles in many NoC architectures.

A NoC router makes the path-selection and forwarding decisions for packets.

A switch fabric or crossbar inside the router connects selected input and output ports.

For example, suppose a router has five ports:

```
           North
             |
             |
West ------ Router ------ East
             |
             |
           South
             |
             |
           Local PE
```

The router may receive a packet from the local processing element and forward it toward the East output.

### **Router's basic operation**

1. Receive the packet or flit.
2. Read its destination information.
3. Calculate the required output direction.
4. Arbitrate for an available output.
5. Forward the data through the selected output port.

## **5. NoC Topologies**

A topology describes how routers and processing elements are connected.

Common NoC topologies include:

- Mesh.
- Torus.
- Ring.
- Tree.
- Crossbar.

### **A. Mesh Topology**

A mesh is one of the most common NoC topologies.

Routers are organised in rows and columns.

Each router may connect to its north, south, east, west neighbours and a local processing element.

Example: 4 × 4 mesh.

**4 × 4 mesh topology**

Each node represents an on-chip router. Links connect adjacent routers. Coordinates are written as (X,Y).

Advantages:

- Regular structure.
- Easy to scale into rows and columns.
- Relatively simple routing logic.
- Multiple paths may exist between certain routers.

Disadvantages:

- Corner and edge routers have fewer neighbours.
- Some traffic patterns create congestion around central regions.
- The number of hops can increase for distant destinations.

### **B. Torus Topology**

A torus is a mesh with wraparound links.

In an ordinary mesh, the leftmost column is not directly connected to the rightmost column.

In a torus, those boundary columns are connected.

Likewise, the top and bottom rows are connected.

```
Mesh:

R0 ---- R1 ---- R2
 |       |       |
R3 ---- R4 ---- R5
 |       |       |
R6 ---- R7 ---- R8

Torus:

Left and right edges wrap around.
Top and bottom edges also wrap around.
```

Advantages:

- More routing choices.
- Can reduce the average number of hops.
- Offers additional paths.

Disadvantages:

- More complicated physical wiring.
- Wraparound links may be expensive.
- Deadlock-free routing needs careful design.

### **C. Ring Topology**

In a ring, each router is connected to two neighbouring routers.

```
      R0 -------- R1
     /              \
    R3              R2
     \              /
      R5 -------- R4
```

Advantages:

- Simple structure.
- Low number of connections per router.
- Useful for small systems.

Disadvantages:

- Traffic may need to pass through many routers.
- Congestion can occur on heavily used ring links.
- Failure of a link or router can affect the ring unless redundancy is provided.

### **D. Tree Topology**

Routers are arranged hierarchically.

```
             Root
            /    \
          R1      R2
         / \     / \
        R3 R4   R5 R6
```

Advantages:

- Structured hierarchical organisation.
- Convenient for certain traffic patterns.
- Can support scalable layouts.

Disadvantages:

- Upper-level links may become bottlenecks.
- Failure of a high-level router can isolate a large portion of the network.

### **E. Crossbar Topology**

In a crossbar, components can connect through a switching matrix.

Conceptually:

```
          Output 1  Output 2  Output 3
Input 1      X         X         X
Input 2      X         X         X
Input 3      X         X         X
```

A crossbar can allow multiple independent input-output connections simultaneously, provided they do not conflict for the same output.

Advantages:

- Low contention between independent input-output pairs.
- Direct connectivity in small systems.
- Predictable switching behaviour.

Disadvantages:

- Hardware cost grows quickly as port count increases.
- The crossbar becomes complex for large numbers of ports.
- Physical area and power may become expensive.

### **Topology comparison**

| Topology | Structure        | Main advantage       | Main limitation                             |
| -------- | ---------------- | -------------------- | ------------------------------------------- |
| Mesh     | Grid             | Regular and scalable | Central congestion and increasing hop count |
| Torus    | Wraparound mesh  | More paths           | More complex wiring                         |
| Ring     | Circular         | Simple structure     | Potentially many hops                       |
| Tree     | Hierarchical     | Structured scaling   | Upper-level bottlenecks                     |
| Crossbar | Switching matrix | High connectivity    | Expensive as size increases                 |

## **6. XY Deterministic Routing**

XY routing is a deterministic routing algorithm commonly used in two-dimensional mesh NoCs.

It is a type of dimension-order routing.

The basic rule is:

1. Move along the X direction first.
2. Once the destination X coordinate is reached, move along the Y direction.

### **Example**

Suppose a packet starts at:

```
Source      = (0,0)
Destination = (2,3)
```

Using XY routing:

First move in the X direction:

```
(0,0) → (1,0) → (2,0)
```

Then move in the Y direction:

```
(2,0) → (2,1) → (2,2) → (2,3)
```

Complete route:

```
(0,0)
  |
  v
(1,0)
  |
  v
(2,0)
  |
  v
(2,1)
  |
  v
(2,2)
  |
  v
(2,3)
```

**XY routing — X first, then Y**

Highlighted path: (0,0) → (1,0) → (2,0) → (2,1) → (2,2) → (2,3).

### **Why is XY routing useful?**

- Simple route calculation.
- Predictable paths.
- Low routing-computation complexity.
- Can provide deadlock-free routing in a standard mesh when dimension-order constraints are correctly applied.

### **Limitation**

XY routing does not dynamically avoid a congested path.

If an alternative route is less congested, basic XY routing still follows its fixed dimension order.

The simplicity and deadlock properties of dimension-order XY routing, alongside its limited response to congestion, are discussed in NoC routing research.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.mdpi.com\&sz=32)

MDPI+1



## **7. Adaptive Routing**

Adaptive routing allows a NoC router to choose between multiple possible paths based on network conditions or other routing information.

For example, suppose a packet can travel through:

```
Path A: 4 hops, high congestion
Path B: 5 hops, low congestion
```

A minimal routing algorithm might always choose the shortest path.

An adaptive routing algorithm may choose Path B because it has less congestion.

### **Example**

```
          Router A
          /      \
         /        \
   Congested     Less congested
     Path A        Path B
         \        /
          \      /
         Destination
```

### **Advantages**

- Can avoid congested links.
- Can balance traffic.
- Can improve throughput under certain traffic patterns.

### **Disadvantages**

- More complicated routing logic.
- Can cause packets to take non-minimal paths.
- Deadlock and livelock prevention become more challenging.
- Requires careful resource and path-management rules.

## **8. Static vs Dynamic Routing in NoC**

### **Static routing**

The route is determined in advance and does not change according to current traffic conditions.

Example:

XY routing always moves in the X direction first, then Y.

### **Dynamic routing**

The path can be selected or changed based on runtime conditions such as congestion, link availability or faults.

### **Comparison**

| Feature              | Static routing   | Dynamic routing          |
| -------------------- | ---------------- | ------------------------ |
| Route decision       | Predefined       | Runtime-dependent        |
| Congestion awareness | Usually limited  | Can be included          |
| Hardware complexity  | Lower            | Often higher             |
| Adaptability         | Limited          | Greater                  |
| Example              | Basic XY routing | Adaptive minimal routing |

## **9. Packet-Switched vs Circuit-Switched Communication in NoC**

### **A. Packet-Switched Communication**

Data is divided into packets.

Each packet contains information such as:

- Destination.
- Control information.
- Payload.

Routers forward packets through the network.

Different packets can share the same physical links at different times.

```
PE A → Router → Router → PE B
       Packet-switched path
```

Advantages:

- Flexible sharing of resources.
- Efficient for bursty traffic.
- Different communication flows can use the network.

Disadvantages:

- Routing and buffering overhead.
- Variable latency.
- Congestion can occur.

### **B. Circuit-Switched Communication**

A dedicated path is established or reserved between source and destination before data transfer.

The reserved resources are then used for communication.

```
PE A → Router 1 → Router 2 → PE B
       Reserved path
```

Advantages:

- Predictable transmission resources after setup.
- Can offer predictable latency and bandwidth under appropriate assumptions.
- Useful for sustained communication patterns.

Disadvantages:

- Setup overhead.
- Reserved resources may remain unused.
- Less flexible when traffic is irregular.

### **Comparison**

| Feature             | Packet Switching                              | Circuit Switching            |
| ------------------- | --------------------------------------------- | ---------------------------- |
| Path setup          | Generally no dedicated end-to-end reservation | Required                     |
| Resource usage      | Shared dynamically                            | Reserved for the circuit     |
| Latency             | Can vary with contention                      | More predictable after setup |
| Resource efficiency | Good for bursty traffic                       | Can waste reserved capacity  |
| Flexibility         | High                                          | Lower                        |

## **10. Store-and-Forward Switching**

Store-and-forward switching is a packet-transfer technique in which an intermediate router receives the entire packet before forwarding it.

### **Working**

Suppose a packet moves through three routers.

```
PE A → Router 1 → Router 2 → Router 3 → PE B
```

Router 1:

1. Receives the full packet.
2. Stores it in its buffer.
3. Determines the output port.
4. Forwards the packet.

Router 2 repeats the process.

### **Advantage**

The complete packet is available before forwarding, so the router can perform packet-level checks before sending it onward.

### **Disadvantage**

It requires enough buffering to store the packet, and it adds per-hop storage delay.

If the packet has \\(L\\) bits and the link transmission rate is \\(R\\) bits per second, the ideal serialization time is:

\\[ \boxed{T\_{serialization}=\frac{L}{R}} \\]

In a store-and-forward network, a packet normally must be fully received before its forwarding transmission begins at the next router.

## **11. Wormhole Switching**

Wormhole switching is widely studied and used in NoC architectures.

Instead of treating the packet as one large unit for buffering, the packet is divided into smaller flow-control units called **flits**.

Flit stands for flow-control digit.

### **Types of flits**

A packet commonly consists of:

- Head flit.
- Body flits.
- Tail flit.

The head flit contains routing information and initiates route selection.

The body flits carry the remaining packet information.

The tail flit marks the end of the packet and releases the resources according to the design's protocol.

### **Working**

```
Packet:

+------+--------+--------+--------+
| Head | Body 1 | Body 2 | Tail   |
+------+--------+--------+--------+
```

In wormhole switching, the flits can pipeline through successive routers.

The head flit determines the route; following flits advance along the reserved path.

**Wormhole switching**

**Router 1**

**H**

**B1**

**B2**

**T**

Flits move through the link in sequence

**Router 2**

**H**

**B1**

**B2**

**T**

**Router 3 / Destination**

### **Advantages**

- Requires less per-router buffering than store-and-forward switching in many designs.
- Allows pipelined packet transfer.
- Can reduce packet-transfer latency under suitable conditions.

### **Disadvantages**

- Blocking a head flit can hold resources across multiple routers.
- Congestion can propagate backward.
- Requires careful flow control and deadlock prevention.

Research on NoC router design describes wormhole switching as a flit-pipelined technique in which the header flit selects the route and subsequent flits follow the reserved path.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.mdpi.com\&sz=32)

MDPI+1



## **12. Store-and-Forward vs Wormhole Switching**

| Feature            | Store-and-Forward                                       | Wormhole                                               |
| ------------------ | ------------------------------------------------------- | ------------------------------------------------------ |
| Transfer unit      | Whole packet                                            | Flits                                                  |
| Buffer requirement | Whole packet may need to be stored                      | Smaller per-router buffers may be sufficient           |
| Forwarding         | After full packet reception                             | Flits are pipelined                                    |
| Latency            | Can accumulate full-packet reception delays at each hop | Can be lower under appropriate conditions              |
| Blocking           | Packet waits at router                                  | A blocked header can hold resources over multiple hops |

## **13. Virtual Channels**

Virtual channels (VCs) are multiple logical channels implemented over a shared physical communication channel.

Each virtual channel has its own logical buffer and flow-control state.

### **Why are virtual channels required?**

Suppose Packet A is waiting for an output port that is busy.

Packet B is queued behind Packet A in the same FIFO buffer.

Even when Packet B could use another output port, it may be blocked behind Packet A.

This is called head-of-line blocking.

Virtual channels can provide separate queues so that Packet B can make progress independently.

### **Example**

```
          One Physical Link
                 |
       +---------+---------+
       |                   |
 Virtual Channel 1   Virtual Channel 2
       |                   |
    Packet A             Packet B
       |                   |
   Blocked path         Free path
```

Packets belonging to separate virtual channels can make progress independently when their requested resources are available.

### **Advantages**

- Reduces certain forms of head-of-line blocking.
- Allows multiple logical traffic flows over one physical link.
- Supports more flexible flow-control management.
- Can help design deadlock-free routing schemes.

### **Important distinction**

Virtual channels do not automatically guarantee deadlock freedom. The routing algorithm and the dependencies between virtual channels must be designed correctly.

They also do not simply multiply the physical link's bandwidth. Multiple virtual channels share the same physical communication capacity.

## **14. Deadlock**

A deadlock occurs when packets are permanently blocked because each is waiting for resources held by another packet.

### **Example**

Consider two packets:

- Packet A holds resource R1 and waits for R2.
- Packet B holds resource R2 and waits for R1.

```
Packet A:
Holds R1 → Waiting for R2

Packet B:
Holds R2 → Waiting for R1
```

Neither packet can progress.

This is a deadlock.

### **The four necessary conditions for deadlock**

The classical conditions are:

**1. Mutual Exclusion**

A resource can be held by only one packet or process at a time.

**2. Hold and Wait**

A packet holds one resource while waiting for another.

**3. No Preemption**

A resource cannot simply be taken away from the packet holding it.

**4. Circular Wait**

A cycle exists in which each packet waits for a resource held by the next packet.

### **Deadlock in NoC**

In NoC designs, a useful way to analyse deadlock is the channel dependency graph.

A cycle in the channel dependency graph indicates a potential deadlock condition, assuming the corresponding channels can be acquired and held in that order.

### **How can NoC deadlocks be prevented?**

- Restrict the order in which channels can be acquired.
- Use deadlock-free routing algorithms such as properly applied dimension-order routing.
- Use virtual channels to break resource dependency cycles.
- Provide escape channels with safe routing rules.

### **XY routing and deadlock**

In a standard mesh, XY routing follows a fixed dimension order: X first, then Y.

This restriction prevents certain cyclic channel dependencies, giving a deadlock-free routing scheme under the usual mesh assumptions.

## **15. Livelock**

Livelock occurs when packets continue moving or being rerouted but never reach their destinations.

Unlike deadlock, packets are not necessarily stationary. They may repeatedly make progress locally but fail to complete delivery.

### **Example**

```
Packet:
Router A → Router B → Router C → Router A
```

The packet keeps moving but never reaches the destination.

Livelock can arise in some adaptive routing approaches if routing decisions repeatedly send packets away from their destinations.

### **Deadlock vs Livelock**

| Feature          | Deadlock                   | Livelock                                  |
| ---------------- | -------------------------- | ----------------------------------------- |
| Packet behaviour | Stuck waiting              | Continues moving or being rerouted        |
| Progress         | No progress                | Local movement but no successful delivery |
| Resource use     | Resources remain held      | Resources may be repeatedly consumed      |
| Example          | Circular wait for channels | Packet repeatedly takes detours           |

## **16. NoC Performance Metrics**

Three important metrics are:

1. Latency.
2. Throughput.
3. Power consumption.

### **A. Latency**

Latency is the time taken for data to travel from its source to its destination.

In a NoC, latency can include:

- Network-interface delay.
- Packet injection delay.
- Router processing delay.
- Buffering delay.
- Link transmission delay.
- Contention and waiting delay.
- Destination ejection delay.

A simple conceptual expression is:

\\[ T\_{total}=T\_{injection}+T\_{network}+T\_{ejection} \\]

For a simple path with \\(H\\) hops, a rough model might be:

\\[ T\_{network}\approx H(T\_{router}+T\_{link})+T\_{queueing} \\]

This is a simplified model; actual latency depends on the switching method, packet size, flit pipeline and flow-control design.

**Example:**

Two routes are available:

| Path | Hops | Approximate router/link cost per hop |
| ---- | ---- | ------------------------------------ |
| A    | 3    | 2 ns                                 |
| B    | 5    | 2 ns                                 |

Ignoring queueing and other differences:

Path A:

\\[ 3\times2=6\text{ ns} \\]

Path B:

\\[ 5\times2=10\text{ ns} \\]

Path A has lower estimated network latency.

### **B. Throughput**

Throughput is the amount of data successfully delivered per unit time.

A simple expression is:

\\[ \boxed{\text{Throughput}=\frac{\text{Delivered data}}{\text{Time}}} \\]

Example:

A NoC delivers 10 GB of data in 2 seconds.

\\[ \text{Throughput}=\frac{10}{2}=5\text{ GB/s} \\]

In NoC design, throughput may be measured per link, per router, per network or across the entire chip.

It depends on:

- Link bandwidth.
- Number of active links.
- Traffic pattern.
- Router design.
- Buffer size.
- Congestion.
- Arbitration policy.

### **C. Power Consumption**

NoC routers and links consume power during data transfer.

Power is important because modern processors must balance performance, energy consumption and heat.

Dynamic power is often approximated by:

\\[ \boxed{P\_{dynamic}\approx\alpha C V^2f} \\]

Where:

- \\(\alpha\\) = switching activity factor.
- \\(C\\) = effective capacitance.
- \\(V\\) = supply voltage.
- \\(f\\) = switching frequency.

There is also leakage power, which exists even when circuits are not actively switching.

Important NoC power contributors include:

- Router buffers.
- Crossbar switching.
- Arbitration and control logic.
- Inter-router links.
- Clock distribution.

A design can provide high throughput but consume too much power for its target chip.

## **17. NoC Example — CPU Sending Data to an AI Accelerator**

Consider a System-on-Chip with:

- CPU.
- GPU.
- AI accelerator.
- Memory controller.
- Several routers.

The CPU wants to send a matrix to the AI accelerator.

### **Communication process**

1. CPU generates the data.
2. Its network interface creates a packet.
3. The packet is split into flits if required.
4. A router reads the packet destination.
5. XY or another routing algorithm selects a path.
6. Intermediate routers forward the flits.
7. The destination network interface receives the packet.
8. It reconstructs and delivers the data to the AI accelerator.

```
+-----+       +--------+       +--------+       +-----+
| CPU | ----> | Router | ----> | Router | ----> | NPU |
+-----+       +--------+       +--------+       +-----+
                 |                  |
               Flits              Routing
                 |                  |
                 +------------------+
```

This example brings together almost every topic in Section 3.8:

- Processing elements.
- Network interfaces.
- Routers and links.
- Mesh topology.
- XY routing.
- Packet switching.
- Wormhole switching.
- Virtual channels.
- Latency and throughput.

## **18. NoC vs Traditional IP Networking**

| Feature             | Traditional IP network                         | Network-on-Chip                                                           |
| ------------------- | ---------------------------------------------- | ------------------------------------------------------------------------- |
| Communication scope | Between hosts and networks                     | Inside an integrated circuit                                              |
| Main devices        | Hosts, routers, switches                       | Processing elements, NIs, NoC routers                                     |
| Addressing          | IP addresses                                   | Destination identifiers, coordinates or implementation-specific addresses |
| Data unit           | IP packet                                      | Packet divided into flits in many designs                                 |
| Routing             | IP routing and forwarding                      | XY, adaptive and other NoC algorithms                                     |
| Physical medium     | Copper, fibre, radio                           | On-chip wires and interconnects                                           |
| Main concerns       | Reachability, routing, congestion, reliability | Latency, throughput, power, area and deadlock                             |
| Examples            | Internet, LAN, data centre                     | CPU/GPU/SoC interconnect                                                  |

**Important takeaway:** A NoC router and an Internet router both forward data, but their architecture, scale, timing constraints and routing mechanisms are different.

# **Unit 3 — Final Revision and Interview Preparation**

## **1. Important differences to remember**

| Comparison                            | Core distinction                                                                                  |
| ------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Routing vs forwarding                 | Routing determines paths; forwarding moves packets to the next hop.                               |
| Connectionless vs connection-oriented | Connectionless sends independent datagrams; connection-oriented establishes a logical path first. |
| Network ID vs Host ID                 | Network ID identifies the subnet; Host ID identifies an interface within it.                      |
| Private vs public IP                  | Private addresses are for private networks; public addresses are globally routable.               |
| IPv4 vs IPv6                          | 32-bit addressing versus 128-bit addressing.                                                      |
| MTU vs packet size                    | MTU is the maximum supported packet size for a link; packet size is the actual packet length.     |
| Subnetting vs supernetting            | Subnetting divides a network; supernetting combines contiguous networks.                          |
| CIDR vs classful addressing           | CIDR uses flexible prefixes; classful addressing uses fixed classes.                              |
| NAT vs PAT                            | NAT maps addresses; PAT maps addresses and transport ports.                                       |
| ARP vs RARP                           | ARP resolves IPv4 to MAC; RARP historically resolved MAC to IPv4.                                 |
| Ping vs traceroute                    | Ping checks reachability; traceroute discovers the hop sequence.                                  |
| Static vs dynamic routing             | Manually configured routes versus routes learned through protocols.                               |
| Distance Vector vs Link State         | Neighbours advertise distances versus routers build a topology database.                          |
| RIP vs OSPF                           | Distance-vector versus link-state routing.                                                        |
| IGP vs BGP                            | Internal routing within an AS versus inter-AS routing.                                            |
| XY vs adaptive routing                | Fixed dimension order versus runtime path selection.                                              |
| Deadlock vs livelock                  | Packets are stuck waiting versus packets keep moving without reaching the destination.            |
| Store-and-forward vs wormhole         | Full packet buffering versus pipelined flit forwarding.                                           |

## **2. Important formulas — keep these ready**

### **IP addressing and subnetting**

| Topic                    | Formula                                     |
| ------------------------ | ------------------------------------------- |
| IPv4 address size        | \\(32\\) bits                               |
| IPv6 address size        | \\(128\\) bits                              |
| Host bits                | \\(h=32-\text{prefix length}\\)             |
| Total IPv4 addresses     | \\(2^h\\)                                   |
| Traditional usable hosts | \\(2^h-2\\)                                 |
| Borrowed subnet bits     | \\(s=\text{new prefix}-\text{old prefix}\\) |
| Number of subnets        | \\(2^s\\)                                   |
| Block size               | \\(256-\text{interesting mask octet}\\)     |
| Network address          | IP address AND subnet mask                  |
| Broadcast address        | Network address with all host bits set to 1 |

### **Routing**

| Topic                           | Formula / rule                               |
| ------------------------------- | -------------------------------------------- |
| Bellman-Ford principle          | \\(D_x(y)=\min_v\\{c(x,v)+D_v(y)\\}\\)       |
| Next-hop selection              | Forward toward the selected route's next hop |
| Longest Prefix Match            | Select the most specific matching prefix     |
| RIP maximum reachable hop count | 15                                           |
| RIP unreachable metric          | 16                                           |
| IPv4 default route              | `0.0.0.0/0`                                  |
| IPv6 default route              | `::/0`                                       |

### **NoC**

| Topic                          | Formula                                      |
| ------------------------------ | -------------------------------------------- |
| Fully connected pairwise links | \\(N(N-1)/2\\)                               |
| Throughput                     | Delivered data / time                        |
| Serialization time             | \\(L/R\\)                                    |
| Approximate network latency    | \\(H(T\_{router}+T\_{link})+T\_{queueing}\\) |
| Approximate dynamic power      | \\(\alpha CV^2f\\)                           |

## **3. Practical commands to try on your computer**

These commands connect the theory with actual network behaviour.

| Task                        | Windows              | Linux                   |
| --------------------------- | -------------------- | ----------------------- |
| Display IP configuration    | `ipconfig /all`      | `ip addr`               |
| Display routing table       | `route print`        | `ip route`              |
| Display ARP/neighbour cache | `arp -a`             | `ip neigh`              |
| Test connectivity           | `ping 8.8.8.8`       | `ping 8.8.8.8`          |
| Trace packet path           | `tracert google.com` | `traceroute google.com` |
| Release DHCP address        | `ipconfig /release`  | DHCP-client dependent   |
| Renew DHCP address          | `ipconfig /renew`    | DHCP-client dependent   |

### **Suggested hands-on experiment**

Run these commands on your laptop:

1. `ipconfig /all` — identify your IPv4 address, subnet mask, default gateway and DNS server.
2. `arp -a` — inspect your local IP-to-MAC mappings.
3. `route print` — find the default route and connected networks.
4. `ping 8.8.8.8` — test IP connectivity.
5. `tracert google.com` — examine the sequence of responding routers.

Then answer these questions:

- Is your laptop using a private or public IPv4 address?
- Which IP address is your default gateway?
- What is your subnet's network address?
- How many usable host addresses are available?
- Does your ARP table contain an entry for the default gateway?
- Which route is used to reach an external destination?

This is an excellent way to connect subnetting, ARP, ICMP, TTL, routing and forwarding.

## **4. Highest-priority topics for interviews**

I would give the following topics additional practice because they test both conceptual understanding and practical problem-solving.

**Priority 1 — Master thoroughly**

Very high

- CIDR and subnetting numerical problems.
- Network address, broadcast address and usable hosts.
- VLSM and route aggregation.
- Longest Prefix Match.
- Routing table and next-hop forwarding.
- ARP and default gateway behaviour.
- NAT vs PAT.
- TCP connection vs IP connectionless service.

  **Priority 2 — Understand algorithms**

  High
- Bellman-Ford and Dijkstra.
- Distance Vector vs Link State.
- Count-to-infinity and split horizon.
- RIP vs OSPF vs BGP.
- OSPF areas.
- BGP AS_PATH and LOCAL_PREF.

  **Priority 3 — Understand architecture**

  Specialized
- NoC topology.
- XY routing.
- Adaptive routing.
- Wormhole vs store-and-forward.
- Virtual channels.
- Deadlock and livelock.
- Latency, throughput and power.

## **5. Technical references and standards**

These are useful primary references for deeper study:

| Subject         | Reference                                                                                                                                                      |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| IPv4            | [RFC 791 — Internet Protocol](https://www.rfc-editor.org/info/rfc791/?utm_source=chatgpt.com)                                                                  |
| IPv6            | [RFC 8200 — IPv6 Specification](https://www.rfc-editor.org/rfc/rfc8200.html?utm_source=chatgpt.com)                                                            |
| IPv6 addressing | [RFC 4291 — IPv6 Addressing Architecture](https://www.rfc-editor.org/rfc/rfc4291.html?utm_source=chatgpt.com)                                                  |
| CIDR            | [RFC 4632 — CIDR](https://www.rfc-editor.org/info/rfc4632/?utm_source=chatgpt.com)                                                                             |
| ARP             | [RFC 826 — Address Resolution Protocol](https://www.rfc-editor.org/info/rfc826/?utm_source=chatgpt.com)                                                        |
| ICMP            | [RFC 792 — ICMP](https://www.rfc-editor.org/info/rfc792/?utm_source=chatgpt.com)                                                                               |
| DHCP            | [RFC 2131 — DHCP](https://www.rfc-editor.org/info/rfc2131/?utm_source=chatgpt.com)                                                                             |
| NAT             | [RFC 3022 — Traditional NAT](https://www.rfc-editor.org/info/rfc3022/?utm_source=chatgpt.com)                                                                  |
| RIP             | [RFC 2453 — RIP Version 2](https://www.rfc-editor.org/rfc/rfc2453.html?utm_source=chatgpt.com)                                                                 |
| OSPF            | [RFC 2328 — OSPF Version 2](https://www.rfc-editor.org/info/rfc2328/?utm_source=chatgpt.com)                                                                   |
| BGP             | [RFC 4271 — BGP-4](https://www.rfc-editor.org/info/rfc4271/?utm_source=chatgpt.com)                                                                            |
| NoC             | [A Survey of Research and Practices of Network-on-Chip](https://www.ecb.torontomu.ca/~courses/ee8205/Data-Sheets/noc/NoC-Survey-06.pdf?utm_source=chatgpt.com) |
