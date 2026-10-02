# C++ / Linux File System — Topic List

## 1. File System Fundamentals
- [ ] What is a File System
- [ ] Files
- [ ] Directories
- [ ] File Names
- [ ] File Paths
- [ ] Absolute Paths
- [ ] Relative Paths
- [ ] File Extensions
- [ ] File Metadata
- [ ] File Types
- [ ] File Descriptors
- [ ] Inodes
- [ ] File System Blocks

## 2. Linux File System Hierarchy
- [ ] `/`
- [ ] `/bin`
- [ ] `/sbin`
- [ ] `/etc`
- [ ] `/home`
- [ ] `/root`
- [ ] `/tmp`
- [ ] `/var`
- [ ] `/usr`
- [ ] `/dev`
- [ ] `/proc`
- [ ] `/sys`
- [ ] `/run`
- [ ] `/mnt`
- [ ] `/opt`

## 3. File Types
- [ ] Regular Files
- [ ] Directories
- [ ] Symbolic Links
- [ ] Hard Links
- [ ] Character Devices
- [ ] Block Devices
- [ ] Named Pipes
- [ ] Unix Domain Sockets

## 4. File Permissions
- [ ] Read Permission
- [ ] Write Permission
- [ ] Execute Permission
- [ ] User
- [ ] Group
- [ ] Others
- [ ] `chmod`
- [ ] `chown`
- [ ] `chgrp`
- [ ] Permission Bits
- [ ] `umask`
- [ ] SUID
- [ ] SGID
- [ ] Sticky Bit

## 5. File System Calls
- [ ] `open()`
- [ ] `close()`
- [ ] `read()`
- [ ] `write()`
- [ ] `pread()`
- [ ] `pwrite()`
- [ ] `lseek()`
- [ ] `fsync()`
- [ ] `fdatasync()`
- [ ] `ftruncate()`
- [ ] `dup()`
- [ ] `dup2()`
- [ ] `fcntl()`
- [ ] `ioctl()`

## 6. File Open Modes
- [ ] `O_RDONLY`
- [ ] `O_WRONLY`
- [ ] `O_RDWR`
- [ ] `O_CREAT`
- [ ] `O_EXCL`
- [ ] `O_TRUNC`
- [ ] `O_APPEND`
- [ ] `O_NONBLOCK`
- [ ] `O_SYNC`
- [ ] `O_CLOEXEC`

## 7. File Descriptors
- [ ] File Descriptor Concept
- [ ] `stdin`
- [ ] `stdout`
- [ ] `stderr`
- [ ] File Descriptor Table
- [ ] File Descriptor Limits
- [ ] Descriptor Duplication
- [ ] Descriptor Inheritance
- [ ] Descriptor Closing
- [ ] File Descriptor Leaks

## 8. File Metadata
- [ ] `stat()`
- [ ] `fstat()`
- [ ] `lstat()`
- [ ] File Size
- [ ] File Permissions
- [ ] Owner
- [ ] Group
- [ ] Inode Number
- [ ] Access Time
- [ ] Modification Time
- [ ] Change Time
- [ ] File Type

## 9. Directory Operations
- [ ] `mkdir()`
- [ ] `rmdir()`
- [ ] `opendir()`
- [ ] `readdir()`
- [ ] `closedir()`
- [ ] `chdir()`
- [ ] `getcwd()`
- [ ] Directory Traversal
- [ ] Recursive Directory Traversal

## 10. File Management
- [ ] `rename()`
- [ ] `unlink()`
- [ ] `link()`
- [ ] `symlink()`
- [ ] `readlink()`
- [ ] File Creation
- [ ] File Deletion
- [ ] File Copying
- [ ] File Moving
- [ ] Temporary Files

## 11. Hard Links & Symbolic Links
- [ ] Hard Links
- [ ] Symbolic Links
- [ ] Inode Sharing
- [ ] Link Count
- [ ] `link()`
- [ ] `unlink()`
- [ ] `symlink()`
- [ ] `readlink()`
- [ ] Hard Link vs Symbolic Link

## 12. File I/O
- [ ] Buffered I/O
- [ ] Unbuffered I/O
- [ ] Sequential I/O
- [ ] Random Access I/O
- [ ] Blocking I/O
- [ ] Non-Blocking I/O
- [ ] Partial Reads
- [ ] Partial Writes
- [ ] EOF
- [ ] File Offset

## 13. C++ File I/O
- [ ] `std::ifstream`
- [ ] `std::ofstream`
- [ ] `std::fstream`
- [ ] `open()`
- [ ] `close()`
- [ ] `read()`
- [ ] `write()`
- [ ] `seekg()`
- [ ] `seekp()`
- [ ] `tellg()`
- [ ] `tellp()`
- [ ] Binary File I/O
- [ ] Text File I/O

## 14. Memory-Mapped Files
- [ ] `mmap()`
- [ ] `munmap()`
- [ ] `mprotect()`
- [ ] File-Backed Memory
- [ ] Anonymous Mapping
- [ ] Shared Mapping
- [ ] Private Mapping
- [ ] Memory-Mapped I/O
- [ ] Copy-on-Write
- [ ] Mapping Permissions

