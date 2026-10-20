

# SQL Joins — Deep Practical & Theoretical Notes

> A comprehensive reference for SQL JOINs, covering theory, practical usage, query patterns, cardinality, NULL behavior, performance, and interview concepts.

---

## Table of Contents

1. [What Is a JOIN?](#1-what-is-a-join)
2. [Why Do We Need JOINs?](#2-why-do-we-need-joins)
3. [JOINs and Relational Algebra](#3-joins-and-relational-algebra)
4. [INNER JOIN](#4-inner-join)
5. [LEFT JOIN](#5-left-join)
6. [RIGHT JOIN](#6-right-join)
7. [FULL OUTER JOIN](#7-full-outer-join)
8. [CROSS JOIN](#8-cross-join)
9. [SELF JOIN](#9-self-join)
10. [Multiple JOINs](#10-multiple-joins)
11. [JOIN Conditions](#11-join-conditions)
12. [JOIN Filtering — ON vs WHERE](#12-join-filtering--on-vs-where)
13. [JOIN Ordering](#13-join-ordering)
14. [Cartesian Product](#14-cartesian-product)
15. [Row Multiplication](#15-row-multiplication)
16. [JOIN vs Subquery](#16-join-vs-subquery)
17. [EXISTS, Semi-JOIN and Anti-JOIN](#17-exists-semi-join-and-anti-join)
18. [JOIN Cardinality](#18-join-cardinality)
19. [JOINs and NULL](#19-joins-and-null)
20. [JOIN vs UNION](#20-join-vs-union)
21. [JOIN Performance](#21-join-performance)
22. [JOIN Algorithms](#22-join-algorithms)
23. [Indexes and JOINs](#23-indexes-and-joins)
24. [Query Grain](#24-query-grain)
25. [Real-World JOIN Patterns](#25-real-world-join-patterns)
26. [JOIN Cheat Sheet](#26-join-cheat-sheet)
27. [JOIN Debugging Strategy](#27-join-debugging-strategy)
28. [Interview-Level Decision Guide](#28-interview-level-decision-guide)
29. [Final Mental Model](#29-final-mental-model)

---

# 1. What Is a JOIN?

A **JOIN combines rows from two or more tables based on a relationship between them**.

Consider:

### `employees`

| employee_id | name | department_id | manager_id |
|---:|---|---:|---:|
| 1 | Alice | 10 | NULL |
| 2 | Bob | 20 | 1 |
| 3 | Charlie | 10 | 1 |
| 4 | David | 30 | 2 |
| 5 | Emma | NULL | 2 |

### `departments`

| department_id | department_name |
|---:|---|
| 10 | Engineering |
| 20 | HR |
| 30 | Sales |
| 40 | Finance |

The relationship is:

```text
employees.department_id
        ↓
departments.department_id
```

A typical JOIN condition is:

```sql
employees.department_id = departments.department_id
```

This is called the **join condition**.

---

# 2. Why Do We Need JOINs?

Relational databases generally avoid storing the same information repeatedly.

Instead of storing:

```text
employee
--------------------------------
Alice    | Engineering
Bob      | HR
Charlie  | Engineering
David    | Sales
```

we normalize the data:

```text
employees
--------------------------------
employee_id | name | department_id

departments
--------------------------------
department_id | department_name
```

Then JOIN the tables when the combined information is needed:

```sql
SELECT
    e.name,
    d.department_name
FROM employees e
JOIN departments d
    ON e.department_id = d.department_id;
```

Result:

| name | department_name |
|---|---|
| Alice | Engineering |
| Bob | HR |
| Charlie | Engineering |
| David | Sales |

### Core idea

> A JOIN says: **"Combine rows from these relations when they satisfy this relationship."**

---

# 3. JOINs and Relational Algebra

SQL JOINs are closely related to **relational algebra**, the mathematical foundation of relational databases.

A useful conceptual model is:

```text
Cartesian Product
       ↓
Apply matching condition
       ↓
JOIN result
```

For example:

```text
A
---
1
2

B
---
1
3
```

Cartesian product:

```text
1,1
1,3
2,1
2,3
```

Apply:

```text
A.id = B.id
```

Result:

```text
1,1
```

Conceptually:

```text
A JOIN B
```

can be understood as combining relations and retaining rows satisfying the join predicate.

This mental model is especially useful for understanding:

- INNER JOIN
- CROSS JOIN
- duplicate rows
- many-to-many relationships
- Cartesian products
- JOIN performance

---

# 4. INNER JOIN

## Definition

An **INNER JOIN returns only rows that have matching rows in both tables**.

Syntax:

```sql
SELECT columns
FROM table1
INNER JOIN table2
    ON table1.column = table2.column;
```

`INNER` is optional:

```sql
JOIN
```

means:

```sql
INNER JOIN
```

in standard usage.

---

## Example

```sql
SELECT
    e.name,
    d.department_name
FROM employees e
INNER JOIN departments d
    ON e.department_id = d.department_id;
```

Result:

| name | department_name |
|---|---|
| Alice | Engineering |
| Bob | HR |
| Charlie | Engineering |
| David | Sales |

Emma disappears because:

```text
Emma.department_id = NULL
```

and Finance disappears because it has no employee.

### Mental model

```text
A ∩ B
```

Only matching rows survive.

---

## When to Use INNER JOIN

Use it when you only care about records that have a valid match.

### Customers who placed orders

```sql
SELECT
    c.customer_name,
    o.order_id
FROM customers c
INNER JOIN orders o
    ON c.customer_id = o.customer_id;
```

### Employees belonging to departments

```sql
SELECT
    e.name,
    d.department_name
FROM employees e
JOIN departments d
    ON e.department_id = d.department_id;
```

### Products with categories

```sql
SELECT
    p.product_name,
    c.category_name
FROM products p
JOIN categories c
    ON p.category_id = c.category_id;
```

---

# 5. LEFT JOIN

Also called:

**LEFT OUTER JOIN**

Syntax:

```sql
SELECT ...
FROM A
LEFT JOIN B
    ON A.key = B.key;
```

The key rule:

> **Every row from the LEFT table is preserved.**

Matching rows from the right table are attached.

If there is no match:

```text
right-side columns = NULL
```

---

## Example

```sql
SELECT
    e.name,
    d.department_name
FROM employees e
LEFT JOIN departments d
    ON e.department_id = d.department_id;
```

Result:

| name | department_name |
|---|---|
| Alice | Engineering |
| Bob | HR |
| Charlie | Engineering |
| David | Sales |
| Emma | NULL |

Emma remains because `employees` is the left table.

---

## Extremely Common Use Case

> Find all customers, including customers who have never placed an order.

```sql
SELECT
    c.customer_id,
    c.customer_name,
    o.order_id
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id;
```

Customers without orders have:

```text
order_id = NULL
```

Then:

```sql
SELECT
    c.customer_id,
    c.customer_name
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
WHERE o.order_id IS NULL;
```

This is an **anti-join pattern**.

---

## LEFT JOIN Mental Model

```text
LEFT TABLE
   │
   ├── matching right rows → combine
   │
   └── no match → keep left row + NULLs
```

---

# 6. RIGHT JOIN

RIGHT JOIN is the mirror image of LEFT JOIN.

Syntax:

```sql
SELECT ...
FROM A
RIGHT JOIN B
    ON A.id = B.id;
```

The **right table is preserved**.

Example:

```sql
SELECT
    e.name,
    d.department_name
FROM employees e
RIGHT JOIN departments d
    ON e.department_id = d.department_id;
```

Result:

| name | department_name |
|---|---|
| Alice | Engineering |
| Charlie | Engineering |
| Bob | HR |
| David | Sales |
| NULL | Finance |

Finance remains because `departments` is the right table.

---

## LEFT JOIN vs RIGHT JOIN

These are logically equivalent if you reverse the table order.

This:

```sql
FROM employees e
LEFT JOIN departments d
    ON e.department_id = d.department_id
```

is equivalent in intent to:

```sql
FROM departments d
RIGHT JOIN employees e
    ON e.department_id = d.department_id
```

Most developers prefer **LEFT JOIN** because it is usually easier to read.

> You can almost always rewrite a RIGHT JOIN as a LEFT JOIN.

---

# 7. FULL OUTER JOIN

A:

```sql
FULL OUTER JOIN
```

returns:

> **All rows from both tables.**

Matched rows are combined.

Unmatched rows receive `NULL` on the missing side.

Example:

```sql
SELECT
    e.name,
    d.department_name
FROM employees e
FULL OUTER JOIN departments d
    ON e.department_id = d.department_id;
```

Result conceptually:

| name | department_name |
|---|---|
| Alice | Engineering |
| Bob | HR |
| Charlie | Engineering |
| David | Sales |
| Emma | NULL |
| NULL | Finance |

So:

```text
employees-only rows     → preserved
departments-only rows   → preserved
matching rows            → combined
```

### Mental model

```text
FULL OUTER JOIN
= matched rows
+ unmatched left rows
+ unmatched right rows
```

---

## Practical Use Case: Reconciliation

Suppose two systems contain transactions:

```text
system_a_transactions
system_b_transactions
```

You want to find:

- transactions in both systems
- transactions only in A
- transactions only in B

```sql
SELECT
    a.transaction_id AS a_id,
    b.transaction_id AS b_id
FROM system_a a
FULL OUTER JOIN system_b b
    ON a.transaction_id = b.transaction_id;
```

Then:

```sql
WHERE a.transaction_id IS NULL
   OR b.transaction_id IS NULL;
```

finds discrepancies.

Common uses:

- financial reconciliation
- data migration
- ETL validation
- data quality checks
- system synchronization

> **Note:** Some database systems, notably MySQL, do not provide native `FULL OUTER JOIN`. It can be emulated with combinations of `LEFT JOIN`, `RIGHT JOIN`, and `UNION`.

---

# 8. CROSS JOIN

A CROSS JOIN produces the **Cartesian product**.

Syntax:

```sql
SELECT *
FROM A
CROSS JOIN B;
```

Every row of A is paired with every row of B.

---

## Example

Table A:

| color |
|---|
| Red |
| Blue |

Table B:

| size |
|---|
| S |
| M |
| L |

Query:

```sql
SELECT *
FROM colors
CROSS JOIN sizes;
```

Result:

| color | size |
|---|---|
| Red | S |
| Red | M |
| Red | L |
| Blue | S |
| Blue | M |
| Blue | L |

Total:

```text
2 × 3 = 6 rows
```

---

## Cartesian Product Formula

If:

```text
A has M rows
B has N rows
```

then:

```text
A CROSS JOIN B
```

produces:

```text
M × N rows
```

Example:

```text
10,000 customers
×
100 products
=
1,000,000 rows
```

---

## Practical CROSS JOIN Use Cases

CROSS JOIN is useful when you intentionally need combinations.

### Product variants

```text
Colors × Sizes
```

### Calendar generation

```text
Dates × Stores
```

### Experiment combinations

```text
Users × Experiments
```

### Scheduling

```text
Employees × Available Shifts
```

Example:

```sql
SELECT
    e.employee_id,
    s.shift_date
FROM employees e
CROSS JOIN shifts s;
```

This generates every possible employee/shift combination.

---

# 9. SELF JOIN

A **SELF JOIN** means joining a table to itself.

There is no special `SELF JOIN` keyword.

You simply use aliases.

Example:

```sql
SELECT
    e.name AS employee,
    m.name AS manager
FROM employees e
LEFT JOIN employees m
    ON e.manager_id = m.employee_id;
```

Here:

```text
employees e
```

represents one logical copy, while:

```text
employees m
```

represents another.

---

## Example

`employees`:

| employee_id | name | manager_id |
|---:|---|---:|
| 1 | Alice | NULL |
| 2 | Bob | 1 |
| 3 | Charlie | 1 |
| 4 | David | 2 |

Query:

```sql
SELECT
    e.name AS employee,
    m.name AS manager
FROM employees e
LEFT JOIN employees m
    ON e.manager_id = m.employee_id;
```

Result:

| employee | manager |
|---|---|
| Alice | NULL |
| Bob | Alice |
| Charlie | Alice |
| David | Bob |

---

## SELF JOIN Use Cases

### 1. Employee hierarchy

```text
employee → manager → director → VP
```

### 2. Finding duplicate records

```sql
SELECT
    a.id,
    b.id
FROM users a
JOIN users b
    ON a.email = b.email
   AND a.id < b.id;
```

### 3. Comparing rows

Find products with the same price:

```sql
SELECT
    p1.name,
    p2.name
FROM products p1
JOIN products p2
    ON p1.price = p2.price
   AND p1.id < p2.id;
```

The condition:

```sql
p1.id < p2.id
```

prevents:

- matching a row with itself
- returning both `(A,B)` and `(B,A)`

---

# 10. Multiple JOINs

Real-world SQL frequently joins more than two tables.

Suppose:

```text
customers
    ↓
orders
    ↓
order_items
    ↓
products
    ↓
categories
```

Query:

```sql
SELECT
    c.customer_name,
    o.order_id,
    p.product_name,
    oi.quantity,
    cat.category_name
FROM customers c
JOIN orders o
    ON c.customer_id = o.customer_id
JOIN order_items oi
    ON o.order_id = oi.order_id
JOIN products p
    ON oi.product_id = p.product_id
JOIN categories cat
    ON p.category_id = cat.category_id;
```

---

## How to Think About Multiple JOINs

Conceptually:

```text
customers
    JOIN orders
         ↓
customer + order rows
         ↓
    JOIN order_items
         ↓
customer + order + item rows
         ↓
    JOIN products
         ↓
customer + order + item + product
         ↓
    JOIN categories
         ↓
final result
```

The critical question after every JOIN is:

> **What does one row represent now?**

That is the beginning of understanding **query grain**.

---

# 11. JOIN Conditions

The `ON` clause defines **how rows are related**.

Typical example:

```sql
ON e.department_id = d.department_id
```

This is an **equi-join** because it uses `=`.

JOIN conditions don't have to use equality.

---

## Non-Equi JOIN

You can use:

```text
>
<
>=
<=
<>
BETWEEN
```

Example: salary ranges.

`employees`:

| name | salary |
|---|---:|
| Alice | 70000 |
| Bob | 45000 |

`salary_grades`:

| grade | min_salary | max_salary |
|---|---:|---:|
| A | 60000 | 100000 |
| B | 40000 | 59999 |

Query:

```sql
SELECT
    e.name,
    e.salary,
    g.grade
FROM employees e
JOIN salary_grades g
    ON e.salary BETWEEN g.min_salary AND g.max_salary;
```

Result:

| name | salary | grade |
|---|---:|---|
| Alice | 70000 | A |
| Bob | 45000 | B |

This is a **range join**.

---

## Interval / Overlap JOIN

A useful advanced pattern:

```sql
ON a.start_date <= b.end_date
AND a.end_date >= b.start_date
```

This detects overlapping intervals.

Used in:

- scheduling
- booking systems
- temporal data
- employment periods
- subscription periods
- pricing validity periods

---

## Composite JOIN Conditions

Sometimes one column isn't enough.

Example:

```sql
ON a.student_id = b.student_id
AND a.course_id = b.course_id
```

Both conditions must match.

This is common when a logical relationship is identified by a **composite key**.

---

# 12. JOIN Filtering — ON vs WHERE

This is one of the most important SQL concepts.

Consider:

```sql
SELECT *
FROM employees e
LEFT JOIN departments d
    ON e.department_id = d.department_id
WHERE d.department_name = 'Engineering';
```

A common mistake is to think:

> "LEFT JOIN means all employees."

Not anymore.

The `WHERE` clause can remove rows that the LEFT JOIN preserved.

---

## Why?

Suppose Emma has:

```text
Emma | NULL
```

After the LEFT JOIN:

```text
Emma | NULL
```

Then:

```sql
WHERE d.department_name = 'Engineering'
```

requires:

```text
NULL = 'Engineering'
```

which is not TRUE.

So Emma disappears.

---

## Version A — Filter in ON

```sql
SELECT *
FROM employees e
LEFT JOIN departments d
    ON e.department_id = d.department_id
   AND d.department_name = 'Engineering';
```

All employees remain.

Conceptually:

```text
Alice    Engineering
Bob      NULL
Charlie  Engineering
David    NULL
Emma     NULL
```

---

## Version B — Filter in WHERE

```sql
SELECT *
FROM employees e
LEFT JOIN departments d
    ON e.department_id = d.department_id
WHERE d.department_name = 'Engineering';
```

Result:

```text
Alice    Engineering
Charlie  Engineering
```

---

## Critical Rule

For an OUTER JOIN:

```text
ON
```

controls **which rows match**.

```text
WHERE
```

controls **which final rows survive**.

This distinction is fundamental.

---

## INNER JOIN: ON vs WHERE

For INNER JOIN:

```sql
FROM A
JOIN B
    ON A.id = B.id
WHERE B.status = 'active'
```

versus:

```sql
FROM A
JOIN B
    ON A.id = B.id
   AND B.status = 'active'
```

often produce the same result.

For OUTER JOINs, moving predicates between `ON` and `WHERE` can **change the result**.

---

# 13. JOIN Ordering

There are two different concepts:

### 1. Written JOIN order

The order in which you write joins.

### 2. Physical execution order

The order selected by the database optimizer.

These are not necessarily the same.

---

## Query Optimizer

Suppose:

```text
customers = 10 million rows
orders    = 100 million rows
products  = 50,000 rows
```

You write:

```sql
SELECT ...
FROM customers c
JOIN orders o ...
JOIN products p ...
```

The database optimizer may internally execute these operations in a different order to reduce the amount of work.

It considers things such as:

- indexes
- table statistics
- estimated cardinality
- selectivity
- join algorithms
- available memory
- predicates

---

## JOIN Reordering

For INNER JOINs, the optimizer generally has substantial freedom to reorder joins because the operations are associative/commutative under relational semantics.

Conceptually:

```text
(A JOIN B) JOIN C
```

can often be rearranged as:

```text
A JOIN (B JOIN C)
```

But OUTER JOINs are different.

For example:

```text
A LEFT JOIN B
```

cannot simply be reordered with arbitrary operations without potentially changing the result.

---

# 14. Cartesian Product

A Cartesian product means:

> **Every row from A combines with every row from B.**

Mathematically:

```text
A × B
```

If:

```text
A = 5 rows
B = 10 rows
```

then:

```text
A × B = 50 rows
```

---

## Accidental Cartesian Products

A common bug is forgetting the relationship between tables.

Explicitly:

```sql
SELECT *
FROM employees e
CROSS JOIN departments d;
```

produces a Cartesian product.

Older comma syntax:

```sql
SELECT *
FROM employees e, departments d
WHERE e.department_id = d.department_id;
```

is effectively an INNER JOIN when the predicate is present.

But:

```sql
SELECT *
FROM employees e, departments d;
```

creates a Cartesian product.

---

## Why Cartesian Products Are Dangerous

Suppose:

```text
employees = 50,000
departments = 100
```

Then:

```text
50,000 × 100 = 5,000,000 rows
```

A larger accidental product can cause:

- huge memory usage
- slow queries
- CPU spikes
- disk spills
- temporary tables
- poor application performance

---

# 15. Row Multiplication

This is one of the most important real-world JOIN problems.

Suppose:

```text
customers
---------
Alice

orders
------
Order 1
Order 2

payments
--------
Payment 1
Payment 2
```

Query:

```sql
SELECT *
FROM customers c
JOIN orders o
    ON c.id = o.customer_id
JOIN payments p
    ON c.id = p.customer_id;
```

Alice has:

```text
2 orders
2 payments
```

The result can become:

```text
Order 1 | Payment 1
Order 1 | Payment 2
Order 2 | Payment 1
Order 2 | Payment 2
```

That's:

```text
2 × 2 = 4 rows
```

This is **row multiplication**.

---

## Why Row Multiplication Causes Bugs

Suppose:

```text
Order 1 = ₹100
Order 2 = ₹200
```

Correct order total:

```text
₹300
```

But after joining to two payments, each order may appear twice.

Then:

```text
₹100 × 2 + ₹200 × 2 = ₹600
```

An aggregation can therefore become incorrect.

---

## Fix: Aggregate Before Joining

A common approach is to aggregate each one-to-many relation separately:

```sql
SELECT
    c.id,
    o.total_orders,
    p.total_payments
FROM customers c
LEFT JOIN (
    SELECT
        customer_id,
        SUM(amount) AS total_orders
    FROM orders
    GROUP BY customer_id
) o
    ON c.id = o.customer_id
LEFT JOIN (
    SELECT
        customer_id,
        SUM(amount) AS total_payments
    FROM payments
    GROUP BY customer_id
) p
    ON c.id = p.customer_id;
```

Now each derived table produces at most one row per customer.

---

# 16. JOIN vs Subquery

This is a major SQL design question.

Suppose you want customers who have placed an order.

### JOIN

```sql
SELECT DISTINCT c.customer_id
FROM customers c
JOIN orders o
    ON c.customer_id = o.customer_id;
```

### EXISTS

```sql
SELECT c.customer_id
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

These express related but different ideas.

---

## JOIN Means "Combine Rows"

Use JOIN when you actually need data from both relations.

```sql
SELECT
    c.name,
    o.order_date
FROM customers c
JOIN orders o
    ON c.id = o.customer_id;
```

You need:

```text
customers
+
orders
```

---

## EXISTS Means "Does a Matching Row Exist?"

If you only need to know whether a customer has an order:

```sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

This avoids generating multiple output rows for a customer who has many orders.

---

## JOIN vs EXISTS

Suppose Alice has 5 orders.

JOIN can produce:

```text
Alice
Alice
Alice
Alice
Alice
```

If you only need a yes/no answer, that's unnecessary.

`EXISTS` naturally expresses:

```text
TRUE / FALSE
```

---

## JOIN vs IN

You could also write:

```sql
SELECT *
FROM customers
WHERE customer_id IN (
    SELECT customer_id
    FROM orders
);
```

Conceptually:

```text
IN       → membership
EXISTS   → existence
JOIN     → row combination
```

Modern optimizers may transform these into similar execution strategies, so avoid simplistic rules such as:

> "JOIN is always faster."

It isn't.

---

# 17. EXISTS, Semi-JOIN and Anti-JOIN

## Semi-JOIN

A **semi-join** returns rows from A when at least one matching row exists in B.

SQL representation:

```sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

Notice:

- customer rows are returned
- order rows are not returned
- a customer appears once even if they have many orders

---

## Anti-JOIN

An **anti-join** returns rows from A that have **no matching row** in B.

Two common SQL implementations:

### LEFT JOIN pattern

```sql
SELECT c.*
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
WHERE o.order_id IS NULL;
```

### NOT EXISTS pattern

```sql
SELECT c.*
FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

The `NOT EXISTS` form is often especially clear when the requirement is:

> "Return rows for which no related row exists."

---

# 18. JOIN Cardinality

**Cardinality** asks:

> How many rows can each row match?

Common relationships:

```text
1 → 1
1 → many
many → 1
many → many
```

Before writing a JOIN, identify the relationship.

---

## One-to-One

Example:

```text
users
user_profiles
```

Each user has one profile.

```sql
SELECT
    u.username,
    p.phone,
    p.address
FROM users u
JOIN user_profiles p
    ON u.user_id = p.user_id;
```

Potential cardinality:

```text
1 user → 1 profile
```

---

## One-to-Many

Very common:

```text
customer
   ↓
many orders
```

```sql
SELECT
    c.name,
    o.order_id
FROM customers c
JOIN orders o
    ON c.customer_id = o.customer_id;
```

One customer can appear multiple times:

```text
Alice | Order 1
Alice | Order 2
Alice | Order 3
```

This is expected.

---

## Many-to-Many

Example:

```text
students
courses
```

A student can take many courses, and a course can contain many students.

Use a junction table:

```text
students
   ↓
enrollments
   ↓
courses
```

Schema:

```text
students
---------
student_id
name

enrollments
-----------
student_id
course_id

courses
-------
course_id
name
```

Query:

```sql
SELECT
    s.name,
    c.name
FROM students s
JOIN enrollments e
    ON s.student_id = e.student_id
JOIN courses c
    ON e.course_id = c.course_id;
```

---

# 19. JOINs and NULL

SQL uses **three-valued logic**:

```text
TRUE
FALSE
UNKNOWN
```

A critical rule:

```sql
NULL = NULL
```

does **not** evaluate to TRUE.

It evaluates to UNKNOWN.

Therefore, this:

```sql
WHERE department_id = NULL
```

is wrong.

Use:

```sql
WHERE department_id IS NULL
```

or:

```sql
WHERE department_id IS NOT NULL
```

---

## NULL and JOINs

Suppose:

```text
employees.department_id = NULL
```

and:

```text
departments.department_id = NULL
```

Then:

```sql
ON e.department_id = d.department_id
```

does not match those rows.

This is a common source of confusion.

---

# 20. JOIN vs UNION

Do not confuse JOIN with UNION.

## JOIN

Combines columns **horizontally**.

```text
A                     B

id | name             id | salary
1  | Alice            1  | 50000
```

JOIN:

```text
id | name  | salary
1  | Alice | 50000
```

Think:

```text
JOIN → wider
```

---

## UNION

Combines rows **vertically**.

```text
A
---
1
2

B
---
3
4
```

UNION:

```text
1
2
3
4
```

Think:

```text
UNION → taller
```

---

# 21. JOIN Performance

JOIN performance depends on many factors.

Important factors include:

### 1. Indexes

Join keys may benefit from indexes.

### 2. Selectivity

How much does a predicate reduce the number of rows?

### 3. Cardinality

How many matching rows exist per row?

### 4. Data distribution

Are values evenly distributed or highly skewed?

### 5. Statistics

The optimizer relies heavily on statistics.

### 6. Join algorithm

Possible algorithms include:

- nested loop
- hash join
- merge join

### 7. Predicate placement

Especially important with OUTER JOINs.

### 8. Result size

Sometimes the query genuinely needs to produce millions of rows.

---

# 22. JOIN Algorithms

The database does not necessarily compare every row with every other row in a naive way.

Common JOIN algorithms:

1. **Nested Loop Join**
2. **Hash Join**
3. **Merge Join**

---

## Nested Loop Join

Conceptually:

```text
for each row in A:
    for each row in B:
        if A.key = B.key:
            output
```

Naively:

```text
O(A × B)
```

But indexes can make nested-loop joins extremely efficient.

---

## Index Nested Loop Join

Suppose:

```text
customers = 1 million
orders.customer_id = indexed
```

The database can:

```text
take one customer
      ↓
use index to find matching orders
      ↓
process next customer
```

Instead of scanning the entire orders table for every customer.

---

## Hash Join

Hash joins are commonly effective for equality joins.

For:

```sql
A JOIN B
ON A.id = B.id
```

the database can build a hash structure from one side:

```text
key → row(s)
```

Then scan the other side and perform hash lookups.

Conceptually:

```text
Build hash table
       ↓
Scan other table
       ↓
Hash lookup
       ↓
Output matches
```

---

## Merge Join

A merge join works particularly well when both inputs are sorted by the join key.

Example:

```text
A:
1
3
5
8

B:
2
3
5
9
```

The database walks through both sorted sequences and finds:

```text
3
5
```

This can be very efficient when sorted inputs are already available.

---

# 23. Indexes and JOINs

Suppose:

```sql
SELECT *
FROM customers c
JOIN orders o
    ON c.customer_id = o.customer_id;
```

Typically:

```text
customers.customer_id
```

is already indexed if it is a primary key.

But:

```text
orders.customer_id
```

may need an index.

For example:

```sql
CREATE INDEX idx_orders_customer_id
ON orders(customer_id);
```

Whether an index is beneficial depends on the database engine and workload, but foreign-key columns are common candidates.

---

## Important Performance Principle

Do not assume:

```text
JOIN = slow
Subquery = fast
```

or:

```text
JOIN = fast
Subquery = slow
```

Both are oversimplifications.

Use the **execution plan** when performance matters:

```sql
EXPLAIN
SELECT ...;
```

and, where supported:

```sql
EXPLAIN ANALYZE
SELECT ...;
```

---

# 24. Query Grain

This is one of the most important concepts in analytics SQL.

**Grain** means:

> **What does one row in the result represent?**

Examples:

```text
customer-level
→ one row = one customer

order-level
→ one row = one order

order-item-level
→ one row = one order item
```

---

## Why Grain Matters

Suppose your target is:

```text
one row per customer
```

But you join:

```text
customers
JOIN orders
JOIN order_items
```

Your intermediate result might become:

```text
one row per order item
```

If you then aggregate a customer-level value, you may accidentally multiply it.

Always ask:

> **What is the grain before and after every JOIN?**

This single habit prevents many serious SQL bugs.

---

# 25. Real-World JOIN Patterns

## Pattern 1 — Basic Relationship

```sql
SELECT ...
FROM A
JOIN B
    ON A.id = B.a_id;
```

Use when only matching rows matter.

---

## Pattern 2 — Preserve Everything From A

```sql
SELECT ...
FROM A
LEFT JOIN B
    ON A.id = B.a_id;
```

Use when every A row must appear.

---

## Pattern 3 — Find Missing B

```sql
SELECT A.*
FROM A
LEFT JOIN B
    ON A.id = B.a_id
WHERE B.id IS NULL;
```

Use for:

- customers without orders
- products never sold
- employees without departments

---

## Pattern 4 — Existence

```sql
SELECT A.*
FROM A
WHERE EXISTS (
    SELECT 1
    FROM B
    WHERE B.a_id = A.id
);
```

Use when you only care whether a matching row exists.

---

## Pattern 5 — Non-Existence

```sql
SELECT A.*
FROM A
WHERE NOT EXISTS (
    SELECT 1
    FROM B
    WHERE B.a_id = A.id
);
```

Use when no matching row should exist.

---

## Pattern 6 — Hierarchy

```sql
SELECT
    e.name,
    m.name
FROM employees e
LEFT JOIN employees m
    ON e.manager_id = m.employee_id;
```

Use for:

- employee-manager structures
- parent-child relationships
- organizational trees

---

## Pattern 7 — Many-to-Many

```sql
SELECT ...
FROM A
JOIN junction j
    ON A.id = j.a_id
JOIN B
    ON j.b_id = B.id;
```

Use for:

- students ↔ courses
- users ↔ roles
- products ↔ tags
- posts ↔ categories

---

# 26. JOIN Cheat Sheet

| JOIN | Matching | Left unmatched | Right unmatched |
|---|---|---|---|
| `INNER JOIN` | ✅ | ❌ | ❌ |
| `LEFT JOIN` | ✅ | ✅ | ❌ |
| `RIGHT JOIN` | ✅ | ❌ | ✅ |
| `FULL OUTER JOIN` | ✅ | ✅ | ✅ |
| `CROSS JOIN` | Every combination | N/A | N/A |
| `SELF JOIN` | Depends on join type | Depends | Depends |

---

# 27. LEFT JOIN for Missing Data

Memorize this pattern:

```sql
SELECT a.*
FROM A a
LEFT JOIN B b
    ON a.id = b.a_id
WHERE b.id IS NULL;
```

It means:

> **Find rows in A that have no corresponding row in B.**

### Customers without orders

```sql
SELECT c.*
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
WHERE o.order_id IS NULL;
```

### Products never sold

```sql
SELECT p.*
FROM products p
LEFT JOIN order_items oi
    ON p.product_id = oi.product_id
WHERE oi.product_id IS NULL;
```

### Employees without departments

```sql
SELECT e.*
FROM employees e
LEFT JOIN departments d
    ON e.department_id = d.department_id
WHERE d.department_id IS NULL;
```

---

# 28. LEFT JOIN + COUNT

Question:

> Show every customer and their number of orders, including customers with zero orders.

Correct:

```sql
SELECT
    c.customer_id,
    c.name,
    COUNT(o.order_id) AS order_count
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
GROUP BY
    c.customer_id,
    c.name;
```

Why:

```sql
COUNT(o.order_id)
```

instead of:

```sql
COUNT(*)
```

For a customer without orders:

```text
o.order_id = NULL
```

and:

```sql
COUNT(o.order_id)
```

returns:

```text
0
```

But:

```sql
COUNT(*)
```

counts the preserved LEFT JOIN row and can return `1`.

This distinction is extremely important.

---

# 29. JOIN Filtering With Dates

Suppose:

> Show every customer, but only attach orders from 2026.

Correct:

```sql
SELECT
    c.name,
    o.order_id,
    o.order_date
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
   AND o.order_date >= '2026-01-01'
   AND o.order_date < '2027-01-01';
```

This preserves customers who had no 2026 orders.

If you instead write:

```sql
SELECT
    c.name,
    o.order_id,
    o.order_date
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
WHERE o.order_date >= '2026-01-01'
  AND o.order_date < '2027-01-01';
```

customers without qualifying orders are removed.

### Rule

If the requirement is:

> **"Keep the left entity even when no matching filtered row exists."**

the right-table filter often belongs in `ON`.

---

# 30. Duplicate Rows After JOIN

Suppose:

```text
departments
-----------
10 | Engineering

employees
---------
1 | Alice   | 10
2 | Bob     | 10
3 | Charlie | 10
```

Query:

```sql
SELECT *
FROM departments d
JOIN employees e
    ON d.department_id = e.department_id;
```

Engineering appears three times.

That is not necessarily a SQL bug.

It represents the real relationship:

```text
1 department → 3 employees
```

The result is correctly representing the one-to-many relationship.

---

## Don't Automatically Use DISTINCT

A common beginner reaction:

> "My JOIN gives duplicates, so I'll use DISTINCT."

That can hide the real problem.

Instead ask:

> **Why are multiple rows being generated?**

If you only need existence:

```sql
WHERE EXISTS (...)
```

may be better.

If you need one row per entity:

```sql
GROUP BY
```

may be appropriate.

If the relationship is genuinely one-to-many, multiple rows may be correct.

---

# 31. JOIN Conditions and NULL

Suppose:

```text
employee.department_id = NULL
```

You might expect:

```text
NULL = NULL
```

to be TRUE.

It isn't.

SQL uses three-valued logic:

```text
TRUE
FALSE
UNKNOWN
```

Therefore:

```sql
NULL = NULL
```

produces UNKNOWN.

This is why:

```sql
WHERE department_id = NULL
```

does not work.

Use:

```sql
WHERE department_id IS NULL;
```

---

# 32. Joining on Multiple Conditions

Example:

```sql
SELECT *
FROM orders o
JOIN shipments s
    ON o.order_id = s.order_id
   AND o.customer_id = s.customer_id;
```

Both conditions must be satisfied.

Think:

```text
condition1 AND condition2
```

This is useful when matching on a composite logical key.

---

# 33. OR in JOIN Conditions

You can technically write:

```sql
JOIN B
    ON A.x = B.x
    OR A.y = B.y
```

But be careful.

OR conditions can:

- produce unexpected matches
- multiply rows
- complicate optimization
- make indexing less effective
- make the query harder to reason about

When possible, investigate whether the relationship can be expressed more cleanly.

---

# 34. Joining on Expressions

You can write:

```sql
ON LOWER(a.email) = LOWER(b.email)
```

or:

```sql
ON DATE(a.created_at) = b.date
```

But functions applied to columns can sometimes prevent efficient use of ordinary indexes, depending on the database.

Possible alternatives include:

- normalized stored values
- generated/computed columns
- functional indexes
- better schema design

---

# 35. Foreign Keys and JOINs

A common relationship is:

```text
departments.department_id
        ↑
        │
employees.department_id
```

Usually:

```text
departments.department_id
```

is a **PRIMARY KEY**.

And:

```text
employees.department_id
```

is a **FOREIGN KEY**.

Then:

```sql
JOIN departments d
    ON e.department_id = d.department_id
```

follows the relational relationship.

---

# 36. A Real E-Commerce Query

Requirement:

> Find customers who purchased laptops in 2026 and show how much they spent.

Tables:

```text
customers
orders
order_items
products
categories
```

Query:

```sql
SELECT
    c.customer_id,
    c.name,
    SUM(oi.quantity * oi.unit_price) AS total_spent
FROM customers c
JOIN orders o
    ON c.customer_id = o.customer_id
JOIN order_items oi
    ON o.order_id = oi.order_id
JOIN products p
    ON oi.product_id = p.product_id
JOIN categories cat
    ON p.category_id = cat.category_id
WHERE o.order_date >= '2026-01-01'
  AND o.order_date < '2027-01-01'
  AND cat.category_name = 'Laptops'
GROUP BY
    c.customer_id,
    c.name
ORDER BY total_spent DESC;
```

This demonstrates:

- multiple JOINs
- join conditions
- filtering
- aggregation
- grouping
- ordering
- one-to-many relationships

---

# 37. Analytics Example — Every Department

Requirement:

> Show every department and the number of employees, including empty departments.

```sql
SELECT
    d.department_id,
    d.department_name,
    COUNT(e.employee_id) AS employee_count
FROM departments d
LEFT JOIN employees e
    ON d.department_id = e.department_id
GROUP BY
    d.department_id,
    d.department_name
ORDER BY employee_count DESC;
```

Why start with:

```sql
FROM departments d
```

instead of:

```sql
FROM employees e
```

Because the requirement is:

> **Every department must appear.**

The **driving table** should usually represent the population that must be preserved.

---

# 38. JOIN Debugging Strategy

When a JOIN produces an unexpected result:

## Step 1 — Inspect each table independently

```sql
SELECT *
FROM customers;
```

## Step 2 — Check row counts

```sql
SELECT COUNT(*)
FROM customers;
```

## Step 3 — Check key uniqueness

```sql
SELECT
    customer_id,
    COUNT(*)
FROM customers
GROUP BY customer_id
HAVING COUNT(*) > 1;
```

## Step 4 — Add one JOIN

```sql
SELECT ...
FROM A
JOIN B
    ON ...;
```

## Step 5 — Check the row count

If the count unexpectedly explodes, investigate the cardinality.

## Step 6 — Add the next JOIN

Repeat.

### Best practice

Do not write an 8-table JOIN and debug everything at once.

Build it incrementally.

---

# 39. The Grain of a Query

Before writing a complex JOIN, explicitly state:

```text
One result row represents __________.
```

Examples:

```text
one row = one customer
one row = one order
one row = one order item
one row = one customer per month
one row = one product per category
```

Then make sure every JOIN and aggregation preserves or intentionally changes that grain.

This is one of the most powerful habits in SQL analytics.

---

# 40. Advanced Performance Principle — Selectivity

**Selectivity** describes how effectively a predicate reduces the number of rows.

Suppose:

```text
orders = 100 million rows
```

and:

```sql
WHERE status = 'CANCELLED'
```

reduces that to:

```text
100,000 rows
```

That filter is highly selective.

Highly selective predicates can dramatically reduce the amount of data that later operations need to process.

Modern optimizers may perform **predicate pushdown**, applying filters as early as possible when doing so is semantically valid.

---

# 41. Advanced Performance Principle — Cardinality Estimates

The optimizer tries to estimate:

```text
How many rows will this operation produce?
```

For example:

```text
customers
   ↓ JOIN
orders
```

If the optimizer incorrectly estimates the number of matching rows, it may choose a poor execution plan.

Statistics help the optimizer estimate:

- number of rows
- distinct values
- data distribution
- selectivity

This is one reason keeping database statistics current matters.

---

# 42. Advanced Performance Principle — Intermediate Results

A query can be slow not because the final result is large, but because an intermediate JOIN creates a huge temporary result.

Example:

```text
customers
JOIN orders
JOIN order_items
JOIN payments
```

Even if the final query returns:

```text
1,000 customers
```

the intermediate result could contain millions of rows.

This is why **cardinality and grain** matter so much.

---

# 43. A Practical Decision Framework

Before writing a JOIN, ask:

### Question 1 — Do I need columns from another table?

If yes:

```text
JOIN
```

### Question 2 — Do I only need to know whether a related row exists?

Consider:

```text
EXISTS
```

### Question 3 — Do I need rows with no match?

Use:

```text
LEFT JOIN
```

or:

```text
FULL OUTER JOIN
```

depending on which side(s) must be preserved.

### Question 4 — Do I need every possible combination?

Use:

```text
CROSS JOIN
```

### Question 5 — Is the relationship hierarchical?

Consider:

```text
SELF JOIN
```

or, for arbitrary-depth hierarchies, a recursive CTE.

### Question 6 — Is this a many-to-many relationship?

Look for:

```text
junction / bridge table
```

### Question 7 — Could the JOIN multiply rows?

Check cardinality before aggregating.

---

# 44. The Most Important JOIN Patterns to Memorize

## Basic relationship

```sql
SELECT ...
FROM A
JOIN B
    ON A.id = B.a_id;
```

## Keep everything from A

```sql
SELECT ...
FROM A
LEFT JOIN B
    ON A.id = B.a_id;
```

## Find rows in A with no B

```sql
SELECT A.*
FROM A
LEFT JOIN B
    ON A.id = B.a_id
WHERE B.id IS NULL;
```

## Existence

```sql
SELECT A.*
FROM A
WHERE EXISTS (
    SELECT 1
    FROM B
    WHERE B.a_id = A.id
);
```

## Non-existence

```sql
SELECT A.*
FROM A
WHERE NOT EXISTS (
    SELECT 1
    FROM B
    WHERE B.a_id = A.id
);
```

## Hierarchy

```sql
SELECT
    e.name AS employee,
    m.name AS manager
FROM employees e
LEFT JOIN employees m
    ON e.manager_id = m.employee_id;
```

## Many-to-many

```sql
SELECT ...
FROM A
JOIN junction j
    ON A.id = j.a_id
JOIN B
    ON j.b_id = B.id;
```

---

# 45. Final Practical Example

Requirement:

> Show every customer, their total number of orders, total amount spent on completed orders, and latest order date. Customers with no orders should still appear.

```sql
SELECT
    c.customer_id,
    c.customer_name,
    COUNT(o.order_id) AS total_orders,
    COALESCE(
        SUM(
            CASE
                WHEN o.status = 'COMPLETED'
                THEN o.amount
                ELSE 0
            END
        ),
        0
    ) AS completed_amount,
    MAX(o.order_date) AS latest_order_date
FROM customers c
LEFT JOIN orders o
    ON c.customer_id = o.customer_id
GROUP BY
    c.customer_id,
    c.customer_name;
```

This one query combines:

```text
LEFT JOIN
+
JOIN condition
+
NULL handling
+
conditional aggregation
+
COUNT
+
SUM
+
MAX
+
GROUP BY
+
query grain
```

### Why LEFT JOIN?

Because:

```text
Every customer must appear.
```

### Why `COUNT(o.order_id)`?

Because customers without orders should get:

```text
0
```

instead of being counted as one preserved JOIN row.

### Why `CASE`?

Because only completed orders should contribute to the completed amount.

### Why `COALESCE`?

Because an aggregate can return `NULL` when there are no qualifying values, while the business requirement is usually to display `0`.

---

# 46. Final JOIN Cheat Sheet

```text
                    SQL JOINs
                       │
        ┌──────────────┼───────────────┐
        │              │               │
     Matching      Preservation     Combination
        │              │               │
        ↓              ↓               ↓
   INNER JOIN      LEFT JOIN       CROSS JOIN
                   RIGHT JOIN
                   FULL JOIN
```

### INNER JOIN

```text
Only matching rows
```

### LEFT JOIN

```text
Everything from left
+ matching right rows
```

### RIGHT JOIN

```text
Everything from right
+ matching left rows
```

### FULL OUTER JOIN

```text
Everything from both sides
```

### CROSS JOIN

```text
Every possible combination
```

### SELF JOIN

```text
A table joined to itself
```

### Multiple JOINs

```text
Chain related tables together
```

### EXISTS

```text
Does a related row exist?
```

### NOT EXISTS

```text
Does no related row exist?
```

---

# 47. Interview-Level Comparison

| Requirement | Best starting point |
|---|---|
| Only matching records | `INNER JOIN` |
| Keep every left record | `LEFT JOIN` |
| Keep every right record | `RIGHT JOIN` |
| Keep everything from both | `FULL OUTER JOIN` |
| Every combination | `CROSS JOIN` |
| Hierarchical relationship | `SELF JOIN` |
| Multiple related entities | Multiple `JOIN`s |
| Check whether related row exists | `EXISTS` |
| Find rows with no related row | `NOT EXISTS` / anti-join |
| Reconcile two datasets | `FULL OUTER JOIN` |
| Generate combinations | `CROSS JOIN` |
| Avoid one-to-many multiplication when only checking existence | `EXISTS` |

---

# 48. The Deepest Mental Model

If you remember only one framework, use this:

```text
                         JOIN
                           │
             ┌─────────────┴─────────────┐
             │                           │
       RELATIONSHIP                  CARDINALITY
             │                           │
         ON clause                 1:1 / 1:N / N:N
             │                           │
             └─────────────┬─────────────┘
                           ↓
                     ROW MATCHING
                           ↓
                WHAT GETS PRESERVED?
                           │
             ┌─────────────┼─────────────┐
             ↓             ↓             ↓
          INNER          LEFT          FULL
        matches only   preserve A   preserve A+B
                           │
                           ↓
                      WHERE clause
                           │
                           ↓
                   final row filtering
```

For every JOIN, ask:

```text
1. What is the relationship?
2. What is the join key?
3. What is the cardinality?
4. Which rows must be preserved?
5. Where should filtering happen?
6. What is the grain of the result?
7. Could the JOIN multiply rows?
8. Do I actually need a JOIN, or is EXISTS better?
```

---

# 49. What You Should Master

## Tier 1 — Essential

- `INNER JOIN`
- `LEFT JOIN`
- `ON` vs `WHERE`
- primary key / foreign key relationships
- one-to-many relationships
- NULL behavior
- finding unmatched records

## Tier 2 — Strong SQL

- `RIGHT JOIN`
- `FULL OUTER JOIN`
- `CROSS JOIN`
- `SELF JOIN`
- multiple JOINs
- composite join conditions
- non-equi joins
- `EXISTS` / `NOT EXISTS`
- semi-joins
- anti-joins

## Tier 3 — Advanced

- cardinality
- query grain
- row multiplication
- join selectivity
- join algorithms
- hash join
- nested-loop join
- merge join
- optimizer join reordering
- execution plans
- index strategy
- predicate pushdown
- reconciliation patterns

---

# 50. The Single Most Important Habit

> ## Before every JOIN, determine the grain and cardinality.

If you know:

```text
customer → orders = 1:N
order → order_items = 1:N
product → category = N:1
student ↔ course = N:N
```

you can predict how your result set will behave **before executing the query**.

That is the point where SQL JOINs stop being syntax to memorize and become something you can reason about systematically.

---

## Quick Reference

```text
INNER JOIN
    → matching rows only

LEFT JOIN
    → preserve left table

RIGHT JOIN
    → preserve right table

FULL OUTER JOIN
    → preserve both tables

CROSS JOIN
    → Cartesian product

SELF JOIN
    → table joined to itself

ON
    → defines matching relationship

WHERE
    → filters final result

EXISTS
    → matching row exists

NOT EXISTS
    → matching row does not exist

CARDINALITY
    → how many rows can match?

GRAIN
    → what does one output row represent?

CARTESIAN PRODUCT
    → every row × every row

ROW MULTIPLICATION
    → one-to-many/many-to-many JOINs can increase rows

EXPLAIN
    → inspect the optimizer's execution plan
```

---

> **Core principle:** JOINs are fundamentally about **relationships, preservation, cardinality, and grain**. Once those four ideas are clear, the SQL syntax becomes much easier to reason about.
