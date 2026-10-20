# SQL Subqueries --- Deep Practical & Theoretical Notes

> A deep, practical, interview-ready guide to SQL subqueries, their
> execution model, NULL behavior, correlated queries, `IN`, `EXISTS`,
> `NOT EXISTS`, and subquery-vs-JOIN decisions.

------------------------------------------------------------------------

## Table of Contents

1.  [What Is a Subquery?](#1-what-is-a-subquery)
2.  [Why Do We Need Subqueries?](#2-why-do-we-need-subqueries)
3.  [Where Can a Subquery Appear?](#3-where-can-a-subquery-appear)
4.  [Scalar Subquery](#4-scalar-subquery)
5.  [Single-Row Subquery](#5-single-row-subquery)
6.  [Multi-Row Subquery](#6-multi-row-subquery)
7.  [Correlated Subquery](#7-correlated-subquery)
8.  [Nested Subquery](#8-nested-subquery)
9.  [`IN` Subquery](#9-in-subquery)
10. [`EXISTS` Subquery](#10-exists-subquery)
11. [`NOT EXISTS`](#11-not-exists)
12. [`NOT IN` and the NULL Trap](#12-not-in-and-the-null-trap)
13. [Subquery in `SELECT`](#13-subquery-in-select)
14. [Subquery in `FROM` --- Derived
    Table](#14-subquery-in-from--derived-table)
15. [Subquery in `HAVING`](#15-subquery-in-having)
16. [Subqueries in `INSERT`, `UPDATE`, and
    `DELETE`](#16-subqueries-in-insert-update-and-delete)
17. [Uncorrelated vs Correlated
    Subqueries](#17-uncorrelated-vs-correlated-subqueries)
18. [How SQL Engines Execute
    Subqueries](#18-how-sql-engines-execute-subqueries)
19. [Subquery vs JOIN](#19-subquery-vs-join)
20. [Subquery vs CTE](#20-subquery-vs-cte)
21. [Performance Considerations](#21-performance-considerations)
22. [Common Mistakes](#22-common-mistakes)
23. [Real-World Patterns](#23-real-world-patterns)
24. [Subqueries and Window
    Functions](#24-subqueries-and-window-functions)
25. [Advanced Mental Models](#25-advanced-mental-models)
26. [Interview Questions](#26-interview-questions)
27. [Quick Cheat Sheet](#27-quick-cheat-sheet)
28. [Final Mental Model](#28-final-mental-model)

------------------------------------------------------------------------

# 1. What Is a Subquery?

A **subquery** is a SQL query written inside another SQL statement.

In other words:

> A subquery is a query whose result is used by another query.

Basic structure:

``` sql
SELECT ...
FROM ...
WHERE column operator (
    SELECT ...
    FROM ...
    WHERE ...
);
```

The inner query is the **subquery**.

The outer query is the **outer query** or **parent query**.

------------------------------------------------------------------------

## 1.1 Simple Example

Suppose we have:

### `employees`

    employee_id name      salary   department_id
  ------------- ------- -------- ---------------
              1 Alice      70000              10
              2 Bob        90000              20
              3 Carol      80000              10
              4 David      60000              30

We want:

> Find employees whose salary is greater than the average salary.

First, conceptually:

``` sql
SELECT AVG(salary)
FROM employees;
```

Suppose the result is:

``` text
75000
```

Then:

``` sql
SELECT *
FROM employees
WHERE salary > 75000;
```

A subquery combines those two ideas:

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```

The inner query calculates the average.

The outer query compares every employee's salary against that result.

------------------------------------------------------------------------

# 2. Why Do We Need Subqueries?

Subqueries are useful when the result of one query becomes an input to
another query.

They are particularly useful for:

-   comparing against an aggregate
-   checking whether related records exist
-   filtering using another table
-   finding missing records
-   finding maximum/minimum values
-   performing conditional calculations
-   building intermediate result sets
-   expressing multi-step logic
-   updating or deleting rows based on another query

------------------------------------------------------------------------

## 2.1 Think of a Subquery as a Temporary Answer

A useful mental model is:

``` text
Inner query
    ↓
produces a result
    ↓
outer query uses that result
```

Example:

``` sql
SELECT *
FROM products
WHERE price > (
    SELECT AVG(price)
    FROM products
);
```

Logical idea:

``` text
Calculate average price
        ↓
      1250
        ↓
Find products where price > 1250
```

------------------------------------------------------------------------

## 2.2 Subqueries Let You Express Multi-Step Logic

Without a subquery, you might need to calculate something separately.

For example:

``` sql
SELECT AVG(salary)
FROM employees;
```

Then manually insert the result:

``` sql
SELECT *
FROM employees
WHERE salary > 75000;
```

With a subquery:

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```

The database performs the complete operation as one statement.

------------------------------------------------------------------------

# 3. Where Can a Subquery Appear?

Subqueries are not restricted to `WHERE`.

They can appear in several places.

------------------------------------------------------------------------

## 3.1 Subquery in `WHERE`

Probably the most common form.

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

------------------------------------------------------------------------

## 3.2 Subquery in `SELECT`

``` sql
SELECT
    name,
    salary,
    (
        SELECT AVG(salary)
        FROM employees
    ) AS company_avg_salary
FROM employees;
```

Example output:

  name      salary   company_avg_salary
  ------- -------- --------------------
  Alice      70000                75000
  Bob        90000                75000
  Carol      80000                75000

The scalar subquery produces one value.

------------------------------------------------------------------------

## 3.3 Subquery in `FROM`

A subquery in `FROM` is commonly called a:

-   derived table
-   inline view

Example:

``` sql
SELECT department_id, avg_salary
FROM (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) AS dept_stats;
```

The inner query produces a temporary relational result.

------------------------------------------------------------------------

## 3.4 Subquery in `HAVING`

``` sql
SELECT department_id, AVG(salary) AS avg_salary
FROM employees
GROUP BY department_id
HAVING AVG(salary) > (
    SELECT AVG(salary)
    FROM employees
);
```

This finds departments whose average salary exceeds the company-wide
average.

------------------------------------------------------------------------

## 3.5 Subquery in `UPDATE`

``` sql
UPDATE employees
SET salary = salary * 1.10
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

------------------------------------------------------------------------

## 3.6 Subquery in `DELETE`

``` sql
DELETE FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE status = 'CLOSED'
);
```

Be extremely careful with `DELETE` + subqueries.

Always test the corresponding `SELECT` first:

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE status = 'CLOSED'
);
```

------------------------------------------------------------------------

# 4. Scalar Subquery

A **scalar subquery** returns:

> Exactly one column and at most one row as a single value.

Typical usage:

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```

The subquery:

``` sql
SELECT AVG(salary)
FROM employees;
```

returns one value.

For example:

``` text
75000
```

Therefore:

``` sql
salary > 75000
```

------------------------------------------------------------------------

## 4.1 Scalar Subquery Shape

Conceptually:

``` text
+----------------+
|      75000     |
+----------------+
```

One row.

One column.

One value.

------------------------------------------------------------------------

## 4.2 Scalar Subquery in `SELECT`

``` sql
SELECT
    name,
    salary,
    (
        SELECT AVG(salary)
        FROM employees
    ) AS avg_salary
FROM employees;
```

The same scalar value can appear for every outer row.

------------------------------------------------------------------------

## 4.3 Scalar Subquery with `MAX`

Find employees with the highest salary:

``` sql
SELECT *
FROM employees
WHERE salary = (
    SELECT MAX(salary)
    FROM employees
);
```

Notice an important detail:

The subquery returns **one value**, but the outer query can return
**multiple rows**.

If three employees share the maximum salary, all three are returned.

------------------------------------------------------------------------

## 4.4 Scalar Subquery Returning Zero Rows

Behavior can depend on SQL context and database engine, but in a scalar
expression a subquery that produces no row commonly evaluates to `NULL`.

Example concept:

``` sql
SELECT (
    SELECT salary
    FROM employees
    WHERE employee_id = 999999
);
```

If no employee exists, the scalar result is generally:

``` text
NULL
```

Always consult the specific database documentation when relying on
edge-case behavior.

------------------------------------------------------------------------

## 4.5 Scalar Subquery Returning Multiple Rows

This is a major mistake:

``` sql
SELECT *
FROM employees
WHERE salary = (
    SELECT salary
    FROM employees
    WHERE department_id = 10
);
```

Suppose department 10 contains:

``` text
70000
80000
```

The subquery returns two rows.

But `=` expects a scalar value.

Conceptually:

``` text
salary = ???
```

There is no single value.

Many database systems therefore raise an error similar to:

``` text
Subquery returned more than 1 row
```

The correct solution may be:

``` sql
WHERE salary IN (
    SELECT salary
    FROM employees
    WHERE department_id = 10
);
```

or perhaps:

``` sql
WHERE salary > (
    SELECT MAX(salary)
    FROM employees
    WHERE department_id = 10
);
```

The correct choice depends on the business requirement.

------------------------------------------------------------------------

# 5. Single-Row Subquery

A **single-row subquery** returns at most one row.

A common example:

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT salary
    FROM employees
    WHERE employee_id = 5
);
```

The inner query is expected to identify one employee.

------------------------------------------------------------------------

## 5.1 Single-Row Operators

Single-row subqueries commonly work with:

``` text
=
<>
!=
>
<
>=
<=
```

Example:

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT salary
    FROM employees
    WHERE employee_id = 5
);
```

------------------------------------------------------------------------

## 5.2 Scalar vs Single-Row

These terms are closely related but not identical.

### Scalar subquery

Returns:

``` text
1 row × 1 column
```

Example:

``` sql
SELECT MAX(salary)
FROM employees;
```

### Single-row subquery

Means the result contains:

``` text
at most 1 row
```

A single-row result could conceptually contain multiple columns:

``` text
1 row × N columns
```

Some SQL dialects support row-value comparisons such as:

``` sql
WHERE (department_id, salary) = (
    SELECT department_id, salary
    FROM employees
    WHERE employee_id = 5
);
```

For everyday SQL, however, scalar subqueries are the most common
single-row form.

------------------------------------------------------------------------

# 6. Multi-Row Subquery

A **multi-row subquery** returns multiple rows.

Example:

``` sql
SELECT department_id
FROM departments
WHERE location = 'Delhi';
```

Suppose it returns:

``` text
10
20
30
```

You cannot normally write:

``` sql
WHERE department_id = (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

because the inner query returns multiple values.

Instead:

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

------------------------------------------------------------------------

## 6.1 Operators for Multi-Row Subqueries

Common choices:

### `IN`

``` sql
WHERE department_id IN (
    SELECT department_id
    FROM departments
);
```

### `EXISTS`

``` sql
WHERE EXISTS (
    SELECT 1
    FROM departments d
    WHERE d.department_id = e.department_id
);
```

### `ANY` / `SOME`

Example:

``` sql
WHERE salary > ANY (
    SELECT salary
    FROM employees
    WHERE department_id = 10
);
```

Meaning:

> Salary is greater than at least one value returned by the subquery.

### `ALL`

``` sql
WHERE salary > ALL (
    SELECT salary
    FROM employees
    WHERE department_id = 10
);
```

Meaning:

> Salary is greater than every value returned by the subquery.

`ANY` and `ALL` are less commonly used than `IN` and `EXISTS`, but they
are important for advanced SQL.

------------------------------------------------------------------------

# 7. Correlated Subquery

A **correlated subquery** is a subquery that references a column from
the outer query.

This is one of the most important subquery concepts.

Example:

``` sql
SELECT
    e.employee_id,
    e.name,
    e.salary,
    e.department_id
FROM employees e
WHERE e.salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

Look carefully:

``` sql
e2.department_id = e.department_id
```

`e2` belongs to the inner query.

`e` belongs to the outer query.

Therefore, the inner query depends on the current outer row.

------------------------------------------------------------------------

## 7.1 What Is Happening Conceptually?

For Alice:

``` text
Alice.department_id = 10
```

The subquery effectively becomes:

``` sql
SELECT AVG(salary)
FROM employees
WHERE department_id = 10;
```

Suppose:

``` text
department 10 average = 75000
```

Then:

``` text
Alice salary > 75000?
```

For Bob:

``` text
Bob.department_id = 20
```

The subquery conceptually becomes:

``` sql
SELECT AVG(salary)
FROM employees
WHERE department_id = 20;
```

Then Bob's salary is compared against department 20's average.

So the logical dependency is:

``` text
Outer row
   ↓
inner query uses outer-row value
   ↓
inner result
   ↓
outer predicate
```

------------------------------------------------------------------------

## 7.2 Correlated Subquery Example

Find employees who earn more than their own department's average:

``` sql
SELECT *
FROM employees e
WHERE e.salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

This is a classic interview problem.

------------------------------------------------------------------------

## 7.3 Correlated `EXISTS`

A very common pattern:

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

Meaning:

> Return customers for whom at least one order exists.

The subquery references:

``` sql
c.customer_id
```

from the outer query.

Therefore it is correlated.

------------------------------------------------------------------------

## 7.4 Why Use `SELECT 1`?

Inside:

``` sql
EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
)
```

The actual selected value is not important.

`EXISTS` asks only:

``` text
Does at least one matching row exist?
```

So these are usually equivalent in intent:

``` sql
EXISTS (
    SELECT 1
    FROM orders o
    WHERE ...
)
```

``` sql
EXISTS (
    SELECT *
    FROM orders o
    WHERE ...
)
```

``` sql
EXISTS (
    SELECT o.order_id
    FROM orders o
    WHERE ...
)
```

`SELECT 1` is simply a clear convention.

------------------------------------------------------------------------

# 8. Nested Subquery

A **nested subquery** is a subquery inside another subquery.

Example:

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE manager_id IN (
        SELECT employee_id
        FROM employees
        WHERE salary > 100000
    )
);
```

There are multiple query levels:

``` text
Outer query
    ↓
Subquery level 1
    ↓
Subquery level 2
```

------------------------------------------------------------------------

## 8.1 Understanding Nested Logic

Suppose:

``` sql
SELECT employee_id
FROM employees
WHERE salary > 100000;
```

returns managers:

``` text
101
102
```

Then:

``` sql
SELECT department_id
FROM departments
WHERE manager_id IN (101, 102);
```

returns:

``` text
10
20
```

Then:

``` sql
SELECT *
FROM employees
WHERE department_id IN (10, 20);
```

returns employees in those departments.

------------------------------------------------------------------------

## 8.2 Deep Nesting Can Hurt Readability

This:

``` sql
SELECT *
FROM A
WHERE x IN (
    SELECT x
    FROM B
    WHERE y IN (
        SELECT y
        FROM C
        WHERE z IN (
            SELECT z
            FROM D
        )
    )
);
```

may be valid SQL.

But it is difficult to understand and maintain.

Often a CTE can make the logic clearer:

``` sql
WITH d_values AS (
    SELECT z
    FROM D
),
c_values AS (
    SELECT y
    FROM C
    WHERE z IN (SELECT z FROM d_values)
),
b_values AS (
    SELECT x
    FROM B
    WHERE y IN (SELECT y FROM c_values)
)
SELECT *
FROM A
WHERE x IN (SELECT x FROM b_values);
```

A CTE is not automatically faster; it is primarily a
readability/organization tool unless the database optimizer treats it
specially.

------------------------------------------------------------------------

# 9. `IN` Subquery

`IN` checks whether a value belongs to a set of values.

Basic syntax:

``` sql
SELECT ...
FROM table1
WHERE column IN (
    SELECT column
    FROM table2
);
```

Example:

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

Meaning:

> Return employees whose department is one of the departments located in
> Delhi.

------------------------------------------------------------------------

## 9.1 `IN` Is About Membership

Conceptually:

``` text
Subquery result:

10
20
30
```

Then:

``` sql
department_id IN (10, 20, 30)
```

is evaluated.

------------------------------------------------------------------------

## 9.2 `IN` vs Multiple `OR`s

This:

``` sql
WHERE department_id IN (10, 20, 30)
```

is conceptually similar to:

``` sql
WHERE department_id = 10
   OR department_id = 20
   OR department_id = 30
```

For a subquery:

``` sql
WHERE department_id IN (
    SELECT department_id
    FROM departments
)
```

the set is generated dynamically.

------------------------------------------------------------------------

## 9.3 `IN` and Duplicate Values

Suppose the subquery returns:

``` text
10
10
10
20
```

For membership testing, duplicates generally do not change the logical
result.

The question is simply:

``` text
Is 10 present?
```

Therefore:

``` sql
WHERE x IN (...)
```

does not inherently multiply outer rows.

This is an important difference from a normal `JOIN`.

------------------------------------------------------------------------

# 10. `EXISTS` Subquery

`EXISTS` asks:

> Does the subquery return at least one row?

Syntax:

``` sql
WHERE EXISTS (
    SELECT 1
    FROM ...
    WHERE ...
)
```

------------------------------------------------------------------------

## 10.1 Basic Example

Find customers who have at least one order:

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

------------------------------------------------------------------------

## 10.2 Why `EXISTS` Is Powerful

Suppose customer 101 has 100 orders.

A normal join:

``` sql
SELECT c.*
FROM customers c
JOIN orders o
    ON o.customer_id = c.customer_id;
```

can produce:

``` text
Customer 101
Customer 101
Customer 101
...
100 rows
```

If your actual requirement is:

> Does this customer have at least one order?

then those 100 rows are unnecessary.

`EXISTS` expresses the requirement directly:

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

The logical result contains the customer once.

------------------------------------------------------------------------

## 10.3 `EXISTS` Is a Semi-Join Concept

Relationally, `EXISTS` often corresponds to a **semi-join**.

A semi-join means:

> Keep rows from the left side when a match exists on the right side.

It does not return columns from the right side.

Conceptually:

``` text
Customers
    │
    │ existence check
    ▼
Orders
```

Keep the customer if a matching order exists.

------------------------------------------------------------------------

## 10.4 `EXISTS` Does Not Care About Selected Columns

These are logically existence tests:

``` sql
EXISTS (
    SELECT 1
    FROM orders o
    WHERE ...
)
```

``` sql
EXISTS (
    SELECT *
    FROM orders o
    WHERE ...
)
```

``` sql
EXISTS (
    SELECT o.order_id
    FROM orders o
    WHERE ...
)
```

The important part is whether at least one row satisfies the predicate.

------------------------------------------------------------------------

## 10.5 `EXISTS` and Early Termination

Conceptually, once a matching row has been found, the existence
condition is already true.

For example:

``` sql
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

The database does not logically need to count all matching orders.

However:

> Do not assume a literal procedural "run the subquery from scratch for
> every row" implementation.

Modern optimizers can transform `EXISTS` into efficient semi-join
strategies.

------------------------------------------------------------------------

# 11. `NOT EXISTS`

`NOT EXISTS` asks:

> Does the subquery return zero rows?

Example:

``` sql
SELECT c.*
FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

Meaning:

> Find customers who have never placed an order.

This is an **anti-join / anti-semi-join** pattern.

------------------------------------------------------------------------

## 11.1 Finding Missing Relationships

Examples:

### Customers without orders

``` sql
SELECT c.*
FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

### Products never sold

``` sql
SELECT p.*
FROM products p
WHERE NOT EXISTS (
    SELECT 1
    FROM order_items oi
    WHERE oi.product_id = p.product_id
);
```

### Employees with no assigned project

``` sql
SELECT e.*
FROM employees e
WHERE NOT EXISTS (
    SELECT 1
    FROM employee_projects ep
    WHERE ep.employee_id = e.employee_id
);
```

------------------------------------------------------------------------

# 12. `NOT IN` and the NULL Trap

This is one of the most important SQL subquery pitfalls.

Consider:

``` sql
SELECT *
FROM employees
WHERE department_id NOT IN (
    SELECT department_id
    FROM departments
);
```

What happens if the subquery returns:

``` text
10
20
NULL
```

SQL uses **three-valued logic**:

``` text
TRUE
FALSE
UNKNOWN
```

The presence of `NULL` can cause `NOT IN` to behave unexpectedly.

------------------------------------------------------------------------

## 12.1 Why?

Suppose:

``` sql
department_id = 30
```

and the subquery produces:

``` text
10
20
NULL
```

Conceptually:

``` text
30 <> 10  → TRUE
30 <> 20  → TRUE
30 <> NULL → UNKNOWN
```

The complete `NOT IN` condition becomes affected by the `UNKNOWN`
result.

In a `WHERE` clause, only `TRUE` passes.

Therefore, rows you expected to receive may disappear.

------------------------------------------------------------------------

## 12.2 Safer Pattern: `NOT EXISTS`

Instead of:

``` sql
WHERE department_id NOT IN (
    SELECT department_id
    FROM departments
);
```

often prefer:

``` sql
WHERE NOT EXISTS (
    SELECT 1
    FROM departments d
    WHERE d.department_id = e.department_id
);
```

This compares matching rows directly and avoids the same `NOT IN`/`NULL`
trap.

------------------------------------------------------------------------

## 12.3 Important Rule

When checking:

> "There is no matching row"

think:

``` sql
NOT EXISTS
```

rather than blindly using:

``` sql
NOT IN
```

especially when the subquery column may contain `NULL`.

------------------------------------------------------------------------

# 13. Subquery in `SELECT`

A subquery can be used as a calculated column.

Example:

``` sql
SELECT
    e.name,
    e.salary,
    (
        SELECT AVG(e2.salary)
        FROM employees e2
    ) AS company_avg_salary
FROM employees e;
```

------------------------------------------------------------------------

## 13.1 Scalar Subquery per Output Row

The subquery returns one scalar value:

``` text
75000
```

The outer query can attach it to every row.

Conceptually:

``` text
Alice  70000  75000
Bob    90000  75000
Carol  80000  75000
```

------------------------------------------------------------------------

## 13.2 Correlated Scalar Subquery

A more advanced example:

``` sql
SELECT
    e.name,
    e.salary,
    (
        SELECT AVG(e2.salary)
        FROM employees e2
        WHERE e2.department_id = e.department_id
    ) AS department_avg
FROM employees e;
```

Now the scalar value depends on the current employee's department.

------------------------------------------------------------------------

## 13.3 Performance Warning

A correlated scalar subquery may be expensive if the optimizer cannot
efficiently transform it.

For example:

``` sql
SELECT
    e.name,
    (
        SELECT AVG(e2.salary)
        FROM employees e2
        WHERE e2.department_id = e.department_id
    )
FROM employees e;
```

If there are millions of employees, repeatedly calculating department
averages can be costly in a naive execution strategy.

A grouped query plus join may be clearer and more efficient:

``` sql
SELECT
    e.name,
    e.salary,
    d.avg_salary
FROM employees e
JOIN (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) d
    ON d.department_id = e.department_id;
```

Again:

> Do not assume the JOIN is automatically faster. Check the execution
> plan.

------------------------------------------------------------------------

# 14. Subquery in `FROM` --- Derived Table

A subquery in `FROM` creates an intermediate relational result.

Example:

``` sql
SELECT
    department_id,
    avg_salary
FROM (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) AS dept_stats;
```

The inner query produces:

    department_id   avg_salary
  --------------- ------------
               10        75000
               20        95000
               30        60000

The outer query treats that result like a table.

------------------------------------------------------------------------

## 14.1 Why Use a Derived Table?

Useful when you want to:

1.  aggregate first
2.  filter the aggregate result later
3.  transform a dataset
4.  join an intermediate result
5.  break a complex query into logical stages

------------------------------------------------------------------------

## 14.2 Example: Departments Above Company Average

``` sql
SELECT *
FROM (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) AS dept_stats
WHERE avg_salary > (
    SELECT AVG(salary)
    FROM employees
);
```

Logical stages:

``` text
employees
    ↓
GROUP BY department
    ↓
department averages
    ↓
compare each department average
    ↓
company average
```

------------------------------------------------------------------------

# 15. Subquery in `HAVING`

`HAVING` filters groups after aggregation.

Example:

``` sql
SELECT
    department_id,
    AVG(salary) AS avg_salary
FROM employees
GROUP BY department_id
HAVING AVG(salary) > (
    SELECT AVG(salary)
    FROM employees
);
```

Meaning:

> Return departments whose average salary is greater than the overall
> company average salary.

------------------------------------------------------------------------

## 15.1 Why `HAVING` Instead of `WHERE`?

`WHERE` filters individual rows before grouping.

`HAVING` filters groups after grouping.

Conceptually:

``` text
FROM
 ↓
WHERE
 ↓
GROUP BY
 ↓
HAVING
 ↓
SELECT
```

So:

``` sql
WHERE salary > ...
```

works at row level.

While:

``` sql
HAVING AVG(salary) > ...
```

works at group level.

------------------------------------------------------------------------

# 16. Subqueries in `INSERT`, `UPDATE`, and `DELETE`

Subqueries are extremely useful in DML.

------------------------------------------------------------------------

## 16.1 `INSERT ... SELECT`

This is technically a query feeding another statement:

``` sql
INSERT INTO high_salary_employees (
    employee_id,
    name,
    salary
)
SELECT
    employee_id,
    name,
    salary
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```

This copies employees earning above average into another table.

------------------------------------------------------------------------

## 16.2 `UPDATE` with Subquery

Give employees in a particular department a raise:

``` sql
UPDATE employees
SET salary = salary * 1.10
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

------------------------------------------------------------------------

## 16.3 Correlated `UPDATE`

A more advanced pattern:

``` sql
UPDATE employees e
SET salary = (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
)
WHERE e.salary < (
    SELECT AVG(e3.salary)
    FROM employees e3
    WHERE e3.department_id = e.department_id
);
```

This kind of statement is powerful but can become difficult to reason
about.

Always test the corresponding `SELECT` first.

------------------------------------------------------------------------

## 16.4 `DELETE` with `NOT EXISTS`

Delete customers with no orders:

``` sql
DELETE FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

This is an anti-join pattern.

Before executing it, test:

``` sql
SELECT c.*
FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

------------------------------------------------------------------------

# 17. Uncorrelated vs Correlated Subqueries

This distinction is fundamental.

------------------------------------------------------------------------

## 17.1 Uncorrelated Subquery

An uncorrelated subquery does not reference the outer query.

Example:

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```

The inner query is independent.

------------------------------------------------------------------------

## 17.2 Correlated Subquery

A correlated subquery references the outer query.

Example:

``` sql
SELECT *
FROM employees e
WHERE salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

The inner query depends on:

``` sql
e.department_id
```

from the outer query.

------------------------------------------------------------------------

## 17.3 Comparison

  Property                  Uncorrelated            Correlated
  ------------------------- ----------------------- ------------------------------
  References outer query?   No                      Yes
  Independent?              Yes                     No
  Result can be reused?     Often easier            Depends
  Typical use               Global aggregate        Per-row/per-group comparison
  Example                   Company average         Department average
  Optimization              Often straightforward   May require decorrelation

------------------------------------------------------------------------

# 18. How SQL Engines Execute Subqueries

This is where theory becomes important.

A common beginner mental model is:

``` text
Run outer query row 1
    ↓
Run inner query
    ↓
Run outer query row 2
    ↓
Run inner query
    ↓
...
```

That is a useful **logical model** for understanding correlated
subqueries, but it is not necessarily how the database physically
executes the query.

Modern query optimizers can transform SQL substantially.

------------------------------------------------------------------------

## 18.1 Logical Query vs Physical Execution

You write:

``` sql
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
)
```

The optimizer may transform this into a semi-join-like execution
strategy.

The database might use:

-   nested-loop join
-   hash-based strategy
-   index lookup
-   bitmap strategy
-   semi-join
-   anti-join
-   materialization
-   predicate pushdown
-   subquery decorrelation

The exact behavior depends on:

-   database engine
-   indexes
-   statistics
-   table sizes
-   selectivity
-   constraints
-   query structure
-   optimizer version

------------------------------------------------------------------------

## 18.2 Subquery Decorrelation

Suppose you write:

``` sql
SELECT *
FROM employees e
WHERE EXISTS (
    SELECT 1
    FROM departments d
    WHERE d.department_id = e.department_id
);
```

The optimizer may recognize that this is essentially an existence join.

Conceptually:

``` text
EXISTS
  ↓
semi-join
```

Therefore:

> A correlated subquery is not automatically slow.

The important question is what execution plan the database chooses.

------------------------------------------------------------------------

## 18.3 Materialization

A database may choose to materialize an intermediate subquery result.

Conceptually:

``` text
Run subquery
    ↓
store intermediate result
    ↓
use result
```

But whether and when this happens is engine-specific.

Do not assume every subquery creates a temporary physical table.

------------------------------------------------------------------------

# 19. Subquery vs JOIN

This is one of the most common SQL interview questions.

The answer is:

> Neither is universally better. The correct choice depends on the
> semantics, cardinality, readability, and optimizer plan.

------------------------------------------------------------------------

## 19.1 Example: Find Customers With Orders

Using `EXISTS`:

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

Using `JOIN`:

``` sql
SELECT DISTINCT c.*
FROM customers c
JOIN orders o
    ON o.customer_id = c.customer_id;
```

Both can answer:

> Which customers have at least one order?

But the semantics are different.

------------------------------------------------------------------------

## 19.2 Why `EXISTS` Is Often More Expressive

The business requirement is:

> Keep the customer if a matching order exists.

`EXISTS` directly expresses that:

``` sql
WHERE EXISTS (...)
```

The `JOIN` version creates matching row combinations first and then
needs:

``` sql
DISTINCT
```

to remove duplicates.

------------------------------------------------------------------------

## 19.3 Example: We Need Columns from the Other Table

Suppose we need:

``` text
customer name
order date
order amount
```

Then a `JOIN` is natural:

``` sql
SELECT
    c.name,
    o.order_date,
    o.amount
FROM customers c
JOIN orders o
    ON o.customer_id = c.customer_id;
```

`EXISTS` is not the natural choice because `EXISTS` only answers whether
a matching row exists.

------------------------------------------------------------------------

## 19.4 Example: We Only Need Existence

Requirement:

> Find customers who have placed at least one order.

Prefer a direct existence expression:

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

------------------------------------------------------------------------

## 19.5 `IN` vs `JOIN`

Suppose:

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

Possible join:

``` sql
SELECT e.*
FROM employees e
JOIN departments d
    ON d.department_id = e.department_id
WHERE d.location = 'Delhi';
```

These can often be optimized similarly.

But be careful:

If the join side contains duplicate keys, a plain `JOIN` can multiply
rows.

`IN` is membership logic and does not have that same row-multiplication
behavior.

------------------------------------------------------------------------

## 19.6 Example of Accidental Row Multiplication

Suppose:

``` text
employees
---------
employee_id
1

departments
-----------
department_id
10
10
```

Then:

``` sql
SELECT e.*
FROM employees e
JOIN departments d
    ON d.department_id = e.department_id;
```

can return the employee twice.

But:

``` sql
SELECT *
FROM employees e
WHERE e.department_id IN (
    SELECT department_id
    FROM departments
);
```

still returns the employee once.

Therefore:

> Never replace a subquery with a JOIN mechanically without checking
> cardinality.

------------------------------------------------------------------------

## 19.7 Comparison Table

  Requirement                                          Natural choice
  ---------------------------------------------------- --------------------------------
  Need columns from both tables                        `JOIN`
  Only need to know whether a match exists             `EXISTS`
  Need to know whether no match exists                 `NOT EXISTS`
  Check membership in a set                            `IN`
  Compare against aggregate                            Scalar subquery
  Build intermediate relation                          Derived table / CTE
  Complex multi-stage transformation                   CTE often improves readability
  Need to preserve left-side rows with optional data   `LEFT JOIN`

------------------------------------------------------------------------

# 20. Subquery vs CTE

A CTE is introduced using:

``` sql
WITH ...
```

Example:

``` sql
WITH department_stats AS (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
)
SELECT *
FROM department_stats;
```

A derived-table version is:

``` sql
SELECT *
FROM (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) AS department_stats;
```

They can express similar logical transformations.

------------------------------------------------------------------------

## 20.1 Why Use a CTE?

CTEs are often easier to read when the query has multiple stages.

Example:

``` sql
WITH department_stats AS (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
),
above_average_departments AS (
    SELECT department_id
    FROM department_stats
    WHERE avg_salary > 80000
)
SELECT e.*
FROM employees e
WHERE e.department_id IN (
    SELECT department_id
    FROM above_average_departments
);
```

The logic becomes:

``` text
employees
    ↓
department_stats
    ↓
above_average_departments
    ↓
final employees
```

------------------------------------------------------------------------

## 20.2 CTE Does Not Automatically Mean Faster

A common misconception:

``` text
CTE = optimized temporary table
```

Not necessarily.

Depending on the database engine, a CTE may be:

-   inlined
-   materialized
-   optimized away
-   transformed

Always inspect the execution plan when performance matters.

------------------------------------------------------------------------

# 21. Performance Considerations

Subqueries can be extremely efficient.

They can also become expensive if poorly structured.

The key principle:

> Do not optimize based solely on whether syntax contains a subquery.
> Optimize based on the actual execution plan and data distribution.

------------------------------------------------------------------------

## 21.1 Index Correlation Columns

For:

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

An index such as:

``` text
orders(customer_id)
```

can be highly useful.

------------------------------------------------------------------------

## 21.2 Index Filtering Columns

For:

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

Useful indexes may include:

``` text
departments(location, department_id)
employees(department_id)
```

The optimal index depends on the actual workload and database engine.

------------------------------------------------------------------------

## 21.3 Correlated Subquery Warning

Consider:

``` sql
SELECT *
FROM employees e
WHERE salary > (
    SELECT AVG(salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

If the optimizer cannot decorrelate this efficiently, repeated work may
occur.

A grouped derived table can sometimes be preferable:

``` sql
SELECT e.*
FROM employees e
JOIN (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) d
    ON d.department_id = e.department_id
WHERE e.salary > d.avg_salary;
```

This computes department averages as a grouped relation.

But again:

> Measure rather than assume.

------------------------------------------------------------------------

## 21.4 Avoid Unnecessary Scalar Repetition

Suppose a scalar subquery computes an expensive aggregate:

``` sql
SELECT
    e.name,
    (
        SELECT expensive_calculation(...)
        FROM ...
    ) AS value
FROM employees e;
```

If the calculation is logically the same for every employee, consider
whether it can be computed once through:

-   a derived table
-   a CTE
-   a window function
-   a join
-   a pre-aggregated relation

------------------------------------------------------------------------

## 21.5 Use `EXISTS` for Existence Logic

If you only need to know whether a matching row exists:

``` sql
WHERE EXISTS (...)
```

is often a clearer expression than:

``` sql
JOIN ...
DISTINCT
```

It also avoids accidentally returning multiple copies of the outer row.

------------------------------------------------------------------------

## 21.6 Always Inspect the Execution Plan

Useful tools depend on your database:

``` sql
EXPLAIN
```

or:

``` sql
EXPLAIN ANALYZE
```

For performance debugging, inspect:

-   access method
-   index usage
-   estimated rows
-   actual rows
-   join strategy
-   scans
-   sorts
-   hash operations
-   loops
-   materialization
-   filter selectivity

------------------------------------------------------------------------

# 22. Common Mistakes

------------------------------------------------------------------------

## 22.1 Using `=` with a Multi-Row Subquery

Wrong:

``` sql
WHERE department_id = (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

if multiple departments can be in Delhi.

Potential fix:

``` sql
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

------------------------------------------------------------------------

## 22.2 Forgetting Correlation

Suppose you want:

> Employees earning above their department average.

Incorrect:

``` sql
SELECT *
FROM employees e
WHERE e.salary > (
    SELECT AVG(salary)
    FROM employees
);
```

This compares against the **company average**, not the department
average.

Correct:

``` sql
SELECT *
FROM employees e
WHERE e.salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

------------------------------------------------------------------------

## 22.3 Confusing `EXISTS` with `IN`

`IN` asks:

> Is this value a member of this result set?

`EXISTS` asks:

> Does at least one matching row exist?

Example:

``` sql
WHERE department_id IN (...)
```

vs:

``` sql
WHERE EXISTS (...)
```

Both can solve similar problems, but they express different concepts.

------------------------------------------------------------------------

## 22.4 `NOT IN` with NULL

Potentially dangerous:

``` sql
WHERE id NOT IN (
    SELECT id
    FROM other_table
);
```

If `other_table.id` contains `NULL`, the result may be surprising.

Often safer:

``` sql
WHERE NOT EXISTS (
    SELECT 1
    FROM other_table o
    WHERE o.id = t.id
);
```

------------------------------------------------------------------------

## 22.5 Assuming Every Correlated Subquery Is Slow

This is an oversimplification.

The optimizer may transform a correlated query into an efficient
join-like plan.

The correct approach:

``` text
Understand query
    ↓
Check indexes/statistics
    ↓
Inspect execution plan
    ↓
Measure
    ↓
Rewrite if necessary
```

------------------------------------------------------------------------

## 22.6 Replacing Every Subquery with a JOIN

Bad optimization instinct:

``` text
subquery = slow
JOIN = fast
```

This is not universally true.

A JOIN can introduce:

-   duplicate rows
-   row multiplication
-   unnecessary columns
-   more complicated grouping
-   semantic changes

Choose the construct that matches the intended relational operation.

------------------------------------------------------------------------

# 23. Real-World Patterns

Here are several patterns worth memorizing.

------------------------------------------------------------------------

## 23.1 Employees Above Company Average

``` sql
SELECT *
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```

**Type:** scalar subquery.

------------------------------------------------------------------------

## 23.2 Employees Above Department Average

``` sql
SELECT *
FROM employees e
WHERE salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

**Type:** correlated scalar subquery.

------------------------------------------------------------------------

## 23.3 Employees in Departments Located in Delhi

``` sql
SELECT *
FROM employees
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

**Type:** multi-row `IN` subquery.

------------------------------------------------------------------------

## 23.4 Customers With at Least One Order

``` sql
SELECT c.*
FROM customers c
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

**Type:** correlated `EXISTS`.

------------------------------------------------------------------------

## 23.5 Customers With No Orders

``` sql
SELECT c.*
FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

**Type:** correlated `NOT EXISTS`.

------------------------------------------------------------------------

## 23.6 Products Never Ordered

``` sql
SELECT p.*
FROM products p
WHERE NOT EXISTS (
    SELECT 1
    FROM order_items oi
    WHERE oi.product_id = p.product_id
);
```

**Type:** anti-semi-join pattern.

------------------------------------------------------------------------

## 23.7 Highest-Paid Employee

``` sql
SELECT *
FROM employees
WHERE salary = (
    SELECT MAX(salary)
    FROM employees
);
```

------------------------------------------------------------------------

## 23.8 All Employees With Maximum Salary

The above query correctly returns all employees tied for maximum salary.

That is better than:

``` sql
SELECT *
FROM employees
ORDER BY salary DESC
LIMIT 1;
```

when the requirement is:

> Return every employee tied for the highest salary.

------------------------------------------------------------------------

## 23.9 Departments Above Company Average

``` sql
SELECT
    department_id,
    AVG(salary) AS avg_salary
FROM employees
GROUP BY department_id
HAVING AVG(salary) > (
    SELECT AVG(salary)
    FROM employees
);
```

------------------------------------------------------------------------

## 23.10 Find Duplicate Email Addresses

A subquery approach:

``` sql
SELECT *
FROM users u
WHERE u.email IN (
    SELECT email
    FROM users
    GROUP BY email
    HAVING COUNT(*) > 1
);
```

The inner query finds duplicated emails.

The outer query retrieves the actual rows.

------------------------------------------------------------------------

## 23.11 Find Latest Order Per Customer

A correlated subquery approach:

``` sql
SELECT o.*
FROM orders o
WHERE o.order_date = (
    SELECT MAX(o2.order_date)
    FROM orders o2
    WHERE o2.customer_id = o.customer_id
);
```

This can return multiple rows if a customer has multiple orders with the
same maximum timestamp.

If you need exactly one row per customer, a window function with
deterministic tie-breaking is often better.

------------------------------------------------------------------------

## 23.12 Employees Managed by Highly Paid Managers

``` sql
SELECT *
FROM employees
WHERE manager_id IN (
    SELECT employee_id
    FROM employees
    WHERE salary > 150000
);
```

This is a self-referencing subquery.

------------------------------------------------------------------------

# 24. Subqueries and Window Functions

Many problems traditionally solved with correlated subqueries can be
solved more naturally with window functions.

------------------------------------------------------------------------

## 24.1 Department Average

Correlated subquery:

``` sql
SELECT
    e.*,
    (
        SELECT AVG(e2.salary)
        FROM employees e2
        WHERE e2.department_id = e.department_id
    ) AS department_avg
FROM employees e;
```

Window function:

``` sql
SELECT
    e.*,
    AVG(salary) OVER (
        PARTITION BY department_id
    ) AS department_avg
FROM employees e;
```

The window version is often clearer when you want both:

-   the original row
-   an aggregate over that row's group

------------------------------------------------------------------------

## 24.2 Rank Employees Within Department

A subquery is possible, but a window function is much more natural:

``` sql
SELECT
    employee_id,
    name,
    department_id,
    salary,
    RANK() OVER (
        PARTITION BY department_id
        ORDER BY salary DESC
    ) AS salary_rank
FROM employees;
```

------------------------------------------------------------------------

## 24.3 General Rule

If your problem sounds like:

> For each row, calculate something over the rows related to it.

consider:

``` text
Window function
```

before writing a complicated correlated subquery.

But do not treat window functions as universal replacements. `EXISTS`
and `NOT EXISTS` remain excellent for existence/anti-existence logic.

------------------------------------------------------------------------

# 25. Advanced Mental Models

This section is about understanding subqueries at a deeper relational
level.

------------------------------------------------------------------------

## 25.1 A Scalar Subquery Is a Value

Think:

``` sql
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
)
```

as:

``` text
salary > VALUE
```

The subquery produces a value.

------------------------------------------------------------------------

## 25.2 An `IN` Subquery Is a Set

Think:

``` sql
WHERE department_id IN (
    SELECT department_id
    FROM departments
)
```

as:

``` text
department_id ∈ SET
```

The subquery produces a set of candidate values.

------------------------------------------------------------------------

## 25.3 An `EXISTS` Subquery Is a Boolean Predicate

Think:

``` sql
WHERE EXISTS (...)
```

as:

``` text
Does a matching row exist?
```

The subquery is being used as a logical condition.

------------------------------------------------------------------------

## 25.4 A Derived Table Is a Relation

Think:

``` sql
FROM (
    SELECT ...
) AS x
```

as:

``` text
subquery → temporary logical relation
```

The outer query treats the result like a table.

------------------------------------------------------------------------

## 25.5 `NOT EXISTS` Is Absence

Think:

``` sql
WHERE NOT EXISTS (...)
```

as:

``` text
There is no matching row.
```

This is one of the cleanest ways to express anti-join logic.

------------------------------------------------------------------------

# 26. Interview Questions

## Q1. What is a subquery?

A query nested inside another SQL statement whose result is used by the
outer statement.

------------------------------------------------------------------------

## Q2. What is a scalar subquery?

A subquery that returns one column and at most one row as a single
value.

------------------------------------------------------------------------

## Q3. What happens if a scalar subquery returns multiple rows?

A scalar context generally raises a cardinality error because a single
value was expected.

------------------------------------------------------------------------

## Q4. What is a correlated subquery?

A subquery that references columns from the outer query.

Example:

``` sql
SELECT *
FROM employees e
WHERE salary > (
    SELECT AVG(salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
);
```

------------------------------------------------------------------------

## Q5. Is a correlated subquery always slower?

No.

The optimizer may transform it into an efficient join-like plan.

Use the execution plan and actual measurements.

------------------------------------------------------------------------

## Q6. What is the difference between `IN` and `EXISTS`?

`IN` tests whether a value belongs to a result set.

`EXISTS` tests whether at least one matching row exists.

------------------------------------------------------------------------

## Q7. Why is `NOT EXISTS` often preferred over `NOT IN`?

Because `NOT IN` has tricky three-valued-logic behavior when the
subquery contains `NULL`.

`NOT EXISTS` expresses anti-existence directly and generally avoids that
specific trap.

------------------------------------------------------------------------

## Q8. Does `EXISTS` care what the subquery selects?

Usually no.

These are existence tests:

``` sql
EXISTS (SELECT 1 ...)
```

``` sql
EXISTS (SELECT * ...)
```

``` sql
EXISTS (SELECT id ...)
```

The important thing is whether at least one row satisfies the predicate.

------------------------------------------------------------------------

## Q9. Can a subquery appear in `SELECT`?

Yes.

Example:

``` sql
SELECT
    name,
    (SELECT AVG(salary) FROM employees) AS avg_salary
FROM employees;
```

------------------------------------------------------------------------

## Q10. Can a subquery appear in `FROM`?

Yes.

It becomes a derived table / inline view:

``` sql
SELECT *
FROM (
    SELECT ...
) AS x;
```

------------------------------------------------------------------------

## Q11. Can subqueries be used in `UPDATE` and `DELETE`?

Yes.

Example:

``` sql
UPDATE employees
SET salary = salary * 1.10
WHERE department_id IN (
    SELECT department_id
    FROM departments
    WHERE location = 'Delhi'
);
```

------------------------------------------------------------------------

## Q12. Which is faster: JOIN or subquery?

There is no universal answer.

Modern optimizers can transform many logically equivalent forms into
similar physical plans.

Compare execution plans and measurements.

------------------------------------------------------------------------

## Q13. Why can replacing `IN` with `JOIN` change results?

Because a `JOIN` can multiply rows when the joined table has multiple
matching rows.

`IN` is membership logic and does not inherently multiply the outer
result.

------------------------------------------------------------------------

## Q14. What is a semi-join?

A relational operation that returns rows from one side when a matching
row exists on the other side.

`EXISTS` is commonly implemented conceptually as a semi-join.

------------------------------------------------------------------------

## Q15. What is an anti-join?

An operation that returns rows for which no matching row exists on the
other side.

`NOT EXISTS` naturally expresses this.

------------------------------------------------------------------------

# 27. Quick Cheat Sheet

  -----------------------------------------------------------------------
  Requirement                         Recommended pattern
  ----------------------------------- -----------------------------------
  Compare with one calculated value   Scalar subquery

  Compare with one row                Single-row subquery

  Match against multiple values       `IN`

  Check whether any match exists      `EXISTS`

  Check that no match exists          `NOT EXISTS`

  Compare against every value         `ALL`

  Compare against at least one value  `ANY` / `SOME`

  Calculate per-group value for each  Window function or grouped relation
  row                                 

  Build intermediate table-like       Derived table / CTE
  result                              

  Need columns from related table     `JOIN`

  Only need relationship existence    `EXISTS`

  Find missing relationships          `NOT EXISTS`

  Avoid `NOT IN` NULL issues          `NOT EXISTS`
  -----------------------------------------------------------------------

------------------------------------------------------------------------

# 28. Final Mental Model

The easiest way to master subqueries is to stop thinking of them as one
giant category.

Think about **what the inner query produces**.

------------------------------------------------------------------------

## 28.1 Produces One Value

Use a scalar subquery:

``` sql
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
)
```

Mental model:

``` text
SUBQUERY → VALUE
```

------------------------------------------------------------------------

## 28.2 Produces a Set of Values

Use `IN`:

``` sql
WHERE department_id IN (
    SELECT department_id
    FROM departments
)
```

Mental model:

``` text
SUBQUERY → SET
```

------------------------------------------------------------------------

## 28.3 Answers Whether a Row Exists

Use `EXISTS`:

``` sql
WHERE EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
)
```

Mental model:

``` text
SUBQUERY → BOOLEAN
```

------------------------------------------------------------------------

## 28.4 Answers Whether a Row Does Not Exist

Use `NOT EXISTS`:

``` sql
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
)
```

Mental model:

``` text
SUBQUERY → ABSENCE
```

------------------------------------------------------------------------

## 28.5 Produces a Table-Like Result

Use a derived table:

``` sql
FROM (
    SELECT ...
) AS x
```

Mental model:

``` text
SUBQUERY → RELATION
```

------------------------------------------------------------------------

## 28.6 Depends on the Current Outer Row

Use a correlated subquery:

``` sql
WHERE e.salary > (
    SELECT AVG(e2.salary)
    FROM employees e2
    WHERE e2.department_id = e.department_id
)
```

Mental model:

``` text
OUTER ROW
    ↓
SUBQUERY
    ↓
ROW-SPECIFIC RESULT
```

------------------------------------------------------------------------

# The Ultimate Subquery Decision Framework

When you see a problem, ask these questions in order:

``` text
1. Do I need a single calculated value?
        ↓
      Scalar subquery

2. Do I need to test membership in a set?
        ↓
      IN

3. Do I only care whether a matching row exists?
        ↓
      EXISTS

4. Do I care that no matching row exists?
        ↓
      NOT EXISTS

5. Does the inner query depend on the current outer row?
        ↓
      Correlated subquery

6. Do I need the inner result as a table-like relation?
        ↓
      Derived table / CTE

7. Do I need columns from both relations?
        ↓
      JOIN

8. Am I doing per-group calculations?
        ↓
      Consider GROUP BY or window functions

9. Could NULL affect NOT IN?
        ↓
      Prefer NOT EXISTS or explicitly handle NULL

10. Is performance important?
        ↓
      Inspect EXPLAIN / EXPLAIN ANALYZE
```

------------------------------------------------------------------------

# One-Page Summary

``` text
SUBQUERY
│
├── Scalar
│   └── One value
│       └── salary > (SELECT AVG(...))
│
├── Single-row
│   └── One row
│       └── =, >, <, >=, <=
│
├── Multi-row
│   ├── IN
│   ├── ANY / SOME
│   └── ALL
│
├── Correlated
│   └── References outer query
│
├── Nested
│   └── Subquery inside another subquery
│
├── EXISTS
│   └── At least one matching row
│       └── Semi-join concept
│
├── NOT EXISTS
│   └── No matching row
│       └── Anti-join concept
│
└── Derived Table
    └── Subquery in FROM
        └── Table-like intermediate result
```

------------------------------------------------------------------------

# Final Takeaway

A subquery is not merely "a query inside a query."

The important question is:

> **What role is the inner query playing?**

It may be producing:

``` text
VALUE
SET
BOOLEAN
ABSENCE
RELATION
```

Once that becomes intuitive, the syntax becomes much easier to choose.

The most important practical distinctions are:

``` text
Scalar       → one value
IN           → membership
EXISTS       → existence
NOT EXISTS   → non-existence
Correlated   → depends on outer row
Derived      → table-like intermediate result
JOIN          → combine rows/columns
Window        → calculate across related rows without collapsing them
```

And the most important safety rule is:

``` text
NOT IN + NULL
        ↓
      be careful
        ↓
NOT EXISTS
```

The most important performance rule is:

``` text
Do not assume:
subquery = slow
JOIN     = fast

Instead:

SQL
 ↓
Optimizer
 ↓
Execution Plan
 ↓
Actual Runtime
```

That is the level at which SQL performance should be reasoned about.
