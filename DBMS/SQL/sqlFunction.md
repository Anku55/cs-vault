# 5. SQL Functions

SQL functions are built-in operations that allow you to **calculate values, transform data, manipulate strings, work with numbers, and handle dates and timestamps**.

Functions can be used in:

- `SELECT`
- `WHERE`
- `HAVING`
- `ORDER BY`
- `GROUP BY`
- `CASE`
- `JOIN` conditions

---

## 5.1 Aggregate Functions

Aggregate functions operate on **multiple rows** and return a **single result**.

### COUNT

Counts rows or non-NULL values.

```sql
-- Count all rows
SELECT COUNT(*)
FROM employees;
```

```sql
-- Count non-NULL employee names
SELECT COUNT(name)
FROM employees;
```

```sql
-- Count distinct departments
SELECT COUNT(DISTINCT department_id)
FROM employees;
```

> `COUNT(*)` counts rows, while `COUNT(column)` ignores `NULL` values.

---

### SUM

Returns the total of a numeric column.

```sql
SELECT SUM(salary)
FROM employees;
```

```sql
SELECT department_id, SUM(salary)
FROM employees
GROUP BY department_id;
```

---

### AVG

Returns the average of numeric values.

```sql
SELECT AVG(salary)
FROM employees;
```

```sql
SELECT department_id, AVG(salary)
FROM employees
GROUP BY department_id;
```

> `AVG()` ignores `NULL` values.

---

### MIN

Returns the smallest value.

```sql
SELECT MIN(salary)
FROM employees;
```

It can also be used with dates and strings:

```sql
SELECT MIN(hire_date)
FROM employees;
```

---

### MAX

Returns the largest value.

```sql
SELECT MAX(salary)
FROM employees;
```

```sql
SELECT MAX(hire_date)
FROM employees;
```

---

### Aggregate Functions Summary

| Function | Purpose |
|---|---|
| `COUNT()` | Counts rows or values |
| `SUM()` | Calculates total |
| `AVG()` | Calculates average |
| `MIN()` | Returns minimum value |
| `MAX()` | Returns maximum value |

### Aggregate Functions with `GROUP BY`

```sql
SELECT department_id,
       COUNT(*) AS employee_count,
       SUM(salary) AS total_salary,
       AVG(salary) AS average_salary,
       MIN(salary) AS minimum_salary,
       MAX(salary) AS maximum_salary
FROM employees
GROUP BY department_id;
```

This is one of the most important patterns in SQL:

```text
Rows
  ↓
GROUP BY
  ↓
Aggregate Functions
  ↓
One result per group
```

---

# 5.2 String Functions

String functions are used to **modify, format, search, and extract text**.

---

## LOWER

Converts text to lowercase.

```sql
SELECT LOWER(name)
FROM employees;
```

Example:

```text
'Ankush Sharma'
      ↓
'ankush sharma'
```

Useful for case-insensitive comparisons:

```sql
SELECT *
FROM users
WHERE LOWER(email) = 'user@example.com';
```

---

## UPPER

Converts text to uppercase.

```sql
SELECT UPPER(name)
FROM employees;
```

Example:

```text
'Ankush Sharma'
      ↓
'ANKUSH SHARMA'
```

---

## LENGTH

Returns the number of characters in a string.

```sql
SELECT LENGTH(name)
FROM employees;
```

Example:

```text
'Ankush'
   ↓
6
```

Can be used for filtering:

```sql
SELECT *
FROM users
WHERE LENGTH(username) > 5;
```

---

## CONCAT

Combines multiple strings.

```sql
SELECT CONCAT(first_name, ' ', last_name) AS full_name
FROM employees;
```

Example:

```text
first_name = 'Ankush'
last_name  = 'Sharma'

CONCAT(...)
    ↓

'Ankush Sharma'
```

You can also use PostgreSQL's string concatenation operator:

```sql
SELECT first_name || ' ' || last_name AS full_name
FROM employees;
```

---

## SUBSTRING

Extracts part of a string.

```sql
SELECT SUBSTRING('PostgreSQL' FROM 1 FOR 4);
```

Result:

```text
Post
```

Another common form:

```sql
SELECT SUBSTRING(name, 1, 3)
FROM employees;
```

---

## TRIM

Removes leading and trailing spaces.

```sql
SELECT TRIM('   Hello World   ');
```

Result:

```text
Hello World
```

Useful when cleaning user input:

```sql
SELECT TRIM(name)
FROM users;
```

You can also remove specific characters:

```sql
SELECT TRIM(BOTH '-' FROM '---hello---');
```

---

## REPLACE

Replaces occurrences of a substring.

```sql
SELECT REPLACE('Hello World', 'World', 'SQL');
```

Result:

```text
Hello SQL
```

Example:

```sql
SELECT REPLACE(phone_number, '-', '')
FROM users;
```

---

### String Functions Summary

