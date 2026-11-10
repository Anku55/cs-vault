# 🧩 Backtracking — LeetCode Practice Roadmap

A structured roadmap to master **Backtracking** through LeetCode problems, progressing from basic recursion and subsets to advanced constraint-based search, state compression, and combinatorial optimization.

---

## 📌 What is Backtracking?

Backtracking is an algorithmic technique in which we build a solution step by step, and whenever we realize that the current choice cannot lead to a valid solution, we undo that choice and try another one.

**In simple words:**

`Try → Explore → If wrong, Undo → Try another option`

### Core Backtracking Template

```cpp
void backtrack(/* state */)
{
    if(/* base case */)
    {
        // Store or process the solution
        return;
    }

    for(/* each possible choice */)
    {
        // Choose

        backtrack(/* updated state */);

        // Undo the choice
    }
}
```

---

## 🗺️ Roadmap Overview

| Group | Topic | Problems |
|---|---|---:|
| 1 | Basic Recursion, Subsets & Binary Choices | 8 |
| 2 | Permutations, Arrangements & Ordering | 8 |
| 3 | Combinations, Target Sum & Subsequence Search | 9 |
| 4 | String Partitioning & Expression Generation | 8 |
| 5 | Grid, Graph & Path-Based Backtracking | 7 |
| 6 | Partitioning Into Buckets / Distribution | 6 |
| 7 | Advanced State Search, Game & Structural Backtracking | 16 |
| | **Total** | **62** |

---

# 🟢 Group 1 — Basic Recursion, Subsets & Binary Choices

**Start here**

