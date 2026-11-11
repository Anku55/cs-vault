# Custom Sorting in C++

Custom sorting is used when the default `sort()` order is not enough.

In C++, we can provide a **custom comparator** to `sort()`.

---

# 1. Basic `sort()`

For a vector of integers:

```cpp
vector<int> v = {5, 2, 8, 1, 3};

sort(v.begin(), v.end());
```

Output:

```text
1 2 3 5 8
```

By default, `sort()` uses ascending order.

---

# 2. Descending Order

We can use `greater<int>()`:

```cpp
sort(v.begin(), v.end(), greater<int>());
```

Output:

```text
8 5 3 2 1
```

---

# 3. Custom Comparator

A comparator tells `sort()`:

> Should `a` come before `b`?

Example:

```cpp
bool comp(int a, int b) {
    return a > b;
}
```

Then:

```cpp
sort(v.begin(), v.end(), comp);
```

This sorts in descending order.

The important idea is:

```cpp
return true;
```

means:

```text
a should come before b
```

---

# 4. Comparator Rule

A comparator should use a **strict comparison**.

Good:

```cpp
return a < b;
```

Good:

```cpp
return a > b;
```

Bad:

```cpp
return a <= b;
```

Bad:

```cpp
return a >= b;
```

Use `<` or `>` rather than `<=` or `>=`.

---

# 5. Comparator with `pair`

Suppose:

```cpp
vector<pair<int,int>> v = {
    {3, 10},
    {1, 20},
    {2, 15},
    {3, 5}
};
```

We want to sort by the first element.

```cpp
bool comp(const pair<int,int>& a,
          const pair<int,int>& b) {

    return a.first < b.first;
}
```

Then:

```cpp
sort(v.begin(), v.end(), comp);
```

Output:

```text
{1,20}
{2,15}
{3,10}
{3,5}
```

---

# 6. Multiple Conditions in Comparator

This is one of the most important patterns in competitive programming.

Suppose we want:

1. First element ascending.
2. If first elements are equal, second element descending.

```cpp
bool comp(const pair<int,int>& a,
          const pair<int,int>& b) {

    if(a.first != b.first)
        return a.first < b.first;

    return a.second > b.second;
}
```

Example:

```text
Before:

{3,10}
{1,20}
{3,5}
{2,15}

After:

{1,20}
{2,15}
{3,10}
{3,5}
```

The general pattern is:

```cpp
if(primary values are different)
    return primary comparison;

return secondary comparison;
```

---

# 7. Custom Comparator with Struct

We can create our own structure.

```cpp
struct Student {
    string name;
    int marks;
};
```

Suppose we want higher marks first:

```cpp
bool comp(const Student& a,
          const Student& b) {

    return a.marks > b.marks;
}
```

Then:

```cpp
sort(students.begin(), students.end(), comp);
```

---

# 8. Lambda Comparator

Instead of creating a separate function, we can use a lambda.

```cpp
sort(v.begin(), v.end(), [](int a, int b) {
    return a > b;
});
```

For pairs:

```cpp
sort(v.begin(), v.end(), [](auto& a, auto& b) {

    if(a.first != b.first)
        return a.first < b.first;

    return a.second > b.second;
});
```

---

# 9. Important Question to Ask

Whenever writing a comparator, ask:

> **Should `a` come before `b`?**

For example:

```cpp
return a < b;
```

means:

```text
a comes before b
```

And:

```cpp
return a > b;
```

means:

```text
a comes before b when a is larger
```

---

# 10. Why Do We Sometimes Store the Original Index?

Sometimes we don't want to sort the original value directly.

Instead, we calculate some **derived/transformed value** and sort using that value.

Example:

```text
Original number → transformed number

12 → 45
25 → 31
31 → 17
```

We want to sort according to the transformed value, but finally return:

```text
original numbers
```

In that situation, store:

```cpp
{
    transformedValue,
    originalIndex
}
```

