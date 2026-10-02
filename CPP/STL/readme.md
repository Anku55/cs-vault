# C++ STL — Topic List

## 1. STL Fundamentals
- [ ] STL Overview
- [ ] Containers
- [ ] Iterators
- [ ] Algorithms
- [ ] Functors
- [ ] Function Objects
- [ ] Allocators
- [ ] Iterator Categories
- [ ] Complexity of STL Operations

## 2. `std::array`
- [ ] `std::array`
- [ ] Initialization
- [ ] `size()`
- [ ] `empty()`
- [ ] `front()`
- [ ] `back()`
- [ ] `at()`
- [ ] `operator[]`
- [ ] `fill()`
- [ ] `data()`
- [ ] Iteration

## 3. `std::vector` 🔥
- [ ] `vector`
- [ ] Initialization
- [ ] `push_back()`
- [ ] `emplace_back()`
- [ ] `pop_back()`
- [ ] `insert()`
- [ ] `erase()`
- [ ] `clear()`
- [ ] `resize()`
- [ ] `reserve()`
- [ ] `capacity()`
- [ ] `shrink_to_fit()`
- [ ] `front()`
- [ ] `back()`
- [ ] `data()`
- [ ] Iterator Invalidation
- [ ] Vector Memory Layout
- [ ] Capacity Growth
- [ ] `vector<bool>`

## 4. `std::string`
- [ ] `string`
- [ ] Construction
- [ ] `size()`
- [ ] `length()`
- [ ] `empty()`
- [ ] `append()`
- [ ] `insert()`
- [ ] `erase()`
- [ ] `replace()`
- [ ] `substr()`
- [ ] `find()`
- [ ] `rfind()`
- [ ] `compare()`
- [ ] `c_str()`
- [ ] `data()`
- [ ] `getline()`
- [ ] String Concatenation
- [ ] String Views

## 5. `std::deque`
- [ ] `deque`
- [ ] `push_front()`
- [ ] `push_back()`
- [ ] `pop_front()`
- [ ] `pop_back()`
- [ ] Random Access
- [ ] Insert / Erase
- [ ] Iterator Invalidation
- [ ] `deque` vs `vector`

## 6. `std::list`
- [ ] `list`
- [ ] `push_front()`
- [ ] `push_back()`
- [ ] `insert()`
- [ ] `erase()`
- [ ] `remove()`
- [ ] `remove_if()`
- [ ] `sort()`
- [ ] `reverse()`
- [ ] `unique()`
- [ ] `merge()`
- [ ] Splicing
- [ ] Iterator Invalidation
- [ ] `list` vs `vector`

## 7. `std::forward_list`
- [ ] `forward_list`
- [ ] `push_front()`
- [ ] `insert_after()`
- [ ] `erase_after()`
- [ ] `remove()`
- [ ] `sort()`
- [ ] `reverse()`
- [ ] `merge()`
- [ ] `splice_after()`
- [ ] `forward_list` vs `list`

## 8. Associative Containers
### `std::set`
- [ ] `set`
- [ ] Insertion
- [ ] Deletion
- [ ] Search
- [ ] `find()`
- [ ] `count()`
- [ ] `lower_bound()`
- [ ] `upper_bound()`
- [ ] `equal_range()`

### `std::multiset`
- [ ] `multiset`
- [ ] Duplicate Elements
- [ ] Search
- [ ] Range Queries

### `std::map`
- [ ] `map`
- [ ] Key-Value Pairs
- [ ] `insert()`
- [ ] `emplace()`
- [ ] `erase()`
- [ ] `find()`
- [ ] `count()`
- [ ] `operator[]`
- [ ] `at()`
- [ ] `lower_bound()`
- [ ] `upper_bound()`
- [ ] `equal_range()`

### `std::multimap`
- [ ] `multimap`
- [ ] Duplicate Keys
- [ ] Range Queries
- [ ] `equal_range()`

## 9. Unordered Containers 
### `std::unordered_set`
- [ ] Hash Tables
- [ ] `unordered_set`
- [ ] Insert
- [ ] Erase
- [ ] Find
- [ ] Bucket Count
- [ ] Load Factor
- [ ] `reserve()`
- [ ] `rehash()`

### `std::unordered_map`
- [ ] `unordered_map`
- [ ] Hashing
- [ ] Key-Value Storage
- [ ] `operator[]`
- [ ] `at()`
- [ ] `find()`
- [ ] `insert()`
- [ ] `emplace()`
- [ ] `erase()`
- [ ] Buckets
- [ ] Load Factor
- [ ] Rehashing
- [ ] Custom Hash

