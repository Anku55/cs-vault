# SQL Constraints --- Deep Practical & Theoretical Notes

> Section 12 of the SQL notes series. A comprehensive guide to database
> constraints, keys, integrity rules, referential actions, constraint
> violations, schema design, practical SQL, debugging, and
> interview-level distinctions.

------------------------------------------------------------------------

# 12. Constraints

SQL constraints are rules enforced by the database to keep data valid,
consistent, and structurally correct.

Without constraints, an application could accidentally insert:

-   duplicate identifiers
-   missing required values
-   invalid domain values
-   references to nonexistent records
-   inconsistent relationships
-   values that violate business rules

Constraints move part of data validation from application code into the
database.

Core constraints:

``` text
PRIMARY KEY
FOREIGN KEY
UNIQUE
NOT NULL
CHECK
DEFAULT
```

Foreign-key actions additionally control what happens when referenced
rows are updated or deleted:

``` text
ON DELETE CASCADE
ON DELETE SET NULL
ON DELETE RESTRICT
ON UPDATE CASCADE
```

------------------------------------------------------------------------

# 12.1 Why Constraints Exist

Suppose we have:

``` text
customers
-------------------------
customer_id | name
-------------------------
1           | Alice
2           | Bob
```

and:

``` text
orders
--------------------------------
order_id | customer_id | amount
--------------------------------
101      | 1           | 500
102      | 2           | 300
```

The database should prevent:

``` text
customer_id = 999
```

in `orders` if customer `999` does not exist.

Similarly, it should prevent:

``` text
customer_id = NULL
```

if every order must belong to a customer.

And it should prevent:

``` text
customer_id = 1
```

from appearing twice in a table intended to identify customers uniquely.

Constraints formalize these rules.

------------------------------------------------------------------------

# 12.2 Constraint vs Validation

Application validation:

``` text
application
    ↓
validate
    ↓
database
```

Database constraint:

``` text
application
    ↓
database
    ↓
constraint enforcement
```

Application validation is useful for:

-   friendly error messages
-   UI validation
-   early feedback
-   business workflow

Database constraints are essential for:

-   protecting data integrity
-   preventing invalid writes from any client
-   enforcing invariants centrally
-   protecting against bugs in multiple applications

A robust system commonly uses both.

------------------------------------------------------------------------

# 12.3 Main Types of Constraints

  Constraint      Main purpose
  --------------- ---------------------------------------
  `PRIMARY KEY`   uniquely identifies each row
  `FOREIGN KEY`   enforces relationships between tables
  `UNIQUE`        prevents duplicate values
  `NOT NULL`      requires a value
  `CHECK`         enforces a Boolean condition
  `DEFAULT`       supplies a value when one is omitted

These constraints can be combined.

Example:

``` sql
CREATE TABLE employees (
    employee_id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    salary DECIMAL(12,2) CHECK (salary >= 0),
    status VARCHAR(20) DEFAULT 'active'
);
```

------------------------------------------------------------------------

# 12.4 PRIMARY KEY

A primary key uniquely identifies each row in a table.

Example:

``` sql
CREATE TABLE customers (
    customer_id INTEGER PRIMARY KEY,
    name VARCHAR(100)
);
```

Here:

``` text
customer_id
```

is the primary key.

It provides an identity for each row.

------------------------------------------------------------------------

# 12.5 Properties of a Primary Key

A primary key must provide uniqueness and non-null identity.

Conceptually:

``` text
PRIMARY KEY
    =
UNIQUE
+
NOT NULL
```

For example:

``` text
customer_id
-----------
1
2
3
```

is valid.

But:

``` text
1
2
2
```

violates uniqueness.

And:

``` text
1
NULL
3
```

violates the primary-key requirement.

A table has one primary-key constraint, although that key may contain
multiple columns.

------------------------------------------------------------------------

# 12.6 Primary Key Syntax

Column-level:

``` sql
CREATE TABLE customers (
    customer_id INTEGER PRIMARY KEY,
    name VARCHAR(100)
);
```

Table-level:

``` sql
CREATE TABLE customers (
    customer_id INTEGER,
    name VARCHAR(100),
    CONSTRAINT pk_customers
        PRIMARY KEY (customer_id)
);
```

Table-level syntax is particularly useful for composite keys and
explicit constraint naming.

------------------------------------------------------------------------

# 12.7 Composite Primary Key

A primary key can consist of multiple columns.

Example:

``` sql
CREATE TABLE order_items (
    order_id INTEGER,
    product_id INTEGER,
    quantity INTEGER,

    PRIMARY KEY (order_id, product_id)
);
```

The combination must be unique.

This means:

``` text
(100, 10)
(100, 11)
(101, 10)
```

is valid.

But:

``` text
(100, 10)
(100, 10)
```

is not.

------------------------------------------------------------------------

# 12.8 Composite Key Meaning

With:

``` sql
PRIMARY KEY (order_id, product_id)
```

neither column necessarily needs to be globally unique.

Instead:

``` text
(order_id, product_id)
```

as a pair must be unique.

This is common in:

-   order-item tables
-   enrollment tables
-   many-to-many relationships
-   association tables

------------------------------------------------------------------------

# 12.9 Surrogate vs Natural Primary Keys

A primary key can be:

### Natural key

A real-world attribute:

``` text
email
ISBN
national identifier
```

### Surrogate key

An artificial identifier:

``` text
customer_id
employee_id
product_id
```

Example:

``` sql
customer_id BIGINT PRIMARY KEY
```

Surrogate keys are common because they tend to be:

-   compact
-   stable
-   independent of business meaning
-   convenient for foreign keys

But the correct choice depends on the domain.

------------------------------------------------------------------------

# 12.10 Auto-Generated Primary Keys

Databases provide different mechanisms:

``` text
IDENTITY
AUTO_INCREMENT
GENERATED AS IDENTITY
SEQUENCE
SERIAL
```

depending on the SQL dialect.

