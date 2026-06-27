# Critique — 2026-06-27 Go slices & maps

**Result:** `Unique`, `Intersection`, `Merge` green; `Chunk` unsolved (returns `nil`).
3 of 4 with genuinely good instincts on the map work. `Chunk` is a single, fixable
mental-model bug — analyzed in depth below since you asked.

## What was strong

- **`Unique`** — textbook map-as-set: read `mappers[val]`, append only when unseen, mark seen.
  Order preserved, empty handled. Nothing to change.
  - Micro-note: `map[int]bool` is fine; `map[int]struct{}` is the zero-byte-value variant some
    Go folks prefer for pure sets (`if _, seen := m[val]; !seen { ... }`). Purely stylistic.
- **`Intersection`** — you reached for the right *performance* shape on your own: build a set
  from one slice, scan the other → O(n+m) instead of O(n·m). That instinct is the valuable part.
- **`Merge`** — correct and clean. `THE_MAP[key] = THE_MAP[key] + val` leans on the same
  zero-value read from your last critique. Good.

## Two things that passed but are worth knowing

1. **`Merge` is not O(n²).** Your comment worries about that, but the nested loops walk *each
   key-value pair exactly once* across all maps — total work is linear in the number of entries.
   That's optimal; there's no faster way to read every pair. Nested loops ≠ quadratic when the
   inner loop is over a *different* collection each time.

2. **`Intersection` passes by luck of the test data.** Two latent deviations from the spec:
   - The spec says *"order they first appear in a."* You scan the **smaller** slice and then
     `sort.Ints(result)`. The tests all expect ascending output, so sorting happens to match —
     but if a case wanted `a = [3,1,2] → [3,1,2]` order, you'd fail. Drop the sort and iterate
     **`a`** (not "smaller"), checking membership in a set built from `b`.
   - Dedup depends on the scanned slice having no duplicates. If the slice you scan had repeats,
     you'd emit repeats. Iterating `a` with a "seen" set (or the set-from-`b` doubling as it)
     closes that hole.

## `Chunk` — where the model slipped (your real ask)

Your commented-out attempt:
```go
result := make([][]int, size)   // (1)
chunkSize := len(s) / size      // (2)
numberOfChunks := len(s) / chunkSize
for idx := range chunkSize / len(s) {  // (3)
```

Three compounding confusions:

1. **`size` already *is* the chunk size** — it's the count of elements per group, handed to you.
   `make([][]int, size)` preallocates the *outer* slice to `size` slots, but the outer length
   should be the *number of chunks*, not the elements-per-chunk. Mixing those two is the root bug.
2. **`chunkSize := len(s) / size`** then re-derives something you were already given, and names
   the *number of chunks* "chunkSize". The naming inversion is what tangled the rest.
3. **`for idx := range chunkSize / len(s)`** — integer division of a small number by a larger one
   is `0`, so the loop never runs and you return empty. That's the mechanical reason it produced nothing.

**The approach that dissolves all of this:** don't pre-size anything. Start with an empty result
and *walk the input in steps of `size`*, slicing one window each step. The only trick is clamping
the last window so it doesn't run past the end. Illustrated on a different shape so it's a pattern,
not today's answer:

```go
// chunking [10,20,30,40,50] by 2  -> windows at i = 0,2,4
out := [][]int{}
for i := 0; i < len(xs); i += size {   // step by size, not by 1
    end := i + size
    if end > len(xs) {                 // clamp the ragged final window
        end = len(xs)
    }
    out = append(out, xs[i:end])       // s[i:end] is the window
}
```

Two things to internalize:
- **`s[i:end]` is a slice expression** — it gives you elements `i` up to (not including) `end`.
  This is the everyday Go tool for "give me a window of a slice."
- **Guard `size <= 0` first** and return the empty `[][]int{}` — that also sidesteps the
  divide-by-zero / `make` panic you'd otherwise hit.

## Verdict

Confidence: **3/5** for slices & maps. Your set/perf instincts are ahead of where I'd expect —
the gap is specifically **slice index/window arithmetic** (`s[i:end]`, stepping by `size`,
clamping the tail). Today's exercise isolates exactly that and builds back up to `Chunk`.
