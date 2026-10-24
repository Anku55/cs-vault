# 🧠 Dynamic Programming — Complete LeetCode Pattern Roadmap

> A structured DP roadmap organized by **patterns, sub-patterns, observational patterns, tricky patterns, and representative LeetCode problems**.
>


# 📚 Table of Contents

- [1. Fibonacci / Basic 1D DP](#1-fibonacci--basic-1d-dp)
- [2. Take / Skip DP](#2-take--skip-dp)
- [3. Kadane / Maximum Subarray](#3-kadane--maximum-subarray-dp)
- [4. State Machine DP](#4-state-machine-dp)
- [5. 0/1 Knapsack](#5-01-knapsack)
- [6. Subset Sum](#6-subset-sum)
- [7. Unbounded Knapsack](#7-unbounded-knapsack)
- [8. Counting DP](#8-counting-dp)
- [9. Grid DP](#9-grid-dp)
- [10. Multi-Agent Grid DP](#10-multi-agent-grid-dp)
- [11. LCS / Two-String DP](#11-lcs--two-string-dp)
- [12. Longest Common Substring](#12-longest-common-substring)
- [13. Edit Distance](#13-edit-distance-family)
- [14. LIS / Subsequence DP](#14-lis--subsequence-dp)
- [15. LIS + Binary Search](#15-lis--binary-search)
- [16. Palindrome DP](#16-palindrome-dp)
- [17. String Partition DP](#17-string-partition-dp)
- [18. Interval DP](#18-interval-dp)
- [19. Catalan / Parenthesization DP](#19-catalan--parenthesization-dp)
- [20. Tree DP](#20-tree-dp)
- [21. Tree Rerooting DP](#21-tree-rerooting-dp)
- [22. DAG / Graph DP](#22-dag--graph-dp)
- [23. DP + Binary Search](#23-dp--binary-search)
- [24. DP + Prefix Sum](#24-dp--prefix-sum)
- [25. DP + Monotonic Deque](#25-dp--monotonic-deque)
- [26. DP + Heap](#26-dp--heap)
- [27. Bitmask DP](#27-bitmask-dp)
- [28. TSP-Style DP](#28-tsp-style-dp)
- [29. Digit DP](#29-digit-dp)
- [30. Probability DP](#30-probability-dp)
- [31. Expected Value DP](#31-expected-value-dp)
- [32. Game DP](#32-game-dp)
- [33. String State / Automaton DP](#33-string-state--automaton-dp)
- [34. Automaton / Forbidden Pattern DP](#34-automaton--forbidden-pattern-dp)
- [35. Finite-State DP](#35-finite-state-dp)
- [36. Profile / Broken Profile DP](#36-profile--broken-profile-dp)
- [37. Multi-Dimensional Resource DP](#37-multi-dimensional-resource-dp)
- [38. DP + Graph State](#38-dp--graph-state)
- [39. DP + Shortest Path](#39-dp--shortest-path)
- [40. Scheduling DP](#40-scheduling-dp)
- [41. Divide & Conquer DP Optimization](#41-divide--conquer-dp-optimization)
- [42. Convex Hull Trick / Li Chao DP](#42-convex-hull-trick--li-chao-type-dp)
- [43. SOS / Subset Transform DP](#43-sos--subset-transform-dp)
- [44. Tree + State DP](#44-tree--state-dp)
- [45. Answer Ending at i](#45-answer-ending-at-i)
- [46. Answer Starting from i](#46-answer-starting-from-i)
- [47. Choose Left or Right](#47-choose-left-or-right)
- [48. Split Somewhere](#48-split-somewhere)
- [49. Previous Compatible Element](#49-previous-compatible-element)
- [50. How Many Ways](#50-how-many-ways)
- [51. Can We Reach X](#51-can-we-reach-x)
- [52. DP Disguised as Backtracking](#52-dp-disguised-as-backtracking)
- [53. DP Disguised as Greedy](#53-dp-disguised-as-greedy)
- [Core 40 Problem Set](#-core-40-dp-problem-set)

---

# 1. Fibonacci / Basic 1D DP

### Pattern

```text
dp[i] depends on previous few states
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 70 | Climbing Stairs | ⭐⭐⭐⭐⭐ |
| 509 | Fibonacci Number | ⭐⭐⭐ |
| 746 | Min Cost Climbing Stairs | ⭐⭐⭐⭐⭐ |
| 198 | House Robber | ⭐⭐⭐⭐⭐ |
| 213 | House Robber II | ⭐⭐⭐⭐⭐ |
| 91 | Decode Ways | ⭐⭐⭐⭐⭐ |
| 343 | Integer Break | ⭐⭐⭐⭐ |
| 279 | Perfect Squares | ⭐⭐⭐⭐ |
| 338 | Counting Bits | ⭐⭐⭐ |
| 1025 | Divisor Game | ⭐⭐⭐ |
| 1137 | N-th Tribonacci Number | ⭐⭐ |
| 790 | Domino and Tromino Tiling | ⭐⭐⭐⭐ |
| 1269 | Number of Ways to Stay in the Same Place | ⭐⭐⭐⭐ |
| 276 | Paint Fence | ⭐⭐⭐ |
| 377 | Combination Sum IV | ⭐⭐⭐⭐⭐ |

### Recommended

```text
70 → 746 → 198 → 213 → 91 → 377
```

---

# 2. Take / Skip DP

### Recognition

At every element:

```text
take
  OR
skip
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 198 | House Robber | ⭐⭐⭐⭐⭐ |
| 213 | House Robber II | ⭐⭐⭐⭐⭐ |
| 740 | Delete and Earn | ⭐⭐⭐⭐ |
| 2140 | Solving Questions With Brainpower | ⭐⭐⭐⭐⭐ |
| 1911 | Maximum Alternating Subsequence Sum | ⭐⭐⭐⭐ |
| 1626 | Best Team With No Conflicts | ⭐⭐⭐⭐ |
| 689 | Maximum Sum of 3 Non-Overlapping Subarrays | ⭐⭐⭐⭐ |

### Key Idea

Many apparently different problems reduce to:

```cpp
dp[i] = max(
    take_current + dp[previous_valid],
    skip_current
);
```

---

# 3. Kadane / Maximum Subarray DP

### State

```cpp
dp[i] = best subarray ending at i
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 53 | Maximum Subarray | ⭐⭐⭐⭐⭐ |
| 918 | Maximum Circular Subarray | ⭐⭐⭐⭐ |
| 152 | Maximum Product Subarray | ⭐⭐⭐⭐⭐ |
| 121 | Best Time to Buy and Sell Stock | ⭐⭐⭐⭐⭐ |
| 1746 | Maximum Subarray Sum After One Operation | ⭐⭐⭐ |

### Important Observation

For Maximum Product Subarray:

```text
maximum ending here
minimum ending here
```

are both required.

---

# 4. State Machine DP

### Core State

```text
day × state
```

Examples:

```text
holding
not holding
cooldown
transactions remaining
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 121 | Best Time to Buy and Sell Stock | ⭐⭐⭐⭐⭐ |
| 122 | Best Time to Buy and Sell Stock II | ⭐⭐⭐⭐⭐ |
| 123 | Best Time to Buy and Sell Stock III | ⭐⭐⭐⭐⭐ |
| 188 | Best Time to Buy and Sell Stock IV | ⭐⭐⭐⭐⭐ |
| 309 | Stock with Cooldown | ⭐⭐⭐⭐⭐ |
| 714 | Stock with Transaction Fee | ⭐⭐⭐⭐⭐ |
| 256 | Paint House | ⭐⭐⭐⭐ |
| 265 | Paint House II | ⭐⭐⭐⭐ |
| 1473 | Paint House III | ⭐⭐⭐ |
| 801 | Minimum Swaps to Make Sequences Increasing | ⭐⭐⭐⭐ |

### Recommended Order

```text
121 → 122 → 309 → 714 → 123 → 188
```

---

# 5. 0/1 Knapsack

### Definition

Each item can be used **at most once**.

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 416 | Partition Equal Subset Sum | ⭐⭐⭐⭐⭐ |
| 494 | Target Sum | ⭐⭐⭐⭐⭐ |
| 1049 | Last Stone Weight II | ⭐⭐⭐⭐⭐ |
| 474 | Ones and Zeroes | ⭐⭐⭐⭐ |
| 698 | Partition to K Equal Sum Subsets | ⭐⭐⭐⭐ |
| 879 | Profitable Schemes | ⭐⭐⭐⭐ |
| 956 | Tallest Billboard | ⭐⭐⭐⭐ |
| 2035 | Partition Array Into Two Arrays to Minimize Sum Difference | ⭐⭐⭐⭐ |

### Key Idea

```text
Each item:
    take
    skip
```

---

# 6. Subset Sum

### Problems

| LC | Problem |
|---:|---|
| 416 | Partition Equal Subset Sum ⭐⭐⭐⭐⭐ |
| 494 | Target Sum ⭐⭐⭐⭐⭐ |
| 1049 | Last Stone Weight II ⭐⭐⭐⭐⭐ |
| 698 | Partition to K Equal Sum Subsets ⭐⭐⭐⭐ |
| 473 | Matchsticks to Square ⭐⭐⭐⭐ |
| 2035 | Partition Array Into Two Arrays to Minimize Sum Difference ⭐⭐⭐⭐ |
| 474 | Ones and Zeroes ⭐⭐⭐⭐ |

### Mental Bucket

```text
416
494
1049
2035
```

→ **Subset / Partition Knapsack**

---

# 7. Unbounded Knapsack

### Definition

An item can be used **multiple times**.

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 322 | Coin Change | ⭐⭐⭐⭐⭐ |
| 518 | Coin Change II | ⭐⭐⭐⭐⭐ |
| 279 | Perfect Squares | ⭐⭐⭐⭐ |
| 377 | Combination Sum IV | ⭐⭐⭐⭐⭐ |
| 343 | Integer Break | ⭐⭐⭐⭐ |
| 983 | Minimum Cost For Tickets | ⭐⭐⭐⭐ |

### Critical Difference

```cpp
// 0/1 Knapsack
for (int w = W; w >= weight; --w)

// Unbounded Knapsack
for (int w = weight; w <= W; ++w)
```

---

# 8. Counting DP

### Question Type

> **How many ways?**

### Problems

| LC | Problem |
|---:|---|
| 70 | Climbing Stairs |
| 91 | Decode Ways ⭐⭐⭐⭐⭐ |
| 518 | Coin Change II ⭐⭐⭐⭐⭐ |
| 377 | Combination Sum IV ⭐⭐⭐⭐⭐ |
| 62 | Unique Paths ⭐⭐⭐⭐⭐ |
| 63 | Unique Paths II ⭐⭐⭐⭐ |
| 115 | Distinct Subsequences ⭐⭐⭐⭐⭐ |
| 1155 | Number of Dice Rolls With Target Sum ⭐⭐⭐⭐ |
| 1639 | Number of Ways to Form a Target String ⭐⭐⭐⭐ |
| 920 | Number of Music Playlists ⭐⭐⭐ |
| 1575 | Count All Possible Routes ⭐⭐⭐ |

---

# 9. Grid DP

### Basic State

```cpp
dp[r][c]
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 62 | Unique Paths | ⭐⭐⭐⭐⭐ |
| 63 | Unique Paths II | ⭐⭐⭐⭐⭐ |
| 64 | Minimum Path Sum | ⭐⭐⭐⭐⭐ |
| 120 | Triangle | ⭐⭐⭐⭐⭐ |
| 931 | Minimum Falling Path Sum | ⭐⭐⭐⭐ |
| 1289 | Minimum Falling Path Sum II | ⭐⭐⭐ |
| 221 | Maximal Square | ⭐⭐⭐⭐⭐ |
| 1277 | Count Square Submatrices With All Ones | ⭐⭐⭐⭐ |
| 174 | Dungeon Game | ⭐⭐⭐⭐⭐ |
| 741 | Cherry Pickup | ⭐⭐⭐ |
| 1463 | Cherry Pickup II | ⭐⭐⭐⭐ |
| 2328 | Number of Increasing Paths in a Grid | ⭐⭐⭐⭐ |
| 1444 | Number of Ways of Cutting a Pizza | ⭐⭐⭐ |

---

# 10. Multi-Agent Grid DP

### Typical State

```cpp
dp[r1][c1][r2][c2]
```

But:

```text
r1 + c1 = r2 + c2
```

can reduce dimensions.

### Problems

| LC | Problem |
|---:|---|
| 741 | Cherry Pickup ⭐⭐⭐⭐⭐ |
| 1463 | Cherry Pickup II ⭐⭐⭐⭐⭐ |
| 1301 | Number of Paths With Max Score ⭐⭐⭐⭐ |
| 1937 | Maximum Number of Points With Cost ⭐⭐⭐⭐ |

---

# 11. LCS / Two-String DP

### Core State

```cpp
dp[i][j]
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 1143 | Longest Common Subsequence | ⭐⭐⭐⭐⭐ |
| 1035 | Uncrossed Lines | ⭐⭐⭐⭐ |
| 583 | Delete Operation for Two Strings | ⭐⭐⭐⭐ |
| 1092 | Shortest Common Supersequence | ⭐⭐⭐⭐ |
| 712 | Minimum ASCII Delete Sum | ⭐⭐⭐⭐ |
| 115 | Distinct Subsequences | ⭐⭐⭐⭐⭐ |
| 97 | Interleaving String | ⭐⭐⭐⭐⭐ |
| 72 | Edit Distance | ⭐⭐⭐⭐⭐ |
| 516 | Longest Palindromic Subsequence | ⭐⭐⭐⭐⭐ |
| 392 | Is Subsequence | ⭐⭐ |

### Recommended Order

```text
1143
 ↓
583
 ↓
1092
 ↓
115
 ↓
97
 ↓
72
```

---

# 12. Longest Common Substring

### Substring ≠ Subsequence

For substring:

```cpp
if (a[i] == b[j])
    dp[i][j] = dp[i-1][j-1] + 1;
else
    dp[i][j] = 0;
```

### Problems

| LC | Problem |
|---:|---|
| 718 | Maximum Length of Repeated Subarray ⭐⭐⭐⭐ |
| 1044 | Longest Duplicate Substring ⭐⭐⭐ |
| 5 | Longest Palindromic Substring ⭐⭐⭐⭐⭐ |

---

# 13. Edit Distance Family

| LC | Problem |
|---:|---|
| 72 | Edit Distance ⭐⭐⭐⭐⭐ |
| 583 | Delete Operation for Two Strings ⭐⭐⭐⭐ |
| 712 | Minimum ASCII Delete Sum ⭐⭐⭐⭐ |
| 161 | One Edit Distance ⭐⭐⭐ |
| 115 | Distinct Subsequences ⭐⭐⭐⭐⭐ |
| 1092 | Shortest Common Supersequence ⭐⭐⭐⭐ |

---

# 14. LIS / Subsequence DP

### State

```cpp
dp[i] = best subsequence ending at i
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 300 | Longest Increasing Subsequence | ⭐⭐⭐⭐⭐ |
| 673 | Number of LIS | ⭐⭐⭐⭐⭐ |
| 354 | Russian Doll Envelopes | ⭐⭐⭐⭐⭐ |
| 646 | Maximum Length of Pair Chain | ⭐⭐⭐⭐ |
| 1027 | Longest Arithmetic Subsequence | ⭐⭐⭐⭐ |
| 1218 | Longest Arithmetic Subsequence of Given Difference | ⭐⭐⭐⭐ |
| 376 | Wiggle Subsequence | ⭐⭐⭐⭐ |
| 1626 | Best Team With No Conflicts | ⭐⭐⭐⭐ |
| 1048 | Longest String Chain | ⭐⭐⭐⭐ |
| 368 | Largest Divisible Subset | ⭐⭐⭐⭐ |

---

# 15. LIS + Binary Search

### Optimization

```text
O(n²) LIS
   ↓
O(n log n)
```

### Problems

| LC | Problem |
|---:|---|
| 300 | Longest Increasing Subsequence ⭐⭐⭐⭐⭐ |
| 354 | Russian Doll Envelopes ⭐⭐⭐⭐⭐ |
| 1671 | Minimum Number of Removals to Make Mountain Array ⭐⭐⭐⭐ |
| 1964 | Longest Valid Obstacle Course at Each Position ⭐⭐⭐⭐ |
| 2389 | Longest Subsequence With Limited Sum ⭐⭐⭐ |

---

# 16. Palindrome DP

### State

```cpp
dp[l][r]
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 5 | Longest Palindromic Substring | ⭐⭐⭐⭐⭐ |
| 647 | Palindromic Substrings | ⭐⭐⭐⭐⭐ |
| 516 | Longest Palindromic Subsequence | ⭐⭐⭐⭐⭐ |
| 1312 | Minimum Insertion Steps to Make a String Palindrome | ⭐⭐⭐⭐ |
| 132 | Palindrome Partitioning II | ⭐⭐⭐⭐⭐ |
| 131 | Palindrome Partitioning | ⭐⭐⭐ |
| 1216 | Valid Palindrome III | ⭐⭐⭐ |
| 2131 | Longest Palindrome by Concatenating Two Letter Words | ⭐⭐⭐ |

---

# 17. String Partition DP

### State

```cpp
dp[i] = answer for prefix [0...i]
```

### Problems

| LC | Problem |
|---:|---|
| 139 | Word Break ⭐⭐⭐⭐⭐ |
| 140 | Word Break II ⭐⭐⭐⭐ |
| 132 | Palindrome Partitioning II ⭐⭐⭐⭐⭐ |
| 93 | Restore IP Addresses ⭐⭐⭐ |
| 1416 | Restore the Array ⭐⭐⭐⭐ |
| 1043 | Partition Array for Maximum Sum ⭐⭐⭐⭐ |
| 91 | Decode Ways ⭐⭐⭐⭐⭐ |
| 2707 | Extra Characters in a String ⭐⭐⭐⭐ |

---

# 18. Interval DP

### Core State

```cpp
dp[l][r]
```

### Typical Recurrence

```cpp
dp[l][r] = best over k:
    dp[l][k] + dp[k][r] + cost
```

### Problems

| LC | Problem | Importance |
|---:|---|---|
| 312 | Burst Balloons | ⭐⭐⭐⭐⭐ |
| 1039 | Minimum Score Triangulation | ⭐⭐⭐⭐ |
| 1547 | Minimum Cost to Cut a Stick | ⭐⭐⭐⭐⭐ |
| 664 | Strange Printer | ⭐⭐⭐⭐ |
| 1406 | Stone Game III | ⭐⭐⭐⭐ |
| 1000 | Merge Stones | ⭐⭐⭐⭐ |
| 546 | Remove Boxes | ⭐⭐⭐ |
| 132 | Palindrome Partitioning II | ⭐⭐⭐⭐⭐ |
| 87 | Scramble String | ⭐⭐⭐ |
| 1000 | Minimum Cost to Merge Stones | ⭐⭐⭐⭐ |

### Must Do

```text
312
1039
1547
664
1000
```

---

# 19. Catalan / Parenthesization DP

| LC | Problem |
|---:|---|
| 96 | Unique Binary Search Trees ⭐⭐⭐⭐ |
| 95 | Unique Binary Search Trees II ⭐⭐⭐ |
| 22 | Generate Parentheses ⭐⭐⭐ |
| 241 | Different Ways to Add Parentheses ⭐⭐⭐⭐ |
| 1039 | Minimum Score Triangulation ⭐⭐⭐⭐ |
| 312 | Burst Balloons ⭐⭐⭐⭐⭐ |

---

# 20. Tree DP 🌳

### State

```cpp
dp[node][state]
```

### Problems

| LC | Problem |
|---:|---|
| 337 | House Robber III ⭐⭐⭐⭐⭐ |
| 124 | Binary Tree Maximum Path Sum ⭐⭐⭐⭐⭐ |
| 543 | Diameter of Binary Tree ⭐⭐⭐⭐⭐ |
| 1339 | Maximum Product of Splitted Binary Tree ⭐⭐⭐ |
| 968 | Binary Tree Cameras ⭐⭐⭐⭐⭐ |
| 1372 | Longest ZigZag Path ⭐⭐⭐⭐ |
| 1373 | Maximum Sum BST in Binary Tree ⭐⭐⭐⭐ |
| 894 | All Possible Full Binary Trees ⭐⭐⭐ |
| 979 | Distribute Coins in Binary Tree ⭐⭐⭐⭐ |

---

# 21. Tree Rerooting DP

### Recognition

Question asks something for **every possible root**.

### Problems

| LC | Problem |
|---:|---|
| 834 | Sum of Distances in Tree ⭐⭐⭐⭐⭐ |
| 310 | Minimum Height Trees ⭐⭐⭐ |
| 1376 | Time Needed to Inform All Employees ⭐⭐⭐ |
| 2467 | Most Profitable Path in a Tree ⭐⭐⭐⭐ |
| 2581 | Count Number of Possible Root Nodes ⭐⭐⭐⭐ |

### Canonical

```text
834 — Sum of Distances in Tree
```

---

# 22. DAG / Graph DP

| LC | Problem |
|---:|---|
| 329 | Longest Increasing Path in a Matrix ⭐⭐⭐⭐⭐ |
| 787 | Cheapest Flights Within K Stops ⭐⭐⭐⭐⭐ |
| 1136 | Parallel Courses ⭐⭐⭐⭐ |
| 2050 | Parallel Courses III ⭐⭐⭐⭐ |
| 1059 | All Paths From Source Lead to Destination ⭐⭐⭐ |
| 1976 | Number of Ways to Arrive at Destination ⭐⭐⭐ |
| 1575 | Count All Possible Routes ⭐⭐⭐⭐ |
| 1928 | Minimum Cost to Reach Destination in Time ⭐⭐⭐⭐ |
| 847 | Shortest Path Visiting All Nodes ⭐⭐⭐⭐⭐ |

---

# 23. DP + Binary Search

### Recognition

```text
sort intervals
+
find previous compatible interval
+
DP
```

### Problems

| LC | Problem |
|---:|---|
| 1235 | Maximum Profit in Job Scheduling ⭐⭐⭐⭐⭐ |
| 1751 | Maximum Number of Events That Can Be Attended II ⭐⭐⭐⭐ |
| 1187 | Make Array Strictly Increasing ⭐⭐⭐⭐ |
| 1671 | Minimum Number of Removals to Make Mountain Array ⭐⭐⭐⭐ |
| 354 | Russian Doll Envelopes ⭐⭐⭐⭐⭐ |
| 887 | Super Egg Drop ⭐⭐⭐⭐ |
| 2218 | Maximum Value of K Coins From Piles ⭐⭐⭐⭐ |

### Canonical

```text
1235 — Maximum Profit in Job Scheduling
```

---

# 24. DP + Prefix Sum

| LC | Problem |
|---:|---|
| 1444 | Number of Ways of Cutting a Pizza ⭐⭐⭐⭐ |
| 1473 | Paint House III ⭐⭐⭐ |
| 1269 | Number of Ways to Stay in the Same Place ⭐⭐⭐⭐ |
| 920 | Number of Music Playlists ⭐⭐⭐ |
| 1425 | Constrained Subsequence Sum ⭐⭐⭐⭐⭐ |

---

# 25. DP + Monotonic Deque

### Pattern

```text
dp[i] = best value inside sliding window
```

### Problems

| LC | Problem |
|---:|---|
| 1425 | Constrained Subsequence Sum ⭐⭐⭐⭐⭐ |
| 1696 | Jump Game VI ⭐⭐⭐⭐⭐ |
| 239 | Sliding Window Maximum ⭐⭐⭐⭐⭐ |
| 55 | Jump Game ⭐⭐⭐⭐⭐ |
| 45 | Jump Game II ⭐⭐⭐⭐⭐ |

### Key Pattern

```text
DP + Sliding Window Maximum
```

---

# 26. DP + Heap

| LC | Problem | Main Idea |
|---:|---|---|
| 1425 | Constrained Subsequence Sum | DP + Heap/Deque |
| 1751 | Maximum Number of Events II | DP + Search |
| 1696 | Jump Game VI | DP + Heap/Deque |
| 857 | Minimum Cost to Hire K Workers | Greedy + Heap |
| 502 | IPO | Greedy + Heap |

> Not every problem in this category is strictly DP. The purpose is to recognize when DP can be replaced or combined with another optimization technique.

---

# 27. Bitmask DP

### State

```cpp
dp[mask]
```

or:

```cpp
dp[mask][last]
```

### Problems

| LC | Problem |
|---:|---|
| 847 | Shortest Path Visiting All Nodes ⭐⭐⭐⭐⭐ |
| 464 | Can I Win ⭐⭐⭐⭐ |
| 698 | Partition to K Equal Sum Subsets ⭐⭐⭐⭐ |
| 473 | Matchsticks to Square ⭐⭐⭐⭐ |
| 1125 | Smallest Sufficient Team ⭐⭐⭐⭐ |
| 1349 | Maximum Students Taking Exam ⭐⭐⭐ |
| 1434 | Number of Ways to Wear Different Hats ⭐⭐⭐ |
| 1595 | Minimum Cost to Connect Two Groups of Points ⭐⭐⭐ |
| 1601 | Maximum Number of Achievable Transfer Requests ⭐⭐⭐ |
| 1723 | Minimum Time to Finish All Jobs ⭐⭐⭐⭐ |
| 1659 | Maximize Grid Happiness ⭐⭐⭐ |
| 1655 | Distribute Repeating Integers ⭐⭐⭐ |
| 1799 | Maximize Score After N Operations ⭐⭐⭐⭐ |

---

# 28. TSP-Style DP

### State

```cpp
dp[mask][last]
```

### Problems

| LC | Problem |
|---:|---|
| 847 | Shortest Path Visiting All Nodes ⭐⭐⭐⭐⭐ |
| 1947 | Maximum Compatibility Score Sum ⭐⭐⭐⭐ |
| 1125 | Smallest Sufficient Team ⭐⭐⭐⭐ |
| 1349 | Maximum Students Taking Exam ⭐⭐⭐ |
| 1595 | Minimum Cost to Connect Two Groups ⭐⭐⭐ |

---

# 29. Digit DP 🔢

### Typical State

```cpp
dp[pos][tight][started][state]
```

### Problems

| LC | Problem |
|---:|---|
| 600 | Non-negative Integers Without Consecutive Ones ⭐⭐⭐⭐ |
| 902 | Numbers At Most N Given Digit Set ⭐⭐⭐⭐ |
| 1012 | Numbers With Repeated Digits ⭐⭐⭐⭐⭐ |
| 357 | Count Numbers With Unique Digits ⭐⭐⭐⭐ |
| 2719 | Count Integers With Digit Sum in Range ⭐⭐⭐⭐ |
| 233 | Number of Digit One ⭐⭐⭐⭐ |
| 2801 | Count Stepping Numbers in Range ⭐⭐⭐ |

### Recommended

```text
600
 ↓
357
 ↓
902
 ↓
1012
 ↓
2719
```

---

# 30. Probability DP 🎲

| LC | Problem |
|---:|---|
| 688 | Knight Probability in Chessboard ⭐⭐⭐⭐⭐ |
| 808 | Soup Servings ⭐⭐⭐⭐ |
| 837 | New 21 Game ⭐⭐⭐⭐ |
| 1230 | Toss Strange Coins ⭐⭐⭐⭐ |
| 1227 | Airplane Seat Assignment Probability ⭐⭐⭐ |
| 799 | Champagne Tower ⭐⭐⭐ |
| 1223 | Dice Roll Simulation ⭐⭐⭐ |

---

# 31. Expected Value DP

| LC | Problem |
|---:|---|
| 837 | New 21 Game ⭐⭐⭐⭐⭐ |
| 808 | Soup Servings ⭐⭐⭐⭐ |
| 688 | Knight Probability ⭐⭐⭐⭐ |
| 1227 | Airplane Seat Assignment Probability ⭐⭐⭐ |
| 1510 | Stone Game IV ⭐⭐⭐⭐ |

---

# 32. Game DP 🎮

### Core Idea

Instead of maximizing absolute score:

```text
maximize score difference
```

### Problems

| LC | Problem |
|---:|---|
| 486 | Predict the Winner ⭐⭐⭐⭐⭐ |
| 877 | Stone Game ⭐⭐⭐⭐ |
| 1140 | Stone Game II ⭐⭐⭐⭐ |
| 1406 | Stone Game III ⭐⭐⭐⭐⭐ |
| 1510 | Stone Game IV ⭐⭐⭐⭐ |
| 1563 | Stone Game V ⭐⭐⭐ |
| 464 | Can I Win ⭐⭐⭐⭐ |
| 913 | Cat and Mouse ⭐⭐⭐ |

---

# 33. String State / Automaton DP

| LC | Problem |
|---:|---|
| 10 | Regular Expression Matching ⭐⭐⭐⭐⭐ |
| 44 | Wildcard Matching ⭐⭐⭐⭐⭐ |
| 97 | Interleaving String ⭐⭐⭐⭐⭐ |
| 87 | Scramble String ⭐⭐⭐ |
| 139 | Word Break ⭐⭐⭐⭐⭐ |
| 1639 | Number of Ways to Form Target String ⭐⭐⭐⭐ |
| 664 | Strange Printer ⭐⭐⭐⭐ |
| 115 | Distinct Subsequences ⭐⭐⭐⭐⭐ |

---

# 34. Automaton / Forbidden Pattern DP

### Recognition

The problem imposes restrictions such as:

```text
cannot have consecutive X
cannot exceed K occurrences
certain transitions forbidden
```

### Problems

| LC | Problem |
|---:|---|
| 10 | Regular Expression Matching |
| 44 | Wildcard Matching |
| 552 | Student Attendance Record II ⭐⭐⭐⭐ |
| 600 | Non-negative Integers Without Consecutive Ones |
| 1220 | Count Vowels Permutation ⭐⭐⭐⭐ |
| 1269 | Number of Ways to Stay in the Same Place ⭐⭐⭐⭐ |

---

# 35. Finite-State DP

### Example State

```text
day × absences × consecutive_late
```

### Problems

| LC | Problem |
|---:|---|
| 551 | Student Attendance Record I |
| 552 | Student Attendance Record II ⭐⭐⭐⭐ |
| 1220 | Count Vowels Permutation ⭐⭐⭐⭐ |
| 1269 | Number of Ways to Stay in the Same Place ⭐⭐⭐⭐ |

---

# 36. Profile / Broken Profile DP

### State

```cpp
dp[column][mask]
```

### Problems

| LC | Problem |
|---:|---|
| 790 | Domino and Tromino Tiling |
| 1240 | Tiling a Rectangle with the Fewest Squares ⭐⭐⭐ |
| 1349 | Maximum Students Taking Exam ⭐⭐⭐ |
| 1411 | Number of Ways to Paint N × 3 Grid ⭐⭐⭐⭐ |

---

# 37. Multi-Dimensional Resource DP

### Problems

| LC | Problem |
|---:|---|
| 474 | Ones and Zeroes ⭐⭐⭐⭐⭐ |
| 879 | Profitable Schemes ⭐⭐⭐⭐ |
| 1473 | Paint House III ⭐⭐⭐ |
| 1751 | Maximum Number of Events II ⭐⭐⭐⭐ |
| 2218 | Maximum Value of K Coins From Piles ⭐⭐⭐⭐ |
| 1478 | Allocate Mailboxes ⭐⭐⭐ |

### Recognition

Multiple independent constraints:

```text
index
+
budget
+
count
+
capacity
```

→ likely multidimensional DP.

---

# 38. DP + Graph State

### Problems

| LC | Problem |
|---:|---|
| 787 | Cheapest Flights Within K Stops ⭐⭐⭐⭐⭐ |
| 847 | Shortest Path Visiting All Nodes ⭐⭐⭐⭐⭐ |
| 1928 | Minimum Cost to Reach Destination in Time ⭐⭐⭐⭐ |
| 1976 | Number of Ways to Arrive at Destination ⭐⭐⭐ |
| 1575 | Count All Possible Routes ⭐⭐⭐⭐ |
| 329 | Longest Increasing Path in a Matrix ⭐⭐⭐⭐⭐ |

---

# 39. DP + Shortest Path

### Important Observation

Sometimes the correct solution is not pure DP.

It may become:

```text
DP
↓
Graph
↓
Shortest Path
```

### Problems

| LC | Problem |
|---:|---|
| 787 | Cheapest Flights Within K Stops ⭐⭐⭐⭐⭐ |
| 1928 | Minimum Cost to Reach Destination in Time ⭐⭐⭐⭐ |
| 1976 | Number of Ways to Arrive at Destination ⭐⭐⭐ |
| 1631 | Path With Minimum Effort ⭐⭐⭐⭐ |
| 778 | Swim in Rising Water ⭐⭐⭐⭐ |

---

# 40. Scheduling DP

### Problems

| LC | Problem |
|---:|---|
| 1235 | Maximum Profit in Job Scheduling ⭐⭐⭐⭐⭐ |
| 1751 | Maximum Number of Events II ⭐⭐⭐⭐ |
| 1335 | Minimum Difficulty of a Job Schedule ⭐⭐⭐⭐ |
| 2218 | Maximum Value of K Coins From Piles ⭐⭐⭐⭐ |
| 630 | Course Schedule III ⭐⭐⭐ |

---

# 41. Divide & Conquer DP Optimization

### General Form

```text
dp[i][j] = min/max over k:
           dp[i-1][k] + cost(k, j)
```

### Problems

| LC | Problem |
|---:|---|
| 1547 | Minimum Cost to Cut a Stick ⭐⭐⭐⭐ |
| 1478 | Allocate Mailboxes ⭐⭐⭐⭐ |
| 1335 | Minimum Difficulty of a Job Schedule ⭐⭐⭐⭐ |
| 664 | Strange Printer ⭐⭐⭐ |
| 312 | Burst Balloons ⭐⭐⭐⭐⭐ |

---

# 42. Convex Hull Trick / Li Chao-Type DP

### Level

```text
Advanced Competitive Programming
```

LeetCode does not provide a clean progression for CHT/Li Chao DP.

Related optimization-style problems:

```text
1000 — Minimum Cost to Merge Stones
1478 — Allocate Mailboxes
1547 — Minimum Cost to Cut a Stick
1335 — Minimum Difficulty of a Job Schedule
```

> Learn CHT/Li Chao separately rather than forcing every interval optimization into this category.

---

# 43. SOS / Subset Transform DP

### Level

```text
Advanced Competitive Programming
```

Related bitmask problems:

| LC | Problem |
|---:|---|
| 1125 | Smallest Sufficient Team |
| 1349 | Maximum Students Taking Exam |
| 1434 | Number of Ways to Wear Different Hats |
| 1601 | Maximum Achievable Transfer Requests |
| 1659 | Maximize Grid Happiness |

---

# 44. Tree + State DP

| LC | Problem |
|---:|---|
| 337 | House Robber III ⭐⭐⭐⭐⭐ |
| 968 | Binary Tree Cameras ⭐⭐⭐⭐⭐ |
| 124 | Binary Tree Maximum Path Sum ⭐⭐⭐⭐⭐ |
| 979 | Distribute Coins in Binary Tree ⭐⭐⭐⭐ |
| 1373 | Maximum Sum BST in Binary Tree ⭐⭐⭐⭐ |
| 1372 | Longest ZigZag Path ⭐⭐⭐⭐ |
| 1339 | Maximum Product of Splitted Binary Tree ⭐⭐⭐ |

---

# 45. Observational Pattern — Answer Ending at `i`

### Recognition

If the problem asks:

> longest/best/maximum **ending at i**

think:

```cpp
dp[i]
```

### Problems

| LC | Problem |
|---:|---|
| 300 | Longest Increasing Subsequence |
| 673 | Number of LIS |
| 1027 | Longest Arithmetic Subsequence |
| 1218 | Longest Arithmetic Subsequence of Given Difference |
| 1048 | Longest String Chain |
| 376 | Wiggle Subsequence |
| 53 | Maximum Subarray |

---

# 46. Observational Pattern — Answer Starting From `i`

### Recognition

```cpp
dp[i] = answer from i → end
```

### Problems

| LC | Problem |
|---:|---|
| 139 | Word Break |
| 55 | Jump Game |
| 45 | Jump Game II |
| 91 | Decode Ways |
| 983 | Minimum Cost For Tickets |
| 1416 | Restore the Array |
| 877 | Stone Game |
| 1140 | Stone Game II |
| 1406 | Stone Game III |

---

# 47. Observational Pattern — Choose Left or Right

### Recognition

If choices are:

```text
left
OR
right
```

investigate:

```cpp
dp[l][r]
```

### Problems

| LC | Problem |
|---:|---|
| 486 | Predict the Winner ⭐⭐⭐⭐⭐ |
| 877 | Stone Game |
| 1140 | Stone Game II |
| 1406 | Stone Game III |
| 312 | Burst Balloons |
| 1039 | Minimum Score Triangulation |
| 664 | Strange Printer |

---

# 48. Observational Pattern — Split Somewhere

### Recognition

If the problem says:

> partition/split at position `k`

think:

```cpp
for (int k = l; k < r; k++)
```

### Problems

| LC | Problem |
|---:|---|
| 312 | Burst Balloons |
| 1039 | Minimum Score Triangulation |
| 1547 | Minimum Cost to Cut a Stick |
| 664 | Strange Printer |
| 1000 | Merge Stones |
| 132 | Palindrome Partitioning II |
| 1043 | Partition Array for Maximum Sum |
| 1335 | Minimum Difficulty of a Job Schedule |

---

# 49. Observational Pattern — Previous Compatible Element

### Recognition

Typical input:

```text
jobs
intervals
events
```

Need:

```text
current + best compatible previous
```

Think:

```text
Sorting
+
Binary Search
+
DP
```

### Problems

| LC | Problem |
|---:|---|
| 1235 | Maximum Profit in Job Scheduling ⭐⭐⭐⭐⭐ |
| 1751 | Maximum Number of Events II ⭐⭐⭐⭐ |
| 354 | Russian Doll Envelopes ⭐⭐⭐⭐⭐ |

---

# 50. Observational Pattern — "How Many Ways?"

Immediately consider **Counting DP**.

### Problems

```text
70
91
62
63
518
377
115
1155
920
1639
1575
```

Typical state:

```cpp
dp[state] = number_of_ways
```

---

# 51. Observational Pattern — "Can We Reach X?"

Immediately consider **Boolean DP**.

### Problems

| LC | Problem |
|---:|---|
| 416 | Partition Equal Subset Sum |
| 139 | Word Break |
| 55 | Jump Game |
| 494 | Target Sum |
| 698 | Partition to K Equal Sum Subsets |
| 473 | Matchsticks to Square |
| 97 | Interleaving String |
| 10 | Regular Expression Matching |
| 44 | Wildcard Matching |

Typical state:

```cpp
dp[state] = true / false;
```

---

# 52. DP Disguised as Backtracking

### Recognition

If recursion repeatedly reaches the **same state**, memoization can convert it into DP.

### Problems

| LC | Problem |
|---:|---|
| 139 | Word Break |
| 464 | Can I Win |
| 698 | Partition to K Equal Sum Subsets |
| 473 | Matchsticks to Square |
| 97 | Interleaving String |
| 87 | Scramble String |
| 10 | Regular Expression Matching |
| 691 | Stickers to Spell Word |

### Key Question

```text
"Have I already solved this exact state?"
```

If yes:

```text
Memoization → DP
```

---

# 53. DP Disguised as Greedy

Useful for understanding when greedy is valid or invalid.

| LC | Problem |
|---:|---|
| 55 | Jump Game |
| 45 | Jump Game II |
| 198 | House Robber |
| 53 | Maximum Subarray |
| 121 | Best Time to Buy/Sell Stock |
| 1043 | Partition Array for Maximum Sum |
| 1235 | Maximum Profit in Job Scheduling |

### Key Question

> Is a local optimal choice guaranteed to lead to a globally optimal solution?

If not, investigate DP.

---

# 🏆 Core 40 DP Problem Set

If you don't want to solve 200+ problems immediately, start with these.

## Level 1 — Foundation

```text
70   Climbing Stairs
746  Min Cost Climbing Stairs
198  House Robber
213  House Robber II
53   Maximum Subarray
152  Maximum Product Subarray
121  Best Time to Buy/Sell Stock
91   Decode Ways
377  Combination Sum IV
```

---

## Level 2 — Knapsack

```text
416  Partition Equal Subset Sum
494  Target Sum
1049 Last Stone Weight II
322  Coin Change
518  Coin Change II
279  Perfect Squares
474  Ones and Zeroes
698  Partition to K Equal Sum Subsets
473  Matchsticks to Square
879  Profitable Schemes
```

---

## Level 3 — Grid

```text
62   Unique Paths
64   Minimum Path Sum
221  Maximal Square
174  Dungeon Game
1463 Cherry Pickup II
329  Longest Increasing Path in a Matrix
```

---

## Level 4 — Subsequence

```text
300  Longest Increasing Subsequence
673  Number of LIS
354  Russian Doll Envelopes
1027 Longest Arithmetic Subsequence
1048 Longest String Chain
```

---

## Level 5 — String DP

```text
1143 Longest Common Subsequence
583  Delete Operation for Two Strings
1092 Shortest Common Supersequence
115  Distinct Subsequences
72   Edit Distance
97   Interleaving String
```

---

## Level 6 — Partition / Interval

```text
139  Word Break
132  Palindrome Partitioning II
312  Burst Balloons
1039 Minimum Score Triangulation
1547 Minimum Cost to Cut a Stick
```

---

## Level 7 — State Machine

```text
122  Stock II
309  Stock Cooldown
714  Stock Transaction Fee
123  Stock III
188  Stock IV
```

---

## Level 8 — Tree

```text
337  House Robber III
124  Binary Tree Maximum Path Sum
543  Diameter of Binary Tree
968  Binary Tree Cameras
834  Sum of Distances in Tree
```

---

## Level 9 — Bitmask

```text
464  Can I Win
473  Matchsticks to Square
698  Partition to K Equal Sum Subsets
847  Shortest Path Visiting All Nodes
1125 Smallest Sufficient Team
1349 Maximum Students Taking Exam
1601 Maximum Achievable Transfer Requests
1723 Minimum Time to Finish All Jobs
1799 Maximize Score After N Operations
```

---

## Level 10 — Advanced

```text
600  Non-negative Integers Without Consecutive Ones
357  Count Numbers With Unique Digits
902  Numbers At Most N Given Digit Set
1012 Numbers With Repeated Digits
2719 Count Integers With Digit Sum in Range

688  Knight Probability in Chessboard
808  Soup Servings
837  New 21 Game

486  Predict the Winner
1140 Stone Game II
1406 Stone Game III
1510 Stone Game IV

1235 Maximum Profit in Job Scheduling
1425 Constrained Subsequence Sum
1696 Jump Game VI
1751 Maximum Number of Events II
```

---

# 🧭 Recommended DP Learning Order

```text
                         DYNAMIC PROGRAMMING
                                  │
          ┌───────────────────────┼───────────────────────┐
          ↓                       ↓                       ↓
       1D DP                  Knapsack                Grid DP
          │                       │                       │
          ↓                       ↓                       ↓
      Take/Skip              0/1 / Unbounded         2D State
          │                       │                       │
          └───────────────────────┼───────────────────────┘
                                  ↓
                         SUBSEQUENCE DP
                         ┌────────┴────────┐
                         ↓                 ↓
                        LIS               LCS
                         │                 │
                         ↓                 ↓
                 Binary Search        String DP
                                           │
                    ┌──────────────────────┼───────────────────┐
                    ↓                      ↓                   ↓
               Palindrome            Edit Distance       Partition
                    │                      │                   │
                    └──────────────────────┼───────────────────┘
                                           ↓
                                     Interval DP
                                           │
                          ┌────────────────┴───────────────┐
                          ↓                                ↓
                       Tree DP                       State Machine
                          │                                │
                          ↓                                ↓
                     Rerooting                        Stock DP
                          │
                          └──────────────┬─────────────────┐
                                         ↓                 ↓
                                     Bitmask            Digit DP
                                         │                 │
                                         └────────┬────────┘
                                                  ↓
                                           Advanced DP
```

---

# 🔥 The Most Important Pattern Buckets

If you remember nothing else, remember these:

```text
1.  Take / Skip
2.  1D DP
3.  Kadane
4.  0/1 Knapsack
5.  Unbounded Knapsack
6.  Grid DP
7.  LIS
8.  LCS
9.  String DP
10. Palindrome DP
11. Partition DP
12. Interval DP
13. State Machine DP
14. Tree DP
15. Rerooting DP
16. Bitmask DP
17. Digit DP
18. Game DP
19. DP + Binary Search
20. DP + Monotonic Queue
21. DP + Graph
22. Probability DP
```

---

# 🧠 Most Important Observational Triggers

| If the problem says... | Think... |
|---|---|
| **ending at `i`** | `dp[i]` |
| **starting from `i`** | suffix DP |
| **take or skip** | 1D / Knapsack |
| **use item once** | 0/1 Knapsack |
| **use item repeatedly** | Unbounded Knapsack |
| **how many ways?** | Counting DP |
| **can we reach?** | Boolean DP |
| **left or right** | Interval DP |
| **split at `k`** | Interval / Partition DP |
| **previous compatible** | Binary Search + DP |
| **every possible root** | Rerooting DP |
| **subset of elements** | Knapsack / Bitmask |
| **small `n`, subsets** | Bitmask DP |
| **digits from `0...N`** | Digit DP |
| **multiple states per day** | State Machine DP |
| **two strings** | LCS / Edit Distance |
| **palindrome** | `dp[l][r]` |
| **tree + choices** | Tree DP |
| **same recursion state repeats** | Memoization |
| **sliding window maximum in recurrence** | Monotonic Deque |
| **intervals + maximum profit** | DP + Binary Search |
| **probability / expected value** | Probability DP |
| **two players optimal** | Game DP |
| **graph + bounded steps/resources** | Graph State DP |

---

# 🎯 Final Strategy

Don't try to memorize:

```text
"LC 312 = Burst Balloons"
```

Instead memorize:

```text
"Interval + split at k"
        ↓
     Interval DP
        ↓
dp[l][r]
        ↓
try every k
```

Similarly:

```text
"LC 416"
        ↓
Subset / Partition
        ↓
0/1 Knapsack
        ↓
dp[target]
```

```text
"LC 1143"
        ↓
Two Strings
        ↓
LCS
        ↓
dp[i][j]
```

```text
"LC 1235"
        ↓
Intervals / Jobs
        ↓
Previous Compatible Job
        ↓
Binary Search + DP
```

```text
"LC 834"
        ↓
Tree
        ↓
Answer for Every Root
        ↓
Rerooting DP
```

```text
"LC 1012"
        ↓
Numbers ≤ N
        ↓
Digit-by-Digit
        ↓
Digit DP
```

