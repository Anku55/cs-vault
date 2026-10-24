
# Graph DSA Roadmap — LeetCode Patterns, Subpatterns & Observations

> Goal: Learn to identify the **graph pattern** before thinking about implementation.
>
> For every pattern, ask:
>
> 1. What exactly are the **vertices**?
> 2. What exactly are the **edges**?
> 3. Is the graph **directed or undirected**?
> 4. Is it **weighted or unweighted**?
> 5. Is the graph **static or changing**?
> 6. Do I need **reachability, shortest path, connectivity, ordering, cycle detection, optimization, or counting**?
> 7. Can I transform the problem into a graph?

---

# 1. Graph Traversal — DFS / BFS

## Core idea

The most basic graph pattern:

```text
Start from a node
→ visit neighbors
→ continue recursively / using queue
→ maintain visited
```

### Observational triggers

Look for:

- "visit all reachable"
- "connected"
- "can reach"
- "number of components"
- "explore region"
- "spread"
- "mark all cells"
- "find whether path exists"

### Problems

| # | LeetCode |
|---|---|
| 1 | [1971] Find if Path Exists in Graph |
| 2 | [733] Flood Fill |
| 3 | [200] Number of Islands |
| 4 | [695] Max Area of Island |
| 5 | [463] Island Perimeter |
| 6 | [841] Keys and Rooms |
| 7 | [547] Number of Provinces |
| 8 | [130] Surrounded Regions |
| 9 | [417] Pacific Atlantic Water Flow |
| 10 | [529] Minesweeper |
| 11 | [733] Flood Fill |
| 12 | [841] Keys and Rooms |
| 13 | [797] All Paths From Source to Target |
| 14 | [1020] Number of Enclaves |
| 15 | [1254] Number of Closed Islands |
| 16 | [1905] Count Sub Islands |
| 17 | [463] Island Perimeter |
| 18 | [305] Number of Islands II |
| 19 | [694] Number of Distinct Islands |
| 20 | [711] Number of Distinct Islands II |

---

# 2. Grid as a Graph

This is one of the most important graph observations.

A matrix is often secretly:

```text
cell = node
adjacent cell = edge
```

Usually 4 directions:

```text
up
down
left
right
```

Sometimes 8 directions.

### Observational triggers

If you see:

- matrix
- grid
- island
- region
- neighboring cells
- spread
- shortest path through cells
- connected area

immediately ask:

> "Is every cell a graph node?"

### Problems

- [200] Number of Islands
- [695] Max Area of Island
- [733] Flood Fill
- [130] Surrounded Regions
- [417] Pacific Atlantic Water Flow
- [994] Rotting Oranges
- [1091] Shortest Path in Binary Matrix
- [1162] As Far from Land as Possible
- [542] 01 Matrix
- [1926] Nearest Exit from Entrance in Maze
- [490] The Maze
- [505] The Maze II
- [499] The Maze III
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [864] Shortest Path to Get All Keys
- [1036] Escape a Large Maze
- [934] Shortest Bridge
- [827] Making A Large Island
- [1020] Number of Enclaves
- [1254] Number of Closed Islands
- [1905] Count Sub Islands
- [130] Surrounded Regions
- [286] Walls and Gates
- [529] Minesweeper
- [1219] Path with Maximum Gold

---

# 3. Connected Components

## Pattern

For an undirected graph:

```text
for every node:
    if not visited:
        DFS/BFS
        components++
```

### Observational triggers

- "number of groups"
- "number of clusters"
- "how many disconnected..."
- "minimum operations to connect..."
- "same component"
- "friends"
- "provinces"

### Problems

- [547] Number of Provinces
- [323] Number of Connected Components in an Undirected Graph
- [200] Number of Islands
- [695] Max Area of Island
- [684] Redundant Connection
- [1319] Number of Operations to Make Network Connected
- [990] Satisfiability of Equality Equations
- [721] Accounts Merge
- [1202] Smallest String With Swaps
- [947] Most Stones Removed with Same Row or Column
- [1020] Number of Enclaves
- [1254] Number of Closed Islands
- [1905] Count Sub Islands
- [305] Number of Islands II
- [1101] The Earliest Moment When Everyone Become Friends

---

# 4. Number of Islands / Region Counting

This deserves its own pattern because it appears constantly.

### Variations

```text
count components
maximum component
component size
component perimeter
component shape
component comparison
```

### Problems

- [200] Number of Islands
- [695] Max Area of Island
- [463] Island Perimeter
- [694] Number of Distinct Islands
- [711] Number of Distinct Islands II
- [1020] Number of Enclaves
- [1254] Number of Closed Islands
- [1905] Count Sub Islands
- [827] Making A Large Island
- [305] Number of Islands II
- [934] Shortest Bridge
- [1568] Minimum Number of Days to Disconnect Island

---

# 5. Multi-Source BFS

## Core observation

If **multiple nodes start spreading simultaneously**, initialize the queue with all of them.

```text
queue = all sources

while queue:
    expand simultaneously
```

### Observational triggers

- "nearest"
- "distance from nearest..."
- "spread simultaneously"
- "all rotten oranges"
- "nearest 0"
- "nearest land"
- "multiple starting points"

### Problems

- [994] Rotting Oranges
- [542] 01 Matrix
- [1162] As Far from Land as Possible
- [286] Walls and Gates
- [317] Shortest Distance from All Buildings
- [1765] Map of Highest Peak
- [934] Shortest Bridge
- [2812] Find the Safest Path in a Grid
- [542] 01 Matrix
- [1091] Shortest Path in Binary Matrix
- [863] All Nodes Distance K in Binary Tree
- [847] Shortest Path Visiting All Nodes

### Key observation

Do NOT run BFS separately from every source.

Instead:

```text
all sources → queue initially
```

This changes:

```text
O(S * V)
```

into roughly:

```text
O(V + E)
```

for standard multi-source BFS.

---

# 6. BFS Shortest Path in Unweighted Graph

## Core rule

For an unweighted graph:

```text
BFS = shortest number of edges
```

### Problems

- [127] Word Ladder
- [752] Open the Lock
- [1091] Shortest Path in Binary Matrix
- [1926] Nearest Exit from Entrance in Maze
- [433] Minimum Genetic Mutation
- [815] Bus Routes
- [847] Shortest Path Visiting All Nodes
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [909] Snakes and Ladders
- [1036] Escape a Large Maze
- [542] 01 Matrix
- [994] Rotting Oranges
- [1162] As Far from Land as Possible
- [863] All Nodes Distance K in Binary Tree
- [934] Shortest Bridge
- [752] Open the Lock
- [773] Sliding Puzzle
- [864] Shortest Path to Get All Keys
- [847] Shortest Path Visiting All Nodes

---

# 7. BFS on an Implicit Graph

Sometimes edges aren't explicitly given.

Example:

```text
"hit"
→ "hot"
→ "dot"
→ "dog"
→ "cog"
```

The graph is generated dynamically.

### Observational trigger

Ask:

> "Can I define a state, and generate all states reachable in one move?"

Then:

```text
state = node
valid transition = edge
```

### Problems

- [127] Word Ladder
- [126] Word Ladder II
- [752] Open the Lock
- [433] Minimum Genetic Mutation
- [773] Sliding Puzzle
- [909] Snakes and Ladders
- [864] Shortest Path to Get All Keys
- [847] Shortest Path Visiting All Nodes
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [1036] Escape a Large Maze

---

# 8. BFS with State

One of the most important advanced graph patterns.

Normal BFS:

```text
visited[node]
```

State BFS:

```text
visited[node][state]
```

Examples:

```text
(node, keys)
(node, remaining_eliminations)
(node, parity)
(node, mask)
(node, fuel)
(node, stops)
```

### Problems

- [864] Shortest Path to Get All Keys
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [787] Cheapest Flights Within K Stops
- [847] Shortest Path Visiting All Nodes
- [1654] Minimum Jumps to Reach Home
- [815] Bus Routes
- [1129] Shortest Path with Alternating Colors
- [934] Shortest Bridge
- [1036] Escape a Large Maze
- [1368] Minimum Cost to Make at Least One Valid Path in a Grid

---

# 9. DFS for Path Enumeration

Different from simple reachability.

Here you need to explore **all possible paths**.

### Problems

- [797] All Paths From Source to Target
- [980] Unique Paths III
- [212] Word Search II
- [79] Word Search
- [51] N-Queens
- [37] Sudoku Solver
- [46] Permutations
- [77] Combinations
- [39] Combination Sum
- [22] Generate Parentheses

Graph connection:

```text
current state
→ choose next state
→ recurse
→ backtrack
```

---

# 10. Cycle Detection — Undirected Graph

## DFS observation

For undirected graph:

```text
neighbor == parent
```

is NOT a cycle.

But:

```text
visited neighbor != parent
```

means cycle.

### Problems

- [684] Redundant Connection
- [261] Graph Valid Tree
- [323] Number of Connected Components in an Undirected Graph
- [785] Is Graph Bipartite?
- [886] Possible Bipartition
- [207] Course Schedule
- [1059] All Paths from Source Lead to Destination
- [1559] Detect Cycles in 2D Grid
- [2360] Longest Cycle in a Graph
- [2685] Count the Number of Complete Components

---

# 11. Cycle Detection — Directed Graph

Directed graphs require a different concept.

Use:

```text
0 = unvisited
1 = currently visiting
2 = completely processed
```

A cycle exists if:

```text
neighbor is currently visiting
```

### Problems

- [207] Course Schedule
- [210] Course Schedule II
- [802] Find Eventual Safe States
- [1059] All Paths from Source Lead to Destination
- [2360] Longest Cycle in a Graph
- [2050] Parallel Courses III
- [1462] Course Schedule IV
- [1559] Detect Cycles in 2D Grid
- [2115] Find All Possible Recipes from Given Supplies
- [2192] All Ancestors of a Node in a Directed Acyclic Graph

---

# 12. Bipartite Graph

## Core observation

Can we divide vertices into two groups such that:

```text
every edge connects opposite groups
```

Equivalent:

```text
2-coloring
```

Use:

```text
color[node] = 0 / 1
```

### Problems

- [785] Is Graph Bipartite?
- [886] Possible Bipartition
- [1042] Flower Planting With No Adjacent
- [2493] Divide Nodes Into the Maximum Number of Groups
- [1222] Queens That Can Attack the King
- [2039] The Time When the Network Becomes Idle
- [785] Is Graph Bipartite?
- [886] Possible Bipartition

### Key trick

Odd cycle ⇒ not bipartite.

---