## 15. File System Architecture
- [ ] File System Interface
- [ ] Virtual File System (VFS)
- [ ] File System Driver
- [ ] Inode
- [ ] Directory Entry
- [ ] Superblock
- [ ] Data Blocks
- [ ] Metadata
- [ ] Block Allocation
- [ ] Free Space Management

## 16. Inodes
- [ ] Inode Structure
- [ ] Inode Number
- [ ] File Metadata
- [ ] File Ownership
- [ ] Permission Bits
- [ ] Timestamps
- [ ] Link Count
- [ ] Data Block References
- [ ] Direct Blocks
- [ ] Indirect Blocks
- [ ] Double Indirect Blocks
- [ ] Triple Indirect Blocks

## 17. Storage & Blocks
- [ ] Disk Blocks
- [ ] Block Size
- [ ] Block Allocation
- [ ] Free Blocks
- [ ] Block Bitmap
- [ ] Data Blocks
- [ ] Metadata Blocks
- [ ] Fragmentation
- [ ] Internal Fragmentation
- [ ] External Fragmentation

## 18. Journaling
- [ ] Journaling
- [ ] Journal
- [ ] Metadata Journaling
- [ ] Ordered Journaling
- [ ] Writeback Journaling
- [ ] Crash Recovery
- [ ] Transaction
- [ ] Filesystem Consistency

## 19. File System Types
- [ ] ext4
- [ ] XFS
- [ ] Btrfs
- [ ] ZFS
- [ ] tmpfs
- [ ] procfs
- [ ] sysfs
- [ ] FAT
- [ ] NTFS
- [ ] File System Comparison

## 20. Mounting
- [ ] Mount Point
- [ ] `mount`
- [ ] `umount`
- [ ] `/etc/fstab`
- [ ] Mount Options
- [ ] Bind Mounts
- [ ] Read-Only Mounts
- [ ] Filesystem Mounting

## 21. Virtual File Systems
- [ ] VFS
- [ ] `/proc`
- [ ] `/sys`
- [ ] `/dev`
- [ ] `/tmp`
- [ ] Virtual Files
- [ ] Device Files
- [ ] Kernel Information via Filesystem

## 22. File Locking & Concurrency
- [ ] File Race Conditions
- [ ] File Locking
- [ ] Advisory Locks
- [ ] Mandatory Locks
- [ ] `flock()`
- [ ] `fcntl()` Locks
- [ ] Shared Locks
- [ ] Exclusive Locks
- [ ] Concurrent File Access

## 23. File System Performance
- [ ] Page Cache
- [ ] Buffer Cache
- [ ] Read Cache
- [ ] Write Cache
- [ ] Disk I/O
- [ ] I/O Scheduling
- [ ] Sequential vs Random I/O
- [ ] `fsync()`
- [ ] `O_DIRECT`
- [ ] Read-Ahead
- [ ] Writeback
- [ ] I/O Bottlenecks

## 24. Advanced I/O
- [ ] `select()`
- [ ] `poll()`
- [ ] `epoll()`
- [ ] Asynchronous I/O
- [ ] `io_uring`
- [ ] Scatter/Gather I/O
- [ ] `readv()`
- [ ] `writev()`
- [ ] `preadv()`
- [ ] `pwritev()`
- [ ] `sendfile()`
- [ ] `splice()`

## 25. File System Reliability
- [ ] Atomic File Operations
- [ ] Crash Consistency
- [ ] Data Durability
- [ ] `fsync()`
- [ ] Atomic Rename
- [ ] Temporary Files
- [ ] File Corruption
- [ ] Recovery
- [ ] Checksums
- [ ] Journaling

## 26. File System Security
- [ ] File Permissions
- [ ] Ownership
- [ ] Access Control
- [ ] ACLs
- [ ] Capability-Based Security
- [ ] Path Traversal
- [ ] Symlink Attacks
- [ ] TOCTOU
- [ ] Secure Temporary Files
- [ ] Sandboxing

## 27. Storage Concepts
- [ ] HDD
- [ ] SSD
- [ ] NVMe
- [ ] Disk Latency
- [ ] IOPS
- [ ] Throughput
- [ ] Storage Hierarchy
- [ ] RAID Basics
- [ ] Persistent Storage

## 28. Redis / Systems Project Relevant 🔥
- [ ] File Descriptors
- [ ] Binary File I/O
- [ ] Sequential Writes
- [ ] Random Access
- [ ] Append-Only Files
- [ ] File Persistence
- [ ] Write-Ahead Logging
- [ ] Memory-Mapped Files
- [ ] `fsync()`
- [ ] Crash Recovery
- [ ] Atomic Rename
- [ ] File Locking
- [ ] Serialization
- [ ] Deserialization
- [ ] Snapshotting
- [ ] Background Persistence
- [ ] File Compaction
- [ ] Log Compaction
- [ ] Disk-Based Data Structures
- [ ] Recovery After Crash

## 29. File System Projects
- [ ] File Explorer
- [ ] `find`-Like Utility
- [ ] `du`-Like Utility
- [ ] `cp`-Like Utility
- [ ] `mv`-Like Utility
- [ ] `cat`-Like Utility
- [ ] File Search Engine
- [ ] File Watcher
- [ ] File Backup Tool
- [ ] Key-Value Store Using Files
- [ ] Write-Ahead Log
- [ ] Append-Only Log
- [ ] Mini File System
- [ ] In-Memory File System
- [ ] Persistent Key-Value Store