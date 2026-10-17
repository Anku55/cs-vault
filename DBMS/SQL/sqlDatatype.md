# 2. SQL Data Types

SQL data types define what kind of data a column can store.

> **Note:** This section primarily follows **PostgreSQL** terminology. Some data types and operators, such as `JSONB`, `ARRAY`, `ILIKE`, and `TIMESTAMPTZ`, are PostgreSQL-specific or have different behavior across database systems.

---

## 2.1 Numeric Data Types

### INTEGER

Stores whole numbers.

```sql
CREATE TABLE students (
    age INTEGER
);
```

Example:

```text
10
21
100
-5
```

---

### BIGINT

Stores large whole numbers.

```sql
CREATE TABLE accounts (
    account_id BIGINT
);
```

Useful when `INTEGER` is not large enough, such as large IDs or counters.

---

### SMALLINT

Stores smaller whole numbers using less storage.

```sql
CREATE TABLE ratings (
    rating SMALLINT
);
```

Useful for values with a relatively small range.

---

### DECIMAL

Stores exact decimal numbers.

```sql
CREATE TABLE products (
    price DECIMAL(10, 2)
);
```

Example:

```text
999.99
1250.50
```

`DECIMAL(10,2)` means:

```text
10 → total number of digits
 2 → digits after decimal point
```

> 💡 Use `DECIMAL` for values such as **money** where exact precision matters.

---

### NUMERIC

`NUMERIC` is equivalent to `DECIMAL` in PostgreSQL.

```sql
price NUMERIC(10, 2)
```

Both are useful for exact numeric calculations.

---

### REAL

Stores floating-point numbers with approximately 6 decimal digits of precision.

```sql
temperature REAL;
```

Example:

```text
36.5
98.6
```

Because it uses floating-point representation, it should generally not be used for exact financial calculations.

---

### DOUBLE PRECISION

Stores floating-point numbers with higher precision than `REAL`.

```sql
distance DOUBLE PRECISION;
```

Useful for scientific calculations and measurements where approximate floating-point values are acceptable.

---

## 2.2 Boolean

### BOOLEAN

Stores `TRUE` or `FALSE`.

```sql
CREATE TABLE users (
    is_active BOOLEAN
);
```

Example:

```sql
INSERT INTO users (is_active)
VALUES (TRUE);
```

Possible values:

```text
TRUE
FALSE
NULL
```

---

# 2.3 Character / String Data Types

### CHAR

Stores a fixed-length string.

```sql
code CHAR(5);
```

Example:

```text
'ABCDE'
```

Useful when values always have a fixed length.

---

### VARCHAR

Stores variable-length strings.

```sql
name VARCHAR(100);
```

Example:

```text
'Ankush'
'Rahul'
'Computer Science'
```

`VARCHAR(100)` allows up to 100 characters.

---

### TEXT

Stores variable-length text.

```sql
description TEXT;
```

Useful for longer text such as:

- Descriptions
- Comments
- Articles
- Messages

Example:

```sql
CREATE TABLE posts (
    content TEXT
);
```

---

# 2.4 Date & Time Data Types

### DATE

Stores a calendar date.

```sql
birth_date DATE;
```

Example:

```text
2026-10-17
```

---

### TIME

Stores a time of day.

```sql
start_time TIME;
```

Example:

```text
14:30:00
```

---

### TIMESTAMP

Stores date and time **without timezone information**.

```sql
created_at TIMESTAMP;
```

Example:

```text
2026-10-17 14:30:00
```

---

### TIMESTAMPTZ

Stores a timestamp with timezone semantics.

```sql
created_at TIMESTAMPTZ;
```

Example:

```text
2026-10-17 14:30:00+05:30
```

Useful when building applications that operate across different time zones.

> ⚠️ PostgreSQL's `TIMESTAMPTZ` does not literally preserve the original timezone label/offset. PostgreSQL stores the instant and converts it for display according to the session timezone.

---

# 2.5 UUID

