# Codeforces Graph Pattern Roadmap

---

# 1. Graph Traversal — DFS / BFS / Connected Components

### Pattern
- Basic DFS/BFS
- Connected components
- Reachability
- Component size
- Multi-source traversal
- Grid as graph
- Implicit graph

### Problems

1. **115A — Party**
2. **500A — New Year Transportation**
3. **580C — Kefa and Park**
4. **893C — Rumor**
5. **217A — Ice Skating**
6. **1829E — The Lakes**
7. **377A — Maze**
8. **1167C — News Distribution**
9. **977E — Cyclic Components**
10. **755C — PolandBall and Forest**
11. **1029C — Maximal Intersection**
12. **1294F — Three Paths on a Tree**
13. **1106D — Lunar New Year and a Wander**
14. **1360E — Polygon**
15. **1196F — K-th Path**
16. **1472C — Long Jumps**
17. **1833E — Round Dance**
18. **1485D — Multiples and Power Differences**
19. **987F — AND Graph**
20. **920E — Connected Components?**

The beginner graph set on Codeforces strongly features problems such as Party, New Year Transportation, Kefa and Park, Ice Skating, Cyclic Components and News Distribution. [Codeforces](https://codeforces.com/problemset?locale=en\&order=BY_SOLVED_DESC\&tags=graphs\&utm_source=chatgpt.com)

---

# 2. DFS Tree / Parent / Depth Observation

### Observation

> If the problem asks about ancestry, descendants, depth, subtree, or “go from parent to child”, build a DFS tree.

### Problems

1. **115A — Party**
2. **580C — Kefa and Park**
3. **161D — Distance in Tree**
4. **1324F — Maximum White Subtree**
5. **1336A — Linova and Kingdom**
6. **274B — Zero Tree**
7. **1153E — Serval and Snake**
8. **1118F1 — Tree Cutting (Easy Version)**
9. **1118F2 — Tree Cutting (Hard Version)**
10. **1092F — Tree with Maximum Cost**
11. **1187E — Tree Painting**
12. **763A — Timofey and a Tree**
13. **1029F — Multicolored Markers**
14. **1029E — Tree with Small Distances**
15. **570D — Tree Requests**

---

# 3. Cycle Detection — Undirected Graph

### Pattern

For an undirected graph:

- DFS + parent
- DSU
- Degree observation
- Component edge count

### Problems

1. **977E — Cyclic Components**
2. **1209D — Cow and Snacks**
3. **755C — PolandBall and Forest**
4. **1029D — Concatenated Numbers**
5. **468B — Two Sets**
6. **659E — New Reform**
7. **1176E — Cover it!**
8. **550D — Regular Bridge**
9. **1242B — 0-1 MST**
10. **25D — Roads not only in Berland**
11. **1399E1 — Weights Division (easy)**
12. **1399E2 — Weights Division (hard)**

### Important observation

If a connected component has:

```text
number of edges == number of vertices
```

then for a simple connected graph it is often useful to investigate whether the component is **unicyclic**.

Codeforces' **977E — Cyclic Components** is the canonical example. [Codeforces](https://codeforces.com/problemset/problem/977/E?locale=en\&mobile=false\&utm_source=chatgpt.com)

---

# 4. Cycle Detection — Directed Graph

### Pattern

- DFS coloring
- 3-state DFS
- Back edge
- Topological-sort failure

States:

```text
0 = unvisited
1 = currently in recursion stack
2 = completely processed
```

### Problems

1. **510C — Fox And Names**
2. **999E — Reachability from the Capital**
3. **1179A — Valeriy and Deque**
4. **919E — Congruence Equation**
5. **986A — Fair**
6. **1029F — Multicolored Markers**
7. **1598E — Air Conditioners**
8. **1867B — XOR Palindromes**
9. **825E — Minimal Labels**
10. **1385E — Directing Edges**

---

# 5. Bipartite Graph / 2-Coloring

### Pattern

If you see:

- divide vertices into two groups
- enemies/friends
- alternating colors
- parity
- odd/even distance
- constraints of form `x != y`

think:

> **Bipartite graph / 2-coloring**

### Problems

1. **862B — Mahmoud and Ehab and the Bipartiteness**
2. **1176E — Cover it!**
3. **115A — Party**
4. **468B — Two Sets**
5. **687A — NP-Hard Problem**
6. **1000E — We Need More Bosses**
7. **813F — Bipartite Checking**
8. **1105D — Kilani and the Game**
9. **1490D — Permutation Transformation**
10. **1354E — Graph Coloring**
11. **1615D — Xor-Paths**
12. **1498B — Box is Pull**
13. **1971F — Circle Perimeter**
14. **1366D — Two Divisors**
15. **1270D — Strange Device**

### Key observation

```text
Odd cycle exists
        ↓
Not bipartite

No odd cycle
        ↓
2-colorable
```

---

# 6. DSU — Basic Component Merging

### Pattern

Think DSU when operations look like:

```text
merge(u,v)
same component?
number of components?
component size?
connect these groups?
```

### Problems

1. **1167C — News Distribution**
2. **755C — PolandBall and Forest**
3. **25D — Roads not only in Berland**
4. **1209D — Cow and Snacks**
5. **1242B — 0-1 MST**
6. **500B — New Year Permutation**
7. **1139C — Edgy Trees**
8. **1534C — Little Alawn's Puzzle**
9. **891C — Envy**
10. **1108F — MST Unification**
11. **1213G — Path Queries**
12. **609E — Minimum spanning tree for each edge**
13. **1245D — Shichikuji and Power Grid**
14. **1553G — Common Divisor Graph**
15. **1706E — Qpwoeirut And Vertices**

Codeforces has a very large DSU/graph problem pool; for example, its current DSU+graph problemset includes 1167C, 755C, 25D, 1209D, 1242B, 609E and many more. [Codeforces](https://codeforces.com/problemset?locale=en\&tags=dsu%2Cgraphs\&utm_source=chatgpt.com)

---

# 7. DSU — Detect Cycle / Redundant Edge

### Observation

Instead of explicitly detecting a cycle:

```cpp
if (find(u) == find(v))
    // adding this edge creates a cycle
```

### Problems

1. **25D — Roads not only in Berland**
2. **1209D — Cow and Snacks**
3. **755C — PolandBall and Forest**
4. **891C — Envy**
5. **1108F — MST Unification**
6. **609E — Minimum spanning tree for each edge**
7. **1242B — 0-1 MST**
8. **500B — New Year Permutation**
9. **1553G — Common Divisor Graph**
10. **1534C — Little Alawn's Puzzle**

---

# 8. DSU — Offline Connectivity

### Pattern

Queries ask connectivity after:

- adding edges
- deleting edges
- processing edges backward
- thresholding edges

Think:

> **Offline + DSU**

### Problems

1. **1213G — Path Queries**
2. **891C — Envy**
3. **1108F — MST Unification**
4. **609E — Minimum spanning tree for each edge**
5. **1217F — Forced Online Queries Problem**
6. **1167C — News Distribution**
7. **500B — New Year Permutation**
8. **1242B — 0-1 MST**
9. **920E — Connected Components?**
10. **1698C — 3SUM**

---

# 9. Minimum Spanning Tree — Kruskal

### Pattern

If you need:

```text
connect all vertices
minimum total edge cost
n - 1 edges
```

think:

> **MST → Kruskal + DSU**

### Problems

1. **609E — Minimum spanning tree for each edge**
2. **1108F — MST Unification**
3. **1245D — Shichikuji and Power Grid**
4. **1242B — 0-1 MST**
5. **891C — Envy**
6. **733F — Drivers Dissatisfaction**
7. **956D — Logo Turtle**
8. **1508C — Complete the MST**
9. **1633E — Spanning Tree Queries**
10. **2081D — MST in Modulo Graph**
11. **959E — Mahmoud and Ehab and the xor-MST**
12. **1000E — We Need More Bosses**
13. **1019A — Elections**
14. **1498B — Box is Pull**
15. **125E — MST Company**

Codeforces' MST-tagged problemset contains examples including 609E, 1242B, 1245D and several advanced MST variants. [Codeforces](https://codeforces.com/problemset?locale=en\&tags=graphs%2Cmst\&utm_source=chatgpt.com)

---

# 10. MST — Kruskal Reconstruction / Merge Tree

### Observation

Sometimes DSU is not only used to build the MST.

You can create a **new tree representing component merges**.

Useful for:

- minimum edge threshold
- maximum edge on path
- connectivity at threshold `x`
- component size after threshold
- query over merged components

### Problems

1. **609E — Minimum spanning tree for each edge**
2. **1213G — Path Queries**
3. **891C — Envy**
4. **1108F — MST Unification**
5. **1253F — Cheap Robot**
6. **1706E — Qpwoeirut And Vertices**
7. **888G — Xor-MST**
8. **959E — Mahmoud and Ehab and the xor-MST**
9. **1508C — Complete the MST**
10. **1633E — Spanning Tree Queries**

---

# 11. Shortest Path — BFS

### Pattern

Unweighted graph:

```text
edge cost = 1
```

Then:

> **BFS = shortest path**

### Problems

1. **520B — Two Buttons**
2. **1106D — Lunar New Year and a Wander**
3. **1365D — Solve The Maze**
4. **1594D — The Number of Pairs**
5. **1272E — Nearest Opposite Parity**
6. **1353E — K-th Not Divisible by n**
8. **1355B — Young Explorers**
9. **1407D — Discrete Centrifugal Jumps**
10. **1485D — Multiples and Power Differences**
11. **1660C — Get an Even String**
12. **1534E — Lost Array**

---

# 12. Multi-Source BFS

### Observation

If many nodes are **simultaneously sources**:

```text
distance from nearest hospital
nearest monster
nearest special node
spread from multiple locations
```

initialize all sources into the queue.

### Problems

1. **1081D — Maximum Distance**
2. **1105D — Kilani and the Game**
3. **1270D — Strange Device**
4. **1365D — Solve The Maze**
5. **1350E — Orac and Game of Life**
6. **1514D — Cut and Stick**
7. **1520E — Arranging The Sheep**
8. **1485F — Copy or Prefix Sum**
9. **1334E — Divisors and Graph**
10. **986C — AND Graph**

---

# 13. 0-1 BFS

### Pattern

Edge weights are only:

```text
0 or 1
```

Use:

```text
deque
```

instead of Dijkstra.

### Problems

1. **173B — Chamber of Secrets**
2. **590C — Three States**
3. **1272E — Nearest Opposite Parity**
4. **1063B — Labyrinth**
5. **938D — Buy a Ticket**
6. **1063C — Dwarves, Hats and Extrasensory Abilities**
7. **1695D — Tree Queries**
8. **1791F — Range Update Point Query**
9. **1019A — Elections**
10. **1242B — 0-1 MST**

**Important:** Codeforces often hides 0-1 BFS behind transformations. Do not only search for explicit edge weights `0/1`.

---

# 14. Dijkstra — Standard

### Pattern

```text
weighted graph
non-negative weights
single source
```

→ **Dijkstra**

### Problems

1. **20C — Dijkstra?**
2. **545E — Paths and Trees**
3. **449B — Jzzhu and Cities**
4. **938D — Buy a Ticket**
5. **1486D — Max Median**
6. **1473E — Minimum Path**
7. **1500D — Build The Bridge**
8. **1547G — How Many Paths?**
9. **1263E — Editor**
10. **1851G — Vlad and the Mountains**
11. **1915G — Bicycles**
12. **1472G — Moving to the Capital**
13. **1076D — Edge Deletion**
14. **1266D — Decreasing Debts**
15. **1547E — Air Conditioners**

Codeforces' shortest-path set includes canonical problems such as **20C — Dijkstra?**, **449B — Jzzhu and Cities**, **938D — Buy a Ticket**, and others. [Codeforces](https://codeforces.com/problemset/page/1?order=BY_SOLVED_DESC\&tags=graphs%2C1900-2300\&utm_source=chatgpt.com)

---

# 15. Dijkstra — State Graph

### Observation

Your actual state isn't just:

```text
node
```

It is:

```text
(node, state)
```

Examples:

```text
(node, coupons_used)
(node, parity)
(node, number_of_special_edges)
(node, fuel)
(node, direction)
```

### Problems

1. **1473E — Minimum Path**
2. **938D — Buy a Ticket**
3. **1851G — Vlad and the Mountains**
4. **1915G — Bicycles**
5. **1547E — Air Conditioners**
6. **1486E — Paired Payment**
7. **1486D — Max Median**
8. **1263F — Economic Difficulties**
9. **1076D — Edge Deletion**
10. **1196F — K-th Path**
11. **1486C2 — Guessing the Greatest**
12. **786B — Weird Journey**

---

# 16. Dijkstra — Reverse Graph Observation

### Observation

Question asks:

```text
distance TO target
```

Instead of running from every node:

```text
reverse all edges
run Dijkstra(target)
```

### Problems

1. **20C — Dijkstra?**
2. **1486E — Paired Payment**
3. **1076D — Edge Deletion**
4. **545E — Paths and Trees**
5. **449B — Jzzhu and Cities**
6. **786B — Weird Journey**
7. **1473E — Minimum Path**
8. **1547G — How Many Paths?**
9. **1915G — Bicycles**
10. **1851G — Vlad and the Mountains**

---

# 17. Shortest Path — DAG

### Pattern

If graph is a DAG:

```text
topological sort
+
DP
```

No need for Dijkstra if all edges respect DAG structure.

### Problems

1. **505E — Mr. Kitayuta vs Bamboos**
2. **919D — Substring**
3. **960F — Pathwalks**
4. **1472G — Moving to the Capital**
5. **1547G — How Many Paths?**
6. **721C — Journey**
7. **977F — Consecutive Subsequence**
8. **264C — Choosing Balls**
9. **1000E — We Need More Bosses**
10. **825E — Minimal Labels**

Codeforces graph guides explicitly classify DAG problems around topological sorting + graph DP. [Codeforces](https://codeforces.com/blog/entry/92820?locale=en\&utm_source=chatgpt.com)

---

# 18. Topological Sort

### Pattern

Look for:

- prerequisites
- dependencies
- ordering
- task scheduling
- directed acyclic dependency graph

### Problems

1. **510C — Fox And Names**
2. **1385E — Directing Edges**
3. **825E — Minimal Labels**
4. **919D — Substring**
5. **721C — Journey**
6. **999E — Reachability from the Capital**
7. **1093D — Beautiful Graph**
8. **1572A — Book**
9. **1867B — XOR Palindromes**
10. **104891E — Inverse Topological Sort**
11. **1679D — Toss a Coin to Your Graph**
12. **1547G — How Many Paths?**

---

# 19. DAG DP — Longest / Number of Paths

### Pattern

Once topological order exists:

```text
dp[v] = best answer reaching v
```

or

```text
dp[v] += dp[u]
```

### Problems

1. **919D — Substring**
2. **721C — Journey**
3. **960F — Pathwalks**
4. **1547G — How Many Paths?**
5. **1472G — Moving to the Capital**
6. **1679D — Toss a Coin to Your Graph**
7. **1000E — We Need More Bosses**
8. **1108E1 — Array and Segments**
9. **505B — Mr. Kitayuta's Colorful Graph**
10. **1385E — Directing Edges**

---

# 20. SCC — Strongly Connected Components

### Pattern

Directed graph + mutual reachability:

```text
u can reach v
v can reach u
```

Think:

> **SCC**

Algorithms:

- Kosaraju
- Tarjan

### Problems

1. **999E — Reachability from the Capital**
2. **427C — Checkposts**
3. **999F — Please, another Queries on Array?**
4. **1497E2 — Square-free division**
5. **1029F — Multicolored Markers**
6. **1187E — Tree Painting**
7. **1472G — Moving to the Capital**
8. **1547G — How Many Paths?**
9. **1139E — Maximize Mex**
10. **1494C — 1D Sokoban**
11. **1320B — Journey Planning**
12. **1239E — Turtle**

### Canonical SCC problems to prioritize

```text
427C
999E
1547G
1472G
```

Codeforces' graph topic guides explicitly list SCC and applications such as reachability and condensation graphs. [Codeforces](https://codeforces.com/blog/entry/92820?locale=en\&utm_source=chatgpt.com)

---

# 21. SCC Condensation Graph

### Observation

After finding SCCs:

```text
original graph
       ↓
compress every SCC
       ↓
DAG
```

This is one of the **most important graph observations** on Codeforces.

### Problems

1. **427C — Checkposts**
2. **999E — Reachability from the Capital**
3. **1547G — How Many Paths?**
4. **1472G — Moving to the Capital**
5. **475B — Strongly Connected City**
6. **1242C — Sum Balance**
7. **1172B — Nauuo and Circle**
8. **115E — Linear Kingdom**
9. **1029F — Multicolored Markers**
10. **141E — Clearing Up**

---

# 22. Bridges — Critical Edges

### Observation

If removing an edge disconnects the graph:

> **Bridge**

Typical technique:

```text
tin[u]
low[u]
```

### Problems

1. **1000E — We Need More Bosses**
2. **118E — Bertown Roads**
3. **652E — Pursuit For Artifacts**
4. **610D — Vika and Segments**
5. **732F — Tourist Reform**
6. **1000E — We Need More Bosses**
7. **1276B — Two Fairs**
8. **796D — Police Stations**
9. **827F — High Load**
10. **962F — Simple Cycles Edges**
11. **1076D — Edge Deletion**
12. **1450E — Capitalism**

Codeforces explicitly treats “does removing an edge disconnect the graph?” as the bridge-recognition trigger. [Codeforces](https://codeforces.com/blog/entry/156481?locale=en\&utm_source=chatgpt.com)

---

# 23. Articulation Points — Critical Vertices

### Observation

If removing a vertex can disconnect the graph:

> **Articulation point**

### Problems

1. **1239D — Catowice City**
2. **118E — Bertown Roads**
3. **1000E — We Need More Bosses**
4. **732F — Tourist Reform**
5. **962F — Simple Cycles Edges**
6. **487E — Tourists**
7. **700E — Cool Slogans**
8. **652E — Pursuit For Artifacts**
9. **610D — Vika and Segments**
10. **796D — Police Stations**

---

# 24. Bridge Tree / 2-Edge-Connected Components

### Pattern

First:

```text
find bridges
```

Then:

```text
remove bridges
```

Each remaining component becomes a node.

The result is a **tree**.

### Problems

1. **1000E — We Need More Bosses**
2. **732F — Tourist Reform**
3. **652E — Pursuit For Artifacts**
4. **118E — Bertown Roads**
5. **962F — Simple Cycles Edges**
6. **610D — Vika and Segments**
7. **1276B — Two Fairs**
8. **796D — Police Stations**
9. **1707C — DFS Trees**
10. **1076D — Edge Deletion**

---

# 25. Euler Path / Euler Circuit

### Observation

Every edge must be used exactly once.

Think:

> **Hierholzer**

Conditions for undirected Euler path:

```text
0 odd-degree → Euler circuit
2 odd-degree → Euler path
>2 → impossible
```

### Problems

1. **508D — Tanya and Password**
2. **501C — Misha and Forest**
3. **510E — Fox And Dinner**
4. **29C — Mail Stamps**
5. **605E — Intergalaxy Trips**
6. **1270G — Subset with Zero Sum**
7. **138D — World of Darkraft**
8. **1179A — Valeriy and Deque**
9. **1364D — Ehab's Last Corollary**
10. **1490G — Old Floppy Drive**

### Must know

```text
degree condition
+
connectivity
+
Hierholzer
```

---

# 26. Functional Graph

### Pattern

Every node has exactly one outgoing edge:

```text
f(x)
```

Structure:

```text
cycle
  ↑
trees feeding into cycle
```

### Problems

1. **1029F — Multicolored Markers**
2. **859E — Desk Disorder**
3. **1167C — News Distribution**
4. **1833E — Round Dance**
5. **939A — Love Triangle**
6. **711D — Directed Roads**
7. **103B — Cthulhu**
8. **1450A — Avoid Trygub**
9. **1320B — Journey**
10. **1485F — Copy or Prefix Sum**
11. **1270E — Two Arrays and Swaps**
12. **1860E — Fast and Fat**

### Critical observation

If:

```text
outdegree(v) = 1
```

don't immediately treat it as a generic graph.

Think:

> **Functional graph → cycles + trees**

---

# 27. Functional Graph — Cycle Detection / Cycle Length

### Problems

1. **939A — Love Triangle**
2. **711D — Directed Roads**
3. **1833E — Round Dance**
4. **859E — Desk Disorder**
5. **1029F — Multicolored Markers**
6. **1167C — News Distribution**
7. **1320B — Journey**
8. **742B — Arpa’s obvious problem and Mehrdad’s terrible solution**
9. **979B — Treasure Hunt**
10. **977F — Consecutive Subsequence**

---

# 28. Functional Graph — Binary Lifting

### Pattern

Repeated application:

```text
f(x)
f(f(x))
f(f(f(x)))
...
```

means:

> **binary lifting**

### Problems

1. **702E — Analysis of Pathes in Functional Graph**
2. **932E — Team Work**
3. **877E — Danil and a Part-time Job**
4. **832D — Misha, Grisha and Underground**
5. **1029F — Multicolored Markers**
6. **1167C — News Distribution**
7. **1320B — Journey**
8. **1516D — Cut**
9. **1547F — Array Stabilization**
10. **1485E — Move and Swap**

---

# 29. Tree Diameter

### Pattern

Questions involving:

```text
farthest node
maximum distance
longest path
tree eccentricity
```

Think:

> **Tree diameter**

Classic:

```text
BFS/DFS from arbitrary node
        ↓
farthest u
        ↓
BFS/DFS from u
        ↓
farthest v
        ↓
u-v = diameter
```

### Problems

1. **1294F — Three Paths on a Tree**
2. **1092F — Tree with Maximum Cost**
3. **1187E — Tree Painting**
4. **1089E — Alice and the Unfinished Chessboard**
5. **101B — Ciel and Duel**
6. **827A — String Reconstruction**
7. **1092D2 — Great Vova Wall**
8. **1029F — Multicolored Markers**
9. **1081D — Maximum Distance**
10. **911F — Tree Destruction**
11. **115A — Party**
12. **1060E — Sergey and Subway**

---

# 30. Tree Diameter — “Three/Four Special Nodes” Observation

### Pattern

If problem asks for:

```text
maximum total distance
choose 2/3/4 vertices
cover tree
maximize union of paths
```

look at the diameter first.

### Problems

1. **1294F — Three Paths on a Tree**
2. **911F — Tree Destruction**
3. **1081D — Maximum Distance**
4. **1060E — Sergey and Subway**
5. **1187E — Tree Painting**
6. **1092F — Tree with Maximum Cost**
7. **1089E — Alice and the Unfinished Chessboard**
8. **1254D — Fedor Runs for President**
9. **1092E — Minimal Diameter Forest**
10. **1245F — Daniel and Spring Cleaning**

---

# 31. LCA — Lowest Common Ancestor

### Pattern

Queries involving:

```text
distance(u,v)
ancestor
path between u and v
common ancestor
```

→ **LCA**

### Problems

1. **1328E — Tree Queries**
2. **832D — Misha, Grisha and Underground**
3. **191C — Fools and Roads**
4. **609E — Minimum spanning tree for each edge**
5. **587C — Duff in the Army**
6. **519E — A and B and Team Training**
7. **208E — Blood Cousins**
8. **696B — Puzzles**
9. **609B — C. Load Testing**
10. **1006E — Military Problem**
11. **715C — Digit Tree**
12. **1076E — Vasya and a Tree**

Codeforces graph roadmaps specifically identify LCA as the standard pattern for tree ancestor/path structure. [Codeforces](https://codeforces.com/blog/entry/156481?locale=en\&utm_source=chatgpt.com)

---

# 32. Binary Lifting

### Observation

If queries repeatedly ask:

```text
k-th ancestor
jump k steps
ancestor at distance k
```

use:

```text
up[v][j]
```

### Problems

1. **208E — Blood Cousins**
2. **1328E — Tree Queries**
3. **832D — Misha, Grisha and Underground**
4. **1006E — Military Problem**
5. **519E — A and B and Team Training**
6. **587C — Duff in the Army**
7. **1076E — Vasya and a Tree**
8. **1516D — Cut**
9. **702E — Analysis of Pathes in Functional Graph**
10. **609E — Minimum spanning tree for each edge**
11. **1702G1 — Passable Paths**
12. **1702G2 — Passable Paths (hard version)**

---

# 33. Tree Path Queries

### Pattern

Queries like:

```text
u → v
sum on path
minimum on path
maximum on path
number of marked nodes
update path
```

Possible tools:

```text
LCA
Binary Lifting
HLD
Euler Tour
Segment Tree
```

### Problems

1. **191C — Fools and Roads**
2. **1328E — Tree Queries**
3. **587C — Duff in the Army**
4. **1076E — Vasya and a Tree**
5. **570D — Tree Requests**
6. **1006E — Military Problem**
7. **1353E — K-th Not Divisible by n**
8. **696B — Puzzles**
9. **1213G — Path Queries**
10. **915F — Imbalance Value of a Tree**
11. **1702G1 — Passable Paths**
12. **1702G2 — Passable Paths**

---

# 34. HLD — Heavy-Light Decomposition

### Pattern

Tree + path queries + updates.

Think:

> **Decompose one path into O(log n) heavy chains.**

### Problems

1. **343D — Water Tree**
2. **1172E — Nauuo and ODT**
4. **1009F — Dominant Indices**
5. **741D — Arpa’s letter-marked tree and Mehrdad's Dokhtar-kosh paths**
6. **613D — Kingdom and its Cities**
7. **343E — Pumping Stations**
8. **504E — Misha and LCP on Tree**
9. **877E — Danil and a Part-time Job**
10. **1187E — Tree Painting**
11. **587C — Duff in the Army**
12. **208E — Blood Cousins**

---

# 35. Euler Tour / Flatten Tree

### Pattern

Subtree queries:

```text
subtree(v)
```

can become:

```text
tin[v] ... tout[v]
```

in an Euler traversal.

Then use:

```text
Fenwick
Segment Tree
Prefix Sum
```

### Problems

1. **877E — Danil and a Part-time Job**
2. **1006E — Military Problem**
3. **1076E — Vasya and a Tree**
4. **570D — Tree Requests**
5. **620E — New Year Tree**
6. **383C — Propagating tree**
7. **893F — Subtree Minimum Query**
8. **696B — Puzzles**
9. **208E — Blood Cousins**
10. **274B — Zero Tree**
11. **1324F — Maximum White Subtree**
12. **1187E — Tree Painting**

---

# 36. Tree DP

### Pattern

Answer depends on children:

```text
dp[u] = combine(dp[child])
```

### Problems

1. **1324F — Maximum White Subtree**
2. **274B — Zero Tree**
3. **1187E — Tree Painting**
4. **1092F — Tree with Maximum Cost**
5. **1336A — Linova and Kingdom**
6. **1249F — Maximum Weight Subset**
7. **1248E — Let Them Slide**
8. **763A — Timofey and a Tree**
9. **1083A — The Fair Nut and the Best Path**
10. **911F — Tree Destruction**
11. **607A — Chain Reaction**
12. **1118F1 — Tree Cutting**
13. **1118F2 — Tree Cutting**

---

# 37. Rerooting DP

### Observation

You calculated:

```text
answer for root = 1
```

but need:

```text
answer for every possible root
```

Then:

> **rerooting DP**

Typical transition:

```text
ans[child] = ans[parent] + contribution_change
```

### Problems

1. **1092F — Tree with Maximum Cost**
2. **1187E — Tree Painting**
3. **1324F — Maximum White Subtree**
4. **1336A — Linova and Kingdom**
5. **833B — The Bakery**
6. **1320E — Treeland and Viruses**
7. **600E — Lomsat gelral**
8. **763A — Timofey and a Tree**
9. **1083A — The Fair Nut and the Best Path**
10. **911F — Tree Destruction**
11. **1249F — Maximum Weight Subset**
12. **1294F — Three Paths on a Tree**

---

# 38. Complement Graph

### Observation

Graph gives edges that **do not** represent the relationship you care about.

Instead of explicitly constructing:

```text
complement graph
```

traverse unvisited nodes while skipping original neighbors.

### Problems

1. **920E — Connected Components?**
2. **1242B — 0-1 MST**
3. **1391E — 505**
4. **990D — Graph And Its Complement**
5. **986C — AND Graph**
6. **1598E — Air Conditioners**
7. **1179A — Valeriy and Deque**
8. **1245D — Shichikuji and Power Grid**
9. **1553G — Common Divisor Graph**
10. **2081D — MST in Modulo Graph**

---

# 39. Implicit Graph

### Pattern

Vertices exist, but edges aren't explicitly given.

Examples:

```text
number → numbers obtainable by operation
string → strings after transformation
state → next state
mask → neighboring mask
```

### Problems

1. **520B — Two Buttons**
2. **500A — New Year Transportation**
3. **1272E — Nearest Opposite Parity**
4. **986C — AND Graph**
5. **1242B — 0-1 MST**
6. **1105D — Kilani and the Game**
7. **1350E — Orac and Game of Life**
8. **1485E — Move and Swap**
9. **939A — Love Triangle**
10. **711D — Directed Roads**
11. **1553G — Common Divisor Graph**
12. **2081D — MST in Modulo Graph**

---

# 40. Graph Compression

### Observation

If many vertices behave identically / are mutually reachable:

```text
compress them
```

into one node.

Common compression types:

```text
SCC compression
DSU compression
Bridge compression
equal-value compression
```

### Problems

1. **427C — Checkposts**
2. **999E — Reachability from the Capital**
3. **1000E — We Need More Bosses**
4. **732F — Tourist Reform**
5. **118E — Bertown Roads**
6. **1553G — Common Divisor Graph**
7. **1242B — 0-1 MST**
8. **987F — AND Graph**
9. **609E — Minimum spanning tree for each edge**
10. **1707C — DFS Trees**

---

# 41. Bitmask Graph

### Pattern

Vertex = mask.

Edges based on:

```text
AND
OR
XOR
submask
supermask
```

### Problems

1. **986C — AND Graph**
2. **987F — AND Graph**
3. **888G — Xor-MST**
4. **959E — Mahmoud and Ehab and the xor-MST**
5. **1615D — Xor-Paths**
6. **1285E — Delete a Segment**
7. **1520F1 — Guess the K-th Zero**
8. **453B — Little Pony and Harmony Chest**
9. **1013E — Hills**
10. **165E — Compatible Numbers**
11. **1225E — Rock Is Push**
12. **850D — Wisteria**

---

# 42. XOR Graph / XOR-MST

### Pattern

When edge cost looks like:

```text
a[i] XOR a[j]
```

do **not** construct all edges.

Think:

```text
bitwise trie
divide by highest differing bit
XOR-MST
```

### Problems

1. **888G — Xor-MST**
2. **959E — Mahmoud and Ehab and the xor-MST**
3. **986C — AND Graph**
4. **987F — AND Graph**
5. **1615D — Xor-Paths**
6. **1625D — Binary Spiders**
7. **706D — Vasiliy's Multiset**
8. **923C — Perfect Security**
9. **665E — Beautiful Subarrays**
10. **845G — Shortest Path Problem?**
11. **1285D — Dr. Evil Underscores**
12. **742E — Arpa's obvious problem and Mehrdad's terrible solution**

---

# 43. Graph + Greedy

### Observation

Graph doesn't necessarily mean graph algorithm only.

Often:

```text
graph structure
+
greedy choice
```

### Problems

1. **1245D — Shichikuji and Power Grid**
2. **449B — Jzzhu and Cities**
3. **609E — Minimum spanning tree for each edge**
4. **1108F — MST Unification**
5. **733F — Drivers Dissatisfaction**
6. **1133F2 — Spanning Tree with One Fixed Degree**
7. **1515F — Phoenix and Earthquake**
8. **1537F — Figure Fixing**
9. **2029D — Cool Graph**
10. **2081D — MST in Modulo Graph**
11. **2117G — Omg Graph**
12. **2208D2 — Tree Orientation**

---

# 44. Graph + Binary Search

### Observation

Question:

> “Can this be done with maximum edge ≤ X?”

If yes/no is monotonic:

```text
binary search answer
+
graph feasibility check
```

### Problems

1. **1679D — Toss a Coin to Your Graph**
2. **1851G — Vlad and the Mountains**
3. **2057E1 — Another Exercise on Graphs**
4. **2057E2 — Another Exercise on Graphs**
5. **1706E — Qpwoeirut And Vertices**
6. **1633E — Spanning Tree Queries**
7. **1550F — Jumping Around**
8. **938G — Shortest Path Queries**
9. **1245D — Shichikuji and Power Grid**
10. **609E — Minimum spanning tree for each edge**
11. **901C — Bipartite Segments**
12. **2275G — Copper Squander**

---

# 45. Graph + DP

### Pattern

Graph controls transitions:

```text
dp[u] = answer obtained from neighbors
```

Common structures:

```text
DAG DP
Tree DP
SCC + DP
Functional graph DP
Shortest path + DP
```

### Problems

1. **919D — Substring**
2. **721C — Journey**
3. **1472G — Moving to the Capital**
4. **1547G — How Many Paths?**
5. **960F — Pathwalks**
6. **1324F — Maximum White Subtree**
7. **1092F — Tree with Maximum Cost**
8. **1187E — Tree Painting**
9. **274B — Zero Tree**
10. **1249F — Maximum Weight Subset**
11. **711D — Directed Roads**
12. **982C — Cut 'em all!**

---

# 46. Graph + Counting

### Observation

Graph asks:

```text
how many ways?
number of valid paths?
number of components?
number of arrangements?
```

Often:

```text
graph decomposition
+
combinatorics / DP
```

### Problems

1. **427C — Checkposts**
2. **711D — Directed Roads**
3. **1547G — How Many Paths?**
4. **1139C — Edgy Trees**
5. **1534C — Little Alawn's Puzzle**
6. **1172B — Nauuo and Circle**
7. **1242C — Sum Balance**
8. **859E — Desk Disorder**
9. **1472G — Moving to the Capital**
10. **987F — AND Graph**

---

# 47. Graph Coloring

### Pattern

Not only bipartite.

Possible colors:

```text
2 colors
k colors
parity
constraints
DSATUR-like greedy
```

### Problems

1. **862B — Mahmoud and Ehab and the bipartiteness**
2. **687A — NP-Hard Problem**
3. **1354E — Graph Coloring**
4. **1093D — Beautiful Graph**
5. **1615D — Xor-Paths**
6. **1272D — Remove One Element**
7. **1537F — Figure Fixing**
8. **891C — Envy**
9. **813F — Bipartite Checking**
10. **2260F — Edge Three-Coloring**
11. **2135C — By the Assignment**
12. **2022E1 — Billetes MX**

---

# 48. Graph Matching

### Pattern

If problem says:

```text
pair vertices
match left side to right side
each vertex used at most once
maximum pairing
```

think:

> **Bipartite Matching / Flow**

### Problems

1. **468B — Two Sets**
2. **468C — Strange Calculation and Cats**
3. **1140E — Palindromes**
4. **1244E — Pens and Pencils**
5. **1153E — Serval and Snake**
6. **1198B — Welfare State**
7. **1637E — Best Pair**
8. **2207F — Hanabi**
9. **1549B — Gregor and the Pawn Game**
10. **1082G — Petya and Graph**

---

# 49. Max Flow

### Pattern

Think flow when:

```text
capacity
source → sink
maximum number
minimum cut
assignment
matching
```

### Problems

1. **653D — Delivery Bears**
2. **546E — Soldier and Traveling**
3. **808E — Selling Souvenirs**
4. **1139D — Steps to One**
5. **498C — Array and Operations**
6. **827D — Best Edge**
7. **990G — GCD Counting**
8. **1156D — 0-1-Tree**
9. **1549B — Gregor and the Pawn Game**
10. **2207F — Hanabi**
11. **176E — Archaeology**
12. **1082G — Petya and Graph**

---

# 50. Flow Disguised as Matching

### Observation

If:

```text
each left item chooses one right item
capacity = 1
```

you may have a bipartite matching problem even though the statement never says "matching".

### Problems

1. **1549B — Gregor and the Pawn Game**
2. **468B — Two Sets**
3. **1082G — Petya and Graph**
4. **2207F — Hanabi**
5. **1638E — Colorful Operations**
6. **101B — Ciel and Duel**
7. **1099F — Cookies**
8. **1139D — Steps to One**
9. **546E — Soldier and Traveling**
10. **653D — Delivery Bears**

---

# 51. 2-SAT

### Pattern

Boolean constraints such as:

```text
x OR y
x → y
not x → y
choose exactly one
```

Convert to implication graph:

```text
(x OR y)
=> (!x -> y)
=> (!y -> x)
```

Then SCC.

### Problems

1. **776B — Sherlock and his girlfriend**
2. **1215F — Radio Stations**
3. **228E — Is it rated?**
4. **776D — The Door Problem**
5. **1428E — Carrots for Rabbits**
6. **1029F — Multicolored Markers**
7. **1472G — Moving to the Capital**
8. **1391E — 505**
9. **206E — Computer Network**
10. **2022E1 — Billetes MX**
11. **1254D — Fedor Runs for President**
12. **1354E — Graph Coloring**

---

# 52. Game on Graph

### Pattern

Vertices = game states.

Edges = legal moves.

Then:

```text
winning[u] =
    exists child that is losing
```

### Problems

1. **138D — World of Darkraft**
2. **936B — Sleepy Game**
3. **455B — A Lot of Games**
4. **603C — Lieges of Legendre**
5. **317B — Game of Credit Cards**
6. **1738F — Connectivity Addicts**
7. **731F — Video Cards**
8. **1242C — Sum Balance**
9. **1466E — Apollo versus Pan**
10. **1598E — Air Conditioners**

---

# 53. Shortest Path + “Which Edges Can Be on a Shortest Path?”

### Observation

Run:

```text
distS[]
distT[]
```

For edge:

```text
u -> v
```

it can be part of a shortest path if:

```text
distS[u] + w + distT[v] == distS[t]
```

This is a **very important Codeforces observation**.

### Problems

1. **449B — Jzzhu and Cities**
2. **1076D — Edge Deletion**
3. **545E — Paths and Trees**
4. **20C — Dijkstra?**
5. **938D — Buy a Ticket**
6. **786B — Weird Journey**
7. **1473E — Minimum Path**
8. **1547G — How Many Paths?**
9. **1196F — K-th Path**
10. **1915G — Bicycles**

---

# 54. Shortest Path + Remove / Keep Edges

### Pattern

Question asks:

```text
remove edges while preserving shortest paths
```

or:

```text
minimum edges needed to preserve distances
```

### Problems

1. **1076D — Edge Deletion**
2. **449B — Jzzhu and Cities**
3. **545E — Paths and Trees**
4. **609E — Minimum spanning tree for each edge**
5. **20C — Dijkstra?**
6. **938D — Buy a Ticket**
7. **1473E — Minimum Path**
8. **786B — Weird Journey**
9. **1547G — How Many Paths?**
10. **1196F — K-th Path**

---

# 55. “Can Reach?” → Reverse Thinking

### Observation

If the question is:

> Which nodes can eventually reach `X`?

reverse the graph and start BFS/DFS from `X`.

### Problems

1. **999E — Reachability from the Capital**
2. **1320B — Journey**
3. **1547G — How Many Paths?**
4. **427C — Checkposts**
5. **919D — Substring**
6. **721C — Journey**
7. **1385E — Directing Edges**
8. **510C — Fox And Names**
9. **1472G — Moving to the Capital**
10. **1187E — Tree Painting**

---

# 56. “All Nodes Reach One Node” → Sink / Reverse Graph

### Problems

1. **999E — Reachability from the Capital**
2. **427C — Checkposts**
3. **1547G — How Many Paths?**
4. **1320B — Journey**
5. **1472G — Moving to the Capital**
6. **1385E — Directing Edges**
7. **919D — Substring**
8. **721C — Journey**
9. **475B — Strongly Connected City**
10. **1000E — We Need More Bosses**

---

# 57. Graph Degree Observation

### Pattern

Sometimes you don't need DFS.

Look at:

```text
degree[v]
```

and derive the answer.

### Problems

1. **550D — Regular Bridge**
2. **501C — Misha and Forest**
3. **977E — Cyclic Components**
4. **1833E — Round Dance**
5. **1209D — Cow and Snacks**
6. **115A — Party**
7. **1133F2 — Spanning Tree with One Fixed Degree**
8. **1515F — Phoenix and Earthquake**
9. **1944A — Destroying Bridges**
10. **2258F — Plus Minus Tree**

---

# 58. Graph Edge Count Observation

### Extremely useful

For a connected undirected graph:

```text
tree:
m = n - 1

unicyclic:
m = n

more cycles:
m > n
```

### Problems

1. **977E — Cyclic Components**
2. **1209D — Cow and Snacks**
3. **25D — Roads not only in Berland**
4. **1833E — Round Dance**
5. **755C — PolandBall and Forest**
6. **550D — Regular Bridge**
7. **1139C — Edgy Trees**
8. **1000E — We Need More Bosses**
9. **732F — Tourist Reform**
10. **118E — Bertown Roads**

---

# 59. Tree Is Just a Special Graph

### Observation

When you see:

```text
n vertices
n-1 edges
connected
```

immediately think:

> **TREE**

Then exploit:

```text
unique path
no cycle
DFS
subtree
diameter
LCA
rerooting
tree DP
```

### Problems

1. **1294F — Three Paths on a Tree**
2. **1324F — Maximum White Subtree**
3. **1092F — Tree with Maximum Cost**
4. **1187E — Tree Painting**
5. **911F — Tree Destruction**
6. **1336A — Linova and Kingdom**
7. **274B — Zero Tree**
8. **600E — Lomsat gelral**
9. **1083A — The Fair Nut and the Best Path**
10. **1249F — Maximum Weight Subset**
11. **570D — Tree Requests**
12. **1006E — Military Problem**

---

# 60. “Tree + Extra Edge” Observation

### Pattern

Graph is almost a tree:

```text
n vertices
n edges
```

Often:

> **exactly one cycle**

This leads to:

```text
find cycle
+
treat cycle specially
+
trees attached to cycle
```

### Problems

1. **977E — Cyclic Components**
2. **1029F — Multicolored Markers**
3. **1833E — Round Dance**
4. **550D — Regular Bridge**
5. **711D — Directed Roads**
6. **859E — Desk Disorder**
7. **609E — Minimum spanning tree for each edge**
8. **1209D — Cow and Snacks**
9. **25D — Roads not only in Berland**
10. **1242B — 0-1 MST**

---

# 61. Graph + Parity

### Pattern

Look for:

```text
odd/even distance
xor
parity constraint
alternating edges
```

Usually:

```text
BFS/DFS coloring
```

or

```text
DSU with parity
```

### Problems

1. **862B — Mahmoud and Ehab and the Bipartiteness**
2. **687A — NP-Hard Problem**
3. **1093D — Beautiful Graph**
4. **1354E — Graph Coloring**
5. **1615D — Xor-Paths**
6. **813F — Bipartite Checking**
7. **468B — Two Sets**
8. **1537F — Figure Fixing**
9. **2260F — Edge Three-Coloring**
10. **2135C — By the Assignment**

---

# 62. DSU With Extra Information

### Pattern

DSU doesn't have to store only:

```cpp
parent[]
size[]
```

It can maintain:

```text
parity
potential
xor
difference
color
weight
```

### Problems

1. **742B — Arpa’s obvious problem and Mehrdad's terrible solution**
2. **1615D — Xor-Paths**
3. **1209D — Cow and Snacks**
4. **891C — Envy**
5. **813F — Bipartite Checking**
6. **1537F — Figure Fixing**
7. **1553G — Common Divisor Graph**
8. **500B — New Year Permutation**
9. **1534C — Little Alawn's Puzzle**
10. **755C — PolandBall and Forest**

---

# 63. Offline Queries + Sort Edges

### Observation

If queries ask:

```text
answer using edges <= X
```

sort:

```text
edges by weight
queries by X
```

Then add edges using DSU.

### Problems

1. **1213G — Path Queries**
2. **891C — Envy**
3. **609E — Minimum spanning tree for each edge**
4. **1108F — MST Unification**
5. **1851G — Vlad and the Mountains**
6. **1706E — Qpwoeirut And Vertices**
7. **1633E — Spanning Tree Queries**
8. **1550F — Jumping Around**
9. **2117G — Omg Graph**
10. **2081D — MST in Modulo Graph**

---

# 64. Graph + Sweep / Sort Events

### Pattern

Events are edges/vertices appearing in some order.

Instead of recomputing the graph:

```text
sort events
maintain DSU / graph state
```

### Problems

1. **1213G — Path Queries**
2. **891C — Envy**
3. **1108F — MST Unification**
4. **609E — Minimum spanning tree for each edge**
5. **1706E — Qpwoeirut And Vertices**
6. **1851G — Vlad and the Mountains**
7. **1633E — Spanning Tree Queries**
8. **2117G — Omg Graph**
9. **1550F — Jumping Around**
10. **2021E2 — Digital Village**

---

# 65. Advanced Graph — Cactus

### Pattern

A graph where every edge belongs to at most one simple cycle.

Think:

```text
tree-like
+
individual cycles
```

### Problems

1. **980F — Cactus to Tree**
2. **962F — Simple Cycles Edges**
3. **932F — Escape Through Leaf**
4. **981C — Useful Decomposition**
5. **1000E — We Need More Bosses**
6. **732F — Tourist Reform**
7. **118E — Bertown Roads**
8. **982C — Cut 'em all!**
9. **1029F — Multicolored Markers**
10. **627F — Island Puzzle**

---

# 66. Advanced Graph — 2-Edge / 2-Vertex Connectivity

### Problems

1. **1000E — We Need More Bosses**
2. **732F — Tourist Reform**
3. **118E — Bertown Roads**
4. **962F — Simple Cycles Edges**
5. **487E — Tourists**
6. **652E — Pursuit For Artifacts**
7. **1707C — DFS Trees**
8. **610D — Vika and Segments**
9. **796D — Police Stations**
10. **1239D — Catowice City**

---

# 67. Advanced Graph — Link-Cut / Dynamic Tree Thinking

### Pattern

When the graph is a tree but operations dynamically:

```text
cut edge
link edge
query path
change root
```

think beyond ordinary DFS.

### Problems

1. **1172E — Nauuo and ODT**
2. **343D — Water Tree**
3. **613D — Kingdom and its Cities**
4. **504E — Misha and LCP on Tree**
5. **587C — Duff in the Army**
6. **877E — Danil and a Part-time Job**
7. **343E — Pumping Stations**
8. **1187E — Tree Painting**
9. **1009F — Dominant Indices**
10. **741D — Arpa’s letter-marked tree and Mehrdad's Dokhtar-kosh paths**

---

# 68. Graph + Segment Tree

### Pattern

Graph/tree traversal converts the problem into:

```text
range query
range update
point query
path query
```

### Problems

1. **877E — Danil and a Part-time Job**
2. **620E — New Year Tree**
3. **383C — Propagating tree**
4. **1076E — Vasya and a Tree**
5. **570D — Tree Requests**
6. **600E — Lomsat gelral**
7. **343D — Water Tree**
8. **741D — Arpa’s letter-marked tree**
9. **504E — Misha and LCP on Tree**
10. **587C — Duff in the Army**

---

# 69. Graph + Fenwick Tree

### Pattern

Euler flattening:

```text
subtree = interval
```

then Fenwick.

### Problems

1. **877E — Danil and a Part-time Job**
2. **383C — Propagating tree**
3. **1076E — Vasya and a Tree**
4. **1006E — Military Problem**
5. **570D — Tree Requests**
6. **620E — New Year Tree**
7. **696B — Puzzles**
8. **208E — Blood Cousins**
9. **600E — Lomsat gelral**
10. **1324F — Maximum White Subtree**

---

# 70. The Most Important “Graph Observation” Problems

If your goal is **Codeforces pattern recognition**, these are particularly valuable:

| Problem | Main observation |
|---|---|
| **977E — Cyclic Components** | Component + degree/cycle |
| **1209D — Cow and Snacks** | DSU + redundant edges |
| **1242B — 0-1 MST** | Complement graph |
| **920E — Connected Components?** | Complement graph traversal |
| **986C — AND Graph** | Implicit graph + bitmask |
| **987F — AND Graph** | Bitmask connectivity |
| **1000E — We Need More Bosses** | Bridges → bridge tree |
| **732F — Tourist Reform** | Bridges + component orientation |
| **427C — Checkposts** | SCC condensation |
| **999E — Reachability from the Capital** | SCC/reachability |
| **1547G — How Many Paths?** | SCC + DAG DP |
| **919D — Substring** | Topological sort + DP |
| **1472G — Moving to the Capital** | DAG/SCC-like graph DP |
| **1679D — Toss a Coin to Your Graph** | Binary search + cycle/DAG |
| **609E — MST for each edge** | MST + max edge on path |
| **1213G — Path Queries** | Sort edges + DSU |
| **891C — Envy** | Kruskal groups + DSU |
| **1108F — MST Unification** | Equal-weight MST handling |
| **449B — Jzzhu and Cities** | Dijkstra + redundant shortest paths |
| **1076D — Edge Deletion** | Dijkstra tree reconstruction |
| **1473E — Minimum Path** | State-expanded Dijkstra |
| **20C — Dijkstra?** | Standard shortest path |
| **1328E — Tree Queries** | LCA + deepest node observation |
| **191C — Fools and Roads** | Tree path counting |
| **1092F — Tree with Maximum Cost** | Rerooting DP |
| **1187E — Tree Painting** | Rerooting |
| **1294F — Three Paths on a Tree** | Diameter + tree observation |
| **911F — Tree Destruction** | Diameter-based greedy |
| **570D — Tree Requests** | Euler tour + parity bitmask |
| **877E — Danil and a Part-time Job** | Euler tour + segment tree |
| **741D — Arpa's Tree** | Tree path + bitmask |
| **702E — Analysis of Pathes in Functional Graph** | Functional graph + binary lifting |
| **711D — Directed Roads** | Functional graph + cycles |
| **1833E — Round Dance** | Functional graph components |
| **508D — Tanya and Password** | Euler path in directed graph |
| **550D — Regular Bridge** | Degree/constructive observation |
| **776D — The Door Problem** | 2-SAT |
| **228E — Is it rated?** | 2-SAT |
| **546E — Soldier and Traveling** | Flow |
| **653D — Delivery Bears** | Flow + binary search |

---

# Recommended Order 🔥

Don't solve these 70 categories randomly. For **Codeforces graph mastery**, I would use this progression:

```text
PHASE 1 — FUNDAMENTALS
│
├── DFS / BFS
├── Connected Components
├── Cycle Detection
├── Bipartite
└── Grid / Implicit Graph


PHASE 2 — DSU
│
├── Basic DSU
├── Cycle Detection with DSU
├── Offline Connectivity
├── DSU + Sorting
└── DSU with Extra Information


PHASE 3 — SHORTEST PATH
│
├── BFS shortest path
├── Multi-source BFS
├── 0-1 BFS
├── Dijkstra
├── State Dijkstra
├── Reverse Dijkstra
└── DAG shortest path


PHASE 4 — DAG
│
├── Topological Sort
├── DAG DP
├── Longest Path
├── Number of Paths
└── SCC → Condensation → DAG


PHASE 5 — MST
│
├── Kruskal
├── Prim
├── MST + DSU
├── MST + LCA
├── MST for every edge
├── Equal-weight MST
└── XOR MST


PHASE 6 — SCC / CONNECTIVITY
│
├── SCC
├── Condensation
├── Bridges
├── Articulation Points
├── Bridge Tree
└── 2-Edge Connectivity


PHASE 7 — TREES
│
├── Tree Diameter
├── Tree DP
├── Rerooting
├── Euler Tour
├── LCA
├── Binary Lifting
├── Path Queries
└── HLD


PHASE 8 — SPECIAL GRAPHS
│
├── Functional Graph
├── Complement Graph
├── Bitmask Graph
├── XOR Graph
├── Cactus
└── Implicit Graph


PHASE 9 — ADVANCED
│
├── 2-SAT
├── Matching
├── Max Flow
├── Min Cut
├── Graph + Binary Search
├── Graph + DP
└── Dynamic Trees
```

This ordering is consistent with the usual Codeforces progression from DFS/BFS, connectivity, shortest paths, DSU/MST and topological sort toward 0-1 BFS, SCC, bridges, articulation points, diameter, LCA, HLD and flow. [Codeforces](https://codeforces.com/blog/entry/156481?locale=en\&utm_source=chatgpt.com)

