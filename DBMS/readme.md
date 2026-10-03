# Database / DBMS Roadmap

## 1. Database Fundamentals

- Database
- DBMS
- RDBMS
- File System vs DBMS
- Advantages of DBMS
- Disadvantages of DBMS
- DBMS Architecture
- Database Users
- Database Administrator
- Database Schema
- Database Instance
- Metadata
- Data Models
- Hierarchical Model
- Network Model
- Relational Model
- Object-Oriented Model
- Physical Data Independence
- Logical Data Independence
- Three-Schema Architecture
- Data Abstraction

---

## 2. Relational Model

- Relation
- Tuple
- Attribute
- Domain
- Degree
- Cardinality
- Relation Schema
- Relation Instance
- Keys
- Primary Key
- Candidate Key
- Super Key
- Alternate Key
- Foreign Key
- Composite Key
- Surrogate Key
- Natural Key
- Referential Integrity

---

## 3. ER Model

- Entity
- Entity Set
- Attributes
- Simple Attributes
- Composite Attributes
- Single-Valued Attributes
- Multi-Valued Attributes
- Derived Attributes
- Key Attributes
- Relationships
- Relationship Sets
- Degree of Relationship
- Cardinality
- Participation Constraints
- Total Participation
- Partial Participation
- Weak Entity
- Strong Entity
- ER Diagram
- ER to Relational Mapping

---

## 4. SQL Fundamentals

- SQL
- DDL
- DML
- DQL
- DCL
- TCL
- CREATE
- ALTER
- DROP
- TRUNCATE
- INSERT
- SELECT
- UPDATE
- DELETE
- WHERE
- ORDER BY
- GROUP BY
- HAVING
- DISTINCT
- LIMIT
- OFFSET
- AS
- Aliases
- NULL
- IS NULL
- IS NOT NULL
- COALESCE
- CASE

---

## 5. SQL Operators

- Arithmetic Operators
- Comparison Operators
- Logical Operators
- AND
- OR
- NOT
- IN
- NOT IN
- BETWEEN
- LIKE
- ILIKE
- EXISTS
- ANY
- ALL

---

## 6. SQL Functions

### Aggregate Functions

- COUNT
- SUM
- AVG
- MIN
- MAX

### String Functions

- CONCAT
- LENGTH
- LOWER
- UPPER
- SUBSTRING
- TRIM
- REPLACE

### Numeric Functions

- ROUND
- CEIL
- FLOOR
- ABS
- MOD

### Date/Time Functions

- CURRENT_DATE
- CURRENT_TIME
- CURRENT_TIMESTAMP
- Date arithmetic
- Date extraction

---

## 7. SQL Joins

- INNER JOIN
- LEFT JOIN
- RIGHT JOIN
- FULL OUTER JOIN
- CROSS JOIN
- SELF JOIN
- Multiple Joins
- Join Conditions
- Join Filtering
- Join Ordering
- Cartesian Product
- JOIN vs Subquery

---

## 8. Subqueries

- Subquery
- Scalar Subquery
- Single-Row Subquery
- Multi-Row Subquery
- Correlated Subquery
- Nested Subquery
- IN Subquery
- EXISTS Subquery
- NOT EXISTS
- Subquery vs JOIN

---

## 9. Set Operations

- UNION
- UNION ALL
- INTERSECT
- EXCEPT
- Set Operation Rules
- Duplicate Handling

---

## 10. Common Table Expressions

- CTE
- WITH
- Multiple CTEs
- Recursive CTE
- CTE vs Subquery
- CTE vs Temporary Table

---

## 11. Window Functions

- Window Functions
- OVER()
- PARTITION BY
- ORDER BY
- ROW_NUMBER
- RANK
- DENSE_RANK
- NTILE
- LAG
- LEAD
- FIRST_VALUE
- LAST_VALUE
- Running Total
- Moving Average
- Window Frames

---

## 12. Constraints

- PRIMARY KEY
- FOREIGN KEY
- UNIQUE
- NOT NULL
- CHECK
- DEFAULT
- Referential Integrity
- ON DELETE CASCADE
- ON DELETE SET NULL
- ON DELETE RESTRICT
- ON UPDATE CASCADE
- Constraint Violations

---

## 13. Database Normalization

- Functional Dependency
- Partial Dependency
- Transitive Dependency
- Data Redundancy
- Insertion Anomaly
- Update Anomaly
- Deletion Anomaly
- 1NF
- 2NF
- 3NF
- BCNF
- 4NF
- 5NF
- Denormalization
- Normalization vs Denormalization

---

