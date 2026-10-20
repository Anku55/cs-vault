# SQL Set Operations & Common Table Expressions --- Deep Practical & Theoretical Notes

> Sections 9 and 10 of the SQL notes series. Covers set operations,
> duplicate handling, column/type rules, `NULL` behavior, CTEs,
> recursive CTEs, execution concepts, performance, and interview-level
> comparisons.

------------------------------------------------------------------------

# 9. Set Operations

Set operations allow SQL queries to combine the results of multiple
`SELECT` statements.

The core set operators are:

``` text
UNION
UNION ALL
INTERSECT
EXCEPT
```

Instead of joining columns from different tables, set operations combine
**result rows** vertically.

------------------------------------------------------------------------

## 9.1 The Fundamental Difference: JOIN vs Set Operation

A `JOIN` combines columns from related rows:

``` text
Table A             Table B
   │                   │
   └────── JOIN ───────┘
              ↓
      wider result
```

A set operation combines rows from separate query results:

``` text
Query A
   ↓
 rows
   +
Query B
   ↓
 rows
   =
combined result
```

Think:

``` text
JOIN       → combine columns horizontally
UNION      → combine rows vertically
```

Example:

``` sql
SELECT employee_id, name
FROM employees

UNION

SELECT contractor_id, name
FROM contractors;
```

Both queries produce the same two-column shape, so their rows can be
combined.

------------------------------------------------------------------------

# 9.2 `UNION`

`UNION` combines the results of two or more queries and removes
duplicate rows.

Basic syntax:

``` sql
SELECT column1, column2
FROM table1

UNION

SELECT column1, column2
FROM table2;
```

Example:

``` sql
SELECT email
FROM customers

UNION

SELECT email
FROM employees;
```

If:

``` text
customers:
alice@example.com
bob@example.com
carol@example.com

employees:
bob@example.com
david@example.com
```

Then:

``` sql
SELECT email
FROM customers

UNION

SELECT email
FROM employees;
```

produces:

``` text
alice@example.com
bob@example.com
carol@example.com
david@example.com
```

`bob@example.com` appears only once.

------------------------------------------------------------------------

## 9.2.1 `UNION` Means Distinct Combined Results

Conceptually:

``` text
Result A
   +
Result B
   ↓
combined rows
   ↓
remove duplicates
   ↓
final result
```

This is why `UNION` may require additional work compared with
`UNION ALL`.

------------------------------------------------------------------------

## 9.2.2 Example with Complete Rows

Suppose:

``` text
employees
--------------------------------
id   name
1    Alice
2    Bob

contractors
--------------------------------
id   name
2    Bob
3    Carol
```

Query:

``` sql
SELECT id, name
FROM employees

UNION

SELECT id, name
FROM contractors;
```

Result:

``` text
1  Alice
2  Bob
3  Carol
```

The duplicate row:

``` text
2  Bob
```

is removed.

------------------------------------------------------------------------

# 9.3 `UNION ALL`

`UNION ALL` combines result sets **without removing duplicates**.

Example:

``` sql
SELECT email
FROM customers

UNION ALL

SELECT email
FROM employees;
```

Using the earlier data:

``` text
customers:
alice@example.com
bob@example.com
carol@example.com

employees:
bob@example.com
david@example.com
```

Result:

``` text
alice@example.com
bob@example.com
carol@example.com
bob@example.com
david@example.com
```

`bob@example.com` appears twice.

------------------------------------------------------------------------

## 9.3.1 Why `UNION ALL` Is Often Faster

`UNION ALL` does not need to perform duplicate elimination.

Conceptually:

``` text
UNION ALL:

A ──┐
    ├── concatenate ──→ result
B ──┘
```

While:

``` text
UNION:

A ──┐
    ├── concatenate ──→ duplicate elimination ──→ result
B ──┘
```

Duplicate elimination may involve operations such as:

-   sorting
-   hashing
-   memory allocation
-   temporary/intermediate structures

The exact physical implementation is database-specific.

Therefore:

> If duplicates are valid and should be preserved, prefer `UNION ALL`.

Do not use `UNION` merely out of habit.

------------------------------------------------------------------------

# 9.4 `INTERSECT`

`INTERSECT` returns rows that appear in **both** result sets.

Example:

``` sql
SELECT email
FROM customers

INTERSECT

SELECT email
FROM employees;
```

If:

``` text
customers:
alice@example.com
bob@example.com
carol@example.com

employees:
bob@example.com
david@example.com
```

Result:

``` text
bob@example.com
```

The logical meaning is:

``` text
A ∩ B
```

where `∩` represents intersection.

------------------------------------------------------------------------

## 9.4.1 `INTERSECT` Is About Common Rows

Think:

``` text
Query A
   ↓
{1, 2, 3}

Query B
   ↓
{2, 3, 4}

INTERSECT
   ↓
{2, 3}
```

Only rows appearing in both inputs survive.

------------------------------------------------------------------------

## 9.4.2 `INTERSECT` Usually Returns Distinct Rows

Set operations such as `UNION`, `INTERSECT`, and `EXCEPT` are
traditionally distinct-producing operations unless a dialect provides
explicit alternatives such as `INTERSECT ALL`.

Therefore:

``` sql
SELECT 1
UNION
SELECT 1;
```

returns one row:

``` text
1
```

Similarly:

``` sql
SELECT 1
INTERSECT
SELECT 1;
```

returns:

``` text
1
```

not two copies.

------------------------------------------------------------------------

# 9.5 `EXCEPT`

`EXCEPT` returns rows from the first query that do not appear in the
second query.

Example:

``` sql
SELECT email
FROM customers

EXCEPT

SELECT email
FROM employees;
```

If:

``` text
customers:
alice@example.com
bob@example.com
carol@example.com

employees:
bob@example.com
david@example.com
```

Result:

``` text
alice@example.com
carol@example.com
```

The mathematical idea is:

``` text
A − B
```

Meaning:

> Keep rows in A that are not in B.

------------------------------------------------------------------------

## 9.5.1 Order Matters

