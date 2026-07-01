# Two Pointers — the converge (interview track, step 2a-iii)

**Pattern:** Two Pointers. **Today's new idea:** the *converge loop* — same `for l < r` frame you
already nailed, but now you move **only one** pointer each step, chosen by a comparison. This is the
move behind most sorted-array two-pointer problems.

## Task (~15 min)

Implement `PairWithTarget(s []int, target int) []int` in `twopointers.go`:

- `s` is sorted **ascending**.
- Return `{i, j}` with `i < j` — the indices of a pair whose values sum to `target`.
- If no pair sums to `target`, return an empty (non-nil) `[]int{}`.

**Constraint:** one pointer at each end, converging. **No nested loop, no map.** The point is the
single-pointer-per-step decision.

## Acceptance criteria

`./daily/run.sh` is green. Tests cover a pair at the front, the two ends, pairs needing convergence
from both sides, an interior pair, no-pair cases, empty and single-element slices, and negatives.

## The thing to get right

The body is a three-way decision on `s[l] + s[r]` versus `target`:
- too small → the only way to grow the sum is to raise the low end,
- too big → the only way to shrink it is to lower the high end,
- equal → you're done, return the pair.
Because `s` is sorted, each move can't skip the answer. Empty result must be `[]int{}`, not `nil`
(tests use `reflect.DeepEqual`).

## Larger arc

Final rung of **2a** in `INTERVIEW_PATTERNS.md`. Next: **2b** the same-direction write-index idiom
(in-place remove-duplicates / move-zeroes), then **2c** a real problem (valid palindrome). No
solution or hints here — that's yours to write.