# 13. Topological Sort — Kahn's Algorithm

## Observation

If you see:

```text
prerequisite
dependency
before
after
course
build order
task dependency
```

think:

> Directed graph + topological ordering.

Kahn:

```text
indegree = 0
→ queue
→ remove node
→ decrease neighbors' indegree
```

### Problems

- [207] Course Schedule
- [210] Course Schedule II
- [310] Minimum Height Trees
- [444] Sequence Reconstruction
- [630] Course Schedule III
- [1136] Parallel Courses
- [2050] Parallel Courses III
- [2115] Find All Possible Recipes from Given Supplies
- [2392] Build a Matrix With Conditions
- [1857] Largest Color Value in a Directed Graph
- [2192] All Ancestors of a Node in a Directed Acyclic Graph
- [1203] Sort Items by Groups Respecting Dependencies
- [1591] Strange Printer II
- [2603] Collect Coins in a Tree
- [1462] Course Schedule IV

---

# 14. Topological Sort — DFS

Use DFS postorder:

```text
DFS(node)
→ DFS(neighbors)
→ push node
```

Reverse result.

### Problems

- [207] Course Schedule
- [210] Course Schedule II
- [802] Find Eventual Safe States
- [1857] Largest Color Value in a Directed Graph
- [2192] All Ancestors of a Node in a Directed Acyclic Graph
- [1059] All Paths from Source Lead to Destination
- [2050] Parallel Courses III
- [2115] Find All Possible Recipes from Given Supplies

---

# 15. DAG Dynamic Programming

A DAG is a graph where topological order gives a natural DP order.

```text
topological order
→ dp[node]
→ transition to neighbors
```

### Problems

- [329] Longest Increasing Path in a Matrix
- [787] Cheapest Flights Within K Stops
- [2050] Parallel Courses III
- [1857] Largest Color Value in a Directed Graph
- [1976] Number of Ways to Arrive at Destination
- [2322] Minimum Score of a Path Between Two Cities
- [1786] Number of Restricted Paths From First to Last Node
- [797] All Paths From Source to Target
- [1136] Parallel Courses
- [1377] Frog Position After T Seconds
- [1514] Path with Maximum Probability

---

# 16. Shortest Path — Which Algorithm?

This is a critical recognition pattern.

| Graph | Algorithm |
|---|---|
| Unweighted | BFS |
| Weight 0/1 | 0-1 BFS |
| Positive weights | Dijkstra |
| Negative edges | Bellman-Ford |
| DAG | Topological DP |
| All pairs | Floyd-Warshall |
| Multiple sources | Multi-source BFS / Dijkstra |
| Exactly K edges | DP / Bellman-Ford style |
| At most K stops | State DP / modified Dijkstra |

---

# 17. Dijkstra — Basic

## Condition

Use when edge weights are:

```text
non-negative
```

Typical implementation:

```text
min heap
```

### Problems

- [743] Network Delay Time
- [1514] Path with Maximum Probability
- [1631] Path With Minimum Effort
- [787] Cheapest Flights Within K Stops
- [882] Reachable Nodes In Subdivided Graph
- [1334] Find the City With the Smallest Number of Neighbors at a Threshold Distance
- [1976] Number of Ways to Arrive at Destination
- [2045] Second Minimum Time to Reach Destination
- [2203] Minimum Weighted Subgraph With the Required Paths
- [2290] Minimum Obstacle Removal to Reach the Corner
- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [778] Swim in Rising Water
- [1928] Minimum Cost to Reach Destination in Time
- [2662] Minimum Cost of a Path With Special Roads
- [2699] Modify Graph Edge Weights
- [2812] Find the Safest Path in a Grid
- [3123] Find Edges in Shortest Paths

---

# 18. Dijkstra + State

This is a very important advanced pattern.

Instead of:

```text
dist[node]
```

use:

```text
dist[node][state]
```

Examples:

```text
node + remaining fuel
node + number of discounts
node + time
node + parity
node + mask
```

### Problems

- [787] Cheapest Flights Within K Stops
- [1928] Minimum Cost to Reach Destination in Time
- [2290] Minimum Obstacle Removal to Reach the Corner
- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [2662] Minimum Cost of a Path With Special Roads
- [2093] Minimum Cost to Reach City With Discounts
- [2203] Minimum Weighted Subgraph With the Required Paths
- [2714] Find Shortest Path with K Hops
- [2045] Second Minimum Time to Reach Destination
- [882] Reachable Nodes In Subdivided Graph

---

# 19. 0-1 BFS

## Observation

Edge weights are only:

```text
0 or 1
```

Then:

```text
deque
```

instead of priority queue.

Rules:

```text
weight 0 → push_front
weight 1 → push_back
```

### Problems

- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner
- [1263] Minimum Moves to Move a Box to Their Target Location
- [934] Shortest Bridge
- [542] 01 Matrix
- [1210] Minimum Moves to Reach Target with Rotations
- [1928] Minimum Cost to Reach Destination in Time
- [2812] Find the Safest Path in a Grid

> The cleanest LeetCode practice problems here are **1368** and **2290**.

---

# 20. Bellman-Ford / K-Edge Shortest Path

## Observation

When the problem says:

```text
at most K edges
exactly K edges
negative edges
relax edges repeatedly
```

Bellman-Ford-style DP becomes relevant.

### Problems

- [787] Cheapest Flights Within K Stops
- [743] Network Delay Time
- [787] Cheapest Flights Within K Stops
- [2045] Second Minimum Time to Reach Destination
- [1928] Minimum Cost to Reach Destination in Time
- [2699] Modify Graph Edge Weights
- [2203] Minimum Weighted Subgraph With the Required Paths

---

# 21. Floyd-Warshall / All-Pairs Shortest Path

## Observation

If you need:

```text
distance between every pair
```

think:

```text
dp[k][i][j]
```

or in-place:

```text
dist[i][j]
```

### Problems

- [1334] Find the City With the Smallest Number of Neighbors at a Threshold Distance
- [1462] Course Schedule IV
- [399] Evaluate Division
- [2959] Number of Possible Sets of Closing Branches
- [2642] Design Graph With Shortest Path Calculator
- [1615] Maximal Network Rank
- [2976] Minimum Cost to Convert String I
- [743] Network Delay Time

---

# 22. Minimum Spanning Tree — Kruskal

## Core idea

Sort edges by weight:

```text
smallest edge first
```

Use DSU to avoid cycles.

```text
if(u and v are different components):
    union
```

### Problems

- [1584] Min Cost to Connect All Points
- [1135] Connecting Cities With Minimum Cost
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [1697] Checking Existence of Edge Length Limited Paths
- [1168] Optimize Water Distribution in a Village
- [261] Graph Valid Tree
- [684] Redundant Connection
- [1101] The Earliest Moment When Everyone Become Friends
- [2421] Number of Good Paths
- [2092] Find All People With Secret

---

# 23. Minimum Spanning Tree — Prim

## Observation

Start from one vertex and repeatedly take the cheapest edge that connects the current tree to a new vertex.

### Problems

- [1584] Min Cost to Connect All Points
- [1135] Connecting Cities With Minimum Cost
- [1168] Optimize Water Distribution in a Village
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1928] Minimum Cost to Reach Destination in Time
- [2093] Minimum Cost to Reach City With Discounts

### Recognition

If problem asks:

> "Connect all nodes with minimum total cost."

think:

```text
MST
```

not shortest path.

---

# 24. DSU / Union-Find

## Core operations

```text
find(x)
union(x, y)
```

Useful for:

- connectivity
- merging groups
- cycle detection
- dynamic connectivity
- MST
- equivalence classes

### Problems

- [547] Number of Provinces
- [684] Redundant Connection
- [685] Redundant Connection II
- [1319] Number of Operations to Make Network Connected
- [721] Accounts Merge
- [947] Most Stones Removed with Same Row or Column
- [990] Satisfiability of Equality Equations
- [1061] Lexicographically Smallest Equivalent String
- [1202] Smallest String With Swaps
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [1584] Min Cost to Connect All Points
- [1697] Checking Existence of Edge Length Limited Paths
- [1722] Minimize Hamming Distance After Swap Operations
- [2421] Number of Good Paths
- [2492] Minimum Score of a Path Between Two Cities
- [2709] Greatest Common Divisor Traversal
- [305] Number of Islands II
- [1101] The Earliest Moment When Everyone Become Friends
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph
- [952] Largest Component Size by Common Factor
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable

---

# 25. DSU + Grid

Very important variation.

Instead of DFS-ing after every update:

```text
activate cell
→ union with active neighbors
→ update component count
```

### Problems

- [305] Number of Islands II
- [827] Making A Large Island
- [959] Regions Cut By Slashes
- [803] Bricks Falling When Hit
- [947] Most Stones Removed with Same Row or Column
- [200] Number of Islands
- [130] Surrounded Regions

---

# 26. DSU + Equivalence / String

## Observation

If statements say:

```text
a == b
x equivalent to y
swap allowed
same group
```

DSU is often the hidden structure.

### Problems

- [990] Satisfiability of Equality Equations
- [1061] Lexicographically Smallest Equivalent String
- [721] Accounts Merge
- [1202] Smallest String With Swaps
- [1722] Minimize Hamming Distance After Swap Operations
- [737] Sentence Similarity II
- [839] Similar String Groups
- [947] Most Stones Removed with Same Row or Column

---

# 27. Redundant Edge / Find Cycle Using DSU

### Observation

If adding an edge connects two nodes already connected:

```text
cycle
```

### Problems

- [684] Redundant Connection
- [685] Redundant Connection II
- [261] Graph Valid Tree
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [1319] Number of Operations to Make Network Connected
- [1584] Min Cost to Connect All Points

---

# 28. Bridges — Tarjan

A bridge is an edge whose removal disconnects the graph.

### Observation

DFS maintains:

```text
tin[u]
low[u]
```

For edge:

```text
u → v
```

if:

```text
low[v] > tin[u]
```

then:

```text
(u,v) is a bridge
```

### Problems

- [1192] Critical Connections in a Network
- [1568] Minimum Number of Days to Disconnect Island
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [2492] Minimum Score of a Path Between Two Cities
- [2092] Find All People With Secret
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph

> **1192 is the must-do Tarjan problem.**

---

# 29. Articulation Points / Cut Vertices

A vertex whose removal disconnects the graph.

### Observation

Same Tarjan framework:

```text
tin[]
low[]
```

but reason about vertices instead of edges.

### LeetCode practice