Example in a dialect supporting identity columns:

``` sql
CREATE TABLE customers (
    customer_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

The exact syntax is database-specific.

------------------------------------------------------------------------

# 12.11 FOREIGN KEY

A foreign key creates a relationship between a child table and a parent
table.

Example:

``` sql
CREATE TABLE customers (
    customer_id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE orders (
    order_id INTEGER PRIMARY KEY,
    customer_id INTEGER,
    amount DECIMAL(12,2),

    FOREIGN KEY (customer_id)
        REFERENCES customers(customer_id)
);
```

Here:

``` text
customers.customer_id
        ↑
        │ referenced key
        │
orders.customer_id
        │
        ↓
foreign key
```

------------------------------------------------------------------------

# 12.12 Parent and Child Tables

Foreign-key terminology:

``` text
Parent table
    ↓
contains referenced key

Child table
    ↓
contains foreign key
```

Example:

``` text
customers
    customer_id  ← parent key

orders
    customer_id  ← child foreign key
```

The child row cannot normally reference a parent key that does not
exist.

------------------------------------------------------------------------

# 12.13 Referential Integrity

Referential integrity means references between tables remain valid.

If:

``` text
orders.customer_id
```

references:

``` text
customers.customer_id
```

then a non-null value such as:

``` text
customer_id = 42
```

must correspond to an allowed referenced parent row.

Therefore:

``` sql
INSERT INTO orders (order_id, customer_id, amount)
VALUES (1001, 42, 500);
```

fails if customer `42` does not exist and the foreign-key constraint
requires a matching parent.

------------------------------------------------------------------------

# 12.14 Referential Integrity Example

Parent:

``` text
customers
----------------
customer_id
----------------
1
2
3
```

Child:

``` text
orders
----------------
order_id | customer_id
----------------
101      | 1
102      | 2
103      | 3
```

This is valid.

But:

``` text
104 | 999
```

is invalid if `999` is not present in the referenced parent key.

------------------------------------------------------------------------

# 12.15 Foreign Keys Can Reference UNIQUE Keys

A foreign key does not necessarily have to reference the primary key.

It can reference a suitable unique key, depending on the database and
constraints.

Example:

``` sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);
```

Another table may reference:

``` text
users.email
```

if the database permits it and the referenced column satisfies the
required uniqueness rules.

In relational design, the referenced columns must identify at most one
parent row.

------------------------------------------------------------------------

# 12.16 Foreign Key Columns and NULL

A foreign key may often contain `NULL` unless it is also declared:

``` sql
NOT NULL
```

Example:

``` sql
manager_id INTEGER,
FOREIGN KEY (manager_id)
    REFERENCES employees(employee_id)
```

Then:

``` text
manager_id = NULL
```

can represent:

``` text
no manager
```

This does not violate the foreign-key relationship because there is no
non-null parent value being claimed.

If the relationship is mandatory:

``` sql
manager_id INTEGER NOT NULL
```

then `NULL` is prohibited.

------------------------------------------------------------------------

# 12.17 PRIMARY KEY vs FOREIGN KEY

  Feature                Primary Key         Foreign Key
  ---------------------- ------------------- ------------------------
  Purpose                identifies row      references another row
  Uniqueness             required            not generally required
  NULL                   not allowed         often allowed
  Number per table       one PK constraint   multiple FKs possible
  Creates relationship   not by itself       yes
  Main concept           identity            referential integrity

Example:

``` text
customers
customer_id PK

orders
order_id     PK
customer_id  FK
```

------------------------------------------------------------------------

# 12.18 UNIQUE Constraint

`UNIQUE` prevents duplicate values in a constrained column or column
combination.

Example:

``` sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE
);
```

This prevents two rows from having the same email under the database's
uniqueness semantics.

------------------------------------------------------------------------

# 12.19 UNIQUE vs PRIMARY KEY

Both enforce uniqueness, but their purposes differ.

``` text
PRIMARY KEY
    → row identity

UNIQUE
    → alternate uniqueness rule
```

Example:

``` sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    username VARCHAR(50) UNIQUE
);
```

Here:

``` text
user_id
    → primary identity

email
    → unique business attribute

username
    → another unique attribute
