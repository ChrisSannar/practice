# Two Pointers — the write index (interview track, step 2b)

**Pattern:** Two Pointers, a new *shape*. So far both pointers started at opposite ends and converged.
Now **both start at the left**: a slow **write** pointer trails a fast **read** pointer. This is the
idiom behind in-place array edits — you overwrite the front of the slice as you scan it.

## Tasks (~20 min, 4 of them — do them in order)

In `writeindex.go` — all four are the same idiom, growing in difficulty. **Do task 0 first**: it
isolates the bare mechanic so tasks 1–3 are just new keep-rules on top of it.

0. `KeepPositives(s []int) int` — **the warm-up.** Compact `s` in place so the first `k` elements
   are the positive values (`> 0`) in order; return `k`. The keep-rule looks only at the *current*
   element. Get this and the rest are variations.
1. `RemoveElement(s []int, val int) int` — remove every occurrence of `val` in place; the first `k`
   elements are the survivors (order preserved), return `k`. Same as the warm-up, keep-rule is
   `!= val`. `s` is *not* sorted.
2. `MoveZeroes(s []int)` — move every `0` to the end, keeping the relative order of the non-zero
   values. Mutate in place, return nothing. (Pack the non-zeroes forward, then fill the tail with 0.)
3. `RemoveDuplicates(s []int) int` — **the hard one.** `s` is sorted ascending. Compact it in place
   so the first `k` elements are the unique values in order; return `k`. Here the keep-rule depends
   on the *previous kept* value (`s[read] != s[write-1]`), not just the current element.

**Constraint:** one pass, in place — no second slice, no map. The read pointer visits every element;
the write pointer only advances when you decide to keep a value. In `KeepPositives`/`RemoveElement`
the two pointers are unrelated — resist comparing `s[slow]` to `s[fast]`; that's a different pattern.

## Acceptance criteria

`./daily/run.sh` is green. Tests cover: mixed / zero / all / none / empty / single (KeepPositives);
some-match / none / all / empty / single (RemoveElement); interleaved / leading / no / all zeroes,
empty, trailing (MoveZeroes); dups at various densities, already-unique, empty, single, all-same
(RemoveDuplicates).

## The thing to get right

The mental model: **read** asks "should this element be kept?"; **write** is "where the next kept
element goes." They advance independently — that's the whole trick, and it's a different rhythm from
the ends-converge loop. For `MoveZeroes`, once every non-zero is packed to the front, the tail slots
from `write` to the end get filled with `0`.

## Larger arc

Step **2b** in `INTERVIEW_PATTERNS.md`. Next: **2c** a real problem that composes what you've built
(valid palindrome / two-sum-II). That closes out Two Pointers and we move to **Prefix Sum**. No
solution or hints here — that's yours to write.