- [1568] Minimum Number of Days to Disconnect Island
- [1192] Critical Connections in a Network
- [2322] Minimum Score After Removals on a Tree
- [2492] Minimum Score of a Path Between Two Cities
- [1489] Critical and Pseudo-Critical Edges in MST

LeetCode has relatively limited direct articulation-point problems, so **1192 + Tarjan implementation practice** is particularly valuable.

---

# 30. Strongly Connected Components

## Directed graph

Two vertices belong to the same SCC if:

```text
u → v
and
v → u
```

Common algorithms:

```text
Kosaraju
Tarjan
```

### Problems / practice

- [802] Find Eventual Safe States
- [207] Course Schedule
- [210] Course Schedule II
- [1192] Critical Connections in a Network
- [2127] Maximum Employees to Be Invited to a Meeting
- [2360] Longest Cycle in a Graph
- [1557] Minimum Number of Vertices to Reach All Nodes
- [2192] All Ancestors of a Node in a Directed Acyclic Graph

> LeetCode does not have as many explicit SCC questions as competitive-programming platforms. Learn the algorithm independently after these problems.

---

# 31. Eventual Safe States

This is a beautiful graph observation.

A node is safe if:

```text
every path eventually reaches terminal node
```

Equivalent:

```text
node does not belong to / lead into a cycle
```

Can solve using:

```text
reverse graph + outdegree
```

### Problems

- [802] Find Eventual Safe States
- [207] Course Schedule
- [210] Course Schedule II
- [1059] All Paths from Source Lead to Destination
- [2360] Longest Cycle in a Graph
- [2127] Maximum Employees to Be Invited to a Meeting

---

# 32. Reverse Graph Trick

Extremely important observational pattern.

Instead of asking:

```text
Who can reach X?
```

reverse edges and ask:

```text
Who is reachable from X?
```

### Problems

- [802] Find Eventual Safe States
- [310] Minimum Height Trees
- [2192] All Ancestors of a Node in a Directed Acyclic Graph
- [2127] Maximum Employees to Be Invited to a Meeting
- [1462] Course Schedule IV
- [1514] Path with Maximum Probability
- [1976] Number of Ways to Arrive at Destination

---

# 33. Shortest Path with Multiple Sources — Dijkstra

Multi-source isn't limited to BFS.

If sources have weighted edges:

```text
push every source into min heap
```

### Problems

- [1514] Path with Maximum Probability
- [2812] Find the Safest Path in a Grid
- [743] Network Delay Time
- [1334] Find the City With the Smallest Number of Neighbors at a Threshold Distance
- [2203] Minimum Weighted Subgraph With the Required Paths

---

# 34. Maximum Probability Path

Instead of:

```text
dist[u] + weight
```

use:

```text
prob[u] * edgeProbability
```

and maximize.

### Problems

- [1514] Path with Maximum Probability
- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid

### Observation

Sometimes "shortest path" is not literally minimum sum.

The objective could be:

```text
maximize product
maximize minimum edge
minimize maximum edge
```

Still a graph optimization problem.

---

# 35. Minimax Path

## Observation

Path cost is:

```text
max(edge costs on path)
```

and you minimize that maximum.

### Problems

- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner

Possible approaches:

```text
Dijkstra
Binary Search + BFS/DFS
MST
```

---

# 36. Binary Search on Answer + Graph Traversal

This is a major hybrid pattern.

Suppose answer is:

```text
minimum X
```

Ask:

> "If X is fixed, can I check whether a valid path exists?"

Then:

```text
binary search X
+
BFS / DFS / Dijkstra
```

### Problems

- [778] Swim in Rising Water
- [1631] Path With Minimum Effort
- [1102] Path With Maximum Minimum Value
- [1870] Minimum Speed to Arrive on Time
- [1552] Magnetic Force Between Two Balls
- [2812] Find the Safest Path in a Grid
- [774] Minimize Max Distance to Gas Station
- [2616] Minimize the Maximum Difference of Pairs

---

# 37. Shortest Path + Binary Search

### Pattern

```text
binary search answer
→ graph feasibility check
```

### Problems

- [778] Swim in Rising Water
- [1631] Path With Minimum Effort
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner

---

# 38. Shortest Path with K Stops

This is a classic "graph + state" problem.

Normal shortest path:

```text
dist[node]
```

But here:

```text
dist[node][stops]
```

### Problems

- [787] Cheapest Flights Within K Stops
- [2045] Second Minimum Time to Reach Destination
- [1928] Minimum Cost to Reach Destination in Time
- [2714] Find Shortest Path with K Hops
- [2093] Minimum Cost to Reach City With Discounts

---

# 39. Cheapest Path with Resource Constraint

State becomes:

```text
(node, resource)
```

Examples:

```text
(node, time)
(node, fuel)
(node, stops)
(node, coupons)
```

### Problems

- [787] Cheapest Flights Within K Stops
- [1928] Minimum Cost to Reach Destination in Time
- [2093] Minimum Cost to Reach City With Discounts
- [871] Minimum Number of Refueling Stops
- [2577] Minimum Time to Visit a Cell In a Grid
- [2662] Minimum Cost of a Path With Special Roads

---

# 40. Graph + Bitmask

Extremely important advanced pattern.

State:

```text
(node, mask)
```

where mask represents visited/collected entities.

### Problems

- [847] Shortest Path Visiting All Nodes
- [864] Shortest Path to Get All Keys
- [980] Unique Paths III
- [1494] Parallel Courses II
- [1125] Smallest Sufficient Team
- [943] Find the Shortest Superstring
- [691] Stickers to Spell Word
- [1655] Distribute Repeating Integers
- [1349] Maximum Students Taking Exam
- [1434] Number of Ways to Wear Different Hats to Each Other
- [1723] Find Minimum Time to Finish All Jobs

---

# 41. TSP / Hamiltonian Path on Graph

## Observation

Need to visit:

```text
all nodes exactly / at least once
```

and N is small.

Then:

```text
dp[mask][node]
```

### Problems

- [847] Shortest Path Visiting All Nodes
- [943] Find the Shortest Superstring
- [980] Unique Paths III
- [1494] Parallel Courses II
- [1125] Smallest Sufficient Team
- [1723] Find Minimum Time to Finish All Jobs

---

# 42. Graph + Bitmask BFS

Very clean pattern:

```text
BFS state:
(node, mask)
```

### Problems

- [847] Shortest Path Visiting All Nodes
- [864] Shortest Path to Get All Keys
- [980] Unique Paths III
- [773] Sliding Puzzle
- [943] Find the Shortest Superstring

---

# 43. Graph Coloring

### Basic idea

Assign colors while respecting constraints.

### Problems

- [785] Is Graph Bipartite?
- [886] Possible Bipartition
- [1042] Flower Planting With No Adjacent
- [2493] Divide Nodes Into the Maximum Number of Groups
- [1591] Strange Printer II
- [1948] Delete Duplicate Folders in System
- [51] N-Queens
- [37] Sudoku Solver

---

# 44. Constraint Graph

Sometimes the graph is built from constraints.

Example:

```text
A < B
B < C
```

or:

```text
x == y
x != y
```

### Problems

- [990] Satisfiability of Equality Equations
- [399] Evaluate Division
- [1061] Lexicographically Smallest Equivalent String
- [737] Sentence Similarity II
- [886] Possible Bipartition
- [269] Alien Dictionary
- [2050] Parallel Courses III
- [1462] Course Schedule IV

---

# 45. Evaluate Division / Weighted Graph

## Observation

Equation:

```text
a / b = 2
```

means:

```text
a → b : 2
b → a : 1/2
```

Then query becomes graph traversal.

### Problems

- [399] Evaluate Division
- [2307] Check for Contradictions in Equations
- [2311] Longest Binary Subsequence Less Than or Equal to K
- [737] Sentence Similarity II
- [1061] Lexicographically Smallest Equivalent String

---

# 46. Graph Transformation / Word Ladder

The important observation:

```text
word = node
one-character transformation = edge
```

### Problems

- [127] Word Ladder
- [126] Word Ladder II
- [433] Minimum Genetic Mutation
- [2059] Minimum Operations to Convert Number
- [752] Open the Lock
- [773] Sliding Puzzle
- [127] Word Ladder
- [433] Minimum Genetic Mutation

---

# 47. Bidirectional BFS

Useful when:

```text
source → target
```

and branching factor is high.

Instead of exploring:

```text
O(b^d)
```

from one side, search from both.

### Problems

- [127] Word Ladder
- [126] Word Ladder II
- [752] Open the Lock
- [773] Sliding Puzzle
- [1036] Escape a Large Maze
- [1654] Minimum Jumps to Reach Home
- [2059] Minimum Operations to Convert Number

---

# 48. Maze Graph

The cells themselves are graph nodes, but movement may have special rules.

### Problems

- [490] The Maze
- [505] The Maze II
- [499] The Maze III
- [1926] Nearest Exit from Entrance in Maze
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [1210] Minimum Moves to Reach Target with Rotations
- [1263] Minimum Moves to Move a Box to Their Target Location
- [773] Sliding Puzzle

---

# 49. Graph with Moving Object State

The graph node isn't just:

```text
position
```

It may be:

```text
(position, direction)
(position, orientation)
(position, box position)
```

### Problems

- [1210] Minimum Moves to Reach Target with Rotations
- [1263] Minimum Moves to Move a Box to Their Target Location
- [864] Shortest Path to Get All Keys
- [1654] Minimum Jumps to Reach Home
- [773] Sliding Puzzle
- [752] Open the Lock

---

# 50. Graph + Direction State

Sometimes reaching the same cell from different directions is different.

### Problems

- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [1210] Minimum Moves to Reach Target with Rotations
- [505] The Maze II
- [499] The Maze III
- [2577] Minimum Time to Visit a Cell In a Grid

---

# 51. Eulerian Path / Trail

## Observation

Need to use:

```text
every edge exactly once
```

Think:

```text
Eulerian path
```

### Problems

- [332] Reconstruct Itinerary
- [753] Cracking the Safe
- [2097] Valid Arrangement of Pairs
- [1743] Restore the Array From Adjacent Pairs
- [2097] Valid Arrangement of Pairs
- [491] Non-decreasing Subsequences
- [332] Reconstruct Itinerary

### Key condition

Eulerian path concerns **edges**, not vertices.

---

# 52. Reconstruct Itinerary / Hierholzer

Classic pattern:

```text
use every directed edge exactly once
```

### Problems

- [332] Reconstruct Itinerary
- [2097] Valid Arrangement of Pairs
- [753] Cracking the Safe
- [1743] Restore the Array From Adjacent Pairs
- [2097] Valid Arrangement of Pairs

