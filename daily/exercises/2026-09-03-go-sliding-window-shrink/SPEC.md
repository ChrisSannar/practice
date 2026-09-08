# Sliding Window — fixing the shrink loop (pattern 3 — rung 4b, re-drill)

**Why this day exists.** 2026-08-28's `LongestUniqueLen` / `LengthOfLongestSubstring` passed their
tests, but on review the shrink step wasn't actually there: on a duplicate, that attempt deleted
only the *incoming* value and jumped `left` straight to the current index — silently dropping every
element between the old `left` and the jump target, without ever checking whether *they* were still
legitimately inside the window. That undercounts whenever the earlier occurrence of the duplicate
isn't sitting right at `left`. Today isolates exactly that step before re-attempting the two driving
functions.

No solution or hints beyond naming the mechanic. That's yours to write.

## Tasks (do them in order — each one builds on the previous)

In `shrink.go`:

1. **`ShrinkPastDuplicate(a []int, left, right int) int`** — `a[left:right]` (i.e. indices
   `left` through `right-1`) is a window with **no duplicates in it**. `a[right]` is a candidate to
   add next, and it *might* duplicate some single element already in that window. Return the new
   `left` after removing `a[left]`, then `a[left+1]`, and so on — one index at a time, advancing
   `left` by exactly one per removal — stopping the instant the element that duplicated `a[right]`
   has been removed. If nothing in `a[left:right]` equals `a[right]`, return `left` unchanged. This
   is the isolated primitive — no set, no driving loop, just "how far does `left` have to walk."
   - `ShrinkPastDuplicate([]int{1, 2, 5, 9, 5}, 0, 4) -> 3` (the duplicate `5` is at index 2, so
     indices 0, 1, 2 all have to go before `left` can stop)
   - `ShrinkPastDuplicate([]int{5, 1, 2, 9, 5}, 0, 4) -> 1` (the duplicate is at index 0 — one
     removal clears it)
   - `ShrinkPastDuplicate([]int{1, 2, 3, 4, 5}, 0, 4) -> 0` (no duplicate of `a[4]`; `left` doesn't
     move)
2. **`LongestUniqueLen(a []int) int`** — same contract as 2026-08-28: grow a window from the right
   into a set of its current values; on a duplicate, shrink from the left **one element at a time**
   (this is task 1's mechanic, now driving the actual loop) until it clears, then add the incoming
   element; track the largest window length. `[]int{}` ⇒ `0`.
   - `LongestUniqueLen([]int{0, 1, 2, 0, 3}) -> 4` (the earlier `0` is *not* adjacent to the new
     window's start — a jump-based shrink undercounts this to 3; the correct one-step-at-a-time
     shrink gets 4, from `[1, 2, 0, 3]`)
3. **`LengthOfLongestSubstring(s string) int`** — the real 4b problem (LeetCode 3), identical
   mechanic over the runes of `s`. `""` ⇒ `0`.
   - `LengthOfLongestSubstring("dvdf") -> 3` (the classic case: the repeated `d` is not adjacent to
     the window's current left edge — `"vdf"` is the answer, not `"vd"`)

## Acceptance criteria

`./daily/run.sh 2026-09-03` is green. Tests specifically include non-adjacent-repeat cases (the
`dvdf` shape, and its int-array equivalent) alongside the usual empty/single/all-same/all-distinct/
early-vs-late-best/unicode-multibyte coverage from 2026-08-28.

## The thing to get right

`left` moves **one index at a time**, and every element it passes over must actually be removed
from the set — never jump `left` straight to a target index without walking (and evicting) every
position in between. The set's contents and `[left, right]` must always agree with each other.

## Larger arc

Pattern **3 (Sliding Window)** in `INTERVIEW_PATTERNS.md`. This is still rung **4b** — it isn't
marked done in the syllabus until this shrink loop is airtight — after which **4c** (minimum window
substring) escalates the same shrink loop to need-counts vs window-counts.
