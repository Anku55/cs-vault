# SQL Window Functions --- Deep Practical & Theoretical Notes

> Section 11 of the SQL notes series. A deep guide to SQL window
> functions, `OVER()`, partitions, ordering, ranking, offsets,
> analytical values, running totals, moving averages, window frames,
> execution concepts, common mistakes, and interview patterns.

------------------------------------------------------------------------

# 11. Window Functions

Window functions are one of the most powerful features in modern SQL.

They allow you to perform calculations across a set of related rows
**without collapsing those rows into a single output row**.

This is the defining idea.

Compare:

``` sql
GROUP BY
```

with:

``` sql
WINDOW FUNCTION
```

`GROUP BY` changes the grain of the result.

A window function normally preserves the original rows while adding
analytical information to them.

------------------------------------------------------------------------

# 11.1 The Core Idea

Suppose we have:

``` text
employees
------------------------------------------------
employee_id | name  | department | salary
------------------------------------------------
1           | Alice | IT         | 90000
2           | Bob   | IT         | 80000
3           | Carol | IT         | 70000
4           | David | HR         | 60000
5           | Eve   | HR         | 50000
```

A normal aggregate:

``` sql
SELECT
    department,
    AVG(salary) AS avg_salary
FROM employees
GROUP BY department;
```

produces:

``` text
department | avg_salary
-----------|-----------
IT         | 80000
HR         | 55000
```

The original employee rows disappear.

A window function:

``` sql
SELECT
    employee_id,
    name,
    department,
    salary,
    AVG(salary) OVER (
        PARTITION BY department
    ) AS department_avg
FROM employees;
```

produces:

``` text
employee_id | name  | department | salary | department_avg
------------|-------|------------|--------|---------------
1           | Alice | IT         | 90000  | 80000
2           | Bob   | IT         | 80000  | 80000
3           | Carol | IT         | 70000  | 80000
4           | David | HR         | 60000  | 55000
5           | Eve   | HR         | 50000  | 55000
```

The employee rows remain.

That is the central power of window functions.

------------------------------------------------------------------------

# 11.2 Window Functions vs `GROUP BY`

This distinction is critical.

## `GROUP BY`

``` sql
SELECT
    department,
    AVG(salary)
FROM employees
GROUP BY department;
```

Conceptually:

``` text
many rows
    ↓
group
    ↓
one row per group
```

## Window Function

``` sql
SELECT
    employee_id,
    department,
    salary,
    AVG(salary) OVER (
        PARTITION BY department
    )
FROM employees;
```

Conceptually:

``` text
many rows
    ↓
analyze each row's window
    ↓
keep every original row
```

Mental model:

``` text
GROUP BY
    → collapse

WINDOW FUNCTION
    → annotate
```

------------------------------------------------------------------------

# 11.3 What Is a Window?

A window is the set of rows that a window function can see for a
particular output row.

For example:

``` sql
AVG(salary) OVER (
    PARTITION BY department
)
```

For Alice:

``` text
Window = all IT employees
```

For David:

``` text
Window = all HR employees
```

The window can be controlled using:

``` text
PARTITION BY
ORDER BY
frame specification
```

These determine which rows participate in the calculation.

------------------------------------------------------------------------

# 11.4 The `OVER()` Clause

The `OVER()` clause tells SQL that a function should operate as a window
function.

Basic syntax:

``` sql
function(...) OVER (...)
```

Example:

``` sql
AVG(salary) OVER ()
```

This calculates an average over the relevant complete result set.

Another example:

``` sql
AVG(salary) OVER (
    PARTITION BY department
)
```

This calculates a separate average for each department.

------------------------------------------------------------------------

# 11.5 Empty `OVER()`

An empty `OVER()` means there is no explicit partition or ordering.

Example:

``` sql
SELECT
    employee_id,
    salary,
    AVG(salary) OVER () AS company_avg
FROM employees;
```

Every row receives the same company-wide average.

If salaries are:

``` text
90000
80000
70000
60000
50000
```

the average is:

``` text
70000
```

Result:

``` text
employee | salary | company_avg
---------|--------|------------
Alice    | 90000  | 70000
Bob      | 80000  | 70000
Carol    | 70000  | 70000
David    | 60000  | 70000
Eve      | 50000  | 70000
```

------------------------------------------------------------------------

# 11.6 `PARTITION BY`

`PARTITION BY` divides the result into independent groups for the window
calculation.

Example:

``` sql
AVG(salary) OVER (
    PARTITION BY department
)
```

Conceptually:

``` text
All employees
     │
     ├── IT partition
     │     ├── Alice
     │     ├── Bob
     │     └── Carol
     │
     └── HR partition
           ├── David
           └── Eve
```

The function is evaluated separately inside each partition.

------------------------------------------------------------------------

# 11.7 `PARTITION BY` Does Not Collapse Rows

This is extremely important.

``` sql
SELECT
    name,
    department,
    salary,
    AVG(salary) OVER (
        PARTITION BY department
    ) AS avg_salary
FROM employees;
```

Still returns:

``` text
one row per employee
```

`PARTITION BY` is not equivalent to:

``` sql
GROUP BY department
```

It creates logical windows rather than collapsing the result.

------------------------------------------------------------------------

# 11.8 Multiple Partition Columns

You can partition by multiple columns:

``` sql
PARTITION BY department, job_title
```

Example:

``` sql
AVG(salary) OVER (
    PARTITION BY department, job_title
)
```

Now each unique combination forms a partition:

``` text
IT + Engineer
IT + Manager
HR + Recruiter
HR + Manager
```

------------------------------------------------------------------------

# 11.9 `ORDER BY` Inside `OVER()`

The `ORDER BY` inside the window definition determines the logical
ordering of rows within the window.

Example:

``` sql
ROW_NUMBER() OVER (
    ORDER BY salary DESC
)
```

This assigns numbers according to salary:

``` text
salary | row_number
-------|-----------
90000  | 1
80000  | 2
70000  | 3
```

With partitions:

``` sql
ROW_NUMBER() OVER (
    PARTITION BY department
    ORDER BY salary DESC
)
```

the numbering restarts for each department.

------------------------------------------------------------------------

# 11.10 Outer `ORDER BY` vs Window `ORDER BY`

These are different.

Consider:

``` sql
SELECT
    name,
    salary,
    ROW_NUMBER() OVER (
        ORDER BY salary DESC
    ) AS rn
FROM employees
ORDER BY name;
```

The window `ORDER BY` determines:

``` text
how row_number is assigned
```

The final `ORDER BY` determines:

``` text
how the result is displayed
```

So the output could be displayed alphabetically while the ranking is
still based on salary.

Mental model:

``` text
OVER(ORDER BY ...)
    → calculation order

final ORDER BY
    → presentation order
```

------------------------------------------------------------------------

# 11.11 Window Function Categories

Window functions can broadly be grouped into:

``` text
Ranking
    ROW_NUMBER
    RANK
    DENSE_RANK
    NTILE

Offset / navigation
    LAG
    LEAD

Value functions
    FIRST_VALUE
    LAST_VALUE

Aggregate window functions
    SUM
    AVG
    COUNT
    MIN
    MAX

Statistical / analytical functions
    database-specific functions
```

This section focuses on the requested core functions.

------------------------------------------------------------------------

# 11.12 `ROW_NUMBER()`

`ROW_NUMBER()` assigns a unique sequential number to each row in the
window.

Syntax:

``` sql
ROW_NUMBER() OVER (
    ORDER BY ...
)
```

Example:

``` sql
SELECT
    name,
    salary,
    ROW_NUMBER() OVER (
        ORDER BY salary DESC
    ) AS row_num
FROM employees;
```

Result:

``` text
name  | salary | row_num
------|--------|--------
Alice | 90000  | 1
Bob   | 80000  | 2
Carol | 70000  | 3
David | 60000  | 4
Eve   | 50000  | 5
```

------------------------------------------------------------------------

# 11.13 `ROW_NUMBER()` with `PARTITION BY`

``` sql
ROW_NUMBER() OVER (
    PARTITION BY department
    ORDER BY salary DESC
)
```

Example:

``` text
department | name  | salary | row_number
-----------|-------|--------|-----------
IT         | Alice | 90000  | 1
IT         | Bob   | 80000  | 2
IT         | Carol | 70000  | 3
HR         | David | 60000  | 1
HR         | Eve   | 50000  | 2
```

The numbering restarts for each partition.

------------------------------------------------------------------------

# 11.14 Common Use Case: Top 1 Row Per Group

Suppose you want the highest-paid employee in every department.

Use:

``` sql
WITH ranked_employees AS (
    SELECT
        employee_id,
        name,
        department,
        salary,
        ROW_NUMBER() OVER (
            PARTITION BY department
            ORDER BY salary DESC
        ) AS rn
    FROM employees
)
SELECT *
FROM ranked_employees
WHERE rn = 1;
```

This is one of the most important practical window-function patterns.

------------------------------------------------------------------------

# 11.15 Why a CTE Is Often Needed

You generally cannot write:

``` sql
SELECT
    ...,
    ROW_NUMBER() OVER (...) AS rn
FROM employees
WHERE rn = 1;
```

because the window function result is not normally available to the
`WHERE` clause at that logical processing stage.

Instead:

``` text
inner query
    ↓
calculate row_number
    ↓
outer query
    ↓
filter rn
```

Use:

-   CTE
-   derived table
-   another query layer

------------------------------------------------------------------------

# 11.16 `RANK()`

`RANK()` assigns the same rank to tied values.

Example:

``` sql
SELECT
    name,
    salary,
    RANK() OVER (
        ORDER BY salary DESC
    ) AS salary_rank
FROM employees;
```

Suppose salaries are:

``` text
90000
80000
80000
70000
```

Then:

``` text
salary | rank
-------|-----
90000  | 1
80000  | 2
80000  | 2
70000  | 4
```

Notice the gap.

There is no rank `3`.

------------------------------------------------------------------------

# 11.17 `DENSE_RANK()`

`DENSE_RANK()` also gives tied values the same rank, but does not leave
gaps.

Example:

``` sql
DENSE_RANK() OVER (
    ORDER BY salary DESC
)
```

For:

``` text
90000
80000
80000
70000
```

result:

``` text
salary | dense_rank
-------|-----------
90000  | 1
80000  | 2
80000  | 2
70000  | 3
```

------------------------------------------------------------------------

# 11.18 `ROW_NUMBER` vs `RANK` vs `DENSE_RANK`

This is one of the most common SQL interview questions.

Suppose:

``` text
salary
------
100
90
90
80
70
70
```

Results:

``` text
salary | ROW_NUMBER | RANK | DENSE_RANK
-------|------------|------|------------
100    | 1          | 1    | 1
90     | 2          | 2    | 2
90     | 3          | 2    | 2
80     | 4          | 4    | 3
70     | 5          | 5    | 4
70     | 6          | 5    | 4
```

Mental model:

``` text
ROW_NUMBER
    → always unique sequence

RANK
    → ties share rank, gaps appear

DENSE_RANK
    → ties share rank, no gaps
```

------------------------------------------------------------------------

# 11.19 When to Use Each Ranking Function

Use `ROW_NUMBER()` when:

> You need exactly one unique sequence position per row.

Typical use:

``` text
latest record per customer
top one per group
deduplication
pagination
```

Use `RANK()` when:

> Competition-style ranking with gaps is meaningful.

Example:

``` text
1st
2nd
2nd
4th
```

Use `DENSE_RANK()` when:

> You want distinct ranking levels without gaps.

Example:

``` text
1
2
2
3
```

------------------------------------------------------------------------

# 11.20 Determinism and Ties

Consider:

``` sql
ROW_NUMBER() OVER (
    ORDER BY salary DESC
)
```

What happens when two employees have the same salary?

The database may have freedom in how those tied rows are ordered unless
additional ordering columns are provided.

For deterministic ordering:

``` sql
ROW_NUMBER() OVER (
    ORDER BY salary DESC, employee_id
)
```

Now ties are broken by `employee_id`.

This is especially important in:

-   pagination
-   deduplication
-   latest-record selection
-   ETL pipelines

------------------------------------------------------------------------

# 11.21 `NTILE()`

`NTILE(n)` divides rows into approximately equal-sized buckets.

Example:

``` sql
NTILE(4) OVER (
    ORDER BY salary
)
```

creates four buckets:

``` text
1 → first quartile
2 → second quartile
3 → third quartile
4 → fourth quartile
```

It is useful for:

-   quartiles
-   customer segmentation
-   score bands
-   approximate percentile groups
-   salary bands

------------------------------------------------------------------------

# 11.22 `NTILE()` Example

Suppose there are 10 rows:

``` sql
SELECT
    name,
    salary,
    NTILE(4) OVER (
        ORDER BY salary
    ) AS quartile
FROM employees;
```

The rows are distributed as evenly as possible among four buckets.

Because 10 is not divisible by 4, the buckets cannot all contain exactly
the same number of rows.

The database distributes the remainder according to its `NTILE`
semantics.

------------------------------------------------------------------------

# 11.23 `NTILE` Is Not the Same as Percentile

`NTILE(4)` creates four row groups.

It does not mean:

``` text
exact statistical quartile calculation
```

For more precise percentile/statistical analysis, databases often
provide functions such as:

``` text
PERCENTILE_CONT
PERCENTILE_DISC
CUME_DIST
PERCENT_RANK
```

These are database-specific or dialect-dependent.

------------------------------------------------------------------------

# 11.24 `LAG()`

`LAG()` accesses a previous row in the window.

Syntax:

``` sql
LAG(expression) OVER (
    ORDER BY ...
)
```

Example:

``` sql
SELECT
    order_date,
    revenue,
    LAG(revenue) OVER (
        ORDER BY order_date
    ) AS previous_revenue
FROM daily_sales;
```

Result conceptually:

``` text
date       | revenue | previous_revenue
-----------|---------|-----------------
Jan 1      | 100     | NULL
Jan 2      | 120     | 100
Jan 3      | 90      | 120
Jan 4      | 150     | 90
```

------------------------------------------------------------------------

# 11.25 `LAG()` for Change Detection

A common pattern:

``` sql
SELECT
    order_date,
    revenue,
    revenue
      - LAG(revenue) OVER (
            ORDER BY order_date
        ) AS revenue_change
FROM daily_sales;
```

Result:

``` text
date  | revenue | change
------|---------|-------
Jan 1 | 100     | NULL
Jan 2 | 120     | 20
Jan 3 | 90      | -30
Jan 4 | 150     | 60
```

------------------------------------------------------------------------

# 11.26 `LAG()` for Percentage Growth

``` sql
SELECT
    order_date,
    revenue,
    LAG(revenue) OVER (
        ORDER BY order_date
    ) AS previous_revenue,
    (
        revenue - LAG(revenue) OVER (
            ORDER BY order_date
        )
    ) / NULLIF(
        LAG(revenue) OVER (
            ORDER BY order_date
        ),
        0
    ) * 100 AS growth_pct
FROM daily_sales;
```

A cleaner approach may calculate `LAG()` once in a CTE:

``` sql
WITH sales_with_previous AS (
    SELECT
        order_date,
        revenue,
        LAG(revenue) OVER (
            ORDER BY order_date
        ) AS previous_revenue
    FROM daily_sales
)
SELECT
    order_date,
    revenue,
    previous_revenue,
    (revenue - previous_revenue)
        / NULLIF(previous_revenue, 0) * 100 AS growth_pct
FROM sales_with_previous;
```

This is easier to read and often easier to maintain.

------------------------------------------------------------------------

# 11.27 `LAG()` with an Offset

Default:

``` sql
LAG(revenue)
```

means:

``` text
previous row
```

You can specify an offset:

``` sql
LAG(revenue, 2)
```

meaning:

``` text
two rows earlier
```

Example:

``` sql
LAG(revenue, 7) OVER (
    ORDER BY order_date
)
```

can be useful for:

``` text
day-over-day comparisons
week-over-week comparisons
period-over-period comparisons
```

The interpretation depends on the row ordering and data grain.

------------------------------------------------------------------------

# 11.28 `LEAD()`

`LEAD()` is the opposite of `LAG()`.