## 14. Functional Dependencies

- Functional Dependency
- Trivial Dependency
- Non-Trivial Dependency
- Full Functional Dependency
- Partial Dependency
- Transitive Dependency
- Attribute Closure
- Candidate Key from Functional Dependencies
- Armstrong's Axioms
- Decomposition
- Lossless Decomposition
- Dependency Preservation

---

## 15. Transactions

- Transaction
- Transaction States
- Active
- Partially Committed
- Committed
- Failed
- Aborted
- Transaction Lifecycle
- BEGIN
- COMMIT
- ROLLBACK
- SAVEPOINT

---

## 16. ACID Properties

- Atomicity
- Consistency
- Isolation
- Durability
- Atomicity Examples
- Consistency Constraints
- Isolation Problems
- Durability Mechanisms

---

## 17. Transaction Schedules

- Schedule
- Serial Schedule
- Non-Serial Schedule
- Concurrent Schedule
- Recoverable Schedule
- Cascadeless Schedule
- Strict Schedule
- Conflict
- Conflict Serializability
- View Serializability
- Precedence Graph

---

## 18. Concurrency Control

- Concurrent Transactions
- Race Conditions
- Lost Update
- Dirty Read
- Non-Repeatable Read
- Phantom Read
- Write Skew
- Concurrency Problems
- Lock-Based Protocols
- Timestamp-Based Protocols
- Optimistic Concurrency Control
- Pessimistic Concurrency Control
- MVCC

---

## 19. Isolation Levels

- Read Uncommitted
- Read Committed
- Repeatable Read
- Serializable
- Isolation Level Tradeoffs
- Dirty Reads
- Non-Repeatable Reads
- Phantom Reads
- Serialization Failures

---

## 20. Locks

- Shared Lock
- Exclusive Lock
- Row-Level Lock
- Table-Level Lock
- Intent Lock
- Lock Compatibility
- Lock Acquisition
- Lock Release
- Lock Contention
- Lock Escalation
- Two-Phase Locking
- Strict 2PL
- SELECT FOR UPDATE

---

## 21. Deadlocks

- Deadlock
- Necessary Conditions for Deadlock
- Mutual Exclusion
- Hold and Wait
- No Preemption
- Circular Wait
- Deadlock Detection
- Deadlock Prevention
- Deadlock Avoidance
- Deadlock Recovery
- Deadlock Timeout
- Handling Deadlocks

---

## 22. Indexing

- What is an Index?
- Why Indexes?
- Index Lookup
- Full Table Scan
- Primary Index
- Secondary Index
- Clustered Index
- Non-Clustered Index
- Dense Index
- Sparse Index
- Unique Index
- Composite Index
- Covering Index
- Partial Index
- Expression Index
- Index Selectivity
- Index Cardinality
- Index Maintenance
- Index Tradeoffs

---

## 23. B-Tree / B+ Tree

- B-Tree
- B+ Tree
- B-Tree vs B+ Tree
- Nodes
- Keys
- Pointers
- Internal Nodes
- Leaf Nodes
- Tree Height
- Searching
- Insertion
- Deletion
- Node Splitting
- Node Merging
- Range Queries
- Why Databases Use B+ Trees

---

## 24. Hash Indexing

- Hash Index
- Hash Function
- Buckets
- Collision
- Collision Resolution
- Hash Index vs B-Tree
- Equality Queries
- Range Queries

---

## 25. Query Processing

- SQL Parsing
- Query Parsing
- Query Validation
- Query Optimization
- Query Planning
- Query Execution
- Execution Plan
- Query Operators
- Selection
- Projection
- Join
- Sorting
- Aggregation

---

## 26. Query Optimization

- Query Optimizer
- Cost-Based Optimization
- Rule-Based Optimization
- Query Statistics
- Cardinality Estimation
- Selectivity
- Join Ordering
- Index Selection
- Predicate Pushdown
- Projection Pushdown
- Query Rewriting
- EXPLAIN
- EXPLAIN ANALYZE

---

## 27. Join Algorithms

- Nested Loop Join
- Block Nested Loop Join
- Hash Join
- Sort-Merge Join
- Join Cost
- Join Ordering
- Join Optimization

---

## 28. Storage Management

- Database Pages
- Blocks
- Records
- Rows
- Heap Files
- Data Files
- Page Layout
- Slotted Pages
- Free Space
- Record Organization
- Row Store
- Column Store

---

## 29. Buffer Management

- Buffer Pool
- Buffer Manager
- Pages in Memory
- Page Replacement
- Cache Hit
- Cache Miss
- Dirty Pages
- Page Eviction
- LRU
- Clock Algorithm
- Disk I/O