---

# 53. In-Degree / Out-Degree Observation

Before choosing an algorithm, count:

```text
indegree
outdegree
```

Useful for:

- topological sort
- Eulerian path
- safe states
- dependency resolution
- source/sink identification

### Problems

- [207] Course Schedule
- [210] Course Schedule II
- [332] Reconstruct Itinerary
- [802] Find Eventual Safe States
- [2097] Valid Arrangement of Pairs
- [1557] Minimum Number of Vertices to Reach All Nodes
- [1136] Parallel Courses
- [2115] Find All Possible Recipes from Given Supplies

---

# 54. Source / Sink Graph Observation

In a DAG:

```text
indegree = 0 → source
outdegree = 0 → sink
```

### Problems

- [1557] Minimum Number of Vertices to Reach All Nodes
- [310] Minimum Height Trees
- [802] Find Eventual Safe States
- [207] Course Schedule
- [210] Course Schedule II
- [2192] All Ancestors of a Node in a Directed Acyclic Graph
- [1059] All Paths from Source Lead to Destination

---

# 55. Tree as a Graph

Every tree is an undirected graph with:

```text
N nodes
N - 1 edges
connected
acyclic
```

This lets you use graph techniques.

### Problems

- [261] Graph Valid Tree
- [310] Minimum Height Trees
- [834] Sum of Distances in Tree
- [1245] Tree Diameter
- [1519] Number of Nodes in the Sub-Tree With the Same Label
- [1443] Minimum Time to Collect All Apples in a Tree
- [1376] Time Needed to Inform All Employees
- [2385] Amount of Time for Binary Tree to Be Infected
- [863] All Nodes Distance K in Binary Tree

---

# 56. Tree Diameter

## Observation

Diameter:

```text
longest path between two nodes
```

Common trick:

```text
BFS/DFS from arbitrary node
→ farthest A

BFS/DFS from A
→ farthest B

A → B = diameter
```

### Problems

- [1245] Tree Diameter
- [543] Diameter of Binary Tree
- [687] Longest Univalue Path
- [1522] Diameter of N-Ary Tree
- [2246] Longest Path With Different Adjacent Characters
- [3203] Find Minimum Diameter After Merging Two Trees

---

# 57. Tree Rerooting / Sum of Distances

A major graph/tree observation:

Calculate answer for one root, then derive answers for neighboring roots.

### Problems

- [834] Sum of Distances in Tree
- [310] Minimum Height Trees
- [2581] Count Number of Possible Root Nodes
- [2049] Count Nodes With the Highest Score
- [2925] Maximum Score After Applying Operations on a Tree
- [2858] Minimum Edge Reversals So Every Node Is Reachable
- [3241] Time Taken to Mark All Nodes

---

# 58. Tree DP

Although trees are technically graphs, tree DP is a major graph pattern.

### Problems

- [337] House Robber III
- [124] Binary Tree Maximum Path Sum
- [968] Binary Tree Cameras
- [543] Diameter of Binary Tree
- [1372] Longest ZigZag Path in a Binary Tree
- [979] Distribute Coins in Binary Tree
- [834] Sum of Distances in Tree
- [1519] Number of Nodes in the Sub-Tree With the Same Label
- [1443] Minimum Time to Collect All Apples in a Tree
- [2246] Longest Path With Different Adjacent Characters
- [2458] Height of Binary Tree After Subtree Removal Queries
- [2641] Cousins in Binary Tree II
- [2925] Maximum Score After Applying Operations on a Tree

---

# 59. Tree + Rerooting

### Problems

- [834] Sum of Distances in Tree
- [2581] Count Number of Possible Root Nodes
- [2858] Minimum Edge Reversals So Every Node Is Reachable
- [3241] Time Taken to Mark All Nodes
- [310] Minimum Height Trees
- [2246] Longest Path With Different Adjacent Characters

---

# 60. Functional Graph

A functional graph means:

```text
every node has at most/exactly one outgoing edge
```

Example:

```text
next[i]
```

This creates:

```text
cycle
+
trees feeding into cycle
```

### Problems

- [2360] Longest Cycle in a Graph
- [2876] Count Visited Nodes in a Directed Graph
- [2359] Find Closest Node to Given Two Nodes
- [2127] Maximum Employees to Be Invited to a Meeting
- [142] Linked List Cycle II
- [457] Circular Array Loop
- [2204] Distance to a Cycle in Undirected Graph

---

# 61. Functional Graph — Cycle + Tail

Classic structure:

```text
      → cycle
     /
tree
```

Useful techniques:

```text
DFS state
visited timestamp
cycle detection
reverse tree DP
```

### Problems

- [2360] Longest Cycle in a Graph
- [2876] Count Visited Nodes in a Directed Graph
- [2127] Maximum Employees to Be Invited to a Meeting
- [2359] Find Closest Node to Given Two Nodes
- [457] Circular Array Loop
- [2204] Distance to a Cycle in Undirected Graph

---

# 62. Longest Cycle in Directed Graph

### Problems

- [2360] Longest Cycle in a Graph
- [2876] Count Visited Nodes in a Directed Graph
- [2127] Maximum Employees to Be Invited to a Meeting
- [457] Circular Array Loop
- [1059] All Paths from Source Lead to Destination

---

# 63. Graph + Greedy

Some graph problems look like greedy but require graph structure.

### Problems

- [630] Course Schedule III
- [1402] Reducing Dishes
- [1584] Min Cost to Connect All Points
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [871] Minimum Number of Refueling Stops
- [1383] Maximum Performance of a Team
- [1514] Path with Maximum Probability

---

# 64. Graph + Heap

Heap often appears with:

```text
Dijkstra
Prim
best next candidate
priority expansion
```

### Problems

- [743] Network Delay Time
- [1514] Path with Maximum Probability
- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [882] Reachable Nodes In Subdivided Graph
- [1334] Find the City With the Smallest Number of Neighbors at a Threshold Distance
- [1928] Minimum Cost to Reach Destination in Time
- [2203] Minimum Weighted Subgraph With the Required Paths
- [2290] Minimum Obstacle Removal to Reach the Corner
- [2699] Modify Graph Edge Weights
- [2812] Find the Safest Path in a Grid

---

# 65. Graph + Binary Search + BFS

### Problems

- [778] Swim in Rising Water
- [1631] Path With Minimum Effort
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [2616] Minimize the Maximum Difference of Pairs

---

# 66. Graph + Binary Search + Dijkstra

### Problems

- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [2812] Find the Safest Path in a Grid
- [1928] Minimum Cost to Reach Destination in Time
- [2662] Minimum Cost of a Path With Special Roads

---

# 67. Minimum Cut / Max Flow

LeetCode has relatively few direct max-flow questions, but this is an important advanced graph family.

### Problems

- [1820] Maximum Number of Accepted Invitations
- [1349] Maximum Students Taking Exam
- [1947] Maximum Compatibility Score Sum
- [1595] Minimum Cost to Connect Two Groups of Points
- [2172] Maximum AND Sum of Array
- [1254] Number of Closed Islands
- [956] Tallest Billboard

For serious max-flow preparation, LeetCode alone is not sufficient; supplement with standard **Dinic / Edmonds-Karp** problems.

---

# 68. Bipartite Matching

## Observation

Two disjoint sets:

```text
left nodes
right nodes
```

Need maximum matching.

### Problems

- [1820] Maximum Number of Accepted Invitations
- [1349] Maximum Students Taking Exam
- [1947] Maximum Compatibility Score Sum
- [1595] Minimum Cost to Connect Two Groups of Points
- [1125] Smallest Sufficient Team
- [1434] Number of Ways to Wear Different Hats to Each Other

---

# 69. Assignment / Matching with Bitmask

For small N, instead of full matching:

```text
dp[mask]
```

### Problems

- [1947] Maximum Compatibility Score Sum
- [1125] Smallest Sufficient Team
- [1434] Number of Ways to Wear Different Hats to Each Other
- [1349] Maximum Students Taking Exam
- [1595] Minimum Cost to Connect Two Groups of Points
- [1494] Parallel Courses II

---

# 70. Alien Dictionary / Ordering from Comparisons

Classic graph construction:

```text
word1
word2
```

Find first differing character:

```text
c1 → c2
```

Then topological sort.

### Problems

- [269] Alien Dictionary
- [953] Verifying an Alien Dictionary
- [1203] Sort Items by Groups Respecting Dependencies
- [1857] Largest Color Value in a Directed Graph
- [210] Course Schedule II

> **269 is the canonical advanced problem.**

---

# 71. Graph + String Dependencies

### Problems

- [269] Alien Dictionary
- [2115] Find All Possible Recipes from Given Supplies
- [1203] Sort Items by Groups Respecting Dependencies
- [2050] Parallel Courses III
- [2192] All Ancestors of a Node in a Directed Acyclic Graph
- [1857] Largest Color Value in a Directed Graph

---

# 72. Recipe / Dependency Graph

## Observation

```text
ingredient → recipe
```

When an ingredient becomes available:

```text
decrease missing count
```

This is essentially topological sorting.

### Problems

- [2115] Find All Possible Recipes from Given Supplies
- [207] Course Schedule
- [210] Course Schedule II
- [2050] Parallel Courses III
- [1136] Parallel Courses
- [1203] Sort Items by Groups Respecting Dependencies

---

# 73. Course Schedule Pattern

One of the most important graph interview families.

### Variants

```text
Can finish?
Return ordering?
Prerequisite queries?
Parallel scheduling?
Minimum completion time?
```

### Problems

- [207] Course Schedule
- [210] Course Schedule II
- [1462] Course Schedule IV
- [2050] Parallel Courses III
- [1136] Parallel Courses
- [1494] Parallel Courses II
- [630] Course Schedule III
- [1203] Sort Items by Groups Respecting Dependencies

---

# 74. Graph Scheduling

DAG + time/cost.

### Problems

- [2050] Parallel Courses III
- [1136] Parallel Courses
- [630] Course Schedule III
- [1494] Parallel Courses II
- [1203] Sort Items by Groups Respecting Dependencies
- [1857] Largest Color Value in a Directed Graph

---

# 75. Longest Path in DAG

Unlike general graphs, longest path in DAG is tractable.

### Pattern

```text
topological order
→ dp[node]
→ dp[next] = max(dp[next], dp[node] + ...)
```

### Problems

- [2050] Parallel Courses III
- [1857] Largest Color Value in a Directed Graph
- [329] Longest Increasing Path in a Matrix
- [797] All Paths From Source to Target
- [1136] Parallel Courses
- [1377] Frog Position After T Seconds

