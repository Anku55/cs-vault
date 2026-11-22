# **Unit 2 — Data Link Layer and High-Speed Interconnects**

This unit explains how data moves reliably across a **single network link**, how devices share a communication medium, how Ethernet switches forward frames, and how high-speed hardware components transfer data between chips.

A useful way to understand the whole unit is to follow the journey of a message:

**How data travels through a network**

**Application data**

The information created by an application

Transport layer adds its header

**Network packet**

Source and destination IP addresses

Data link layer adds its header and trailer

**Data link frame**

Source and destination MAC addresses, control information and error-checking information

Physical layer transmits bits as signals

**Bits on the medium**

Electrical, optical or radio signals

The **Data Link Layer** is responsible for communication between directly connected network interfaces. It does not normally deliver data all the way across the Internet; routers and the Network Layer handle forwarding between different networks.

# **2.1 Data Link Layer Fundamentals**

## **1. What is the Data Link Layer?**

The Data Link Layer is the second layer of the OSI model. It provides communication between two devices connected through the same local link.

Its main job is to take a packet received from the Network Layer, place it inside a frame, and transmit that frame over the physical medium.

**Example: Sending data from a laptop to a Wi-Fi router**

Suppose your laptop is connected to your home Wi-Fi and you open a website.

1. The application generates a request.
2. The Network Layer creates an IP packet addressed to the destination.
3. The Data Link Layer puts that packet inside a Wi-Fi frame.
4. The Wi-Fi interface transmits the frame through radio signals.
5. The router receives the frame, processes it, and forwards the IP packet toward the Internet.

The Wi-Fi frame is used for the local wireless link. When the router forwards the packet over a different link, it creates a new link-layer frame appropriate for that next link.

## **2. Responsibilities of the Data Link Layer**

### **A. Framing**

The Network Layer gives the Data Link Layer a packet. The Data Link Layer wraps it in a frame.

A frame provides boundaries and control information so that the receiver knows where the packet begins and ends.

**Encapsulation at the Data Link Layer**

**Network-layer packet**

Add link-layer header and trailer

**Header**

**Packet / Payload**

**Trailer**

Typical conceptual frame layout. Exact fields differ by protocol.

### **B. Physical Addressing**

The Data Link Layer uses physical addresses, commonly called **MAC addresses**, to identify network interfaces on a local link.

For example:

- Laptop MAC: `A4:5E:60:12:34:56`
- Router Wi-Fi MAC: `80:2A:A8:11:22:33`

On an Ethernet LAN, a switch uses destination MAC addresses to decide which port should receive a frame.

### **C. Error Detection**

Bits can get corrupted while travelling through a communication medium.

For example, the sender transmits:

`10110010`

But the receiver obtains:

`10100010`

The Data Link Layer can use error-detection methods such as:

- Parity
- Checksum
- Cyclic Redundancy Check (CRC)

These methods help the receiver detect corruption. They do not all correct it; many protocols discard a damaged frame and rely on retransmission or a higher layer.

### **D. Flow Control**

Flow control prevents a fast sender from overwhelming a slower receiver.

Imagine:

- Sender can transmit 100 frames per second.
- Receiver can process only 20 frames per second.

Without flow control, frames may accumulate in the receiver's buffer and eventually be lost.

Mechanisms such as Stop-and-Wait and Sliding Window regulate how much data can be sent before acknowledgements are received.

### **E. Error Control**

Error control deals with lost, damaged or duplicated frames.

It can involve:

- Error detection
- Acknowledgements
- Timers
- Retransmissions
- Sequence numbers

For example, if a sender transmits frame 5 but receives no acknowledgement before the timeout, it may retransmit frame 5.

### **F. Medium Access Control**

When many devices share one communication medium, the Data Link Layer needs rules that control who can transmit and when.

Examples:

- Wi-Fi uses CSMA/CA-based channel access.
- Traditional shared Ethernet used CSMA/CD.
- Token-passing networks give transmission permission to the device holding a token.

## **3. Two Sublayers: LLC and MAC**

In the IEEE 802 architecture, the Data Link Layer is commonly divided into two sublayers:

**Network Layer**

**Data Link Layer**

## **LLC**

Logical Link Control

Interface to upper-layer protocols

## **MAC**

Media Access Control

Frame addressing and access to the medium

**Physical Layer**

### **Logical Link Control (LLC)**

LLC is the upper sublayer.

Its purpose is to provide a logical interface between the Network Layer and the underlying MAC mechanism.

Its traditional responsibilities include:

- Identifying which upper-layer protocol is carried.
- Providing logical link-control functions.
- Supporting different network-layer protocols over the same underlying link technology.

**Example:** A local network can carry IPv4 and IPv6 traffic over Ethernet. The link-layer format needs a way to identify what type of network-layer payload is inside a frame.

In modern Ethernet, the EtherType field is commonly used for this identification. LLC/SNAP formats also exist, but you should not assume every Ethernet frame uses an LLC header.

### **Media Access Control (MAC)**

MAC is the lower sublayer.

It deals with how a device uses the communication medium.

Responsibilities include:

- MAC addressing
- Frame transmission and reception
- Medium-access rules
- Frame-level control specific to the link technology

**Example:** In Wi-Fi, multiple devices use the same radio channel. MAC rules coordinate transmission so that devices reduce the likelihood of interfering with one another.

## **4. Framing and Frame Structure**

A **frame** is the Data Link Layer's protocol data unit (PDU).

A general conceptual structure is:

```
+------------------+----------------------+------------------+
|    Header        |       Payload        |     Trailer      |
+------------------+----------------------+------------------+
| Addresses and    | Network-layer packet | Error detection  |
| control fields   | or other data        | information      |
+------------------+----------------------+------------------+
```

Not every protocol uses exactly this layout. Ethernet, Wi-Fi, PPP and HDLC have different frame formats.

### **Real-world Ethernet frame**

A simplified Ethernet frame contains:

```
+----------+----------+----------+----------+-------------+---------+
| Preamble | SFD      | Dest MAC | Src MAC  | Type/Length | Payload |
+----------+----------+----------+----------+-------------+---------+
                                                            |
                                                            v
                                                   +----------------+
                                                   | FCS (CRC-32)   |
                                                   +----------------+
```

The Ethernet preamble and Start Frame Delimiter (SFD) help the receiver synchronize with the incoming frame. They are transmitted on the medium but are not normally counted as part of the Ethernet MAC frame itself.

### **Physical addressing: MAC address**

A MAC address identifies a network interface at the link layer. Ethernet MAC addresses are generally 48 bits long.

Example:

`00:1A:2B:3C:4D:5E`

It consists of six hexadecimal octets:

| Octet | Value |
| ----- | ----- |
| 1     | `00`  |
| 2     | `1A`  |
| 3     | `2B`  |
| 4     | `3C`  |
| 5     | `4D`  |
| 6     | `5E`  |

Each hexadecimal digit represents 4 bits. Therefore:

\\[ 6 \times 8 = 48 \text{ bits} \\]

The first part commonly identifies an organizationally assigned address block, while the remaining bits identify an interface within that block. The exact interpretation depends on the address type and allocation rules.

### **MAC address types**

**Unicast:** Identifies one destination interface.

Example: A switch forwards a unicast frame to the port associated with the destination MAC address.

**Multicast:** Identifies a group of interfaces.

Example: IPv6 uses link-layer multicast to deliver certain control messages to a group of nodes.

**Broadcast:** Identifies all stations on the local Ethernet broadcast domain.

The standard Ethernet broadcast MAC address is:

`FF:FF:FF:FF:FF:FF`

A broadcast frame is delivered to all stations in that Layer 2 broadcast domain, subject to network configuration and filtering.

### **MAC address vs IP address**

| Feature     | MAC address                              | IP address                                 |
| ----------- | ---------------------------------------- | ------------------------------------------ |
| OSI layer   | Data Link                                | Network                                    |
| Purpose     | Local link delivery                      | Logical addressing and routing             |
| Common form | 48-bit Ethernet address                  | IPv4: 32 bits; IPv6: 128 bits              |
| Used by     | Ethernet switches and network interfaces | Hosts and routers                          |
| Scope       | Link-local                               | Can be used across interconnected networks |
| Example     | `00:1A:2B:3C:4D:5E`                      | `192.168.1.10`                             |

**Real-world example: Opening a website**

Your laptop wants to send an IP packet to a web server outside your local network.

1. The IP packet has the web server's IP address as its destination.
2. Your laptop determines that the destination is outside its local subnet.
3. It sends the packet to the MAC address of its default gateway (your router).
4. The router removes the incoming link-layer frame and forwards the IP packet using a new frame on the next link.

The destination IP can remain the web server's IP address, while the link-layer destination MAC changes at each routed hop.

**Important:** MAC addressing and IP addressing solve different problems. A MAC address is used for local-link delivery; an IP address supports logical addressing and routing between networks.

# **2.2 Framing Techniques**

## **1. What is framing?**

The Physical Layer transmits a continuous stream of bits. The receiver needs a way to identify where one frame ends and the next frame begins.

**Framing** is the process of dividing a continuous stream of bits into identifiable units called frames.

Consider this bit stream:

```
1011010011101010110011010110...
```

Without framing, the receiver may not know which bits belong to which message.

Framing solves this problem by introducing boundaries.

```
Continuous bit stream:

... | FRAME 1 | FRAME 2 | FRAME 3 | FRAME 4 | ...
```

Different Data Link Layer protocols use different techniques to identify frame boundaries.

## **2. Character Count**

In the character-count technique, the frame header contains a field that indicates the number of bytes in the frame.

### **Example**

Suppose a frame contains 5 bytes.

```
+----------------+-----------------------------+
| Length = 5     | Data                        |
+----------------+-----------------------------+
| 00000101       | A  B  C  D  E                |
+----------------+-----------------------------+
```

The receiver reads the length field and knows that the frame contains five bytes.

### **Working**

1. The sender inserts the frame length in the header.
2. The receiver reads the length field.
3. It counts the specified number of bytes.
4. It identifies the end of the frame.

### **Problem with character count**

Suppose the length field itself becomes corrupted.

The sender transmits:

```
Length = 5
Data   = A B C D E
```

But the receiver reads:

```
Length = 8
```

The receiver may consume bytes belonging to the next frame as part of the current frame.

This can destroy synchronization between sender and receiver.

**Key point:** Character count is easy to understand, but a corrupted count can cause framing errors.

## **3. Byte Stuffing**

Byte stuffing is used in byte-oriented protocols.

A special byte, called a **flag byte**, marks the beginning and end of a frame.

Suppose the flag byte is:

```
FLAG = 7E
```

The receiver treats `7E` as a frame delimiter.

But what happens when the actual data contains `7E`?

The receiver might incorrectly interpret it as the end of the frame.

### **Solution: Escape byte**

The sender inserts an escape byte before a reserved byte appearing inside the payload.

For illustration:

```
FLAG = 7E
ESC  = 7D
```

Suppose the original data is:

```
A  B  7E  C  D
```

Since `7E` is reserved, the sender escapes it:

```
A  B  7D  7E  C  D
```

The transmitted frame becomes:

```
+------+---------------------------+------+
| 7E   | A B 7D 7E C D             | 7E   |
+------+---------------------------+------+
 Start           Data                End
```

The receiver performs **byte destuffing**:

1. It detects the starting flag.
2. It reads the payload.
3. When it sees the escape byte, it treats the following byte as data.
4. It removes the escape byte.
5. It detects the ending flag.

### **Real-world application**

PPP uses an escape mechanism for asynchronous links so that flag-like and other special control bytes can be transmitted transparently. Its exact escaping rules are specified in its framing standard.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



## **4. Bit Stuffing**

Bit stuffing is used in bit-oriented protocols such as HDLC.

Instead of escaping entire bytes, the sender inserts an extra bit into the bit stream.

### **Why is bit stuffing needed?**

HDLC uses this special flag:

```
01111110
```

Suppose the same pattern occurs inside the payload.

The receiver might mistake payload bits for a frame boundary.

### **Solution**

The sender inserts a `0` after every sequence of five consecutive `1` bits in the payload.

This prevents the payload from accidentally containing the complete flag pattern.

### **Example**

Original data:

```
1111101111110
```

Apply bit stuffing:

- After the first five consecutive `1` bits, insert `0`.
- After the next five consecutive `1` bits, insert `0`.

Stuffed data:

```
111110011111010
```

**Bit stuffing concept**

**Original payload**

# **11111 0 11111 0**

Insert a 0 after each group of five consecutive 1s

**Stuffed payload**

# **1111100 1111100**

The receiver removes a `0` that follows five consecutive `1` bits in the stuffed data.

### **Important distinction**

- **Byte stuffing:** Inserts an escape byte.
- **Bit stuffing:** Inserts an extra bit.