### `std::unordered_multiset`
- [ ] `unordered_multiset`

### `std::unordered_multimap`
- [ ] `unordered_multimap`

## 10. Container Adapters
### `std::stack`
- [ ] `stack`
- [ ] `push()`
- [ ] `emplace()`
- [ ] `pop()`
- [ ] `top()`
- [ ] `empty()`
- [ ] `size()`

### `std::queue`
- [ ] `queue`
- [ ] `push()`
- [ ] `emplace()`
- [ ] `pop()`
- [ ] `front()`
- [ ] `back()`
- [ ] `empty()`

### `std::priority_queue`
- [ ] `priority_queue`
- [ ] Max Heap
- [ ] Min Heap
- [ ] `push()`
- [ ] `emplace()`
- [ ] `pop()`
- [ ] `top()`
- [ ] Custom Comparator
- [ ] Custom Object Priority

## 11. `std::pair` & `std::tuple`
- [ ] `pair`
- [ ] `make_pair()`
- [ ] `.first`
- [ ] `.second`
- [ ] `tuple`
- [ ] `make_tuple()`
- [ ] `get<>`
- [ ] Structured Bindings
- [ ] `tie()`
- [ ] `ignore`

## 12. Iterators 
- [ ] Iterator Basics
- [ ] `begin()`
- [ ] `end()`
- [ ] `cbegin()`
- [ ] `cend()`
- [ ] `rbegin()`
- [ ] `rend()`
- [ ] `crbegin()`
- [ ] `crend()`
- [ ] Iterator Dereferencing
- [ ] Iterator Arithmetic
- [ ] Iterator Invalidation
- [ ] Const Iterators

### Iterator Categories
- [ ] Input Iterator
- [ ] Output Iterator
- [ ] Forward Iterator
- [ ] Bidirectional Iterator
- [ ] Random Access Iterator
- [ ] Contiguous Iterator

## 13. STL Algorithms 
### Searching
- [ ] `find()`
- [ ] `find_if()`
- [ ] `find_if_not()`
- [ ] `binary_search()`
- [ ] `lower_bound()`
- [ ] `upper_bound()`
- [ ] `equal_range()`

### Sorting
- [ ] `sort()`
- [ ] `stable_sort()`
- [ ] `partial_sort()`
- [ ] `nth_element()`
- [ ] `is_sorted()`
- [ ] `is_sorted_until()`

### Modification
- [ ] `copy()`
- [ ] `copy_if()`
- [ ] `move()`
- [ ] `fill()`
- [ ] `fill_n()`
- [ ] `replace()`
- [ ] `replace_if()`
- [ ] `swap()`
- [ ] `iter_swap()`

### Removal
- [ ] `remove()`
- [ ] `remove_if()`
- [ ] `unique()`
- [ ] Erase-Remove Idiom

### Reordering
- [ ] `reverse()`
- [ ] `rotate()`
- [ ] `shuffle()`
- [ ] `partition()`
- [ ] `stable_partition()`

## 14. Numeric Algorithms
- [ ] `accumulate()`
- [ ] `reduce()`
- [ ] `inner_product()`
- [ ] `partial_sum()`
- [ ] `adjacent_difference()`
- [ ] `iota()`
- [ ] `gcd()`
- [ ] `lcm()`

## 15. Min / Max Algorithms
- [ ] `min()`
- [ ] `max()`
- [ ] `min_element()`
- [ ] `max_element()`
- [ ] `minmax()`
- [ ] `minmax_element()`

## 16. Set Algorithms
- [ ] `set_union()`
- [ ] `set_intersection()`
- [ ] `set_difference()`
- [ ] `set_symmetric_difference()`
- [ ] `includes()`

## 17. Heap Algorithms
- [ ] `make_heap()`
- [ ] `push_heap()`
- [ ] `pop_heap()`
- [ ] `sort_heap()`
- [ ] `is_heap()`
- [ ] `is_heap_until()`

## 18. Permutation Algorithms
- [ ] `next_permutation()`
- [ ] `prev_permutation()`
- [ ] `is_permutation()`

## 19. Lambda & Functors
- [ ] Lambda Functions
- [ ] Lambda Parameters
- [ ] Lambda Return Type
- [ ] Capture by Value
- [ ] Capture by Reference
- [ ] Generic Lambdas
- [ ] Mutable Lambdas
- [ ] Function Objects
- [ ] Custom Comparators