UUID stands for **Universally Unique Identifier**.

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY
);
```

Example:

```text
550e8400-e29b-41d4-a716-446655440000
```

Useful for distributed systems and applications where globally unique identifiers are desirable.

Example using PostgreSQL:

```sql
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT
);
```

---

# 2.6 JSON

Stores JSON data.

```sql
CREATE TABLE users (
    profile JSON
);
```

Example:

```json
{
    "name": "Ankush",
    "age": 21,
    "city": "Delhi"
}
```

Useful when data has a flexible or semi-structured format.

---

# 2.7 JSONB

`JSONB` stores JSON in PostgreSQL's binary/decomposed format.

```sql
CREATE TABLE users (
    profile JSONB
);
```

Example:

```sql
INSERT INTO users (profile)
VALUES (
    '{"name": "Ankush", "age": 21}'
);
```

`JSONB` is generally preferred over `JSON` in PostgreSQL when you need to query, index, or manipulate JSON data frequently.

---

# 2.8 ARRAY

Stores multiple values in a single column.

```sql
CREATE TABLE students (
    name TEXT,
    skills TEXT[]
);
```

Insert:

```sql
INSERT INTO students (name, skills)
VALUES (
    'Ankush',
    ARRAY['C', 'C++', 'SQL', 'Python']
);
```

You can also use:

```sql
skills TEXT[] DEFAULT '{}';
```

Example value:

```text
{C,C++,SQL,Python}
```

> 💡 Arrays are useful in PostgreSQL, but for relational data, a separate table is often preferable when you need to query individual elements extensively.

---

# 2.9 ENUM

Defines a custom type with a fixed set of allowed values.

```sql
CREATE TYPE user_status AS ENUM (
    'active',
    'inactive',
    'banned'
);
```

Use it:

```sql
CREATE TABLE users (
    status user_status
);
```

Valid:

```sql
INSERT INTO users (status)
VALUES ('active');
```

Invalid:

```sql
INSERT INTO users (status)
VALUES ('unknown');
```

> ⚠️ ENUMs are convenient for small, stable sets of values. For frequently changing categories, a lookup table can be more flexible.

---

# 3. Constraints

Constraints enforce rules on table data.

They help maintain:

- Data integrity
- Data consistency
- Valid relationships
- Valid values
- Uniqueness

---

## 3.1 PRIMARY KEY

Uniquely identifies each row.

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100)
);
```

A primary key:

- Must be unique
- Cannot be `NULL`
- Identifies a row

Example:

```text
id | name
---+--------
1  | Ankush
2  | Rahul
3  | Priya
```

---

# 3.2 FOREIGN KEY

Creates a relationship between two tables.

```sql
CREATE TABLE departments (
    id INTEGER PRIMARY KEY,
    name TEXT
);
```

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name TEXT,
    department_id INTEGER,

    FOREIGN KEY (department_id)
        REFERENCES departments(id)
);
```

Relationship:

```text
departments
     │
     │ id
     ↓
students.department_id
```

---

# 3.3 UNIQUE

Ensures that values are unique.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE
);
```

This prevents duplicate email values.

---

# 3.4 NOT NULL

Prevents a column from containing `NULL`.

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);
```

This is invalid:

```sql
INSERT INTO students (id, name)
VALUES (1, NULL);
```

---

# 3.5 CHECK

Ensures that a condition is satisfied.

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    age INTEGER CHECK (age >= 18)
);
```

Invalid:

```sql
INSERT INTO students (id, age)
VALUES (1, 15);
```

Another example:

```sql
salary NUMERIC CHECK (salary >= 0)
```

---

# 3.6 DEFAULT

Provides a default value when one isn't supplied.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    is_active BOOLEAN DEFAULT TRUE
);
```

Then:

```sql
INSERT INTO users (id)
VALUES (1);
```

The database automatically uses:

```text
is_active = TRUE
```

---

# 3.7 Composite Primary Key

A primary key consisting of multiple columns.

```sql
CREATE TABLE enrollments (
    student_id INTEGER,
    course_id INTEGER,

    PRIMARY KEY (student_id, course_id)
);
```

The combination must be unique:

```text
student_id | course_id
-----------+----------
1          | 101
1          | 102
2          | 101
```

But this would be a duplicate:

```text
1 | 101
1 | 101  ← duplicate combination
```

Commonly used for **many-to-many relationships**.

---

# 3.8 Composite UNIQUE Constraint

A `UNIQUE` constraint involving multiple columns.

```sql
CREATE TABLE schedules (
    room_id INTEGER,
    schedule_date DATE,

    UNIQUE (room_id, schedule_date)
);
```

This prevents the same room from being assigned to the same date more than once.

---

# 3.9 ON DELETE CASCADE

Automatically deletes related child rows when the parent row is deleted.

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
```

If:

```text
users
  ↓
User 1 deleted
```

then related orders for that user are also deleted.

```text
User
 ↓
Orders
 ↓
Deleted automatically
```

> ⚠️ Use `CASCADE` carefully because deleting a parent can delete many child rows.

---

# 3.10 ON DELETE SET NULL

Sets the foreign key to `NULL` when the referenced parent is deleted.

