---
description: Generate today's bite-sized TDD practice exercise (Go/Python/TS) and track progress
argument-hint: "[concept | continue | (empty to recommend)]"
allowed-tools: Bash(./daily/run.sh:*), Bash(date:*), Bash(find:*), Bash(git diff:*), Bash(git status:*), Read, Write, Edit
---

You are running the **daily practice ritual** for this repo. The user owns the learning —
your job is ONLY to manufacture a well-scoped TDD workout and to calibrate its difficulty from how
past attempts went. Do not teach the concept and do not write the solution.

**Volume: at least 3 exercises per day.** Each day's folder must contain **≥3 small problems**
(distinct functions/tasks), sized so the whole set is completable in ~20 minutes (each one ~5–7 min).
They should build on each other — ideally consecutive micro-steps of the active track's ladder, or
a primitive plus two variations that drill the same idea.

**Active track: interview patterns.** The current learning goal is `daily/INTERVIEW_PATTERNS.md` —
a Go syllabus of 15 coding-interview patterns, each broken into primitive micro-steps that ladder
up to a real problem. Read it at the start of every run. Treat it as the spine for `continue` and
empty/recommend mode, and **always build the prerequisite primitives before the full problem**.
When you finish a step, tick its `[ ]`→`[x]` checkbox in that file (and flip the pattern's status
☐→◐→☑) so the syllabus reflects progress. A named concept argument still overrides the track.

Argument: `$ARGUMENTS`

Date: by default run `date +%F` and use that as `<date>` everywhere below. **If the argument
contains an explicit date** (a `YYYY-MM-DD` token), use that as `<date>` instead — this lets you
generate or grade an exercise for a day other than today. Strip the date token out before
interpreting the rest of the argument as the concept / `continue` / empty.

## Step 1 — Close out the previous exercise (always do this first)

1. Find the most recent exercise: `find daily/exercises -mindepth 1 -maxdepth 1 -type d | sort | tail -1`.
2. If one exists and `PROGRESS.md` shows it as `in-progress`:
   - Run its tests: `./daily/run.sh <its-date-prefix>`.
   - Look at the user's attempt: `git diff -- daily/exercises/<that-folder>/` and read the changed files.
   - Decide the outcome: **passed** (tests green), **failed** (red but attempted), or **abandoned** (untouched stub).
   - **Write a critique** to `CRITIQUE.md` in that exercise's folder (skip only if the stub was untouched).
     Read the user's actual solution and assess it: what they got right / already idiomatic, where to
     sharpen (idiom, stdlib choices, performance, edge cases), and a 1–5 confidence verdict for the
     concept. Be specific and show the better version inline. Critique the code, not the person; these
     are sharpening notes, not bug reports unless tests actually fail.
   - Update `PROGRESS.md`:
     - Set that log entry's Outcome and add a one-line note on how it went and how the next task should adjust.
     - Update the concept's row in the Concepts table: bump `Times`, set `Last result`, adjust `Confidence` (1–5),
       set `Last practiced` = its date, and compute `Next review` using the spaced-repetition rule
       (pass → step the interval up 1→3→7→16→35→90 days; fail/abandoned → reset to +1 day).
3. If there's no prior in-progress exercise, skip straight to Step 2.

## Step 2 — Decide what to give them today

Parse `$ARGUMENTS`:
- **A concept** (e.g. `go channels`, `python iterators`, `ts mapped types`): use it.
  Infer the language from the concept. If the concept doesn't imply one of Go/Python/TypeScript,
  ask the user which of the three to use.
- **`continue`**: advance the active track — give the next unchecked micro-step in
  `INTERVIEW_PATTERNS.md` (the next `[ ]` within the current pattern, or the first step of the next
  pattern). Honor the laddering: never jump to a `?c` problem step if its `?a`/`?b` primitives
  aren't done.
- **Empty**: **recommend** — first, from the `PROGRESS.md` Concepts table, resurface any concept
  whose `Next review` is on or before today (lowest `Confidence` first). If none are due, advance
  the active track per `INTERVIEW_PATTERNS.md` (same as `continue`).

Only Go, Python, and TypeScript are in scope.

## Step 3 — Calibrate difficulty

Use the concept's history in `PROGRESS.md`:
- Passed easily last time → push harder: a new facet, an edge case, or fewer scaffolds.
- Failed or struggled → shrink it: isolate the exact sticking point in a simpler task.
- New concept → start with the smallest meaningful slice.
Keep the whole set of **≥3 exercises** completable in **~20 minutes** total. If the concept is
bigger, scope today to one slice of the ladder (its next 3 rungs) and note the larger arc in the
spec so `continue` can pick up the rest.

## Step 4 — Generate today's exercise folder

Create `daily/exercises/<date>-<lang>-<slug>/` with:

- **`SPEC.md`** — a short problem statement: the goal, the **≥3 tasks** for today (each a distinct
  function), concrete acceptance criteria, and a line naming the larger arc this ladders into. No
  hints, no solution outline.
- **A stub file** with a `// TODO` / `# TODO` markered signature for **each of the ≥3 tasks**. Above
  every stubbed function, include a **concrete input → output example** in the doc comment (e.g.
  `// BuildPrefixSum([1, 2, 3]) -> [0, 1, 3, 6]`) so the contract is unambiguous without reading the
  tests. This is a standard part of every stub, not optional.
- **Failing tests** that encode the acceptance criteria for every task. All tests MUST fail against
  the stubs.

Per-language conventions (match existing repo style):

- **Go** — add a `go.mod` (`module daily/<date>-<lang>-<slug>`, `go 1.25.6`), a `<slug>.go`
  stub (package matching the test), and `<slug>_test.go` using stdlib `testing` (table-driven).
- **Python** — a `<slug>.py` stub and a `test_<slug>.py` using stdlib `unittest` (no third-party deps).
- **TypeScript** — a `<slug>.ts` stub and a `<slug>.test.ts` using `bun test`'s
  `import { test, expect } from "bun:test"`.

## Step 5 — Record and hand off

1. Prepend a new entry to the `PROGRESS.md` **Log** (Outcome: `in-progress`) and add/update the
   concept's row in the Concepts table (Last practiced = today, Next review = today, set a starting Confidence).
2. Tell the user, briefly: the concept, the folder path, and to run `./daily/run.sh` to see RED.
   Do not reveal how to implement it.

## Hard rules

- **At least 3 exercises per day** (distinct tasks in the folder).
- **Every stub function has an input → output example** in its doc comment.
- Never write or sketch the solution. Folders contain only the spec + failing tests.
- Tests must fail initially (true RED).
- ~20 minutes for the whole set; split bigger concepts across days.
- Only Go, Python, TypeScript.