## 20. Function Utilities
- [ ] `std::function`
- [ ] `std::bind`
- [ ] `std::invoke`
- [ ] `std::ref`
- [ ] `std::cref`
- [ ] `std::mem_fn`

## 21. Comparators
- [ ] `std::less`
- [ ] `std::greater`
- [ ] `std::equal_to`
- [ ] `std::not_equal_to`
- [ ] Custom Comparators
- [ ] Lambda Comparators
- [ ] Three-Way Comparison
- [ ] `operator<=>`

## 22. Allocators & Memory
- [ ] STL Allocators
- [ ] `std::allocator`
- [ ] Custom Allocators
- [ ] Allocator-Aware Containers
- [ ] Memory Allocation
- [ ] Object Construction
- [ ] Object Destruction

## 23. Modern STL
- [ ] `std::optional`
- [ ] `std::variant`
- [ ] `std::any`
- [ ] `std::string_view`
- [ ] `std::span`
- [ ] `std::ranges`
- [ ] Range-Based Algorithms
- [ ] Views
- [ ] Concepts with STL

## 24. C++20 Ranges
- [ ] `std::ranges`
- [ ] `ranges::sort`
- [ ] `ranges::find`
- [ ] `ranges::filter_view`
- [ ] `ranges::transform_view`
- [ ] `ranges::take_view`
- [ ] `ranges::drop_view`
- [ ] Range Pipelines
- [ ] Lazy Evaluation
- [ ] Range Concepts

## 25. STL & Memory
- [ ] Contiguous Memory
- [ ] Dynamic Allocation
- [ ] Iterator Invalidation
- [ ] Reference Invalidation
- [ ] Move Semantics
- [ ] Copy Semantics
- [ ] `emplace` vs `insert`
- [ ] `reserve` vs `resize`
- [ ] Container Memory Overhead

## 26. STL Complexity
- [ ] Big-O of Containers
- [ ] Vector Complexity
- [ ] List Complexity
- [ ] Deque Complexity
- [ ] Map Complexity
- [ ] Set Complexity
- [ ] Unordered Map Complexity
- [ ] Stack Complexity
- [ ] Queue Complexity
- [ ] Priority Queue Complexity
- [ ] Algorithm Complexity

## 27. Container Selection 
- [ ] `vector` vs `array`
- [ ] `vector` vs `deque`
- [ ] `vector` vs `list`
- [ ] `map` vs `unordered_map`
- [ ] `set` vs `unordered_set`
- [ ] `map` vs `multimap`
- [ ] `set` vs `multiset`
- [ ] `stack` vs `vector`
- [ ] `queue` vs `deque`
- [ ] When to Use Each Container

## 28. STL for Systems Programming
- [ ] `vector` for Buffers
- [ ] `deque` for Queues
- [ ] `unordered_map` for Hash Tables
- [ ] `map` for Ordered Data
- [ ] `set` for Unique Keys
- [ ] `priority_queue` for Scheduling
- [ ] `queue` for Work Queues
- [ ] `string_view` for Zero-Copy Parsing
- [ ] `span` for Non-Owning Buffers
- [ ] `optional` for Optional Results
- [ ] `variant` for Protocol/Data Types
- [ ] Custom Allocators
- [ ] Iterator Invalidation
- [ ] Memory Efficiency

## 29. Redis-Relevant STL 
- [ ] `unordered_map`
- [ ] `map`
- [ ] `vector`
- [ ] `deque`
- [ ] `list`
- [ ] `set`
- [ ] `unordered_set`
- [ ] `priority_queue`
- [ ] `string`
- [ ] `string_view`
- [ ] `optional`
- [ ] `variant`
- [ ] Iterators
- [ ] Custom Hash
- [ ] Custom Comparator
- [ ] Allocators
- [ ] Memory Efficiency
- [ ] Container Complexity
- [ ] Iterator Invalidation
- [ ] Range Algorithms

## 30. STL Practice
- [ ] Implement Dynamic Array
- [ ] Implement Stack
- [ ] Implement Queue
- [ ] Implement Priority Queue
- [ ] Implement Linked List
- [ ] Implement Hash Table
- [ ] Implement Binary Search Tree
- [ ] Implement Iterator
- [ ] Implement Custom Comparator
- [ ] Implement Custom Allocator
- [ ] Implement LRU Cache
- [ ] Implement Thread-Safe Queue
- [ ] Implement In-Memory Key-Value Store