Example:

```cpp
struct Node {
    int value;
    int index;
};
```

Then:

```cpp
arr.push_back({transformedValue, i});
```

After sorting:

```cpp
nums[arr[i].index]
```

gives the original value.

---

# 11. Example 1 — 1451. Rearrange Words in a Sentence

## Problem

Rearrange words according to their length.

Rules:

1. Shorter words come first.
2. If two words have the same length, preserve their original order.
3. The first word should start with uppercase.

Example:

```text
Input:

"Leetcode is cool"
```

Words:

```text
Leetcode → length 8 → index 0
is       → length 2 → index 1
cool     → length 4 → index 2
```

After sorting:

```text
is
cool
Leetcode
```

Answer:

```text
"Is cool leetcode"
```

## Code

```cpp
class Solution {
public:

    struct Word {
        string s;
        int index;
    };

    static bool comp(const Word& a, const Word& b) {

        // Shorter word first
        if(a.s.length() != b.s.length())
            return a.s.length() < b.s.length();

        // Same length -> original order
        return a.index < b.index;
    }

    string arrangeWords(string text) {

        // Convert first character to lowercase
        text[0] = tolower(text[0]);

        stringstream ss(text);

        vector<Word> words;

        string word;
        int index = 0;

        // Split sentence
        while(ss >> word) {

            words.push_back({word, index});

            index++;
        }

        // Custom sorting
        sort(words.begin(), words.end(), comp);

        // First word should start with uppercase
        words[0].s[0] = toupper(words[0].s[0]);

        string ans;

        for(int i = 0; i < words.size(); i++) {

            if(i > 0)
                ans += " ";

            ans += words[i].s;
        }

        return ans;
    }
};
```

## Comparator

The important part:

```cpp
static bool comp(const Word& a, const Word& b) {

    if(a.s.length() != b.s.length())
        return a.s.length() < b.s.length();

    return a.index < b.index;
}
```

This means:

```text
Primary condition:
    shorter word first

Secondary condition:
    smaller original index first
```

---

# 12. Example 2 — 2191. Sort the Jumbled Numbers

## Problem

Each digit is transformed according to `mapping`.

For example:

```text
mapping = [8,9,4,0,2,1,3,5,7,6]
```

Then:

```text
0 → 8
1 → 9
2 → 4
3 → 0
...
```

We need to sort numbers according to their transformed values.

But we must return the **original numbers**, not the transformed numbers.

---

## Example

Suppose:

```text
nums = [991, 338, 38]
```

We calculate the transformed values and store:

```text
mapped value     original index
--------------------------------
...              0
...              1
...              2
```

So we use:

```cpp
struct mp {
    int num;
    int index;
};
```

---

## Code

```cpp
class Solution {
public:

    struct mp {
        int num;
        int index;
    };

    static bool comp(const mp& a, const mp& b) {

        // Smaller mapped number first
        if(a.num != b.num)
            return a.num < b.num;

        // Same mapped number -> original order
        return a.index < b.index;
    }

    vector<int> sortJumbled(vector<int>& mapping,
                             vector<int>& nums) {

        vector<mp> arr;

        for(int i = 0; i < nums.size(); i++) {

            int t = nums[i];

            int j = 1;
            int number = 0;

            // Special case: number is 0
            if(t == 0) {
                number = mapping[0];
            }

            while(t > 0) {

                int digit = mapping[t % 10];

                t = t / 10;

                number += j * digit;

                j = j * 10;
            }

            // Store transformed number + original index
            arr.push_back({number, i});
        }

        // Sort using transformed number
        sort(arr.begin(), arr.end(), comp);

        vector<int> ans;

        // Return ORIGINAL numbers
        for(int i = 0; i < arr.size(); i++) {

            ans.push_back(nums[arr[i].index]);
        }

        return ans;
    }
};
```

---

# 13. Critical Difference in 2191

Notice this:

```cpp
arr.push_back({number, i});
```

Here:

```text
number = transformed number
i      = original index
```

After sorting:

```cpp
sort(arr.begin(), arr.end(), comp);
```

`arr` contains the numbers in their **sorted transformed order**.

But we don't return:

```cpp
arr[i].num
```

because that is the transformed number.

We return:

```cpp
nums[arr[i].index]
```

because that gives the original number.

### Remember

```text
transformed value → used for sorting

original index → used for retrieving original value
```

---

# 14. General Pattern

This pattern appears frequently in competitive programming.

```cpp
struct Node {
    int key;
    int index;
};
```

Build:

```cpp
vector<Node> arr;

for(int i = 0; i < nums.size(); i++) {

    int key = calculateSomething(nums[i]);

    arr.push_back({key, i});
}
```

Sort:

```cpp
sort(arr.begin(), arr.end(), comp);
```

Retrieve:

```cpp
nums[arr[i].index]
```

This is sometimes called **decorate → sort → undecorate**.

---

# 15. Common Comparator Templates

## Ascending

```cpp
bool comp(int a, int b) {
    return a < b;
}
```

## Descending

```cpp
bool comp(int a, int b) {
    return a > b;
}
```

## First ascending, second ascending

```cpp
bool comp(pair<int,int>& a,
          pair<int,int>& b) {

    if(a.first != b.first)
        return a.first < b.first;

    return a.second < b.second;
}
```

## First ascending, second descending

```cpp
bool comp(pair<int,int>& a,
          pair<int,int>& b) {

    if(a.first != b.first)
        return a.first < b.first;

    return a.second > b.second;
}
```

## Struct

```cpp
bool comp(const Node& a,
          const Node& b) {

    if(a.key != b.key)
        return a.key < b.key;

    return a.index < b.index;
}
```

---

# 16. `sort()` vs `stable_sort()`

Sometimes equal elements need to maintain their original order.

C++ provides:

```cpp
stable_sort()
```

Example:

```cpp
stable_sort(v.begin(), v.end(), comp);
```

`stable_sort()` preserves the relative order of elements that are considered equal by the comparator.

For example:

```text
Before:

A(5)
B(2)
C(5)
D(3)
```

Sort by number:

```text
B(2)
D(3)
A(5)
C(5)
```

Notice:

```text
A before C
```

is preserved.

---

# 17. When Should You Use Custom Sorting?

Use a custom comparator when the problem says things like:

- Sort by length
- Sort by frequency
- Sort by distance
- Sort by one value, then another
- Sort according to a custom rule
- Sort objects based on a calculated value
- Sort numbers according to a mapping
- Preserve original order for equal values
- Sort intervals using custom criteria

---

# 18. Mental Model

Whenever you see:

> "Sort according to X"

Think:

```text
Can I calculate X for every element?
        ↓
Store X
        ↓
Write comparator using X
        ↓
Sort
```

For example:

```text
Sort students by marks
        ↓
compare marks

Sort words by length
        ↓
compare length

Sort numbers by mapped value
        ↓
compare mapped value

Sort jobs by deadline
        ↓
compare deadline
```

---

# Key Takeaways

### 1. Comparator answers:

```text
Should a come before b?
```

### 2. Multiple conditions:

```cpp
if(primary differs)
    return primary comparison;

return secondary comparison;
```

### 3. Never use:

```cpp
<=
>=
```

for normal `sort()` comparators.

Use strict comparisons:

```cpp
<
>
```

### 4. When sorting transformed values:

Store:

```cpp
{transformedValue, originalIndex}
```

Then retrieve:

```cpp
originalArray[originalIndex]
```

### 5. Two important problems practiced:

```text
1451 — Rearrange Words in a Sentence
2191 — Sort the Jumbled Numbers
```

Both demonstrate the same fundamental idea:

> **Define exactly what determines which element should come first, then encode that rule in the comparator.**