---

## 30. Database Storage

- Disk Storage
- SSD
- HDD
- Sequential I/O
- Random I/O
- Disk Blocks
- Pages
- Storage Layout
- I/O Cost
- Memory vs Disk
- Buffer Pool

---

## 31. Database Recovery

- Failure Types
- Transaction Failure
- System Crash
- Media Failure
- Crash Recovery
- Log-Based Recovery
- Write-Ahead Logging
- WAL
- Undo
- Redo
- Checkpoints
- Recovery Manager
- Shadow Paging
- ARIES

---

## 32. Write-Ahead Logging

- WAL Concept
- WAL Records
- Log Sequence Number
- Redo
- Undo
- Checkpoints
- WAL and Durability
- WAL and Crash Recovery
- WAL and Replication

---

## 33. Views

- Views
- Creating Views
- Updating Views
- View Security
- Materialized Views
- Materialized View Refresh
- View vs Table
- View vs Materialized View

---

## 34. Stored Procedures

- Stored Procedures
- Functions
- Parameters
- Return Values
- Procedure Execution
- Advantages
- Disadvantages
- When to Use Procedures

---

## 35. Triggers

- Trigger
- BEFORE Trigger
- AFTER Trigger
- INSERT Trigger
- UPDATE Trigger
- DELETE Trigger
- Row-Level Trigger
- Statement-Level Trigger
- Trigger Use Cases
- Trigger Problems

---

## 36. PostgreSQL

- PostgreSQL Architecture
- PostgreSQL Installation
- psql
- Databases
- Schemas
- Roles
- Users
- Permissions
- PostgreSQL Data Types
- UUID
- JSONB
- Arrays
- ENUM
- Sequences
- Identity Columns
- PostgreSQL Indexes
- PostgreSQL Transactions
- PostgreSQL Extensions
- PostgreSQL EXPLAIN

---

## 37. MySQL

- MySQL Architecture
- MySQL Storage Engines
- InnoDB
- MyISAM
- MySQL Transactions
- MySQL Indexes
- MySQL Isolation Levels
- MySQL EXPLAIN
- MySQL Replication

---

## 38. SQL vs NoSQL

- SQL Databases
- NoSQL Databases
- Relational Databases
- Key-Value Databases
- Document Databases
- Wide-Column Databases
- Graph Databases
- SQL vs NoSQL Tradeoffs
- Schema-on-Write
- Schema-on-Read
- Structured Data
- Semi-Structured Data

---

## 39. NoSQL Databases

- Key-Value Stores
- Document Databases
- Column-Family Databases
- Graph Databases
- Redis
- MongoDB
- Cassandra
- DynamoDB
- Data Modeling in NoSQL
- Denormalization
- Eventual Consistency

---

## 40. CAP Theorem

- Consistency
- Availability
- Partition Tolerance
- Network Partition
- CAP Tradeoffs
- CP Systems
- AP Systems
- CAP Misconceptions

---

## 41. BASE

- Basically Available
- Soft State
- Eventual Consistency
- BASE vs ACID

---

## 42. Distributed Databases

- Distributed Database
- Replication
- Sharding
- Partitioning
- Distributed Transactions
- Distributed Consistency
- Quorum
- Leader
- Follower
- Leader Election
- Failover
- Consensus
- Two-Phase Commit

---

## 43. Replication

- Database Replication
- Primary-Replica
- Leader-Follower
- Read Replicas
- Synchronous Replication
- Asynchronous Replication
- Replication Lag
- Failover
- Read Scaling
- High Availability

---

## 44. Partitioning

- Horizontal Partitioning
- Vertical Partitioning
- Range Partitioning
- Hash Partitioning
- List Partitioning
- Partition Pruning
- Partition Management
- Partitioned Indexes

---

## 45. Sharding

- Database Sharding
- Shard Key
- Hash Sharding
- Range Sharding
- Consistent Hashing
- Shard Distribution
- Hot Shards
- Cross-Shard Queries
- Cross-Shard Transactions
- Rebalancing

---

## 46. Caching

- Why Caching?
- Cache Hit
- Cache Miss
- Cache Hit Ratio
- Cache-Aside
- Read-Through
- Write-Through
- Write-Behind
- Cache Invalidation
- TTL
- Cache Stampede
- Cache Penetration
- Cache Avalanche
- Cache Consistency

---

## 47. Redis

