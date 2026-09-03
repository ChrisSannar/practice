# Daily Practice — Progress

This file is the memory of the daily TDD practice system. `/daily` reads it to decide
what to give you next and to space out reviews; it appends to it after each session.
You can edit it by hand any time (bump a confidence score, retire a concept, add a note).

**Active track:** interview patterns — see `INTERVIEW_PATTERNS.md` for the syllabus and laddering.
This file tracks per-concept results/confidence/review dates; that file tracks where we are in the
15-pattern arc.

## Spaced-repetition rule

Each concept has a **Next review** date. When you pass an exercise, its interval grows;
when you fail, it resets so the concept comes back fast.

| Outcome | Next interval |
|---------|---------------|
| pass    | 1 → 3 → 7 → 16 → 35 → 90 days (step up one level each pass) |
| fail / abandoned | reset to 1 day |

`Confidence` is a 1–5 self/observed rating (1 = shaky, 5 = solid). In **recommend** mode,
`/daily` prefers concepts whose `Next review` is on or before today (lowest confidence first).

## Concepts

<!-- One row per concept. /daily maintains this table. -->

| Concept | Lang | Times | Last result | Confidence | Last practiced | Next review |
|---------|------|-------|-------------|------------|----------------|-------------|
| Go basics (vars/slices/strings/maps/loops) | Go | 1 | passed | 4 | 2026-06-25 | 2026-06-28 |
| Go slices & maps (chunk/unique/set/merge) | Go | 1 | failed (Chunk) | 3 | 2026-06-27 | 2026-06-28 |
| Go slice windowing (Take/Drop/Chunk) | Go | 1 | failed (Chunk) | 3 | 2026-06-28 | 2026-06-29 |
| Two pointers — ends mechanic (interview 2a-i/ii) | Go | 1 | passed | 4 | 2026-06-30 | 2026-07-02 |
| Two pointers — converge (interview 2a-iii) | Go | 1 | passed | 4 | 2026-07-01 | 2026-07-03 |
| Two pointers — write index (interview 2b) | Go | 1 | passed | 4 | 2026-07-02 | 2026-07-05 |
| Prefix sum (interview 3a/3b/3c-i/3c-ii) | Go | 1 | abandoned (3c deferred) | 3 | 2026-07-07 | 2026-08-19 |
| Sliding window — fixed 4a + 4b-pre (interview track) | Go | 3 | passed (slide healed) | 4 | 2026-08-15 | 2026-09-04 |
| Sliding window — seen-set window 4b (interview track) | Go | 1 | in-progress | 3 | 2026-08-28 | 2026-08-28 |

## Log

<!-- Reverse-chronological. Newest entry on top. /daily prepends one entry per exercise. -->

### 2026-08-28 — Sliding Window — seen-set window (Go) — interview track 4b
- Folder: `exercises/2026-08-28-go-sliding-window-uniques/`
- Outcome: in-progress
- Notes: `continue` from the now-solid slide/expand-shrink base into 4b (longest substring without
  repeating chars). 4b's new ingredient is the **seen-set** — the exact "maintain a state, check
  membership" reflex deferred at 3c — so the day ladders it back in gently rather than opening cold on
  the full problem. Three tasks: (1) `FirstRepeatedIndex` — pure set membership, no arithmetic, the
  gentle 3c re-entry; (2) `LongestUniqueLen` — the seen-set *driving* an expand/shrink window over
  ints (combines the healed 4b-pre mechanic with task-1 membership); (3) `LengthOfLongestSubstring` —
  the real LC 3 interview problem, same mechanic over a string. Confidence started at 3 because the
  seen-set is the known conceptual wall; will bump to 4 and step the interval if the window-with-set
  clicks, or reset + isolate the membership-inside-a-shrinking-window step if it fights. **Next**: if
  green → 4c minimum-window-substring (need-counts + window-counts). If task 2/3 stall on the shrink
  loop, that shrink-until-duplicate-clears step is the isolated sticking point for the next drill.

### 2026-08-15 — Sliding Window — slide re-drill + variable-window intro (Go) — interview track 4a-refresher / 4b-pre
- Folder: `exercises/2026-08-15-go-sliding-window-slide/`
- Outcome: passed (slide healed) — all four green, and the O(1) slide (`sum += a[i] - a[i-k]`) with
  separated `best`/`bestIdx` state actually appears this time; the large forcing test passes in ms.
