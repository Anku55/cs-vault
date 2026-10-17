##  Table of Contents

- [3. Constraints](#3-constraints)
  - [3.1 PRIMARY KEY](#31-primary-key)
  - [3.2 FOREIGN KEY](#32-foreign-key)
  - [3.3 UNIQUE](#33-unique)
  - [3.4 NOT NULL](#34-not-null)
  - [3.5 CHECK](#35-check)
  - [3.6 DEFAULT](#36-default)
  - [3.7 Composite Primary Key](#37-composite-primary-key)
  - [3.8 Composite UNIQUE Constraint](#38-composite-unique-constraint)
  - [3.9 ON DELETE CASCADE](#39-on-delete-cascade)
  - [3.10 ON DELETE SET NULL](#310-on-delete-set-null)
  - [3.11 ON DELETE RESTRICT](#311-on-delete-restrict)
  - [3.12 ON UPDATE CASCADE](#312-on-update-cascade)
  - [3.13 Constraint Violations](#313-constraint-violations)

- [4. SQL Filtering](#4-sql-filtering)
  - [4.1 Comparison Operators](#41-comparison-operators)
  - [4.2 AND](#42-and)
  - [4.3 OR](#43-or)
  - [4.4 NOT](#44-not)
  - [4.5 IN](#45-in)
  - [4.6 NOT IN](#46-not-in)
  - [4.7 BETWEEN](#47-between)
  - [4.8 LIKE](#48-like)
  - [4.9 ILIKE](#49-ilike)
  - [4.10 EXISTS](#410-exists)
  - [4.11 NOT EXISTS](#411-not-exists)
  - [4.12 ANY](#412-any)
  - [4.13 ALL](#413-all)


# 3. Constraints

Constraints are rules applied to table columns to maintain **data integrity, consistency, and valid relationships**.

---

## 3.1 PRIMARY KEY

Uniquely identifies each row in a table.

- Must be unique
- Cannot contain `NULL`
- A table can have only one primary key constraint

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100)
);
```

### Example

```sql
INSERT INTO students (id, name)
VALUES (1, 'Ankush');

INSERT INTO students (id, name)
VALUES (1, 'Rahul');
```

The second statement fails because `id = 1` already exists.

---

## 3.2 FOREIGN KEY

Creates a relationship between two tables.

```sql
CREATE TABLE departments (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100),
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
    ▼
students.department_id
```

A foreign key ensures that the referenced value exists in the parent table, unless `NULL` is allowed.

---

## 3.3 UNIQUE

Ensures that values in a column are unique.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE
);
```

This is valid:

```sql
INSERT INTO users (id, email)
VALUES (1, 'ankush@gmail.com');
```

This fails:

```sql
INSERT INTO users (id, email)
VALUES (2, 'ankush@gmail.com');
```

> **Note:** In PostgreSQL, a normal `UNIQUE` constraint allows multiple `NULL` values.

---

## 3.4 NOT NULL

Prevents a column from containing `NULL`.

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

This fails:

```sql
INSERT INTO students (id, name)
VALUES (1, NULL);
```

Use `NOT NULL` when a value is required.

---

## 3.5 CHECK

Ensures that a condition is satisfied.

```sql
CREATE TABLE students (
    id INTEGER PRIMARY KEY,
    age INTEGER CHECK (age >= 18)
);
```

Valid:

```sql
INSERT INTO students (id, age)
VALUES (1, 21);
```

Invalid:

```sql
INSERT INTO students (id, age)
VALUES (2, 15);
```

### Multiple Conditions

```sql
CREATE TABLE products (
    id INTEGER PRIMARY KEY,
    price NUMERIC CHECK (price >= 0),
    quantity INTEGER CHECK (quantity >= 0)
);
```

---

## 3.6 DEFAULT

Provides a default value when no value is supplied.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE
);
```

Now:

```sql
INSERT INTO users (id, name)
VALUES (1, 'Ankush');
```

Automatically results in:

```text
is_active = TRUE
```

---

## 3.7 Composite Primary Key

A primary key consisting of **multiple columns**.

```sql
CREATE TABLE enrollments (
    student_id INTEGER,
    course_id INTEGER,

    PRIMARY KEY (student_id, course_id)
);
```

Valid:

```text
student_id | course_id
-----------+----------
1          | 101
1          | 102
2          | 101
```

Invalid:

```text
student_id | course_id
-----------+----------
1          | 101
1          | 101  ← duplicate
```

### Common Use

Composite primary keys are frequently used for **many-to-many relationship tables**.

```text
Students
   │
   │
   ▼
Enrollments
   ▲
   │
   │
Courses
```

---

## 3.8 Composite UNIQUE Constraint

A `UNIQUE` constraint involving multiple columns.

```sql
CREATE TABLE schedules (
    room_id INTEGER,
    schedule_date DATE,

    UNIQUE (room_id, schedule_date)
);
```

This prevents the same room from being scheduled more than once on the same date.

Valid:

```text
room_id | schedule_date
--------+--------------
101     | 2026-10-17
101     | 2026-10-18
102     | 2026-10-17
```

Invalid:

```text
room_id | schedule_date
--------+--------------
101     | 2026-10-17
101     | 2026-10-17  ← duplicate combination
```

---

## 3.9 ON DELETE CASCADE

Automatically deletes child rows when the referenced parent row is deleted.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY
);

CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
```

If:

```sql
DELETE FROM users
WHERE id = 1;
```

then all orders belonging to user `1` are automatically deleted.

```text
User 1
  │
  ├── Order 101
  ├── Order 102
  └── Order 103
        │
        ▼
      DELETE
        │
        ▼
All related orders deleted
```

> ⚠️ Use `CASCADE` carefully because one delete can remove many related rows.

---

## 3.10 ON DELETE SET NULL

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

If manager `10` is deleted:

```text
manager_id = 10
       ↓
manager_id = NULL
```

The employee record remains.

> The foreign key column must allow `NULL`.

---

## 3.11 ON DELETE RESTRICT

Prevents deletion of a parent row if dependent child rows exist.

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT
);
```

If user `1` has orders:

```sql
DELETE FROM users
WHERE id = 1;
```

The database rejects the deletion.

```text
Parent
  │
  ├── Child
  ├── Child
  └── Child

DELETE parent
      ↓
   REJECTED
```

Useful when child records must not become orphaned.

---

## 3.12 ON UPDATE CASCADE

Automatically updates foreign key values when the referenced key changes.

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

>  Primary keys are normally stable, so `ON UPDATE CASCADE` is less commonly needed than `ON DELETE` actions.

---

## 3.13 Constraint Violations

A constraint violation occurs when an operation violates a defined constraint.

### PRIMARY KEY violation

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY
);

INSERT INTO users (id)
VALUES (1);

INSERT INTO users (id)
VALUES (1);
```

 The second insert violates the primary key.

---

### UNIQUE violation

```sql
CREATE TABLE users (
    email VARCHAR(255) UNIQUE
);

INSERT INTO users (email)
VALUES ('a@gmail.com');

INSERT INTO users (email)
VALUES ('a@gmail.com');
```

 Duplicate value violates `UNIQUE`.

---

### NOT NULL violation

```sql
CREATE TABLE users (
    name TEXT NOT NULL
);

INSERT INTO users (name)
VALUES (NULL);
```

 `NULL` violates `NOT NULL`.

---

### CHECK violation

```sql
CREATE TABLE products (
    price NUMERIC CHECK (price >= 0)
);

INSERT INTO products (price)
VALUES (-100);
```

❌ Value violates the `CHECK` condition.

---

### FOREIGN KEY violation

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES users(id)
);

INSERT INTO orders (id, user_id)
VALUES (1, 999);
```

If user `999` does not exist:

❌ Foreign key constraint violation.

---

# 4. SQL Filtering

Filtering is used to select only rows that satisfy specific conditions.

Most filtering is performed using:

```sql
WHERE
```

Example:

```sql
SELECT *
FROM students
WHERE age >= 18;
```

---

## 4.1 Comparison Operators

| Operator | Meaning | Example |
|---|---|---|
| `=` | Equal | `age = 20` |
| `<>` | Not equal | `age <> 20` |
| `!=` | Not equal | `age != 20` |
| `>` | Greater than | `age > 20` |
| `<` | Less than | `age < 20` |
| `>=` | Greater/equal | `age >= 20` |
| `<=` | Less/equal | `age <= 20` |

Example:

```sql
SELECT *
FROM students
WHERE age >= 18;
```

---

## 4.2 AND

All conditions must be `TRUE`.

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
city = 'Delhi'
```

Both conditions must match.

---

## 4.3 OR

At least one condition must be `TRUE`.

```sql
SELECT *
FROM students
WHERE city = 'Delhi'
   OR city = 'Mumbai';
```

This returns students from either city.

---

## 4.4 NOT

Negates a condition.

```sql
SELECT *
FROM students
WHERE NOT city = 'Delhi';
```

Equivalent:

```sql
SELECT *
FROM students
WHERE city <> 'Delhi';
```

`NOT` can also be used with:

```sql
NOT IN
NOT BETWEEN
NOT LIKE
NOT EXISTS
```

---

## 4.5 IN

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

## 4.6 NOT IN

Checks whether a value does **not** exist in a list.

```sql
SELECT *
FROM students
WHERE city NOT IN ('Delhi', 'Mumbai');
```

This excludes students from Delhi and Mumbai.

> ⚠️ `NOT IN` requires care when `NULL` can occur, because SQL's three-valued logic can cause unexpected results.

---

## 4.7 BETWEEN

Checks whether a value is within an **inclusive range**.

```sql
SELECT *
FROM students
WHERE age BETWEEN 18 AND 25;
```

Includes both:

```text
18
...
25
```

Equivalent to:

```sql
SELECT *
FROM students
WHERE age >= 18
  AND age <= 25;
```

---

## 4.8 LIKE

Used for pattern matching.

### `%`

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

### `_`

Matches exactly one character.

```sql
SELECT *
FROM students
WHERE name LIKE 'A___';
```

Matches four-character names beginning with `A`.

### Common Patterns

| Pattern | Meaning |
|---|---|
| `'A%'` | Starts with `A` |
| `'%A'` | Ends with `A` |
| `'%A%'` | Contains `A` |
| `'A___'` | Starts with `A`, exactly 4 characters |
| `'_____'` | Exactly 5 characters |

---

## 4.9 ILIKE

`ILIKE` performs **case-insensitive pattern matching** in PostgreSQL.

```sql
SELECT *
FROM students
WHERE name ILIKE 'ankush';
```

Matches:

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

> **PostgreSQL:** `ILIKE`  
> **Other databases:** Case-insensitive matching may use different syntax or collations.

---

## 4.10 EXISTS

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

> Find students who have at least one enrollment.

Conceptually:

```text
Student
   │
   ▼
Does enrollment exist?
   │
 ┌─┴─┐
YES  NO
 │    │
 ▼    X
Return
```

`EXISTS` is commonly used with correlated subqueries.

---

## 4.11 NOT EXISTS

Checks whether a subquery returns **zero rows**.

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

> Find students who have no enrollments.

---

## 4.12 ANY

Compares a value with values returned by a subquery.

The condition is `TRUE` if it matches **at least one** value.

```sql
SELECT *
FROM employees
WHERE salary > ANY (
    SELECT salary
    FROM employees
    WHERE department = 'HR'
);
```

Suppose HR salaries are:

```text
30000
40000
50000
```

Then:

```text
salary > ANY (30000, 40000, 50000)
```

means the salary must be greater than **at least one** of those values.

Conceptually:

```text
salary > 30000
     OR
salary > 40000
     OR
salary > 50000
```

---

## 4.13 ALL

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

If HR salaries are:

```text
30000
40000
50000
```

then the employee must have:

```text
salary > 30000
AND
salary > 40000
AND
salary > 50000
```

In other words, the salary must be greater than the **maximum HR salary** in this example.

---

# 🔥 Combining Filters

Multiple filtering techniques can be combined.

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

This query means:

```text
Age is between 18 and 25
        AND
City is Delhi, Mumbai, or Pune
        AND
Name starts with A/a
        ↓
Sort by age descending
```

---