```

------------------------------------------------------------------------

# 12.20 Multiple UNIQUE Constraints

A table can have multiple unique constraints.

Example:

``` sql
CREATE TABLE employees (
    employee_id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    employee_code VARCHAR(50) UNIQUE
);
```

The table has:

``` text
1 primary key
2 unique constraints
```

This is perfectly normal.

------------------------------------------------------------------------

# 12.21 Composite UNIQUE Constraint

Example:

``` sql
CREATE TABLE enrollments (
    student_id INTEGER,
    course_id INTEGER,

    UNIQUE (student_id, course_id)
);
```

This prevents duplicate enrollment of the same student in the same
course.

But:

``` text
student 1 + course 10
student 1 + course 11
student 2 + course 10
```

are valid.

------------------------------------------------------------------------

# 12.22 UNIQUE and NULL

`NULL` handling for `UNIQUE` constraints is database-dependent.

Some databases allow multiple `NULL` values because `NULL` represents an
unknown/non-value and is not treated as equal to another `NULL`.

Other systems or configurations may behave differently.

Therefore, if the business rule is:

> every row must have a unique non-null email

use:

``` sql
email VARCHAR(255) NOT NULL UNIQUE
```

rather than relying on `UNIQUE` alone.

------------------------------------------------------------------------

# 12.23 NOT NULL

`NOT NULL` requires a column to contain a value.

Example:

``` sql
CREATE TABLE employees (
    employee_id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

This is invalid:

``` sql
INSERT INTO employees (employee_id, name)
VALUES (1, NULL);
```

because `name` cannot be null.

------------------------------------------------------------------------

# 12.24 NULL Does Not Mean Empty String

These are different:

``` text
NULL
''
'   '
0
```

`NULL` means the value is absent/unknown according to SQL's three-valued
logic.

An empty string is a string value.

Zero is a numeric value.

Do not use:

``` text
NULL = ''
```

as a conceptual equivalence.

------------------------------------------------------------------------

# 12.25 CHECK Constraint

A `CHECK` constraint enforces a Boolean condition.

Example:

``` sql
CREATE TABLE products (
    product_id INTEGER PRIMARY KEY,
    price DECIMAL(12,2),
    CHECK (price >= 0)
);
```

This prevents negative prices under normal constraint semantics.

------------------------------------------------------------------------

# 12.26 More CHECK Examples

Age:

``` sql
CHECK (age >= 18)
```

Quantity:

``` sql
CHECK (quantity > 0)
```

Status:

``` sql
CHECK (status IN ('active', 'inactive'))
```

Salary:

``` sql
CHECK (salary >= 0)
```

Date relationship:

``` sql
CHECK (end_date >= start_date)
```

The exact behavior of `CHECK` expressions involving `NULL` follows SQL's
three-valued logic and database semantics, so `NOT NULL` may also be
required when the value itself must exist.

------------------------------------------------------------------------

# 12.27 CHECK Is Not a Full Business Rule Engine

A `CHECK` constraint is excellent for row-level conditions such as:

``` text
price >= 0
quantity > 0
end_date >= start_date
```

But complex cross-row or cross-table rules may require:

-   foreign keys
-   unique constraints
-   triggers
-   indexed/generated structures
-   application logic
-   stored procedures
-   transactions

depending on the rule and database.

Do not force every business rule into a `CHECK`.

------------------------------------------------------------------------

# 12.28 DEFAULT

A `DEFAULT` supplies a value when an insert omits the column.

Example:

``` sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    status VARCHAR(20) DEFAULT 'active'
);
```

Then:

``` sql
INSERT INTO users (user_id)
VALUES (1);
```

can result in:

``` text
user_id | status
--------|--------
1       | active
```

------------------------------------------------------------------------

# 12.29 DEFAULT Does Not Mean NOT NULL

These are different:

``` sql
status VARCHAR(20) DEFAULT 'active'
```

and:

``` sql
status VARCHAR(20) NOT NULL DEFAULT 'active'
```

The first provides a default when the column is omitted, but may still
permit explicit `NULL` if the database semantics allow it.

The second both:

``` text
provides a default
+
rejects NULL
```

when the `NOT NULL` constraint applies.

------------------------------------------------------------------------

# 12.30 DEFAULT and Explicit NULL

A common conceptual mistake is assuming:

``` sql
INSERT INTO users (user_id, status)
VALUES (1, NULL);
```

will use the default.

Usually, an explicit `NULL` is a request to store `NULL`, not an
instruction to use the default.

The default generally applies when the column is omitted from the insert
or when database-specific default syntax is used.

------------------------------------------------------------------------

# 12.31 Constraint Naming

Constraints should often be explicitly named in production schemas.

Example:

``` sql
CREATE TABLE employees (
    employee_id INTEGER,
    email VARCHAR(255) NOT NULL,

    CONSTRAINT pk_employees
        PRIMARY KEY (employee_id),

    CONSTRAINT uq_employees_email
        UNIQUE (email),

    CONSTRAINT chk_employees_email
        CHECK (email <> '')
);
```

Named constraints make errors and schema maintenance easier.

------------------------------------------------------------------------

# 12.32 Foreign Key Naming

Example:

``` sql
CREATE TABLE orders (
    order_id INTEGER PRIMARY KEY,
    customer_id INTEGER,

    CONSTRAINT fk_orders_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(customer_id)
);
```

A naming convention might be:

``` text
pk_<table>
fk_<child>_<parent>
uq_<table>_<column>
chk_<table>_<rule>
```

Exact conventions are project-specific.

------------------------------------------------------------------------

# 12.33 ON DELETE Actions

A foreign key can define what happens when the referenced parent row is
deleted.

Common options:

``` text
ON DELETE CASCADE
ON DELETE SET NULL
ON DELETE RESTRICT
```

The choice represents a data-modeling decision.

------------------------------------------------------------------------

# 12.34 ON DELETE CASCADE

Example:

``` sql
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON DELETE CASCADE
```

Meaning:

> If the parent customer is deleted, related child rows are
> automatically deleted.

Conceptually:

``` text
customer
   │
   ├── order 101
   ├── order 102
   └── order 103

DELETE customer
       ↓
delete related orders
```

------------------------------------------------------------------------

# 12.35 When CASCADE Makes Sense

Good use cases can include tightly owned dependent data.

Examples:

``` text
user
 └── temporary preferences

order
 └── order_items
```

If the child has no meaningful independent existence without the parent,
cascade may be appropriate.

------------------------------------------------------------------------

# 12.36 CASCADE Danger

Cascade deletion can be powerful and dangerous.

A single:

``` sql
DELETE FROM customers
WHERE customer_id = 1;
```

could trigger deletion of:

``` text
customer
   ↓
orders
   ↓
order_items
   ↓
other dependent records
```

depending on the schema.

Always understand the dependency graph before enabling cascading
deletes.

------------------------------------------------------------------------

# 12.37 ON DELETE SET NULL

Example:

``` sql
FOREIGN KEY (manager_id)
REFERENCES employees(employee_id)
ON DELETE SET NULL
```

Meaning:

> If the referenced employee is deleted, set the child foreign-key value
> to `NULL`.

Example:

``` text
Before:

employee_id | manager_id
------------|-----------
10          | 5
11          | 5
12          | 7
```

Delete employee `5`.

After:

``` text
employee_id | manager_id
------------|-----------
10          | NULL
11          | NULL
12          | 7
```

The child rows survive.

------------------------------------------------------------------------

# 12.38 SET NULL Requires Nullable Foreign Key

This is an important requirement.

If:

``` sql
manager_id INTEGER NOT NULL
```

then:

``` sql
ON DELETE SET NULL
```

cannot successfully set `manager_id` to `NULL`.

The schema is logically inconsistent.

Use:

``` sql
manager_id INTEGER
```

if `NULL` is the intended post-delete state.

------------------------------------------------------------------------

# 12.39 When SET NULL Makes Sense

Use `SET NULL` when the relationship is optional and the child should
survive without the parent.

Examples:

``` text
employee → manager
article → author
ticket → assigned_agent
record → deleted_owner
```

The exact business meaning matters.

------------------------------------------------------------------------

# 12.40 ON DELETE RESTRICT

`RESTRICT` prevents deletion of the parent while dependent child rows
exist.

Conceptually:

``` text
customer
   ↓
orders exist

DELETE customer
       ↓
REJECT
```

The database protects the relationship.

Example:

``` sql
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON DELETE RESTRICT
```

Exact default behavior and support can vary by database.

------------------------------------------------------------------------

# 12.41 RESTRICT vs CASCADE

``` text
CASCADE
    delete parent
        ↓
    delete children

RESTRICT
    delete parent
        ↓
    reject if children exist
```

Use `RESTRICT` when deleting the parent should require explicit cleanup
or reassignment of dependent records.

------------------------------------------------------------------------

# 12.42 ON DELETE Action Comparison

  Action       Parent deleted           Child row
  ------------ ------------------------ -------------------
  `CASCADE`    allowed                  deleted
  `SET NULL`   allowed                  FK becomes `NULL`
  `RESTRICT`   rejected if referenced   unchanged

Conceptually:

``` text
CASCADE
    "delete dependents too"

SET NULL
    "keep dependents, remove relationship"

RESTRICT
    "do not delete parent while relationship exists"
```

------------------------------------------------------------------------

# 12.43 ON UPDATE CASCADE

Foreign keys can also define what happens when a referenced key changes.

Example:

``` sql
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON UPDATE CASCADE
```

If:

``` text
customers.customer_id
```

changes from:

``` text
100 → 200
```

the related foreign-key values in child rows can automatically change:

``` text
orders.customer_id
100 → 200
```

------------------------------------------------------------------------

# 12.44 Why ON UPDATE CASCADE Is Less Common

Primary keys are often designed to be stable.

For example:

``` text
customer_id = 12345
```

normally should not change.

Therefore:

``` text
ON UPDATE CASCADE
```

is less frequently needed than:

``` text
ON DELETE CASCADE
```

Still, it can be useful when referenced keys are mutable and the
database supports the desired action.

------------------------------------------------------------------------

# 12.45 ON UPDATE CASCADE Example

Parent:

``` text
customers
----------------
customer_id | name
----------------
10          | Alice
```

Child:

``` text
orders
----------------
order_id | customer_id
----------------
1001     | 10
1002     | 10
```

Update:

``` sql
UPDATE customers
SET customer_id = 20
WHERE customer_id = 10;
```

With:

``` sql
ON UPDATE CASCADE
```

the child rows become:

``` text
order_id | customer_id
---------|------------
1001     | 20
1002     | 20
```

------------------------------------------------------------------------

# 12.46 Referential Actions Are Data-Model Decisions

Do not choose:

``` text
CASCADE
SET NULL
RESTRICT
```

simply because one is convenient.

Ask:

> What should happen to the child entity if the parent disappears?

### If child has no independent meaning:

``` text
CASCADE
```

### If child survives but relationship should disappear:

``` text
SET NULL
```

### If parent deletion must be explicitly controlled:

``` text
RESTRICT
```

------------------------------------------------------------------------

# 12.47 Constraint Violations

A constraint violation occurs when a database operation attempts to
create a state that violates a defined constraint.

Examples:

``` text
duplicate primary key
duplicate unique value
NULL in NOT NULL column
invalid CHECK condition
missing foreign-key parent
invalid cascading action
```

The database rejects the operation rather than storing invalid data.

------------------------------------------------------------------------

# 12.48 Primary Key Violation

Suppose:

``` sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY
);
```

Existing:

``` text
user_id
--------
1
```

Then:

``` sql
INSERT INTO users (user_id)
VALUES (1);
```

violates the primary key.

Reason:

``` text
1 already identifies an existing row
```

------------------------------------------------------------------------

# 12.49 UNIQUE Violation

Given:

``` sql
email VARCHAR(255) UNIQUE
```

and existing:

``` text
alice@example.com
```

this fails:

``` sql
INSERT INTO users (email)
VALUES ('alice@example.com');
```

The exact error class/message is database-specific, but the semantic
problem is a uniqueness violation.

------------------------------------------------------------------------

# 12.50 NOT NULL Violation

Given:

``` sql
name VARCHAR(100) NOT NULL
```

this fails:

``` sql
INSERT INTO employees (name)
VALUES (NULL);
```

because the column requires a value.

------------------------------------------------------------------------

# 12.51 CHECK Violation

Given:

``` sql
CHECK (salary >= 0)
```

this fails:

``` sql
INSERT INTO employees (salary)
VALUES (-5000);
```

because:

``` text
-5000 >= 0
```

is false.

------------------------------------------------------------------------

# 12.52 FOREIGN KEY Violation

Given:

``` sql
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
```

if customer `999` does not exist:

``` sql
INSERT INTO orders (
    order_id,
    customer_id
)
VALUES (
    100,
    999
);
```

fails.

This protects referential integrity.

------------------------------------------------------------------------

# 12.53 Delete Constraint Violation

Suppose:

``` text
customers
    ↓
orders
```

and the foreign key uses restrictive delete behavior.

If orders exist:

``` sql
DELETE FROM customers
WHERE customer_id = 1;
```

may fail because child rows still reference customer `1`.

The correct action may be to:

-   delete child rows
-   reassign child rows
-   use an appropriate cascade policy
-   use soft deletion
-   remove the parent only after dependencies are resolved

------------------------------------------------------------------------

# 12.54 Constraint Enforcement Is Usually Statement-Level

When an operation violates a constraint, the database normally rejects
the operation.

For example:

``` sql
UPDATE employees
SET salary = -1000
WHERE employee_id = 10;
```

If:

``` sql
CHECK (salary >= 0)
```

exists, the update cannot create the invalid state.

Exact transactional behavior and timing can vary by database, including
support for deferred constraints.

------------------------------------------------------------------------

# 12.55 Immediate vs Deferred Constraints

Some database systems support deferred constraint checking.

Conceptually:

``` text
Immediate
    check during statement / relevant operation

Deferred
    check later, often at transaction commit
```

This can be useful for complex transactional operations where
intermediate states temporarily violate a constraint but the final
committed state is valid.

Support is database-specific.

Do not assume every SQL database supports deferred constraints.

------------------------------------------------------------------------

# 12.56 Constraints and Transactions

Constraints work together with transactions.

Example:

``` sql
BEGIN;

INSERT INTO orders (...);

UPDATE inventory
SET quantity = quantity - 1
WHERE product_id = 10;

COMMIT;
```

If a constraint violation occurs, the application may need to roll back
the transaction or otherwise handle the failed statement according to
the database's transaction semantics.

The key principle:

> Constraints prevent invalid database states from being committed
> according to the database's enforcement semantics.

------------------------------------------------------------------------

# 12.57 Constraint Interaction

Constraints can work together.

Example:

``` sql
CREATE TABLE products (
    product_id INTEGER PRIMARY KEY,
    sku VARCHAR(50) NOT NULL UNIQUE,
    price DECIMAL(12,2) NOT NULL
        CHECK (price >= 0),
    status VARCHAR(20) NOT NULL
        DEFAULT 'active'
);
```

Here:

``` text
PRIMARY KEY
    → product identity

NOT NULL
    → required SKU, price, status

UNIQUE
    → SKU cannot duplicate

CHECK
    → price cannot be negative

DEFAULT
    → omitted status becomes active
```

------------------------------------------------------------------------

# 12.58 Full Parent-Child Example

``` sql
CREATE TABLE customers (
    customer_id INTEGER PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE orders (
    order_id INTEGER PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    amount DECIMAL(12,2) NOT NULL
        CHECK (amount >= 0),
    status VARCHAR(20) NOT NULL
        DEFAULT 'pending',

    CONSTRAINT fk_orders_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(customer_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);
```

This schema expresses several rules:

``` text
customer_id
    → unique identity

email
    → required and unique

name
    → required

order_id
    → unique order identity

customer_id in orders
    → required

customer_id in orders
    → must reference a customer

amount
    → required and non-negative

status
    → required and defaults to pending

customer deletion
    → blocked while orders reference it

customer key update
    → propagated to orders
```

------------------------------------------------------------------------

# 12.59 Designing a Good Constraint Set

For each column, ask:

### Identity

``` text
Does this uniquely identify the row?
```

Use:

``` text
PRIMARY KEY
```

### Requiredness

``` text
Must this value exist?
```

Use:

``` text
NOT NULL
```

### Uniqueness

``` text
Can two rows share this value?
```

Use:

``` text
UNIQUE
```

### Domain validity

``` text
What values are logically allowed?
```

Use:

``` text
CHECK
```

### Relationship

``` text
Must this value reference another table?
```

Use:

``` text
FOREIGN KEY
```

### Default

``` text
What value should be used when omitted?
```

Use:

``` text
DEFAULT
```

------------------------------------------------------------------------

# 12.60 Constraints vs Indexes

Constraints and indexes are related but not identical.

A constraint expresses a data-integrity rule.

An index is primarily a data-access structure.

For example:

``` sql
email VARCHAR(255) UNIQUE
```

means:

``` text
email values must satisfy uniqueness
```

The database may implement that uniqueness using an index or an
index-like internal structure.

But:

``` text
UNIQUE constraint
```

and:

``` text
ordinary index
```

are conceptually different.

An ordinary index does not necessarily enforce uniqueness.

------------------------------------------------------------------------

# 12.61 Constraints vs Triggers

A constraint expresses declarative integrity.

A trigger executes procedural logic in response to database events.

Use constraints for straightforward rules:

``` text
salary >= 0
email unique
customer exists
```

Use triggers only when the rule genuinely requires procedural behavior
that constraints cannot express cleanly.

Triggers introduce:

-   hidden side effects
-   additional debugging complexity
-   ordering concerns
-   performance considerations

Prefer declarative constraints when possible.

------------------------------------------------------------------------

# 12.62 Constraints vs Application Logic

Consider:

``` text
email must be unique
```

Application-only validation:

``` text
SELECT ...
if not exists:
    INSERT
```

This can suffer from race conditions.

Two transactions might both observe:

``` text
email does not exist
```

and then both attempt to insert.

A database-level:

``` sql
UNIQUE
```

constraint provides the authoritative integrity guarantee.

Application code should still handle the resulting constraint violation
gracefully.

------------------------------------------------------------------------

# 12.63 Constraint Race Conditions

This is a key production concept.

Suppose two transactions execute:

``` text
Transaction A:
check email → available

Transaction B:
check email → available

A inserts
B inserts
```

Without database enforcement, duplicates may appear.

With:

``` sql
UNIQUE(email)
```

the database ensures only a valid unique state survives according to its
concurrency and transaction semantics.

This is why:

> Application validation is not a substitute for database constraints.

------------------------------------------------------------------------

# 12.64 Soft Delete and Foreign Keys

Many production systems avoid physically deleting important parent rows.

Instead:

``` sql
deleted_at TIMESTAMP NULL
```

can represent logical deletion.

Then:

``` text
customer
    remains physically present
    but is logically deleted
```

This can reduce problems caused by:

``` text
ON DELETE CASCADE
```

and preserve historical relationships.

However, soft deletion creates its own design considerations:

-   unique constraints
-   filtering deleted rows
-   foreign-key semantics
-   storage growth
-   archival strategy

------------------------------------------------------------------------

# 12.65 Choosing ON DELETE Behavior

Use this decision tree:

``` text
Should children disappear with parent?
        │
        ├── Yes
        │     ↓
        │   CASCADE
        │
        └── No
             ↓
      Should children survive
      without relationship?
             │
             ├── Yes → SET NULL
             │
             └── No  → RESTRICT
```

This is a conceptual decision framework, not a universal schema rule.

------------------------------------------------------------------------

# 12.66 Common Mistake: CASCADE Everywhere

Do not blindly use:

``` sql
ON DELETE CASCADE
```

on every foreign key.

Ask:

``` text
Is the child truly dependent on the parent?
```

For:

``` text
order → order_item
```

cascade may make sense.

For:

``` text
customer → financial transaction
```

automatic deletion may be unacceptable.

The correct action depends on domain semantics, audit requirements,
retention rules, and legal/compliance constraints.

------------------------------------------------------------------------

# 12.67 Common Mistake: SET NULL with NOT NULL

This is contradictory:

``` sql
manager_id INTEGER NOT NULL,

FOREIGN KEY (manager_id)
REFERENCES employees(employee_id)
ON DELETE SET NULL
```

If the manager is deleted, the database needs to set:

``` text
manager_id = NULL
```

but:

``` text
NOT NULL
```

prohibits that.

Design the relationship consistently.

------------------------------------------------------------------------

# 12.68 Common Mistake: Assuming DEFAULT Protects Against NULL

This:

``` sql
status VARCHAR(20) DEFAULT 'active'
```

does not necessarily mean:

``` text
status can never be NULL
```

If `NULL` is invalid, use:

``` sql
status VARCHAR(20)
    NOT NULL
    DEFAULT 'active'
```

------------------------------------------------------------------------

# 12.69 Common Mistake: Using UNIQUE Without Considering NULL

If a business rule requires:

``` text
every row has a unique identifier value
```

consider:

``` sql
NOT NULL UNIQUE
```

rather than relying only on:

``` sql
UNIQUE
```

because `NULL` semantics differ from ordinary values and can vary across
database systems.

------------------------------------------------------------------------

# 12.70 Common Mistake: Forgetting Composite-Key Semantics

Given:

``` sql
UNIQUE (country_code, phone_number)
```

this does not mean:

``` text
country_code globally unique
phone_number globally unique
```

It means:

``` text
(country_code, phone_number)
```

must be unique as a combination.

------------------------------------------------------------------------

# 12.71 Common Mistake: Assuming Foreign Keys Are Unique

This is valid:

``` text
customer_id
-----------
1
1
1
2
2
```

in an orders table.

A foreign key normally does not require uniqueness.

Many child rows can reference one parent.

This represents a:

``` text
one-to-many
```

relationship.

If the child reference itself must be unique, add a:

``` sql
UNIQUE
```

constraint.

------------------------------------------------------------------------

# 12.72 One-to-One Relationship

Example:

``` sql
CREATE TABLE user_profiles (
    profile_id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE,

    FOREIGN KEY (user_id)
        REFERENCES users(user_id)
);
```

The:

``` sql
UNIQUE(user_id)
```

constraint ensures each user can have at most one profile.

Without it, multiple profiles could reference the same user.

------------------------------------------------------------------------

# 12.73 One-to-Many Relationship

Typical schema:

``` sql
CREATE TABLE customers (
    customer_id INTEGER PRIMARY KEY
);

CREATE TABLE orders (
    order_id INTEGER PRIMARY KEY,
    customer_id INTEGER NOT NULL,

    FOREIGN KEY (customer_id)
        REFERENCES customers(customer_id)
);
```

There can be:

``` text
1 customer
    ↓
many orders
```

because the foreign key is not unique.

------------------------------------------------------------------------

# 12.74 Many-to-Many Relationship

Use a junction table:

``` sql
CREATE TABLE students (
    student_id INTEGER PRIMARY KEY
);

CREATE TABLE courses (
    course_id INTEGER PRIMARY KEY
);

CREATE TABLE enrollments (
    student_id INTEGER,
    course_id INTEGER,

    PRIMARY KEY (student_id, course_id),

    FOREIGN KEY (student_id)
        REFERENCES students(student_id),

    FOREIGN KEY (course_id)
        REFERENCES courses(course_id)
);
```

The composite primary key prevents duplicate enrollment pairs.

------------------------------------------------------------------------

# 12.75 Constraint Dependencies

A useful dependency graph:

``` text
customers
    │
    │ referenced by
    ↓
orders
    │
    │ referenced by
    ↓
order_items
```

Deleting:

``` text
customers
```

may therefore affect:

``` text
orders
```

which may affect:

``` text
order_items
```

The referential actions determine whether these operations:

``` text
cascade
fail
or null out relationships
```

------------------------------------------------------------------------

# 12.76 Circular Foreign Keys

Some schemas can create cycles:

``` text
A → B
B → A
```

These can complicate:

-   insertion
-   deletion
-   updates
-   transaction ordering
-   deferred constraints

Avoid unnecessary cycles.

When cycles are required by the domain, carefully design:

-   nullable foreign keys
-   insertion order
-   transaction boundaries
-   deferred constraints if supported

------------------------------------------------------------------------

# 12.77 Constraint Introspection

When debugging production schemas, inspect the database metadata to
determine:

``` text
which constraints exist
constraint names
columns involved
referenced tables
referential actions
```

Most database systems expose metadata through:

``` text
INFORMATION_SCHEMA
system catalogs
database-specific metadata views
```

The exact queries differ by SQL dialect.

Never assume the schema matches an outdated application model.

------------------------------------------------------------------------

# 12.78 Debugging a Constraint Violation

When an insert/update/delete fails:

### Step 1 --- Identify the constraint

Determine whether it is:

``` text
PRIMARY KEY
UNIQUE
NOT NULL
CHECK
FOREIGN KEY
```

### Step 2 --- Inspect the offending values

Ask:

``` text
Which value caused the violation?
```

### Step 3 --- Inspect existing rows

For uniqueness:

``` text
Does the value already exist?
```

For foreign keys:

``` text
Does the referenced parent exist?
```

### Step 4 --- Check schema semantics

Verify:

``` text
NULLability
data types
CHECK expression
foreign-key action
```

### Step 5 --- Decide whether data or schema is wrong

The fix may be:

``` text
correct the input
```

or:

``` text
correct the constraint
```

Do not weaken constraints merely to make an error disappear.

------------------------------------------------------------------------

# 12.79 Example Debugging: Foreign Key

Error scenario:

``` sql
INSERT INTO orders (
    order_id,
    customer_id
)
VALUES (100, 999);
```

Constraint:

``` sql
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
```

Debug:

``` sql
SELECT *
FROM customers
WHERE customer_id = 999;
```

If no row exists, the insert violates referential integrity.

Possible fixes:

``` text
insert the parent first
use an existing customer_id
correct the input
```

Do not disable the foreign key simply to allow invalid data.

------------------------------------------------------------------------

# 12.80 Example Debugging: UNIQUE

Error:

``` sql
INSERT INTO users (email)
VALUES ('alice@example.com');
```

Constraint:

``` sql
UNIQUE(email)
```

Check:

``` sql
SELECT user_id, email
FROM users
WHERE email = 'alice@example.com';
```

If a row exists, determine whether:

``` text
duplicate insertion
```

or:

``` text
incorrect uniqueness rule
```

is the actual problem.

------------------------------------------------------------------------

# 12.81 Example Debugging: CHECK

Given:

``` sql
CHECK (quantity > 0)
```

and:

``` sql
INSERT INTO order_items (quantity)
VALUES (0);
```

The constraint violation indicates that the data violates the domain
rule.

Ask:

``` text
Is zero quantity actually invalid?
```

If yes:

``` text
fix input
```

If no:

``` text
revisit schema rule
```

------------------------------------------------------------------------

# 12.82 Constraints and Data Quality

A well-designed schema can enforce important invariants automatically.

For example:

``` sql
CREATE TABLE payments (
    payment_id BIGINT PRIMARY KEY,
    order_id BIGINT NOT NULL,
    amount DECIMAL(12,2) NOT NULL
        CHECK (amount > 0),
    status VARCHAR(20) NOT NULL
        CHECK (status IN ('pending', 'paid', 'failed'))
);
```

The database now prevents several classes of bad data.

------------------------------------------------------------------------

# 12.83 Constraint Design Principle

A useful principle is:

> Store only states that are valid according to the domain model.

Instead of allowing:

``` text
salary = -5000
```

and relying on every application to remember:

``` text
salary >= 0
```

encode the rule:

``` sql
CHECK (salary >= 0)
```

This is called enforcing an invariant.

------------------------------------------------------------------------

# 12.84 Constraints as Invariants

An invariant is a condition that should remain true.

Examples:

``` text
customer_id is unique
```

``` text
salary >= 0
```

``` text
every order references a valid customer
```

``` text
email cannot be NULL
```

``` text
student-course pair cannot repeat
```

Constraints make these invariants executable at the database layer.

------------------------------------------------------------------------

# 12.85 Constraint Enforcement and Concurrency

Constraints are particularly important under concurrent writes.

Suppose two transactions try to create the same unique username
simultaneously.

Application-level:

``` text
check
insert
```

is vulnerable to races.

A database-level:

``` sql
UNIQUE(username)
```

provides the authoritative integrity boundary.

The database's concurrency control determines which transaction succeeds
and which receives the constraint violation.

------------------------------------------------------------------------

# 12.86 Primary Key vs Unique --- Interview Comparison

  Question                PRIMARY KEY         UNIQUE
  ----------------------- ------------------- --------------------------------
  Identifies row?         yes                 can provide alternate identity
  Duplicate values?       no                  no
  NULL allowed?           no                  database-dependent
  Multiple constraints?   one PK constraint   multiple
  Can be composite?       yes                 yes
  Common purpose          row identity        alternate uniqueness

A table can therefore have:

``` text
1 primary key
+
many unique constraints
```

------------------------------------------------------------------------

# 12.87 Foreign Key --- Interview Comparison

A foreign key:

``` text
does not normally make values unique
```

It enforces:

``` text
relationship validity
```

Example:

``` text
customers
1
2

orders
1
1
1
2
```

This is valid.

The same customer can be referenced by many orders.

------------------------------------------------------------------------

# 12.88 Constraint Selection Cheat Sheet

``` text
Need unique row identity?
    → PRIMARY KEY

Need required value?
    → NOT NULL

Need unique alternate value?
    → UNIQUE

Need value/domain rule?
    → CHECK

Need parent-child relationship?
    → FOREIGN KEY

Need automatic omitted value?
    → DEFAULT
```

For parent deletion:

``` text
Delete children too?
    → ON DELETE CASCADE

Keep children, remove relationship?
    → ON DELETE SET NULL

Prevent parent deletion while referenced?
    → ON DELETE RESTRICT
```

For referenced-key changes:

``` text
Propagate key changes?
    → ON UPDATE CASCADE
```

------------------------------------------------------------------------

# 12.89 Complete Example

``` sql
CREATE TABLE departments (
    department_id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE employees (
    employee_id INTEGER PRIMARY KEY,

    email VARCHAR(255)
        NOT NULL
        UNIQUE,

    department_id INTEGER,

    salary DECIMAL(12,2)
        NOT NULL
        CHECK (salary >= 0),

    status VARCHAR(20)
        NOT NULL
        DEFAULT 'active',

    CONSTRAINT fk_employees_department
        FOREIGN KEY (department_id)
        REFERENCES departments(department_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);
```

This expresses:

``` text
departments.department_id
    → primary key

departments.name
    → required and unique

employees.employee_id
    → primary key

employees.email
    → required and unique

employees.salary
    → required and non-negative

employees.status
    → required and defaults to active

employees.department_id
    → optional foreign key

department deletion
    → employee survives, department_id becomes NULL

department key update
    → employee references update
```

------------------------------------------------------------------------

# 12.90 Constraint Design Checklist

Before finalizing a table, ask:

### Identity

``` text
What uniquely identifies one row?
```

### Requiredness

``` text
Which fields are mandatory?
```

### Uniqueness

``` text
Which attributes cannot duplicate?
```

### Domain

``` text
Which values are invalid?
```

### Relationships

``` text
Which columns reference other tables?
```

### Delete behavior

``` text
What happens if the parent is deleted?
```

### Update behavior

``` text
What happens if the referenced key changes?
```

### Defaults

``` text
What should happen when a value is omitted?
```

### Null semantics

``` text
Does NULL mean something meaningful?
```

### Concurrency

``` text
Could application-only validation race?
```

------------------------------------------------------------------------

# 12.91 Interview Questions

## Q1. What is a constraint?

A database rule that restricts allowed data states and helps maintain
data integrity.

------------------------------------------------------------------------

## Q2. What is a primary key?

A constraint that uniquely identifies each row in a table.

------------------------------------------------------------------------

## Q3. Can a table have multiple primary keys?

No. A table has one primary-key constraint, although it may be composite
and contain multiple columns.

------------------------------------------------------------------------

## Q4. Can a table have multiple UNIQUE constraints?

Yes.

------------------------------------------------------------------------

## Q5. Can a foreign key contain duplicate values?

Yes. A foreign key normally does not imply uniqueness.

------------------------------------------------------------------------

## Q6. Can a foreign key contain NULL?

Often yes, unless `NOT NULL` is also specified and subject to the
database's normal constraint semantics.

------------------------------------------------------------------------

## Q7. What is referential integrity?

The rule that valid foreign-key references must correspond to valid
referenced parent keys, subject to SQL's null and referential-action
semantics.

------------------------------------------------------------------------

## Q8. Difference between CASCADE and RESTRICT?

``` text
CASCADE
    parent deletion propagates to children

RESTRICT
    parent deletion is rejected while dependent references exist
```

------------------------------------------------------------------------

## Q9. What does SET NULL do?

It keeps the child row but sets its foreign-key value to `NULL` when the
referenced parent is deleted.

------------------------------------------------------------------------

## Q10. What requirement does SET NULL create?

The foreign-key column must be capable of storing `NULL`.

------------------------------------------------------------------------

## Q11. Does DEFAULT prevent NULL?

No. Use `NOT NULL` as well when `NULL` is invalid.

------------------------------------------------------------------------

## Q12. What is a CHECK constraint?

A constraint that requires an expression to satisfy the database's check
semantics.

Example:

``` sql
CHECK (salary >= 0)
```

------------------------------------------------------------------------

## Q13. What is a composite key?

A key consisting of multiple columns whose combination provides
uniqueness.

------------------------------------------------------------------------

## Q14. Why use a database UNIQUE constraint if the application already checks uniqueness?

Because application checks can race under concurrent transactions. The
database constraint provides authoritative enforcement.

------------------------------------------------------------------------

## Q15. What happens when a constraint is violated?

The database rejects the operation according to the constraint and
transaction semantics of the database system.

------------------------------------------------------------------------

# 12.92 Final Mental Model

Think of constraints as different layers of protection:

``` text
PRIMARY KEY
    ↓
"Who is this row?"

UNIQUE
    ↓
"Can this value occur more than once?"

NOT NULL
    ↓
"Must this value exist?"

CHECK
    ↓
"Is this value valid?"

DEFAULT
    ↓
"What should be used if omitted?"

FOREIGN KEY
    ↓
"Does this relationship point to a valid parent?"
```

Then referential actions answer:

``` text
What happens if the parent changes?

ON DELETE CASCADE
    → delete children

ON DELETE SET NULL
    → preserve children, remove reference

ON DELETE RESTRICT
    → prevent parent deletion

ON UPDATE CASCADE
    → propagate referenced-key changes
```

The deepest principle is:

> **Constraints turn assumptions about valid data into rules the
> database can enforce.**

A strong SQL schema does not merely describe what data looks like. It
defines which states are allowed to exist.

------------------------------------------------------------------------

# 12.93 Final Compact Reference

``` sql
-- Primary key
PRIMARY KEY (id)

-- Unique
UNIQUE (email)

-- Required
email VARCHAR(255) NOT NULL

-- Domain validation
CHECK (salary >= 0)

-- Default
status VARCHAR(20) DEFAULT 'active'

-- Foreign key
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)

-- Delete dependent rows
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON DELETE CASCADE

-- Preserve child but remove reference
FOREIGN KEY (manager_id)
REFERENCES employees(employee_id)
ON DELETE SET NULL

-- Prevent parent deletion while referenced
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON DELETE RESTRICT

-- Propagate referenced-key updates
FOREIGN KEY (customer_id)
REFERENCES customers(customer_id)
ON UPDATE CASCADE
```

------------------------------------------------------------------------

# Conclusion

SQL constraints are a foundational part of relational database design.

The key distinction is:

``` text
PRIMARY KEY
    → identity

UNIQUE
    → uniqueness

NOT NULL
    → requiredness

CHECK
    → valid domain

DEFAULT
    → automatic value

FOREIGN KEY
    → relationship integrity
```

Foreign-key actions then define relationship behavior:

``` text
CASCADE
    → propagate deletion

SET NULL
    → preserve child, remove reference

RESTRICT
    → block deletion

UPDATE CASCADE
    → propagate key changes
```

The practical design rule is:

> **If a rule must always be true for the data, prefer enforcing it at
> the database level whenever the relational model can express it
> cleanly.**

Application validation improves user experience; database constraints
protect the integrity of the actual data.