- Critique: `exercises/2026-08-15-go-sliding-window-slide/CRITIQUE.md`
- Notes: Today is exactly 4a's `Next review` date (2026-08-15, +3-day pass interval from 2026-08-12),
  and the user explicitly asked to "hone home the missing points from the critique" — i.e. the O(1)
  slide reflex that 4a brute-forced as O(n·k). So this is a deliberate self-reset (critique option 1)
  folded into a `continue`: three tasks that ladder the slide mechanic back in.
  1. `WindowSums` — the slide in pure form (one running `sum`, mutated). Includes a **large forcing
     test** (n=100000, k=50000): O(n) slide finishes in µs; an O(n·k) recompute loop visibly hangs
     for several seconds — that's the only real mechanism to break the `sum(a[i:i+k])` habit, which
     the critique warned would otherwise compound ("skipping it twice is a habit"). Stubs fail fast
     (return nil/0, no loop), so `run.sh` shows RED instantly without hanging.
  2. `MaxSumKAndStart` — 4a's job redone right in one pass: running `sum` + `best` + `bestIdx`, no
     `sum()` helper. Combined value+index signature so it isn't a paste of the 4a brute-force.
  3. `MinLenSumAtLeast` — the variable-window expand/shrink primitive (LC 209, positives only):
     grow right, shrink left while sum ≥ target, track min length. This is the bridge into 4b's
     mechanic — **deliberately sum-only, no seen-set**, because the seen-set (3c) is the user's
     separate known wall; introduce grow/shrink first, defer the seen-set to the real 4b day.
  Confidence held at 3 (will reassess on close — if the large test passes in ms and the slide state
  actually appears, bump to 4 and step the interval; if brute-forced again, drop to 2 and reset to +1).
  **Next**: if green via slide → tick 4a's caveat resolved and attempt 4b (longest substring without
  repeating chars, which adds the seen-set). If brute-forced again → the hole is a habit; isolate the
  slide on an even smaller drill before touching 4b.

### 2026-08-12 — Sliding Window / fixed (Go) — interview track 4a
- Folder: `exercises/2026-08-12-go-sliding-window-fixed/`
- Outcome: passed (all three green) — but see caveat
- Critique: `exercises/2026-08-12-go-sliding-window-fixed/CRITIQUE.md`
- Notes: Test correctness is solid — all three functions pass, including the trap cases (all-negatives,
  k==1, k==len(a), earliest-tie). **Caveat:** the O(1) slide mechanic — the whole point of 4a — did
  not land. Solution uses `sum(a[i:i+k])` inside the loop, recomputing each window from scratch
  (O(n·k)), instead of seed + add-right/subtract-left (O(n)). Running-sum + running-best state never
  appears. Confidence set to 3 (not 4) for this reason; `Next review` scheduled at +3 days per the
  pass-interval rule, but the critique offers the user an explicit self-reset path to re-attempt the
  slide sooner since the pattern is the missing piece, not the answer. User also caught a genuine
  arithmetic bug in my spec's `TestMaxSumKStart` (`want 3` → `want 2`); fixed correctly.
  **Next**: 4b variable window (expand right / shrink left on violation) — a distinct mechanic, gets
  its own day; if implemented with a proper running sum, the fixed-window slide heals as a special case.

### 2026-08-12 — Prefix Sum (Go) — partial close, 3c deferred
- Folder: `exercises/2026-07-07-go-prefix-sum/`
- Outcome: abandoned (3c deferred) — **3a/3b passed**, **3c-i/3c-ii untouched**
- Notes: User hit the conceptual wall on the running-sum + seen-set trick several times across sessions
  and chose to move on rather than keep beating the head, per `INTERVIEW_PATTERNS.md`'s "the plan, not
  a contract" rule. `BuildPrefixSum` and `RangeSum` are correct and idiomatic — confidence 4 on the
  *mechanical* foundation (index shift, `pre[j+1] - pre[i]` identity). The algorithmic leap to
  `HasSubarraySum` / `CountSubarraysWithSum` (LC 560 via running-sum + seen-set / seen-map) did not click
  and is **deferred as a future review rung**, same shape as Two Pointers 2c. Working-tree scaffolds
  (debug `fmt.Println`s, no logic) were reverted before this closure so the next `/daily` grades a clean
  deliberate-deferral, not a "failed attempt." Confidence for the *whole* concept set to 3 (primitives
  solid, application not yet attempted successfully). `Next review` bumped to +7 days (2026-08-19) rather
  than the strict +1-day-fail interval, since this is a deliberate deferral with the primitives still
  solid, not a true fail — gives a real runway before resurfacing. **Next: Pattern 3, Sliding Window**.

### 2026-07-07 — Prefix Sum (Go) — interview track 3a/3b/3c-i/3c-ii
- Folder: `exercises/2026-07-07-go-prefix-sum/`
- Outcome: abandoned (3c deferred) — closed 2026-08-12; see that entry above
- Notes: New pattern, four tasks at user's request (pace change — confident in two pointers, wants
  to move faster). Ladders inside one day since prefix sum is small: `BuildPrefixSum` (the array) →
  `RangeSum` (O(1) query off it) → `HasSubarraySum` (boolean existence via "seen running sums", the
  core trick) → `CountSubarraysWithSum` (leetcode 560, same trick but a frequency map instead of a
  set, and counting instead of stopping at first hit). Closes out pattern 2 in one sitting if it goes
  well; if `CountSubarraysWithSum` fights, that's the isolated sticking point for next time.