These are not generally equivalent:

``` sql
A EXCEPT B
```

and:

``` sql
B EXCEPT A
```

Example:

``` text
A = {1, 2, 3}
B = {2, 3, 4}
```

Then:

``` text
A EXCEPT B = {1}
B EXCEPT A = {4}
```

Therefore:

> `EXCEPT` is directional.

------------------------------------------------------------------------

# 9.6 `UNION` vs `UNION ALL`

This distinction is fundamental.

  Feature                `UNION`                 `UNION ALL`
  ---------------------- ----------------------- ----------------------------
  Combines rows          Yes                     Yes
  Removes duplicates     Yes                     No
  Preserves duplicates   No                      Yes
  Usually more work      Yes                     Usually no
  Typical use            Distinct combined set   Append/concatenate results

Example:

``` sql
SELECT id FROM A
UNION
SELECT id FROM B;
```

versus:

``` sql
SELECT id FROM A
UNION ALL
SELECT id FROM B;
```

------------------------------------------------------------------------

# 9.7 Set Operation Rules

Set operations have important compatibility rules.

The queries being combined must produce compatible result shapes.

------------------------------------------------------------------------

## 9.7.1 Same Number of Columns

This is valid:

``` sql
SELECT id, name
FROM employees

UNION

SELECT id, name
FROM contractors;
```

Both return:

``` text
2 columns
```

This is invalid:

``` sql
SELECT id, name
FROM employees

UNION

SELECT id
FROM contractors;
```

The first query returns two columns.

The second returns one.

The result shape cannot be combined.

------------------------------------------------------------------------

## 9.7.2 Corresponding Columns Must Be Type-Compatible

Example:

``` sql
SELECT employee_id
FROM employees

UNION

SELECT contractor_id
FROM contractors;
```

If both IDs are compatible integer types, this is straightforward.

But problematic combinations can occur:

``` sql
SELECT employee_id
FROM employees

UNION

SELECT employee_name
FROM contractors;
```

Whether a database accepts such a combination depends on its
type-conversion rules.

The key principle:

> Corresponding columns must have compatible data types.

Do not rely on implicit conversions when correctness matters.

------------------------------------------------------------------------

## 9.7.3 Column Position Matters

Consider:

``` sql
SELECT employee_id, name
FROM employees

UNION

SELECT name, employee_id
FROM contractors;
```

The database does not infer that you intended:

``` text
id ↔ id
name ↔ name
```

The set operation combines:

``` text
column 1 with column 1
column 2 with column 2
```

Therefore, the second query should normally be:

``` sql
SELECT employee_id, name
FROM contractors;
```

------------------------------------------------------------------------

# 9.8 Column Names in Set Operations

The output column names are generally determined by the first `SELECT`.

Example:

``` sql
SELECT employee_id AS id, name AS person_name
FROM employees

UNION

SELECT contractor_id AS contractor_id, name AS contractor_name
FROM contractors;
```

The final result generally uses the first query's output names:

``` text
id
person_name
```

The aliases from later branches do not normally rename the final
columns.

This is another reason to make the first query's output schema
intentional.

------------------------------------------------------------------------

# 9.9 `ORDER BY` with Set Operations

When ordering the final combined result, the `ORDER BY` normally belongs
at the end:

``` sql
SELECT id, name
FROM employees

UNION

SELECT id, name
FROM contractors

ORDER BY name;
```

The intent is:

``` text
query A
   +
query B
   ↓
combined result
   ↓
ORDER BY
```

An `ORDER BY` inside an individual branch generally requires special
syntax or a derived table depending on the database.

------------------------------------------------------------------------

## 9.9.1 `ORDER BY` Does Not Define Set Membership

Ordering is presentation behavior.

It does not change which rows belong to the set.

Think:

``` text
Set operation → determines rows
ORDER BY      → determines display order
```

------------------------------------------------------------------------

# 9.10 `LIMIT` / `TOP` with Set Operations

Be careful with row-limiting clauses.

For example:

``` sql
SELECT id
FROM employees

UNION ALL

SELECT id
FROM contractors

LIMIT 10;
```

typically means:

> Combine both result sets, then limit the final result to 10 rows.

If you want each branch independently limited, use derived tables or
dialect-specific syntax.

For example:

``` sql
SELECT id
FROM (
    SELECT id
    FROM employees
    ORDER BY id
    LIMIT 10
) e

UNION ALL

SELECT id
FROM (
    SELECT id
    FROM contractors
    ORDER BY id
    LIMIT 10
) c;
```

Exact syntax varies by SQL dialect.

------------------------------------------------------------------------

# 9.11 Duplicate Handling

Duplicate behavior is one of the most important topics in set
operations.

------------------------------------------------------------------------

## 9.11.1 `UNION`

Duplicates removed:

``` sql
SELECT 1
UNION
SELECT 1;
```

Result:

``` text
1
```

------------------------------------------------------------------------

## 9.11.2 `UNION ALL`

Duplicates preserved:

``` sql
SELECT 1
UNION ALL
SELECT 1;
```

Result:

``` text
1
1
```

------------------------------------------------------------------------

## 9.11.3 Duplicate Rows Are Compared as Rows

Suppose:

``` sql
SELECT id, name
FROM A

UNION

SELECT id, name
FROM B;
```

These two rows are duplicates:

``` text
1 | Alice
1 | Alice
```

But these are different:

``` text
1 | Alice
1 | Bob
```

because the complete row values differ.

------------------------------------------------------------------------

## 9.11.4 Duplicate Detection Uses the Complete Output Row

Consider:

``` text
A:
1 | Alice | Delhi

B:
1 | Alice | Mumbai
```

Even though:

``` text
id = 1
name = Alice
```

are the same, the complete rows differ because the third column differs.

Therefore `UNION` keeps both:

``` text
1 | Alice | Delhi
1 | Alice | Mumbai
```

------------------------------------------------------------------------

# 9.12 `NULL` and Set Operations

`NULL` deserves special attention.

Consider:

``` sql
SELECT NULL
UNION
SELECT NULL;
```

Set-style duplicate elimination treats the two `NULL` values as
belonging to the same duplicate row for purposes of the set operation.

The result is effectively one row containing `NULL`.

This is different from ordinary SQL comparison semantics, where:

``` sql
NULL = NULL
```

does not evaluate to `TRUE`.

This distinction is important:

> Set duplicate semantics are not identical to ordinary `=` predicate
> semantics.

------------------------------------------------------------------------

# 9.13 `INTERSECT` and `NULL`

Similarly, if both result sets contain corresponding `NULL` values,
set-operation semantics can consider those rows matching for
intersection purposes.

Conceptually:

``` text
A = {NULL, 1}
B = {NULL, 2}

A INTERSECT B
```

produces:

``` text
NULL
```

plus no `1` or `2`.

Exact details can vary in edge cases by SQL dialect, but standard
set-operation semantics treat corresponding nulls as matching for
duplicate elimination and set comparison.

------------------------------------------------------------------------

# 9.14 `EXCEPT` and `NULL`

Consider:

``` text
A = {NULL, 1}
B = {NULL}
```

Then:

``` text
A EXCEPT B
```

returns:

``` text
1
```

The `NULL` row is considered present in both sets and is therefore
removed.

This differs significantly from naïve thinking based on:

``` sql
NULL = NULL
```

------------------------------------------------------------------------

# 9.15 `UNION` vs `OR`

These can sometimes solve similar-looking problems, but they are not the
same operation.

Example:

``` sql
SELECT *
FROM employees
WHERE department_id = 10
   OR department_id = 20;
```

could be conceptually expressed as:

``` sql
SELECT *
FROM employees
WHERE department_id = 10

UNION

SELECT *
FROM employees
WHERE department_id = 20;
```

Because `UNION` removes duplicates, the results may match.

But `UNION ALL` would preserve duplicates if the branches overlap.

For simple predicates, `OR` is generally clearer.

------------------------------------------------------------------------

# 9.16 `UNION` vs `JOIN`

Suppose:

``` text
employees
id | name

contractors
id | name
```

If you want one combined list:

``` sql
SELECT id, name
FROM employees

UNION ALL

SELECT id, name
FROM contractors;
```

You want a vertical combination.

If you want:

``` text
employee
department
manager
```

from related tables, use a `JOIN`.

Mental model:

``` text
UNION:
A
B
C
↓
A + B + C rows

JOIN:
A columns + B columns
```

------------------------------------------------------------------------

# 9.17 `INTERSECT` vs `INNER JOIN`

Suppose:

``` sql
SELECT customer_id
FROM customers_2025

INTERSECT

SELECT customer_id
FROM customers_2026;
```

This asks:

> Which customer IDs occur in both result sets?

A join could also express the relationship:

``` sql
SELECT DISTINCT a.customer_id
FROM customers_2025 a
JOIN customers_2026 b
    ON b.customer_id = a.customer_id;
```

But the semantics differ.

`INTERSECT` expresses:

> Common rows between two query results.

The `JOIN` expresses:

> Match rows from two relations.

The query optimizer may produce similar physical strategies in some
cases, but the SQL communicates different intent.

------------------------------------------------------------------------

# 9.18 `EXCEPT` vs `NOT EXISTS`

These can often express similar anti-set logic.

Using `EXCEPT`:

``` sql
SELECT customer_id
FROM customers

EXCEPT

SELECT customer_id
FROM orders;
```

Meaning:

> Customer IDs that appear in customers but not in orders.

Using `NOT EXISTS`:

``` sql
SELECT c.customer_id
FROM customers c
WHERE NOT EXISTS (
    SELECT 1
    FROM orders o
    WHERE o.customer_id = c.customer_id
);
```

The second form is often more convenient when you need additional
columns from `customers`.

The first is elegant when the problem is naturally:

``` text
set A minus set B
```

------------------------------------------------------------------------

# 9.19 Set Operation Decision Guide

Use:

``` text
UNION
```

when:

> Combine results and remove duplicates.

Use:

``` text
UNION ALL
```

when:

> Combine results and preserve duplicates.

Use:

``` text
INTERSECT
```

when:

> Return common rows.

Use:

``` text
EXCEPT
```

when:

> Return rows in the first result but not the second.

Use:

``` text
JOIN
```

when:

> Combine related columns from different relations.

------------------------------------------------------------------------

# 9.20 Set Operations --- Interview Checklist

Remember these rules:

``` text
1. Same number of columns
2. Compatible data types
3. Position matters
4. First SELECT usually determines output names
5. UNION removes duplicates
6. UNION ALL preserves duplicates
7. INTERSECT returns common rows
8. EXCEPT returns left-only rows
9. EXCEPT is directional
10. Final ORDER BY normally belongs at the end
11. Set operations work on complete result rows
12. NULL has special set-comparison semantics
```

------------------------------------------------------------------------

# 10. Common Table Expressions

A **Common Table Expression**, or **CTE**, is a named query expression
introduced with:

``` sql
WITH
```

A CTE allows you to define an intermediate result and then reference it
from the main query.

Basic syntax:

``` sql
WITH cte_name AS (
    SELECT ...
)
SELECT ...
FROM cte_name;
```

------------------------------------------------------------------------

# 10.1 Why CTEs Exist

Without a CTE, a complex query can become deeply nested:

``` sql
SELECT ...
FROM (
    SELECT ...
    FROM (
        SELECT ...
        FROM ...
    ) x
) y;
```

A CTE can turn that into:

``` sql
WITH first_step AS (
    SELECT ...
),
second_step AS (
    SELECT ...
    FROM first_step
)
SELECT ...
FROM second_step;
```

The query becomes easier to reason about.

------------------------------------------------------------------------

# 10.2 Basic CTE Example

Suppose we want department-level average salaries.

Without CTE:

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

With CTE:

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

The CTE gives the intermediate result a name.

------------------------------------------------------------------------

# 10.3 `WITH`

The `WITH` clause introduces one or more CTEs.