---

# 76. Count Number of Shortest Paths

Don't only store:

```text
dist[node]
```

also store:

```text
ways[node]
```

Transition:

```text
shorter:
    dist = newDist
    ways = ways[u]

equal:
    ways += ways[u]
```

### Problems

- [1976] Number of Ways to Arrive at Destination
- [1514] Path with Maximum Probability
- [1786] Number of Restricted Paths From First to Last Node
- [3123] Find Edges in Shortest Paths
- [743] Network Delay Time
- [2045] Second Minimum Time to Reach Destination

---

# 77. Restricted Shortest Paths

Calculate shortest distances first, then only follow edges satisfying a restriction.

### Problems

- [1786] Number of Restricted Paths From First to Last Node
- [1976] Number of Ways to Arrive at Destination
- [3123] Find Edges in Shortest Paths
- [2045] Second Minimum Time to Reach Destination
- [2203] Minimum Weighted Subgraph With the Required Paths

---

# 78. Second Shortest Path

Maintain:

```text
best
second_best
```

instead of only shortest.

### Problems

- [2045] Second Minimum Time to Reach Destination
- [743] Network Delay Time
- [1976] Number of Ways to Arrive at Destination
- [3123] Find Edges in Shortest Paths
- [1786] Number of Restricted Paths From First to Last Node

---

# 79. Shortest Path Tree / Edges in Shortest Paths

First compute:

```text
distFromSource
distFromTarget
```

Then an edge:

```text
u → v
```

is on a shortest path if:

```text
distS[u] + w + distT[v] == distS[target]
```

### Problems

- [3123] Find Edges in Shortest Paths
- [1976] Number of Ways to Arrive at Destination
- [1786] Number of Restricted Paths From First to Last Node
- [743] Network Delay Time
- [2203] Minimum Weighted Subgraph With the Required Paths

---

# 80. Multi-Dijkstra

Run Dijkstra from:

```text
source
target
```

or multiple important nodes.

### Problems

- [2203] Minimum Weighted Subgraph With the Required Paths
- [3123] Find Edges in Shortest Paths
- [1786] Number of Restricted Paths From First to Last Node
- [1976] Number of Ways to Arrive at Destination
- [1334] Find the City With the Smallest Number of Neighbors at a Threshold Distance

---

# 81. Graph State = Position + Remaining Resource

This is one of the highest-value observations.

Ask:

> "If I reach the same node with different remaining resources, are those states actually different?"

If yes:

```text
state = (node, resource)
```

### Problems

- [787] Cheapest Flights Within K Stops
- [1928] Minimum Cost to Reach Destination in Time
- [2093] Minimum Cost to Reach City With Discounts
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [864] Shortest Path to Get All Keys
- [871] Minimum Number of Refueling Stops
- [2577] Minimum Time to Visit a Cell In a Grid

---

# 82. Graph State = Position + Direction

### Problems

- [1210] Minimum Moves to Reach Target with Rotations
- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [499] The Maze III
- [505] The Maze II
- [2577] Minimum Time to Visit a Cell In a Grid
- [1263] Minimum Moves to Move a Box to Their Target Location

---

# 83. Graph State = Position + Mask

### Problems

- [847] Shortest Path Visiting All Nodes
- [864] Shortest Path to Get All Keys
- [980] Unique Paths III
- [943] Find the Shortest Superstring
- [1494] Parallel Courses II
- [1125] Smallest Sufficient Team

---

# 84. Graph State = Position + Time

### Problems

- [1928] Minimum Cost to Reach Destination in Time
- [2577] Minimum Time to Visit a Cell In a Grid
- [2045] Second Minimum Time to Reach Destination
- [1091] Shortest Path in Binary Matrix
- [743] Network Delay Time

---

# 85. Graph + Time-Dependent Edges

Sometimes an edge isn't always available.

Examples:

```text
traffic lights
waiting time
periodic availability
```

### Problems

- [2045] Second Minimum Time to Reach Destination
- [2577] Minimum Time to Visit a Cell In a Grid
- [1928] Minimum Cost to Reach Destination in Time
- [2039] The Time When the Network Becomes Idle
- [1514] Path with Maximum Probability

---

# 86. Graph + Waiting

The optimal move may be:

```text
move
OR
wait
```

### Problems

- [2045] Second Minimum Time to Reach Destination
- [2577] Minimum Time to Visit a Cell In a Grid
- [1928] Minimum Cost to Reach Destination in Time
- [2039] The Time When the Network Becomes Idle

---

# 87. Graph + Obstacles

### Problems

- [1091] Shortest Path in Binary Matrix
- [1293] Shortest Path in a Grid with Obstacles Elimination
- [2290] Minimum Obstacle Removal to Reach the Corner
- [2812] Find the Safest Path in a Grid
- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1263] Minimum Moves to Move a Box to Their Target Location

---

# 88. Graph + Dynamic Activation

Nodes/edges become available over time.

### Problems

- [305] Number of Islands II
- [803] Bricks Falling When Hit
- [1101] The Earliest Moment When Everyone Become Friends
- [1970] Last Day Where You Can Still Cross
- [778] Swim in Rising Water
- [959] Regions Cut By Slashes

---

# 89. Reverse Time + DSU

A very useful trick.

Instead of:

```text
remove nodes forward
```

do:

```text
add nodes backward
```

because DSU handles additions much better than deletions.

### Problems

- [803] Bricks Falling When Hit
- [1970] Last Day Where You Can Still Cross
- [305] Number of Islands II
- [1101] The Earliest Moment When Everyone Become Friends

---

# 90. Offline Queries + DSU

Process queries in sorted order / reverse order.

### Problems

- [1697] Checking Existence of Edge Length Limited Paths
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [2076] Process Restricted Friend Requests
- [1847] Closest Room
- [2421] Number of Good Paths
- [1722] Minimize Hamming Distance After Swap Operations

---

# 91. MST + Edge Classification

Important advanced MST observation:

For every edge:

```text
critical?
pseudo-critical?
neither?
```

### Problems

- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1584] Min Cost to Connect All Points
- [1135] Connecting Cities With Minimum Cost
- [1168] Optimize Water Distribution in a Village
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable

---

# 92. Critical Edge Observation

An edge is critical if:

```text
removing it increases MST cost
or makes MST impossible
```

### Main problem

- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree

Related:

- [1192] Critical Connections in a Network
- [684] Redundant Connection
- [685] Redundant Connection II

---

# 93. Graph + Lowest Common Ancestor

Tree is a graph.

LCA becomes useful when queries ask about:

```text
ancestor
distance
path
lowest common node
```

### Problems

- [1483] Kth Ancestor of a Tree Node
- [1766] Tree of Coprimes
- [2846] Minimum Edge Weight Equilibrium Queries in a Tree
- [2322] Minimum Score After Removals on a Tree
- [2277] Closest Node to Path in Tree
- [2096] Step-By-Step Directions From a Binary Tree Node to Another
- [2646] Minimize the Total Price of the Trips

---

# 94. Binary Lifting

### Problems

- [1483] Kth Ancestor of a Tree Node
- [2846] Minimum Edge Weight Equilibrium Queries in a Tree
- [1766] Tree of Coprimes
- [2646] Minimize the Total Price of the Trips
- [2277] Closest Node to Path in Tree

---

# 95. Tree Path Queries

### Problems

- [2846] Minimum Edge Weight Equilibrium Queries in a Tree
- [2646] Minimize the Total Price of the Trips
- [2322] Minimum Score After Removals on a Tree
- [2277] Closest Node to Path in Tree
- [2096] Step-By-Step Directions From a Binary Tree Node to Another

---

# 96. Graph + Prefix / Subtree Aggregation

On a tree:

```text
DFS
→ aggregate child information
→ calculate parent
```

### Problems

- [834] Sum of Distances in Tree
- [1519] Number of Nodes in the Sub-Tree With the Same Label
- [2049] Count Nodes With the Highest Score
- [2458] Height of Binary Tree After Subtree Removal Queries
- [2925] Maximum Score After Applying Operations on a Tree
- [2646] Minimize the Total Price of the Trips

---

# 97. Graph + Backtracking

When the graph is small and paths/states must be enumerated.

### Problems

- [797] All Paths From Source to Target
- [980] Unique Paths III
- [79] Word Search
- [212] Word Search II
- [51] N-Queens
- [37] Sudoku Solver
- [332] Reconstruct Itinerary
- [491] Non-decreasing Subsequences

---

# 98. Hamiltonian Path / Visit Every Node

### Observation

If:

```text
N <= ~15
```

and need to remember which vertices were visited:

```text
bitmask
```

### Problems

- [847] Shortest Path Visiting All Nodes
- [980] Unique Paths III
- [943] Find the Shortest Superstring
- [1494] Parallel Courses II
- [1125] Smallest Sufficient Team
- [1723] Find Minimum Time to Finish All Jobs

---

# 99. Graph + Backtracking + Memoization

State:

```text
(node, mask)
```

or:

```text
(position, remaining)
```

Memoize.

### Problems

- [847] Shortest Path Visiting All Nodes
- [864] Shortest Path to Get All Keys
- [980] Unique Paths III
- [943] Find the Shortest Superstring
- [1494] Parallel Courses II
- [1125] Smallest Sufficient Team

---

# 100. "Can Reach?" Graph Observation

Before thinking about shortest path, ask:

> Do I actually care about the distance?

If not:

```text
DFS/BFS
```

is enough.

### Problems

- [1971] Find if Path Exists in Graph
- [841] Keys and Rooms
- [547] Number of Provinces
- [261] Graph Valid Tree
- [1557] Minimum Number of Vertices to Reach All Nodes
- [802] Find Eventual Safe States
- [1059] All Paths from Source Lead to Destination

---

# 101. "Minimum Number of Moves?" Observation

Ask whether every move has equal cost.

If yes:

```text
BFS
```

If costs are:

```text
0/1 → 0-1 BFS
positive → Dijkstra
```

### Problems

- [127] Word Ladder
- [752] Open the Lock
- [1091] Shortest Path in Binary Matrix
- [773] Sliding Puzzle
- [1926] Nearest Exit from Entrance in Maze
- [994] Rotting Oranges
- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner
- [743] Network Delay Time

---

# 102. "Minimum Cost?" Observation

Do NOT automatically use Dijkstra.

First ask:

```text
Are weights:
0/1?
positive?
negative?
DAG?
```

Then choose:

```text
0/1 BFS
Dijkstra
Bellman-Ford
DAG DP
```

