# Sliding Window — the O(1) slide, honed (pattern 3 — 4a refresher + 4b-pre)

**Why this day exists.** The 4a pass got the right *answers* but bypassed the *pattern*: each window's
sum was recomputed with `sum(a[i:i+k])` inside the loop → **O(n·k)**, and the running-sum +
running-best state never appeared. The entire reason Sliding Window exists as a named pattern is the
**one-add / one-subtract slide** — seed the first window's sum, then for each step right: `sum += a[i]
- a[i-k]`. Once that reflex is in, the window's sum is O(1) and the decision layered on top (max, min,
where, longest, shortest) is a separate, tiny piece of state. Today installs that reflex, re-applies
it to the 4a job done right, then extends it to the **variable-width** shape (expand right / shrink
left) — the bridge into 4b.

No solution or hints beyond naming the mechanic. That's yours to write.

## Tasks (do them in order — each one builds on the previous)

In `slidingwindow.go`:

1. **`WindowSums(a []int, k int) []int`** — return the sums of every length-`k` window,
   left to right. This is the slide in pure form: the *only* thing to track is one running `sum`
   you mutate by adding the new right element and subtracting the old left one. Result length is
   `len(a)-k+1`. If `k <= 0` or `k > len(a)`, return `nil`.
2. **`MaxSumKAndStart(a []int, k int) (int, int)`** — the 4a job, done right in **one pass**: a
   single running `sum` (seeded then slid), a running `best`, and a running `bestIdx` carried
   alongside. Return `(maxSum, earliestStart)`. On ties the earliest start wins. If `k <= 0` or
   `k > len(a)`, return `(0, -1)`. No `sum()` helper, no per-window recompute.
3. **`MinLenSumAtLeast(a []int, target int) int`** — the variable-window primitive. Expand the
   right end, adding to a running `sum`; whenever `sum >= target`, try to **shrink from the left**
   (subtract the left element, advance left) to find a shorter window that still meets target, and
   track the minimum length seen. You may **assume all elements are positive**, so growing always
   raises the sum and shrinking always lowers it — that monotonicity is what makes the two-pointer
   expand/shrink correct. Return the minimum length, or `0` if no subarray reaches `target`
   (`target <= 0` ⇒ `0`).

## Acceptance criteria

`./daily/run.sh` is green. Tests cover: ordinary mixed positives/negatives, all-negatives (the
seed must come from a real window, not `0`), `k == 1`, `k == len(a)`, ties (earliest wins), invalid
`k` sentinels, and — for `WindowSums` — a **large input** (`n = 100000`, `k = 50000`). That large
test exists on purpose: an O(n) seed+slide finishes in microseconds; an O(n·k) recompute loop takes
several seconds and visibly hangs. If the run hangs there, the 4a recompute habit is back — use the
slide.

## The thing to get right

One running `sum` variable that you **mutate in place**, never a fresh sum over a slice. Tasks 1–2
only slide it (fixed width); task 3 grows and shrinks it (variable width). The decision state (`best`,
`bestIdx`, `minLen`) is a separate, tiny thing you carry next to `sum` — that separation is the
pattern.

## Larger arc

Pattern **3 (Sliding Window)** in `INTERVIEW_PATTERNS.md`. Today heals the 4a slide gap and installs
the expand/shrink primitive that **4b** (longest substring without repeating chars) builds on —
except 4b swaps the running sum for a *seen-set*, which is a separate, previously-deferred sticking
point (3c), so today sums only. The seen-set day comes after this mechanic is solid.