## **5. Physical Layer Coding Violations**

Some Physical Layer encoding schemes reserve particular signal patterns that are not used for ordinary data.

A protocol may use one of these reserved patterns to indicate a frame boundary or special control event.

For example, if a particular encoding uses certain code groups for data, an otherwise-invalid code group can be reserved to indicate a control symbol.

The receiver recognizes the special pattern without needing to escape a sequence from the payload.

### **Advantage**

This technique can identify boundaries or control information without inserting extra bytes into ordinary data.

### **Limitation**

It is possible only when the Physical Layer encoding provides unused or reserved symbols.

## **6. Frame Delimiting**

Frame delimiting means identifying the start and end of a frame.

Different approaches include:

| Technique        | How the boundary is found                      |
| ---------------- | ---------------------------------------------- |
| Character count  | Length field                                   |
| Byte stuffing    | Special bytes and escape bytes                 |
| Bit stuffing     | Special bit pattern protected by inserted bits |
| Coding violation | Reserved Physical Layer symbols                |

### **Quick comparison**

| Technique        | Unit manipulated | Main idea                          |
| ---------------- | ---------------- | ---------------------------------- |
| Character count  | Bytes            | Specify frame length               |
| Byte stuffing    | Bytes            | Escape reserved bytes              |
| Bit stuffing     | Bits             | Insert 0 after five consecutive 1s |
| Coding violation | Encoded symbols  | Use reserved symbols for control   |

**Remember:** Framing is about identifying boundaries. Error detection is about checking whether the received frame was corrupted. They solve different problems.

# **2.3 Error Detection and Correction**

High-priority topic

This is one of the most important sections of the unit because it combines conceptual questions, numerical problems and practical communication-system design.

## **1. Why do transmission errors occur?**

Data travels through physical communication channels such as:

- Copper cables
- Fibre-optic cables
- Wireless radio channels
- Chip-to-chip electrical links

During transmission, noise, interference, attenuation, crosstalk or hardware faults can alter the transmitted bits.

Example:

```
Sender transmits:    10110110
Receiver receives:   10100110
                           ^
                    One bit changed
```

The sender transmitted `1` at one position, but the receiver received `0`.

This is called a transmission error.

## **2. Types of Transmission Errors**

### **A. Single-bit error**

A single-bit error occurs when exactly one bit in a data unit changes.

Example:

```
Original:   10110010
Received:   10100010
                  ^
```

Only one bit has changed.

### **B. Burst error**

A burst error occurs when two or more bits within a span of bits are affected.

Example:

```
Original:   1011001011101010
Received:   1010110010101010
                 <---->
             Error span
```

A burst error may affect multiple bits. Not every bit inside the span must necessarily be incorrect.

### **Difference**

| Feature            | Single-bit error  | Burst error                                |
| ------------------ | ----------------- | ------------------------------------------ |
| Affected bits      | Exactly one       | Two or more                                |
| Error distribution | One position      | Spread over a span                         |
| Common concern     | Isolated bit flip | Noise or interference affecting a sequence |
| Example            | `1010` → `1000`   | `10110010` → `10001110`                    |

Wireless links, storage interfaces and high-speed serial communication systems all need techniques to detect and sometimes correct these errors.

## **3. Error Detection vs Error Correction**

Before studying individual methods, understand this distinction:

**Error detection:** The receiver determines that the received data is probably or definitely invalid, according to the code's detection capability.

**Error correction:** The receiver uses redundant information to identify and correct errors, without necessarily asking the sender to retransmit.

**Two ways to deal with corrupted data**

**Error detection**

Detect corruption

Often discard the damaged frame and request retransmission.

**Error correction**

Recover data

Use redundant bits or symbols to reconstruct the original data.

## **4. Parity Check**

Parity is one of the simplest error-detection techniques.

The sender adds an extra bit, called a **parity bit**, to make the number of `1` bits either even or odd.

### **A. Even parity**

In even parity, the total number of `1` bits, including the parity bit, must be even.

Example:

Data:

```
1011001
```

Number of 1s = 4.

Since 4 is already even:

```
Parity bit = 0
```

Transmitted data:

```
10110010
```

### **B. Odd parity**

In odd parity, the total number of `1` bits must be odd.

Using the same data:

```
1011001
```

Number of 1s = 4.

Four is even, so the parity bit must be 1.

Transmitted data:

```
10110011
```

### **Working example**

Suppose even parity is used.

Sender transmits:

```
10110010
```

The receiver receives:

```
10100010
```

Count the received `1` bits:

- Received number of ones = 3.
- Expected even parity.
- Three is odd.

Therefore, the receiver detects an error.

### **Limitation of parity**

A simple parity bit cannot detect every error.

For example, if two bits change, the number of `1` bits may still have the expected parity.

```
Transmitted: 10110010
Received:    10010000
```

Two bits changed. The even-parity condition remains satisfied, so the error can go undetected.

**Exam point:** A single parity bit detects every odd number of bit errors, but cannot detect an even number of bit errors.

## **5. Two-Dimensional Parity**

Two-dimensional parity improves on simple parity by arranging data into rows and columns.

The sender calculates parity for every row and every column.

### **Example**

Consider three rows of four data bits each:

```
             Data bits
          C1  C2  C3  C4
        +---+---+---+---+
Row 1   | 1 | 0 | 1 | 1 |
        +---+---+---+---+
Row 2   | 1 | 1 | 0 | 0 |
        +---+---+---+---+
Row 3   | 1 | 0 | 0 | 1 |
        +---+---+---+---+
```

Using even parity:

**Row parity**

- Row 1: `1011` has three ones, so parity = 1.
- Row 2: `1100` has two ones, so parity = 0.
- Row 3: `1001` has two ones, so parity = 0.

**Column parity**

- Column 1: `111` has three ones, so parity = 1.
- Column 2: `010` has one one, so parity = 1.
- Column 3: `100` has one one, so parity = 1.
- Column 4: `101` has two ones, so parity = 0.

The complete layout becomes:

```
          C1  C2  C3  C4 | Row parity
        +---+---+---+---+---+
Row 1   | 1 | 0 | 1 | 1 | 1 |
        +---+---+---+---+---+
Row 2   | 1 | 1 | 0 | 0 | 0 |
        +---+---+---+---+---+
Row 3   | 1 | 0 | 0 | 1 | 0 |
        +---+---+---+---+---+
        | 1 | 1 | 1 | 0 |
        +---+---+---+---+
          Column parity
```

### **How can it detect and correct an error?**

Suppose Row 2, Column 3 changes from 0 to 1.

```
          C1  C2  C3  C4
        +---+---+---+---+
Row 1   | 1 | 0 | 1 | 1 |
        +---+---+---+---+
Row 2   | 1 | 1 | 1 | 0 |  <-- Row parity fails
        +---+---+---+---+
Row 3   | 1 | 0 | 0 | 1 |
        +---+---+---+---+
                  ^
             Column 3 parity fails
```

The receiver identifies:

- The incorrect row: Row 2.
- The incorrect column: Column 3.

Their intersection identifies the corrupted bit.

This allows correction of a single-bit error under the scheme's assumptions.

### **Limitation**

Two-dimensional parity is not guaranteed to correct arbitrary multiple-bit errors. Some error patterns can produce ambiguous results or escape detection.

## **6. Checksum**

A checksum is calculated by adding data words and sending a complement of the resulting sum.

The receiver performs the same addition, including the received checksum, and checks the result.

The Internet checksum used by protocols such as IPv4, TCP and UDP is based on 16-bit one's-complement arithmetic.

### **Working example using 8-bit words**

For a simpler illustration, use two 8-bit words:

```
Word 1 = 10011001
Word 2 = 11100010
```

Convert them to hexadecimal:

```
10011001 = 0x99
11100010 = 0xE2
```

### **Step 1: Add the words**

```
  10011001
+ 11100010
----------
1 01111011
```

The result has a carry beyond 8 bits.

### **Step 2: Wrap around the carry**

Add the carry bit to the lower 8 bits:

```
01111011
+      1
--------
01111100
```

### **Step 3: Take the one's complement**

Invert every bit:

```
01111100
--------
10000011
```

Therefore:

```
Checksum = 10000011
```

### **Step 4: Receiver verification**

The receiver adds the original data words and the checksum using one's-complement arithmetic.

The expected result is:

```
11111111
```

That indicates no error was detected.

**Important:** This is a simplified 8-bit teaching example. Internet checksums normally operate on 16-bit words.

### **Limitation**

Checksums are useful and inexpensive, but some different error patterns can produce the same checksum. CRC is generally more effective at detecting common burst-error patterns.

## **7. Cyclic Redundancy Check (CRC)**

CRC is one of the most important error-detection techniques.

It is widely used in communication and storage systems because it detects many common transmission-error patterns very effectively.

### **Core idea**

CRC treats a bit sequence as a polynomial over binary arithmetic, also called GF(2).

The sender divides the data polynomial by a predetermined **generator polynomial** using modulo-2 division.

In modulo-2 arithmetic:

- Addition is XOR.
- Subtraction is XOR.
- There are no carries or borrows.

XOR truth table:

| A | B | A XOR B |
| - | - | ------- |
| 0 | 0 | 0       |
| 0 | 1 | 1       |
| 1 | 0 | 1       |
| 1 | 1 | 0       |

### **CRC sender procedure**

1. Choose a generator polynomial.
2. Append zeros to the data. The number of zeros equals the degree of the generator.
3. Divide the resulting bit sequence by the generator using modulo-2 division.
4. Take the remainder.
5. Replace the appended zeros with the remainder.
6. Transmit the resulting codeword.

### **Numerical problem**

Suppose:

```
Data = 1101
Generator = 1011
```

The generator has four bits, so its degree is 3.

Therefore, append three zeros:

```
1101 → 1101000
```

Now perform modulo-2 division using XOR.

```
              1011
          ------------
1011 ) 1101000
       1011
       ----
        01100
         1011
         ----
          01110
           1011
           ----
            0101
             0000
             ----
              001
```

The remainder is:

```
CRC = 001
```

Therefore, the transmitted codeword is:

```
Original data = 1101
CRC remainder = 001

Transmitted codeword = 1101001
```

### **Receiver procedure**

The receiver divides the received codeword by the same generator.

```
Received codeword = 1101001
Generator         = 1011
```

If the remainder is zero, the receiver treats the frame as passing the CRC check.

If the remainder is non-zero, an error is detected.

### **What is a generator polynomial?**

A generator polynomial defines the CRC calculation.

For the generator:

```
1011
```

The polynomial is:

\\[ G(x)=x^3+x+1 \\]

The bit sequence `1011` represents the coefficients of:

\\[ 1x^3+0x^2+1x+1 \\]

### **Why is CRC powerful?**

CRC can detect:

- All single-bit errors for commonly used generator polynomials.
- Many multiple-bit errors.
- Many burst errors.
- All burst errors whose length is no greater than the degree of the generator, for a generator with a nonzero constant term.

The exact error-detection capability depends on the generator polynomial.

### **CRC vs Checksum vs Parity**

| Feature               | Parity              | Checksum                 | CRC                             |
| --------------------- | ------------------- | ------------------------ | ------------------------------- |
| Complexity            | Very low            | Low                      | Higher                          |
| Error detection       | Limited             | Good for many errors     | Strong for many common patterns |
| Burst-error detection | Limited             | Depends on error pattern | Particularly effective          |
| Main operation        | Bit counting        | Addition and complement  | Modulo-2 polynomial division    |
| Example use           | Simple serial links | IP, TCP, UDP             | Ethernet FCS                    |

**Remember:** CRC is primarily an error-detection method. Ethernet's FCS uses CRC-32; a receiver that detects a bad FCS generally discards the frame rather than correcting its payload.

## **8. Forward Error Correction (FEC)**

So far, we have discussed techniques that detect errors.

But consider a communication system where retransmission is expensive or impossible.

Examples:

- Satellite communication
- Deep-space communication
- Real-time streaming
- High-speed serial links
- Optical communication

If the receiver detects an error, waiting for retransmission may be undesirable or impossible.

FEC solves this by adding redundant information that enables the receiver to correct certain errors directly.

### **Basic working**

```
Sender:
Data → Add redundant bits → Transmit

                         ↓

Receiver:
Receive → Decode redundancy → Correct supported errors
```

For example, suppose the original message is:

```
1011
```

An error-correcting encoder produces a longer codeword:

```
1011010
```

The extra bits provide information that helps locate an error.

The exact code depends on the coding scheme.

## **9. Hamming Distance**

Hamming distance is the number of bit positions at which two binary strings of equal length differ.

### **Example**

```
A = 101101
B = 100001
```

Compare each position:

```
A = 1 0 1 1 0 1
B = 1 0 0 0 0 1
        ^ ^
```

