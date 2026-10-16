# SQL Revision Notes

## Databases

**Definition:**  
A database is an organized collection of data that can be stored, managed, and retrieved efficiently.

**Example:**
```sql
CREATE DATABASE college;

USE college;
```

---

## Tables

**Definition:**  
A table is used to store data in rows and columns.

**Example:**
```sql
CREATE TABLE students (
    id INT,
    name VARCHAR(50),
    age INT
);
```

---

## Insert Rows

**Definition:**  
`INSERT INTO` is used to add new rows to a table.

**Example:**
```sql
INSERT INTO students (id, name, age)
VALUES
    (1, 'Ankush', 20),
    (2, 'Rahul', 21),
    (3, 'Aman', 19);
```

---

## SELECT

**Definition:**  
`SELECT` is used to retrieve data from a table.

**Example:**
```sql
SELECT name, age
FROM students;
```

To retrieve all columns:

```sql
SELECT * FROM students;
```

---

## UPDATE

**Definition:**  
`UPDATE` is used to modify existing data in a table.

**Example:**
```sql
UPDATE students
SET age = 21
WHERE id = 1;
```

---

## DELETE

**Definition:**  
`DELETE` is used to remove rows from a table.

**Example:**
```sql
DELETE FROM students
WHERE id = 3;
```

---

## AUTOCOMMIT

**Definition:**  
Autocommit automatically saves each SQL statement as a transaction.

**Example:**
```sql
SET autocommit = 1;
```

---

## COMMIT

**Definition:**  
`COMMIT` permanently saves the changes made in the current transaction.

**Example:**
```sql
START TRANSACTION;

UPDATE students
SET age = 22
WHERE id = 1;

COMMIT;
```

---

## ROLLBACK

**Definition:**  
`ROLLBACK` undoes changes made in the current transaction that have not been committed.

**Example:**
```sql
START TRANSACTION;

UPDATE students
SET age = 25
WHERE id = 1;

ROLLBACK;
```

The update is undone.

---

## CURRENT_DATE()

**Definition:**  
`CURRENT_DATE()` returns the current date.

**Example:**
```sql
SELECT CURRENT_DATE();
```

**Output:**
```text
2026-10-16
```

---

## CURRENT_TIME()

**Definition:**  
`CURRENT_TIME()` returns the current time.

**Example:**
```sql
SELECT CURRENT_TIME();
```

**Output:**
```text
23:48:10
```

---

## UNIQUE

**Definition:**  
`UNIQUE` ensures that values in a column cannot be duplicated.

**Example:**
```sql
CREATE TABLE users (
    id INT,
    email VARCHAR(100) UNIQUE
);
```

This is not allowed:

```text
1 | ankush@gmail.com
2 | ankush@gmail.com
```

because the email must be unique.

---

## NOT NULL

**Definition:**  
`NOT NULL` ensures that a column cannot contain `NULL` values.

**Example:**
```sql
CREATE TABLE students (
    id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    age INT
);
```

This is invalid:

```sql
INSERT INTO students (id, name, age)
VALUES (1, NULL, 20);
```

because `name` cannot be `NULL`.

---

## CHECK

**Definition:**  
`CHECK` ensures that values satisfy a specified condition.

**Example:**
```sql
CREATE TABLE students (
    id INT,
    name VARCHAR(50),
    age INT CHECK (age >= 18)
);
```

Valid:

```sql
INSERT INTO students
VALUES (1, 'Ankush', 20);
```

Invalid:

```sql
INSERT INTO students
VALUES (2, 'Rahul', 15);
```

because `age` must be at least `18`.

---

## DEFAULT

**Definition:**  
`DEFAULT` automatically assigns a specified value when no value is provided.

**Example:**
```sql
CREATE TABLE users (
    id INT,
    name VARCHAR(50),
    status VARCHAR(20) DEFAULT 'active'
);
```

```sql
INSERT INTO users (id, name)
VALUES (1, 'Ankush');
```

The database automatically sets:

```text
status = 'active'
```
## PRIMARY KEY

**Definition:**  
A `PRIMARY KEY` is a column (or combination of columns) that **uniquely identifies each row** in a table. It cannot contain `NULL` values or duplicate values.

**Example:**
```sql
CREATE TABLE students (
    id INT PRIMARY KEY,
    name VARCHAR(50),
    age INT
);
```

```sql
INSERT INTO students (id, name, age)
VALUES (1, 'Ankush', 20);
```

Each student must have a unique `id`.

---

## AUTO_INCREMENT

**Definition:**  
`AUTO_INCREMENT` automatically generates a **new numeric value** for a column whenever a new row is inserted.

**Example:**
```sql
CREATE TABLE students (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    age INT
);
```

Now you don't need to provide the `id`:

```sql
INSERT INTO students (name, age)
VALUES ('Ankush', 20);

INSERT INTO students (name, age)
VALUES ('Rahul', 21);
```

The database automatically generates:

```text
id | name
---|-------
1  | Ankush
2  | Rahul
```

**Common usage:**
```sql
id INT PRIMARY KEY AUTO_INCREMENT
```

This makes `id` **unique** and automatically generates its value.