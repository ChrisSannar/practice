# Go basics — slices & maps

**Goal:** level up from standalone helpers to the slice/map manipulation you reach for
constantly: building new slices, using maps as sets, and merging maps.

**Today's slice (~15–20 min):** implement the four stubs in `slicesmaps.go`:

| Function | What it exercises |
|----------|-------------------|
| `Chunk(s []int, size int) [][]int` | building a slice-of-slices, slicing windows |
| `Unique(s []int) []int` | map-as-set while preserving first-seen order |
| `Intersection(a, b []int) []int` | set membership, ordering from `a`, dedup |
| `Merge(maps ...map[string]int) map[string]int` | variadic params, summing shared keys |

**Contracts to honor (the tests enforce these):**
- `Chunk`: the last chunk may be shorter; `size <= 0` returns an empty `[][]int{}`; an
  empty input returns `[][]int{}` (not nil).
- `Unique` / `Intersection`: results keep the order values first appear in the *first* slice,
  with duplicates removed; empty/no-overlap cases return an empty (non-nil) slice.
- `Merge`: with no maps, return an empty map; shared keys have their values **summed**.

**Acceptance criteria:** `./daily/run.sh` is green — every case passes.

Heads-up: several cases check **empty (non-nil)** results via `reflect.DeepEqual`, so a `nil`
return won't pass where an empty `[]int{}` / `[][]int{}` / `map[string]int{}` is expected.

**Larger arc:** Go fundamentals refresher → (today) slices & maps → next rungs: error handling
& multiple returns, then structs & methods, then goroutines/channels. `/daily continue` to climb.