### 2026-07-02 — Two pointers / same-direction write index (Go) — interview track 2b
- Folder: `exercises/2026-07-02-go-two-pointers-write-index/`
- Outcome: passed (all four green) — `KeepPositives`, `RemoveElement`, `MoveZeroes`, `RemoveDuplicates`
- Critique: `exercises/2026-07-02-go-two-pointers-write-index/CRITIQUE.md`
- Notes: Gradient fix landed — `KeepPositives` warm-up unstuck it. All four correct and idiomatic;
  notably `RemoveDuplicates` correctly switches to the *other* write-index convention (write = index
  of last kept value, not next open slot) because its keep-rule needs the previous kept value — a
  real sign the pattern clicked, not just pattern-matched. User is confident enough to skip 2c
  (palindrome/two-sum-II problem) and move to the next pattern; 2c parked as a future review rung.

### 2026-07-01 — Two pointers / converge (Go) — interview track 2a-iii
- Folder: `exercises/2026-07-01-go-two-pointers-converge/`
- Outcome: passed (green) — `PairWithTarget`
- Critique: `exercises/2026-07-01-go-two-pointers-converge/CRITIQUE.md`
- Notes: Converge nailed — correct three-way decision, moved only one pointer per step, `for l < r`
  boundary, non-nil empty. Also caught a genuine bug in the test's negatives case (two valid pairs;
  converge returns the outer one). Only cleanup: a leftover `fmt.Println` debug line. Pattern 1
  (Two Pointers converge family) is solid → 2b introduces the same-direction shape.

### 2026-06-30 — Two pointers / ends mechanic (Go) — interview track 2a-i & 2a-ii
- Folder: `exercises/2026-06-30-go-two-pointers-ends/`
- Outcome: passed (both green) — `IsPalindrome` + `ReverseInPlace`
- Critique: `exercises/2026-06-30-go-two-pointers-ends/CRITIQUE.md`
- Notes: Ends-converge mechanic clicked — airtight `for l < r`, idiomatic tuple swap, early-return
  on mismatch, empty/single fall out for free. Only sharpen is style (step belongs in the `for`
  post clause; drop the dead `fmt` import). Ready for the which-pointer decision → 2a-iii.

### 2026-06-28 — Go slice windowing (Go)
- Folder: `exercises/2026-06-28-go-slice-windows/`
- Outcome: failed (Chunk) — Take and Drop green (both were red before), Chunk red
- Critique: `exercises/2026-06-28-go-slice-windows/CRITIQUE.md`
- Notes: Big win — single-bound clamping (`Take`/`Drop`) is now solid. `Chunk` still failed on two
  bugs: loop condition inverted (`idx+size >= len(s)` should be `<=`) and an unconditional tail
  append that adds a spurious empty group on exact multiples. Residual gap is loop-termination
  reasoning, not slice syntax — which the two-pointer converge loop (next task) drills head-on, so
  no third Chunk re-drill needed for now.

### 2026-06-27 — Go slices & maps (Go)
- Folder: `exercises/2026-06-27-go-slices-maps/`
- Outcome: failed (Chunk) — Unique, Intersection, Merge green; Chunk left as `nil`
- Critique: `exercises/2026-06-27-go-slices-maps/CRITIQUE.md`
- Notes: Strong map/perf instincts (O(n) Intersection, clean Unique). Real gap is slice
  index/window arithmetic (`s[i:end]`, stepping by `size`, clamping the tail). Also: Merge
  isn't O(n²) (it's linear); Intersection passes by luck (sorts instead of honoring a-order).
  Next task isolates windowing → see 2026-06-28.

### 2026-06-25 — Go basics (Go)
- Folder: `exercises/2026-06-25-go-basics/`
- Outcome: passed (all six green, including Unicode `Reverse` and empty-map `WordCount`)
- Critique: `exercises/2026-06-25-go-basics/CRITIQUE.md`
- Notes: Back in the flow cleanly. Gaps are idiom, not fundamentals — map zero-value
  increment (skip the existence check), `strconv.Itoa` over `fmt.Sprintf`, in-place
  two-index reverse, `strings.ContainsRune`. Next task (`continue`): deeper slices/maps,
  less scaffolding.

<!--
### YYYY-MM-DD — <concept> (<lang>)
- Folder: `exercises/YYYY-MM-DD-<lang>-<slug>/`
- Outcome: passed | failed | abandoned | in-progress
- Notes: <what went well / where you struggled / how the next task should adjust>
-->