They differ at two positions.

Therefore:

\\[ d(A,B)=2 \\]

### **Minimum Hamming distance**

The minimum Hamming distance of a code is the smallest distance between any two different valid codewords.

It is written as:

\\[ d\_{min} \\]

This value determines how many errors a code can detect or correct.

**Error detection capability:**

\\[ \boxed{\text{Detect up to } d\_{min}-1 \text{ errors}} \\]

**Error correction capability:**

\\[ \boxed{\text{Correct up to } \left\lfloor\frac{d\_{min}-1}{2}\right\rfloor \text{ errors}} \\]

### **Example**

Suppose a code has:

\\[ d\_{min}=5 \\]

It can detect:

\\[ 5-1=4 \text{ errors} \\]

It can correct:

\\[ \left\lfloor\frac{5-1}{2}\right\rfloor=2 \text{ errors} \\]

### **Why does correction require more distance?**

Imagine two valid codewords are very close to each other.

If a bit changes, the corrupted word might become indistinguishable from another valid codeword.

Larger Hamming distance gives the receiver more room to identify the original codeword.

## **10. Hamming Code**

Hamming code is an error-correcting code that adds parity bits at carefully selected positions.

A commonly studied version is **Hamming (7,4)**:

- 4 data bits.
- 3 parity bits.
- 7 total transmitted bits.

It can correct any one-bit error in a codeword.

### **Step 1: Calculate the number of parity bits**

Let:

- \\(m\\) = number of data bits.
- \\(r\\) = number of parity bits.

The number of parity bits must satisfy:

\\[ \boxed{2^r \geq m+r+1} \\]

Suppose:

\\[ m=4 \\]

Try \\(r=3\\):

\\[ 2^3 \geq 4+3+1 \\]

\\[ 8\geq8 \\]

So three parity bits are sufficient.

### **Step 2: Place the parity bits**

Parity bits are placed at positions that are powers of 2:

```
Position:  1  2  3  4  5  6  7
Type:      P1 P2 D1 P4 D2 D3 D4
```

Suppose the data is:

```
1011
```

Place it in positions 3, 5, 6 and 7.

```
Position:  1  2  3  4  5  6  7
Bit:       ?  ?  1  ?  0  1  1
```

### **Step 3: Calculate parity bits**

Assume even parity.

**P1 checks positions 1, 3, 5 and 7.**

These positions contain:

```
P1, 1, 0, 1
```

There are two data ones, so:

```
P1 = 0
```

**P2 checks positions 2, 3, 6 and 7.**

```
P2, 1, 1, 1
```

There are three data ones, so:

```
P2 = 1
```

**P4 checks positions 4, 5, 6 and 7.**

```
P4, 0, 1, 1
```

There are two data ones, so:

```
P4 = 0
```

### **Step 4: Construct the final codeword**

```
Position:  1  2  3  4  5  6  7
Bit:       0  1  1  0  0  1  1
```

The transmitted codeword, written from position 1 to position 7, is:

```
0110011
```

### **Step 5: Detect and correct an error**

Suppose position 6 gets corrupted:

```
Sent:       0 1 1 0 0 1 1
Received:   0 1 1 0 0 0 1
                         ^
```

The receiver recalculates the parity checks.

Let:

- \\(S_1\\) be the result for P1's group.
- \\(S_2\\) be the result for P2's group.
- \\(S_4\\) be the result for P4's group.

For this corrupted codeword:

```
S1 = 0
S2 = 1
S4 = 1
```

The error position is calculated as:

\\[ S= S_1+2S_2+4S_4 \\]

\\[ S=0+2(1)+4(1) \\]

\\[ \boxed{S=6} \\]

Therefore, bit 6 is corrupted.

The receiver flips bit 6 back to its original value.

### **What is a syndrome?**

The syndrome is the pattern of parity-check failures.

It identifies the position of an error in a single-error-correcting Hamming code.

**Important:** The basic Hamming (7,4) code corrects a single-bit error. It does not guarantee correction of arbitrary multiple-bit errors. An extended Hamming code adds an overall parity bit to provide single-error correction and double-error detection (SECDED).

## **11. Redundancy Bits Calculation**

For a Hamming code, the number of parity bits is calculated using:

\\[ \boxed{2^r\geq m+r+1} \\]

### **Example 1**

Suppose there are 8 data bits.

Find the minimum number of parity bits.

Try \\(r=3\\):

\\[ 2^3=8 \\]

\\[ m+r+1=8+3+1=12 \\]

Not sufficient.

Try \\(r=4\\):

\\[ 2^4=16 \\]

\\[ 8+4+1=13 \\]

Sufficient.

Therefore:

\\[ \boxed{r=4} \\]

Total codeword length:

\\[ n=m+r=8+4=12 \\]

### **Example 2**

Suppose there are 11 data bits.

Try \\(r=4\\):

\\[ 2^4=16 \\]

\\[ 11+4+1=16 \\]

Therefore, four parity bits are sufficient.

## **12. Reed-Solomon Codes**

Reed-Solomon is a powerful block error-correcting code.

Unlike the basic Hamming code examples that operate on individual bits, Reed-Solomon codes operate on **symbols**, where a symbol may contain multiple bits.

For example, a symbol could be a byte containing 8 bits.

### **Basic idea**

The sender:

1. Divides the data into symbols.
2. Adds redundant symbols.
3. Encodes them using a mathematical polynomial construction.
4. Sends the resulting codeword.

The receiver uses the redundant symbols to identify and correct supported errors.

### **Example**

Consider a data transmission system that sends a series of bytes.

Suppose a noisy channel corrupts several bytes.

A Reed-Solomon decoder can use the extra encoded symbols to reconstruct the original data, provided the error count is within the code's correction capacity.

### **Real-world applications**

Reed-Solomon coding has been used in:

- Optical discs, including CDs and DVDs.
- Data storage systems.
- Digital broadcasting.
- Communication links.
- Certain satellite and wireless systems.

### **Why is it useful?**

It is especially effective when errors occur in groups of bits or entire symbols.

For example, a burst error that damages several consecutive bits may corrupt multiple symbols. Reed-Solomon coding is well suited to handling symbol-level errors and can also use erasure information when the locations of missing or damaged symbols are known.

## **13. BCH Codes**

BCH stands for Bose–Chaudhuri–Hocquenghem.

BCH codes are a family of cyclic error-correcting codes.

They can be designed to correct a specified number of bit errors.

### **Basic idea**

A BCH encoder adds redundancy according to a generator polynomial.

The decoder calculates a syndrome and uses it to locate errors.

### **Applications**

BCH codes are used in areas such as:

- Flash memory.
- Digital communication.
- Storage devices.
- Hardware systems requiring predictable error-correction capability.

## **14. Hamming Codes in Hardware Communication**

Hamming codes are useful in digital hardware where individual bits must be protected against errors.

### **Example: RAM protection**

A memory system stores a data word together with parity bits.

When the word is read:

1. The memory controller reads the data and parity bits.
2. It recalculates the parity checks.
3. It determines whether a correctable error occurred.
4. It corrects the bit if the error is within the code's capability.

This is one form of **ECC memory**.

A common design is SECDED: Single Error Correction and Double Error Detection.

## **15. Comparison of Error-Correction Techniques**

| Technique        | What it operates on | Main capability                           | Typical application               |
| ---------------- | ------------------- | ----------------------------------------- | --------------------------------- |
| Hamming          | Bits                | Corrects a single-bit error in basic form | Memory, simple digital links      |
| Extended Hamming | Bits                | Corrects one bit and detects two          | ECC memory                        |
| BCH              | Bits                | Configurable multiple-bit correction      | Flash and communications          |
| Reed-Solomon     | Multi-bit symbols   | Corrects symbol errors                    | Storage and digital communication |

### **Important formulas to remember**

| Concept             | Formula                                         |
| ------------------- | ----------------------------------------------- |
| Hamming distance    | Number of differing bit positions               |
| Error detection     | \\(d\_{min}-1\\)                                |
| Error correction    | \\(\left\lfloor (d\_{min}-1)/2 \right\rfloor\\) |
| Hamming parity bits | \\(2^r\geq m+r+1\\)                             |

**Exam focus:** Practise parity calculations, two-dimensional parity, checksum addition, CRC division, Hamming-code construction, syndrome calculation and redundancy-bit problems.

# **2.4 Flow Control and Error Control**

## **1. Why is Flow Control Required?**

Flow control regulates the rate at which a sender sends data to a receiver.

Imagine:

- Sender transmits 100 frames per second.
- Receiver can process only 20 frames per second.

The receiver's buffer starts filling up.

```
Sender                         Receiver
  |                               |
  | Frame 1 --------------------> |
  | Frame 2 --------------------> |
  | Frame 3 --------------------> |  Processing slowly
  | Frame 4 --------------------> |
  | Frame 5 --------------------> |
  |                               |
  |                       Buffer fills
  |                               |
  |                       Frames may be lost
```

Flow control prevents the sender from overwhelming the receiver.

## **2. Stop-and-Wait Protocol**

Stop-and-Wait is the simplest flow-control protocol.

The sender transmits one frame and waits for permission to send the next one.

### **Working**

1. Sender transmits Frame 1.
2. Receiver receives Frame 1.
3. Receiver sends an acknowledgement.
4. Sender receives the acknowledgement.
5. Sender transmits Frame 2.

```
Sender                         Receiver
  |                               |
  |--------- Frame 1 -----------> |
  |                               |
  | <----------- ACK -------------|
  |                               |
  |--------- Frame 2 -----------> |
  |                               |
  | <----------- ACK -------------|
  |                               |
  |--------- Frame 3 -----------> |
```

### **Advantage**

- Very simple.
- Prevents receiver overload.

### **Disadvantage**

The sender wastes time waiting for acknowledgements.

When the propagation delay is large, link utilization can be very low.

## **3. Stop-and-Wait ARQ**

ARQ stands for **Automatic Repeat reQuest**.

Stop-and-Wait ARQ adds error recovery to the basic Stop-and-Wait flow-control protocol.

It uses:

- Frames
- ACKs
- Sequence numbers
- Timers
- Retransmissions

### **Why are sequence numbers required?**

Suppose the sender transmits Frame 0.

The receiver receives it and sends ACK 1.

But the ACK is lost.

The sender does not know that the frame was received successfully. It times out and sends the frame again.

Without sequence numbers, the receiver might accept the retransmitted frame as new data and deliver it twice.

To avoid duplicates, the sender and receiver use sequence numbers, commonly alternating between 0 and 1 in Stop-and-Wait.

### **Example**

```
Sender                         Receiver
  |                               |
  |-------- Frame 0 ------------> |
  |                               |
  | <----------- ACK 1 -----------|
  |                               |
  |-------- Frame 1 ------------> |
  |                               |
  | <----------- ACK 0 -----------|
```

The acknowledgement value commonly identifies the next expected sequence number.

### **What happens when a frame is lost?**

```
Sender                         Receiver
  |                               |
  |-------- Frame 0 -----X         |
  |                               |
  |          Timeout              |
  |                               |
  |-------- Frame 0 ------------> |
  |                               |
  | <----------- ACK 1 -----------|
```

The sender retransmits Frame 0 after the timer expires.

### **What happens when an ACK is lost?**

```
Sender                         Receiver
  |                               |
  |-------- Frame 0 ------------> |
  |                               |
  | <------ ACK 1 -------X         |
  |                               |
  |          Timeout              |
  |                               |
  |-------- Frame 0 ------------> |
  |                               |
  |  Receiver recognizes duplicate|
  |                               |
  | <----------- ACK 1 -----------|
```

The receiver recognizes that Frame 0 is a duplicate and does not deliver its payload again.

## **4. Go-Back-N ARQ**

Go-Back-N is a Sliding Window ARQ protocol.

The sender can transmit multiple frames without waiting for an individual acknowledgement after each frame.

The receiver generally accepts only the next expected frame in sequence.

### **Working**

Suppose the sender transmits frames:

```
0, 1, 2, 3, 4
```

Now suppose Frame 2 is lost.

```
Sender                         Receiver
  |                               |
  |--------- Frame 0 -----------> |
  |--------- Frame 1 -----------> |
  |--------- Frame 2 -----X        |
  |--------- Frame 3 -----------> |
  |--------- Frame 4 -----------> |
  |                               |
  |                   Frames 3 and 4
  |                   are discarded
  |                               |
  | <----- ACK for next expected 2 |
```

The receiver is waiting for Frame 2. It does not accept later frames as in-order data.

When the sender times out for the missing frame, it retransmits Frame 2 and all subsequent outstanding frames.

```
Retransmission:

Sender                         Receiver
  |                               |
  |--------- Frame 2 -----------> |
  |--------- Frame 3 -----------> |
  |--------- Frame 4 -----------> |
```