Basic form:

``` sql
WITH cte_name AS (
    SELECT ...
)
SELECT ...
FROM cte_name;
```

Multiple CTEs:

``` sql
WITH
cte1 AS (
    SELECT ...
),
cte2 AS (
    SELECT ...
    FROM cte1
)
SELECT ...
FROM cte2;
```

------------------------------------------------------------------------

# 10.4 CTE Is Not Automatically a Physical Temporary Table

This is a crucial distinction.

A CTE is primarily a **query expression**, not necessarily a stored
table.

Writing:

``` sql
WITH x AS (
    SELECT *
    FROM employees
)
SELECT *
FROM x;
```

does not inherently mean:

``` text
create physical temporary table x
insert rows
query x
drop x
```

The optimizer may instead inline or transform the CTE.

Physical behavior depends on the database engine and query.

------------------------------------------------------------------------

# 10.5 CTE Example: Employees Above Department Average

``` sql
WITH department_stats AS (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
)
SELECT
    e.employee_id,
    e.name,
    e.salary,
    d.avg_salary
FROM employees e
JOIN department_stats d
    ON d.department_id = e.department_id
WHERE e.salary > d.avg_salary;
```

Logical stages:

``` text
employees
    ↓
GROUP BY department
    ↓
department_stats
    ↓
JOIN back to employees
    ↓
filter employees above department average
```

------------------------------------------------------------------------

# 10.6 Multiple CTEs

Multiple CTEs are separated by commas.

Example:

``` sql
WITH
department_stats AS (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
),
high_salary_departments AS (
    SELECT department_id
    FROM department_stats
    WHERE avg_salary > 80000
)
SELECT e.*
FROM employees e
JOIN high_salary_departments h
    ON h.department_id = e.department_id;
```

The dependency graph is:

``` text
employees
    ↓
department_stats
    ↓
high_salary_departments
    ↓
final query
```

------------------------------------------------------------------------

# 10.7 Multiple CTEs as a Data Pipeline

A useful way to think about CTEs:

``` text
raw data
   ↓
CTE 1: filter
   ↓
CTE 2: aggregate
   ↓
CTE 3: enrich
   ↓
CTE 4: rank
   ↓
final result
```

Example:

``` sql
WITH filtered_orders AS (
    SELECT *
    FROM orders
    WHERE order_date >= DATE '2026-01-01'
),
customer_totals AS (
    SELECT
        customer_id,
        SUM(amount) AS total_spend
    FROM filtered_orders
    GROUP BY customer_id
),
ranked_customers AS (
    SELECT
        customer_id,
        total_spend,
        RANK() OVER (
            ORDER BY total_spend DESC
        ) AS spending_rank
    FROM customer_totals
)
SELECT *
FROM ranked_customers
WHERE spending_rank <= 10;
```

This is much easier to read than deeply nested derived tables.

------------------------------------------------------------------------

# 10.8 CTE Column Naming

You can explicitly name CTE columns in some SQL dialects:

``` sql
WITH department_stats (department_id, avg_salary) AS (
    SELECT
        department_id,
        AVG(salary)
    FROM employees
    GROUP BY department_id
)
SELECT *
FROM department_stats;
```

Alternatively, aliases inside the query are usually clearer:

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

------------------------------------------------------------------------

# 10.9 CTE Scope

A normal CTE exists only for the statement immediately following the
`WITH` clause.

Example:

``` sql
WITH x AS (
    SELECT 1 AS value
)
SELECT *
FROM x;
```

After this statement, `x` is not a permanent database object.

You cannot generally run:

``` sql
SELECT *
FROM x;
```

in a separate statement and expect the CTE to still exist.

------------------------------------------------------------------------

# 10.10 CTE vs Subquery

A CTE and a subquery can often express the same logical transformation.

------------------------------------------------------------------------

## 10.10.1 Subquery Version

``` sql
SELECT *
FROM (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
) d
WHERE d.avg_salary > 80000;
```

------------------------------------------------------------------------

## 10.10.2 CTE Version

``` sql
WITH department_stats AS (
    SELECT
        department_id,
        AVG(salary) AS avg_salary
    FROM employees
    GROUP BY department_id
)
SELECT *
FROM department_stats
WHERE avg_salary > 80000;
```

The CTE version is often easier to read.

------------------------------------------------------------------------

## 10.10.3 When CTEs Are Better for Readability

Use a CTE when:

-   the intermediate logic is substantial
-   there are multiple stages
-   the query needs descriptive names
-   the same logical intermediate result is referenced multiple times
-   recursion is required
-   you want to separate transformation stages

Use a simple subquery when:

-   the intermediate query is short
-   it is used only once
-   nesting remains readable
-   adding a CTE would make a tiny query unnecessarily verbose

------------------------------------------------------------------------

# 10.11 CTE vs Temporary Table

These are fundamentally different concepts.

### CTE

``` sql
WITH x AS (
    SELECT ...
)
SELECT ...
FROM x;
```

Usually:

-   statement-scoped
-   logical query expression
-   not inherently persistent
-   optimizer may inline or materialize depending on engine

### Temporary table

Example syntax varies by database:

``` sql
CREATE TEMPORARY TABLE temp_x AS
SELECT ...;
```

Usually:

-   physically stores rows in a temporary relation
-   can be referenced by multiple statements
-   has a lifecycle controlled by the database/session/transaction
-   can sometimes be indexed
-   can be useful for multi-step workflows

------------------------------------------------------------------------

# 10.12 CTE vs Temporary Table --- Comparison

  -----------------------------------------------------------------------------
  Feature                 CTE                     Temporary Table
  ----------------------- ----------------------- -----------------------------
  Main purpose            Query organization      Materialized intermediate
                                                  data

  Typical scope           One statement           Session/transaction/defined
                                                  lifecycle

  Persistent object       No                      Temporary database object

  Can index directly?     Usually no              Often yes

  Can be reused across    No                      Yes
  statements?                                     

  Can be recursive?       Yes                     Not inherently

  Physical                No                      Yes, as a stored temporary
  materialization                                 relation
  guaranteed?                                     

  Good for readability    Excellent               Less direct

  Good for large          Sometimes               Often useful
  multi-step workflows                            
  -----------------------------------------------------------------------------

