# Sliding Window (fixed) — pattern 3, step 4a

**Pattern:** Fixed-size sliding window. Seed the sum of the first `k` elements, then slide the
window one step right at a time: add the new right element, subtract the old left element. Each
window's sum is O(1) — no recompute. The decision layered on top varies (track max, track min,
track *where* the max is).

This is the **first** micro-step of Sliding Window. The variable-window shape (expand right /
shrink left on violation) is a different mechanic and is parked for the next session — today is
just the fixed-window slide.

## Tasks (do them in order — each one builds on the previous)

In `slidingwindow.go`:

1. `MaxSumK(a []int, k int) int` — return the **maximum** sum of any contiguous subarray of
   length `k`. If `k <= 0` or `k > len(a)`, return `0`.
2. `MinSumK(a []int, k int) int` — return the **minimum** sum of any contiguous subarray of
   length `k`. Same seed-and-slide mechanic, opposite comparison. Same invalid-`k` rule: return `0`.
3. `MaxSumKStart(a []int, k int) int` — return the **starting index** of the window whose sum is
   maximum (i.e. the answer to task 1, but reporting *where* it lives, not its value). If multiple
   windows share the max, return the **earliest** index. If `k <= 0` or `k > len(a)`, return `-1`.

## Acceptance criteria

`./daily/run.sh` is green. Tests cover: ordinary mixed positives/negatives, an all-negatives array
(so the true max is negative — the seed matters), single-element windows (`k == 1`), window equal
to the whole array (`k == len(a)`), ties (earliest wins), and the invalid-`k` sentinels.

## The thing to get right

Seed the first window's sum with a plain loop. Then for each slide position `i` from `k` to
`len(a)-1`: `sum += a[i] - a[i-k]`. Decide the best after each slide (or fold the seed in as the
initial best — your call, both work). The only genuinely new reflex over what you've already done
(range sum off a prefix slice) is that you don't *store* the sums — you keep a single running
`sum` and a single running `best`, O(1) space beyond the input. Task 3 adds one more piece of
state to carry alongside `best`.

## Larger arc

Pattern **3 (Sliding Window)** in `INTERVIEW_PATTERNS.md`, step **4a**. Next up: **4b** variable
window (expand right, shrink left on violation — different mechanic, drills a different reflex).
No solution or hints here — that's yours to write.