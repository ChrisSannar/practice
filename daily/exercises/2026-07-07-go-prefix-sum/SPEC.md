# Prefix Sum — pattern 2 (interview track 3a/3b/3c-i/3c-ii)

**Pattern:** Prefix Sum. Precompute cumulative sums once so range-sum queries become O(1) instead of
O(n) each, and so "does some subarray sum to X" becomes a single pass with a hash set/map instead of
checking every subarray (O(n²)).

Pace note: you called Two Pointers solid, so this is **four tasks in one day** — all four are small
and build directly on each other; take the time you need.

## Tasks (do them in order — each one is used by the next)

In `prefixsum.go`:

1. `BuildPrefixSum(a []int) []int` — return `pre` where `pre[0] = 0` and
   `pre[i] = pre[i-1] + a[i-1]` for `i = 1..len(a)`. So `pre` has length `len(a)+1`, and `pre[i]` is
   the sum of the first `i` elements of `a`.
2. `RangeSum(pre []int, i, j int) int` — given a prefix slice built by task 1, return the sum of
   `a[i..j]` **inclusive**, in O(1). (Work out the formula from what `pre[i]` and `pre[j+1]` each
   represent — no loop needed.)
3. `HasSubarraySum(a []int, target int) bool` — does **any contiguous subarray** of `a` sum to
   `target`? One pass: track a running sum and a set of running sums you've already seen. At each
   step, ask "have I seen `running - target` before?" — if yes, the subarray between that point and
   here sums to `target`.
4. `CountSubarraysWithSum(a []int, target int) int` — same trick as task 3, but **count** every
   subarray that sums to `target` instead of stopping at the first one. The set becomes a
   `map[int]int` of how many times each running sum has occurred, and instead of a boolean check you
   add `counts[running-target]` to the result at every step (there may be several earlier points
   with the same running sum, each giving a different valid subarray).

## Acceptance criteria

`./daily/run.sh` is green. Tests cover: `BuildPrefixSum`/`RangeSum` on ordinary, negative, single, and
empty inputs; `HasSubarraySum`/`CountSubarraysWithSum` on positives, negatives, no-match, and empty.

## The thing to get right

Tasks 1–2 are pure bookkeeping — get the prefix-array index shift right (`pre[i]` covers the first
`i` elements, so `pre[j+1] - pre[i]` is the sum of indices `i..j`). Tasks 3–4 are the actual pattern:
the running-sum-as-you-go idea, checked against a seen-set/seen-map, is what unlocks a huge class of
subarray-sum problems in O(n) instead of O(n²). Task 4 is the only genuinely new idea over task 3 —
counting instead of a yes/no, which means the seen-set becomes a seen-map of *counts*, and you must
initialize it with `{0: 1}` (an empty prefix, i.e. "sum so far is 0," has occurred once before you've
read anything) so a subarray starting at index 0 is counted correctly.

## Larger arc

Pattern **2 (Prefix Sum)** in `INTERVIEW_PATTERNS.md`, steps 3a/3b/3c-i/3c-ii. Task 4 is
`subarray-sum-equals-k` (LeetCode 560) — closes out this pattern. Next up: **Sliding Window**.
No solution or hints here — that's yours to write.
