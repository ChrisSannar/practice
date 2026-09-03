# Critique — 2026-08-12 — Sliding Window / fixed (Go)

Verdict: **passed** (all three test functions green). Confidence: **3/5** — tests pass, but the core
reflex the pattern is meant to install did not land yet. See below.

## What's right

- All three functions correct on the cases provided, including the edge cases that catch the
  common traps: all-negatives (`MaxSumK([-4,-2,-3,-1], 2) = -4`, not 0 — the seed must not start at
  0), `k == 1`, `k == len(a)`, and the earliest-on-tie rule for `MaxSumKStart`.
- Sentinel handling is consistent and correct: `0` for the value-returning functions, `-1` for the
  index-returning one, all guarded up front with the same `k <= 0 || k > len(a)` check.
- `MaxSumKStart` correctly uses strict `>` for the comparison so the earliest window wins on ties.
  A `>=` there would have silently returned the *latest* tie — easy bug to make, you didn't.
- The shape is consistent across all three — the same seed-and-scan loop, just a different thing
  tracked — which is the right mental model: the slide is mechanic, the decision is separate.

## Where to sharpen — the big one

**This is not the sliding-window pattern.** The whole point of step 4a — the reflex this pattern
is supposed to install — is the **O(1) slide**: seed the first window's sum, then for each step
right add the new element and subtract the one falling off the left, so each window's sum is one
add and one subtract instead of a fresh sum over `k` elements.

Your solution calls `sum(a[i : i+k])` inside the loop, recomputing the window from scratch every
time. That's **O(n·k)**, not O(n). For the small test inputs it is invisible; on an interview
problem with n = 10^5 and k = 10^4 it blows up where the O(n) slide would not. Worse, it sidesteps
the actual skill: the running-sum + running-best state is the thing you're meant tointernalize,
and it never appears here.

The shape the spec calls out, made concrete on `MaxSumK`:

```go
func MaxSumK(a []int, k int) int {
    if k <= 0 || k > len(a) {
        return 0
    }
    sum := 0
    for i := 0; i < k; i++ { sum += a[i] }   // seed
    best := sum
    for i := k; i < len(a); i++ {           // slide
        sum += a[i] - a[i-k] 
        if sum > best { best = sum }
    }
    return best
}
```

Notice: no `math.MaxInt` sentinel needed — seeding `best := sum` from the first window handles
the all-negatives case for free, because the true max is *some* real window. The same shape works
for `MinSumK` with the comparison flipped, and for `MaxSumKStart` by carrying `idx` alongside
`best` and updating both together. `MinSumK` and `MaxSumKStart` currently have the same recompute
shape — each one is a chance to drill the slide, and each one currently misses it.

## Smaller notes

- `math.MaxInt * -1` works but `math.MinInt` is the direct constant. Seeding from the first window
  (above) sidesteps the choice entirely and is the more idiomatic pattern-shaped answer.
- Loop condition `i+k < len(a)+1` is correct but reads awkwardly. Idiomatic forms: `i+k <= len(a)`
  or `i <= len(a)-k`. Under the slide shape the loop becomes `for i := k; i < len(a); i++` which
  needs no arithmetic at all.
- The shared `sum` helper is fine but disappears under the slide approach — it's only needed
  because of the recompute.

## How the next task should adjust

This is a pass on test correctness but **not a pass on the pattern**. The O(n·k) recompute solution
would not survive an interview's follow-up "what's the time complexity?" question. Two honest options:

1. **Self-reset before the interval elapses.** Edit the Prefix Sum row in `PROGRESS.md` to confidence
   2 and `Next review` to tomorrow (2026-08-13), so `/daily` empty-mode resurfaces this pattern
   immediately and you re-attempt the slide specifically. The tests already in place are good — the
   slide passes them just as the brute force did.
2. **Fold the slide into 4b.** Variable window (next session) uses the same running-sum reflex but
   with a grow/shrink decision. If you implement 4b with a proper running sum, the fixed-window slide
   will click as a special case, and 4a's weakness heals itself. Lower risk, but it depends on 4b
   landing — if 4b *also* gets brute-forced, the hole compounds.

Default: option 1. The whole point of this pattern is the one-line slide; skipping it once is fine,
skipping it twice is a habit.

## Test fix you made

`TestMaxSumKStart` case "max window starts at index 3": original spec had `want 3`, you corrected to
`want 2`. You were right. `a = [2, 1, 5, 1, 3, 2], k = 3`: window sums are 8 / 7 / **9** / 6, max
at index 2. The original `3` was an arithmetic error in the spec — good catch, and thanks for fixing
the test rather than working backwards to match the wrong expected value.