### Problems

- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner
- [743] Network Delay Time
- [787] Cheapest Flights Within K Stops
- [1928] Minimum Cost to Reach Destination in Time
- [2662] Minimum Cost of a Path With Special Roads
- [2093] Minimum Cost to Reach City With Discounts
- [2203] Minimum Weighted Subgraph With the Required Paths

---

# 103. "Maximum Possible Minimum" Observation

Usually:

```text
maximize the minimum value along path
```

Think:

```text
binary search
+
BFS/DFS
```

or:

```text
Dijkstra variant
```

### Problems

- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [778] Swim in Rising Water
- [1631] Path With Minimum Effort

---

# 104. "Minimum Possible Maximum" Observation

Usually:

```text
minimize maximum edge/cost
```

### Problems

- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner

---

# 105. "All Paths Eventually..." Observation

If wording says:

```text
every path
eventually reaches
cannot get stuck
safe
```

think:

```text
cycle detection
reverse graph
topological elimination
```

### Problems

- [802] Find Eventual Safe States
- [1059] All Paths from Source Lead to Destination
- [207] Course Schedule
- [210] Course Schedule II
- [2360] Longest Cycle in a Graph
- [2127] Maximum Employees to Be Invited to a Meeting

---

# 106. "Prerequisite" Observation

Immediately think:

```text
directed graph
```

Then:

```text
cycle?
topological ordering?
DAG DP?
```

### Problems

- [207] Course Schedule
- [210] Course Schedule II
- [1462] Course Schedule IV
- [1136] Parallel Courses
- [2050] Parallel Courses III
- [1494] Parallel Courses II
- [1203] Sort Items by Groups Respecting Dependencies

---

# 107. "Connect Everything with Minimum Cost"

Think:

```text
MST
```

NOT:

```text
Dijkstra
```

### Problems

- [1584] Min Cost to Connect All Points
- [1135] Connecting Cities With Minimum Cost
- [1168] Optimize Water Distribution in a Village
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable

---

# 108. "Merge Groups"

Think:

```text
DSU
```

### Problems

- [721] Accounts Merge
- [737] Sentence Similarity II
- [839] Similar String Groups
- [947] Most Stones Removed with Same Row or Column
- [990] Satisfiability of Equality Equations
- [1061] Lexicographically Smallest Equivalent String
- [1202] Smallest String With Swaps
- [1722] Minimize Hamming Distance After Swap Operations
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph

---

# 109. "Remove Maximum Edges While Keeping Connectivity"

Think:

```text
DSU
+
process shared edges first
```

### Problems

- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [684] Redundant Connection
- [685] Redundant Connection II
- [1319] Number of Operations to Make Network Connected
- [1584] Min Cost to Connect All Points

---

# 110. "Number of Ways to Reach"

This is not necessarily just BFS.

Potential state:

```text
shortest distance
+
number of ways
```

### Problems

- [1976] Number of Ways to Arrive at Destination
- [1786] Number of Restricted Paths From First to Last Node
- [70] Climbing Stairs
- [63] Unique Paths II
- [980] Unique Paths III
- [1575] Count All Possible Routes
- [1514] Path with Maximum Probability

---

# 111. Graph + Counting

### Problems

- [1976] Number of Ways to Arrive at Destination
- [1786] Number of Restricted Paths From First to Last Node
- [2421] Number of Good Paths
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph
- [2876] Count Visited Nodes in a Directed Graph
- [2685] Count the Number of Complete Components
- [2285] Maximum Total Importance of Roads
- [1575] Count All Possible Routes

---

# 112. Complete Components

A component is complete when:

```text
edges = n(n-1)/2
```

### Problems

- [2685] Count the Number of Complete Components
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph
- [547] Number of Provinces
- [323] Number of Connected Components in an Undirected Graph
- [2421] Number of Good Paths

---

# 113. Component Size + Combinatorics

After finding component size:

```text
size
```

sometimes answer is:

```text
C(size, 2)
```

### Problems

- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph
- [947] Most Stones Removed with Same Row or Column
- [2421] Number of Good Paths
- [2685] Count the Number of Complete Components
- [547] Number of Provinces

---

# 114. Graph Degree Tricks

Sometimes you don't need traversal immediately.

Look at:

```text
degree[i]
```

### Problems

- [997] Find the Town Judge
- [1791] Find Center of Star Graph
- [1615] Maximal Network Rank
- [2285] Maximum Total Importance of Roads
- [2493] Divide Nodes Into the Maximum Number of Groups
- [2359] Find Closest Node to Given Two Nodes
- [332] Reconstruct Itinerary

---

# 115. Star Graph Observation

If one node connects to everyone:

```text
center
```

### Problems

- [1791] Find Center of Star Graph
- [997] Find the Town Judge
- [1615] Maximal Network Rank
- [2285] Maximum Total Importance of Roads
- [997] Find the Town Judge

---

# 116. Network Rank / Degree-Based Graph Problems

### Problems

- [1615] Maximal Network Rank
- [2285] Maximum Total Importance of Roads
- [997] Find the Town Judge
- [1791] Find Center of Star Graph
- [1557] Minimum Number of Vertices to Reach All Nodes

---

# 117. Graph + Sorting Edges

Sorting edges is useful for:

```text
Kruskal
offline DSU
threshold queries
edge activation
MST classification
```

### Problems

- [1584] Min Cost to Connect All Points
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1697] Checking Existence of Edge Length Limited Paths
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [1101] The Earliest Moment When Everyone Become Friends
- [2076] Process Restricted Friend Requests
- [2421] Number of Good Paths

---

# 118. Graph + Offline Queries

### Problems

- [1697] Checking Existence of Edge Length Limited Paths
- [2076] Process Restricted Friend Requests
- [1847] Closest Room
- [2421] Number of Good Paths
- [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
- [1722] Minimize Hamming Distance After Swap Operations

---

# 119. Graph + Monotonic Property

If feasibility has a property like:

```text
if possible(X)
then possible(X+1)
```

you can binary search.

### Problems

- [778] Swim in Rising Water
- [1631] Path With Minimum Effort
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [1970] Last Day Where You Can Still Cross

---

# 120. Reverse Process Trick

Sometimes forward operation is hard:

```text
delete
remove
block
disconnect
```

but reverse is easy:

```text
add
activate
connect
```

### Problems

- [803] Bricks Falling When Hit
- [1970] Last Day Where You Can Still Cross
- [305] Number of Islands II
- [1101] The Earliest Moment When Everyone Become Friends

---

# 121. Graph + Simulation

Some problems require simulating graph state.

### Problems

- [994] Rotting Oranges
- [2039] The Time When the Network Becomes Idle
- [1380] Lucky Numbers in a Matrix
- [2127] Maximum Employees to Be Invited to a Meeting
- [909] Snakes and Ladders
- [289] Game of Life

---

# 122. Network Propagation

Think:

```text
source
→ spread
→ delay
→ all nodes
```

### Problems

- [743] Network Delay Time
- [2039] The Time When the Network Becomes Idle
- [1376] Time Needed to Inform All Employees
- [994] Rotting Oranges
- [1514] Path with Maximum Probability
- [1976] Number of Ways to Arrive at Destination

---

# 123. Graph + Infection / Spread

### Problems

- [994] Rotting Oranges
- [2039] The Time When the Network Becomes Idle
- [1376] Time Needed to Inform All Employees
- [2385] Amount of Time for Binary Tree to Be Infected
- [863] All Nodes Distance K in Binary Tree
- [1162] As Far from Land as Possible

---

# 124. Graph + Probability

### Problems

- [1514] Path with Maximum Probability
- [1377] Frog Position After T Seconds
- [688] Knight Probability in Chessboard
- [1466] Reorder Routes to Make All Paths Lead to the City Zero
- [1976] Number of Ways to Arrive at Destination

---

# 125. Graph + DFS Probability

### Problems

- [1377] Frog Position After T Seconds
- [1514] Path with Maximum Probability
- [399] Evaluate Division
- [688] Knight Probability in Chessboard
- [1466] Reorder Routes to Make All Paths Lead to the City Zero

---

# 126. Reorder Edges / Directed Tree

A particularly useful tree observation:

Given an undirected-looking tree with directed edges, count how many edges need reversal so every node can reach root.

### Problems

- [1466] Reorder Routes to Make All Paths Lead to the City Zero
- [2858] Minimum Edge Reversals So Every Node Is Reachable
- [1519] Number of Nodes in the Sub-Tree With the Same Label
- [834] Sum of Distances in Tree
- [310] Minimum Height Trees

---

# 127. Edge Reversal Cost

Represent:

```text
u → v : cost 0
v → u : cost 1
```

Then solve shortest path.

### Problems

- [1466] Reorder Routes to Make All Paths Lead to the City Zero
- [2858] Minimum Edge Reversals So Every Node Is Reachable
- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [2290] Minimum Obstacle Removal to Reach the Corner

---

# 128. Graph + 0/1 Cost Transformation

A very useful trick:

If an operation has:

```text
cost 0
```

and another:

```text
cost 1
```

model them as edges and use:

```text
0-1 BFS
```

### Problems

- [1368] Minimum Cost to Make at Least One Valid Path in a Grid
- [1466] Reorder Routes to Make All Paths Lead to the City Zero
- [2858] Minimum Edge Reversals So Every Node Is Reachable
- [2290] Minimum Obstacle Removal to Reach the Corner

---

# 129. Graph + Minimum Cut Intuition

Even when you don't implement max-flow, learn to recognize:

```text
minimum set of edges/nodes whose removal disconnects source and target
```

### Problems

- [1192] Critical Connections in a Network
- [1568] Minimum Number of Days to Disconnect Island
- [1489] Critical and Pseudo-Critical Edges in MST
- [1320] Minimum Distance to Type a Word Using Two Fingers
- [1820] Maximum Number of Accepted Invitations

---

# 130. "One Edge Can Change the Answer"

This pattern appears in:

```text
remove one edge
add one edge
change one edge
reverse one edge
```

### Problems

- [684] Redundant Connection
- [685] Redundant Connection II
- [1489] Critical and Pseudo-Critical Edges in MST
- [2699] Modify Graph Edge Weights
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [1466] Reorder Routes to Make All Paths Lead to the City Zero

---

# 131. Add One Edge / Make Graph Connected

### Problems

- [1319] Number of Operations to Make Network Connected
- [1579] Remove Max Number of Edges to Keep Graph Fully Traversable
- [684] Redundant Connection
- [685] Redundant Connection II
- [261] Graph Valid Tree

---

# 132. Graph Valid Tree

Classic combination:

```text
connected
+
N - 1 edges
+
no cycle
```

### Problems

- [261] Graph Valid Tree
- [684] Redundant Connection
- [1319] Number of Operations to Make Network Connected
- [323] Number of Connected Components in an Undirected Graph
- [1584] Min Cost to Connect All Points

---

# 133. Connected + N-1 Edges Observation

For an undirected graph:

```text
connected + n-1 edges
⇔ tree
```

### Problems

- [261] Graph Valid Tree
- [684] Redundant Connection
- [1319] Number of Operations to Make Network Connected
- [1135] Connecting Cities With Minimum Cost
- [1584] Min Cost to Connect All Points

---

# 134. Graph Complement Thinking

Sometimes the stated graph is the complement of the graph you actually need to traverse.

### Problems

- [261] Graph Valid Tree
- [886] Possible Bipartition
- [2493] Divide Nodes Into the Maximum Number of Groups
- [928] Minimize Malware Spread
- [924] Minimize Malware Spread
- [1559] Detect Cycles in 2D Grid

---

# 135. Malware / Infection Components

The key observation is often:

```text
infected component
+
number of initially infected nodes
```

### Problems

- [924] Minimize Malware Spread
- [928] Minimize Malware Spread II
- [947] Most Stones Removed with Same Row or Column
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph
- [2421] Number of Good Paths

---

# 136. Graph + Component Contribution

Find components first, then calculate their contribution.

### Problems

- [924] Minimize Malware Spread
- [928] Minimize Malware Spread II
- [2316] Count Unreachable Pairs of Nodes in an Undirected Graph
- [947] Most Stones Removed with Same Row or Column
- [2421] Number of Good Paths
- [2685] Count the Number of Complete Components

---

# 137. Graph + DSU + Value Sorting

Sort nodes/edges by value, activate gradually, union active neighbors.

### Problems

- [2421] Number of Good Paths
- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1102] Path With Maximum Minimum Value
- [1970] Last Day Where You Can Still Cross

