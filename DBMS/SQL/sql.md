#  SQL Fundamentals
---

##  Table of Contents

- [1. CREATE DATABASE](#1-create-database)
- [2. CREATE TABLE](#2-create-table)
- [3. ALTER TABLE](#3-alter-table)
- [4. DROP TABLE](#4-drop-table)
- [5. TRUNCATE TABLE](#5-truncate-table)
- [6. INSERT](#6-insert)
- [7. SELECT](#7-select)
- [8. UPDATE](#8-update)
- [9. DELETE](#9-delete)
- [10. WHERE](#10-where)
- [11. ORDER BY](#11-order-by)
- [12. LIMIT](#12-limit)
- [13. OFFSET](#13-offset)
- [14. DISTINCT](#14-distinct)
- [15. Aliases](#15-aliases)
- [16. NULL](#16-null)
- [17. IS NULL](#17-is-null)
- [18. IS NOT NULL](#18-is-not-null)
- [19. COALESCE](#19-coalesce)
- [20. CASE](#20-case)
- [21. SQL Comments](#21-sql-comments)
- [22. Complete Example](#22-complete-example)
- [23. SQL Query Execution Order](#23-sql-query-execution-order)
- [24. What to Learn Next](#24-what-to-learn-next)

---

# 1. CREATE DATABASE

Creates a new database.

```sql
CREATE DATABASE college;
```

### Create only if it doesn't already exist

```sql
CREATE DATABASE IF NOT EXISTS college;
```

### Select the database

```sql
USE college;
```

---

# 2. CREATE TABLE

Creates a new table.

```sql
CREATE TABLE students (
    id INT,
    name VARCHAR(100),
    age INT,
    city VARCHAR(50)
);
```

### Using constraints

```sql
CREATE TABLE students (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    city VARCHAR(50)
);
```

### Common Data Types

| Data Type | Example |
|---|---|
| `INT` | `21` |
| `DECIMAL(10,2)` | `99.99` |
| `VARCHAR(100)` | `'Ankush'` |
| `DATE` | `'2026-10-17'` |
| `BOOLEAN` | `TRUE` |

---

# 3. ALTER TABLE

Used to modify an existing table.

## Add a Column

```sql
ALTER TABLE students
ADD email VARCHAR(100);
```

## Modify a Column

```sql
ALTER TABLE students
MODIFY age SMALLINT;
```

## Rename a Column

```sql
ALTER TABLE students
RENAME COLUMN name TO full_name;
```

## Drop a Column

```sql
ALTER TABLE students
DROP COLUMN email;
```

---

# 4. DROP TABLE

Deletes the **entire table**, including its structure and data.

```sql
DROP TABLE students;
```

After executing this, the table no longer exists.

> ⚠️ **Warning:** `DROP TABLE` is destructive.

---

# 5. TRUNCATE TABLE

Removes **all rows** from a table while keeping the table structure.

```sql
TRUNCATE TABLE students;
```

### Difference

| Command | Deletes Data | Deletes Structure |
|---|---:|---:|
| `DROP TABLE` | ✅ | ✅ |
| `TRUNCATE TABLE` | ✅ | ❌ |

---

# 6. INSERT

Adds data to a table.

## Insert One Row

```sql
INSERT INTO students (id, name, age, city)
VALUES (1, 'Ankush', 21, 'Delhi');
```

## Insert Multiple Rows

```sql
INSERT INTO students (id, name, age, city)
VALUES
    (2, 'Rahul', 22, 'Mumbai'),
    (3, 'Priya', 20, 'Delhi'),
    (4, 'Aman', 23, 'Pune');
```

---

# 7. SELECT

Used to retrieve data.

## Select All Columns

```sql
SELECT *
FROM students;
```

## Select Specific Columns

```sql
SELECT name, age
FROM students;
```

## Use Expressions

```sql
SELECT name, age + 1
FROM students;
```

---

# 8. UPDATE

Used to modify existing records.

```sql
UPDATE students
SET age = 22
WHERE id = 1;
```

### Update Multiple Columns

```sql
UPDATE students
SET age = 23,
    city = 'Chandigarh'
WHERE id = 1;
```

> ⚠️ **Always be careful when using `UPDATE` without `WHERE`.**

```sql
UPDATE students
SET age = 25;
```

This updates **every row**.

---

# 9. DELETE

Used to remove rows.

```sql
DELETE FROM students
WHERE id = 3;
```

### Delete Multiple Rows

```sql
DELETE FROM students
WHERE age < 21;
```

### Delete All Rows

```sql
DELETE FROM students;
```

This removes all rows but **does not remove the table**.

> ⚠️ Always check your `WHERE` condition before executing `DELETE`.

---

# 10. WHERE

Filters rows based on a condition.

```sql
SELECT *
FROM students
WHERE age > 20;
```

## Comparison Operators

| Operator | Meaning |
|---|---|
| `=` | Equal |
| `<>` | Not equal |
| `!=` | Not equal |
| `>` | Greater than |
| `<` | Less than |
| `>=` | Greater than or equal |
| `<=` | Less than or equal |

### Example

```sql
SELECT *
FROM students
WHERE city = 'Delhi';
```

```sql
SELECT *
FROM students
WHERE age >= 21;
```

## Multiple Conditions

### AND

```sql
SELECT *
FROM students
WHERE age >= 21
  AND city = 'Delhi';
```

### OR

```sql
SELECT *
FROM students
WHERE city = 'Delhi'
   OR city = 'Mumbai';
```

---

# 11. ORDER BY

Used to sort query results.

## Ascending Order

```sql
SELECT *
FROM students
ORDER BY age ASC;
```

`ASC` is the default.

## Descending Order

```sql
SELECT *
FROM students
ORDER BY age DESC;
```

## Multiple Columns

```sql
SELECT *
FROM students
ORDER BY city ASC, age DESC;
```

This means:

1. Sort by `city` ascending.
2. For the same city, sort by `age` descending.

---

# 12. LIMIT

Restricts the number of rows returned.

```sql
SELECT *
FROM students
LIMIT 3;
```

### Example: Top 5 Oldest Students

```sql
SELECT *
FROM students
ORDER BY age DESC
LIMIT 5;
```

---

# 13. OFFSET

Skips a specified number of rows.

```sql
SELECT *
FROM students
LIMIT 5 OFFSET 10;
```

This means:

```text
Skip 10 rows
↓
Return the next 5 rows
```

### Pagination Example

```text
Page 1 → LIMIT 10 OFFSET 0
Page 2 → LIMIT 10 OFFSET 10
Page 3 → LIMIT 10 OFFSET 20
Page 4 → LIMIT 10 OFFSET 30
```

---

# 14. DISTINCT

Removes duplicate values from the result.

Suppose the table contains:

```text
city
---------
Delhi
Mumbai
Delhi
Pune
Mumbai
```

Query:

```sql
SELECT DISTINCT city
FROM students;
```

Result:

```text
Delhi
Mumbai
Pune
```

### Multiple Columns

```sql
SELECT DISTINCT city, age
FROM students;
```

> `DISTINCT` applies to the **combination of selected columns**.

---

# 15. Aliases

Aliases provide temporary names for columns or tables.

## Column Alias

```sql
SELECT name AS student_name
FROM students;
```

You can also omit `AS`:

```sql
SELECT name student_name
FROM students;
```

## Expression Alias

```sql
SELECT age + 1 AS next_year_age
FROM students;
```

## Table Alias

```sql
SELECT s.name, s.age
FROM students AS s;
```

Table aliases become especially useful when working with `JOIN`.

---

# 16. NULL

`NULL` represents a **missing or unknown value**.

It is **not** the same as:

```text
0
''
FALSE
```

Example:

```text
id | name   | email
---+--------+----------------
1  | Ankush | ankush@gmail.com
2  | Rahul  | NULL
```

Rahul's email is `NULL`.

---

# 17. IS NULL

Used to find rows containing `NULL`.

```sql
SELECT *
FROM students
WHERE email IS NULL;
```

### ❌ Incorrect

```sql
WHERE email = NULL;
```

Use:

```sql
WHERE email IS NULL;
```

---

# 18. IS NOT NULL

Used to find rows where a value exists.

```sql
SELECT *
FROM students
WHERE email IS NOT NULL;
```

Example:

```sql
SELECT name
FROM students
WHERE phone IS NOT NULL;
```

---

# 19. COALESCE

Returns the **first non-NULL value**.

```sql
COALESCE(value1, value2, value3)
```

### Example

```sql
SELECT
    name,
    COALESCE(email, 'No Email') AS email
FROM students;
```

If `email` is `NULL`, the result becomes:

```text
No Email
```

### Multiple Fallbacks

```sql
SELECT
    name,
    COALESCE(phone, email, 'No Contact') AS contact
FROM students;
```

SQL checks:

```text
phone
  ↓
Is it NULL?
  ↓
email
  ↓
Is it NULL?
  ↓
'No Contact'
```

---

# 20. CASE

`CASE` provides conditional logic in SQL.

It is similar to:

```text
if
else if
else
```

## Basic Example

```sql
SELECT
    name,
    age,
    CASE
        WHEN age >= 18 THEN 'Adult'
        ELSE 'Minor'
    END AS category
FROM students;
```

### Multiple Conditions

```sql
SELECT
    name,
    marks,
    CASE
        WHEN marks >= 90 THEN 'A'
        WHEN marks >= 80 THEN 'B'
        WHEN marks >= 70 THEN 'C'
        WHEN marks >= 60 THEN 'D'
        ELSE 'F'
    END AS grade
FROM students;
```

### CASE with UPDATE

```sql
UPDATE students
SET category =
    CASE
        WHEN age >= 18 THEN 'Adult'
        ELSE 'Minor'
    END;
```

---

# 21. SQL Comments

Comments are ignored by SQL.

## Single-Line Comment

```sql
-- This is a comment

SELECT *
FROM students;
```

## MySQL `#` Comment

```sql
# This is a comment

SELECT *
FROM students;
```

## Multi-Line Comment

```sql
/*
   This is a
   multi-line comment
*/

SELECT *
FROM students;
```

---

# 22. Complete Example

Let's combine several concepts into one example.

## Create Database

```sql
CREATE DATABASE college;

USE college;
```

## Create Table

```sql
CREATE TABLE students (
    id INT PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    city VARCHAR(50),
    email VARCHAR(100)
);
```

## Insert Data

```sql
INSERT INTO students
    (id, name, age, city, email)
VALUES
    (1, 'Ankush', 21, 'Delhi', 'ankush@gmail.com'),
    (2, 'Rahul', 19, 'Mumbai', NULL),
    (3, 'Priya', 22, 'Delhi', 'priya@gmail.com'),
    (4, 'Aman', 20, 'Pune', NULL),
    (5, 'Neha', 23, 'Mumbai', 'neha@gmail.com');
```

## Query

```sql
SELECT
    name AS student_name,
    age,
    COALESCE(email, 'No Email') AS email_status,
    CASE
        WHEN age >= 21 THEN 'Senior'
        ELSE 'Junior'
    END AS level
FROM students
WHERE city = 'Delhi'
ORDER BY age DESC
LIMIT 10;
```

This query demonstrates:

- `SELECT`
- Aliases
- `COALESCE`
- `CASE`
- `FROM`
- `WHERE`
- `ORDER BY`
- `LIMIT`

---

# 23. SQL Query Execution Order

When writing SQL, remember that the order you **write** a query isn't exactly the order in which SQL logically processes it.

Example:

```sql
SELECT name, age
FROM students
WHERE age >= 20
ORDER BY age DESC
LIMIT 5;
```

### Logical Processing Order

```text
FROM
  ↓
WHERE
  ↓
GROUP BY
  ↓
HAVING
  ↓
SELECT
  ↓
DISTINCT
  ↓
ORDER BY
  ↓
LIMIT / OFFSET
```

This becomes extremely important when learning:

- `GROUP BY`
- `HAVING`
- Subqueries
- CTEs
- Window Functions

---



## ⭐ Quick Reference

| Command | Purpose |
|---|---|
| `CREATE DATABASE` | Create database |
| `CREATE TABLE` | Create table |
| `ALTER TABLE` | Modify table structure |
| `DROP TABLE` | Delete table |
| `TRUNCATE TABLE` | Remove all rows |
| `INSERT` | Add data |
| `SELECT` | Retrieve data |
| `UPDATE` | Modify data |
| `DELETE` | Remove data |
| `WHERE` | Filter rows |
| `ORDER BY` | Sort results |
| `LIMIT` | Limit results |
| `OFFSET` | Skip rows |
| `DISTINCT` | Remove duplicates |
| `AS` | Create aliases |
| `IS NULL` | Find missing values |
| `IS NOT NULL` | Find non-missing values |
| `COALESCE()` | Replace/fallback for `NULL` |
| `CASE` | Conditional logic |
| `--` | Single-line comment |
| `/* */` | Multi-line comment |

---

## 📌 Key Takeaways

1. **`DROP` deletes the table; `TRUNCATE` keeps the table.**
2. **Always use `WHERE` carefully with `UPDATE` and `DELETE`.**
3. **Use `IS NULL`, not `= NULL`.**
4. **`COALESCE()` is useful for handling missing values.**
5. **`CASE` provides conditional logic.**
6. **`ORDER BY ... LIMIT` is a common pattern for top-N queries.**
7. **`LIMIT ... OFFSET` is commonly used for pagination.**
8. **Aliases improve query readability.**
9. Learn the **logical query processing order** early—it will save you confusion later.
10. After these fundamentals, focus heavily on **aggregations and JOINs**.

---