- Redis Architecture
- Strings
- Lists
- Sets
- Sorted Sets
- Hashes
- Streams
- Pub/Sub
- TTL
- Expiration
- Transactions
- Lua Scripts
- Persistence
- RDB
- AOF
- Replication
- Redis Cluster

---

## 48. Database Connection Management

- Database Connections
- Connection Lifecycle
- Connection Pooling
- Maximum Connections
- Minimum Connections
- Idle Connections
- Connection Timeout
- Connection Reuse
- Connection Exhaustion
- Pool Monitoring

---

## 49. Database Migrations

- Database Migration
- Migration Files
- Up Migration
- Down Migration
- Migration Versioning
- Schema Evolution
- Migration Rollback
- Migration Conflicts
- Production Migrations
- Zero-Downtime Migrations

---

## 50. Database Security

- Authentication
- Authorization
- Users
- Roles
- Permissions
- Least Privilege
- SQL Injection
- Parameterized Queries
- Prepared Statements
- Password Hashing
- Secrets Management
- TLS
- Encryption at Rest
- Encryption in Transit

---

## 51. Database Performance

- Query Performance
- Index Optimization
- Query Optimization
- Connection Pool Tuning
- Slow Query Analysis
- Lock Contention
- Deadlocks
- CPU Bottlenecks
- Memory Bottlenecks
- Disk I/O Bottlenecks
- Cache Performance
- Database Profiling

---

## 52. Database Monitoring

- Query Latency
- Query Throughput
- Error Rate
- Connection Count
- Connection Pool Usage
- Cache Hit Ratio
- Lock Contention
- Deadlocks
- Replication Lag
- Disk Usage
- CPU Usage
- Memory Usage
- Slow Queries

---

## 53. Backup and Recovery

- Database Backup
- Full Backup
- Incremental Backup
- Differential Backup
- Point-in-Time Recovery
- WAL Archiving
- Restore
- Disaster Recovery
- Recovery Point Objective
- Recovery Time Objective
- Backup Testing

---

## 54. Database Design Patterns

- Repository Pattern
- Unit of Work
- Data Mapper
- Active Record
- DAO
- CQRS
- Event Sourcing
- Transactional Outbox
- Saga Pattern
- Read Model
- Write Model

---

## 55. Database Testing

- Unit Testing Database Code
- Integration Testing
- Test Database
- Test Containers
- Migration Testing
- Query Testing
- Constraint Testing
- Transaction Testing
- Concurrent Transaction Testing
- Failure Testing
- Performance Testing

---

# Interview-Focused DBMS Topics

## Must Know ⭐⭐⭐⭐⭐

- DBMS vs RDBMS
- Three-Schema Architecture
- Data Independence
- Keys
- Primary Key
- Foreign Key
- Candidate Key
- Super Key
- ER Model
- Normalization
- 1NF
- 2NF
- 3NF
- BCNF
- Functional Dependencies
- SQL
- Joins
- Subqueries
- Aggregations
- GROUP BY
- HAVING
- Indexing
- B+ Tree
- Transactions
- ACID
- Serializability
- Concurrency Control
- Isolation Levels
- Locks
- Deadlocks
- MVCC
- Query Optimization
- SQL vs NoSQL
- CAP Theorem

# SQL Interview Topics ⭐⭐⭐⭐⭐

- SELECT
- WHERE
- GROUP BY
- HAVING
- ORDER BY
- DISTINCT
- LIMIT
- Aggregate Functions
- Joins
- Subqueries
- CTE
- Window Functions
- CASE
- NULL Handling
- UNION
- UNION ALL
- EXISTS
- IN
- Primary Keys
- Foreign Keys
- Constraints
- Indexes
- Transactions

# Backend Interview Topics ⭐⭐⭐⭐⭐

- Transactions
- ACID
- Isolation Levels
- Locks
- Deadlocks
- MVCC
- Indexing
- B+ Trees
- Query Optimization
- EXPLAIN
- Connection Pooling
- Database Migrations
- Replication
- Read Replicas
- Sharding
- Partitioning
- Caching
- Redis
- Consistency
- CAP Theorem
- Eventual Consistency
- Idempotency
- Distributed Transactions
- Outbox Pattern

# Advanced System Design Topics ⭐⭐⭐⭐

- Database Replication
- Sharding
- Consistent Hashing
- Partitioning
- Distributed Transactions
- Two-Phase Commit
- Consensus
- Quorum
- Leader Election
- Eventual Consistency
- WAL
- MVCC
- Storage Engines
- Buffer Pool
- Query Planner
- Database Recovery
- CQRS
- Event Sourcing
- Transactional Outbox
- Saga Pattern