```sql
CREATE TABLE employees (
    id INTEGER PRIMARY KEY,
    manager_id INTEGER,

    FOREIGN KEY (manager_id)
        REFERENCES employees(id)
        ON DELETE SET NULL
);
```

If the manager is deleted:

```text
manager_id
    ↓
NULL
```

The employee remains.

> The foreign key column must allow `NULL`.

---

# 3.11 ON DELETE RESTRICT

Prevents deletion of a referenced parent when dependent rows exist.

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT
);
```

If a user has orders, deleting that user is rejected.

Conceptually:

```text
Parent has children?
        │
     YES
        ↓
DELETE rejected
```

This is useful when child records must not become orphaned.

---

# 3.12 ON UPDATE CASCADE

Automatically updates referencing foreign keys when the referenced key changes.

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON UPDATE CASCADE
);
```

If:

```text
users.id
100 → 200
```

then:

```text
orders.user_id
100 → 200
```

automatically.

> 💡 In practice, primary keys are usually stable, so `ON UPDATE CASCADE` is less frequently needed than `ON DELETE` actions.

---

# 3.13 Constraint Violations

A constraint violation occurs when data breaks a defined rule.

### PRIMARY KEY violation

```sql
INSERT INTO students (id, name)
VALUES (1, 'Ankush');

INSERT INTO students (id, name)
VALUES (1, 'Rahul');
```

The second insert violates the primary key.

---

### UNIQUE violation

```sql
INSERT INTO users (email)
VALUES ('test@gmail.com');

INSERT INTO users (email)
VALUES ('test@gmail.com');
```

The second insert violates `UNIQUE`.

---

### NOT NULL violation

```sql
INSERT INTO students (name)
VALUES (NULL);
```

If `name` is `NOT NULL`, the query fails.

---

### CHECK violation

```sql
INSERT INTO students (age)
VALUES (-10);
```

If:

```sql
CHECK (age >= 0)
```

exists, the insert fails.

---

### FOREIGN KEY violation

```sql
INSERT INTO orders (user_id)
VALUES (999);
```

If user `999` doesn't exist, the foreign key constraint can reject the insert.

---

# 4. SQL Filtering

Filtering determines **which rows should be returned**.

Most filtering conditions are used with `WHERE`.

---

# 4.1 Comparison Operators

Common comparison operators:

| Operator | Meaning |
|---|---|
| `=` | Equal |
| `<>` | Not equal |
| `!=` | Not equal |
| `>` | Greater than |
| `<` | Less than |
| `>=` | Greater than or equal |
| `<=` | Less than or equal |

Example:

```sql
SELECT *
FROM students
WHERE age >= 18;
```

---

# 4.2 AND

All conditions must be true.

```sql
SELECT *
FROM students
WHERE age >= 18
  AND city = 'Delhi';
```

Meaning:

```text
age >= 18
AND
city = Delhi
```

---

# 4.3 OR

At least one condition must be true.

```sql
SELECT *
FROM students
WHERE city = 'Delhi'
   OR city = 'Mumbai';
```

---

# 4.4 NOT

Negates a condition.

```sql
SELECT *
FROM students
WHERE NOT city = 'Delhi';
```

Equivalent in this simple case:

```sql
SELECT *
FROM students
WHERE city <> 'Delhi';
```

`NOT` becomes particularly useful with operators such as:

```sql
NOT IN
NOT BETWEEN
NOT LIKE
NOT EXISTS
```

---

# 4.5 IN

Checks whether a value exists in a list.

Instead of:

```sql
SELECT *
FROM students
WHERE city = 'Delhi'
   OR city = 'Mumbai'
   OR city = 'Pune';
```

Use:

```sql
SELECT *
FROM students
WHERE city IN ('Delhi', 'Mumbai', 'Pune');
```

Much cleaner.

---

# 4.6 NOT IN

Checks that a value does **not** exist in a list.

```sql
SELECT *
FROM students
WHERE city NOT IN ('Delhi', 'Mumbai');
```

> ⚠️ Be careful when the compared column or list contains `NULL`. `NOT IN` and `NULL` can produce unintuitive results because SQL uses three-valued logic.

---

# 4.7 BETWEEN

Checks whether a value falls within an inclusive range.

```sql
SELECT *
FROM students
WHERE age BETWEEN 18 AND 25;
```

This includes:

```text
18
19
20
...
24
25
```

Equivalent conceptually to:

```sql
WHERE age >= 18
  AND age <= 25
```

---

# 4.8 LIKE

Performs pattern matching.

## `%`

Matches zero or more characters.

```sql
SELECT *
FROM students
WHERE name LIKE 'A%';
```

Matches:

```text
Ankush
Aman
Aarav
```

---

## `_`