Exact lifecycle and behavior vary by SQL dialect.

------------------------------------------------------------------------

# 10.13 Recursive CTE

A **recursive CTE** is a CTE that references itself.

It is used for hierarchical or graph-like data.

Common examples:

-   organizational hierarchies
-   category trees
-   folder structures
-   bill of materials
-   dependency graphs
-   parent-child relationships
-   graph traversal

------------------------------------------------------------------------

# 10.14 Recursive CTE Structure

A recursive CTE usually contains two logical parts:

``` text
Anchor member
    +
Recursive member
```

Generic syntax:

``` sql
WITH RECURSIVE cte_name AS (
    -- anchor
    SELECT ...

    UNION ALL

    -- recursive member
    SELECT ...
    FROM cte_name
    JOIN ...
)
SELECT *
FROM cte_name;
```

Some database systems use `WITH RECURSIVE`; others have dialect-specific
syntax.

------------------------------------------------------------------------

# 10.15 Anchor Member

The anchor query establishes the starting rows.

Example:

``` sql
SELECT
    employee_id,
    name,
    manager_id,
    0 AS level
FROM employees
WHERE manager_id IS NULL
```

This finds top-level employees.

------------------------------------------------------------------------

# 10.16 Recursive Member

The recursive part finds the next level.

``` sql
SELECT
    e.employee_id,
    e.name,
    e.manager_id,
    org.level + 1
FROM employees e
JOIN org
    ON e.manager_id = org.employee_id
```

The recursive CTE feeds its own previous results into the next
iteration.

------------------------------------------------------------------------

# 10.17 Full Recursive Hierarchy Example

Suppose:

``` text
employee_id | name   | manager_id
------------|--------|-----------
1           | Alice  | NULL
2           | Bob    | 1
3           | Carol  | 1
4           | David  | 2
5           | Eve    | 2
6           | Frank  | 4
```

Hierarchy:

``` text
Alice
├── Bob
│   ├── David
│   │   └── Frank
│   └── Eve
└── Carol
```

Recursive CTE:

``` sql
WITH RECURSIVE org AS (
    -- Anchor
    SELECT
        employee_id,
        name,
        manager_id,
        0 AS level
    FROM employees
    WHERE manager_id IS NULL

    UNION ALL

    -- Recursive member
    SELECT
        e.employee_id,
        e.name,
        e.manager_id,
        org.level + 1
    FROM employees e
    JOIN org
        ON e.manager_id = org.employee_id
)
SELECT *
FROM org
ORDER BY level, employee_id;
```

------------------------------------------------------------------------

# 10.18 How Recursive CTEs Work Conceptually

Think in iterations.

Initial anchor:

``` text
Iteration 0:

Alice
```

Recursive step:

``` text
Iteration 1:

Bob
Carol
```

Next:

``` text
Iteration 2:

David
Eve
```

Next:

``` text
Iteration 3:

Frank
```

Next:

``` text
Iteration 4:

nothing
```

When no new rows are produced, recursion terminates.

Conceptually:

``` text
Anchor
  ↓
Level 1
  ↓
Level 2
  ↓
Level 3
  ↓
...
  ↓
No new rows
  ↓
Stop
```

The exact duplicate handling and recursion semantics can vary by
database and query structure.

------------------------------------------------------------------------

# 10.19 `UNION ALL` in Recursive CTEs

Recursive CTEs commonly use:

``` sql
UNION ALL
```

between the anchor and recursive member:

``` sql
WITH RECURSIVE org AS (
    ...
    UNION ALL
    ...
)
```

Why?

Because recursion is intentionally generating rows across levels.

Using duplicate-eliminating set semantics can change the behavior and
cost.

However, recursion over graphs introduces a major issue:

> Cycles.

------------------------------------------------------------------------

# 10.20 Recursive CTE and Cycles

Consider:

``` text
A → B
B → C
C → A
```

This creates a cycle.

A naïve recursive traversal can continue indefinitely or until the
database's recursion limit is reached.

For graph traversal, you may need:

-   cycle detection
-   visited-node tracking
-   depth limits
-   database-specific cycle clauses/features

A simple defensive technique is to track a path or visited set.

For example, conceptually:

``` text
current node
+
visited nodes
```

and stop recursion when the next node is already visited.

The exact implementation is database-specific.

------------------------------------------------------------------------

# 10.21 Recursive CTE for Numbers

Recursive CTEs can generate sequences.

Example:

``` sql
WITH RECURSIVE numbers AS (
    SELECT 1 AS n

    UNION ALL

    SELECT n + 1
    FROM numbers
    WHERE n < 10
)
SELECT *
FROM numbers;
```

Result:

``` text
1
2
3
4
5
6
7
8
9
10
```

This is useful for understanding recursion, although many databases
provide dedicated sequence/generator functions that may be more
appropriate for production workloads.

------------------------------------------------------------------------

# 10.22 Recursive CTE for Date Ranges

A recursive CTE can also generate dates.

Conceptually:

``` sql
WITH RECURSIVE dates AS (
    SELECT DATE '2026-01-01' AS day

    UNION ALL

    SELECT day + INTERVAL '1 day'
    FROM dates
    WHERE day < DATE '2026-01-10'
)
SELECT *
FROM dates;
```

Date arithmetic syntax varies substantially between SQL dialects.

The concept is:

``` text
start date
   ↓
add one interval
   ↓
add one interval
   ↓
...
```

------------------------------------------------------------------------

# 10.23 Recursive CTE for Category Trees

Suppose:

``` text
categories
---------------------------
id | name       | parent_id

1  | Electronics | NULL
2  | Phones      | 1
3  | Laptops     | 1
4  | Android     | 2
5  | iPhone      | 2
```

A recursive CTE can traverse:

``` text
Electronics
├── Phones
│   ├── Android
│   └── iPhone
└── Laptops
```