---

# 138. "Activate by Threshold"

Process:

```text
value <= threshold
```

or:

```text
value >= threshold
```

### Problems

- [2421] Number of Good Paths
- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1102] Path With Maximum Minimum Value
- [2812] Find the Safest Path in a Grid
- [1970] Last Day Where You Can Still Cross

---

# 139. Graph + Sorted Values

### Problems

- [2421] Number of Good Paths
- [1631] Path With Minimum Effort
- [778] Swim in Rising Water
- [1102] Path With Maximum Minimum Value
- [1970] Last Day Where You Can Still Cross
- [329] Longest Increasing Path in a Matrix

---

# 140. Increasing Path in Grid

A grid becomes a DAG if movement only goes to:

```text
strictly larger
```

because cycles become impossible.

### Problems

- [329] Longest Increasing Path in a Matrix
- [2713] Maximum Strictly Increasing Cells in a Matrix
- [2328] Number of Increasing Paths in a Grid
- [2684] Maximum Number of Moves in a Grid
- [417] Pacific Atlantic Water Flow

---

# 141. Grid DAG

### Problems

- [329] Longest Increasing Path in a Matrix
- [2328] Number of Increasing Paths in a Grid
- [2684] Maximum Number of Moves in a Grid
- [2713] Maximum Strictly Increasing Cells in a Matrix
- [1976] Number of Ways to Arrive at Destination

---

# 142. Graph + Memoized DFS

If graph is acyclic or states have no problematic cycles:

```text
DFS(state)
memo[state]
```

### Problems

- [329] Longest Increasing Path in a Matrix
- [797] All Paths From Source to Target
- [1377] Frog Position After T Seconds
- [1575] Count All Possible Routes
- [1857] Largest Color Value in a Directed Graph
- [2328] Number of Increasing Paths in a Grid
- [2684] Maximum Number of Moves in a Grid

---

# 143. Number of Increasing Paths

### Problems

- [2328] Number of Increasing Paths in a Grid
- [329] Longest Increasing Path in a Matrix
- [2713] Maximum Strictly Increasing Cells in a Matrix
- [2684] Maximum Number of Moves in a Grid
- [417] Pacific Atlantic Water Flow

---

# 144. Graph + Memoization + DAG

### Problems

- [329] Longest Increasing Path in a Matrix
- [2328] Number of Increasing Paths in a Grid
- [1857] Largest Color Value in a Directed Graph
- [1575] Count All Possible Routes
- [797] All Paths From Source to Target
- [2684] Maximum Number of Moves in a Grid

---

# 145. Graph + Constraints + Topological DP

### Problems

- [1857] Largest Color Value in a Directed Graph
- [2050] Parallel Courses III
- [1136] Parallel Courses
- [1203] Sort Items by Groups Respecting Dependencies
- [2192] All Ancestors of a Node in a Directed Acyclic Graph
- [2115] Find All Possible Recipes from Given Supplies

---

# 146. Color DP on DAG

A node has a color/value and we need the maximum frequency along a path.

### Problems

- [1857] Largest Color Value in a Directed Graph
- [2050] Parallel Courses III
- [329] Longest Increasing Path in a Matrix
- [2328] Number of Increasing Paths in a Grid
- [797] All Paths From Source to Target

---

# 147. Graph + Frequency / Aggregation

### Problems

- [1857] Largest Color Value in a Directed Graph
- [1519] Number of Nodes in the Sub-Tree With the Same Label
- [2421] Number of Good Paths
- [2049] Count Nodes With the Highest Score
- [2285] Maximum Total Importance of Roads
- [2685] Count the Number of Complete Components

---

# 148. Special Graph: Complete Graph

When every pair is connected, don't explicitly build:

```text
O(N²)
```

unless N is small.

### Problems

- [1615] Maximal Network Rank
- [2285] Maximum Total Importance of Roads
- [924] Minimize Malware Spread
- [1489] Critical and Pseudo-Critical Edges in MST
- [1584] Min Cost to Connect All Points

---

# 149. Hidden Graph

Sometimes no adjacency list exists.

You must construct:

```text
node
→ possible next states
```

### Problems

- [127] Word Ladder
- [752] Open the Lock
- [773] Sliding Puzzle
- [2059] Minimum Operations to Convert Number
- [864] Shortest Path to Get All Keys
- [847] Shortest Path Visiting All Nodes
- [909] Snakes and Ladders
- [433] Minimum Genetic Mutation

---

# 150. The Most Important Graph Observational Patterns

When reading a problem, train yourself to automatically map phrases to algorithms.

| Problem wording / observation | Likely technique |
|---|---|
| "can reach" | DFS/BFS |
| "number of groups" | Components / DSU |
| "connected" | DFS/BFS/DSU |
| "minimum number of moves" | BFS |
| "shortest unweighted path" | BFS |
| "0/1 cost" | 0-1 BFS |
| "positive edge weights" | Dijkstra |
| "negative edges" | Bellman-Ford |
| "all pairs shortest path" | Floyd-Warshall |
| "prerequisite" | Topological sort |
| "dependency" | Topological sort |
| "ordering" | Topological sort |
| "every path eventually..." | Cycle / safe states |
| "two groups" | Bipartite |
| "odd cycle" | Bipartite |
| "connect all with minimum cost" | MST |
| "merge groups" | DSU |
| "redundant edge" | DSU |
| "critical connection" | Bridges / Tarjan |
| "all edges exactly once" | Eulerian |
| "visit every node" | Bitmask DP/BFS |
| "all keys" | Bitmask BFS |
| "K stops" | State shortest path |
| "limited time/fuel" | State shortest path |
| "maximum minimum" | Binary search / maximin |
| "minimum maximum" | Binary search / minimax |
| "multiple starting points" | Multi-source BFS |
| "nearest source" | Multi-source BFS |
| "grid" | Grid graph |
| "island" | Grid DFS/BFS/DSU |
| "equivalent" | DSU |
| "equations" | Weighted graph / DSU |
| "word transformation" | Implicit BFS |
| "one character change" | Implicit graph |
| "cycle in directed graph" | 3-state DFS / topo |
| "cycle in undirected graph" | parent DFS / DSU |
| "every node has one outgoing edge" | Functional graph |
| "longest cycle" | Functional graph |
| "tree + path queries" | LCA / binary lifting |
| "tree + all roots" | Rerooting DP |
| "distance from every node" | Rerooting / multi-source |
| "edge reversal" | 0-1 BFS / rerooting |
| "remove nodes forward" | Reverse process |
| "queries with thresholds" | Offline sorting + DSU |
| "activate by value" | DSU / threshold |
| "count shortest paths" | Dijkstra + ways |
| "second shortest" | 2-distance Dijkstra |
| "DAG + longest path" | Topological DP |
| "increasing grid" | DAG DP |
| "same component contribution" | Components + combinatorics |

---

# 151. Graph Algorithm Decision Tree

Use this during contests/interviews.

```text
START
 │
 ├── Is it a graph?
 │     │
 │     ├── Explicit edges?
 │     │       │
 │     │       └── YES
 │     │
 │     └── Hidden graph?
 │             └── Define state + transitions
 │
 ├── Need simple reachability?
 │       └── DFS / BFS
 │
 ├── Need connected components?
 │       └── DFS / BFS / DSU
 │
 ├── Need cycle?
 │       │
 │       ├── Undirected → parent DFS / DSU
 │       └── Directed → 3-state DFS / Kahn
 │
 ├── Need shortest path?
 │       │
 │       ├── Unweighted → BFS
 │       ├── 0/1 weights → 0-1 BFS
 │       ├── Positive weights → Dijkstra
 │       ├── Negative edges → Bellman-Ford
 │       ├── DAG → DP / topo
 │       └── All pairs → Floyd-Warshall
 │
 ├── Need ordering?
 │       └── Topological sort
 │
 ├── Need minimum cost to connect ALL?
 │       └── MST
 │
 ├── Need merge/connectivity?
 │       └── DSU
 │
 ├── Need critical edges?
 │       └── Tarjan bridges
 │
 ├── Need two groups?
 │       └── Bipartite
 │
 ├── Need every edge exactly once?
 │       └── Eulerian path
 │
 ├── Need visit every node?
 │       └── Bitmask BFS/DP
 │
 ├── Need state + shortest path?
 │       └── BFS/Dijkstra on (node,state)
 │
 ├── Multiple sources?
 │       └── Multi-source BFS/Dijkstra
 │
 ├── Threshold answer?
 │       └── Binary Search + Graph
 │
 ├── Dynamic connectivity?
 │       └── DSU
 │
 ├── Tree path queries?
 │       └── LCA / Binary Lifting
 │
 └── Tree answers for every root?
         └── Rerooting DP
```