It accesses a future row.

Example:

``` sql
SELECT
    order_date,
    revenue,
    LEAD(revenue) OVER (
        ORDER BY order_date
    ) AS next_revenue
FROM daily_sales;
```

Result:

``` text
date  | revenue | next_revenue
------|---------|-------------
Jan 1 | 100     | 120
Jan 2 | 120     | 90
Jan 3 | 90      | 150
Jan 4 | 150     | NULL
```

------------------------------------------------------------------------

# 11.29 `LAG` vs `LEAD`

``` text
LAG
    current row ← previous row

LEAD
    current row → next row
```

Use `LAG()` for:

``` text
previous value
previous event
previous status
previous period
```

Use `LEAD()` for:

``` text
next value
next event
next status
next period
```

------------------------------------------------------------------------

# 11.30 `FIRST_VALUE()`

`FIRST_VALUE()` returns the first value according to the window
ordering/frame.

Example:

``` sql
SELECT
    name,
    department,
    salary,
    FIRST_VALUE(name) OVER (
        PARTITION BY department
        ORDER BY salary DESC
    ) AS highest_paid_employee
FROM employees;
```

For every employee in a department, the function can return the name of
the highest-paid employee according to the specified window.

------------------------------------------------------------------------

# 11.31 `FIRST_VALUE()` Example

Suppose IT salaries are:

``` text
Alice   90000
Bob     80000
Carol   70000
```

With:

``` sql
FIRST_VALUE(name) OVER (
    PARTITION BY department
    ORDER BY salary DESC
)
```

each IT row can receive:

``` text
Alice
```

because Alice is first according to descending salary.

------------------------------------------------------------------------

# 11.32 `LAST_VALUE()` --- The Important Trap

`LAST_VALUE()` is one of the most frequently misunderstood window
functions.

A naïve query:

``` sql
LAST_VALUE(salary) OVER (
    PARTITION BY department
    ORDER BY salary
)
```

does not necessarily return the final salary in the entire department.

Why?

Because the **window frame** matters.

The function returns the last value in its current frame, not
automatically the last value in the entire partition.

This makes window frames essential.

------------------------------------------------------------------------

# 11.33 Window Frames

A window can be thought of as having two levels:

``` text
PARTITION
    ↓
set of related rows

FRAME
    ↓
subset of the partition relevant to the current row
```

For example:

``` sql
SUM(amount) OVER (
    ORDER BY order_date
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

The partition may contain all rows.

The frame for each row contains:

``` text
first row
through
current row
```

------------------------------------------------------------------------

# 11.34 Frame Syntax

General syntax:

``` sql
ROWS BETWEEN start AND end
```

or:

``` sql
RANGE BETWEEN start AND end
```

or, in supported dialects:

``` sql
GROUPS BETWEEN start AND end
```

Common boundaries:

``` text
UNBOUNDED PRECEDING
n PRECEDING
CURRENT ROW
n FOLLOWING
UNBOUNDED FOLLOWING
```

------------------------------------------------------------------------

# 11.35 `ROWS`

`ROWS` describes a physical/logical number of rows relative to the
current row.

Example:

``` sql
ROWS BETWEEN 2 PRECEDING AND CURRENT ROW
```

means:

``` text
current row
+
previous 2 rows
```

So the frame contains up to three rows.

------------------------------------------------------------------------

# 11.36 `RANGE`

`RANGE` is value-based rather than simply row-count-based.

It considers the ordering values and peer groups according to SQL's
frame semantics.

This distinction becomes especially important when the `ORDER BY` column
contains duplicates.

Do not casually substitute:

``` text
ROWS
```

and:

``` text
RANGE
```

without understanding the data.

------------------------------------------------------------------------

# 11.37 `GROUPS`

Some databases support:

``` sql
GROUPS
```

which operates in terms of peer groups rather than individual physical
rows.

This is especially useful when the ordering column contains ties.

Support and exact behavior vary by database.

------------------------------------------------------------------------

# 11.38 Common Frame: Running Total

A classic running total is:

``` sql
SUM(amount) OVER (
    ORDER BY order_date
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

Suppose:

``` text
date | amount
-----|-------
Jan1 | 100
Jan2 | 200
Jan3 | 150
Jan4 | 50
```

Running total:

``` text
date | amount | running_total
-----|--------|--------------
Jan1 | 100    | 100
Jan2 | 200    | 300
Jan3 | 150    | 450
Jan4 | 50     | 500
```

------------------------------------------------------------------------

# 11.39 Running Total with `PARTITION BY`

Suppose sales are tracked by store.

``` sql
SUM(amount) OVER (
    PARTITION BY store_id
    ORDER BY sale_date
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

The running total restarts for every store.

Conceptually:

``` text
Store A:
100 → 150 → 220

Store B:
50 → 90 → 140
```

------------------------------------------------------------------------

# 11.40 Running Total Without Explicit `ROWS`

You may encounter:

``` sql
SUM(amount) OVER (
    ORDER BY order_date
)
```

This can work as a cumulative calculation, but the exact default frame
depends on the database and function/context.

In particular, default framing and peer behavior can make results
surprising when the ordering column has duplicate values.

For precise analytical SQL, explicitly specify the frame when frame
semantics matter:

``` sql
ROWS BETWEEN UNBOUNDED PRECEDING
         AND CURRENT ROW
```

This makes the intended behavior clearer.

------------------------------------------------------------------------

# 11.41 Moving Average

A moving average calculates the average over a sliding window.

Example:

``` sql
AVG(revenue) OVER (
    ORDER BY order_date
    ROWS BETWEEN 2 PRECEDING
             AND CURRENT ROW
)
```

This produces a three-row moving average:

``` text
current row
+
previous row
+
two rows back
```

Example:

``` text
revenue:
100
120
140
160
```

Moving average:

``` text
100
110
120
140
```

The first rows have fewer available preceding rows, so their frames are
smaller.

------------------------------------------------------------------------

# 11.42 Seven-Period Moving Average

For daily data:

``` sql
AVG(revenue) OVER (
    ORDER BY order_date
    ROWS BETWEEN 6 PRECEDING
             AND CURRENT ROW
)
```

This represents up to seven rows:

``` text
current day
+
previous 6 rows
```

Important:

> `ROWS BETWEEN 6 PRECEDING AND CURRENT ROW` means seven rows, not
> necessarily seven calendar days.

If dates are missing, the distinction matters.

For true calendar-day windows, a value-based/range-based strategy or
calendar table may be more appropriate, depending on the database.

------------------------------------------------------------------------

# 11.43 Window Frame: `UNBOUNDED PRECEDING`

``` sql
ROWS BETWEEN UNBOUNDED PRECEDING
         AND CURRENT ROW
```

means:

``` text
from the beginning of the frame
through the current row
```

This is the canonical running-total frame.

------------------------------------------------------------------------

# 11.44 Window Frame: `CURRENT ROW`

``` sql
CURRENT ROW
```

means the current row according to the frame semantics.

Its interaction with `ROWS`, `RANGE`, and `GROUPS` differs.

For `ROWS`, it refers to the current physical/logical row position.

For `RANGE` and `GROUPS`, peer-group behavior can matter.

------------------------------------------------------------------------

# 11.45 Window Frame: `UNBOUNDED FOLLOWING`

``` sql
ROWS BETWEEN CURRENT ROW
         AND UNBOUNDED FOLLOWING
```

means:

``` text
current row
through the end
```

This can be useful for calculations involving future rows.

------------------------------------------------------------------------

# 11.46 `LAST_VALUE()` Correctly

Suppose you want the final salary in the ordered partition.

You can explicitly define the frame:

``` sql
LAST_VALUE(salary) OVER (
    PARTITION BY department
    ORDER BY salary
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND UNBOUNDED FOLLOWING
)
```

Now the frame spans the entire partition.

Therefore `LAST_VALUE()` can return the final value according to the
specified ordering.

This is a critical example of why window frames matter.

------------------------------------------------------------------------

# 11.47 `FIRST_VALUE()` vs `LAST_VALUE()`

``` text
FIRST_VALUE
    → first value in the relevant frame

LAST_VALUE
    → last value in the relevant frame
```

Not:

``` text
FIRST_VALUE
    → first value in entire partition automatically

LAST_VALUE
    → last value in entire partition automatically
```

The frame determines the scope.

------------------------------------------------------------------------

# 11.48 Ranking Does Not Usually Need a Frame

Ranking functions such as:

``` sql
ROW_NUMBER()
RANK()
DENSE_RANK()
NTILE()
```

are primarily driven by:

``` text
PARTITION BY
ORDER BY
```

The concept of a frame is much more important for:

``` text
aggregate window functions
FIRST_VALUE
LAST_VALUE
```

and other frame-sensitive analytical functions.

------------------------------------------------------------------------

# 11.49 Window Aggregates

Many ordinary aggregate functions can become window functions:

``` sql
SUM(...)
AVG(...)
COUNT(...)
MIN(...)
MAX(...)
```

by adding:

``` sql
OVER(...)
```

Example:

``` sql
SUM(amount) OVER ()
```

company/global total.

Example:

``` sql
SUM(amount) OVER (
    PARTITION BY customer_id
)
```

customer total repeated on every customer row.

Example:

``` sql
SUM(amount) OVER (
    PARTITION BY customer_id
    ORDER BY order_date
)
```

customer running total.

------------------------------------------------------------------------

# 11.50 Group Aggregate vs Window Aggregate

Consider:

``` sql
SELECT
    customer_id,
    SUM(amount)
FROM orders
GROUP BY customer_id;
```

Output:

``` text
customer | total
---------|------
A        | 300
B        | 500
```

Window version:

``` sql
SELECT
    order_id,
    customer_id,
    amount,
    SUM(amount) OVER (
        PARTITION BY customer_id
    ) AS customer_total
FROM orders;
```

Output:

``` text
order | customer | amount | customer_total
------|----------|--------|---------------
1     | A        | 100    | 300
2     | A        | 200    | 300
3     | B        | 500    | 500
```

The aggregate value is attached to each detail row.

------------------------------------------------------------------------

# 11.51 Running Total vs Total

These are different:

``` sql
SUM(amount) OVER (
    PARTITION BY customer_id
)
```

gives the complete customer total on every row.

Whereas:

``` sql
SUM(amount) OVER (
    PARTITION BY customer_id
    ORDER BY order_date
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

gives the cumulative total up to the current row.

------------------------------------------------------------------------

# 11.52 Percentage of Total

A powerful analytical pattern:

``` sql
SELECT
    product_id,
    revenue,
    revenue
      / SUM(revenue) OVER () * 100 AS pct_of_total
FROM sales;
```

This allows each row to be compared with the overall total without
collapsing the result.

------------------------------------------------------------------------

# 11.53 Percentage of Group Total

``` sql
SELECT
    department,
    employee_id,
    salary,
    salary
      / SUM(salary) OVER (
            PARTITION BY department
        ) * 100 AS pct_of_department_salary
FROM employees;
```

Now each employee's salary is expressed as a percentage of their
department's total salary.

------------------------------------------------------------------------

# 11.54 Difference from Previous Row

Using `LAG()`:

``` sql
SELECT
    order_date,
    revenue,
    revenue
      - LAG(revenue) OVER (
            ORDER BY order_date
        ) AS change
FROM daily_sales;
```

This pattern is fundamental for time-series analysis.

------------------------------------------------------------------------

# 11.55 Difference from Next Row

Using `LEAD()`:

``` sql
SELECT
    order_date,
    revenue,
    LEAD(revenue) OVER (
        ORDER BY order_date
    ) - revenue AS next_change
FROM daily_sales;
```

------------------------------------------------------------------------

# 11.56 Top N Per Group

A classic problem:

> Find the top 3 highest-paid employees in each department.

Solution:

``` sql
WITH ranked AS (
    SELECT
        employee_id,
        name,
        department,
        salary,
        ROW_NUMBER() OVER (
            PARTITION BY department
            ORDER BY salary DESC
        ) AS rn
    FROM employees
)
SELECT *
FROM ranked
WHERE rn <= 3;
```

If ties should all receive the same rank, consider:

``` sql
RANK()
```

or:

``` sql
DENSE_RANK()
```

depending on the desired semantics.

------------------------------------------------------------------------

# 11.57 Latest Row Per Group

Suppose there are multiple records per customer and you need the latest
one.

``` sql
WITH ranked AS (
    SELECT
        customer_id,
        order_id,
        order_date,
        amount,
        ROW_NUMBER() OVER (
            PARTITION BY customer_id
            ORDER BY order_date DESC, order_id DESC
        ) AS rn
    FROM orders
)
SELECT *
FROM ranked
WHERE rn = 1;
```

The second ordering column:

``` sql
order_id DESC
```

acts as a deterministic tie-breaker when dates are identical.

------------------------------------------------------------------------

# 11.58 Deduplication with `ROW_NUMBER()`

Suppose duplicate customer records exist:

``` text
customer_id | email | created_at
```

You can retain one row per email:

``` sql
WITH ranked AS (
    SELECT
        *,
        ROW_NUMBER() OVER (
            PARTITION BY email
            ORDER BY created_at DESC
        ) AS rn
    FROM customers
)
SELECT *
FROM ranked
WHERE rn = 1;
```

This means:

> Within each email group, keep the newest row.

This is a very common ETL/data-cleaning pattern.

------------------------------------------------------------------------

# 11.59 Detecting Changes Between Consecutive Rows

Suppose customer status changes over time.

``` sql
WITH status_history AS (
    SELECT
        customer_id,
        event_time,
        status,
        LAG(status) OVER (
            PARTITION BY customer_id
            ORDER BY event_time
        ) AS previous_status
    FROM customer_events
)
SELECT *
FROM status_history
WHERE previous_status IS DISTINCT FROM status;
```

The exact `IS DISTINCT FROM` syntax is dialect-dependent.

Conceptually:

``` text
current status
vs
previous status
```

This helps detect state transitions.

------------------------------------------------------------------------

# 11.60 Session/Gap Analysis

`LAG()` can be used to compare timestamps.

Conceptually:

``` sql
LAG(event_time) OVER (
    PARTITION BY user_id
    ORDER BY event_time
)
```

Then calculate:

``` text
current_event_time - previous_event_time
```

If the gap exceeds a threshold, a new session can be identified.

This is a common analytics pattern.

------------------------------------------------------------------------

# 11.61 Cumulative Maximum

``` sql
MAX(revenue) OVER (
    ORDER BY order_date
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
) AS running_max
```

This tells you the highest observed value up to each point.

Useful for:

-   peak tracking
-   financial analysis
-   drawdown calculations
-   monitoring systems

------------------------------------------------------------------------

# 11.62 Running Minimum

Similarly:

``` sql
MIN(value) OVER (
    ORDER BY event_time
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

------------------------------------------------------------------------

# 11.63 Moving Sum

``` sql
SUM(amount) OVER (
    ORDER BY order_date
    ROWS BETWEEN 6 PRECEDING
             AND CURRENT ROW
)
```

This calculates a rolling seven-row sum.

------------------------------------------------------------------------

# 11.64 Moving Maximum

``` sql
MAX(value) OVER (
    ORDER BY event_date
    ROWS BETWEEN 6 PRECEDING
             AND CURRENT ROW
)
```

Useful for:

``` text
rolling peak
rolling SLA
rolling traffic maximum
```

------------------------------------------------------------------------

# 11.65 `NULL` Behavior

Window functions interact with `NULL` according to the semantics of the
underlying function and database.

For example:

``` sql
AVG(value) OVER (...)
```

generally ignores `NULL` values, just as ordinary `AVG()` does.

`LAG()` can return `NULL` when:

-   there is no previous row
-   the previous row's value is itself `NULL`

Therefore:

``` sql
LAG(value)
```

returning `NULL` does not automatically tell you whether:

``` text
there was no previous row
```

or:

``` text
the previous value was NULL
```

This distinction can matter.

------------------------------------------------------------------------

# 11.66 Default Values with `LAG()` / `LEAD()`

Many SQL dialects support a default argument:

``` sql
LAG(value, 1, 0)
```

Conceptually:

``` text
previous value
or 0 if no previous row exists
```

Example:

``` sql
LAG(revenue, 1, 0) OVER (
    ORDER BY order_date
)
```

Syntax and support should be checked for the target database.

Be careful about replacing missing historical data with `0` when `NULL`
has a different business meaning.

------------------------------------------------------------------------

# 11.67 Window Function Logical Processing

A useful simplified model of SQL query processing is:

``` text
FROM
 ↓
JOIN
 ↓
WHERE
 ↓
GROUP BY
 ↓
HAVING
 ↓
window functions
 ↓
SELECT output
 ↓
ORDER BY
 ↓
LIMIT/OFFSET
```

The exact formal logical processing order is more nuanced, but this
model explains an important practical fact:

> You generally cannot reference a window-function alias directly in
> `WHERE`.

For example:

``` sql
SELECT
    *,
    ROW_NUMBER() OVER (...) AS rn
FROM employees
WHERE rn = 1;
```

usually fails.

Instead:

``` sql
WITH ranked AS (
    SELECT
        *,
        ROW_NUMBER() OVER (...) AS rn
    FROM employees
)
SELECT *
FROM ranked
WHERE rn = 1;
```

------------------------------------------------------------------------

# 11.68 Window Functions and `WHERE`

If you want to filter using a window result, use another query layer.

Pattern:

``` text
base query
   ↓
window calculation
   ↓
outer query
   ↓
filter
```

Example:

``` sql
WITH ranked AS (
    SELECT
        *,
        RANK() OVER (
            PARTITION BY department
            ORDER BY salary DESC
        ) AS rnk
    FROM employees
)
SELECT *
FROM ranked
WHERE rnk <= 3;
```

------------------------------------------------------------------------

# 11.69 `QUALIFY`

Some SQL databases support:

``` sql
QUALIFY
```

which allows filtering after window functions.

Example:

``` sql
SELECT
    *,
    ROW_NUMBER() OVER (
        PARTITION BY customer_id
        ORDER BY order_date DESC
    ) AS rn
FROM orders
QUALIFY rn = 1;
```

This can be much cleaner than a CTE.

However:

> `QUALIFY` is not universally supported.

For portable SQL, the CTE/derived-table pattern is safer.

------------------------------------------------------------------------

# 11.70 Named Windows

Some SQL dialects support a `WINDOW` clause that lets you define a
reusable window specification.

Conceptually:

``` sql
SELECT
    employee_id,
    salary,
    RANK() OVER w AS salary_rank,
    DENSE_RANK() OVER w AS dense_salary_rank
FROM employees
WINDOW w AS (
    PARTITION BY department
    ORDER BY salary DESC
);
```

This reduces repetition.

Support varies by database.

------------------------------------------------------------------------

# 11.71 Window Function Composition

Window functions generally cannot simply be nested inside one another:

``` sql
SUM(
    ROW_NUMBER() OVER (...)
) OVER (...)
```

is not a normal valid pattern.

Instead, use multiple query layers:

``` text
query 1
    ↓
calculate window function
    ↓
query 2
    ↓
calculate another window function
```

Example:

``` sql
WITH ranked AS (
    SELECT
        *,
        ROW_NUMBER() OVER (
            PARTITION BY department
            ORDER BY salary DESC
        ) AS rn
    FROM employees
)
SELECT
    *,
    SUM(salary) OVER (
        PARTITION BY department
    ) AS department_total
FROM ranked;
```

------------------------------------------------------------------------

# 11.72 Window Functions vs Self Join

Before window functions became common, many analytical queries were
expressed using self joins.

For example:

``` text
current row
    JOIN
previous row
```

Window functions often make such logic simpler:

``` sql
LAG(value) OVER (
    ORDER BY event_time
)
```

instead of manually joining each row to its predecessor.

Benefits include:

-   clearer intent
-   fewer join conditions
-   easier maintenance
-   often better optimization opportunities

But always inspect the execution plan for performance-critical
workloads.

------------------------------------------------------------------------

# 11.73 Window Functions vs Correlated Subqueries

Many correlated subqueries can be replaced with window functions.

Example:

``` sql
SELECT
    e.*,
    (
        SELECT AVG(e2.salary)
        FROM employees e2
        WHERE e2.department = e.department
    ) AS department_avg
FROM employees e;
```

Window equivalent:

``` sql
SELECT
    e.*,
    AVG(salary) OVER (
        PARTITION BY department
    ) AS department_avg
FROM employees e;
```

The window-function version is usually clearer for this analytical
pattern.

------------------------------------------------------------------------

# 11.74 Window Functions vs `GROUP BY`

Use `GROUP BY` when:

``` text
you want fewer rows
```

Use a window function when:

``` text
you want analytical information
while keeping detail rows
```

Example:

``` text
GROUP BY:
5,000 orders
   ↓
500 customers
```

Window:

``` text
5,000 orders
   ↓
5,000 orders + customer-level analytics
```

------------------------------------------------------------------------

# 11.75 Window Frames --- Detailed Mental Model

Imagine a partition:

``` text
Row 1
Row 2
Row 3
Row 4
Row 5
```

For Row 4:

``` sql
ROWS BETWEEN 2 PRECEDING AND CURRENT ROW
```

the frame is:

``` text
Row 2
Row 3
Row 4
```

For Row 2:

``` text
Row 1
Row 2
```

because there is no Row 0.

For Row 1:

``` text
Row 1
```

This is why moving averages near the beginning of a dataset often use
fewer observations than the full window size.

------------------------------------------------------------------------

# 11.76 Frame Examples

## Entire Partition

``` sql
ROWS BETWEEN UNBOUNDED PRECEDING
         AND UNBOUNDED FOLLOWING
```

``` text
[first row ... current row ... last row]
```

------------------------------------------------------------------------

## Beginning Through Current Row

``` sql
ROWS BETWEEN UNBOUNDED PRECEDING
         AND CURRENT ROW
```

``` text
[first row ... current row]
```

Typical use:

``` text
running total
running average
running maximum
```

------------------------------------------------------------------------

## Previous 3 Rows + Current

``` sql
ROWS BETWEEN 3 PRECEDING
         AND CURRENT ROW
```

``` text
[row-3, row-2, row-1, current]
```

Up to four rows.

------------------------------------------------------------------------

## Current + Next 3

``` sql
ROWS BETWEEN CURRENT ROW
         AND 3 FOLLOWING
```

Useful for forward-looking calculations.

------------------------------------------------------------------------

# 11.77 Centered Moving Average

A centered five-row moving average:

``` sql
AVG(value) OVER (
    ORDER BY event_time
    ROWS BETWEEN 2 PRECEDING
             AND 2 FOLLOWING
)
```

Conceptually:

``` text
2 previous
+
current
+
2 following
```

This is useful in smoothing and time-series analysis.

------------------------------------------------------------------------

# 11.78 `ROWS` vs Calendar Time

Suppose:

``` text
Jan 1
Jan 2
Jan 7
Jan 8
```

Then:

``` sql
ROWS BETWEEN 2 PRECEDING AND CURRENT ROW
```

means:

``` text
previous 2 records
```

not:

``` text
previous 2 calendar days
```

This distinction is extremely important in real-world analytics.

For calendar-aware analysis, you may need:

-   a calendar/date dimension
-   `RANGE` with appropriate ordering
-   dialect-specific interval frames
-   explicit date joins

depending on the database.

------------------------------------------------------------------------

# 11.79 Duplicate Ordering Values and Window Frames

Suppose:

``` text
date | amount
-----|-------
Jan1 | 100
Jan1 | 200
Jan2 | 300
```

If you use:

``` sql
ORDER BY date
```

there are peers with the same ordering value.

The behavior of:

``` text
ROWS
RANGE
GROUPS
```

can differ significantly.

For deterministic row-level sequencing, consider adding a unique
tie-breaker:

``` sql
ORDER BY date, transaction_id
```

This is especially important for:

``` text
ROW_NUMBER
running calculations
LAG/LEAD
```

------------------------------------------------------------------------

# 11.80 Practical Example: Sales Dashboard

Suppose:

``` text
daily_sales
----------------------------
sale_date | revenue
```

A useful dashboard query:

``` sql
SELECT
    sale_date,
    revenue,

    SUM(revenue) OVER (
        ORDER BY sale_date
        ROWS BETWEEN UNBOUNDED PRECEDING
                 AND CURRENT ROW
    ) AS running_revenue,

    AVG(revenue) OVER (
        ORDER BY sale_date
        ROWS BETWEEN 6 PRECEDING
                 AND CURRENT ROW
    ) AS seven_day_avg,

    LAG(revenue) OVER (
        ORDER BY sale_date
    ) AS previous_day_revenue

FROM daily_sales
ORDER BY sale_date;
```

One query now provides:

``` text
daily revenue
running revenue
7-row moving average
previous-day revenue
```

This is the practical power of window functions.

------------------------------------------------------------------------

# 11.81 Practical Example: Customer Analytics

``` sql
WITH customer_orders AS (
    SELECT
        customer_id,
        order_id,
        order_date,
        amount,

        ROW_NUMBER() OVER (
            PARTITION BY customer_id
            ORDER BY order_date DESC, order_id DESC
        ) AS latest_rank,

        SUM(amount) OVER (
            PARTITION BY customer_id
        ) AS customer_total,

        LAG(amount) OVER (
            PARTITION BY customer_id
            ORDER BY order_date, order_id
        ) AS previous_order_amount

    FROM orders
)
SELECT *
FROM customer_orders;
```

This provides:

-   latest-order ranking
-   customer lifetime/order total
-   previous-order comparison

without collapsing the order-level rows.

------------------------------------------------------------------------

# 11.82 Practical Example: Department Salary Analysis

``` sql
SELECT
    employee_id,
    name,
    department,
    salary,

    RANK() OVER (
        PARTITION BY department
        ORDER BY salary DESC
    ) AS salary_rank,

    AVG(salary) OVER (
        PARTITION BY department
    ) AS department_avg,

    salary
        - AVG(salary) OVER (
            PARTITION BY department
        ) AS difference_from_avg,

    salary
        / SUM(salary) OVER (
            PARTITION BY department
        ) * 100 AS salary_share

FROM employees;
```

This produces a rich analytical dataset while keeping one row per
employee.

------------------------------------------------------------------------

# 11.83 Performance Considerations

Window functions can be computationally expensive because the database
may need to:

-   partition rows
-   sort rows
-   maintain frame state
-   compute analytical values

Performance often depends heavily on:

``` text
PARTITION BY
ORDER BY
data volume
cardinality
indexes
sort operations
memory
```

------------------------------------------------------------------------

# 11.84 Why `ORDER BY` Can Be Expensive

Consider:

``` sql
ROW_NUMBER() OVER (
    PARTITION BY department
    ORDER BY salary DESC
)
```

The database may need to establish the required ordering inside each
department.

For millions of rows, this can involve substantial sorting or related
execution work.

Indexes can sometimes help, but whether they do depends on:

-   database engine
-   query plan
-   existing filters
-   index structure
-   physical data organization

Do not assume an index automatically eliminates all window-function
sorting.

------------------------------------------------------------------------

# 11.85 Filtering Before Windowing

Suppose you only care about 2026 data.

Prefer filtering early when semantically valid:

``` sql
SELECT
    ...,
    ROW_NUMBER() OVER (...)
FROM orders
WHERE order_date >= DATE '2026-01-01';
```

The window operates over the rows surviving the `WHERE` clause.

This can substantially reduce the data the window operation must
process.

But always verify that filtering before the window does not change the
intended analytical population.

------------------------------------------------------------------------

# 11.86 Window Functions and Indexes

Suppose:

``` sql
ROW_NUMBER() OVER (
    PARTITION BY customer_id
    ORDER BY order_date
)
```

A potentially relevant index might begin with:

``` text
(customer_id, order_date)
```

But whether that index materially improves performance depends on the
optimizer and database engine.

A window query is not automatically "indexed" just because an index
exists on the partition column.

------------------------------------------------------------------------

# 11.87 Avoid Unnecessary Window Calculations

This:

``` sql
SELECT
    *,
    SUM(amount) OVER () AS total,
    AVG(amount) OVER () AS average,
    MAX(amount) OVER () AS maximum
FROM huge_table;
```

may be perfectly reasonable if those values are needed.

But avoid calculating analytical columns that are never used downstream.

Window functions can be expensive on large datasets.

------------------------------------------------------------------------

# 11.88 Common Mistake: Confusing `RANK` with `ROW_NUMBER`

Wrong assumption:

``` text
ROW_NUMBER = ranking with ties
```

No.

`ROW_NUMBER()` always assigns distinct sequence numbers.

For:

``` text
100
100
90
```

you can get:

``` text
ROW_NUMBER:
1
2
3
```

while:

``` text
RANK:
1
1
3
```

and:

``` text
DENSE_RANK:
1
1
2
```

------------------------------------------------------------------------

# 11.89 Common Mistake: Filtering a Window Alias in `WHERE`

Wrong:

``` sql
SELECT
    *,
    ROW_NUMBER() OVER (
        PARTITION BY customer_id
        ORDER BY order_date DESC
    ) AS rn
FROM orders
WHERE rn = 1;
```

Use:

``` sql
WITH ranked AS (
    SELECT
        *,
        ROW_NUMBER() OVER (
            PARTITION BY customer_id
            ORDER BY order_date DESC
        ) AS rn
    FROM orders
)
SELECT *
FROM ranked
WHERE rn = 1;
```

Or use `QUALIFY` where supported.

------------------------------------------------------------------------

# 11.90 Common Mistake: Assuming `LAST_VALUE()` Means Last Row

This:

``` sql
LAST_VALUE(value) OVER (
    ORDER BY date
)
```

may return the current row's value or another frame-relative value
rather than the final partition value.

If you need the final value across the full partition, explicitly define
the frame:

``` sql
ROWS BETWEEN UNBOUNDED PRECEDING
         AND UNBOUNDED FOLLOWING
```

------------------------------------------------------------------------

# 11.91 Common Mistake: Ignoring Ties

This:

``` sql
ROW_NUMBER() OVER (
    ORDER BY salary DESC
)
```

does not specify a deterministic tie-breaker.

Better:

``` sql
ROW_NUMBER() OVER (
    ORDER BY salary DESC, employee_id
)
```

------------------------------------------------------------------------

# 11.92 Common Mistake: Assuming `ROWS` Means Time

``` sql
ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
```

means:

``` text
previous six rows
```

not necessarily:

``` text
previous six days
```

Data grain matters.

------------------------------------------------------------------------

# 11.93 Common Mistake: Using `GROUP BY` When Detail Rows Are Needed

Suppose you need:

``` text
order-level data
+
customer total
```

This:

``` sql
GROUP BY customer_id
```

may collapse the orders.

Instead:

``` sql
SUM(amount) OVER (
    PARTITION BY customer_id
)
```

preserves the order rows.

------------------------------------------------------------------------

# 11.94 Common Mistake: Using a Window Function When a Simple Aggregate Is Better

If you only need:

``` text
one row per department
```

then:

``` sql
GROUP BY department
```

is usually clearer than:

``` sql
SELECT DISTINCT department,
       AVG(salary) OVER (PARTITION BY department)
FROM employees;
```

Use the tool matching the desired result grain.

------------------------------------------------------------------------

# 11.95 Window Functions and Query Grain

Always ask:

> What does one output row represent?

Examples:

``` text
one row = employee
one row = order
one row = customer
one row = day
```

Then ask:

> Do I want to preserve that grain?

If yes:

``` text
window function
```

may be appropriate.

If no:

``` text
GROUP BY
```

may be appropriate.

This "grain-first" thinking prevents many SQL mistakes.

------------------------------------------------------------------------

# 11.96 Window Functions and NULL Ordering

The placement of `NULL` values in:

``` sql
ORDER BY
```

can vary by database.

Some systems support explicit syntax such as:

``` sql
NULLS FIRST
```

or:

``` sql
NULLS LAST
```

For portable analytical logic, be aware that null ordering is
dialect-sensitive.

This matters for:

``` text
ROW_NUMBER
RANK
LAG
LEAD
FIRST_VALUE
LAST_VALUE
```

------------------------------------------------------------------------

# 11.97 Advanced Pattern: Top N with Ties

If you need:

> top 3 salary ranks per department, including ties

consider:

``` sql
WITH ranked AS (
    SELECT
        *,
        DENSE_RANK() OVER (
            PARTITION BY department
            ORDER BY salary DESC
        ) AS rnk
    FROM employees
)
SELECT *
FROM ranked
WHERE rnk <= 3;
```

This is different from:

``` sql
ROW_NUMBER() <= 3
```

because `ROW_NUMBER()` may arbitrarily select among tied rows.

------------------------------------------------------------------------

# 11.98 Advanced Pattern: First Event Per User

``` sql
WITH ranked AS (
    SELECT
        user_id,
        event_id,
        event_time,
        ROW_NUMBER() OVER (
            PARTITION BY user_id
            ORDER BY event_time, event_id
        ) AS rn
    FROM events
)
SELECT *
FROM ranked
WHERE rn = 1;
```

------------------------------------------------------------------------

# 11.99 Advanced Pattern: Previous Status

``` sql
SELECT
    user_id,
    event_time,
    status,
    LAG(status) OVER (
        PARTITION BY user_id
        ORDER BY event_time
    ) AS previous_status
FROM user_status_history;
```

------------------------------------------------------------------------

# 11.100 Advanced Pattern: Next Event

``` sql
SELECT
    user_id,
    event_time,
    LEAD(event_time) OVER (
        PARTITION BY user_id
        ORDER BY event_time
    ) AS next_event_time
FROM events;
```

Then:

``` text
next_event_time - event_time
```

can be used to measure time until the next event.

------------------------------------------------------------------------

# 11.101 Advanced Pattern: Running Balance

For transactions:

``` text
credit
debit
```

calculate:

``` sql
SUM(
    CASE
        WHEN transaction_type = 'credit' THEN amount
        WHEN transaction_type = 'debit' THEN -amount
    END
) OVER (
    PARTITION BY account_id
    ORDER BY transaction_time, transaction_id
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
) AS running_balance
```

This is a common financial-analysis pattern.

------------------------------------------------------------------------

# 11.102 Advanced Pattern: Running Percentage

Conceptually:

``` sql
SUM(amount) OVER (
    ORDER BY event_time
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
/
SUM(amount) OVER ()
* 100
```

This tells you what percentage of the final total has accumulated by
each row.

------------------------------------------------------------------------

# 11.103 Advanced Pattern: Contribution Within Group

``` sql
SELECT
    department,
    employee_id,
    salary,
    salary
        / SUM(salary) OVER (
            PARTITION BY department
        ) AS contribution_ratio
FROM employees;
```

This is common in:

-   financial reporting
-   sales analysis
-   resource allocation
-   portfolio analysis

------------------------------------------------------------------------

# 11.104 Advanced Pattern: Compare to Group Average

``` sql
SELECT
    employee_id,
    department,
    salary,
    AVG(salary) OVER (
        PARTITION BY department
    ) AS avg_salary,
    salary
        - AVG(salary) OVER (
            PARTITION BY department
        ) AS variance
FROM employees;
```

------------------------------------------------------------------------

# 11.105 Advanced Pattern: Detect New Groups

Suppose a new session starts when the gap exceeds 30 minutes.

Conceptually:

``` sql
WITH events_with_previous AS (
    SELECT
        user_id,
        event_time,
        LAG(event_time) OVER (
            PARTITION BY user_id
            ORDER BY event_time
        ) AS previous_time
    FROM events
)
SELECT
    *,
    CASE
        WHEN previous_time IS NULL THEN 1
        WHEN event_time - previous_time > INTERVAL '30 minutes'
            THEN 1
        ELSE 0
    END AS new_session
FROM events_with_previous;
```

A second window calculation can then cumulatively sum `new_session` to
assign session IDs.

Exact interval syntax varies by database.

------------------------------------------------------------------------

# 11.106 Two-Stage Window Analysis

Complex analytical SQL often follows:

``` text
Stage 1
    calculate LAG / ROW_NUMBER / etc.
        ↓
Stage 2
    derive flags or metrics
        ↓
Stage 3
    apply another window
```

Example concept:

``` text
events
  ↓
LAG()
  ↓
identify session boundaries
  ↓
SUM(flag) OVER (...)
  ↓
session_id
```

This is a powerful pattern for event analytics.

------------------------------------------------------------------------

# 11.107 Window Function Execution Mental Model

A useful conceptual model:

``` text
FROM / JOIN
      ↓
filter rows
      ↓
group/aggregate where applicable
      ↓
produce the logical row set
      ↓
partition rows
      ↓
order rows within windows
      ↓
apply window function/frame
      ↓
produce annotated rows
      ↓
final ORDER BY
```

The physical database engine may implement this very differently.

The purpose of the model is to understand semantics.

------------------------------------------------------------------------

# 11.108 The Three Main Window Controls

Most window expressions can be understood through:

``` text
PARTITION BY
    Which group?

ORDER BY
    In what logical sequence?

FRAME
    Which portion of that ordered partition?
```

For example:

``` sql
SUM(amount) OVER (
    PARTITION BY customer_id
    ORDER BY order_date
    ROWS BETWEEN 6 PRECEDING
             AND CURRENT ROW
)
```

means:

``` text
Which group?
    → customer

What order?
    → order date

Which rows?
    → current + previous 6 rows
```

This is the most useful way to reason about complex windows.

------------------------------------------------------------------------

# 11.109 Window Function Cheat Sheet

``` text
ROW_NUMBER()
    unique sequential number

RANK()
    ties share rank
    gaps appear

DENSE_RANK()
    ties share rank
    no gaps

NTILE(n)
    divide rows into n buckets

LAG()
    previous row

LEAD()
    next row

FIRST_VALUE()
    first value in frame

LAST_VALUE()
    last value in frame

SUM() OVER(...)
    window sum

AVG() OVER(...)
    window average

COUNT() OVER(...)
    window count

MIN() OVER(...)
    window minimum

MAX() OVER(...)
    window maximum
```

------------------------------------------------------------------------

# 11.110 Ranking Cheat Sheet

For:

``` text
100
90
90
80
```

``` text
ROW_NUMBER   RANK   DENSE_RANK
-----------  -----  ----------
1            1      1
2            2      2
3            2      2
4            4      3
```

Remember:

``` text
ROW_NUMBER → unique positions

RANK       → competition ranking

DENSE_RANK → compact ranking
```

------------------------------------------------------------------------

# 11.111 Navigation Cheat Sheet

``` text
LAG(x)
    previous x

LAG(x, 2)
    x from two rows earlier

LEAD(x)
    next x

LEAD(x, 2)
    x from two rows later
```

Always define the relevant ordering:

``` sql
LAG(x) OVER (
    PARTITION BY ...
    ORDER BY ...
)
```

------------------------------------------------------------------------

# 11.112 Frame Cheat Sheet

``` text
UNBOUNDED PRECEDING
    beginning

CURRENT ROW
    current position/group depending on frame semantics

n PRECEDING
    n prior rows/ranges/groups depending on frame type

n FOLLOWING
    n future rows/ranges/groups

UNBOUNDED FOLLOWING
    end
```

Common patterns:

``` sql
-- Running total
ROWS BETWEEN UNBOUNDED PRECEDING
         AND CURRENT ROW
```

``` sql
-- 7-row moving average
ROWS BETWEEN 6 PRECEDING
         AND CURRENT ROW
```

``` sql
-- Entire partition
ROWS BETWEEN UNBOUNDED PRECEDING
         AND UNBOUNDED FOLLOWING
```

``` sql
-- Centered 5-row window
ROWS BETWEEN 2 PRECEDING
         AND 2 FOLLOWING
```

------------------------------------------------------------------------

# 11.113 Most Important Window Patterns

## Pattern 1 --- Top N per group

``` sql
ROW_NUMBER() OVER (
    PARTITION BY group_col
    ORDER BY metric DESC
)
```

------------------------------------------------------------------------

## Pattern 2 --- Ranking with ties

``` sql
RANK() OVER (
    PARTITION BY group_col
    ORDER BY metric DESC
)
```

------------------------------------------------------------------------

## Pattern 3 --- Previous value

``` sql
LAG(value) OVER (
    PARTITION BY group_col
    ORDER BY time_col
)
```

------------------------------------------------------------------------

## Pattern 4 --- Next value

``` sql
LEAD(value) OVER (
    PARTITION BY group_col
    ORDER BY time_col
)
```

------------------------------------------------------------------------

## Pattern 5 --- Running total

``` sql
SUM(value) OVER (
    PARTITION BY group_col
    ORDER BY time_col
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

------------------------------------------------------------------------

## Pattern 6 --- Moving average

``` sql
AVG(value) OVER (
    ORDER BY time_col
    ROWS BETWEEN 6 PRECEDING
             AND CURRENT ROW
)
```

------------------------------------------------------------------------

## Pattern 7 --- Group total while preserving details

``` sql
SUM(value) OVER (
    PARTITION BY group_col
)
```

------------------------------------------------------------------------

## Pattern 8 --- Percentage of group

``` sql
value
/
SUM(value) OVER (
    PARTITION BY group_col
)
```

------------------------------------------------------------------------

# 11.114 Interview Questions

## Q1. What is a window function?

A function that performs a calculation across a related set of rows
while generally preserving the individual rows in the result.

------------------------------------------------------------------------

## Q2. What does `OVER()` do?

It defines the window over which a window function operates.

------------------------------------------------------------------------

## Q3. What does `PARTITION BY` do?

It divides rows into independent groups for the window calculation
without collapsing them.

------------------------------------------------------------------------

## Q4. What is the difference between `PARTITION BY` and `GROUP BY`?

`GROUP BY` collapses rows into groups.

`PARTITION BY` creates logical windows while retaining rows.

------------------------------------------------------------------------

## Q5. Difference between `ROW_NUMBER`, `RANK`, and `DENSE_RANK`?

``` text
ROW_NUMBER
    unique sequence

RANK
    ties + gaps

DENSE_RANK
    ties + no gaps
```

------------------------------------------------------------------------

## Q6. Difference between `LAG` and `LEAD`?

``` text
LAG  → previous row
LEAD → next row
```

------------------------------------------------------------------------

## Q7. What is a window frame?

The subset of the ordered partition that a frame-sensitive window
calculation uses for the current row.

------------------------------------------------------------------------

## Q8. What is a running total?

A cumulative aggregate where the frame begins at the first row and ends
at the current row.

Typical expression:

``` sql
SUM(value) OVER (
    ORDER BY date
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)
```

------------------------------------------------------------------------

## Q9. Why is `LAST_VALUE()` confusing?

Because it returns the last value in the current window frame, not
necessarily the last value in the entire partition.

------------------------------------------------------------------------

## Q10. How do you get the latest row per customer?

A common approach:

``` sql
ROW_NUMBER() OVER (
    PARTITION BY customer_id
    ORDER BY event_time DESC
)
```

then filter:

``` sql
rn = 1
```

using a CTE, derived table, or `QUALIFY` where supported.

------------------------------------------------------------------------

## Q11. Can a window function be used directly in `WHERE`?

Generally no.

Use another query layer or `QUALIFY` where supported.

------------------------------------------------------------------------

## Q12. What is `NTILE()`?

It divides ordered rows into approximately equal-sized buckets.

------------------------------------------------------------------------

## Q13. Is `ROWS BETWEEN 6 PRECEDING AND CURRENT ROW` always a 7-day window?

No.

It is a seven-row window, assuming the current row exists. Missing dates
can make it different from seven calendar days.

------------------------------------------------------------------------

## Q14. Why add a tie-breaker to `ROW_NUMBER()`?

To make the ordering deterministic when the primary ordering values are
equal.

Example:

``` sql
ORDER BY salary DESC, employee_id
```

------------------------------------------------------------------------

# 11.115 Final Decision Framework

When you encounter an analytical SQL problem, ask:

``` text
Do I need to preserve individual rows?
        │
        ├── No → GROUP BY / aggregation
        │
        └── Yes
             ↓
       Do I need ranking?
             │
             ├── unique sequence → ROW_NUMBER
             ├── ties + gaps → RANK
             └── ties + no gaps → DENSE_RANK
             ↓
       Do I need previous/next row?
             │
             ├── previous → LAG
             └── next → LEAD
             ↓
       Do I need buckets?
             │
             └── NTILE
             ↓
       Do I need first/last value?
             │
             ├── FIRST_VALUE
             └── LAST_VALUE
             ↓
       Do I need cumulative/sliding analytics?
             │
             ├── running → frame ending CURRENT ROW
             └── moving → bounded PRECEDING/FOLLOWING frame
```

------------------------------------------------------------------------

# 11.116 Final Mental Model

The most important thing to understand is not the individual functions.

It is this:

``` text
WINDOW FUNCTION
       │
       ├── PARTITION BY
       │       ↓
       │   "Which rows belong together?"
       │
       ├── ORDER BY
       │       ↓
       │   "In what sequence?"
       │
       └── FRAME
               ↓
          "Which part of that
           ordered partition matters
           for this row?"
```

Then choose the operation:

``` text
ROW_NUMBER
    → position

RANK
    → ranking with gaps

DENSE_RANK
    → ranking without gaps

NTILE
    → buckets

LAG
    → previous row

LEAD
    → next row

FIRST_VALUE
    → first value

LAST_VALUE
    → last value

SUM / AVG / COUNT / MIN / MAX
    → analytical aggregates
```

The deeper mental model is:

``` text
GROUP BY
    → reduce many rows into fewer rows

WINDOW FUNCTION
    → keep the rows and analyze their neighborhood
```

And the three questions that solve most window-function problems are:

``` text
1. What is my partition?
2. What is my ordering?
3. What is my frame?
```

Once those three are clear, the SQL expression usually becomes
straightforward.

------------------------------------------------------------------------

# 11.117 Final Compact Reference

``` sql
-- Ranking
ROW_NUMBER() OVER (
    PARTITION BY ...
    ORDER BY ...
)

RANK() OVER (
    PARTITION BY ...
    ORDER BY ...
)

DENSE_RANK() OVER (
    PARTITION BY ...
    ORDER BY ...
)

NTILE(4) OVER (
    PARTITION BY ...
    ORDER BY ...
)

-- Previous / next
LAG(value) OVER (
    PARTITION BY ...
    ORDER BY ...
)

LEAD(value) OVER (
    PARTITION BY ...
    ORDER BY ...
)

-- First / last
FIRST_VALUE(value) OVER (
    PARTITION BY ...
    ORDER BY ...
)

LAST_VALUE(value) OVER (
    PARTITION BY ...
    ORDER BY ...
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND UNBOUNDED FOLLOWING
)

-- Running total
SUM(value) OVER (
    PARTITION BY ...
    ORDER BY ...
    ROWS BETWEEN UNBOUNDED PRECEDING
             AND CURRENT ROW
)

-- Moving average
AVG(value) OVER (
    PARTITION BY ...
    ORDER BY ...
    ROWS BETWEEN 6 PRECEDING
             AND CURRENT ROW
)
```

------------------------------------------------------------------------

# Conclusion

Window functions transform SQL from a language that primarily retrieves
and aggregates data into a powerful analytical language.

The essential concepts are:

``` text
OVER()
    ↓
defines the window

PARTITION BY
    ↓
separates independent groups

ORDER BY
    ↓
defines analytical sequence

FRAME
    ↓
defines the relevant subset of the ordered window
```

Then:

``` text
ROW_NUMBER → unique position
RANK       → ranking with gaps
DENSE_RANK → ranking without gaps
NTILE      → buckets
LAG        → previous row
LEAD       → next row
FIRST_VALUE→ first value
LAST_VALUE → last value
SUM        → running/cumulative totals
AVG        → moving averages
```

The key practical principle is:

> **Use `GROUP BY` when you want to collapse rows; use window functions
> when you want to analyze rows while keeping their detail.**

That distinction, combined with a strong understanding of
`PARTITION BY`, `ORDER BY`, and window frames, is the foundation for
advanced SQL analytics.