Hence the name **Go-Back-N**: the sender goes back to the missing frame and retransmits it and the later outstanding frames.

### **Advantages**

- More efficient than Stop-and-Wait when the link has significant delay.
- Receiver logic is relatively simple.
- Multiple frames can be in transit.

### **Disadvantages**

- Correctly received later frames may need to be retransmitted.
- A single lost frame can cause multiple retransmissions.

## **5. Selective Repeat ARQ**

Selective Repeat also uses a Sliding Window, but it is more efficient when some frames are lost.

The receiver can accept and buffer out-of-order frames.

Only missing or damaged frames are retransmitted.

### **Example**

Sender transmits:

```
0, 1, 2, 3, 4
```

Frame 2 is lost.

```
Sender                         Receiver
  |                               |
  |--------- Frame 0 -----------> |
  |--------- Frame 1 -----------> |
  |--------- Frame 2 -----X        |
  |--------- Frame 3 -----------> |
  |--------- Frame 4 -----------> |
  |                               |
  |                  Stores 3 and 4
  |                  in its buffer
  |                               |
  | <------- NAK / missing 2 ------|
  |                               |
  |--------- Frame 2 -----------> |
  |                               |
  |                  Reorders 0-4
  |                  and delivers data
```

Only Frame 2 needs retransmission.

### **Advantages**

- More efficient than Go-Back-N when losses are isolated.
- Avoids retransmitting correctly received frames.

### **Disadvantages**

- Requires receiver buffering.
- Requires more complex sequence-number and window management.

## **6. Sliding Window Protocol**

The Sliding Window Protocol allows the sender to transmit a group of frames before waiting for acknowledgements.

Imagine the sequence-number space is:

```
0 1 2 3 4 5 6 7 8 9 ...
```

Suppose the sender's window size is four.

It may initially transmit:

```
[0 1 2 3] 4 5 6 7 8 ...
```

After receiving acknowledgements, the window moves forward.

```
Before ACKs:

[0 1 2 3] 4 5 6 7 8

After acknowledging 0 and 1:

0 1 [2 3 4 5] 6 7 8
```

This is called a sliding window because the permitted range of sequence numbers moves forward as frames are acknowledged.

### **Sender window**

The sender window represents frames the sender is currently permitted to transmit or maintain as outstanding.

### **Receiver window**

The receiver window represents sequence numbers the receiver is willing to accept.

The two window sizes differ depending on the ARQ protocol.

| Protocol          | Sender window  | Receiver window |
| ----------------- | -------------- | --------------- |
| Stop-and-Wait ARQ | 1              | 1               |
| Go-Back-N         | Greater than 1 | 1               |
| Selective Repeat  | Greater than 1 | Greater than 1  |

## **7. Sequence Numbers**

Sequence numbers identify individual frames.

They help detect:

- Lost frames
- Duplicate frames
- Out-of-order frames
- Retransmissions

If a sequence-number field contains \\(m\\) bits, the number of possible sequence numbers is:

\\[ \boxed{2^m} \\]

### **Example**

For a 3-bit sequence number:

\\[ 2^3=8 \\]

Possible sequence numbers:

```
0, 1, 2, 3, 4, 5, 6, 7
```

After 7, numbering wraps around to 0.

This is called **modulo-8 sequence numbering**.

### **Window-size rules**

These are important for numerical and conceptual questions.

**Go-Back-N:**

For an \\(m\\)-bit sequence number:

\\[ \boxed{W_s \leq 2^m-1} \\]

The receiver window is normally:

\\[ \boxed{W_r=1} \\]

**Selective Repeat:**

To avoid ambiguity between old and new frames after sequence-number wraparound:

\\[ \boxed{W_s \leq 2^{m-1}} \\]

\\[ \boxed{W_r \leq 2^{m-1}} \\]

### **Example**

Suppose the sequence number has 3 bits.

Then:

\\[ 2^3=8 \\]

Maximum Go-Back-N sender window:

\\[ W_s\leq 7 \\]

Maximum Selective Repeat sender and receiver windows:

\\[ W_s,W_r\leq4 \\]

## **8. ACK and NAK**

### **ACK — Acknowledgement**

An ACK indicates successful receipt of data, or may indicate the next sequence number expected.

Example:

```
Receiver sends ACK 3
```

In a cumulative-acknowledgement protocol, this can mean that all frames up to sequence number 2 have been received in order and the receiver expects Frame 3 next.

The exact interpretation depends on the protocol.

### **NAK — Negative Acknowledgement**

A NAK indicates that a frame was damaged or that a particular frame is missing.

Example:

```
Receiver: Frame 4 is missing.
Receiver sends NAK 4.
```

The sender may retransmit Frame 4.

Not every ARQ protocol uses an explicit NAK. Some protocols rely on duplicate acknowledgements and timeouts.

## **9. Piggybacking**

Piggybacking is a technique used in bidirectional communication.

Instead of sending a separate ACK frame, the receiver includes acknowledgement information in an outgoing data frame.

### **Example**

Suppose A and B are communicating in both directions.

Without piggybacking:

```
A -------------------- Data --------------------> B
A <------------------- ACK ---------------------- B

A <------------------- Data ---------------------- B
A -------------------- ACK ----------------------> B
```

With piggybacking:

```
A -------------------- Data --------------------> B
A <------------------- Data + ACK ---------------- B
A -------------------- Data + ACK ----------------> B
```

### **Advantage**

It reduces the number of separate control frames.

### **Limitation**

If the receiver has no data to send, it cannot wait indefinitely to piggyback the acknowledgement. A timer or other acknowledgement rule is needed.

## **10. Timeout and Retransmission**

When a sender transmits a frame, it starts a timer.

If the expected acknowledgement does not arrive before the timeout, the sender may retransmit.

```
Sender                         Receiver
  |                               |
  |---------- Frame 5 ----------> |
  |                               |
  |          Start timer           |
  |                               |
  | <--------- ACK 6 -----X        |
  |                               |
  |          Timer expires        |
  |                               |
  |---------- Frame 5 ----------> |
```

Timeout values matter:

- Too short: unnecessary retransmissions.
- Too long: slow recovery after a lost frame.

In real communication systems, timer selection may depend on observed round-trip time and protocol-specific algorithms.

## **11. Flow Control vs Congestion Control**

These concepts are often confused.

**Flow control** protects the receiver from receiving more data than it can handle.

**Congestion control** protects the network from receiving more traffic than its resources can handle.

| Feature        | Flow control                             | Congestion control                    |
| -------------- | ---------------------------------------- | ------------------------------------- |
| Protects       | Receiver                                 | Network                               |
| Main problem   | Receiver cannot process data fast enough | Network is overloaded                 |
| Focus          | Sender-to-receiver rate                  | Traffic entering the network          |
| Example        | TCP receive window                       | TCP congestion window                 |
| Main objective | Prevent receiver-buffer overflow         | Prevent or relieve network congestion |

### **Real-world example**

Imagine a video server sending data to a laptop.

- The laptop has limited free receive-buffer space: this is a flow-control concern.
- The network path contains a congested router: this is a congestion-control concern.

Both conditions can limit how much data the sender should transmit.

# **2.5 Data Link Layer Protocols**

Two important protocols in this section are:

1. HDLC
2. PPP

Both provide framing and link-level communication functions, but they have different purposes and capabilities.

## **1. HDLC Protocol**

HDLC stands for **High-Level Data Link Control**.

It is a bit-oriented Data Link Layer protocol used for communication over point-to-point and multipoint links.

Its design includes:

- Frame delimiting
- Physical addressing in the HDLC frame format
- Error detection
- Flow and error control
- Numbered and unnumbered frame types

### **HDLC flag**

HDLC uses the following flag:

```
01111110
```

This sequence marks a frame boundary.

### **HDLC frame structure**

```
+------+---------+---------+-------------+------+------+
| Flag | Address | Control | Information | FCS  | Flag |
+------+---------+---------+-------------+------+------+
```

| Field       | Purpose                                               |
| ----------- | ----------------------------------------------------- |
| Flag        | Marks the beginning or end of a frame                 |
| Address     | Identifies the station or destination                 |
| Control     | Identifies frame type and carries control information |
| Information | Carries user data when present                        |
| FCS         | Frame Check Sequence for error detection              |

### **Bit stuffing in HDLC**

HDLC uses bit stuffing to prevent the payload from accidentally containing its flag sequence.

After five consecutive `1` bits in the frame contents, the sender inserts a `0`.

The receiver removes that inserted bit.

This ensures the flag pattern can be recognized reliably.

## **2. HDLC Frame Types**

HDLC has three major frame types.

### **A. I-Frame — Information Frame**

I-frames carry actual user data.

They also contain sequence and acknowledgement information.

Typical functions:

- Carry data.
- Identify the transmitted frame sequence number.
- Carry acknowledgement information.

Example:

```
A ---------------- I-frame carrying data ----------------> B
```

### **B. S-Frame — Supervisory Frame**

S-frames are used for flow control and error control.

They do not carry ordinary user data in their information field.

Common supervisory functions include:

| Frame | Meaning                                                       |
| ----- | ------------------------------------------------------------- |
| RR    | Receive Ready                                                 |
| RNR   | Receive Not Ready                                             |
| REJ   | Reject; requests retransmission beginning at a frame          |
| SREJ  | Selective Reject; requests retransmission of a specific frame |

### **C. U-Frame — Unnumbered Frame**

U-frames handle link-management and unnumbered control functions.

Examples include:

- Link setup.
- Link disconnection.
- Mode setting.
- Unnumbered information transfer.

### **Comparison**

| Frame   | Main purpose            | Carries user data?              |
| ------- | ----------------------- | ------------------------------- |
| I-frame | Data transfer           | Yes                             |
| S-frame | Supervision and control | No ordinary user payload        |
| U-frame | Link management         | Depends on the specific U-frame |

## **3. PPP Protocol**

PPP stands for **Point-to-Point Protocol**.

It is a Data Link Layer protocol designed to carry network-layer packets over point-to-point links.

Examples include certain serial and WAN links.

PPP provides:

- Framing
- Protocol identification
- Error detection using FCS
- Link establishment and configuration
- Network-layer protocol configuration

PPP does not provide a general MAC addressing system like Ethernet because it operates over a point-to-point link.

### **PPP frame structure**

PPP commonly uses an HDLC-like framing format:

```
+------+---------+---------+----------+-------------+------+
| Flag | Address | Control | Protocol | Information | FCS  |
+------+---------+---------+----------+-------------+------+
```

The closing flag follows the FCS.

| Field       | Typical size | Purpose                              |
| ----------- | ------------ | ------------------------------------ |
| Flag        | 1 byte       | Frame delimiter, `0x7E`              |
| Address     | 1 byte       | Commonly `0xFF`                      |
| Control     | 1 byte       | Commonly `0x03`                      |
| Protocol    | 1 or 2 bytes | Identifies the encapsulated protocol |
| Information | Variable     | Carries the packet                   |
| FCS         | 2 or 4 bytes | Error detection                      |

The PPP framing format and its flag, address, control and FCS fields are defined in RFC 1662.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



## **4. LCP — Link Control Protocol**

LCP stands for Link Control Protocol.

It is used to establish, configure, test and terminate a PPP link.

Its work includes:

1. Negotiating link options.
2. Checking whether the peer supports required options.
3. Testing the link.
4. Terminating the link when communication is finished.

### **Example**

Two routers establish a PPP connection.

```
Router A                     Router B
   |                             |
   |------ LCP Configure-Req --->|
   |<----- LCP Configure-Ack ----|
   |                             |
   |       Link established      |
```

The routers negotiate how the PPP link will operate.

## **5. NCP — Network Control Protocol**

NCP stands for Network Control Protocol.

It configures the Network Layer protocols that will use the PPP link.

Examples include:

- IPCP for IPv4 configuration.
- IPv6CP for IPv6 configuration.

### **Why are both LCP and NCP required?**

Think of a PPP connection as a link that needs two kinds of configuration:

- LCP configures the link itself.
- NCP configures network-layer protocols that will run over the link.

**PPP connection establishment**

1. **Physical connection**

   The point-to-point link becomes available.
2. **LCP negotiation**

   The two endpoints establish and configure the PPP link.
3. **Authentication, if configured**

   The peers authenticate using an agreed method.
4. **NCP negotiation**

   Network-layer protocol configuration is completed.
5. **Data transfer**

   Network-layer packets travel through PPP frames.

## **6. HDLC vs PPP**

