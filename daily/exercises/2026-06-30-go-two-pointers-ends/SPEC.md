# Two Pointers — the ends mechanic (interview track, steps 2a-i & 2a-ii)

**Pattern:** Two Pointers. **Today's primitive:** just the *mechanic* — two pointers, one at each
end of a slice, walking toward the middle while `l < r`. No decisions about which one to move yet
(both move every step). Two tiny functions that share the exact same loop skeleton.

## Task (~15 min)

In `twopointers.go`:

1. `IsPalindrome(s []int) bool` — does `s` read the same forwards and backwards? Compare the two
   ends; on a mismatch it's not a palindrome. Empty and single-element slices are palindromes.
2. `ReverseInPlace(s []int)` — reverse `s` in place (mutate the caller's slice, return nothing).
   Swap the two ends, then step inward.

Both use the same shape: `l := 0`, `r := len(s)-1`, loop `for l < r`, do the work on `s[l]`/`s[r]`,
then `l++; r--`. Write one and the other is muscle memory.

## Acceptance criteria

`./daily/run.sh` is green. Tests cover odd/even lengths, empty, single, and two-element slices.

## The thing to get right

The whole exercise is the loop boundary: **`for l < r`** (not `<=` — when they meet or cross,
you're done). That's the same loop-termination instinct from `Chunk`, in its simplest possible form.

## Larger arc

Steps **2a-i / 2a-ii** of Two Pointers in `INTERVIEW_PATTERNS.md`. These establish the ends-converge
mechanic. Next (**2a-iii**) adds the one new idea: moving *only one* pointer based on a comparison
(`PairWithTarget` on a sorted slice). Then 2b (same-direction write index) and 2c (a real problem).
No solution or hints here — that's yours to write.
