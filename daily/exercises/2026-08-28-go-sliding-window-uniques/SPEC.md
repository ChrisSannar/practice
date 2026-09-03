# Sliding Window — the seen-set window (pattern 3 — rung 4b)

**Why this day exists.** The slide (O(1) fixed window) and the expand/shrink (variable window with a
running sum) are now solid — the 2026-08-15 re-drill went green with the real slide present. The next
rung, **4b (longest substring without repeating characters)**, needs one new ingredient the running
sum didn't: a **seen-set**. You expand the window right, but the "is this legal?" check is now
*membership* ("have I already got this element in the window?"), and the shrink condition is "keep
dropping from the left until the duplicate is gone." That membership reflex is the same idea that
stalled at 3c, so today ladders it in gently — a pure set primitive first, then the set driving a
window over ints, then the real string problem — instead of opening cold on LeetCode 3.

No solution or hints beyond naming the mechanic. That's yours to write.

## Tasks (do them in order — each one builds on the previous)

In `uniques.go`:

1. **`FirstRepeatedIndex(a []int) int`** — the pure membership primitive, no window, no arithmetic.
   Walk left to right keeping a set of values seen so far; return the index of the first element
   that's already in the set, or `-1` if all distinct. Just: maintain a set, check membership.
2. **`LongestUniqueLen(a []int) int`** — the set now *drives a window*. Grow the window from the
   right, adding each element to a set of the window's current values. When the incoming element is
   already in the set, **shrink from the left** — remove `a[left]` from the set and advance `left` —
   repeating until the duplicate is cleared, then add the incoming element. Track the largest window
   length. Combines the (now solid) expand/shrink with task 1's membership. `[]int{}` ⇒ `0`.
3. **`LengthOfLongestSubstring(s string) int`** — the real 4b problem (LeetCode 3): identical
   mechanic, but the window holds the **runes** of `s`. Iterate runes (not bytes) so multibyte
   characters count as one. `""` ⇒ `0`.

## Acceptance criteria

`./daily/run.sh 2026-08-28` is green. Tests cover: empty and single inputs, all-same (answer 1),
all-distinct (answer = length), an interior repeat that forces the window to shrink past several
elements, best-window-early vs best-window-late, and — for the string case — multibyte unicode (so a
byte-based solution is caught).

## The thing to get right

The **shrink loop**. On a duplicate you don't reset the window or step left once — you keep removing
from the left *until the specific duplicate is gone*, and only then extend right. The set is the
window's contents; `left` and `right` are its bounds; the max length is a separate tiny piece of
state carried alongside — same "state + membership check" separation as the running-sum windows, with
a set instead of a sum.

## Larger arc

Pattern **3 (Sliding Window)** in `INTERVIEW_PATTERNS.md`. This is rung **4b**; it installs the
seen-set window that the final rung **4c (minimum window substring)** builds on — 4c swaps the plain
set for *need-counts vs window-counts*, the last escalation of this pattern.