This is one of the strongest practical use cases for recursive CTEs.

------------------------------------------------------------------------

# 10.24 Recursive CTE Depth

Production recursive queries should consider:

``` text
How deep can the hierarchy become?
```

A depth column is useful:

``` sql
0 AS level
```

then:

``` sql
level + 1
```

You can also impose a limit:

``` sql
WHERE org.level < 20
```

This can prevent runaway recursion, although it is not a substitute for
proper cycle handling.

------------------------------------------------------------------------

# 10.25 CTE Materialization

A CTE can sometimes be materialized by the database, depending on the
engine and query.

Conceptually:

``` text
CTE
 ↓
materialize intermediate result
 ↓
reuse result
```

Other times:

``` text
CTE
 ↓
inline into outer query
 ↓
optimize as one larger expression
```

Some databases also expose explicit materialization controls.

The important principle is:

> A CTE is a logical abstraction; its physical execution strategy is
> optimizer-dependent unless the database/dialect explicitly specifies
> materialization behavior.

------------------------------------------------------------------------

# 10.26 Reusing a CTE

Suppose:

``` sql
WITH expensive_data AS (
    SELECT ...
)
SELECT ...
FROM expensive_data a
JOIN expensive_data b
    ON ...;
```

A common misconception is:

``` text
CTE defined once
=
result definitely computed once
```

Not necessarily.

Depending on the engine, the optimizer may:

-   materialize once
-   inline references
-   optimize each reference independently
-   choose another strategy

If you require a deliberately materialized intermediate dataset for
repeated multi-statement use, a temporary table may be more appropriate.

------------------------------------------------------------------------

# 10.27 CTE Performance

CTEs are not inherently:

``` text
faster
```

or:

``` text
slower
```

than subqueries.

They are primarily a way to express query structure.

Performance depends on:

-   optimizer
-   cardinality
-   indexes
-   predicates
-   joins
-   aggregation
-   materialization decisions
-   statistics
-   database engine

Always inspect:

``` sql
EXPLAIN
```

or:

``` sql
EXPLAIN ANALYZE
```

when performance matters.

------------------------------------------------------------------------

# 10.28 CTE vs Subquery --- Detailed Comparison

  -----------------------------------------------------------------------
  Property                CTE                     Subquery
  ----------------------- ----------------------- -----------------------
  Syntax                  `WITH name AS (...)`    `(SELECT ...)`

  Readability             Often better for        Good for small logic
                          complex logic           

  Naming                  Explicit name           Alias required for
                                                  derived table

  Multiple stages         Excellent               Can become deeply
                                                  nested

  Recursive queries       Yes                     Not in the same
                                                  standard CTE form

  Reuse within statement  Can reference CTE name  Usually repeat the
                          multiple times, subject subquery
                          to dialect rules        

  Scope                   Statement               Expression/query
                                                  location

  Materialization         No                      No
  guaranteed?                                     

  Performance             Optimizer-dependent     Optimizer-dependent
  -----------------------------------------------------------------------

------------------------------------------------------------------------

# 10.29 CTE vs Temporary Table --- Detailed Comparison

Use a **CTE** when:

``` text
I need a logical intermediate result
for one SQL statement.
```

Use a **temporary table** when:

``` text
I need an actual temporary relation
across multiple statements or need
physical indexing/materialization.
```

Example CTE workflow:

``` sql
WITH filtered_orders AS (
    SELECT *
    FROM orders
    WHERE amount > 1000
)
SELECT *
FROM filtered_orders;
```

Temporary-table workflow:

``` sql
CREATE TEMPORARY TABLE filtered_orders AS
SELECT *
FROM orders
WHERE amount > 1000;

CREATE INDEX idx_filtered_orders_customer
ON filtered_orders(customer_id);

SELECT ...
FROM filtered_orders;

SELECT ...
FROM filtered_orders;
```

The second pattern is useful when the intermediate dataset needs to
survive across multiple statements and potentially be indexed.

------------------------------------------------------------------------

# 10.30 CTE vs Temporary Table --- Decision Table

  Requirement                                    CTE           Temporary Table
  ------------------------------------------ ------------ -------------------------
  One statement                                   ✅              Possible
  Multiple statements                             ❌                 ✅
  Readability                                     ✅              Moderate
  Recursive hierarchy                             ✅       Not the primary feature
  Explicit indexing                           Usually ❌             ✅
  Guaranteed physical intermediate storage        ❌                 ✅
  Reuse across statements                         ❌                 ✅
  Complex multi-step ETL-style workflow       Sometimes             Often
  Avoid persistent schema objects                 ✅                 ✅

Exact behavior depends on the database system.

------------------------------------------------------------------------

# 10.31 CTE + Window Function

CTEs pair extremely well with window functions.

Example:

``` sql
WITH ranked_employees AS (
    SELECT
        employee_id,
        name,
        department_id,
        salary,
        RANK() OVER (
            PARTITION BY department_id
            ORDER BY salary DESC
        ) AS salary_rank
    FROM employees
)
SELECT *
FROM ranked_employees
WHERE salary_rank <= 3;
```

The CTE creates a named intermediate result containing rankings.

The outer query filters the window-function result.

This is a very common SQL pattern.

------------------------------------------------------------------------

# 10.32 CTE + Aggregation

Example:

``` sql
WITH department_stats AS (
    SELECT
        department_id,
        COUNT(*) AS employee_count,
        AVG(salary) AS avg_salary,
        MAX(salary) AS max_salary
    FROM employees
    GROUP BY department_id
)
SELECT *
FROM department_stats
WHERE employee_count >= 5
  AND avg_salary > 80000;
```

This cleanly separates:

``` text
aggregation
    ↓
filtering
```

------------------------------------------------------------------------

# 10.33 CTE + Set Operations

CTEs can also make set operations easier to understand.

Example:

``` sql
WITH active_customers AS (
    SELECT customer_id
    FROM customers
    WHERE status = 'active'
),
recent_buyers AS (
    SELECT DISTINCT customer_id
    FROM orders
    WHERE order_date >= DATE '2026-01-01'
)
SELECT customer_id
FROM active_customers

INTERSECT

SELECT customer_id
FROM recent_buyers;
```

The result is:

> Active customers who bought something recently.

------------------------------------------------------------------------

# 10.34 Multiple CTEs + Set Operations

You can build a more sophisticated pipeline:

``` sql
WITH
active_customers AS (
    SELECT customer_id
    FROM customers
    WHERE status = 'active'
),
recent_buyers AS (
    SELECT DISTINCT customer_id
    FROM orders
    WHERE order_date >= DATE '2026-01-01'
),
inactive_buyers AS (
    SELECT customer_id
    FROM recent_buyers

    EXCEPT

    SELECT customer_id
    FROM active_customers
)
SELECT *
FROM inactive_buyers;
```

The query becomes a sequence of named relational transformations.

------------------------------------------------------------------------

# 10.35 Common CTE Mistakes

## Mistake 1: Treating a CTE as a permanent table

Wrong mental model:

``` text
WITH x AS (...)
```

creates a reusable database table.

Correct:

``` text
CTE = statement-scoped query expression
```

------------------------------------------------------------------------

## Mistake 2: Assuming CTE = Materialized Result

Not necessarily.

The optimizer may inline it.

------------------------------------------------------------------------

## Mistake 3: Assuming CTE = Faster

A CTE is not a performance keyword.

It is primarily a query-organization mechanism.

------------------------------------------------------------------------

## Mistake 4: Overusing CTEs

This:

``` sql
WITH a AS (
    SELECT 1 AS x
)
SELECT *
FROM a;
```

may be perfectly valid, but adding a CTE for trivial logic can make
simple queries harder to read.

Use abstraction where it improves clarity.

------------------------------------------------------------------------

## Mistake 5: Creating Too Many CTE Layers

This:

``` text
CTE 1
 ↓
CTE 2
 ↓
CTE 3
 ↓
CTE 4
 ↓
CTE 5
 ↓
CTE 6
```

can become difficult to maintain if every stage is tiny and poorly
named.

Good CTE names should communicate purpose:

``` text
filtered_orders
customer_totals
ranked_customers
```

rather than:

``` text
x
y
z
tmp1
tmp2
```

------------------------------------------------------------------------

# 10.36 Common Recursive CTE Mistakes

## Mistake 1: Missing Termination Condition

Bad:

``` sql
SELECT n + 1
FROM numbers;
```

There is no stopping condition.

Correct pattern:

``` sql
WHERE n < 100;
```

------------------------------------------------------------------------

## Mistake 2: Ignoring Cycles

Hierarchies and graphs can contain:

``` text
A → B → C → A
```

Without cycle detection, recursion can become non-terminating.

------------------------------------------------------------------------

## Mistake 3: Forgetting the Recursive Join Condition

A recursive CTE needs a relationship between the current result and the
next rows.

For hierarchy:

``` sql
ON e.manager_id = org.employee_id
```

Without the correct relationship, you can accidentally generate a huge
Cartesian-like result.

------------------------------------------------------------------------

## Mistake 4: Not Tracking Depth

For complex recursive queries, adding:

``` sql
level
```

is extremely useful for:

-   debugging
-   limiting depth
-   formatting trees
-   analyzing hierarchy depth

------------------------------------------------------------------------

# 10.37 Set Operations + CTE + Subquery Together

These concepts often appear together in real SQL.

Example:

``` sql
WITH active_customers AS (
    SELECT customer_id
    FROM customers
    WHERE status = 'active'
),
recent_buyers AS (
    SELECT DISTINCT customer_id
    FROM orders
    WHERE order_date >= DATE '2026-01-01'
)
SELECT c.*
FROM customers c
WHERE c.customer_id IN (
    SELECT customer_id
    FROM active_customers

    INTERSECT

    SELECT customer_id
    FROM recent_buyers
);
```

Here we have:

``` text
CTE
 ↓
CTE
 ↓
INTERSECT
 ↓
IN subquery
 ↓
outer query
```

This illustrates why learning SQL features independently is useful, but
understanding how they compose is even more important.

------------------------------------------------------------------------

# 10.38 Set Operations vs JOIN vs Subquery vs CTE

A useful comparison:

  Tool              Primary idea
  ----------------- -------------------------------------------
  `JOIN`            Combine related rows/columns
  `UNION`           Combine result sets and remove duplicates
  `UNION ALL`       Append result sets
  `INTERSECT`       Common result rows
  `EXCEPT`          Rows in A but not B
  Scalar subquery   Produce one value
  `IN` subquery     Membership
  `EXISTS`          Existence
  `NOT EXISTS`      Non-existence
  CTE               Name a query expression
  Recursive CTE     Traverse hierarchical/recursive data
  Temporary table   Materialize temporary relational data

------------------------------------------------------------------------

# 10.39 Practical Decision Framework

When writing SQL, ask:

``` text
Do I need columns from another table?
        ↓
      JOIN

Do I need to append rows from another query?
        ↓
      UNION ALL

Do I need appended rows without duplicates?
        ↓
      UNION

Do I need rows common to both queries?
        ↓
      INTERSECT

Do I need rows in A but not B?
        ↓
      EXCEPT

Do I need one calculated value?
        ↓
      Scalar subquery

Do I need membership?
        ↓
      IN

Do I need existence?
        ↓
      EXISTS

Do I need non-existence?
        ↓
      NOT EXISTS

Do I need a named intermediate query?
        ↓
      CTE

Do I need recursion/hierarchy traversal?
        ↓
      Recursive CTE

Do I need an intermediate dataset across multiple statements?
        ↓
      Temporary table
```

------------------------------------------------------------------------

# 10.40 Interview-Level Mental Model

## Set Operations

Think:

``` text
UNION
    A ∪ B
    combine + deduplicate

UNION ALL
    A + B
    combine + preserve duplicates

INTERSECT
    A ∩ B
    common rows

EXCEPT
    A − B
    left-only rows
```