| Feature                     | HDLC                              | PPP                                                   |
| --------------------------- | --------------------------------- | ----------------------------------------------------- |
| Full form                   | High-Level Data Link Control      | Point-to-Point Protocol                               |
| Main purpose                | Link-level framing and control    | Carry network-layer packets over point-to-point links |
| Orientation                 | Bit-oriented                      | HDLC-like framing                                     |
| Addressing                  | Uses an address field             | Commonly uses an all-stations address value           |
| Error detection             | FCS                               | FCS                                                   |
| Link configuration          | Depends on HDLC variant           | LCP                                                   |
| Network-layer configuration | Not PPP-style NCP                 | NCP                                                   |
| Authentication              | Depends on variant and deployment | Can support authentication mechanisms                 |

**Important:** PPP uses HDLC-like framing, but PPP is not the same as the full HDLC protocol. RFC 1662 explicitly describes PPP's HDLC-like framing rather than claiming that every HDLC feature is part of PPP.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.rfc-editor.org\&sz=32)

RFC Editor+1



# **2.6 Medium Access Control**

## **1. What is Medium Access Control?**

Imagine 20 devices connected to the same communication medium.

If all devices transmit at the same time, their transmissions may interfere.

Medium Access Control defines the rules for sharing a common communication channel.

It answers:

**Which device is allowed to transmit, and when?**

MAC protocols are broadly divided into:

1. Random access protocols.
2. Controlled access protocols.

## **2. Random Access Protocols**

In random access protocols, devices compete for access to the channel.

A device attempts transmission when it has data to send, following the protocol's rules.

Common examples:

- Pure ALOHA
- Slotted ALOHA
- CSMA
- CSMA/CD
- CSMA/CA

## **3. Pure ALOHA**

Pure ALOHA is a simple random access protocol.

A device transmits whenever it has data.

It does not first check whether another device is transmitting.

### **Working**

```
Device A:       Frame A -------------------->
Device B:             Frame B -------------->
Device C:                    Frame C ------->
```

If transmissions overlap, a collision occurs.

The affected frames may be corrupted and must be retransmitted.

### **Collision handling**

1. The device transmits.
2. It waits for an acknowledgement.
3. If no acknowledgement arrives, it assumes a collision or failure.
4. It waits for a random backoff period.
5. It retransmits.

### **Vulnerable period**

In Pure ALOHA, a frame is vulnerable to collision for a period equal to:

\\[ 2T \\]

where \\(T\\) is the frame transmission time.

This is because another frame can begin transmitting up to one frame time before or after the original frame begins.

### **Throughput**

The normalized throughput of Pure ALOHA is:

\\[ \boxed{S=Ge^{-2G}} \\]

where:

- \\(S\\) = throughput.
- \\(G\\) = average number of transmission attempts per frame time.

Maximum throughput:

\\[ \boxed{S\_{max}=\frac{1}{2e}\approx0.184} \\]

So the maximum theoretical channel efficiency is approximately 18.4%.

## **4. Slotted ALOHA**

Slotted ALOHA improves Pure ALOHA by dividing time into equal slots.

Devices are allowed to transmit only at the beginning of a slot.

```
Time →

+--------+--------+--------+--------+
| Slot 1 | Slot 2 | Slot 3 | Slot 4 |
+--------+--------+--------+--------+
```

If two devices transmit in the same slot, a collision occurs.

But transmissions beginning in different slots cannot overlap.

### **Vulnerable period**

For Slotted ALOHA:

\\[ T \\]

The vulnerable period is reduced from \\(2T\\) to \\(T\\).

### **Throughput**

\\[ \boxed{S=Ge^{-G}} \\]

Maximum throughput:

\\[ \boxed{S\_{max}=\frac{1}{e}\approx0.368} \\]

The theoretical maximum is approximately 36.8%.

### **Pure ALOHA vs Slotted ALOHA**

| Feature            | Pure ALOHA   | Slotted ALOHA       |
| ------------------ | ------------ | ------------------- |
| Transmission time  | Any time     | Beginning of a slot |
| Synchronization    | Not required | Required            |
| Vulnerable period  | \\(2T\\)     | \\(T\\)             |
| Maximum throughput | 18.4%        | 36.8%               |

## **5. CSMA — Carrier Sense Multiple Access**

CSMA improves on ALOHA by requiring a device to listen before transmitting.

**Carrier sense** means checking whether the communication channel is currently occupied.

### **Working**

1. Device checks the channel.
2. If the channel is idle, it may transmit.
3. If the channel is busy, it waits according to the protocol.
4. It attempts transmission later.

### **Why can collisions still happen?**

Consider two devices at opposite ends of a network.

```
Device A ------------------------ Device B
```

At almost the same time:

- A checks the channel and sees it is idle.
- B checks the channel and also sees it is idle.
- Both begin transmitting.
- Their signals meet and overlap.

This can happen because of propagation delay.

Therefore, listening before transmitting reduces collisions but does not eliminate them.

## **6. Types of CSMA**

### **A. 1-Persistent CSMA**

When the channel becomes idle, the device transmits immediately with probability 1.

If the channel is busy, it continuously senses it until it becomes idle.

**Advantage:** Low waiting time when the channel becomes free.

**Disadvantage:** Multiple waiting devices may transmit simultaneously when the channel becomes idle.

### **B. Non-Persistent CSMA**

If the channel is busy, the device waits for a random period before sensing again.

**Advantage:** Can reduce repeated collisions.

**Disadvantage:** Can introduce additional delay and leave the channel idle unnecessarily.

### **C. p-Persistent CSMA**

Typically used with slotted channels.

When the channel is idle:

- Transmit with probability \\(p\\).
- Defer to the next slot with probability \\(1-p\\).

It balances immediate transmission against collision risk.

## **7. CSMA/CD**

CSMA/CD stands for **Carrier Sense Multiple Access with Collision Detection**.

It was used in traditional shared, half-duplex Ethernet.

### **How it works**

1. Listen to the channel.
2. If idle, transmit.
3. While transmitting, monitor the medium.
4. If a collision is detected, stop transmission.
5. Send a jam signal.
6. Wait for a random backoff interval.
7. Retransmit.

### **Diagram**

```
Device A                         Device B
   |                                |
   |------ Start transmission       |
   |                                |
   |                  Start transmission
   |                                |
   |<--------- Collision ---------->|
   |                                |
   |-------- Jam signal ------------|
   |                                |
   |         Random backoff         |
   |                                |
   |-------- Retransmit ------------|
```

### **Binary Exponential Backoff**

After a collision, a device waits for a random number of time slots.

The range increases after repeated collisions.

For the first collision, it chooses from:

```
0 or 1 slots
```

For the second:

```
0 to 3 slots
```

For the third:

```
0 to 7 slots
```

The backoff range is capped in the protocol.

The goal is to reduce the chance that colliding devices retransmit at exactly the same time again.

### **Important modern distinction**

Traditional shared half-duplex Ethernet used CSMA/CD.

Modern switched full-duplex Ethernet links do not have the same shared-medium collision problem and do not use CSMA/CD for normal frame transmission.

## **8. CSMA/CA**

CSMA/CA stands for **Carrier Sense Multiple Access with Collision Avoidance**.

It is used in Wi-Fi networks.

Unlike traditional half-duplex Ethernet collision detection, a wireless transmitter generally cannot reliably detect a collision while transmitting.

Instead, Wi-Fi uses techniques to reduce the probability of collisions.

### **Working**

1. A station senses the wireless channel.
2. If the channel is idle, it waits for a required inter-frame period.
3. It chooses a random backoff.
4. The backoff counter decreases while the channel remains idle.
5. If the channel becomes busy, the counter pauses.
6. When the counter reaches zero, the station transmits.
7. The receiver returns an ACK if the frame is received successfully.

```
Wi-Fi station                  Access Point
     |                              |
     |--- Sense channel             |
     |--- Wait / random backoff     |
     |--- Data frame -------------->|
     |                              |
     |<----------------------- ACK --|
```

### **RTS/CTS**

Wi-Fi can also use RTS/CTS, particularly where it helps with hidden-terminal problems.

RTS = Request to Send.

CTS = Clear to Send.

```
Station A                       Access Point
   |                                |
   |--------- RTS ----------------->|
   |<-------- CTS ------------------|
   |                                |
   |--------- DATA ---------------->|
   |<-------- ACK ------------------|
```

Other stations that hear the reservation information defer their transmissions for the indicated duration.

## **9. Controlled Access Protocols**

In controlled access protocols, stations coordinate access instead of competing freely.

Common methods include:

- Polling
- Token passing

### **A. Polling**

A central controller asks devices one by one whether they want to transmit.

```
             Controller
             /   |   \
            /    |    \
          A      B     C
```

Working:

1. Controller polls Device A.
2. A transmits if it has data.
3. Controller polls Device B.
4. B transmits if it has data.
5. The sequence continues.

**Advantage:** Avoids collisions when implemented correctly.

**Disadvantages:**

- Polling overhead.
- A central controller can become a single point of failure.
- Devices may wait while other devices are being polled.

### **B. Token Passing**

A special control frame called a token circulates among devices.

Only the device holding the token can transmit.

```
      A --------> B
      ^           |
      |           v
      D <---------C
```

Suppose Device B has the token.

It can transmit data. After completing its allowed transmission, it passes the token to Device C.

**Advantage:** Controlled, predictable access without ordinary contention collisions.

**Disadvantage:** Token-management overhead and potential problems if the token is lost or duplicated.

## **10. Collision and Collision Domains**

A collision occurs when two or more transmissions interfere with one another on a shared half-duplex communication medium.

A **collision domain** is the part of a network where simultaneous transmissions can collide.

### **Hub example**

A hub repeats received electrical signals to its other ports.

All connected devices share the same collision domain.

```
          PC A
            |
            |
PC B ------ HUB ------ PC C
            |
            |
          PC D
```

All four PCs share the same collision domain.

A switch is different. It forwards frames to specific ports and gives each full-duplex port a separate collision-free link.

**Important distinction:**

- Collision domains are primarily a concern with shared half-duplex media.
- Broadcast domains describe which interfaces receive a Layer 2 broadcast.
- Routers and VLAN boundaries separate broadcast domains.

# **2.7 Ethernet and Switching**

## **1. What is Ethernet?**

Ethernet is a family of technologies used to create wired local area networks.

It is standardized primarily under the IEEE 802.3 family.

Ethernet defines functions such as:

- Frame format
- MAC addressing
- Physical-layer signalling options
- Link speeds
- Media requirements

Ethernet is widely used in:

- Homes
- Offices
- Data centres
- Campus networks
- Industrial networks

## **2. Ethernet Architecture**

A simple Ethernet LAN looks like this:

```
                Router
                  |
                  |
              Ethernet
               Switch
          /       |       \
         /        |        \
       PC 1      PC 2     Server
```

Each endpoint connects to the switch through an Ethernet link.

When PC 1 sends a frame to the server:

1. It puts the destination MAC address in the Ethernet header.
2. The switch receives the frame.
3. The switch examines the source MAC address and learns where PC 1 is connected.
4. It examines the destination MAC address.
5. It forwards the frame out the appropriate port if the destination is known.

## **3. Ethernet Frame Format**

The Ethernet frame is one of the most important diagrams in this unit.

```
+----------+-----+-----------+-----------+---------------+-----------+------+
| Preamble | SFD | Dest MAC  | Src MAC   | Type/Length   | Payload   | FCS  |
+----------+-----+-----------+-----------+---------------+-----------+------+
| 7 bytes  | 1 B | 6 bytes   | 6 bytes   | 2 bytes       | 46-1500 B | 4 B  |
+----------+-----+-----------+-----------+---------------+-----------+------+
```

### **A. Preamble**

Size: 7 bytes.

It helps the receiving interface synchronize with the incoming signal.

### **B. SFD — Start Frame Delimiter**

Size: 1 byte.

It marks the end of the preamble and the beginning of the frame's MAC header.

### **C. Destination MAC**

Size: 6 bytes.

Identifies the destination interface or group.

### **D. Source MAC**

Size: 6 bytes.

Identifies the source interface.

### **E. Type/Length**

Size: 2 bytes.

It can identify the encapsulated protocol or indicate a payload length, depending on the value and Ethernet format.

Common EtherType values include:

| Protocol | EtherType |
| -------- | --------- |
| IPv4     | `0x0800`  |
| ARP      | `0x0806`  |
| IPv6     | \`0x86DD  |

### **F. Payload**

Carries the encapsulated information, usually an upper-layer packet.

Traditional Ethernet supports a minimum payload size of 46 bytes and a maximum of 1500 bytes for the common Ethernet MAC frame format. Padding is used when necessary to reach the minimum frame size.

### **G. FCS — Frame Check Sequence**

Size: 4 bytes.

Contains a CRC-32 value used for error detection.

The receiver checks the FCS to determine whether the frame was corrupted.

### **Ethernet frame size**

The standard Ethernet MAC frame size is commonly:

- Minimum: 64 bytes.
- Maximum: 1518 bytes without an 802.1Q VLAN tag.

These counts include the destination MAC through the FCS, but exclude the preamble and SFD.

An 802.1Q VLAN tag adds four bytes to the frame, giving 1522 bytes for the usual tagged maximum frame size.

## **4. Ethernet MAC Address**

An Ethernet MAC address is a 48-bit address.

Example:

```
A4:5E:60:12:34:56
```

It is generally represented as six hexadecimal octets.

The switch uses destination MAC addresses to determine where frames should be forwarded.

### **MAC address types**

| Type      | Purpose                           |
| --------- | --------------------------------- |
| Unicast   | One destination                   |
| Multicast | A group of destinations           |
| Broadcast | All devices in a broadcast domain |

Ethernet's broadcast address:

```
FF:FF:FF:FF:FF:FF
```

## **5. Ethernet Speed Standards**

Ethernet has evolved from relatively slow shared networks to extremely high-speed switched links.

| Ethernet family      | Nominal speed | Common name |
| -------------------- | ------------- | ----------- |
| Ethernet             | 10 Mbps       | 10BASE-T    |
| Fast Ethernet        | 100 Mbps      | 100BASE-TX  |
| Gigabit Ethernet     | 1 Gbps        | 1000BASE-T  |
| 2.5 Gigabit Ethernet | 2.5 Gbps      | 2.5GBASE-T  |
| 5 Gigabit Ethernet   | 5 Gbps        | 5GBASE-T    |
| 10 Gigabit Ethernet  | 10 Gbps       | 10GBASE-T   |
| 25 Gigabit Ethernet  | 25 Gbps       | 25GbE       |
| 40 Gigabit Ethernet  | 40 Gbps       | 40GbE       |
| 100 Gigabit Ethernet | 100 Gbps      | 100GbE      |
| 400 Gigabit Ethernet | 400 Gbps      | 400GbE      |

Different variants use different copper, fibre, connector, signalling and distance specifications.

For example, 1000BASE-T uses twisted-pair copper cabling, while many higher-speed data-centre links use optical transceivers and fibre.

## **6. Full-Duplex vs Half-Duplex Ethernet**

### **Half-duplex**

A device can transmit or receive at a given time, but not both simultaneously.

Traditional shared Ethernet used this arrangement.

Collisions were possible, so CSMA/CD was required.

### **Full-duplex**

A device can transmit and receive simultaneously.

Modern switched Ethernet links normally use full-duplex operation.

```
Half-duplex:

Device A  <------>  Device B
One direction at a time

Full-duplex:

Device A  -------->  Device B
Device A  <--------  Device B
Simultaneous transmission and reception
```

### **Comparison**

| Feature                             | Half-duplex                         | Full-duplex                                        |
| ----------------------------------- | ----------------------------------- | -------------------------------------------------- |
| Transmit and receive simultaneously | No                                  | Yes                                                |
| Collisions on a shared link         | Possible                            | Not in normal point-to-point full-duplex operation |
| CSMA/CD                             | Used in traditional shared Ethernet | Not required                                       |
| Performance                         | Lower in shared environments        | Better for modern switched links                   |

## **7. Ethernet Switches**

An Ethernet switch operates primarily at the Data Link Layer.

It receives Ethernet frames and forwards them according to its MAC address table.

### **Example network**

```
            Ethernet Switch
           /       |        \
          /        |         \
        PC A      PC B      Server
       Port 1    Port 2      Port 3
```

Suppose PC A wants to send data to the server.

The switch examines the destination MAC and determines which port leads to the server.

If it already knows the destination port, it forwards the frame there instead of sending it everywhere.

## **8. MAC Address Table**

A switch maintains a MAC address table.

A simplified example:

| MAC address         | Port   |
| ------------------- | ------ |
| `AA:AA:AA:AA:AA:AA` | Port 1 |
| `BB:BB:BB:BB:BB:BB` | Port 2 |
| `CC:CC:CC:CC:CC:CC` | Port 3 |

The table tells the switch which port is associated with each learned MAC address.

### **How does a switch learn addresses?**

Suppose PC A sends a frame to the server.

```
PC A                         Switch
 |                               |
 |--------- Frame -------------->|
 | Source MAC = AA:AA:AA:AA:AA:AA
 | Destination MAC = CC:CC:CC:CC:CC:CC
```

The switch learns:

```
AA:AA:AA:AA:AA:AA → Port 1
```

It learns the **source MAC address**, not the destination address, from the received frame.

It then looks up the destination MAC.

This is called MAC address learning.

## **9. Switching: Forwarding, Flooding and Filtering**

Switches typically use three basic decisions.

### **A. Forwarding**

When the destination MAC address is known and is on another port, the switch forwards the frame only through that port.

Example:

```
PC A → Switch → Server
```

Only the server's associated port receives the unicast frame.

### **B. Flooding**

When a destination MAC address is unknown, the switch floods the frame out appropriate ports except the port on which it arrived.

Broadcast frames are also flooded within their VLAN, subject to normal filtering rules.

```
               Switch
             /   |   \
            /    |    \
          PC A  PC B  Server
```

If a frame from PC A has an unknown destination MAC, the switch may send it to PC B and the server.

### **C. Filtering**

If the destination is known to be on the same port the frame arrived on, the switch does not forward the frame to other ports.

This prevents unnecessary traffic.

### **Summary**

| Operation  | Meaning                                                  |
| ---------- | -------------------------------------------------------- |
| Forwarding | Send frame to the correct destination port               |
| Flooding   | Send frame to multiple appropriate ports                 |
| Filtering  | Do not forward the frame where forwarding is unnecessary |

## **10. Bridge Operation**

A bridge connects multiple LAN segments and forwards frames based on MAC addresses.

It performs the basic functions of a Layer 2 switch:

1. Learn source MAC addresses.
2. Associate addresses with ports.
3. Forward known unicast frames.
4. Flood unknown unicast and broadcasts appropriately.
5. Filter frames that do not need to cross the bridge.

Modern Ethernet switches can be understood as multiport bridges with hardware-accelerated forwarding.

## **11. Switch vs Bridge vs Hub**

| Feature              | Hub                | Bridge                                   | Switch                    |
| -------------------- | ------------------ | ---------------------------------------- | ------------------------- |
| OSI layer            | Physical           | Data Link                                | Data Link                 |
| Forwarding basis     | Repeats signals    | MAC address                              | MAC address               |
| Learns MAC addresses | No                 | Yes                                      | Yes                       |
| Ports                | Multiple           | Usually few                              | Multiple                  |
| Collision domains    | One shared domain  | Separated by bridge ports                | Separated by switch ports |
| Broadcast handling   | Repeats broadcasts | Forwards broadcasts within connected LAN | Floods within the VLAN    |

### **Real-world example**

A hub repeats all incoming signals, so every connected device shares the same collision domain.

A switch creates separate point-to-point links for its ports and forwards frames selectively.

This gives better performance and reduces unnecessary traffic.

## **12. Collision Domain vs Broadcast Domain**

This distinction is essential.

### **Collision domain**

The region in which transmissions can collide on a shared medium.

A switch separates collision domains between its ports in ordinary switched Ethernet.

### **Broadcast domain**

The group of devices that receive a Layer 2 broadcast frame.

A normal Layer 2 switch forwards broadcasts throughout a VLAN.

Routers and VLAN boundaries separate broadcast domains.

```
                 Router
               /        \
              /          \
          Switch A      Switch B
         /      \       /       \
       PC 1    PC 2   PC 3     PC 4

Separate routed networks / broadcast domains
```

Within each switched VLAN, broadcasts are distributed across that VLAN's member ports.

## **13. VLAN — Virtual LAN**

VLAN stands for Virtual Local Area Network.

It logically divides a physical switched network into multiple Layer 2 broadcast domains.

### **Why VLANs are required**

Imagine an organization with:

- HR department.
- Engineering department.
- Finance department.

All employees may be connected to the same physical switch infrastructure.

But the organization may want each department in a separate logical network.

VLANs allow this separation without requiring physically separate switches for every group.

### **Diagram**

```
             Physical Switch
           +-------------------+
           |                   |
           | VLAN 10           |
           | HR                |
           |                   |
           | VLAN 20           |
           | Engineering       |
           |                   |
           | VLAN 30           |
           | Finance           |
           +-------------------+
```

The devices in different VLANs are in different Layer 2 broadcast domains.

Communication between VLANs generally requires a Layer 3 routing function.

### **Access port**

An access port typically carries traffic for one VLAN.

Example:

```
PC of HR → Switch Port 1 → VLAN 10
```

### **Trunk port**

A trunk carries traffic for multiple VLANs between network devices.

It typically uses 802.1Q tagging to identify the VLAN to which each frame belongs.

```
Switch A ================= Switch B
          Trunk link
      VLAN 10, 20 and 30
```

## **14. IEEE 802.1Q VLAN Tagging**

IEEE 802.1Q defines a VLAN tagging method for Ethernet frames.

A four-byte VLAN tag is inserted into the Ethernet frame after the source MAC address and before the Type/Length field.

### **Tagged Ethernet frame**

```
+---------+---------+----------+-----------+-------------+---------+
| Dest MAC| Src MAC | VLAN Tag | Type/Length| Payload    | FCS     |
+---------+---------+----------+-----------+-------------+---------+
```

The 802.1Q tag contains:

- TPID: Tag Protocol Identifier, commonly `0x8100`.
- PCP: Priority Code Point, 3 bits.
- DEI: Drop Eligible Indicator, 1 bit.
- VID: VLAN Identifier, 12 bits.

The VID field identifies the VLAN.

A 12-bit VID has 4096 possible values, although not all values are available for ordinary VLAN assignment.

**Real-world example:** A trunk link between two switches can carry frames for VLAN 10, VLAN 20 and VLAN 30. The VLAN tag tells the receiving switch which VLAN each tagged frame belongs to.

## **15. Spanning Tree Protocol (STP)**

STP is used to prevent Layer 2 switching loops.

### **Why are loops dangerous?**

Consider a network with redundant links:

```
       Switch A
       /      \
      /        \
 Switch B ---- Switch C
```

The redundant physical paths improve availability.

However, Ethernet broadcasts and unknown unicast frames can circulate indefinitely if Layer 2 loops are not controlled.

This can cause:

- Broadcast storms.
- Repeated frame circulation.
- MAC address-table instability.
- Severe network disruption.

### **How STP solves the problem**

STP logically blocks some redundant paths while keeping them available as backups.

It creates a loop-free logical topology.

### **Basic STP operation**

1. Switches exchange Bridge Protocol Data Units (BPDUs).
2. They elect a Root Bridge.
3. Each non-root switch selects a Root Port.
4. Designated Ports are selected for network segments.
5. Redundant paths are placed into a blocking state.

```
             Root Bridge
             Switch A
             /      \
            /        \
       Switch B      Switch C
            \        /
             \      /
             Blocked
            redundant path
```

The blocked path can be activated if the active topology changes, depending on the STP variant and convergence process.

### **Root Bridge**

The Root Bridge is the central reference point for STP path calculations.

The switch with the best Bridge ID becomes the root.

The Bridge ID contains a priority component and a MAC address component.

### **Root Port**

Each non-root switch chooses one Root Port: the port providing the lowest-cost path toward the Root Bridge.

### **Designated Port**

A designated port is selected for a network segment to provide the best path toward the Root Bridge for that segment.

Other redundant ports may be blocked to prevent loops.

### **STP vs RSTP**

RSTP stands for Rapid Spanning Tree Protocol.

It improves convergence time compared with classic STP.

| Feature       | STP                   | RSTP                                                    |
| ------------- | --------------------- | ------------------------------------------------------- |
| Standard      | IEEE 802.1D family    | IEEE 802.1w, incorporated into later bridging standards |
| Convergence   | Generally slower      | Generally faster                                        |
| Purpose       | Prevent Layer 2 loops | Prevent loops with faster recovery                      |
| Basic concept | Loop-free topology    | Loop-free topology                                      |

**Key point:** VLANs divide a LAN into broadcast domains. STP prevents loops in bridged networks. They solve different problems, but are commonly used together.

# **2.8 High-Speed Interconnects for VLSI Systems**

Hardware-oriented section

This section is different from traditional computer networking.

The previous sections focused on communication between computers and network devices.

This section focuses on communication **inside electronic systems**, including communication between chips, processors, memory controllers, cameras, displays and other hardware components.

VLSI stands for **Very Large Scale Integration**. It refers to integrating a large number of electronic components, especially transistors, into a chip.

## **1. What is a High-Speed Interconnect?**

An interconnect is the communication path that transfers signals between electronic components.

Examples:

- CPU to memory.
- Camera sensor to processor.
- Processor to display.
- Chip to chip.
- FPGA to external interface.
- Memory controller to storage interface.

### **Example: Smartphone camera**

A smartphone camera sensor captures image data.

That data must travel to the application processor.

```
+-------------+       High-speed       +------------------+
| Camera      | =====================> | Application      |
| Sensor      |  Serial data lanes     | Processor        |
+-------------+                       +------------------+
```

The connection must carry a large amount of data with low delay and controlled power consumption.

Ordinary low-speed interfaces would not provide adequate throughput for high-resolution, high-frame-rate image data.

This is why high-speed serial interconnect technologies are used.

## **2. Parallel vs Serial Communication**

### **A. Parallel Communication**

Parallel communication sends multiple bits simultaneously over separate data wires.

Suppose we want to transfer 8 bits:

```
Transmitter                     Receiver

