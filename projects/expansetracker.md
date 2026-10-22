C++ Expense Tracker
A console-based expense tracking application built with C++17. I built this project to apply core C++ concepts such as Object-Oriented Programming, STL containers, file handling, input validation, and basic data persistence in a practical application.
Overview
The Expense Tracker allows users to record and analyze their expenses from a command-line interface.

Each expense contains:
- ID
- Amount
- Category
- Date
- Description
The application can:
- Add expenses
- View saved expenses
- Calculate monthly expenditure
- Analyze spending by category
- Identify the highest-spending category
- Save expenses to a local file
- Load expenses when the application starts
- Validate amounts and dates
- Handle malformed records more safely
The project is intentionally a console application. I wanted to focus on the application logic and C++ design before adding a GUI or database.
Tech Stack
- Language: C++17
- Compiler: GCC / MinGW
- Standard Library: STL
- IDE: VS Code
- Platform: Windows
- Version Control: Git / GitHub
- Storage: Local text file
No external libraries or database are required.
Project Structure
Expense-Tracker/
│
├── main.cpp
├── Expense.h
├── Expense.cpp
├── ExpenseTracker.h
├── ExpenseTracker.cpp
├── expenses.txt
├── README.md
└── ...
main.cpp
Handles the program entry point, menu, and user interaction.
Expense
Represents a single expense and stores its:
- ID
- Amount
- Category
- Date
- Description
ExpenseTracker
Handles the main application logic, including:
- Adding expenses
- Displaying expenses
- Monthly calculations
- Category analysis
- Input validation
- Loading data
- Saving data
Design
I kept the design relatively simple with two main classes.
                ┌──────────────┐
                │   main.cpp   │
                │ Menu / Input │
                └──────┬───────┘
                       │
                       ▼
             ┌──────────────────┐
             │  ExpenseTracker  │
             │                  │
             │ Business Logic   │
             │ Validation       │
             │ Aggregation      │
             │ Persistence      │
             └────────┬─────────┘
                      │
                      ▼
             ┌──────────────────┐
             │ vector<Expense>  │
             └────────┬─────────┘
                      │
                      ▼
             ┌──────────────────┐
             │   expenses.txt   │
             └──────────────────┘

Category Analysis
        │
        ▼
map<string, double>
The main idea is to keep the expense data separate from the application logic.
Data Structures
vector<Expense>
I use:
std::vector<Expense> expenses;
because the number of expenses is not fixed.

A vector works well because the application frequently needs to:
- Add expenses
- Iterate through all expenses
- Calculate totals
- Display records
Complexity
Operation	Complexity
Add at end	Amortized O(1)
Iterate	O(n)
Monthly calculation	O(n)
Linear ID search	O(n)
Storage	O(n)


I chose a vector instead of a linked list because the application mainly performs sequential access and iteration. A vector also provides good cache locality and simple STL integration.
map<string, double>
For category analysis, I use:
std::map<std::string, double> categoryTotals;
The map stores the total amount spent for each category.

For example:
Food       → ₹250.50
Travel     → ₹250.50
Shopping   → ₹120.00
Bills      → ₹500.00
While processing expenses:
categoryTotals[expense.getCategory()] += expense.getAmount();
This gives me a straightforward way to aggregate spending by category.

If:
- n = number of expenses
- k = number of unique categories
then:
- Building category totals: O(n log k)
- Finding the maximum: O(k)
- Extra aggregation space: O(k)
I chose map because the category-to-total relationship naturally fits a key-value structure, and map also keeps keys ordered.

An unordered_map would also be a valid choice and could provide average O(1) lookup.
File Persistence
One of the goals of the project was to make sure expenses don't disappear when the application closes.

The application saves expenses to:
expenses.txt
and loads them again when it starts.

The file begins with a version header:
EXPENSE_TRACKER_V1
This identifies the expected file format and gives the application a way to distinguish the current format from a future version.
Custom File Format
A simple delimiter-only format can become problematic when text fields contain the delimiter.

For example:
1|250.5|Food|2026-10-11|Lunch | Dinner
The | inside the description makes simple parsing ambiguous.

To avoid this, the current format uses length-prefixed strings for text fields.

Example:
1|250.5|15:Food and Dining10:2026-10-1118:Lunch with friends
The parser knows exactly how many characters belong to each string.

This means characters such as:
|
:
can safely appear inside categories and descriptions.
Data Safety During Saving
The application uses a temporary-file / backup-style saving approach instead of blindly replacing the original data file.

The goal is to reduce the chance of losing the existing data if a write is interrupted or fails.

This is not intended to be a fully transactional or crash-proof persistence system. If the application grows into a larger system, I would use a database or a more robust persistence layer.
Input Validation
I added validation so that the application does not assume every user input is correct.
Amount Validation
The amount must be:
- Positive
- Finite
- Within the configured maximum
The current maximum is:
1,000,000,000
This prevents invalid values such as zero, negative amounts, NaN, infinity, and unreasonably large values.
Date Validation
Dates follow:
YYYY-MM-DD
The application also checks whether the date is actually valid.