### Core Concepts
- Pick / Not Pick
- Recursion tree
- Base cases
- Path building
- Binary decision recursion

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 1 | Sum of All Subset XOR Totals | [1863](https://leetcode.com/problems/sum-of-all-subset-xor-totals/) | Easy | [ ] |
| 2 | Letter Case Permutation | [784](https://leetcode.com/problems/letter-case-permutation/) | Medium | [ ] |
| 3 | Generate Binary Strings Without Adjacent Zeros | [3211](https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/) | Medium | [ ] |
| 4 | The k-th Lexicographical String of All Happy Strings of Length n | [1415](https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/) | Medium | [ ] |
| 5 | Numbers With Same Consecutive Differences | [967](https://leetcode.com/problems/numbers-with-same-consecutive-differences/) | Medium | [ ] |
| 6 | Count Numbers with Unique Digits | [357](https://leetcode.com/problems/count-numbers-with-unique-digits/) | Medium | [ ] |
| 7 | Maximum Rows Covered by Columns | [2397](https://leetcode.com/problems/maximum-rows-covered-by-columns/) | Medium | [ ] |
| 8 | Maximum Split of Positive Even Integers | [2178](https://leetcode.com/problems/maximum-split-of-positive-even-integers/) | Medium | [ ] |

---

# 🟡 Group 2 — Permutations, Arrangements & Ordering

### Core Concepts
- `visited[]` array
- Used elements
- Swap-based recursion
- Duplicate handling
- Permutation generation
- Ordering constraints

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 9 | Beautiful Arrangement | [526](https://leetcode.com/problems/beautiful-arrangement/) | Medium | [ ] |
| 10 | Number of Squareful Arrays | [996](https://leetcode.com/problems/number-of-squareful-arrays/) | Hard | [ ] |
| 11 | Circular Permutation in Binary Representation | [1238](https://leetcode.com/problems/circular-permutation-in-binary-representation/) | Medium | [ ] |
| 12 | Construct Smallest Number From DI String | [2375](https://leetcode.com/problems/construct-smallest-number-from-di-string/) | Medium | [ ] |
| 13 | Construct the Lexicographically Largest Valid Sequence | [1718](https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/) | Medium | [ ] |
| 14 | Number of Valid Move Combinations On Chessboard | [2056](https://leetcode.com/problems/number-of-valid-move-combinations-on-chessboard/) | Hard | [ ] |
| 15 | Next Greater Numerically Balanced Number | [2048](https://leetcode.com/problems/next-greater-numerically-balanced-number/) | Medium | [ ] |
| 16 | Next Special Palindrome Number | [3646](https://leetcode.com/problems/next-special-palindrome-number/) | Hard | [ ] |

---

# 🟡 Group 3 — Combinations, Target Sum & Subsequence Search

**Very Important**

### Core Concepts
- `startIndex`
- Target sum
- Duplicate skipping
- Combination construction
- Subsequences
- Pruning

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 17 | Combination Sum III | [216](https://leetcode.com/problems/combination-sum-iii/) | Medium | [ ] |
| 18 | Non-decreasing Subsequences | [491](https://leetcode.com/problems/non-decreasing-subsequences/) | Medium | [ ] |
| 19 | Maximum Length of a Concatenated String with Unique Characters | [1239](https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/) | Medium | [ ] |
| 20 | Split a String Into the Max Number of Unique Substrings | [1593](https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/) | Medium | [ ] |
| 21 | Maximum Product of the Length of Two Palindromic Subsequences | [2002](https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/) | Medium | [ ] |
| 22 | Longest Subsequence Repeated k Times | [2014](https://leetcode.com/problems/longest-subsequence-repeated-k-times/) | Hard | [ ] |
| 23 | Maximum Points in an Archery Competition | [2212](https://leetcode.com/problems/maximum-points-in-an-archery-competition/) | Medium | [ ] |
| 24 | Maximum Strength of a Group | [2708](https://leetcode.com/problems/maximum-strength-of-a-group/) | Medium | [ ] |
| 25 | Balanced K-Factor Decomposition | [3669](https://leetcode.com/problems/balanced-k-factor-decomposition/) | Medium | [ ] |

---

# 🟠 Group 4 — String Partitioning & Expression Generation

### Core Concepts
- Choosing split points
- Validating substrings
- Maintaining expression state
- Recursive string construction
- Handling operators
- Leading zero constraints

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 26 | Restore IP Addresses | [93](https://leetcode.com/problems/restore-ip-addresses/) | Medium | [ ] |
| 27 | Split Array into Fibonacci Sequence | [842](https://leetcode.com/problems/split-array-into-fibonacci-sequence/) | Medium | [ ] |
| 28 | Additive Number | [306](https://leetcode.com/problems/additive-number/) | Medium | [ ] |
| 29 | Ambiguous Coordinates | [816](https://leetcode.com/problems/ambiguous-coordinates/) | Medium | [ ] |
| 30 | Splitting a String Into Descending Consecutive Values | [1849](https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/) | Medium | [ ] |
| 31 | Partition String Into Minimum Beautiful Substrings | [2767](https://leetcode.com/problems/partition-string-into-minimum-beautiful-substrings/) | Medium | [ ] |
| 32 | Expression Add Operators | [282](https://leetcode.com/problems/expression-add-operators/) | Hard | [ ] |
| 33 | Verbal Arithmetic Puzzle | [1307](https://leetcode.com/problems/verbal-arithmetic-puzzle/) | Hard | [ ] |

---

# 🔴 Group 5 — Grid, Graph & Path-Based Backtracking

### Core Concepts
- DFS + Backtracking
- `visited[]`
- Mark / Unmark
- Direction arrays
- Path tracking
- Graph constraints

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 34 | Binary Tree Paths | [257](https://leetcode.com/problems/binary-tree-paths/) | Easy | [ ] |
| 35 | Word Search | [79](https://leetcode.com/problems/word-search/) | Medium | [ ] |
| 36 | Unique Paths III | [980](https://leetcode.com/problems/unique-paths-iii/) | Hard | [ ] |
| 37 | Word Search II | [212](https://leetcode.com/problems/word-search-ii/) | Hard | [ ] |
| 38 | Word Ladder II | [126](https://leetcode.com/problems/word-ladder-ii/) | Hard | [ ] |
| 39 | Maximum Path Quality of a Graph | [2065](https://leetcode.com/problems/maximum-path-quality-of-a-graph/) | Hard | [ ] |
| 40 | Sliding Puzzle | [773](https://leetcode.com/problems/sliding-puzzle/) | Hard | [ ] |

> **Prerequisite:** LeetCode 79 — Word Search is recommended before LeetCode 212 — Word Search II.

---

# 🔴 Group 6 — Partitioning Into Buckets / Distribution

### Core Concepts
- Assigning elements into groups
- Symmetry pruning
- Bucket states
- Constraint optimization
- Sorting for pruning
- Backtracking with capacities

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 41 | Matchsticks to Square | [473](https://leetcode.com/problems/matchsticks-to-square/) | Medium | [ ] |
| 42 | Partition to K Equal Sum Subsets | [698](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/) | Medium | [ ] |
| 43 | Fair Distribution of Cookies | [2305](https://leetcode.com/problems/fair-distribution-of-cookies/) | Medium | [ ] |
| 44 | Distribute Repeating Integers | [1655](https://leetcode.com/problems/distribute-repeating-integers/) | Hard | [ ] |
| 45 | Minimum Number of Work Sessions to Finish the Tasks | [1986](https://leetcode.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/) | Medium | [ ] |
| 46 | Find Minimum Time to Finish All Jobs | [1723](https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/) | Hard | [ ] |

---

# 🔥 Group 7 — Advanced State Search, Game & Structural Backtracking

### Core Concepts
- Branch and Bound
- Bitmasking
- Memoization
- State compression
- Constraint-based search
- Assignment optimization
- Game search
- Recursive tree construction

| # | Problem | LeetCode | Difficulty | Status |
|---|---|---|---|---|
| 47 | Unique Binary Search Trees II | [95](https://leetcode.com/problems/unique-binary-search-trees-ii/) | Medium | [ ] |
| 48 | Iterator for Combination | [1286](https://leetcode.com/problems/iterator-for-combination/) | Medium | [ ] |
| 49 | Pyramid Transition Matrix | [756](https://leetcode.com/problems/pyramid-transition-matrix/) | Medium | [ ] |
| 50 | Shopping Offers | [638](https://leetcode.com/problems/shopping-offers/) | Medium | [ ] |
| 51 | Stickers to Spell Word | [691](https://leetcode.com/problems/stickers-to-spell-word/) | Hard | [ ] |
| 52 | 24 Game | [679](https://leetcode.com/problems/24-game/) | Hard | [ ] |
| 53 | Tiling a Rectangle with the Fewest Squares | [1240](https://leetcode.com/problems/tiling-a-rectangle-with-the-fewest-squares/) | Hard | [ ] |
| 54 | Maximum Score Words Formed by Letters | [1255](https://leetcode.com/problems/maximum-score-words-formed-by-letters/) | Hard | [ ] |
| 55 | Probability of a Two Boxes Having The Same Number of Distinct Balls | [1467](https://leetcode.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/) | Hard | [ ] |
| 56 | Maximize Score After N Operations | [1799](https://leetcode.com/problems/maximize-score-after-n-operations/) | Hard | [ ] |
| 57 | Maximum Compatibility Score Sum | [1947](https://leetcode.com/problems/maximum-compatibility-score-sum/) | Medium | [ ] |
| 58 | Maximum Good People Based on Statements | [2151](https://leetcode.com/problems/maximum-good-people-based-on-statements/) | Hard | [ ] |
| 59 | Closest Dessert Cost | [1774](https://leetcode.com/problems/closest-dessert-cost/) | Medium | [ ] |
| 60 | Minimum Moves to Spread Stones Over Grid | [2850](https://leetcode.com/problems/minimum-moves-to-spread-stones-over-grid/) | Medium | [ ] |
| 61 | Minimum Time to Break Locks I | [3376](https://leetcode.com/problems/minimum-time-to-break-locks-i/) | Medium | [ ] |
| 62 | Word Squares II | [3799](https://leetcode.com/problems/word-squares-ii/) | Hard | [ ] |

---

# 🧠 Backtracking Patterns — Revision Checklist

Use this checklist to track mastery of the underlying patterns.

- [ ] Pick / Not Pick recursion
- [ ] Subset generation
- [ ] Permutation generation
- [ ] Combination generation
- [ ] `startIndex` pattern
- [ ] `visited[]` pattern
- [ ] Swap-based recursion
- [ ] Duplicate handling
- [ ] String partitioning
- [ ] Expression generation
- [ ] Grid DFS + Backtracking
- [ ] Graph path search
- [ ] Bucket assignment
- [ ] Symmetry pruning
- [ ] Branch and Bound
- [ ] Bitmask + Backtracking
- [ ] Memoization on states
- [ ] Constraint-based search
- [ ] Assignment optimization

---

# 📈 Progress Tracker

| Metric | Progress |
|---|---:|
| Total Problems | 62 |
| Solved | 0 |
| Remaining | 62 |
| Completion | 0% |

**Goal:** Solve all 62 problems while understanding the recursion tree, state transitions, pruning strategy, and time complexity of each solution.

---

## 🎯 Practice Strategy

1. Understand the pattern before solving the problem.
2. Try solving independently before checking hints.
3. Draw the recursion tree for difficult problems.
4. Identify the state, choices, base case, and backtracking step.
5. Apply pruning wherever possible.
6. Analyze time and space complexity.
7. Revisit difficult problems after a few days.

> **The objective is not merely to solve 62 problems. The objective is to recognize backtracking patterns and apply them independently in new problems.**

