# Two Pointers — the write index (interview track, step 2b)

**Pattern:** Two Pointers, a new *shape*. So far both pointers started at opposite ends and converged.
Now **both start at the left**: a slow **write** pointer trails a fast **read** pointer. This is the
idiom behind in-place array edits — you overwrite the front of the slice as you scan it.

## Tasks (~20 min, 3 of them)

In `writeindex.go` — all three are the same idiom, growing in freedom:

1. `RemoveDuplicates(s []int) int` — `s` is sorted ascending. Compact it in place so the first `k`
   elements are the unique values in order; return `k`. Elements past index `k` don't matter.
2. `MoveZeroes(s []int)` — move every `0` to the end, keeping the relative order of the non-zero
   values. Mutate in place, return nothing.
3. `RemoveElement(s []int, val int) int` — remove every occurrence of `val` in place; the first `k`
   elements are the survivors (order preserved), return `k`. `s` is *not* sorted.

**Constraint:** one pass, in place — no second slice, no map. The read pointer visits every element;
the write pointer only advances when you decide to keep a value.

## Acceptance criteria

`./daily/run.sh` is green. Tests cover: dups at various densities, already-unique, empty, single,
all-same (RemoveDuplicates); interleaved / leading / no / all zeroes, empty, trailing (MoveZeroes);
and some-match / none / all / empty / single (RemoveElement).

## The thing to get right

The mental model: **read** asks "should this element be kept?"; **write** is "where the next kept
element goes." They advance independently — that's the whole trick, and it's a different rhythm from
the ends-converge loop. For `MoveZeroes`, once every non-zero is packed to the front, the tail slots
from `write` to the end get filled with `0`.

## Larger arc

Step **2b** in `INTERVIEW_PATTERNS.md`. Next: **2c** a real problem that composes what you've built
(valid palindrome / two-sum-II). That closes out Two Pointers and we move to **Prefix Sum**. No
solution or hints here — that's yours to write.