------------------------------------------------------------------------

## CTEs

Think:

``` text
WITH
    ↓
name an intermediate query
    ↓
use that query in the main statement
```

------------------------------------------------------------------------

## Recursive CTE

Think:

``` text
Anchor
   ↓
Recursive step
   ↓
Recursive step
   ↓
Recursive step
   ↓
termination
```

------------------------------------------------------------------------

# 10.41 Interview Questions

## Q1. What is the difference between `UNION` and `UNION ALL`?

`UNION` removes duplicate result rows.

`UNION ALL` preserves duplicates.

`UNION ALL` is generally preferable when duplicate elimination is
unnecessary.

------------------------------------------------------------------------

## Q2. What are the rules for set operations?

The participating queries must return:

-   compatible numbers of columns
-   compatible data types
-   corresponding columns in matching positions

------------------------------------------------------------------------

## Q3. What does `INTERSECT` do?

Returns rows common to both query results.

------------------------------------------------------------------------

## Q4. What does `EXCEPT` do?

Returns rows from the first query that are absent from the second query.

------------------------------------------------------------------------

## Q5. Is `EXCEPT` symmetric?

No.

``` text
A EXCEPT B
```

is generally different from:

``` text
B EXCEPT A
```

------------------------------------------------------------------------

## Q6. What is a CTE?

A named query expression introduced with `WITH`, normally scoped to a
single SQL statement.

------------------------------------------------------------------------

## Q7. Is a CTE a temporary table?

No.

A CTE is a logical query expression. A temporary table is an actual
temporary relation that can be materialized and usually reused across
multiple statements.

------------------------------------------------------------------------

## Q8. Is a CTE always materialized?

No.

Physical behavior is database- and optimizer-dependent.

------------------------------------------------------------------------

## Q9. Is a CTE always faster than a subquery?

No.

A CTE mainly improves organization and readability. Performance depends
on the optimizer and execution plan.

------------------------------------------------------------------------

## Q10. What is a recursive CTE?

A CTE whose recursive member references the CTE itself, allowing
hierarchical or recursive data to be traversed.

------------------------------------------------------------------------

## Q11. What are the two components of a recursive CTE?

Usually:

``` text
Anchor member
+
Recursive member
```

------------------------------------------------------------------------

## Q12. Why can recursive CTEs become dangerous?

Because recursive data can contain cycles or unexpectedly deep paths.

Use:

-   termination conditions
-   depth limits where appropriate
-   cycle detection when required

------------------------------------------------------------------------

## Q13. When should I use a temporary table instead of a CTE?

When you need an intermediate dataset across multiple statements, need
to explicitly materialize it, or need capabilities such as indexes that
a CTE does not provide directly.

------------------------------------------------------------------------

# 10.42 One-Page Cheat Sheet

``` text
SET OPERATIONS
────────────────────────────────────────

UNION
    combine results
    remove duplicates

UNION ALL
    combine results
    preserve duplicates

INTERSECT
    common rows

EXCEPT
    rows in first result but not second


SET RULES
────────────────────────────────────────

✓ Same number of columns
✓ Compatible data types
✓ Position matters
✓ First query generally determines output names
✓ Final ORDER BY normally applies to final result


CTE
────────────────────────────────────────

WITH name AS (
    SELECT ...
)
SELECT ...
FROM name;


MULTIPLE CTEs
────────────────────────────────────────

WITH
a AS (...),
b AS (...),
c AS (...)
SELECT ...
FROM c;


RECURSIVE CTE
────────────────────────────────────────

WITH RECURSIVE tree AS (
    -- anchor

    SELECT ...

    UNION ALL

    -- recursive member

    SELECT ...
    FROM tree
    JOIN ...
)
SELECT *
FROM tree;


CTE vs SUBQUERY
────────────────────────────────────────

CTE:
    named
    easier for complex multi-step logic
    recursive support
    statement scoped

Subquery:
    inline
    concise
    excellent for small transformations


CTE vs TEMP TABLE
────────────────────────────────────────

CTE:
    logical query expression
    one statement
    no guaranteed materialization

TEMP TABLE:
    actual temporary relation
    multiple statements
    can often be indexed
    explicit materialization
```

------------------------------------------------------------------------

# Final Mental Model

The easiest way to remember these concepts is to classify them by the
problem they solve.

## Set Operations

``` text
Need to combine rows from independent query results?
        ↓
Set operation
```

Then choose:

``` text
UNION       → combine + deduplicate
UNION ALL   → combine + preserve duplicates
INTERSECT   → common rows
EXCEPT      → first-only rows
```

------------------------------------------------------------------------

## CTEs

``` text
Need to give a complex intermediate query a name?
        ↓
CTE
```

Think:

``` text
WITH
    ↓
named intermediate relation
    ↓
main query
```

------------------------------------------------------------------------

## Recursive CTEs

``` text
Need to repeatedly follow a relationship?
        ↓
Recursive CTE
```

Think:

``` text
start
  ↓
children
  ↓
grandchildren
  ↓
next level
  ↓
...
  ↓
stop
```

------------------------------------------------------------------------

## The Bigger SQL Mental Model

At this stage, the major SQL building blocks can be viewed as:

``` text
                    SQL
                     │
       ┌─────────────┼─────────────┐
       │             │             │
    Combine       Filter        Transform
       │             │             │
    JOIN         WHERE        SELECT
    UNION        HAVING       GROUP BY
    INTERSECT   EXISTS        WINDOW
    EXCEPT      IN            CTE
    UNION ALL   SUBQUERY      RECURSION
```

And the key conceptual distinction is:

``` text
JOIN
    → combine related relations horizontally

SET OPERATION
    → combine query results vertically

SUBQUERY
    → use one query's result inside another query

CTE
    → name and organize a query expression

RECURSIVE CTE
    → repeatedly derive rows from previous rows
```

Once these distinctions become intuitive, complex SQL stops looking like
a collection of unrelated keywords and starts looking like a small set
of composable relational operations.
