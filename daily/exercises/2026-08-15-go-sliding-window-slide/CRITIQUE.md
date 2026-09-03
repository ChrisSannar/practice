# Critique — 2026-08-15 Sliding Window (slide re-drill + 4b-pre)

**Verdict: PASS — confidence 4.** All four tests green, and this time the *pattern* is actually here,
not just the answers. This is the win the whole day was scheduled for.

## What landed (the point of the day)

The O(1) slide reflex that 4a bypassed is now in your hands:

```go
sum := 0
for i := range k { sum += a[i] }      // seed the first window once
sums = append(sums, sum)
for i := k; i < len(a); i++ {
    sum += a[i] - a[i-k]              // one add, one subtract — O(1) per step
    sums = append(sums, sum)
}
```

That `sum += a[i] - a[i-k]` is exactly what was missing on 2026-08-12, when every window was
recomputed with `sum(a[i:i+k])` (O(n·k)). The large forcing test (n=100000, k=50000) passing in
milliseconds is the proof the recompute habit is gone.

In `MaxSumKAndStart` you also kept the **decision state separate** from the sum — `maxSum`,
`maxSumIdx` carried alongside the running `sum`, updated only on a strict `>` so the earliest tie
wins. That separation (sum is one thing; best/where is another) *is* the sliding-window pattern.
Nicely done.

`MinLenSumAtLeast` is a correct variable-width window: expand right into `runningSum`, and
`for runningSum >= target { shrink left }` — the right shape, and you leaned on the positives-only
monotonicity exactly as intended.

## Sharpen (all style, no bugs)

- **Dead code:** `PrintArr` is a leftover debug scaffold — delete it before it rides along into the
  next file.
- `make([]int, 0)` in `WindowSums` works, but `var sums []int` is the idiomatic nil-friendly form
  (and matches your `return nil` guard).
- In `MinLenSumAtLeast` you track `rightIdx-leftIdx` and return `minLen + 1`. Correct, but the more
  common framing is to store the true length `right-left+1` directly — one fewer "+1 at the end"
  fact to remember under pressure. Minor.

## Where this leaves the ladder

4a's caveat is **resolved** — the slide is a reflex now, and `MinLenSumAtLeast` completes **4b-pre**.
Next is **4b** (longest substring without repeating chars), whose new ingredient is the **seen-set** —
the same "maintain a state, check membership" idea that stalled at 3c. Next day ladders that back in
gently on top of this now-solid expand/shrink base.
