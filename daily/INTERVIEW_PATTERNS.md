# Interview Patterns — the active track

**Objective:** get fluent in the recurring patterns behind most coding-interview problems, in
**Go**, so that an unseen problem decomposes into "oh, this is pattern X" instead of a blank page.
Source: [15 Coding Interview Patterns](https://dev.to/somadevtoo/coding-interviews-was-hard-until-i-learned-these-patterns-2ji7).

This file is the **syllabus** the `/daily` command walks. `PROGRESS.md` is still the day-to-day
memory (results, confidence, spaced repetition); this file is the fixed map of *where we're going*.
If a session falls apart, this is the doc to re-read.

## How the laddering works

Each pattern is broken into **micro-steps** before the real problem. We never open with the
LeetCode-style problem cold. The sequence for every pattern:

1. **Primitive drills** — the mechanical sub-skills the pattern is made of (e.g. "move a write
   index", "pop while stack top is smaller"), each a ≤15-min exercise. Multiple days if needed.
2. **The pattern, assembled** — combine the primitives on a clean, minimal problem.
3. **A representative problem** — a real interview-style question that the pattern unlocks.

Rules carried over from the daily system: **≤15 min per exercise**, no solutions or hints in the
folder (just SPEC + failing tests), one pattern may span several days. A step is only "done" when
its tests are green; `/daily continue` advances to the next step, `/daily` (empty) may resurface an
earlier step that's due for review.

Difficulty calibration: if a step passes easily, skip ahead or fold two primitives into one day; if
a step fails, split it smaller and isolate the sticking point (this is exactly how the slice
windowing recovery drill was built).

## Go prerequisites (introduced just-in-time, not all up front)

These get their own primitive-drill days the first time a pattern needs them:

- **Slice as stack / queue** — `append` to push; `s[len(s)-1]` + reslice to pop; `s[0]` + `s[1:]`
  to dequeue. (Needed by: Monotonic Stack, BFS, Matrix Traversal.)
- **`container/heap`** — implementing `heap.Interface` (`Len/Less/Swap/Push/Pop`) for a min/max
  heap. (Needed by: Top K.)
- **`sort.Slice`** — sorting structs by a field with a less-func. (Needed by: Intervals.)
- **Linked-list node** — `type ListNode struct { Val int; Next *ListNode }`, building and walking
  one. (Needed by: Fast & Slow Pointers, In-Place Reversal.)
- **Tree node** — `type TreeNode struct { Val int; Left, Right *TreeNode }`, recursion over it.
  (Needed by: Tree Traversal, DFS, BFS.)
- **`map` as a set / counter** — already comfortable (covered in earlier daily work).

## The 15 patterns, in order

Status legend: ☐ not started · ◐ in progress · ☑ solid. Update as we go.

### 1. Two Pointers ☑ (converge + write-index solid; 2c skipped by user call)
Build on the slice index work already done.
- [x] **2a-i** Ends mechanic, no decisions — `IsPalindrome`: `l`/`r` from both ends, `for l < r`,
  compare and move *both* inward. → `exercises/2026-06-30-go-two-pointers-ends/` (passed)
- [x] **2a-ii** Same mechanic + a swap — `ReverseInPlace`: ends, swap, move both inward.
  → `exercises/2026-06-30-go-two-pointers-ends/` (passed)
- [x] **2a-iii** The converge: move *only one* pointer based on `sum < / > / == target` on a sorted
  slice. (`PairWithTarget`, return the index pair.) This is the one genuinely new idea.
  → `exercises/2026-07-01-go-two-pointers-converge/` (passed)
- [x] **2b** Same-direction write index: `KeepPositives`/`RemoveElement`/`MoveZeroes`/`RemoveDuplicates`
  — the "slow write pointer, fast read pointer" idiom, both conventions (next-open-slot and
  last-kept-index). → `exercises/2026-07-02-go-two-pointers-write-index/` (passed)
- [ ] **2c** Problem: valid palindrome (skip non-alphanumerics) and/or two-sum-II. **Skipped for now**
  — user is confident in the primitives (converge + write-index both landed clean) and asked to move
  on to the next pattern. Revisit 2c later as a review rung if it comes up due for spaced repetition.

### 2. Prefix Sum ◐ (primitives solid; 3c deferred by user call — revisit as review rung)
Build on the slice index work already done.
- [x] **3a** Build a prefix-sum slice from an input (`pre[i] = pre[i-1] + a[i-1]`, length n+1).
  → `exercises/2026-07-07-go-prefix-sum/` (passed)
- [x] **3b** Range-sum query: answer `sum(i..j)` in O(1) from the prefix slice.
  → `exercises/2026-07-07-go-prefix-sum/` (passed)
- [ ] **3c-i** Existence check: does any contiguous subarray sum to `target`? (running sum + a *set*
  of seen prefix sums — the "have I seen `running - target` before" trick, boolean only.)
  **Deferred** — user hit the conceptual wall on the running-sum + seen-set trick across multiple
  sessions and chose to move on rather than keep stalling. The mechanical foundation (3a/3b) is
  solid; this is the actual algorithmic insight that didn't click. Revisit later as a review rung
  if it comes up due for spaced repetition or once Sliding Window / similar running-state patterns
  reinforce the "maintain a state and check membership against it" reflex.
- [ ] **3c-ii** Problem: subarray-sum-equals-k (leetcode 560) — same trick but *counting*, so the set
  becomes a frequency map and every match adds `counts[running-target]` instead of stopping at the
  first hit. This is the real interview problem the primitives above unlock.
  **Deferred** alongside 3c-i — same conceptual sticking point. Revisit with 3c-i.

### 3. Sliding Window ◐
- [x] **4a** Fixed window: max sum of a size-k subarray by add-right / subtract-left (no recompute).
  → `exercises/2026-08-12-go-sliding-window-fixed/` (passed — **caveat: solved O(n·k) recompute, not
  the O(n) slide; the slide mechanic did not land. Review sooner than the interval if 4b doesn't
  heal it.**) **Slide re-drill on 2026-08-15** (`exercises/2026-08-15-go-sliding-window-slide/`) — user
  asked to hone the missing point from the critique directly: isolate the slide, re-apply to 4a done
  right, then introduce variable-width expand/shrink with a running sum only. **Resolved 2026-08-28:
  that re-drill went green with the real O(1) slide present — the recompute habit is healed.**
- [x] **4b-pre** Variable-window sum primitive: expand right / shrink left while sum ≥ target, track
  min length (positives only). The expand/shrink mechanic with a running sum — no seen-set — so it
  drills the grow/shrink reflex without re-triggering the 3c seen-set wall.
  → `exercises/2026-08-15-go-sliding-window-slide/` (passed — expand/shrink landed)
- [ ] **4b** Variable window: longest substring without repeating chars — expand right, shrink left
  on violation, track best. **Needs a seen-set** (the deferred 3c sticking point); attempt only after
  4b-pre's expand/shrink is solid.
- [ ] **4c** Problem: minimum window substring (need-counts + window-counts).

### 4. Fast & Slow Pointers ☐
- [ ] **5a** Prereq: build & traverse a singly linked list (`ListNode`).
- [ ] **5b** Floyd cycle detection: does the list have a cycle?
- [ ] **5c** Problem: find the cycle's start node, or happy-number.

### 5. LinkedList In-Place Reversal ☐
- [ ] **6a** Iterative full reversal: the `prev`/`curr`/`next` three-pointer dance.
- [ ] **6b** Reverse a sublist between positions m and n.
- [ ] **6c** Problem: swap nodes in pairs.

### 6. Monotonic Stack ☐
- [ ] **7a** Prereq: slice-as-stack push/pop; peek the top safely.
- [ ] **7b** Next-greater-element: maintain a decreasing stack of indices, pop while smaller.
- [ ] **7c** Problem: daily temperatures.

### 7. Top 'K' Elements ☐
- [ ] **8a** Prereq: implement `container/heap` for a min-heap of ints.
- [ ] **8b** Kth largest via a size-k min-heap (push, pop when over k).
- [ ] **8c** Problem: top-k-frequent (count map + heap).

### 8. Overlapping Intervals ☐
- [ ] **9a** Prereq: `sort.Slice` intervals by start.
- [ ] **9b** Merge overlapping intervals after sorting.
- [ ] **9c** Problem: insert-interval, or meeting-rooms-II (min rooms).

### 9. Modified Binary Search ☐
- [ ] **10a** Vanilla binary search with airtight bounds (`lo`, `hi`, `mid`, no off-by-one).
- [ ] **10b** Leftmost / rightmost insertion point (lower/upper bound).
- [ ] **10c** Problem: search in rotated sorted array, or find-peak-element.

### 10. Binary Tree Traversal ☐
- [ ] **11a** Prereq: `TreeNode`; recursive in-order traversal into a slice.
- [ ] **11b** Pre-order and post-order; then iterative in-order with an explicit stack.
- [ ] **11c** Problem: max depth and validate-BST.

### 11. Depth-First Search ☐
- [ ] **12a** DFS template on a tree with an accumulator (root-to-leaf path collection).
- [ ] **12b** Path-sum (does any root-to-leaf path equal target?).
- [ ] **12c** Problem: number-of-islands (DFS flood on a grid) — bridges into Matrix Traversal.

### 12. Breadth-First Search ☐
- [ ] **13a** Prereq: slice-as-queue; level-order traversal flat into one slice.
- [ ] **13b** Level-by-level using a size snapshot per level (`[][]int`).
- [ ] **13c** Problem: min-depth or shortest-path-in-binary-matrix.

### 13. Matrix Traversal ☐
- [ ] **14a** Iterate a 2D grid; enumerate the 4 neighbors with bounds checks.
- [ ] **14b** Flood fill (DFS or BFS over a grid).
- [ ] **14c** Problem: rotting-oranges or surrounded-regions.

### 14. Backtracking ☐
- [ ] **15a** The choose / recurse / un-choose template: all subsets of a set.
- [ ] **15b** Permutations of distinct elements.
- [ ] **15c** Problem: combination-sum.

### 15. Dynamic Programming ☐
- [ ] **16a** 1D memoization → tabulation: climbing-stairs.
- [ ] **16b** 1D state choice: house-robber or coin-change.
- [ ] **16c** Problem: longest-increasing-subsequence; first taste of 2D (grid unique-paths).

## Notes
- Patterns 1–3 (two pointers, prefix sum, sliding window) are pure slice/array work and continue
  directly from the `Take/Drop/Chunk` windowing drill — start here.
- Patterns 4–5 need linked lists; 10–12 need trees; 7 needs heaps — those Go prereqs become their
  own `?a` micro-step days.
- This list is the plan, not a contract. Reorder or insert remedial steps freely based on how the
  daily attempts go — that calibration is the whole point.
</content>
</invoke>