| Function | Purpose |
|---|---|
| `LOWER()` | Converts text to lowercase |
| `UPPER()` | Converts text to uppercase |
| `LENGTH()` | Returns string length |
| `CONCAT()` | Combines strings |
| `SUBSTRING()` | Extracts part of a string |
| `TRIM()` | Removes leading/trailing characters |
| `REPLACE()` | Replaces text |

---

# 5.3 Numeric Functions

Numeric functions perform **mathematical and numerical operations**.

---

## ROUND

Rounds a number to a specified number of decimal places.

```sql
SELECT ROUND(123.4567, 2);
```

Result:

```text
123.46
```

Example:

```sql
SELECT ROUND(AVG(salary), 2)
FROM employees;
```

---

## CEIL

Rounds a number **up** to the nearest integer.

```sql
SELECT CEIL(12.3);
```

Result:

```text
13
```

```sql
SELECT CEIL(12.9);
```

Result:

```text
13
```

---

## FLOOR

Rounds a number **down** to the nearest integer.

```sql
SELECT FLOOR(12.9);
```

Result:

```text
12
```

---

## ABS

Returns the absolute value of a number.

```sql
SELECT ABS(-25);
```

Result:

```text
25
```

Useful when calculating differences:

```sql
SELECT ABS(actual_amount - expected_amount) AS difference
FROM transactions;
```

---

## MOD

Returns the remainder after division.

```sql
SELECT MOD(10, 3);
```

Result:

```text
1
```

Equivalent PostgreSQL operator:

```sql
SELECT 10 % 3;
```

Example — find even numbers:

```sql
SELECT *
FROM numbers
WHERE MOD(number, 2) = 0;
```

---

### Numeric Functions Summary

| Function | Purpose |
|---|---|
| `ROUND()` | Rounds a number |
| `CEIL()` | Rounds upward |
| `FLOOR()` | Rounds downward |
| `ABS()` | Returns absolute value |
| `MOD()` | Returns remainder |

---

# 5.4 Date Functions

Date functions are used to **retrieve, manipulate, extract, truncate, and compare dates and timestamps**.

---

## CURRENT_DATE

Returns the current date.

```sql
SELECT CURRENT_DATE;
```

Example result:

```text
2026-10-17
```

Find employees hired today:

```sql
SELECT *
FROM employees
WHERE hire_date = CURRENT_DATE;
```

---

## CURRENT_TIMESTAMP

Returns the current date and time, including the current timestamp.

```sql
SELECT CURRENT_TIMESTAMP;
```

---

## NOW()

Returns the current date and time.

```sql
SELECT NOW();
```

In PostgreSQL:

```sql
SELECT NOW();
```

and

```sql
SELECT CURRENT_TIMESTAMP;
```

are commonly used for the same purpose.

---

## DATE_PART

Extracts a specific component from a date or timestamp.

```sql
SELECT DATE_PART('year', CURRENT_DATE);
```

Example:

```text
2026
```

Extract the month:

```sql
SELECT DATE_PART('month', CURRENT_DATE);
```

Extract the day:

```sql
SELECT DATE_PART('day', CURRENT_DATE);
```

Example using a table:

```sql
SELECT name,
       DATE_PART('year', hire_date) AS hire_year
FROM employees;
```

Common fields include:

```text
year
month
day
hour
minute
second
quarter
week
```

---

## DATE_TRUNC

Truncates a date or timestamp to a specified precision.

```sql
SELECT DATE_TRUNC('month', CURRENT_TIMESTAMP);
```

Example:

```text
2026-10-17 14:35:22
        ↓
2026-10-01 00:00:00
```

Truncate to year:

```sql
SELECT DATE_TRUNC('year', CURRENT_TIMESTAMP);
```

Truncate to day:

```sql
SELECT DATE_TRUNC('day', CURRENT_TIMESTAMP);
```

Common use case — monthly aggregation:

```sql
SELECT DATE_TRUNC('month', order_date) AS month,
       SUM(amount) AS revenue
FROM orders
GROUP BY DATE_TRUNC('month', order_date)
ORDER BY month;
```

This is an extremely common analytical SQL pattern.

---

# 5.5 Interval Operations

An `INTERVAL` represents a **duration of time**.

You can add or subtract intervals from dates and timestamps.

### Add Days

```sql
SELECT CURRENT_DATE + INTERVAL '7 days';
```

### Subtract Days

```sql
SELECT CURRENT_DATE - INTERVAL '7 days';
```

### Add Months

```sql
SELECT CURRENT_DATE + INTERVAL '1 month';
```

### Add Years

```sql
SELECT CURRENT_DATE + INTERVAL '1 year';
```

### Add Hours

```sql
SELECT CURRENT_TIMESTAMP + INTERVAL '2 hours';
```

---

### Date Difference

You can calculate differences between dates.

```sql
SELECT CURRENT_DATE - hire_date AS days_employed
FROM employees;
```

For timestamps:

```sql
SELECT CURRENT_TIMESTAMP - created_at AS account_age
FROM users;
```

---

# 5.6 Date Comparisons

Dates can be compared using standard comparison operators.

```sql
=
<>
>
<
>=
<=
```

Example:

```sql
SELECT *
FROM employees
WHERE hire_date > '2025-01-01';
```

### Before a Date

```sql
SELECT *
FROM orders
WHERE order_date < '2026-01-01';
```

### After a Date

```sql
SELECT *
FROM orders
WHERE order_date >= '2026-01-01';
```

### Date Range

```sql
SELECT *
FROM orders
WHERE order_date BETWEEN '2026-01-01' AND '2026-01-31';
```

### Recent Records

```sql
SELECT *
FROM orders
WHERE created_at >= CURRENT_TIMESTAMP - INTERVAL '7 days';
```

---

# 5.7 Functions with `GROUP BY`

Functions become particularly powerful when combined with `GROUP BY`.

```sql
SELECT department_id,
       COUNT(*) AS employee_count,
       ROUND(AVG(salary), 2) AS average_salary,
       MIN(salary) AS minimum_salary,
       MAX(salary) AS maximum_salary
FROM employees
GROUP BY department_id;
```

---

# 5.8 Functions with `WHERE`

Functions can also be used to filter records.

```sql
SELECT *
FROM users
WHERE LENGTH(username) > 8;
```

```sql
SELECT *
FROM employees
WHERE LOWER(department) = 'engineering';
```

```sql
SELECT *
FROM orders
WHERE order_date >= CURRENT_DATE - INTERVAL '30 days';
```

---

# 5.9 Functions with `ORDER BY`

Functions can be used to determine sorting logic.

```sql
SELECT name, salary
FROM employees
ORDER BY ROUND(salary, -3) DESC;
```

Or:

```sql
SELECT name
FROM users
ORDER BY LENGTH(name) DESC;
```

---

# 5.10 Function Categories — Quick Reference

| Category | Functions |
|---|---|
| **Aggregate** | `COUNT`, `SUM`, `AVG`, `MIN`, `MAX` |
| **String** | `LOWER`, `UPPER`, `LENGTH`, `CONCAT`, `SUBSTRING`, `TRIM`, `REPLACE` |
| **Numeric** | `ROUND`, `CEIL`, `FLOOR`, `ABS`, `MOD` |
| **Date** | `CURRENT_DATE`, `CURRENT_TIMESTAMP`, `NOW()`, `DATE_PART`, `DATE_TRUNC` |
| **Date Operations** | `INTERVAL`, date arithmetic, date comparisons |

---

# 5.11 Function Execution Pattern

A useful mental model:

```text
Raw Data
   ↓
Function
   ↓
Transformed Value
   ↓
Filter / Group / Sort / Display
```

For example:

```sql
SELECT department_id,
       ROUND(AVG(salary), 2) AS avg_salary
FROM employees
WHERE salary > 50000
GROUP BY department_id
ORDER BY avg_salary DESC;
```

Conceptually:

```text
employees
    ↓
WHERE
    ↓
salary > 50000
    ↓
GROUP BY department_id
    ↓
AVG(salary)
    ↓
ROUND(..., 2)
    ↓
ORDER BY avg_salary DESC
```

---

# 5.12 SQL Functions Checklist

### Aggregate Functions

- [x] `COUNT()`
- [x] `SUM()`
- [x] `AVG()`
- [x] `MIN()`
- [x] `MAX()`

### String Functions

- [x] `LOWER()`
- [x] `UPPER()`
- [x] `LENGTH()`
- [x] `CONCAT()`
- [x] `SUBSTRING()`
- [x] `TRIM()`
- [x] `REPLACE()`

### Numeric Functions

- [x] `ROUND()`
- [x] `CEIL()`
- [x] `FLOOR()`
- [x] `ABS()`
- [x] `MOD()`

### Date Functions

- [x] `CURRENT_DATE`
- [x] `CURRENT_TIMESTAMP`
- [x] `NOW()`
- [x] `DATE_PART()`
- [x] `DATE_TRUNC()`
- [x] `INTERVAL` operations
- [x] Date comparisons

---

## 🧠 Key Takeaways

```text
Aggregate Functions
    ↓
COUNT | SUM | AVG | MIN | MAX

String Functions
    ↓
LOWER | UPPER | LENGTH | CONCAT
SUBSTRING | TRIM | REPLACE

Numeric Functions
    ↓
ROUND | CEIL | FLOOR | ABS | MOD

Date Functions
    ↓
CURRENT_DATE | CURRENT_TIMESTAMP | NOW
DATE_PART | DATE_TRUNC | INTERVAL
```

These functions form the foundation for the next major SQL topics:

```text
SQL Functions
     ↓
GROUP BY
     ↓
HAVING
     ↓
JOINs
     ↓
Subqueries
     ↓
CTEs
     ↓
Window Functions
```