For example:
2026-13-01 → invalid
2025-02-29 → invalid
2024-02-29 → valid
Leap years follow the standard calendar rule:
Divisible by 400
OR
Divisible by 4 but not divisible by 100
Examples:
2000-02-29 → valid
1900-02-29 → invalid
2024-02-29 → valid
2025-02-29 → invalid
Monthly Expense Calculation
For a monthly summary, the application checks each expense against the requested month and year.

Conceptually:
For every expense:
    Check year
    Check month
    If both match:
        Add amount to total
Complexity
Time  → O(n)
Space → O(1)
assuming only a running total is required.
Highest-Spending Category
Category analysis first creates a map of:
category → total spending
Then the application finds the maximum total.

Ties are handled explicitly.

For example:
Food       ₹250.50
Travel     ₹250.50
Shopping   ₹100.00
The result includes both:
Food
Travel
rather than selecting only one category.

Because monetary calculations use floating-point values in the current version, a configured tolerance is used when comparing totals.
Why double for Money?
I currently use:
double
because this is a learning project and it keeps the implementation straightforward.

However, double is not ideal for financial calculations because binary floating-point cannot represent every decimal value exactly.

For a more production-oriented implementation, I would store money as integer minor units.

For example:
₹250.50
could be represented as:
25050 paise
This would avoid many floating-point precision issues.
Debugging: Windows CRLF Issue
One of the practical issues I encountered was related to Windows line endings.

Windows commonly uses:
\r\n
while Unix-style systems generally use:
\n
When reading the file, getline() removes the newline character but a trailing \r can remain in some reading modes.

That caused a valid header comparison to fail because the value being compared could effectively contain an extra carriage-return character.

I fixed this by normalizing a trailing carriage return before validating the header and record lines.

After recompiling, the application successfully loaded the existing records on Windows.

This was a useful debugging experience because the issue was caused by the interaction between the file format and the operating system rather than by the main application logic.
Testing
I manually tested the main application flows, including:
- Adding expenses
- Displaying expenses
- Persistence across application restarts
- Category aggregation
- Highest-category ties
- Monthly calculations
- Invalid input
- Malformed records
- Categories and descriptions containing | and :
- Windows CRLF handling
I also verified that a March 2026 monthly query returned no matching expenses for the current October 2026 dataset.

These are manual/integration checks, not a comprehensive automated test suite.
Planned Testing Improvements
I would add automated tests for:
- Date validation
- Amount validation
- Monthly calculations
- Category aggregation
- Serialization
- Deserialization
- Malformed input handling
Integration tests should use temporary test files so that they never modify real expense data.
Complexity Summary
For:
n = number of expenses
k = number of unique categories
Operation	Complexity
Add expense	Amortized O(1)
Display expenses	O(n)
Monthly summary	O(n)
Category aggregation	O(n log k)
Find maximum category	O(k)
In-memory storage	O(n)
Category aggregation storage	O(k)


Disk I/O and console output add their own overhead and are not included in these core collection/calculation complexities.
Why I Chose This Design
I didn't want to over-engineer a small project.

For the current requirements:
vector + map + local file
is enough.

I could have introduced a database, multiple service layers, a GUI, authentication, or networking, but that would add complexity without solving a requirement of the current application.

If the requirements grow, the architecture should grow with them.
Limitations
This is a learning project, not a production financial application.

Current limitations include:
1. Local text-file storage
   - No database
   - Limited querying
   - No concurrent users
2. double for monetary values
   - Can introduce floating-point precision issues
3. Console interface
   - No graphical interface
4. Limited automated testing
   - Current testing is mainly manual/integration testing
5. Single-user design
   - No authentication or multi-user support
6. Limited CRUD functionality
   - The current version focuses on adding, viewing, and analyzing expenses
Future Improvements
If I continue developing this project, I would prioritize:
1. Automated Tests
Add unit tests for validation, calculations, and persistence.
2. Better Currency Representation
Replace double with integer minor units such as paise.
3. Database Persistence
Move to SQLite if the application needs:
- More complex queries
- Larger datasets
- Transactions
- Better recovery
- Concurrent access
4. Additional Features
Potential additions include:
- Edit expense
- Delete expense
- Date-range filtering
- Budget limits
- Search
- Reports
- CSV export
What I Learned
This project helped me move from learning individual C++ concepts to using them together in an actual application.

The main things I practised were:
- Object-Oriented Programming
- Encapsulation
- STL containers
- vector
- map
- References and const
- File handling
- Serialization and deserialization
- Input validation
- Date validation
- Error handling
- Time and space complexity
- Git and GitHub
- Debugging platform-specific issues
One of the more useful lessons was that building an application is not only about implementing features. Handling invalid input, corrupted data, persistence, and unexpected platform behavior is also part of writing reliable software.
Running the Project
Compile
Using GCC / MinGW:
g++ -std=c++17 -Wall -Wextra -Wpedantic main.cpp Expense.cpp ExpenseTracker.cpp -o expense_tracker
Run
On Windows:
.\expense_tracker.exe
The application expects expenses.txt in the current working directory for existing data. If there is no existing data file, the application can start with an empty expense list.
Project Status
Current status: Functional learning project

The core expense management, analysis, validation, and persistence functionality is implemented.

The next focus areas would be automated testing, stronger monetary representation, and improved persistence.
Author
Ankush

Built as a C++ learning project to practise software development concepts through a practical application.

If you find an issue or have a suggestion, feel free to open an issue or submit a pull request.