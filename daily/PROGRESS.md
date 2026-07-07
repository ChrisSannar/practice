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
| Prefix sum (interview 3a/3b/3c-i/3c-ii) | Go | 0 | in-progress | — | 2026-07-07 | 2026-07-07 |

## Log

<!-- Reverse-chronological. Newest entry on top. /daily prepends one entry per exercise. -->

### 2026-07-07 — Prefix Sum (Go) — interview track 3a/3b/3c-i/3c-ii
- Folder: `exercises/2026-07-07-go-prefix-sum/`
- Outcome: in-progress
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