---

# 152. Highest Priority Graph Problems

If you don't have time to solve everything, solve these first.

## Tier 1 — Absolute Must Do

### Traversal

1. [200] Number of Islands
2. [695] Max Area of Island
3. [547] Number of Provinces
4. [1971] Find if Path Exists in Graph
5. [841] Keys and Rooms

### BFS

6. [994] Rotting Oranges
7. [127] Word Ladder
8. [752] Open the Lock
9. [1091] Shortest Path in Binary Matrix
10. [542] 01 Matrix

### Topological Sort

11. [207] Course Schedule
12. [210] Course Schedule II
13. [802] Find Eventual Safe States
14. [2115] Find All Possible Recipes from Given Supplies
15. [2050] Parallel Courses III

### DSU

16. [684] Redundant Connection
17. [1319] Number of Operations to Make Network Connected
18. [721] Accounts Merge
19. [1202] Smallest String With Swaps
20. [1584] Min Cost to Connect All Points

### Dijkstra

21. [743] Network Delay Time
22. [1514] Path with Maximum Probability
23. [1631] Path With Minimum Effort
24. [787] Cheapest Flights Within K Stops
25. [778] Swim in Rising Water

### MST

26. [1584] Min Cost to Connect All Points
27. [1489] Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
28. [1135] Connecting Cities With Minimum Cost
29. [1168] Optimize Water Distribution in a Village
30. [1579] Remove Max Number of Edges to Keep Graph Fully Traversable

### Advanced

31. [1192] Critical Connections in a Network
32. [847] Shortest Path Visiting All Nodes
33. [864] Shortest Path to Get All Keys
34. [2360] Longest Cycle in a Graph
35. [834] Sum of Distances in Tree

---

# 153. Tier 2 — Strong Interview Preparation

36. [886] Possible Bipartition
37. [785] Is Graph Bipartite?
38. [310] Minimum Height Trees
39. [399] Evaluate Division
40. [433] Minimum Genetic Mutation
41. [815] Bus Routes
42. [773] Sliding Puzzle
43. [1293] Shortest Path in a Grid with Obstacles Elimination
44. [1368] Minimum Cost to Make at Least One Valid Path in a Grid
45. [2290] Minimum Obstacle Removal to Reach the Corner
46. [1928] Minimum Cost to Reach Destination in Time
47. [2203] Minimum Weighted Subgraph With the Required Paths
48. [1976] Number of Ways to Arrive at Destination
49. [1786] Number of Restricted Paths From First to Last Node
50. [3123] Find Edges in Shortest Paths
51. [2421] Number of Good Paths
52. [1697] Checking Existence of Edge Length Limited Paths
53. [2076] Process Restricted Friend Requests
54. [305] Number of Islands II
55. [803] Bricks Falling When Hit

---

# 154. Tier 3 — Advanced / Hard

56. [1483] Kth Ancestor of a Tree Node
57. [2846] Minimum Edge Weight Equilibrium Queries in a Tree
58. [2858] Minimum Edge Reversals So Every Node Is Reachable
59. [3241] Time Taken to Mark All Nodes
60. [2581] Count Number of Possible Root Nodes
61. [1857] Largest Color Value in a Directed Graph
62. [2699] Modify Graph Edge Weights
63. [2812] Find the Safest Path in a Grid
64. [2577] Minimum Time to Visit a Cell In a Grid
65. [2662] Minimum Cost of a Path With Special Roads
66. [2093] Minimum Cost to Reach City With Discounts
67. [2876] Count Visited Nodes in a Directed Graph
68. [2127] Maximum Employees to Be Invited to a Meeting
69. [1203] Sort Items by Groups Respecting Dependencies
70. [1494] Parallel Courses II
71. [1349] Maximum Students Taking Exam
72. [1595] Minimum Cost to Connect Two Groups of Points
73. [1820] Maximum Number of Accepted Invitations
74. [1947] Maximum Compatibility Score Sum
75. [943] Find the Shortest Superstring

---

# 155. The 20 Graph Patterns You MUST Master

If your goal is interviews/placements, prioritize these:

```text
1. DFS/BFS traversal
2. Grid as graph
3. Connected components
4. Cycle detection — undirected
5. Cycle detection — directed
6. Bipartite graph
7. Multi-source BFS
8. BFS shortest path
9. BFS + state
10. Topological sorting
11. DAG DP
12. Dijkstra
13. Dijkstra + state
14. 0-1 BFS
15. DSU
16. Kruskal / MST
17. Bridges / Tarjan
18. Bitmask graph
19. Functional graph
20. Tree rerooting / LCA
```

---

# 156. The 10 Most Important Graph Observations

These are arguably more important than memorizing algorithms.

## Observation 1

```text
Unweighted shortest path
→ BFS
```

## Observation 2

```text
0/1 edge cost
→ 0-1 BFS
```

## Observation 3

```text
Positive weighted edges
→ Dijkstra
```

## Observation 4

```text
Prerequisite / dependency
→ Directed graph + topological sort
```

## Observation 5

```text
Connect all nodes with minimum total cost
→ MST
```

## Observation 6

```text
Merge groups / connectivity under additions
→ DSU
```

## Observation 7

```text
Multiple starting points spreading
→ Multi-source BFS
```

## Observation 8

```text
Same node but different resources
→ State graph
```

Example:

```text
(node, stops)
(node, mask)
(node, fuel)
(node, time)
```

## Observation 9

```text
Every path eventually reaches destination
→ Cycle / topological reasoning
```

## Observation 10

```text
Answer is threshold-dependent
→ Binary Search + Graph
```

---

# 157. Final Graph Master Checklist

Before coding a graph problem, ask these in order:

```text
1. What is a node?

2. What is an edge?

3. Directed or undirected?

4. Weighted or unweighted?

5. Static or dynamic?

6. Do I need:
   - reachability?
   - components?
   - cycle?
   - shortest path?
   - longest path?
   - ordering?
   - connectivity?
   - minimum cost?
   - maximum cost?
   - counting?
   - probability?

7. Is the graph hidden inside:
   - grid?
   - string?
   - state?
   - transformation?
   - tree?
   - dependency system?

8. If shortest path:
   - BFS?
   - 0-1 BFS?
   - Dijkstra?
   - Bellman-Ford?
   - DAG DP?
   - Floyd?

9. If connectivity:
   - DFS/BFS?
   - DSU?
   - reverse DSU?

10. If optimization:
    - MST?
    - Dijkstra?
    - binary search + graph?

11. Does state matter?

    dist[node]
    vs
    dist[node][state]

12. Can I reverse the graph?

13. Can I process the graph offline?

14. Can I sort edges/nodes by value?

15. Is this secretly a DAG?

16. Is this secretly a functional graph?

17. Is this secretly a tree?

18. Is there a cycle?
    If yes, does the cycle itself matter?

19. Can I reduce the problem to components?

20. Can I solve the feasibility version first?
```

---

# Recommended Graph Learning Order

Do NOT learn graph algorithms randomly.

Follow this sequence:

```text
PHASE 1 — Fundamentals
    ↓
Graph representation
    ↓
DFS
    ↓
BFS
    ↓
Grid DFS/BFS
    ↓
Connected Components
    ↓
Cycle Detection

PHASE 2 — Structural Graphs
    ↓
Bipartite
    ↓
Topological Sort
    ↓
DAG
    ↓
DAG DP
    ↓
Functional Graph

PHASE 3 — Connectivity
    ↓
DSU
    ↓
Kruskal
    ↓
Prim
    ↓
Offline DSU

PHASE 4 — Shortest Paths
    ↓
BFS shortest path
    ↓
Multi-source BFS
    ↓
0-1 BFS
    ↓
Dijkstra
    ↓
Dijkstra + State
    ↓
Bellman-Ford
    ↓
Floyd-Warshall

PHASE 5 — Advanced Graph
    ↓
Bridges
    ↓
Articulation Points
    ↓
SCC
    ↓
Eulerian Path
    ↓
LCA
    ↓
Binary Lifting
    ↓
Rerooting DP

PHASE 6 — Advanced State Graph
    ↓
Bitmask BFS
    ↓
Bitmask DP
    ↓
Graph + resource
    ↓
Graph + time
    ↓
Graph + binary search
    ↓
Graph + offline queries
```

# Core 50 Graph Problem Set

If you want a compact set that covers almost every important pattern:

```text
1.   200  Number of Islands
2.   695  Max Area of Island
3.   547  Number of Provinces
4.   1971 Find if Path Exists in Graph
5.   841  Keys and Rooms
6.   994  Rotting Oranges
7.   542  01 Matrix
8.   127  Word Ladder
9.   752  Open the Lock
10.  1091 Shortest Path in Binary Matrix

11.  785  Is Graph Bipartite?
12.  886  Possible Bipartition
13.  207  Course Schedule
14.  210  Course Schedule II
15.  802  Find Eventual Safe States

16.  684  Redundant Connection
17.  721  Accounts Merge
18.  1202 Smallest String With Swaps
19.  1319 Number of Operations to Make Network Connected
20.  1584 Min Cost to Connect All Points

21.  743  Network Delay Time
22.  1514 Path with Maximum Probability
23.  1631 Path With Minimum Effort
24.  787  Cheapest Flights Within K Stops
25.  778  Swim in Rising Water

26.  1368 Minimum Cost to Make at Least One Valid Path
27.  2290 Minimum Obstacle Removal
28.  1928 Minimum Cost to Reach Destination in Time
29.  2203 Minimum Weighted Subgraph
30.  1976 Number of Ways to Arrive at Destination

31.  1489 Critical/Pseudo-Critical MST Edges
32.  1579 Remove Max Edges
33.  1697 Checking Edge Length Limited Paths
34.  2421 Number of Good Paths
35.  305  Number of Islands II

36.  1192 Critical Connections
37.  310  Minimum Height Trees
38.  834  Sum of Distances in Tree
39.  1483 Kth Ancestor
40.  2858 Minimum Edge Reversals

41.  847  Shortest Path Visiting All Nodes
42.  864  Shortest Path to Get All Keys
43.  2360 Longest Cycle
44.  2876 Count Visited Nodes
45.  2127 Maximum Employees to Be Invited

46.  1857 Largest Color Value
47.  2115 Find All Possible Recipes
48.  2699 Modify Graph Edge Weights
49.  2577 Minimum Time to Visit a Cell
50.  2812 Find the Safest Path in a Grid
```