D0 ----------------------------> D0
D1 ----------------------------> D1
D2 ----------------------------> D2
D3 ----------------------------> D3
D4 ----------------------------> D4
D5 ----------------------------> D5
D6 ----------------------------> D6
D7 ----------------------------> D7
```

All eight bits can travel during the same transfer interval.

### **Advantages**

- Multiple bits transferred simultaneously.
- Simple concept.
- Can provide high throughput over short distances at suitable rates.

### **Disadvantages**

At high speeds, multiple parallel wires create serious signal-integrity challenges.

For example:

- Clock skew.
- Timing mismatch between wires.
- Crosstalk.
- Routing complexity.
- Increased pin count.

### **B. Serial Communication**

Serial communication sends bits sequentially over one or more serial lanes.

```
Transmitter                       Receiver

Serial lane ====================> Serial lane

Bits: 1 → 0 → 1 → 1 → 0 → 0 → 1
```

### **Advantages**

- Fewer wires or pins for a given connection.
- Easier routing for many high-speed interconnect designs.
- Can support very high data rates.
- Often reduces power and physical interface complexity compared with very wide parallel buses.

### **Disadvantages**

- Requires serialization and deserialization.
- Clock recovery or clock distribution may be required.
- High-speed signal integrity becomes important.
- The receiver may need equalization and clock-recovery circuits.

### **Comparison**

| Feature              | Parallel                      | Serial                                  |
| -------------------- | ----------------------------- | --------------------------------------- |
| Data transfer        | Multiple bits at once         | Bits sent sequentially per lane         |
| Number of data wires | More                          | Fewer                                   |
| Routing              | More complex for wide buses   | Often simpler                           |
| Clock skew           | Significant concern           | Depends on clocking architecture        |
| High-speed design    | Challenging across many lanes | Commonly preferred for high-speed links |

**Important:** Serial does not automatically mean faster. A well-designed parallel bus can be very fast. Serial interfaces are often chosen because they scale more conveniently at high data rates and over practical board or cable distances.

## **3. SerDes — Serializer / Deserializer**

SerDes is a circuit architecture that converts parallel data into serial data and converts received serial data back into parallel data.

It is widely used in high-speed communication interfaces.

### **Why do we need SerDes?**

Suppose a chip internally produces 16 bits of data at a time.

Sending 16 bits across 16 separate wires can be expensive and difficult at high speeds.

Instead, the transmitter converts the parallel data into a serial bit stream.

At the receiver, another circuit reconstructs parallel data.

### **Working**

**SerDes architecture**

**Parallel input data**

# **1011 0101**

**Serializer**

Converts parallel data into serial bits

**High-speed serial link**

# **1 → 0 → 1 → 1 → 0 → 1 → 0 → 1**

**Deserializer**

Reconstructs parallel data

**Parallel output data**

# **1011 0101**

### **Example**

A chip has a 16-bit parallel interface operating at a clock frequency of 100 MHz.

Its theoretical raw data rate is:

\\[ 16 \times 100\\,MHz \\]

\\[ =1600\\,Mbps \\]

\\[ =1.6\\,Gbps \\]

A SerDes link can serialize this data onto fewer high-speed lanes.

The actual serial line rate depends on encoding overhead, framing, protocol requirements and implementation details.

### **Real-world applications**

SerDes is found in many high-speed interfaces, including:

- PCI Express.
- SATA.
- USB high-speed implementations.
- Ethernet physical-layer implementations.
- Fibre Channel.
- Chip-to-chip links.
- Automotive camera links.

## **4. LVDS — Low-Voltage Differential Signaling**

LVDS is a low-voltage differential signalling method.

Instead of sending a signal using one wire referenced to ground, it uses a pair of wires carrying complementary signals.

### **Basic idea**

```
             Transmitter                 Receiver

             +---------+                 +---------+
Data bit --->|         |==== Wire P ====>|         |
             | Driver  |                 | Receiver|
             |         |==== Wire N ====>|         |
             +---------+                 +---------+
```

The receiver looks at the voltage difference between the two wires.

\\[ V\_{diff}=V_P-V_N \\]

The logical value is determined by the polarity of this voltage difference.

### **Why differential signaling?**

A disturbance may affect both wires in a similar way.

The receiver measures their difference, so some of that common-mode noise cancels out.

This improves noise immunity.

### **Advantages**

- Good noise immunity.
- Low voltage swing.
- High-speed operation.
- Reduced electromagnetic interference compared with many single-ended arrangements.
- Suitable for board-level high-speed links.

### **Real-world examples**

LVDS has been used in:

- Display interfaces.
- Camera interfaces.
- Industrial equipment.
- High-speed data acquisition.
- Board-to-board communication.

## **5. MIPI D-PHY**

MIPI D-PHY is a high-speed physical-layer interface commonly used to connect camera sensors and displays to application processors.

MIPI is an industry organisation that develops interface specifications for mobile and related devices.

D-PHY is commonly used with:

- MIPI CSI-2 for camera data.
- MIPI DSI-2 for display data.

### **Smartphone camera example**

```
+----------------+       MIPI D-PHY       +------------------+
| Camera Sensor  | =====================> | Application      |
|                |  High-speed lanes      | Processor        |
+----------------+                        +------------------+
```

When a smartphone camera captures an image:

1. The sensor produces pixel data.
2. The interface formats and transmits that data.
3. D-PHY carries the high-speed electrical signals.
4. The application processor receives the data.
5. The image-processing pipeline processes the image.

### **D-PHY lanes**

A typical D-PHY arrangement uses:

- One clock lane.
- One or more data lanes.

The data lanes allow the interface to scale aggregate bandwidth.

D-PHY supports both high-speed and low-power operating states, which helps balance data throughput and power usage.

### **Real-world importance**

D-PHY is widely used in cameras and displays in smartphones, tablets, wearables, automotive systems and other embedded devices. MIPI describes it as a synchronous physical layer with a clock lane and scalable data lanes.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.mipi.org\&sz=32)

MIPI+1



**Note on versions:** MIPI lists D-PHY v3.6, released in September 2025, as its current version on the specification page. The D-PHY specification itself is available through MIPI Alliance membership.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.mipi.org\&sz=32)

MIPI



## **6. CML — Current Mode Logic**

CML stands for Current Mode Logic.

It is a high-speed signalling technique used in circuits such as:

- Clock distribution.
- High-speed serial transceivers.
- SerDes blocks.
- Optical communication interfaces.
- Chip-to-chip links.

### **Basic concept**

CML uses a controlled current through a differential output structure.

A simplified conceptual diagram:

```
             Transmitter
          +----------------+
          |                |
          | Differential   |
          | current driver |
          +----------------+
               |      |
               |      |
             OUT+    OUT-
               |      |
               +------+
                  |
             Differential
             receiver
```

The circuit steers current between output paths to create a differential voltage across the receiver's termination.

### **Why is CML useful?**

- Supports high-speed switching.
- Can provide controlled signal transitions.
- Works well in high-frequency, high-speed circuit designs.
- Commonly used in serial transceiver implementations.

### **LVDS vs CML**

| Feature        | LVDS                                        | CML                                              |
| -------------- | ------------------------------------------- | ------------------------------------------------ |
| Full form      | Low-Voltage Differential Signaling          | Current Mode Logic                               |
| Signalling     | Differential voltage                        | Current-steering differential circuit            |
| Typical use    | Data and clock links                        | High-speed clocks and serial links               |
| Design concern | Differential voltage levels and termination | Bias current, termination and common-mode levels |

## **7. Differential Signaling**

Differential signaling is a method in which information is represented by the voltage difference between two conductors.

### **Single-ended signalling**

```
Signal wire -----------------> Receiver
Ground     ------------------> Reference
```

The receiver measures one signal relative to a reference.

### **Differential signalling**

```
Positive wire ----------------> Receiver
Negative wire ----------------> Receiver
```

The receiver measures the difference.

\\[ V\_{diff}=V_P-V_N \\]

### **Why is it effective?**

Suppose noise adds the same voltage to both wires.

```
Before noise:

Wire P = 1.0 V
Wire N = 0.8 V

Difference = 0.2 V
```

After common-mode noise of 0.1 V:

```
Wire P = 1.1 V
Wire N = 0.9 V

Difference = 0.2 V
```

The difference remains unchanged.

This is why differential signalling provides good immunity to common-mode noise.

### **Important terms**

**Common-mode voltage:** The average voltage of the two conductors.

\\[ V\_{cm}=\frac{V_P+V_N}{2} \\]

**Differential voltage:** The voltage difference between the conductors.

\\[ V\_{diff}=V_P-V_N \\]

**Differential impedance:** The impedance seen between the two conductors, an important design parameter for high-speed links.

## **8. High-Speed Serial Links**

A high-speed serial link transfers data at a high bit rate through one or more serial channels.

Examples:

- PCI Express.
- SATA.
- USB.
- High-speed Ethernet.
- MIPI interfaces.
- Chip-to-chip SerDes links.

### **Typical link architecture**

```
+----------------+                       +----------------+
| Transmitter    |                       | Receiver       |
|                |                       |                |
| Encoder        |                       | Decoder        |
| Serializer     |                       | Deserializer   |
| Driver         |                       | Equalizer      |
+-------+--------+                       +-------+--------+
        |                                        ^
        |                                        |
        +========== Serial channel ==============+
```

The exact blocks differ from one interface to another.

### **Common challenges**

- Signal attenuation.
- Reflections.
- Crosstalk.
- Jitter.
- Electromagnetic interference.
- Clock recovery.
- Equalization.
- Bit errors.

## **9. Clock and Data Recovery (CDR)**

CDR stands for Clock and Data Recovery.

In many high-speed serial links, the receiver must determine the timing of incoming bits.

The transmitter sends a serial data stream. The receiver gets the electrical signal, but it may not receive a separate clock wire.

CDR estimates and reconstructs the sampling clock from transitions in the received data.

### **Example**

Suppose the received data is:

```
1 0 1 1 0 0 1 0
```

The receiver needs to know the appropriate times to sample each bit.

If it samples too early or too late, it may read the wrong value.

CDR tracks the timing and helps choose the correct sampling instant.

### **Simplified diagram**

```
Transmitter                     Receiver

Data stream ===================> Signal
                                 |
                                 v
                            +---------+
                            |   CDR   |
                            +---------+
                                 |
                                 v
                         Recovered clock
                                 |
                                 v
                         Correct data sampling
```

### **Why is CDR important?**

It helps the receiver maintain synchronization with the incoming data stream.

Without accurate timing recovery, even a clean electrical signal can be sampled incorrectly.

Some interfaces transmit a separate clock, while others embed timing information in the encoded data stream.

## **10. Signal Integrity Basics**

Signal integrity refers to preserving the quality and timing of electrical signals as they travel through an interconnect.

At low frequencies, a wire may behave approximately like an ordinary connection.

At high speeds, the electrical behaviour of the connection matters greatly.

A PCB trace, connector, cable and package can affect the signal.

### **A. Attenuation**

Attenuation means signal strength decreases as the signal travels.

```
Transmitter                         Receiver

Strong signal --------------------> Weaker signal
```

The receiver may struggle to distinguish different bit values when the received signal becomes too weak.

### **B. Reflection**

Reflection occurs when a signal encounters an impedance discontinuity.

For example, an improperly terminated transmission line can reflect part of the signal back toward the transmitter.

This can distort the waveform.

### **C. Crosstalk**

Crosstalk occurs when a signal on one conductor affects another nearby conductor.

```
Signal line 1  =====================>
                       |
                       | Unwanted coupling
                       v
Signal line 2  =====================>
```

It can introduce unwanted noise into a neighbouring signal.

### **D. Jitter**

Jitter is variation in the timing of signal transitions.

Ideally, transitions occur at perfectly regular time intervals.

In real systems, the timing can vary.

```
Ideal:

|----|----|----|----|----|

With jitter:

|--- |-----|--- |------|
```

Too much jitter can make the receiver sample at the wrong time.

### **E. Eye Diagram**

An eye diagram is a visualization used to evaluate the quality of a high-speed digital signal.

It is generated by superimposing many bit periods of a signal.

**Conceptual eye diagram**

Conceptual illustration only. A real eye diagram contains many overlaid signal transitions.

A large, open eye generally indicates better signal quality and a wider sampling margin.

A closed or narrow eye suggests more noise, distortion or timing uncertainty.

### **F. Termination**

Termination means matching the electrical characteristics of the interconnect to reduce signal reflections.

Correct termination is critical for many high-speed links.

## **11. Putting It All Together: Camera-to-Processor Example**

Consider a smartphone camera.

```
+--------------+      +------------------+      +------------------+
| Camera Sensor| ---> | MIPI D-PHY       | ---> | Application      |
|              |      | Serial Interface |      | Processor        |
+--------------+      +------------------+      +------------------+
                              |
                              v
                    High-speed serial data
                    Differential signalling
                    Clock and data synchronization
```

The complete process:

1. The camera sensor captures an image.
2. The sensor produces digital pixel data.
3. The serializer/interface logic prepares the data for transmission.
4. High-speed serial lanes carry the data to the processor.
5. The receiver recovers the data and its timing.
6. The processor receives the image data.
7. The image-processing pipeline processes the frame.

This example connects several topics from this section:

- SerDes: conversion between parallel and serial data.
- Differential signalling: improved noise immunity.
- MIPI D-PHY: physical link between camera and processor.
- CDR: recovery of timing in interfaces that use embedded clocking.
- Signal integrity: preservation of signal quality over the physical connection.

**The key idea:** High-speed interconnect design is not only about transmitting more bits per second. It must also control power, noise, timing, physical routing and error rates.

# **Final Revision Sheet**

Use this section to connect the major concepts before an exam.

## **1. Important comparisons**

| Concepts                             | Main difference                                                                                                                         |
| ------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------- |
| LLC vs MAC                           | LLC provides the logical interface to upper layers; MAC handles medium access and MAC-level framing/addressing functions.               |
| MAC vs IP                            | MAC supports local-link delivery; IP supports logical addressing and routing.                                                           |
| Byte stuffing vs bit stuffing        | Byte stuffing inserts escape bytes; bit stuffing inserts bits.                                                                          |
| Error detection vs correction        | Detection identifies corruption; correction attempts to recover the original data.                                                      |
| Checksum vs CRC                      | Checksum uses one's-complement addition; CRC uses polynomial division over GF(2).                                                       |
| Hamming vs Reed-Solomon              | Hamming operates on bits; Reed-Solomon operates on multi-bit symbols.                                                                   |
| Flow control vs congestion control   | Flow control protects the receiver; congestion control protects the network.                                                            |
| Go-Back-N vs Selective Repeat        | Go-Back-N retransmits the missing frame and subsequent outstanding frames; Selective Repeat retransmits only missing or damaged frames. |
| HDLC vs PPP                          | HDLC is a link-control protocol; PPP carries network-layer packets over point-to-point links using HDLC-like framing.                   |
| CSMA/CD vs CSMA/CA                   | Collision detection was used in traditional shared Ethernet; collision avoidance is used by Wi-Fi.                                      |
| Hub vs switch                        | A hub repeats signals; a switch learns MAC addresses and selectively forwards frames.                                                   |
| Collision domain vs broadcast domain | A collision domain concerns conflicting transmissions; a broadcast domain concerns the reach of Layer 2 broadcasts.                     |
| VLAN vs STP                          | VLAN separates broadcast domains; STP prevents Layer 2 loops.                                                                           |
| Parallel vs serial                   | Parallel transfers multiple bits at once; serial transfers bits sequentially per lane.                                                  |
| LVDS vs CML                          | LVDS is a low-voltage differential signalling standard; CML is a current-steering circuit/signalling family.                            |

## **2. Important numerical formulas**

| Topic                            | Formula                                             |
| -------------------------------- | --------------------------------------------------- |
| MAC address length               | 48 bits for Ethernet MAC addresses                  |
| IPv4 address length              | 32 bits                                             |
| IPv6 address length              | 128 bits                                            |
| Number of sequence numbers       | \\(2^m\\)                                           |
| Go-Back-N maximum sender window  | \\(2^m-1\\)                                         |
| Selective Repeat maximum window  | \\(2^{m-1}\\)                                       |
| Hamming parity-bit condition     | \\(2^r\geq m+r+1\\)                                 |
| Error detection capability       | \\(d\_{min}-1\\)                                    |
| Error correction capability      | \\(\left\lfloor\frac{d\_{min}-1}{2}\right\rfloor\\) |
| Pure ALOHA throughput            | \\(S=Ge^{-2G}\\)                                    |
| Slotted ALOHA throughput         | \\(S=Ge^{-G}\\)                                     |
| Pure ALOHA maximum throughput    | \\(1/(2e)\approx18.4\\%\\)                          |
| Slotted ALOHA maximum throughput | \\(1/e\approx36.8\\%\\)                             |

Here:

- \\(m\\) = number of data bits in the Hamming-code calculation.
- \\(r\\) = number of parity bits.
- \\(d\_{min}\\) = minimum Hamming distance.
- \\(G\\) = average attempted transmissions per frame time.
- \\(S\\) = normalized throughput.

## **3. Which diagrams should you practise drawing?**

These are particularly useful for theory examinations and viva questions:

Priority 1

**Error detection and correction**

[PPT - PART III PowerPoint Presentation, free download - ID:879821](https://images.openai.com/static-rsc-4/zZMOpYulZwe4qo7n-CMWVPK-Apaiezhb9HRKp-zCt8ouY_fo4pFDqJTjqnl7MHmjN1iJmPe2XwSvxlD-9h5_r4rFS1-P8khKcrL-BfVf4Zs_SPyGo49FgrDhHdkR92SzDS_tGZ3sATuRkK5lcXFejM0gWADkRp9mJDgiaAVrmvs?purpose=inline)

CRC division

[PPT - Hamming Code PowerPoint Presentation, free download - ID:5545001](https://images.openai.com/static-rsc-4/OHtSNI_FJpvWygeKMLYCvQty1btEHYDjh6vqmnXoldkwoDEHwcS270_qcZiEIS-o3t6CxRqViq7UEWnOV_he8mvVDbh8AvFk57GgtzTdXezVFNJPlQwb7B4qInQXytR8pNHIqqJMv9-IwQLNgArmcIehk0RZA9oP291TsrUxh8M?purpose=inline)

Hamming code

Priority 2

**Flow and error control**

[PPT - Computer Networks: Introduction PowerPoint Presentation, free download - ID:6110882](https://images.openai.com/static-rsc-4/2iejmwv1pVpVo7vz44JqQCzZIu7ShsKX9Y_YptOYsr0zSE_nMgcog-Kl_JdLfFEBQuyqTmD6NN8ySx2vN4SmzUaGZOg6lIN55zXVrF-MJv9cPDsEkw9GYvcrLK1DmejXF-TfI3kUsKp1DTuBEpE2UhMFVwuCHOS4cjU8jEeT8pI?purpose=inline)

Go-Back-N ARQ

[Noted of Reliable Delivery in Computer Networks | PDF](https://images.openai.com/static-rsc-4/o1KmJogokRJnPkpyizkjVG65Yh72fcmg_POdSw3uWYPEyJSaPjwSY9G-5PzA_JVF3xaGP1Mh1kXbypzfSNeSiy214BMNjzNxEzegLmR-XRiVm6-zFVGP2g7k1kbAtxQ1OFo73BL4-YYHp3PbuDtm1XAaylCLK83AHYiKq6aj5f8?purpose=inline)

Selective Repeat ARQ

Priority 3

**Ethernet and switching**

[Frame Ethernet](https://images.openai.com/static-rsc-4/Bzc9cxmDO4cpebariJ2-GR7BZ6LVEIil77hTDaieEADUeztOSl1-plluebwDpvRONxDtSYtGJfQUjASQ7NSTr6hlFaUZvY1qCrifS9p4PylDQ4yJXCUiXl-ezwUZOKsENUjqNiqOFixxbuUBoUJcfac1TISjpfOfimGWQ0_1_NM?purpose=inline)

Ethernet frame

[Ithy - Unraveling Cisco's Spanning Tree Protocol: Your Key to Loop-Free Networks](https://images.openai.com/static-rsc-4/atDqS_Ml0KtqkMdjbqLH_-eHasi3dfNDYkLbxWUg-ahvT3WUBs8KasENsTFBY68u1Gk-tLN1GgPf8otEjkGvb21tPOLD30yX2a1Wn3PyfAQrQCPdikNHXVzirNkxLk0Nedrw1bkfuP_G7CdVCaHtAIuVsLYG1Q9f0VLdS0utpFI?purpose=inline)

Spanning Tree Protocol

**High-speed interconnects**

[Know Your 400G Transceiver | Juniper Networks](https://images.openai.com/static-rsc-4/QoyNPEuBe8mH8EvNsCIJ1wmEJHVtEs4m1XQHmCH9BMzp9oV0ceY8M8okUu-Zq7AR6REZTb3Dn5dMm5045BwRWXo-O8jYXmW5NS0OLRMqhWZ6eJBDdw73iHSLgek3Bo7YPImKB317OZilL3CujRFPpam5qZFDQ1uafblRGBig3EI?purpose=inline)

[AN-1177: LVDS and M-LVDS Circuit Implementation Guide | Analog Devices](https://images.openai.com/static-rsc-4/gSFeSfMRNAD4N0YUsefb1hFrvQ2e0lZtmG1cLKQU1YhGOYwn4dYUZBy4pFuITYjH6GO21Sp91TFoFqHI3EYFycwD6QE1-0UDJKYinjSWbdKAH4S2xZQWbyeWBdywvmjFyDMR0-_05jefX_5fNWrEuiP7i4UB6IxSkNroJe384a4?purpose=inline)

[MIPI协议及应用-腾讯云开发者社区-腾讯云](https://images.openai.com/static-rsc-4/4zoLdbCrI_vv2vUJUsX2ztFxcgGsq8KJcyQYKshPxGvqm65smkPA5M6Uoj3BwcWNiD3Fyp6Cqr1NHQaGv2EDJzg0pocWxyS1j5kt2CuEstTcoKRigWgPPXjsgii7YKjmLyR4Lug0kmuUJOnFXB6Yfzqru6CwzbiIXxRaJYce0aE?purpose=inline)

9

## **4. Recommended practice order**

I recommend studying and practising the unit in this order:

1. **Framing:** Character count, byte stuffing and bit stuffing.
2. **Error detection:** Parity, two-dimensional parity, checksum and CRC.
3. **Error correction:** Hamming distance, Hamming code, redundancy bits, BCH and Reed-Solomon.
4. **Flow and error control:** Stop-and-Wait, Go-Back-N, Selective Repeat and sliding windows.
5. **Data Link protocols:** HDLC, PPP, LCP and NCP.
6. **MAC protocols:** ALOHA, CSMA, CSMA/CD and CSMA/CA.
7. **Ethernet and switching:** MAC learning, forwarding, VLAN and STP.
8. **VLSI interconnects:** SerDes, LVDS, MIPI D-PHY, CML, CDR and signal integrity.

For the numerical portion, give extra practice to **CRC, checksum, parity, Hamming-code construction, syndrome calculation, Hamming distance, sequence-number windows and ALOHA throughput**.

## **5. Technical references**

These are useful for moving beyond textbook explanations:

- [RFC 1662 — PPP in HDLC-like Framing](https://www.rfc-editor.org/rfc/rfc1662.html?utm_source=chatgpt.com): PPP frame structure, flags, FCS and escaping.
- [RFC 1661 — The Point-to-Point Protocol](https://www.rfc-editor.org/info/rfc1661/?utm_source=chatgpt.com): PPP architecture and protocol operation.
- [IEEE 802.3 — Ethernet Standard](https://standards.ieee.org/ieee/802/10422/?utm_source=chatgpt.com): Ethernet MAC and PHY standard.
- [IEEE 802.1Q — Bridges and Bridged Networks](https://standards.ieee.org/ieee/802.1q/10323/?utm_source=chatgpt.com): VLANs and bridged networks.
- [Cisco — LAN Switching and STP](https://www.cisco.com/c/en/us/support/docs/lan-switching/ethernet/12006-chapter22.html?utm_source=chatgpt.com): MAC learning, forwarding, filtering, flooding and switching loops.
- [Texas Instruments — LVDS SerDes Design Guidelines](https://www.ti.com/lit/an/snla008a/snla008a.pdf?ts=1776257057445\&utm_source=chatgpt.com): Serializer/deserializer and differential-link design.
- [MIPI Alliance — D-PHY](https://www.mipi.org/specifications/d-phy?utm_source=chatgpt.com): Camera/display physical-layer interface.