Matches exactly one character.

```sql
SELECT *
FROM students
WHERE name LIKE 'A___';
```

Matches four-character names beginning with `A`.

Example:

```text
Aman
```

---

### More LIKE Examples

Starts with `A`:

```sql
WHERE name LIKE 'A%'
```

Ends with `sh`:

```sql
WHERE name LIKE '%sh'
```

Contains `an`:

```sql
WHERE name LIKE '%an%'
```

Exactly 5 characters:

```sql
WHERE name LIKE '_____'
```

---

# 4.9 ILIKE

`ILIKE` performs **case-insensitive pattern matching** in PostgreSQL.

```sql
SELECT *
FROM students
WHERE name ILIKE 'ankush';
```

This can match:

```text
Ankush
ANKUSH
ankush
AnKuSh
```

Pattern matching also works:

```sql
SELECT *
FROM students
WHERE name ILIKE 'a%';
```

> `ILIKE` is PostgreSQL-specific. Other databases may use different approaches for case-insensitive matching.

---

# 4.10 EXISTS

Checks whether a subquery returns **at least one row**.

Example:

```sql
SELECT *
FROM students s
WHERE EXISTS (
    SELECT 1
    FROM enrollments e
    WHERE e.student_id = s.id
);
```

Meaning:

```text
Does this student have at least one enrollment?
        │
    ┌───┴───┐
   YES      NO
    │        │
 RETURN    IGNORE
```

`EXISTS` is commonly used with **correlated subqueries**.

---

# 4.11 NOT EXISTS

Checks whether a subquery returns **no rows**.

```sql
SELECT *
FROM students s
WHERE NOT EXISTS (
    SELECT 1
    FROM enrollments e
    WHERE e.student_id = s.id
);
```

Meaning:

> Find students who have **no enrollments**.

---

# 4.12 ANY

Compares a value with the result of a subquery and returns true if the comparison is true for **at least one** value.

Example:

```sql
SELECT *
FROM employees
WHERE salary > ANY (
    SELECT salary
    FROM employees
    WHERE department = 'HR'
);
```

Meaning:

> Find employees whose salary is greater than **at least one HR employee's salary**.

Conceptually:

```text
salary > ANY (1000, 2000, 3000)

True if:
salary > 1000
OR
salary > 2000
OR
salary > 3000
```

---

# 4.13 ALL

Compares a value with every value returned by a subquery.

```sql
SELECT *
FROM employees
WHERE salary > ALL (
    SELECT salary
    FROM employees
    WHERE department = 'HR'
);
```

Meaning:

> Find employees whose salary is greater than **every HR employee's salary**.

Conceptually:

```text
salary > ALL (1000, 2000, 3000)

Must satisfy:

salary > 1000
AND
salary > 2000
AND
salary > 3000
```

---

#  Filtering Summary

| Operator | Purpose | Example |
|---|---|---|
| `=` | Equal | `age = 20` |
| `<>` | Not equal | `age <> 20` |
| `>` | Greater than | `age > 20` |
| `<` | Less than | `age < 20` |
| `>=` | Greater/equal | `age >= 20` |
| `<=` | Less/equal | `age <= 20` |
| `AND` | All conditions true | `age > 18 AND city = 'Delhi'` |
| `OR` | At least one true | `city = 'Delhi' OR city = 'Pune'` |
| `NOT` | Negate condition | `NOT city = 'Delhi'` |
| `IN` | Match list | `city IN ('Delhi', 'Pune')` |
| `NOT IN` | Exclude list | `city NOT IN ('Delhi', 'Pune')` |
| `BETWEEN` | Inclusive range | `age BETWEEN 18 AND 25` |
| `LIKE` | Pattern matching | `name LIKE 'A%'` |
| `ILIKE` | Case-insensitive pattern matching | `name ILIKE 'a%'` |
| `EXISTS` | Subquery returns rows | `EXISTS (...)` |
| `NOT EXISTS` | Subquery returns no rows | `NOT EXISTS (...)` |
| `ANY` | Match at least one subquery value | `salary > ANY (...)` |
| `ALL` | Match every subquery value | `salary > ALL (...)` |

---

#  Putting Filtering Together

You can combine multiple operators.

```sql
SELECT
    name,
    age,
    city
FROM students
WHERE age BETWEEN 18 AND 25
  AND city IN ('Delhi', 'Mumbai', 'Pune')
  AND name ILIKE 'a%'
ORDER BY age DESC;
```

This means:

```text
Age between 18 and 25
        AND
City is Delhi, Mumbai, or Pune
        AND
Name starts with A/a
        ↓
Sort by age descending
```

---

